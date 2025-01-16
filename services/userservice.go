package services

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"net"
	"strconv"
	"strings"
)

func GetUserRole(u models.UserModel) string {
	i := u.Roles[0]
	return i.Slug
}

// getInterfaceByName 获取指定名称的网卡
func GetInterfaceByName(name string) *net.Interface {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, iface := range interfaces {
		if iface.Name == name {
			return &iface
		}
	}
	return nil
	//, errors.New(fmt.Sprintf("未找到名为 %s 的网卡", name))
}

// getIPv4Addresses 获取网卡的所有IPv4地址
func GetIPv4Addresses(iface *net.Interface) string {
	addrs, err := iface.Addrs()
	if err != nil {
		return "nil"
	}
	var ips []string
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}
	if len(ips) == 0 {
		return ""
		//, errors.New(fmt.Sprintf("网卡 %s 没有IPv4地址", iface.Name))
	}
	ip := strings.Join(ips, ".")
	return ip
}

func GetStudentNameMap(u models.UserModel) (results []map[string]interface{}) {
	results, err := GetDb().Query("SELECT std_id FROM user_student as us\njoin goadmin_users as gu on us.user_id = gu.id where gu.id = " + strconv.Itoa(int(u.Id)) + ";")
	if err != nil {
		return
	}
	return
}

func GetStudentName(u models.UserModel) interface{} {
	var stu []interface{}
	// 遍历切片
	for _, item := range GetStudentNameMap(u) {
		// 获取每个 map 的键和值
		for _, value := range item {
			stu = append(stu, value)
		}
	}
	return stu[0]
}

func TransItoStr(a interface{}) string {
	// 类型断言，确保 value 是 int64 类型
	if intValue, ok := a.(int64); ok {
		// 使用 strconv.FormatInt 将 int64 转换为 string
		strValue := strconv.FormatInt(intValue, 10) // 10 表示十进制
		fmt.Println(strValue)                       // 输出
		return strValue
	} else {
		fmt.Println("value is not an int64")
		return "error!"
	}
}
