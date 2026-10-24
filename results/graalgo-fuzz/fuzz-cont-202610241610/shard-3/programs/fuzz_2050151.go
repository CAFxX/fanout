package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int32; F1 int64; F2 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{514, 65535, 65535, 65535}
    var v3 float32 = 3.14159
    var v4 int16 = -607
    v5 := make(map[uint32]string)
    var v6 uint8 = 157
    if true {
        var v7 int32 = 192
        _ = v7
    } else {
        v8 := GenS1{F0: -1}
        _ = v8
    }
    fmt.Println(v4)
    v9 := GenS2{F0: -118, F1: 321}
    fmt.Println(v6)
    delete(v5, 65535)
    if (4294967295 != 107) {
        var v10 int64 = 1
        if (2147483647 == -412) {
            if true {
                var v11 float64 = (1e10 - (1.5 + 1e10))
                if (false || (("a" <= ("" + (("" + ("\n\t" + "foo bar")) + "hello"))) || (!true))) {
                    switch 127 {
                    case 4294967295:
                        _ = len(v5)
                    default:
                        for fi12 := 0; fi12 < 19; fi12++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v1[(fi12 / (((fi12 << ((317) & 63))) | 1))] = 2147483647
                            _ = fi12
                        }
                    }
                } else {
                    v13 := GenS1{F0: 66}
                    _ = v13
                }
                for fi14 := 0; fi14 < 7; fi14++ {
                    switch (fi14 & fi14) {
                    case 608:
                        fmt.Println(v3)
                    case -9223372036854775808:
                        switch (-(fi14)) {
                        case 4294967295:
                            var v15 int = fi14
                            _ = v15
                        case 2147483647:
                            fmt.Println(v4, v11)
                        case 128:
                            fmt.Println(v3, v11, v11)
                        default:
                            var v16 float64 = v11
                            _ = v16
                        }
                    }
                    _ = fi14
                }
                _ = v11
            }
            v4 += (v4 << ((1) & 63))
        } else {
            v3 -= v3
            var v17 float64 = (0.0 + (3.14159 + (float64(((v3 + (v3 * (-2.25 * v3))) / (v3 - 0.1))) / ((float64(v3) * 3.14159) - 0.1))))
            v18 := GenS3{F0: 1, F1: v10, F2: (^(v6))}
            _ = v17
            _ = v18
        }
        var v19 bool = (false && true)
        _ = v10
        _ = v19
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        if false {
            var v21 int = -2147483648
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 float64 = 0.1
                switch v21 {
                case -329:
                    if true {
                        v3 = (1.5 + (v3 + v3))
                    }
                }
                _ = fi22
                _ = v23
            }
            _ = v21
        } else {
            if (false && (("" + "hello") != "\n\t")) {
                for fi24 := 0; fi24 < 13; fi24++ {
                    var v25 float64 = 1e10
                    if len(v2) > 0 {
                        v2[0] = 913
                    }
                    _ = fi24
                    _ = v25
                }
            }
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            var v27 uint8 = v6
            _ = fi26
            _ = v27
        }
        _ = fi20
    }
    v6 = (v6 - v6)
    fmt.Println(v4, v6, v6)
    if (1 < 4294967295) {
        fmt.Println(v6, v4, v3)
        switch ("a" + "\n\t") {
        case "foo bar":
            var v28 int64 = -1
            _ = v28
        case "":
            v3 = ((1e10 / v3) * (float32(0.0) / 1e10))
        default:
            v3 = (v3 - v3)
        }
        var v29 int8 = -1
        _ = v29
    } else {
        var v30 float64 = 1.5
        v3 = -2.25
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
