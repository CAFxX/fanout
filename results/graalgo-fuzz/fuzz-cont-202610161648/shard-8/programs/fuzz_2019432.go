package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 int32, p1 int32) int32 {
    fmt.Println(p0, p1)
    if ("a" < "foo bar") {
        p0 += 42
        v6 := []int{4294967295, 32768, 9223372036854775807, 42}
        _ = v6
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 1
    v3 := []int{-9223372036854775808, 128, 1}
    v4 := make(map[string]bool)
    var v5 int8 = -128
    var v7 string = "a"
    switch -2147483648 {
    case -288:
        v8 := GenS1{F0: ((128 != 32767) || (-2.25 != (1e10 * float32((1.5 * float64(0.1))))))}
        _ = v8
    case 2147483647:
        var v9 uint8 = 0
        _ = v9
    case 32768:
        if true {
            var v10 uint = 42
            var v11 int = -2147483648
            _ = v10
            _ = v11
        } else {
            v2 = v2
        }
    default:
        var v12 int32 = -2147483648
        _ = v12
    }
    var v13 uint = 186
    for fi14 := 0; fi14 < 12; fi14++ {
        v3 = append(v3, fi14)
        _ = fi14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int16 = 32767
        v2 += v2
        switch v7 {
        case "":
            var v17 int32 = -2147483648
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    switch 300 {
    case 2147483647:
        var v18 uint8 = 42
        _ = v18
    }
    v19 := GenS2{F0: v5, F1: 4294967295}
    var v20 uint8 = 127
    v7 = v7
    var v21 float32 = 0.0
    v22 := GenS1{F0: (uint16(v13) != 42)}
    var v23 int = -1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
