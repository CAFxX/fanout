package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint8; F2 string }
type GenS5 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint16; F1 uint32 }
func genFn0(p0 bool, p1 uint64) int8 {
    v6 := []int{-756, 1}
    for fi7 := 0; fi7 < 19; fi7++ {
        p1 = (^(p1))
        _ = fi7
    }
    if (42 <= 42) {
        if p0 {
            fmt.Println(p1, p0)
        } else {
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 uint8 = 227
                p1 -= (^(p1))
                fmt.Println(p1, p1, fi8)
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 bool = p0
                v11 = ((p1 * p1) != p1)
                switch "a" {
                case "":
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v13 := GenS1{F0: ""}
                        var v14 int32 = 2147483647
                        _ = len(v6)
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case "hello":
                    p0 = (908 != 44)
                }
                _ = fi10
                _ = v11
            }
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(p1)
    }
    v15 := GenS1{F0: "\n\t"}
    _ = p0
    _ = p1
    _ = v15
    _ = v6
    return 1
}
func genFn1(p0 uint64) uint {
    var v16 string = "\n\t"
    if (false && ((-9223372036854775808 >= -2147483648) || false)) {
        v16 = v16
        var v17 float32 = 1.5
        _ = v17
    } else {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 int32 = -2147483648
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := []int{(^(fi20))}
                p0 = (p0 & (p0 >> ((471) & 63)))
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
    }
    v22 := GenS2{F0: 935, F1: v16, F2: 128}
    v23 := GenS3{F0: (true || (!true)), F1: 2147483648}
    _ = p0
    _ = v16
    _ = v22
    _ = v23
    return 255
}

