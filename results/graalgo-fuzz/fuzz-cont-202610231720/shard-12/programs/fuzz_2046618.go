package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 string; F2 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16; F1 bool; F2 uint16 }
type GenS11 struct { F0 uint }
func genFn0() uint64 {
    var v5 float32 = -2.25
    if ((false && true) || (!(93 <= 128))) {
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 int32 = -2147483648
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 1; fi8++ {
            switch ((("hello" + "hello") + "foo bar") + "foo bar") {
            case "a":
                v9 := GenS1{F0: 70}
                _ = v9
            case "\n\t":
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 string = "a"
                    var v12 uint8 = 252
                    v13 := GenS2{F0: fi10, F1: 0}
                    fmt.Println(fi8)
                    var v14 float32 = -2.25
                    fmt.Println(v12, fi8, fi10)
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            default:
                v15 := GenS2{F0: 992, F1: 42}
                _ = v15
            }
            _ = fi8
        }
        if (false || (0 >= -9223372036854775808)) {
            v16 := []int{9223372036854775807}
            _ = v16
        } else {
            for fi17 := 0; fi17 < 13; fi17++ {
                v18 := []int{(fi17 / ((fi17) | 1))}
                _ = fi17
                _ = v18
            }
        }
    } else {
        for fi19 := 0; fi19 < 7; fi19++ {
            for fi20 := 0; fi20 < 2; fi20++ {
                v5 += (3.14159 / (((3.14159 / v5) + (v5 - ((-2.25 + v5) + 0.0))) / v5))
                v5 = 1e10
                fmt.Println(fi19, fi20)
                v5 += 0.0
                _ = fi20
            }
            var v21 uint32 = 42
            _ = fi19
            _ = v21
        }
        v22 := []int{327}
        _ = v22
    }
    v23 := GenS2{F0: 32767, F1: 1}
    _ = v23
    _ = v5
    return 732
}

