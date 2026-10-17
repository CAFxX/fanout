package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 string }

func main() {
    var v1 uint64 = 973
    var v2 int8 = -72
    v3 := make(map[int]uint32)
    v4 := []uint32{586, 429, 0}
    var v5 uint16 = 255
    v6 := make(map[int]uint32)
    v5 -= 554
    switch (("" + "\n\t") + "foo bar") {
    case "\n\t":
        v1 = 911
    case "foo bar":
        var v7 uint16 = 2
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 1; fi8++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if true {
                if false {
                    v9 := GenS1{F0: true, F1: 364}
                    fmt.Println(fi8, v2, v2)
                    _ = v9
                } else {
                    if (-40 == 20) {
                        v10 := GenS2{F0: ((false || ((1 > 1) && false)) && false), F1: -85}
                        v11 := GenS3{F0: ("\n\t" + "foo bar")}
                        _ = v10
                        _ = v11
                    }
                    if false {
                        v5 = v5
                    }
                    v1 += (v1 - 18446744073709551615)
                }
            }
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 int8 = v2
                if (!false) {
                    var v14 bool = (v5 != (v5 | (v5 + v5)))
                    v5 = 255
                    if ((((127 <= 1) && v14) || v14) || v14) {
                        v5 += 255
                        if ((v13 / ((1) | 1)) >= (v13 >> ((238) & 63))) {
                            v2 = v13
                        }
                    } else {
                        switch "" {
                        case "hello":
                            var v15 bool = v14
                            _ = v15
                        case "a":
                            var v16 uint16 = 762
                            _ = v16
                        }
                        fmt.Println(v14)
                        if v14 {
                            _ = len(v3)
                            v2 = v2
                            v4 = append(v4, 68)
                        } else {
                            fmt.Println(v1)
                            fmt.Println(v5, v1, fi8)
                            v1 = v1
                        }
                    }
                    _ = v14
                }
                _ = fi12
                _ = v13
            }
            _ = fi8
        }
    }
    v2 += (v2 + v2)
    var v17 int8 = (v2 | v2)
    var v18 int64 = int64(v5)
    delete(v6, -29)
    if (!true) {
        for fi19 := 0; fi19 < 10; fi19++ {
            fmt.Println(fi19)
            v18 = -645
            switch ("" + "foo bar") {
            case "hello":
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 int64 = v18
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi19
        }
    } else {
        switch 2147483647 {
        case 0:
            if (true || true) {
                v5 -= v5
            }
        }
    }
    v1 -= v1
    var v22 int32 = 1
    switch 9223372036854775807 {
    case 98:
        v2 = v17
    default:
        switch 4294967295 {
        case 9223372036854775807:
            var v23 string = "hello"
            _ = v23
        case 155:
            switch "a" {
            case "a":
                switch "foo bar" {
                case "a":
                    var v24 uint8 = 107
                    _ = v24
                case "foo bar":
                    v25 := GenS1{F0: (!true), F1: 749}
                    _ = v25
                default:
                    var v26 int64 = -536
                    _ = v26
                }
            }
        case 32767:
            switch "hello" {
            case "a":
                var v27 uint32 = 119
                _ = v27
            case "":
                v5 = 562
            case "\n\t":
                switch 2147483647 {
                case 32767:
                    v5 -= v5
                case 127:
                    var v28 int64 = (v18 & v18)
                    _ = v28
                case -2147483648:
                    var v29 uint32 = 42
                    _ = v29
                default:
                    v17 -= v17
                }
            }
        }
    }
    var v30 float32 = 1e10
    switch ("foo bar" + "hello") {
    case "foo bar":
        var v31 float64 = 0.1
        _ = v31
    default:
        v22 += (v22 + (v22 & 2147483647))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
