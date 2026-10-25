package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32; F1 uint; F2 bool }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint8; F1 string; F2 uint32 }
type GenS5 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint64, p1 uint8, p2 int) int16 {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint8 = (p1 | (p1 - p1))
        v8 := []int{(p2 - (fi6 + fi6)), -285}
        var v9 int = 1
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint16 = uint16(p1)
    v11 := []int{32768, p2, 2147483647}
    switch "foo bar" {
    case "\n\t":
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 int = fi13
                switch (p2 >> ((4294967295) & 63)) {
                case 9223372036854775807:
                    var v15 int8 = -1
                    _ = v15
                case 0:
                    v14 = (v14 % ((fi13) | 1))
                default:
                    var v16 uint16 = v10
                    _ = v16
                }
                if ((42 == int64(p0)) || true) {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    if len(v11) > 0 {
                        v11[0] = (fi12 * 2147483647)
                    }
                }
                p0 -= p0
                var v17 int = 35
                v10 -= (v10 & (v10 * v10))
                _ = fi13
                _ = v14
                _ = v17
            }
            p0 += p0
            var v18 int64 = -9223372036854775808
            _ = fi12
            _ = v18
        }
    case "hello":
        if ((!true) || ((p1 << ((751) & 63)) <= (^(p1)))) {
            var v19 int8 = 127
            if true {
                switch "hello" {
                case "a":
                    v10 = v10
                default:
                    v20 := GenS1{F0: (p1 | 128), F1: true}
                    _ = v20
                }
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                for fi21 := 0; fi21 < 16; fi21++ {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        if ((!true) && (-550 <= -32768)) {
                            var v23 int32 = 0
                            _ = v23
                        } else {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            var v24 string = "hello"
                            v24 = v24
                            _ = v24
                        }
                        _ = fi22
                    }
                    var v25 uint8 = p1
                    if (9223372036854775807 == 42) {
                        if len(v11) > 0 {
                            v11[0] = p2
                        }
                        var v26 float32 = (0.1 - ((3.14159 * 1e10) * 0.1))
                        _ = v26
                    }
                    _ = fi21
                    _ = v25
                }
            } else {
                _ = len(v11)
            }
            v27 := GenS2{F0: 653, F1: 18446744073709551615, F2: true}
            _ = v19
            _ = v27
        } else {
            var v28 string = ("a" + (("foo bar" + "") + "hello"))
            for fi29 := 0; fi29 < 17; fi29++ {
                for fi30 := 0; fi30 < 4; fi30++ {
                    v31 := GenS1{F0: 2, F1: false}
                    _ = fi30
                    _ = v31
                }
                v32 := GenS1{F0: p1, F1: false}
                _ = fi29
                _ = v32
            }
            var v33 uint8 = 255
            _ = v28
            _ = v33
        }
    case "foo bar":
        p1 -= (p1 + 0)
    default:
        if ((("hello" + "\n\t") + (("foo bar" + "\n\t") + "foo bar")) < "") {
            fmt.Println(p1, p1)
        } else {
            switch p2 {
            case -9223372036854775808:
                fmt.Println(v10, p0)
            case 128:
                for fi34 := 0; fi34 < 19; fi34++ {
                    p2 = p2
                    var v35 uint8 = (p1 % (((p1 - (p1 + 42))) | 1))
                    _ = fi34
                    _ = v35
                }
            default:
                if true {
                    p1 -= p1
                    for fi36 := 0; fi36 < 19; fi36++ {
                        v10 -= v10
                        p1 = p1
                        _ = fi36
                    }
                    fmt.Println(p0)
                } else {
                    var v37 float32 = ((1.5 / 1.5) + 0.1)
                    switch p2 {
                    case -2147483648:
                        for fi38 := 0; fi38 < 3; fi38++ {
                            switch ("" + "hello") {
                            case "hello":
                                var v39 int32 = -654
                                _ = v39
                            case "\n\t":
                                v40 := GenS1{F0: (^(p1)), F1: (v37 <= v37)}
                                _ = v40
                            case "foo bar":
                                v41 := GenS1{F0: p1, F1: (!false)}
                                _ = v41
                            }
                            p0 = (p0 & p0)
                            if (0.1 != 0.1) {
                                switch p2 {
                                case 67:
                                    p1 -= 128
                                case 42:
                                    v42 := GenS2{F0: 1, F1: 1, F2: ((1 <= 860) || (!false))}
                                    _ = v42
                                case 0:
                                    var v43 string = ("" + "hello")
                                    _ = v43
                                }
                                v10 += 255
                                fmt.Println(p2, v10)
                            }
                            _ = fi38
                        }
                    case 18:
                        fmt.Println(v37, v10, p2)
                    }
                    var v44 bool = true
                    _ = v37
                    _ = v44
                }
            }
            p0 += (p0 << ((694) & 63))
            for fi45 := 0; fi45 < 12; fi45++ {
                for fi46 := 0; fi46 < 18; fi46++ {
                    p2 += p2
                    fmt.Println(p2)
                    _ = fi46
                }
                for fi47 := 0; fi47 < 8; fi47++ {
                    v48 := GenS3{F0: -1, F1: 292}
                    v49 := GenS2{F0: 42, F1: 751, F2: ((p1 | 69) < (p1 * p1))}
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                if len(v11) > 0 {
                    v11[0] = (p2 % ((fi45) | 1))
                }
                _ = fi45
            }
        }
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint = 2
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    v5 := []string{("\n\t" + ("" + ("\n\t" + "\n\t"))), "\n\t"}
    v50 := GenS4{F0: 128, F1: "a", F2: 215}
    var v51 int32 = -1
    if len(v5) > 0 {
        v5[0] = "a"
    }
    if true {
        var v52 int8 = -18
        v51 = (v1 | (-(v1)))
        _ = v52
    }
    var v53 float64 = 1.5
    var v54 float64 = 1.5
    var v55 string = "\n\t"
    v53 += float64(0.1)
    if (42 <= 236) {
        fmt.Println(v2)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        if true {
            if false {
                _ = len(v4)
                v3[4294967295] = 2147483648
                fmt.Println(v2, v53)
            } else {
                for fi56 := 0; fi56 < 2; fi56++ {
                    var v57 uint8 = 0
                    _ = fi56
                    _ = v57
                }
                fmt.Println(v1, v2, v1)
            }
        }
        if (42 != 67) {
            v1 = (v1 % ((-2147483648) | 1))
        }
    }
    for fi58 := 0; fi58 < 16; fi58++ {
        v54 = 3.14159
        if len(v5) > 0 {
            v5[0] = (v55 + "foo bar")
        }
        _ = fi58
    }
    var v59 float64 = (-2.25 * float64(0.0))
    var v60 int = 32768
    v61 := GenS5{F0: 1, F1: 536}
    fmt.Println(v59)
    for fi62 := 0; fi62 < 6; fi62++ {
        var v63 int8 = -1
        var v64 uint = (^(v2))
        if false {
            switch v55 {
            case "hello":
                var v65 int8 = v63
                _ = v65
            }
            v1 -= v51
            v60 = fi62
        } else {
            v5 = append(v5, "\n\t")
        }
        _ = fi62
        _ = v63
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v53
    _ = v54
    _ = v55
    _ = v59
    _ = v60
    _ = v61
}
