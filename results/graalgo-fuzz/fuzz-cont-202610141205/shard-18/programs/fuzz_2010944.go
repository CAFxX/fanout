package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 string; F2 uint32 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 uint; F1 int16; F2 uint32 }
type GenS10 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS11 struct { F0 int8; F1 uint8; F2 string }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int32; F1 int16; F2 bool }
type GenS14 struct { F0 uint32 }
func genFn0() bool {
    switch ("a" + "hello") {
    case "\n\t":
        var v7 uint = 255
        _ = v7
    case "foo bar":
        switch 32767 {
        case 127:
            var v8 float64 = (1e10 + 3.14159)
            _ = v8
        case 1:
            v9 := GenS1{F0: -7}
            _ = v9
        }
    default:
        switch ("hello" + "") {
        case "a":
            switch ((("" + "") + "\n\t") + "foo bar") {
            case "foo bar":
                if false {
                    var v10 int8 = -1
                    for fi11 := 0; fi11 < 1; fi11++ {
                        var v12 bool = (!true)
                        if (!(!true)) {
                            v13 := GenS2{F0: -9223372036854775808}
                            _ = v13
                        }
                        _ = fi11
                        _ = v12
                    }
                    if (false && false) {
                        for fi14 := 0; fi14 < 2; fi14++ {
                            v10 += v10
                            _ = fi14
                        }
                    }
                    _ = v10
                }
            default:
                for fi15 := 0; fi15 < 6; fi15++ {
                    switch (fi15 % (((-(fi15))) | 1)) {
                    case 32768:
                        if false {
                            fmt.Println(fi15)
                            for fi16 := 0; fi16 < 16; fi16++ {
                                if false {
                                    var v17 string = (("a" + "hello") + "")
                                    var v18 uint64 = 2
                                    _ = v17
                                    _ = v18
                                } else {
                                    v19 := []int{fi16}
                                    if true {
                                        if true {
                                            v20 := GenS3{F0: 1, F1: "\n\t", F2: 0}
                                            var v21 int32 = 2147483647
                                            v22 := GenS4{F0: true, F1: false}
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        } else {
                                            var v23 uint8 = 0
                                            var v24 int8 = int8(fi15)
                                            if len(v19) > 0 {
                                                v19[0] = fi16
                                            }
                                            _ = v23
                                            _ = v24
                                        }
                                    } else {
                                        if (0 > 2147483647) {
                                            if true {
                                                var v25 int = -2147483648
                                                v25 = (fi15 - v25)
                                                var v26 uint64 = 4294967295
                                                _ = v25
                                                _ = v26
                                            } else {
                                                v27 := GenS3{F0: 42, F1: ("a" + ""), F2: 2147483647}
                                                var v28 int32 = 451
                                                v29 := GenS4{F0: (false && ((1.5 <= float32(0.1)) || (914 < -1))), F1: true}
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            }
                                            var v30 int32 = 42
                                            v30 -= (v30 << ((65535) & 63))
                                            _ = v30
                                        } else {
                                            if true {
                                                switch (fi16 + (fi15 & (fi16 + fi16))) {
                                                case 32768:
                                                    var v31 uint64 = 596
                                                    _ = v31
                                                case 2147483647:
                                                    var v32 uint16 = 255
                                                    _ = v32
                                                case 127:
                                                    var v33 float64 = -2.25
                                                    _ = v33
                                                }
                                                fmt.Println(fi16, fi16)
                                            }
                                            if true {
                                                switch "" {
                                                case "":
                                                    v34 := GenS5{F0: -1}
                                                    _ = v34
                                                default:
                                                    var v35 int = fi16
                                                    _ = v35
                                                }
                                            }
                                        }
                                    }
                                    var v36 uint64 = 4294967295
                                    switch fi15 {
                                    case 961:
                                        var v37 int8 = -113
                                        _ = v37
                                    case 32768:
                                        v38 := GenS2{F0: 1}
                                        _ = v38
                                    case 32767:
                                        var v39 uint16 = 1
                                        v39 -= (v39 + uint16(v36))
                                        v36 = v36
                                        _ = v39
                                    default:
                                        var v40 uint = 4294967295
                                        _ = v40
                                    }
                                    _ = v19
                                    _ = v36
                                }
                                _ = fi16
                            }
                            var v41 uint16 = 1
                            _ = v41
                        } else {
                            v42 := GenS3{F0: 42, F1: "foo bar", F2: 4294967295}
                            _ = v42
                        }
                    case 892:
                        fmt.Println(fi15)
                    default:
                        var v43 float32 = 0.0
                        _ = v43
                    }
                    v44 := GenS3{F0: 32767, F1: "", F2: 2}
                    v45 := GenS3{F0: 0, F1: "\n\t", F2: 1}
                    _ = fi15
                    _ = v44
                    _ = v45
                }
            }
        case "foo bar":
            var v46 uint16 = 255
            _ = v46
        case "hello":
            var v47 int32 = 277
            _ = v47
        }
    }
    var v48 int64 = -1
    switch "foo bar" {
    case "a":
        v49 := []int{-9223372036854775808, -1, -1, 543}
        _ = v49
    default:
        v48 = (v48 + 493)
    }
    _ = v48
    return false
}

