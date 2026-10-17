package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int16 }
func genFn0() bool {
    var v4 int32 = -2147483648
    if (-2.25 != 3.14159) {
        v5 := []int{-934, 772, 32768, -9223372036854775808}
        _ = v5
    } else {
        var v6 uint = 2
        v7 := []int{32768, -1, 127}
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: true, F1: 533, F2: 32767}
    _ = v4
    _ = v8
    return (1 > 2)
}
func genFn1(p0 uint) int16 {
    p0 = 2
    var v9 uint16 = 42
    _ = p0
    _ = v9
    return -627
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ""
    var v3 int32 = 2147483647
    var v10 uint64 = 0
    fmt.Println(v10, v2)
    var v11 float32 = 1.5
    v12 := GenS1{F0: ((2 <= 0) && true), F1: 1, F2: 0}
    v13 := []int{-543, 9223372036854775807, -589}
    var v14 int = 42
    v15 := GenS1{F0: false, F1: 0, F2: genFn1(778)}
    v10 = 495
    var v16 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
