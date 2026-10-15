package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int8; F1 uint8 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int = 42
    var v3 int = (v2 / (((v2 ^ -330)) | 1))
    v4 := []int{(v2 - 9223372036854775807)}
    var v5 int64 = v1
    var v6 bool = (42 <= 235)
    if len(v4) > 0 {
        v4[0] = (v3 / (((v2 - (v2 - (v3 ^ -1)))) | 1))
    }
    _ = len(v4)
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint32 = 2147483648
            if (!(!(0.1 > 1e10))) {
                v6 = v6
            }
            if len(v4) > 0 {
                v4[0] = (fi7 ^ v3)
            }
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 344, F1: v1}
        _ = fi7
        _ = v10
    }
    if v6 {
        if (!v6) {
            for fi11 := 0; fi11 < 8; fi11++ {
                switch "hello" {
                case "":
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v13 := GenS2{F0: 69, F1: 185}
                        _ = fi12
                        _ = v13
                    }
                case "hello":
                    v6 = v6
                case "a":
                    var v14 int64 = 0
                    _ = v14
                default:
                    v5 = -1
                }
                _ = fi11
            }
            var v15 bool = true
            _ = v15
        } else {
            fmt.Println(v1)
        }
    } else {
        switch (v2 | 1) {
        case 0:
            var v16 uint8 = 2
            _ = v16
        case 32768:
            var v17 uint16 = 32767
            _ = v17
        default:
            fmt.Println(v5)
        }
    }
    v18 := GenS1{F0: v3, F1: v5}
    switch v2 {
    case 42:
        for fi19 := 0; fi19 < 1; fi19++ {
            if v6 {
                v5 -= 470
                v20 := GenS1{F0: (v3 * v2), F1: (v1 | v1)}
                _ = v20
            }
            if (920 > -382) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi19
        }
    default:
        fmt.Println(v5)
    }
    v21 := GenS1{F0: v2, F1: (v5 | v5)}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v22 float32 = (3.14159 + (0.0 / 0.1))
    var v23 uint = 42
    for fi24 := 0; fi24 < 10; fi24++ {
        v22 += ((3.14159 - (1e10 * (float32(3.14159) + -2.25))) - (float32((1.5 / (1e10 - 3.14159))) * v22))
        v5 += v5
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
