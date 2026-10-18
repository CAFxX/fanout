package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int32) bool {
    fmt.Println(p0)
    var v7 bool = true
    switch "foo bar" {
    case "":
        v7 = v7
    default:
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(fi8, v7)
            v9 := GenS1{F0: fi8, F1: (!(v7 && true)), F2: 9223372036854775807}
            _ = fi8
            _ = v9
        }
    }
    fmt.Println(v7, p0)
    _ = p0
    _ = v7
    return true
}
func genFn1(p0 int8, p1 int8, p2 bool) uint32 {
    if (false && p2) {
        v10 := GenS1{F0: -589, F1: p2, F2: 851}
        var v11 int8 = (p1 - 0)
        _ = v10
        _ = v11
    }
    fmt.Println(p0, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint8 = 142
    var v2 float64 = 1e10
    var v3 uint32 = 2147483647
    v4 := []string{"hello", "\n\t", "a", "a"}
    var v5 int32 = 440
    v6 := []int{}
    v12 := GenS1{F0: 0, F1: (false || (false && (-1 <= -1))), F2: 0}
    var v13 uint64 = 18446744073709551615
    if (((v13 - 42) >= 0) && true) {
        if false {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int64 = 42
                v16 := GenS1{F0: 127, F1: (!(!true)), F2: 996}
                if true {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v18 := GenS1{F0: fi17, F1: (!(!false)), F2: v15}
                        if (104 >= 2) {
                            var v19 float32 = 0.1
                            if (137 == 32767) {
                                var v20 float64 = v2
                                var v21 float64 = (v2 + v2)
                                _ = v20
                                _ = v21
                            } else {
                                var v22 string = "hello"
                                fmt.Println(v15, v2, v15)
                                switch v22 {
                                case "hello":
                                    fmt.Println(v5, v2, v5)
                                    fmt.Println(v1, v5)
                                    v6 = append(v6, (-(fi14)))
                                case "":
                                    if true {
                                        var v23 uint8 = 255
                                        _ = v23
                                    }
                                case "a":
                                    if false {
                                        v24 := GenS1{F0: fi14, F1: ((v3 & (v3 | (v3 % ((v3) | 1)))) < v3), F2: (^(v15))}
                                        _ = v24
                                    } else {
                                        v25 := GenS1{F0: fi17, F1: (v2 <= 1.5), F2: v15}
                                        if true {
                                            fmt.Println(v13)
                                        } else {
                                            if (836 != 0) {
                                                v3 = (v3 & uint32(v13))
                                                var v26 uint8 = (^(v1))
                                                v27 := GenS1{F0: fi17, F1: false, F2: v15}
                                                v3 = (v3 << ((65535) & 63))
                                                fmt.Println(v22, v22)
                                                _ = v26
                                                _ = v27
                                            } else {
                                                var v28 int8 = -84
                                                if ("a" <= v22) {
                                                    var v29 string = "foo bar"
                                                    _ = v29
                                                }
                                                v30 := GenS1{F0: (fi17 + int(v3)), F1: (false && false), F2: (^(v15))}
                                                if true {
                                                    var v31 uint = 255
                                                    _ = v31
                                                } else {
                                                    var v32 float32 = v19
                                                    if ((v13 & 854) < v13) {
                                                        v28 = v28
                                                        var v33 uint16 = 1
                                                        if (v1 > uint8(v13)) {
                                                            v4 = append(v4, "")
                                                            v32 = (v19 / 3.14159)
                                                        } else {
                                                            v2 -= v2
                                                        }
                                                        _ = v33
                                                    }
                                                    _ = v32
                                                }
                                                _ = v28
                                                _ = v30
                                            }
                                            v19 = v19
                                        }
                                        _ = v25
                                    }
                                }
                                _ = v22
                            }
                            _ = v19
                        } else {
                            var v34 int16 = -704
                            v1 -= 155
                            _ = v34
                        }
                        v35 := GenS2{F0: "\n\t", F1: v1}
                        var v36 bool = true
                        _ = fi17
                        _ = v18
                        _ = v35
                        _ = v36
                    }
                    switch 9223372036854775807 {
                    case 127:
                        _ = len(v4)
                    case -715:
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v1)
                        }
                    }
                } else {
                    v4 = append(v4, "foo bar")
                    for fi37 := 0; fi37 < 17; fi37++ {
                        if true {
                            if len(v4) > 0 {
                                v4[0] = (((("a" + "foo bar") + ("" + ("" + ("foo bar" + ("\n\t" + "\n\t"))))) + ("hello" + "")) + "hello")
                            }
                            switch 42 {
                            case 236:
                                var v38 uint8 = (v1 / ((v1) | 1))
                                v2 -= (((1e10 * (float64((-2.25 + 0.0)) - (1e10 / float64(-2.25)))) - ((v2 + v2) + v2)) - float64(0.0))
                                if len(v6) > 0 {
                                    v6[0] = (fi14 ^ fi14)
                                }
                                fmt.Println(v3, v2)
                                var v39 float64 = -2.25
                                _ = v38
                                _ = v39
                            }
                            switch (fi14 << ((2) & 63)) {
                            case -2147483648:
                                fmt.Println(v5, v13)
                            case -292:
                                if false {
                                    v15 = (v15 + -9223372036854775808)
                                    if ((v5 & v5) <= -836) {
                                        if len(v4) > 0 {
                                            v4[0] = "hello"
                                        }
                                        var v40 string = ("\n\t" + ("hello" + "foo bar"))
                                        _ = v40
                                    }
                                }
                            case -334:
                                if false {
                                    v1 = (v1 & uint8(v13))
                                    _ = len(v4)
                                    if false {
                                        v5 = (v5 % (((v5 | 42)) | 1))
                                    } else {
                                        var v41 int32 = v5
                                        _ = v41
                                    }
                                } else {
                                    v15 += -270
                                    fmt.Println(v3, v2)
                                }
                            }
                        } else {
                            var v42 int = fi37
                            if false {
                                v43 := GenS3{F0: 250}
                                _ = v43
                            }
                            _ = v42
                        }
                        var v44 int64 = v15
                        _ = fi37
                        _ = v44
                    }
                    fmt.Println(v5, v2, v5)
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            if (!true) {
                for fi45 := 0; fi45 < 3; fi45++ {
                    for fi46 := 0; fi46 < 12; fi46++ {
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                        _ = fi46
                    }
                    v2 -= v2
                    v47 := GenS4{F0: 255}
                    _ = fi45
                    _ = v47
                }
            }
            for fi48 := 0; fi48 < 1; fi48++ {
                v2 -= v2
                _ = fi48
            }
        }
        var v49 uint8 = v1
        _ = v49
    }
    v2 += 3.14159
    var v50 string = ("" + "\n\t")
    var v51 float64 = (((1e10 * 1e10) / v2) - ((1.5 - v2) / (v2 - v2)))
    switch 128 {
    case 128:
        if false {
            fmt.Println(v50, v1, v50)
            v51 += v2
        } else {
            var v52 uint = 255
            fmt.Println(v1)
            _ = v52
        }
    case -9223372036854775808:
        for fi53 := 0; fi53 < 5; fi53++ {
            v13 = v13
            _ = fi53
        }
    case 813:
        var v54 int64 = 9223372036854775807
        _ = v54
    default:
        var v55 uint16 = 1
        _ = v55
    }
    v4 = append(v4, ("\n\t" + ("foo bar" + v50)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
