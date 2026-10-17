package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int8; F1 uint64; F2 int }
type GenS4 struct { F0 int; F1 string; F2 int8 }
type GenS5 struct { F0 int8; F1 uint16; F2 int8 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint; F1 int16; F2 int16 }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 uint64 }
func genFn0(p0 uint8, p1 bool, p2 int16) uint16 {
    fmt.Println(p1)
    fmt.Println(p2, p0, p0)
    p0 = (p0 | (p0 * (p0 % (((p0 + 2)) | 1))))
    v7 := GenS1{F0: 2, F1: "\n\t"}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 409
}
func genFn1(p0 uint, p1 int16, p2 int64) uint64 {
    var v8 uint16 = 748
    var v9 bool = true
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(v8)
        v11 := GenS1{F0: 0, F1: ("a" + "foo bar")}
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS1{F0: 18446744073709551615, F1: "a"}
            var v14 bool = (-80 > int32(p1))
            v15 := []int{(fi10 - 109), -1}
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    v8 = (v8 * (^(v8)))
    fmt.Println(p2, v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 269
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[int]int)
    var v3 uint8 = 0
    var v4 uint32 = 996
    v5 := []int{557}
    v6 := make(map[int]string)
    v16 := GenS1{F0: 0, F1: ("foo bar" + "")}
    var v17 float32 = (v1 / (v1 / (v1 * ((v1 + 1.5) * v1))))
    if false {
        switch -136 {
        case -9223372036854775808:
            v1 = v17
        case 0:
            v18 := v2[-699]
            _ = v18
        case 1:
            fmt.Println(v1, v17, v17)
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v19 := GenS2{F0: "hello", F1: 610}
        v20 := GenS2{F0: ("\n\t" + ("\n\t" + "")), F1: 1}
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        if (true || true) {
            v17 = v17
        } else {
            _ = len(v6)
        }
        _ = fi21
    }
    var v22 uint = 94
    for fi23 := 0; fi23 < 12; fi23++ {
        switch "" {
        case "":
            if (178 != v3) {
                var v24 uint8 = v3
                for fi25 := 0; fi25 < 3; fi25++ {
                    _ = len(v5)
                    v26 := GenS2{F0: "foo bar", F1: 1000}
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
        }
        _ = fi23
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "" {
    case "":
        var v27 int8 = 24
        _ = v27
    case "a":
        v22 = 42
    case "foo bar":
        if false {
            v28 := GenS1{F0: 673, F1: "a"}
            _ = v28
        } else {
            if len(v5) > 0 {
                v5[0] = -9223372036854775808
            }
            if false {
                fmt.Println(v17, v1, v22)
                v22 += 65535
            } else {
                if false {
                    v1 -= (((v17 - float32(0.1)) / v17) / float32(-2.25))
                }
                if (!(616 > v4)) {
                    switch "\n\t" {
                    case "\n\t":
                        if (32767 != -2147483648) {
                            fmt.Println(v3, v4, v17)
                        }
                    case "a":
                        switch -400 {
                        case 128:
                            var v29 uint32 = 65535
                            _ = v29
                        case 2147483647:
                            for fi30 := 0; fi30 < 20; fi30++ {
                                v3 -= (v3 | 72)
                                _ = fi30
                            }
                        case -276:
                            for fi31 := 0; fi31 < 9; fi31++ {
                                var v32 int32 = 672
                                _ = fi31
                                _ = v32
                            }
                        default:
                            var v33 float32 = (v17 - (float32(float64(float32(1.5))) - 1e10))
                            _ = v33
                        }
                    case "hello":
                        var v34 string = "\n\t"
                        _ = v34
                    default:
                        var v35 uint8 = (v3 % ((v3) | 1))
                        _ = v35
                    }
                }
                if false {
                    for fi36 := 0; fi36 < 18; fi36++ {
                        v37 := GenS1{F0: 18446744073709551615, F1: ("" + ("hello" + "foo bar"))}
                        _ = fi36
                        _ = v37
                    }
                    v1 = -2.25
                    var v38 uint16 = 255
                    _ = v38
                }
            }
        }
    }
    var v39 uint64 = 42
    switch ("" + ("foo bar" + ("foo bar" + ("" + "\n\t")))) {
    case "":
        if false {
            if false {
                if (v22 < (v22 << ((18446744073709551615) & 63))) {
                    if true {
                        if (!false) {
                            v1 = 1e10
                            fmt.Println(v39, v4, v39)
                        }
                        if false {
                            for fi40 := 0; fi40 < 11; fi40++ {
                                v41 := GenS3{F0: -123, F1: v39, F2: (fi40 ^ (fi40 >> (((v22 / ((v22) | 1))) & 63)))}
                                for fi42 := 0; fi42 < 15; fi42++ {
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v43 float64 = 3.14159
                                        _ = v43
                                    case "a":
                                        var v44 int8 = 0
                                        _ = v44
                                    default:
                                        var v45 bool = (952 == -1)
                                        _ = v45
                                    }
                                    delete(v6, 9223372036854775807)
                                    if (!true) {
                                        v3 += 2
                                        if false {
                                            v4 -= (^(v4))
                                        } else {
                                            v3 -= v3
                                            switch "foo bar" {
                                            case "hello":
                                                v17 -= float32((1.5 / (0.1 + (0.0 - (1e10 - ((float64((0.0 + 3.14159)) + 3.14159) * 0.1))))))
                                            }
                                            if false {
                                                v4 += 582
                                            } else {
                                                var v46 int = (fi40 ^ fi40)
                                                v47 := GenS2{F0: "\n\t", F1: 42}
                                                _ = v46
                                                _ = v47
                                            }
                                            v39 = uint64(v22)
                                            v48 := GenS3{F0: 127, F1: (v39 / ((4294967295) | 1)), F2: 32768}
                                            var v49 uint = v22
                                            _ = v48
                                            _ = v49
                                        }
                                    } else {
                                        fmt.Println(v1, v17)
                                    }
                                    _ = fi42
                                }
                                _ = fi40
                                _ = v41
                            }
                            for fi50 := 0; fi50 < 12; fi50++ {
                                var v51 uint64 = (v39 + genFn1(uint(v39), 0, 42))
                                _ = fi50
                                _ = v51
                            }
                        }
                    } else {
                        for fi52 := 0; fi52 < 2; fi52++ {
                            for fi53 := 0; fi53 < 13; fi53++ {
                                if false {
                                    fmt.Println(v17)
                                    v39 = v39
                                    switch fi53 {
                                    case 23:
                                        v54 := GenS1{F0: (^(v39)), F1: ("" + ("foo bar" + ("hello" + ("\n\t" + "foo bar"))))}
                                        _ = v54
                                    case -1:
                                        v2[821] = int(v3)
                                    case -710:
                                        var v55 int32 = 0
                                        _ = v55
                                    }
                                    fmt.Println(v3)
                                    if true {
                                        v17 = v17
                                        if (!((^(v3)) <= (^(v3)))) {
                                            if true {
                                                if len(v5) > 0 {
                                                    v5[0] = fi53
                                                }
                                            } else {
                                                v56 := GenS4{F0: fi53, F1: (("hello" + "foo bar") + "foo bar"), F2: -41}
                                                var v57 string = "\n\t"
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                _ = v56
                                                _ = v57
                                            }
                                        } else {
                                            if (-107 <= 63) {
                                                v2[(fi52 * (fi52 & (fi52 >> ((v22) & 63))))] = (fi52 + (fi52 - fi53))
                                                delete(v2, -774)
                                                fmt.Println(v3)
                                            }
                                            var v58 uint32 = 65535
                                            if false {
                                                fmt.Println(fi52)
                                                v1 += 3.14159
                                                v59 := v6[128]
                                                _ = v59
                                            }
                                            _ = v58
                                        }
                                    } else {
                                        switch (("hello" + ("\n\t" + ("\n\t" + ("\n\t" + "")))) + "\n\t") {
                                        case "\n\t":
                                            fmt.Println(v39, v3)
                                        case "hello":
                                            v3 -= 1
                                        case "a":
                                            _ = len(v5)
                                        default:
                                            v39 = (v39 / (((v39 * (^(v39)))) | 1))
                                        }
                                        v3 -= v3
                                        _ = len(v5)
                                    }
                                    var v60 bool = (v39 >= (v39 >> ((708) & 63)))
                                    _ = v60
                                }
                                _ = fi53
                            }
                            fmt.Println(v4)
                            _ = fi52
                        }
                    }
                    v39 = (v39 % (((v39 % (((^(v39))) | 1))) | 1))
                }
                fmt.Println(v4, v3, v3)
            } else {
                switch ("\n\t" + "hello") {
                case "\n\t":
                    if (true || false) {
                        switch -348 {
                        case -823:
                            v61 := GenS5{F0: -31, F1: 42, F2: 42}
                            _ = v61
                        case 42:
                            var v62 uint8 = v3
                            _ = v62
                        default:
                            if (!(2 <= 42)) {
                                switch "foo bar" {
                                case "\n\t":
                                    var v63 int16 = 0
                                    _ = v63
                                case "a":
                                    var v64 uint8 = 1
                                    _ = v64
                                case "foo bar":
                                    var v65 bool = (42 >= -2147483648)
                                    _ = v65
                                default:
                                    var v66 bool = false
                                    _ = v66
                                }
                                v67 := GenS1{F0: (v39 / ((v39) | 1)), F1: ("a" + "")}
                                _ = v67
                            } else {
                                if true {
                                    if false {
                                        var v68 uint16 = uint16(v22)
                                        _ = v68
                                    } else {
                                        var v69 int16 = 32767
                                        if true {
                                            v70 := GenS4{F0: 413, F1: ("hello" + "foo bar"), F2: 1}
                                            v39 -= (^(v39))
                                            v71 := GenS3{F0: -128, F1: genFn1((^(v22)), v69, 335), F2: 429}
                                            _ = v70
                                            _ = v71
                                        }
                                        _ = v69
                                    }
                                } else {
                                    var v72 string = (("" + ("\n\t" + (("a" + "\n\t") + "a"))) + "")
                                    var v73 uint32 = (v4 & (v4 % ((0) | 1)))
                                    _ = len(v2)
                                    _ = v72
                                    _ = v73
                                }
                                for fi74 := 0; fi74 < 5; fi74++ {
                                    var v75 uint64 = (^(v39))
                                    _ = fi74
                                    _ = v75
                                }
                                var v76 int32 = -940
                                _ = v76
                            }
                        }
                        for fi77 := 0; fi77 < 18; fi77++ {
                            fmt.Println(v39)
                            var v78 string = ("\n\t" + "")
                            var v79 int8 = 100
                            _ = fi77
                            _ = v78
                            _ = v79
                        }
                    }
                case "hello":
                    switch -470 {
                    case -790:
                        for fi80 := 0; fi80 < 6; fi80++ {
                            if true {
                                fmt.Println(v1, v1, v22)
                            }
                            v81 := GenS6{F0: "", F1: 0}
                            v2[(^(fi80))] = fi80
                            _ = fi80
                            _ = v81
                        }
                    case 127:
                        v82 := GenS6{F0: "foo bar", F1: -1}
                        _ = v82
                    case -9223372036854775808:
                        v83 := GenS7{F0: "hello", F1: ((!(false && (!((v39 / ((v39) | 1)) != genFn1((^(v22)), 42, 31))))) && (false || ((^(v4)) < v4)))}
                        _ = v83
                    }
                default:
                    if true {
                        var v84 bool = false
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v85 bool = ("hello" <= ("" + "\n\t"))
                        _ = v84
                        _ = v85
                    } else {
                        v1 += v17
                        fmt.Println(v39)
                        if len(v5) > 0 {
                            v5[0] = 128
                        }
                    }
                }
                fmt.Println(v1)
                if (true && true) {
                    _ = len(v6)
                    for fi86 := 0; fi86 < 7; fi86++ {
                        v87 := GenS8{F0: (v22 + uint(v3)), F1: -278, F2: int16(fi86)}
                        _ = fi86
                        _ = v87
                    }
                } else {
                    v88 := GenS6{F0: "", F1: 0}
                    v89 := GenS1{F0: v39, F1: ""}
                    v90 := GenS6{F0: "a", F1: 0}
                    _ = v88
                    _ = v89
                    _ = v90
                }
            }
            fmt.Println(v22, v39, v22)
            v91 := GenS9{F0: v4, F1: 42}
            _ = v91
        } else {
            v5 = append(v5, 32767)
            var v92 uint64 = v39
            v93 := GenS2{F0: "a", F1: -1}
            _ = v92
            _ = v93
        }
    default:
        v94 := GenS10{F0: (v39 - (v39 - 2))}
        _ = v94
    }
    for fi95 := 0; fi95 < 1; fi95++ {
        if ((v22 | v22) < uint(v3)) {
            var v96 int8 = -1
            _ = v96
        }
        _ = fi95
    }
    fmt.Println(v3, v3, v39)
    var v97 uint32 = (^(v4))
    switch 2147483647 {
    case 32767:
        var v98 int16 = -488
        _ = v98
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v97:", v97)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v97
}
