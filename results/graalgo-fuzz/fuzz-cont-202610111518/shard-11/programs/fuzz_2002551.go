package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 bool }

func main() {
    var v1 int16 = 32767
    var v2 int16 = -1
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]string)
    v5 := []uint32{}
    v2 += v2
    v2 = 0
    if ((false || false) || (65535 == 2)) {
        if ((0.0 - float64(0.0)) == 1e10) {
            _ = len(v4)
            var v6 int32 = 42
            _ = v6
        }
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS1{F0: 313}
            fmt.Println(v1)
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(v1, v1, v1)
        var v10 uint16 = 255
        _ = fi9
        _ = v10
    }
    v1 = (v1 | v1)
    v2 -= 1
    for fi11 := 0; fi11 < 17; fi11++ {
        if (2 >= 1) {
            if true {
                v12 := v4[(^(fi11))]
                switch (fi11 >> ((4294967295) & 63)) {
                case -919:
                    switch v12 {
                    case "":
                        if (false || false) {
                            var v13 uint8 = 192
                            _ = v13
                        } else {
                            v14 := GenS2{F0: (^(v1)), F1: (18446744073709551615 <= 660)}
                            if len(v5) > 0 {
                                v5[0] = 664
                            }
                            v15 := GenS2{F0: v2, F1: true}
                            _ = v14
                            _ = v15
                        }
                    case "foo bar":
                        switch (fi11 % ((fi11) | 1)) {
                        case 4294967295:
                            for fi16 := 0; fi16 < 3; fi16++ {
                                if false {
                                    fmt.Println(v2)
                                    v17 := GenS1{F0: 255}
                                    if len(v5) > 0 {
                                        v5[0] = 582
                                    }
                                    _ = v17
                                }
                                _ = fi16
                            }
                        case -773:
                            var v18 uint16 = 450
                            _ = v18
                        default:
                            fmt.Println(v12, v1)
                        }
                    case "hello":
                        v1 = (v1 & v1)
                    }
                case -754:
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v20 := GenS1{F0: 290}
                        var v21 uint64 = 691
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                default:
                    fmt.Println(v1, v2, v2)
                }
                _ = v12
            } else {
                if true {
                    if false {
                        switch fi11 {
                        case -873:
                            for fi22 := 0; fi22 < 3; fi22++ {
                                v3[65535] = 4294967295
                                _ = fi22
                            }
                        }
                    }
                    delete(v3, 4294967295)
                }
            }
        }
        _ = fi11
    }
    var v23 float64 = 0.0
    fmt.Println(v2, v23, v23)
    fmt.Println(v2)
    v2 = (-(v2))
    v24 := GenS1{F0: 18446744073709551615}
    switch "a" {
    case "\n\t":
        var v25 int64 = 196
        _ = v25
    case "a":
        for fi26 := 0; fi26 < 10; fi26++ {
            v5 = append(v5, 2147483647)
            _ = fi26
        }
    default:
        var v27 float64 = (v23 - ((v23 - v23) + v23))
        _ = v27
    }
    v23 = float64(0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
