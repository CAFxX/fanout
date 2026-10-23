package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int8; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int32 }

func main() {
    var v1 float32 = (0.1 * 3.14159)
    var v2 float64 = 0.0
    var v3 float64 = 1e10
    v4 := []int{336}
    v1 -= 1.5
    var v5 float64 = v3
    var v6 int64 = 276
    for fi7 := 0; fi7 < 2; fi7++ {
        v4 = append(v4, (fi7 << ((65535) & 63)))
        _ = fi7
    }
    var v8 uint = 42
    fmt.Println(v5, v5)
    if (0 != 0) {
        for fi9 := 0; fi9 < 18; fi9++ {
            switch "" {
            case "":
                v3 = ((((v5 / (v2 / 0.1)) * 1e10) / (v2 * v2)) + v5)
            case "hello":
                for fi10 := 0; fi10 < 2; fi10++ {
                    if (0 != 35) {
                        v4 = append(v4, 4294967295)
                        _ = len(v4)
                        v11 := GenS1{F0: -77, F1: 0}
                        _ = v11
                    } else {
                        fmt.Println(v8, v5, v5)
                        fmt.Println(v6)
                    }
                    if true {
                        switch (^(fi9)) {
                        case 777:
                            v12 := GenS1{F0: 81, F1: 554}
                            _ = v12
                        case 240:
                            v13 := GenS2{F0: 1, F1: 569, F2: 42}
                            _ = v13
                        }
                        v1 += float32((((float64(v1) + ((float64(v1) - v2) - 1e10)) / -2.25) - (v5 + -2.25)))
                        v1 -= v1
                    } else {
                        v4 = append(v4, 32768)
                        fmt.Println(v2, v6, v3)
                        if (18446744073709551615 <= 846) {
                            fmt.Println(v8, v3, v1)
                            var v14 uint64 = 65
                            if (!false) {
                                var v15 int = (fi10 >> ((v8) & 63))
                                _ = v15
                            }
                            _ = v14
                        } else {
                            var v16 uint16 = 42
                            _ = v16
                        }
                    }
                    v17 := GenS3{F0: 18}
                    _ = fi10
                    _ = v17
                }
            case "a":
                v4 = append(v4, (fi9 - fi9))
            default:
                switch fi9 {
                case -753:
                    fmt.Println(v5, v1, v1)
                case 32768:
                    var v18 int8 = -6
                    _ = v18
                case -2147483648:
                    v2 += (3.14159 - ((float64((float32(3.14159) + 3.14159)) * v3) * v2))
                }
            }
            if true {
                if false {
                    v8 = (v8 & 1)
                }
                fmt.Println(v1)
                var v19 int16 = 892
                _ = v19
            }
            v20 := GenS3{F0: 1}
            _ = fi9
            _ = v20
        }
        v21 := GenS4{F0: "\n\t", F1: 1}
        if ((938 != 515) || (false && (685 < 42))) {
            v2 = v5
            v22 := GenS3{F0: 843}
            switch "\n\t" {
            case "":
                v4 = append(v4, -9223372036854775808)
            default:
                fmt.Println(v6, v6, v8)
            }
            _ = v22
        }
        _ = v21
    } else {
        v2 = float64(v1)
        if len(v4) > 0 {
            v4[0] = -9223372036854775808
        }
        v2 += v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
