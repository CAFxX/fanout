package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS4 struct { F0 int32; F1 int; F2 uint32 }
type GenS5 struct { F0 int }

func main() {
    v1 := []string{"a", ""}
    var v2 int8 = -128
    var v3 int = 1
    var v4 float32 = -2.25
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 14; fi6++ {
        v2 = (v2 >> ((1) & 63))
        _ = fi6
    }
    if (!false) {
        if (!true) {
            var v7 int64 = 1
            _ = v7
        }
    } else {
        v3 += (v3 / ((v3) | 1))
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int64 = 1
            if ((v3 >> ((635) & 63)) >= 2147483647) {
                fmt.Println(v4, v2, v10)
            }
            fmt.Println(v3)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if (v2 >= v2) {
        switch -2147483648 {
        case -803:
            fmt.Println(v3, v3)
        default:
            v4 += (((0.0 / v4) + 0.0) + 0.0)
        }
        v2 -= (v2 ^ (-(v2)))
        switch "a" {
        case "hello":
            v5[v3] = (("\n\t" + "\n\t") + "foo bar")
        default:
            delete(v5, 127)
        }
    } else {
        for fi11 := 0; fi11 < 6; fi11++ {
            fmt.Println(v4, v4)
            var v12 int64 = 9223372036854775807
            if true {
                var v13 string = ""
                switch v13 {
                case "\n\t":
                    if true {
                        v13 = "\n\t"
                        _ = len(v1)
                        v14 := GenS1{F0: 0}
                        _ = v14
                    }
                case "a":
                    v5[(v3 + v3)] = v13
                default:
                    switch (v13 + v13) {
                    case "foo bar":
                        for fi15 := 0; fi15 < 15; fi15++ {
                            var v16 float64 = float64((float32(3.14159) * (v4 + ((0.0 * 3.14159) / (3.14159 - v4)))))
                            v3 += (fi11 >> ((86) & 63))
                            v12 -= 1
                            var v17 uint16 = 32767
                            if false {
                                v18 := v5[int(v12)]
                                v2 -= (v2 - (v2 >> ((1) & 63)))
                                var v19 uint64 = 18446744073709551615
                                _ = v18
                                _ = v19
                            } else {
                                v1 = append(v1, (v13 + v13))
                            }
                            v4 = (v4 / v4)
                            v3 += fi15
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                }
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        v20 := GenS2{F0: (v3 / ((v3) | 1)), F1: (!false), F2: 1}
        _ = v20
    }
    var v21 int16 = 287
    v2 = (^(v2))
    if (478 != 2) {
        v3 -= (v3 >> ((735) & 63))
        v22 := GenS3{F0: 42, F1: 1, F2: 1000}
        _ = v22
    } else {
        v3 -= v3
        if false {
            if (!true) {
                v3 = v3
                var v23 int8 = v2
                v3 -= (v3 << ((255) & 63))
                _ = v23
            } else {
                switch "a" {
                case "hello":
                    v4 = (v4 - v4)
                default:
                    var v24 string = "a"
                    _ = v24
                }
                v4 = 0.1
            }
            v25 := GenS4{F0: 42, F1: 1, F2: 2}
            if (-1 > int(v21)) {
                v4 = 1.5
            } else {
                for fi26 := 0; fi26 < 19; fi26++ {
                    v21 = (-(v21))
                    _ = len(v1)
                    var v27 uint16 = 65535
                    _ = fi26
                    _ = v27
                }
            }
            _ = v25
        }
        var v28 uint8 = 128
        _ = v28
    }
    if false {
        v29 := GenS3{F0: 1, F1: 989, F2: 1}
        v30 := GenS5{F0: v3}
        _ = v29
        _ = v30
    }
    v4 += ((v4 / (1e10 * (v4 / v4))) + (v4 * v4))
    v1 = append(v1, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
