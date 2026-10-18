package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int64; F2 int64 }
func genFn0() int64 {
    if true {
        var v7 uint16 = 1
        _ = v7
    } else {
        var v8 uint16 = 255
        fmt.Println(v8)
        v8 -= (v8 >> ((608) & 63))
        _ = v8
    }
    switch -235 {
    case -271:
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
    case 2147483647:
        var v11 int32 = 1
        _ = v11
    case -202:
        switch 1 {
        case 127:
            v12 := []int{222}
            _ = v12
        case 0:
            v13 := GenS1{F0: 18446744073709551615, F1: 128}
            _ = v13
        }
    }
    if true {
        for fi14 := 0; fi14 < 11; fi14++ {
            switch "" {
            case "":
                fmt.Println(fi14)
            case "foo bar":
                fmt.Println(fi14)
            }
            _ = fi14
        }
    } else {
        v15 := GenS2{F0: 255}
        v16 := GenS2{F0: 255}
        _ = v15
        _ = v16
    }
    var v17 uint32 = 65535
    v17 -= (^(v17))
    _ = v17
    return 9223372036854775807
}
func genFn1(p0 string, p1 uint32, p2 int8) int {
    v18 := GenS2{F0: 1}
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := []int{}
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 1
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]uint32)
    v3 := []uint32{}
    var v4 int64 = 9223372036854775807
    v5 := []string{"hello", "\n\t"}
    var v6 float64 = 0.0
    v1 = 529
    var v21 uint16 = 184
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 string = ""
        v2[(fi22 ^ fi22)] = 2147483648
        fmt.Println(v21)
        _ = fi22
        _ = v23
    }
    v24 := GenS3{F0: 2, F1: v4, F2: -1}
    var v25 int8 = 42
    v2[32767] = (v1 % ((v1) | 1))
    for fi26 := 0; fi26 < 18; fi26++ {
        v6 = v6
        _ = fi26
    }
    fmt.Println(v21, v6, v25)
    var v27 float64 = v6
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
