package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 bool; F1 uint; F2 int }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint16, p2 int16) int {
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: 127, F1: 0}
        _ = fi6
        _ = v7
    }
    var v8 int16 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 319
}
func genFn1() int16 {
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 string = ("" + ("hello" + ""))
        v10 = ("" + "\n\t")
        _ = fi9
        _ = v10
    }
    if (true && true) {
        v11 := GenS2{F0: (-1 != 1), F1: 4294967295, F2: 32767}
        _ = v11
    }
    v12 := GenS1{F0: 127, F1: 42}
    var v13 int64 = -1
    _ = v12
    _ = v13
    return 32767
}

func main() {
    var v1 float32 = 3.14159
    var v2 int8 = 23
    v3 := make(map[string]bool)
    v4 := []uint32{844, 0, 1}
    var v5 int32 = 2147483647
    for fi14 := 0; fi14 < 15; fi14++ {
        if false {
            if false {
                v2 = (v2 % (((v2 << ((255) & 63))) | 1))
                v15 := GenS1{F0: 203, F1: 321}
                if true {
                    var v16 float64 = ((-2.25 / 1.0) / -2.25)
                    _ = v16
                } else {
                    delete(v3, "\n\t")
                    fmt.Println(fi14, v1)
                }
                _ = v15
            }
            switch fi14 {
            case -2147483648:
                v2 -= (-(v2))
            case -1:
                if false {
                    if true {
                        for fi17 := 0; fi17 < 10; fi17++ {
                            var v18 int32 = -2147483648
                            switch fi17 {
                            case 0:
                                var v19 int8 = (-(v2))
                                _ = v19
                            case -2147483648:
                                fmt.Println(fi14, v5, v5)
                            case 465:
                                v18 += v5
                            default:
                                v18 = (v18 - int32(fi17))
                            }
                            var v20 uint8 = 37
                            _ = fi17
                            _ = v18
                            _ = v20
                        }
                    }
                } else {
                    for fi21 := 0; fi21 < 2; fi21++ {
                        v1 += (0.1 + v1)
                        _ = fi21
                    }
                    var v22 uint32 = 2147483648
                    for fi23 := 0; fi23 < 11; fi23++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if (!true) {
                            v24 := GenS2{F0: (((!(18446744073709551615 == 2)) && false) || ((-(v2)) > 100)), F1: 65535, F2: fi14}
                            _ = v24
                        }
                        v25 := GenS1{F0: 0, F1: 9223372036854775807}
                        var v26 uint64 = 4294967295
                        v26 += v26
                        _ = fi23
                        _ = v25
                        _ = v26
                    }
                    _ = v22
                }
            }
            for fi27 := 0; fi27 < 5; fi27++ {
                switch "" {
                case "\n\t":
                    v28 := GenS3{F0: (-(v2)), F1: 2}
                    _ = v28
                case "foo bar":
                    fmt.Println(v2, v5, fi27)
                default:
                    var v29 bool = (1e10 > (-2.25 + ((((3.14159 / 3.14159) / 1.0) / 1.5) * 0.1)))
                    _ = v29
                }
                var v30 float64 = 0.0
                if (((v1 - (v1 / 0.1)) <= float32(float64(float32(v30)))) && false) {
                    v31 := GenS4{F0: (v2 + v2), F1: 255}
                    var v32 uint32 = 4294967295
                    var v33 int8 = 42
                    _ = v31
                    _ = v32
                    _ = v33
                } else {
                    v2 = 0
                    v34 := GenS2{F0: false, F1: 2, F2: genFn0(18, 32767, 0)}
                    switch (fi14 + (fi14 ^ genFn0(93, 0, -961))) {
                    case 42:
                        if ((1 >= 1) && false) {
                            _ = len(v3)
                            if len(v4) > 0 {
                                v4[0] = 826
                            }
                            var v35 uint32 = 4294967295
                            _ = v35
                        } else {
                            var v36 bool = ("a" != (("foo bar" + "a") + ("\n\t" + ("" + "a"))))
                            if v36 {
                                var v37 uint = 42
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v37
                            } else {
                                v36 = (!((false && v36) && true))
                                v38 := GenS3{F0: v2, F1: 4294967295}
                                v39 := GenS5{F0: 2147483647}
                                _ = v38
                                _ = v39
                            }
                            v40 := GenS1{F0: 2, F1: 9223372036854775807}
                            v4 = append(v4, 0)
                            if (!v36) {
                                switch "\n\t" {
                                case "":
                                    var v41 float64 = float64(v1)
                                    _ = v41
                                case "a":
                                    var v42 string = "a"
                                    _ = v42
                                case "foo bar":
                                    var v43 uint16 = 65535
                                    _ = v43
                                }
                            }
                            _ = v36
                            _ = v40
                        }
                        v1 = (v1 * v1)
                    case 0:
                        v1 = v1
                    default:
                        fmt.Println(fi27)
                    }
                    _ = v34
                }
                _ = fi27
                _ = v30
            }
        }
        _ = fi14
    }
    if false {
        fmt.Println(v2)
        v1 += v1
        var v44 int = 0
        _ = v44
    } else {
        switch 1 {
        case 42:
            fmt.Println(v1, v2, v5)
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v4 = append(v4, 2147483648)
    v5 -= v5
    for fi45 := 0; fi45 < 12; fi45++ {
        v46 := GenS6{F0: 512}
        v47 := GenS6{F0: -1}
        for fi48 := 0; fi48 < 5; fi48++ {
            var v49 uint16 = 32767
            switch genFn0(82, v49, 42) {
            case 1:
                var v50 uint16 = v49
                _ = v50
            }
            _ = fi48
            _ = v49
        }
        _ = fi45
        _ = v46
        _ = v47
    }
    var v51 int64 = 1
    fmt.Println(v2)
    var v52 bool = true
    for fi53 := 0; fi53 < 1; fi53++ {
        var v54 float64 = 1.5
        var v55 string = ""
        var v56 uint64 = 0
        _ = fi53
        _ = v54
        _ = v55
        _ = v56
    }
    if true {
        switch 0 {
        case -1:
            v57 := GenS7{F0: v5}
            _ = v57
        case -574:
            for fi58 := 0; fi58 < 7; fi58++ {
                var v59 float32 = 1e10
                if len(v4) > 0 {
                    v4[0] = 2147483648
                }
                _ = fi58
                _ = v59
            }
        case -155:
            _ = len(v3)
        default:
            for fi60 := 0; fi60 < 10; fi60++ {
                fmt.Println(v5)
                _ = fi60
            }
        }
        for fi61 := 0; fi61 < 13; fi61++ {
            if len(v4) > 0 {
                v4[0] = 562
            }
            _ = fi61
        }
        v62 := GenS4{F0: (v2 << ((854) & 63)), F1: 1}
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
}
