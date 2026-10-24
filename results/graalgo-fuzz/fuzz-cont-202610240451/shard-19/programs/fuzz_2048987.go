package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
type GenS2 struct { F0 bool }

func main() {
    var v1 int64 = 1
    var v2 bool = (2 != 842)
    v3 := make(map[int]string)
    v4 := []int{128, 128}
    var v5 uint32 = 42
    var v6 float64 = (1e10 - -2.25)
    fmt.Println(v5, v6)
    if v2 {
        delete(v3, -438)
    } else {
        _ = len(v3)
        for fi7 := 0; fi7 < 14; fi7++ {
            v4 = append(v4, -1)
            var v8 uint32 = (v5 / ((v5) | 1))
            v8 -= (v5 ^ v8)
            _ = fi7
            _ = v8
        }
        switch -1 {
        case 9223372036854775807:
            var v9 string = ("a" + "")
            _ = v9
        case 127:
            if v2 {
                if (100 <= 30) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch "foo bar" {
                    case "hello":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case "":
                        v6 = float64((0.1 - (float32(v6) - (3.14159 * -2.25))))
                    default:
                        v1 = (v1 >> ((19) & 63))
                    }
                } else {
                    var v10 float32 = 0.0
                    fmt.Println(v5, v2)
                    _ = v10
                }
                var v11 float64 = (1e10 / ((v6 * (float64(0.1) / 1e10)) + 1e10))
                var v12 uint64 = 4294967295
                _ = v11
                _ = v12
            } else {
                if (!(!((((0 >= -1) && (-1 == (v1 / (((v1 >> ((2) & 63))) | 1)))) && ((v5 << ((42) & 63)) < v5)) && v2))) {
                    var v13 int16 = 32767
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        v15 := GenS1{F0: uint8(v5), F1: 32767, F2: 65535}
                        _ = fi14
                        _ = v15
                    }
                }
                if len(v4) > 0 {
                    v4[0] = -1
                }
            }
        }
    }
    v16 := GenS2{F0: v2}
    if (-128 > -1) {
        delete(v3, 9223372036854775807)
    }
    v5 = v5
    fmt.Println(v1, v1)
    fmt.Println(v1)
    v17 := v3[9223372036854775807]
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
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
