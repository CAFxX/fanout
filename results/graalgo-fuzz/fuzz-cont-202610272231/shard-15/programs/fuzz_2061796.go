package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint64 = 0
    v2 := []int{-989, 9223372036854775807, -9223372036854775808, 128}
    var v3 float32 = ((float32((0.1 * -2.25)) / 1.5) + 3.14159)
    var v4 int64 = 42
    var v5 uint8 = 128
    var v6 int32 = 729
    var v7 int64 = v4
    var v8 int32 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3, v6)
    v9 := GenS1{F0: "\n\t"}
    switch 0 {
    case 0:
        v10 := GenS1{F0: "foo bar"}
        _ = v10
    case -1:
        v11 := GenS1{F0: ("hello" + "\n\t")}
        _ = v11
    case 4294967295:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if (v3 >= v3) {
        if true {
            switch -1 {
            case 1:
                v4 += v7
            case -429:
                for fi12 := 0; fi12 < 1; fi12++ {
                    v13 := GenS1{F0: "\n\t"}
                    _ = fi12
                    _ = v13
                }
            case -1:
                fmt.Println(v1, v3)
            default:
                for fi14 := 0; fi14 < 17; fi14++ {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        if (2 <= 569) {
                            _ = len(v2)
                        }
                        _ = fi15
                    }
                    _ = fi14
                }
            }
            for fi16 := 0; fi16 < 12; fi16++ {
                switch fi16 {
                case 32767:
                    v6 -= (v8 ^ (v8 ^ v6))
                }
                if true {
                    var v17 int = fi16
                    _ = v17
                }
                fmt.Println(v4, fi16)
                _ = fi16
            }
        }
        v18 := GenS2{F0: true, F1: 12}
        v2 = append(v2, -447)
        _ = v18
    }
    switch 42 {
    case 32767:
        fmt.Println(v7, v3)
    case 1:
        _ = len(v2)
    default:
        if (-32768 <= -327) {
            v8 -= -859
        }
    }
    v4 += v4
    v19 := GenS3{F0: -128}
    v20 := GenS3{F0: 1}
    var v21 int16 = 1000
    if (!(!true)) {
        switch 32768 {
        case -483:
            if len(v2) > 0 {
                v2[0] = 930
            }
        case 4294967295:
            for fi22 := 0; fi22 < 16; fi22++ {
                v23 := GenS2{F0: true, F1: 128}
                _ = fi22
                _ = v23
            }
        default:
            v8 -= 801
        }
        v3 += ((3.14159 * -2.25) / v3)
    }
    var v24 int = 127
    v6 -= v6
    v2 = append(v2, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
