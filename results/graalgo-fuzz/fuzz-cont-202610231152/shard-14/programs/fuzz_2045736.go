package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint32 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int32; F1 int32; F2 int }
func genFn0(p0 uint16) int8 {
    fmt.Println(p0)
    p0 -= (p0 * 1)
    p0 = p0
    _ = p0
    return -1
}
func genFn1() uint64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 bool = true
        _ = fi7
        _ = v8
    }
    var v9 int32 = 2147483647
    var v10 string = (("" + "\n\t") + "\n\t")
    _ = v10
    _ = v9
    return 1
}

func main() {
    v1 := []int{}
    var v2 float64 = 1.5
    var v3 string = "a"
    var v4 int64 = -9223372036854775808
    var v5 string = (v3 + v3)
    var v6 uint16 = 1
    switch 732 {
    case 715:
        fmt.Println(v6)
    case 42:
        v11 := GenS1{F0: 127, F1: 1, F2: 845}
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v1 = append(v1, fi12)
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint8 = 35
            switch (fi13 & (fi12 ^ fi12)) {
            case 127:
                v15 := GenS1{F0: -19, F1: 1, F2: 2147483648}
                _ = v15
            case 128:
                var v16 string = v3
                _ = v16
            case 0:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi13
            _ = v14
        }
        fmt.Println(v5, v2, v3)
        _ = fi12
    }
    var v17 int = -92
    if (!false) {
        var v18 int = (v17 ^ (v17 * 42))
        _ = v18
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 float64 = 0.0
            _ = len(v1)
            var v21 int8 = 42
            _ = fi19
            _ = v20
            _ = v21
        }
        v3 = ("\n\t" + (v5 + (("hello" + (v5 + v5)) + "foo bar")))
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v23 := GenS1{F0: 0, F1: 78, F2: 4294967295}
        v3 = ((v5 + ("\n\t" + "hello")) + (v5 + "foo bar"))
        _ = fi22
        _ = v23
    }
    v6 += (v6 << ((255) & 63))
    for fi24 := 0; fi24 < 9; fi24++ {
        v3 = "foo bar"
        if len(v1) > 0 {
            v1[0] = v17
        }
        _ = fi24
    }
    var v25 uint8 = 128
    v2 -= 0.0
    if (v3 < v5) {
        if false {
            fmt.Println(v4, v2, v17)
            for fi26 := 0; fi26 < 4; fi26++ {
                for fi27 := 0; fi27 < 15; fi27++ {
                    if true {
                        v28 := GenS1{F0: -1, F1: 78, F2: 2}
                        switch "hello" {
                        case "hello":
                            v5 = v3
                            var v29 uint64 = 18446744073709551615
                            if (v2 < 3.14159) {
                                v17 -= 209
                                v25 = v25
                            } else {
                                v30 := GenS1{F0: 0, F1: genFn0(42), F2: 0}
                                var v31 uint32 = 242
                                _ = v30
                                _ = v31
                            }
                            _ = v29
                        case "foo bar":
                            v5 = (("\n\t" + v3) + "hello")
                            v32 := GenS1{F0: -128, F1: 42, F2: 815}
                            _ = v32
                        default:
                            if false {
                                var v33 int64 = -9223372036854775808
                                _ = v33
                            }
                        }
                        fmt.Println(fi27, v4)
                        _ = v28
                    }
                    v25 += uint8(v6)
                    v17 += (v17 << ((255) & 63))
                    var v34 float64 = v2
                    v35 := GenS1{F0: -122, F1: -128, F2: 2147483647}
                    _ = fi27
                    _ = v34
                    _ = v35
                }
                for fi36 := 0; fi36 < 3; fi36++ {
                    if ((3.14159 == (0.1 - 1e10)) || false) {
                        if false {
                            v37 := GenS1{F0: 2, F1: 100, F2: 770}
                            v38 := GenS2{F0: "hello", F1: (v25 ^ v25), F2: (2147483647 < 0)}
                            if true {
                                switch fi26 {
                                case 32768:
                                    fmt.Println(v17, v17)
                                }
                                if true {
                                    _ = len(v1)
                                    fmt.Println(v5, v6)
                                }
                            } else {
                                fmt.Println(v17)
                            }
                            if false {
                                if ((!false) || true) {
                                    var v39 uint32 = 0
                                    var v40 int = v17
                                    var v41 string = "hello"
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                } else {
                                    fmt.Println(fi36)
                                    var v42 float32 = (((0.1 / 1.0) + 0.0) * float32(v2))
                                    switch (fi26 ^ int(v25)) {
                                    case 32768:
                                        v5 = (v5 + "foo bar")
                                    }
                                    _ = v42
                                }
                                fmt.Println(fi26, fi36, fi26)
                                v6 += 65535
                            } else {
                                v43 := GenS3{F0: genFn0(334), F1: int32(v17)}
                                if false {
                                    switch "a" {
                                    case "a":
                                        var v44 uint8 = (v25 / ((v25) | 1))
                                        _ = v44
                                    case "\n\t":
                                        fmt.Println(v17)
                                    default:
                                        var v45 int32 = 2147483647
                                        _ = v45
                                    }
                                } else {
                                    if (0 <= 2) {
                                        v46 := GenS1{F0: 1, F1: 56, F2: 42}
                                        fmt.Println(v25, v17, v6)
                                        v2 = v2
                                        var v47 int16 = 1
                                        switch (fi36 & int(v6)) {
                                        case 9223372036854775807:
                                            v4 = 1
                                        case 2147483647:
                                            var v48 int16 = (^(v47))
                                            _ = v48
                                        case -9223372036854775808:
                                            switch (v3 + ("foo bar" + v3)) {
                                            case "\n\t":
                                                v4 = int64(v6)
                                            case "":
                                                fmt.Println(fi36, fi36)
                                            }
                                        }
                                        _ = v46
                                        _ = v47
                                    }
                                }
                                _ = v43
                            }
                            v49 := GenS4{F0: -87}
                            _ = v37
                            _ = v38
                            _ = v49
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v5 = v3
                    if (2147483647 <= 42) {
                        if len(v1) > 0 {
                            v1[0] = -1
                        }
                        var v50 float64 = v2
                        _ = len(v1)
                        _ = v50
                    } else {
                        v51 := GenS4{F0: -119}
                        var v52 uint = 2
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v51
                        _ = v52
                    }
                    _ = fi36
                }
                _ = fi26
            }
            v25 = v25
        } else {
            fmt.Println(v3, v6)
        }
        switch v17 {
        case 362:
            v53 := GenS2{F0: v3, F1: uint8(v6), F2: false}
            _ = v53
        case 2147483647:
            var v54 bool = false
            _ = v54
        case 42:
            for fi55 := 0; fi55 < 6; fi55++ {
                for fi56 := 0; fi56 < 18; fi56++ {
                    v4 -= (v4 * v4)
                    v57 := GenS3{F0: -92, F1: 42}
                    _ = fi56
                    _ = v57
                }
                var v58 float64 = v2
                switch 658 {
                case 9223372036854775807:
                    for fi59 := 0; fi59 < 9; fi59++ {
                        var v60 int64 = (v4 % ((-70) | 1))
                        _ = fi59
                        _ = v60
                    }
                case -9223372036854775808:
                    if (float32(v2) > 1e10) {
                        v58 -= (float64((1.5 / (0.1 * 0.1))) + v58)
                        v61 := GenS5{F0: true, F1: (v25 - v25)}
                        fmt.Println(v2, v25, v3)
                        _ = v61
                    }
                default:
                    if (!false) {
                        for fi62 := 0; fi62 < 3; fi62++ {
                            switch "hello" {
                            case "hello":
                                switch (fi62 * 4294967295) {
                                case 1:
                                    v63 := GenS6{F0: -341, F1: 42, F2: (v17 * fi62)}
                                    _ = v63
                                case -9223372036854775808:
                                    fmt.Println(v58)
                                }
                            }
                            v4 = v4
                            _ = fi62
                        }
                    }
                }
                _ = fi55
                _ = v58
            }
        }
        v5 = (v3 + ((v3 + "foo bar") + v3))
    }
    if (375 >= 255) {
        fmt.Println(v25, v17)
    } else {
        v6 -= (^(v6))
        v1 = append(v1, (v17 / (((v17 << ((0) & 63))) | 1)))
    }
    if true {
        v6 = 0
    }
    if false {
        v3 = ((v5 + "foo bar") + "a")
        v64 := GenS4{F0: -102}
        v3 = (v3 + v5)
        _ = v64
    } else {
        v1 = append(v1, (-(v17)))
    }
    switch ("hello" + ((v5 + v3) + (v5 + "foo bar"))) {
    case "\n\t":
        v2 -= 0.1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
