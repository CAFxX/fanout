package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 string }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 int32 }

func main() {
    var v1 int8 = 10
    var v2 float64 = 0.1
    var v3 uint = 65535
    var v4 int16 = -32768
    var v5 int16 = (v4 - (v4 * int16(v3)))
    var v6 uint32 = 1
    if ((false || false) || (v3 >= 2)) {
        if (true && (true && false)) {
            v7 := GenS1{F0: (0 <= 1), F1: (0 != 546), F2: "hello"}
            switch 4294967295 {
            case 2147483647:
                v6 = (^(v6))
            default:
                v8 := []int{0}
                _ = v8
            }
            if false {
                v1 = v1
                v5 = v4
            } else {
                v9 := GenS1{F0: (v2 == float64((0.1 * (3.14159 + (-2.25 / 1.5))))), F1: false, F2: "a"}
                fmt.Println(v4, v1)
                var v10 uint16 = 65535
                _ = v10
                _ = v9
            }
            _ = v7
        } else {
            var v11 string = ("" + ("" + "hello"))
            var v12 bool = true
            if (((!false) && (v12 || false)) || v12) {
                fmt.Println(v5)
                for fi13 := 0; fi13 < 12; fi13++ {
                    v14 := []int{}
                    _ = fi13
                    _ = v14
                }
            }
            _ = v11
            _ = v12
        }
    } else {
        v1 += (v1 % ((v1) | 1))
        var v15 float64 = 0.0
        v16 := []int{2147483647, 571}
        _ = v15
        _ = v16
    }
    var v17 uint32 = (v6 | (v6 % ((v6) | 1)))
    for fi18 := 0; fi18 < 19; fi18++ {
        if ((v3 | (^(v3))) <= (v3 - 65535)) {
            v19 := []int{1}
            switch 0 {
            case 4294967295:
                var v20 float64 = v2
                _ = v20
            case -328:
                var v21 int16 = v4
                _ = v21
            case 127:
                var v22 uint8 = 42
                _ = v22
            }
            _ = v19
        } else {
            v23 := GenS2{F0: -1, F1: 18446744073709551615}
            var v24 uint = v3
            _ = v23
            _ = v24
        }
        var v25 int32 = 1
        for fi26 := 0; fi26 < 19; fi26++ {
            v27 := GenS1{F0: true, F1: (float64(1.5) <= v2), F2: (("hello" + "hello") + "hello")}
            var v28 uint16 = 255
            v29 := []int{fi26, (fi18 * 32768), (fi26 << (((v3 * (v3 & 42))) & 63)), 32767}
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi18
        _ = v25
    }
    var v30 int32 = 1
    var v31 int16 = v5
    var v32 uint16 = 42
    fmt.Println(v31, v4, v1)
    v2 = 1.5
    v33 := GenS3{F0: 501}
    fmt.Println(v2, v17, v30)
    for fi34 := 0; fi34 < 17; fi34++ {
        for fi35 := 0; fi35 < 19; fi35++ {
            var v36 int = (-(fi34))
            var v37 bool = (1.5 >= ((3.14159 - float32(v2)) * 1e10))
            if true {
                v1 = (v1 * v1)
            } else {
                if (false && v37) {
                    v38 := GenS4{F0: (v4 - (v31 / ((1000) | 1))), F1: (v30 - (v30 >> ((uint(v32)) & 63))), F2: v32}
                    _ = v38
                }
            }
            _ = fi35
            _ = v36
            _ = v37
        }
        v6 = (v6 / ((493) | 1))
        switch (fi34 * (fi34 * fi34)) {
        case -9223372036854775808:
            v31 = v4
        case 2147483647:
            if true {
                v5 -= v4
            } else {
                var v39 int16 = v31
                fmt.Println(v39)
                var v40 int16 = v39
                _ = v39
                _ = v40
            }
        case 9223372036854775807:
            var v41 string = "a"
            _ = v41
        default:
            v6 = uint32(v3)
        }
        _ = fi34
    }
    v42 := []int{32768, 1, 1, 919}
    switch -279 {
    case -1:
        for fi43 := 0; fi43 < 6; fi43++ {
            v31 = (v4 % ((-653) | 1))
            _ = fi43
        }
    case 42:
        fmt.Println(v30, v30, v32)
    }
    v2 += 3.14159
    v17 += v17
    v32 -= uint16(v3)
    if ((v6 + 98) != (v6 << (((^(v3))) & 63))) {
        v6 -= uint32(v3)
        var v44 int32 = v30
        for fi45 := 0; fi45 < 18; fi45++ {
            for fi46 := 0; fi46 < 12; fi46++ {
                if true {
                    v47 := GenS3{F0: 18446744073709551615}
                    if ("foo bar" <= "") {
                        if ((1e10 - 3.14159) != 3.14159) {
                            switch "" {
                            case "a":
                                switch (fi45 | fi45) {
                                case 9223372036854775807:
                                    switch 0 {
                                    case -628:
                                        switch "hello" {
                                        case "\n\t":
                                            if false {
                                                switch (("hello" + "hello") + "") {
                                                case "a":
                                                    if (167 > (v6 << (((v3 << ((v3) & 63))) & 63))) {
                                                        fmt.Println(v32, v32)
                                                        if len(v42) >= 2 {
                                                            v42 = v42[1:len(v42)]
                                                        }
                                                        if (1.5 <= 0.0) {
                                                            v4 += v5
                                                            v31 -= (v4 >> ((18446744073709551615) & 63))
                                                            if false {
                                                                if false {
                                                                    var v48 int16 = -32768
                                                                    v3 = v3
                                                                    var v49 int8 = v1
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                                fmt.Println(v5)
                                                                v31 += (v4 % ((-32768) | 1))
                                                            } else {
                                                                v2 -= v2
                                                            }
                                                            v50 := GenS4{F0: v4, F1: (v30 - (v44 / ((v44) | 1))), F2: v32}
                                                            _ = v50
                                                        }
                                                    } else {
                                                        var v51 float32 = 0.0
                                                        v42 = append(v42, fi45)
                                                        switch ("foo bar" + ("hello" + "")) {
                                                        case "hello":
                                                            v2 += 0.1
                                                        case "a":
                                                            v52 := GenS4{F0: 1000, F1: v44, F2: v32}
                                                            _ = v52
                                                        default:
                                                            if (-270 == 9223372036854775807) {
                                                                fmt.Println(v3, v5)
                                                                if false {
                                                                    v31 = v4
                                                                    var v53 string = (("" + ("hello" + "\n\t")) + "a")
                                                                    fmt.Println(v44)
                                                                    _ = v53
                                                                }
                                                                v6 += 2147483647
                                                                var v54 int64 = 42
                                                                switch fi46 {
                                                                case 127:
                                                                    var v55 uint32 = v17
                                                                    _ = v55
                                                                }
                                                                if (!false) {
                                                                    var v56 int16 = v4
                                                                    _ = v56
                                                                }
                                                                _ = v54
                                                            } else {
                                                                v42 = append(v42, 982)
                                                                switch int(v4) {
                                                                case 128:
                                                                    v3 -= v3
                                                                default:
                                                                    _ = len(v42)
                                                                }
                                                                fmt.Println(v1)
                                                            }
                                                            if ((("foo bar" + ("foo bar" + "a")) == "\n\t") && false) {
                                                                var v57 uint = v3
                                                                v58 := GenS3{F0: uint64(v32)}
                                                                _ = len(v42)
                                                                _ = v57
                                                                _ = v58
                                                            } else {
                                                                if len(v42) >= 2 {
                                                                    v42 = v42[1:len(v42)]
                                                                }
                                                                var v59 int8 = v1
                                                                _ = v59
                                                            }
                                                        }
                                                        _ = v51
                                                    }
                                                default:
                                                    var v60 int16 = v4
                                                    _ = v60
                                                }
                                            }
                                            var v61 uint64 = 4294967295
                                            var v62 float64 = 1e10
                                            _ = v61
                                            _ = v62
                                        case "foo bar":
                                            var v63 int = -9223372036854775808
                                            _ = v63
                                        case "":
                                            if false {
                                                var v64 bool = (!(int16(v1) < (v4 ^ -608)))
                                                fmt.Println(v44, v31)
                                                if len(v42) > 0 {
                                                    v42[0] = (fi46 + (fi45 >> (((v3 | (v3 % ((v3) | 1)))) & 63)))
                                                }
                                                v30 = (-(v30))
                                                _ = v64
                                            } else {
                                                var v65 float32 = 0.1
                                                if true {
                                                    if (!false) {
                                                        if (((v4 % (((v4 ^ v31)) | 1)) != v5) || true) {
                                                            if false {
                                                                if len(v42) >= 2 {
                                                                    v42 = v42[1:len(v42)]
                                                                }
                                                                v6 += v6
                                                                _ = len(v42)
                                                            } else {
                                                                switch (("\n\t" + "\n\t") + ((("a" + "foo bar") + ("a" + "hello")) + "a")) {
                                                                case "foo bar":
                                                                    var v66 string = "\n\t"
                                                                    _ = v66
                                                                case "\n\t":
                                                                    v67 := GenS5{F0: v3, F1: 0}
                                                                    _ = v67
                                                                default:
                                                                    var v68 int32 = (v44 / (((-(v44))) | 1))
                                                                    _ = v68
                                                                }
                                                            }
                                                        }
                                                        v69 := GenS3{F0: 4294967295}
                                                        _ = v69
                                                    }
                                                    v70 := GenS2{F0: (v5 << (((v3 << (((v3 - 4294967295)) & 63))) & 63)), F1: 490}
                                                    var v71 bool = ((42 > 2) || true)
                                                    v6 -= (v17 + (v17 - 2))
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v72 uint64 = 187
                                                        _ = v72
                                                    case "\n\t":
                                                        if (v44 != v44) {
                                                            var v73 int64 = -1
                                                            if v71 {
                                                                var v74 float64 = (v2 - v2)
                                                                v32 -= (v32 ^ (v32 & (v32 * (v32 | (v32 | 0)))))
                                                                _ = v74
                                                            } else {
                                                                switch (fi46 | 2147483647) {
                                                                case 32768:
                                                                    v17 = (v17 >> ((583) & 63))
                                                                }
                                                                if ((v71 && ((int16(v30) >= v31) || (234 < 1))) || (v17 > (v17 * v6))) {
                                                                    v75 := GenS4{F0: int16(v32), F1: int32(v1), F2: 2}
                                                                    if v71 {
                                                                        var v76 int32 = v30
                                                                        v6 = v17
                                                                        v1 = (v1 | v1)
                                                                        _ = v76
                                                                    }
                                                                    v3 -= 42
                                                                    _ = v75
                                                                }
                                                            }
                                                            v30 -= int32(fi46)
                                                            _ = v73
                                                        } else {
                                                            v77 := GenS5{F0: 2, F1: (^(v4))}
                                                            _ = v77
                                                        }
                                                    default:
                                                        var v78 int32 = v30
                                                        _ = v78
                                                    }
                                                    _ = v70
                                                    _ = v71
                                                }
                                                if true {
                                                    v30 += (v44 | v44)
                                                    if true {
                                                        v42 = append(v42, fi46)
                                                        fmt.Println(v44, v31)
                                                        var v79 int32 = (v44 & (v44 & 333))
                                                        _ = v79
                                                    }
                                                } else {
                                                    v42 = append(v42, 1)
                                                    fmt.Println(fi46, v31, v44)
                                                    var v80 float32 = (1e10 - (((1.5 - ((v65 * v65) / (0.1 + 0.1))) + v65) - 1.5))
                                                    _ = v80
                                                }
                                                _ = v65
                                            }
                                        }
                                    case 1:
                                        var v81 float32 = 1.5
                                        _ = v81
                                    }
                                }
                            case "foo bar":
                                v4 = 1
                            case "hello":
                                v1 = (v1 << ((255) & 63))
                            default:
                                v31 += 0
                            }
                            v4 += (v4 << (((v3 - 18446744073709551615)) & 63))
                            v82 := GenS1{F0: (102 >= 0), F1: true, F2: "hello"}
                            _ = v82
                        } else {
                            switch (fi45 ^ fi45) {
                            case 32768:
                                if (v31 != (-(v4))) {
                                    var v83 string = "a"
                                    _ = v83
                                }
                            default:
                                v2 += float64(3.14159)
                            }
                            if len(v42) >= 2 {
                                v42 = v42[1:len(v42)]
                            }
                        }
                        if false {
                            v17 = (v17 / ((v17) | 1))
                            var v84 int = 2147483647
                            v42 = append(v42, v84)
                            _ = v84
                        }
                        if len(v42) > 0 {
                            v42[0] = 32768
                        }
                    } else {
                        var v85 int = fi45
                        switch ((((("hello" + "") + "hello") + "") + "hello") + "\n\t") {
                        case "foo bar":
                            v86 := GenS2{F0: int16(v3), F1: 2}
                            _ = v86
                        default:
                            v4 = (v4 | (v4 / (((^(v4))) | 1)))
                        }
                        var v87 int64 = 9223372036854775807
                        _ = v85
                        _ = v87
                    }
                    switch (^(fi46)) {
                    case -9223372036854775808:
                        v88 := GenS6{F0: v30}
                        _ = v88
                    case 760:
                        _ = len(v42)
                        v4 += v31
                    case 247:
                        if len(v42) > 0 {
                            v42[0] = fi46
                        }
                    }
                    _ = v47
                }
                switch "a" {
                case "hello":
                    fmt.Println(v30)
                case "a":
                    var v89 uint = 2
                    _ = v89
                case "foo bar":
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v30, v31, v2)
                    }
                default:
                    fmt.Println(v5)
                }
                _ = fi46
            }
            _ = fi45
        }
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
