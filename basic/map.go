package main

import (
	"fmt"
	"reflect"
)

func main() {
	// map, 해시 테이블을구현한 자료구조, 파이썬의 딕셔너리, 루비 해시와 같다.
	var idMap map[int]string // nil map
	fmt.Println(idMap) // map[]

	// make 함수를 이용한 map 생성
	madeMap := make(map[int]string)
	fmt.Println(madeMap) // map[]

	// 기본값으로 초기화 된 map 생성
	company := map[string]string {
		"G": "GOOGLE",
		"F": "FACEBOOK",
		"M": "MICROSOFT",
	}
	fmt.Println(company) // map[F:FACEBOOK G:GOOGLE M:MICROSOFT]

	// map을 초기화 해 주지 않으면 "assignment to entry in nil map" 에러가 발생한다.
	idMap = map[int]string{}

	// make() 함수로 생성한 map은 초기화가 되어 있다.
	madeMap[0] = "test_package"

	// map, 추가 혹은 갱신
	idMap[404] = "Not Found"
	idMap[500] = "Client Error"
	fmt.Println(idMap) // map[404:Not Found 500:Client Error]
	idMap[500] = "Interval Server Error"
	fmt.Println(idMap) // map[404:Not Found 500:Interval Server Error]

	// map 사용
	fmt.Println(idMap[404]) // Not Found

	// 존재하지 않는 key에 대해서 value의 경우 zero value를 반환,
	fmt.Println(reflect.TypeOf(idMap[403])) // ""
	// reference의 경우 nil을 반환한다.
	referenceMap := map[int]*int{}
	fmt.Println(referenceMap[1]) // <nil>
	// key에 대한 value가 존재 여부를 두 번째 리턴값에 반환한다.
	status, exist := idMap[403]
	fmt.Println(status, exist) // "" false

	// map range
	for key, value := range idMap {
		fmt.Println(key, value)
	} // 404 Not Found 500 Interval Server Error
}
