package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint8 = 125
    v2 := make(map[int]int)
    var v3 uint = 0
    var v4 float32 = 1e10
    v5 := make(map[int]int)
    v3 = 18446744073709551615
    var v6 uint32 = 2147483648
    v7 := []int{32767, 32768, 32768}
    v3 += v3
    switch -9223372036854775808 {
    case 32767:
        switch 0 {
        case -265:
            var v8 int = -729
            _ = v8
        default:
            v9 := GenS1{F0: 1, F1: 1}
            _ = v9
        }
    default:
        switch 169 {
        case -239:
            if true {
                switch 128 {
                case 30:
                    if false {
                        fmt.Println(v6)
                        for fi10 := 0; fi10 < 13; fi10++ {
                            v5[(fi10 * fi10)] = -1
                            _ = fi10
                        }
                    }
                }
                delete(v2, -2147483648)
                if false {
                    if len(v7) > 0 {
                        v7[0] = 32768
                    }
                    _ = len(v7)
                }
            } else {
                var v11 int64 = 9223372036854775807
                var v12 float32 = ((-2.25 / v4) / v4)
                if (!false) {
                    var v13 uint = 18446744073709551615
                    switch 2147483647 {
                    case 700:
                        fmt.Println(v3, v4, v13)
                    default:
                        var v14 string = ("\n\t" + "foo bar")
                        _ = v14
                    }
                    _ = v13
                } else {
                    var v15 string = ""
                    v12 -= -2.25
                    _ = len(v7)
                    _ = v15
                }
                _ = v11
                _ = v12
            }
        default:
            if true {
                _ = len(v7)
                fmt.Println(v1, v3, v4)
                v16 := GenS2{F0: 0}
                _ = v16
            } else {
                v6 += v6
                v1 += (^(v1))
                for fi17 := 0; fi17 < 6; fi17++ {
                    v6 += v6
                    v1 -= (v1 & v1)
                    for fi18 := 0; fi18 < 17; fi18++ {
                        var v19 float64 = float64((float32(0.1) + (v4 / (((((float32(((-2.25 + 0.0) * 0.0)) - ((1e10 - ((1e10 + v4) + (0.0 / (-2.25 * v4)))) * ((0.1 * 1.5) - v4))) / 1e10) / v4) / v4) + v4))))
                        v4 += (v4 * v4)
                        _ = len(v7)
                        _ = fi18
                        _ = v19
                    }
                    _ = fi17
                }
            }
        }
    }
    v20 := GenS2{F0: -102}
    v21 := GenS1{F0: 18446744073709551615, F1: 2}
    var v22 float32 = 0.1
    var v23 int64 = 967
    v1 = (v1 % (((v1 & uint8(v3))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