func main() {
    var v1 bool = true
    var v2 int32 = 0
    v3 := []int{-837}
    v4 := []string{"", (("foo bar" + "\n\t") + ""), (("foo bar" + "") + ("" + ""))}
    v1 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch -2147483648 {
    case 128:
        for fi24 := 0; fi24 < 19; fi24++ {
            for fi25 := 0; fi25 < 7; fi25++ {
                v1 = (float32(1.5) < ((1.5 / ((3.14159 - 0.0) + ((float32(1e10) - (3.14159 + 1e10)) - float32(0.1)))) + -2.25))
                _ = fi25
            }
            _ = fi24
        }
    case 1:
        v1 = (!v1)
    default:
        fmt.Println(v1)
    }
    v26 := GenS3{F0: 42, F1: -32768}
    if v1 {
        v2 -= (v2 ^ v2)
        v27 := GenS4{F0: 1000}
        v1 = (((("\n\t" + "a") + "a") <= ("a" + (("\n\t" + ("\n\t" + "a")) + "a"))) && v1)
        _ = v27
    }
    switch 1 {
    case -9223372036854775808:
        for fi28 := 0; fi28 < 16; fi28++ {
            switch "a" {
            case "foo bar":
                v1 = v1
            case "hello":
                var v29 bool = ((fi28 & fi28) <= fi28)
                _ = v29
            case "\n\t":
                if (v1 && v1) {
                    for fi30 := 0; fi30 < 13; fi30++ {
                        v2 -= v2
                        _ = fi30
                    }
                    switch 32767 {
                    case 0:
                        for fi31 := 0; fi31 < 8; fi31++ {
                            if v1 {
                                if true {
                                    var v32 int8 = 0
                                    _ = v32
                                }
                            } else {
                                if false {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                if (!(2 <= 0)) {
                                    if false {
                                        if (!true) {
                                            v1 = v1
                                        } else {
                                            var v33 bool = v1
                                            _ = v33
                                        }
                                        v3 = append(v3, 9223372036854775807)
                                        var v34 float64 = 0.0
                                        _ = v34
                                    } else {
                                        var v35 uint32 = 362
                                        _ = v35
                                    }
                                    fmt.Println(v2, v1, fi31)
                                } else {
                                    var v36 uint32 = 2
                                    v36 += v36
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v36
                                }
                                fmt.Println(fi28, v2, v2)
                            }
                            v37 := GenS3{F0: 0, F1: 32767}
                            v38 := GenS5{F0: 18446744073709551615, F1: "a", F2: 42}
                            switch fi28 {
                            case -668:
                                v3 = append(v3, fi31)
                            case 4294967295:
                                var v39 int16 = int16(v2)
                                _ = v39
                            case 42:
                                v4 = append(v4, "")
                            }
                            _ = fi31
                            _ = v37
                            _ = v38
                        }
                    default:
                        if ((v1 && true) && (1000 < 369)) {
                            var v40 float32 = 3.14159
                            _ = v40
                        } else {
                            if (-1 > 42) {
                                if len(v3) > 0 {
                                    v3[0] = (fi28 & fi28)
                                }
                                if (779 >= 255) {
                                    switch 2147483647 {
                                    case 1:
                                        if (0 > 4294967295) {
                                            fmt.Println(v2, v1)
                                        }
                                    case -9223372036854775808:
                                        for fi41 := 0; fi41 < 4; fi41++ {
                                            v3 = append(v3, 863)
                                            _ = fi41
                                        }
                                    case -649:
                                        if ((true && (!v1)) || (!("" < ("" + ("foo bar" + "foo bar"))))) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v3 = append(v3, (fi28 % (((fi28 - (-(fi28)))) | 1)))
                                            if v1 {
                                                switch (fi28 << ((717) & 63)) {
                                                case -9223372036854775808:
                                                    if v1 {
                                                        var v42 int16 = -32768
                                                        for fi43 := 0; fi43 < 8; fi43++ {
                                                            var v44 float64 = (1.5 + (3.14159 * (0.1 + 0.0)))
                                                            _ = fi43
                                                            _ = v44
                                                        }
                                                        _ = v42
                                                    } else {
                                                        _ = len(v3)
                                                        if false {
                                                            v45 := GenS6{F0: ((("hello" + ((("hello" + "a") + ("foo bar" + "hello")) + ("hello" + ("\n\t" + (("" + "foo bar") + "foo bar"))))) + (((("foo bar" + "a") + ("foo bar" + "a")) + "") + "a")) + "\n\t")}
                                                            _ = v45
                                                        } else {
                                                            for fi46 := 0; fi46 < 11; fi46++ {
                                                                var v47 uint32 = 162
                                                                v48 := GenS1{F0: 1}
                                                                _ = fi46
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                            var v49 uint64 = genFn0()
                                                            _ = v49
                                                        }
                                                        v4 = append(v4, "foo bar")
                                                    }
                                                case -546:
                                                    for fi50 := 0; fi50 < 9; fi50++ {
                                                        v1 = (v1 && false)
                                                        switch fi28 {
                                                        case 2147483647:
                                                            v51 := GenS7{F0: 4294967295}
                                                            _ = v51
                                                        case 32768:
                                                            v52 := GenS3{F0: -638, F1: 0}
                                                            _ = v52
                                                        default:
                                                            v53 := GenS8{F0: 42, F1: 1000, F2: 266}
                                                            _ = v53
                                                        }
                                                        _ = fi50
                                                    }
                                                default:
                                                    if false {
                                                        _ = len(v4)
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        for fi54 := 0; fi54 < 11; fi54++ {
                                                            if false {
                                                                if len(v3) > 0 {
                                                                    v3[0] = fi28
                                                                }
                                                            } else {
                                                                v55 := GenS2{F0: (fi54 & 128), F1: -1}
                                                                v56 := GenS9{F0: 256}
                                                                v57 := GenS10{F0: 1, F1: ("\n\t" != "foo bar"), F2: 877}
                                                                _ = v55
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                            var v58 uint8 = 127
                                                            if v1 {
                                                                var v59 float64 = ((0.1 / 1.0) / float64(float32(-2.25)))
                                                                _ = v59
                                                            }
                                                            v58 += 2
                                                            var v60 float64 = 3.14159
                                                            _ = fi54
                                                            _ = v58
                                                            _ = v60
                                                        }
                                                    }
                                                }
                                                _ = len(v3)
                                            }
                                        } else {
                                            switch 483 {
                                            case -2147483648:
                                                switch fi28 {
                                                case 128:
                                                    if (451 != 698) {
                                                        v61 := GenS11{F0: 255}
                                                        v62 := GenS6{F0: "hello"}
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                case 4294967295:
                                                    switch -224 {
                                                    case 42:
                                                        if len(v3) > 0 {
                                                            v3[0] = fi28
                                                        }
                                                    default:
                                                        switch (fi28 ^ (-(fi28))) {
                                                        case 127:
                                                            v2 += -2147483648
                                                        case 4294967295:
                                                            v1 = v1
                                                        default:
                                                            v1 = (128 == 223)
                                                        }
                                                    }
                                                }
                                            case 42:
                                                for fi63 := 0; fi63 < 12; fi63++ {
                                                    if (1.5 == 1e10) {
                                                        fmt.Println(v2)
                                                        v2 = (-(v2))
                                                        v1 = (v1 || (v1 || v1))
                                                    } else {
                                                        v2 += 302
                                                        var v64 float32 = 0.0
                                                        var v65 int8 = -128
                                                        var v66 float64 = 0.0
                                                        v67 := GenS1{F0: 2}
                                                        _ = v64
                                                        _ = v65
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                    _ = len(v3)
                                                    if v1 {
                                                        v2 = v2
                                                        v68 := GenS1{F0: 53}
                                                        _ = v68
                                                    }
                                                    _ = fi63
                                                }
                                            case 308:
                                                var v69 uint = 1
                                                _ = v69
                                            default:
                                                var v70 int16 = 32767
                                                _ = v70
                                            }
                                            if (1.5 > (0.0 - 0.1)) {
                                                v1 = v1
                                                v2 = (-(v2))
                                            }
                                            var v71 int = 128
                                            _ = v71
                                        }
                                    default:
                                        v72 := GenS9{F0: 255}
                                        _ = v72
                                    }
                                }
                                v73 := GenS5{F0: 0, F1: ("hello" + ("\n\t" + ("foo bar" + ""))), F2: 42}
                                _ = v73
                            } else {
                                v3 = append(v3, (^(fi28)))
                                if len(v4) > 0 {
                                    v4[0] = ("a" + ("hello" + (("" + "a") + "foo bar")))
                                }
                            }
                        }
                    }
                    v74 := GenS6{F0: ""}
                    _ = v74
                }
            }
            _ = fi28
        }
    case 96:
        var v75 float32 = (1e10 + 3.14159)
        _ = v75
    case -162:
        if v1 {
            var v76 uint32 = 0
            _ = v76
        }
    }
    if v1 {
        for fi77 := 0; fi77 < 6; fi77++ {
            for fi78 := 0; fi78 < 11; fi78++ {
                fmt.Println(v2)
                if v1 {
                    var v79 bool = true
                    _ = v79
                }
                _ = fi78
            }
            v1 = ("a" > "foo bar")
            v1 = v1
            _ = fi77
        }
        v2 -= v2
        for fi80 := 0; fi80 < 16; fi80++ {
            var v81 uint64 = 104
            fmt.Println(v2, v1, fi80)
            var v82 uint8 = 255
            _ = fi80
            _ = v81
            _ = v82
        }
    } else {
        v83 := GenS7{F0: 93}
        fmt.Println(v1)
        _ = v83
    }
    fmt.Println(v1)
    var v84 bool = (!v1)
    v84 = v84
    var v85 uint8 = 42
    if ((1e10 != (3.14159 - 0.1)) && (genFn0() == 18446744073709551615)) {
        if (4294967295 > 262) {
            if true {
                v4 = append(v4, ("hello" + "hello"))
                v84 = (!(!v1))
            } else {
                var v86 uint8 = (v85 ^ v85)
                v87 := GenS2{F0: -1, F1: 0}
                _ = v86
                _ = v87
            }
        }
        v84 = true
        for fi88 := 0; fi88 < 20; fi88++ {
            v89 := GenS5{F0: 857, F1: ("a" + ("hello" + "hello")), F2: (v85 >> ((791) & 63))}
            v2 = v2
            _ = fi88
            _ = v89
        }
    }
    if (!(128 > -849)) {
        fmt.Println(v1)
        if v84 {
            fmt.Println(v1)
            var v90 int32 = (v2 & (v2 << ((255) & 63)))
            _ = v90
        }
    }
    var v91 int64 = 867
    fmt.Println(v85, v1, v2)
    v84 = (0 > 552)
    v92 := GenS1{F0: (^(v85))}
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v84
    _ = v85
    _ = v91
    _ = v92
}
