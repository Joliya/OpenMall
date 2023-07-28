/**
 * @Author: jinpeng zhang
 * @Date: 2023/7/25 22:04
 * @Description:
 */

package string_util

func IsEmpty(str string) bool {
	return str == ""
}

func IsNil(str interface{}) bool {
	return str == nil
}

func IsNotNil(str interface{}) bool {
	return !IsNil(str)
}
