package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 int }
func genFn0() uint16 {
    v5 := GenS1{F0: 42}
    if (!true) {
        switch "\n\t" {
        case "\n\t":
            var v6 int16 = -227
            _ = v6
        case "":
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := []int{1, fi7, (fi7 << ((420) & 63)), 460}
                v8 = append(v8, fi7)
                var v9 int64 = int64(fi7)
                _ = fi7
                _ = v8
                _ = v9
            }
        }
    } else {
        switch "hello" {
        case "":
            var v10 int8 = -78
            _ = v10
        case "hello":
            v11 := []int{-549, 128, -1, 211}
            _ = v11
        }
        var v12 int8 = 100
        v13 := GenS1{F0: 65535}
        _ = v12
        _ = v13
    }
    _ = v5
    return 677
}
func genFn1(p0 uint32, p1 uint32) string {
    v14 := GenS2{F0: 802}
    for fi15 := 0; fi15 < 6; fi15++ {
        if false {
            var v16 int8 = 42
            var v17 uint8 = 41
            _ = v16
            _ = v17
        } else {
            fmt.Println(fi15, fi15)
            var v18 uint64 = 4294967295
            v19 := GenS2{F0: 255}
            _ = v18
            _ = v19
        }
        p1 = 42
        _ = fi15
    }
    fmt.Println(p0)
    if true {
        var v20 uint16 = 32767
        v21 := GenS2{F0: v20}
        var v22 int64 = 1
        _ = v20
        _ = v21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v14
    return ("hello" + "foo bar")
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 42
    var v3 float32 = 3.14159
    var v4 float64 = (3.14159 / 1.0)
    for fi23 := 0; fi23 < 17; fi23++ {
        fmt.Println(fi23)
        v24 := []int{(fi23 ^ fi23), 32768}
        _ = fi23
        _ = v24
    }
    if false {
        if (!("foo bar" == "foo bar")) {
            fmt.Println(v2, v3)
            v3 = ((v3 * float32(1e10)) / (v3 / (v3 - 1.5)))
        } else {
            for fi25 := 0; fi25 < 9; fi25++ {
                v4 -= ((v4 / v4) + 1.5)
                v26 := GenS3{F0: 42, F1: 255}
                _ = fi25
                _ = v26
            }
            if true {
                v3 = (v3 * (1e10 / 3.14159))
                v27 := GenS1{F0: 65535}
                _ = v27
            } else {
                fmt.Println(v3)
                var v28 bool = ((!(false && (("a" < "a") || true))) || (!false))
                v29 := []int{1, -2147483648, -1}
                _ = v28
                _ = v29
            }
        }
        switch "" {
        case "hello":
            var v30 int64 = -9223372036854775808
            _ = v30
        default:
            var v31 float32 = (v3 + (((((v3 / v3) * 3.14159) * v3) + ((1.5 - ((v3 / float32(1e10)) * v3)) * v3)) + (0.1 - 0.1)))
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        v33 := GenS3{F0: 1, F1: 10}
        _ = fi32
        _ = v33
    }
    var v34 uint32 = 2147483647
    for fi35 := 0; fi35 < 16; fi35++ {
        v2 += 65535
        if (!true) {
            var v36 uint = 65535
            if false {
                var v37 int16 = -635
                if (255 <= 1) {
                    switch 128 {
                    case -1:
                        for fi38 := 0; fi38 < 9; fi38++ {
                            switch "" {
                            case "foo bar":
                                if (!false) {
                                    v39 := GenS1{F0: 2147483648}
                                    _ = v39
                                } else {
                                    var v40 int = -2147483648
                                    v41 := GenS2{F0: (v2 | v2)}
                                    _ = v40
                                    _ = v41
                                }
                            default:
                                fmt.Println(fi35)
                            }
                            _ = fi38
                        }
                    case 9223372036854775807:
                        v42 := GenS4{F0: -2147483648}
                        _ = v42
                    default:
                        v43 := []int{fi35, fi35}
                        _ = v43
                    }
                    v4 = float64(v3)
                }
                switch ((("" + "\n\t") + (("hello" + "hello") + "hello")) + "hello") {
                case "a":
                    fmt.Println(fi35)
                case "":
                    v34 += (v34 * v34)
                default:
                    for fi44 := 0; fi44 < 14; fi44++ {
                        v45 := GenS2{F0: v2}
                        if true {
                            if false {
                                v46 := GenS3{F0: 6, F1: 255}
                                _ = len(v1)
                                var v47 int = (fi35 >> ((978) & 63))
                                _ = v46
                                _ = v47
                            }
                            var v48 uint32 = v34
                            fmt.Println(v36)
                            _ = v48
                        } else {
                            v49 := GenS5{F0: v37, F1: 255, F2: 1}
                            _ = v49
                        }
                        if true {
                            if (1 > 1) {
                                if true {
                                    v2 += (v2 >> ((42) & 63))
                                    if true {
                                        if true {
                                            fmt.Println(v36, v4)
                                            if ((v2 < v2) || ((v37 >> ((v36) & 63)) > (v37 & (v37 | v37)))) {
                                                v50 := GenS6{F0: 70, F1: 275}
                                                _ = v50
                                            }
                                            v4 += ((v4 / ((3.14159 / float64(v3)) - (0.1 - 1e10))) + v4)
                                        } else {
                                            fmt.Println(fi35, fi44)
                                            var v51 bool = false
                                            _ = v51
                                        }
                                        switch "\n\t" {
                                        case "":
                                            v52 := v1[v34]
                                            _ = v52
                                        default:
                                            v2 += (v2 * (v2 * (v2 >> ((v36) & 63))))
                                        }
                                        v37 = v37
                                    }
                                }
                            }
                            delete(v1, (v34 / ((4294967295) | 1)))
                        }
                        _ = fi44
                        _ = v45
                    }
                }
                _ = v37
            } else {
                _ = len(v1)
            }
            for fi53 := 0; fi53 < 3; fi53++ {
                fmt.Println(v3, v36, v34)
                var v54 float64 = ((0.1 / 1.0) + float64(v3))
                var v55 int32 = 1
                _ = fi53
                _ = v54
                _ = v55
            }
            _ = v36
        }
        _ = fi35
    }
    v56 := []int{-1, 1, 128}
    v57 := GenS7{F0: 0, F1: 2}
    fmt.Println(v2, v34, v4)
    var v58 float64 = v4
    v3 = 1e10
    fmt.Println(v58, v34, v3)
    v56 = append(v56, 128)
    if (42 > 65535) {
        switch 1 {
        case 0:
            v34 = (^(v34))
        default:
            if true {
                if false {
                    _ = len(v56)
                    fmt.Println(v34)
                    switch 963 {
                    case -9223372036854775808:
                        var v59 uint32 = 42
                        _ = v59
                    default:
                        v3 -= float32(v58)
                    }
                } else {
                    var v60 uint64 = 18446744073709551615
                    var v61 uint8 = 1
                    _ = v60
                    _ = v61
                }
            }
        }
        v62 := v1[(^(v34))]
        _ = v62
    } else {
        v63 := v1[v34]
        _ = v63
    }
    var v64 int32 = -2147483648
    v2 = (v2 + v2)
    var v65 uint = 940
    v66 := GenS8{F0: 2147483647}
    fmt.Println(v58, v3, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v58:", v58)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v56
    _ = v57
    _ = v58
    _ = v64
    _ = v65
    _ = v66
}
