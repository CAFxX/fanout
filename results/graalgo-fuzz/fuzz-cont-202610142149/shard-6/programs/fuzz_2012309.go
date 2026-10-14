package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int64; F1 int32; F2 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 int }
type GenS8 struct { F0 bool }

func main() {
    var v1 int32 = 0
    var v2 uint32 = 103
    var v3 uint32 = (v2 % (((v2 - 2147483647)) | 1))
    v4 := make(map[uint32]bool)
    var v5 uint16 = 2
    var v6 uint64 = 42
    switch -162 {
    case -119:
        if false {
            for fi7 := 0; fi7 < 17; fi7++ {
                for fi8 := 0; fi8 < 8; fi8++ {
                    if ((false && true) && ((v3 ^ v3) > v3)) {
                        v9 := []int{(fi8 & fi7), (fi7 + (fi8 + fi8)), int(v5)}
                        v3 -= 2147483648
                        v10 := GenS1{F0: -571, F1: fi7, F2: 373}
                        _ = v10
                        _ = v9
                    } else {
                        v5 -= v5
                        var v11 int8 = 42
                        _ = v11
                    }
                    if ("foo bar" == "a") {
                        v4[(^(v3))] = true
                    } else {
                        var v12 int8 = 1
                        _ = v12
                    }
                    _ = fi8
                }
                if true {
                    switch fi7 {
                    case 9223372036854775807:
                        var v13 uint64 = (v6 / (((v6 % (((v6 | v6)) | 1))) | 1))
                        _ = v13
                    case 1:
                        for fi14 := 0; fi14 < 3; fi14++ {
                            v15 := []int{}
                            v1 -= -226
                            _ = fi14
                            _ = v15
                        }
                    case -2147483648:
                        var v16 int = (fi7 ^ fi7)
                        _ = v16
                    }
                } else {
                    v17 := []int{(^(fi7)), 42, (fi7 * 2147483647)}
                    switch ("hello" + "a") {
                    case "hello":
                        var v18 int32 = v1
                        _ = v18
                    case "a":
                        var v19 float64 = (-2.25 + (0.0 - 3.14159))
                        _ = v19
                    case "":
                        var v20 uint = 42
                        _ = v20
                    }
                    v5 = uint16(v6)
                    _ = v17
                }
                _ = fi7
            }
            for fi21 := 0; fi21 < 4; fi21++ {
                v5 = v5
                switch ("a" + "") {
                case "a":
                    fmt.Println(v3, v3, v5)
                case "":
                    if (float64(0.0) > (3.14159 * -2.25)) {
                        var v22 int16 = 0
                        var v23 float64 = (0.0 / 1e10)
                        _ = v22
                        _ = v23
                    }
                case "foo bar":
                    v24 := []int{}
                    _ = v24
                }
                v1 = (v1 + -564)
                _ = fi21
            }
        } else {
            v1 = v1
        }
    case 2147483647:
        var v25 uint16 = v5
        _ = v25
    case -366:
        v1 = -2147483648
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 string = "foo bar"
        v28 := GenS2{F0: v2}
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS3{F0: 9223372036854775807, F1: (v6 & 4294967295)}
    var v30 bool = (-48 < 0)
    v31 := []int{-528, -9223372036854775808, -725}
    for fi32 := 0; fi32 < 9; fi32++ {
        if false {
            var v33 uint64 = (v6 << ((1) & 63))
            v3 -= uint32(v6)
            _ = v33
        } else {
            switch (fi32 % ((int(v6)) | 1)) {
            case -2147483648:
                for fi34 := 0; fi34 < 4; fi34++ {
                    var v35 int16 = -354
                    v6 = (v6 % ((2) | 1))
                    v6 -= 1
                    v36 := v4[2147483648]
                    v37 := GenS2{F0: v2}
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
            v38 := GenS2{F0: v3}
            _ = v38
        }
        _ = fi32
    }
    if (((v3 | (v2 * 537)) > (v3 >> ((255) & 63))) && true) {
        _ = len(v31)
        switch ("a" + "") {
        case "foo bar":
            for fi39 := 0; fi39 < 9; fi39++ {
                v40 := GenS3{F0: -9223372036854775808, F1: (v6 & 240)}
                switch (fi39 * (fi39 >> ((486) & 63))) {
                case -9223372036854775808:
                    v41 := GenS3{F0: 9223372036854775807, F1: 273}
                    _ = v41
                case 42:
                    if false {
                        fmt.Println(v2, v6)
                        if v30 {
                            fmt.Println(v2)
                            var v42 uint8 = 42
                            if true {
                                if false {
                                    for fi43 := 0; fi43 < 9; fi43++ {
                                        v44 := GenS3{F0: -9223372036854775808, F1: uint64(v2)}
                                        var v45 uint64 = (v6 % ((uint64(v42)) | 1))
                                        delete(v4, (v2 >> ((2) & 63)))
                                        _ = fi43
                                        _ = v44
                                        _ = v45
                                    }
                                    for fi46 := 0; fi46 < 3; fi46++ {
                                        fmt.Println(v2, v5, v42)
                                        switch ("a" + "\n\t") {
                                        case "a":
                                            v47 := GenS4{F0: 42, F1: v5}
                                            _ = v47
                                        }
                                        switch fi46 {
                                        case -2147483648:
                                            var v48 uint64 = (v6 >> ((65535) & 63))
                                            _ = v48
                                        case 4294967295:
                                            if true {
                                                switch fi46 {
                                                case 42:
                                                    switch fi39 {
                                                    case -106:
                                                        fmt.Println(v2, fi46)
                                                    case 46:
                                                        v49 := GenS4{F0: 1, F1: (v5 << ((42) & 63))}
                                                        _ = v49
                                                    default:
                                                        v5 = (v5 * v5)
                                                    }
                                                }
                                            }
                                        case -9223372036854775808:
                                            if len(v31) > 0 {
                                                v31[0] = (fi46 / ((fi39) | 1))
                                            }
                                        }
                                        _ = fi46
                                    }
                                }
                                fmt.Println(v5)
                                for fi50 := 0; fi50 < 1; fi50++ {
                                    v51 := GenS1{F0: 674, F1: fi50, F2: 0}
                                    switch (-(fi50)) {
                                    case 32768:
                                        if false {
                                            if (!v30) {
                                                v2 = v2
                                                var v52 uint8 = 42
                                                switch ("\n\t" + "") {
                                                case "foo bar":
                                                    var v53 float64 = (1e10 * (float64(0.0) / (1e10 - 0.0)))
                                                    _ = v53
                                                case "":
                                                    var v54 string = "hello"
                                                    _ = v54
                                                case "hello":
                                                    var v55 int64 = -1
                                                    _ = v55
                                                default:
                                                    v1 += v1
                                                }
                                                _ = v52
                                            } else {
                                                v30 = (((^(v2)) <= (v2 % ((v3) | 1))) || (255 != 2))
                                                _ = len(v31)
                                                if v30 {
                                                    switch "" {
                                                    case "\n\t":
                                                        if v30 {
                                                            if len(v31) > 0 {
                                                                v31[0] = fi50
                                                            }
                                                            delete(v4, (v2 ^ (v2 ^ 679)))
                                                            var v56 int8 = int8(v3)
                                                            _ = v56
                                                        } else {
                                                            v57 := v4[(v3 ^ v3)]
                                                            v3 += (v2 * (^(v2)))
                                                            var v58 float64 = 0.0
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                    case "a":
                                                        v6 -= 1
                                                    default:
                                                        v5 += v5
                                                    }
                                                    var v59 float32 = -2.25
                                                    v42 = 42
                                                    v59 += (3.14159 * 0.1)
                                                    _ = v59
                                                }
                                            }
                                            var v60 float64 = float64(0.1)
                                            _ = v60
                                        } else {
                                            if v30 {
                                                if true {
                                                    delete(v4, v2)
                                                }
                                                var v61 uint64 = (v6 >> ((uint(v5)) & 63))
                                                _ = v61
                                            } else {
                                                if true {
                                                    v30 = true
                                                    var v62 bool = (v30 || (1.5 >= 1.5))
                                                    if v30 {
                                                        fmt.Println(v30, v3)
                                                    }
                                                    _ = v62
                                                } else {
                                                    var v63 uint32 = v2
                                                    _ = v63
                                                }
                                            }
                                            v64 := GenS3{F0: 9223372036854775807, F1: (v6 | (v6 ^ 42))}
                                            var v65 uint32 = (v2 ^ 0)
                                            _ = v64
                                            _ = v65
                                        }
                                    case 1:
                                        fmt.Println(v6)
                                        if len(v31) > 0 {
                                            v31[0] = (fi39 + (-(fi50)))
                                        }
                                    case -2147483648:
                                        fmt.Println(v1, v5, v3)
                                    }
                                    switch 32767 {
                                    case -442:
                                        switch "hello" {
                                        case "":
                                            fmt.Println(v5)
                                        case "hello":
                                            fmt.Println(v2)
                                            v6 = v6
                                            if ((!false) || (fi50 < (fi39 % (((fi39 + (fi39 >> ((4294967295) & 63)))) | 1)))) {
                                                fmt.Println(v1)
                                            } else {
                                                fmt.Println(v6)
                                                v66 := GenS1{F0: 489, F1: fi39, F2: 1}
                                                fmt.Println(v3, fi50)
                                                v67 := GenS2{F0: 65535}
                                                v68 := GenS5{F0: -9223372036854775808, F1: (v1 ^ (v1 + v1)), F2: v1}
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            }
                                        }
                                    case -2147483648:
                                        var v69 bool = ((v42 == v42) && (!((v2 % ((1) | 1)) < (^(v2)))))
                                        _ = v69
                                    case -1:
                                        switch (fi50 >> ((uint(v3)) & 63)) {
                                        case 2147483647:
                                            v30 = v30
                                        case -9223372036854775808:
                                            switch "" {
                                            case "":
                                                fmt.Println(fi50, v5)
                                            case "foo bar":
                                                fmt.Println(v2, v30, v6)
                                            case "a":
                                                v70 := GenS6{F0: v3}
                                                _ = v70
                                            }
                                            v42 += 144
                                            fmt.Println(v3, v42)
                                        }
                                    }
                                    fmt.Println(v42)
                                    v71 := GenS1{F0: -212, F1: (^(fi39)), F2: 65535}
                                    _ = fi50
                                    _ = v51
                                    _ = v71
                                }
                            } else {
                                if true {
                                    fmt.Println(v3, v6)
                                    var v72 uint32 = 676
                                    v31 = append(v31, fi39)
                                    _ = v72
                                }
                                for fi73 := 0; fi73 < 18; fi73++ {
                                    v6 = v6
                                    _ = fi73
                                }
                            }
                            _ = v42
                        } else {
                            if true {
                                if (!v30) {
                                    v5 = v5
                                } else {
                                    v2 += v3
                                    fmt.Println(v3, v2, v6)
                                    v2 = (^(v2))
                                }
                                if len(v31) >= 2 {
                                    v31 = v31[1:len(v31)]
                                }
                                for fi74 := 0; fi74 < 13; fi74++ {
                                    fmt.Println(v2)
                                    v75 := GenS7{F0: v1, F1: int(v6)}
                                    _ = fi74
                                    _ = v75
                                }
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    v76 := GenS5{F0: -1, F1: (-(v1)), F2: (v1 >> ((2) & 63))}
                                    _ = v76
                                case "a":
                                    switch (fi39 | int(v1)) {
                                    case 42:
                                        v77 := GenS8{F0: false}
                                        _ = v77
                                    default:
                                        var v78 int64 = 42
                                        _ = v78
                                    }
                                default:
                                    v79 := GenS2{F0: v3}
                                    _ = v79
                                }
                                if v30 {
                                    delete(v4, (^(v3)))
                                    v6 = v6
                                }
                                v6 = (^(v6))
                            }
                            var v80 int8 = -128
                            _ = v80
                        }
                    } else {
                        v1 = v1
                        switch "a" {
                        case "":
                            v1 = (v1 / (((v1 & int32(v5))) | 1))
                        case "foo bar":
                            if ((0 == -1) || ((v2 ^ (v3 - v2)) < (^(v2)))) {
                                v81 := GenS7{F0: (-(v1)), F1: (fi39 + fi39)}
                                _ = v81
                            }
                        case "a":
                            if len(v31) >= 2 {
                                v31 = v31[1:len(v31)]
                            }
                        default:
                            v31 = append(v31, fi39)
                        }
                        fmt.Println(v3, v5)
                    }
                case 4294967295:
                    var v82 bool = v30
                    _ = v82
                default:
                    var v83 int16 = -32768
                    _ = v83
                }
                _ = fi39
                _ = v40
            }
        }
    }
    var v84 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v84
}
