package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 string; F2 int64 }
type GenS4 struct { F0 int64; F1 bool; F2 uint8 }
func genFn0(p0 bool, p1 bool, p2 int64) int32 {
    v6 := GenS1{F0: 2147483648}
    p2 -= (-(p2))
    var v7 string = ((((("foo bar" + "foo bar") + (("" + (((("\n\t" + ("" + "hello")) + "hello") + (("a" + "foo bar") + "hello")) + "a")) + "a")) + "hello") + "\n\t") + "a")
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 uint8, p1 uint64) int32 {
    var v8 uint8 = (p0 - p0)
    v9 := GenS2{F0: 65535}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int16 = -824
        v12 := GenS1{F0: 139}
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint16 = 42
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return genFn0(true, ((-1 <= 2147483647) || true), -400)
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 988
    var v3 uint = 0
    var v4 uint16 = 0
    var v5 int64 = 1
    v3 = (v3 ^ v3)
    v15 := GenS3{F0: -1, F1: "foo bar", F2: v5}
    var v16 uint64 = 0
    if (42 >= 2) {
        var v17 string = "foo bar"
        _ = v17
    }
    v3 = v3
    v18 := GenS3{F0: -9223372036854775808, F1: "\n\t", F2: (v1 / ((-1) | 1))}
    switch int(v5) {
    case 240:
        var v19 int8 = -128
        _ = v19
    default:
        if false {
            fmt.Println(v4, v2, v4)
        }
    }
    var v20 float32 = 0.0
    v21 := GenS1{F0: 4294967295}
    v22 := GenS4{F0: 42, F1: (146 == -9223372036854775808), F2: 224}
    var v23 string = "a"
    v24 := []int{-742}
    var v25 string = ("" + v23)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
