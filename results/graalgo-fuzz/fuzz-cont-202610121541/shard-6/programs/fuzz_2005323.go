package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int32; F1 bool }

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    delete(v3, 845)
    fmt.Println(v1)
    v1 -= (v1 + (v1 % ((v1) | 1)))
    v1 = (v1 >> ((915) & 63))
    v4 := GenS1{F0: -32768}
    var v5 float32 = ((3.14159 - ((-2.25 - float32(-2.25)) - 1e10)) * float32((0.0 * 1e10)))
    var v6 int64 = -9223372036854775808
    if true {
        v1 = v1
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(v6)
        if (-1 > 0) {
            switch ("hello" + ("a" + "foo bar")) {
            case "\n\t":
                fmt.Println(v1)
            case "a":
                for fi8 := 0; fi8 < 9; fi8++ {
                    v9 := GenS1{F0: -32768}
                    _ = fi8
                    _ = v9
                }
            case "hello":
                fmt.Println(v6)
            }
            for fi10 := 0; fi10 < 13; fi10++ {
                fmt.Println(v1, v6)
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v1, v6, v6)
                case "":
                    v11 := GenS2{F0: true, F1: (v6 >> ((4294967295) & 63))}
                    var v12 float64 = ((-2.25 * 1e10) + 0.0)
                    v1 = (v1 - (v1 >> ((2) & 63)))
                    _ = v11
                    _ = v12
                default:
                    switch fi7 {
                    case -2147483648:
                        v13 := GenS3{F0: int32(v6), F1: (!(false && false))}
                        _ = v13
                    default:
                        switch "" {
                        case "":
                            if (("" + "") >= "hello") {
                                v6 -= v6
                            }
                        case "foo bar":
                            v14 := v3[503]
                            _ = v14
                        default:
                            fmt.Println(v1)
                        }
                        v5 += v5
                        var v15 uint8 = 151
                        _ = v15
                    }
                }
                if (1.5 <= 1.5) {
                    v5 = 1e10
                }
                _ = fi10
            }
            v2[2147483647] = "a"
        } else {
            if (false && false) {
                v5 = 0.0
            } else {
                v6 += 42
                var v16 int16 = 1
                v6 += v6
                _ = v16
            }
            var v17 int64 = -765
            _ = v17
        }
        v3[0] = 2147483647
        _ = fi7
    }
    fmt.Println(v6)
    var v18 float32 = (v5 / (float32(float64((-2.25 * v5))) - (v5 * (v5 * v5))))
    fmt.Println(v6, v6, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
