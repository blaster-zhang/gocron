package utils

import (
    "crypto/md5"
    "encoding/hex"
    "math/rand"
    "time"
    "runtime"
    "github.com/Tang-RoseChild/mahonia"
    "strings"
    "os"
    "fmt"
)

// 生成长度为length的随机字符串
func RandString(length int64) string {
    sources := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    result := []byte{}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    sourceLength := len(sources)
    var i int64 = 0
    for ; i < length; i++ {
        result = append(result, sources[r.Intn(sourceLength)])
    }

    return string(result)
}

// 生成32位MD5摘要
func Md5(str string) string {
    m := md5.New()
    m.Write([]byte(str))

    return hex.EncodeToString(m.Sum(nil))
}

// 生成0-max之间随机数
func RandNumber(max int) int {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    return r.Intn(max)
}

// 判断当前系统是否是windows
func IsWindows() bool {
    return runtime.GOOS == "windows"
}

// GBK编码转换为UTF8
func GBK2UTF8(s string) (string, bool) {
    dec := mahonia.NewDecoder("gbk")

    return dec.ConvertStringOK(s)
}

// 批量替换字符串
func ReplaceStrings(s string, old []string, replace []string) string  {
    if s == "" {
        return s
    }
    if len(old) != len(replace) {
        return s
    }

    for i, v := range old {
        s = strings.Replace(s, v, replace[i], 1000)
    }

    return s
}

func InStringSlice(slice []string, element string) bool {
    element = strings.TrimSpace(element)
    for _, v := range slice {
        if strings.TrimSpace(v) == element{
            return true
        }
    }

    return false
}

// 转义json特殊字符
func EscapeJson(s string) string  {
    specialChars := []string{"\\", "\b","\f", "\n", "\r", "\t", "\"",}
    replaceChars := []string{ "\\\\", "\\b", "\\f", "\\n", "\\r", "\\t", "\\\"",}

    return ReplaceStrings(s, specialChars, replaceChars)
}

// 判断文件是否存在及是否有权限访问
func FileExist(file string) bool {
    _, err := os.Stat(file)
    if os.IsNotExist(err) {
        return false
    }
    if os.IsPermission(err) {
        return false
    }

    return true
}

// 格式化环境变量
func FormatUnixEnv(key, value string) string {
    return fmt.Sprintf("export %s=%s; ", key, value)
}