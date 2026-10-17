package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint = 42
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 int64 = -9223372036854775808
    v5 := make(map[string]int)
    var v6 bool = (0.0 > float64(0.1))
    v7 := GenS1{F0: (^(v1))}
    var v8 bool = v6
    v4 = v4
    v9 := GenS1{F0: 1}
    switch 32767 {
    case -645:
        var v10 uint64 = 948
        _ = v10
    case -1:
        for fi11 := 0; fi11 < 5; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                _ = len(v2)
                if (!(0.0 <= -2.25)) {
                    switch fi12 {
                    case 32768:
                        delete(v3, 399)
                    case 127:
                        v13 := GenS2{F0: 1, F1: v1}
                        _ = v13
                    }
                    v6 = (v1 < v1)
                }
                _ = fi12
            }
            var v14 int = fi11
            var v15 int16 = 828
            _ = fi11
            _ = v14
            _ = v15
        }
    case -896:
        v6 = v8
    default:
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 string = "hello"
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        v4 -= (-(v4))
        _ = fi18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        v4 -= (v4 ^ (v4 % ((9223372036854775807) | 1)))
        if (-1 <= 127) {
            switch fi19 {
            case 128:
                switch ("a" + ("foo bar" + "a")) {
                case "a":
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 int = fi19
                        _ = fi20
                        _ = v21
                    }
                case "\n\t":
                    if v6 {
                        v22 := GenS3{F0: 42, F1: 354}
                        if true {
                            v1 -= v1
                        } else {
                            var v23 uint64 = 4294967295
                            _ = v23
                        }
                        _ = v22
                    } else {
                        var v24 uint8 = 2
                        _ = v24
                    }
                default:
                    if v6 {
                        var v25 uint64 = 135
                        v2[fi19] = 2147483647
                        _ = v25
                    }
                }
            default:
                if v6 {
                    switch fi19 {
                    case 778:
                        var v26 int8 = -116
                        _ = v26
                    case -523:
                        v8 = v6
                    }
                    v8 = v6
                }
            }
        } else {
            var v27 float64 = (1.5 * (1e10 * (-2.25 / 1.5)))
            _ = v27
        }
        _ = fi19
    }
    var v28 uint16 = 42
    if (((4294967295 > 128) || (2 != (v1 / ((uint(v28)) | 1)))) || false) {
        var v29 float64 = (0.1 / 1.5)
        var v30 uint32 = 42
        v31 := []int{1, 42, 0, int(v30)}
        _ = v29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        fmt.Println(v28, v6, v8)
        var v33 int32 = -2147483648
        fmt.Println(v1, v6, v6)
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