func main() {
    var v1 uint = 17
    var v2 int64 = -9223372036854775808
    v3 := []string{"\n\t", ("" + "a"), ""}
    var v4 bool = (127 <= 42)
    v5 := make(map[uint32]int)
    v5[42] = 9223372036854775807
    var v24 uint16 = 32767
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 uint16 = v24
        _ = fi25
        _ = v26
    }
    if false {
        if v4 {
            _ = len(v3)
            if (1 >= -32768) {
                var v27 string = "hello"
                for fi28 := 0; fi28 < 17; fi28++ {
                    var v29 uint32 = 853
                    var v30 int64 = 9223372036854775807
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v1 = 781
                _ = v27
            } else {
                v24 = (v24 - 1)
                if v4 {
                    for fi31 := 0; fi31 < 1; fi31++ {
                        v32 := GenS3{F0: v4, F1: 600}
                        _ = fi31
                        _ = v32
                    }
                    if true {
                        switch 998 {
                        case 32767:
                            v24 -= (v24 | (v24 % (((v24 >> ((genFn1(510)) & 63))) | 1)))
                        }
                    }
                }
            }
        }
        v1 -= 255
    }
    if ((128 < 132) || true) {
        var v33 int64 = (v2 / ((v2) | 1))
        if (0 < 1) {
            switch 9223372036854775807 {
            case -9223372036854775808:
                if (!(!(!v4))) {
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    v1 -= (v1 << ((v1) & 63))
                    v4 = v4
                }
            }
            v34 := v5[892]
            _ = v34
        } else {
            v5[2147483647] = 32768
            var v35 uint64 = 162
            _ = v35
        }
        _ = v33
    }
    fmt.Println(v2, v24)
    v24 = v24
    v36 := GenS2{F0: v24, F1: "foo bar", F2: 127}
    var v37 string = ""
    if v4 {
        v38 := GenS1{F0: v37}
        _ = len(v3)
        _ = v38
    } else {
        var v39 int32 = 0
        _ = v39
    }
    if true {
        v4 = (!true)
        var v40 int16 = 1
        _ = v40
    } else {
        if true {
            if v4 {
                switch 1 {
                case 32768:
                    v37 = ("hello" + "a")
                case -1:
                    v41 := GenS4{F0: -45, F1: 0, F2: "\n\t"}
                    _ = v41
                default:
                    v42 := GenS5{F0: v24, F1: -52, F2: 2}
                    _ = v42
                }
                fmt.Println(v1, v2)
                for fi43 := 0; fi43 < 10; fi43++ {
                    switch "\n\t" {
                    case "foo bar":
                        var v44 uint8 = 2
                        _ = v44
                    case "hello":
                        for fi45 := 0; fi45 < 6; fi45++ {
                            v46 := GenS3{F0: false, F1: 42}
                            if len(v3) > 0 {
                                v3[0] = v37
                            }
                            if v4 {
                                v47 := GenS4{F0: 1, F1: 1, F2: v37}
                                if v4 {
                                    if true {
                                        var v48 uint8 = 255
                                        _ = len(v5)
                                        v49 := GenS1{F0: (v37 + "a")}
                                        _ = v48
                                        _ = v49
                                    }
                                    var v50 bool = ("a" > "\n\t")
                                    if v50 {
                                        var v51 bool = false
                                        v2 = 743
                                        v50 = v50
                                        _ = v51
                                    } else {
                                        if (((v1 + uint(v24)) != (v1 << ((v1) & 63))) && (0 == 2)) {
                                            fmt.Println(v4, fi45)
                                        }
                                        v37 = "foo bar"
                                        v52 := GenS5{F0: (v24 - (v24 | v24)), F1: 42, F2: 687}
                                        _ = v52
                                    }
                                    _ = v50
                                } else {
                                    switch v37 {
                                    case "\n\t":
                                        v1 += genFn1(509)
                                    }
                                    fmt.Println(v24, fi45, fi43)
                                }
                                if ((v1 >> (((v1 % (((v1 + 18446744073709551615)) | 1))) & 63)) != (v1 >> (((v1 << ((v1) & 63))) & 63))) {
                                    v53 := GenS6{F0: 0, F1: (v24 > (^(v24)))}
                                    _ = v53
                                } else {
                                    v2 += v2
                                }
                                _ = v47
                            } else {
                                fmt.Println(v1, fi43, fi45)
                                if true {
                                    switch -9223372036854775808 {
                                    case -413:
                                        v5[2147483648] = fi45
                                    case 282:
                                        switch -1 {
                                        case -2147483648:
                                            var v54 float64 = (1e10 - (-2.25 * 0.1))
                                            _ = v54
                                        default:
                                            if (float64(3.14159) != 1.5) {
                                                var v55 float64 = ((-2.25 - 1.5) - 1e10)
                                                if v4 {
                                                    v56 := GenS2{F0: v24, F1: "foo bar", F2: 255}
                                                    delete(v5, 2147483648)
                                                    v1 -= (^(v1))
                                                    _ = v56
                                                }
                                                v1 = (v1 * (v1 >> ((v1) & 63)))
                                                _ = v55
                                            }
                                        }
                                    case 32768:
                                        fmt.Println(fi45)
                                    default:
                                        v3 = append(v3, v37)
                                    }
                                    fmt.Println(fi45, v24)
                                } else {
                                    v5[42] = fi43
                                    v37 = v37
                                    var v57 int8 = 100
                                    _ = v57
                                }
                                v58 := GenS1{F0: (v37 + "foo bar")}
                                _ = v58
                            }
                            _ = fi45
                            _ = v46
                        }
                    case "":
                        var v59 int = fi43
                        _ = v59
                    }
                    _ = fi43
                }
            } else {
                for fi60 := 0; fi60 < 16; fi60++ {
                    v1 = 92
                    v37 = "\n\t"
                    v2 = v2
                    _ = fi60
                }
                if (1.5 <= 3.14159) {
                    for fi61 := 0; fi61 < 9; fi61++ {
                        fmt.Println(v2)
                        _ = fi61
                    }
                    for fi62 := 0; fi62 < 12; fi62++ {
                        switch ("a" + "") {
                        case "":
                            var v63 int64 = 1
                            _ = v63
                        default:
                            v64 := GenS4{F0: -1, F1: 251, F2: v37}
                            _ = v64
                        }
                        v24 += (v24 / (((v24 << ((v1) & 63))) | 1))
                        _ = fi62
                    }
                }
                for fi65 := 0; fi65 < 13; fi65++ {
                    var v66 float64 = 0.1
                    switch ((v37 + (((v37 + v37) + (v37 + v37)) + v37)) + "hello") {
                    case "foo bar":
                        switch (fi65 << (((v1 * v1)) & 63)) {
                        case 368:
                            fmt.Println(v37)
                        case -1:
                            v5[uint32(v24)] = (^(fi65))
                        case 9223372036854775807:
                            var v67 float32 = -2.25
                            _ = v67
                        default:
                            var v68 bool = v4
                            _ = v68
                        }
                    case "hello":
                        v69 := GenS3{F0: v4, F1: 736}
                        _ = v69
                    case "":
                        fmt.Println(v2, v1)
                    default:
                        for fi70 := 0; fi70 < 13; fi70++ {
                            v66 = ((-2.25 + v66) / 1e10)
                            switch (fi70 ^ fi70) {
                            case -919:
                                if (v4 || (v66 != ((1e10 + 1e10) + float64(float32(((float64(float32(v66)) - v66) / (float64((0.1 + ((float32(0.0) - 1e10) + 1.5))) * (v66 + (1e10 + float64(((1.5 + 1e10) * (((1.5 - -2.25) / (1e10 + 0.0)) * 3.14159)))))))))))) {
                                    var v71 int = (fi65 / ((516) | 1))
                                    switch v37 {
                                    case "hello":
                                        delete(v5, 1)
                                    case "foo bar":
                                        switch (-(fi65)) {
                                        case 32767:
                                            v72 := GenS1{F0: v37}
                                            v4 = v4
                                            _ = v72
                                        }
                                    }
                                    v73 := GenS7{F0: 0, F1: true}
                                    v74 := v5[260]
                                    var v75 int16 = 0
                                    if (!((2147483648 >= 42) || v4)) {
                                        fmt.Println(v1)
                                    }
                                    _ = v71
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                } else {
                                    if (true || (v4 && true)) {
                                        if true {
                                            switch v37 {
                                            case "\n\t":
                                                var v76 float64 = v66
                                                _ = v76
                                            case "a":
                                                switch (fi70 / ((576) | 1)) {
                                                case 2147483647:
                                                    v77 := GenS8{F0: 2, F1: 42}
                                                    if ((0.0 >= 1.5) || ((v2 >> ((v1) & 63)) >= v2)) {
                                                        v1 += 2
                                                    }
                                                    _ = v77
                                                case 706:
                                                    var v78 int32 = -1
                                                    _ = v78
                                                default:
                                                    var v79 uint16 = v24
                                                    _ = v79
                                                }
                                            default:
                                                fmt.Println(fi65, fi70)
                                            }
                                            fmt.Println(v66, v66, v4)
                                        }
                                        v80 := GenS9{F0: 2}
                                        _ = v80
                                    }
                                    v81 := GenS8{F0: 38, F1: 487}
                                    _ = v81
                                }
                            case 142:
                                var v82 int64 = (v2 - (-(v2)))
                                _ = v82
                            }
                            _ = fi70
                        }
                    }
                    _ = fi65
                    _ = v66
                }
            }
        }
        for fi83 := 0; fi83 < 9; fi83++ {
            v84 := v5[2147483648]
            _ = fi83
            _ = v84
        }
    }
    for fi85 := 0; fi85 < 20; fi85++ {
        v86 := v5[1]
        var v87 int32 = 0
        _ = fi85
        _ = v86
        _ = v87
    }
    v88 := GenS3{F0: (!(0.1 <= float64((((3.14159 - 3.14159) / (((0.1 + 3.14159) - (float32(1e10) - float32(1e10))) * float32(1.5))) + (-2.25 - float32(0.0)))))), F1: 144}
    if true {
        _ = len(v3)
        delete(v5, 1)
        if v4 {
            var v89 bool = (-128 > 0)
            _ = len(v3)
            v89 = ((-2.25 + float64(0.1)) > ((0.1 - -2.25) + 0.0))
            _ = v89
        }
    } else {
        if v4 {
            var v90 int32 = 2147483647
            _ = v90
        }
        for fi91 := 0; fi91 < 8; fi91++ {
            if (!v4) {
                var v92 uint32 = 883
                _ = v92
            } else {
                v37 = v37
                fmt.Println(v37, v24)
                if ((("\n\t" + (v37 + v37)) + (("\n\t" + (v37 + v37)) + "a")) >= v37) {
                    if false {
                        v2 -= (v2 & (-(v2)))
                        var v93 uint = v1
                        fmt.Println(v93)
                        _ = v93
                    }
                }
            }
            _ = fi91
        }
    }
    v2 = 42
    v94 := GenS10{F0: (v24 & 42), F1: 4294967295}
    if false {
        v24 = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v88
    _ = v94
}
