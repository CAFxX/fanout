package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint; F2 int8 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 uint64; F1 int64; F2 string }
type GenS7 struct { F0 int; F1 uint16; F2 uint32 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 bool; F1 uint64; F2 bool }
func genFn0() string {
    switch 1 {
    case 128:
        if ((1.5 + ((1e10 / float32(1.5)) * 0.1)) <= ((float32(1e10) * 0.0) * (1e10 * 3.14159))) {
            var v5 uint64 = 18446744073709551615
            _ = v5
        } else {
            v6 := GenS1{F0: 127}
            var v7 int8 = 127
            if (!(!true)) {
                v8 := GenS1{F0: 127}
                switch 4294967295 {
                case 4294967295:
                    v9 := GenS1{F0: 255}
                    _ = v9
                default:
                    var v10 uint8 = 60
                    _ = v10
                }
                _ = v8
            }
            _ = v6
            _ = v7
        }
    case 576:
        switch -1 {
        case 739:
            v11 := GenS1{F0: 164}
            _ = v11
        }
    case -391:
        var v12 bool = true
        _ = v12
    }
    var v13 uint32 = 1
    var v14 int = -2147483648
    switch (-(v14)) {
    case -9223372036854775808:
        var v15 uint8 = 0
        _ = v15
    default:
        if true {
            v16 := []int{32768, (^(v14)), -2147483648}
            fmt.Println(v13, v14)
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 int = 507
                for fi19 := 0; fi19 < 4; fi19++ {
                    if false {
                        var v20 int32 = -2147483648
                        v13 = (v13 - v13)
                        _ = v20
                    } else {
                        v21 := GenS2{F0: -9223372036854775808, F1: 4294967295, F2: 0}
                        v14 = v18
                        _ = v21
                    }
                    var v22 uint8 = 2
                    v16 = append(v16, (^(fi17)))
                    _ = fi19
                    _ = v22
                }
                _ = fi17
                _ = v18
            }
            _ = v16
        } else {
            for fi23 := 0; fi23 < 13; fi23++ {
                for fi24 := 0; fi24 < 11; fi24++ {
                    v25 := []int{}
                    v26 := GenS1{F0: 75}
                    fmt.Println(v13)
                    var v27 uint = 1
                    if true {
                        v13 = (v13 / (((v13 + v13)) | 1))
                        if (v27 != (v27 ^ 18446744073709551615)) {
                            v14 = v14
                            v13 = v13
                            v27 = v27
                        } else {
                            if ((65535 < 0) && (388 < 255)) {
                                v13 -= (v13 << ((18446744073709551615) & 63))
                                var v28 uint64 = uint64(v27)
                                v29 := GenS3{F0: (v27 << (((v27 >> (((v27 | v27)) & 63))) & 63)), F1: uint8(v27), F2: 65535}
                                _ = v28
                                _ = v29
                            } else {
                                if true {
                                    if false {
                                        v30 := GenS4{F0: -1, F1: v27, F2: 127}
                                        _ = v30
                                    }
                                    v27 += v27
                                    v31 := GenS5{F0: 1000, F1: v27}
                                    _ = v31
                                }
                            }
                            switch "hello" {
                            case "":
                                if (true && false) {
                                    if true {
                                        var v32 uint8 = 230
                                        var v33 string = "foo bar"
                                        _ = v32
                                        _ = v33
                                    }
                                } else {
                                    switch -959 {
                                    case -9223372036854775808:
                                        v27 -= uint(v13)
                                    case 4294967295:
                                        var v34 uint32 = (^(v13))
                                        _ = v34
                                    case -277:
                                        if false {
                                            fmt.Println(fi24, fi24)
                                        }
                                    default:
                                        v13 = v13
                                    }
                                    v13 = (v13 ^ (v13 << ((0) & 63)))
                                }
                            case "foo bar":
                                v35 := GenS4{F0: -94, F1: 4294967295, F2: 127}
                                _ = v35
                            case "hello":
                                v36 := GenS6{F0: uint64(v27), F1: 935, F2: ("foo bar" + "hello")}
                                _ = v36
                            }
                            if ((-1 == 1000) && false) {
                                var v37 int64 = int64(v27)
                                v14 -= (fi23 % ((128) | 1))
                                if len(v25) > 0 {
                                    v25[0] = fi23
                                }
                                _ = v37
                            } else {
                                v25 = append(v25, fi23)
                            }
                        }
                    } else {
                        var v38 uint8 = 1
                        _ = v38
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                var v39 int = fi23
                for fi40 := 0; fi40 < 4; fi40++ {
                    v14 -= -772
                    var v41 int32 = -2147483648
                    switch ("a" + "hello") {
                    case "hello":
                        if false {
                            v42 := GenS4{F0: 1, F1: 970, F2: -128}
                            _ = v42
                        }
                    case "\n\t":
                        v43 := GenS7{F0: v39, F1: 42, F2: (v13 & v13)}
                        _ = v43
                    }
                    _ = fi40
                    _ = v41
                }
                _ = fi23
                _ = v39
            }
            v14 = v14
            v44 := []int{(v14 ^ -963)}
            _ = v44
        }
    }
    v14 = (v14 & 32767)
    _ = v13
    _ = v14
    return ""
}
func genFn1() uint32 {
    var v45 int8 = 0
    var v46 uint = 255
    for fi47 := 0; fi47 < 15; fi47++ {
        v48 := GenS1{F0: 127}
        _ = fi47
        _ = v48
    }
    v46 = (^(v46))
    _ = v45
    _ = v46
    return 65535
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = false
    var v3 int32 = 1
    var v4 float64 = (-2.25 + (-2.25 * ((((0.1 - 3.14159) / 3.14159) / 0.1) * 0.1)))
    v2 = (v2 || (!v2))
    v49 := GenS1{F0: 42}
    v2 = v2
    var v50 uint64 = 4294967295
    switch -1 {
    case 375:
        v3 += (v3 * 765)
    case 887:
        for fi51 := 0; fi51 < 1; fi51++ {
            fmt.Println(fi51, fi51, v2)
            for fi52 := 0; fi52 < 11; fi52++ {
                fmt.Println(v4, v2)
                _ = fi52
            }
            _ = fi51
        }
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        var v54 float32 = 3.14159
        _ = fi53
        _ = v54
    }
    if (true && v2) {
        var v55 uint8 = 255
        v55 += 2
        fmt.Println(v2, v2)
        _ = v55
    }
    if (255 <= 0) {
        v56 := []int{}
        _ = v56
    } else {
        var v57 string = "hello"
        for fi58 := 0; fi58 < 10; fi58++ {
            v59 := GenS8{F0: fi58, F1: 1}
            if (!v2) {
                switch "" {
                case "a":
                    v60 := []int{fi58}
                    _ = v60
                case "\n\t":
                    v4 = v4
                default:
                    var v61 int32 = (-(v3))
                    _ = v61
                }
                if (345 <= 999) {
                    var v62 float32 = 1e10
                    for fi63 := 0; fi63 < 7; fi63++ {
                        var v64 uint16 = 255
                        _ = fi63
                        _ = v64
                    }
                    _ = v62
                }
                var v65 int16 = 1000
                _ = v65
            } else {
                var v66 int64 = 1
                fmt.Println(v3, v50)
                switch v57 {
                case "":
                    var v67 bool = (((v50 & v50) == v50) && (!v2))
                    _ = v67
                }
                _ = v66
            }
            v68 := []int{(fi58 - (fi58 - fi58))}
            _ = fi58
            _ = v59
            _ = v68
        }
        _ = v57
    }
    v1[""] = ((1.5 >= float64(0.1)) && v2)
    if v2 {
        v69 := []int{0, -2147483648, 1}
        var v70 string = ("hello" + ("\n\t" + ""))
        for fi71 := 0; fi71 < 18; fi71++ {
            v72 := GenS2{F0: 1, F1: v50, F2: v50}
            var v73 uint16 = 149
            _ = fi71
            _ = v72
            _ = v73
        }
        _ = v69
        _ = v70
    } else {
        switch (((("a" + "") + "\n\t") + "foo bar") + ("\n\t" + ("a" + ("foo bar" + "hello")))) {
        case "\n\t":
            fmt.Println(v50, v4, v4)
        case "foo bar":
            v2 = v2
        case "a":
            v74 := GenS9{F0: v2, F1: v50, F2: v2}
            _ = v74
        default:
            switch "foo bar" {
            case "\n\t":
                switch 42 {
                case -302:
                    v4 = v4
                case 128:
                    var v75 int = -900
                    _ = v75
                default:
                    v76 := GenS3{F0: 262, F1: 128, F2: 1}
                    _ = v76
                }
            case "a":
                switch "hello" {
                case "":
                    var v77 int8 = 79
                    _ = v77
                default:
                    v3 -= v3
                }
            case "":
                var v78 int32 = 820
                _ = v78
            }
        }
        var v79 float32 = -2.25
        v50 += v50
        _ = v79
    }
    for fi80 := 0; fi80 < 18; fi80++ {
        v81 := GenS1{F0: 177}
        if v2 {
            v82 := GenS3{F0: 65535, F1: 62, F2: 1}
            fmt.Println(fi80)
            _ = v82
        } else {
            var v83 bool = v2
            var v84 uint8 = 42
            var v85 uint32 = 992
            _ = v83
            _ = v84
            _ = v85
        }
        if (3.14159 != 3.14159) {
            for fi86 := 0; fi86 < 8; fi86++ {
                var v87 uint8 = 255
                fmt.Println(fi80, v50)
                _ = fi86
                _ = v87
            }
        }
        _ = fi80
        _ = v81
    }
    var v88 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v88
}
