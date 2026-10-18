package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint; F2 uint8 }
type GenS4 struct { F0 int32; F1 string; F2 int64 }

func main() {
    var v1 uint32 = 0
    var v2 uint8 = 128
    var v3 bool = true
    if (!v3) {
        fmt.Println(v2, v1)
        if ((2 >= 255) || (!(uint8(v1) < 1))) {
            v4 := GenS1{F0: -828, F1: 0}
            fmt.Println(v2, v1, v1)
            _ = v4
        } else {
            v1 += (v1 + v1)
            v2 = uint8(v1)
        }
        switch -972 {
        case 4294967295:
            var v5 uint64 = 585
            _ = v5
        default:
            v2 += v2
        }
    }
    v6 := GenS2{F0: -128}
    if v3 {
        var v7 uint8 = v2
        fmt.Println(v7)
        v2 = 42
        _ = v7
    } else {
        fmt.Println(v3)
        v8 := GenS2{F0: 73}
        var v9 bool = (9223372036854775807 != 484)
        _ = v8
        _ = v9
    }
    switch 4294967295 {
    case 127:
        var v10 uint8 = v2
        _ = v10
    case -1:
        var v11 uint = 992
        _ = v11
    }
    if (!(v3 || ((-128 > 100) || (42 == -493)))) {
        v3 = (42 <= 0)
        var v12 int = 2147483647
        v13 := GenS3{F0: v12, F1: 255, F2: v2}
        _ = v12
        _ = v13
    } else {
        v2 -= (v2 >> ((2) & 63))
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v2, v3, v3)
        case "":
            v3 = (18446744073709551615 >= 659)
        }
    }
    v14 := GenS4{F0: -1, F1: "", F2: 42}
    for fi15 := 0; fi15 < 6; fi15++ {
        v1 -= v1
        if ((false || v3) || v3) {
            fmt.Println(fi15, v3)
        } else {
            fmt.Println(v2, v2, v1)
        }
        _ = fi15
    }
    var v16 int64 = 1
    var v17 string = ((("a" + (("hello" + "foo bar") + ("" + ("\n\t" + "a")))) + "") + ("" + "\n\t"))
    v18 := []int{1, 128, -1, 32767}
    v1 = v1
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    if len(v18) > 0 {
        v18[0] = 42
    }
    var v19 float32 = float32((1.5 + (((0.1 - 0.0) * (float64(1e10) / 1.0)) + 1.5)))
    fmt.Println(v3, v1, v3)
    switch -2147483648 {
    case -2147483648:
        fmt.Println(v3, v17)
    default:
        switch 740 {
        case -925:
            fmt.Println(v17, v17, v16)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
}
