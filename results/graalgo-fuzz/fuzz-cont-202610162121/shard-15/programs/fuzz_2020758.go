package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]string)
    v3 := []int{127, 9223372036854775807}
    var v4 float32 = 1.5
    var v5 int16 = 129
    var v6 uint = 255
    switch 32768 {
    case 4294967295:
        var v7 uint16 = 714
        _ = v7
    case -1:
        v6 += v6
    case -776:
        v5 = (-(v5))
    default:
        if (!false) {
            var v8 int16 = -528
            _ = v8
        }
    }
    v9 := GenS1{F0: 42, F1: (-9223372036854775808 > -815)}
    switch (("\n\t" + "hello") + "hello") {
    case "foo bar":
        if true {
            var v10 uint32 = 0
            if (((v1 - 0.0) - (v1 + v1)) <= v1) {
                switch -2147483648 {
                case 32768:
                    if false {
                        var v11 bool = (-65 > 21)
                        if len(v3) > 0 {
                            v3[0] = -646
                        }
                        _ = v11
                    }
                case -342:
                    _ = len(v2)
                default:
                    if (127 != 0) {
                        var v12 float64 = ((((3.14159 + v1) - float64((v4 + v4))) - v1) - float64((1e10 / v4)))
                        switch 1 {
                        case 127:
                            switch 0 {
                            case 505:
                                v13 := GenS1{F0: 127, F1: ("" == "\n\t")}
                                _ = v13
                            case 32768:
                                fmt.Println(v4, v12, v10)
                            case 9223372036854775807:
                                fmt.Println(v1)
                            }
                        case 2147483647:
                            var v14 uint8 = 145
                            _ = v14
                        default:
                            var v15 string = (("a" + "foo bar") + ("foo bar" + ""))
                            _ = v15
                        }
                        v6 -= (v6 - 2)
                        _ = v12
                    }
                }
                var v16 uint = (^(v6))
                _ = v16
            } else {
                v3 = append(v3, -9223372036854775808)
            }
            v4 -= (v4 / 1.5)
            _ = v10
        }
    case "":
        v17 := GenS2{F0: 1}
        _ = v17
    case "a":
        fmt.Println(v5, v4, v1)
    }
    v18 := GenS2{F0: 535}
    v4 = -2.25
    fmt.Println(v1)
    v19 := GenS1{F0: 2, F1: (!true)}
    var v20 float64 = (v1 - ((v1 / (float64(float32(v1)) - (v1 / ((0.0 - (float64(1.5) + ((v1 - float64(float32(float64((((v4 * ((v4 / v4) / -2.25)) * 1e10) + (((v4 / float32(v1)) - ((v4 / 1e10) + v4)) / (float32((v1 + v1)) - (v4 * v4)))))))) * (v1 + (float64(v4) + -2.25))))) / v1)))) - v1))
    fmt.Println(v20)
    v20 = float64(v4)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
