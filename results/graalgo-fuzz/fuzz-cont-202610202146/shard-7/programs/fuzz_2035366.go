package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 int64; F2 int }
type GenS8 struct { F0 bool }
func genFn0(p0 bool, p1 int32, p2 int16) int32 {
    var v7 int8 = 1
    var v8 string = ("a" + "a")
    switch (v8 + "a") {
    case "":
        var v9 int = -2147483648
        _ = v9
    default:
        v10 := []int{32767, 128, 42, 32768}
        _ = v10
    }
    v7 += (v7 * int8(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -55
}
func genFn1(p0 uint32, p1 int) int64 {
    switch "\n\t" {
    case "":
        v11 := GenS1{F0: "", F1: p0}
        _ = v11
    case "foo bar":
        fmt.Println(p0, p1)
    case "a":
        if true {
            v12 := GenS2{F0: -1, F1: p1}
            v13 := []int{}
            _ = v12
            _ = v13
        } else {
            fmt.Println(p0, p0)
        }
    }
    p1 += p1
    if true {
        if (368 == -63) {
            v14 := []int{(p1 >> ((4294967295) & 63)), p1, (p1 / (((p1 - p1)) | 1))}
            _ = v14
        } else {
            fmt.Println(p0, p0)
            fmt.Println(p1, p1)
            var v15 int64 = 1
            _ = v15
        }
        v16 := GenS1{F0: "hello", F1: 943}
        var v17 uint = 42
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: "", F1: 2}
    if false {
        v19 := []int{}
        p1 -= -1
        for fi20 := 0; fi20 < 2; fi20++ {
            p0 += p0
            var v21 string = ""
            p1 -= (fi20 >> ((255) & 63))
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v18
    return -1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    v3 := []string{"", "\n\t", ("hello" + ("a" + "foo bar")), "a"}
    var v4 float64 = (0.0 * 0.1)
    var v5 bool = false
    var v6 float64 = float64(1e10)
    delete(v2, -871)
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 uint8 = 0
            v4 -= (v6 / 1e10)
            delete(v1, "")
            var v25 int16 = 507
            v5 = (!(802 <= 65535))
            v25 = (v25 >> ((1) & 63))
            _ = fi23
            _ = v24
            _ = v25
        }
        v6 = 3.14159
        var v26 uint32 = 4294967295
        _ = fi22
        _ = v26
    }
    v4 -= v4
    var v27 string = ("foo bar" + "foo bar")
    v4 -= ((0.0 + 1.5) * (1e10 - 1.5))
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 int = (fi28 | (fi28 - 4294967295))
        var v30 uint32 = 4294967295
        v27 = v27
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := GenS3{F0: -880, F1: 42}
    for fi32 := 0; fi32 < 16; fi32++ {
        fmt.Println(v4, v6)
        fmt.Println(v6)
        _ = fi32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        v5 = (((42 != genFn0((-2.25 != 3.14159), 2147483647, 1)) && true) && false)
        _ = fi33
    }
    var v34 uint16 = 255
    v5 = (1 != 1)
    delete(v2, 1)
    v27 = v27
    for fi35 := 0; fi35 < 19; fi35++ {
        if v5 {
            for fi36 := 0; fi36 < 8; fi36++ {
                v37 := GenS2{F0: 1000, F1: fi35}
                if ("a" >= v27) {
                    switch 9223372036854775807 {
                    case 315:
                        v34 -= (v34 * (v34 & (v34 - (v34 * (^(v34))))))
                    default:
                        var v38 int32 = 933
                        _ = v38
                    }
                    var v39 int8 = 100
                    _ = v39
                } else {
                    v34 = 324
                    v27 = (v27 + v27)
                    var v40 int8 = -128
                    var v41 uint16 = 65535
                    switch fi35 {
                    case -2147483648:
                        v5 = (2147483647 >= 2147483648)
                    case 4294967295:
                        v42 := GenS2{F0: -32768, F1: (^(fi36))}
                        _ = v42
                    case 128:
                        var v43 float32 = 0.0
                        _ = v43
                    default:
                        var v44 uint8 = 1
                        _ = v44
                    }
                    var v45 int = fi36
                    _ = v40
                    _ = v41
                    _ = v45
                }
                _ = fi36
                _ = v37
            }
            var v46 uint64 = 1
            if (true && (1 > -116)) {
                switch fi35 {
                case -320:
                    for fi47 := 0; fi47 < 15; fi47++ {
                        var v48 bool = (fi47 > (fi35 & 128))
                        fmt.Println(fi47, v48)
                        _ = fi47
                        _ = v48
                    }
                case 127:
                    var v49 uint64 = v46
                    _ = v49
                default:
                    v50 := GenS2{F0: -32768, F1: fi35}
                    _ = v50
                }
                if (true && true) {
                    v51 := v1[v27]
                    _ = v51
                }
            } else {
                var v52 int8 = 0
                if ((fi35 + 32768) != (fi35 + (fi35 - fi35))) {
                    switch v27 {
                    case "a":
                        fmt.Println(v5)
                    case "hello":
                        for fi53 := 0; fi53 < 8; fi53++ {
                            v6 = ((-2.25 + float64(float32((v4 * v4)))) / (v4 - v4))
                            _ = fi53
                        }
                    case "\n\t":
                        for fi54 := 0; fi54 < 8; fi54++ {
                            var v55 uint = 1
                            var v56 bool = (v34 >= v34)
                            fmt.Println(v34, fi35, v5)
                            _ = fi54
                            _ = v55
                            _ = v56
                        }
                    default:
                        for fi57 := 0; fi57 < 14; fi57++ {
                            if true {
                                var v58 uint16 = v34
                                v59 := GenS1{F0: "foo bar", F1: uint32(v34)}
                                _ = v58
                                _ = v59
                            }
                            var v60 uint = 58
                            if (!(v60 == v60)) {
                                switch "foo bar" {
                                case "foo bar":
                                    v61 := GenS4{F0: 128, F1: (true && v5)}
                                    _ = v61
                                case "a":
                                    v5 = ((v34 * (v34 / ((v34) | 1))) <= 264)
                                case "":
                                    var v62 int32 = 0
                                    _ = v62
                                default:
                                    var v63 uint16 = (v34 % ((32767) | 1))
                                    _ = v63
                                }
                                if v5 {
                                    fmt.Println(v27, v5)
                                    var v64 float32 = 3.14159
                                    _ = v64
                                } else {
                                    if v5 {
                                        v65 := GenS5{F0: 0}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if v5 {
                                            v6 = v6
                                        }
                                        _ = v65
                                    } else {
                                        switch (v27 + v27) {
                                        case "":
                                            var v66 uint16 = 963
                                            _ = v66
                                        case "\n\t":
                                            var v67 int = fi57
                                            _ = v67
                                        default:
                                            if false {
                                                var v68 float64 = (0.0 + v4)
                                                v69 := v2[(fi57 % ((0) | 1))]
                                                v68 += (v6 - 1e10)
                                                _ = v68
                                                _ = v69
                                            } else {
                                                if v5 {
                                                    if len(v3) > 0 {
                                                        v3[0] = v27
                                                    }
                                                } else {
                                                    if (((61 != 42) || (255 != 42)) || (((3.14159 * 0.1) * 1e10) > (-2.25 - 0.1))) {
                                                        fmt.Println(v46, v6, v5)
                                                        if false {
                                                            var v70 bool = false
                                                            v71 := GenS4{F0: 1, F1: false}
                                                            _ = v70
                                                            _ = v71
                                                        } else {
                                                            v5 = ("" != "a")
                                                            if v5 {
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v27 + "a")
                                                                }
                                                                v72 := GenS2{F0: 0, F1: 9223372036854775807}
                                                                if ((0.1 / (float32(v4) - -2.25)) > 0.0) {
                                                                    v60 = (v60 & 4294967295)
                                                                    var v73 int64 = 0
                                                                    switch ((v27 + "a") + "") {
                                                                    case "foo bar":
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    default:
                                                                        delete(v1, "a")
                                                                    }
                                                                    _ = v73
                                                                }
                                                                _ = v72
                                                            } else {
                                                                v74 := v1[(v27 + (("\n\t" + (("" + v27) + "foo bar")) + ("hello" + v27)))]
                                                                var v75 bool = false
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                            v34 -= v34
                                                        }
                                                        switch (v27 + ((("" + "foo bar") + v27) + (v27 + (v27 + v27)))) {
                                                        case "hello":
                                                            if (1 == 0) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = ("hello" + (v27 + v27))
                                                                }
                                                                if true {
                                                                    v76 := GenS4{F0: 1, F1: true}
                                                                    v5 = (!(42 >= 0))
                                                                    _ = v76
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = "hello"
                                                                }
                                                            }
                                                        case "\n\t":
                                                            if v5 {
                                                                switch fi35 {
                                                                case 0:
                                                                    if (!(false && (128 < 44))) {
                                                                        var v77 uint8 = 167
                                                                        _ = v77
                                                                    } else {
                                                                        v6 = v6
                                                                        var v78 int32 = -1
                                                                        _ = v78
                                                                    }
                                                                case 128:
                                                                    v52 = (-(v52))
                                                                }
                                                                v5 = v5
                                                            } else {
                                                                v79 := v2[(fi35 - fi57)]
                                                                v80 := GenS6{F0: 2147483648}
                                                                _ = v79
                                                                _ = v80
                                                            }
                                                        }
                                                    }
                                                    v1[(v27 + (("" + ("foo bar" + (v27 + (v27 + "\n\t")))) + (v27 + "foo bar")))] = v27
                                                }
                                            }
                                        }
                                        if (2147483647 == genFn0(v5, 2147483647, -437)) {
                                            var v81 float32 = 3.14159
                                            v1["a"] = (v27 + ("\n\t" + v27))
                                            _ = v81
                                        }
                                        var v82 float32 = 1.5
                                        _ = v82
                                    }
                                    v46 += (^(v46))
                                }
                            }
                            _ = fi57
                            _ = v60
                        }
                    }
                    _ = len(v3)
                }
                var v83 uint64 = 18446744073709551615
                _ = v52
                _ = v83
            }
            _ = v46
        }
        _ = len(v2)
        v84 := GenS7{F0: 116, F1: genFn1(2147483648, fi35), F2: (fi35 >> ((1) & 63))}
        _ = fi35
        _ = v84
    }
    for fi85 := 0; fi85 < 16; fi85++ {
        v86 := GenS8{F0: true}
        _ = fi85
        _ = v86
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
