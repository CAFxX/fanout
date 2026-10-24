package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint }
func genFn0() uint16 {
    var v5 int = -590
    v6 := []int{v5, 9223372036854775807, v5}
    _ = v5
    _ = v6
    return 65535
}
func genFn1() string {
    v7 := []int{}
    if (false || false) {
        if (407 != -2147483648) {
            var v8 int64 = 9223372036854775807
            v8 += -9223372036854775808
            _ = v8
        } else {
            v9 := GenS1{F0: 1, F1: -128, F2: 65535}
            _ = v9
        }
    }
    _ = v7
    return ""
}

func main() {
    var v1 int8 = 100
    var v2 int16 = 738
    var v3 uint8 = 178
    var v4 float32 = 3.14159
    if true {
        v2 = v2
        if (!(4294967295 != 2147483648)) {
            if false {
                v1 -= (^(v1))
                var v10 uint = 261
                _ = v10
            } else {
                var v11 uint = 2
                var v12 int8 = v1
                _ = v11
                _ = v12
            }
        } else {
            v3 = 1
        }
        fmt.Println(v3, v3, v3)
    } else {
        fmt.Println(v3, v2)
    }
    v13 := []int{669, 4294967295}
    var v14 int32 = -2147483648
    var v15 float32 = -2.25
    v16 := GenS1{F0: 255, F1: int8(v2), F2: 2}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v17 bool = true
    var v18 uint8 = 42
    v19 := GenS1{F0: 65535, F1: (^(v1)), F2: 255}
    fmt.Println(v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
