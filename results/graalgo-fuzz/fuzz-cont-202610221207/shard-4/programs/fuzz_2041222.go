package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{}
    var v3 uint16 = 2
    var v4 float64 = ((-2.25 / ((1e10 - -2.25) - float64(0.1))) + float64(v1))
    var v5 uint32 = 4294967295
    v6 := []string{"", (((("" + "a") + ("hello" + "a")) + ("foo bar" + "hello")) + "a"), "foo bar"}
    v7 := GenS1{F0: 58}
    v5 += v5
    var v8 float64 = (((v4 + 1.5) + (v4 - v4)) - 0.0)
    v2 = append(v2, (v5 >> ((18446744073709551615) & 63)))
    var v9 uint = 18446744073709551615
    v9 += v9
    var v10 int64 = 42
    if (false && (!true)) {
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    } else {
        if true {
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := GenS1{F0: 1}
                fmt.Println(v8, fi11)
                v9 = v9
                _ = fi11
                _ = v12
            }
        } else {
            var v13 float32 = -2.25
            if (0 < 42) {
                switch "" {
                case "a":
                    switch 127 {
                    case -2147483648:
                        v14 := GenS2{F0: 117, F1: (v5 >> ((uint(v3)) & 63))}
                        _ = v14
                    case 178:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        v9 = (v9 % ((691) | 1))
                    }
                case "":
                    var v15 bool = ((!(32767 > 32767)) || (true && false))
                    _ = v15
                case "foo bar":
                    var v16 float64 = 1e10
                    _ = v16
                default:
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v2 = append(v2, v5)
                        _ = fi17
                    }
                }
                var v18 int32 = int32(v9)
                for fi19 := 0; fi19 < 10; fi19++ {
                    var v20 int8 = -52
                    var v21 uint64 = 596
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = v18
            }
            var v22 uint16 = (v3 & (v3 % ((v3) | 1)))
            _ = v13
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
