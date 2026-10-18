package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int16 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint; F1 uint16; F2 int32 }
type GenS8 struct { F0 int8; F1 int8; F2 int }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 uint; F1 int }
type GenS11 struct { F0 uint8; F1 int8 }
type GenS12 struct { F0 int16 }

func main() {
    var v1 int8 = -1
    var v2 bool = (!(4294967295 >= 388))
    var v3 uint = 559
    for fi4 := 0; fi4 < 1; fi4++ {
        v3 = (v3 << (((^(v3))) & 63))
        fmt.Println(v1)
        _ = fi4
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        v3 = (v3 - (v3 | (v3 | 2)))
        _ = fi5
    }
    var v6 uint8 = 2
    for fi7 := 0; fi7 < 18; fi7++ {
        v2 = true
        v8 := []int{487, fi7, 473}
        v2 = (!v2)
        _ = fi7
        _ = v8
    }
    if (3.14159 < 1.5) {
        for fi9 := 0; fi9 < 6; fi9++ {
            v1 = 100
            _ = fi9
        }
        v10 := []int{2147483647, int(v6), 9223372036854775807, -276}
        var v11 int16 = 1
        _ = v10
        _ = v11
    }
    var v12 int32 = -2147483648
    var v13 uint32 = 625
    var v14 uint32 = (v13 % ((v13) | 1))
    v15 := GenS1{F0: 1}
    fmt.Println(v6, v1)
    if v2 {
        switch ("foo bar" + "foo bar") {
        case "hello":
            v13 += (v14 / ((219) | 1))
        case "\n\t":
            for fi16 := 0; fi16 < 18; fi16++ {
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := GenS2{F0: 2, F1: v2, F2: 0}
                    var v19 int32 = (-(v12))
                    if (0 < 42) {
                        v20 := []int{}
                        v21 := GenS2{F0: 4294967295, F1: (false || v2), F2: 42}
                        v20 = append(v20, int(v1))
                        _ = v20
                        _ = v21
                    } else {
                        if (("hello" + (("foo bar" + "") + "hello")) != "") {
                            v22 := []int{(fi16 + 1), (fi16 & -9223372036854775808), fi16}
                            if v2 {
                                v2 = ("a" < ("foo bar" + ""))
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            }
                            if (!(v2 || true)) {
                                fmt.Println(fi17, v2, v6)
                            } else {
                                switch 352 {
                                case 42:
                                    v1 += v1
                                default:
                                    var v23 string = "foo bar"
                                    var v24 int16 = 1
                                    fmt.Println(v12)
                                    _ = v23
                                    _ = v24
                                }
                                fmt.Println(v6, v1, v14)
                                v14 = v14
                            }
                            fmt.Println(fi17)
                            fmt.Println(v14)
                            if (("a" + "foo bar") > "foo bar") {
                                v25 := GenS2{F0: 2, F1: (float64(float32(-2.25)) != (1e10 / 1.0)), F2: -32768}
                                var v26 uint64 = 2
                                v27 := GenS1{F0: (v1 + 127)}
                                _ = v25
                                _ = v26
                                _ = v27
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    var v28 int16 = 42
                                    _ = v28
                                case "a":
                                    if (1 == 42) {
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                        if v2 {
                                            var v29 uint8 = 0
                                            var v30 string = "a"
                                            _ = v29
                                            _ = v30
                                        } else {
                                            if len(v22) > 0 {
                                                v22[0] = 128
                                            }
                                            if v2 {
                                                if v2 {
                                                    var v31 int = (fi16 << ((686) & 63))
                                                    _ = v31
                                                } else {
                                                    var v32 int64 = 0
                                                    _ = v32
                                                }
                                                var v33 uint32 = (v14 * v13)
                                                v1 = v1
                                                v13 -= 2147483648
                                                _ = v33
                                            } else {
                                                v34 := GenS1{F0: 82}
                                                v14 = v13
                                                if v2 {
                                                    v2 = (!((v14 % ((v13) | 1)) == v13))
                                                }
                                                _ = v34
                                            }
                                        }
                                        if (false && v2) {
                                            var v35 int = (fi16 + fi17)
                                            _ = v35
                                        }
                                        var v36 int = 1
                                        if true {
                                            fmt.Println(v19)
                                            switch (("\n\t" + "\n\t") + "foo bar") {
                                            case "":
                                                var v37 int64 = int64(v19)
                                                _ = v37
                                            }
                                            var v38 uint32 = 1
                                            _ = v38
                                        }
                                        _ = v36
                                    } else {
                                        v39 := GenS1{F0: v1}
                                        _ = len(v22)
                                        v40 := GenS3{F0: 1000, F1: fi17}
                                        fmt.Println(v3, v19, v12)
                                        var v41 uint32 = (v14 / (((v13 ^ (^(v13)))) | 1))
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                case "\n\t":
                                    v13 = v14
                                default:
                                    switch "\n\t" {
                                    case "":
                                        var v42 uint16 = 181
                                        _ = v42
                                    case "hello":
                                        switch (("foo bar" + "foo bar") + "a") {
                                        case "":
                                            var v43 int16 = 0
                                            _ = v43
                                        case "hello":
                                            if (false && true) {
                                                switch 467 {
                                                case -439:
                                                    var v44 int = fi16
                                                    _ = v44
                                                case -698:
                                                    if true {
                                                        v45 := GenS4{F0: v3, F1: v14, F2: uint32(v3)}
                                                        v46 := GenS5{F0: (("foo bar" + "") + ("" + "a")), F1: (fi17 * int(v6))}
                                                        fmt.Println(v14)
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        switch "foo bar" {
                                                        case "":
                                                            if ((v3 << ((1) & 63)) <= v3) {
                                                                v47 := GenS4{F0: (^(v3)), F1: uint32(v6), F2: (v14 * 564)}
                                                                if len(v22) >= 2 {
                                                                    v22 = v22[1:len(v22)]
                                                                }
                                                                v14 += v13
                                                                _ = v47
                                                            } else {
                                                                fmt.Println(v12, v12)
                                                                v2 = (!false)
                                                            }
                                                        case "hello":
                                                            _ = len(v22)
                                                        case "foo bar":
                                                            switch "a" {
                                                            case "":
                                                                fmt.Println(fi16, v14, v2)
                                                            case "hello":
                                                                if len(v22) > 0 {
                                                                    v22[0] = fi17
                                                                }
                                                            default:
                                                                if ("foo bar" >= "a") {
                                                                    var v48 int32 = -726
                                                                    _ = v48
                                                                } else {
                                                                    if false {
                                                                        var v49 uint = uint(v13)
                                                                        _ = v49
                                                                    } else {
                                                                        var v50 bool = (!v2)
                                                                        _ = v50
                                                                    }
                                                                    if len(v22) >= 2 {
                                                                        v22 = v22[1:len(v22)]
                                                                    }
                                                                    fmt.Println(v3, v14, v3)
                                                                    v51 := GenS2{F0: 18446744073709551615, F1: (4294967295 < 144), F2: 32767}
                                                                    _ = v51
                                                                }
                                                            }
                                                            v2 = (v2 || v2)
                                                            var v52 int16 = 1000
                                                            _ = v52
                                                        default:
                                                            if v2 {
                                                                var v53 uint64 = 42
                                                                var v54 bool = v2
                                                                if len(v22) > 0 {
                                                                    v22[0] = -2147483648
                                                                }
                                                                _ = v53
                                                                _ = v54
                                                            } else {
                                                                v14 -= (v14 - (v14 ^ (v14 << ((v3) & 63))))
                                                            }
                                                        }
                                                        var v55 float32 = 0.1
                                                        _ = v55
                                                    }
                                                }
                                            } else {
                                                if (275 == 42) {
                                                    var v56 uint8 = 0
                                                    var v57 uint = 65535
                                                    _ = v56
                                                    _ = v57
                                                }
                                            }
                                        case "a":
                                            if ("hello" <= "") {
                                                if (65535 <= 255) {
                                                    fmt.Println(v6, fi17)
                                                }
                                                v2 = (v3 > v3)
                                            }
                                        default:
                                            v3 = 18446744073709551615
                                        }
                                        var v58 uint64 = 1
                                        v19 += v12
                                        _ = v58
                                    case "\n\t":
                                        if len(v22) > 0 {
                                            v22[0] = 2147483647
                                        }
                                    default:
                                        fmt.Println(v3)
                                    }
                                }
                                v2 = v2
                            }
                            v22 = append(v22, (fi17 + fi16))
                            if true {
                                _ = len(v22)
                                v59 := GenS2{F0: 469, F1: false, F2: -333}
                                v6 = 90
                                switch "\n\t" {
                                case "\n\t":
                                    var v60 uint = uint(v6)
                                    _ = v60
                                }
                                v6 -= (v6 % ((42) | 1))
                                _ = v59
                            }
                            _ = v22
                        } else {
                            v19 -= (v12 * (^(v19)))
                            var v61 uint64 = 995
                            _ = v61
                        }
                        if (v3 <= v3) {
                            if (0.1 < -2.25) {
                                switch fi17 {
                                case 4294967295:
                                    if (("foo bar" + (("\n\t" + "foo bar") + ("foo bar" + ""))) > ("a" + "a")) {
                                        v13 -= (v14 | v14)
                                        v62 := []int{}
                                        var v63 uint = v3
                                        _ = v62
                                        _ = v63
                                    } else {
                                        var v64 uint8 = v6
                                        v65 := []int{}
                                        _ = v64
                                        _ = v65
                                    }
                                case -462:
                                    fmt.Println(v1, v14)
                                case -638:
                                    switch 196 {
                                    case 500:
                                        v66 := []int{32768, 127}
                                        _ = v66
                                    }
                                default:
                                    v3 += 2
                                }
                            }
                            v67 := []int{(fi17 >> ((255) & 63)), 128}
                            if true {
                                switch 9223372036854775807 {
                                case 2147483647:
                                    var v68 int = (fi17 ^ fi16)
                                    if (!(((v12 * (v12 << ((v3) & 63))) <= (v12 % (((^(v12))) | 1))) && ((1e10 * 3.14159) > (float32(0.0) * -2.25)))) {
                                        if v2 {
                                            fmt.Println(v1, fi17)
                                            v69 := GenS6{F0: (602 > 0)}
                                            switch 742 {
                                            case 378:
                                                switch (fi16 | (v68 ^ (^(fi16)))) {
                                                case 128:
                                                    fmt.Println(v68)
                                                }
                                            default:
                                                var v70 int8 = v1
                                                _ = v70
                                            }
                                            var v71 int32 = -1
                                            v3 += (^(v3))
                                            switch (v68 >> ((1) & 63)) {
                                            case 32768:
                                                fmt.Println(v19, fi16)
                                            case -9223372036854775808:
                                                if len(v67) > 0 {
                                                    v67[0] = v68
                                                }
                                            default:
                                                fmt.Println(v12, v71)
                                            }
                                            var v72 int64 = -546
                                            switch (fi16 >> (((v3 + 735)) & 63)) {
                                            case -711:
                                                v19 -= -2147483648
                                            default:
                                                if false {
                                                    switch (fi16 - fi17) {
                                                    case 545:
                                                        fmt.Println(v68)
                                                    }
                                                    if len(v67) > 0 {
                                                        v67[0] = (fi16 / (((^(fi16))) | 1))
                                                    }
                                                }
                                            }
                                            if (v71 >= (v19 ^ 428)) {
                                                v72 = v72
                                                v67 = append(v67, (v68 & (fi17 & (^(v68)))))
                                                v73 := GenS3{F0: int16(v13), F1: 256}
                                                _ = v73
                                            }
                                            var v74 uint32 = uint32(v6)
                                            v75 := GenS3{F0: 0, F1: -785}
                                            _ = v69
                                            _ = v71
                                            _ = v72
                                            _ = v74
                                            _ = v75
                                        } else {
                                            if len(v67) > 0 {
                                                v67[0] = fi16
                                            }
                                            switch -1 {
                                            case -9223372036854775808:
                                                var v76 int16 = 32767
                                                _ = v76
                                            case -775:
                                                v6 += v6
                                            case 42:
                                                if ((v1 + (v1 ^ 127)) < (v1 - (v1 - -128))) {
                                                    v14 -= v13
                                                }
                                            }
                                            var v77 string = "a"
                                            _ = v77
                                        }
                                    } else {
                                        v3 = (v3 ^ (^(v3)))
                                        var v78 int16 = 0
                                        _ = v78
                                    }
                                    var v79 string = "\n\t"
                                    _ = v68
                                    _ = v79
                                case 49:
                                    v19 = v19
                                    if (v14 != (^(v13))) {
                                        if v2 {
                                            if len(v67) >= 2 {
                                                v67 = v67[1:len(v67)]
                                            }
                                            var v80 int = fi16
                                            v81 := GenS7{F0: (^(v3)), F1: 255, F2: v12}
                                            _ = v80
                                            _ = v81
                                        }
                                        v14 -= v13
                                    } else {
                                        switch (fi16 >> ((v3) & 63)) {
                                        case 4294967295:
                                            var v82 uint8 = v6
                                            _ = v82
                                        case -151:
                                            v83 := GenS1{F0: v1}
                                            _ = v83
                                        case 127:
                                            v3 = (^(v3))
                                        }
                                    }
                                default:
                                    var v84 int32 = v12
                                    _ = v84
                                }
                            }
                            _ = v67
                        } else {
                            v12 -= v19
                            fmt.Println(v3)
                            var v85 int = (fi17 * 653)
                            _ = v85
                        }
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
        default:
            for fi86 := 0; fi86 < 20; fi86++ {
                var v87 int8 = 127
                v88 := []int{(fi86 << ((v3) & 63))}
                for fi89 := 0; fi89 < 6; fi89++ {
                    if len(v88) >= 2 {
                        v88 = v88[1:len(v88)]
                    }
                    var v90 uint64 = 18446744073709551615
                    v2 = v2
                    _ = fi89
                    _ = v90
                }
                _ = fi86
                _ = v87
                _ = v88
            }
        }
        v14 = (v14 & 666)
    }
    switch ("\n\t" + ("a" + "a")) {
    case "":
        v14 = 2147483648
    }
    var v91 float32 = 1.5
    v2 = ("a" < "\n\t")
    v92 := []int{42, 1, -643, 0}
    v14 += 0
    switch "foo bar" {
    case "foo bar":
        switch -9223372036854775808 {
        case 291:
            if len(v92) >= 2 {
                v92 = v92[1:len(v92)]
            }
        case 127:
            v93 := GenS8{F0: v1, F1: v1, F2: -653}
            _ = v93
        default:
            if (v91 <= 3.14159) {
                for fi94 := 0; fi94 < 4; fi94++ {
                    v14 -= v13
                    v95 := GenS9{F0: 598, F1: (v14 * (v13 % (((v13 - 2)) | 1)))}
                    _ = fi94
                    _ = v95
                }
                var v96 int8 = v1
                fmt.Println(v13)
                _ = v96
            } else {
                v97 := GenS10{F0: (^(v3)), F1: 127}
                switch "hello" {
                case "":
                    var v98 int8 = 23
                    _ = v98
                case "foo bar":
                    if ("hello" < "hello") {
                        var v99 int16 = 1000
                        _ = v99
                    } else {
                        v100 := GenS11{F0: 128, F1: (v1 & (v1 & 0))}
                        switch 368 {
                        case -1:
                            for fi101 := 0; fi101 < 20; fi101++ {
                                for fi102 := 0; fi102 < 15; fi102++ {
                                    v103 := GenS12{F0: 42}
                                    var v104 int64 = 9223372036854775807
                                    _ = fi102
                                    _ = v103
                                    _ = v104
                                }
                                _ = fi101
                            }
                        case 32767:
                            var v105 int8 = (v1 - (v1 + 35))
                            _ = v105
                        }
                        v14 = (v13 ^ v13)
                        _ = v100
                    }
                default:
                    if len(v92) > 0 {
                        v92[0] = int(v3)
                    }
                }
                _ = v97
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    fmt.Println("slicelen v92:", len(v92), v92)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v91
    _ = v92
}
