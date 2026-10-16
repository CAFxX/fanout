package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int16; F2 bool }
func genFn0(p0 uint32, p1 int16, p2 uint8) uint8 {
    p1 = 893
    fmt.Println(p2)
    v7 := GenS1{F0: -562, F1: -377, F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 47
}
func genFn1() int32 {
    v8 := []int{1, -9223372036854775808, -1, 4294967295}
    if (!false) {
        if (1 == 2) {
            var v9 uint16 = 65535
            var v10 int8 = -128
            _ = v10
            _ = v9
        }
        v11 := GenS2{F0: ""}
        var v12 int32 = -272
        _ = v11
        _ = v12
    }
    var v13 uint8 = 163
    _ = v13
    _ = v8
    return 1
}

func main() {
    var v1 uint32 = 65535
    var v2 bool = (42 != 142)
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    var v5 string = "a"
    var v6 float32 = (3.14159 * (0.1 + 0.0))
    var v14 int16 = -551
    var v15 int8 = 42
    var v16 uint = 2
    if true {
        switch -360 {
        case 127:
            var v17 int = -9223372036854775808
            _ = v17
        default:
            v18 := GenS1{F0: 42, F1: -125, F2: (v16 * v16)}
            _ = v18
        }
    } else {
        var v19 uint8 = 42
        v16 = v16
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v5 = v5
        var v21 int32 = int32(v16)
        _ = fi20
        _ = v21
    }
    var v22 string = v5
    var v23 uint = 600
    if false {
        fmt.Println(v6)
    } else {
        if v2 {
            v24 := GenS3{F0: 0, F1: -150, F2: (false || (-2147483648 >= 4294967295))}
            v3[128] = 32768
            v25 := GenS3{F0: 9223372036854775807, F1: (v14 | 32767), F2: (genFn1() == 42)}
            _ = v24
            _ = v25
        }
        var v26 uint16 = 546
        _ = len(v3)
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
