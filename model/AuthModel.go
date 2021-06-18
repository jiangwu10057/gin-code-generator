/**
* @Author chenwu
* @Date  2021/06/18 09:56
**/
package model

type AuthModel struct {

}

func (AuthModel) TableName() string {
	return "AUTH"
}