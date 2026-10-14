package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 bool; F1 bool }

func main() {
    v1 := []int{0}
    v2 := []uint32{}
    var v3 string = "\n\t"
    v4 := []string{}
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint8 = 1
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v3 = v3
        switch ((v3 + (v3 + v3)) + v3) {
        case "":
            fmt.Println(v3, fi7)
        case "foo bar":
            for fi8 := 0; fi8 < 6; fi8++ {
                if false {
                    v3 = (v3 + v3)
                    var v9 string = v3
                    _ = v9
                } else {
                    v3 = (v3 + "")
                    var v10 float32 = 1.5
                    fmt.Println(fi8, fi8)
                    if (148 > 233) {
                        fmt.Println(fi7, fi7, v3)
                    }
                    if false {
                        var v11 uint16 = 42
                        var v12 int32 = -431
                        _ = v11
                        _ = v12
                    } else {
                        v10 += (v10 + v10)
                    }
                    _ = v10
                }
                switch 127 {
                case 32768:
                    fmt.Println(v3, fi7, fi7)
                case 0:
                    if true {
                        v13 := GenS1{F0: 18446744073709551615, F1: 167, F2: false}
                        v3 = v3
                        _ = v13
                    } else {
                        fmt.Println(fi8, v3)
                        switch "a" {
                        case "a":
                            v14 := GenS2{F0: 127, F1: -128, F2: 544}
                            _ = v14
                        case "foo bar":
                            switch "hello" {
                            case "hello":
                                var v15 uint8 = 128
                                _ = v15
                            case "":
                                v16 := GenS3{F0: 1, F1: 636}
                                _ = v16
                            }
                        default:
                            v4 = append(v4, v3)
                        }
                        var v17 float64 = (-2.25 * (0.1 - (((1.5 * 3.14159) * ((1.5 / -2.25) / (3.14159 + 0.0))) - ((1e10 / 1.5) * 0.0))))
                        _ = v17
                    }
                default:
                    var v18 uint8 = 62
                    _ = v18
                }
                _ = fi8
            }
        default:
            v3 = "\n\t"
        }
        _ = fi7
    }
    v3 = "\n\t"
    var v19 uint32 = 2147483647
    switch 127 {
    case 2147483647:
        for fi20 := 0; fi20 < 17; fi20++ {
            v3 = ("a" + (("a" + (v3 + v3)) + "foo bar"))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v21 string = "hello"
            _ = fi20
            _ = v21
        }
    case 579:
        fmt.Println(v19)
    }
    _ = len(v4)
    v19 = (v19 & (^(v19)))
    v22 := GenS2{F0: 1, F1: -1, F2: 2}
    v19 = (v19 >> ((4294967295) & 63))
    switch 729 {
    case 4294967295:
        v3 = ("" + "")
    default:
        for fi23 := 0; fi23 < 9; fi23++ {
            if false {
                v24 := GenS2{F0: 1, F1: 0, F2: 327}
                if false {
                    v19 += 42
                } else {
                    var v25 float64 = 3.14159
                    var v26 string = v3
                    _ = v25
                    _ = v26
                }
                var v27 int8 = 0
                _ = v24
                _ = v27
            }
            v3 = "foo bar"
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 int32 = 0
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi28
                _ = v29
            }
            _ = fi23
        }
    }
    var v30 string = v3
    switch 42 {
    case -9223372036854775808:
        v3 = v3
    default:
        v31 := GenS4{F0: 235, F1: 29}
        _ = v31
    }
    var v32 uint64 = 42
    v19 -= v19
    switch -643 {
    case 563:
        v33 := GenS1{F0: uint(v19), F1: 128, F2: true}
        _ = v33
    case -546:
        v34 := GenS5{F0: (!false), F1: (946 <= 68)}
        _ = v34
    }
    if (128 > 0) {
        v32 = 18446744073709551615
    } else {
        _ = len(v2)
    }
    var v35 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v32
    _ = v35
    _ = v4
}
