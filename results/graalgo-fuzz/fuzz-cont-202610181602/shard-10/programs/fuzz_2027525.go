package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int; F2 int64 }
type GenS3 struct { F0 int; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16 }
func genFn0(p0 int16, p1 int8) int32 {
    var v7 int32 = 2147483647
    switch ((("\n\t" + "foo bar") + "") + ("hello" + "")) {
    case "\n\t":
        v7 = v7
    case "a":
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: p1}
            _ = fi8
            _ = v9
        }
    default:
        var v10 int16 = (p0 - -32768)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 uint8, p1 int64) int {
    v11 := GenS1{F0: int8(p1)}
    var v12 uint32 = 84
    v13 := GenS1{F0: 100}
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint32 = 4294967295
        _ = fi14
        _ = v15
    }
    var v16 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    return -1
}

func main() {
    var v1 uint16 = 65535
    var v2 bool = true
    var v3 bool = v2
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    var v6 string = ("" + "foo bar")
    if v2 {
        if (-434 == 1) {
            var v17 int8 = 42
            for fi18 := 0; fi18 < 12; fi18++ {
                if false {
                    var v19 int16 = 32767
                    _ = v19
                } else {
                    v20 := GenS1{F0: (^(v17))}
                    _ = v20
                }
                v21 := GenS2{F0: 0, F1: 4294967295, F2: -9223372036854775808}
                _ = fi18
                _ = v21
            }
            _ = v17
        }
        fmt.Println(v1, v1, v1)
        if v3 {
            var v22 uint8 = 2
            v3 = (!((((3.14159 * (0.1 * -2.25)) * 3.14159) + 3.14159) == 3.14159))
            _ = v22
        }
    }
    var v23 uint8 = 1
    if (574 == 309) {
        var v24 float32 = float32((0.0 - 3.14159))
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 uint64 = 42
        v2 = (!false)
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 uint8 = (v23 >> ((4294967295) & 63))
        if true {
            v3 = (!true)
            var v29 int64 = 42
            _ = v29
        }
        if ((0.1 <= 1.5) || (65535 < 2147483647)) {
            v2 = (!v2)
            var v30 uint = 252
            _ = v30
        } else {
            if v2 {
                v23 = v23
                switch (fi27 >> ((97) & 63)) {
                case 127:
                    var v31 bool = false
                    _ = v31
                case 42:
                    v32 := GenS3{F0: -875, F1: 879, F2: 1000}
                    _ = v32
                default:
                    var v33 int = (fi27 >> ((0) & 63))
                    _ = v33
                }
            } else {
                fmt.Println(v6, v2)
                if (genFn1(v23, -1) > fi27) {
                    v1 -= v1
                    v1 -= (v1 % (((^(v1))) | 1))
                    var v34 int8 = 0
                    _ = v34
                }
                v3 = v2
            }
        }
        _ = fi27
        _ = v28
    }
    if true {
        fmt.Println(v3)
        for fi35 := 0; fi35 < 16; fi35++ {
            v2 = true
            _ = fi35
        }
    } else {
        v6 = "\n\t"
        v3 = (2 == 999)
        switch 2147483647 {
        case -2147483648:
            if v2 {
                var v36 uint8 = (v23 ^ v23)
                _ = v36
            } else {
                var v37 int64 = -709
                var v38 string = v6
                for fi39 := 0; fi39 < 16; fi39++ {
                    if (true && (1e10 > 0.0)) {
                        var v40 uint8 = v23
                        v1 = 42
                        for fi41 := 0; fi41 < 7; fi41++ {
                            var v42 float32 = -2.25
                            _ = fi41
                            _ = v42
                        }
                        _ = v40
                    } else {
                        v43 := GenS4{F0: (v1 >> ((18446744073709551615) & 63)), F1: (v23 >> ((0) & 63)), F2: 32767}
                        _ = v43
                    }
                    if ("a" > ((v6 + "a") + ((v38 + v6) + v6))) {
                        v3 = v2
                        var v44 uint32 = 898
                        _ = v44
                    }
                    _ = fi39
                }
                _ = v37
                _ = v38
            }
        case 127:
            for fi45 := 0; fi45 < 19; fi45++ {
                var v46 bool = (42 < 93)
                _ = fi45
                _ = v46
            }
        case 138:
            switch 39 {
            case 127:
                var v47 int8 = 0
                _ = v47
            case -520:
                v2 = (int8(v1) < -1)
            default:
                delete(v4, 0)
            }
        }
    }
    var v48 int = genFn1(127, 803)
    var v49 int32 = 0
    var v50 int32 = (-(v49))
    for fi51 := 0; fi51 < 14; fi51++ {
        fmt.Println(v2, v2, v23)
        v2 = (!(-2.25 <= 0.1))
        _ = fi51
    }
    switch "" {
    case "hello":
        for fi52 := 0; fi52 < 12; fi52++ {
            if (563 >= 1) {
                var v53 bool = ((v48 | 9223372036854775807) != v48)
                for fi54 := 0; fi54 < 9; fi54++ {
                    if (1 >= 1) {
                        if (4294967295 > 0) {
                            if false {
                                v4[genFn1(1, -1)] = 478
                            }
                            var v55 string = ("a" + "foo bar")
                            _ = v55
                        } else {
                            if (!(1.5 <= 3.14159)) {
                                var v56 int8 = 1
                                var v57 uint = uint(v1)
                                var v58 float32 = 0.0
                                _ = v56
                                _ = v57
                                _ = v58
                            }
                            if v53 {
                                var v59 int16 = 1000
                                _ = v59
                            }
                            var v60 int32 = (v49 - (v50 % ((genFn0(-1, 39)) | 1)))
                            v50 = v50
                            _ = v60
                        }
                    }
                    _ = fi54
                }
                for fi61 := 0; fi61 < 14; fi61++ {
                    var v62 int = fi52
                    if v2 {
                        v62 -= fi52
                    }
                    switch "hello" {
                    case "":
                        var v63 uint = 1
                        _ = v63
                    case "a":
                        if ((fi61 % (((-(v62))) | 1)) < (fi61 ^ v62)) {
                            v64 := []int{genFn1(v23, 792), (fi61 >> ((291) & 63))}
                            _ = v64
                        } else {
                            delete(v5, int(v23))
                            var v65 uint16 = 502
                            if false {
                                v65 = 1
                                var v66 bool = true
                                var v67 uint64 = 806
                                _ = v66
                                _ = v67
                            }
                            _ = v65
                        }
                    case "hello":
                        var v68 bool = ((!((v48 <= fi52) && (true && v2))) && (1 >= 521))
                        _ = v68
                    default:
                        v49 = -1
                    }
                    var v69 float64 = 3.14159
                    if true {
                        var v70 int32 = 2147483647
                        if (v2 && ((-1 > -1) || ((0.0 * 1e10) != ((v69 * float64(0.1)) / float64(3.14159))))) {
                            if (!true) {
                                v53 = (0 >= 0)
                            } else {
                                if v53 {
                                    var v71 string = v6
                                    var v72 uint = 0
                                    if v53 {
                                        v71 = "foo bar"
                                    }
                                    _ = v71
                                    _ = v72
                                } else {
                                    fmt.Println(fi61, fi61)
                                    v53 = false
                                }
                                v73 := GenS5{F0: 2, F1: (fi52 ^ genFn1(v23, -1))}
                                _ = v73
                            }
                            fmt.Println(v23)
                            fmt.Println(fi61, v70)
                        } else {
                            switch fi61 {
                            case 9223372036854775807:
                                v53 = (!true)
                            }
                        }
                        delete(v4, (v48 * 32767))
                        _ = v70
                    }
                    v74 := v5[v62]
                    _ = fi61
                    _ = v62
                    _ = v69
                    _ = v74
                }
                _ = v53
            }
            _ = fi52
        }
    default:
        if (0.1 > ((float64(1.5) * (float64(-2.25) + 1e10)) * (-2.25 - 0.1))) {
            var v75 uint8 = (v23 - v23)
            _ = v75
        } else {
            v76 := GenS6{F0: (v1 * v1)}
            v77 := GenS7{F0: -9223372036854775808, F1: "foo bar"}
            _ = v76
            _ = v77
        }
    }
    for fi78 := 0; fi78 < 3; fi78++ {
        v6 = (v6 + ("a" + "hello"))
        v49 = v50
        switch (v6 + v6) {
        case "hello":
            for fi79 := 0; fi79 < 14; fi79++ {
                if false {
                    fmt.Println(v49, v1)
                    v80 := GenS8{F0: v23}
                    v81 := v4[genFn1((v23 % (((v23 + 51)) | 1)), -112)]
                    _ = v80
                    _ = v81
                }
                var v82 uint32 = uint32(v23)
                if v2 {
                    v48 += genFn1((v23 % ((v23) | 1)), 1)
                    var v83 uint16 = v1
                    switch (v6 + "\n\t") {
                    case "hello":
                        if v2 {
                            var v84 bool = (!((!(419 == 894)) && (64 >= -1)))
                            var v85 uint16 = v1
                            _ = v84
                            _ = v85
                        } else {
                            if false {
                                fmt.Println(v82, v50, v48)
                            } else {
                                var v86 int8 = 100
                                _ = v86
                            }
                            fmt.Println(v1, v2)
                        }
                    case "foo bar":
                        var v87 uint32 = (v82 << ((547) & 63))
                        _ = v87
                    }
                    v88 := GenS9{F0: 1000}
                    _ = v83
                    _ = v88
                }
                _ = fi79
                _ = v82
            }
        case "foo bar":
            var v89 int8 = -17
            _ = v89
        case "\n\t":
            v90 := GenS5{F0: 4294967295, F1: (fi78 << ((255) & 63))}
            _ = v90
        }
        _ = fi78
    }
    v1 = v1
    var v91 string = ""
    switch v91 {
    case "":
        var v92 bool = v2
        _ = v92
    case "foo bar":
        if (!true) {
            switch (-(v48)) {
            case 635:
                v93 := GenS6{F0: (v1 / ((65535) | 1))}
                _ = v93
            case 9223372036854775807:
                var v94 uint16 = (v1 << ((0) & 63))
                _ = v94
            case -2147483648:
                v2 = v2
            }
            var v95 uint32 = 1
            if (!v2) {
                for fi96 := 0; fi96 < 19; fi96++ {
                    switch (fi96 * genFn1(v23, -1)) {
                    case -881:
                        var v97 float64 = float64(1.5)
                        _ = v97
                    case -2147483648:
                        fmt.Println(v49, fi96, v1)
                    default:
                        for fi98 := 0; fi98 < 12; fi98++ {
                            v99 := []int{(fi98 << ((255) & 63)), genFn1(42, -9223372036854775808)}
                            var v100 int64 = 0
                            _ = fi98
                            _ = v100
                            _ = v99
                        }
                    }
                    for fi101 := 0; fi101 < 10; fi101++ {
                        v102 := GenS7{F0: 9223372036854775807, F1: "a"}
                        _ = fi101
                        _ = v102
                    }
                    _ = fi96
                }
                var v103 uint64 = 4294967295
                _ = v103
            }
            _ = v95
        }
    case "\n\t":
        v50 = -2147483648
    }
    if ((1e10 / (1.5 - ((1e10 + 3.14159) * float64(1e10)))) == float64(((3.14159 - (-2.25 + -2.25)) / 0.1))) {
        v4[(v48 >> ((255) & 63))] = 42
    }
    var v104 int32 = v50
    v104 = (v50 - (v49 * v50))
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v104
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v91
}
