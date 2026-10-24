package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
func genFn0(p0 int16, p1 uint, p2 uint64) int32 {
    var v5 uint16 = 65535
    v6 := GenS1{F0: p0, F1: ""}
    var v7 uint16 = v5
    var v8 int = 832
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -71
}

func main() {
    v1 := []uint32{2147483648, 922, 2147483648, 2}
    v2 := []string{"a", "hello", "foo bar"}
    var v3 int64 = -9223372036854775808
    v4 := make(map[int]bool)
    fmt.Println(v3)
    fmt.Println(v3)
    v9 := GenS2{F0: 127}
    for fi10 := 0; fi10 < 17; fi10++ {
        v3 += (v3 - v3)
        _ = fi10
    }
    var v11 uint32 = 1
    switch "hello" {
    case "hello":
        var v12 int32 = -2147483648
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v3 -= (-(v3))
        var v14 uint = 1
        _ = fi13
        _ = v14
    }
    var v15 uint64 = 547
    v16 := GenS3{F0: 882}
    fmt.Println(v11, v11, v11)
    if false {
        if true {
            fmt.Println(v3, v3)
            switch ("a" + "foo bar") {
            case "":
                v3 = -149
            case "hello":
                v17 := GenS4{F0: true}
                _ = v17
            case "foo bar":
                var v18 int32 = -1
                _ = v18
            }
            if (1 > -2147483648) {
                v11 -= 0
            } else {
                _ = len(v4)
            }
        }
        var v19 string = ("hello" + "foo bar")
        v11 = v11
        _ = v19
    } else {
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 int16 = 42
            var v22 uint64 = v15
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    if len(v1) > 0 {
        v1[0] = v11
    }
    fmt.Println(v11, v3, v11)
    v4[-948] = false
    var v23 string = ("foo bar" + ("hello" + "foo bar"))
    v24 := GenS1{F0: -1, F1: v23}
    if true {
        var v25 int8 = 100
        var v26 uint8 = 128
        _ = v25
        _ = v26
    }
    v11 -= v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
