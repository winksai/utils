package main

import (
	"fmt"
	"log"
	"utils/utils"
)

func main() {

	// 初始化密码管理器
	// 注意：密钥应该从安全的地方获取，而不是硬编码
	key := []byte("A_32_byte_key1234567890121234567")
	pm, err := utils.NewPasswordManager(key)
	if err != nil {
		log.Fatalf("Failed to create password manager: %v", err)
	}

	// 示例1: 加密解密
	password := "my-secret-password-123"
	fmt.Printf("Original password: %s\n", password)

	// 加密
	encrypted, err := pm.EncryptPassword(password)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}
	fmt.Printf("Encrypted password: %s\n", encrypted)

	// 解密
	decrypted, err := pm.DecryptPassword(encrypted)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}
	fmt.Printf("Decrypted password: %s\n", decrypted)

	// 示例2: 密码哈希
	fmt.Println("\nPassword hashing example:")
	userPassword := "user-password-456"

	// 生成哈希
	hash, err := pm.HashPassword(userPassword)
	if err != nil {
		log.Fatalf("Hashing failed: %v", err)
	}
	fmt.Printf("Password hash: %s\n", hash)

	// 验证密码
	fmt.Println("Verifying correct password:", pm.VerifyPassword(userPassword, hash))
	fmt.Println("Verifying wrong password:", pm.VerifyPassword("wrong-password", hash))

}
