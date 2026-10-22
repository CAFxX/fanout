package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32) int8 {
    var v5 uint32 = p0
    v6 := []int{-9223372036854775808, 42, 0}
    _ = p0
    _ = v5
    _ = v6
    return 127
}

func main() {
    var v1 uint16 = 0
    var v2 int64 = 0
    var v3 float64 = (1e10 - -2.25)
    v4 := []uint32{}
    var v7 int64 = -639
    if true {
        fmt.Println(v7, v3, v7)
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 32768, F1: 348}
            switch (fi8 << ((65535) & 63)) {
            case -1:
                v1 -= (v1 % (((^(v1))) | 1))
            case 679:
                v10 := GenS2{F0: 18446744073709551615}
                _ = v10
            case -9223372036854775808:
                v11 := GenS2{F0: 4294967295}
                _ = v11
            }
            v12 := GenS1{F0: (^(fi8)), F1: 1000}
            _ = fi8
            _ = v12
            _ = v9
        }
        if len(v4) > 0 {
            v4[0] = 42
        }
    } else {
        v3 = -2.25
    }
    var v13 uint8 = 125
    switch "hello" {
    case "foo bar":
        v2 = (v2 * v7)
    case "hello":
        var v14 bool = true
        _ = v14
    case "a":
        if ((("" + "foo bar") + ("hello" + ("\n\t" + "\n\t"))) < "a") {
            var v15 float64 = v3
            v4 = append(v4, 2147483648)
            _ = v15
        } else {
            switch 0 {
            case -2147483648:
                var v16 bool = ((3.14159 / 1.0) <= -2.25)
                _ = v16
            }
        }
    default:
        var v17 int16 = -151
        _ = v17
    }
    if false {
        var v18 float64 = v3
        if (!false) {
            var v19 float64 = (v3 / float64(0.1))
            var v20 uint32 = 65535
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            for fi22 := 0; fi22 < 14; fi22++ {
                v1 = 255
                fmt.Println(v2)
                v2 = 42
                _ = fi22
            }
            _ = fi21
        }
        _ = v18
    }
    v3 -= (v3 * (float64(0.0) * float64(0.1)))
    if false {
        if len(v4) > 0 {
            v4[0] = 153
        }
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 uint16 = (v1 << ((2) & 63))
            _ = fi23
            _ = v24
        }
        v7 = (v2 % ((v7) | 1))
    }
    var v25 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
