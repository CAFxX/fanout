package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 bool; F1 int64; F2 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 uint64; F1 uint32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 string }
func genFn0(p0 int32, p1 int) int {
    switch (-(p1)) {
    case 691:
        if true {
            var v4 bool = (false || (0.0 < 1.5))
            var v5 float64 = 3.14159
            fmt.Println(v5, p0, v4)
            _ = v4
            _ = v5
        } else {
            for fi6 := 0; fi6 < 16; fi6++ {
                if false {
                    v7 := []int{32767, -753, (fi6 % (((fi6 >> ((0) & 63))) | 1)), fi6}
                    p0 -= (p0 << ((255) & 63))
                    _ = v7
                }
                fmt.Println(p1, fi6)
                _ = fi6
            }
        }
    case 2147483647:
        var v8 int16 = -32768
        _ = v8
    default:
        v9 := []int{-1, p1, (p1 | 32767), p1}
        _ = v9
    }
    switch ("\n\t" + "foo bar") {
    case "a":
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 float64 = 0.0
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 float64 = float64(float32((v11 * v11)))
                var v14 uint32 = 714
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
        }
    default:
        var v15 bool = true
        _ = v15
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"foo bar", "foo bar", "\n\t"}
    var v3 uint = 255
    var v16 int = 927
    switch ("a" + "\n\t") {
    case "\n\t":
        v16 = v16
    case "foo bar":
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 uint = 0
            v18 = v18
            _ = fi17
            _ = v18
        }
    case "hello":
        var v19 string = "foo bar"
        _ = v19
    default:
        for fi20 := 0; fi20 < 11; fi20++ {
            switch ((("foo bar" + "a") + "foo bar") + ((("foo bar" + "a") + ("a" + ("foo bar" + "foo bar"))) + "")) {
            case "a":
                v16 += v16
            case "":
                if true {
                    v21 := GenS1{F0: 150, F1: 2147483647}
                    fmt.Println(fi20, v3)
                    _ = v21
                } else {
                    v22 := GenS2{F0: genFn0(-1, (fi20 * -1))}
                    if true {
                        v3 = 18446744073709551615
                        var v23 string = "a"
                        v3 = v3
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 10; fi24++ {
                        var v25 bool = false
                        _ = fi24
                        _ = v25
                    }
                    _ = v22
                }
            default:
                if (32767 >= 32767) {
                    fmt.Println(v16, fi20, v16)
                    var v26 bool = true
                    _ = v26
                }
            }
            _ = fi20
        }
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := GenS3{F0: (60 <= 42)}
        var v29 int8 = 100
        _ = fi27
        _ = v28
        _ = v29
    }
    switch "a" {
    case "hello":
        if true {
            v30 := GenS1{F0: 2, F1: -2147483648}
            _ = v30
        }
    case "":
        switch "\n\t" {
        case "foo bar":
            for fi31 := 0; fi31 < 10; fi31++ {
                v3 += (v3 & v3)
                for fi32 := 0; fi32 < 6; fi32++ {
                    var v33 int8 = -128
                    switch "\n\t" {
                    case "foo bar":
                        v34 := GenS4{F0: "a"}
                        _ = v34
                    case "":
                        var v35 int32 = -1
                        _ = v35
                    }
                    v36 := GenS3{F0: (v33 >= (v33 + v33))}
                    var v37 uint = (v3 & v3)
                    _ = len(v1)
                    _ = fi32
                    _ = v33
                    _ = v36
                    _ = v37
                }
                if true {
                    var v38 float32 = -2.25
                    var v39 float64 = 1e10
                    _ = v38
                    _ = v39
                } else {
                    fmt.Println(v3, v16, v3)
                    for fi40 := 0; fi40 < 12; fi40++ {
                        var v41 int16 = 42
                        switch "a" {
                        case "\n\t":
                            v16 = v16
                        }
                        fmt.Println(v3, v16)
                        _ = fi40
                        _ = v41
                    }
                }
                _ = fi31
            }
        case "hello":
            for fi42 := 0; fi42 < 3; fi42++ {
                var v43 int64 = -9223372036854775808
                for fi44 := 0; fi44 < 15; fi44++ {
                    if (2 != 255) {
                        var v45 int32 = 0
                        if true {
                            switch (fi44 / ((fi44) | 1)) {
                            case 0:
                                fmt.Println(v16, fi44, v43)
                            case 1:
                                fmt.Println(v45, v3)
                            case 745:
                                v3 = 65535
                            default:
                                switch "hello" {
                                case "a":
                                    if ((!false) && false) {
                                        var v46 uint64 = 42
                                        fmt.Println(fi44, v3)
                                        _ = v46
                                    }
                                case "foo bar":
                                    v47 := GenS3{F0: (!false)}
                                    _ = v47
                                default:
                                    if false {
                                        v2 = append(v2, "foo bar")
                                        v48 := GenS5{F0: ("" + "foo bar"), F1: (v45 * int32(v3))}
                                        _ = v48
                                    }
                                    var v49 int64 = v43
                                    _ = v49
                                }
                            }
                        }
                        var v50 uint = 255
                        v16 = (-(fi44))
                        _ = v45
                        _ = v50
                    } else {
                        switch (^(fi44)) {
                        case 128:
                            switch (fi44 % ((1) | 1)) {
                            case -1:
                                if true {
                                    v51 := GenS6{F0: (!false), F1: 42}
                                    v1[2147483647] = (false || false)
                                    _ = v51
                                } else {
                                    v52 := GenS7{F0: (!(500 < 2147483648)), F1: v43, F2: ((v3 * (v3 >> (((^(v3))) & 63))) < 2)}
                                    switch genFn0(42, (-(v16))) {
                                    case 42:
                                        v1[4294967295] = (164 > 10)
                                    case 4294967295:
                                        v3 = (v3 ^ (^(v3)))
                                    case 526:
                                        _ = len(v2)
                                    }
                                    _ = v52
                                }
                            case -594:
                                v16 += genFn0(-584, (fi44 + fi44))
                            case 32768:
                                fmt.Println(fi42, v16, fi42)
                            }
                            v3 -= v3
                            if len(v2) > 0 {
                                v2[0] = ("" + "foo bar")
                            }
                            var v53 uint8 = 128
                            _ = v53
                        case 0:
                            var v54 string = "foo bar"
                            _ = v54
                        case 9223372036854775807:
                            if (2 <= 945) {
                                v3 -= v3
                                var v55 int16 = 32767
                                v56 := GenS8{F0: 42}
                                _ = v55
                                _ = v56
                            } else {
                                v57 := GenS9{F0: 1, F1: 830}
                                var v58 uint32 = 968
                                _ = v57
                                _ = v58
                            }
                        default:
                            var v59 string = "hello"
                            _ = v59
                        }
                        v16 += genFn0(-1, v16)
                        if (true && (genFn0(42, -1000) <= v16)) {
                            var v60 bool = (!false)
                            v43 -= v43
                            var v61 string = ""
                            v62 := GenS9{F0: 2, F1: -373}
                            _ = v60
                            _ = v61
                            _ = v62
                        }
                        if true {
                            var v63 string = "hello"
                            if (32768 > (fi44 - -1)) {
                                var v64 int64 = v43
                                v65 := GenS10{F0: 4294967295, F1: 601}
                                v2 = append(v2, "\n\t")
                                var v66 uint32 = 219
                                var v67 uint16 = 255
                                _ = v64
                                _ = v65
                                _ = v66
                                _ = v67
                            } else {
                                v68 := GenS11{F0: 4294967295}
                                var v69 int32 = -2147483648
                                if ((((v69 | v69) == v69) && true) && true) {
                                    delete(v1, 343)
                                    var v70 int32 = v69
                                    var v71 int32 = (v70 << (((^(v3))) & 63))
                                    _ = v70
                                    _ = v71
                                }
                                _ = v68
                                _ = v69
                            }
                            var v72 float64 = 1e10
                            _ = v63
                            _ = v72
                        } else {
                            v43 += -1
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v43 = (v43 - -1)
                        var v73 float32 = (1.5 + (1.5 + 0.1))
                        v16 -= 4294967295
                        _ = v73
                    }
                    var v74 uint32 = 65535
                    _ = fi44
                    _ = v74
                }
                fmt.Println(v16)
                _ = fi42
                _ = v43
            }
        default:
            v3 = (v3 * (v3 * (v3 - (v3 / ((v3) | 1)))))
        }
    default:
        var v75 uint32 = 4294967295
        _ = v75
    }
    var v76 string = ((("" + "\n\t") + "") + "hello")
    v76 = "foo bar"
    fmt.Println(v76)
    var v77 int8 = 100
    _ = len(v1)
    v3 += (v3 + 1)
    fmt.Println(v76, v76, v16)
    switch v16 {
    case 387:
        v16 = v16
    case 9223372036854775807:
        var v78 int16 = 32767
        _ = v78
    case 352:
        if (!false) {
            var v79 int8 = v77
            v16 += v16
            _ = v79
        } else {
            for fi80 := 0; fi80 < 11; fi80++ {
                _ = len(v1)
                _ = fi80
            }
        }
    default:
        v76 = "foo bar"
    }
    switch 712 {
    case 2147483647:
        v76 = v76
    case 127:
        if (!(548 < v3)) {
            if true {
                var v81 uint32 = 0
                v82 := v1[v81]
                _ = v81
                _ = v82
            } else {
                var v83 int32 = 42
                var v84 bool = true
                _ = v83
                _ = v84
            }
            if (1e10 >= ((3.14159 / (3.14159 - (0.0 * (0.1 + (0.1 - float64((float32((((0.0 / 1e10) - (0.1 * (0.1 * (1e10 / 1.0)))) + (3.14159 - -2.25))) - (1e10 * 1e10)))))))) / (0.0 - 1.5))) {
                v85 := GenS12{F0: v76}
                v76 = "foo bar"
                fmt.Println(v77, v77, v3)
                _ = v85
            }
        } else {
            for fi86 := 0; fi86 < 10; fi86++ {
                var v87 uint8 = 0
                var v88 int8 = (v77 / ((v77) | 1))
                _ = fi86
                _ = v87
                _ = v88
            }
            if ((42 >= -1) || true) {
                if (!true) {
                    if len(v2) > 0 {
                        v2[0] = v76
                    }
                } else {
                    var v89 uint32 = 45
                    var v90 int = genFn0(2147483647, 25)
                    _ = len(v1)
                    _ = v89
                    _ = v90
                }
                switch ((v76 + v76) + ("a" + ("foo bar" + v76))) {
                case "a":
                    for fi91 := 0; fi91 < 14; fi91++ {
                        v92 := GenS8{F0: 2}
                        switch v76 {
                        case "":
                            v93 := v1[1]
                            _ = v93
                        }
                        _ = fi91
                        _ = v92
                    }
                case "foo bar":
                    v94 := GenS8{F0: 42}
                    _ = v94
                }
            } else {
                v95 := GenS1{F0: 0, F1: -206}
                v76 = (v76 + ("a" + v76))
                _ = len(v2)
                _ = v95
            }
            v16 = (v16 * v16)
        }
    case -458:
        var v96 uint32 = 0
        _ = v96
    }
    if (0 >= (-(v77))) {
        var v97 uint32 = 0
        var v98 float64 = 1e10
        _ = v97
        _ = v98
    } else {
        var v99 int = v16
        _ = v99
    }
    v76 = v76
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v76
    _ = v77
}
