package utils

// PasswordManager 密码管理器
type PasswordManager struct {
	encryptionKey []byte
}

// NewPasswordManager 创建新的密码管理器
func NewPasswordManager(key []byte) (*PasswordManager, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}
	return &PasswordManager{encryptionKey: key}, nil
}

// EncryptPassword 加密密码
func (pm *PasswordManager) EncryptPassword(password string) (string, error) {
	return EncryptString(password, pm.encryptionKey)
}

// DecryptPassword 解密密码
func (pm *PasswordManager) DecryptPassword(encryptedPassword string) (string, error) {
	return DecryptString(encryptedPassword, pm.encryptionKey)
}

// HashPassword 生成密码哈希
func (pm *PasswordManager) HashPassword(password string) (string, error) {
	return HashPassword(password)
}

// VerifyPassword 验证密码与哈希是否匹配
func (pm *PasswordManager) VerifyPassword(password, hash string) bool {
	return CheckPasswordHash(password, hash)
}
