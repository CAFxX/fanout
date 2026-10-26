package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint }

func main() {
    var v1 uint64 = 934
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]string)
    _ = len(v3)
    v4 := []int{32768, 32767, int(v1)}
    v1 -= 651
    var v5 string = "\n\t"
    if true {
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: uint16(v1)}
            v8 := GenS2{F0: v1, F1: 127}
            switch v5 {
            case "hello":
                if len(v4) > 0 {
                    v4[0] = fi6
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 int16 = -486
        _ = v9
    }
    var v10 int16 = 32767
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 uint = 788
            switch (v5 + v5) {
            case "foo bar":
                v4 = append(v4, 32767)
            case "a":
                v3[0] = ((v5 + v5) + v5)
            case "hello":
                var v14 int8 = -72
                _ = v14
            }
            fmt.Println(v1, v13)
            _ = fi12
            _ = v13
        }
        if false {
            switch v5 {
            case "\n\t":
                v15 := GenS3{F0: 65535, F1: 128, F2: 18446744073709551615}
                _ = v15
            case "hello":
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint16 = 32767
                    _ = fi16
                    _ = v17
                }
            case "foo bar":
                switch fi11 {
                case -228:
                    for fi18 := 0; fi18 < 10; fi18++ {
                        fmt.Println(v5)
                        fmt.Println(fi11)
                        fmt.Println(fi18, fi18)
                        v19 := v3[1]
                        _ = fi18
                        _ = v19
                    }
                case -1:
                    v1 -= (^(v1))
                case 4294967295:
                    fmt.Println(v10, v5)
                default:
                    var v20 uint16 = 0
                    _ = v20
                }
            }
            switch 32767 {
            case 32767:
                v1 = (^(v1))
            case 1:
                v21 := GenS1{F0: 255}
                _ = v21
            default:
                v10 += 68
            }
        }
        _ = fi11
    }
    var v22 int64 = v2
    v3[626] = v5
    var v23 int8 = 79
    var v24 int8 = int8(v22)
    if false {
        for fi25 := 0; fi25 < 6; fi25++ {
            v5 = (v5 + v5)
            v5 = ""
            var v26 int64 = 893
            _ = fi25
            _ = v26
        }
        var v27 float32 = ((-2.25 + (((1e10 + (1e10 / ((1e10 / 3.14159) + -2.25))) * 1.5) / 0.1)) + -2.25)
        _ = v27
    }
    v3[2147483648] = "a"
    v3[1] = "hello"
    var v28 bool = (v2 != v22)
    var v29 int8 = -49
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
