package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 string; F1 uint16 }
func genFn0(p0 int, p1 uint, p2 uint32) int32 {
    v7 := []int{p0, p0}
    if len(v7) > 0 {
        v7[0] = (p0 + (p0 ^ p0))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}

func main() {
    var v1 string = "hello"
    var v2 bool = (false && true)
    v3 := make(map[uint32]int)
    var v4 int16 = 1000
    v5 := make(map[int]bool)
    v6 := []uint32{2, 65535, 2147483647, 699}
    var v8 string = ("a" + "a")
    switch "" {
    case "foo bar":
        var v9 uint32 = 2147483648
        _ = v9
    case "a":
        if len(v6) > 0 {
            v6[0] = 1
        }
    default:
        delete(v5, -1)
    }
    v10 := GenS1{F0: -521}
    for fi11 := 0; fi11 < 16; fi11++ {
        v4 += v4
        v12 := GenS1{F0: 0}
        v13 := GenS2{F0: -9223372036854775808, F1: v2}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = len(v3)
    switch "\n\t" {
    case "hello":
        var v14 int64 = 42
        _ = v14
    default:
        if v2 {
            var v15 int8 = 127
            delete(v3, 2147483647)
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        switch v1 {
        case "foo bar":
            switch (-(fi16)) {
            case 4294967295:
                v8 = (((v1 + v1) + v8) + v8)
            case 0:
                _ = len(v6)
            case 128:
                _ = len(v3)
            }
        case "":
            var v17 float64 = 3.14159
            _ = v17
        case "hello":
            v5[fi16] = (v2 && (!v2))
        }
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS3{F0: (^(fi18)), F1: -75}
            switch (-(fi18)) {
            case -1:
                fmt.Println(fi16, fi16, fi18)
            case 127:
                v20 := GenS2{F0: 1, F1: false}
                _ = v20
            case 2147483647:
                _ = len(v6)
            default:
                v3[42] = 4294967295
            }
            fmt.Println(v1, v4, v8)
            _ = fi18
            _ = v19
        }
        v2 = false
        _ = fi16
    }
    v4 = v4
    if (-1 < 42) {
        for fi21 := 0; fi21 < 10; fi21++ {
            delete(v5, fi21)
            if v2 {
                v2 = v2
                var v22 uint = 1
                v1 = v1
                _ = v22
            }
            _ = fi21
        }
        delete(v5, 128)
    } else {
        v23 := GenS4{F0: 42, F1: 84}
        v24 := GenS5{F0: "\n\t", F1: 716}
        _ = v23
        _ = v24
    }
    v1 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
