package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint8; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 string; F2 int64 }
type GenS7 struct { F0 int; F1 uint8; F2 uint64 }
type GenS8 struct { F0 uint32; F1 uint8; F2 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 0.1
    var v3 int16 = 0
    v4 := []string{("a" + ""), "\n\t"}
    var v5 float64 = 1e10
    v6 := make(map[int]int)
    v7 := GenS1{F0: 1}
    v2 = (v2 + (v5 * v5))
    v5 = (v5 * (v5 - (v2 * v5)))
    if (false || (false || (2147483648 == 0))) {
        _ = len(v1)
        if false {
            v8 := GenS2{F0: 1, F1: 0, F2: (true || ("\n\t" < "foo bar"))}
            delete(v6, 4294967295)
            _ = v8
        }
        v5 -= (float64(((0.1 / float32(1.5)) - (1e10 / 1e10))) * (1e10 + v2))
    }
    var v9 uint16 = 0
    v3 += (v3 >> ((65535) & 63))
    var v10 uint8 = 225
    v3 = v3
    var v11 string = ("a" + ("hello" + ""))
    if true {
        delete(v6, 9223372036854775807)
        if true {
            v12 := GenS2{F0: 822, F1: v10, F2: (-128 > 100)}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 4; fi13++ {
                for fi14 := 0; fi14 < 17; fi14++ {
                    fmt.Println(fi14)
                    _ = fi14
                }
                for fi15 := 0; fi15 < 6; fi15++ {
                    fmt.Println(v10)
                    _ = fi15
                }
                fmt.Println(v10, v5, v3)
                _ = fi13
            }
        }
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 float64 = ((v5 - float64(1.5)) + ((v5 * 1.5) - 3.14159))
        if (false && true) {
            var v18 float64 = v5
            _ = v18
        } else {
            if false {
                var v19 int8 = 1
                v9 -= (v9 / (((v9 - (v9 | (v9 - (v9 >> ((190) & 63)))))) | 1))
                var v20 uint64 = 42
                _ = v19
                _ = v20
            } else {
                v21 := GenS3{F0: fi16}
                _ = v21
            }
            for fi22 := 0; fi22 < 20; fi22++ {
                if false {
                    var v23 bool = (127 != 1)
                    v24 := GenS4{F0: 365}
                    fmt.Println(v23, v5, v5)
                    _ = v23
                    _ = v24
                }
                fmt.Println(v9, v10)
                switch (v11 + v11) {
                case "\n\t":
                    var v25 uint8 = (v10 & 2)
                    _ = v25
                case "":
                    switch (fi22 >> ((488) & 63)) {
                    case 0:
                        fmt.Println(v2)
                    case 171:
                        fmt.Println(v2, v9, v11)
                    case 42:
                        var v26 int = -1
                        _ = v26
                    default:
                        if false {
                            var v27 int64 = 0
                            _ = v27
                        }
                    }
                case "hello":
                    fmt.Println(v5, v5)
                default:
                    var v28 int8 = 42
                    _ = v28
                }
                v11 = v11
                switch fi16 {
                case 42:
                    var v29 int64 = 749
                    _ = v29
                case 222:
                    if (true && (v3 < (v3 % (((^(v3))) | 1)))) {
                        fmt.Println(v9, v3)
                        v30 := GenS3{F0: (fi22 * fi16)}
                        _ = v30
                    } else {
                        var v31 float64 = float64(1.5)
                        if false {
                            if true {
                                v32 := GenS4{F0: (^(v9))}
                                if len(v4) > 0 {
                                    v4[0] = "a"
                                }
                                _ = v32
                            }
                        } else {
                            switch fi16 {
                            case 42:
                                switch (fi16 ^ 32767) {
                                case 1:
                                    v33 := v6[fi22]
                                    fmt.Println(v3)
                                    v33 = (fi16 - (fi22 << ((42) & 63)))
                                    _ = v33
                                case -1:
                                    v9 -= 251
                                case 907:
                                    v5 -= (1e10 + v5)
                                }
                            case 4294967295:
                                v10 += 255
                            case 128:
                                var v34 int32 = -1
                                _ = v34
                            }
                        }
                        _ = v31
                    }
                case 2147483647:
                    if true {
                        fmt.Println(v9, v10, v10)
                        v9 -= v9
                    } else {
                        if true {
                            var v35 int = 0
                            v36 := GenS2{F0: 2147483647, F1: v10, F2: (((((0.0 * float64(3.14159)) - (v17 / (v2 - v17))) * float64(float32(v5))) - ((-2.25 * 1.5) + 1.5)) >= v17)}
                            var v37 int = fi16
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                    }
                default:
                    var v38 float32 = 1.5
                    _ = v38
                }
                _ = fi22
            }
            var v39 int16 = v3
            _ = v39
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v10, v11, v11)
    v3 += 32767
    v9 = (v9 % (((v9 >> ((4294967295) & 63))) | 1))
    switch ("foo bar" + (v11 + "foo bar")) {
    case "foo bar":
        switch 0 {
        case 32767:
            var v40 float32 = (0.0 - ((float32(3.14159) * float32((v5 + 0.1))) - 1e10))
            _ = v40
        case -538:
            if (!true) {
                fmt.Println(v10, v10, v9)
                v10 -= v10
            } else {
                fmt.Println(v10, v2, v3)
                for fi41 := 0; fi41 < 12; fi41++ {
                    for fi42 := 0; fi42 < 19; fi42++ {
                        fmt.Println(v9)
                        fmt.Println(fi42, fi41)
                        var v43 int8 = 0
                        _ = fi42
                        _ = v43
                    }
                    _ = fi41
                }
                fmt.Println(v11)
            }
        case 884:
            if false {
                for fi44 := 0; fi44 < 17; fi44++ {
                    if len(v4) > 0 {
                        v4[0] = v11
                    }
                    for fi45 := 0; fi45 < 5; fi45++ {
                        delete(v6, fi45)
                        _ = fi45
                    }
                    v5 = float64(-2.25)
                    _ = fi44
                }
            }
        default:
            if ((2147483647 <= 0) && false) {
                v46 := GenS2{F0: int32(v9), F1: v10, F2: ((v11 + v11) <= "\n\t")}
                v47 := GenS4{F0: uint16(v10)}
                _ = v46
                _ = v47
            }
        }
    case "\n\t":
        for fi48 := 0; fi48 < 7; fi48++ {
            v5 += (v2 * v5)
            v49 := v6[fi48]
            if false {
                if (2147483647 >= 42) {
                    var v50 uint = 213
                    _ = v50
                } else {
                    v5 = v5
                    if true {
                        var v51 float64 = 1.5
                        v1[(v49 >> ((18446744073709551615) & 63))] = true
                        _ = v51
                    }
                    fmt.Println(fi48, v3)
                }
                for fi52 := 0; fi52 < 6; fi52++ {
                    if (true && true) {
                        v53 := GenS5{F0: 0}
                        _ = v53
                    }
                    _ = fi52
                }
                v9 = 255
            } else {
                var v54 uint8 = v10
                _ = v54
            }
            _ = fi48
            _ = v49
        }
    case "":
        v2 = v5
    default:
        var v55 uint16 = 255
        _ = v55
    }
    if (((v9 << ((2) & 63)) < 400) || (!(!(v10 == 128)))) {
        switch "" {
        case "foo bar":
            _ = len(v4)
        case "a":
            _ = len(v4)
        default:
            if true {
                for fi56 := 0; fi56 < 7; fi56++ {
                    switch 32767 {
                    case 9223372036854775807:
                        var v57 int8 = 0
                        _ = v57
                    case -172:
                        if (463 > (v9 << ((4294967295) & 63))) {
                            for fi58 := 0; fi58 < 12; fi58++ {
                                var v59 float64 = (float64(-2.25) - v5)
                                _ = fi58
                                _ = v59
                            }
                            v11 = (("\n\t" + "foo bar") + v11)
                            if true {
                                switch v11 {
                                case "\n\t":
                                    var v60 uint = 255
                                    _ = v60
                                default:
                                    var v61 uint32 = 65535
                                    _ = v61
                                }
                                var v62 int32 = 0
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v62
                            }
                        } else {
                            v63 := GenS5{F0: 587}
                            if true {
                                if false {
                                    _ = len(v1)
                                }
                            } else {
                                switch 127 {
                                case 32767:
                                    fmt.Println(v9, v2)
                                case -936:
                                    fmt.Println(v2, v2, v10)
                                default:
                                    v64 := GenS4{F0: uint16(v10)}
                                    _ = v64
                                }
                                var v65 uint64 = 18446744073709551615
                                _ = v65
                            }
                            var v66 uint32 = 0
                            _ = v63
                            _ = v66
                        }
                    case 2147483647:
                        switch (fi56 | 964) {
                        case -2147483648:
                            for fi67 := 0; fi67 < 1; fi67++ {
                                fmt.Println(v3)
                                fmt.Println(v5, v10)
                                _ = len(v6)
                                _ = fi67
                            }
                        case 4294967295:
                            var v68 int32 = 817
                            _ = v68
                        default:
                            if false {
                                switch fi56 {
                                case 128:
                                    for fi69 := 0; fi69 < 17; fi69++ {
                                        var v70 uint = 42
                                        if true {
                                            v3 -= 1
                                        } else {
                                            var v71 int64 = 834
                                            _ = v71
                                        }
                                        fmt.Println(v5, v5)
                                        _ = fi69
                                        _ = v70
                                    }
                                case 411:
                                    for fi72 := 0; fi72 < 16; fi72++ {
                                        switch (fi56 | (fi56 >> ((657) & 63))) {
                                        case 127:
                                            if ((v3 * 32767) > (v3 * 1000)) {
                                                v10 = (v10 ^ v10)
                                                var v73 string = ("\n\t" + "a")
                                                _ = v73
                                            }
                                        }
                                        var v74 float64 = v5
                                        switch fi72 {
                                        case 45:
                                            v9 += v9
                                        case -580:
                                            if (32767 != v9) {
                                                v5 -= 1e10
                                                v9 += v9
                                                var v75 int64 = -345
                                                fmt.Println(v2, v75, v2)
                                                if (((v75 - v75) <= v75) && true) {
                                                    delete(v6, fi72)
                                                    v11 = "\n\t"
                                                    v76 := GenS6{F0: int(v75), F1: "", F2: v75}
                                                    _ = v76
                                                }
                                                _ = v75
                                            }
                                        default:
                                            var v77 uint16 = v9
                                            _ = v77
                                        }
                                        fmt.Println(v74, v3)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v11 = v11
                                        switch v11 {
                                        case "hello":
                                            v78 := GenS7{F0: (fi72 / (((fi72 + (fi72 - (fi72 ^ (fi56 << ((uint(v10)) & 63)))))) | 1)), F1: 127, F2: 0}
                                            _ = v78
                                        }
                                        v1[1] = false
                                        _ = fi72
                                        _ = v74
                                    }
                                case -2147483648:
                                    var v79 uint16 = 0
                                    _ = v79
                                }
                                v80 := GenS8{F0: 4294967295, F1: (v10 * v10), F2: 121}
                                _ = v80
                            }
                        }
                    }
                    _ = fi56
                }
                var v81 uint = 712
                _ = v81
            } else {
                v2 += (v2 / 3.14159)
                v9 -= (v9 % ((v9) | 1))
            }
        }
    } else {
        v4 = append(v4, v11)
        v10 = (v10 + (v10 - (v10 << ((2) & 63))))
        var v82 uint16 = (^(v9))
        _ = v82
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
