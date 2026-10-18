package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint8 }
func genFn0(p0 uint32, p1 int) string {
    if (67 < 1) {
        if (-32768 != -32768) {
            var v4 float32 = 1.5
            v5 := GenS1{F0: "hello", F1: p1, F2: 255}
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 int64 = 34
                _ = fi6
                _ = v7
            }
            _ = v4
            _ = v5
        }
    }
    if (!(uint8(p0) >= 222)) {
        p0 = p0
    }
    switch "a" {
    case "foo bar":
        if (18446744073709551615 < 1) {
            v8 := GenS1{F0: "foo bar", F1: -244, F2: 0}
            _ = v8
        } else {
            var v9 uint64 = 18446744073709551615
            var v10 float32 = float32(0.0)
            _ = v10
            _ = v9
        }
    case "hello":
        var v11 bool = (p0 <= 382)
        _ = v11
    case "":
        var v12 float32 = 3.14159
        _ = v12
    }
    _ = p0
    _ = p1
    return "a"
}

func main() {
    var v1 int8 = -89
    v2 := make(map[int]string)
    var v3 bool = true
    v2[1000] = ("" + "hello")
    v13 := v2[128]
    switch "foo bar" {
    case "hello":
        var v14 uint = 2
        _ = v14
    default:
        if v3 {
            for fi15 := 0; fi15 < 10; fi15++ {
                v3 = (!v3)
                _ = fi15
            }
        }
    }
    _ = len(v2)
    switch 4294967295 {
    case 32767:
        _ = len(v2)
    }
    if v3 {
        var v16 int64 = -9223372036854775808
        fmt.Println(v1)
        _ = v16
    }
    v17 := []int{651}
    var v18 bool = v3
    v13 = v13
    fmt.Println(v13, v18)
    v3 = (-916 > 0)
    fmt.Println(v13)
    for fi19 := 0; fi19 < 8; fi19++ {
        fmt.Println(v3, v3)
        _ = fi19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        if (!(0 < 127)) {
            delete(v2, fi20)
        } else {
            var v21 bool = true
            _ = v21
        }
        switch fi20 {
        case -1:
            for fi22 := 0; fi22 < 9; fi22++ {
                v1 += -96
                _ = fi22
            }
        case 42:
            v1 = (v1 - (v1 - v1))
        case 32767:
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        }
        fmt.Println(v3, fi20)
        _ = fi20
    }
    switch 32767 {
    case 127:
        switch v13 {
        case "foo bar":
            var v23 int16 = 1000
            _ = v23
        default:
            var v24 bool = (1 == -1)
            _ = v24
        }
    }
    if v18 {
        fmt.Println(v1)
    } else {
        var v25 int = 0
        v18 = ("foo bar" > v13)
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
