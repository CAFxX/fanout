package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint32 = 90
    var v2 bool = true
    var v3 int = -9223372036854775808
    v4 := make(map[string]string)
    var v5 float32 = 3.14159
    v6 := GenS1{F0: 1, F1: (v3 - v3), F2: 128}
    if v2 {
        v5 += (v5 / float32((1.5 * 1e10)))
        var v7 bool = v2
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 string = ("\n\t" + "")
            if ((v9 + "a") >= v9) {
                if v7 {
                    v10 := GenS1{F0: 2147483647, F1: 708, F2: 1}
                    v3 = (-(v3))
                    _ = v10
                }
                v7 = (!v7)
            } else {
                switch (fi8 / ((1) | 1)) {
                case -1:
                    var v11 bool = (315 == 42)
                    _ = v11
                case -977:
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v9 = v9
                        _ = fi12
                    }
                }
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        fmt.Println(v2)
        v5 -= (v5 * float32(1e10))
    }
    var v13 float64 = (1e10 + 1.5)
    v1 += 2147483648
    v13 = (0.1 * (v13 / float64(v5)))
    v14 := []int{(v3 * (v3 / (((v3 & v3)) | 1))), (v3 >> ((18446744073709551615) & 63)), 0}
    v14 = append(v14, (v3 << ((255) & 63)))
    if (128 <= 42) {
        fmt.Println(v13, v1)
        fmt.Println(v5)
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 int32 = 42
            v17 := GenS2{F0: ((v16 / (((v16 >> ((42) & 63))) | 1)) <= v16)}
            var v18 uint16 = 2
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    }
    var v19 int16 = 42
    if (!v2) {
        v19 = (^(v19))
        var v20 float64 = float64(0.0)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 9; fi21++ {
            v22 := GenS1{F0: 42, F1: fi21, F2: 1}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            for fi23 := 0; fi23 < 6; fi23++ {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                fmt.Println(v13)
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
        var v24 int8 = 0
        _ = v24
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        switch "a" {
        case "hello":
            v5 = (((((v5 / (-2.25 + v5)) - v5) - 3.14159) * float32(v13)) - v5)
        case "":
            v5 += v5
        case "foo bar":
            v2 = true
        default:
            v14 = append(v14, (-(v3)))
        }
        _ = fi25
    }
    if v2 {
        v13 -= float64(0.0)
    }
    var v26 uint8 = uint8(v1)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
