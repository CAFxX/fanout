package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64; F1 uint; F2 string }
type GenS9 struct { F0 int8; F1 int }
type GenS10 struct { F0 int32; F1 int64 }
type GenS11 struct { F0 uint16; F1 int32; F2 int32 }
type GenS12 struct { F0 uint32; F1 int16 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 string; F1 int; F2 uint16 }
type GenS15 struct { F0 int32; F1 bool }
type GenS16 struct { F0 int; F1 uint8 }

func main() {
    var v1 uint64 = 826
    var v2 uint = 18446744073709551615
    v3 := make(map[string]bool)
    var v4 float64 = 1e10
    for fi5 := 0; fi5 < 18; fi5++ {
        if (127 != 1) {
            v6 := GenS1{F0: -1, F1: "a", F2: (^(v1))}
            var v7 string = "a"
            switch v7 {
            case "a":
                switch v7 {
                case "a":
                    fmt.Println(v2, v4)
                case "\n\t":
                    v2 += v2
                default:
                    v8 := GenS1{F0: int32(v1), F1: ((v7 + ("foo bar" + (v7 + ("hello" + v7)))) + "\n\t"), F2: 610}
                    _ = v8
                }
            case "foo bar":
                v9 := GenS2{F0: 0}
                _ = v9
            case "\n\t":
                if (true && true) {
                    var v10 float64 = 1.5
                    var v11 string = (("" + v7) + (v7 + "hello"))
                    v12 := GenS3{F0: ((v2 << (((v2 ^ 71)) & 63)) > 2), F1: 42}
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    v4 -= v4
                    v4 += float64((3.14159 / 3.14159))
                    for fi13 := 0; fi13 < 12; fi13++ {
                        if (!(-1 != -947)) {
                            if true {
                                var v14 uint16 = 42
                                v1 -= (v1 * v1)
                                _ = v14
                            }
                        } else {
                            if (!true) {
                                v15 := v3[v7]
                                var v16 bool = false
                                v2 += v2
                                _ = v15
                                _ = v16
                            }
                            v4 = float64(float32(v4))
                        }
                        if ((1.5 <= -2.25) || (true && true)) {
                            v17 := []int{fi5, fi5, (fi13 / (((fi5 / (((fi5 - (fi13 + fi13))) | 1))) | 1))}
                            switch "" {
                            case "hello":
                                if (!false) {
                                    fmt.Println(v4, fi13)
                                    var v18 uint32 = 65535
                                    var v19 bool = (505 > v2)
                                    _ = v18
                                    _ = v19
                                }
                            default:
                                if true {
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                } else {
                                    v20 := GenS3{F0: true, F1: -305}
                                    var v21 int64 = 503
                                    _ = v20
                                    _ = v21
                                }
                            }
                            switch (fi5 & (fi13 & (fi13 * fi5))) {
                            case -1:
                                if (fi5 >= (-(fi13))) {
                                    v22 := GenS3{F0: (!(1 == 104)), F1: -9223372036854775808}
                                    if (true || false) {
                                        if true {
                                            switch v7 {
                                            case "foo bar":
                                                var v23 int8 = 98
                                                _ = v23
                                            case "":
                                                if len(v17) >= 2 {
                                                    v17 = v17[1:len(v17)]
                                                }
                                            default:
                                                _ = len(v17)
                                            }
                                            var v24 int8 = 61
                                            if (3.14159 <= float32(v4)) {
                                                var v25 uint = v2
                                                var v26 int8 = v24
                                                _ = v25
                                                _ = v26
                                            } else {
                                                _ = len(v3)
                                                v1 -= v1
                                                if len(v17) >= 2 {
                                                    v17 = v17[1:len(v17)]
                                                }
                                            }
                                            _ = v24
                                        } else {
                                            var v27 uint32 = 1
                                            v28 := GenS4{F0: true, F1: v7}
                                            _ = v27
                                            _ = v28
                                        }
                                        fmt.Println(fi5, v1, v2)
                                    }
                                    v4 = ((-2.25 + 0.1) * 1e10)
                                    v29 := GenS5{F0: v2}
                                    _ = v22
                                    _ = v29
                                } else {
                                    v30 := GenS6{F0: v1}
                                    _ = len(v3)
                                    v4 -= v4
                                    var v31 int32 = 0
                                    var v32 uint16 = 18
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                            case 42:
                                v33 := GenS7{F0: -2147483648}
                                _ = v33
                            case 1:
                                v34 := GenS3{F0: true, F1: -394}
                                _ = v34
                            default:
                                v35 := GenS7{F0: 128}
                                _ = v35
                            }
                            _ = v17
                        }
                        if true {
                            v36 := GenS5{F0: uint(v1)}
                            v37 := GenS7{F0: (fi5 & 42)}
                            v4 = (v4 - (float64((3.14159 * 0.1)) + 3.14159))
                            _ = v36
                            _ = v37
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(fi13)
                            case "a":
                                fmt.Println(v1, v7, v1)
                            case "":
                                v38 := GenS8{F0: (v1 >> ((v2) & 63)), F1: (v2 % ((2) | 1)), F2: v7}
                                _ = v38
                            }
                            var v39 int8 = -1
                            var v40 float64 = (0.0 - 0.1)
                            _ = v39
                            _ = v40
                        }
                        var v41 uint8 = 128
                        v42 := GenS2{F0: 42}
                        _ = fi13
                        _ = v41
                        _ = v42
                    }
                }
            }
            _ = v6
            _ = v7
        } else {
            var v43 string = "hello"
            v44 := GenS9{F0: 42, F1: fi5}
            _ = v43
            _ = v44
        }
        switch "hello" {
        case "hello":
            if (false || false) {
                for fi45 := 0; fi45 < 12; fi45++ {
                    if true {
                        v2 = 1
                    }
                    v46 := GenS10{F0: 2147483647, F1: 1}
                    v47 := []int{fi5}
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                _ = len(v3)
            }
        case "a":
            v48 := GenS11{F0: 0, F1: 1, F2: 42}
            _ = v48
        }
        v49 := []int{}
        _ = fi5
        _ = v49
    }
    var v50 int16 = -1
    var v51 uint = v2
    v52 := v3[("hello" + "\n\t")]
    var v53 int64 = 9223372036854775807
    if false {
        for fi54 := 0; fi54 < 20; fi54++ {
            var v55 int8 = int8(v1)
            _ = fi54
            _ = v55
        }
        v56 := GenS8{F0: v1, F1: v51, F2: ("hello" + ("hello" + "hello"))}
        if ((!v52) || v52) {
            fmt.Println(v1, v50, v50)
        } else {
            v57 := GenS12{F0: 2147483647, F1: v50}
            _ = v57
        }
        _ = v56
    }
    switch "hello" {
    case "\n\t":
        for fi58 := 0; fi58 < 5; fi58++ {
            switch (fi58 * (fi58 / ((782) | 1))) {
            case 127:
                var v59 uint16 = 1
                _ = v59
            default:
                fmt.Println(v1, v1, v4)
            }
            v52 = (1.5 <= float64(((1.5 / 0.1) + 1e10)))
            v60 := []int{fi58, -2147483648}
            _ = fi58
            _ = v60
        }
    case "hello":
        switch 128 {
        case 32767:
            var v61 bool = false
            _ = v61
        case 4294967295:
            v62 := GenS9{F0: 100, F1: -1}
            _ = v62
        default:
            v63 := v3["hello"]
            _ = v63
        }
    default:
        for fi64 := 0; fi64 < 19; fi64++ {
            fmt.Println(v53, v1)
            v2 = v51
            v51 = (v2 & (^(v51)))
            _ = fi64
        }
    }
    switch "" {
    case "a":
        v52 = v52
    case "hello":
        var v65 uint8 = 207
        _ = v65
    }
    switch "hello" {
    case "\n\t":
        if (-470 != 42) {
            var v66 float64 = 0.1
            _ = v66
        } else {
            v50 += (v50 * 32767)
        }
    }
    if v52 {
        for fi67 := 0; fi67 < 10; fi67++ {
            v2 -= (v51 + v2)
            for fi68 := 0; fi68 < 2; fi68++ {
                switch 32768 {
                case 9223372036854775807:
                    fmt.Println(v51, fi67)
                }
                if false {
                    v69 := GenS13{F0: 0}
                    _ = v69
                }
                v70 := []int{fi68}
                v51 -= v2
                v51 -= (v51 | 4294967295)
                if len(v70) > 0 {
                    v70[0] = (fi68 | -2147483648)
                }
                if (255 != 485) {
                    if len(v70) > 0 {
                        v70[0] = (fi68 ^ int(v1))
                    }
                    if (true || v52) {
                        if v52 {
                            fmt.Println(v4, v51)
                            v2 = 782
                        } else {
                            var v71 string = (("hello" + (("foo bar" + "") + ((("\n\t" + "a") + ("\n\t" + ("" + "foo bar"))) + "hello"))) + "")
                            if len(v70) >= 2 {
                                v70 = v70[1:len(v70)]
                            }
                            v72 := GenS2{F0: -528}
                            _ = v71
                            _ = v72
                        }
                        v50 -= 1
                    } else {
                        if (!(!(2 <= 110))) {
                            var v73 int16 = v50
                            _ = len(v3)
                            v50 = (v50 ^ (v50 << ((v2) & 63)))
                            _ = v73
                        } else {
                            if false {
                                v2 = (v51 << ((42) & 63))
                                v74 := GenS14{F0: "hello", F1: fi67, F2: 469}
                                _ = v74
                            }
                            v75 := GenS7{F0: -2147483648}
                            _ = v75
                        }
                    }
                    v76 := GenS15{F0: 914, F1: false}
                    fmt.Println(v52, v53)
                    _ = v76
                }
                _ = fi68
                _ = v70
            }
            _ = fi67
        }
        for fi77 := 0; fi77 < 15; fi77++ {
            if v52 {
                if true {
                    fmt.Println(v50)
                    var v78 uint32 = 360
                    _ = v78
                } else {
                    v50 = (v50 ^ (^(v50)))
                    fmt.Println(v50, v2)
                    v50 += v50
                }
                var v79 uint64 = (v1 + v1)
                for fi80 := 0; fi80 < 8; fi80++ {
                    switch "\n\t" {
                    case "a":
                        switch ("hello" + (("foo bar" + "hello") + "hello")) {
                        case "\n\t":
                            switch ("foo bar" + ("" + ("a" + ((("\n\t" + ("a" + ("\n\t" + "\n\t"))) + "a") + ("hello" + "\n\t"))))) {
                            case "hello":
                                var v81 float64 = 3.14159
                                _ = v81
                            case "\n\t":
                                var v82 int32 = int32(v53)
                                _ = v82
                            case "foo bar":
                                v2 = 124
                            default:
                                var v83 string = "hello"
                                _ = v83
                            }
                        case "":
                            var v84 uint = v2
                            _ = v84
                        case "hello":
                            var v85 string = "hello"
                            _ = v85
                        default:
                            v86 := []int{}
                            _ = v86
                        }
                    case "foo bar":
                        fmt.Println(v79, v1)
                    case "\n\t":
                        v87 := GenS16{F0: (fi77 << (((v51 << ((v2) & 63))) & 63)), F1: 2}
                        _ = v87
                    }
                    var v88 uint = (v2 >> (((v2 & v2)) & 63))
                    fmt.Println(v53, v51, v50)
                    _ = fi80
                    _ = v88
                }
                _ = v79
            }
            v1 = 339
            switch "\n\t" {
            case "":
                fmt.Println(v52)
            case "hello":
                v1 += v1
            default:
                var v89 uint = v51
                _ = v89
            }
            _ = fi77
        }
    }
    for fi90 := 0; fi90 < 3; fi90++ {
        v1 = (^(v1))
        if (!true) {
            var v91 float32 = -2.25
            _ = v91
        } else {
            for fi92 := 0; fi92 < 16; fi92++ {
                if v52 {
                    v93 := []int{fi92, fi90, 2147483647}
                    _ = v93
                } else {
                    v53 = 9223372036854775807
                }
                v94 := GenS16{F0: fi90, F1: 3}
                _ = fi92
                _ = v94
            }
        }
        v52 = (!((!(-56 <= 42)) || (!(v52 || (1 >= 2147483647)))))
        _ = fi90
    }
    v95 := []int{2147483647, -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v95:", len(v95), v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v95
}
