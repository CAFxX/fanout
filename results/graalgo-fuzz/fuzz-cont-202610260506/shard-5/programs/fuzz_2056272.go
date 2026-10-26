package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 1
    var v3 uint8 = 2
    var v4 uint16 = 65535
    fmt.Println(v4)
    if ((v4 - (v4 % ((65535) | 1))) != 0) {
        for fi5 := 0; fi5 < 14; fi5++ {
            v6 := GenS1{F0: 0}
            _ = fi5
            _ = v6
        }
    }
    v7 := GenS1{F0: 18446744073709551615}
    switch "a" {
    case "foo bar":
        if true {
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 int = fi8
                _ = fi8
                _ = v9
            }
            v10 := GenS2{F0: (!(0 <= -128)), F1: -105}
            _ = v10
        }
    default:
        v11 := []int{-9223372036854775808}
        _ = v11
    }
    _ = len(v1)
    v4 = (v4 - (^(v4)))
    if true {
        delete(v1, 128)
        v4 = (v4 % ((v4) | 1))
        _ = len(v1)
    } else {
        var v12 int8 = -81
        for fi13 := 0; fi13 < 19; fi13++ {
            v4 = (v4 / ((v4) | 1))
            v14 := []int{(fi13 + 32767), (fi13 % ((int(v3)) | 1)), (fi13 | 9223372036854775807)}
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            if (!false) {
                if true {
                    fmt.Println(v3, v2)
                    if true {
                        v16 := GenS1{F0: 18446744073709551615}
                        _ = v16
                    }
                    var v17 uint = 18446744073709551615
                    _ = v17
                } else {
                    var v18 uint8 = v3
                    switch 128 {
                    case 127:
                        var v19 uint16 = v4
                        _ = v19
                    default:
                        var v20 float64 = (((0.0 * (float64(0.1) * (((float64(0.0) + (((0.1 + 0.0) + (1.5 * (1e10 + 0.1))) - ((-2.25 - (-2.25 * (1.5 * 0.0))) * 0.1))) / (1.5 + float64(float32(1.5)))) / (1e10 + (0.1 / 1.5))))) * 1.5) - 0.1)
                        _ = v20
                    }
                    _ = v18
                }
                v3 += (v3 ^ (v3 + (v3 % ((2) | 1))))
            } else {
                v1[(fi15 / ((fi15) | 1))] = fi15
            }
            var v21 uint = 1
            _ = fi15
            _ = v21
        }
        _ = v12
    }
    _ = len(v1)
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
