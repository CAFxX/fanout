package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint; F1 int64; F2 uint16 }
type GenS9 struct { F0 bool; F1 int16; F2 string }
type GenS10 struct { F0 int8 }
func genFn0() bool {
    v4 := GenS1{F0: true, F1: 2147483647}
    var v5 float32 = (float32(3.14159) + 3.14159)
    v6 := []int{}
    _ = v4
    _ = v5
    _ = v6
    return (true && true)
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []string{}
    var v3 int16 = -18
    v1 = 42
    var v7 float64 = 3.14159
    fmt.Println(v1, v7, v1)
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: (!true), F1: 2147483647}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v7 = 0.0
        var v11 uint64 = 4294967295
        _ = fi10
        _ = v11
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    v1 -= (^(v1))
    v7 += 3.14159
    v1 = v1
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 int8 = -128
            switch (((("a" + "\n\t") + "") + ("\n\t" + ("foo bar" + ""))) + "\n\t") {
            case "a":
                var v15 uint64 = 1
                _ = v15
            }
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v3 += 892
        v17 := GenS1{F0: (false || false), F1: -2147483648}
        _ = fi16
        _ = v17
    }
    if (42 >= 32767) {
        var v18 uint8 = 42
        v18 -= 107
        _ = v18
    } else {
        switch "" {
        case "":
            v7 -= v7
        case "\n\t":
            var v19 int = 32768
            _ = v19
        case "foo bar":
            var v20 float64 = float64(float32((v7 + float64((1e10 - 1.5)))))
            _ = v20
        }
        switch ("\n\t" + ("foo bar" + "hello")) {
        case "hello":
            v7 += (float64(0.0) / float64((float32((float64(((0.1 - 1e10) / (1.5 / ((float32(v7) - (3.14159 * (1e10 - 3.14159))) + 0.1)))) - 1.5)) * 0.1)))
        case "":
            for fi21 := 0; fi21 < 11; fi21++ {
                for fi22 := 0; fi22 < 8; fi22++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v23 float32 = float32(v7)
                    if (fi21 == 9223372036854775807) {
                        v24 := GenS2{F0: 827}
                        switch "hello" {
                        case "a":
                            if false {
                                fmt.Println(v3, fi22)
                                var v25 bool = false
                                v1 -= 2147483648
                                _ = v25
                            }
                        case "":
                            var v26 float64 = v7
                            _ = v26
                        case "hello":
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v27 uint8 = 127
                                _ = v27
                            } else {
                                fmt.Println(v23, v23, v7)
                            }
                        }
                        _ = v24
                    } else {
                        if false {
                            var v28 uint = 973
                            if (int8(v28) > 100) {
                                v1 += (v1 - (v1 | 1))
                                if (!(!((!((-418 >= 42) || false)) && true))) {
                                    _ = len(v2)
                                    var v29 int8 = 55
                                    v30 := GenS1{F0: false, F1: 2147483647}
                                    if (-2.25 >= 0.0) {
                                        if len(v2) > 0 {
                                            v2[0] = ("foo bar" + "hello")
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        var v31 int8 = v29
                                        v32 := GenS1{F0: (-9223372036854775808 != 1), F1: -386}
                                        if true {
                                            switch fi21 {
                                            case 32767:
                                                if true {
                                                    if false {
                                                        if true {
                                                            if true {
                                                                if true {
                                                                    fmt.Println(fi21)
                                                                    if (false && (!((v28 % ((v28) | 1)) > 4294967295))) {
                                                                        v33 := GenS2{F0: 255}
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            var v34 string = ("" + "")
                                                                            var v35 uint = v28
                                                                            v36 := GenS2{F0: 2}
                                                                            fmt.Println(v7)
                                                                            _ = v34
                                                                            _ = v35
                                                                            _ = v36
                                                                        case "a":
                                                                            fmt.Println(v7, v1, v23)
                                                                        default:
                                                                            v7 = float64(v23)
                                                                        }
                                                                        if ((^(v28)) >= v28) {
                                                                            fmt.Println(fi21, v23)
                                                                            var v37 uint8 = 127
                                                                            fmt.Println(fi21, fi22)
                                                                            v38 := GenS3{F0: 42}
                                                                            _ = v37
                                                                            _ = v38
                                                                        } else {
                                                                            var v39 float64 = (float64(v23) + ((float64(v23) + (1.5 - -2.25)) * v7))
                                                                            _ = v39
                                                                        }
                                                                        _ = v33
                                                                    } else {
                                                                        v40 := GenS2{F0: 131}
                                                                        fmt.Println(v3)
                                                                        v28 += uint(v1)
                                                                        _ = v40
                                                                    }
                                                                    v41 := GenS3{F0: 0}
                                                                    v42 := GenS4{F0: 2}
                                                                    v3 = v3
                                                                    _ = v41
                                                                    _ = v42
                                                                } else {
                                                                    v43 := GenS5{F0: 0}
                                                                    var v44 uint64 = 18446744073709551615
                                                                    var v45 float32 = 3.14159
                                                                    _ = v43
                                                                    _ = v44
                                                                    _ = v45
                                                                }
                                                            }
                                                            v46 := GenS1{F0: false, F1: -54}
                                                            fmt.Println(v28, v3, v31)
                                                            _ = v46
                                                        } else {
                                                            if false {
                                                                v7 += (v7 - -2.25)
                                                            }
                                                            if false {
                                                                fmt.Println(v23)
                                                            }
                                                            var v47 float64 = 0.0
                                                            _ = v47
                                                        }
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = "foo bar"
                                                    }
                                                } else {
                                                    if ((0.1 >= v23) || false) {
                                                        v2 = append(v2, ("hello" + (("hello" + "") + ("a" + "foo bar"))))
                                                        var v48 int16 = -546
                                                        switch 127 {
                                                        case 4294967295:
                                                            v7 = 1.5
                                                        case 128:
                                                            if (!false) {
                                                                var v49 uint8 = 128
                                                                _ = v49
                                                            } else {
                                                                v1 = 0
                                                                switch fi21 {
                                                                case -1:
                                                                    v29 -= v31
                                                                default:
                                                                    if (!(((v31 | int8(v28)) >= 42) && (((!((!(!false)) || ((true || false) || (-1 > int64(v31))))) && (0 < 526)) && ((true || (0 > 711)) && true)))) {
                                                                        var v50 float64 = 3.14159
                                                                        v51 := GenS6{F0: -1}
                                                                        v48 = v3
                                                                        _ = v50
                                                                        _ = v51
                                                                    }
                                                                }
                                                                var v52 uint8 = 255
                                                                _ = v52
                                                            }
                                                        }
                                                        if (!true) {
                                                            v1 = (v1 & v1)
                                                            if false {
                                                                v2 = append(v2, "a")
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                        } else {
                                                            v53 := GenS7{F0: fi21, F1: 2}
                                                            v23 += 0.0
                                                            _ = v53
                                                        }
                                                        _ = v48
                                                    } else {
                                                        var v54 uint8 = uint8(v1)
                                                        _ = v54
                                                    }
                                                }
                                            case -246:
                                                v31 -= (v29 * (v31 + v29))
                                            case -2147483648:
                                                v3 += 0
                                            default:
                                                var v55 uint8 = 127
                                                _ = v55
                                            }
                                            v56 := GenS6{F0: -1}
                                            _ = v56
                                        }
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = v29
                                    _ = v30
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = ((("foo bar" + ("foo bar" + ("foo bar" + "foo bar"))) + (("foo bar" + "\n\t") + ("\n\t" + "foo bar"))) + "foo bar")
                                    }
                                }
                            } else {
                                var v57 uint8 = 255
                                _ = v57
                            }
                            if ((4294967295 > 517) || (uint64(v28) > 4294967295)) {
                                if false {
                                    v23 -= (float32(((float64(float32((0.0 - v7))) - (v7 + 0.0)) - v7)) - float32(v7))
                                } else {
                                    v1 += (^(v1))
                                    if true {
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                        v2 = append(v2, "foo bar")
                                    } else {
                                        var v58 uint = v28
                                        _ = v58
                                    }
                                    v3 = -32768
                                }
                                v23 = (float32(-2.25) - (((v23 * v23) * float32(1e10)) - v23))
                                if false {
                                    var v59 uint8 = 255
                                    _ = v59
                                }
                                v2 = append(v2, "\n\t")
                            } else {
                                v3 -= (v3 & (v3 + (v3 * v3)))
                                v60 := GenS8{F0: uint(v1), F1: 9223372036854775807, F2: 0}
                                v61 := GenS5{F0: 0}
                                v2 = append(v2, "foo bar")
                                v2 = append(v2, "")
                                _ = v60
                                _ = v61
                            }
                            var v62 int = fi22
                            _ = v28
                            _ = v62
                        }
                        if true {
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v1, v23, v1)
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        } else {
                            if (v3 > 1) {
                                v63 := GenS9{F0: false, F1: (v3 | (v3 >> ((18446744073709551615) & 63))), F2: ("a" + "\n\t")}
                                _ = v63
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                var v64 uint = 1
                                _ = v64
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v23 += 1e10
                        }
                    }
                    fmt.Println(fi21, v7)
                    v65 := GenS1{F0: true, F1: -2147483648}
                    _ = fi22
                    _ = v23
                    _ = v65
                }
                v1 = v1
                _ = fi21
            }
        case "\n\t":
            v7 = v7
        }
    }
    v66 := GenS10{F0: -128}
    var v67 float64 = (((0.0 - 0.1) - v7) / ((v7 - v7) - v7))
    _ = len(v2)
    for fi68 := 0; fi68 < 9; fi68++ {
        v69 := GenS1{F0: (-9223372036854775808 != 42), F1: 627}
        var v70 float32 = ((1.5 - float32(3.14159)) / 1e10)
        v70 = 3.14159
        _ = fi68
        _ = v69
        _ = v70
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v66
    _ = v67
    _ = v7
}
