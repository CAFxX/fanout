package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 uint; F2 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8; F1 uint }
type GenS9 struct { F0 int8; F1 uint16 }
type GenS10 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint16, p1 int32, p2 uint32) int64 {
    p1 = (p1 - (p1 % ((p1) | 1)))
    var v6 uint = 430
    var v7 float64 = 1.5
    if (42 <= -814) {
        fmt.Println(p0, p2, p1)
        var v8 uint = v6
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 int8) string {
    if true {
        v9 := []int{-932, 2147483647, 4294967295}
        _ = v9
    }
    fmt.Println(p0)
    switch "" {
    case "hello":
        var v10 uint = 0
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        p0 -= p0
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := []int{127, (fi11 >> ((801) & 63))}
            if (int64(fi11) > 683) {
                fmt.Println(p0, fi11, fi12)
                p0 -= 63
                var v14 float64 = 3.14159
                v14 = v14
                fmt.Println(fi12, p0)
                var v15 uint16 = 255
                v16 := GenS1{F0: p0, F1: 1, F2: ((((0.0 * (v14 + ((v14 + (v14 - (0.0 * -2.25))) * v14))) / (v14 - (v14 - ((v14 - 0.0) - float64(((0.0 / (-2.25 * float32((v14 / 3.14159)))) - float32(0.0))))))) * v14) > v14)}
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    return ("\n\t" + (("\n\t" + (("" + (("hello" + "foo bar") + "")) + (("foo bar" + "\n\t") + ""))) + ""))
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    var v5 int64 = 42
    if false {
        v3[2147483647] = 4294967295
        var v17 int = -9223372036854775808
        _ = v17
    } else {
        v2 = false
    }
    fmt.Println(v2)
    v18 := v4[((("" + "foo bar") + "") + "hello")]
    switch ("a" + "") {
    case "hello":
        v19 := GenS2{F0: -1, F1: 42}
        _ = v19
    case "foo bar":
        if true {
            v20 := GenS3{F0: 100}
            v18 = false
            v21 := GenS2{F0: 681, F1: 32767}
            _ = v20
            _ = v21
        }
    }
    var v22 uint8 = 73
    v5 = 9223372036854775807
    v2 = ((1 < -128) || (v22 == 255))
    var v23 string = ("foo bar" + "a")
    switch -448 {
    case -328:
        var v24 uint = 65535
        _ = v24
    default:
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 float64 = 0.0
            if (true && (612 >= 477)) {
                v22 = 94
                for fi27 := 0; fi27 < 11; fi27++ {
                    v5 = (v5 % ((42) | 1))
                    delete(v4, v23)
                    v4[v23] = v2
                    _ = fi27
                }
                switch v23 {
                case "":
                    if (v5 <= (^(v5))) {
                        var v28 int = fi25
                        for fi29 := 0; fi29 < 6; fi29++ {
                            v30 := GenS2{F0: -639, F1: (v28 / (((fi25 ^ fi29)) | 1))}
                            _ = fi29
                            _ = v30
                        }
                        _ = v28
                    } else {
                        v31 := []int{fi25, fi25, 42, (fi25 | (fi25 / ((fi25) | 1)))}
                        for fi32 := 0; fi32 < 6; fi32++ {
                            v23 = v23
                            _ = fi32
                        }
                        if (((v23 + ((("hello" + ("\n\t" + "a")) + (v23 + "")) + v23)) + v23) > v23) {
                            for fi33 := 0; fi33 < 4; fi33++ {
                                v5 = v5
                                v26 -= v26
                                _ = len(v31)
                                _ = fi33
                            }
                        }
                        _ = v31
                    }
                case "a":
                    _ = len(v1)
                default:
                    if ((42 >= 24) || (100 == 100)) {
                        var v34 uint32 = 32
                        if (!true) {
                            if ((fi25 << ((799) & 63)) <= (fi25 << ((671) & 63))) {
                                v26 = v26
                            }
                            if v18 {
                                var v35 int = 127
                                if (false && false) {
                                    switch fi25 {
                                    case 128:
                                        var v36 float32 = 3.14159
                                        _ = v36
                                    case 9223372036854775807:
                                        fmt.Println(v2)
                                    case -925:
                                        var v37 uint64 = 1
                                        _ = v37
                                    default:
                                        v22 = v22
                                    }
                                    v38 := GenS4{F0: 18446744073709551615, F1: 1, F2: (v22 & (v22 - 225))}
                                    for fi39 := 0; fi39 < 11; fi39++ {
                                        v18 = true
                                        v23 = (v23 + (("" + "foo bar") + v23))
                                        if ((v34 >> ((129) & 63)) >= 1) {
                                            v22 = (^(v22))
                                            fmt.Println(fi39, v2, v23)
                                        }
                                        v40 := v3[v34]
                                        if (!v2) {
                                            var v41 uint16 = 426
                                            v42 := []int{}
                                            _ = v41
                                            _ = v42
                                        } else {
                                            var v43 int8 = 100
                                            if (!(1 != 65535)) {
                                                v44 := v1[32768]
                                                v45 := []int{}
                                                _ = v44
                                                _ = v45
                                            } else {
                                                v40 += v35
                                                v46 := GenS5{F0: ("\n\t" + (((v23 + (((v23 + v23) + v23) + "")) + v23) + "hello"))}
                                                var v47 bool = true
                                                if v47 {
                                                    v48 := []int{-2147483648, -773}
                                                    v43 -= 42
                                                    _ = v48
                                                } else {
                                                    fmt.Println(v5, fi39)
                                                    if (!(v22 >= 177)) {
                                                        switch (fi39 ^ fi25) {
                                                        case 32767:
                                                            v49 := GenS5{F0: ((((("foo bar" + ("" + (v23 + (v23 + v23)))) + v23) + (v23 + (v23 + (v23 + "foo bar")))) + v23) + v23)}
                                                            _ = v49
                                                        case 694:
                                                            var v50 int16 = 1
                                                            _ = v50
                                                        case 42:
                                                            v18 = false
                                                        default:
                                                            v51 := GenS3{F0: v43}
                                                            _ = v51
                                                        }
                                                        v52 := GenS5{F0: v23}
                                                        _ = v52
                                                    } else {
                                                        var v53 uint8 = (v22 * v22)
                                                        var v54 uint = 217
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                    var v55 string = "hello"
                                                    fmt.Println(fi39, v2, v34)
                                                    v23 = (v23 + (v55 + v23))
                                                    _ = v55
                                                }
                                                var v56 uint16 = 102
                                                _ = v46
                                                _ = v47
                                                _ = v56
                                            }
                                            _ = v43
                                        }
                                        var v57 uint = 404
                                        _ = fi39
                                        _ = v40
                                        _ = v57
                                    }
                                    _ = v38
                                } else {
                                    v58 := []int{fi25, v35}
                                    _ = v58
                                }
                                _ = v35
                            }
                        }
                        fmt.Println(fi25)
                        _ = v34
                    } else {
                        var v59 uint32 = 2147483647
                        if v2 {
                            switch fi25 {
                            case 790:
                                for fi60 := 0; fi60 < 8; fi60++ {
                                    var v61 float64 = v26
                                    var v62 int8 = -128
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                }
                            case 128:
                                var v63 uint64 = 453
                                _ = v63
                            }
                            var v64 int64 = (v5 / ((42) | 1))
                            if v2 {
                                if v2 {
                                    fmt.Println(v5, v18)
                                    var v65 int16 = -289
                                    switch fi25 {
                                    case 1:
                                        v22 -= (v22 << ((42) & 63))
                                    default:
                                        var v66 uint = 1
                                        _ = v66
                                    }
                                    _ = v65
                                }
                            } else {
                                for fi67 := 0; fi67 < 18; fi67++ {
                                    v4["foo bar"] = (!false)
                                    v23 = v23
                                    v1[(^(fi25))] = -2147483648
                                    _ = fi67
                                }
                                fmt.Println(v2, v59, v64)
                            }
                            _ = v64
                        } else {
                            v68 := []int{}
                            _ = v68
                        }
                        if v18 {
                            fmt.Println(v5)
                            for fi69 := 0; fi69 < 8; fi69++ {
                                v70 := GenS2{F0: 1000, F1: fi25}
                                v71 := GenS6{F0: 779, F1: (438 >= 0)}
                                switch (("" + (v23 + (v23 + v23))) + ("foo bar" + v23)) {
                                case "hello":
                                    fmt.Println(fi25)
                                case "a":
                                    var v72 uint8 = (v22 & (^(v22)))
                                    if (!true) {
                                        v73 := GenS7{F0: 65}
                                        _ = v73
                                    } else {
                                        v22 = v22
                                        var v74 uint8 = (v22 >> ((1) & 63))
                                        v5 = v5
                                        _ = v74
                                    }
                                    _ = v72
                                case "foo bar":
                                    var v75 int32 = -2147483648
                                    _ = v75
                                default:
                                    var v76 uint16 = uint16(v59)
                                    v76 += uint16(v22)
                                    _ = len(v4)
                                    v77 := GenS8{F0: 42, F1: 4294967295}
                                    _ = v76
                                    _ = v77
                                }
                                switch (^(fi69)) {
                                case 2147483647:
                                    v23 = "foo bar"
                                case 4294967295:
                                    var v78 int16 = 1000
                                    v79 := []int{(fi25 & 32767), fi69, fi69, fi69}
                                    if v18 {
                                        v26 = float64(1e10)
                                        fmt.Println(v22, v23)
                                        switch -756 {
                                        case 9223372036854775807:
                                            v80 := v4["hello"]
                                            _ = v80
                                        case 4294967295:
                                            if len(v79) > 0 {
                                                v79[0] = 81
                                            }
                                            fmt.Println(fi69, v2)
                                            fmt.Println(v26, v22)
                                            v5 -= (v5 * -733)
                                        default:
                                            var v81 uint = 0
                                            _ = v81
                                        }
                                    } else {
                                        switch v23 {
                                        case "a":
                                            var v82 uint8 = (v22 & (v22 / ((222) | 1)))
                                            _ = v82
                                        case "hello":
                                            if false {
                                                if true {
                                                    v23 = (v23 + ("\n\t" + v23))
                                                    v18 = (255 >= 228)
                                                    v83 := GenS9{F0: -19, F1: 255}
                                                    _ = v83
                                                }
                                            } else {
                                                var v84 uint = 65535
                                                _ = v84
                                            }
                                        default:
                                            delete(v1, fi25)
                                        }
                                        v59 -= 42
                                        switch int(v59) {
                                        case 0:
                                            fmt.Println(v78, v59, v18)
                                        case -2147483648:
                                            switch v23 {
                                            case "foo bar":
                                                var v85 int32 = -1
                                                switch v23 {
                                                case "\n\t":
                                                    v5 = (v5 + v5)
                                                case "a":
                                                    if v2 {
                                                        switch v23 {
                                                        case "":
                                                            switch int(v22) {
                                                            case -40:
                                                                if false {
                                                                    var v86 uint16 = 255
                                                                    fmt.Println(v18, fi69, v78)
                                                                    fmt.Println(v59, v18)
                                                                    v23 = (v23 + "")
                                                                    _ = v86
                                                                } else {
                                                                    fmt.Println(v26, v59)
                                                                    var v87 int8 = 100
                                                                    var v88 uint8 = v22
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                            case -9223372036854775808:
                                                                var v89 bool = true
                                                                _ = v89
                                                            default:
                                                                var v90 string = (v23 + v23)
                                                                _ = v90
                                                            }
                                                        default:
                                                            if ((v5 - genFn0(939, v85, v59)) != v5) {
                                                                v91 := GenS10{F0: v85, F1: 42}
                                                                fmt.Println(v85)
                                                                var v92 int16 = 1
                                                                _ = v91
                                                                _ = v92
                                                            } else {
                                                                v22 = 1
                                                                fmt.Println(v59, fi69)
                                                            }
                                                        }
                                                    } else {
                                                        v22 = 58
                                                        var v93 float32 = (-2.25 * (-2.25 * (3.14159 + 1e10)))
                                                        _ = v93
                                                    }
                                                }
                                                if false {
                                                    switch (fi69 | (fi69 << ((65535) & 63))) {
                                                    case -2147483648:
                                                        v78 += (v78 - v78)
                                                    default:
                                                        v59 += (v59 / (((v59 >> ((359) & 63))) | 1))
                                                    }
                                                    var v94 int = (fi69 | 9223372036854775807)
                                                    fmt.Println(v18)
                                                    v59 += (^(v59))
                                                    _ = v94
                                                }
                                                _ = v85
                                            case "hello":
                                                var v95 int64 = genFn0(0, 1, 588)
                                                _ = v95
                                            }
                                        default:
                                            if ((!false) || false) {
                                                v78 -= (v78 << ((18446744073709551615) & 63))
                                            }
                                        }
                                    }
                                    v3[v59] = fi25
                                    _ = v78
                                    _ = v79
                                case 0:
                                    v96 := GenS9{F0: 127, F1: 0}
                                    _ = v96
                                }
                                _ = fi69
                                _ = v70
                                _ = v71
                            }
                            v97 := v1[611]
                            _ = v97
                        } else {
                            switch v23 {
                            case "foo bar":
                                v98 := []int{(fi25 | (fi25 * 610)), (^(fi25)), fi25}
                                _ = v98
                            }
                            var v99 int64 = v5
                            v100 := GenS6{F0: 687, F1: (((v22 >> ((264) & 63)) > v22) && ((v59 ^ (v59 + v59)) <= uint32(v22)))}
                            _ = v100
                            _ = v99
                        }
                        _ = v59
                    }
                }
            }
            _ = fi25
            _ = v26
        }
    }
    for fi101 := 0; fi101 < 12; fi101++ {
        if (!(!true)) {
            v5 += (v5 / ((v5) | 1))
        }
        _ = fi101
    }
    v18 = v18
    delete(v4, ("\n\t" + v23))
    v18 = v18
    fmt.Println(v5, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
