package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint; F2 int64 }
type GenS5 struct { F0 string; F1 int8; F2 int64 }
type GenS6 struct { F0 int16; F1 uint32; F2 int }
type GenS7 struct { F0 uint32; F1 int }
func genFn0(p0 bool) uint32 {
    if (float64(0.1) != ((0.1 + (0.0 / (((float64(0.0) * 0.0) / 1.0) - 1e10))) * 3.14159)) {
        for fi4 := 0; fi4 < 11; fi4++ {
            var v5 uint32 = 702
            var v6 uint64 = 42
            v7 := GenS1{F0: 127, F1: 305}
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS2{F0: 42}
                fmt.Println(p0, fi8, p0)
                v11 := GenS2{F0: 2147483647}
                var v12 uint16 = 65535
                v12 -= 135
                fmt.Println(fi9)
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            v13 := GenS1{F0: 127, F1: 159}
            p0 = ((-29 != 82) || (p0 && (3.14159 == (0.1 - 0.1))))
            _ = fi8
            _ = v13
        }
        v14 := GenS3{F0: 42}
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := []int{(fi15 << ((65535) & 63)), (fi15 >> ((65535) & 63)), 32768, fi15}
            p0 = (p0 && p0)
            p0 = (42 > 2)
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    switch "hello" {
    case "":
        var v17 float64 = 3.14159
        _ = v17
    case "a":
        var v18 uint64 = 18446744073709551615
        _ = v18
    default:
        if p0 {
            if (692 <= 0) {
                v19 := []int{}
                v20 := GenS2{F0: 4294967295}
                for fi21 := 0; fi21 < 16; fi21++ {
                    fmt.Println(fi21, fi21)
                    v19 = append(v19, 9223372036854775807)
                    _ = fi21
                }
                _ = v19
                _ = v20
            } else {
                v22 := GenS4{F0: 128, F1: 1, F2: -1}
                if (42 == 42) {
                    fmt.Println(p0)
                    for fi23 := 0; fi23 < 13; fi23++ {
                        for fi24 := 0; fi24 < 18; fi24++ {
                            v25 := GenS4{F0: fi23, F1: 65535, F2: 1}
                            var v26 float32 = 1.5
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        p0 = ((!((354 != 39) && ((p0 || (((!p0) && p0) && false)) && (false || p0)))) && (!(-2147483648 <= 533)))
                        _ = fi23
                    }
                    v27 := GenS5{F0: ("a" + "a"), F1: 127, F2: -178}
                    _ = v27
                }
                var v28 int32 = 25
                _ = v22
                _ = v28
            }
            var v29 uint = 65535
            if p0 {
                v29 = v29
                v30 := []int{-1, -383, 943}
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 uint16 = 42
                    var v33 uint = (v29 % ((1) | 1))
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = v30
            } else {
                for fi34 := 0; fi34 < 10; fi34++ {
                    switch "foo bar" {
                    case "":
                        var v35 int32 = -370
                        _ = v35
                    case "foo bar":
                        var v36 int32 = 1
                        _ = v36
                    }
                    var v37 uint64 = 571
                    v29 = 18446744073709551615
                    _ = fi34
                    _ = v37
                }
                if (!p0) {
                    switch 128 {
                    case 4294967295:
                        for fi38 := 0; fi38 < 11; fi38++ {
                            if p0 {
                                v39 := []int{fi38, fi38}
                                _ = v39
                            } else {
                                v29 = v29
                                v29 = v29
                            }
                            _ = fi38
                        }
                    case 193:
                        var v40 int8 = 0
                        _ = v40
                    case -748:
                        if p0 {
                            v29 += (v29 & v29)
                            fmt.Println(v29)
                        }
                    default:
                        var v41 uint16 = 32767
                        _ = v41
                    }
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        v42 := GenS3{F0: 225}
                        _ = v42
                    default:
                        var v43 float64 = float64((((1e10 / 3.14159) - ((1e10 - 1.5) - (1e10 * 0.0))) / -2.25))
                        _ = v43
                    }
                    var v44 float32 = 1e10
                    _ = v44
                } else {
                    v29 -= v29
                }
                for fi45 := 0; fi45 < 14; fi45++ {
                    var v46 int = fi45
                    fmt.Println(p0)
                    _ = fi45
                    _ = v46
                }
            }
            _ = v29
        } else {
            v47 := []int{0, 32768}
            _ = v47
        }
    }
    _ = p0
    return 65535
}
func genFn1() string {
    var v48 bool = (0 > 65535)
    switch -221 {
    case -487:
        v48 = true
    case 319:
        var v49 string = "foo bar"
        _ = v49
    default:
        var v50 int8 = 0
        _ = v50
    }
    if false {
        for fi51 := 0; fi51 < 19; fi51++ {
            v52 := []int{-392, fi51, 128}
            _ = fi51
            _ = v52
        }
        var v53 int = -335
        if v48 {
            v48 = v48
            v54 := GenS3{F0: -128}
            v48 = true
            _ = v54
        } else {
            var v55 int32 = 1
            _ = v55
        }
        _ = v53
    } else {
        v48 = true
        var v56 uint64 = 42
        switch -9223372036854775808 {
        case 1:
            fmt.Println(v48)
        case 127:
            v48 = v48
        default:
            var v57 uint32 = 347
            _ = v57
        }
        _ = v56
    }
    for fi58 := 0; fi58 < 18; fi58++ {
        v59 := []int{fi58, fi58}
        _ = fi58
        _ = v59
    }
    _ = v48
    return ""
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]string)
    var v3 int8 = 127
    v60 := GenS6{F0: 1, F1: genFn0(((("a" == "foo bar") || (v3 <= v3)) || (4294967295 >= 4294967295))), F2: 9223372036854775807}
    v3 = (-(v3))
    for fi61 := 0; fi61 < 7; fi61++ {
        var v62 bool = (!false)
        _ = fi61
        _ = v62
    }
    v63 := []int{-291}
    switch -9223372036854775808 {
    case -1:
        v3 += v3
    case -2147483648:
        v64 := v2[0]
        _ = v64
    case -258:
        var v65 int64 = -596
        _ = v65
    }
    if true {
        if false {
            var v66 float32 = (1.5 * -2.25)
            var v67 string = "\n\t"
            v67 = "a"
            _ = v66
            _ = v67
        }
        if (!false) {
            for fi68 := 0; fi68 < 17; fi68++ {
                var v69 int32 = -2147483648
                _ = fi68
                _ = v69
            }
            switch "" {
            case "foo bar":
                var v70 float64 = -2.25
                _ = v70
            }
        }
    } else {
        var v71 uint8 = 223
        v71 += 100
        for fi72 := 0; fi72 < 5; fi72++ {
            switch ("a" + "foo bar") {
            case "\n\t":
                for fi73 := 0; fi73 < 1; fi73++ {
                    if (((float64(1e10) * 1e10) + v1) > (0.0 / v1)) {
                        v71 = 37
                    } else {
                        delete(v2, 2147483647)
                        fmt.Println(fi73, fi72)
                        var v74 int64 = -9223372036854775808
                        var v75 uint = 255
                        _ = v74
                        _ = v75
                    }
                    var v76 int64 = 1
                    _ = fi73
                    _ = v76
                }
            case "a":
                v77 := GenS4{F0: fi72, F1: 1, F2: -332}
                _ = v77
            case "":
                v78 := GenS5{F0: "hello", F1: (-(v3)), F2: 9223372036854775807}
                _ = v78
            default:
                if false {
                    v79 := GenS3{F0: 0}
                    fmt.Println(v1, v1)
                    var v80 uint = 456
                    _ = v79
                    _ = v80
                }
            }
            if (!true) {
                var v81 string = "a"
                _ = v81
            } else {
                for fi82 := 0; fi82 < 6; fi82++ {
                    v71 = (^(v71))
                    v71 += 127
                    v83 := GenS4{F0: -1, F1: 2, F2: -9223372036854775808}
                    _ = fi82
                    _ = v83
                }
                v63 = append(v63, fi72)
                switch fi72 {
                case 1:
                    fmt.Println(v3, fi72)
                case 9223372036854775807:
                    for fi84 := 0; fi84 < 10; fi84++ {
                        fmt.Println(fi84, v3)
                        fmt.Println(fi84)
                        _ = fi84
                    }
                case 568:
                    if (2147483648 < 4294967295) {
                        for fi85 := 0; fi85 < 15; fi85++ {
                            v86 := GenS3{F0: -9223372036854775808}
                            v71 += (v71 | 99)
                            fmt.Println(fi72)
                            _ = fi85
                            _ = v86
                        }
                    }
                }
            }
            _ = fi72
        }
        _ = v71
    }
    if ((0.0 >= 0.0) || false) {
        var v87 uint64 = 829
        _ = v87
    } else {
        for fi88 := 0; fi88 < 13; fi88++ {
            if false {
                v1 = (float64(1.5) + float64((1e10 + (3.14159 + 1e10))))
            }
            _ = fi88
        }
        if len(v63) >= 2 {
            v63 = v63[1:len(v63)]
        }
    }
    delete(v2, 339)
    fmt.Println(v3, v3)
    v1 += float64(0.1)
    if ((((v1 / 0.1) + ((float64(1.5) + 0.1) - (v1 / (v1 * float64(0.0))))) * (float64(3.14159) + (-2.25 / 0.1))) <= 1e10) {
        if false {
            var v89 int8 = v3
            v3 -= v3
            _ = v89
        }
        v90 := GenS7{F0: 0, F1: -1}
        _ = v90
    }
    fmt.Println(v3)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v63:", len(v63), v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v60
    _ = v63
}
