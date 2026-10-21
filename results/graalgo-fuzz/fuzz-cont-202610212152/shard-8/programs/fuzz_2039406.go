package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint; F1 bool; F2 int64 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 int8; F1 int; F2 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 uint16 }
type GenS11 struct { F0 int; F1 int32 }
func genFn0(p0 int16, p1 uint) int16 {
    var v7 uint8 = 0
    p0 -= 1
    v8 := GenS1{F0: 32767}
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int64 = -1
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS2{F0: (!((-2147483648 < -181) || (60 != 42)))}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -706
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 2
    var v3 uint32 = 654
    var v4 string = ("a" + (("a" + ((("foo bar" + "a") + "foo bar") + "\n\t")) + "foo bar"))
    var v5 int64 = 9223372036854775807
    v6 := make(map[string]bool)
    var v13 int = -71
    fmt.Println(v4, v4, v3)
    var v14 bool = ((!true) || false)
    v15 := GenS3{F0: ((v14 || ("a" <= "foo bar")) && v14), F1: v2}
    fmt.Println(v4, v13, v14)
    var v16 uint64 = (v2 / ((v2) | 1))
    var v17 uint = 4294967295
    switch -725 {
    case 4294967295:
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 int16 = 781
            v20 := GenS4{F0: v14, F1: v4}
            _ = fi18
            _ = v19
            _ = v20
        }
    case -718:
        for fi21 := 0; fi21 < 17; fi21++ {
            if true {
                v6[(("" + v4) + "foo bar")] = v14
                fmt.Println(fi21)
            } else {
                v22 := []int{(^(fi21)), v13, v13}
                if ((v14 && (42 >= (v17 | (^(v17))))) && (277 == 764)) {
                    v23 := GenS3{F0: v14, F1: (v2 >> ((v17) & 63))}
                    for fi24 := 0; fi24 < 17; fi24++ {
                        var v25 uint8 = 128
                        if false {
                            if true {
                                v13 = fi21
                            }
                        }
                        var v26 int = fi21
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    _ = v23
                } else {
                    v3 += v3
                    v27 := v1[(v4 + "foo bar")]
                    v17 += (v17 ^ 4294967295)
                    _ = v27
                }
                if (-2.25 >= ((3.14159 * ((float32((0.0 * -2.25)) * ((0.0 * (float32(0.1) + 3.14159)) * 1.5)) * (1e10 - -2.25))) * 0.1)) {
                    v22 = append(v22, (fi21 | 32768))
                    var v28 int16 = 42
                    for fi29 := 0; fi29 < 13; fi29++ {
                        switch fi21 {
                        case 4294967295:
                            var v30 int32 = 42
                            _ = v30
                        }
                        v31 := GenS5{F0: (^(v17)), F1: ((v3 == v3) && ((true && (-997 > -9223372036854775808)) || false)), F2: v5}
                        switch 42 {
                        case 128:
                            v3 = (v3 / (((^(v3))) | 1))
                        default:
                            v17 += (v17 ^ (v17 ^ v17))
                        }
                        var v32 uint32 = (v3 - (^(v3)))
                        _ = fi29
                        _ = v31
                        _ = v32
                    }
                    _ = v28
                } else {
                    switch (v13 | v13) {
                    case 275:
                        switch v4 {
                        case "a":
                            var v33 int32 = 1
                            _ = v33
                        }
                    case -9223372036854775808:
                        for fi34 := 0; fi34 < 5; fi34++ {
                            var v35 uint = (^(v17))
                            switch (v13 >> ((v17) & 63)) {
                            case 32768:
                                fmt.Println(v2)
                            }
                            _ = fi34
                            _ = v35
                        }
                    case 1:
                        var v36 uint8 = 25
                        _ = v36
                    default:
                        for fi37 := 0; fi37 < 9; fi37++ {
                            v16 += (^(v16))
                            _ = len(v22)
                            var v38 uint64 = v16
                            _ = fi37
                            _ = v38
                        }
                    }
                }
                _ = v22
            }
            switch "foo bar" {
            case "":
                delete(v6, (v4 + "\n\t"))
            case "hello":
                var v39 string = v4
                _ = v39
            }
            for fi40 := 0; fi40 < 4; fi40++ {
                var v41 uint8 = 255
                if v14 {
                    if ((v3 * (v3 << ((v17) & 63))) < v3) {
                        fmt.Println(v17, v41, v14)
                        switch v4 {
                        case "":
                            if (363 == 0) {
                                var v42 uint16 = 729
                                v43 := GenS6{F0: v42, F1: 4294967295}
                                if true {
                                    v17 = (v17 * uint(v16))
                                    switch "a" {
                                    case "a":
                                        v44 := GenS7{F0: -63, F1: (fi40 - fi40), F2: v41}
                                        _ = v44
                                    }
                                    v17 += 18446744073709551615
                                    switch ("" + v4) {
                                    case "hello":
                                        v45 := []int{(fi40 ^ (v13 / (((v13 << (((v17 | (v17 << ((224) & 63)))) & 63))) | 1))), 189}
                                        _ = v45
                                    case "\n\t":
                                        v2 = v16
                                    default:
                                        v46 := GenS3{F0: (-817 != 42), F1: v2}
                                        _ = v46
                                    }
                                    var v47 uint64 = (v16 / ((0) | 1))
                                    fmt.Println(v16, v5)
                                    v48 := []int{fi21, (fi21 / (((-(fi21))) | 1)), 9223372036854775807, (fi40 + v13)}
                                    if v14 {
                                        _ = len(v6)
                                    }
                                    _ = v47
                                    _ = v48
                                } else {
                                    v1["\n\t"] = true
                                    v42 = 1
                                    fmt.Println(v17, v17)
                                    if (32767 != 1000) {
                                        var v49 float32 = 3.14159
                                        _ = v49
                                    } else {
                                        switch int(v41) {
                                        case 0:
                                            var v50 float64 = (3.14159 - 3.14159)
                                            _ = v50
                                        case 32768:
                                            fmt.Println(v16, v42)
                                        case -9223372036854775808:
                                            if (v41 != uint8(v42)) {
                                                var v51 bool = true
                                                _ = v51
                                            }
                                        }
                                        v13 = fi21
                                        var v52 string = v4
                                        _ = v52
                                    }
                                }
                                v2 += 1
                                _ = v42
                                _ = v43
                            } else {
                                var v53 int32 = 42
                                if (("\n\t" <= v4) || ((fi40 + (fi21 - (fi21 ^ -299))) >= fi21)) {
                                    if (!v14) {
                                        var v54 string = (v4 + v4)
                                        v54 = "hello"
                                        v55 := GenS8{F0: 9223372036854775807}
                                        _ = v54
                                        _ = v55
                                    }
                                } else {
                                    v56 := []int{(fi21 % (((fi21 * fi40)) | 1)), v13, -1}
                                    if (v14 && v14) {
                                        if (1 <= (^(v17))) {
                                            var v57 string = ""
                                            fmt.Println(v53, v4)
                                            var v58 bool = (!(0 < (v17 & 245)))
                                            _ = v57
                                            _ = v58
                                        } else {
                                            v14 = v14
                                            var v59 uint64 = (v16 - v16)
                                            _ = v59
                                        }
                                        v2 = (v16 ^ (v2 + (^(v2))))
                                        fmt.Println(v13, fi40, v4)
                                        v41 = 128
                                        fmt.Println(v2, v5)
                                        if len(v56) >= 2 {
                                            v56 = v56[1:len(v56)]
                                        }
                                    }
                                    _ = v56
                                }
                                var v60 float32 = -2.25
                                _ = v53
                                _ = v60
                            }
                        case "foo bar":
                            v61 := []int{(v13 ^ fi40), 128, int(v17)}
                            _ = v61
                        case "\n\t":
                            v13 += -739
                        }
                    } else {
                        switch fi21 {
                        case -2147483648:
                            var v62 int32 = 42
                            _ = v62
                        case 800:
                            fmt.Println(v14, v14)
                        case 4294967295:
                            switch (fi40 - 1) {
                            case 0:
                                switch v13 {
                                case 32768:
                                    if v14 {
                                        var v63 int8 = 0
                                        var v64 float32 = 1e10
                                        if (((((-2.25 / (0.1 / 3.14159)) / 0.1) == -2.25) || (uint16(v16) <= 771)) || (45 > -992)) {
                                            v5 -= 668
                                            v65 := GenS3{F0: v14, F1: uint64(v17)}
                                            _ = v65
                                        } else {
                                            switch (fi21 - 0) {
                                            case 9223372036854775807:
                                                v63 = v63
                                            default:
                                                v66 := GenS9{F0: v16}
                                                _ = v66
                                            }
                                            v67 := GenS10{F0: int16(v3), F1: 65535}
                                            v2 = 65
                                            _ = v67
                                        }
                                        _ = v63
                                        _ = v64
                                    }
                                }
                            }
                        default:
                            fmt.Println(v14, v17)
                        }
                        v68 := GenS9{F0: (v2 >> ((1) & 63))}
                        _ = v68
                    }
                    if true {
                        switch v4 {
                        case "hello":
                            fmt.Println(v5)
                        }
                    } else {
                        if (2147483647 != 42) {
                            var v69 int32 = 0
                            v69 = v69
                            switch v4 {
                            case "a":
                                fmt.Println(v41)
                            case "foo bar":
                                v69 = -2147483648
                            case "\n\t":
                                v14 = true
                            }
                            _ = v69
                        }
                        if (!false) {
                            fmt.Println(fi40, v16)
                            var v70 uint = (v17 ^ v17)
                            _ = v70
                        }
                        fmt.Println(v16, v4, v4)
                    }
                }
                _ = fi40
                _ = v41
            }
            _ = fi21
        }
    case -9223372036854775808:
        if v14 {
            fmt.Println(v14, v14, v4)
        } else {
            v5 = (v5 * (v5 / (((v5 | (v5 ^ (v5 + (v5 - (-(v5))))))) | 1)))
            for fi71 := 0; fi71 < 18; fi71++ {
                if (v14 || (-98 >= 42)) {
                    v14 = (!v14)
                    _ = len(v1)
                }
                _ = fi71
            }
        }
    }
    var v72 uint32 = v3
    for fi73 := 0; fi73 < 16; fi73++ {
        v74 := GenS11{F0: v13, F1: -1}
        for fi75 := 0; fi75 < 13; fi75++ {
            v76 := GenS8{F0: (^(v5))}
            switch (v4 + v4) {
            case "foo bar":
                v77 := v6[("\n\t" + v4)]
                _ = v77
            case "\n\t":
                v3 = v3
            default:
                var v78 string = v4
                _ = v78
            }
            _ = fi75
            _ = v76
        }
        _ = fi73
        _ = v74
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v72:", v72)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
}
