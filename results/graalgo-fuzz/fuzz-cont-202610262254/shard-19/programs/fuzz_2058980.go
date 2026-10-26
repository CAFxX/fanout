package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int32; F2 int8 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int32; F1 int16; F2 int32 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint }
func genFn0() int64 {
    var v4 uint16 = 255
    var v5 uint64 = 18446744073709551615
    v4 -= 913
    if false {
        v6 := []int{}
        var v7 uint32 = 271
        for fi8 := 0; fi8 < 16; fi8++ {
            if false {
                switch (fi8 ^ (fi8 - 127)) {
                case 32767:
                    v4 += v4
                case -840:
                    fmt.Println(v5, v4, fi8)
                case 9223372036854775807:
                    for fi9 := 0; fi9 < 9; fi9++ {
                        var v10 int16 = -903
                        _ = fi9
                        _ = v10
                    }
                default:
                    if ("" < "\n\t") {
                        v11 := GenS1{F0: (v4 + 32767)}
                        fmt.Println(v7, v4, v4)
                        _ = v11
                    } else {
                        v4 = uint16(v7)
                        if (true && true) {
                            _ = len(v6)
                            _ = len(v6)
                        } else {
                            if ("hello" == "a") {
                                _ = len(v6)
                            } else {
                                if true {
                                    _ = len(v6)
                                    for fi12 := 0; fi12 < 14; fi12++ {
                                        v6 = append(v6, fi12)
                                        v4 = (^(v4))
                                        v13 := GenS2{F0: v5, F1: -2147483648, F2: 1}
                                        _ = fi12
                                        _ = v13
                                    }
                                }
                                for fi14 := 0; fi14 < 5; fi14++ {
                                    fmt.Println(v7, fi14)
                                    if (!((v4 * v4) < v4)) {
                                        switch (fi8 - fi14) {
                                        case 301:
                                            switch -1 {
                                            case 42:
                                                var v15 uint16 = 0
                                                _ = v15
                                            default:
                                                _ = len(v6)
                                            }
                                        case 704:
                                            switch 2147483647 {
                                            case 4294967295:
                                                var v16 uint16 = v4
                                                _ = v16
                                            default:
                                                fmt.Println(v5)
                                            }
                                        default:
                                            var v17 uint16 = 513
                                            _ = v17
                                        }
                                    }
                                    v5 = v5
                                    _ = fi14
                                }
                            }
                        }
                        if true {
                            var v18 uint8 = 39
                            var v19 bool = (-43 == 100)
                            var v20 uint16 = v4
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    }
                }
                if false {
                    if (18446744073709551615 >= 255) {
                        v5 = 437
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    if (226 == 2) {
                        fmt.Println(v5)
                    } else {
                        if (!true) {
                            var v21 int32 = -1
                            for fi22 := 0; fi22 < 12; fi22++ {
                                var v23 bool = (0.0 <= ((0.0 - float64((-2.25 * 0.0))) + float64(1.5)))
                                fmt.Println(v21, v5, v4)
                                var v24 int8 = 1
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            _ = v21
                        } else {
                            v5 = 2
                            _ = len(v6)
                            var v25 int8 = 42
                            _ = v25
                        }
                        switch 42 {
                        case -460:
                            switch fi8 {
                            case 707:
                                if len(v6) > 0 {
                                    v6[0] = (fi8 & 2147483647)
                                }
                            default:
                                for fi26 := 0; fi26 < 20; fi26++ {
                                    v27 := GenS2{F0: 0, F1: 2147483647, F2: int8(fi8)}
                                    var v28 uint = 42
                                    if (((((((float32((0.0 + 0.0)) * -2.25) - 1.5) / (0.0 - 0.1)) * 1.5) / -2.25) >= -2.25) || (("" + "hello") >= "")) {
                                        if len(v6) > 0 {
                                            v6[0] = fi8
                                        }
                                    } else {
                                        var v29 uint = v28
                                        if (false || ((-(fi26)) == fi26)) {
                                            v5 -= 42
                                            var v30 int32 = -133
                                            _ = v30
                                        }
                                        v31 := GenS1{F0: (^(v4))}
                                        switch (-(fi26)) {
                                        case 9223372036854775807:
                                            v29 = (v28 % ((uint(v7)) | 1))
                                        case -632:
                                            v4 = v4
                                        }
                                        _ = v29
                                        _ = v31
                                    }
                                    v32 := GenS1{F0: uint16(v7)}
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                    _ = v32
                                }
                            }
                        case -797:
                            var v33 uint32 = (^(v7))
                            _ = v33
                        case 41:
                            var v34 string = (("a" + "foo bar") + "")
                            _ = v34
                        }
                        v6 = append(v6, (fi8 >> ((uint(v5)) & 63)))
                    }
                } else {
                    v35 := GenS3{F0: ("" + ("hello" + ("" + "hello"))), F1: 1}
                    var v36 int8 = 1
                    _ = v35
                    _ = v36
                }
            }
            v37 := GenS4{F0: 42, F1: 170}
            _ = fi8
            _ = v37
        }
        _ = v6
        _ = v7
    }
    v38 := GenS5{F0: 766, F1: -23, F2: 288}
    _ = v38
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint64) bool {
    for fi39 := 0; fi39 < 7; fi39++ {
        v40 := []int{(fi39 - -2147483648), (fi39 - fi39)}
        _ = fi39
        _ = v40
    }
    v41 := []int{-1, 758, 42, 32767}
    p0 += p0
    _ = p0
    _ = v41
    return (-2.25 < (1e10 + 1e10))
}

func main() {
    var v1 int32 = 1
    v2 := []string{"", "", "hello"}
    var v3 int16 = -32768
    switch -32 {
    case 127:
        for fi42 := 0; fi42 < 8; fi42++ {
            v43 := GenS4{F0: -128, F1: 42}
            _ = fi42
            _ = v43
        }
    case 353:
        v3 += v3
    default:
        v1 = (v1 << ((42) & 63))
    }
    v44 := GenS6{F0: (!(true && (913 != 2147483647))), F1: 4294967295, F2: 2}
    v1 = -2147483648
    var v45 int8 = -128
    v1 -= (v1 / ((v1) | 1))
    var v46 uint8 = 253
    var v47 int16 = v3
    v45 = v45
    var v48 float64 = 1.5
    if (2 == 42) {
        v47 += v47
    }
    var v49 int8 = v45
    v50 := GenS6{F0: true, F1: 4294967295, F2: 65535}
    switch -2147483648 {
    case 42:
        v48 += 1e10
    default:
        v48 += 1.5
    }
    if (!(v49 == (^(v49)))) {
        var v51 float64 = (v48 - v48)
        _ = v51
    }
    var v52 int16 = v47
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v52
}
