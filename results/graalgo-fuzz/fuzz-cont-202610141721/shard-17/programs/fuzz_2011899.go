package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int; F2 int32 }
type GenS5 struct { F0 int16; F1 int16; F2 uint }
type GenS6 struct { F0 int; F1 uint32; F2 int8 }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 uint64; F1 string; F2 bool }
type GenS9 struct { F0 int8; F1 int32; F2 string }
type GenS10 struct { F0 string; F1 uint8 }
type GenS11 struct { F0 int16 }

func main() {
    var v1 uint32 = 2
    var v2 int = -2147483648
    var v3 int16 = -32768
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 8; fi5++ {
        v3 += v3
        _ = fi5
    }
    var v6 int64 = 9223372036854775807
    if false {
        v7 := v4[(v1 + v1)]
        switch (^(v2)) {
        case 1:
            v4[v1] = false
        default:
            v1 += v1
        }
        _ = v7
    } else {
        fmt.Println(v2, v6)
        fmt.Println(v3)
    }
    if ("foo bar" != ("hello" + "hello")) {
        var v8 bool = (201 != 244)
        _ = v8
    }
    var v9 string = "hello"
    v10 := []int{v2}
    var v11 float32 = ((float32((1.5 - (1.5 + 3.14159))) + 1.5) * -2.25)
    fmt.Println(v1)
    for fi12 := 0; fi12 < 8; fi12++ {
        if (-2147483648 == 931) {
            _ = len(v4)
            var v13 int = v2
            if (((false || true) || (!(v1 <= v1))) || true) {
                for fi14 := 0; fi14 < 7; fi14++ {
                    v2 -= v2
                    _ = fi14
                }
                v1 -= v1
                v2 += (v13 / ((9223372036854775807) | 1))
            }
            _ = v13
        } else {
            v9 = ""
            for fi15 := 0; fi15 < 8; fi15++ {
                if (100 >= 127) {
                    switch fi12 {
                    case -1:
                        v16 := GenS1{F0: fi12}
                        _ = v16
                    case -906:
                        fmt.Println(fi15, v6, v2)
                    default:
                        fmt.Println(v1, v9, fi12)
                    }
                    _ = len(v10)
                } else {
                    fmt.Println(v11, fi12, v2)
                }
                switch (v2 + (v2 - int(v6))) {
                case 128:
                    if (72 == -1) {
                        v2 = fi12
                    }
                case 32768:
                    if ((v6 - 42) >= v6) {
                        if false {
                            if len(v10) > 0 {
                                v10[0] = fi12
                            }
                        } else {
                            switch -641 {
                            case 9223372036854775807:
                                delete(v4, 42)
                            default:
                                switch 32767 {
                                case -2147483648:
                                    _ = len(v4)
                                case -418:
                                    fmt.Println(v1)
                                default:
                                    v2 = fi15
                                }
                            }
                        }
                        if (42 >= 255) {
                            v11 = (1.5 / v11)
                            var v17 uint64 = 195
                            switch "hello" {
                            case "":
                                v10 = append(v10, v2)
                            case "hello":
                                var v18 float32 = ((v11 * 1.5) / v11)
                                _ = v18
                            default:
                                if (888 < 42) {
                                    switch 2147483647 {
                                    case -2147483648:
                                        v19 := GenS1{F0: v2}
                                        _ = v19
                                    case 319:
                                        fmt.Println(v3, v17, fi15)
                                    default:
                                        switch 689 {
                                        case -1:
                                            if true {
                                                v6 -= 469
                                                if (true || false) {
                                                    v20 := GenS1{F0: (fi15 & (v2 - -449))}
                                                    v21 := GenS1{F0: fi15}
                                                    if (498 <= 65535) {
                                                        fmt.Println(fi12, v6, v1)
                                                        var v22 string = ("foo bar" + v9)
                                                        _ = v22
                                                    } else {
                                                        v23 := GenS1{F0: fi15}
                                                        v10 = append(v10, fi15)
                                                        if ((v17 | uint64(v1)) > v17) {
                                                            v24 := GenS2{F0: v1, F1: 1000}
                                                            var v25 uint64 = 42
                                                            delete(v4, 686)
                                                            _ = v24
                                                            _ = v25
                                                        } else {
                                                            if ((v6 / ((v6) | 1)) != v6) {
                                                                if len(v10) >= 2 {
                                                                    v10 = v10[1:len(v10)]
                                                                }
                                                                if (206 < (v17 % ((0) | 1))) {
                                                                    var v26 bool = false
                                                                    delete(v4, 42)
                                                                    _ = v26
                                                                } else {
                                                                    switch 128 {
                                                                    case 0:
                                                                        var v27 uint = 42
                                                                        _ = v27
                                                                    case 42:
                                                                        if true {
                                                                            switch v9 {
                                                                            case "":
                                                                                v9 = v9
                                                                            case "foo bar":
                                                                                if ((false && false) || false) {
                                                                                    v1 = (v1 ^ (v1 & v1))
                                                                                    fmt.Println(v6, v3)
                                                                                }
                                                                            case "a":
                                                                                fmt.Println(v6, fi12, fi12)
                                                                                var v28 int8 = -103
                                                                                _ = v28
                                                                            }
                                                                            var v29 int8 = -1
                                                                            var v30 uint32 = v1
                                                                            _ = v29
                                                                            _ = v30
                                                                        }
                                                                    }
                                                                    var v31 float32 = v11
                                                                    var v32 int8 = int8(v17)
                                                                    _ = v31
                                                                    _ = v32
                                                                }
                                                                var v33 int16 = (-(v3))
                                                                v9 = v9
                                                                _ = v33
                                                            }
                                                            switch v2 {
                                                            case 42:
                                                                v4[42] = (42 >= -2147483648)
                                                            case 973:
                                                                v11 = v11
                                                            default:
                                                                v34 := GenS2{F0: v1, F1: v3}
                                                                _ = v34
                                                            }
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                        }
                                                        if true {
                                                            if (1 < 0) {
                                                                fmt.Println(v6, fi15, v9)
                                                                if true {
                                                                    if true {
                                                                        if ((((v9 + v9) > ((v9 + v9) + ((v9 + v9) + v9))) || true) && (791 <= -2147483648)) {
                                                                            v9 = ((v9 + (v9 + v9)) + "a")
                                                                            var v35 int64 = v6
                                                                            v2 += (fi15 ^ (v2 % ((v2) | 1)))
                                                                            _ = v35
                                                                        } else {
                                                                            fmt.Println(v17, v11, v6)
                                                                            fmt.Println(v17, v6, v1)
                                                                            v36 := GenS3{F0: 0, F1: 42}
                                                                            _ = v36
                                                                        }
                                                                        v11 = v11
                                                                    } else {
                                                                        fmt.Println(v1)
                                                                        switch v9 {
                                                                        case "foo bar":
                                                                            v37 := GenS4{F0: 127, F1: v2, F2: 0}
                                                                            _ = v37
                                                                        case "a":
                                                                            v38 := GenS5{F0: v3, F1: (^(v3)), F2: 4294967295}
                                                                            _ = v38
                                                                        }
                                                                    }
                                                                    v39 := GenS6{F0: 42, F1: 42, F2: 27}
                                                                    v40 := v4[v1]
                                                                    v41 := v4[v1]
                                                                    _ = v39
                                                                    _ = v40
                                                                    _ = v41
                                                                }
                                                            }
                                                            switch v9 {
                                                            case "foo bar":
                                                                if ((v9 + v9) <= v9) {
                                                                    v11 += 0.0
                                                                    v11 += v11
                                                                }
                                                            case "\n\t":
                                                                var v42 int16 = 1000
                                                                _ = v42
                                                            default:
                                                                var v43 int = fi12
                                                                _ = v43
                                                            }
                                                        }
                                                        _ = v23
                                                    }
                                                    _ = v20
                                                    _ = v21
                                                }
                                            } else {
                                                v11 -= v11
                                                switch fi15 {
                                                case 0:
                                                    switch v9 {
                                                    case "\n\t":
                                                        var v44 int8 = -81
                                                        _ = v44
                                                    case "a":
                                                        if false {
                                                            v45 := GenS7{F0: -1, F1: v9}
                                                            if true {
                                                                var v46 uint64 = 870
                                                                _ = v46
                                                            } else {
                                                                if ((0.0 / (1e10 / 3.14159)) == 0.0) {
                                                                    v17 = v17
                                                                    var v47 uint8 = 42
                                                                    switch v9 {
                                                                    case "a":
                                                                        var v48 uint32 = 65535
                                                                        _ = v48
                                                                    default:
                                                                        v47 += (v47 / ((uint8(v17)) | 1))
                                                                    }
                                                                    _ = v47
                                                                } else {
                                                                    v10 = append(v10, v2)
                                                                    switch 127 {
                                                                    case -2147483648:
                                                                        switch v9 {
                                                                        case "hello":
                                                                            var v49 float64 = float64(1e10)
                                                                            _ = v49
                                                                        }
                                                                    default:
                                                                        v6 = int64(v3)
                                                                    }
                                                                }
                                                            }
                                                            switch (fi12 << ((65535) & 63)) {
                                                            case -312:
                                                                v3 -= v3
                                                            default:
                                                                v9 = (v9 + v9)
                                                            }
                                                            _ = v45
                                                        }
                                                    default:
                                                        v17 += (v17 & 4294967295)
                                                    }
                                                case -9223372036854775808:
                                                    switch "\n\t" {
                                                    case "a":
                                                        v2 += (fi12 + fi15)
                                                    case "\n\t":
                                                        v9 = "foo bar"
                                                    case "":
                                                        var v50 int16 = 32767
                                                        _ = v50
                                                    }
                                                default:
                                                    switch 187 {
                                                    case 0:
                                                        var v51 float32 = 0.1
                                                        _ = v51
                                                    default:
                                                        v52 := GenS3{F0: 1, F1: v17}
                                                        _ = v52
                                                    }
                                                }
                                            }
                                        case 0:
                                            v4[(v1 % ((498) | 1))] = (149 > 2)
                                        default:
                                            if (!true) {
                                                if false {
                                                    var v53 uint64 = v17
                                                    v9 = v9
                                                    _ = v53
                                                } else {
                                                    v54 := GenS8{F0: 1, F1: "\n\t", F2: ((!(float32(1e10) == 0.1)) && false)}
                                                    v55 := GenS1{F0: (^(fi15))}
                                                    _ = v54
                                                    _ = v55
                                                }
                                                fmt.Println(fi12, v1, v3)
                                                fmt.Println(fi15, v3)
                                            } else {
                                                v10 = append(v10, (fi15 / ((-1) | 1)))
                                                v56 := v4[4294967295]
                                                fmt.Println(v11)
                                                _ = v56
                                            }
                                        }
                                    }
                                } else {
                                    v9 = ""
                                    var v57 uint32 = 457
                                    var v58 float64 = (1.5 / 3.14159)
                                    _ = v57
                                    _ = v58
                                }
                            }
                            v59 := GenS5{F0: (-(v3)), F1: 0, F2: 18446744073709551615}
                            v2 += int(v3)
                            v60 := GenS6{F0: 127, F1: (v1 >> ((uint(v17)) & 63)), F2: 1}
                            _ = v17
                            _ = v59
                            _ = v60
                        } else {
                            fmt.Println(fi12)
                            var v61 uint = 426
                            _ = v61
                        }
                    } else {
                        v62 := GenS9{F0: -1, F1: 0, F2: v9}
                        _ = v62
                    }
                }
                _ = fi15
            }
            var v63 uint = 908
            _ = v63
        }
        _ = fi12
    }
    v64 := GenS10{F0: (v9 + v9), F1: 128}
    v1 += 65535
    v6 -= (v6 >> ((143) & 63))
    delete(v4, (v1 % ((v1) | 1)))
    fmt.Println(v3)
    switch "\n\t" {
    case "\n\t":
        if true {
            if false {
                var v65 float64 = 1.5
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                if (938 >= 18446744073709551615) {
                    v4[(v1 + v1)] = true
                    v2 = 9223372036854775807
                }
                _ = v65
            }
            if (v2 != v2) {
                for fi66 := 0; fi66 < 18; fi66++ {
                    var v67 bool = false
                    v68 := v4[v1]
                    _ = fi66
                    _ = v67
                    _ = v68
                }
                var v69 int8 = -115
                _ = len(v10)
                _ = v69
            } else {
                if true {
                    delete(v4, (v1 | v1))
                    if (42 > 32767) {
                        for fi70 := 0; fi70 < 9; fi70++ {
                            var v71 int16 = v3
                            v72 := GenS11{F0: v71}
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                    } else {
                        var v73 float64 = (0.1 / 1.0)
                        _ = v73
                    }
                    fmt.Println(v1)
                } else {
                    var v74 float32 = v11
                    v1 += v1
                    _ = v74
                }
            }
        }
    case "":
        var v75 bool = false
        _ = v75
    default:
        v2 = v2
    }
    v76 := v4[v1]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v64
    _ = v76
    _ = v9
}
