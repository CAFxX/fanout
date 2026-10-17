package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int; F1 int32; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 string; F1 uint32; F2 int8 }
func genFn0(p0 int) int32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(p0, fi7)
        _ = fi7
    }
    var v8 uint8 = 1
    v8 = (v8 | 127)
    p0 += (p0 % ((p0) | 1))
    _ = p0
    _ = v8
    return 42
}
func genFn1() uint {
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 float32 = ((-2.25 * ((0.1 - ((0.1 - -2.25) * 1e10)) / 1.0)) - 1.5)
        for fi11 := 0; fi11 < 9; fi11++ {
            v10 += -2.25
            var v12 string = "a"
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 int32 = 1
    var v14 string = "\n\t"
    v15 := GenS1{F0: 186}
    v14 = "\n\t"
    _ = v13
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 float32 = 3.14159
    var v2 string = "hello"
    var v3 int64 = 9223372036854775807
    var v4 uint8 = 239
    var v5 bool = (-9223372036854775808 < 2147483647)
    var v6 uint16 = 0
    if ((v4 != v4) && v5) {
        v16 := []int{4294967295, 32768, -1, 4294967295}
        v17 := GenS1{F0: (v6 << ((42) & 63))}
        _ = v16
        _ = v17
    }
    v18 := []int{127, 1, -939}
    fmt.Println(v6)
    for fi19 := 0; fi19 < 1; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 int = 21
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        if (-128 >= 1) {
            v2 = v2
            for fi23 := 0; fi23 < 15; fi23++ {
                v1 = v1
                switch v2 {
                case "foo bar":
                    fmt.Println(v5, v1)
                case "a":
                    var v24 bool = v5
                    var v25 int = (fi22 + 1)
                    _ = v24
                    _ = v25
                default:
                    if (v4 < v4) {
                        switch ("\n\t" + "") {
                        case "\n\t":
                            if len(v18) > 0 {
                                v18[0] = (fi22 / ((32768) | 1))
                            }
                        default:
                            switch (v2 + ("a" + v2)) {
                            case "\n\t":
                                v18 = append(v18, (fi23 << ((18446744073709551615) & 63)))
                            }
                        }
                        var v26 int64 = (v3 + v3)
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        _ = v26
                    }
                }
                _ = fi23
            }
            var v27 uint16 = v6
            _ = v27
        } else {
            var v28 int8 = -128
            var v29 float32 = float32((((1e10 - 1e10) - 0.0) + 1e10))
            v5 = (((v2 + "\n\t") + "foo bar") <= ((v2 + v2) + "\n\t"))
            _ = v28
            _ = v29
        }
        var v30 uint = 166
        _ = fi22
        _ = v30
    }
    v4 = v4
    fmt.Println(v1)
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := GenS1{F0: (v6 << ((4294967295) & 63))}
        fmt.Println(v2, v2)
        v18 = append(v18, fi31)
        _ = fi31
        _ = v32
    }
    if (0 < 1) {
        v5 = true
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        if true {
            switch (^(fi33)) {
            case 0:
                v18 = append(v18, int(v4))
            case 4294967295:
                var v34 int16 = 1
                _ = v34
            case 1:
                for fi35 := 0; fi35 < 5; fi35++ {
                    v36 := GenS1{F0: (v6 | 32767)}
                    fmt.Println(v6, v1, fi35)
                    v37 := GenS2{F0: 0, F1: v2}
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
            v38 := GenS1{F0: 1}
            var v39 float32 = (v1 * ((v1 - 0.0) + v1))
            _ = v38
            _ = v39
        }
        if (793 > 2147483648) {
            v3 = (v3 << ((678) & 63))
            if true {
                v5 = v5
            } else {
                if (!v5) {
                    if (670 <= -1) {
                        for fi40 := 0; fi40 < 19; fi40++ {
                            v41 := GenS2{F0: -128, F1: v2}
                            _ = fi40
                            _ = v41
                        }
                    } else {
                        if (v5 && v5) {
                            v42 := GenS3{F0: 824}
                            v5 = true
                            v43 := GenS4{F0: false, F1: 4294967295, F2: v6}
                            _ = v42
                            _ = v43
                        } else {
                            var v44 int64 = (v3 % ((v3) | 1))
                            switch 32768 {
                            case -907:
                                if ((!(true && (1 != v44))) || true) {
                                    for fi45 := 0; fi45 < 16; fi45++ {
                                        v46 := GenS5{F0: -124, F1: 0, F2: 46}
                                        var v47 uint64 = 4294967295
                                        _ = fi45
                                        _ = v46
                                        _ = v47
                                    }
                                    v48 := GenS6{F0: genFn1()}
                                    _ = v48
                                }
                            case 0:
                                switch ("hello" + v2) {
                                case "foo bar":
                                    fmt.Println(v2, v44)
                                case "\n\t":
                                    if (!v5) {
                                        if (-11 > 0) {
                                            switch (^(fi33)) {
                                            case -707:
                                                v49 := GenS6{F0: 65535}
                                                _ = v49
                                            case 32768:
                                                if len(v18) > 0 {
                                                    v18[0] = fi33
                                                }
                                            case -1:
                                                v50 := GenS6{F0: uint(v4)}
                                                _ = v50
                                            }
                                            if ((v2 + v2) >= "") {
                                                if len(v18) > 0 {
                                                    v18[0] = 127
                                                }
                                            }
                                        }
                                        for fi51 := 0; fi51 < 4; fi51++ {
                                            var v52 string = ("\n\t" + ("hello" + v2))
                                            var v53 int64 = v3
                                            _ = fi51
                                            _ = v52
                                            _ = v53
                                        }
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                    } else {
                                        v6 = (v6 | (v6 % ((2) | 1)))
                                    }
                                case "hello":
                                    v5 = v5
                                default:
                                    for fi54 := 0; fi54 < 2; fi54++ {
                                        if (false && v5) {
                                            var v55 uint = 10
                                            _ = v55
                                        }
                                        _ = fi54
                                    }
                                }
                            case 471:
                                var v56 uint8 = (v4 + 96)
                                _ = v56
                            default:
                                v57 := GenS7{F0: (v2 + ("hello" + "a")), F1: 75, F2: -1}
                                _ = v57
                            }
                            if v5 {
                                v4 = 237
                                v1 = ((v1 - (float32(1e10) + ((v1 - v1) + 1e10))) * 3.14159)
                            }
                            _ = v44
                        }
                        if v5 {
                            for fi58 := 0; fi58 < 11; fi58++ {
                                fmt.Println(v2, v1, fi33)
                                _ = fi58
                            }
                        } else {
                            switch (fi33 % ((0) | 1)) {
                            case 90:
                                for fi59 := 0; fi59 < 17; fi59++ {
                                    v6 = (^(v6))
                                    v60 := GenS7{F0: v2, F1: 2, F2: -1}
                                    _ = fi59
                                    _ = v60
                                }
                            case 557:
                                v5 = v5
                            }
                            if false {
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                if (-133 >= 739) {
                                    var v61 int32 = -2147483648
                                    _ = v61
                                }
                            }
                        }
                    }
                } else {
                    v62 := GenS2{F0: -42, F1: v2}
                    var v63 uint64 = 2
                    v2 = v2
                    _ = v62
                    _ = v63
                }
                _ = len(v18)
            }
        }
        _ = fi33
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
