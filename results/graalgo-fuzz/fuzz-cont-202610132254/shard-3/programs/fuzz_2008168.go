package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 float64 = (((0.1 / ((1e10 * -2.25) - 0.0)) / 1.5) / float64(3.14159))
    v4 := make(map[int]uint32)
    v3 = float64(float32(1.5))
    _ = len(v1)
    v5 := v2["foo bar"]
    fmt.Println(v5)
    v3 = (((((v3 - (1.5 + float64((((3.14159 / 1.5) + 1.5) / 0.1)))) / v3) - float64(1e10)) + v3) / 1.5)
    v6 := []int{(v5 >> ((255) & 63)), (v5 - 128), (v5 & v5), v5}
    var v7 uint16 = 32767
    v5 = 24
    if (0 <= 9223372036854775807) {
        fmt.Println(v3)
        for fi8 := 0; fi8 < 9; fi8++ {
            switch "foo bar" {
            case "foo bar":
                if (-614 >= 471) {
                    fmt.Println(v7)
                } else {
                    v6 = append(v6, (^(v5)))
                    for fi9 := 0; fi9 < 4; fi9++ {
                        fmt.Println(fi9)
                        _ = len(v4)
                        if len(v6) > 0 {
                            v6[0] = fi8
                        }
                        _ = fi9
                    }
                }
            default:
                v10 := GenS1{F0: 42, F1: ("foo bar" + "")}
                _ = v10
            }
            switch fi8 {
            case -9223372036854775808:
                v7 -= v7
            case -937:
                for fi11 := 0; fi11 < 7; fi11++ {
                    fmt.Println(v5)
                    _ = fi11
                }
            }
            _ = fi8
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            v7 -= (v7 ^ 65535)
            _ = fi12
        }
    } else {
        v5 += v5
    }
    if (!(false && true)) {
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := GenS1{F0: 1, F1: "a"}
            for fi15 := 0; fi15 < 3; fi15++ {
                _ = len(v6)
                var v16 uint64 = 621
                var v17 bool = (65535 == 2147483647)
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi13
            _ = v14
        }
        var v18 string = ((("foo bar" + "") + "a") + "\n\t")
        _ = v18
    } else {
        if (false && true) {
            for fi19 := 0; fi19 < 6; fi19++ {
                switch "a" {
                case "hello":
                    var v20 uint = 4294967295
                    _ = v20
                default:
                    switch "" {
                    case "hello":
                        for fi21 := 0; fi21 < 2; fi21++ {
                            v22 := GenS1{F0: 4294967295, F1: ""}
                            var v23 int8 = 42
                            var v24 bool = (-9223372036854775808 < 1)
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    default:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                }
                _ = fi19
            }
        } else {
            fmt.Println(v3)
        }
        fmt.Println(v5, v7, v7)
    }
    v7 = v7
    v7 -= 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
