package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 string; F1 uint; F2 int64 }
type GenS3 struct { F0 uint32; F1 bool; F2 string }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int32; F1 int8; F2 int }
type GenS6 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 uint64; F1 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int32; F1 int }
type GenS11 struct { F0 uint8; F1 int }
func genFn0(p0 uint64) int8 {
    var v6 uint8 = 2
    v7 := []int{9223372036854775807, -9223372036854775808}
    if len(v7) > 0 {
        v7[0] = -2147483648
    }
    fmt.Println(p0)
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: 226, F1: (p0 * 0)}
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS1{F0: v6, F1: (p0 * p0)}
            var v12 uint32 = 1
            var v13 int64 = -9223372036854775808
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 int32, p1 uint64) int {
    if (true || (127 > 24)) {
        var v14 float64 = 1.5
        _ = v14
    }
    var v15 int = 662
    if true {
        v16 := GenS1{F0: 1, F1: 4294967295}
        _ = v16
    } else {
        v15 = (v15 * 32767)
        v17 := GenS2{F0: (("" + ((((("foo bar" + ("hello" + ("foo bar" + ("" + "")))) + "") + "") + ("a" + "hello")) + ("\n\t" + "hello"))) + "a"), F1: 4294967295, F2: 570}
        if false {
            if ((255 == 213) || (3.14159 == (1.5 - 1.5))) {
                if (!(5 > 42)) {
                    v18 := GenS2{F0: "hello", F1: 358, F2: 123}
                    _ = v18
                }
                v15 = v15
            }
        } else {
            v19 := []int{}
            _ = v19
        }
        _ = v17
    }
    var v20 uint16 = 0
    var v21 int8 = 42
    _ = p0
    _ = p1
    _ = v15
    _ = v20
    _ = v21
    return (v15 / ((v15) | 1))
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]string)
    var v3 bool = ((42 >= 1) || (-768 <= 42))
    var v4 bool = false
    var v5 float64 = (-2.25 - 3.14159)
    if v3 {
        fmt.Println(v1, v3)
        fmt.Println(v4)
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 bool = true
            if false {
                if (v5 == (v5 - v5)) {
                    fmt.Println(v1)
                }
            }
            _ = fi22
            _ = v23
        }
    }
    v5 += v5
    v24 := []int{-1, -1, 42, 0}
    var v25 int16 = 0
    var v26 uint = 65535
    v25 = v25
    v3 = (v25 < 42)
    if (-876 == 9223372036854775807) {
        if (42 != 0) {
            if false {
                var v27 int16 = 1000
                switch 42 {
                case 128:
                    for fi28 := 0; fi28 < 13; fi28++ {
                        if v4 {
                            if (!false) {
                                for fi29 := 0; fi29 < 14; fi29++ {
                                    fmt.Println(v4, fi28)
                                    v3 = false
                                    _ = fi29
                                }
                                for fi30 := 0; fi30 < 12; fi30++ {
                                    v31 := GenS1{F0: 128, F1: 0}
                                    v1 -= (v1 * genFn0(476))
                                    _ = fi30
                                    _ = v31
                                }
                                var v32 float64 = v5
                                _ = v32
                            } else {
                                if true {
                                    v33 := GenS3{F0: 740, F1: v4, F2: ("" + "")}
                                    v34 := GenS4{F0: -481, F1: 71}
                                    _ = v33
                                    _ = v34
                                }
                            }
                            v35 := GenS4{F0: 42, F1: 447}
                            _ = v35
                        }
                        v36 := GenS5{F0: 2147483647, F1: v1, F2: genFn1(-1, 18446744073709551615)}
                        if true {
                            v37 := GenS6{F0: -343, F1: 255, F2: 42}
                            if v4 {
                                if false {
                                    var v38 float32 = 1.5
                                    fmt.Println(v4, v27)
                                    v5 = (v5 - float64(((1e10 / ((v38 / v38) - ((3.14159 + v38) - v38))) / (v38 * float32(v5)))))
                                    _ = v38
                                } else {
                                    if (true && false) {
                                        v4 = v4
                                        if v3 {
                                            if true {
                                                if len(v24) >= 2 {
                                                    v24 = v24[1:len(v24)]
                                                }
                                            }
                                            fmt.Println(v3, v3, fi28)
                                        }
                                        if true {
                                            var v39 int64 = -678
                                            v40 := GenS4{F0: (-(v39)), F1: uint32(v26)}
                                            _ = v39
                                            _ = v40
                                        } else {
                                            delete(v2, (((("hello" + "\n\t") + "foo bar") + "hello") + ("a" + "\n\t")))
                                            if (!((^(v27)) > (v25 - 0))) {
                                                var v41 int8 = (v1 * (^(v1)))
                                                var v42 uint32 = 1
                                                _ = v41
                                                _ = v42
                                            }
                                        }
                                    } else {
                                        v24 = append(v24, (fi28 % ((fi28) | 1)))
                                        v43 := GenS6{F0: 204, F1: 32767, F2: 25}
                                        switch int(v1) {
                                        case 9223372036854775807:
                                            if v4 {
                                                var v44 uint32 = 42
                                                _ = v44
                                            } else {
                                                var v45 uint = 1
                                                var v46 int32 = -1
                                                _ = v45
                                                _ = v46
                                            }
                                        case 128:
                                            v5 = -2.25
                                        case 220:
                                            v47 := GenS7{F0: (fi28 - genFn1(192, 313)), F1: 65535}
                                            _ = v47
                                        }
                                        _ = v43
                                    }
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                }
                                for fi48 := 0; fi48 < 6; fi48++ {
                                    fmt.Println(fi28, v25, v5)
                                    v4 = false
                                    _ = fi48
                                }
                                fmt.Println(fi28, fi28)
                            }
                            _ = v37
                        }
                        _ = fi28
                        _ = v36
                    }
                default:
                    if (!(1 <= 415)) {
                        if (v4 && (!(!true))) {
                            fmt.Println(v27, v27, v26)
                            for fi49 := 0; fi49 < 16; fi49++ {
                                if false {
                                    if v3 {
                                        if true {
                                            for fi50 := 0; fi50 < 16; fi50++ {
                                                switch "hello" {
                                                case "\n\t":
                                                    v51 := GenS3{F0: 1, F1: (!v4), F2: "hello"}
                                                    _ = v51
                                                case "":
                                                    var v52 float64 = v5
                                                    _ = v52
                                                default:
                                                    fmt.Println(v4, v3, v3)
                                                }
                                                if v3 {
                                                    var v53 uint32 = 4294967295
                                                    var v54 int = genFn1(512, uint64(v53))
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    switch int(v1) {
                                                    case 2147483647:
                                                        var v55 int16 = (v25 / (((v25 << (((v26 % ((255) | 1))) & 63))) | 1))
                                                        _ = v55
                                                    case 32767:
                                                        var v56 uint16 = 2
                                                        _ = v56
                                                    case 500:
                                                        v27 = v25
                                                    }
                                                }
                                                var v57 int32 = 0
                                                _ = fi50
                                                _ = v57
                                            }
                                            if (-2.25 == -2.25) {
                                                if v4 {
                                                    if false {
                                                        fmt.Println(v26, v26, v1)
                                                        v4 = (v5 != float64(1.5))
                                                    } else {
                                                        for fi58 := 0; fi58 < 4; fi58++ {
                                                            var v59 int = 277
                                                            v4 = v3
                                                            _ = fi58
                                                            _ = v59
                                                        }
                                                        v3 = true
                                                    }
                                                } else {
                                                    v27 = (v25 >> (((^(v26))) & 63))
                                                    v4 = (v26 < (v26 << ((v26) & 63)))
                                                    if v3 {
                                                        var v60 bool = v4
                                                        v4 = v3
                                                        _ = v60
                                                    } else {
                                                        var v61 int32 = -402
                                                        fmt.Println(v5)
                                                        _ = v61
                                                    }
                                                }
                                            } else {
                                                v3 = false
                                                if v3 {
                                                    for fi62 := 0; fi62 < 6; fi62++ {
                                                        switch (("hello" + "\n\t") + "hello") {
                                                        case "foo bar":
                                                            v26 -= (^(v26))
                                                        case "hello":
                                                            v4 = (554 == 492)
                                                        case "\n\t":
                                                            var v63 float32 = 0.0
                                                            _ = v63
                                                        }
                                                        var v64 int8 = 127
                                                        _ = fi62
                                                        _ = v64
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        if v4 {
                                            if false {
                                                switch 32767 {
                                                case 4294967295:
                                                    var v65 bool = (-9223372036854775808 >= -9223372036854775808)
                                                    _ = v65
                                                case 42:
                                                    fmt.Println(v25)
                                                }
                                            }
                                            for fi66 := 0; fi66 < 20; fi66++ {
                                                v3 = ((v1 / ((v1) | 1)) > (-(v1)))
                                                v1 = (v1 | 100)
                                                v24 = append(v24, (fi49 + fi66))
                                                _ = fi66
                                            }
                                        } else {
                                            for fi67 := 0; fi67 < 10; fi67++ {
                                                v5 += (v5 - float64(-2.25))
                                                var v68 int16 = (-(v27))
                                                fmt.Println(v4, fi49, v5)
                                                v3 = (9223372036854775807 > 349)
                                                if (-2147483648 == 1) {
                                                    if v4 {
                                                        v69 := GenS8{F0: 382, F1: ("" + "")}
                                                        _ = v69
                                                    }
                                                    v70 := GenS9{F0: 42}
                                                    v1 += genFn0(1)
                                                    var v71 int = fi49
                                                    _ = v70
                                                    _ = v71
                                                }
                                                _ = fi67
                                                _ = v68
                                            }
                                        }
                                    }
                                    v1 = v1
                                    v24 = append(v24, (-(fi49)))
                                }
                                v72 := GenS3{F0: 42, F1: v3, F2: "\n\t"}
                                var v73 uint32 = 1
                                _ = fi49
                                _ = v72
                                _ = v73
                            }
                        }
                        v26 -= (v26 << ((2) & 63))
                        fmt.Println(v4, v25, v5)
                    }
                }
                _ = v27
            } else {
                v3 = (1000 < 265)
                for fi74 := 0; fi74 < 8; fi74++ {
                    for fi75 := 0; fi75 < 9; fi75++ {
                        var v76 uint32 = 923
                        v77 := GenS8{F0: 4294967295, F1: "a"}
                        _ = fi75
                        _ = v76
                        _ = v77
                    }
                    _ = fi74
                }
            }
            v1 -= (v1 >> (((v26 & 65535)) & 63))
            var v78 uint = (v26 + v26)
            _ = v78
        }
        switch ("hello" + ("" + "\n\t")) {
        case "hello":
            v79 := GenS5{F0: -2147483648, F1: (^(v1)), F2: genFn1(42, 0)}
            _ = v79
        case "":
            v80 := GenS10{F0: -2147483648, F1: genFn1(2147483647, 1)}
            _ = v80
        case "\n\t":
            v81 := GenS11{F0: 37, F1: 725}
            _ = v81
        default:
            var v82 int32 = 1
            _ = v82
        }
        var v83 int16 = (-(v25))
        _ = v83
    }
    switch -1 {
    case -84:
        v84 := GenS4{F0: -1, F1: 4294967295}
        _ = v84
    default:
        v25 += v25
    }
    var v85 float32 = (0.0 * ((float32(v5) * 1.5) / 3.14159))
    var v86 int8 = 127
    if false {
        var v87 int = int(v26)
        _ = v87
    }
    for fi88 := 0; fi88 < 10; fi88++ {
        fmt.Println(v5)
        var v89 int = int(v26)
        _ = fi88
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v85
    _ = v86
}
