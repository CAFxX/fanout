package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 bool; F1 bool; F2 int64 }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 uint; F1 int; F2 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 string; F1 int16; F2 string }
type GenS13 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS14 struct { F0 bool; F1 uint16 }
type GenS15 struct { F0 int8; F1 string }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int64; F1 int16 }
type GenS19 struct { F0 uint; F1 int; F2 string }

func main() {
    var v1 uint64 = 515
    v2 := make(map[string]uint32)
    var v3 int16 = 962
    v4 := make(map[int]int)
    var v5 int = -120
    if (42 <= 1) {
        if (!((v5 - 2147483647) >= v5)) {
            switch v5 {
            case 32767:
                var v6 float32 = 0.0
                _ = v6
            case 1:
                var v7 int = (v5 - (v5 * v5))
                _ = v7
            default:
                var v8 string = "a"
                _ = v8
            }
            var v9 bool = true
            _ = v9
        } else {
            if (9223372036854775807 < -264) {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 int = (fi10 / ((fi10) | 1))
                    v1 = v1
                    _ = fi10
                    _ = v11
                }
                v12 := []int{(v5 / (((-(v5))) | 1))}
                var v13 string = "hello"
                _ = v12
                _ = v13
            }
            var v14 int = v5
            var v15 bool = false
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            v5 = (v5 + 125)
            if true {
                if ("hello" < "hello") {
                    v17 := GenS1{F0: 42}
                    switch (("\n\t" + "\n\t") + "foo bar") {
                    case "\n\t":
                        var v18 float64 = 1.5
                        _ = v18
                    default:
                        fmt.Println(v1, v3)
                    }
                    _ = v17
                }
            }
            _ = fi16
        }
        if (0 <= (-(v3))) {
            for fi19 := 0; fi19 < 17; fi19++ {
                if false {
                    v20 := GenS2{F0: 2147483647, F1: v1}
                    v21 := []int{fi19, v5, (v5 / ((fi19) | 1)), fi19}
                    for fi22 := 0; fi22 < 9; fi22++ {
                        v5 += (v5 + (^(v5)))
                        switch "foo bar" {
                        case "a":
                            v1 += v1
                        }
                        _ = fi22
                    }
                    _ = v20
                    _ = v21
                } else {
                    v1 = (^(v1))
                    for fi23 := 0; fi23 < 3; fi23++ {
                        v24 := GenS3{F0: 127}
                        var v25 int8 = -98
                        switch "" {
                        case "hello":
                            switch (fi23 + fi23) {
                            case -93:
                                v26 := []int{(fi23 - -499), 9223372036854775807, 423}
                                _ = v26
                            case -1:
                                v27 := []int{fi19, 128, fi19}
                                _ = v27
                            default:
                                if true {
                                    switch fi19 {
                                    case 1:
                                        v28 := GenS4{F0: -757, F1: 65535}
                                        _ = v28
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        _ = len(v2)
                                    case "foo bar":
                                        var v29 int = (fi23 | 798)
                                        _ = v29
                                    case "a":
                                        v30 := GenS3{F0: 42}
                                        _ = v30
                                    default:
                                        v5 += (-(fi23))
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        var v31 uint8 = 127
                                        _ = v31
                                    case "":
                                        var v32 int64 = -761
                                        _ = v32
                                    default:
                                        var v33 uint = 42
                                        _ = v33
                                    }
                                } else {
                                    v34 := GenS2{F0: -286, F1: v1}
                                    _ = v34
                                }
                                var v35 uint = 65535
                                _ = v35
                            }
                        case "a":
                            v36 := GenS5{F0: true, F1: ((true || true) || false), F2: -9223372036854775808}
                            _ = v36
                        default:
                            v5 -= (fi19 + fi23)
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
                switch (fi19 / (((v5 | v5)) | 1)) {
                case -2147483648:
                    fmt.Println(fi19, v3)
                case 4294967295:
                    var v37 int16 = (v3 % ((1000) | 1))
                    _ = v37
                default:
                    for fi38 := 0; fi38 < 2; fi38++ {
                        if (485 >= 65535) {
                            if false {
                                if true {
                                    fmt.Println(v1, v1)
                                    switch "\n\t" {
                                    case "\n\t":
                                        if (false || ((3.14159 + (3.14159 * 1.5)) >= 1.5)) {
                                            if true {
                                                v39 := GenS6{F0: 128, F1: 887}
                                                switch "" {
                                                case "foo bar":
                                                    var v40 int32 = 42
                                                    _ = v40
                                                case "":
                                                    delete(v4, v5)
                                                case "a":
                                                    v41 := []int{}
                                                    _ = v41
                                                default:
                                                    if (false && true) {
                                                        _ = len(v2)
                                                    } else {
                                                        v1 = 18446744073709551615
                                                        v1 -= 2
                                                        var v42 int8 = 100
                                                        _ = v42
                                                    }
                                                }
                                                fmt.Println(v1, v5)
                                                _ = v39
                                            }
                                            if (!(0 > 2)) {
                                                delete(v2, ("foo bar" + ((("hello" + "") + ((("foo bar" + "") + "hello") + (("" + "hello") + "foo bar"))) + "a")))
                                            } else {
                                                v5 += 4294967295
                                                if false {
                                                    v43 := GenS7{F0: 42}
                                                    _ = v43
                                                } else {
                                                    v44 := v2["hello"]
                                                    switch "a" {
                                                    case "foo bar":
                                                        if false {
                                                            var v45 int64 = 42
                                                            var v46 int16 = v3
                                                            v5 -= 92
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    }
                                                    v47 := GenS8{F0: 316}
                                                    _ = v44
                                                    _ = v47
                                                }
                                                v5 = 532
                                                switch (-(fi38)) {
                                                case 802:
                                                    v48 := GenS9{F0: 42, F1: 255}
                                                    _ = v48
                                                case 32767:
                                                    _ = len(v4)
                                                }
                                                if (true && false) {
                                                    v1 -= v1
                                                } else {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        fmt.Println(v1, fi38)
                                                        v49 := GenS10{F0: 2, F1: (v5 >> ((338) & 63)), F2: 255}
                                                        var v50 uint16 = 65535
                                                        _ = v49
                                                        _ = v50
                                                    case "":
                                                        v51 := GenS6{F0: 1, F1: -409}
                                                        _ = v51
                                                    case "\n\t":
                                                        switch v5 {
                                                        case -2147483648:
                                                            v3 -= v3
                                                        case -1:
                                                            var v52 uint32 = 2
                                                            v3 += (v3 << ((198) & 63))
                                                            if false {
                                                                v52 = v52
                                                                _ = len(v2)
                                                            } else {
                                                                if false {
                                                                    switch v5 {
                                                                    case 42:
                                                                        var v53 string = ("hello" + "a")
                                                                        _ = v53
                                                                    case 531:
                                                                        if (-1 < 1) {
                                                                            var v54 uint8 = 2
                                                                            switch -110 {
                                                                            case 32768:
                                                                                v1 = v1
                                                                            case 394:
                                                                                var v55 uint8 = v54
                                                                                fmt.Println(v52, fi38, v54)
                                                                                v56 := []int{int(v3), v5, 796, 9223372036854775807}
                                                                                _ = v55
                                                                                _ = v56
                                                                            default:
                                                                                v5 -= fi19
                                                                            }
                                                                            switch fi19 {
                                                                            case -233:
                                                                                v57 := []int{(fi19 & (fi19 * (fi38 << ((4294967295) & 63))))}
                                                                                _ = v57
                                                                            case 42:
                                                                                v1 += v1
                                                                            case -64:
                                                                                v54 = uint8(v1)
                                                                                v54 = v54
                                                                            }
                                                                            _ = v54
                                                                        }
                                                                    }
                                                                    v58 := GenS11{F0: false}
                                                                    _ = v58
                                                                }
                                                                v1 -= (v1 / (((v1 + (^(v1)))) | 1))
                                                                v59 := v4[fi38]
                                                                _ = v59
                                                            }
                                                            _ = v52
                                                        default:
                                                            var v60 int32 = 241
                                                            _ = v60
                                                        }
                                                    }
                                                    v61 := GenS6{F0: 143, F1: -1}
                                                    _ = v61
                                                }
                                            }
                                        }
                                        v62 := []int{int(v1)}
                                        v5 += (fi19 / ((fi19) | 1))
                                        _ = v62
                                    case "a":
                                        switch "" {
                                        case "a":
                                            if true {
                                                v63 := v4[(-(fi19))]
                                                if false {
                                                    v64 := []int{v5, -714, v63, v5}
                                                    v65 := GenS3{F0: 118}
                                                    _ = v64
                                                    _ = v65
                                                }
                                                _ = v63
                                            } else {
                                                if ("foo bar" > "a") {
                                                    var v66 string = "hello"
                                                    v5 = -466
                                                    v1 += (v1 & v1)
                                                    _ = v66
                                                }
                                                v3 = -70
                                                v3 = 0
                                            }
                                            v1 = 80
                                            var v67 float32 = float32(1.5)
                                            _ = v67
                                        case "hello":
                                            var v68 int16 = 32767
                                            _ = v68
                                        case "":
                                            fmt.Println(fi19, fi38, v1)
                                            switch "hello" {
                                            case "\n\t":
                                                v69 := []int{-1, fi19}
                                                _ = v69
                                            case "a":
                                                switch "hello" {
                                                case "foo bar":
                                                    v1 = 42
                                                }
                                            case "hello":
                                                v3 = (v3 ^ 42)
                                            }
                                            v4[(^(fi38))] = (fi38 - v5)
                                            v70 := GenS12{F0: "", F1: (^(v3)), F2: (("" + (("a" + "") + "foo bar")) + "hello")}
                                            _ = v70
                                        }
                                    }
                                    delete(v2, "foo bar")
                                }
                                v71 := GenS6{F0: 94, F1: -1}
                                if (!true) {
                                    v3 = v3
                                    v3 += v3
                                    fmt.Println(fi19, v3, v3)
                                }
                                _ = v71
                            }
                            fmt.Println(v1)
                        }
                        v72 := GenS11{F0: (0 <= 42)}
                        _ = fi38
                        _ = v72
                    }
                }
                _ = fi19
            }
            v73 := []int{(-(v5)), (-(v5)), v5}
            _ = v73
        }
    }
    v5 = v5
    v74 := GenS6{F0: 101, F1: -1}
    v75 := GenS3{F0: 128}
    _ = len(v2)
    for fi76 := 0; fi76 < 17; fi76++ {
        var v77 bool = true
        if (255 != 1) {
            switch "foo bar" {
            case "hello":
                if (90 <= -2147483648) {
                    switch v5 {
                    case -2147483648:
                        v3 += (v3 % ((v3) | 1))
                    case 127:
                        delete(v4, 32768)
                    case -1:
                        for fi78 := 0; fi78 < 4; fi78++ {
                            var v79 bool = ((4294967295 > 255) && false)
                            _ = fi78
                            _ = v79
                        }
                    default:
                        _ = len(v2)
                    }
                    v80 := GenS13{F0: 4294967295, F1: 42, F2: 42}
                    v81 := GenS14{F0: v77, F1: 42}
                    _ = v80
                    _ = v81
                }
            }
            v77 = false
            var v82 float32 = 1.5
            _ = v82
        }
        _ = fi76
        _ = v77
    }
    v4[(-(v5))] = v5
    switch (v5 % (((v5 >> ((18446744073709551615) & 63))) | 1)) {
    case 88:
        var v83 uint = 42
        _ = v83
    case 385:
        var v84 bool = (57 < 255)
        _ = v84
    case -9223372036854775808:
        for fi85 := 0; fi85 < 12; fi85++ {
            var v86 int32 = -1
            _ = fi85
            _ = v86
        }
    }
    v87 := GenS15{F0: -68, F1: "foo bar"}
    if true {
        v1 = v1
        for fi88 := 0; fi88 < 9; fi88++ {
            switch (("hello" + "") + "") {
            case "hello":
                v5 = fi88
            case "foo bar":
                if true {
                    switch ("" + ("\n\t" + ((("a" + "hello") + "\n\t") + ("a" + "a")))) {
                    case "a":
                        v89 := GenS4{F0: int16(fi88), F1: 785}
                        _ = v89
                    }
                    v90 := GenS16{F0: 255}
                    _ = v90
                } else {
                    fmt.Println(v1)
                }
            case "":
                v91 := []int{v5}
                _ = v91
            default:
                fmt.Println(v3)
            }
            var v92 int8 = -1
            _ = fi88
            _ = v92
        }
        v93 := GenS17{F0: (-(v3))}
        _ = v93
    } else {
        v3 = v3
        if false {
            v5 = 910
            var v94 int16 = v3
            _ = v94
        } else {
            delete(v2, "hello")
            v5 -= (v5 & v5)
            if false {
                v95 := []int{v5, (v5 / ((0) | 1)), -615, -241}
                var v96 uint = 234
                v97 := GenS7{F0: (v5 ^ 0)}
                _ = v95
                _ = v96
                _ = v97
            }
        }
    }
    v1 = (v1 % ((42) | 1))
    if false {
        var v98 float64 = ((float64(3.14159) * 0.0) - 1.5)
        _ = v98
    } else {
        v99 := GenS18{F0: -9223372036854775808, F1: v3}
        if (42 >= 32767) {
            v100 := GenS19{F0: 42, F1: v5, F2: "a"}
            _ = v100
        }
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
    _ = v75
    _ = v87
}
