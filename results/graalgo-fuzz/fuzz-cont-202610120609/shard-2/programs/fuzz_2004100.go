package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int16; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS5 struct { F0 int; F1 bool; F2 int32 }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 int64; F1 int16; F2 int32 }
type GenS8 struct { F0 uint8 }
func genFn0() uint32 {
    switch "foo bar" {
    case "hello":
        v4 := []int{32767, -2147483648}
        _ = v4
    case "a":
        var v5 int64 = 0
        _ = v5
    case "foo bar":
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 int8 = 87
            v8 := GenS1{F0: (v7 >> ((4294967295) & 63))}
            v9 := []int{}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    default:
        v10 := GenS1{F0: -54}
        _ = v10
    }
    var v11 int16 = -1
    _ = v11
    return 2147483648
}
func genFn1(p0 int) int {
    p0 -= p0
    for fi12 := 0; fi12 < 3; fi12++ {
        p0 -= (fi12 % (((^(fi12))) | 1))
        _ = fi12
    }
    var v13 int16 = 1
    _ = p0
    _ = v13
    return p0
}

func main() {
    var v1 string = "hello"
    var v2 float64 = -2.25
    v3 := make(map[string]bool)
    var v14 int16 = 1
    if (!true) {
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 float32 = 1.5
            _ = fi15
            _ = v16
        }
    } else {
        fmt.Println(v2)
        switch ((v1 + (("hello" + v1) + (v1 + ("hello" + "hello")))) + "foo bar") {
        case "":
            switch 42 {
            case 180:
                v14 += v14
            case 32768:
                fmt.Println(v2, v14, v2)
            case -720:
                if false {
                    var v17 uint64 = 258
                    _ = v17
                }
            }
        }
        fmt.Println(v2, v14)
    }
    v14 -= (v14 | (v14 << ((65535) & 63)))
    switch -1 {
    case 128:
        v18 := []int{42, -9223372036854775808, 0}
        _ = v18
    case 127:
        switch 809 {
        case 42:
            if true {
                if false {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        var v20 int32 = -1
                        _ = fi19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 18; fi21++ {
                        var v22 uint8 = 250
                        _ = fi21
                        _ = v22
                    }
                    v23 := GenS2{F0: 1, F1: (-(v14)), F2: 65535}
                    _ = v23
                } else {
                    v24 := GenS3{F0: false}
                    for fi25 := 0; fi25 < 5; fi25++ {
                        switch -962 {
                        case 1:
                            var v26 uint8 = 65
                            _ = v26
                        }
                        var v27 uint8 = 1
                        var v28 int = (^(fi25))
                        _ = fi25
                        _ = v27
                        _ = v28
                    }
                    if (18446744073709551615 >= 42) {
                        fmt.Println(v2)
                        fmt.Println(v1, v14, v2)
                        var v29 int16 = (-(v14))
                        _ = v29
                    } else {
                        fmt.Println(v1)
                    }
                    _ = v24
                }
                v14 = v14
            }
        default:
            _ = len(v3)
        }
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        if (!true) {
            var v31 bool = (!((113 == -128) || (!(1e10 >= 1e10))))
            var v32 uint32 = 2147483648
            switch (^(fi30)) {
            case -9223372036854775808:
                v1 = (v1 + v1)
            case -1:
                v33 := []int{genFn1(448), (fi30 ^ genFn1(fi30)), genFn1(genFn1(genFn1((fi30 + (fi30 | -975))))), (fi30 & genFn1(128))}
                _ = v33
            default:
                v32 -= 2
            }
            _ = v31
            _ = v32
        } else {
            fmt.Println(v2, fi30)
            v14 -= v14
            var v34 float32 = 1e10
            _ = v34
        }
        var v35 int = fi30
        var v36 int = (^(fi30))
        _ = fi30
        _ = v35
        _ = v36
    }
    v37 := []int{42, 32767}
    switch ("" + (v1 + "\n\t")) {
    case "a":
        fmt.Println(v1, v2, v2)
    default:
        v2 = v2
    }
    if false {
        if false {
            if (1 <= 484) {
                if len(v37) >= 2 {
                    v37 = v37[1:len(v37)]
                }
                v38 := GenS4{F0: 32767, F1: (false || (v2 < v2)), F2: 1}
                _ = v38
            } else {
                var v39 int64 = 535
                var v40 string = v1
                v41 := GenS4{F0: 0, F1: (!true), F2: 0}
                _ = v39
                _ = v40
                _ = v41
            }
        } else {
            switch 542 {
            case -895:
                v2 = float64(((1e10 * 0.0) * 0.0))
            case 127:
                fmt.Println(v14)
            default:
                switch 32767 {
                case 557:
                    switch 32767 {
                    case 0:
                        v1 = (v1 + v1)
                    case 32767:
                        var v42 uint32 = 42
                        _ = v42
                    case -2147483648:
                        if true {
                            var v43 uint8 = 9
                            if true {
                                v2 -= (float64((-2.25 + 0.0)) + v2)
                            }
                            var v44 uint = 4294967295
                            _ = v43
                            _ = v44
                        }
                    default:
                        var v45 bool = (!(!(false && (1 == 9223372036854775807))))
                        _ = v45
                    }
                case 4294967295:
                    v14 = (v14 ^ -32768)
                case 42:
                    fmt.Println(v14, v2)
                default:
                    if len(v37) > 0 {
                        v37[0] = -843
                    }
                }
            }
        }
    }
    var v46 uint32 = 42
    var v47 bool = true
    fmt.Println(v2, v47, v14)
    v48 := v3[("hello" + "a")]
    v49 := GenS4{F0: 1, F1: (((-9223372036854775808 == 9223372036854775807) || (18446744073709551615 > 683)) || false), F2: 2}
    switch 32767 {
    case 2147483647:
        v50 := GenS4{F0: 32767, F1: (!(!((v1 + "a") > v1))), F2: 0}
        _ = v50
    case 1:
        if (!(!v48)) {
            switch 1 {
            case 127:
                v46 += (v46 - v46)
            }
        }
    }
    fmt.Println(v47, v1, v2)
    v14 -= (v14 >> ((uint(v46)) & 63))
    for fi51 := 0; fi51 < 8; fi51++ {
        if v48 {
            switch 42 {
            case 32768:
                var v52 uint32 = 226
                _ = v52
            case 98:
                v1 = ("\n\t" + (("a" + v1) + v1))
            case 1:
                if (((!false) || true) || (42 > 0)) {
                    for fi53 := 0; fi53 < 16; fi53++ {
                        if v47 {
                            v47 = (float32(((v2 - (float64((0.0 + -2.25)) - (v2 + v2))) * v2)) != (0.1 - 1e10))
                            v47 = ((65535 > 4294967295) && (fi53 <= (fi53 | fi53)))
                            switch fi51 {
                            case 127:
                                v54 := GenS3{F0: v47}
                                _ = v54
                            case -1:
                                v2 += (0.1 + (-2.25 / 0.1))
                            case 9223372036854775807:
                                v48 = v48
                            }
                        } else {
                            switch v1 {
                            case "a":
                                v2 -= v2
                            case "\n\t":
                                if len(v37) > 0 {
                                    v37[0] = (^(fi53))
                                }
                            case "hello":
                                if v47 {
                                    if (0 == 1) {
                                        _ = len(v3)
                                        fmt.Println(fi51, v47)
                                    } else {
                                        var v55 uint64 = 0
                                        _ = v55
                                    }
                                    v48 = v47
                                }
                            default:
                                switch 1 {
                                case 4294967295:
                                    v1 = ""
                                case 39:
                                    var v56 string = (v1 + v1)
                                    _ = v56
                                case -767:
                                    _ = len(v3)
                                default:
                                    var v57 int = 32767
                                    _ = v57
                                }
                            }
                            v48 = (!((405 != 1) || (!(v1 < v1))))
                            switch (fi53 - -872) {
                            case -502:
                                v47 = (!(fi51 > (fi51 - (fi51 >> ((18446744073709551615) & 63)))))
                            }
                        }
                        fmt.Println(fi51, v47, v47)
                        v1 = ("" + "")
                        switch (-(fi51)) {
                        case -2147483648:
                            if v47 {
                                switch (v1 + ("hello" + v1)) {
                                case "\n\t":
                                    v14 = v14
                                }
                                var v58 int = (fi51 - (fi51 + (fi51 % (((fi53 / (((fi51 ^ (fi51 % ((-624) | 1)))) | 1))) | 1))))
                                _ = v58
                            } else {
                                v59 := v3[("hello" + v1)]
                                _ = v59
                            }
                        case -204:
                            v60 := GenS5{F0: (fi53 & (fi51 << ((694) & 63))), F1: false, F2: 404}
                            _ = v60
                        default:
                            v61 := GenS6{F0: 413, F1: 9223372036854775807}
                            _ = v61
                        }
                        v62 := GenS3{F0: (((((1e10 * ((-2.25 * float32((1e10 * v2))) / 1.0)) + float32(1.5)) / 1.5) / 0.1) == 3.14159)}
                        _ = fi53
                        _ = v62
                    }
                    var v63 int64 = -9223372036854775808
                    v64 := GenS7{F0: 1, F1: (^(v14)), F2: -82}
                    _ = v63
                    _ = v64
                } else {
                    if v47 {
                        v65 := v3[""]
                        if (-378 >= -2147483648) {
                            var v66 uint64 = 680
                            _ = v66
                        } else {
                            _ = len(v37)
                            _ = len(v3)
                        }
                        _ = v65
                    } else {
                        v1 = v1
                        if v47 {
                            var v67 uint64 = 18446744073709551615
                            _ = len(v3)
                            if v47 {
                                fmt.Println(v67, v67, v47)
                                v68 := GenS2{F0: -1, F1: v14, F2: 11}
                                _ = v68
                            } else {
                                v69 := GenS8{F0: uint8(v67)}
                                var v70 uint64 = (v67 | 42)
                                v14 -= v14
                                _ = v69
                                _ = v70
                            }
                            _ = v67
                        }
                        fmt.Println(v1, v1, fi51)
                    }
                    if v48 {
                        var v71 float32 = 1e10
                        v48 = (((100 != 127) || (2 <= 255)) && ((fi51 * -9223372036854775808) != fi51))
                        _ = v71
                    } else {
                        delete(v3, v1)
                    }
                }
            default:
                var v72 int64 = 1
                _ = v72
            }
            if (!(v2 <= 0.0)) {
                var v73 uint = 18446744073709551615
                if v47 {
                    if (!false) {
                        v74 := GenS1{F0: 42}
                        for fi75 := 0; fi75 < 20; fi75++ {
                            v1 = v1
                            fmt.Println(v46, v47)
                            _ = fi75
                        }
                        _ = v74
                    }
                } else {
                    if v47 {
                        var v76 bool = v48
                        v37 = append(v37, (-(fi51)))
                        _ = v76
                    }
                }
                var v77 uint = (v73 & 42)
                _ = v73
                _ = v77
            }
            if (true && false) {
                v1 = v1
            }
        } else {
            v2 -= v2
            v78 := GenS8{F0: 240}
            _ = v78
        }
        v2 = v2
        var v79 int32 = 42
        _ = fi51
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
