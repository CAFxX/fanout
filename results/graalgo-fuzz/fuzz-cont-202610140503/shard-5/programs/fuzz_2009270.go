package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 497
    var v3 uint16 = v2
    var v4 int8 = 100
    v3 += v2
    var v5 uint32 = 922
    v6 := []int{-9223372036854775808, 32767, int(v2), -9223372036854775808}
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 0}
        var v9 int8 = (^(v4))
        switch "" {
        case "":
            v5 = (v5 / ((65535) | 1))
        case "foo bar":
            _ = len(v1)
        default:
            if false {
                v5 = 813
                if false {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        var v11 int16 = 167
                        fmt.Println(fi10)
                        v12 := GenS1{F0: v11}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    var v13 int32 = -2147483648
                    _ = v13
                } else {
                    v5 = v5
                }
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v14 uint8 = 42
    v1[v5] = "foo bar"
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := v1[(v5 >> ((18446744073709551615) & 63))]
        v6 = append(v6, (fi15 % ((fi15) | 1)))
        _ = fi15
        _ = v16
    }
    var v17 uint64 = 720
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int = int(v2)
        v20 := GenS2{F0: (^(v3)), F1: (^(v4)), F2: 1}
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: (v2 << ((255) & 63)), F1: int8(v17), F2: 4294967295}
    _ = len(v6)
    var v22 uint32 = v5
    v23 := GenS3{F0: "\n\t"}
    if (!(3.14159 <= (0.1 * float64(0.1)))) {
        switch 128 {
        case -9223372036854775808:
            for fi24 := 0; fi24 < 4; fi24++ {
                for fi25 := 0; fi25 < 9; fi25++ {
                    v4 += v4
                    v26 := GenS1{F0: -744}
                    _ = fi25
                    _ = v26
                }
                fmt.Println(fi24)
                _ = fi24
            }
        }
        v4 = v4
    } else {
        v5 += (^(v22))
        fmt.Println(v22, v3)
    }
    var v27 float64 = (0.0 * 3.14159)
    fmt.Println(v27, v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
