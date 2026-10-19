package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 int32; F1 int16; F2 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 bool, p1 int8, p2 uint32) int32 {
    var v6 int32 = 460
    var v7 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 uint32, p1 bool, p2 uint16) string {
    if false {
        for fi8 := 0; fi8 < 9; fi8++ {
            p0 = (p0 >> ((1) & 63))
            fmt.Println(fi8, p0)
            _ = fi8
        }
    }
    p1 = true
    if true {
        var v9 int16 = 32767
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}

func main() {
    var v1 int = 0
    var v2 string = (("\n\t" + ("a" + "")) + "\n\t")
    var v3 uint = 18446744073709551615
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    v10 := []int{-9223372036854775808, -938}
    fmt.Println(v2, v3, v2)
    fmt.Println(v3, v2)
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS1{F0: (fi11 * 4294967295), F1: (!true), F2: 1}
        if (!true) {
            var v13 int = fi11
            _ = v13
        }
        v14 := GenS2{F0: genFn0((true && true), -128, 4294967295), F1: 32767, F2: 994}
        _ = fi11
        _ = v12
        _ = v14
    }
    fmt.Println(v3)
    v10 = append(v10, (v1 | v1))
    fmt.Println(v2)
    if (-1 <= -128) {
        v1 = -9223372036854775808
    } else {
        switch (v1 % ((v1) | 1)) {
        case 0:
            var v15 int16 = 451
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := GenS2{F0: 491, F1: 1, F2: -1}
            var v19 int32 = 2147483647
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    v1 = -769
    v2 = v2
    switch v1 {
    case 128:
        v1 += v1
    case -1:
        v20 := GenS3{F0: genFn0(false, -128, 2147483648)}
        _ = v20
    case -497:
        v3 = v3
    }
    fmt.Println(v1)
    switch -878 {
    case 42:
        switch v1 {
        case 2147483647:
            fmt.Println(v1, v2)
        }
    case 915:
        fmt.Println(v2, v2)
    default:
        var v21 string = v2
        _ = v21
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v22 float32 = ((1e10 + 0.0) * ((0.1 / 1.0) + (1e10 + -2.25)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
