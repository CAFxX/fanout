package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint; F2 int }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int }
func genFn0(p0 uint16, p1 int8) uint8 {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 uint16 = (p0 * (p0 ^ p0))
        for fi6 := 0; fi6 < 10; fi6++ {
            v7 := GenS1{F0: 42, F1: fi4, F2: (!(false && ((!(v5 == 598)) && (false || (!(0 < (^(v5))))))))}
            if (3.14159 >= float64(0.1)) {
                switch "hello" {
                case "hello":
                    v5 = 2
                case "\n\t":
                    if (!(1 <= -1)) {
                        switch "\n\t" {
                        case "a":
                            p0 = (p0 - 2)
                        case "":
                            v5 += p0
                        case "hello":
                            v8 := GenS2{F0: -697}
                            _ = v8
                        }
                        p1 = (p1 / ((0) | 1))
                        p0 += 1
                    }
                }
                var v9 float32 = (float32(-2.25) * 0.1)
                _ = v9
            }
            var v10 int64 = -1
            _ = fi6
            _ = v10
            _ = v7
        }
        switch ("hello" + ("" + "hello")) {
        case "a":
            for fi11 := 0; fi11 < 19; fi11++ {
                switch fi4 {
                case 0:
                    p0 = v5
                case 9223372036854775807:
                    var v12 uint32 = 2147483648
                    _ = v12
                case 32767:
                    v13 := []int{}
                    _ = v13
                }
                if true {
                    switch (fi11 % ((fi4) | 1)) {
                    case 127:
                        if (p0 <= 911) {
                            p0 -= (^(p0))
                            var v14 uint8 = 127
                            fmt.Println(p0, v5)
                            v5 = 656
                            var v15 uint = 255
                            _ = v14
                            _ = v15
                        } else {
                            if false {
                                v16 := GenS3{F0: fi4, F1: 127}
                                v17 := GenS4{F0: 277, F1: 4294967295}
                                if (true || true) {
                                    v18 := []int{(fi4 - fi11), (-(fi11)), (fi4 << ((111) & 63)), 0}
                                    _ = v18
                                }
                                if true {
                                    if true {
                                        fmt.Println(v5, p1)
                                    } else {
                                        if false {
                                            fmt.Println(p1, p1, fi11)
                                            var v19 uint64 = 4294967295
                                            var v20 uint8 = 115
                                            _ = v19
                                            _ = v20
                                        } else {
                                            v21 := GenS1{F0: (p1 % ((p1) | 1)), F1: fi11, F2: true}
                                            var v22 uint16 = v5
                                            var v23 float64 = (0.1 - -2.25)
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                    v24 := []int{(^(fi11)), (fi11 & -9223372036854775808), fi4, fi11}
                                    var v25 float32 = 3.14159
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    _ = v24
                                    _ = v25
                                } else {
                                    p1 = -43
                                    v26 := GenS4{F0: 2, F1: 42}
                                    _ = v26
                                }
                                switch (^(fi11)) {
                                case 1:
                                    fmt.Println(p0, fi11, v5)
                                }
                                _ = v16
                                _ = v17
                            } else {
                                fmt.Println(fi11, p0)
                                var v27 uint16 = p0
                                _ = v27
                            }
                            var v28 float32 = ((1.5 / (-2.25 + 1e10)) + ((0.0 + 1.5) / 0.1))
                            _ = v28
                        }
                    }
                }
                switch "hello" {
                case "":
                    p1 -= p1
                case "hello":
                    fmt.Println(fi4)
                default:
                    v29 := GenS2{F0: 1}
                    _ = v29
                }
                _ = fi11
            }
        case "foo bar":
            v30 := GenS5{F0: -1, F1: 1, F2: fi4}
            _ = v30
        case "\n\t":
            p1 += p1
        }
        _ = fi4
        _ = v5
    }
    v31 := []int{}
    var v32 float32 = float32(float64((((0.0 + 0.1) * (-2.25 + -2.25)) - 0.1)))
    _ = p0
    _ = p1
    _ = v31
    _ = v32
    return 42
}
func genFn1(p0 int8, p1 uint) uint16 {
    fmt.Println(p1)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{191, 4294967295, 2147483647}
    var v3 float32 = float32(1.5)
    if false {
        v3 -= (0.1 / 0.1)
        var v33 string = ("foo bar" + "\n\t")
        if true {
            fmt.Println(v33, v33)
        } else {
            v34 := GenS2{F0: 42}
            for fi35 := 0; fi35 < 12; fi35++ {
                if (127 == 0) {
                    v36 := GenS5{F0: -380, F1: 255, F2: 2147483647}
                    v37 := v1["hello"]
                    for fi38 := 0; fi38 < 17; fi38++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch v33 {
                        case "":
                            var v39 float32 = float32((0.0 - 1e10))
                            _ = v39
                        case "\n\t":
                            var v40 bool = (63 == 0)
                            _ = v40
                        case "hello":
                            v33 = (((v33 + v33) + "foo bar") + (v33 + v33))
                        default:
                            if (!false) {
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                                switch fi38 {
                                case -1:
                                    var v41 int8 = 1
                                    _ = v41
                                case 463:
                                    var v42 float32 = ((v3 + (v3 / 3.14159)) + 3.14159)
                                    _ = v42
                                default:
                                    fmt.Println(fi38, v3, v33)
                                    _ = len(v1)
                                    fmt.Println(v37)
                                }
                            }
                        }
                        var v43 uint32 = 4294967295
                        switch fi35 {
                        case 4294967295:
                            fmt.Println(v43, v37)
                        case 1:
                            var v44 float32 = v3
                            _ = v44
                        case -251:
                            switch v33 {
                            case "foo bar":
                                switch "" {
                                case "a":
                                    if false {
                                        if (225 <= -1) {
                                            v37 = -2147483648
                                            v37 = 42
                                            v33 = v33
                                        }
                                        switch (v37 * (fi35 % ((32767) | 1))) {
                                        case 447:
                                            v43 -= (v43 * v43)
                                        case 128:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case 16:
                                            var v45 int64 = -9223372036854775808
                                            _ = v45
                                        default:
                                            switch "hello" {
                                            case "hello":
                                                var v46 int64 = 42
                                                _ = v46
                                            }
                                        }
                                        v43 -= 1
                                    } else {
                                        if false {
                                            v47 := GenS3{F0: 257, F1: 11}
                                            fmt.Println(v37, fi35)
                                            _ = v47
                                        }
                                        v48 := GenS4{F0: v43, F1: (v43 ^ (v43 ^ (v43 % (((v43 % ((v43) | 1))) | 1))))}
                                        v49 := v1[""]
                                        var v50 uint16 = 504
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    }
                                }
                            case "\n\t":
                                if len(v2) > 0 {
                                    v2[0] = v43
                                }
                                fmt.Println(fi38)
                            case "hello":
                                var v51 uint16 = 796
                                _ = v51
                            default:
                                var v52 uint32 = v43
                                _ = v52
                            }
                        default:
                            v1[v33] = (v37 / (((^(fi35))) | 1))
                            delete(v1, ((("foo bar" + (v33 + "hello")) + "a") + ((((("a" + "hello") + ((v33 + (("\n\t" + "") + ((v33 + (v33 + (v33 + ("hello" + v33)))) + v33))) + "\n\t")) + v33) + v33) + "")))
                            var v53 uint8 = 12
                            if (v33 <= v33) {
                                var v54 float32 = (float32(1.5) - (-2.25 * ((v3 * float32(1.5)) / 1.0)))
                                v1[v33] = 9223372036854775807
                                _ = v54
                            }
                            _ = v53
                        }
                        if false {
                            v3 += v3
                            var v55 bool = true
                            var v56 float64 = 1.5
                            _ = v55
                            _ = v56
                        } else {
                            var v57 uint64 = 18446744073709551615
                            var v58 uint8 = genFn0(0, 1)
                            _ = v57
                            _ = v58
                        }
                        fmt.Println(v43)
                        switch fi38 {
                        case 441:
                            v59 := GenS1{F0: -41, F1: v37, F2: (v3 <= v3)}
                            _ = v59
                        case 128:
                            switch v33 {
                            case "foo bar":
                                v60 := GenS6{F0: 0, F1: 419}
                                _ = v60
                            }
                        case 2147483647:
                            var v61 float32 = v3
                            _ = v61
                        default:
                            var v62 uint = 941
                            _ = v62
                        }
                        _ = fi38
                        _ = v43
                    }
                    _ = v36
                    _ = v37
                }
                v33 = v33
                _ = fi35
            }
            v3 -= float32(0.0)
            _ = v34
        }
        _ = v33
    } else {
        v3 = v3
        v3 -= (float32((0.0 * (-2.25 - 3.14159))) / (-2.25 - (v3 * (1.5 + v3))))
        v63 := GenS7{F0: false, F1: 4294967295, F2: 42}
        _ = v63
    }
    v64 := GenS7{F0: (-559 < 2147483647), F1: 2147483648, F2: 123}
    var v65 uint8 = 255
    _ = len(v1)
    fmt.Println(v65)
    v3 = v3
    for fi66 := 0; fi66 < 10; fi66++ {
        var v67 float32 = float32(((3.14159 * 1e10) / 1e10))
        for fi68 := 0; fi68 < 12; fi68++ {
            var v69 float32 = 1e10
            _ = fi68
            _ = v69
        }
        _ = fi66
        _ = v67
    }
    if (42 != 880) {
        v70 := GenS3{F0: 42, F1: v65}
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = v70
    }
    switch 128 {
    case -9223372036854775808:
        v71 := GenS8{F0: 763, F1: -290}
        _ = v71
    case 2147483647:
        v72 := GenS6{F0: 304, F1: 685}
        _ = v72
    case -599:
        var v73 uint64 = 4294967295
        _ = v73
    }
    switch "" {
    case "\n\t":
        for fi74 := 0; fi74 < 12; fi74++ {
            for fi75 := 0; fi75 < 3; fi75++ {
                var v76 float32 = v3
                var v77 uint = 255
                _ = len(v1)
                _ = fi75
                _ = v76
                _ = v77
            }
            for fi78 := 0; fi78 < 16; fi78++ {
                if false {
                    var v79 uint16 = 42
                    _ = v79
                }
                v80 := GenS9{F0: -613}
                if (genFn0(255, -1) <= v65) {
                    v81 := GenS10{F0: 65535}
                    _ = v81
                }
                _ = fi78
                _ = v80
            }
            v3 = ((0.0 - -2.25) + v3)
            _ = fi74
        }
    case "":
        switch 2147483647 {
        case 133:
            var v82 uint = 0
            _ = v82
        }
    }
    v65 = (v65 - (v65 << ((787) & 63)))
    var v83 uint64 = 42
    var v84 uint16 = 42
    v85 := GenS11{F0: 32768}
    v65 = (^(v65))
    v83 += (^(v83))
    for fi86 := 0; fi86 < 17; fi86++ {
        var v87 int = 0
        delete(v1, ("\n\t" + ("\n\t" + ("" + "foo bar"))))
        _ = fi86
        _ = v87
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v65:", v65)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v64
    _ = v65
    _ = v83
    _ = v84
    _ = v85
}
