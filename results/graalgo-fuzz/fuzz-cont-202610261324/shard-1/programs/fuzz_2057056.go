package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 bool }
type GenS2 struct { F0 string; F1 int16 }

func main() {
    var v1 int = -1
    var v2 int8 = 100
    v3 := make(map[string]uint32)
    var v4 float64 = ((0.1 - 1e10) * float64((0.0 - (1.5 * 0.0))))
    v5 := []int{}
    var v6 int64 = -1
    fmt.Println(v2, v2)
    v7 := GenS1{F0: -138, F1: 0, F2: ((^(v6)) != (v6 / ((v6) | 1)))}
    if true {
        var v8 string = "foo bar"
        _ = v8
    }
    var v9 int8 = (v2 ^ (v2 * (v2 | (^(v2)))))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch v1 {
    case 67:
        fmt.Println(v4, v4, v6)
    case 2147483647:
        if (((0 != 32767) || (true && false)) || ((((1e10 * 1.5) - float32((v4 + v4))) - -2.25) == -2.25)) {
            v1 = (v1 + v1)
            for fi10 := 0; fi10 < 18; fi10++ {
                fmt.Println(fi10, fi10)
                fmt.Println(fi10, v4)
                v1 = (v1 / (((v1 + (v1 ^ 42))) | 1))
                _ = fi10
            }
            switch (v1 | (v1 << ((2) & 63))) {
            case -1:
                if (2 >= 42) {
                    v11 := GenS1{F0: -591, F1: 100, F2: (2147483647 != 1)}
                    _ = v11
                } else {
                    v12 := GenS2{F0: "foo bar", F1: 1}
                    _ = v12
                }
            default:
                _ = len(v3)
            }
        } else {
            if ((float64(-2.25) - v4) > -2.25) {
                var v13 float32 = ((1.5 * float32(((v4 + ((float64((((1e10 * (-2.25 / (1.5 / 1.0))) / 1e10) - -2.25)) - float64((((((1e10 * 1e10) + 0.1) - 1e10) * 1.5) * 3.14159))) + ((v4 + 1e10) - v4))) / 1.5))) / 1.0)
                if len(v5) > 0 {
                    v5[0] = 988
                }
                v6 += v6
                _ = v13
            } else {
                v2 += -128
                v9 = (-(v2))
                v14 := GenS2{F0: (("" + ("hello" + "")) + "hello"), F1: 801}
                _ = v14
            }
        }
    }
    switch "hello" {
    case "a":
        for fi15 := 0; fi15 < 20; fi15++ {
            if false {
                v1 = 42
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 bool = true
                _ = fi16
                _ = v17
            }
            switch fi15 {
            case -488:
                fmt.Println(fi15, fi15)
            default:
                var v18 bool = false
                _ = v18
            }
            _ = fi15
        }
    default:
        v9 -= (v2 >> ((804) & 63))
    }
    fmt.Println(v9, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
