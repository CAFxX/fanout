package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32; F1 int64; F2 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 string; F1 int32; F2 uint8 }
type GenS11 struct { F0 string; F1 uint16 }
func genFn0(p0 int16, p1 int16) int64 {
    p0 = (^(p0))
    var v4 int8 = 1
    var v5 bool = (false || (!true))
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(v4)
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint = 4294967295
            _ = fi7
            _ = v8
        }
        var v9 uint16 = 206
        _ = fi6
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint8, p1 string, p2 int16) int32 {
    v10 := GenS1{F0: 42, F1: 1}
    v11 := []int{}
    for fi12 := 0; fi12 < 19; fi12++ {
        p1 = ((((p1 + ((("hello" + p1) + ((p1 + (("hello" + p1) + ("foo bar" + "hello"))) + "hello")) + p1)) + p1) + p1) + p1)
        if (p1 <= "a") {
            if (!(uint32(p0) == 42)) {
                for fi13 := 0; fi13 < 13; fi13++ {
                    if true {
                        var v14 bool = (!false)
                        p1 = "foo bar"
                        _ = len(v11)
                        _ = v14
                    }
                    _ = fi13
                }
                fmt.Println(p0, p2, p2)
            } else {
                for fi15 := 0; fi15 < 16; fi15++ {
                    if len(v11) > 0 {
                        v11[0] = fi12
                    }
                    _ = len(v11)
                    v16 := GenS2{F0: false}
                    _ = fi15
                    _ = v16
                }
            }
        }
        _ = fi12
    }
    v17 := GenS2{F0: ((p1 + ("hello" + p1)) >= p1)}
    v18 := GenS2{F0: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    return -2147483648
}

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v19 := GenS3{F0: 2, F1: 0}
    if (-1 < 1000) {
        for fi20 := 0; fi20 < 19; fi20++ {
            switch ("foo bar" + ("" + "hello")) {
            case "\n\t":
                switch "a" {
                case "":
                    switch 1 {
                    case 1:
                        for fi21 := 0; fi21 < 18; fi21++ {
                            var v22 float32 = 3.14159
                            delete(v3, ("" + ((("" + "") + ("foo bar" + "")) + (("" + "") + ("hello" + (("foo bar" + ("a" + ("" + (("" + "a") + "\n\t")))) + ("" + "\n\t")))))))
                            _ = fi21
                            _ = v22
                        }
                    case -2147483648:
                        v1 = (v1 * genFn0(0, 722))
                    case -754:
                        v1 = (^(v1))
                    }
                case "\n\t":
                    switch "hello" {
                    case "":
                        _ = len(v2)
                    case "a":
                        switch ("hello" + "foo bar") {
                        case "":
                            switch (fi20 * fi20) {
                            case 471:
                                v23 := GenS4{F0: 2}
                                _ = v23
                            case 42:
                                var v24 uint16 = 255
                                _ = v24
                            }
                        case "hello":
                            for fi25 := 0; fi25 < 14; fi25++ {
                                fmt.Println(v1, fi20, fi20)
                                v26 := []int{}
                                fmt.Println(v1, v1)
                                _ = fi25
                                _ = v26
                            }
                        default:
                            switch (fi20 * fi20) {
                            case 109:
                                v27 := []int{(fi20 + fi20), (^(fi20)), -505, fi20}
                                _ = v27
                            case -9223372036854775808:
                                v28 := GenS3{F0: 513, F1: 210}
                                _ = v28
                            }
                        }
                    case "\n\t":
                        switch "foo bar" {
                        case "\n\t":
                            if true {
                                var v29 int8 = -89
                                var v30 int8 = (v29 << ((404) & 63))
                                _ = v29
                                _ = v30
                            }
                        case "hello":
                            for fi31 := 0; fi31 < 10; fi31++ {
                                var v32 uint = 340
                                var v33 uint16 = 1
                                if (true || false) {
                                    v34 := []int{128}
                                    switch (fi20 - fi31) {
                                    case 768:
                                        if (((v32 * v32) >= v32) && (0 < genFn1(238, "a", -978))) {
                                            v32 += (v32 * uint(v33))
                                        }
                                    default:
                                        v35 := v2[2147483648]
                                        _ = v35
                                    }
                                    _ = v34
                                } else {
                                    v36 := GenS2{F0: (false && ((v33 - v33) != v33))}
                                    fmt.Println(v33, v33, v33)
                                    _ = v36
                                }
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        }
                    }
                case "hello":
                    if false {
                        switch "hello" {
                        case "a":
                            if false {
                                v37 := GenS5{F0: v1}
                                if true {
                                    if true {
                                        for fi38 := 0; fi38 < 16; fi38++ {
                                            v1 += -1
                                            _ = fi38
                                        }
                                        v39 := []int{}
                                        _ = v39
                                    }
                                    var v40 string = "\n\t"
                                    fmt.Println(v40)
                                    _ = v40
                                }
                                v41 := GenS4{F0: 47}
                                _ = v37
                                _ = v41
                            }
                        case "":
                            delete(v3, "foo bar")
                        default:
                            var v42 float64 = 1.5
                            _ = v42
                        }
                    } else {
                        v1 -= v1
                        v43 := GenS6{F0: 18446744073709551615, F1: true}
                        _ = v43
                    }
                }
            case "":
                v1 = v1
            }
            for fi44 := 0; fi44 < 17; fi44++ {
                var v45 int64 = (v1 - (v1 / ((genFn0(-1, 42)) | 1)))
                v1 = -1
                _ = fi44
                _ = v45
            }
            v46 := GenS7{F0: 0}
            _ = fi20
            _ = v46
        }
        var v47 float32 = 1.5
        _ = v47
    } else {
        for fi48 := 0; fi48 < 9; fi48++ {
            fmt.Println(v1, v1)
            _ = fi48
        }
        var v49 int64 = v1
        v50 := []int{4294967295}
        _ = v49
        _ = v50
    }
    if (true && true) {
        fmt.Println(v1)
    }
    if false {
        if true {
            fmt.Println(v1)
        } else {
            v51 := []int{128}
            switch 1 {
            case 4294967295:
                v52 := v3[("foo bar" + ("foo bar" + "\n\t"))]
                _ = v52
            case -888:
                switch "a" {
                case "foo bar":
                    v53 := GenS8{F0: -662, F1: v1, F2: "a"}
                    _ = v53
                case "a":
                    var v54 uint8 = 128
                    _ = v54
                default:
                    fmt.Println(v1)
                }
            case 42:
                v2[1] = ("" + "a")
            }
            _ = v51
        }
    }
    if true {
        switch 32768 {
        case -317:
            var v55 float64 = 1e10
            _ = v55
        default:
            v56 := []int{244, -9223372036854775808, -784}
            _ = v56
        }
    } else {
        v1 = genFn0(804, -1)
        switch 4294967295 {
        case 2147483647:
            var v57 string = "foo bar"
            _ = v57
        case 32768:
            for fi58 := 0; fi58 < 10; fi58++ {
                fmt.Println(fi58, fi58)
                var v59 int64 = (^(v1))
                _ = fi58
                _ = v59
            }
        case 1:
            v60 := GenS9{F0: (true && false)}
            _ = v60
        default:
            if (-32768 > 32767) {
                v1 = (^(v1))
            } else {
                if true {
                    var v61 uint64 = 537
                    fmt.Println(v61, v1)
                    _ = v61
                } else {
                    switch ("" + "foo bar") {
                    case "\n\t":
                        var v62 int16 = 1000
                        _ = v62
                    case "a":
                        switch 621 {
                        case 0:
                            v1 = v1
                        case 32767:
                            if (true || false) {
                                var v63 uint64 = 2
                                v64 := v2[42]
                                fmt.Println(v63, v64)
                                _ = v63
                                _ = v64
                            } else {
                                v65 := GenS6{F0: 761, F1: ((0 < 112) || false)}
                                for fi66 := 0; fi66 < 7; fi66++ {
                                    v67 := GenS10{F0: ("a" + "foo bar"), F1: -1, F2: 255}
                                    v1 -= v1
                                    v1 += 42
                                    _ = fi66
                                    _ = v67
                                }
                                _ = v65
                            }
                        case 2147483647:
                            var v68 int8 = -1
                            _ = v68
                        }
                    default:
                        v1 = 9223372036854775807
                    }
                    var v69 int64 = (v1 ^ (v1 << ((51) & 63)))
                    _ = v69
                }
            }
        }
    }
    delete(v3, "\n\t")
    var v70 bool = false
    var v71 uint32 = 644
    v71 -= (v71 + v71)
    if v70 {
        if (376 > 42) {
            var v72 uint64 = 1
            _ = v72
        }
        var v73 uint8 = 128
        _ = v73
    } else {
        for fi74 := 0; fi74 < 15; fi74++ {
            v70 = v70
            _ = fi74
        }
        v70 = v70
        for fi75 := 0; fi75 < 14; fi75++ {
            v76 := GenS4{F0: 127}
            v71 = v71
            _ = fi75
            _ = v76
        }
    }
    for fi77 := 0; fi77 < 5; fi77++ {
        delete(v2, 2)
        var v78 float32 = 1e10
        v79 := []int{fi77, fi77}
        _ = fi77
        _ = v78
        _ = v79
    }
    v71 = (^(v71))
    switch "hello" {
    case "\n\t":
        fmt.Println(v1, v1, v71)
    case "":
        var v80 bool = (v71 >= v71)
        _ = v80
    default:
        v81 := GenS5{F0: 9223372036854775807}
        _ = v81
    }
    switch 0 {
    case 0:
        v1 = v1
    case 4294967295:
        v82 := v3["\n\t"]
        _ = v82
    case 9223372036854775807:
        for fi83 := 0; fi83 < 5; fi83++ {
            var v84 int = (fi83 % ((1) | 1))
            if ((3.14159 / 1.0) == (float32(((0.1 / ((0.0 - 0.1) + (0.0 / (1.5 / -2.25)))) - 1.5)) * float32((3.14159 + 1.5)))) {
                var v85 int32 = 42
                fmt.Println(v71, fi83)
                v3[("hello" + (("a" + "") + ("\n\t" + "a")))] = (fi83 >> ((255) & 63))
                _ = v85
            }
            if true {
                v86 := GenS11{F0: "hello", F1: 255}
                var v87 string = ((("hello" + ("a" + (("\n\t" + "foo bar") + "\n\t"))) + "") + "a")
                _ = v86
                _ = v87
            }
            _ = fi83
            _ = v84
        }
    default:
        var v88 string = "a"
        _ = v88
    }
    v1 = (v1 + (v1 << ((662) & 63)))
    if (!(42 < -563)) {
        var v89 uint32 = (v71 * (^(v71)))
        _ = v89
    } else {
        fmt.Println(v1)
        v90 := GenS7{F0: 128}
        v91 := v2[v71]
        _ = v90
        _ = v91
    }
    v71 -= (^(v71))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v70
    _ = v71
}
