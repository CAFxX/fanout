package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int8 }

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]bool)
    v3 := []uint32{2147483647, 2147483647}
    var v4 uint16 = v1
    var v5 int16 = 42
    var v6 uint64 = 1
    v5 = (v5 * (v5 | v5))
    var v7 float32 = (float32(float64(((1e10 - 1e10) / -2.25))) - 0.0)
    v8 := GenS1{F0: -128, F1: 0}
    var v9 bool = (true || true)
    if ((v6 / ((42) | 1)) != v6) {
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS1{F0: 100, F1: -410}
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 uint8 = 1
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        var v14 bool = (v7 < 1e10)
        _ = v14
    }
    v4 = 238
    v5 = 673
    if v9 {
        if (255 == v1) {
            v15 := GenS1{F0: int8(v6), F1: 265}
            _ = v15
        }
    } else {
        var v16 int64 = 9223372036854775807
        if true {
            switch "a" {
            case "\n\t":
                if (0 != uint8(v4)) {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        fmt.Println(v5, v9, v4)
                        for fi18 := 0; fi18 < 2; fi18++ {
                            v6 += v6
                            if ((int64(v1) >= (v16 | (v16 * int64(v6)))) && v9) {
                                v3 = append(v3, 696)
                                fmt.Println(fi17)
                                switch "" {
                                case "hello":
                                    v19 := v2[1]
                                    _ = v19
                                case "a":
                                    _ = len(v2)
                                case "\n\t":
                                    v9 = (v9 || v9)
                                default:
                                    v20 := v2[926]
                                    _ = v20
                                }
                            } else {
                                switch ((("a" + "") + ("hello" + "a")) + ("a" + "a")) {
                                case "hello":
                                    var v21 int16 = v5
                                    _ = v21
                                case "\n\t":
                                    switch "foo bar" {
                                    case "foo bar":
                                        v22 := GenS1{F0: -117, F1: -391}
                                        _ = v22
                                    default:
                                        switch (^(fi18)) {
                                        case -584:
                                            v23 := GenS1{F0: 127, F1: int64(v6)}
                                            switch fi17 {
                                            case -408:
                                                var v24 uint16 = (^(v1))
                                                v25 := GenS1{F0: 120, F1: -1}
                                                var v26 int8 = 75
                                                if v9 {
                                                    switch ("\n\t" + "foo bar") {
                                                    case "\n\t":
                                                        var v27 uint32 = 259
                                                        _ = v27
                                                    case "a":
                                                        var v28 int16 = (v5 & v5)
                                                        _ = v28
                                                    case "":
                                                        fmt.Println(v26, fi17)
                                                    default:
                                                        v29 := GenS1{F0: (v26 << ((0) & 63)), F1: v16}
                                                        _ = v29
                                                    }
                                                    var v30 uint32 = uint32(v4)
                                                    _ = v30
                                                }
                                                _ = v24
                                                _ = v25
                                                _ = v26
                                            }
                                            var v31 uint = 4294967295
                                            _ = v23
                                            _ = v31
                                        case 32767:
                                            if v9 {
                                                v4 = (v4 + (v4 % ((v4) | 1)))
                                            } else {
                                                v4 = v1
                                                var v32 uint64 = (v6 ^ (v6 % (((v6 << ((65535) & 63))) | 1)))
                                                v32 -= v6
                                                _ = v32
                                            }
                                        }
                                    }
                                }
                                v3 = append(v3, 4294967295)
                            }
                            v9 = (!v9)
                            _ = fi18
                        }
                        var v33 int16 = (v5 / (((v5 >> ((4294967295) & 63))) | 1))
                        _ = fi17
                        _ = v33
                    }
                    if false {
                        v1 += uint16(v6)
                        switch 9223372036854775807 {
                        case 1:
                            var v34 uint = 386
                            _ = v34
                        default:
                            var v35 int16 = v5
                            _ = v35
                        }
                        v36 := GenS1{F0: -128, F1: (v16 - v16)}
                        _ = v36
                    } else {
                        switch "\n\t" {
                        case "hello":
                            var v37 float32 = (1.5 * float32(((3.14159 / 1e10) * 1e10)))
                            _ = v37
                        case "a":
                            v3 = append(v3, 4294967295)
                        case "foo bar":
                            if true {
                                var v38 uint32 = 2147483647
                                _ = v38
                            } else {
                                v39 := GenS1{F0: 74, F1: (^(v16))}
                                for fi40 := 0; fi40 < 4; fi40++ {
                                    for fi41 := 0; fi41 < 9; fi41++ {
                                        v42 := GenS1{F0: int8(v16), F1: v16}
                                        v43 := GenS1{F0: 18, F1: (v16 >> ((2) & 63))}
                                        _ = fi41
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = fi40
                                }
                                switch ("hello" + "\n\t") {
                                case "a":
                                    switch ("foo bar" + ("a" + "foo bar")) {
                                    case "foo bar":
                                        var v44 bool = ((-2.25 + 1e10) >= 3.14159)
                                        _ = v44
                                    }
                                case "\n\t":
                                    v45 := GenS1{F0: int8(v4), F1: v16}
                                    _ = v45
                                default:
                                    delete(v2, 4294967295)
                                }
                                _ = v39
                            }
                        default:
                            v46 := GenS1{F0: 79, F1: (v16 - (v16 - 0))}
                            _ = v46
                        }
                        var v47 float64 = ((0.1 + -2.25) * float64(1.5))
                        _ = v47
                    }
                    if true {
                        if v9 {
                            fmt.Println(v5, v16)
                            if (v9 && (42 > 18446744073709551615)) {
                                v48 := GenS2{F0: 0, F1: 0}
                                delete(v2, 1)
                                v6 -= (v6 * v6)
                                _ = v48
                            }
                            if v9 {
                                v49 := GenS3{F0: 255}
                                if v9 {
                                    for fi50 := 0; fi50 < 18; fi50++ {
                                        v51 := GenS4{F0: 2147483647, F1: "hello"}
                                        _ = fi50
                                        _ = v51
                                    }
                                    v4 += 42
                                    if (v5 < (v5 / ((-979) | 1))) {
                                        v5 -= -66
                                        for fi52 := 0; fi52 < 15; fi52++ {
                                            v7 = v7
                                            _ = fi52
                                        }
                                        var v53 int32 = 2147483647
                                        _ = v53
                                    }
                                } else {
                                    if (v9 || (!(!v9))) {
                                        var v54 uint8 = 127
                                        switch "" {
                                        case "hello":
                                            fmt.Println(v7, v7, v9)
                                        case "foo bar":
                                            _ = len(v2)
                                        case "a":
                                            if v9 {
                                                var v55 float64 = ((0.0 + (0.1 - (((float64(v7) / -2.25) * ((1.5 + 3.14159) - (float64(0.1) - (-2.25 - (1.5 / 1e10))))) * (1e10 + (0.0 / 0.1))))) * 3.14159)
                                                _ = v55
                                            } else {
                                                var v56 int64 = int64(v6)
                                                var v57 uint = uint(v1)
                                                _ = v56
                                                _ = v57
                                            }
                                        default:
                                            for fi58 := 0; fi58 < 4; fi58++ {
                                                v16 += -855
                                                fmt.Println(v9)
                                                v1 = v1
                                                _ = fi58
                                            }
                                        }
                                        _ = v54
                                    }
                                    if false {
                                        v59 := GenS2{F0: 2, F1: 2}
                                        v9 = true
                                        _ = v59
                                    } else {
                                        if (uint64(v4) > (v6 * v6)) {
                                            for fi60 := 0; fi60 < 10; fi60++ {
                                                v5 = (v5 ^ (v5 & (v5 | v5)))
                                                _ = fi60
                                            }
                                        } else {
                                            _ = len(v3)
                                        }
                                        fmt.Println(v4)
                                    }
                                }
                                switch 2147483647 {
                                case 128:
                                    for fi61 := 0; fi61 < 17; fi61++ {
                                        if false {
                                            v6 -= v6
                                            if v9 {
                                                var v62 string = ("" + "")
                                                v63 := GenS2{F0: 0, F1: 1}
                                                if (uint(v6) <= 18446744073709551615) {
                                                    var v64 int32 = 701
                                                    _ = v64
                                                } else {
                                                    v65 := GenS5{F0: 100}
                                                    _ = v65
                                                }
                                                _ = v62
                                                _ = v63
                                            } else {
                                                for fi66 := 0; fi66 < 7; fi66++ {
                                                    var v67 int32 = 42
                                                    _ = fi66
                                                    _ = v67
                                                }
                                            }
                                        } else {
                                            var v68 uint16 = v1
                                            _ = v68
                                        }
                                        var v69 float32 = (((0.1 * (float32((0.0 - 0.0)) * (((((-2.25 - v7) * (v7 + (1.5 / (0.0 + 0.1)))) + v7) + v7) + v7))) * v7) + ((v7 + 3.14159) / (float32(0.1) - (v7 * float32(0.1)))))
                                        v9 = v9
                                        _ = fi61
                                        _ = v69
                                    }
                                case 857:
                                    v70 := v2[4294967295]
                                    _ = v70
                                default:
                                    for fi71 := 0; fi71 < 12; fi71++ {
                                        var v72 uint64 = v6
                                        var v73 uint8 = 115
                                        _ = fi71
                                        _ = v72
                                        _ = v73
                                    }
                                }
                                _ = v49
                            }
                        } else {
                            v6 -= (v6 * (v6 & 42))
                            v5 += v5
                        }
                        _ = len(v2)
                        v74 := GenS1{F0: 1, F1: -9223372036854775808}
                        _ = v74
                    } else {
                        v9 = (true && false)
                    }
                } else {
                    var v75 uint32 = 704
                    _ = v75
                }
            }
        }
        _ = v16
    }
    v6 += (v6 << ((2) & 63))
    v6 = v6
    if (v9 && v9) {
        fmt.Println(v7)
        fmt.Println(v1, v4, v1)
    }
    v5 = v5
    v6 -= 920
    v7 -= 1.5
    var v76 int16 = (v5 - v5)
    v77 := GenS3{F0: 127}
    for fi78 := 0; fi78 < 7; fi78++ {
        v2[2147483647] = true
        _ = fi78
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v76:", v76)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v76
    _ = v77
    _ = v8
    _ = v9
}
