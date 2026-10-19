package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int64; F1 int16 }

func main() {
    var v1 uint8 = 202
    var v2 float64 = float64(float32(1.5))
    var v3 float64 = float64((0.1 / ((-2.25 - 0.1) + float32(3.14159))))
    var v4 uint64 = 4294967295
    var v5 float32 = ((0.0 - 0.0) * 1e10)
    v6 := []uint32{1}
    v5 = 1.5
    switch -485 {
    case -9223372036854775808:
        v7 := GenS1{F0: 287}
        _ = v7
    }
    fmt.Println(v3)
    if (0 != 42) {
        for fi8 := 0; fi8 < 6; fi8++ {
            if false {
                var v9 int64 = int64(v4)
                _ = v9
            }
            fmt.Println(v5, v2)
            v3 = v3
            _ = fi8
        }
    } else {
        switch ("hello" + "foo bar") {
        case "":
            var v10 int8 = -22
            _ = v10
        case "\n\t":
            _ = len(v6)
        case "hello":
            var v11 uint32 = 1
            _ = v11
        default:
            var v12 uint8 = v1
            _ = v12
        }
        if true {
            fmt.Println(v3)
            var v13 uint = 714
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 uint8 = 225
                var v16 int = -9223372036854775808
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v13
        }
        v17 := GenS1{F0: 1}
        _ = v17
    }
    switch "hello" {
    case "\n\t":
        v6 = append(v6, 0)
    case "":
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 uint32 = 2
            if (false || true) {
                v2 += ((1.5 / v2) - v3)
                v6 = append(v6, (v19 >> ((647) & 63)))
                v20 := GenS1{F0: 255}
                _ = v20
            } else {
                fmt.Println(v2, v5, fi18)
                _ = len(v6)
            }
            _ = fi18
            _ = v19
        }
    case "hello":
        var v21 int16 = 749
        _ = v21
    default:
        var v22 uint = 1
        _ = v22
    }
    if (true && (127 >= 0)) {
        v23 := GenS1{F0: 42}
        v24 := GenS2{F0: 127, F1: 548}
        v25 := GenS1{F0: 42}
        _ = v23
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 7; fi26++ {
            if false {
                if false {
                    fmt.Println(v4, fi26, v1)
                }
                var v27 float64 = ((1e10 + v2) - v3)
                var v28 int8 = 127
                _ = v27
                _ = v28
            }
            _ = fi26
        }
        switch "foo bar" {
        case "\n\t":
            var v29 float64 = 1.5
            _ = v29
        default:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        v30 := GenS3{F0: 9223372036854775807, F1: 0}
        _ = v30
    }
    fmt.Println(v2, v5, v2)
    v1 += 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
