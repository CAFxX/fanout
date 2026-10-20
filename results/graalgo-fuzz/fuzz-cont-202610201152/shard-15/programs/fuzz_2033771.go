package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint; F2 bool }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint; F1 int16; F2 string }

func main() {
    var v1 uint = 4294967295
    var v2 int16 = 1
    var v3 float32 = ((3.14159 / float32(1.5)) * 1.5)
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    v6 := []int{42, 42}
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: (v2 % (((v2 + (v2 % (((-(v2))) | 1)))) | 1)), F1: 2}
        switch fi7 {
        case -1:
            fmt.Println(v1, v3, v1)
        }
        _ = fi7
        _ = v8
    }
    var v9 uint = 0
    switch -377 {
    case 0:
        v10 := GenS2{F0: "a", F1: -1}
        _ = v10
    case 1:
        v3 = (1e10 * v3)
    }
    var v11 int16 = (v2 * (^(v2)))
    var v12 uint = (v9 - v9)
    switch ("hello" + ("a" + "")) {
    case "":
        v13 := GenS2{F0: ("\n\t" + ("hello" + "hello")), F1: 2147483647}
        _ = v13
    case "foo bar":
        v3 -= float32(1e10)
    default:
        v12 = 0
    }
    switch 2147483647 {
    case 377:
        for fi14 := 0; fi14 < 7; fi14++ {
            fmt.Println(v2)
            _ = fi14
        }
    case 2147483647:
        v15 := GenS1{F0: (v2 + (-(v11))), F1: 0}
        _ = v15
    case -2147483648:
        v16 := GenS3{F0: "hello", F1: (v11 / (((v11 ^ (-(v2)))) | 1))}
        _ = v16
    default:
        v3 = v3
    }
    if (128 <= 1) {
        var v17 uint32 = 1
        for fi18 := 0; fi18 < 16; fi18++ {
            if true {
                var v19 int16 = v2
                if true {
                    switch fi18 {
                    case 277:
                        v17 = v17
                    }
                } else {
                    for fi20 := 0; fi20 < 5; fi20++ {
                        v21 := v4[1]
                        v21 = v21
                        _ = fi20
                        _ = v21
                    }
                }
                if (!((v17 % (((v17 & v17)) | 1)) == v17)) {
                    if true {
                        var v22 int8 = 0
                        _ = v22
                    } else {
                        if false {
                            if (v19 <= 0) {
                                if (18446744073709551615 != 784) {
                                    if true {
                                        v19 += (v19 >> ((v9) & 63))
                                        v17 -= (v17 - v17)
                                    }
                                } else {
                                    var v23 float32 = v3
                                    _ = v23
                                }
                            } else {
                                v24 := GenS3{F0: "a", F1: v2}
                                v25 := GenS3{F0: "", F1: -32768}
                                for fi26 := 0; fi26 < 1; fi26++ {
                                    fmt.Println(fi26, fi26)
                                    v3 += v3
                                    v27 := GenS3{F0: "\n\t", F1: (v19 / ((1) | 1))}
                                    _ = fi26
                                    _ = v27
                                }
                                _ = v24
                                _ = v25
                            }
                            v3 = v3
                            for fi28 := 0; fi28 < 11; fi28++ {
                                switch "" {
                                case "a":
                                    if ((^(v1)) >= uint(v17)) {
                                        var v29 float32 = v3
                                        fmt.Println(v1)
                                        if (178 > 42) {
                                            fmt.Println(v2, v9)
                                            if (true && (false || (v17 >= (v17 >> (((v12 / ((v9) | 1))) & 63))))) {
                                                if (true || (false || (!(-9223372036854775808 >= 0)))) {
                                                    v30 := GenS4{F0: uint8(v17), F1: v12, F2: false}
                                                    fmt.Println(v11, fi18)
                                                    fmt.Println(fi28, v17, v3)
                                                    v31 := GenS5{F0: 42, F1: v19}
                                                    if (true && false) {
                                                        var v32 string = "a"
                                                        v33 := GenS6{F0: v32, F1: uint64(v17)}
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    v6 = append(v6, (fi28 << ((v9) & 63)))
                                                }
                                                v19 += v2
                                                var v34 int32 = 1
                                                v1 += 18446744073709551615
                                                fmt.Println(v29, v19)
                                                _ = v34
                                            }
                                        }
                                        _ = v29
                                    }
                                case "":
                                    v12 = (^(v9))
                                case "hello":
                                    var v35 string = (("" + "a") + "foo bar")
                                    var v36 string = ("a" + (((v35 + (v35 + v35)) + v35) + ""))
                                    _ = v35
                                    _ = v36
                                }
                                _ = fi28
                            }
                        }
                        v17 += v17
                        v37 := GenS6{F0: ("hello" + "\n\t"), F1: 702}
                        _ = v37
                    }
                } else {
                    v9 = v9
                }
                _ = v19
            }
            var v38 string = "foo bar"
            switch fi18 {
            case 9223372036854775807:
                if false {
                    v9 += v12
                }
            case 32768:
                for fi39 := 0; fi39 < 9; fi39++ {
                    var v40 float64 = ((-2.25 * (3.14159 - 1.5)) / (1e10 - ((1.5 - 3.14159) + 0.1)))
                    _ = fi39
                    _ = v40
                }
            case 4294967295:
                if false {
                    for fi41 := 0; fi41 < 8; fi41++ {
                        fmt.Println(fi18, v11, v3)
                        var v42 int16 = (-(v11))
                        _ = fi41
                        _ = v42
                    }
                }
            }
            _ = fi18
            _ = v38
        }
        v3 = float32(1e10)
        _ = v17
    }
    switch 32767 {
    case 1:
        v1 += 0
    case 128:
        var v43 uint32 = 2
        _ = v43
    case -700:
        if ((v12 - v1) < 0) {
            switch "a" {
            case "hello":
                v6 = append(v6, -2147483648)
            default:
                for fi44 := 0; fi44 < 16; fi44++ {
                    if true {
                        switch (fi44 | 1) {
                        case 0:
                            for fi45 := 0; fi45 < 19; fi45++ {
                                v46 := GenS3{F0: "foo bar", F1: v2}
                                switch ("hello" + "") {
                                case "foo bar":
                                    var v47 int64 = 0
                                    _ = v47
                                default:
                                    v2 += (v11 * v11)
                                }
                                v48 := GenS1{F0: v2, F1: 65535}
                                _ = fi45
                                _ = v46
                                _ = v48
                            }
                        case 127:
                            var v49 int16 = v11
                            _ = v49
                        default:
                            for fi50 := 0; fi50 < 18; fi50++ {
                                if true {
                                    v12 = (v9 ^ v1)
                                    if (!false) {
                                        v9 = (v12 * v1)
                                        _ = len(v5)
                                    } else {
                                        if false {
                                            v1 -= 304
                                            v51 := GenS6{F0: "hello", F1: 18446744073709551615}
                                            _ = v51
                                        }
                                        v12 += v1
                                    }
                                } else {
                                    v52 := GenS7{F0: 2, F1: -9223372036854775808, F2: 530}
                                    switch int(v11) {
                                    case 42:
                                        switch "a" {
                                        case "hello":
                                            _ = len(v6)
                                        case "a":
                                            v3 = (v3 * 0.1)
                                        case "\n\t":
                                            if false {
                                                v3 -= float32(3.14159)
                                                v9 = v12
                                                fmt.Println(v11, v9, fi44)
                                            } else {
                                                v53 := GenS4{F0: 127, F1: (v12 ^ 255), F2: true}
                                                if (-30 > 126) {
                                                    v54 := GenS8{F0: 205}
                                                    _ = v54
                                                } else {
                                                    fmt.Println(fi44, v12, v9)
                                                    if false {
                                                        v11 = 0
                                                        v55 := GenS9{F0: "\n\t"}
                                                        v56 := GenS9{F0: (("foo bar" + "foo bar") + "hello")}
                                                        fmt.Println(v1)
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                }
                                                _ = v53
                                            }
                                        }
                                        switch (fi50 / (((fi44 << ((v12) & 63))) | 1)) {
                                        case 282:
                                            var v57 int32 = 0
                                            _ = v57
                                        }
                                    }
                                    _ = v52
                                }
                                if false {
                                    v9 -= 65535
                                    v2 -= (-(v2))
                                    if false {
                                        v58 := GenS3{F0: "foo bar", F1: v11}
                                        var v59 int32 = 791
                                        _ = v58
                                        _ = v59
                                    }
                                }
                                var v60 uint16 = 65535
                                _ = fi50
                                _ = v60
                            }
                        }
                        for fi61 := 0; fi61 < 20; fi61++ {
                            if len(v6) > 0 {
                                v6[0] = (fi61 + fi44)
                            }
                            v62 := v5[855]
                            var v63 uint32 = 964
                            v63 = v63
                            _ = fi61
                            _ = v62
                            _ = v63
                        }
                        for fi64 := 0; fi64 < 6; fi64++ {
                            switch fi44 {
                            case -262:
                                v11 = -871
                            case 276:
                                var v65 float32 = v3
                                _ = v65
                            case 428:
                                var v66 uint8 = 42
                                _ = v66
                            }
                            v1 = v9
                            v9 -= v9
                            var v67 int = 32767
                            _ = len(v6)
                            _ = fi64
                            _ = v67
                        }
                    } else {
                        v9 = (v1 << (((^(v12))) & 63))
                    }
                    fmt.Println(v3)
                    if (1 != 9223372036854775807) {
                        v9 = 2
                        var v68 int8 = 100
                        _ = v68
                    } else {
                        var v69 uint = v12
                        var v70 int64 = -9223372036854775808
                        _ = v69
                        _ = v70
                    }
                    _ = fi44
                }
            }
            switch -689 {
            case -390:
                v1 = v12
            case 127:
                var v71 string = "\n\t"
                _ = v71
            case 32767:
                var v72 int8 = -1
                _ = v72
            }
        }
    default:
        delete(v5, 1)
    }
    for fi73 := 0; fi73 < 10; fi73++ {
        var v74 bool = false
        _ = fi73
        _ = v74
    }
    var v75 int32 = 0
    if false {
        switch -605 {
        case 476:
            v11 -= (v2 / ((v2) | 1))
        case 42:
            delete(v4, -9223372036854775808)
        case -7:
            v9 += (v1 >> (((v12 * (^(v1)))) & 63))
        default:
            if true {
                var v76 int16 = v11
                v2 -= int16(v12)
                _ = v76
            }
        }
    } else {
        if false {
            var v77 int8 = -1
            _ = v77
        } else {
            if false {
                for fi78 := 0; fi78 < 19; fi78++ {
                    fmt.Println(v9)
                    for fi79 := 0; fi79 < 14; fi79++ {
                        v80 := v5[42]
                        fmt.Println(v3)
                        if true {
                            v81 := v4[-9223372036854775808]
                            _ = v81
                        } else {
                            _ = len(v6)
                            var v82 bool = true
                            _ = v82
                        }
                        if false {
                            var v83 string = v80
                            v9 = v12
                            var v84 bool = false
                            _ = v83
                            _ = v84
                        }
                        fmt.Println(v75, v3)
                        _ = fi79
                        _ = v80
                    }
                    v2 += v2
                    _ = fi78
                }
            } else {
                v12 -= (v12 | v1)
            }
        }
        v5[42] = "\n\t"
        for fi85 := 0; fi85 < 9; fi85++ {
            switch 308 {
            case -1:
                switch fi85 {
                case 128:
                    v86 := GenS10{F0: v12, F1: v11, F2: "\n\t"}
                    _ = v86
                case -414:
                    v3 += (float32(-2.25) / (((1e10 * -2.25) / (v3 - v3)) / ((3.14159 + (v3 - (1.5 / v3))) * float32(-2.25))))
                }
            case 160:
                fmt.Println(v12, v3, v11)
            }
            _ = fi85
        }
    }
    v9 += (v1 % ((18446744073709551615) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v9
}
