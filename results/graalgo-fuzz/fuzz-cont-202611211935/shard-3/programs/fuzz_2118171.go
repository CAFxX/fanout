package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint32; F1 uint }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]int)
    var v3 bool = true
    var v4 uint32 = 65535
    v5 := make(map[uint32]bool)
    var v6 int32 = 42
    if true {
        v1 -= (float32(1.5) - v1)
    }
    switch 2147483647 {
    case -9223372036854775808:
        var v7 int = -9223372036854775808
        _ = v7
    }
    _ = len(v2)
    if (v3 || (!false)) {
        switch "a" {
        case "\n\t":
            var v8 uint8 = 128
            _ = v8
        case "":
            switch -2147483648 {
            case -506:
                fmt.Println(v4, v3)
            }
        case "foo bar":
            v9 := GenS1{F0: 127, F1: 42}
            _ = v9
        }
        if false {
            var v10 int64 = -9223372036854775808
            fmt.Println(v1, v10)
            v11 := []int{-2147483648, -3, -1, -307}
            _ = v10
            _ = v11
        } else {
            switch 32767 {
            case -749:
                v3 = v3
            case 815:
                v6 = 2147483647
            case 9223372036854775807:
                for fi12 := 0; fi12 < 10; fi12++ {
                    v6 = 0
                    if true {
                        for fi13 := 0; fi13 < 3; fi13++ {
                            v14 := GenS2{F0: ""}
                            if v3 {
                                v4 = (^(v4))
                                v15 := []int{}
                                var v16 int16 = 0
                                var v17 int8 = 120
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            var v18 uint = 18446744073709551615
                            _ = fi13
                            _ = v14
                            _ = v18
                        }
                    }
                    _ = fi12
                }
            default:
                v19 := GenS2{F0: "foo bar"}
                _ = v19
            }
            switch -1 {
            case 2147483647:
                v6 = (v6 | v6)
            }
            v6 = (v6 & (v6 % ((v6) | 1)))
        }
    } else {
        var v20 uint8 = 127
        switch -9223372036854775808 {
        case 32768:
            v21 := GenS2{F0: "a"}
            _ = v21
        case -2147483648:
            var v22 string = "hello"
            _ = v22
        case 1:
            fmt.Println(v1, v1, v20)
        default:
            v23 := GenS3{F0: 42}
            _ = v23
        }
        _ = v20
    }
    fmt.Println(v4, v6, v6)
    v24 := GenS4{F0: (v6 % (((-(v6))) | 1)), F1: -18}
    for fi25 := 0; fi25 < 12; fi25++ {
        for fi26 := 0; fi26 < 11; fi26++ {
            if v3 {
                v6 -= (v6 >> ((42) & 63))
                if true {
                    var v27 int8 = -1
                    if false {
                        var v28 uint32 = (v4 | 673)
                        var v29 uint16 = 490
                        var v30 uint8 = 1
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    _ = v27
                } else {
                    v3 = v3
                    switch fi26 {
                    case 9223372036854775807:
                        delete(v2, v4)
                    default:
                        if false {
                            v31 := GenS5{F0: v4, F1: 902}
                            _ = v31
                        } else {
                            fmt.Println(v3, fi26, v4)
                        }
                    }
                    fmt.Println(v3, v4, fi26)
                }
                _ = len(v5)
                v32 := GenS5{F0: v4, F1: 42}
                _ = v32
            } else {
                switch (("\n\t" + ("a" + "a")) + "\n\t") {
                case "a":
                    if v3 {
                        var v33 uint64 = 1
                        if v3 {
                            fmt.Println(v4, fi26, v3)
                            if (4294967295 <= 557) {
                                v3 = v3
                                _ = len(v2)
                                var v34 int = (fi26 ^ fi26)
                                _ = v34
                            } else {
                                if (v3 || (false && (353 >= 255))) {
                                    var v35 uint16 = 0
                                    v36 := GenS4{F0: -152, F1: 42}
                                    _ = v35
                                    _ = v36
                                }
                            }
                            fmt.Println(v6, v1, v1)
                        } else {
                            if false {
                                v37 := GenS3{F0: 445}
                                _ = v37
                            }
                            var v38 uint64 = 2
                            _ = v38
                        }
                        var v39 uint32 = v4
                        _ = v33
                        _ = v39
                    } else {
                        v40 := GenS4{F0: 0, F1: -846}
                        if ((^(v4)) >= 394) {
                            var v41 uint = 794
                            _ = v41
                        } else {
                            v42 := GenS1{F0: -128, F1: -578}
                            if ("a" <= ("" + "\n\t")) {
                                v6 = v6
                                v43 := GenS5{F0: (^(v4)), F1: 1}
                                _ = v43
                            } else {
                                if v3 {
                                    var v44 int64 = -1
                                    v4 = 2147483648
                                    _ = v44
                                }
                                v45 := v5[v4]
                                v46 := GenS2{F0: "\n\t"}
                                _ = v45
                                _ = v46
                            }
                            _ = v42
                        }
                        _ = v40
                    }
                }
                v3 = false
                v5[v4] = (33 >= 246)
                var v47 bool = ((!true) && false)
                _ = v47
            }
            v6 = (v6 * (v6 / (((v6 + 289)) | 1)))
            _ = fi26
        }
        if ((18446744073709551615 < 18446744073709551615) || (fi25 > 515)) {
            switch 4294967295 {
            case 706:
                var v48 int8 = -1
                _ = v48
            case 127:
                if (1 > 42) {
                    if (("a" != ((("foo bar" + "\n\t") + "") + ("hello" + ("" + (("" + "foo bar") + "foo bar"))))) && (898 > 1)) {
                        fmt.Println(v1, v1)
                    } else {
                        v49 := GenS4{F0: (v6 << ((18446744073709551615) & 63)), F1: 42}
                        switch (-(fi25)) {
                        case 127:
                            v50 := []int{int(v6), (fi25 % (((fi25 | 952)) | 1)), (fi25 + fi25)}
                            _ = v50
                        case 128:
                            for fi51 := 0; fi51 < 15; fi51++ {
                                v1 -= (v1 / (v1 / 1.0))
                                switch "hello" {
                                case "\n\t":
                                    v6 += v6
                                default:
                                    if v3 {
                                        if ((v6 / ((int32(v4)) | 1)) == v6) {
                                            var v52 uint16 = 255
                                            if v3 {
                                                var v53 uint64 = 4294967295
                                                v54 := v5[(v4 << ((255) & 63))]
                                                _ = v53
                                                _ = v54
                                            }
                                            _ = v52
                                        }
                                        v55 := v5[(v4 * v4)]
                                        _ = v55
                                    }
                                }
                                _ = fi51
                            }
                        case 32768:
                            switch ((("a" + (("hello" + "\n\t") + (("" + "") + (("\n\t" + ("" + "foo bar")) + "\n\t")))) + "a") + "hello") {
                            case "a":
                                fmt.Println(fi25, v3, v3)
                            case "\n\t":
                                v56 := []int{fi25, (fi25 * -1), fi25, fi25}
                                _ = v56
                            default:
                                var v57 int16 = 1000
                                _ = v57
                            }
                        }
                        v3 = (v4 < (v4 ^ (v4 & (^(v4)))))
                        _ = v49
                    }
                    v58 := GenS3{F0: -1}
                    for fi59 := 0; fi59 < 1; fi59++ {
                        if (v3 && (false && v3)) {
                            var v60 float64 = float64(3.14159)
                            _ = v60
                        }
                        var v61 uint32 = v4
                        _ = fi59
                        _ = v61
                    }
                    _ = v58
                } else {
                    v4 = 2147483647
                    if true {
                        var v62 int32 = 1
                        if (1 > 18446744073709551615) {
                            var v63 uint16 = 42
                            fmt.Println(v4)
                            switch int(v62) {
                            case -707:
                                var v64 int32 = (v62 ^ (^(v62)))
                                _ = v64
                            case 127:
                                var v65 uint32 = 0
                                _ = v65
                            default:
                                var v66 uint = 42
                                _ = v66
                            }
                            _ = v63
                        } else {
                            if v3 {
                                for fi67 := 0; fi67 < 16; fi67++ {
                                    v68 := v2[(v4 % ((557) | 1))]
                                    _ = fi67
                                    _ = v68
                                }
                                if true {
                                    var v69 int16 = 1
                                    if (42 > 9223372036854775807) {
                                        v6 = 0
                                    } else {
                                        v69 -= v69
                                    }
                                    v1 += 1e10
                                    _ = v69
                                } else {
                                    for fi70 := 0; fi70 < 6; fi70++ {
                                        if v3 {
                                            var v71 uint64 = 2
                                            var v72 uint16 = 65535
                                            _ = v71
                                            _ = v72
                                        } else {
                                            v1 = v1
                                        }
                                        _ = fi70
                                    }
                                    switch "" {
                                    case "a":
                                        _ = len(v2)
                                    case "\n\t":
                                        if (0 > -343) {
                                            var v73 uint = 42
                                            switch fi25 {
                                            case -9223372036854775808:
                                                var v74 uint8 = 222
                                                _ = v74
                                            case 0:
                                                switch 127 {
                                                case 9223372036854775807:
                                                    fmt.Println(v6)
                                                }
                                            case 4294967295:
                                                v73 -= uint(v4)
                                            default:
                                                var v75 int = fi25
                                                _ = v75
                                            }
                                            v4 -= v4
                                            _ = v73
                                        }
                                    default:
                                        var v76 int = fi25
                                        _ = v76
                                    }
                                    if (false || v3) {
                                        var v77 int16 = 42
                                        v78 := v2[(v4 * 65535)]
                                        for fi79 := 0; fi79 < 3; fi79++ {
                                            v4 = v4
                                            _ = fi79
                                        }
                                        _ = v77
                                        _ = v78
                                    }
                                }
                            } else {
                                var v80 float64 = (0.1 + float64(float32((1.5 / (float64((float32(-2.25) - ((v1 / 0.1) * v1))) + ((3.14159 / 0.1) + 0.0))))))
                                _ = v80
                            }
                            for fi81 := 0; fi81 < 15; fi81++ {
                                switch ("\n\t" + "\n\t") {
                                case "hello":
                                    if (-128 >= 0) {
                                        var v82 uint16 = 255
                                        _ = v82
                                    }
                                case "a":
                                    v83 := GenS5{F0: v4, F1: 255}
                                    _ = v83
                                }
                                var v84 uint = 2
                                _ = fi81
                                _ = v84
                            }
                            fmt.Println(v4)
                        }
                        _ = v62
                    }
                    var v85 uint32 = v4
                    _ = v85
                }
            default:
                for fi86 := 0; fi86 < 8; fi86++ {
                    switch fi86 {
                    case 33:
                        v3 = v3
                    default:
                        switch "a" {
                        case "hello":
                            v4 -= 0
                        default:
                            v3 = true
                        }
                    }
                    _ = fi86
                }
            }
        } else {
            if false {
                for fi87 := 0; fi87 < 17; fi87++ {
                    fmt.Println(v3)
                    if v3 {
                        if ((false || v3) && v3) {
                            fmt.Println(v1, fi25, v3)
                            v4 -= v4
                            if v3 {
                                var v88 uint64 = 214
                                var v89 uint8 = 136
                                _ = v88
                                _ = v89
                            }
                        } else {
                            var v90 int = (fi25 & fi87)
                            var v91 int32 = (v6 ^ (v6 | (v6 % (((v6 & 42)) | 1))))
                            fmt.Println(v4, v6)
                            _ = v90
                            _ = v91
                        }
                    } else {
                        v92 := GenS2{F0: (((("foo bar" + ("\n\t" + "foo bar")) + "") + ("foo bar" + "a")) + "hello")}
                        var v93 int32 = -1
                        fmt.Println(fi87)
                        _ = v92
                        _ = v93
                    }
                    v4 = (v4 & (v4 >> ((2) & 63)))
                    _ = fi87
                }
                v94 := GenS5{F0: (^(v4)), F1: 255}
                var v95 uint32 = v4
                _ = v94
                _ = v95
            }
            v96 := []int{(fi25 - 2147483647)}
            var v97 float32 = float32(1e10)
            _ = v96
            _ = v97
        }
        v3 = (v3 || (9223372036854775807 > int64(v6)))
        _ = fi25
    }
    var v98 uint32 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
}
