package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 int; F2 int32 }
func genFn0(p0 int16) int {
    if false {
        v7 := GenS1{F0: 2, F1: p0, F2: 619}
        var v8 bool = (19 == -128)
        _ = v7
        _ = v8
    } else {
        if true {
            v9 := GenS1{F0: 88, F1: p0, F2: -9223372036854775808}
            v10 := GenS1{F0: 1, F1: (^(p0)), F2: -29}
            _ = v10
            _ = v9
        } else {
            var v11 int8 = -1
            for fi12 := 0; fi12 < 6; fi12++ {
                v11 += (v11 << ((65535) & 63))
                _ = fi12
            }
            _ = v11
        }
    }
    v13 := []int{-706, -822, -444}
    var v14 bool = (!false)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    _ = p0
    _ = v13
    _ = v14
    return 9223372036854775807
}
func genFn1() int16 {
    var v15 uint16 = 2
    var v16 uint32 = 2147483648
    var v17 float64 = 3.14159
    switch ((("" + "a") + "foo bar") + "hello") {
    case "\n\t":
        v18 := GenS1{F0: 2, F1: -364, F2: 0}
        _ = v18
    default:
        v19 := GenS2{F0: 128}
        _ = v19
    }
    v20 := GenS2{F0: 152}
    _ = v15
    _ = v16
    _ = v17
    _ = v20
    return -754
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{42, 2147483647, 2, 1}
    var v3 uint16 = 230
    v4 := []int{-316, 128}
    var v5 uint64 = 42
    var v6 uint64 = 1
    var v21 uint32 = 283
    v22 := GenS3{F0: (true && true), F1: true, F2: v3}
    v3 += (v3 - v3)
    v1 = "hello"
    v4 = append(v4, 127)
    var v23 uint16 = 0
    if true {
        var v24 string = "hello"
        if (-536 < 596) {
            var v25 uint64 = (v5 + (v5 ^ v6))
            fmt.Println(v3)
            if (!false) {
                fmt.Println(v5, v5)
                if (true && false) {
                    v1 = v24
                    var v26 bool = (42 < genFn0(1000))
                    _ = v26
                }
                var v27 string = v1
                _ = v27
            }
            _ = v25
        }
        fmt.Println(v1)
        _ = v24
    } else {
        v23 += (v23 | 42)
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 uint16 = (v23 + v23)
        _ = fi28
        _ = v29
    }
    v21 -= (v21 & (v21 * uint32(v23)))
    v3 -= (v23 / (((^(v23))) | 1))
    v30 := GenS3{F0: true, F1: ((((^(v3)) <= 1) || (2147483647 < 4294967295)) || false), F2: 2}
    v3 += (v3 + (v23 - 554))
    if (v3 != v3) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        var v31 int64 = -41
        if false {
            for fi32 := 0; fi32 < 4; fi32++ {
                if ((v23 * v23) == (v3 % (((v3 | (v3 % (((v3 / (((^(v23))) | 1))) | 1)))) | 1))) {
                    v23 += 42
                }
                if true {
                    if (1.5 == (1e10 + (float64((1e10 * 0.0)) * (1e10 * (0.0 - 0.1))))) {
                        fmt.Println(v21, v1, v23)
                        switch v1 {
                        case "a":
                            if (42 >= 1) {
                                var v33 bool = false
                                _ = v33
                            }
                        case "foo bar":
                            for fi34 := 0; fi34 < 20; fi34++ {
                                v6 -= 2
                                v5 = (v5 << ((381) & 63))
                                switch 4294967295 {
                                case 127:
                                    var v35 uint8 = 116
                                    _ = v35
                                }
                                _ = fi34
                            }
                        case "\n\t":
                            switch (v1 + v1) {
                            case "\n\t":
                                var v36 int32 = 42
                                _ = v36
                            case "foo bar":
                                v5 -= (v6 + uint64(v3))
                            }
                        default:
                            v6 = (^(v5))
                        }
                        var v37 int64 = 9223372036854775807
                        _ = v37
                    }
                    if (false && true) {
                        if true {
                            switch (v1 + v1) {
                            case "hello":
                                var v38 int64 = (v31 >> ((26) & 63))
                                _ = v38
                            case "a":
                                if ((v21 << ((466) & 63)) >= (^(v21))) {
                                    if false {
                                        v1 = v1
                                        v2 = append(v2, 778)
                                        v6 = (v6 + v6)
                                    }
                                } else {
                                    for fi39 := 0; fi39 < 10; fi39++ {
                                        fmt.Println(v1, v5, v31)
                                        switch v1 {
                                        case "hello":
                                            var v40 float64 = 3.14159
                                            _ = v40
                                        default:
                                            v6 = 0
                                        }
                                        _ = fi39
                                    }
                                }
                            default:
                                var v41 uint32 = 475
                                _ = v41
                            }
                            for fi42 := 0; fi42 < 20; fi42++ {
                                var v43 int8 = 0
                                _ = fi42
                                _ = v43
                            }
                        } else {
                            var v44 float64 = 3.14159
                            if true {
                                if len(v2) > 0 {
                                    v2[0] = (v21 % ((uint32(v6)) | 1))
                                }
                                v3 = 1
                            } else {
                                v45 := GenS1{F0: uint8(v23), F1: 52, F2: (v31 & (-(v31)))}
                                _ = v45
                            }
                            _ = v44
                        }
                        if (!true) {
                            var v46 uint8 = 1
                            v47 := GenS4{F0: (v6 ^ (v6 & (v5 << ((255) & 63))))}
                            _ = v46
                            _ = v47
                        } else {
                            if (fi32 < (fi32 - (fi32 % ((fi32) | 1)))) {
                                var v48 int32 = -2147483648
                                _ = v48
                            } else {
                                var v49 uint16 = uint16(v5)
                                var v50 uint32 = (^(v21))
                                if len(v2) > 0 {
                                    v2[0] = uint32(v6)
                                }
                                _ = v49
                                _ = v50
                            }
                            switch v1 {
                            case "a":
                                if len(v4) > 0 {
                                    v4[0] = int(v21)
                                }
                            case "foo bar":
                                v3 = v3
                            }
                        }
                        fmt.Println(v5, v23)
                    }
                    v1 = "hello"
                }
                if false {
                    v6 += v5
                } else {
                    v3 = (v23 ^ v23)
                    v6 = (^(v6))
                    v51 := GenS5{F0: 42, F1: (^(fi32)), F2: 0}
                    _ = v51
                }
                _ = fi32
            }
            var v52 int32 = -1
            _ = v52
        } else {
            switch "a" {
            case "":
                v5 += 42
            case "hello":
                v1 = ""
            case "\n\t":
                if ((0.1 / float32((0.1 - 1.5))) < (0.0 + (1.5 * 1.5))) {
                    for fi53 := 0; fi53 < 18; fi53++ {
                        var v54 float64 = float64(((1e10 / 3.14159) - (-2.25 - 1.5)))
                        v31 += v31
                        _ = fi53
                        _ = v54
                    }
                    v6 = 285
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            }
            var v55 uint32 = v21
            _ = v55
        }
        for fi56 := 0; fi56 < 10; fi56++ {
            for fi57 := 0; fi57 < 14; fi57++ {
                v4 = append(v4, (fi57 << ((4294967295) & 63)))
                var v58 int16 = 1
                _ = fi57
                _ = v58
            }
            v59 := GenS1{F0: 174, F1: int16(fi56), F2: v31}
            _ = fi56
            _ = v59
        }
        _ = v31
    }
    var v60 int64 = -779
    v60 -= (v60 - v60)
    var v61 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
