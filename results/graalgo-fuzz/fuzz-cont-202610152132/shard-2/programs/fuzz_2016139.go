package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string; F1 uint64; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint32; F1 int16; F2 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS11 struct { F0 bool; F1 int32; F2 int32 }
type GenS12 struct { F0 int64; F1 int64 }
type GenS13 struct { F0 int8; F1 int16 }
func genFn0(p0 string, p1 uint32) uint {
    var v4 float64 = (3.14159 + 3.14159)
    var v5 uint32 = 544
    v5 -= (v5 - 0)
    if true {
        for fi6 := 0; fi6 < 19; fi6++ {
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 bool = false
                v4 = v4
                if (!false) {
                    p1 = 1
                    v9 := []int{fi6, -1}
                    _ = v9
                } else {
                    if ((2 > 1) || (!v8)) {
                        v10 := GenS1{F0: 0, F1: v5, F2: 0}
                        p0 = "\n\t"
                        var v11 int64 = 463
                        _ = v10
                        _ = v11
                    } else {
                        v12 := []int{(fi7 >> ((613) & 63)), fi6, (fi6 - 32767)}
                        _ = v12
                    }
                    var v13 float64 = ((v4 + v4) * v4)
                    var v14 uint8 = 1
                    _ = v13
                    _ = v14
                }
                switch p0 {
                case "hello":
                    var v15 uint16 = 1
                    _ = v15
                case "foo bar":
                    var v16 float32 = 1.5
                    v17 := []int{(fi6 - (fi6 ^ fi7)), 42, -9223372036854775808}
                    _ = v16
                    _ = v17
                case "a":
                    v18 := GenS1{F0: -119, F1: p1, F2: 255}
                    _ = v18
                default:
                    p0 = "a"
                }
                _ = fi7
                _ = v8
            }
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 string = "\n\t"
                v21 := []int{}
                switch fi19 {
                case -1:
                    if false {
                        v22 := GenS2{F0: "\n\t", F1: 0, F2: 42}
                        _ = v22
                    } else {
                        switch ((v20 + v20) + v20) {
                        case "\n\t":
                            fmt.Println(fi6, p1, v5)
                        default:
                            v23 := GenS2{F0: p0, F1: 4294967295, F2: 131}
                            _ = v23
                        }
                    }
                case 2147483647:
                    v5 = (p1 * (v5 + (^(p1))))
                case -2147483648:
                    v24 := GenS2{F0: v20, F1: 0, F2: 32767}
                    _ = v24
                }
                _ = fi19
                _ = v20
                _ = v21
            }
            v25 := GenS1{F0: -1, F1: 862, F2: 901}
            _ = fi6
            _ = v25
        }
    } else {
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 uint32 = v5
            v4 += 3.14159
            for fi28 := 0; fi28 < 14; fi28++ {
                var v29 int64 = 0
                v30 := GenS3{F0: -1}
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi26
            _ = v27
        }
        fmt.Println(v4, v5, p0)
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 255
}
func genFn1() string {
    v31 := []int{229, 2147483647, 4294967295}
    _ = len(v31)
    var v32 uint64 = 4294967295
    for fi33 := 0; fi33 < 19; fi33++ {
        fmt.Println(fi33, v32)
        _ = fi33
    }
    _ = v31
    _ = v32
    return ("hello" + "hello")
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (!false)
    var v3 uint64 = 2
    if (!v2) {
        for fi34 := 0; fi34 < 14; fi34++ {
            var v35 string = ""
            v36 := GenS1{F0: 127, F1: 4294967295, F2: 2}
            if ((v35 >= ("foo bar" + (v35 + v35))) && v2) {
                v35 = "\n\t"
                if false {
                    switch (^(fi34)) {
                    case -798:
                        v37 := []int{(fi34 + fi34), fi34, fi34, fi34}
                        _ = v37
                    case -561:
                        if true {
                            fmt.Println(v2)
                            switch (v35 + "a") {
                            case "":
                                v3 += (v3 + (v3 / ((v3) | 1)))
                            case "a":
                                var v38 uint32 = 65535
                                _ = v38
                            case "\n\t":
                                for fi39 := 0; fi39 < 8; fi39++ {
                                    var v40 uint32 = 728
                                    v41 := GenS4{F0: -183, F1: (v40 - (^(v40))), F2: -269}
                                    v1[v35] = v2
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                            default:
                                if (304 == 915) {
                                    v2 = (1 <= 13)
                                    v1[("hello" + v35)] = v2
                                    var v42 uint64 = (v3 ^ (v3 - v3))
                                    _ = v42
                                } else {
                                    v43 := GenS4{F0: -1, F1: 4294967295, F2: 42}
                                    _ = v43
                                }
                            }
                            var v44 uint32 = 2
                            _ = v44
                        }
                    case 626:
                        if (!v2) {
                            v3 -= 4294967295
                        } else {
                            for fi45 := 0; fi45 < 9; fi45++ {
                                v2 = true
                                _ = fi45
                            }
                        }
                    }
                } else {
                    switch (fi34 & 0) {
                    case 9223372036854775807:
                        var v46 float32 = 1.5
                        _ = v46
                    case 128:
                        switch (v35 + "a") {
                        case "":
                            var v47 float64 = (3.14159 / (3.14159 / 1.5))
                            _ = v47
                        case "foo bar":
                            fmt.Println(v3)
                        default:
                            v48 := []int{(fi34 ^ (fi34 % ((-313) | 1))), fi34}
                            _ = v48
                        }
                    }
                }
            }
            _ = fi34
            _ = v35
            _ = v36
        }
    } else {
        v49 := GenS1{F0: 89, F1: 42, F2: 32767}
        _ = v49
    }
    v50 := v1[""]
    v51 := GenS5{F0: uint32(v3), F1: 413, F2: (("foo bar" + "") + ("" + "hello"))}
    switch "" {
    case "a":
        v52 := GenS4{F0: 32767, F1: 4294967295, F2: 0}
        _ = v52
    case "":
        if (((0 > -1) || (((float32(3.14159) + 1e10) + 0.1) > 3.14159)) || v2) {
            if false {
                if (!v50) {
                    for fi53 := 0; fi53 < 7; fi53++ {
                        for fi54 := 0; fi54 < 13; fi54++ {
                            switch fi53 {
                            case 9223372036854775807:
                                var v55 int64 = 0
                                var v56 uint16 = 664
                                _ = v55
                                _ = v56
                            case 4294967295:
                                switch "" {
                                case "":
                                    v57 := []int{(fi53 + fi54)}
                                    _ = v57
                                case "a":
                                    var v58 uint16 = 65535
                                    v59 := GenS2{F0: (("\n\t" + "hello") + ""), F1: 42, F2: 0}
                                    _ = v58
                                    _ = v59
                                }
                            }
                            v50 = ((v2 || v50) || v50)
                            if (18446744073709551615 != 605) {
                                v50 = (9223372036854775807 >= 42)
                            } else {
                                fmt.Println(v3, v50)
                                switch ("\n\t" + ("foo bar" + "")) {
                                case "a":
                                    if ((!(!v50)) && (v3 <= (^(v3)))) {
                                        v60 := GenS6{F0: -1}
                                        var v61 int16 = -650
                                        fmt.Println(v61)
                                        _ = v60
                                        _ = v61
                                    } else {
                                        v62 := GenS7{F0: 0, F1: (979 != 992)}
                                        var v63 int = (-(fi54))
                                        v64 := []int{fi54}
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                    }
                                case "\n\t":
                                    fmt.Println(v2)
                                default:
                                    var v65 uint = 2
                                    _ = v65
                                }
                                if (v50 || false) {
                                    var v66 bool = false
                                    v67 := GenS8{F0: 0}
                                    _ = v66
                                    _ = v67
                                } else {
                                    var v68 uint8 = 119
                                    var v69 uint16 = 65535
                                    if (genFn0("a", 65535) != 1) {
                                        v70 := GenS8{F0: 229}
                                        _ = v70
                                    }
                                    _ = v68
                                    _ = v69
                                }
                            }
                            var v71 string = "\n\t"
                            v3 = 18446744073709551615
                            _ = fi54
                            _ = v71
                        }
                        _ = fi53
                    }
                    v2 = v50
                } else {
                    switch 490 {
                    case 715:
                        v72 := v1[((("hello" + "hello") + "hello") + ("a" + "a"))]
                        _ = v72
                    default:
                        if v2 {
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v2, v3)
                            case "a":
                                for fi73 := 0; fi73 < 4; fi73++ {
                                    var v74 uint64 = (v3 | v3)
                                    if true {
                                        var v75 float32 = 0.0
                                        v76 := GenS9{F0: 42, F1: 32767}
                                        _ = v75
                                        _ = v76
                                    } else {
                                        for fi77 := 0; fi77 < 1; fi77++ {
                                            var v78 int8 = 67
                                            var v79 uint16 = 255
                                            if false {
                                                v50 = v2
                                                fmt.Println(fi77, v2, v78)
                                            } else {
                                                v80 := GenS5{F0: 65535, F1: 32767, F2: "a"}
                                                if (!(18 >= 120)) {
                                                    if ((3.14159 < 1.5) || (!(!(-1 == -2147483648)))) {
                                                        var v81 int16 = 701
                                                        v3 -= (^(v3))
                                                        _ = v81
                                                    } else {
                                                        v79 = 42
                                                        v3 = 0
                                                        v1[("hello" + "foo bar")] = v50
                                                    }
                                                } else {
                                                    v1["\n\t"] = ((v3 % ((v74) | 1)) != 42)
                                                }
                                                _ = v80
                                            }
                                            _ = fi77
                                            _ = v78
                                            _ = v79
                                        }
                                        for fi82 := 0; fi82 < 11; fi82++ {
                                            if v2 {
                                                v83 := GenS8{F0: 2}
                                                if (uint8(v74) >= 128) {
                                                    v84 := []int{}
                                                    var v85 string = ((("\n\t" + "\n\t") + (("" + "hello") + "hello")) + ("hello" + ("a" + "")))
                                                    _ = v84
                                                    _ = v85
                                                }
                                                _ = v83
                                            }
                                            _ = fi82
                                        }
                                    }
                                    _ = fi73
                                    _ = v74
                                }
                            }
                        } else {
                            fmt.Println(v50, v2)
                        }
                    }
                    v50 = (!((((0.1 / 1e10) / (1e10 - ((((-2.25 + (((0.1 * 1.5) - 1.5) / (1e10 + 1e10))) - (3.14159 + (0.0 + (1.5 / (float64(-2.25) + 1e10))))) + 0.1) + 1e10))) - 1e10) > 0.0))
                }
                v1["a"] = (!(0 != (^(v3))))
            }
            fmt.Println(v2, v3)
        } else {
            fmt.Println(v50)
        }
    case "hello":
        for fi86 := 0; fi86 < 11; fi86++ {
            v3 += 2
            if v50 {
                fmt.Println(fi86)
                v87 := GenS10{F0: -2147483648, F1: 957, F2: 2147483648}
                _ = v87
            }
            v88 := GenS3{F0: 311}
            _ = fi86
            _ = v88
        }
    }
    v89 := GenS11{F0: false, F1: 2147483647, F2: 2147483647}
    var v90 float64 = (-2.25 * -2.25)
    if (1.5 > 1.5) {
        var v91 int8 = -128
        if (2147483647 >= -557) {
            var v92 uint32 = 4294967295
            if v2 {
                switch "a" {
                case "a":
                    v93 := GenS12{F0: 1, F1: 1}
                    _ = v93
                case "":
                    fmt.Println(v92, v50, v92)
                case "foo bar":
                    v2 = v50
                default:
                    var v94 string = "foo bar"
                    _ = v94
                }
            } else {
                if v50 {
                    var v95 uint64 = v3
                    if v50 {
                        var v96 float32 = 0.0
                        _ = v96
                    }
                    fmt.Println(v92, v50, v3)
                    _ = v95
                } else {
                    v1["hello"] = true
                }
            }
            fmt.Println(v3)
            _ = v92
        }
        v2 = v2
        _ = v91
    }
    var v97 int8 = 127
    v50 = (!(-987 >= 32767))
    v50 = (uint32(v3) == 251)
    for fi98 := 0; fi98 < 16; fi98++ {
        v99 := GenS3{F0: 0}
        var v100 uint64 = v3
        fmt.Println(v2)
        _ = fi98
        _ = v100
        _ = v99
    }
    if false {
        v101 := []int{663, int(v3), 9223372036854775807}
        fmt.Println(v90)
        _ = v101
    } else {
        v97 += -128
        switch 2147483647 {
        case 2147483647:
            v102 := GenS13{F0: v97, F1: 42}
            _ = v102
        default:
            var v103 uint8 = 255
            _ = v103
        }
        if (543 >= -1) {
            var v104 uint16 = 32767
            var v105 int64 = -12
            var v106 string = "hello"
            _ = v104
            _ = v105
            _ = v106
        }
    }
    if ((0 >= 1) || (("a" + ("\n\t" + "\n\t")) >= "foo bar")) {
        if (842 != 176) {
            fmt.Println(v50, v2, v2)
            v107 := []int{}
            _ = v107
        } else {
            v90 -= 3.14159
            var v108 uint = 1
            var v109 float64 = (v90 * 0.1)
            _ = v108
            _ = v109
        }
        var v110 int32 = 42
        v50 = (v50 && (((v90 / (1e10 - v90)) / v90) < v90))
        _ = v110
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v90:", v90)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v89
    _ = v90
    _ = v97
}
