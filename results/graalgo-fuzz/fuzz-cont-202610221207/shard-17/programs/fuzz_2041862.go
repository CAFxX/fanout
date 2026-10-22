package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int64, p1 uint32) int {
    v5 := GenS1{F0: p1, F1: 621}
    v6 := GenS1{F0: p1, F1: 386}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint = 65535
        v9 := GenS1{F0: (^(p1)), F1: 18446744073709551615}
        var v10 int = -78
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := []int{128}
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    _ = v6
    return int(p1)
}
func genFn1() int16 {
    v12 := GenS1{F0: 4294967295, F1: 42}
    if true {
        var v13 bool = (-1 != -2147483648)
        var v14 int64 = 9223372036854775807
        v14 = v14
        _ = v13
        _ = v14
    } else {
        var v15 uint64 = 726
        var v16 bool = (504 <= 1000)
        _ = v15
        _ = v16
    }
    _ = v12
    return -69
}

func main() {
    var v1 uint64 = 2
    var v2 uint8 = 54
    var v3 int = 127
    var v4 int8 = -39
    for fi17 := 0; fi17 < 20; fi17++ {
        v1 += v1
        fmt.Println(v3, fi17)
        _ = fi17
    }
    var v18 float32 = 0.0
    if (18 != 1) {
        v19 := []int{(-(v3))}
        fmt.Println(v4)
        _ = v19
    }
    v20 := GenS1{F0: 1, F1: 189}
    fmt.Println(v3, v1)
    if true {
        var v21 int32 = -2147483648
        var v22 string = "foo bar"
        _ = v21
        _ = v22
    } else {
        var v23 uint16 = 1
        _ = v23
    }
    v24 := []int{genFn0(-9223372036854775808, 2147483647), (v3 / (((v3 - genFn0(-423, 0))) | 1))}
    v1 = (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
