package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }
type GenS2 struct { F0 uint16; F1 string; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 int8, p1 int32) int16 {
    if (0 <= 1) {
        v4 := GenS1{F0: 255, F1: p0, F2: p1}
        if true {
            v5 := []int{}
            p1 = p1
            var v6 uint64 = 18446744073709551615
            _ = v5
            _ = v6
        }
        v7 := GenS1{F0: 18446744073709551615, F1: p0, F2: 1}
        _ = v4
        _ = v7
    }
    v8 := []int{0, 128, int(p0)}
    _ = p0
    _ = p1
    _ = v8
    return 42
}
func genFn1(p0 uint, p1 int64) bool {
    if false {
        for fi9 := 0; fi9 < 7; fi9++ {
            fmt.Println(p1, p1)
            var v10 int16 = genFn0(-128, 0)
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS2{F0: 1, F1: ("hello" + ((("a" + (((("a" + "foo bar") + "foo bar") + "") + "hello")) + ("" + ("foo bar" + "\n\t"))) + "a")), F2: 255}
    v12 := []int{-9223372036854775808, 2147483647, -1, 0}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return false
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 100
    var v3 uint8 = 249
    v2 -= (v2 * v2)
    v3 = (v3 ^ 0)
    var v13 int32 = 1
    var v14 uint64 = 4294967295
    fmt.Println(v13, v13, v13)
    switch "foo bar" {
    case "\n\t":
        v13 -= -175
    case "":
        v15 := GenS2{F0: 2, F1: "foo bar", F2: 65535}
        _ = v15
    case "a":
        switch "foo bar" {
        case "":
            v14 -= v14
        case "hello":
            v16 := []int{128, -2147483648, -245, -815}
            _ = v16
        case "\n\t":
            var v17 bool = (0.0 <= 3.14159)
            _ = v17
        default:
            var v18 float64 = 1.5
            _ = v18
        }
    }
    var v19 uint8 = v3
    v20 := GenS2{F0: 65535, F1: "a", F2: 322}
    if ((!(1 >= 32768)) || false) {
        var v21 uint8 = 42
        var v22 int64 = 269
        fmt.Println(v2)
        _ = v21
        _ = v22
    } else {
        v23 := GenS2{F0: 712, F1: ("foo bar" + ("hello" + "a")), F2: 1}
        _ = v23
    }
    fmt.Println(v13, v14)
    v24 := GenS3{F0: (v19 + v3), F1: 32767}
    for fi25 := 0; fi25 < 10; fi25++ {
        v14 = 27
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
}
