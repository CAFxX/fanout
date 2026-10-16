package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint; F2 bool }
type GenS3 struct { F0 uint; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 int64; F2 uint }
type GenS6 struct { F0 int8; F1 uint16; F2 int8 }
type GenS7 struct { F0 int8; F1 uint16; F2 bool }
type GenS8 struct { F0 int16; F1 int16 }
func genFn0(p0 int8) uint64 {
    switch 9223372036854775807 {
    case 1:
        v5 := GenS1{F0: 128}
        _ = v5
    case -28:
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 int = fi6
            for fi8 := 0; fi8 < 1; fi8++ {
                v7 = fi6
                v9 := []int{(fi6 / ((v7) | 1))}
                switch (fi6 + (fi8 >> ((0) & 63))) {
                case 211:
                    v9 = append(v9, 0)
                case 1:
                    fmt.Println(v7, fi6, p0)
                case 32767:
                    v10 := GenS2{F0: 990, F1: 65535, F2: (!(4294967295 < 1))}
                    _ = v10
                default:
                    var v11 float32 = 0.1
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            v12 := GenS3{F0: 42, F1: 49, F2: 1}
            _ = fi6
            _ = v12
            _ = v7
        }
    }
    v13 := []int{0, -46}
    _ = p0
    _ = v13
    return 0
}

func main() {
    var v1 uint64 = 6
    v2 := make(map[uint32]uint32)
    v3 := []uint32{788, 42, 0}
    var v4 int64 = 0
    if true {
        for fi14 := 0; fi14 < 8; fi14++ {
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 uint16 = 2
                fmt.Println(fi14)
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 4; fi17++ {
                v4 = (-(v4))
                fmt.Println(v4, fi17, v1)
                _ = fi17
            }
            fmt.Println(v1)
            _ = fi14
        }
    } else {
        for fi18 := 0; fi18 < 13; fi18++ {
            v1 = (^(v1))
            v19 := v2[2147483648]
            _ = fi18
            _ = v19
        }
    }
    if true {
        v4 -= (v4 & (v4 | 9223372036854775807))
    }
    var v20 float64 = 1e10
    v21 := GenS4{F0: 2}
    v1 += 18446744073709551615
    fmt.Println(v20, v1)
    v22 := GenS5{F0: genFn0(-128), F1: (v4 / ((-1) | 1)), F2: 557}
    switch 1 {
    case -1:
        v1 = v1
    case 786:
        v1 += 2
    case -2147483648:
        var v23 int16 = 1000
        _ = v23
    default:
        if false {
            v24 := GenS4{F0: 255}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v24
        }
    }
    var v25 uint8 = 103
    if ((390 <= -1) && true) {
        var v26 float64 = (v20 + v20)
        if (1 == v25) {
            v26 -= v20
            var v27 uint8 = (v25 & v25)
            var v28 float32 = 0.1
            _ = v27
            _ = v28
        } else {
            v29 := GenS6{F0: 72, F1: 451, F2: 0}
            delete(v2, 2147483647)
            _ = v29
        }
        _ = v26
    }
    fmt.Println(v4)
    if ((v25 << ((0) & 63)) == 19) {
        switch 319 {
        case 591:
            v25 += (v25 + (v25 >> ((778) & 63)))
        case -9223372036854775808:
            switch 32768 {
            case -886:
                fmt.Println(v20, v1)
            case 2147483647:
                if false {
                    for fi30 := 0; fi30 < 1; fi30++ {
                        var v31 uint = 65535
                        v4 = v4
                        _ = fi30
                        _ = v31
                    }
                }
            default:
                v25 -= (v25 + v25)
            }
        case 429:
            v20 = v20
        default:
            if (!((((v1 >> ((462) & 63)) == v1) || (127 < -128)) || true)) {
                v25 -= (v25 >> ((0) & 63))
                delete(v2, 2)
                v32 := GenS7{F0: 42, F1: 2, F2: (false && (0.0 <= (1e10 / 1.5)))}
                _ = v32
            }
        }
    } else {
        var v33 uint16 = 1
        _ = len(v2)
        var v34 uint64 = genFn0(100)
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 int64 = (v4 << ((1) & 63))
        v2[885] = 65535
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        for fi38 := 0; fi38 < 1; fi38++ {
            v1 -= (^(v1))
            var v39 uint8 = v25
            v4 = v4
            _ = fi38
            _ = v39
        }
        switch "a" {
        case "":
            for fi40 := 0; fi40 < 20; fi40++ {
                if false {
                    var v41 float32 = (-2.25 + 1e10)
                    switch (fi40 >> ((588) & 63)) {
                    case -1:
                        var v42 float64 = v20
                        _ = v42
                    }
                    var v43 int = 1
                    _ = v41
                    _ = v43
                }
                var v44 string = "foo bar"
                _ = fi40
                _ = v44
            }
        case "a":
            v45 := GenS8{F0: 0, F1: 135}
            _ = v45
        case "foo bar":
            var v46 int64 = v4
            _ = v46
        default:
            v25 -= (v25 ^ 255)
        }
        v4 = v4
        _ = fi37
    }
    var v47 uint8 = (v25 + v25)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v47
}
