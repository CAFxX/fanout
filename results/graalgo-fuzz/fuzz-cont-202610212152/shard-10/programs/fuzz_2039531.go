package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint; F1 uint; F2 uint8 }
type GenS9 struct { F0 uint; F1 int16 }
type GenS10 struct { F0 uint; F1 int64; F2 int }
type GenS11 struct { F0 uint32; F1 int64 }
type GenS12 struct { F0 int64; F1 int8 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 42
    var v3 bool = (-128 != -66)
    if (-9223372036854775808 >= 401) {
        if (v3 || (!v3)) {
            fmt.Println(v2, v3)
        } else {
            for fi4 := 0; fi4 < 7; fi4++ {
                if (!v3) {
                    v5 := GenS1{F0: (fi4 ^ fi4)}
                    _ = v5
                } else {
                    if false {
                        fmt.Println(v2, fi4)
                        _ = len(v1)
                    } else {
                        switch (("hello" + ("a" + "hello")) + ("\n\t" + ("foo bar" + "hello"))) {
                        case "\n\t":
                            if v3 {
                                var v6 int32 = 1
                                _ = v6
                            } else {
                                v7 := GenS1{F0: -2147483648}
                                var v8 uint8 = 2
                                if v3 {
                                    for fi9 := 0; fi9 < 2; fi9++ {
                                        switch ("" + ("hello" + ("\n\t" + "hello"))) {
                                        case "":
                                            if (127 >= fi4) {
                                                v8 -= v8
                                                _ = len(v1)
                                            } else {
                                                v2 = fi9
                                            }
                                        case "hello":
                                            v10 := GenS2{F0: true}
                                            _ = v10
                                        }
                                        v11 := []int{}
                                        v2 = (fi4 * fi9)
                                        _ = fi9
                                        _ = v11
                                    }
                                } else {
                                    v8 = (v8 * v8)
                                    v12 := GenS3{F0: 42}
                                    _ = v12
                                }
                                _ = v7
                                _ = v8
                            }
                        case "hello":
                            v2 -= v2
                        case "foo bar":
                            if (2 >= 1) {
                                v3 = (42 != -1)
                            } else {
                                if ("foo bar" == "\n\t") {
                                    v3 = v3
                                    for fi13 := 0; fi13 < 14; fi13++ {
                                        switch fi13 {
                                        case -2147483648:
                                            v2 = (v2 + -1)
                                        case -9223372036854775808:
                                            var v14 int8 = -128
                                            _ = v14
                                        case 32768:
                                            v15 := GenS4{F0: 2}
                                            _ = v15
                                        default:
                                            _ = len(v1)
                                        }
                                        v2 = 4294967295
                                        if (2 <= 557) {
                                            var v16 int32 = -957
                                            v17 := GenS5{F0: v3}
                                            if false {
                                                v16 = v16
                                                fmt.Println(v16)
                                                v3 = (!((!v3) || (0.1 < 0.1)))
                                            } else {
                                                if v3 {
                                                    var v18 int64 = 9223372036854775807
                                                    v19 := GenS3{F0: 42}
                                                    if (!v3) {
                                                        if (1 >= 42) {
                                                            fmt.Println(fi4)
                                                        } else {
                                                            var v20 float32 = 3.14159
                                                            _ = v20
                                                        }
                                                    }
                                                    _ = v18
                                                    _ = v19
                                                }
                                                _ = len(v1)
                                                fmt.Println(v3, v2)
                                            }
                                            _ = v16
                                            _ = v17
                                        } else {
                                            if false {
                                                var v21 int64 = -1
                                                v22 := GenS6{F0: false}
                                                if ((^(fi13)) >= fi4) {
                                                    v23 := []int{v2, (fi4 ^ fi4), (fi4 & v2), 170}
                                                    var v24 uint32 = 42
                                                    _ = v23
                                                    _ = v24
                                                }
                                                _ = v21
                                                _ = v22
                                            } else {
                                                fmt.Println(fi13)
                                                var v25 uint8 = 2
                                                switch (v2 ^ fi13) {
                                                case 9223372036854775807:
                                                    var v26 int = 42
                                                    _ = v26
                                                case 42:
                                                    v27 := GenS7{F0: 42}
                                                    _ = v27
                                                }
                                                fmt.Println(v2)
                                                _ = v25
                                            }
                                            if (-10 <= (fi4 - (^(fi4)))) {
                                                v28 := GenS2{F0: (false || v3)}
                                                fmt.Println(fi4, v3, v2)
                                                v3 = false
                                                delete(v1, fi4)
                                                _ = v28
                                            }
                                        }
                                        _ = fi13
                                    }
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        if true {
                                            fmt.Println(fi4, v3, v3)
                                            v29 := GenS8{F0: 505, F1: 65535, F2: 0}
                                            _ = v29
                                        } else {
                                            for fi30 := 0; fi30 < 17; fi30++ {
                                                if true {
                                                    fmt.Println(v3)
                                                } else {
                                                    var v31 uint64 = 230
                                                    var v32 float32 = float32(3.14159)
                                                    var v33 int = fi4
                                                    var v34 int32 = -2147483648
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                }
                                                var v35 int64 = 42
                                                if true {
                                                    fmt.Println(v35, fi30)
                                                    fmt.Println(fi4, fi30, v3)
                                                    if (v3 || false) {
                                                        v36 := GenS3{F0: 0}
                                                        _ = v36
                                                    }
                                                } else {
                                                    var v37 float32 = 1.5
                                                    v2 += fi4
                                                    v3 = ((v37 + v37) <= v37)
                                                    _ = v37
                                                }
                                                _ = fi30
                                                _ = v35
                                            }
                                            var v38 float64 = 0.0
                                            _ = v38
                                        }
                                    case "hello":
                                        v3 = (0 == 868)
                                    case "a":
                                        var v39 int16 = 0
                                        _ = v39
                                    default:
                                        var v40 uint64 = 42
                                        _ = v40
                                    }
                                    if v3 {
                                        v41 := GenS9{F0: 65535, F1: -1}
                                        fmt.Println(v2)
                                        _ = v41
                                    } else {
                                        v42 := GenS10{F0: 575, F1: 1, F2: (fi4 + -1)}
                                        v2 = v2
                                        _ = v42
                                    }
                                }
                                v43 := GenS11{F0: 65535, F1: -319}
                                v44 := GenS12{F0: -409, F1: 1}
                                _ = v43
                                _ = v44
                            }
                        default:
                            v2 -= 4294967295
                        }
                        if true {
                            var v45 uint16 = 2
                            switch (v2 + (-(v2))) {
                            case -2147483648:
                                switch "\n\t" {
                                case "foo bar":
                                    v2 = (fi4 & fi4)
                                default:
                                    for fi46 := 0; fi46 < 12; fi46++ {
                                        v47 := []int{(v2 + (fi46 >> ((822) & 63))), (v2 + 127)}
                                        v48 := GenS13{F0: 0}
                                        if (2 == 18446744073709551615) {
                                            if ("" != "hello") {
                                                v49 := GenS14{F0: 0}
                                                _ = v49
                                            } else {
                                                fmt.Println(v2)
                                                if len(v47) > 0 {
                                                    v47[0] = (^(v2))
                                                }
                                                v1[(fi4 * -1)] = 197
                                            }
                                        } else {
                                            v50 := GenS9{F0: 153, F1: -32768}
                                            _ = v50
                                        }
                                        _ = fi46
                                        _ = v47
                                        _ = v48
                                    }
                                }
                            }
                            switch ("" + "foo bar") {
                            case "hello":
                                fmt.Println(fi4, v3)
                            case "\n\t":
                                var v51 int8 = 1
                                _ = v51
                            case "":
                                if v3 {
                                    v52 := GenS2{F0: false}
                                    v2 -= -267
                                    fmt.Println(v3)
                                    _ = v52
                                } else {
                                    var v53 float32 = (((0.1 - 0.0) / 0.1) + (0.0 / 3.14159))
                                    v54 := GenS8{F0: 4294967295, F1: 380, F2: 94}
                                    v55 := []int{(-(fi4)), 4294967295, v2, (v2 / ((v2) | 1))}
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                            }
                            _ = v45
                        }
                        v3 = (707 < -193)
                    }
                    for fi56 := 0; fi56 < 13; fi56++ {
                        v57 := GenS15{F0: -32768, F1: "foo bar"}
                        v58 := GenS8{F0: 65535, F1: 18446744073709551615, F2: 255}
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                }
                var v59 uint16 = 1
                _ = fi4
                _ = v59
            }
            var v60 int64 = 1
            _ = v60
        }
    } else {
        v1[v2] = 2147483648
    }
    if true {
        v3 = true
        switch (("hello" + "a") + "") {
        case "foo bar":
            if (-2.25 < (0.1 + (0.1 / 1e10))) {
                var v61 int = 4294967295
                _ = v61
            } else {
                var v62 int8 = 1
                var v63 float32 = 0.0
                v64 := []int{(v2 << ((4294967295) & 63))}
                _ = v62
                _ = v63
                _ = v64
            }
        case "":
            for fi65 := 0; fi65 < 19; fi65++ {
                fmt.Println(v2)
                _ = fi65
            }
        default:
            var v66 float64 = -2.25
            _ = v66
        }
    }
    var v67 uint8 = 2
    v2 = 127
    for fi68 := 0; fi68 < 18; fi68++ {
        fmt.Println(v3, fi68, v67)
        v69 := []int{2147483647, -9223372036854775808, 4294967295, (^(fi68))}
        for fi70 := 0; fi70 < 11; fi70++ {
            var v71 string = ""
            v71 = (("\n\t" + v71) + v71)
            _ = fi70
            _ = v71
        }
        _ = fi68
        _ = v69
    }
    v72 := []int{v2, v2}
    fmt.Println(v3, v3)
    for fi73 := 0; fi73 < 2; fi73++ {
        if (0 > 32767) {
            var v74 uint = 200
            var v75 uint64 = 1
            _ = v74
            _ = v75
        }
        for fi76 := 0; fi76 < 8; fi76++ {
            var v77 uint32 = 2
            _ = len(v1)
            _ = fi76
            _ = v77
        }
        if v3 {
            v2 = v2
            switch "a" {
            case "hello":
                v78 := GenS14{F0: -718}
                _ = v78
            case "foo bar":
                var v79 float32 = 3.14159
                _ = v79
            default:
                fmt.Println(fi73, v3)
            }
            delete(v1, fi73)
        }
        _ = fi73
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v72:", len(v72), v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v67
    _ = v72
}
