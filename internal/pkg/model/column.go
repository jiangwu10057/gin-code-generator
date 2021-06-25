package model

import (
)

type ColumnGetterOperator struct {
	Strategy ColumnGetter
}

func NewColumnOperator() ColumnGetterOperator {
	return ColumnGetterOperator{}
}

func (operator ColumnGetterOperator) SetStrategy(strategy ColumnGetter) (ColumnGetterOperator){
	operator.Strategy = strategy
	return operator
}

func (operator ColumnGetterOperator) Get(tableName string) (error) {
	err := operator.Strategy.Get(tableName)
	// data := operator.Strategy.Result

	return err
}