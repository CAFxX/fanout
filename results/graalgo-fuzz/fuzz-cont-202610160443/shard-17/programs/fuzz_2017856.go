package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 bool; F2 string }

func main() {
    var v1 int = -1
    v2 := []int{-89, (v1 - (v1 ^ (v1 | 32767)))}
    v3 := make(map[string]string)
    var v4 int16 = 1000
    var v5 uint8 = 255
    if false {
        v1 -= (v1 >> ((18446744073709551615) & 63))
    } else {
        v5 -= v5
    }
    var v6 int = (-(v1))
    fmt.Println(v4, v1)
    var v7 string = ""
    var v8 int32 = 0
    fmt.Println(v4)
    switch v7 {
    case "hello":
        switch v1 {
        case -2147483648:
            switch 577 {
            case 32768:
                var v9 int = (v1 & (v6 / ((-1) | 1)))
                _ = v9
            case 668:
                if (("" + ("" + ((((v7 + v7) + v7) + v7) + ((v7 + "foo bar") + "foo bar")))) != v7) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        v10 := GenS1{F0: (true || false)}
                        _ = v10
                    case "foo bar":
                        if (4294967295 <= 886) {
                            v8 += v8
                            for fi11 := 0; fi11 < 17; fi11++ {
                                v5 = v5
                                if (-9223372036854775808 <= -1) {
                                    if ((3.14159 <= 0.1) || (!(v8 != v8))) {
                                        switch v7 {
                                        case "hello":
                                            delete(v3, v7)
                                        }
                                        v12 := GenS1{F0: (false || true)}
                                        fmt.Println(v4, v7)
                                        _ = v12
                                    }
                                    var v13 uint = 0
                                    _ = v13
                                }
                                if ((3.14159 * ((float64(float32(3.14159)) * 3.14159) - -2.25)) >= 1.5) {
                                    for fi14 := 0; fi14 < 4; fi14++ {
                                        var v15 bool = (!true)
                                        _ = fi14
                                        _ = v15
                                    }
                                }
                                _ = fi11
                            }
                            fmt.Println(v4, v4, v8)
                        }
                    default:
                        v5 += (v5 * v5)
                    }
                    var v16 int = 127
                    _ = v16
                }
            }
        case 9223372036854775807:
            v6 = v1
        default:
            v1 = v6
        }
    default:
        var v17 float64 = 1.5
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        switch v1 {
        case 9223372036854775807:
            var v19 float32 = float32(float64((3.14159 + -2.25)))
            _ = v19
        case 128:
            v20 := GenS2{F0: 9223372036854775807, F1: true, F2: v7}
            _ = v20
        case 2147483647:
            var v21 float64 = 0.1
            _ = v21
        }
        if false {
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := GenS2{F0: 0, F1: (!false), F2: v7}
                v8 -= v8
                _ = fi22
                _ = v23
            }
        } else {
            if false {
                v7 = v7
                fmt.Println(v6)
                fmt.Println(v1, v5)
            }
            var v24 float64 = ((-2.25 / 1.0) / (0.1 / (float64((0.0 + (3.14159 + 0.0))) / ((0.0 + 0.1) + ((-2.25 - 0.0) / 0.1)))))
            v8 = v8
            _ = v24
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
