package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint; F2 uint32 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]uint32)
    var v3 int8 = 64
    v4 := make(map[int]string)
    var v5 int8 = 127
    v6 := []int{127, -825}
    fmt.Println(v3, v1)
    _ = len(v6)
    v7 := GenS1{F0: 1}
    var v8 int = 42
    if ((false || ((v5 << ((704) & 63)) >= v5)) || false) {
        fmt.Println(v3)
        if (1 <= 1) {
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(v5, v5, v5)
                _ = fi9
            }
            switch "" {
            case "a":
                switch 42 {
                case 32767:
                    if (!true) {
                        var v10 float64 = float64(0.1)
                        fmt.Println(v3, v8)
                        switch v8 {
                        case 663:
                            var v11 int64 = v1
                            _ = v11
                        case 0:
                            v6 = append(v6, (v8 >> ((0) & 63)))
                        default:
                            fmt.Println(v1, v5, v5)
                        }
                        _ = v10
                    }
                case 0:
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v5 = -105
                        fmt.Println(v3)
                        _ = fi12
                    }
                case -219:
                    for fi13 := 0; fi13 < 20; fi13++ {
                        switch (^(fi13)) {
                        case 42:
                            if false {
                                for fi14 := 0; fi14 < 9; fi14++ {
                                    var v15 string = ("a" + "")
                                    v1 = 483
                                    v3 += 42
                                    switch 0 {
                                    case 2147483647:
                                        var v16 int16 = -495
                                        _ = v16
                                    case 9223372036854775807:
                                        var v17 int8 = (v5 % ((-11) | 1))
                                        _ = v17
                                    case 368:
                                        var v18 int = fi14
                                        _ = v18
                                    }
                                    _ = fi14
                                    _ = v15
                                }
                                var v19 uint32 = 2147483647
                                v3 += v5
                                _ = v19
                            }
                        case 32767:
                            if (false || (0 != 18446744073709551615)) {
                                v3 = 1
                            }
                        }
                        _ = fi13
                    }
                }
            case "foo bar":
                var v20 int = v8
                _ = v20
            case "":
                v1 = 1
            default:
                v6 = append(v6, 552)
            }
        } else {
            var v21 bool = false
            _ = v21
        }
        fmt.Println(v3, v3)
    }
    v22 := GenS2{F0: true, F1: 476, F2: 774}
    v5 += (v5 / ((v5) | 1))
    var v23 float64 = (((0.0 * float64(1e10)) / 1.0) / 1.0)
    var v24 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
