package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint }
type GenS6 struct { F0 uint32; F1 uint64; F2 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16; F1 string; F2 int16 }
type GenS9 struct { F0 int32; F1 string }
type GenS10 struct { F0 bool; F1 uint16 }
type GenS11 struct { F0 uint32 }

func main() {
    var v1 int8 = 127
    var v2 int16 = -32768
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    var v5 int32 = 1
    v6 := []uint32{2147483648, 967, 1, 2}
    for fi7 := 0; fi7 < 1; fi7++ {
        if true {
            v1 = (v1 | v1)
            if (1 >= 2) {
                if false {
                    var v8 int64 = 42
                    for fi9 := 0; fi9 < 2; fi9++ {
                        v8 = -1
                        fmt.Println(fi9, v2, v8)
                        _ = fi9
                    }
                    _ = v8
                }
            } else {
                var v10 int16 = v2
                v10 = (v2 % ((-497) | 1))
                _ = v10
            }
            if false {
                for fi11 := 0; fi11 < 10; fi11++ {
                    switch fi11 {
                    case -9223372036854775808:
                        switch (fi11 / ((fi7) | 1)) {
                        case -930:
                            fmt.Println(fi11, v2, v1)
                        case -763:
                            if true {
                                v2 = (v2 | v2)
                                if (false || true) {
                                    var v12 uint16 = 947
                                    _ = v12
                                } else {
                                    if (true || false) {
                                        if (522 <= 32767) {
                                            v13 := v3[("a" + "foo bar")]
                                            _ = v13
                                        } else {
                                            v14 := v4["a"]
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            _ = v14
                                        }
                                    } else {
                                        v15 := GenS1{F0: 779, F1: (!true)}
                                        _ = v15
                                    }
                                    v16 := GenS1{F0: 2, F1: true}
                                    _ = v16
                                }
                                var v17 int = fi11
                                _ = v17
                            } else {
                                _ = len(v4)
                                var v18 int16 = v2
                                _ = v18
                            }
                        }
                    default:
                        var v19 float32 = 0.0
                        _ = v19
                    }
                    var v20 int16 = v2
                    var v21 uint = 42
                    _ = fi11
                    _ = v20
                    _ = v21
                }
            }
        } else {
            for fi22 := 0; fi22 < 16; fi22++ {
                v1 = (v1 / ((-128) | 1))
                _ = fi22
            }
            fmt.Println(fi7, v5, v1)
            v23 := GenS1{F0: 269, F1: false}
            _ = v23
        }
        var v24 int8 = v1
        _ = fi7
        _ = v24
    }
    fmt.Println(v5, v1, v5)
    switch "foo bar" {
    case "a":
        v2 = v2
    case "\n\t":
        if (425 > 384) {
            for fi25 := 0; fi25 < 6; fi25++ {
                v26 := GenS1{F0: 32767, F1: true}
                var v27 uint32 = 2147483648
                var v28 int = -139
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            v29 := GenS2{F0: 1, F1: 1}
            _ = v29
        }
    case "":
        for fi30 := 0; fi30 < 3; fi30++ {
            v1 += v1
            _ = fi30
        }
    }
    var v31 bool = ((!true) && true)
    v5 = (v5 << ((0) & 63))
    if v31 {
        v2 = v2
        if v31 {
            for fi32 := 0; fi32 < 19; fi32++ {
                var v33 uint32 = 391
                if true {
                    if true {
                        for fi34 := 0; fi34 < 8; fi34++ {
                            if v31 {
                                v6 = append(v6, v33)
                                switch (fi34 / (((^(fi32))) | 1)) {
                                case 0:
                                    switch fi32 {
                                    case 32768:
                                        var v35 int16 = int16(fi34)
                                        _ = v35
                                    case -9223372036854775808:
                                        v36 := v3[("hello" + "foo bar")]
                                        _ = v36
                                    case -135:
                                        v33 = 381
                                    default:
                                        v37 := GenS3{F0: -1}
                                        _ = v37
                                    }
                                case 4294967295:
                                    fmt.Println(v5)
                                case 1:
                                    switch fi32 {
                                    case -627:
                                        v5 += (v5 | v5)
                                    }
                                default:
                                    v38 := GenS4{F0: -663, F1: -9223372036854775808}
                                    _ = v38
                                }
                                _ = len(v6)
                            } else {
                                fmt.Println(v5, v5, fi32)
                                v1 = v1
                                var v39 uint64 = 392
                                _ = v39
                            }
                            _ = fi34
                        }
                        fmt.Println(v1, v5)
                        fmt.Println(v33)
                    }
                    v33 = v33
                } else {
                    var v40 uint64 = 123
                    if false {
                        v2 = v2
                    }
                    _ = v40
                }
                _ = fi32
                _ = v33
            }
            if (160 >= 0) {
                var v41 int8 = -62
                _ = v41
            }
            var v42 string = "a"
            _ = v42
        }
    } else {
        switch "" {
        case "\n\t":
            switch -291 {
            case 0:
                switch 0 {
                case 1:
                    switch 32768 {
                    case 42:
                        var v43 float32 = ((1.5 - 3.14159) / 1.0)
                        _ = v43
                    case -1:
                        for fi44 := 0; fi44 < 18; fi44++ {
                            if v31 {
                                var v45 int32 = (v5 << ((2) & 63))
                                v5 -= v5
                                _ = v45
                            }
                            _ = fi44
                        }
                    case 9223372036854775807:
                        switch ("foo bar" + "a") {
                        case "\n\t":
                            delete(v3, "hello")
                        case "":
                            fmt.Println(v5)
                        }
                    }
                }
            case 623:
                v5 -= v5
            case 127:
                switch ("a" + "\n\t") {
                case "foo bar":
                    switch 32767 {
                    case 1:
                        for fi46 := 0; fi46 < 14; fi46++ {
                            fmt.Println(v2)
                            _ = fi46
                        }
                    case -9223372036854775808:
                        if v31 {
                            if (!true) {
                                v5 += int32(v1)
                                v47 := GenS5{F0: 990, F1: (v2 >> ((255) & 63)), F2: 247}
                                _ = v47
                            }
                        }
                    case 0:
                        v1 += (v1 ^ v1)
                    }
                case "\n\t":
                    for fi48 := 0; fi48 < 12; fi48++ {
                        for fi49 := 0; fi49 < 20; fi49++ {
                            var v50 int16 = v2
                            v31 = true
                            v2 -= (v2 << ((2) & 63))
                            _ = fi49
                            _ = v50
                        }
                        _ = fi48
                    }
                }
            }
        case "hello":
            for fi51 := 0; fi51 < 11; fi51++ {
                if (fi51 != (fi51 >> ((843) & 63))) {
                    var v52 bool = true
                    for fi53 := 0; fi53 < 17; fi53++ {
                        _ = len(v4)
                        v54 := GenS6{F0: 664, F1: 1, F2: v52}
                        _ = fi53
                        _ = v54
                    }
                    _ = v52
                }
                var v55 int64 = 42
                v56 := GenS5{F0: 4294967295, F1: int16(v5), F2: 758}
                _ = fi51
                _ = v55
                _ = v56
            }
        case "foo bar":
            v1 -= v1
        }
        if (1 == 2147483647) {
            v57 := GenS7{F0: v1}
            var v58 int32 = v5
            v59 := GenS8{F0: 1, F1: "foo bar", F2: (v2 << ((18446744073709551615) & 63))}
            _ = v57
            _ = v58
            _ = v59
        } else {
            for fi60 := 0; fi60 < 4; fi60++ {
                if true {
                    var v61 string = "a"
                    var v62 float64 = 1e10
                    _ = v61
                    _ = v62
                }
                switch fi60 {
                case 32768:
                    if (!((v31 && true) && ((float32(1e10) <= -2.25) && ((!(!v31)) && (((!true) && v31) && (!(v31 && v31))))))) {
                        v63 := GenS5{F0: 42, F1: 19, F2: 880}
                        _ = v63
                    }
                case 0:
                    v1 = v1
                case 689:
                    if (float64(-2.25) != float64(float32(3.14159))) {
                        v64 := GenS9{F0: (v5 | (v5 - (v5 & v5))), F1: ("a" + "foo bar")}
                        fmt.Println(v1)
                        v31 = true
                        _ = v64
                    } else {
                        switch fi60 {
                        case 128:
                            for fi65 := 0; fi65 < 15; fi65++ {
                                v66 := GenS9{F0: (v5 >> ((2) & 63)), F1: "a"}
                                var v67 uint32 = 65535
                                v67 += v67
                                _ = fi65
                                _ = v66
                                _ = v67
                            }
                        case 9223372036854775807:
                            for fi68 := 0; fi68 < 15; fi68++ {
                                v5 = v5
                                switch -648 {
                                case 32768:
                                    var v69 int16 = (v2 - (v2 | 1000))
                                    _ = v69
                                case 32767:
                                    switch (fi68 & (fi60 - int(v5))) {
                                    case 1:
                                        delete(v3, "hello")
                                    default:
                                        var v70 uint8 = 241
                                        _ = v70
                                    }
                                default:
                                    v71 := GenS4{F0: 394, F1: -975}
                                    if true {
                                        v72 := GenS10{F0: v31, F1: 32767}
                                        v4[(("foo bar" + "foo bar") + "\n\t")] = v31
                                        switch (fi60 >> ((0) & 63)) {
                                        case 2147483647:
                                            v1 = (^(v1))
                                        default:
                                            var v73 int = fi60
                                            _ = v73
                                        }
                                        _ = v72
                                    }
                                    if (!v31) {
                                        v74 := GenS8{F0: 989, F1: "", F2: 42}
                                        v75 := GenS4{F0: 289, F1: -1}
                                        _ = v74
                                        _ = v75
                                    }
                                    _ = v71
                                }
                                if true {
                                    v76 := GenS2{F0: 42, F1: 1}
                                    var v77 float32 = 0.1
                                    _ = v76
                                    _ = v77
                                }
                                _ = fi68
                            }
                        default:
                            v5 += v5
                        }
                    }
                }
                _ = fi60
            }
        }
        v78 := GenS6{F0: 65535, F1: 18446744073709551615, F2: (-1 <= 0)}
        _ = v78
    }
    for fi79 := 0; fi79 < 10; fi79++ {
        fmt.Println(v1)
        v80 := GenS10{F0: v31, F1: 312}
        _ = fi79
        _ = v80
    }
    var v81 float32 = 1.5
    if v31 {
        v1 = v1
    } else {
        delete(v3, ("\n\t" + ""))
    }
    v81 -= v81
    for fi82 := 0; fi82 < 1; fi82++ {
        var v83 uint64 = 499
        switch fi82 {
        case 32767:
            if true {
                v83 = (v83 + 272)
            }
        case 1:
            var v84 uint = 255
            _ = v84
        case 42:
            if (0 < v2) {
                v85 := GenS11{F0: 65535}
                _ = v85
            }
        default:
            for fi86 := 0; fi86 < 16; fi86++ {
                v87 := GenS1{F0: 255, F1: true}
                v88 := v4["hello"]
                if (42 <= 0) {
                    var v89 int64 = 42
                    v90 := GenS4{F0: v89, F1: -9223372036854775808}
                    v83 += v83
                    _ = v89
                    _ = v90
                }
                _ = fi86
                _ = v87
                _ = v88
            }
        }
        _ = fi82
        _ = v83
    }
    if v31 {
        v81 -= v81
        v1 = (v1 - (v1 % ((127) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v81
}
