package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint }

func main() {
    var v1 float32 = (float32((float64(3.14159) - (((((((3.14159 / 1.5) + float64((0.0 / 1e10))) / 0.1) + 0.1) + ((0.0 - (-2.25 + 0.0)) / (0.1 + 0.0))) / 3.14159) + 3.14159))) - ((((3.14159 - 3.14159) + (3.14159 / 3.14159)) - ((float32(-2.25) - 1e10) * float32(float64((0.1 * (0.0 / 1e10)))))) - 0.0))
    var v2 uint8 = 2
    var v3 int = 9223372036854775807
    v4 := []int{-447}
    v5 := GenS1{F0: false, F1: 1, F2: 680}
    if (1 <= 2) {
        switch ("" + (("foo bar" + "") + ("foo bar" + "\n\t"))) {
        case "hello":
            for fi6 := 0; fi6 < 16; fi6++ {
                var v7 uint32 = 725
                _ = fi6
                _ = v7
            }
        default:
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := GenS2{F0: 42, F1: 1}
                v10 := GenS3{F0: -50, F1: 255}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        fmt.Println(v2, v1, v3)
        if true {
            for fi11 := 0; fi11 < 15; fi11++ {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v1 += (v1 + v1)
                    switch 128 {
                    case 95:
                        if false {
                            switch fi12 {
                            case 128:
                                v4 = append(v4, fi11)
                            default:
                                var v13 int16 = -148
                                _ = v13
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    default:
                        v14 := GenS4{F0: 859}
                        _ = v14
                    }
                    _ = fi12
                }
                if len(v4) > 0 {
                    v4[0] = -842
                }
                _ = fi11
            }
            v1 += -2.25
        } else {
            v15 := GenS1{F0: ((!(0 != 100)) || ((!(42 == 405)) && true)), F1: -747, F2: 922}
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 uint32 = 572
        v1 = v1
        var v18 uint64 = 352
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        v1 -= 1e10
        switch 127 {
        case 2147483647:
            v2 = (v2 * (v2 ^ v2))
        default:
            v4 = append(v4, fi19)
        }
        if len(v4) > 0 {
            v4[0] = v3
        }
        _ = fi19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        fmt.Println(v3, v1)
        if (!(0 < 42)) {
            var v21 uint8 = v2
            _ = v21
        } else {
            if ((!false) || (!true)) {
                fmt.Println(v2)
                v3 -= 127
                var v22 uint8 = 1
                _ = v22
            }
            switch -56 {
            case -468:
                if ((100 < 42) && true) {
                    v1 = v1
                    fmt.Println(v2, fi20)
                }
            default:
                for fi23 := 0; fi23 < 1; fi23++ {
                    if (-1 < -982) {
                        v24 := GenS5{F0: 42, F1: 42}
                        _ = v24
                    }
                    v25 := GenS4{F0: 429}
                    _ = fi23
                    _ = v25
                }
            }
        }
        _ = fi20
    }
    var v26 int = v3
    if ((true && false) && true) {
        var v27 bool = ((((-2.25 + 1.5) + 0.1) + 0.0) != 0.1)
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
