package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 int32, p1 uint64) bool {
    if (2 > 318) {
        var v7 int8 = 100
        _ = v7
    }
    p0 = -2147483648
    fmt.Println(p1, p1)
    if (!false) {
        var v8 int32 = p0
        v9 := GenS1{F0: "\n\t", F1: 1, F2: 36}
        var v10 int64 = 42
        _ = v10
        _ = v8
        _ = v9
    } else {
        v11 := []int{-2147483648, 127, 32768}
        v12 := GenS1{F0: ("foo bar" + ""), F1: 0, F2: 0}
        switch 42 {
        case -512:
            var v13 string = ("a" + "hello")
            _ = v13
        default:
            var v14 string = "hello"
            _ = v14
        }
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    return false
}
func genFn1(p0 uint32, p1 uint16) uint64 {
    var v15 bool = false
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v15
    return 569
}

func main() {
    var v1 float32 = 0.1
    v2 := []int{}
    v3 := []int{1, 32767}
    var v4 int32 = 0
    v5 := make(map[uint32]bool)
    var v6 float64 = -2.25
    v1 = (v1 / v1)
    var v16 int8 = 57
    v17 := v5[618]
    fmt.Println(v16, v4, v6)
    fmt.Println(v16, v16, v6)
    var v18 uint16 = 822
    v19 := GenS2{F0: v18}
    v20 := GenS3{F0: 9223372036854775807, F1: v18}
    var v21 string = ("a" + "a")
    v22 := GenS1{F0: (v21 + "foo bar"), F1: (v16 * v16), F2: 81}
    fmt.Println(v17, v17, v6)
    v23 := GenS2{F0: v18}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
