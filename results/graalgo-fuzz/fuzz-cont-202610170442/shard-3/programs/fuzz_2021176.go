package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint16; F2 int8 }
type GenS5 struct { F0 int32; F1 uint; F2 string }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS8 struct { F0 uint16 }

func main() {
    var v1 float32 = 0.1
    var v2 int = 0
    v3 := []string{}
    var v4 int = v2
    var v5 float64 = 0.0
    if (!false) {
        var v6 uint32 = 0
        _ = v6
    } else {
        var v7 float32 = 1e10
        switch "hello" {
        case "hello":
            if (!((2 <= 4294967295) && true)) {
                var v8 float64 = v5
                v7 -= float32(v8)
                for fi9 := 0; fi9 < 15; fi9++ {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        var v11 uint8 = 1
                        var v12 bool = true
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    var v13 uint64 = 42
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    _ = fi9
                    _ = v13
                }
                _ = v8
            } else {
                v14 := GenS1{F0: -268}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi15 := 0; fi15 < 7; fi15++ {
                    v2 = v4
                    v16 := GenS1{F0: 42}
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
        case "foo bar":
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 int = (-(fi17))
                _ = fi17
                _ = v18
            }
        case "":
            var v19 uint16 = 32767
            _ = v19
        default:
            v20 := GenS2{F0: 42}
            _ = v20
        }
        _ = v7
    }
    v1 -= 3.14159
    v1 = 0.0
    var v21 int16 = -1
    if true {
        for fi22 := 0; fi22 < 7; fi22++ {
            v1 -= ((v1 + v1) + float32((-2.25 / float64(3.14159))))
            if true {
                v2 = fi22
                switch ("hello" + (("foo bar" + "") + ("\n\t" + "foo bar"))) {
                case "hello":
                    fmt.Println(v21, v5, v5)
                }
                for fi23 := 0; fi23 < 14; fi23++ {
                    switch ((("\n\t" + ("\n\t" + "foo bar")) + "foo bar") + ("" + "a")) {
                    case "hello":
                        fmt.Println(v1, fi22, fi22)
                    case "\n\t":
                        switch (fi23 | (v4 + (v4 - v4))) {
                        case 127:
                            var v24 int = fi23
                            _ = v24
                        }
                    case "":
                        v3 = append(v3, "foo bar")
                    default:
                        if false {
                            v5 = v5
                            v1 -= v1
                            v25 := GenS2{F0: 941}
                            _ = v25
                        } else {
                            if true {
                                fmt.Println(v4)
                            } else {
                                switch -722 {
                                case 42:
                                    v3 = append(v3, ("" + "foo bar"))
                                default:
                                    v2 -= (fi23 & 32768)
                                }
                            }
                            switch (v4 * fi23) {
                            case -1:
                                v1 = float32(-2.25)
                            case 127:
                                if (!(274 <= 956)) {
                                    if (!true) {
                                        v21 = (-(v21))
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        switch "hello" {
                                        case "foo bar":
                                            var v26 uint16 = 255
                                            _ = v26
                                        default:
                                            var v27 float64 = 3.14159
                                            _ = v27
                                        }
                                    }
                                    v4 = -9223372036854775808
                                    if (2147483647 >= 2147483647) {
                                        var v28 int8 = 1
                                        _ = v28
                                    }
                                    v3 = append(v3, "a")
                                    if (!true) {
                                        v1 = v1
                                    } else {
                                        fmt.Println(fi23, v1)
                                        v29 := GenS2{F0: -928}
                                        if (!false) {
                                            if true {
                                                var v30 float32 = 1.5
                                                if ((2 != 2) || false) {
                                                    switch v2 {
                                                    case 2147483647:
                                                        var v31 uint8 = 2
                                                        _ = v31
                                                    default:
                                                        v32 := GenS1{F0: 0}
                                                        _ = v32
                                                    }
                                                    v2 = v2
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if (!(!true)) {
                                                        var v33 int16 = 0
                                                        fmt.Println(v5)
                                                        _ = v33
                                                    } else {
                                                        v3 = append(v3, ((((("\n\t" + "") + "\n\t") + ("hello" + "")) + "hello") + ("\n\t" + "hello")))
                                                    }
                                                    v21 += (v21 * (v21 | v21))
                                                }
                                                _ = v30
                                            } else {
                                                fmt.Println(v4, fi23)
                                            }
                                            v5 = (float64(1e10) / float64(v1))
                                        } else {
                                            if (-609 == 1) {
                                                var v34 bool = ((v5 - ((((v5 + v5) / ((v5 / float64(v1)) + v5)) - (v5 / v5)) - 1.5)) >= v5)
                                                _ = v34
                                            }
                                            if false {
                                                var v35 int16 = (v21 >> ((1) & 63))
                                                v1 += v1
                                                switch "foo bar" {
                                                case "":
                                                    if true {
                                                        if true {
                                                            if (39 > 218) {
                                                                v4 -= (fi22 ^ (v4 | 32768))
                                                                v36 := GenS3{F0: 0}
                                                                v1 -= (((v1 * 3.14159) - (3.14159 * 1e10)) / v1)
                                                                _ = v36
                                                            } else {
                                                                var v37 int = fi23
                                                                var v38 uint64 = 2
                                                                var v39 uint16 = 0
                                                                if len(v3) > 0 {
                                                                    v3[0] = (("hello" + ("a" + "a")) + ("foo bar" + "a"))
                                                                }
                                                                _ = v37
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            var v40 float32 = v1
                                                            _ = v40
                                                        }
                                                        v41 := GenS4{F0: "", F1: 965, F2: 100}
                                                        switch -2147483648 {
                                                        case -2147483648:
                                                            v1 = ((1e10 - (1e10 / 1.0)) - 1e10)
                                                        case -337:
                                                            _ = len(v3)
                                                        case 4294967295:
                                                            var v42 uint16 = 255
                                                            _ = v42
                                                        default:
                                                            var v43 int16 = (^(v35))
                                                            _ = v43
                                                        }
                                                        _ = v41
                                                    }
                                                case "hello":
                                                    var v44 float32 = 0.0
                                                    _ = v44
                                                case "foo bar":
                                                    switch fi23 {
                                                    case 2147483647:
                                                        switch fi23 {
                                                        case -571:
                                                            fmt.Println(v35, v4)
                                                        default:
                                                            _ = len(v3)
                                                        }
                                                    case -532:
                                                        switch "" {
                                                        case "\n\t":
                                                            v45 := GenS3{F0: -1}
                                                            _ = v45
                                                        case "foo bar":
                                                            v46 := GenS4{F0: "foo bar", F1: 65535, F2: 0}
                                                            _ = v46
                                                        }
                                                    default:
                                                        if (!(-634 > 42)) {
                                                            v35 = v35
                                                        }
                                                    }
                                                }
                                                _ = v35
                                            } else {
                                                if (765 <= 42) {
                                                    v21 += (v21 & 172)
                                                    if false {
                                                        switch "" {
                                                        case "hello":
                                                            var v47 string = "\n\t"
                                                            if (128 < 128) {
                                                                switch v47 {
                                                                case "hello":
                                                                    var v48 string = ""
                                                                    _ = v48
                                                                default:
                                                                    var v49 uint64 = 42
                                                                    _ = v49
                                                                }
                                                            } else {
                                                                fmt.Println(v21, v4, fi23)
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                v1 = 1.5
                                                            }
                                                            switch (v2 - (v4 ^ (v4 * (v2 * (-(fi22)))))) {
                                                            case -291:
                                                                v1 = v1
                                                            case 0:
                                                                var v50 bool = (false && false)
                                                                _ = v50
                                                            case 2147483647:
                                                                var v51 int16 = v21
                                                                _ = v51
                                                            }
                                                            _ = v47
                                                        case "foo bar":
                                                            v52 := GenS5{F0: 42, F1: 1, F2: ""}
                                                            _ = v52
                                                        case "\n\t":
                                                            v53 := GenS2{F0: -161}
                                                            _ = v53
                                                        }
                                                        var v54 int16 = 1000
                                                        var v55 float32 = -2.25
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        var v56 float32 = -2.25
                                                        switch v4 {
                                                        case -996:
                                                            var v57 uint16 = 199
                                                            _ = v57
                                                        case -406:
                                                            v58 := GenS4{F0: "", F1: 0, F2: -81}
                                                            _ = v58
                                                        case 869:
                                                            v59 := GenS1{F0: 42}
                                                            _ = v59
                                                        }
                                                        var v60 uint64 = 4294967295
                                                        _ = v56
                                                        _ = v60
                                                    }
                                                } else {
                                                    switch fi23 {
                                                    case 4294967295:
                                                        var v61 float64 = ((v5 / float64(float32(-2.25))) + v5)
                                                        _ = v61
                                                    }
                                                    v3 = append(v3, "\n\t")
                                                }
                                            }
                                        }
                                        _ = v29
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                }
                            default:
                                if ((-(v21)) >= v21) {
                                    v1 = float32((0.0 + ((((v5 * float64(v1)) / v5) + ((float64(v1) / 1.5) * 3.14159)) - v5)))
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(v5, v5)
                                    case "":
                                        v62 := GenS1{F0: 42}
                                        _ = v62
                                    default:
                                        v1 -= (0.0 - (1.5 * (0.1 / float32((v5 + (v5 + float64(((3.14159 / v1) / (float32(v5) + v1)))))))))
                                    }
                                    if false {
                                        v2 = fi23
                                        v2 = fi23
                                    }
                                }
                            }
                            var v63 uint16 = 2
                            _ = v63
                        }
                    }
                    _ = fi23
                }
            } else {
                v5 = 0.1
                v64 := GenS5{F0: -668, F1: 42, F2: "\n\t"}
                v2 -= (fi22 << ((42) & 63))
                _ = v64
            }
            _ = fi22
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v21 += (v21 ^ v21)
    } else {
        v65 := GenS6{F0: -32768, F1: 678}
        _ = v65
    }
    v2 = v2
    switch (v4 - 32767) {
    case 1:
        if ((v21 ^ (v21 << ((1) & 63))) > v21) {
            v5 += ((((1e10 * float64((float32(((((v5 - (1e10 - (float64(float32(3.14159)) * v5))) * 1.5) / ((0.1 * v5) - v5)) / (v5 / float64(float32(float64(float32((v5 / 1.0)))))))) / v1))) - v5) + (1.5 * (v5 * 1.5))) + ((float64(((3.14159 - v1) * -2.25)) + ((v5 - 0.0) * ((v5 + (v5 * (3.14159 + v5))) / float64(float32(v5))))) / v5))
            var v66 uint8 = 255
            if false {
                var v67 bool = true
                if ((v2 >> ((18446744073709551615) & 63)) == (-(v2))) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if true {
                        for fi68 := 0; fi68 < 6; fi68++ {
                            var v69 uint32 = uint32(v66)
                            var v70 float64 = 0.1
                            var v71 int32 = 0
                            _ = fi68
                            _ = v69
                            _ = v70
                            _ = v71
                        }
                        if (false && (!false)) {
                            if len(v3) > 0 {
                                v3[0] = ""
                            }
                            var v72 int64 = 1
                            switch "foo bar" {
                            case "hello":
                                v73 := GenS6{F0: v21, F1: 0}
                                _ = v73
                            case "":
                                v74 := GenS6{F0: v21, F1: 2}
                                _ = v74
                            default:
                                var v75 uint16 = 284
                                _ = v75
                            }
                            _ = v72
                        } else {
                            fmt.Println(v66, v66, v5)
                            switch (v4 + v4) {
                            case 4294967295:
                                var v76 uint64 = 2
                                _ = v76
                            case -9223372036854775808:
                                fmt.Println(v4, v67, v4)
                            case 2147483647:
                                for fi77 := 0; fi77 < 20; fi77++ {
                                    var v78 int16 = (v21 >> ((0) & 63))
                                    _ = fi77
                                    _ = v78
                                }
                            }
                        }
                        if (!(((v2 << ((149) & 63)) <= int(v66)) || v67)) {
                            v21 = 42
                        }
                    }
                } else {
                    v79 := GenS1{F0: 1}
                    if v67 {
                        v2 = (v2 | -1)
                        fmt.Println(v67)
                        for fi80 := 0; fi80 < 12; fi80++ {
                            _ = len(v3)
                            if ("\n\t" > "\n\t") {
                                v1 = (v1 / v1)
                                v5 = (float64(0.1) + 3.14159)
                                if (v66 == (v66 % ((v66) | 1))) {
                                    var v81 int16 = 42
                                    if v67 {
                                        fmt.Println(v66, v66, v2)
                                    }
                                    v21 = v81
                                    _ = v81
                                }
                            } else {
                                v21 = 957
                                v82 := GenS3{F0: -9223372036854775808}
                                v67 = v67
                                _ = v82
                            }
                            if v67 {
                                for fi83 := 0; fi83 < 18; fi83++ {
                                    var v84 uint8 = v66
                                    _ = fi83
                                    _ = v84
                                }
                                var v85 uint = 1
                                fmt.Println(v2)
                                _ = v85
                            } else {
                                if false {
                                    fmt.Println(v66)
                                }
                                var v86 string = "\n\t"
                                _ = v86
                            }
                            _ = fi80
                        }
                    }
                    if v67 {
                        v5 -= v5
                    }
                    _ = v79
                }
                _ = v67
            } else {
                v87 := GenS3{F0: 7}
                _ = v87
            }
            _ = v66
        } else {
            v3 = append(v3, "a")
            v1 = 1e10
            var v88 int8 = 52
            _ = v88
        }
    case 128:
        var v89 bool = (("a" + "\n\t") <= "hello")
        _ = v89
    }
    for fi90 := 0; fi90 < 6; fi90++ {
        if (!(-375 >= 1)) {
            switch "\n\t" {
            case "":
                var v91 uint32 = 2147483648
                _ = v91
            case "\n\t":
                v92 := GenS3{F0: 9223372036854775807}
                _ = v92
            case "foo bar":
                fmt.Println(v1)
            }
            if ((^(v21)) < (-(v21))) {
                var v93 int32 = -1
                var v94 bool = true
                var v95 float64 = 0.1
                _ = v93
                _ = v94
                _ = v95
            }
            v4 = (fi90 % ((v4) | 1))
        } else {
            var v96 uint32 = 752
            _ = v96
        }
        _ = fi90
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "a" {
    case "hello":
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    v97 := GenS7{F0: 1, F1: 2147483647, F2: 0}
    v98 := GenS8{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v97
    _ = v98
}
