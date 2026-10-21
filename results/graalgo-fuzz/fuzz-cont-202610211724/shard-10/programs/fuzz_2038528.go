package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint64 }

func main() {
    var v1 float32 = float32(-2.25)
    var v2 bool = true
    var v3 uint = 42
    v4 := GenS1{F0: 42}
    v3 += (v3 - 1)
    var v5 float32 = v1
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{-9223372036854775808, fi6, -9223372036854775808, 32768}
        fmt.Println(fi6, fi6)
        _ = fi6
        _ = v7
    }
    switch 9223372036854775807 {
    case -312:
        for fi8 := 0; fi8 < 20; fi8++ {
            v2 = v2
            _ = fi8
        }
    default:
        switch ("hello" + "hello") {
        case "\n\t":
            v9 := []int{32767, 1}
            _ = v9
        case "a":
            var v10 float64 = 1e10
            _ = v10
        case "":
            var v11 int64 = 700
            _ = v11
        default:
            v1 -= v5
        }
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v1 = 1.5
        v13 := []int{}
        for fi14 := 0; fi14 < 8; fi14++ {
            if false {
                v2 = (-2147483648 >= 0)
                if len(v13) > 0 {
                    v13[0] = fi12
                }
            }
            v5 = v5
            var v15 float32 = (1.5 * (((v1 / 1.5) / (float32(-2.25) + 0.0)) - 0.0))
            if len(v13) > 0 {
                v13[0] = fi12
            }
            if (!true) {
                v15 = float32(((3.14159 * (0.0 + 0.1)) / 1.5))
                v13 = append(v13, (fi12 & (-(fi12))))
                fmt.Println(fi14, fi12)
            }
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v1 -= ((v1 / ((3.14159 * v1) + v1)) - v1)
    var v16 uint8 = 186
    if true {
        for fi17 := 0; fi17 < 20; fi17++ {
            switch "foo bar" {
            case "":
                if false {
                    v18 := []int{(fi17 / (((fi17 << (((v3 << (((v3 ^ (v3 / ((18446744073709551615) | 1)))) & 63))) & 63))) | 1)), fi17, -1, 127}
                    switch "foo bar" {
                    case "a":
                        switch "" {
                        case "":
                            if len(v18) > 0 {
                                v18[0] = (fi17 >> ((0) & 63))
                            }
                        }
                    }
                    v5 -= v5
                    _ = v18
                }
            }
            var v19 int8 = -128
            var v20 int32 = -862
            _ = fi17
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 int = fi21
            _ = fi21
            _ = v22
        }
    } else {
        v23 := GenS2{F0: -9223372036854775808, F1: 2}
        v24 := []int{}
        if (!(!true)) {
            v3 = 255
            for fi25 := 0; fi25 < 17; fi25++ {
                if (v3 <= v3) {
                    v26 := GenS1{F0: -185}
                    _ = v26
                }
                var v27 bool = ((true || true) && (2147483647 <= 1))
                _ = fi25
                _ = v27
            }
        } else {
            v3 += v3
            v2 = v2
        }
        _ = v23
        _ = v24
    }
    var v28 uint8 = v16
    for fi29 := 0; fi29 < 6; fi29++ {
        v5 -= ((v5 + float32(0.0)) + v5)
        for fi30 := 0; fi30 < 19; fi30++ {
            v16 -= (v16 ^ (v16 - (v16 / (((^(v16))) | 1))))
            _ = fi30
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
