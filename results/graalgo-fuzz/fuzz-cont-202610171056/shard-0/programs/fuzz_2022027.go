package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0(p0 uint16, p1 uint32, p2 uint8) int32 {
    v7 := GenS1{F0: (p1 << ((uint(p1)) & 63)), F1: ""}
    var v8 float32 = -2.25
    if (false || (!(false && (!true)))) {
        switch 32767 {
        case -2147483648:
            v9 := []int{1, -1, -9223372036854775808, 128}
            _ = v9
        case -1:
            v10 := []int{}
            _ = v10
        }
        v11 := GenS2{F0: -1, F1: -98}
        _ = v11
    }
    var v12 uint16 = uint16(p1)
    switch -570 {
    case -2147483648:
        v12 = (v12 % ((v12) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint = 0
    var v2 bool = (65535 < 0)
    var v3 int8 = -1
    v4 := []uint32{775}
    v5 := []int{4294967295}
    var v6 uint = 0
    if (-32768 != 1000) {
        var v13 int = -1
        _ = v13
    } else {
        var v14 float32 = -2.25
        switch "a" {
        case "hello":
            v3 += (v3 / ((0) | 1))
        case "\n\t":
            var v15 uint32 = 451
            _ = v15
        case "a":
            v1 = (v1 % (((v1 / (((v1 ^ v6)) | 1))) | 1))
        default:
            var v16 uint64 = 226
            _ = v16
        }
        v17 := GenS2{F0: -9223372036854775808, F1: -1}
        _ = v14
        _ = v17
    }
    switch "foo bar" {
    case "a":
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 uint64 = 938
            v6 = (v1 ^ (^(v1)))
            var v20 string = "hello"
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    fmt.Println(v2)
    var v21 bool = v2
    for fi22 := 0; fi22 < 19; fi22++ {
        v1 -= 65535
        _ = fi22
    }
    fmt.Println(v21)
    var v23 uint32 = 0
    switch "hello" {
    case "foo bar":
        v3 -= v3
    default:
        if len(v4) > 0 {
            v4[0] = 1
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
