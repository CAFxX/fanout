package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = ((-2.25 * ((float64(-2.25) / (0.0 - ((-2.25 + 0.1) * 3.14159))) - 1.5)) * (0.1 / (((-2.25 / 1e10) + 0.1) - (1e10 / 1.0))))
    var v3 int32 = -2147483648
    v4 := make(map[int]bool)
    var v5 uint32 = 65535
    var v6 int = 128
    if (!true) {
        for fi7 := 0; fi7 < 9; fi7++ {
            switch "" {
            case "foo bar":
                var v8 uint16 = 0
                _ = v8
            case "hello":
                var v9 uint32 = 2147483647
                _ = v9
            case "":
                if false {
                    v2 = ((((1e10 + 1.5) - 0.0) - 3.14159) - -2.25)
                    var v10 int = fi7
                    v11 := GenS1{F0: true, F1: (v3 >> ((432) & 63)), F2: 771}
                    _ = v10
                    _ = v11
                }
            default:
                v12 := GenS1{F0: true, F1: v3, F2: 2}
                _ = v12
            }
            var v13 bool = false
            _ = fi7
            _ = v13
        }
        v14 := GenS1{F0: false, F1: v3, F2: 42}
        switch "" {
        case "\n\t":
            fmt.Println(v2, v2)
        default:
            var v15 int = v6
            _ = v15
        }
        _ = v14
    }
    var v16 float64 = v2
    v2 = float64((0.0 * ((1.5 - -2.25) / 1.5)))
    var v17 float32 = 1e10
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS2{F0: (-2.25 == (1e10 + v17))}
        if (4294967295 <= 689) {
            for fi20 := 0; fi20 < 15; fi20++ {
                if false {
                    if false {
                        v16 = (float64(3.14159) + 0.1)
                        var v21 uint = 686
                        if (-1 != -9223372036854775808) {
                            v17 += 1.5
                            fmt.Println(fi18, v6)
                            var v22 bool = true
                            _ = v22
                        }
                        _ = v21
                    } else {
                        fmt.Println(v2, v3, v16)
                        v23 := GenS1{F0: true, F1: -396, F2: 2}
                        _ = v23
                    }
                }
                _ = fi20
            }
            v5 += v5
            v24 := GenS2{F0: (!false)}
            _ = v24
        } else {
            v25 := v4[(^(v6))]
            if v25 {
                var v26 uint16 = 994
                _ = v26
            } else {
                v25 = (v25 || v25)
                v16 = (v16 + 0.0)
                v5 = v5
            }
            v6 += -9223372036854775808
            _ = v25
        }
        _ = fi18
        _ = v19
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := GenS3{F0: 163}
        _ = fi27
        _ = v28
    }
    if false {
        var v29 int = 0
        if (18446744073709551615 == 788) {
            v30 := GenS1{F0: (248 > 209), F1: v3, F2: 2}
            var v31 int16 = 1
            _ = v30
            _ = v31
        } else {
            v32 := v1["\n\t"]
            v33 := GenS1{F0: (false && false), F1: (v3 / ((v3) | 1)), F2: 18446744073709551615}
            var v34 uint16 = 2
            _ = v32
            _ = v33
            _ = v34
        }
        var v35 uint16 = 65535
        _ = v29
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
