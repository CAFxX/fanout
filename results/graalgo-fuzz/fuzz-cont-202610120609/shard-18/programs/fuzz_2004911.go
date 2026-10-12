package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int64; F1 uint32; F2 int8 }
type GenS6 struct { F0 uint64; F1 uint; F2 int16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 1
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 0}
        v6 := []int{fi4}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 1}
    switch -9223372036854775808 {
    case 42:
        v8 := GenS1{F0: 32767}
        _ = v8
    case 697:
        if (!false) {
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint64 = 241
                v2 += 144
                _ = fi9
                _ = v10
            }
        }
    }
    switch "" {
    case "hello":
        if false {
            switch -9223372036854775808 {
            case 9223372036854775807:
                for fi11 := 0; fi11 < 1; fi11++ {
                    v2 += 247
                    switch fi11 {
                    case 1:
                        var v12 float32 = 1e10
                        _ = v12
                    case 42:
                        v2 = v2
                    }
                    _ = fi11
                }
            }
            var v13 int8 = 127
            var v14 bool = (false || false)
            _ = v13
            _ = v14
        }
    case "\n\t":
        switch (((("a" + ("hello" + "")) + ("a" + "")) + "hello") + "\n\t") {
        case "":
            v2 = v2
        case "\n\t":
            v15 := GenS1{F0: 551}
            _ = v15
        case "a":
            var v16 uint = 0
            _ = v16
        }
    case "a":
        for fi17 := 0; fi17 < 10; fi17++ {
            switch 1 {
            case -900:
                var v18 int32 = 952
                _ = v18
            case 127:
                var v19 int64 = -826
                _ = v19
            case 32767:
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 uint32 = 1
                    if false {
                        delete(v1, "hello")
                        var v22 uint16 = 2
                        _ = v22
                    } else {
                        v2 -= v2
                        v21 += v21
                        v2 = v2
                        v21 = 101
                        v1[""] = false
                        var v23 int8 = 1
                        _ = v23
                    }
                    v24 := []int{fi17}
                    _ = fi20
                    _ = v21
                    _ = v24
                }
            default:
                v25 := GenS1{F0: 32767}
                _ = v25
            }
            var v26 float64 = 0.0
            if (1 < 937) {
                v27 := GenS2{F0: -424, F1: -1, F2: 88}
                _ = v27
            }
            _ = fi17
            _ = v26
        }
    default:
        switch -9223372036854775808 {
        case -690:
            if false {
                switch (("" + "") + "hello") {
                case "hello":
                    if false {
                        v2 = 0
                    } else {
                        v2 += v2
                        v28 := GenS2{F0: -847, F1: 72, F2: 2}
                        v29 := GenS2{F0: -771, F1: -140, F2: 65535}
                        _ = v28
                        _ = v29
                    }
                default:
                    v30 := GenS1{F0: 1}
                    _ = v30
                }
                var v31 int64 = 0
                _ = v31
            } else {
                v2 -= (v2 ^ (v2 | 255))
            }
        default:
            var v32 int8 = -86
            _ = v32
        }
    }
    v33 := GenS3{F0: 4294967295}
    switch "a" {
    case "\n\t":
        switch "" {
        case "hello":
            v34 := GenS3{F0: 441}
            _ = v34
        case "a":
            var v35 int8 = 20
            _ = v35
        default:
            var v36 uint = 720
            _ = v36
        }
    case "foo bar":
        for fi37 := 0; fi37 < 8; fi37++ {
            switch fi37 {
            case 32768:
                v2 = 2
            case 128:
                v38 := []int{692}
                _ = v38
            default:
                var v39 float64 = 1e10
                _ = v39
            }
            v2 = (v2 & v2)
            _ = fi37
        }
    default:
        for fi40 := 0; fi40 < 6; fi40++ {
            v2 = v2
            var v41 uint64 = 0
            v42 := GenS4{F0: 0, F1: 0}
            _ = fi40
            _ = v41
            _ = v42
        }
    }
    fmt.Println(v2)
    if (0.0 >= (-2.25 + ((1e10 * (3.14159 + 0.0)) - ((0.1 + 1e10) / -2.25)))) {
        v43 := GenS3{F0: 1}
        v44 := []int{32768}
        v45 := v1[""]
        _ = v43
        _ = v44
        _ = v45
    } else {
        var v46 float64 = 0.0
        switch -467 {
        case 127:
            v47 := []int{}
            _ = v47
        case 1:
            var v48 string = ("a" + "")
            _ = v48
        }
        if (!false) {
            fmt.Println(v46, v46)
            v49 := GenS5{F0: -418, F1: 0, F2: 111}
            v2 = v2
            _ = v49
        } else {
            switch "foo bar" {
            case "hello":
                v50 := []int{}
                _ = v50
            case "":
                fmt.Println(v46)
            case "foo bar":
                v46 += v46
            }
        }
        _ = v46
    }
    var v51 uint32 = 42
    v52 := []int{4294967295}
    v51 += (^(v51))
    for fi53 := 0; fi53 < 13; fi53++ {
        if (0.1 == 0.0) {
            v51 = 2147483648
        } else {
            var v54 uint16 = 553
            var v55 float32 = 1e10
            _ = v54
            _ = v55
        }
        if (78 > -128) {
            if ((0.1 - (float32((3.14159 - (3.14159 * 1e10))) / ((0.0 / (1.5 + 0.0)) - (-2.25 * 3.14159)))) == (0.1 + 3.14159)) {
                fmt.Println(fi53)
            }
            if (v2 == 146) {
                v51 += 936
            }
            delete(v1, "hello")
        } else {
            for fi56 := 0; fi56 < 7; fi56++ {
                v1["\n\t"] = true
                v2 += v2
                _ = fi56
            }
        }
        _ = fi53
    }
    v57 := v3[""]
    for fi58 := 0; fi58 < 9; fi58++ {
        if len(v52) > 0 {
            v52[0] = fi58
        }
        var v59 string = v57
        if false {
            for fi60 := 0; fi60 < 8; fi60++ {
                switch (v59 + "\n\t") {
                case "hello":
                    var v61 uint = 18446744073709551615
                    _ = v61
                case "":
                    v2 -= (^(v2))
                }
                fmt.Println(v59, v57)
                v51 = (v51 + 65535)
                _ = fi60
            }
            var v62 int32 = -1
            _ = v62
        } else {
            var v63 float64 = 1e10
            _ = v63
        }
        _ = fi58
        _ = v59
    }
    var v64 uint = 254
    for fi65 := 0; fi65 < 7; fi65++ {
        if ((uint(v51) != (v64 * v64)) && (4294967295 >= 0)) {
            v66 := GenS6{F0: 460, F1: (v64 | v64), F2: -32768}
            if (((-2.25 + (1e10 - 0.0)) + 0.1) != (((-2.25 + 1.5) * 0.0) - (0.1 / (1e10 + 3.14159)))) {
                fmt.Println(v51, v64)
            } else {
                v1[((v57 + "a") + v57)] = (v51 < v51)
                if true {
                    for fi67 := 0; fi67 < 7; fi67++ {
                        v57 = ("\n\t" + ((("a" + v57) + ((v57 + "") + "hello")) + ("" + "\n\t")))
                        var v68 bool = (!(!false))
                        if true {
                            v69 := v3[("\n\t" + "foo bar")]
                            _ = v69
                        }
                        var v70 int64 = int64(v2)
                        v70 += (-(v70))
                        _ = fi67
                        _ = v68
                        _ = v70
                    }
                    switch (v57 + "hello") {
                    case "hello":
                        v51 = (^(v51))
                    default:
                        v57 = (v57 + v57)
                    }
                    if len(v52) > 0 {
                        v52[0] = 2147483647
                    }
                } else {
                    v71 := GenS3{F0: uint(v2)}
                    v51 -= 524
                    if false {
                        switch fi65 {
                        case -2147483648:
                            v2 -= (v2 - 127)
                        default:
                            v57 = "a"
                        }
                        fmt.Println(v2, v51)
                        var v72 float32 = 3.14159
                        _ = v72
                    }
                    _ = v71
                }
                v57 = ""
            }
            _ = v66
        }
        var v73 float32 = float32(-2.25)
        if true {
            if true {
                v74 := v3[v57]
                var v75 int32 = 42
                _ = v74
                _ = v75
            }
        } else {
            if true {
                fmt.Println(v64, v2, v51)
                var v76 uint16 = 32767
                fmt.Println(fi65)
                _ = v76
            }
            v57 = "\n\t"
        }
        _ = fi65
        _ = v73
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v57:", v57)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v51
    _ = v52
    _ = v57
    _ = v64
    _ = v7
}
