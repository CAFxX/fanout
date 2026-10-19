package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int8 }
func genFn0() uint {
    var v4 int16 = -32768
    if false {
        v5 := []int{4294967295, 127, 128, 4294967295}
        _ = v5
    }
    var v6 uint64 = 18446744073709551615
    _ = v4
    _ = v6
    return 1
}
func genFn1(p0 int16) uint8 {
    v7 := []int{2147483647, 32768, 42, -9223372036854775808}
    fmt.Println(p0)
    v8 := GenS1{F0: 18446744073709551615, F1: 764}
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -1
    var v3 uint32 = 2147483648
    delete(v1, (v3 - (v3 % ((v3) | 1))))
    v3 = v3
    v9 := []int{-299, 0, 9223372036854775807}
    if true {
        var v10 int8 = 0
        _ = v10
    } else {
        var v11 uint64 = 2
        switch "a" {
        case "hello":
            fmt.Println(v3, v11)
        case "foo bar":
            v12 := GenS2{F0: 0}
            _ = v12
        case "":
            v1[v3] = 2147483647
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            v2 += 155
            var v14 string = "a"
            _ = fi13
            _ = v14
        }
        _ = v11
    }
    var v15 int64 = 42
    v15 = (v15 + v15)
    var v16 uint8 = 1
    for fi17 := 0; fi17 < 20; fi17++ {
        fmt.Println(v15, v15)
        _ = fi17
    }
    v18 := GenS2{F0: -1}
    v19 := GenS2{F0: -128}
    var v20 bool = (!(false || true))
    v21 := GenS1{F0: 403, F1: 0}
    for fi22 := 0; fi22 < 17; fi22++ {
        v16 = (v16 / ((v16) | 1))
        var v23 bool = (v20 || v20)
        fmt.Println(fi22)
        _ = fi22
        _ = v23
    }
    var v24 uint = genFn0()
    v25 := v1[v3]
    var v26 uint32 = v25
    var v27 int32 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v9
}
