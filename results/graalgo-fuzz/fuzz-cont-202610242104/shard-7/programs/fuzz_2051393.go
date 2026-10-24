package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64, p1 uint, p2 uint32) int8 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 float64 = 0.1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 1, F1: 175, F2: (p1 ^ (p1 & 28))}
    var v10 int32 = int32(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 117
}
func genFn1(p0 int64, p1 string, p2 string) uint {
    var v11 int = -3
    switch ("hello" + p2) {
    case "":
        v12 := GenS1{F0: -1, F1: 1, F2: 18446744073709551615}
        _ = v12
    case "foo bar":
        if (-1 <= -1) {
            fmt.Println(p0, p1)
        } else {
            p2 = p1
        }
    case "\n\t":
        p0 += p0
    default:
        v13 := GenS1{F0: 773, F1: 42, F2: 665}
        _ = v13
    }
    var v14 uint16 = 2
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 int16 = 32767
        p1 = ("hello" + p2)
        _ = fi15
        _ = v16
    }
    var v17 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v14
    _ = v17
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    v3 := []uint32{}
    v4 := []int{-749, 0, 32767, -9223372036854775808}
    v5 := make(map[string]int)
    v6 := []string{"a", ((("" + "\n\t") + "foo bar") + "a")}
    var v18 int32 = 945
    for fi19 := 0; fi19 < 3; fi19++ {
        v18 += (v18 * (v18 | (v18 << ((712) & 63))))
        fmt.Println(v18, v18)
        _ = fi19
    }
    v20 := GenS2{F0: true}
    v18 -= v18
    v21 := GenS2{F0: false}
    v18 -= v18
    var v22 int64 = 9223372036854775807
    switch 4294967295 {
    case 42:
        v18 = (-(v18))
    }
    if ((false && false) || (0 >= 790)) {
        fmt.Println(v22)
        v23 := GenS3{F0: 1000}
        _ = v23
    } else {
        v24 := GenS1{F0: (v18 >> ((4294967295) & 63)), F1: (v18 & (v18 % ((v18) | 1))), F2: genFn1(-1, "\n\t", ("hello" + "a"))}
        for fi25 := 0; fi25 < 2; fi25++ {
            var v26 int64 = 703
            _ = fi25
            _ = v26
        }
        fmt.Println(v18, v22)
        _ = v24
    }
    var v27 bool = false
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
