package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 bool }
func genFn0(p0 int32, p1 int32) string {
    if ((((1e10 / 0.1) + (0.1 / 0.1)) - 1.5) <= (((-2.25 / 0.1) * float32(1e10)) + 0.1)) {
        fmt.Println(p0, p1)
    } else {
        if (42 <= -9223372036854775808) {
            v7 := GenS1{F0: "foo bar", F1: 42, F2: true}
            _ = v7
        } else {
            var v8 bool = (42 != 0)
            fmt.Println(p0, p1)
            _ = v8
        }
    }
    v9 := []int{-9223372036854775808}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v9
    return ("" + "\n\t")
}
func genFn1() string {
    var v10 uint32 = 4294967295
    switch "foo bar" {
    case "":
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := []int{(fi11 - (fi11 + 202))}
            _ = fi11
            _ = v12
        }
    }
    v13 := GenS1{F0: "hello", F1: 511, F2: false}
    _ = v10
    _ = v13
    return "\n\t"
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 1e10
    var v3 uint8 = 128
    v4 := []int{}
    v5 := make(map[uint32]uint32)
    var v6 string = "a"
    var v14 int32 = 0
    fmt.Println(v14)
    fmt.Println(v6, v14, v6)
    delete(v5, 1)
    v5[2147483648] = 624
    var v15 uint8 = v3
    switch v6 {
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        delete(v5, 0)
    }
    if true {
        v3 = (v15 ^ v3)
    } else {
        _ = len(v1)
        v14 += (v14 % (((v14 * v14)) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
