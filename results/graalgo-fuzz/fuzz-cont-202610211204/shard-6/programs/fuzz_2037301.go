package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 bool; F1 bool; F2 string }
type GenS6 struct { F0 bool; F1 int16; F2 uint32 }

func main() {
    var v1 bool = (false && false)
    var v2 uint32 = 0
    v3 := []int{}
    var v4 int64 = -9223372036854775808
    var v5 uint64 = 448
    var v6 uint8 = 42
    var v7 int = 958
    switch ("foo bar" + "hello") {
    case "foo bar":
        _ = len(v3)
    case "a":
        switch 2147483647 {
        case 891:
            v8 := GenS1{F0: (v2 >= 42)}
            _ = v8
        case 370:
            v7 -= (v7 % (((v7 ^ (^(v7)))) | 1))
        case 42:
            switch "a" {
            case "":
                var v9 uint = 138
                _ = v9
            case "\n\t":
                v4 -= (v4 / ((int64(v7)) | 1))
            case "foo bar":
                if (!(42 == (v7 >> ((778) & 63)))) {
                    switch "a" {
                    case "hello":
                        var v10 int64 = v4
                        _ = v10
                    default:
                        var v11 int8 = 42
                        _ = v11
                    }
                    var v12 int64 = v4
                    v13 := GenS1{F0: v1}
                    _ = v12
                    _ = v13
                }
            }
        }
    case "":
        if v1 {
            v2 += (^(v2))
            v1 = true
        } else {
            if (((v5 ^ (v5 % (((v5 >> ((274) & 63))) | 1))) < 0) && v1) {
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 int32 = 2147483647
                    _ = fi14
                    _ = v15
                }
            } else {
                if (-2.25 >= ((float64(0.1) + 1.5) * 0.0)) {
                    var v16 int64 = v4
                    _ = v16
                } else {
                    v6 -= (v6 * v6)
                    var v17 uint16 = 42
                    _ = v17
                }
                fmt.Println(v5)
                var v18 bool = true
                _ = v18
            }
            for fi19 := 0; fi19 < 5; fi19++ {
                if (!v1) {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        switch (v7 >> ((2) & 63)) {
                        case 862:
                            v2 = v2
                        case 9223372036854775807:
                            v21 := GenS1{F0: ("a" < "a")}
                            _ = v21
                        default:
                            v1 = true
                        }
                        _ = fi20
                    }
                }
                _ = fi19
            }
            if (!(!(2 >= 127))) {
                for fi22 := 0; fi22 < 16; fi22++ {
                    v23 := GenS2{F0: 1}
                    var v24 uint16 = 0
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v25 := GenS2{F0: 773}
                for fi26 := 0; fi26 < 19; fi26++ {
                    v1 = (-1 != 10)
                    var v27 int16 = 1000
                    v28 := GenS1{F0: ((((255 == 136) && true) && true) && (!(v5 <= (v5 & uint64(v6)))))}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = v25
            }
        }
    default:
        switch "foo bar" {
        case "":
            var v29 int = -9223372036854775808
            _ = v29
        case "\n\t":
            if v1 {
                v3 = append(v3, 1)
                var v30 uint16 = 956
                _ = v30
            }
        default:
            v31 := GenS3{F0: 32767, F1: uint32(v5), F2: 42}
            _ = v31
        }
    }
    fmt.Println(v2, v6)
    fmt.Println(v1, v7, v1)
    fmt.Println(v2, v5, v7)
    v32 := GenS4{F0: v2, F1: ""}
    switch ("\n\t" + "foo bar") {
    case "":
        v5 = (v5 + (v5 & (v5 % ((v5) | 1))))
    case "a":
        v2 -= v2
    default:
        fmt.Println(v4, v2)
    }
    fmt.Println(v7, v1, v4)
    for fi33 := 0; fi33 < 6; fi33++ {
        v34 := GenS5{F0: false, F1: false, F2: "\n\t"}
        v35 := GenS3{F0: 1000, F1: 93, F2: 1}
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := GenS6{F0: v1, F1: -469, F2: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
