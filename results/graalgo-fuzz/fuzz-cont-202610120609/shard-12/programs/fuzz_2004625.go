package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 bool; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 uint32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint8 }

func main() {
    var v1 uint32 = 592
    var v2 int = -1
    v3 := []string{(((("\n\t" + "hello") + "hello") + "a") + ("hello" + "a")), "a"}
    var v4 float32 = (-2.25 + 3.14159)
    var v5 int64 = 42
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            if len(v3) > 0 {
                v3[0] = (("foo bar" + "") + "hello")
            }
            var v9 float64 = 1.5
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if false {
        switch "" {
        case "a":
            if false {
                v1 -= v1
                v4 += float32((1e10 + 1e10))
            }
        case "\n\t":
            var v10 uint16 = 32767
            _ = v10
        case "":
            for fi11 := 0; fi11 < 3; fi11++ {
                fmt.Println(v4, v5)
                fmt.Println(v5, v1, v4)
                for fi12 := 0; fi12 < 4; fi12++ {
                    switch -1 {
                    case -2147483648:
                        var v13 int64 = v5
                        _ = v13
                    case 127:
                        fmt.Println(fi12, fi12, fi12)
                    default:
                        var v14 int16 = -258
                        _ = v14
                    }
                    if false {
                        if (((1e10 + float32(((3.14159 - float64((float32(0.1) / ((float32(float64(-2.25)) - (v4 * v4)) + ((float32(-2.25) * 1.5) + ((1.5 + (v4 / 1.0)) - (0.1 / 1.0))))))) * (float64(1.5) - float64(3.14159))))) >= v4) || (983 != -1)) {
                            if (true && false) {
                                v3 = append(v3, ("a" + "\n\t"))
                                v4 -= float32(float64(v4))
                                v2 = fi11
                                if true {
                                    v5 -= v5
                                    switch "hello" {
                                    case "foo bar":
                                        _ = len(v3)
                                    case "":
                                        if true {
                                            if (1000 >= 344) {
                                                v15 := GenS1{F0: "hello"}
                                                switch (fi11 & (v2 ^ 127)) {
                                                case 2147483647:
                                                    var v16 uint64 = 440
                                                    _ = v16
                                                case -428:
                                                    v17 := GenS2{F0: int16(v5), F1: false, F2: true}
                                                    _ = v17
                                                default:
                                                    fmt.Println(v1, v5, v2)
                                                    v4 -= v4
                                                }
                                                v5 = v5
                                                fmt.Println(fi11, v2)
                                                var v18 uint = 2
                                                if true {
                                                    if true {
                                                        v19 := v6["\n\t"]
                                                        v20 := GenS3{F0: 42}
                                                        _ = v19
                                                        _ = v20
                                                    }
                                                    fmt.Println(fi11, v2, fi12)
                                                } else {
                                                    var v21 uint64 = 2
                                                    var v22 float32 = (0.1 * ((float32((((((1e10 - ((-2.25 + (3.14159 * (0.1 / ((0.1 + 1.5) - 0.0)))) * float64(1.5))) + 1e10) / (1e10 / 3.14159)) / 1.5) + (1.5 * 1e10))) / v4) * v4))
                                                    v21 = (^(v21))
                                                    v23 := GenS3{F0: v5}
                                                    _ = v21
                                                    _ = v22
                                                    _ = v23
                                                }
                                                _ = v15
                                                _ = v18
                                            }
                                        }
                                        v2 -= v2
                                        switch v2 {
                                        case -9223372036854775808:
                                            fmt.Println(v2, v1)
                                        case -71:
                                            var v24 bool = (false && (v2 >= fi12))
                                            switch "\n\t" {
                                            case "hello":
                                                var v25 int16 = -642
                                                _ = v25
                                            case "":
                                                var v26 int8 = -128
                                                _ = v26
                                            case "\n\t":
                                                v24 = (((1e10 * 1.5) - 0.0) == ((1e10 * 1e10) / 1.5))
                                            default:
                                                switch (v2 * (-(fi12))) {
                                                case 128:
                                                    v5 -= (-(v5))
                                                case -9223372036854775808:
                                                    v27 := GenS4{F0: 2, F1: 102}
                                                    _ = v27
                                                case 775:
                                                    v24 = (!(!v24))
                                                default:
                                                    switch ("\n\t" + "") {
                                                    case "hello":
                                                        if ((0.1 + v4) < v4) {
                                                            v1 -= v1
                                                            var v28 string = "a"
                                                            _ = v28
                                                        }
                                                    case "a":
                                                        if (2 != 32767) {
                                                            v29 := GenS5{F0: 239}
                                                            _ = v29
                                                        } else {
                                                            v30 := GenS5{F0: 98}
                                                            v5 = v5
                                                            _ = v30
                                                        }
                                                    case "foo bar":
                                                        var v31 int8 = -1
                                                        _ = v31
                                                    default:
                                                        var v32 uint64 = uint64(v1)
                                                        _ = v32
                                                    }
                                                }
                                            }
                                            if (!v24) {
                                                var v33 bool = false
                                                v34 := GenS3{F0: -9223372036854775808}
                                                v6["a"] = (!v24)
                                                v35 := GenS1{F0: "foo bar"}
                                                var v36 uint32 = 0
                                                if len(v3) > 0 {
                                                    v3[0] = ""
                                                }
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                                _ = v36
                                            }
                                            _ = v24
                                        case 1:
                                            v37 := GenS6{F0: 100}
                                            _ = v37
                                        default:
                                            var v38 int = (-(fi12))
                                            _ = v38
                                        }
                                        var v39 uint8 = 127
                                        _ = v39
                                    }
                                    fmt.Println(fi12)
                                } else {
                                    v4 = v4
                                    switch "a" {
                                    case "\n\t":
                                        if (4294967295 < (v1 % (((v1 | v1)) | 1))) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            _ = len(v6)
                                        }
                                    case "":
                                        var v40 int8 = -1
                                        _ = v40
                                    default:
                                        v41 := GenS7{F0: ""}
                                        _ = v41
                                    }
                                    delete(v6, "hello")
                                }
                                v2 = (fi11 + (^(fi11)))
                            } else {
                                var v42 uint = 112
                                var v43 uint32 = 321
                                if (!false) {
                                    fmt.Println(v43)
                                    var v44 int32 = -2147483648
                                    _ = v44
                                }
                                _ = v42
                                _ = v43
                            }
                            if len(v3) > 0 {
                                v3[0] = (("hello" + "foo bar") + "")
                            }
                        }
                        v1 = 181
                        v5 += v5
                        var v45 bool = true
                        var v46 uint32 = 718
                        _ = v45
                        _ = v46
                    } else {
                        v2 = (fi12 >> ((2) & 63))
                        switch "\n\t" {
                        case "a":
                            if (!true) {
                                v4 -= ((3.14159 + v4) + v4)
                                var v47 uint64 = 2
                                v4 = (float32(-2.25) / v4)
                                _ = v47
                            }
                        case "hello":
                            v48 := GenS8{F0: 255}
                            _ = v48
                        case "foo bar":
                            if (228 >= 2) {
                                var v49 bool = false
                                switch "" {
                                case "hello":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    v5 = (-(v5))
                                }
                                _ = v49
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = len(v3)
                                _ = len(v6)
                            }
                            v3 = append(v3, "a")
                            v6["\n\t"] = false
                        default:
                            var v50 int = -9223372036854775808
                            _ = v50
                        }
                    }
                    _ = fi12
                }
                _ = fi11
            }
        }
        v51 := GenS3{F0: v5}
        _ = v51
    }
    var v52 int64 = (v5 >> ((0) & 63))
    fmt.Println(v1)
    for fi53 := 0; fi53 < 12; fi53++ {
        v52 -= (v5 + (-(v52)))
        v54 := GenS1{F0: "foo bar"}
        if (true || true) {
            var v55 uint32 = 2147483647
            v56 := GenS8{F0: 18446744073709551615}
            _ = v55
            _ = v56
        }
        _ = fi53
        _ = v54
    }
    for fi57 := 0; fi57 < 16; fi57++ {
        if len(v3) > 0 {
            v3[0] = ((("" + "\n\t") + ("" + "\n\t")) + ("" + "\n\t"))
        }
        var v58 uint32 = v1
        switch (v2 << ((2) & 63)) {
        case 9223372036854775807:
            var v59 float64 = float64(v4)
            _ = v59
        case 32768:
            var v60 uint32 = v1
            _ = v60
        default:
            var v61 int64 = -1
            _ = v61
        }
        _ = fi57
        _ = v58
    }
    if true {
        if false {
            var v62 uint16 = 65535
            _ = v62
        } else {
            for fi63 := 0; fi63 < 17; fi63++ {
                var v64 uint16 = 32767
                _ = fi63
                _ = v64
            }
            v65 := GenS4{F0: 255, F1: 79}
            _ = v65
        }
        v4 += ((1e10 - v4) / (0.1 + (v4 + float32(0.1))))
    } else {
        v4 = v4
    }
    v52 -= 189
    var v66 uint32 = 801
    for fi67 := 0; fi67 < 17; fi67++ {
        for fi68 := 0; fi68 < 3; fi68++ {
            fmt.Println(v1)
            _ = fi68
        }
        _ = fi67
    }
    for fi69 := 0; fi69 < 15; fi69++ {
        v2 = fi69
        _ = fi69
    }
    v70 := GenS3{F0: v52}
    if true {
        v71 := GenS5{F0: 127}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v71
    }
    if (100 > 1) {
        for fi72 := 0; fi72 < 2; fi72++ {
            if (true && false) {
                switch (v2 + (fi72 & v2)) {
                case 9223372036854775807:
                    var v73 int16 = -32768
                    _ = v73
                case -716:
                    if (!((1.5 / 1.0) < 1.5)) {
                        v74 := GenS9{F0: v5, F1: (v1 & v1)}
                        _ = v74
                    } else {
                        v4 += float32(0.1)
                        var v75 uint32 = (v66 ^ (v66 | (v1 ^ 2147483647)))
                        switch fi72 {
                        case 4294967295:
                            v76 := GenS10{F0: 1}
                            _ = v76
                        case 128:
                            for fi77 := 0; fi77 < 6; fi77++ {
                                delete(v6, "")
                                v78 := GenS11{F0: 128}
                                fmt.Println(v4)
                                _ = fi77
                                _ = v78
                            }
                        }
                        _ = v75
                    }
                case -44:
                    fmt.Println(v1, v5)
                default:
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                }
            } else {
                var v79 uint16 = 606
                _ = v79
            }
            for fi80 := 0; fi80 < 12; fi80++ {
                v52 = -374
                var v81 uint8 = 1
                if true {
                    if false {
                        fmt.Println(v5, fi72, v66)
                    }
                    v66 -= (v66 - v66)
                }
                v81 = v81
                _ = fi80
                _ = v81
            }
            _ = fi72
        }
        fmt.Println(v52, v4)
        for fi82 := 0; fi82 < 6; fi82++ {
            v1 += (v1 >> ((18446744073709551615) & 63))
            _ = fi82
        }
    } else {
        v5 = v52
        if ((1 == 383) && false) {
            v4 = (v4 + float32(1.5))
        }
        for fi83 := 0; fi83 < 8; fi83++ {
            v1 = (v66 << ((42) & 63))
            _ = fi83
        }
    }
    var v84 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v66:", v66)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v66
    _ = v70
    _ = v84
}
