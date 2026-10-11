package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint }
func genFn0() uint16 {
    var v7 int64 = 1
    var v8 int8 = 42
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 bool, p1 int32, p2 int32) int32 {
    var v9 uint8 = 128
    var v10 string = "foo bar"
    v11 := []int{544, -9223372036854775808, 127}
    p0 = (127 <= (^(v9)))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return p2
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 bool = ("" == "a")
    var v4 int32 = 1
    v5 := make(map[int]int)
    var v6 uint = 255
    switch 9223372036854775807 {
    case 128:
        v3 = ((v3 && v3) && true)
    default:
        v12 := GenS1{F0: 1}
        _ = v12
    }
    switch "a" {
    case "":
        var v13 uint32 = 2147483647
        _ = v13
    case "foo bar":
        v6 = v6
    case "\n\t":
        var v14 uint64 = uint64(v6)
        _ = v14
    }
    var v15 uint64 = 2
    var v16 uint = v6
    if ((true && true) || (-128 <= 127)) {
        v17 := []int{0}
        _ = v17
    }
    v18 := []int{-986}
    if v3 {
        var v19 uint16 = 929
        fmt.Println(v4, v16, v16)
        _ = v19
    }
    switch 42 {
    case 1:
        v20 := GenS2{F0: 9223372036854775807, F1: v16}
        _ = v20
    case 42:
        for fi21 := 0; fi21 < 1; fi21++ {
            if (!false) {
                v16 = 628
            }
            v3 = (false || v3)
            var v22 int32 = (v4 << (((v6 & 280)) & 63))
            _ = fi21
            _ = v22
        }
    case -619:
        v15 = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
