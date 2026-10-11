package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 string }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint8 }
func genFn0() string {
    if true {
        var v5 bool = false
        if (127 < -1) {
            v5 = ((-1 < -9223372036854775808) && v5)
            v5 = v5
        }
        var v6 uint8 = 128
        _ = v5
        _ = v6
    }
    var v7 uint64 = 1
    v7 -= (v7 << ((389) & 63))
    fmt.Println(v7)
    _ = v7
    return ""
}
func genFn1(p0 uint, p1 uint32, p2 string) uint8 {
    var v8 int = 562
    var v9 int64 = 214
    fmt.Println(p2)
    fmt.Println(v9, v8, v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 128
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = -2.25
    var v3 int64 = 9223372036854775807
    v4 := []int{-1, -2147483648}
    v3 = 42
    if true {
        var v10 float64 = v2
        for fi11 := 0; fi11 < 2; fi11++ {
            v10 += v10
            fmt.Println(fi11, v10, v3)
            _ = fi11
        }
        _ = v10
    }
    var v12 uint16 = 32767
    v3 -= v3
    v12 -= 65535
    switch ("hello" + "hello") {
    case "hello":
        if true {
            for fi13 := 0; fi13 < 12; fi13++ {
                fmt.Println(v12, v12, fi13)
                _ = fi13
            }
            v14 := GenS1{F0: ("\n\t" + ""), F1: 1, F2: "foo bar"}
            v2 -= (v2 + v2)
            _ = v14
        }
    case "foo bar":
        v15 := GenS2{F0: 705, F1: 127}
        _ = v15
    case "":
        fmt.Println(v3, v2, v12)
    }
    fmt.Println(v3, v3, v12)
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 float64 = float64(3.14159)
        var v18 uint8 = 2
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: (^(v3)), F1: 128}
    v3 += 0
    if false {
        fmt.Println(v12)
    } else {
        v12 = v12
        for fi20 := 0; fi20 < 6; fi20++ {
            v21 := GenS2{F0: v3, F1: fi20}
            var v22 int16 = -1
            v12 -= v12
            _ = fi20
            _ = v21
            _ = v22
        }
        v2 -= v2
    }
    v12 += (^(v12))
    for fi23 := 0; fi23 < 10; fi23++ {
        v24 := GenS3{F0: 128}
        if false {
            v4 = append(v4, (^(fi23)))
            v25 := v1[uint32(v12)]
            v25 = ("\n\t" == "\n\t")
            _ = v25
        }
        for fi26 := 0; fi26 < 13; fi26++ {
            fmt.Println(v2, v12)
            fmt.Println(v2, v2)
            _ = fi26
        }
        _ = fi23
        _ = v24
    }
    v27 := v1[65535]
    var v28 float64 = (0.0 * v2)
    var v29 float32 = -2.25
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
