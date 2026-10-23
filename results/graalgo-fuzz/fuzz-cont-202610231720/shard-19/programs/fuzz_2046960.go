package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int64 }

func main() {
    var v1 uint = 608
    var v2 float32 = 1e10
    v3 := make(map[int]uint32)
    var v4 int32 = 42
    v5 := make(map[string]bool)
    var v6 uint16 = 32767
    switch "\n\t" {
    case "hello":
        v4 -= (v4 + -2147483648)
    case "\n\t":
        for fi7 := 0; fi7 < 13; fi7++ {
            if false {
                switch "hello" {
                case "":
                    _ = len(v3)
                case "\n\t":
                    var v8 float32 = v2
                    _ = v8
                }
                if true {
                    var v9 uint = (^(v1))
                    v1 = (v9 / ((v1) | 1))
                    fmt.Println(v1, v2, v9)
                    _ = v9
                }
                v2 -= (float32((3.14159 + 0.0)) * ((v2 / float32(((float64(((0.1 * (0.1 - (v2 - 0.1))) - (3.14159 / v2))) + (0.1 / (1e10 * float64(3.14159)))) / ((0.1 / ((((1.5 + 3.14159) / 1.5) / 0.1) - 0.0)) * float64(-2.25))))) / (v2 + float32(((1e10 + (0.0 + 0.1)) - ((0.1 / float64(v2)) * float64(v2)))))))
            } else {
                v10 := GenS1{F0: v1, F1: 2147483647, F2: 206}
                _ = v10
            }
            v4 -= int32(fi7)
            _ = fi7
        }
    case "a":
        v4 = (v4 + (v4 & (v4 - v4)))
    }
    if (!(true || true)) {
        if true {
            v6 += v6
            if (v2 != v2) {
                v11 := []int{}
                switch (("" + "") + "hello") {
                case "hello":
                    if false {
                        switch 875 {
                        case 56:
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        }
                        var v12 uint = v1
                        _ = v12
                    }
                case "":
                    v13 := GenS1{F0: 407, F1: 510, F2: 18446744073709551615}
                    _ = v13
                }
                _ = v11
            }
        } else {
            v14 := GenS1{F0: (v1 | v1), F1: v4, F2: 779}
            var v15 int32 = (v4 - (v4 & v4))
            _ = v14
            _ = v15
        }
        v16 := GenS1{F0: v1, F1: 42, F2: 0}
        if false {
            v6 += (v6 + (^(v6)))
            if ((!(int8(v4) > 127)) || ((v4 / ((v4) | 1)) >= v4)) {
                switch 0 {
                case -121:
                    fmt.Println(v4, v6, v1)
                default:
                    var v17 uint64 = 0
                    _ = v17
                }
                if true {
                    v18 := GenS1{F0: (v1 - v1), F1: 873, F2: 2}
                    var v19 float32 = (float32(1.5) / (float32(0.1) * ((float32((float64(v2) - 1e10)) - 1.5) * ((float32(float64(((((v2 + v2) + v2) - v2) + v2))) + v2) + ((1e10 + v2) - 0.0)))))
                    _ = v18
                    _ = v19
                } else {
                    if (255 != 242) {
                        v1 -= v1
                        var v20 uint16 = v6
                        _ = v20
                    } else {
                        fmt.Println(v6)
                        fmt.Println(v2)
                    }
                    v4 += -1
                    var v21 uint = (v1 + v1)
                    _ = v21
                }
                v6 = 42
            } else {
                v2 += v2
            }
            v2 += v2
        }
        _ = v16
    }
    if (599 != 32767) {
        fmt.Println(v4, v2, v6)
        v22 := []int{4294967295, 2147483647, 1, -263}
        v5[("" + ("foo bar" + ("hello" + ("\n\t" + (("a" + "") + "\n\t")))))] = false
        _ = v22
    } else {
        switch "foo bar" {
        case "foo bar":
            v23 := v5["\n\t"]
            _ = v23
        case "":
            v6 = v6
        case "hello":
            switch "a" {
            case "\n\t":
                for fi24 := 0; fi24 < 10; fi24++ {
                    var v25 int = fi24
                    var v26 int64 = 42
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            default:
                v6 += (v6 << (((^(v1))) & 63))
            }
        default:
            v1 += (v1 + v1)
        }
        v27 := GenS1{F0: (v1 + v1), F1: (v4 & v4), F2: 2}
        _ = v27
    }
    v2 = v2
    var v28 uint32 = 330
    fmt.Println(v2, v4)
    switch 1 {
    case 0:
        if false {
            for fi29 := 0; fi29 < 5; fi29++ {
                v4 -= -2147483648
                v30 := GenS2{F0: (v6 ^ v6), F1: 423}
                v3[fi29] = v28
                _ = fi29
                _ = v30
            }
        }
    }
    switch "foo bar" {
    case "foo bar":
        if true {
            fmt.Println(v4, v1, v28)
        }
    default:
        var v31 int32 = (v4 % (((^(v4))) | 1))
        _ = v31
    }
    var v32 bool = false
    if (65535 == v1) {
        var v33 uint16 = (v6 << ((v1) & 63))
        v2 += 1.5
        v6 += v33
        _ = v33
    }
    fmt.Println(v4, v32, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
