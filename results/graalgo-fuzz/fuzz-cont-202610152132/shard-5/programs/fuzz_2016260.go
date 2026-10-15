package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 uint; F2 uint8 }
type GenS8 struct { F0 uint16; F1 uint }
func genFn0() bool {
    var v5 float64 = (0.0 + 3.14159)
    fmt.Println(v5)
    v6 := GenS1{F0: -1, F1: 0, F2: (562 <= 65535)}
    v7 := GenS2{F0: 1, F1: 42}
    _ = v5
    _ = v6
    _ = v7
    return (!true)
}

func main() {
    var v1 int8 = -69
    var v2 float64 = (-2.25 * 1e10)
    var v3 int16 = 1
    var v4 float64 = (3.14159 * -2.25)
    if true {
        v8 := GenS3{F0: 111, F1: -2147483648}
        for fi9 := 0; fi9 < 13; fi9++ {
            if (!true) {
                for fi10 := 0; fi10 < 6; fi10++ {
                    v4 -= (v2 + v4)
                    v11 := []int{(fi10 << ((559) & 63)), fi10, fi9}
                    _ = fi10
                    _ = v11
                }
            }
            v1 += v1
            if (!(872 >= 18446744073709551615)) {
                v12 := []int{-997, (fi9 >> ((42) & 63)), 9223372036854775807}
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v12
            }
            _ = fi9
        }
        _ = v8
    }
    v2 = 0.0
    var v13 uint64 = 966
    fmt.Println(v2)
    switch -657 {
    case 645:
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int8 = 42
            _ = fi14
            _ = v15
        }
    case -1:
        v16 := GenS3{F0: 42, F1: 0}
        _ = v16
    default:
        var v17 int16 = -1
        _ = v17
    }
    switch ("\n\t" + (("hello" + ("\n\t" + ("" + ("foo bar" + "hello")))) + "")) {
    case "\n\t":
        var v18 int16 = 1000
        _ = v18
    case "hello":
        if false {
            var v19 string = "foo bar"
            v20 := GenS3{F0: int64(v3), F1: 187}
            if ("a" > v19) {
                v1 = v1
            }
            _ = v19
            _ = v20
        }
    }
    v21 := []int{2147483647, 9223372036854775807, 4294967295, -750}
    if true {
        for fi22 := 0; fi22 < 2; fi22++ {
            if (false && true) {
                var v23 bool = ((true && true) && (!false))
                _ = v23
            } else {
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            }
            for fi24 := 0; fi24 < 7; fi24++ {
                v3 -= 32767
                v4 -= (v2 / (v2 * float64(0.1)))
                var v25 uint8 = 0
                switch "a" {
                case "hello":
                    if (true || false) {
                        if (42 == 0) {
                            v3 -= v3
                            v26 := GenS4{F0: (v3 << ((42) & 63)), F1: v13, F2: 4294967295}
                            var v27 int64 = -1
                            var v28 uint16 = 32767
                            v29 := GenS5{F0: 0, F1: (v28 - (v28 * v28)), F2: v25}
                            var v30 int64 = v27
                            _ = v26
                            _ = v27
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        var v31 int32 = -1
                        var v32 uint64 = (v13 | v13)
                        _ = v31
                        _ = v32
                    } else {
                        fmt.Println(v4)
                        fmt.Println(fi22)
                    }
                case "foo bar":
                    if (!((v3 % ((v3) | 1)) != 32767)) {
                        v21 = append(v21, fi22)
                        var v33 int = (fi24 / ((-2147483648) | 1))
                        v34 := GenS6{F0: int64(fi24)}
                        _ = v33
                        _ = v34
                    }
                }
                v35 := GenS3{F0: 650, F1: int(v13)}
                var v36 bool = (9223372036854775807 != 1)
                v1 += -63
                _ = fi24
                _ = v25
                _ = v35
                _ = v36
            }
            _ = fi22
        }
        v2 = 0.0
        if true {
            fmt.Println(v1, v4, v4)
        }
    }
    switch "a" {
    case "\n\t":
        if (v13 < v13) {
            if false {
                if (!false) {
                    fmt.Println(v1)
                    if true {
                        _ = len(v21)
                        switch 9223372036854775807 {
                        case 32768:
                            if false {
                                fmt.Println(v4)
                            }
                        case -193:
                            v4 = v4
                        case -1:
                            var v37 uint32 = 77
                            _ = v37
                        }
                    } else {
                        v21 = append(v21, 9223372036854775807)
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                    }
                    fmt.Println(v2, v2, v3)
                } else {
                    if len(v21) > 0 {
                        v21[0] = 32767
                    }
                }
                for fi38 := 0; fi38 < 18; fi38++ {
                    fmt.Println(fi38)
                    _ = fi38
                }
                if (101 > (v1 ^ -1)) {
                    var v39 float64 = 0.0
                    _ = v39
                }
            }
        }
    case "":
        switch ("a" + "a") {
        case "hello":
            var v40 string = ("\n\t" + ("foo bar" + ("hello" + "")))
            _ = v40
        case "foo bar":
            var v41 int32 = 0
            _ = v41
        case "":
            v42 := GenS6{F0: 42}
            _ = v42
        default:
            for fi43 := 0; fi43 < 8; fi43++ {
                _ = len(v21)
                for fi44 := 0; fi44 < 20; fi44++ {
                    if (255 > 42) {
                        v3 = (-(v3))
                        var v45 int32 = int32(v3)
                        if false {
                            var v46 uint64 = 42
                            switch (fi44 % ((2147483647) | 1)) {
                            case -535:
                                var v47 float64 = 3.14159
                                var v48 float32 = 1.5
                                _ = v47
                                _ = v48
                            default:
                                var v49 int = (fi44 * 2147483647)
                                _ = v49
                            }
                            v50 := GenS5{F0: -9223372036854775808, F1: 19, F2: 255}
                            _ = v46
                            _ = v50
                        } else {
                            var v51 uint = 4294967295
                            v1 += v1
                            v52 := GenS7{F0: v45, F1: (v51 ^ (v51 & (v51 << (((v51 * 4294967295)) & 63)))), F2: 127}
                            v45 += v45
                            var v53 uint64 = v13
                            _ = v51
                            _ = v52
                            _ = v53
                        }
                        if true {
                            var v54 int64 = 42
                            _ = len(v21)
                            v1 = (v1 % ((v1) | 1))
                            if true {
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                switch ("" + "a") {
                                case "":
                                    var v55 uint16 = 207
                                    _ = v55
                                default:
                                    v56 := GenS4{F0: v3, F1: v13, F2: 2147483648}
                                    _ = v56
                                }
                                v45 = (v45 / ((int32(v1)) | 1))
                            }
                            if false {
                                v57 := GenS8{F0: 540, F1: 0}
                                _ = v57
                            }
                            v58 := GenS2{F0: (fi44 >> ((0) & 63)), F1: v3}
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            _ = v54
                            _ = v58
                        } else {
                            v21 = append(v21, (fi43 & fi43))
                            var v59 int = fi43
                            _ = v59
                        }
                        _ = v45
                    } else {
                        var v60 float64 = v2
                        if false {
                            v60 += v2
                            var v61 bool = true
                            _ = v61
                        } else {
                            v60 = ((v4 / v60) * -2.25)
                            if (255 >= 396) {
                                var v62 int64 = 278
                                v2 -= 1.5
                                if ((32767 >= 2) && true) {
                                    if (431 <= 255) {
                                        switch (fi44 / ((fi43) | 1)) {
                                        case 32767:
                                            fmt.Println(v13, v3)
                                        case -452:
                                            if false {
                                                switch int(v62) {
                                                case 431:
                                                    v60 += 3.14159
                                                    if len(v21) > 0 {
                                                        v21[0] = fi44
                                                    }
                                                    v63 := GenS6{F0: 0}
                                                    if len(v21) >= 2 {
                                                        v21 = v21[1:len(v21)]
                                                    }
                                                    v62 -= int64(v13)
                                                    var v64 bool = true
                                                    var v65 float64 = v2
                                                    _ = v63
                                                    _ = v64
                                                    _ = v65
                                                }
                                                if ((^(v13)) < 0) {
                                                    var v66 uint16 = 32767
                                                    switch ("hello" + "foo bar") {
                                                    case "\n\t":
                                                        var v67 int8 = (v1 % ((1) | 1))
                                                        _ = v67
                                                    case "hello":
                                                        var v68 float32 = ((0.0 - 1.5) - 1e10)
                                                        _ = v68
                                                    case "a":
                                                        v3 -= (v3 & (v3 >> ((355) & 63)))
                                                        v2 += (0.1 * (((v60 / v60) + (((v60 * float64((0.0 / float32(0.1)))) * v2) - (3.14159 + (((((0.1 + v60) * v60) / v2) * 0.0) * (v2 + 0.0))))) + 1e10))
                                                        var v69 bool = (true && false)
                                                        _ = v69
                                                    default:
                                                        var v70 int32 = 42
                                                        _ = v70
                                                    }
                                                    _ = v66
                                                }
                                            }
                                        case 486:
                                            v62 = (v62 * (v62 / ((9223372036854775807) | 1)))
                                        }
                                        if true {
                                            v4 = 1.5
                                            v3 = (v3 % ((v3) | 1))
                                            var v71 float64 = v4
                                            _ = v71
                                        } else {
                                            v3 = v3
                                            v2 = v2
                                            v3 = (v3 & (v3 + int16(fi43)))
                                        }
                                        switch ((("\n\t" + "\n\t") + "hello") + ("\n\t" + "")) {
                                        case "foo bar":
                                            var v72 int8 = (v1 - (v1 >> ((65535) & 63)))
                                            _ = v72
                                        default:
                                            v73 := GenS5{F0: (^(v62)), F1: 1, F2: 174}
                                            v60 += float64(0.0)
                                            if true {
                                                v4 = float64(1e10)
                                                var v74 int = (^(fi43))
                                                fmt.Println(v13, v3)
                                                _ = v74
                                            } else {
                                                var v75 uint = 292
                                                _ = v75
                                            }
                                            _ = v73
                                        }
                                    }
                                    var v76 uint = 0
                                    var v77 int64 = v62
                                    _ = v76
                                    _ = v77
                                } else {
                                    var v78 uint32 = 344
                                    _ = v78
                                }
                                _ = v62
                            }
                            v21 = append(v21, fi44)
                        }
                        v4 += (0.1 - 0.0)
                        _ = v60
                    }
                    _ = fi44
                }
                v13 = (v13 + (v13 - (v13 << ((1) & 63))))
                _ = fi43
            }
        }
    case "a":
        var v79 string = "\n\t"
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
