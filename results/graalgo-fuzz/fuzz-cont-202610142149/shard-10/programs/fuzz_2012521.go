package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32; F1 int8; F2 uint8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint16 }

func main() {
    var v1 int8 = 1
    var v2 uint16 = 42
    v3 := []int{42, -163, -2147483648}
    var v4 float64 = ((1.5 + 1.5) / 1.0)
    var v5 uint = 4294967295
    v6 := []int{-175, 0}
    for fi7 := 0; fi7 < 18; fi7++ {
        if len(v6) > 0 {
            v6[0] = fi7
        }
        switch fi7 {
        case -9223372036854775808:
            if len(v3) > 0 {
                v3[0] = fi7
            }
        default:
            v8 := GenS1{F0: 1}
            _ = v8
        }
        var v9 float32 = -2.25
        _ = fi7
        _ = v9
    }
    v1 -= -1
    v10 := GenS2{F0: (v2 / (((v2 << ((369) & 63))) | 1)), F1: -1}
    fmt.Println(v5)
    v11 := GenS2{F0: 255, F1: -32768}
    if ((((v2 / ((948) | 1)) > 810) || true) || (1 == 327)) {
        v4 = 3.14159
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(v5)
            var v13 uint32 = 718
            _ = fi12
            _ = v13
        }
    }
    if (128 != 864) {
        if (0 >= 2147483647) {
            switch -407 {
            case 2147483647:
                v6 = append(v6, 32767)
            case 127:
                if false {
                    var v14 string = ""
                    v14 = ("" + v14)
                    _ = v14
                }
            case 128:
                var v15 uint64 = 1
                _ = v15
            default:
                var v16 int32 = -373
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            v2 -= 65535
            _ = fi17
        }
        if len(v3) > 0 {
            v3[0] = 128
        }
    }
    switch "a" {
    case "":
        fmt.Println(v2, v1)
    case "hello":
        for fi18 := 0; fi18 < 5; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                if (!((fi19 <= fi18) && (fi18 <= (^(fi19))))) {
                    v3 = append(v3, (fi19 + 32768))
                    v1 += v1
                    switch fi19 {
                    case 318:
                        if (false || true) {
                            v20 := GenS2{F0: (v2 & (v2 % ((v2) | 1))), F1: -589}
                            switch -247 {
                            case 127:
                                v21 := GenS3{F0: v5, F1: "hello"}
                                _ = v21
                            case 32768:
                                if (true && ((v5 - (v5 << ((v5) & 63))) > 18446744073709551615)) {
                                    if false {
                                        if len(v3) > 0 {
                                            v3[0] = (fi19 >> (((v5 + v5)) & 63))
                                        }
                                    } else {
                                        if (!true) {
                                            fmt.Println(fi19, fi18, v4)
                                        } else {
                                            var v22 uint32 = 2
                                            v23 := GenS1{F0: 400}
                                            switch "hello" {
                                            case "a":
                                                v3 = append(v3, fi19)
                                                _ = len(v6)
                                                var v24 uint16 = v2
                                                _ = v24
                                            case "":
                                                switch fi18 {
                                                case 42:
                                                    v25 := GenS4{F0: 42}
                                                    _ = v25
                                                case 127:
                                                    v26 := GenS1{F0: -337}
                                                    _ = v26
                                                }
                                            case "hello":
                                                fmt.Println(v5)
                                                if len(v6) > 0 {
                                                    v6[0] = fi18
                                                }
                                                if (!(false || (false && (true && ((!(520 == 32767)) && (127 >= fi19)))))) {
                                                    switch fi19 {
                                                    case 128:
                                                        v27 := GenS3{F0: v5, F1: "a"}
                                                        _ = v27
                                                    case -449:
                                                        v5 = v5
                                                        fmt.Println(v5, v22, fi18)
                                                        var v28 uint64 = 1
                                                        _ = v28
                                                    case 9223372036854775807:
                                                        v2 = 1
                                                    default:
                                                        var v29 int64 = -1
                                                        if true {
                                                            var v30 int64 = (-(v29))
                                                            var v31 float64 = (((1.5 * (float64((1.5 + float32(v4))) + v4)) * 3.14159) * v4)
                                                            fmt.Println(fi18)
                                                            _ = v30
                                                            _ = v31
                                                        } else {
                                                            var v32 int = fi19
                                                            _ = v32
                                                        }
                                                        _ = v29
                                                    }
                                                    v2 -= v2
                                                } else {
                                                    var v33 float64 = 0.1
                                                    var v34 int64 = 9223372036854775807
                                                    switch 32768 {
                                                    case 42:
                                                        var v35 uint64 = 819
                                                        _ = v35
                                                    }
                                                    v2 = v2
                                                    _ = v33
                                                    _ = v34
                                                }
                                            }
                                            switch (fi19 * 0) {
                                            case 32767:
                                                v5 += v5
                                            case 32768:
                                                v36 := GenS4{F0: 0}
                                                _ = v36
                                            }
                                            v1 += int8(v22)
                                            v37 := GenS5{F0: (^(v5))}
                                            v22 += v22
                                            _ = v22
                                            _ = v23
                                            _ = v37
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                            default:
                                if len(v6) > 0 {
                                    v6[0] = (fi18 * fi18)
                                }
                                var v38 float64 = v4
                                switch "" {
                                case "a":
                                    var v39 uint64 = 18446744073709551615
                                    _ = v39
                                case "foo bar":
                                    var v40 uint8 = 255
                                    _ = v40
                                case "":
                                    var v41 int64 = 874
                                    _ = v41
                                default:
                                    v2 -= v2
                                    v38 = -2.25
                                    v42 := GenS6{F0: fi18}
                                    var v43 bool = true
                                    if v43 {
                                        v44 := GenS1{F0: 0}
                                        var v45 uint = v5
                                        var v46 uint8 = 181
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    } else {
                                        if (!((v5 ^ v5) <= (v5 * v5))) {
                                            switch ("foo bar" + "\n\t") {
                                            case "hello":
                                                if len(v3) > 0 {
                                                    v3[0] = (fi18 & (fi18 >> ((456) & 63)))
                                                }
                                            case "foo bar":
                                                if true {
                                                    v47 := GenS3{F0: v5, F1: (("foo bar" + ("\n\t" + "\n\t")) + "foo bar")}
                                                    fmt.Println(v43, v4)
                                                    v5 += 42
                                                    _ = v47
                                                } else {
                                                    v48 := GenS7{F0: 177}
                                                    _ = v48
                                                }
                                            }
                                            if (v43 || (!true)) {
                                                v6 = append(v6, (fi19 | (fi19 + fi18)))
                                            } else {
                                                var v49 int64 = 63
                                                if true {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    v50 := GenS4{F0: 14}
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    v6 = append(v6, fi18)
                                                    _ = v50
                                                }
                                                v5 += (^(v5))
                                                _ = v49
                                            }
                                            fmt.Println(v1, v5)
                                        } else {
                                            var v51 uint64 = 287
                                            fmt.Println(v51)
                                            v4 += 0.0
                                            _ = v51
                                        }
                                        v52 := GenS8{F0: 1, F1: v1, F2: 128}
                                        v38 += ((float64((((0.0 / float32(v4)) * -2.25) - (1e10 + (0.0 / 0.1)))) - v38) / (-2.25 + v38))
                                        if (2147483647 <= -2147483648) {
                                            v5 = 1
                                        } else {
                                            switch ("a" + "") {
                                            case "a":
                                                v2 = (^(v2))
                                            }
                                        }
                                        fmt.Println(v5, fi19)
                                        if (v1 >= (v1 >> ((v5) & 63))) {
                                            var v53 uint32 = 2147483648
                                            var v54 float32 = 1e10
                                            fmt.Println(fi19, v38, v53)
                                            _ = v53
                                            _ = v54
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = 895
                                        }
                                        _ = v52
                                    }
                                    _ = v42
                                    _ = v43
                                }
                                _ = v38
                            }
                            v5 += v5
                            _ = v20
                        } else {
                            var v55 int16 = -32768
                            _ = v55
                        }
                    case 32767:
                        v2 = (v2 - v2)
                    default:
                        if false {
                            switch (fi19 | 32768) {
                            case 584:
                                if true {
                                    var v56 int64 = -1
                                    v1 = 127
                                    _ = v56
                                }
                            case -1:
                                switch (fi19 ^ 127) {
                                case 39:
                                    switch "a" {
                                    case "foo bar":
                                        var v57 int8 = -128
                                        _ = v57
                                    default:
                                        fmt.Println(v4)
                                    }
                                case 1:
                                    v3 = append(v3, (fi18 - (fi19 - (fi19 | (fi19 << (((v5 + 65535)) & 63))))))
                                }
                                if false {
                                    fmt.Println(fi19)
                                    var v58 int8 = (^(v1))
                                    v3 = append(v3, 128)
                                    _ = v58
                                }
                            case 493:
                                switch (fi19 - fi19) {
                                case -173:
                                    v59 := GenS4{F0: 1}
                                    _ = v59
                                case 128:
                                    switch -1 {
                                    case -621:
                                        var v60 string = "hello"
                                        _ = v60
                                    case 0:
                                        v1 += 127
                                    case -9223372036854775808:
                                        var v61 string = ""
                                        _ = v61
                                    }
                                case 992:
                                    var v62 string = "a"
                                    _ = v62
                                }
                            }
                            if false {
                                var v63 int = -301
                                switch (-(fi19)) {
                                case -954:
                                    v64 := GenS5{F0: 25}
                                    _ = v64
                                case 993:
                                    switch 128 {
                                    case -1:
                                        var v65 uint64 = 18446744073709551615
                                        v66 := GenS5{F0: v5}
                                        var v67 uint64 = v65
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    case -2147483648:
                                        fmt.Println(v63)
                                        if false {
                                            fmt.Println(fi19)
                                            if (!true) {
                                                fmt.Println(v1, v63, fi19)
                                                v68 := GenS5{F0: (v5 * (v5 >> (((v5 ^ 210)) & 63)))}
                                                var v69 bool = (-303 < -201)
                                                v70 := GenS4{F0: 170}
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                            } else {
                                                if false {
                                                    var v71 uint16 = v2
                                                    v72 := GenS7{F0: 2}
                                                    v4 = v4
                                                    var v73 int = fi18
                                                    _ = v71
                                                    _ = v72
                                                    _ = v73
                                                } else {
                                                    v2 = (^(v2))
                                                    if false {
                                                        v74 := GenS4{F0: 2}
                                                        _ = v74
                                                    }
                                                }
                                                var v75 string = (("foo bar" + "") + "a")
                                                _ = v75
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        switch "" {
                                        case "a":
                                            v76 := GenS5{F0: 705}
                                            _ = v76
                                        }
                                    case 127:
                                        v2 += (^(v2))
                                    default:
                                        if (-1 > 1000) {
                                            switch (v63 << ((255) & 63)) {
                                            case 1:
                                                var v77 uint32 = 16
                                                _ = v77
                                            case 535:
                                                var v78 bool = (true && true)
                                                _ = v78
                                            case 940:
                                                v5 = (^(v5))
                                            default:
                                                v4 += float64(3.14159)
                                            }
                                            fmt.Println(v5, v5, v2)
                                            if (!((v2 - v2) != v2)) {
                                                v79 := GenS9{F0: true}
                                                _ = v79
                                            } else {
                                                v1 += (v1 & (^(v1)))
                                                var v80 uint = (^(v5))
                                                fmt.Println(v4, v5)
                                                v5 += 369
                                                var v81 float64 = (v4 * v4)
                                                fmt.Println(v1)
                                                _ = v80
                                                _ = v81
                                            }
                                            if (3.14159 > 1e10) {
                                                var v82 float32 = float32(v4)
                                                _ = v82
                                            }
                                        }
                                    }
                                default:
                                    switch ("a" + ("a" + ("" + ""))) {
                                    case "foo bar":
                                        v83 := GenS10{F0: (v2 << ((4294967295) & 63))}
                                        _ = v83
                                    case "":
                                        if false {
                                            v84 := GenS9{F0: false}
                                            var v85 float64 = ((((v4 - float64(1.5)) - float64(3.14159)) * v4) * (float64(1.5) + -2.25))
                                            if true {
                                                v85 = v85
                                            }
                                            v63 = 2147483647
                                            var v86 string = (("foo bar" + "\n\t") + "foo bar")
                                            v85 = (v85 * (float64(3.14159) + v85))
                                            fmt.Println(v1, v86)
                                            _ = v84
                                            _ = v85
                                            _ = v86
                                        } else {
                                            var v87 int64 = 566
                                            _ = v87
                                        }
                                    default:
                                        var v88 uint8 = 37
                                        _ = v88
                                    }
                                    var v89 uint8 = 0
                                    v5 += (v5 << ((v5) & 63))
                                    _ = v89
                                }
                                v5 += 255
                                _ = v63
                            }
                        }
                    }
                    _ = len(v3)
                } else {
                    if false {
                        fmt.Println(v5)
                        switch fi19 {
                        case 35:
                            var v90 string = "a"
                            _ = v90
                        }
                        v91 := GenS2{F0: v2, F1: 32767}
                        _ = v91
                    }
                    v92 := GenS7{F0: 0}
                    v1 -= -1
                    _ = v92
                }
                var v93 int16 = -581
                v4 -= v4
                if (181 <= 1) {
                    if false {
                        v93 -= (v93 * 42)
                        v94 := GenS9{F0: ("" <= ("\n\t" + ("foo bar" + "a")))}
                        switch (fi18 >> (((^(v5))) & 63)) {
                        case -484:
                            fmt.Println(v1, fi19, v2)
                        case 32768:
                            v95 := GenS1{F0: -140}
                            _ = v95
                        }
                        _ = v94
                    } else {
                        v96 := GenS9{F0: false}
                        _ = v96
                    }
                } else {
                    v1 += (v1 >> (((v5 * (v5 & v5))) & 63))
                    fmt.Println(fi18)
                }
                var v97 int = (fi18 / ((fi18) | 1))
                _ = fi19
                _ = v93
                _ = v97
            }
            v98 := GenS3{F0: (v5 / (((^(v5))) | 1)), F1: (("hello" + (("hello" + "foo bar") + "hello")) + ("foo bar" + ("\n\t" + "")))}
            _ = fi18
            _ = v98
        }
    case "foo bar":
        switch 42 {
        case -230:
            v1 -= v1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
