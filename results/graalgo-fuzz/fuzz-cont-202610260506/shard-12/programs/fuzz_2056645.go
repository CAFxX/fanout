package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint; F2 uint }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS6 struct { F0 uint32; F1 int16; F2 bool }
func genFn0(p0 int) string {
    p0 -= 4294967295
    p0 = (^(p0))
    v5 := GenS1{F0: -32768}
    p0 = (p0 / ((-9223372036854775808) | 1))
    if ((p0 >> ((4294967295) & 63)) != p0) {
        fmt.Println(p0)
    }
    _ = p0
    _ = v5
    return ""
}

func main() {
    var v1 uint8 = 144
    v2 := make(map[int]uint32)
    var v3 int = 783
    var v4 float32 = (3.14159 + (3.14159 - (1.5 - -2.25)))
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int8 = 75
        switch (("" + ("\n\t" + "\n\t")) + "a") {
        case "a":
            v7 -= 1
        case "\n\t":
            v8 := []int{128, v3}
            _ = v8
        case "foo bar":
            var v9 uint16 = 255
            _ = v9
        default:
            v7 = v7
        }
        _ = fi6
        _ = v7
    }
    v1 = v1
    v10 := GenS2{F0: 257, F1: 0, F2: v1}
    if false {
        if (true && (-1 > 42)) {
            var v11 int16 = 0
            _ = v11
        }
        fmt.Println(v3, v4, v4)
    } else {
        if true {
            if false {
                var v12 uint32 = 4294967295
                v13 := GenS2{F0: 255, F1: 623, F2: v1}
                _ = v12
                _ = v13
            } else {
                var v14 bool = (!true)
                _ = v14
            }
            v15 := GenS3{F0: 1, F1: 65535, F2: 0}
            _ = v15
        } else {
            for fi16 := 0; fi16 < 16; fi16++ {
                for fi17 := 0; fi17 < 13; fi17++ {
                    if true {
                        var v18 string = ""
                        v19 := []int{fi16, -65, 32768, (fi16 >> ((1) & 63))}
                        var v20 uint32 = 42
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        if (1 < 0) {
                            if (-1 == -1) {
                                if (!true) {
                                    v21 := v2[v3]
                                    if (true || false) {
                                        _ = len(v2)
                                        if ("" == ("a" + "hello")) {
                                            fmt.Println(v21)
                                            var v22 int64 = 42
                                            _ = v22
                                        }
                                    } else {
                                        v23 := GenS2{F0: 1, F1: 2, F2: 0}
                                        v24 := GenS3{F0: -1, F1: 0, F2: 4294967295}
                                        v25 := GenS1{F0: -953}
                                        if (42 > 100) {
                                            v3 += fi17
                                            switch "hello" {
                                            case "a":
                                                if (!false) {
                                                    var v26 string = ("a" + ((("foo bar" + "foo bar") + "\n\t") + "foo bar"))
                                                    switch (v26 + ("hello" + ((v26 + v26) + v26))) {
                                                    case "":
                                                        var v27 uint32 = v21
                                                        _ = v27
                                                    default:
                                                        fmt.Println(v1, v3, v3)
                                                    }
                                                    v28 := v2[-2147483648]
                                                    _ = v26
                                                    _ = v28
                                                }
                                            case "hello":
                                                v29 := GenS1{F0: 1000}
                                                _ = v29
                                            case "\n\t":
                                                v1 -= (v1 & v1)
                                            }
                                            fmt.Println(fi16)
                                            if false {
                                                var v30 uint16 = 32767
                                                var v31 int16 = -271
                                                v32 := []int{v3, fi17}
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            } else {
                                                v33 := GenS1{F0: 1000}
                                                if (v21 == (v21 - (v21 ^ 400))) {
                                                    v21 += v21
                                                    fmt.Println(fi17)
                                                } else {
                                                    v34 := []int{fi16, (^(fi17)), (fi16 << ((502) & 63)), fi16}
                                                    switch "hello" {
                                                    case "hello":
                                                        var v35 float32 = float32((3.14159 * float64(((-2.25 / 1.0) / -2.25))))
                                                        _ = v35
                                                    default:
                                                        v4 += v4
                                                    }
                                                    _ = v34
                                                }
                                                v3 -= fi17
                                                _ = v33
                                            }
                                            v3 -= (fi16 << ((24) & 63))
                                        } else {
                                            var v36 int8 = 42
                                            v37 := GenS2{F0: 2, F1: 255, F2: (v1 ^ (v1 ^ (v1 ^ (^(v1)))))}
                                            v1 -= v1
                                            v36 -= 7
                                            v3 = fi17
                                            v4 = 1.5
                                            if true {
                                                if (((v3 + v3) == 127) && (9223372036854775807 > -9223372036854775808)) {
                                                    fmt.Println(v4, v3, v1)
                                                    var v38 uint32 = v21
                                                    _ = v38
                                                }
                                            } else {
                                                v39 := GenS3{F0: 1, F1: 42, F2: 644}
                                                _ = v39
                                            }
                                            _ = v36
                                            _ = v37
                                        }
                                        if false {
                                            var v40 string = ""
                                            if false {
                                                var v41 uint16 = 2
                                                fmt.Println(v41, fi16)
                                                _ = v41
                                            } else {
                                                var v42 float32 = v4
                                                switch v3 {
                                                case 2147483647:
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v43 := GenS4{F0: 931, F1: 1}
                                                        _ = v43
                                                    case "":
                                                        v44 := []int{(^(fi16)), fi17, -343, 290}
                                                        _ = v44
                                                    case "hello":
                                                        var v45 int8 = -125
                                                        _ = v45
                                                    }
                                                case 128:
                                                    v46 := GenS1{F0: 209}
                                                    _ = v46
                                                default:
                                                    var v47 uint64 = 558
                                                    _ = v47
                                                }
                                                _ = v42
                                            }
                                            v48 := GenS3{F0: 9223372036854775807, F1: 104, F2: 65535}
                                            _ = v40
                                            _ = v48
                                        } else {
                                            fmt.Println(v3, v4, fi16)
                                        }
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                    _ = v21
                                } else {
                                    if true {
                                        delete(v2, (-(fi17)))
                                    } else {
                                        v4 -= 3.14159
                                    }
                                }
                                fmt.Println(v4, fi16, v1)
                            } else {
                                v4 -= (float32(0.0) - 0.0)
                                v49 := []int{2147483647}
                                if (("" != "") || true) {
                                    switch "" {
                                    case "":
                                        v50 := GenS5{F0: (v1 << ((18446744073709551615) & 63)), F1: 255, F2: 65}
                                        _ = v50
                                    }
                                    switch ("\n\t" + "foo bar") {
                                    case "foo bar":
                                        var v51 uint8 = (^(v1))
                                        _ = v51
                                    case "a":
                                        var v52 int32 = -478
                                        _ = v52
                                    case "":
                                        fmt.Println(v1, v3)
                                    }
                                    fmt.Println(v3)
                                }
                                v53 := GenS4{F0: 350, F1: 0}
                                if ((false || (1e10 > v4)) && (!false)) {
                                    var v54 string = "hello"
                                    _ = v54
                                }
                                _ = v49
                                _ = v53
                            }
                            v4 -= (0.1 / 0.1)
                            var v55 int = (fi17 ^ (v3 - int(v1)))
                            _ = v55
                        }
                        var v56 int64 = -1
                        var v57 float64 = (3.14159 - 3.14159)
                        _ = v56
                        _ = v57
                    }
                    if false {
                        var v58 string = ("\n\t" + "foo bar")
                        v59 := GenS6{F0: 706, F1: -1, F2: (-1 < 121)}
                        _ = v58
                        _ = v59
                    }
                    _ = fi17
                }
                var v60 bool = false
                _ = fi16
                _ = v60
            }
            v3 = -554
        }
    }
    fmt.Println(v3)
    v3 = (v3 - v3)
    v61 := GenS5{F0: (v1 * v1), F1: uint16(v1), F2: 127}
    v62 := []int{v3}
    var v63 int = v3
    fmt.Println(v3, v4, v1)
    _ = len(v2)
    v64 := GenS2{F0: 42, F1: 65535, F2: (v1 / (((v1 & (^(v1)))) | 1))}
    v63 += 32768
    for fi65 := 0; fi65 < 9; fi65++ {
        var v66 uint32 = 2147483647
        v66 += 2
        delete(v2, (v63 / ((v3) | 1)))
        _ = fi65
        _ = v66
    }
    fmt.Println(v63, v3)
    if true {
        switch "hello" {
        case "\n\t":
            switch v63 {
            case -1:
                v3 -= v63
            }
        case "hello":
            delete(v2, v3)
        case "foo bar":
            if (-2147483648 > v63) {
                if len(v62) > 0 {
                    v62[0] = (-(v63))
                }
                var v67 uint = 276
                if (((18446744073709551615 >= 2) || false) || false) {
                    switch "a" {
                    case "":
                        v62 = append(v62, 4294967295)
                    case "hello":
                        v3 -= -117
                    }
                    v3 += (v3 / (((v3 % (((v3 | v63)) | 1))) | 1))
                    v62 = append(v62, v3)
                }
                _ = v67
            } else {
                fmt.Println(v4)
                if true {
                    v1 = (v1 & v1)
                    fmt.Println(v63)
                    var v68 uint8 = 2
                    _ = v68
                }
                v63 = 478
            }
        default:
            v63 = (v63 & v63)
        }
    } else {
        for fi69 := 0; fi69 < 12; fi69++ {
            var v70 int16 = 42
            var v71 string = "a"
            if true {
                v3 = fi69
            }
            _ = fi69
            _ = v70
            _ = v71
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v62
    _ = v63
    _ = v64
}
