package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint8; F1 int8; F2 uint }
type GenS5 struct { F0 string; F1 uint64; F2 uint }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 bool; F1 uint64; F2 int8 }
func genFn0(p0 int16) int {
    var v4 int16 = p0
    v4 = p0
    v5 := GenS1{F0: 420, F1: 201, F2: ("a" + ("" + "foo bar"))}
    switch 999 {
    case 791:
        var v6 uint16 = 0
        _ = v6
    case -2147483648:
        if true {
            p0 += 155
        } else {
            var v7 int32 = -1
            _ = v7
        }
    case 1:
        v8 := GenS1{F0: 254, F1: 2147483647, F2: "foo bar"}
        _ = v8
    default:
        var v9 string = "foo bar"
        _ = v9
    }
    var v10 uint8 = 42
    _ = p0
    _ = v10
    _ = v4
    _ = v5
    return 128
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 65535
    var v3 float32 = 0.0
    fmt.Println(v3, v2)
    var v11 string = "\n\t"
    for fi12 := 0; fi12 < 7; fi12++ {
        _ = len(v1)
        v13 := GenS1{F0: 1, F1: (v2 & (v2 | 1)), F2: "foo bar"}
        if true {
            v14 := GenS2{F0: false}
            v15 := GenS3{F0: 255, F1: false}
            _ = v14
            _ = v15
        } else {
            v16 := []int{fi12, (-(fi12)), fi12}
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    v2 = 849
    if ((!false) && true) {
        var v17 uint64 = 42
        v18 := GenS4{F0: 0, F1: 0, F2: 2}
        v19 := GenS2{F0: false}
        _ = v17
        _ = v18
        _ = v19
    }
    switch 1 {
    case 4294967295:
        fmt.Println(v11, v2, v2)
    }
    fmt.Println(v3)
    v2 -= (^(v2))
    v3 -= v3
    if true {
        if ((472 == 42) || false) {
            var v20 int8 = 100
            _ = v20
        } else {
            v21 := v1[((v11 + "foo bar") + "\n\t")]
            _ = v21
        }
        v11 = ""
    }
    if false {
        v11 = ""
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        switch 128 {
        case 0:
            v23 := GenS5{F0: v11, F1: 240, F2: 18446744073709551615}
            _ = v23
        case 9223372036854775807:
            var v24 int64 = -9223372036854775808
            _ = v24
        }
        v3 += (((v3 + 3.14159) / v3) - (1e10 + 0.1))
        switch (-(fi22)) {
        case 4294967295:
            v25 := GenS4{F0: 128, F1: -117, F2: 65535}
            _ = v25
        case -2147483648:
            switch fi22 {
            case -220:
                switch v11 {
                case "":
                    if (0 != 1) {
                        v26 := []int{(fi22 >> ((719) & 63))}
                        if (false && false) {
                            var v27 float64 = (1.5 / 3.14159)
                            if len(v26) > 0 {
                                v26[0] = 127
                            }
                            v28 := GenS6{F0: 4294967295, F1: 367}
                            _ = v27
                            _ = v28
                        }
                        _ = v26
                    }
                case "hello":
                    var v29 uint32 = v2
                    _ = v29
                case "foo bar":
                    fmt.Println(fi22, v2, v2)
                default:
                    v30 := GenS7{F0: (-2.25 <= -2.25), F1: 0, F2: 127}
                    _ = v30
                }
            default:
                fmt.Println(v11, fi22)
            }
        default:
            var v31 uint = 846
            _ = v31
        }
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
