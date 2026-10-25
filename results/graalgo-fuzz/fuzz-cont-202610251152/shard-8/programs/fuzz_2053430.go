package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 bool }
type GenS3 struct { F0 uint16; F1 string; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16; F1 uint; F2 uint }
func genFn0() uint8 {
    v7 := []int{-2147483648, 692, 415}
    if false {
        if (!(!(((42 < 42) || (348 <= 65535)) || (!((!(!false)) && (1 >= 42)))))) {
            var v8 int64 = 1
            var v9 float64 = 3.14159
            switch -2147483648 {
            case -1:
                var v10 int = 4294967295
                _ = v10
            default:
                v7 = append(v7, 9223372036854775807)
            }
            _ = v8
            _ = v9
        }
        switch -82 {
        case 247:
            for fi11 := 0; fi11 < 5; fi11++ {
                fmt.Println(fi11)
                if true {
                    var v12 uint32 = 42
                    for fi13 := 0; fi13 < 17; fi13++ {
                        v12 = 2147483647
                        var v14 uint64 = 2
                        _ = fi13
                        _ = v14
                    }
                    switch 42 {
                    case -1:
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    case 1:
                        v15 := GenS1{F0: 961}
                        _ = v15
                    default:
                        for fi16 := 0; fi16 < 12; fi16++ {
                            var v17 float32 = 1e10
                            _ = fi16
                            _ = v17
                        }
                    }
                    _ = v12
                }
                if (true || ((!false) && (35 < 32))) {
                    var v18 float64 = (0.0 * (1e10 + 0.1))
                    _ = v18
                }
                _ = fi11
            }
        case 1:
            v19 := GenS2{F0: 65535, F1: 18446744073709551615, F2: false}
            _ = v19
        case 4294967295:
            if (true && false) {
                var v20 int16 = 945
                v21 := GenS3{F0: 2, F1: ((("foo bar" + "foo bar") + "a") + ("" + "foo bar")), F2: 4294967295}
                v20 += (v20 / ((-283) | 1))
                _ = v20
                _ = v21
            } else {
                var v22 int = 42
                var v23 float64 = 1e10
                v7 = append(v7, -9223372036854775808)
                _ = v22
                _ = v23
            }
        default:
            var v24 uint16 = 0
            _ = v24
        }
        var v25 int = -9223372036854775808
        _ = v25
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if (734 > 2) {
            var v26 uint16 = 690
            for fi27 := 0; fi27 < 4; fi27++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi27
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 uint = 1
                if len(v7) > 0 {
                    v7[0] = 851
                }
                var v30 uint = uint(v26)
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v26
        }
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        if true {
            for fi32 := 0; fi32 < 16; fi32++ {
                var v33 float64 = 0.1
                _ = fi32
                _ = v33
            }
            switch (((("\n\t" + (("" + "foo bar") + "hello")) + (((("foo bar" + "a") + "a") + ("foo bar" + "a")) + ((("a" + ("" + ("\n\t" + "hello"))) + ((((("\n\t" + (("a" + "foo bar") + "hello")) + "\n\t") + "foo bar") + "hello") + "a")) + ("foo bar" + "hello")))) + "\n\t") + (("foo bar" + "foo bar") + "\n\t")) {
            case "\n\t":
                var v34 uint32 = 0
                _ = v34
            case "foo bar":
                var v35 uint64 = 0
                _ = v35
            default:
                if (0 == 0) {
                    if false {
                        var v36 uint32 = 4294967295
                        _ = v36
                    } else {
                        v37 := GenS4{F0: (fi31 & fi31)}
                        _ = v37
                    }
                    if (!true) {
                        var v38 uint64 = 0
                        _ = v38
                    } else {
                        var v39 uint16 = 2
                        for fi40 := 0; fi40 < 1; fi40++ {
                            var v41 uint32 = 2
                            v42 := GenS4{F0: (fi40 / (((fi31 & 2147483647)) | 1))}
                            var v43 uint64 = 179
                            fmt.Println(v43)
                            fmt.Println(v39, v41, v39)
                            _ = fi40
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        _ = v39
                    }
                }
            }
        } else {
            v44 := GenS2{F0: 4294967295, F1: 279, F2: (1e10 == (((-2.25 * float32((float64(((3.14159 * 1.5) + 1.5)) * 0.0))) * (float32(float64((0.0 * float32((3.14159 / (1e10 - (3.14159 - (((float64(float32(0.0)) - ((-2.25 + 0.0) - 1.5)) * 1e10) / 1.5)))))))) - 3.14159)) * 1e10))}
            for fi45 := 0; fi45 < 17; fi45++ {
                v46 := GenS1{F0: 1}
                v47 := GenS2{F0: 1, F1: 982, F2: true}
                _ = len(v7)
                _ = fi45
                _ = v46
                _ = v47
            }
            _ = v44
        }
        if false {
            v7 = append(v7, fi31)
            fmt.Println(fi31)
            for fi48 := 0; fi48 < 9; fi48++ {
                if len(v7) > 0 {
                    v7[0] = fi48
                }
                _ = fi48
            }
        } else {
            if (18446744073709551615 >= 42) {
                v49 := GenS2{F0: 2147483647, F1: 42, F2: (0.1 != 0.0)}
                v7 = append(v7, fi31)
                for fi50 := 0; fi50 < 8; fi50++ {
                    switch ("hello" + "") {
                    case "\n\t":
                        var v51 uint16 = 752
                        _ = v51
                    case "a":
                        var v52 int16 = 435
                        _ = v52
                    case "foo bar":
                        var v53 float32 = 3.14159
                        _ = v53
                    default:
                        if (true && true) {
                            if len(v7) > 0 {
                                v7[0] = fi50
                            }
                        } else {
                            if false {
                                var v54 int = 1
                                _ = v54
                            }
                        }
                    }
                    var v55 float32 = 0.1
                    v56 := GenS5{F0: 142}
                    var v57 uint16 = 563
                    _ = fi50
                    _ = v55
                    _ = v56
                    _ = v57
                }
                _ = v49
            }
            var v58 uint = 1
            _ = v58
        }
        var v59 bool = true
        _ = fi31
        _ = v59
    }
    _ = v7
    return 255
}

func main() {
    var v1 uint64 = 499
    var v2 int = 540
    var v3 uint64 = v1
    var v4 uint32 = 656
    var v5 int = 127
    var v6 int16 = -890
    var v60 int8 = 42
    v61 := GenS4{F0: 698}
    v62 := []int{}
    v3 = (v3 % ((v1) | 1))
    var v63 bool = (v60 >= (v60 << ((4294967295) & 63)))
    v6 = (v6 - -1)
    v1 = v3
    if v63 {
        for fi64 := 0; fi64 < 6; fi64++ {
            v65 := GenS6{F0: 781, F1: uint(v4), F2: 778}
            var v66 bool = (-2147483648 > -570)
            var v67 float32 = 0.0
            _ = fi64
            _ = v65
            _ = v66
            _ = v67
        }
        v1 = (v1 | v1)
    } else {
        var v68 bool = false
        _ = v68
    }
    _ = len(v62)
    if len(v62) > 0 {
        v62[0] = v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
