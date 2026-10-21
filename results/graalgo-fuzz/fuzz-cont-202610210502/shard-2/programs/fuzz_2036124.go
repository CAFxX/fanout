package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 int8; F2 bool }
func genFn0() int64 {
    var v7 int8 = 80
    var v8 int8 = 127
    v8 = v8
    var v9 uint8 = 25
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint = 0
        _ = fi10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    var v3 bool = true
    var v4 bool = v3
    var v5 bool = v3
    var v6 uint16 = 255
    v5 = (!((!((127 >= 1) && ("hello" < (("\n\t" + "foo bar") + "\n\t")))) && (-46 < -35)))
    fmt.Println(v3, v4, v3)
    v4 = (v4 && ((!v4) || (42 != 32767)))
    if (v4 && (!(0.1 == 3.14159))) {
        fmt.Println(v3, v3)
    }
    var v12 float32 = 1e10
    if v3 {
        fmt.Println(v4, v4)
        v13 := []int{}
        _ = v13
    }
    var v14 uint = 2
    v4 = true
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint32 = 2
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 bool = (v14 > v14)
            v5 = v5
            _ = fi17
            _ = v18
        }
        v19 := GenS1{F0: fi15}
        _ = fi15
        _ = v16
        _ = v19
    }
    if (0 == 1) {
        var v20 int8 = 10
        v5 = v3
        fmt.Println(v3)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 3; fi21++ {
            if v3 {
                v22 := GenS2{F0: 2147483647}
                switch (fi21 & (fi21 - -535)) {
                case 552:
                    v23 := GenS2{F0: 2147483648}
                    _ = v23
                case -1:
                    v24 := GenS3{F0: fi21}
                    _ = v24
                }
                switch (-(fi21)) {
                case 128:
                    v3 = (!((v14 >= uint(v6)) || false))
                default:
                    for fi25 := 0; fi25 < 4; fi25++ {
                        v26 := v1[(fi21 / ((fi25) | 1))]
                        switch ("foo bar" + "") {
                        case "":
                            delete(v1, 4294967295)
                        }
                        _ = fi25
                        _ = v26
                    }
                }
                _ = v22
            }
            v27 := GenS2{F0: 0}
            _ = fi21
            _ = v27
        }
    }
    if ("foo bar" == "hello") {
        for fi28 := 0; fi28 < 14; fi28++ {
            fmt.Println(v6, v14, v12)
            _ = fi28
        }
        if v4 {
            v12 = 3.14159
            _ = len(v1)
        } else {
            for fi29 := 0; fi29 < 12; fi29++ {
                var v30 uint32 = 2
                var v31 uint = (v14 & 65535)
                v32 := GenS2{F0: v30}
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                v4 = (!false)
                _ = fi33
            }
            var v34 int64 = genFn0()
            _ = v34
        }
    } else {
        var v35 uint32 = 417
        var v36 uint8 = 127
        _ = v35
        _ = v36
    }
    switch "" {
    case "a":
        if true {
            switch 0 {
            case 127:
                v4 = (v14 > v14)
            }
        }
    case "hello":
        if true {
            _ = len(v2)
        } else {
            v37 := []int{}
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        switch ("" + ("\n\t" + "hello")) {
        case "":
            switch (fi38 % ((fi38) | 1)) {
            case 128:
                var v39 uint32 = 215
                _ = v39
            case 42:
                fmt.Println(v4, v3, v12)
            default:
                v40 := []int{42, (fi38 << ((v14) & 63)), (fi38 + int(v14)), (fi38 ^ (fi38 / (((fi38 / (((-(fi38))) | 1))) | 1)))}
                _ = v40
            }
        }
        switch 32767 {
        case 32767:
            fmt.Println(v6)
        case -9223372036854775808:
            switch ("" + "\n\t") {
            case "a":
                v4 = (v5 && (!v4))
            default:
                switch fi38 {
                case -9223372036854775808:
                    v41 := GenS3{F0: 128}
                    _ = v41
                case 127:
                    var v42 int32 = -2147483648
                    _ = v42
                case 42:
                    fmt.Println(v14, fi38)
                default:
                    if false {
                        if false {
                            v5 = (((!((v3 && (-1 != 607)) && (("" < "\n\t") && false))) || (true || (-2.25 >= v12))) && (v14 <= v14))
                            switch "hello" {
                            case "":
                                switch ("hello" + "foo bar") {
                                case "\n\t":
                                    var v43 uint8 = 0
                                    _ = v43
                                case "":
                                    v6 -= v6
                                case "foo bar":
                                    var v44 uint64 = 1
                                    _ = v44
                                }
                            case "hello":
                                var v45 string = "\n\t"
                                _ = v45
                            case "a":
                                fmt.Println(v14, v6)
                            }
                            v14 -= (v14 << (((v14 - v14)) & 63))
                        } else {
                            if (!false) {
                                v46 := GenS3{F0: (fi38 << ((v14) & 63))}
                                switch 1 {
                                case -2147483648:
                                    v3 = (518 > 0)
                                default:
                                    var v47 uint8 = 161
                                    _ = v47
                                }
                                v3 = (343 < 1)
                                _ = v46
                            } else {
                                switch ("hello" + "hello") {
                                case "":
                                    var v48 int16 = 1
                                    _ = v48
                                case "hello":
                                    if (false && (-123 >= -382)) {
                                        var v49 uint64 = 2
                                        fmt.Println(v3, v3, v4)
                                        _ = v49
                                    } else {
                                        switch -9223372036854775808 {
                                        case 32767:
                                            if (0 < 1) {
                                                for fi50 := 0; fi50 < 16; fi50++ {
                                                    if true {
                                                        var v51 uint8 = 128
                                                        switch (fi50 % ((fi38) | 1)) {
                                                        case 0:
                                                            fmt.Println(v12, fi38)
                                                        case -946:
                                                            v12 += (v12 + 3.14159)
                                                        default:
                                                            fmt.Println(v14, fi50, v51)
                                                        }
                                                        v3 = v4
                                                        fmt.Println(v6, v51)
                                                        if (!false) {
                                                            switch "\n\t" {
                                                            case "":
                                                                if (-9223372036854775808 == 0) {
                                                                    var v52 float64 = 1.5
                                                                    if false {
                                                                        switch "hello" {
                                                                        case "":
                                                                            if (1 <= 93) {
                                                                                if ((!(v14 >= (^(v14)))) || (1 > 587)) {
                                                                                    v4 = (v51 >= (v51 / (((^(v51))) | 1)))
                                                                                }
                                                                                v53 := []int{}
                                                                                _ = v53
                                                                            }
                                                                        }
                                                                        var v54 uint = v14
                                                                        _ = v54
                                                                    } else {
                                                                        fmt.Println(fi50, v51, v4)
                                                                        v55 := []int{(fi50 % (((-(fi38))) | 1)), fi38, -808}
                                                                        if v3 {
                                                                            if true {
                                                                                v56 := GenS1{F0: fi50}
                                                                                _ = v56
                                                                            }
                                                                            v57 := GenS2{F0: 727}
                                                                            v58 := GenS4{F0: fi50, F1: ("foo bar" + "a")}
                                                                            v59 := GenS4{F0: 9223372036854775807, F1: ("hello" + "a")}
                                                                            v14 -= (v14 - 456)
                                                                            _ = v57
                                                                            _ = v58
                                                                            _ = v59
                                                                        }
                                                                        _ = v55
                                                                    }
                                                                    _ = v52
                                                                }
                                                            case "\n\t":
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    if (-651 < 1) {
                                                                        var v60 int16 = 32767
                                                                        _ = v60
                                                                    } else {
                                                                        var v61 int64 = 0
                                                                        var v62 string = "hello"
                                                                        _ = v61
                                                                        _ = v62
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                var v63 uint = 4294967295
                                                                _ = v63
                                                            default:
                                                                var v64 int16 = 1
                                                                _ = v64
                                                            }
                                                            if v3 {
                                                                var v65 uint64 = 0
                                                                fmt.Println(fi50, fi38, fi50)
                                                                v66 := GenS4{F0: 32767, F1: ""}
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                            v67 := []int{fi38}
                                                            _ = v67
                                                        }
                                                        v68 := []int{42, 0}
                                                        v14 = (v14 ^ (v14 << (((^(v14))) & 63)))
                                                        _ = v51
                                                        _ = v68
                                                    } else {
                                                        v69 := GenS2{F0: 65535}
                                                        v70 := GenS4{F0: fi38, F1: (("\n\t" + "") + "a")}
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    if (!false) {
                                                        v71 := GenS5{F0: -1}
                                                        v6 += uint16(v14)
                                                        var v72 float64 = 1.5
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        var v73 bool = (-2.25 > float64(v12))
                                                        switch (^(fi38)) {
                                                        case -595:
                                                            v4 = (-15 == -1)
                                                        case 32768:
                                                            var v74 int16 = 0
                                                            _ = v74
                                                        default:
                                                            v4 = (!v5)
                                                        }
                                                        var v75 bool = true
                                                        _ = v73
                                                        _ = v75
                                                    }
                                                    _ = fi50
                                                }
                                                v14 = (v14 * (^(v14)))
                                                for fi76 := 0; fi76 < 10; fi76++ {
                                                    var v77 uint64 = 42
                                                    _ = fi76
                                                    _ = v77
                                                }
                                            } else {
                                                v3 = (!(!v4))
                                            }
                                        }
                                        v14 = v14
                                        switch -1 {
                                        case 128:
                                            var v78 float32 = 1.5
                                            _ = v78
                                        case 42:
                                            fmt.Println(v4, v12, v12)
                                        default:
                                            v79 := []int{}
                                            _ = v79
                                        }
                                    }
                                case "foo bar":
                                    if v3 {
                                        var v80 string = ("\n\t" + "\n\t")
                                        _ = v80
                                    }
                                default:
                                    var v81 int64 = 42
                                    _ = v81
                                }
                                for fi82 := 0; fi82 < 14; fi82++ {
                                    v3 = v5
                                    if false {
                                        v14 += 1
                                    }
                                    _ = fi82
                                }
                                if v3 {
                                    v5 = (false || false)
                                    switch fi38 {
                                    case 9223372036854775807:
                                        var v83 float64 = 1.5
                                        _ = v83
                                    case 1:
                                        _ = len(v2)
                                    case -9223372036854775808:
                                        v84 := v2["foo bar"]
                                        _ = v84
                                    default:
                                        var v85 uint64 = 2
                                        _ = v85
                                    }
                                    v86 := GenS5{F0: 42}
                                    _ = v86
                                } else {
                                    v14 += (v14 >> ((v14) & 63))
                                    v87 := []int{(-(fi38)), (^(fi38)), 42}
                                    _ = v87
                                }
                            }
                            for fi88 := 0; fi88 < 5; fi88++ {
                                v89 := GenS5{F0: 2147483647}
                                _ = fi88
                                _ = v89
                            }
                            v90 := GenS6{F0: (v6 & v6)}
                            _ = v90
                        }
                        fmt.Println(v3, v3)
                        delete(v1, int(v6))
                    }
                }
            }
        case 128:
            if (!true) {
                var v91 float32 = -2.25
                _ = v91
            }
        default:
            var v92 int8 = 88
            _ = v92
        }
        for fi93 := 0; fi93 < 5; fi93++ {
            switch (("\n\t" + "foo bar") + ("\n\t" + ("a" + ("" + "foo bar")))) {
            case "hello":
                if (4294967295 >= 2147483648) {
                    if (1e10 != (((-2.25 + v12) + v12) * 0.0)) {
                        if (v12 > -2.25) {
                            var v94 uint16 = v6
                            _ = v94
                        }
                        v1[(^(fi93))] = fi93
                    }
                    var v95 uint8 = 217
                    _ = v95
                }
            default:
                var v96 float64 = 3.14159
                _ = v96
            }
            _ = fi93
        }
        _ = fi38
    }
    var v97 int8 = 42
    for fi98 := 0; fi98 < 1; fi98++ {
        v99 := v1[167]
        switch "\n\t" {
        case "\n\t":
            v100 := GenS7{F0: (!false), F1: (v97 - (v97 >> ((4294967295) & 63))), F2: (int32(v6) != 2147483647)}
            _ = v100
        case "hello":
            delete(v2, ("hello" + "hello"))
        }
        _ = fi98
        _ = v99
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v97
}
