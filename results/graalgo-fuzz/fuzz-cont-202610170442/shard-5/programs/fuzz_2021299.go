package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int8; F1 string; F2 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 string }
func genFn0(p0 int, p1 uint, p2 uint16) uint64 {
    var v4 uint8 = 78
    var v5 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 18446744073709551615
}

func main() {
    v1 := []uint32{2147483648, 65535, 2147483647}
    var v2 uint64 = 829
    v3 := []string{}
    v2 += (v2 >> ((255) & 63))
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int = (fi6 + (fi6 ^ 270))
        _ = fi6
        _ = v7
    }
    if (18446744073709551615 > 4294967295) {
        v8 := GenS1{F0: 38, F1: 791}
        _ = v8
    } else {
        _ = len(v1)
    }
    var v9 int16 = -1
    var v10 uint8 = 2
    switch -1 {
    case 2147483647:
        for fi11 := 0; fi11 < 20; fi11++ {
            if (float64(1e10) <= 1.5) {
                if false {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v9 = (v9 | v9)
                        _ = fi12
                    }
                } else {
                    var v13 float32 = -2.25
                    if true {
                        var v14 int32 = 34
                        v13 += (v13 - (v13 - v13))
                        v15 := GenS2{F0: 1, F1: 1}
                        _ = v14
                        _ = v15
                    }
                    _ = v13
                }
                var v16 uint8 = v10
                fmt.Println(v2)
                _ = v16
            }
            v2 = (^(v2))
            v2 -= v2
            _ = fi11
        }
    case -1:
        v2 = v2
    case 624:
        switch -452 {
        case 9223372036854775807:
            if ("a" > ("" + ("hello" + "hello"))) {
                var v17 uint8 = (v10 << ((4294967295) & 63))
                _ = v17
            } else {
                switch "a" {
                case "\n\t":
                    if false {
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v2)
                        case "hello":
                            fmt.Println(v2)
                        default:
                            if (!(false && ((-1 <= (v9 * (v9 | v9))) && false))) {
                                v18 := GenS1{F0: (v10 << ((4294967295) & 63)), F1: 32767}
                                var v19 int = 128
                                var v20 float32 = ((1e10 - float32(3.14159)) - -2.25)
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                v21 := GenS1{F0: v10, F1: 250}
                                v22 := GenS2{F0: 0, F1: v10}
                                fmt.Println(v10, v9)
                                _ = v21
                                _ = v22
                            }
                        }
                    }
                default:
                    v23 := GenS3{F0: -1, F1: "\n\t", F2: (v10 & (v10 ^ 128))}
                    _ = v23
                }
                v2 = uint64(v10)
            }
        case -1:
            v24 := GenS2{F0: 944, F1: v10}
            _ = v24
        case 558:
            if (-9223372036854775808 < 0) {
                fmt.Println(v2, v2)
            }
        }
    }
    switch -444 {
    case 353:
        if false {
            v25 := GenS4{F0: 65535}
            var v26 int16 = (v9 ^ -32768)
            _ = v25
            _ = v26
        }
    }
    var v27 string = "foo bar"
    v28 := GenS1{F0: v10, F1: 617}
    v29 := GenS5{F0: v2, F1: -247}
    for fi30 := 0; fi30 < 12; fi30++ {
        for fi31 := 0; fi31 < 16; fi31++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi31
        }
        fmt.Println(v27, v27)
        _ = fi30
    }
    var v32 int8 = -17
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 uint8 = v10
        v9 = v9
        _ = fi33
        _ = v34
    }
    switch ("\n\t" + ("hello" + (v27 + (("hello" + v27) + "a")))) {
    case "hello":
        v35 := GenS6{F0: "foo bar"}
        _ = v35
    case "":
        var v36 uint8 = 42
        _ = v36
    case "foo bar":
        if true {
            v2 -= v2
        }
    default:
        fmt.Println(v2, v9, v32)
    }
    v32 = 1
    var v37 int = 32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v37
    _ = v9
}
