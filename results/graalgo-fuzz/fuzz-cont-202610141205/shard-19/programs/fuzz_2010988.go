package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint) int8 {
    var v5 uint = p0
    v6 := GenS1{F0: 466, F1: 32767, F2: 2147483647}
    _ = p0
    _ = v5
    _ = v6
    return 61
}
func genFn1(p0 uint32, p1 uint32, p2 uint) string {
    var v7 string = "a"
    v8 := GenS2{F0: -955}
    v9 := GenS3{F0: p2}
    if (p2 >= 0) {
        v10 := []int{-1, 32767, -2147483648}
        _ = v10
    } else {
        v7 = "\n\t"
        fmt.Println(p0, p1, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    var v1 int = 0
    v2 := []string{"hello"}
    var v3 uint8 = 241
    var v4 uint64 = 1
    v11 := GenS3{F0: 42}
    var v12 int32 = -527
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        switch "\n\t" {
        case "foo bar":
            v1 = v1
        case "hello":
            fmt.Println(v3, v3, v12)
        default:
            v13 := GenS4{F0: v4}
            _ = v13
        }
    }
    v14 := GenS5{F0: 42, F1: 42}
    v3 = (^(v3))
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    if (-11 < genFn0(1)) {
        var v15 int8 = -60
        _ = v15
    }
    v16 := GenS1{F0: 1, F1: 0, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
