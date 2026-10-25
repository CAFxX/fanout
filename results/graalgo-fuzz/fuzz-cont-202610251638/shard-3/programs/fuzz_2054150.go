package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 uint8; F1 int; F2 uint }
func genFn0(p0 uint64, p1 uint32, p2 uint8) uint32 {
    switch "hello" {
    case "foo bar":
        var v6 int = 32768
        _ = v6
    }
    v7 := []int{42}
    v8 := GenS1{F0: 0}
    p0 -= (p0 / ((347) | 1))
    if (-111 == -13) {
        var v9 uint32 = uint32(p2)
        _ = v9
    } else {
        var v10 int64 = 1
        v11 := GenS2{F0: 868, F1: (^(v10)), F2: (v10 >> ((255) & 63))}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 % ((p1) | 1))
}
func genFn1(p0 uint32, p1 uint8) uint16 {
    var v12 int8 = 95
    v13 := GenS3{F0: 65535, F1: "hello"}
    v14 := []int{}
    fmt.Println(v12, p0)
    var v15 bool = (true || (9223372036854775807 == 1))
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 uint8 = 34
    v2 := make(map[int]bool)
    var v3 int = 2147483647
    var v4 uint8 = 127
    var v5 float64 = 3.14159
    delete(v2, (v3 / (((^(v3))) | 1)))
    v16 := GenS4{F0: -800, F1: 0}
    var v17 uint8 = (^(v1))
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 uint = 476
        v20 := []int{}
        v1 -= 255
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int32 = 1
    fmt.Println(v1)
    v22 := GenS5{F0: (!(-32768 >= 1))}
    if ((18446744073709551615 < 217) && false) {
        for fi23 := 0; fi23 < 17; fi23++ {
            v5 += (v5 - v5)
            _ = fi23
        }
        if ((false || (false && (2 != 4294967295))) || (((1 <= 404) && false) || (4294967295 <= 18446744073709551615))) {
            fmt.Println(v4)
            v24 := []int{-789, v3}
            v21 = v21
            _ = v24
        } else {
            var v25 uint32 = 2147483647
            var v26 int64 = -1
            fmt.Println(v3)
            _ = v25
            _ = v26
        }
        fmt.Println(v21)
    } else {
        for fi27 := 0; fi27 < 4; fi27++ {
            v28 := GenS1{F0: 4294967295}
            v3 = (fi27 + 4294967295)
            v29 := GenS6{F0: 0, F1: 815}
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    var v30 bool = (412 > genFn1(144, (v1 / (((v17 | 101)) | 1))))
    var v31 int8 = -128
    fmt.Println(v4, v4)
    v32 := GenS5{F0: v30}
    v33 := GenS7{F0: (v4 | v4), F1: (-(v3)), F2: 587}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
