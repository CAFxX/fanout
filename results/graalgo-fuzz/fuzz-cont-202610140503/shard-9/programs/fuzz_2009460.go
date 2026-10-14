package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 bool }
func genFn0() int16 {
    var v5 uint32 = 4294967295
    switch "" {
    case "\n\t":
        var v6 int8 = 1
        _ = v6
    case "hello":
        var v7 uint16 = 278
        _ = v7
    case "a":
        v8 := GenS1{F0: -89, F1: "", F2: 9223372036854775807}
        _ = v8
    }
    _ = v5
    return -554
}
func genFn1() int16 {
    var v9 int32 = -2147483648
    v9 = v9
    v9 = -2147483648
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"hello", "foo bar", "a", ("hello" + "foo bar")}
    v3 := make(map[uint32]bool)
    var v4 uint64 = 1
    switch "\n\t" {
    case "a":
        if true {
            v4 += (v4 << ((255) & 63))
            fmt.Println(v4)
        } else {
            var v10 uint16 = 2
            var v11 int16 = -376
            _ = v10
            _ = v11
        }
    default:
        v4 = 18446744073709551615
    }
    var v12 uint16 = 21
    var v13 int64 = -1
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            v4 = v4
            _ = fi15
        }
        if (4294967295 <= 4294967295) {
            var v16 float32 = 1.5
            v17 := GenS1{F0: 0, F1: ("" + "\n\t"), F2: fi14}
            _ = v16
            _ = v17
        } else {
            if true {
                v12 = (v12 ^ v12)
            } else {
                var v18 uint16 = (v12 + v12)
                v12 = v12
                _ = v18
            }
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            if true {
                switch fi20 {
                case -9223372036854775808:
                    v13 = (v13 / ((v13) | 1))
                default:
                    v21 := GenS2{F0: 42, F1: 42}
                    _ = v21
                }
                v12 += 42
            } else {
                fmt.Println(v4)
            }
            if (984 <= 255) {
                switch (fi20 % (((fi19 >> ((18446744073709551615) & 63))) | 1)) {
                case 32768:
                    if ((v12 * v12) < 65535) {
                        switch "a" {
                        case "a":
                            var v22 bool = (((true || (((-1 < -1) || (!(!(1.5 == -2.25)))) && true)) || true) && false)
                            _ = v22
                        }
                    }
                case 2147483647:
                    fmt.Println(v12, fi20, fi19)
                }
            }
            if true {
                if true {
                    v4 = 1
                    fmt.Println(fi20)
                    v23 := GenS2{F0: -932, F1: uint8(v4)}
                    _ = v23
                }
                switch (fi20 << ((65535) & 63)) {
                case 32768:
                    v13 += (v13 + v13)
                    var v24 uint64 = v4
                    if true {
                        fmt.Println(v24, fi19)
                        var v25 uint32 = 0
                        _ = v25
                    } else {
                        _ = len(v2)
                        if false {
                            fmt.Println(v4, fi19, v24)
                        } else {
                            v26 := GenS2{F0: genFn1(), F1: 204}
                            var v27 int8 = 42
                            if (42 <= (v13 | (v13 << ((1) & 63)))) {
                                v24 += v4
                            }
                            v28 := GenS3{F0: (false || (0 >= 128))}
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        v29 := GenS1{F0: -2147483648, F1: "\n\t", F2: (fi19 * 127)}
                        _ = v29
                    }
                    v30 := v3[2147483647]
                    v12 = v12
                    _ = v24
                    _ = v30
                }
                v13 += (v13 & v13)
                var v31 int64 = (v13 - (v13 & v13))
                fmt.Println(fi20, fi20, v4)
                v12 = (v12 & v12)
                _ = v31
            }
            _ = fi20
        }
        v12 += (v12 >> ((4294967295) & 63))
        _ = fi19
    }
    var v32 uint32 = 989
    switch -1 {
    case -887:
        var v33 int16 = 42
        _ = v33
    case -66:
        var v34 float32 = float32((0.1 + (0.1 * 1e10)))
        _ = v34
    default:
        var v35 int = -1
        _ = v35
    }
    if ((2 == 531) || true) {
        var v36 int32 = -2147483648
        switch "" {
        case "":
            var v37 float32 = 3.14159
            _ = v37
        }
        v4 = v4
        _ = v36
    } else {
        if (1e10 >= float32(float64(float32((((float64(3.14159) * -2.25) / 1.5) / 1.5))))) {
            v12 = (v12 - (v12 | v12))
        }
        delete(v3, v32)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
