package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64, p1 uint32, p2 int32) bool {
    p2 = (p2 >> ((42) & 63))
    var v4 int16 = -656
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return false
}
func genFn1() uint16 {
    if false {
        v5 := []int{-418, -2147483648, -1, -641}
        var v6 float32 = (1e10 - 0.1)
        var v7 float64 = (0.0 * (1.5 / (float64(((v6 - 0.1) * -2.25)) * 0.0)))
        _ = v5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 uint16 = 0
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 132, F1: 18446744073709551615}
    _ = v10
    return 239
}

func main() {
    v1 := []int{-2147483648, 128, -59}
    var v2 int16 = 32767
    v3 := []string{"a"}
    var v11 float32 = 0.1
    fmt.Println(v11, v2)
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            switch "\n\t" {
            case "foo bar":
                v11 = (-2.25 - v11)
            }
            switch ("" + "a") {
            case "":
                v1 = append(v1, (fi12 << ((255) & 63)))
            }
            switch fi12 {
            case 2147483647:
                _ = len(v1)
            case 128:
                var v14 uint = 469
                _ = v14
            case -2147483648:
                v15 := GenS1{F0: 2, F1: 1}
                _ = v15
            }
            _ = fi13
        }
        switch "hello" {
        case "\n\t":
            v11 += float32(1.5)
        case "a":
            var v16 uint = 42
            _ = v16
        }
        _ = fi12
    }
    v17 := GenS2{F0: 1}
    v11 = (v11 * 0.1)
    var v18 float32 = v11
    var v19 float64 = 1e10
    if (-1 <= 127) {
        var v20 int16 = 32767
        for fi21 := 0; fi21 < 15; fi21++ {
            v11 = v18
            _ = fi21
        }
        var v22 uint64 = 42
        _ = v20
        _ = v22
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        v19 -= v19
        v19 = (1e10 / 1.5)
        var v24 int8 = -1
        _ = fi23
        _ = v24
    }
    if false {
        v2 -= (v2 & (-(v2)))
        fmt.Println(v18, v18, v19)
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS1{F0: 489, F1: 1}
            switch (fi25 - -1) {
            case 32767:
                v27 := GenS2{F0: 42}
                _ = v27
            case -9223372036854775808:
                fmt.Println(v18, v2)
            case -2147483648:
                var v28 uint32 = 204
                _ = v28
            default:
                v11 -= v11
            }
            var v29 uint = 1
            _ = fi25
            _ = v26
            _ = v29
        }
    } else {
        var v30 int8 = -5
        _ = v30
    }
    fmt.Println(v11, v18)
    fmt.Println(v19, v11, v18)
    if len(v1) > 0 {
        v1[0] = 0
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