func main() {
    var v1 int32 = -1
    var v2 int8 = -111
    v3 := make(map[int]int)
    var v4 int16 = 0
    var v5 int32 = (v1 / (((v1 >> ((163) & 63))) | 1))
    v6 := make(map[uint32]uint32)
    for fi50 := 0; fi50 < 8; fi50++ {
        v1 += v5
        v51 := []int{(fi50 ^ (fi50 | (fi50 ^ (-(fi50))))), (fi50 >> ((347) & 63))}
        _ = fi50
        _ = v51
    }
    v1 -= (v1 | -652)
    if false {
        v2 -= (^(v2))
        var v52 float32 = 3.14159
        switch "hello" {
        case "foo bar":
            if true {
                for fi53 := 0; fi53 < 1; fi53++ {
                    v54 := []int{-1, (fi53 * fi53), 32768, -9223372036854775808}
                    var v55 int8 = -1
                    var v56 int64 = -390
                    _ = fi53
                    _ = v54
                    _ = v55
                    _ = v56
                }
                v4 += (v4 & v4)
                if (true || false) {
                    for fi57 := 0; fi57 < 1; fi57++ {
                        v58 := GenS6{F0: 2147483648}
                        _ = fi57
                        _ = v58
                    }
                }
            }
        case "hello":
            fmt.Println(v2, v52)
        case "":
            if true {
                var v59 uint8 = 69
                switch -217 {
                case -2147483648:
                    v4 += v4
                default:
                    fmt.Println(v59, v59, v2)
                }
                switch ("a" + "a") {
                case "":
                    v5 += (^(v1))
                case "hello":
                    switch 9223372036854775807 {
                    case 127:
                        v5 = (v5 << ((2) & 63))
                    case 4294967295:
                        if (uint(v59) != 42) {
                            v60 := GenS1{F0: 65}
                            _ = v60
                        }
                    case 1:
                        v61 := []int{9223372036854775807, -9223372036854775808, -9223372036854775808, 2147483647}
                        _ = v61
                    default:
                        v5 -= -880
                    }
                default:
                    if false {
                        for fi62 := 0; fi62 < 11; fi62++ {
                            if false {
                                for fi63 := 0; fi63 < 10; fi63++ {
                                    switch "foo bar" {
                                    case "":
                                        v64 := GenS4{F0: false, F1: (1 <= 100)}
                                        _ = v64
                                    default:
                                        fmt.Println(fi62, fi62, v2)
                                    }
                                    switch ((("hello" + "foo bar") + "hello") + "\n\t") {
                                    case "foo bar":
                                        var v65 float64 = (-2.25 * 3.14159)
                                        _ = v65
                                    default:
                                        v66 := GenS1{F0: v2}
                                        _ = v66
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        v67 := GenS7{F0: 2, F1: 65535}
                                        _ = v67
                                    }
                                    _ = fi63
                                }
                                var v68 string = "hello"
                                _ = v68
                            }
                            _ = fi62
                        }
                        for fi69 := 0; fi69 < 7; fi69++ {
                            v70 := GenS8{F0: -264, F1: 18446744073709551615}
                            if (!(false && false)) {
                                v71 := GenS7{F0: 2, F1: 926}
                                var v72 uint64 = 2
                                v73 := GenS5{F0: (v2 * 42)}
                                _ = v71
                                _ = v72
                                _ = v73
                            }
                            if true {
                                var v74 int16 = (v4 - (v4 - 1000))
                                v75 := GenS9{F0: 1, F1: 1, F2: 2147483647}
                                var v76 int8 = int8(v59)
                                _ = v74
                                _ = v75
                                _ = v76
                            }
                            _ = fi69
                            _ = v70
                        }
                    }
                }
                _ = v59
            } else {
                fmt.Println(v5, v1, v52)
                v77 := GenS4{F0: true, F1: (0 > (v1 | 0))}
                _ = v77
            }
        }
        _ = v52
    } else {
        v78 := GenS10{F0: (v4 / (((v4 & (-(v4)))) | 1)), F1: 128, F2: 861}
        _ = v78
    }
    _ = len(v3)
    fmt.Println(v2)
    _ = len(v6)
    v2 -= v2
    for fi79 := 0; fi79 < 6; fi79++ {
        switch "\n\t" {
        case "a":
            for fi80 := 0; fi80 < 4; fi80++ {
                v81 := GenS11{F0: (v2 ^ (-(v2))), F1: 1, F2: ""}
                _ = fi80
                _ = v81
            }
        }
        v5 = (v1 ^ (v5 * 1))
        fmt.Println(v5, v4, v4)
        _ = fi79
    }
    var v82 bool = (0 <= 66)
    for fi83 := 0; fi83 < 17; fi83++ {
        v1 -= (v1 + (v5 / (((v5 - v5)) | 1)))
        v5 -= (v5 / ((0) | 1))
        var v84 string = ("foo bar" + "hello")
        _ = fi83
        _ = v84
    }
    var v85 uint = 255
    v85 = v85
    for fi86 := 0; fi86 < 10; fi86++ {
        v87 := v6[576]
        if v82 {
            v88 := v6[(v87 ^ v87)]
            v89 := GenS12{F0: (fi86 ^ int(v87))}
            v5 = (^(v5))
            _ = v88
            _ = v89
        }
        _ = fi86
        _ = v87
    }
    v90 := GenS13{F0: (v1 % (((^(v1))) | 1)), F1: v4, F2: (42 != 84)}
    v91 := GenS14{F0: 2147483648}
    if v82 {
        var v92 int = -252
        for fi93 := 0; fi93 < 2; fi93++ {
            _ = len(v6)
            _ = fi93
        }
        _ = v92
    } else {
        v5 = (v5 | v5)
        var v94 int = -339
        _ = v94
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v82:", v82)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v82
    _ = v85
    _ = v90
    _ = v91
}
