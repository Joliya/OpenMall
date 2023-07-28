/**
 * @Author: jinpeng zhang
 * @Date: 2023/7/28 15:16
 * @Description:
 */

package number

import (
	"github.com/shopspring/decimal"
)

func GetYuanFromFen(price int) string {
	d := decimal.New(1, 2)
	result := decimal.NewFromInt(int64(price)).DivRound(d, 2).String()
	return result
}

func GetFenFromYuan(price float64) int64 {
	d := decimal.New(1, 2)  //分转元乘以100
	d1 := decimal.New(1, 0) //乘完之后，保留2为小数，需要这么一个中间参数
	//如下是满足，当乘以100后，仍然有小数位，取四舍五入法后，再取整数部分
	dff := decimal.NewFromFloat(price).Mul(d).DivRound(d1, 0).IntPart()
	return dff
}
