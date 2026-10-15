package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int16 }
type GenS2 struct { F0 uint32; F1 int64; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 int; F2 int }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 int8; F1 uint32; F2 string }

func main() {
    var v1 int32 = -2147483648
    var v2 uint = 1
    v3 := []int{128}
    v4 := make(map[uint32]bool)
    var v5 uint = 18446744073709551615
    var v6 float32 = (((-2.25 - 0.1) / 1.0) + (3.14159 + 1.5))
    var v7 int8 = -66
    if (v1 <= -2147483648) {
        v5 = (v2 & (v5 & (v5 - 18446744073709551615)))
        v8 := GenS1{F0: 287, F1: -1, F2: -1}
        _ = v8
    } else {
        v9 := GenS1{F0: 0, F1: 0, F2: 1000}
        _ = v9
    }
    v10 := GenS2{F0: 2147483648, F1: -1, F2: ("a" + "\n\t")}
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 bool = (true && (v5 != v5))
        _ = fi11
        _ = v12
    }
    v13 := GenS3{F0: 4294967295}
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(v6)
        var v15 int = 9223372036854775807
        for fi16 := 0; fi16 < 9; fi16++ {
            if (!true) {
                fmt.Println(v1, v5)
                v17 := GenS4{F0: -1, F1: v15, F2: (fi14 ^ fi16)}
                _ = v17
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    switch 9223372036854775807 {
    case 272:
        var v18 uint = 2
        _ = v18
    case 1:
        for fi19 := 0; fi19 < 9; fi19++ {
            v5 = (v2 << ((255) & 63))
            v7 = v7
            for fi20 := 0; fi20 < 13; fi20++ {
                if (true || false) {
                    v5 -= 255
                }
                v21 := GenS5{F0: (("" + "a") + ""), F1: v7}
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    default:
        v22 := GenS1{F0: 1000, F1: 609, F2: -1}
        _ = v22
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            var v25 int32 = v1
            v2 += 2
            if (false && (!(!("" != (("hello" + ("\n\t" + "a")) + "a"))))) {
                if true {
                    var v26 int32 = v1
                    v2 -= 65535
                    _ = v26
                } else {
                    fmt.Println(v7)
                }
                switch fi23 {
                case 9223372036854775807:
                    var v27 int = 42
                    _ = v27
                case -707:
                    switch "foo bar" {
                    case "hello":
                        v28 := GenS1{F0: 42, F1: 9223372036854775807, F2: 1000}
                        _ = v28
                    default:
                        var v29 float64 = (-2.25 / ((-2.25 / (float64(v6) / (float64(-2.25) / 3.14159))) - 0.1))
                        _ = v29
                    }
                }
                var v30 uint = v5
                _ = v30
            } else {
                if (false || false) {
                    v7 = (v7 << ((554) & 63))
                    switch "\n\t" {
                    case "\n\t":
                        var v31 int = 626
                        _ = v31
                    case "foo bar":
                        v32 := GenS6{F0: 0, F1: 255, F2: 2}
                        _ = v32
                    default:
                        fmt.Println(v1, fi23)
                    }
                    fmt.Println(v6)
                } else {
                    if true {
                        var v33 uint64 = 30
                        v3 = append(v3, (fi24 & (fi23 >> (((v2 << (((^(v5))) & 63))) & 63))))
                        if (false || ((!true) || false)) {
                            v34 := GenS7{F0: int(v1), F1: (!false)}
                            fmt.Println(v5, v7)
                            _ = v34
                        }
                        _ = v33
                    } else {
                        var v35 int = (^(fi23))
                        _ = v35
                    }
                    v1 = (v1 - -2147483648)
                    var v36 uint64 = 185
                    _ = v36
                }
            }
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    fmt.Println(v7)
    switch "a" {
    case "":
        v1 = v1
    case "foo bar":
        v6 -= v6
    default:
        v37 := GenS8{F0: -39, F1: 1, F2: "\n\t"}
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
