package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint32; F1 string; F2 int64 }

func main() {
    var v1 int8 = 127
    var v2 uint16 = 979
    var v3 int = 42
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]string)
    var v7 uint = 420
    fmt.Println(v2, v1, v7)
    for fi8 := 0; fi8 < 15; fi8++ {
        switch "a" {
        case "":
            fmt.Println(v7)
        }
        _ = fi8
    }
    var v9 float64 = (-2.25 + 3.14159)
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int8 = v1
        v9 = v9
        _ = fi10
        _ = v11
    }
    if ((false && true) && true) {
        v7 = (v7 / ((255) | 1))
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 float32 = 0.1
            var v14 int16 = -87
            v15 := v6["\n\t"]
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        v16 := GenS1{F0: false, F1: (v2 >> (((v7 ^ (v7 / ((v7) | 1)))) & 63))}
        _ = v16
    } else {
        v1 = (v1 % (((-(v1))) | 1))
        var v17 int = 127
        v9 = -2.25
        _ = v17
    }
    var v18 int32 = -2147483648
    v19 := GenS1{F0: (v2 <= v2), F1: (v2 & (v2 + uint16(v7)))}
    if true {
        fmt.Println(v18, v3)
    }
    switch "a" {
    case "hello":
        v20 := []int{}
        _ = v20
    case "a":
        v2 += v2
    default:
        fmt.Println(v1)
    }
    var v21 int16 = 1000
    if (v21 > v21) {
        if (!(!false)) {
            switch (v3 | v3) {
            case 4294967295:
                v18 -= (v18 % (((v18 & (v18 & -1))) | 1))
            default:
                v22 := GenS2{F0: 2147483648, F1: "foo bar", F2: 398}
                _ = v22
            }
            var v23 int = int(v2)
            _ = v23
        } else {
            v18 = v18
            v1 = (v1 * v1)
            _ = len(v5)
        }
    }
    v2 = v2
    if true {
        for fi24 := 0; fi24 < 7; fi24++ {
            if true {
                switch fi24 {
                case 42:
                    v18 = int32(v21)
                case -214:
                    v9 = (v9 - (v9 / float64((((0.1 * float32((v9 / (((1e10 - (v9 / v9)) / 1e10) + v9)))) * (float32(v9) * 0.0)) - 3.14159))))
                default:
                    var v25 float64 = (3.14159 / 1.0)
                    _ = v25
                }
                v1 = (^(v1))
                var v26 uint32 = 895
                _ = v26
            }
            _ = fi24
        }
    } else {
        v27 := GenS2{F0: 2147483647, F1: ("hello" + ("" + (("" + "hello") + ""))), F2: 1}
        v3 -= 32767
        fmt.Println(v2, v18)
        _ = v27
    }
    var v28 uint16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
