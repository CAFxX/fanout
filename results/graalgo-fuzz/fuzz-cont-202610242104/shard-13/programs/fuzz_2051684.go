package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint16; F1 string; F2 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32; F1 uint; F2 int8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int; F1 int32; F2 int32 }
type GenS11 struct { F0 int8; F1 int64; F2 int16 }
func genFn0(p0 bool, p1 uint32, p2 bool) int64 {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 65535}
        var v7 bool = p0
        v8 := []int{}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    if (p2 && true) {
        fmt.Println(p1)
        var v9 int64 = -1
        v10 := GenS1{F0: 733}
        _ = v10
        _ = v9
    }
    var v11 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return v11
}
func genFn1(p0 int, p1 int) bool {
    var v12 uint8 = 255
    if false {
        v13 := []int{4294967295, p0, 748}
        var v14 int16 = -235
        v15 := GenS1{F0: 42}
        _ = v13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v12
    return true
}

func main() {
    var v1 int8 = -44
    v2 := make(map[int]int)
    var v3 float64 = 1.5
    var v4 string = (("\n\t" + ("a" + ("foo bar" + "a"))) + "foo bar")
    switch 9223372036854775807 {
    case 127:
        var v16 int = -711
        _ = v16
    case -2147483648:
        v4 = ""
    default:
        if (0 != 0) {
            for fi17 := 0; fi17 < 10; fi17++ {
                v4 = (v4 + ((("hello" + "") + v4) + (("hello" + v4) + v4)))
                _ = fi17
            }
            switch ("\n\t" + (("hello" + v4) + "a")) {
            case "":
                v18 := GenS2{F0: 4294967295}
                _ = v18
            case "a":
                v3 -= ((v3 + (v3 / -2.25)) / v3)
            case "hello":
                switch 127 {
                case -527:
                    for fi19 := 0; fi19 < 1; fi19++ {
                        if false {
                            fmt.Println(fi19, fi19)
                            v20 := GenS1{F0: 239}
                            _ = v20
                        }
                        _ = fi19
                    }
                default:
                    v21 := []int{76, 1}
                    _ = v21
                }
            default:
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 string = v4
                    _ = fi22
                    _ = v23
                }
            }
        }
    }
    if (genFn0(true, 4294967295, false) <= 1) {
        switch -1 {
        case -273:
            v4 = v4
        case 0:
            v24 := GenS2{F0: 316}
            _ = v24
        case 2147483647:
            for fi25 := 0; fi25 < 9; fi25++ {
                for fi26 := 0; fi26 < 8; fi26++ {
                    v1 -= v1
                    _ = fi26
                }
                _ = fi25
            }
        }
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        switch v4 {
        case "hello":
            var v28 float64 = float64(1e10)
            _ = v28
        case "a":
            v29 := GenS3{F0: true, F1: 2}
            _ = v29
        default:
            for fi30 := 0; fi30 < 1; fi30++ {
                v4 = "foo bar"
                var v31 string = v4
                if true {
                    fmt.Println(v31, v3)
                    var v32 uint8 = 2
                    if (int16(v32) != -818) {
                        v33 := GenS4{F0: (0 <= 957), F1: fi27}
                        _ = v33
                    } else {
                        v4 = v4
                        if (42 == -494) {
                            v34 := GenS1{F0: 255}
                            var v35 uint32 = 2147483647
                            _ = v34
                            _ = v35
                        }
                        var v36 int = (fi27 & (-(fi30)))
                        _ = v36
                    }
                    _ = v32
                }
                _ = fi30
                _ = v31
            }
        }
        var v37 uint32 = 2
        _ = fi27
        _ = v37
    }
    v38 := GenS4{F0: (((true && true) && (!true)) || (65535 <= 867)), F1: 128}
    fmt.Println(v3, v4)
    if (!true) {
        if (!(32767 != -2147483648)) {
            var v39 uint = 65535
            v40 := []int{4294967295, 2147483647, -2147483648}
            var v41 uint64 = 51
            _ = v39
            _ = v40
            _ = v41
        } else {
            v1 += v1
            v42 := GenS5{F0: 2, F1: v4, F2: 873}
            var v43 int64 = -9223372036854775808
            _ = v42
            _ = v43
        }
        var v44 int32 = 921
        v45 := GenS2{F0: 657}
        _ = v44
        _ = v45
    }
    v46 := GenS3{F0: false, F1: 0}
    for fi47 := 0; fi47 < 15; fi47++ {
        var v48 string = (v4 + v4)
        _ = fi47
        _ = v48
    }
    v2[-158] = 824
    v3 = v3
    if (128 > 0) {
        v49 := GenS6{F0: 42}
        for fi50 := 0; fi50 < 8; fi50++ {
            for fi51 := 0; fi51 < 12; fi51++ {
                if true {
                    v52 := GenS7{F0: (fi51 / ((fi51) | 1))}
                    _ = v52
                }
                fmt.Println(v4, v4)
                _ = fi51
            }
            if (false && (true && (4294967295 > 122))) {
                switch "" {
                case "\n\t":
                    for fi53 := 0; fi53 < 7; fi53++ {
                        if true {
                            var v54 string = v4
                            v55 := GenS8{F0: -861, F1: 18446744073709551615, F2: v1}
                            _ = v54
                            _ = v55
                        }
                        if false {
                            v56 := []int{}
                            if (!(!(((v3 >= v3) || (!(2147483647 <= 42))) && (874 > 0)))) {
                                if true {
                                    fmt.Println(v1, v1)
                                    fmt.Println(v1, fi53, v4)
                                    switch v4 {
                                    case "a":
                                        if false {
                                            var v57 uint64 = 0
                                            _ = len(v2)
                                            _ = v57
                                        } else {
                                            var v58 uint = 255
                                            var v59 int64 = 42
                                            _ = v58
                                            _ = v59
                                        }
                                    case "\n\t":
                                        var v60 uint = 18446744073709551615
                                        _ = v60
                                    default:
                                        v4 = "foo bar"
                                    }
                                }
                                var v61 uint16 = 0
                                var v62 int16 = -32768
                                _ = v61
                                _ = v62
                            } else {
                                if (false || true) {
                                    var v63 string = v4
                                    v1 = (v1 | v1)
                                    v64 := GenS7{F0: -2147483648}
                                    _ = v63
                                    _ = v64
                                } else {
                                    var v65 int = (fi53 ^ (-(fi50)))
                                    if false {
                                        fmt.Println(v3, v3, v4)
                                        if true {
                                            v65 = -530
                                            if ((-1 > 0) && true) {
                                                switch (v4 + "foo bar") {
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "a":
                                                        v66 := GenS6{F0: 759}
                                                        _ = v66
                                                    case "\n\t":
                                                        v3 += (v3 / -2.25)
                                                    case "":
                                                        fmt.Println(fi53, fi53)
                                                    default:
                                                        delete(v2, (^(fi53)))
                                                    }
                                                case "a":
                                                    v67 := GenS9{F0: 509}
                                                    _ = v67
                                                default:
                                                    var v68 uint8 = 4
                                                    _ = v68
                                                }
                                                v3 -= (v3 - ((((float64(((-2.25 - 1.5) + (float32((v3 - (1.5 + -2.25))) - (0.1 * 1.5)))) + ((0.1 + v3) - float64(1.5))) - v3) / (float64(0.0) - 3.14159)) / 1.5))
                                            } else {
                                                v65 -= fi53
                                            }
                                        } else {
                                            if true {
                                                v2[-9223372036854775808] = (v65 + 127)
                                                fmt.Println(fi53, fi53)
                                                v3 = v3
                                            } else {
                                                _ = len(v2)
                                                switch v4 {
                                                case "foo bar":
                                                    v3 = (0.0 / 1.5)
                                                case "\n\t":
                                                    var v69 int32 = int32(v1)
                                                    _ = v69
                                                case "hello":
                                                    switch ("foo bar" + (v4 + v4)) {
                                                    case "\n\t":
                                                        v3 = float64((3.14159 - (3.14159 * 0.0)))
                                                    case "hello":
                                                        v56 = append(v56, (^(fi50)))
                                                    }
                                                    if len(v56) >= 2 {
                                                        v56 = v56[1:len(v56)]
                                                    }
                                                    fmt.Println(fi53, fi53, v1)
                                                    v3 += (1e10 * (float64(float32(float64(0.1))) - (v3 - ((v3 / 3.14159) / (((1e10 / (v3 - ((v3 * v3) / 0.1))) - (3.14159 + v3)) - float64(3.14159))))))
                                                    v70 := GenS10{F0: fi53, F1: -1, F2: 42}
                                                    v71 := GenS1{F0: 18446744073709551615}
                                                    v65 -= fi53
                                                    _ = v70
                                                    _ = v71
                                                }
                                                v72 := GenS1{F0: 65535}
                                                _ = v72
                                            }
                                            var v73 string = ("hello" + v4)
                                            _ = v73
                                        }
                                        fmt.Println(v3, fi50, v1)
                                    }
                                    _ = v65
                                }
                                switch (fi50 << ((1) & 63)) {
                                case 180:
                                    if (v3 <= v3) {
                                        v3 = v3
                                        if false {
                                            if ((!false) && true) {
                                                v74 := GenS11{F0: v1, F1: 0, F2: 0}
                                                _ = v74
                                            }
                                            v1 += (v1 << ((42) & 63))
                                        }
                                    }
                                case 81:
                                    fmt.Println(v1, v1)
                                default:
                                    var v75 uint8 = 127
                                    _ = v75
                                }
                                v3 = (-2.25 * float64(((((-2.25 / -2.25) * -2.25) - (1e10 + (1.5 - (-2.25 + 3.14159)))) * 1e10)))
                                fmt.Println(fi50)
                                if true {
                                    fmt.Println(v3, v1, v3)
                                    if len(v56) > 0 {
                                        v56[0] = (fi53 << ((1) & 63))
                                    }
                                } else {
                                    switch 611 {
                                    case -429:
                                        if len(v56) > 0 {
                                            v56[0] = 32767
                                        }
                                    case 9223372036854775807:
                                        var v76 string = v4
                                        var v77 int64 = 42
                                        if (!true) {
                                            var v78 uint32 = 4294967295
                                            _ = v78
                                        }
                                        switch (fi50 % ((42) | 1)) {
                                        case 42:
                                            var v79 int16 = 32767
                                            _ = v79
                                        case 32767:
                                            switch fi50 {
                                            case 32767:
                                                var v80 int16 = 32767
                                                _ = v80
                                            case -2147483648:
                                                v81 := v2[fi50]
                                                _ = v81
                                            }
                                        }
                                        if false {
                                            v1 = (v1 - (^(v1)))
                                            v82 := GenS6{F0: 18446744073709551615}
                                            _ = v82
                                        }
                                        _ = v76
                                        _ = v77
                                    case 0:
                                        switch v4 {
                                        case "\n\t":
                                            if false {
                                                v83 := GenS10{F0: -2147483648, F1: 42, F2: 42}
                                                _ = v83
                                            }
                                        case "foo bar":
                                            var v84 uint64 = 1
                                            _ = v84
                                        case "hello":
                                            v2[fi50] = (fi50 | fi50)
                                        }
                                    default:
                                        switch fi53 {
                                        case -859:
                                            var v85 int = (fi53 | 2147483647)
                                            _ = v85
                                        case 127:
                                            fmt.Println(v4, fi50, fi50)
                                        }
                                    }
                                }
                                var v86 int16 = -32768
                                switch (v4 + v4) {
                                case "hello":
                                    _ = len(v56)
                                case "a":
                                    v87 := GenS2{F0: 956}
                                    _ = v87
                                default:
                                    fmt.Println(v4, v1, v4)
                                }
                                v3 -= v3
                                _ = v86
                            }
                            var v88 uint64 = 0
                            var v89 uint32 = 42
                            _ = v56
                            _ = v88
                            _ = v89
                        }
                        _ = fi53
                    }
                case "a":
                    var v90 int16 = 958
                    _ = v90
                case "":
                    var v91 float32 = 0.1
                    _ = v91
                default:
                    v92 := []int{fi50, fi50, fi50, fi50}
                    _ = v92
                }
                var v93 int32 = 42
                var v94 float32 = -2.25
                _ = v93
                _ = v94
            } else {
                if (fi50 < fi50) {
                    for fi95 := 0; fi95 < 16; fi95++ {
                        var v96 uint = 4294967295
                        var v97 uint8 = 255
                        _ = fi95
                        _ = v96
                        _ = v97
                    }
                    var v98 int16 = 42
                    var v99 float32 = (1.5 * (0.1 + (0.0 - (((1e10 + (1e10 + 3.14159)) - (0.0 + (0.0 + ((1.5 / (((((1.5 * (1e10 * (0.0 - (3.14159 + 0.1)))) * (3.14159 - 0.0)) * (-2.25 * 0.0)) + float32(v3)) * 1e10)) + 1e10)))) - (0.0 + -2.25)))))
                    _ = v98
                    _ = v99
                } else {
                    switch (v4 + v4) {
                    case "a":
                        fmt.Println(v3)
                    case "\n\t":
                        v100 := []int{fi50}
                        _ = v100
                    }
                }
            }
            switch fi50 {
            case -2147483648:
                v3 = 1e10
            case 32767:
                v4 = "\n\t"
            case 128:
                if false {
                    var v101 uint = 4294967295
                    if (!false) {
                        v1 = v1
                        var v102 int32 = -610
                        _ = v102
                    }
                    for fi103 := 0; fi103 < 15; fi103++ {
                        switch (v4 + v4) {
                        case "a":
                            v104 := []int{(fi50 >> (((^(v101))) & 63)), (^(fi103))}
                            _ = v104
                        default:
                            var v105 float32 = (1e10 - 1.5)
                            _ = v105
                        }
                        _ = fi103
                    }
                    _ = v101
                }
            }
            _ = fi50
        }
        var v106 string = v4
        _ = v106
        _ = v49
    } else {
        for fi107 := 0; fi107 < 13; fi107++ {
            v3 += v3
            _ = fi107
        }
        v4 = "hello"
    }
    var v108 string = ("a" + ("hello" + ("a" + (("\n\t" + v4) + "\n\t"))))
    var v109 uint = 2
    if (-32768 != 1) {
        v110 := GenS7{F0: -2147483648}
        switch (((v108 + v108) + v108) + (v108 + (v4 + (v4 + "\n\t")))) {
        case "foo bar":
            _ = len(v2)
        case "hello":
            fmt.Println(v109, v108)
        default:
            v1 += v1
        }
        _ = v110
    }
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v108
    _ = v109
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v46
}
