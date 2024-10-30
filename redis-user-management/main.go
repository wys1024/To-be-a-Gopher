package main

import (
	"context"       // 引入 context 包，用于处理上下文
	"encoding/json" // 引入 json 包，用于 JSON 数据的编码和解码
	"fmt"           // 引入 fmt 包，用于格式化输入输出
	"log"           // 引入 log 包，用于日志记录
	"net/http"      // 引入 net/http 包，用于 HTTP 客户端和服务器

	"github.com/go-redis/redis/v8" // 引入 go-redis 包，用于与 Redis 数据库交互
)

// 声明上下文变量，用于 Redis 操作
var ctx = context.Background()

// 声明 Redis 客户端变量
var rdb *redis.Client

// User 结构体表示系统中的用户
type User struct {
    ID    string `json:"id"`    // 用户 ID
    Name  string `json:"name"`  // 用户姓名
    Email string `json:"email"` // 用户邮箱
}

// 初始化函数，用于创建 Redis 客户端
func init() {
    // 创建一个新的 Redis 客户端
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis 服务器地址
    })

    // 检查与 Redis 的连接是否成功
    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        // 如果连接失败，记录错误并退出
        log.Fatalf("Could not connect to Redis: %v", err)
    }
}

// CreateUser 创建一个新的用户
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User // 声明一个 User 结构体变量

    // 解码请求体中的 JSON 数据到 user 结构体
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        // 如果解码失败，返回 400 错误
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // 将用户信息存储到 Redis 中
    rdb.Set(ctx, user.ID, user, 0) // ID 作为键，user 作为值

    // 返回 201 创建成功的状态码，并将用户信息以 JSON 格式返回
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// GetUser 根据 ID 获取用户信息
func GetUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id") // 从请求的查询参数中获取用户 ID

    // 从 Redis 中获取用户信息
    val, err := rdb.Get(ctx, id).Result()
    if err == redis.Nil {
        // 如果用户不存在，返回 404 错误
        http.Error(w, "User not found", http.StatusNotFound)
        return
    } else if err != nil {
        // 如果获取过程中发生错误，返回 500 错误
        http.Error(w, "Error retrieving user", http.StatusInternalServerError)
        return
    }

    var user User // 声明一个 User 结构体变量
    // 将获取的 JSON 字符串解码到 user 结构体
    json.Unmarshal([]byte(val), &user)

    // 返回 200 状态码，并将用户信息以 JSON 格式返回
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

// UpdateUser 更新已有用户的信息
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    var user User // 声明一个 User 结构体变量

    // 解码请求体中的 JSON 数据到 user 结构体
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        // 如果解码失败，返回 400 错误
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // 更新 Redis 中的用户信息
    rdb.Set(ctx, user.ID, user, 0)

    // 返回 200 状态码，并将更新后的用户信息以 JSON 格式返回
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

// DeleteUser 根据 ID 删除用户
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id") // 从请求的查询参数中获取用户 ID

    // 从 Redis 中删除用户
    err := rdb.Del(ctx, id).Err()
    if err == redis.Nil {
        // 如果用户不存在，返回 404 错误
        http.Error(w, "User not found", http.StatusNotFound)
        return
    } else if err != nil {
        // 如果删除过程中发生错误，返回 500 错误
        http.Error(w, "Error deleting user", http.StatusInternalServerError)
        return
    }

    // 返回 204 状态码，表示成功删除用户且无内容返回
    w.WriteHeader(http.StatusNoContent)
}

func main() {
    // 注册处理函数
    http.HandleFunc("/users", CreateUser)               // 处理用户创建请求
    http.HandleFunc("/users/get", GetUser)              // 处理获取用户请求
    http.HandleFunc("/users/update", UpdateUser)        // 处理用户更新请求
    http.HandleFunc("/users/delete", DeleteUser)        // 处理用户删除请求

    // 启动 HTTP 服务器，监听 8080 端口
    fmt.Println("Server is running at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil)) // 启动服务器并记录致命错误
}
