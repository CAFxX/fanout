package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 uint8; F2 bool }
type GenS7 struct { F0 int8; F1 uint32; F2 int32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int64 }
func genFn0(p0 int32) uint8 {
    var v4 int32 = (p0 - (p0 | (p0 | p0)))
    var v5 int = 9223372036854775807
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{v5}
        _ = fi6
        _ = v7
    }
    var v8 int = (v5 + (v5 & v5))
    if (3.14159 != 3.14159) {
        v9 := GenS1{F0: 0, F1: 203}
        _ = v9
    } else {
        v10 := GenS2{F0: 2, F1: 2}
        v11 := []int{-740, v8, (v5 ^ (v5 >> ((42) & 63)))}
        if len(v11) > 0 {
            v11[0] = (v8 & v8)
        }
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v8
    return 126
}
func genFn1(p0 uint16) int {
    if true {
        switch 4294967295 {
        case 32767:
            if (!(false || false)) {
                p0 = (p0 >> ((18446744073709551615) & 63))
                var v12 uint64 = 662
                v13 := []int{9223372036854775807, -2147483648, -471}
                _ = v12
                _ = v13
            }
        default:
            if (0.0 <= 0.0) {
                var v14 float64 = (float64(float32(0.1)) * 3.14159)
                _ = v14
            }
        }
        var v15 float64 = -2.25
        v16 := GenS3{F0: -1}
        _ = v15
        _ = v16
    }
    if true {
        p0 += p0
        v17 := []int{9223372036854775807, 0, 2147483647}
        for fi18 := 0; fi18 < 19; fi18++ {
            _ = len(v17)
            for fi19 := 0; fi19 < 11; fi19++ {
                if (!true) {
                    p0 -= (^(p0))
                    v20 := GenS3{F0: 2147483647}
                    var v21 bool = false
                    _ = v20
                    _ = v21
                } else {
                    v22 := GenS3{F0: fi19}
                    fmt.Println(p0, p0)
                    switch fi19 {
                    case 2147483647:
                        var v23 uint8 = 27
                        _ = v23
                    case 0:
                        var v24 float32 = (float32(3.14159) + 1e10)
                        _ = v24
                    }
                    _ = v22
                }
                _ = fi19
            }
            for fi25 := 0; fi25 < 11; fi25++ {
                v17 = append(v17, fi25)
                v26 := GenS4{F0: 42}
                var v27 int64 = 1
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi18
        }
        _ = v17
    }
    if true {
        v28 := GenS5{F0: -1}
        p0 = p0
        _ = v28
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 int32 = 84
        var v31 int8 = 0
        v32 := GenS4{F0: 65535}
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    _ = p0
    return 127
}

func main() {
    var v1 uint32 = 2
    var v2 int32 = 903
    var v3 uint = 2
    for fi33 := 0; fi33 < 15; fi33++ {
        if true {
            fmt.Println(v3, fi33)
            v34 := GenS2{F0: v1, F1: 193}
            _ = v34
        } else {
            v35 := []int{(-(fi33)), fi33, fi33}
            if (false || false) {
                v3 = 42
                for fi36 := 0; fi36 < 4; fi36++ {
                    switch 398 {
                    case -253:
                        var v37 uint64 = 178
                        _ = v37
                    case 0:
                        if len(v35) >= 2 {
                            v35 = v35[1:len(v35)]
                        }
                    case 1:
                        fmt.Println(fi33)
                    }
                    if true {
                        v38 := GenS4{F0: (v3 / ((255) | 1))}
                        switch ("" + "a") {
                        case "\n\t":
                            var v39 float32 = 1e10
                            _ = v39
                        case "a":
                            v1 = v1
                        }
                        _ = v38
                    }
                    if true {
                        switch "foo bar" {
                        case "a":
                            var v40 int16 = 634
                            _ = v40
                        case "foo bar":
                            switch "a" {
                            case "a":
                                v35 = append(v35, 9223372036854775807)
                            }
                        case "hello":
                            fmt.Println(v1, fi33)
                        }
                        if true {
                            v1 -= 2147483647
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                        }
                        var v41 uint32 = 0
                        v42 := GenS6{F0: 4294967295, F1: 127, F2: (-1 <= -32768)}
                        _ = v41
                        _ = v42
                    } else {
                        v35 = append(v35, 1)
                        if len(v35) > 0 {
                            v35[0] = 32768
                        }
                    }
                    if false {
                        if len(v35) > 0 {
                            v35[0] = fi36
                        }
                        v43 := GenS1{F0: v1, F1: 90}
                        _ = v43
                    } else {
                        var v44 bool = false
                        v45 := GenS5{F0: 32767}
                        var v46 float64 = 3.14159
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                    switch fi36 {
                    case -1:
                        _ = len(v35)
                    case 32768:
                        var v47 float64 = (0.0 - float64(-2.25))
                        _ = v47
                    case -451:
                        v35 = append(v35, (fi33 | (fi36 * (fi36 ^ 188))))
                    default:
                        var v48 uint32 = (v1 / ((2147483647) | 1))
                        var v49 uint8 = 7
                        _ = v48
                        _ = v49
                    }
                    if (!(-2147483648 == -2147483648)) {
                        var v50 bool = false
                        _ = v50
                    }
                    _ = fi36
                }
            }
            var v51 int16 = -586
            _ = v35
            _ = v51
        }
        v52 := GenS7{F0: 100, F1: (v1 & 94), F2: 2147483647}
        _ = fi33
        _ = v52
    }
    if true {
        var v53 int32 = v2
        v54 := GenS7{F0: 0, F1: 277, F2: -2147483648}
        _ = v53
        _ = v54
    } else {
        v1 += 2147483647
    }
    v55 := []int{-383, 42, 93, 9223372036854775807}
    v56 := GenS8{F0: -535}
    var v57 uint16 = 42
    v58 := GenS9{F0: -9223372036854775808}
    for fi59 := 0; fi59 < 5; fi59++ {
        var v60 uint16 = v57
        var v61 int16 = 1000
        v62 := GenS6{F0: 4294967295, F1: 2, F2: false}
        _ = fi59
        _ = v60
        _ = v61
        _ = v62
    }
    switch 127 {
    case 2147483647:
        var v63 uint32 = (^(v1))
        _ = v63
    default:
        v1 = v1
    }
    var v64 uint8 = 2
    var v65 uint = (v3 * (v3 | v3))
    var v66 uint64 = 1
    var v67 bool = false
    v66 = v66
    if true {
        switch (("foo bar" + "foo bar") + "\n\t") {
        case "foo bar":
            for fi68 := 0; fi68 < 14; fi68++ {
                if (((v57 < (v57 ^ (v57 & uint16(v1)))) || false) && true) {
                    if ((v66 / (((v66 ^ uint64(v3))) | 1)) != v66) {
                        fmt.Println(v67)
                    } else {
                        if (false || (v57 >= v57)) {
                            if v67 {
                                switch (fi68 ^ (fi68 - (fi68 ^ (fi68 ^ (fi68 * 42))))) {
                                case 0:
                                    var v69 uint8 = (v64 ^ (v64 ^ 1))
                                    _ = v69
                                case -964:
                                    var v70 uint16 = (v57 + (^(v57)))
                                    _ = v70
                                case -686:
                                    if len(v55) >= 2 {
                                        v55 = v55[1:len(v55)]
                                    }
                                }
                            }
                        } else {
                            for fi71 := 0; fi71 < 2; fi71++ {
                                v65 += (^(v65))
                                _ = fi71
                            }
                        }
                    }
                    v1 = v1
                }
                _ = fi68
            }
        case "a":
            switch "a" {
            case "hello":
                var v72 int64 = 0
                _ = v72
            case "foo bar":
                v73 := GenS2{F0: (v1 % ((2) | 1)), F1: v64}
                _ = v73
            case "\n\t":
                var v74 int32 = (v2 >> ((v65) & 63))
                _ = v74
            }
        default:
            for fi75 := 0; fi75 < 13; fi75++ {
                if v67 {
                    if v67 {
                        v57 -= (^(v57))
                    } else {
                        var v76 int32 = (^(v2))
                        if len(v55) > 0 {
                            v55[0] = (fi75 >> ((2) & 63))
                        }
                        _ = v76
                    }
                } else {
                    if (v65 <= v65) {
                        fmt.Println(v65, v64)
                    } else {
                        for fi77 := 0; fi77 < 20; fi77++ {
                            var v78 uint32 = 0
                            _ = fi77
                            _ = v78
                        }
                    }
                    v66 -= (v66 * v66)
                    if false {
                        v2 += 2147483647
                    }
                }
                v66 += v66
                _ = fi75
            }
        }
    } else {
        switch 128 {
        case 1:
            var v79 int8 = -31
            _ = v79
        }
        _ = len(v55)
        fmt.Println(v64, v65)
    }
    switch ((("\n\t" + "a") + ((("foo bar" + "a") + "hello") + ("" + ""))) + "") {
    case "":
        for fi80 := 0; fi80 < 9; fi80++ {
            var v81 uint8 = (^(v64))
            v82 := GenS3{F0: fi80}
            var v83 int = (fi80 % ((945) | 1))
            _ = fi80
            _ = v81
            _ = v82
            _ = v83
        }
    default:
        var v84 string = (("hello" + (("a" + "") + "\n\t")) + "\n\t")
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v57:", v57)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v64
    _ = v65
    _ = v66
    _ = v67
}
