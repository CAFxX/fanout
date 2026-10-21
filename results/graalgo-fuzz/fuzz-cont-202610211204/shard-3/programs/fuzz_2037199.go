package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 int8 }

func main() {
    var v1 int8 = -128
    var v2 int64 = 949
    var v3 uint16 = 2
    var v4 bool = (1.5 < 3.14159)
    var v5 uint64 = 0
    v6 := make(map[uint32]bool)
    for fi7 := 0; fi7 < 16; fi7++ {
        switch fi7 {
        case -590:
            var v8 int8 = 127
            _ = v8
        case 266:
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 int = fi7
                var v11 int32 = 2147483647
                v12 := v6[65535]
                v13 := GenS1{F0: -725, F1: v5, F2: 1000}
                v3 = v3
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
                _ = v13
            }
        case -9223372036854775808:
            v14 := GenS2{F0: ""}
            _ = v14
        default:
            v15 := GenS3{F0: 0}
            _ = v15
        }
        _ = fi7
    }
    var v16 uint16 = v3
    v17 := v6[2147483648]
    for fi18 := 0; fi18 < 14; fi18++ {
        switch fi18 {
        case 121:
            for fi19 := 0; fi19 < 5; fi19++ {
                switch "foo bar" {
                case "\n\t":
                    v20 := []int{128, (fi18 << ((457) & 63)), fi19}
                    _ = v20
                case "a":
                    fmt.Println(v1)
                    v21 := GenS4{F0: 42}
                    var v22 uint32 = 42
                    if (!true) {
                        v2 -= v2
                    } else {
                        v5 = 914
                        var v23 int16 = -988
                        v5 -= (^(v5))
                        _ = v23
                    }
                    _ = v21
                    _ = v22
                case "foo bar":
                    v24 := GenS5{F0: 127}
                    _ = v24
                default:
                    v25 := GenS5{F0: 128}
                    _ = v25
                }
                _ = fi19
            }
        case 389:
            v26 := GenS2{F0: "a"}
            _ = v26
        case 1:
            v16 = (v16 % (((v16 ^ v3)) | 1))
        default:
            var v27 uint = 1
            _ = v27
        }
        v16 -= v16
        v28 := []int{(fi18 % (((fi18 * 9223372036854775807)) | 1)), (fi18 ^ (fi18 | -1))}
        _ = fi18
        _ = v28
    }
    var v29 int64 = (v2 | (v2 + (^(v2))))
    for fi30 := 0; fi30 < 1; fi30++ {
        v17 = true
        _ = fi30
    }
    var v31 uint32 = 2
    v17 = (!false)
    fmt.Println(v3, v16)
    fmt.Println(v29)
    var v32 float64 = -2.25
    var v33 int16 = -100
    fmt.Println(v16, v29)
    delete(v6, (v31 % ((0) | 1)))
    v34 := GenS6{F0: 18446744073709551615, F1: -128}
    fmt.Println(v16, v16, v31)
    fmt.Println(v1, v3)
    switch 32768 {
    case 0:
        v3 = v16
    default:
        var v35 int = 42
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
