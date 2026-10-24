package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint }

func main() {
    var v1 string = (("a" + ("\n\t" + "hello")) + "a")
    var v2 int = 32768
    var v3 int8 = -1
    var v4 string = (v1 + ("\n\t" + "hello"))
    if false {
        v5 := []int{v2, -747, v2, (v2 - (v2 + v2))}
        _ = v5
    }
    fmt.Println(v1)
    v6 := GenS1{F0: 42, F1: false}
    v7 := []int{(v2 - v2), v2}
    var v8 uint8 = 54
    v2 += (v2 / ((v2) | 1))
    _ = len(v7)
    if false {
        var v9 float32 = ((1.5 * float32(1.5)) / 1.0)
        _ = v9
    } else {
        if (v3 == (v3 + v3)) {
            var v10 uint64 = 4294967295
            if (2 < uint16(v10)) {
                var v11 int16 = -1
                var v12 uint8 = 2
                _ = v11
                _ = v12
            }
            _ = v10
        }
        v13 := GenS1{F0: 0, F1: true}
        _ = v13
    }
    v14 := GenS1{F0: 1, F1: true}
    for fi15 := 0; fi15 < 14; fi15++ {
        switch "foo bar" {
        case "foo bar":
            v1 = (v4 + v4)
        case "hello":
            v3 = 0
        }
        switch fi15 {
        case -2147483648:
            fmt.Println(v3)
        case 538:
            switch v4 {
            case "hello":
                var v16 bool = false
                _ = v16
            default:
                switch (v1 + v1) {
                case "\n\t":
                    if (true && ((v8 % ((v8) | 1)) < (v8 << ((42) & 63)))) {
                        if len(v7) > 0 {
                            v7[0] = (fi15 - (v2 | 0))
                        }
                        v2 += v2
                        if len(v7) > 0 {
                            v7[0] = fi15
                        }
                    }
                case "hello":
                    for fi17 := 0; fi17 < 11; fi17++ {
                        v3 += (v3 ^ 77)
                        if len(v7) > 0 {
                            v7[0] = 1
                        }
                        _ = fi17
                    }
                case "":
                    switch (^(fi15)) {
                    case -746:
                        for fi18 := 0; fi18 < 19; fi18++ {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = fi18
                        }
                    default:
                        switch (-(fi15)) {
                        case 127:
                            if (("" + "foo bar") == "") {
                                v19 := GenS1{F0: 2, F1: true}
                                _ = v19
                            } else {
                                for fi20 := 0; fi20 < 12; fi20++ {
                                    var v21 float64 = (float64(1.5) / 1.5)
                                    _ = fi20
                                    _ = v21
                                }
                            }
                        case 42:
                            v4 = v4
                        }
                    }
                }
            }
        case 127:
            if true {
                v22 := GenS2{F0: v8, F1: v8}
                var v23 bool = ((!true) || false)
                if len(v7) > 0 {
                    v7[0] = (^(v2))
                }
                _ = v22
                _ = v23
            } else {
                if false {
                    v24 := GenS2{F0: v8, F1: (v8 / ((127) | 1))}
                    _ = v24
                } else {
                    switch v2 {
                    case -34:
                        var v25 uint8 = (^(v8))
                        _ = v25
                    case 32768:
                        if (true || true) {
                            if len(v7) > 0 {
                                v7[0] = v2
                            }
                            v1 = (v1 + (v1 + (v4 + "foo bar")))
                            fmt.Println(fi15, v8, v4)
                        } else {
                            for fi26 := 0; fi26 < 10; fi26++ {
                                v2 -= 0
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                var v27 uint = 4294967295
                                _ = fi26
                                _ = v27
                            }
                            if (!(((v3 * 100) >= v3) && true)) {
                                for fi28 := 0; fi28 < 11; fi28++ {
                                    if false {
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        v29 := GenS1{F0: 42, F1: true}
                                        v4 = ("\n\t" + "")
                                        _ = v29
                                    } else {
                                        fmt.Println(v3, fi15)
                                        if false {
                                            v7 = append(v7, (fi15 | int(v3)))
                                            switch (v4 + "a") {
                                            case "":
                                                if (!false) {
                                                    var v30 uint = 18446744073709551615
                                                    v7 = append(v7, v2)
                                                    _ = v30
                                                } else {
                                                    v31 := GenS3{F0: v8, F1: 0}
                                                    _ = v31
                                                }
                                            case "a":
                                                v3 = -1
                                            default:
                                                v2 += 0
                                            }
                                            v32 := GenS4{F0: 494}
                                            _ = v32
                                        } else {
                                            v4 = "\n\t"
                                        }
                                    }
                                    var v33 uint8 = 42
                                    if ((fi28 << ((565) & 63)) < fi28) {
                                        v7 = append(v7, (v2 * (fi15 / ((2147483647) | 1))))
                                        var v34 float64 = 0.1
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        _ = v34
                                    } else {
                                        var v35 int64 = -9223372036854775808
                                        _ = v35
                                    }
                                    _ = fi28
                                    _ = v33
                                }
                                if true {
                                    v36 := GenS1{F0: 0, F1: (989 < 42)}
                                    var v37 int32 = 42
                                    if (405 >= 18446744073709551615) {
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        v38 := GenS5{F0: 42, F1: 2, F2: 18446744073709551615}
                                        _ = v38
                                    }
                                    _ = v36
                                    _ = v37
                                } else {
                                    for fi39 := 0; fi39 < 12; fi39++ {
                                        _ = len(v7)
                                        _ = fi39
                                    }
                                }
                            } else {
                                fmt.Println(fi15)
                                v3 += (v3 % (((-(v3))) | 1))
                                v1 = v4
                            }
                            var v40 uint8 = (v8 * (v8 % ((42) | 1)))
                            _ = v40
                        }
                    default:
                        if (((0.0 + 0.0) - -2.25) == (0.0 - 0.1)) {
                            fmt.Println(v2)
                        } else {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                    }
                }
                var v41 int32 = -2147483648
                _ = v41
            }
        }
        for fi42 := 0; fi42 < 11; fi42++ {
            if false {
                switch fi15 {
                case 722:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                case -9223372036854775808:
                    var v43 float32 = 1.5
                    _ = v43
                case 127:
                    var v44 int8 = v3
                    _ = v44
                default:
                    var v45 uint16 = 867
                    _ = v45
                }
                if (true || false) {
                    v3 = (v3 + (-(v3)))
                } else {
                    var v46 float32 = 1.5
                    if (false || (false && (2 != 65535))) {
                        var v47 int = fi15
                        var v48 uint16 = 557
                        v49 := GenS3{F0: v8, F1: 572}
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                    _ = v46
                }
            } else {
                v50 := GenS5{F0: 516, F1: 65535, F2: uint(v8)}
                switch "\n\t" {
                case "\n\t":
                    var v51 float32 = 3.14159
                    _ = v51
                case "":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                default:
                    var v52 uint32 = 995
                    _ = v52
                }
                _ = v50
            }
            var v53 uint8 = v8
            _ = fi42
            _ = v53
        }
        _ = fi15
    }
    var v54 uint16 = 1
    v54 += 2
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v6
    _ = v7
    _ = v8
}
