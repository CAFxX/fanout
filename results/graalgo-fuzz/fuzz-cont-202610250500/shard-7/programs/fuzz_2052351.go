package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int8; F2 int16 }

func main() {
    var v1 uint16 = 404
    var v2 string = "hello"
    var v3 int = 42
    var v4 uint16 = 643
    var v5 uint16 = (v4 << ((4294967295) & 63))
    var v6 uint = 68
    var v7 uint8 = 2
    v3 = -1
    for fi8 := 0; fi8 < 7; fi8++ {
        if (!true) {
            switch (v2 + v2) {
            case "":
                switch v3 {
                case -1:
                    if (520 > 42) {
                        for fi9 := 0; fi9 < 16; fi9++ {
                            v10 := GenS1{F0: -1}
                            v3 += fi9
                            v11 := GenS2{F0: 42, F1: 31, F2: 0}
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                        fmt.Println(v5, v5, v4)
                    }
                case -9223372036854775808:
                    if (((v4 % ((v1) | 1)) <= 1) && true) {
                        v12 := []int{(v3 & v3), fi8, v3, 128}
                        _ = v12
                    }
                default:
                    v1 = v1
                }
            case "hello":
                switch (fi8 * -392) {
                case 566:
                    if ((((1.5 + 0.0) / 0.1) == float32(-2.25)) || (!false)) {
                        if ((v3 + (fi8 << ((255) & 63))) != (fi8 + int(v6))) {
                            var v13 uint8 = v7
                            for fi14 := 0; fi14 < 17; fi14++ {
                                if true {
                                    fmt.Println(v1)
                                } else {
                                    v15 := []int{}
                                    _ = v15
                                }
                                fmt.Println(v5, v13, v4)
                                _ = fi14
                            }
                            v16 := GenS1{F0: 738}
                            _ = v13
                            _ = v16
                        }
                    } else {
                        for fi17 := 0; fi17 < 17; fi17++ {
                            var v18 int = int(v7)
                            var v19 int64 = 9223372036854775807
                            var v20 int16 = 496
                            v21 := []int{1, -2147483648}
                            if (42 <= 2) {
                                var v22 int = fi8
                                v23 := GenS2{F0: 2147483647, F1: 0, F2: v20}
                                _ = v22
                                _ = v23
                            } else {
                                fmt.Println(v5, fi17, fi8)
                                var v24 uint16 = v5
                                _ = v24
                            }
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        v25 := []int{}
                        for fi26 := 0; fi26 < 16; fi26++ {
                            v5 = v1
                            _ = fi26
                        }
                        _ = v25
                    }
                case -9223372036854775808:
                    v1 += v4
                default:
                    fmt.Println(v3, v2)
                }
            case "a":
                v27 := GenS1{F0: 361}
                _ = v27
            default:
                v4 -= (v1 % (((v1 + (^(v5)))) | 1))
            }
            var v28 uint = v6
            v4 -= (v4 - (^(v4)))
            _ = v28
        }
        _ = fi8
    }
    v5 = v1
    switch (v3 << (((v6 + v6)) & 63)) {
    case 272:
        switch -386 {
        case 502:
            if (!false) {
                if false {
                    v29 := []int{v3, 4294967295, (v3 & -38), 74}
                    _ = v29
                }
                var v30 int16 = 0
                v6 -= (^(v6))
                _ = v30
            }
        default:
            var v31 int32 = 0
            _ = v31
        }
    case -503:
        for fi32 := 0; fi32 < 6; fi32++ {
            v33 := GenS1{F0: 2147483647}
            _ = fi32
            _ = v33
        }
    case -570:
        for fi34 := 0; fi34 < 11; fi34++ {
            fmt.Println(v6, v6, fi34)
            _ = fi34
        }
    }
    var v35 int64 = -997
    v36 := []int{v3}
    fmt.Println(v3, v6, v35)
    v7 = 0
    v7 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
