package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int64; F2 int32 }
func genFn0(p0 uint64, p1 uint) int64 {
    p0 = 616
    v6 := []int{int(p1), 127}
    v7 := GenS1{F0: -820, F1: p0}
    var v8 int16 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 uint, p1 int8) int32 {
    if true {
        switch ("" + "a") {
        case "hello":
            switch 42 {
            case 42:
                v9 := []int{-1, -1}
                _ = v9
            default:
                var v10 uint16 = 265
                _ = v10
            }
        case "":
            var v11 uint32 = 509
            _ = v11
        default:
            var v12 int32 = 42
            _ = v12
        }
        v13 := []int{4294967295, 1, 127}
        _ = v13
    } else {
        var v14 float64 = (0.0 + 3.14159)
        _ = v14
    }
    switch 127 {
    case 128:
        v15 := GenS1{F0: 837, F1: 0}
        _ = v15
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int8 = -128
    v2 := []string{"a", ""}
    var v3 int64 = -1
    var v4 string = "a"
    v5 := make(map[int]bool)
    if false {
        if true {
            var v16 uint = 255
            var v17 string = ""
            _ = v16
            _ = v17
        } else {
            v18 := GenS1{F0: 42, F1: 1}
            _ = v18
        }
    }
    var v19 int64 = (v3 + v3)
    for fi20 := 0; fi20 < 7; fi20++ {
        if ((-464 != 306) && false) {
            var v21 int64 = genFn0(155, 602)
            v22 := GenS1{F0: -320, F1: 1}
            _ = v21
            _ = v22
        }
        var v23 uint64 = 0
        if (127 == 42) {
            v24 := GenS1{F0: -32768, F1: v23}
            _ = v24
        }
        _ = fi20
        _ = v23
    }
    v25 := GenS2{F0: v19}
    v4 = "foo bar"
    _ = len(v2)
    v26 := GenS3{F0: 88, F1: 9223372036854775807, F2: 2147483647}
    var v27 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
