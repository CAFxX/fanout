package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint8; F1 int64; F2 int32 }
func genFn0(p0 bool, p1 int16) bool {
    p0 = p0
    var v4 int = 128
    _ = p0
    _ = p1
    _ = v4
    return p0
}
func genFn1(p0 int16, p1 int32) uint32 {
    p0 = p0
    v5 := GenS1{F0: 554}
    var v6 int32 = p1
    v7 := []int{2147483647, -677, 598}
    v8 := GenS2{F0: 69, F1: ""}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -1
    v3 := make(map[string]string)
    if (-87 >= -9223372036854775808) {
        v9 := []int{-998}
        _ = v9
    } else {
        v2 += 9223372036854775807
        fmt.Println(v2)
        fmt.Println(v2)
    }
    v10 := GenS2{F0: 128, F1: "\n\t"}
    v2 += -1
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        switch -974 {
        case 42:
            var v11 int64 = 1
            _ = v11
        case 476:
            v12 := GenS3{F0: 0}
            _ = v12
        case -483:
            for fi13 := 0; fi13 < 8; fi13++ {
                v2 = (v2 * (v2 | (v2 ^ v2)))
                _ = fi13
            }
        }
    case "a":
        fmt.Println(v2)
    case "hello":
        var v14 int8 = -128
        _ = v14
    }
    var v15 int64 = 1
    switch ("\n\t" + ("" + ("hello" + (("hello" + "foo bar") + "a")))) {
    case "\n\t":
        v2 = (v15 / ((v15) | 1))
    case "foo bar":
        for fi16 := 0; fi16 < 14; fi16++ {
            v2 = (v15 * 9223372036854775807)
            var v17 int16 = 1000
            _ = fi16
            _ = v17
        }
    }
    v18 := []int{}
    v19 := v1[""]
    var v20 uint16 = 65535
    _ = len(v3)
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 uint32 = 2147483647
        v23 := GenS4{F0: 0, F1: 0, F2: genFn1(1, 1)}
        v20 = (v20 | v20)
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS5{F0: uint8(v20), F1: int64(v20), F2: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
}
