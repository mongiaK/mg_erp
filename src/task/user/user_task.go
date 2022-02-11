/*================================================================
*
*  文件名称：user_task.go
*  创 建 者: mongia
*  创建日期：2022年01月07日
*
================================================================*/

package task

import (
	"errors"
	"strings"

	userpb "pharmacyerp/pb/user"
)

var (
	logicMap map[int]string = map[int]string{
		0: " < ",
		1: " <= ",
		2: " = ",
		3: " >= ",
		4: " > ",
		5: " in ",
		6: " not in ",
		7: " group by ",
		8: " order by ",
		9: " <> ",
	}
)

func getUserFileds(field int) []string {
	if 0 == field {
		return nil
	}

	var ret []string = []string{}
	for key, val := range userpb.UserItem_name {
		val := strings.ToLower(val)
		if 0 != int(key)&field {
			ret = append(ret, val)
		}
	}

	return ret
}

func generateSQLUserField(item *userpb.ConditionItem) (string, bool) {
	keyField, err := userpb.UserItem_name[int32(item.Key)]
	if !err {
		return "", err
	}
	keyField = strings.ToLower(keyField)

	switch item.Logic {
	case userpb.Logic_IN:
	case userpb.Logic_EQUAL:
		break
	case userpb.Logic_LESS:
	case userpb.Logic_LESS_THAN:
	case userpb.Logic_GREAT_THAN:
	case userpb.Logic_GREAT:
		logic, _ := logicMap[int(item.Logic)]
		keyField += logic
		break
	}
	return keyField, true
}

func generateSQLUserValue(item *userpb.ConditionItem) interface{} {
	switch item.Vtype {
	case userpb.ConditionValueType_INT:
		return item.Ivalue
	case userpb.ConditionValueType_STRING:
		return item.Svalue
	case userpb.ConditionValueType_INT_ARRAY:
		return item.Iavalue
	case userpb.ConditionValueType_STRING_ARRAY:
		return item.Savalue
	}
	return nil
}

func generateSQLWhere(cond *userpb.Condition) (map[string]interface{}, error) {
	where := map[string]interface{}{}
	for _, item := range cond.And.Items {
		keyField, err := generateSQLUserField(item)
		keyValue := generateSQLUserValue(item)
		if !err || nil == keyValue {
			return nil, errors.New("[sql] generate and condition failed")
		}

		where[keyField] = keyValue
	}

	orCond := []map[string]interface{}{}
	for _, innerAnd := range cond.Or {
		andInOrCond := map[string]interface{}{}
		for _, item := range innerAnd.Items {
			keyField, err := generateSQLUserField(item)
			keyValue := generateSQLUserValue(item)
			if !err || nil == keyValue {
				return nil, errors.New("[sql] generate or condition failed")
			}
			andInOrCond[keyField] = keyValue
		}

		orCond = append(orCond, andInOrCond)
	}
	if len(orCond) > 0 {
		where["_or"] = orCond
	}

	if nil != cond.Groupby {
		where["_groupby"] = cond.Groupby.Svalue
	}
	if nil != cond.Orderby {
		where["_orderby"] = cond.Orderby.Svalue
	}

	return where, nil
}
