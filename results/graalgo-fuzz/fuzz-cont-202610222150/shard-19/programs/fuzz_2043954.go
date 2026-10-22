package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint; F1 uint8; F2 int }
type GenS7 struct { F0 int16; F1 bool; F2 int64 }
type GenS8 struct { F0 uint32; F1 uint8; F2 bool }
type GenS9 struct { F0 int16; F1 int8; F2 string }
func genFn0() bool {
    var v4 int8 = -128
    v4 -= (v4 >> ((74) & 63))
    switch "foo bar" {
    case "foo bar":
        var v5 uint16 = 255
        _ = v5
    case "a":
        switch "" {
        case "foo bar":
            var v6 uint = 4294967295
            _ = v6
        case "hello":
            v4 = (v4 ^ v4)
        }
    case "\n\t":
        var v7 int16 = 0
        _ = v7
    default:
        var v8 uint = 18446744073709551615
        _ = v8
    }
    v9 := GenS1{F0: 0}
    _ = v4
    _ = v9
    return (-1 < 42)
}
func genFn1(p0 int16) uint16 {
    p0 += p0
    var v10 int = 128
    for fi11 := 0; fi11 < 1; fi11++ {
        v10 += v10
        p0 = 1
        _ = fi11
    }
    switch ("a" + "\n\t") {
    case "hello":
        v12 := GenS2{F0: 65535, F1: 219, F2: 18446744073709551615}
        _ = v12
    }
    _ = p0
    _ = v10
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 1
    var v3 int64 = 1
    var v13 uint64 = 2
    _ = len(v1)
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            if true {
                v3 += v3
                switch 2147483647 {
                case -2147483648:
                    v16 := GenS3{F0: ""}
                    _ = v16
                case 769:
                    v3 = v3
                case 32768:
                    var v17 bool = ("" == "foo bar")
                    v18 := v1[("hello" + "")]
                    switch -9223372036854775808 {
                    case 32768:
                        if (159 >= genFn1(1000)) {
                            var v19 int8 = -128
                            _ = v19
                        } else {
                            v20 := GenS2{F0: 2147483648, F1: 128, F2: 65535}
                            if false {
                                fmt.Println(v2, v2, fi15)
                            } else {
                                var v21 uint = 4294967295
                                if (950 == 732) {
                                    if (v17 && (!((v17 || (v17 || (v21 == (v21 % ((v21) | 1))))) && v17))) {
                                        fmt.Println(v3, fi15)
                                        if ((!v17) && v17) {
                                            var v22 int8 = 127
                                            if (42 <= 32767) {
                                                v22 = v22
                                                if v17 {
                                                    var v23 int = 0
                                                    if v17 {
                                                        v17 = (("a" != ("" + "a")) || (2 > 65535))
                                                        fmt.Println(v22, v22)
                                                    }
                                                    _ = v23
                                                }
                                            }
                                            _ = v22
                                        } else {
                                            var v24 int32 = int32(v13)
                                            var v25 uint8 = 102
                                            _ = v24
                                            _ = v25
                                        }
                                        var v26 uint8 = 127
                                        var v27 float64 = 0.1
                                        if (true || (-533 != int16(v13))) {
                                            if true {
                                                v18 -= 128
                                            } else {
                                                if true {
                                                    delete(v1, "foo bar")
                                                }
                                            }
                                        } else {
                                            if (true && (606 <= 32767)) {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if true {
                                                        v28 := GenS4{F0: fi14}
                                                        var v29 uint16 = 553
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        var v30 float64 = 1e10
                                                        var v31 string = "\n\t"
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                case "a":
                                                    var v32 int = 32767
                                                    _ = v32
                                                default:
                                                    v33 := []int{fi15}
                                                    _ = v33
                                                }
                                            }
                                        }
                                        _ = v26
                                        _ = v27
                                    }
                                    if ((v17 && v17) || (65535 != 0)) {
                                        v3 -= v3
                                        if v17 {
                                            v13 += 125
                                            v34 := GenS4{F0: 42}
                                            if true {
                                                var v35 uint32 = 1
                                                if v17 {
                                                    v18 += (fi15 << ((v21) & 63))
                                                }
                                                fmt.Println(fi14)
                                                v36 := GenS5{F0: "", F1: 65535}
                                                _ = v35
                                                _ = v36
                                            }
                                            _ = v34
                                        } else {
                                            v1["a"] = 128
                                            v21 = (v21 & v21)
                                            v37 := GenS6{F0: 854, F1: 138, F2: 9223372036854775807}
                                            var v38 string = (("a" + "foo bar") + "foo bar")
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                    switch fi14 {
                                    case 32767:
                                        if ((-2.25 * 0.0) <= (-2.25 / (1e10 * (((3.14159 - ((0.1 * 3.14159) / 1.0)) - (1.5 + (0.1 * (1.5 + (-2.25 + 3.14159))))) + 3.14159)))) {
                                            v39 := GenS6{F0: 255, F1: 0, F2: (v18 >> ((65535) & 63))}
                                            if (((!v17) && v17) && v17) {
                                                if (!((-1 <= int16(v18)) || false)) {
                                                    switch "a" {
                                                    case "foo bar":
                                                        v40 := GenS5{F0: "a", F1: 255}
                                                        _ = v40
                                                    }
                                                } else {
                                                    var v41 int32 = (v2 * (-(v2)))
                                                    if (680 > 4294967295) {
                                                        v17 = (376 > 2147483647)
                                                        fmt.Println(v21, v2, v41)
                                                    }
                                                    var v42 bool = v17
                                                    v17 = (!v17)
                                                    _ = v41
                                                    _ = v42
                                                }
                                                switch fi15 {
                                                case -9223372036854775808:
                                                    if (!true) {
                                                        v17 = (4294967295 > 4294967295)
                                                    }
                                                case -384:
                                                    fmt.Println(v17)
                                                default:
                                                    switch 32768 {
                                                    case -211:
                                                        v43 := GenS1{F0: 0}
                                                        _ = v43
                                                    }
                                                    v44 := []int{}
                                                    _ = v44
                                                }
                                                v45 := GenS7{F0: 0, F1: v17, F2: (v3 ^ 42)}
                                                if v17 {
                                                    v46 := GenS8{F0: 42, F1: 14, F2: ((v3 / ((v3) | 1)) <= (v3 & 1))}
                                                    fmt.Println(v18, v3)
                                                    _ = v46
                                                }
                                                _ = v45
                                            } else {
                                                fmt.Println(fi14, v2, v2)
                                                fmt.Println(v21, v21, fi14)
                                                if (!(255 <= 0)) {
                                                    if (32767 == 42) {
                                                        if false {
                                                            v3 += (-(v3))
                                                            fmt.Println(fi14)
                                                        }
                                                        if (2147483647 != 655) {
                                                            fmt.Println(v21, v21)
                                                            var v47 bool = (-112 > -1)
                                                            _ = v47
                                                        }
                                                    } else {
                                                        fmt.Println(v21, v13)
                                                        delete(v1, "hello")
                                                        _ = len(v1)
                                                    }
                                                }
                                            }
                                            v17 = (!(false || v17))
                                            _ = v39
                                        } else {
                                            fmt.Println(v21, fi15, v18)
                                        }
                                    }
                                } else {
                                    var v48 float32 = ((((0.0 - ((3.14159 * 3.14159) * float32(-2.25))) - -2.25) - 0.0) + 0.1)
                                    var v49 float32 = v48
                                    _ = len(v1)
                                    _ = v48
                                    _ = v49
                                }
                                v50 := GenS3{F0: ("a" + "a")}
                                var v51 int16 = -45
                                _ = v21
                                _ = v50
                                _ = v51
                            }
                            fmt.Println(v2)
                            _ = v20
                        }
                    default:
                        var v52 uint64 = v13
                        _ = v52
                    }
                    _ = v17
                    _ = v18
                }
                switch "" {
                case "hello":
                    v3 -= 0
                case "a":
                    switch fi15 {
                    case 1:
                        if (false || (!(2 > 974))) {
                            v53 := GenS4{F0: (fi15 << ((0) & 63))}
                            fmt.Println(v2, v3)
                            _ = v53
                        } else {
                            _ = len(v1)
                            v2 = (v2 ^ -782)
                            v54 := []int{(fi15 << ((42) & 63)), (fi14 - int(v13))}
                            _ = v54
                        }
                    case 515:
                        var v55 uint64 = (^(v13))
                        _ = v55
                    default:
                        switch 0 {
                        case -721:
                            v56 := []int{}
                            _ = v56
                        default:
                            var v57 float32 = (3.14159 - float32((1.5 / 3.14159)))
                            _ = v57
                        }
                        v2 = (^(v2))
                    }
                case "foo bar":
                    v13 = (v13 / (((^(v13))) | 1))
                    v3 += -423
                    switch int(v2) {
                    case 32768:
                        v58 := GenS3{F0: "foo bar"}
                        _ = v58
                    case 1:
                        var v59 uint8 = 128
                        _ = v59
                    }
                default:
                    v13 = (v13 << ((65535) & 63))
                }
            }
            _ = fi15
        }
        v13 += v13
        var v60 int16 = -1
        _ = fi14
        _ = v60
    }
    delete(v1, (((("a" + "hello") + "a") + ("a" + "\n\t")) + "hello"))
    v2 -= v2
    var v61 int = -2147483648
    if false {
        for fi62 := 0; fi62 < 4; fi62++ {
            var v63 int16 = 0
            var v64 bool = (genFn1((v63 + 1)) == 32767)
            for fi65 := 0; fi65 < 12; fi65++ {
                var v66 uint8 = 128
                if v64 {
                    v66 += v66
                    v61 -= (fi65 * (fi62 << ((1) & 63)))
                    v67 := GenS9{F0: -32768, F1: 101, F2: "a"}
                    _ = v67
                }
                _ = fi65
                _ = v66
            }
            _ = fi62
            _ = v63
            _ = v64
        }
    }
    var v68 int16 = 70
    if false {
        v3 = v3
        if (!(222 < 42)) {
            if (("a" == "a") && true) {
                v13 = (v13 / (((v13 & (v13 & v13))) | 1))
                v68 = (-(v68))
            } else {
                v61 -= (v61 | v61)
                v3 = (v3 - (v3 / (((-(v3))) | 1)))
            }
            var v69 int16 = v68
            v70 := GenS2{F0: 2, F1: 42, F2: 1}
            _ = v69
            _ = v70
        } else {
            if (v13 == (v13 >> ((18446744073709551615) & 63))) {
                fmt.Println(v68, v13, v68)
                v13 = v13
                if (v3 >= v3) {
                    var v71 float64 = (((3.14159 + 0.0) + 3.14159) + -2.25)
                    _ = v71
                } else {
                    for fi72 := 0; fi72 < 13; fi72++ {
                        var v73 uint16 = 189
                        v13 -= v13
                        switch ((("" + "a") + ("\n\t" + ("\n\t" + ("foo bar" + (("a" + "\n\t") + "\n\t"))))) + (("foo bar" + "\n\t") + "foo bar")) {
                        case "foo bar":
                            v3 = v3
                        default:
                            fmt.Println(v61)
                        }
                        _ = fi72
                        _ = v73
                    }
                    var v74 int32 = v2
                    if (v68 >= 32767) {
                        var v75 int8 = 42
                        var v76 int = v61
                        _ = v75
                        _ = v76
                    }
                    _ = v74
                }
            }
            _ = len(v1)
        }
    } else {
        v61 = (v61 - v61)
    }
    if true {
        switch v61 {
        case 42:
            v61 = (v61 + (-(v61)))
        case -1:
            var v77 uint = 42
            _ = v77
        case 32767:
            v3 -= 876
        default:
            v13 += v13
        }
        v78 := GenS7{F0: 1000, F1: false, F2: v3}
        v2 = v2
        _ = v78
    } else {
        fmt.Println(v68, v68, v2)
        v2 = 945
        var v79 uint64 = 18446744073709551615
        _ = v79
    }
    for fi80 := 0; fi80 < 17; fi80++ {
        v2 = (v2 - -22)
        _ = fi80
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v61
    _ = v68
}
