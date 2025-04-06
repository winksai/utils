package main

import "regexp"

// 验证中国大陆手机号格式
func ValidateChinesePhoneNumberStrict(phone string) bool {
	// 各运营商号段（截至2023年）
	patterns := []string{
		`^1[3-9]\d{9}$`, // 基础验证
		`^1(3[0-9]|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\d{8}$`, // 详细号段
		`^14[579]\d{8}$`,   // 物联网号段
		`^16[2567]\d{8}$`,  // 虚拟运营商
		`^19[12389]\d{8}$`, // 新号段
	}

	for _, pattern := range patterns {
		matched, _ := regexp.MatchString(pattern, phone)
		if matched {
			return true
		}
	}
	return false
}
