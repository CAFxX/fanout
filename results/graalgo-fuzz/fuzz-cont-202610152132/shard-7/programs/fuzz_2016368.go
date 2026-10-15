package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
func genFn0(p0 uint) uint {
    var v7 float32 = (float32(1.5) * 0.0)
    for fi8 := 0; fi8 < 9; fi8++ {
        if (128 <= 1) {
            fmt.Println(fi8, fi8, p0)
            v7 -= (v7 * 0.0)
        } else {
            v7 = v7
        }
        fmt.Println(fi8)
        if (p0 > p0) {
            var v9 int64 = 0
            v10 := GenS1{F0: 193, F1: true, F2: true}
            v11 := []int{fi8}
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi8
    }
    v12 := []int{-2147483648, 32768, 9223372036854775807, 2147483647}
    var v13 bool = false
    _ = p0
    _ = v12
    _ = v13
    _ = v7
    return (p0 >> ((p0) & 63))
}
func genFn1(p0 int, p1 int16, p2 uint64) bool {
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint64 = 1
        _ = fi14
        _ = v15
    }
    fmt.Println(p1)
    p2 -= 786
    _ = p0
    _ = p1
    _ = p2
    return (float32((-2.25 / 3.14159)) >= 0.1)
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    var v3 int8 = 1
    v4 := []int{0, 32768, -2147483648, 128}
    v5 := make(map[string]uint32)
    v6 := []uint32{0, 4294967295}
    if true {
        fmt.Println(v3)
    }
    if (1 > 991) {
        fmt.Println(v3)
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v1[""] = 0
    v3 = v3
    if true {
        var v16 float32 = 1.5
        for fi17 := 0; fi17 < 7; fi17++ {
            v16 += 0.0
            fmt.Println(v3)
            _ = fi17
        }
        _ = v16
    } else {
        fmt.Println(v3)
        var v18 int16 = 1000
        var v19 int16 = (v18 * v18)
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 int8 = v3
        v21 += v21
        _ = fi20
        _ = v21
    }
    var v22 int64 = 1
    var v23 uint8 = 2
    v23 = 42
    fmt.Println(v22, v22, v22)
    if false {
        for fi24 := 0; fi24 < 6; fi24++ {
            _ = len(v4)
            v22 = (v22 >> ((23) & 63))
            _ = fi24
        }
        fmt.Println(v22)
    } else {
        v23 = v23
        if true {
            var v25 uint8 = v23
            _ = v25
        }
        var v26 string = ("hello" + "\n\t")
        _ = v26
    }
    v22 = int64(v23)
    switch ("foo bar" + "\n\t") {
    case "hello":
        v3 = v3
    case "\n\t":
        var v27 int64 = (v22 >> ((4294967295) & 63))
        _ = v27
    case "foo bar":
        if (false && ("foo bar" != "hello")) {
            v3 -= 1
        } else {
            switch 32768 {
            case 2147483647:
                v3 -= -4
            case -1:
                var v28 int = 149
                _ = v28
            case -900:
                v29 := GenS1{F0: 2, F1: (!(1.5 == 1e10)), F2: true}
                _ = v29
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
