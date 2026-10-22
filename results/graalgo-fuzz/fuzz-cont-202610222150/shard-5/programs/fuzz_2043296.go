package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int16 }
type GenS2 struct { F0 string; F1 uint32; F2 string }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint16; F1 int8; F2 int32 }
type GenS5 struct { F0 int16; F1 uint16; F2 int8 }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int64; F1 int32 }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS11 struct { F0 int8; F1 int8; F2 uint32 }
type GenS12 struct { F0 uint; F1 bool }

func main() {
    var v1 int32 = -555
    v2 := []int{9223372036854775807, 4294967295, -435, -2147483648}
    var v3 uint16 = 65535
    var v4 bool = false
    v5 := make(map[string]int)
    var v6 int8 = 87
    var v7 uint16 = v3
    var v8 uint = 473
    var v9 int = -1
    var v10 uint64 = 4294967295
    if v4 {
        v11 := GenS1{F0: -612, F1: (v8 >> ((v8) & 63)), F2: -32768}
        _ = v11
    } else {
        fmt.Println(v3, v3)
        for fi12 := 0; fi12 < 6; fi12++ {
            if ((^(v10)) < uint64(v3)) {
                v13 := GenS1{F0: 957, F1: v8, F2: int16(fi12)}
                delete(v5, (("foo bar" + "\n\t") + ("" + "a")))
                if true {
                    fmt.Println(fi12)
                } else {
                    v3 = (^(v7))
                    var v14 int32 = -2147483648
                    _ = v14
                }
                _ = v13
            }
            _ = fi12
        }
    }
    v15 := GenS2{F0: "foo bar", F1: 42, F2: (("\n\t" + ("hello" + "a")) + "foo bar")}
    if (42 == 255) {
        fmt.Println(v1, v7, v9)
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 int = int(v1)
            fmt.Println(v4, fi16, v3)
            var v18 uint = v8
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v19 := GenS1{F0: -515, F1: 738, F2: 1000}
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := v5[("" + ("\n\t" + (("\n\t" + "") + ((("hello" + "a") + "foo bar") + "foo bar"))))]
        v5["a"] = v9
        _ = fi20
        _ = v21
    }
    if false {
        v4 = v4
        if v4 {
            switch v9 {
            case -9223372036854775808:
                v22 := v5["a"]
                _ = v22
            case -425:
                var v23 int8 = -128
                _ = v23
            case -73:
                if (-1 >= -1) {
                    fmt.Println(v1)
                    v6 = (v6 & v6)
                } else {
                    var v24 uint8 = 1
                    if (1000 >= 1) {
                        fmt.Println(v3, v10, v8)
                    } else {
                        v3 -= 42
                        v8 += (v8 << (((v8 * (v8 % ((v8) | 1)))) & 63))
                        v25 := GenS2{F0: "hello", F1: 2, F2: "hello"}
                        _ = v25
                    }
                    _ = v24
                }
            }
            v4 = v4
            v26 := GenS3{F0: v10, F1: v4}
            _ = v26
        } else {
            switch 1 {
            case -9223372036854775808:
                for fi27 := 0; fi27 < 5; fi27++ {
                    var v28 uint = 18446744073709551615
                    v29 := v5["\n\t"]
                    var v30 int64 = 1
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
            default:
                var v31 string = "\n\t"
                _ = v31
            }
            if (v4 && true) {
                v7 += (^(v3))
                fmt.Println(v8)
                if len(v2) > 0 {
                    v2[0] = (v9 * v9)
                }
            }
            v4 = ((v10 & v10) >= v10)
        }
        v32 := GenS4{F0: (v3 % (((^(v7))) | 1)), F1: v6, F2: v1}
        _ = v32
    }
    var v33 uint64 = 2
    v33 = (^(v10))
    v34 := GenS1{F0: 9223372036854775807, F1: uint(v10), F2: 1}
    if (!(!false)) {
        v3 = 32767
        var v35 float64 = (0.0 + 1e10)
        _ = v35
    } else {
        for fi36 := 0; fi36 < 2; fi36++ {
            if len(v2) > 0 {
                v2[0] = v9
            }
            _ = fi36
        }
        v37 := GenS2{F0: (("a" + "") + ("a" + ((("foo bar" + "a") + ("" + ("\n\t" + (("a" + ("\n\t" + "\n\t")) + "foo bar")))) + "\n\t"))), F1: 65535, F2: ""}
        for fi38 := 0; fi38 < 17; fi38++ {
            var v39 float32 = -2.25
            for fi40 := 0; fi40 < 16; fi40++ {
                var v41 uint = v8
                if v4 {
                    var v42 int16 = 32767
                    v43 := GenS2{F0: "", F1: 0, F2: (("hello" + "\n\t") + "hello")}
                    v44 := GenS1{F0: 329, F1: (v8 - (v41 & (v8 + (v41 & (v41 >> ((uint(v7)) & 63)))))), F2: -32768}
                    switch fi40 {
                    case 912:
                        var v45 uint32 = 65535
                        _ = v45
                    case -1:
                        if ((v1 - 42) <= (v1 >> ((v8) & 63))) {
                            v9 = 0
                            if v4 {
                                if ((v10 % ((uint64(v3)) | 1)) > v33) {
                                    if v4 {
                                        if ((v42 != v42) && ((v8 * v8) != v41)) {
                                            switch ("\n\t" + "hello") {
                                            case "foo bar":
                                                var v46 uint32 = 752
                                                _ = v46
                                            case "\n\t":
                                                v47 := GenS5{F0: v42, F1: v7, F2: -1}
                                                _ = v47
                                            case "":
                                                var v48 string = ""
                                                _ = v48
                                            }
                                            fmt.Println(v4, v4, fi38)
                                            var v49 float32 = v39
                                            _ = len(v2)
                                            var v50 uint32 = 65535
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                    v51 := GenS6{F0: (v9 << ((65535) & 63)), F1: v1}
                                    _ = v51
                                } else {
                                    if false {
                                        var v52 float32 = (0.1 / (((3.14159 / 3.14159) / (v39 - (v39 * v39))) + (v39 / float32(-2.25))))
                                        var v53 int8 = 97
                                        var v54 uint = uint(v7)
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    }
                                    switch "\n\t" {
                                    case "":
                                        switch "\n\t" {
                                        case "":
                                            v55 := GenS2{F0: ("a" + ""), F1: 4294967295, F2: (("\n\t" + "foo bar") + "foo bar")}
                                            _ = v55
                                        case "a":
                                            var v56 int8 = (v6 + v6)
                                            _ = v56
                                        }
                                    case "hello":
                                        var v57 string = "hello"
                                        _ = v57
                                    default:
                                        switch "\n\t" {
                                        case "":
                                            v58 := v5[(((("" + ((("hello" + ("" + "")) + "foo bar") + ("a" + ""))) + ((("" + "") + "") + "")) + ("" + "foo bar")) + "a")]
                                            _ = v58
                                        }
                                    }
                                    var v59 bool = v4
                                    _ = v59
                                }
                            }
                            if (v4 || v4) {
                                if v4 {
                                    switch fi38 {
                                    case 2147483647:
                                        v60 := v5["a"]
                                        _ = v60
                                    case -900:
                                        _ = len(v2)
                                        switch "" {
                                        case "":
                                            var v61 uint32 = 1
                                            _ = v61
                                        case "\n\t":
                                            v62 := v5["a"]
                                            _ = v62
                                        default:
                                            switch 657 {
                                            case -9223372036854775808:
                                                if (v4 || (v7 >= (^(v7)))) {
                                                    var v63 bool = v4
                                                    var v64 uint16 = v3
                                                    v9 -= fi40
                                                    _ = v63
                                                    _ = v64
                                                }
                                            case 72:
                                                fmt.Println(v4, v3, fi40)
                                            case -702:
                                                switch v9 {
                                                case -855:
                                                    if (v9 >= 128) {
                                                        v7 = 35
                                                        v33 = (v33 >> (((^(v41))) & 63))
                                                        var v65 uint8 = 128
                                                        _ = v65
                                                    } else {
                                                        v8 += v41
                                                        var v66 uint64 = (v10 % (((v10 * 71)) | 1))
                                                        fmt.Println(fi40, v42, fi40)
                                                        _ = v66
                                                    }
                                                case -1:
                                                    v2 = append(v2, (^(fi40)))
                                                }
                                            default:
                                                v67 := GenS7{F0: 1}
                                                _ = v67
                                            }
                                        }
                                    }
                                    var v68 uint16 = 2
                                    _ = v68
                                } else {
                                    v69 := GenS7{F0: 65535}
                                    _ = v69
                                }
                                v42 -= (v42 | (v42 & v42))
                            }
                        }
                    }
                    v70 := GenS7{F0: uint32(v10)}
                    _ = v42
                    _ = v43
                    _ = v44
                    _ = v70
                } else {
                    v71 := GenS3{F0: 18446744073709551615, F1: ((fi40 + fi40) >= (fi38 | 2147483647))}
                    var v72 int = -1
                    _ = v71
                    _ = v72
                }
                if (4294967295 >= (fi40 << ((v8) & 63))) {
                    v5["\n\t"] = v9
                    v39 += v39
                    v73 := GenS8{F0: -32, F1: v1}
                    v74 := GenS2{F0: ((("\n\t" + "a") + "") + "a"), F1: 1, F2: "\n\t"}
                    if v4 {
                        switch (fi38 + (-(v9))) {
                        case 9223372036854775807:
                            var v75 uint64 = 2
                            _ = v75
                        case 128:
                            v39 = v39
                        default:
                            var v76 uint = v41
                            _ = v76
                        }
                        switch "hello" {
                        case "\n\t":
                            switch (("a" + "hello") + (("\n\t" + (("a" + "foo bar") + (("" + "hello") + ""))) + ("\n\t" + ("hello" + "hello")))) {
                            case "foo bar":
                                v77 := GenS9{F0: (^(v3)), F1: ((v4 && true) && (3.14159 <= 3.14159))}
                                _ = v77
                            case "a":
                                var v78 uint32 = 2
                                _ = v78
                            default:
                                var v79 string = "\n\t"
                                if (true && (!((1.5 + 1e10) < float64((v39 - -2.25))))) {
                                    var v80 int8 = v6
                                    var v81 bool = v4
                                    _ = v80
                                    _ = v81
                                } else {
                                    v82 := GenS8{F0: int64(v3), F1: int32(fi38)}
                                    delete(v5, (v79 + (v79 + (v79 + "foo bar"))))
                                    v41 -= (v41 >> ((v8) & 63))
                                    _ = v82
                                }
                                var v83 uint = 2
                                v10 -= v10
                                v7 += 675
                                _ = v79
                                _ = v83
                            }
                        }
                        fmt.Println(v6)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        v84 := GenS10{F0: v1, F1: (v3 % ((973) | 1)), F2: 1}
                        _ = v84
                    }
                    if (1.5 > v39) {
                        fmt.Println(v41)
                        switch ("a" + "foo bar") {
                        case "":
                            switch "a" {
                            case "":
                                if v4 {
                                    v2 = append(v2, (fi40 >> (((^(v41))) & 63)))
                                    v10 = uint64(v3)
                                } else {
                                    v3 -= v7
                                    v85 := GenS11{F0: 0, F1: (v6 & v6), F2: 2147483648}
                                    if ((^(v7)) == v7) {
                                        v4 = true
                                        var v86 int = fi40
                                        if v4 {
                                            v87 := GenS4{F0: 603, F1: int8(v86), F2: v1}
                                            _ = len(v2)
                                            _ = v87
                                        }
                                        _ = v86
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = len(v2)
                                    }
                                    _ = v85
                                }
                            case "hello":
                                var v88 int64 = 42
                                _ = v88
                            default:
                                if false {
                                    v4 = v4
                                    v89 := GenS2{F0: "foo bar", F1: 4294967295, F2: (("" + "a") + ("" + "hello"))}
                                    v8 += v8
                                    v41 += (v41 * (v8 & 1))
                                    v1 = (v1 - (v1 % ((v1) | 1)))
                                    _ = len(v2)
                                    _ = v89
                                }
                            }
                        case "foo bar":
                            v90 := GenS12{F0: v8, F1: ((v39 >= v39) || v4)}
                            _ = v90
                        case "hello":
                            var v91 string = "foo bar"
                            _ = v91
                        default:
                            var v92 bool = v4
                            _ = v92
                        }
                    }
                    delete(v5, "\n\t")
                    _ = v73
                    _ = v74
                }
                _ = fi40
                _ = v41
            }
            _ = fi38
            _ = v39
        }
        _ = v37
    }
    delete(v5, ("hello" + ""))
    switch (^(v9)) {
    case 9223372036854775807:
        var v93 uint64 = v33
        _ = v93
    default:
        var v94 int64 = 9223372036854775807
        _ = v94
    }
    if true {
        var v95 uint32 = 2147483648
        _ = v95
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
