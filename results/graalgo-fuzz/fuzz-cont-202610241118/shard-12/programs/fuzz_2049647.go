package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int; F1 int16; F2 uint }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -78
    var v3 int32 = -658
    var v4 float32 = 1e10
    var v5 string = ""
    var v6 uint = 928
    v7 := []int{-2147483648, 127, 708, -997}
    for fi8 := 0; fi8 < 18; fi8++ {
        if (!false) {
            switch (v5 + "\n\t") {
            case "":
                v3 = (v3 ^ 1)
            case "a":
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 float32 = (v4 - (v4 / -2.25))
                    var v11 int8 = (v2 - (v2 | v2))
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case "\n\t":
                var v12 float64 = -2.25
                _ = v12
            default:
                v5 = ("foo bar" + ("a" + "a"))
            }
        }
        switch (fi8 ^ fi8) {
        case 1:
            v4 -= (float32(float64(v4)) + -2.25)
        case 9223372036854775807:
            var v13 uint8 = 127
            _ = v13
        case -1:
            v6 = v6
        }
        _ = fi8
    }
    v4 -= v4
    v5 = (v5 + v5)
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 int32 = (v3 ^ v3)
        if (615 == 0) {
            fmt.Println(v15, v2, v15)
            var v16 uint16 = 579
            var v17 string = (v5 + (v5 + v5))
            _ = v16
            _ = v17
        }
        var v18 int16 = -1
        _ = fi14
        _ = v15
        _ = v18
    }
    v5 = (("hello" + v5) + "hello")
    v19 := GenS1{F0: -1, F1: (!false)}
    fmt.Println(v2)
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := GenS1{F0: -729, F1: false}
        for fi22 := 0; fi22 < 5; fi22++ {
            if ((float32(3.14159) <= 0.0) || ((78 > 2147483648) && (0 != 2147483648))) {
                var v23 uint8 = 2
                _ = v23
            }
            _ = fi22
        }
        v6 = v6
        _ = fi20
        _ = v21
    }
    if true {
        var v24 float64 = float64((float32(1e10) / (v4 / ((-2.25 / v4) + (v4 - v4)))))
        var v25 uint16 = 577
        switch v5 {
        case "foo bar":
            v1[0] = -9223372036854775808
        case "a":
            v26 := GenS1{F0: -935, F1: false}
            _ = v26
        }
        _ = v24
        _ = v25
    }
    if true {
        var v27 uint8 = 244
        _ = v27
    }
    fmt.Println(v4, v6)
    for fi28 := 0; fi28 < 4; fi28++ {
        v6 -= (v6 % ((494) | 1))
        _ = fi28
    }
    if true {
        var v29 uint16 = 0
        _ = v29
    } else {
        v30 := GenS2{F0: 250, F1: 1000, F2: v6}
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
