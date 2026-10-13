package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint8 }
func genFn0(p0 uint64, p1 bool, p2 uint64) uint8 {
    var v7 float64 = 1e10
    v8 := GenS1{F0: 2147483647, F1: 42}
    if (128 >= 377) {
        switch "a" {
        case "hello":
            v9 := GenS2{F0: 2147483647, F1: 2}
            _ = v9
        default:
            p1 = (133 != 42)
        }
        v10 := []int{802, 0, -1, 702}
        _ = v10
    }
    p2 -= 85
    v11 := GenS3{F0: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int32, p1 int, p2 bool) string {
    fmt.Println(p2, p1)
    var v12 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return ""
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]int)
    var v5 float64 = 1.5
    v6 := make(map[string]string)
    v1 += v1
    v5 -= -2.25
    switch "a" {
    case "a":
        v13 := GenS4{F0: 2147483647, F1: -370, F2: 1}
        _ = v13
    default:
        fmt.Println(v1)
    }
    fmt.Println(v5, v5)
    if (0 <= 2) {
        v14 := []int{(v1 | (^(v1))), (v1 - 0)}
        var v15 float64 = 0.1
        _ = v14
        _ = v15
    } else {
        v5 = v5
        v5 += 3.14159
    }
    var v16 bool = ("foo bar" != (("" + ("" + "foo bar")) + "\n\t"))
    fmt.Println(v5)
    var v17 int64 = 0
    var v18 uint32 = 301
    var v19 int = (v1 | (v1 + v1))
    var v20 uint8 = 255
    var v21 uint16 = 65535
    fmt.Println(v21)
    fmt.Println(v16, v1)
    if v16 {
        v22 := GenS3{F0: 42}
        var v23 int = (^(v1))
        v24 := GenS5{F0: v18, F1: 2, F2: genFn0(4294967295, ((((true && false) || v16) || v16) && true), 0)}
        _ = v22
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
