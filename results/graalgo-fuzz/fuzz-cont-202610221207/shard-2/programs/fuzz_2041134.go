package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int; F1 int; F2 uint64 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 uint; F1 uint64 }
func genFn0(p0 int8) int {
    for fi4 := 0; fi4 < 5; fi4++ {
        fmt.Println(p0)
        var v5 int32 = -2147483648
        _ = fi4
        _ = v5
    }
    p0 = -21
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 4; fi7++ {
        if (208 < (v6 - 255)) {
            fmt.Println(v6)
        } else {
            var v8 uint16 = 106
            fmt.Println(v8)
            if (-9223372036854775808 <= int64(v8)) {
                v9 := []int{}
                var v10 uint32 = 2
                var v11 float64 = float64(0.0)
                _ = v10
                _ = v11
                _ = v9
            }
            _ = v8
        }
        switch "a" {
        case "hello":
            var v12 uint64 = 977
            _ = v12
        case "\n\t":
            v13 := []int{(fi7 >> ((v6) & 63)), fi7, fi7, fi7}
            _ = v13
        case "":
            switch (^(fi7)) {
            case 32767:
                v14 := []int{}
                _ = v14
            case -1:
                v15 := []int{128, 2147483647, (^(fi7))}
                _ = v15
            default:
                var v16 uint = 1
                _ = v16
            }
        }
        v6 = v6
        _ = fi7
    }
    _ = p0
    _ = v6
    return 127
}
func genFn1() uint64 {
    if true {
        var v17 uint16 = 255
        fmt.Println(v17)
        switch "" {
        case "foo bar":
            v18 := GenS1{F0: false}
            _ = v18
        case "a":
            for fi19 := 0; fi19 < 16; fi19++ {
                v17 = 509
                var v20 uint32 = 2147483647
                _ = fi19
                _ = v20
            }
        case "hello":
            var v21 float64 = float64(0.0)
            _ = v21
        }
        _ = v17
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 16; fi23++ {
            v24 := []int{(fi22 - fi22), (fi23 ^ genFn0(-128)), 4294967295}
            _ = fi23
            _ = v24
        }
        v25 := []int{genFn0(0), -2147483648, 1, fi22}
        fmt.Println(fi22)
        _ = fi22
        _ = v25
    }
    return 0
}

func main() {
    var v1 string = ""
    var v2 int = 32767
    v3 := make(map[int]int)
    var v26 int64 = -9223372036854775808
    var v27 int8 = 100
    if false {
        switch (v2 & (v2 * v2)) {
        case -9223372036854775808:
            v28 := GenS1{F0: (1 <= 152)}
            _ = v28
        case -787:
            var v29 uint = 255
            _ = v29
        case 42:
            delete(v3, (^(v2)))
        default:
            switch (v1 + v1) {
            case "hello":
                fmt.Println(v26)
            case "foo bar":
                var v30 bool = true
                _ = v30
            default:
                v31 := GenS2{F0: 2, F1: 364}
                _ = v31
            }
        }
        fmt.Println(v1, v27, v27)
        var v32 string = "foo bar"
        _ = v32
    } else {
        var v33 float64 = 0.1
        _ = v33
    }
    var v34 int16 = -32768
    v35 := GenS2{F0: 577, F1: v34}
    switch v2 {
    case 9223372036854775807:
        if false {
            v36 := []int{}
            _ = v36
        } else {
            fmt.Println(v26)
            var v37 bool = true
            if v37 {
                var v38 uint16 = 486
                for fi39 := 0; fi39 < 20; fi39++ {
                    for fi40 := 0; fi40 < 4; fi40++ {
                        if (v37 || true) {
                            if (!true) {
                                v41 := GenS1{F0: v37}
                                var v42 uint = 42
                                switch (fi39 * (fi40 + -2147483648)) {
                                case -9223372036854775808:
                                    var v43 float64 = 0.0
                                    _ = v43
                                default:
                                    if v37 {
                                        var v44 float64 = 0.0
                                        _ = v44
                                    } else {
                                        var v45 string = v1
                                        if v37 {
                                            var v46 bool = false
                                            var v47 int64 = (v26 / ((1) | 1))
                                            _ = v46
                                            _ = v47
                                        } else {
                                            v48 := GenS1{F0: v37}
                                            if v37 {
                                                v2 -= int(v26)
                                            }
                                            switch (v45 + "foo bar") {
                                            case "\n\t":
                                                if (v34 < 1000) {
                                                    var v49 int8 = v27
                                                    _ = v49
                                                } else {
                                                    v34 += 1
                                                    v2 = int(v34)
                                                }
                                            case "a":
                                                if true {
                                                    switch v1 {
                                                    case "foo bar":
                                                        if ((^(v27)) >= v27) {
                                                            v27 = v27
                                                        }
                                                        if v37 {
                                                            var v50 uint64 = 547
                                                            switch fi39 {
                                                            case -9223372036854775808:
                                                                if ((fi39 | (v2 ^ -453)) != genFn0(v27)) {
                                                                    switch v2 {
                                                                    case 128:
                                                                        v2 += v2
                                                                    }
                                                                } else {
                                                                    v51 := GenS1{F0: (v37 && (!v37))}
                                                                    if (uint(v50) < 5) {
                                                                        v37 = true
                                                                    }
                                                                    v52 := GenS3{F0: v2, F1: 2147483647, F2: genFn1()}
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                            }
                                                            v3[fi40] = (-(fi39))
                                                            _ = v50
                                                        }
                                                        v53 := GenS2{F0: 2147483647, F1: v34}
                                                        _ = v53
                                                    default:
                                                        if (v37 || v37) {
                                                            v42 += (v42 % ((v42) | 1))
                                                            v54 := []int{v2, fi40}
                                                            _ = v54
                                                        } else {
                                                            v55 := []int{}
                                                            v38 = 42
                                                            _ = v55
                                                        }
                                                    }
                                                    v56 := GenS1{F0: true}
                                                    _ = v56
                                                }
                                            }
                                            var v57 string = "a"
                                            _ = v48
                                            _ = v57
                                        }
                                        _ = v45
                                    }
                                }
                                _ = v41
                                _ = v42
                            } else {
                                fmt.Println(fi39, v37)
                                switch fi39 {
                                case 128:
                                    v37 = (1 > 2)
                                case 2147483647:
                                    v1 = v1
                                default:
                                    switch "a" {
                                    case "hello":
                                        v34 += v34
                                    default:
                                        if (!(!(42 >= 0))) {
                                            var v58 float64 = 0.1
                                            var v59 uint8 = 0
                                            v60 := GenS4{F0: 32768, F1: (-(v34))}
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        } else {
                                            var v61 uint16 = (v38 % ((v38) | 1))
                                            if (((v1 + v1) <= v1) || (!v37)) {
                                                fmt.Println(v61)
                                                fmt.Println(v38, fi39)
                                                fmt.Println(v37, fi39, v26)
                                            } else {
                                                v62 := GenS4{F0: genFn0(0), F1: v34}
                                                v34 = (^(v34))
                                                v63 := GenS5{F0: 255, F1: 1}
                                                _ = v62
                                                _ = v63
                                            }
                                            _ = v61
                                        }
                                    }
                                }
                            }
                            if v37 {
                                _ = len(v3)
                            } else {
                                if (0 < 4294967295) {
                                    var v64 uint8 = 255
                                    _ = v64
                                } else {
                                    v2 = genFn0((^(v27)))
                                    var v65 int64 = (v26 / ((v26) | 1))
                                    fmt.Println(v38, fi39)
                                    _ = v65
                                }
                            }
                        }
                        var v66 float64 = 3.14159
                        if v37 {
                            var v67 uint16 = 255
                            var v68 uint16 = (v38 + (v38 * v38))
                            var v69 float32 = 3.14159
                            var v70 string = (v1 + v1)
                            _ = v67
                            _ = v68
                            _ = v69
                            _ = v70
                        } else {
                            var v71 int16 = v34
                            if v37 {
                                if v37 {
                                    switch (v1 + v1) {
                                    case "foo bar":
                                        v2 = fi39
                                    }
                                }
                                var v72 uint64 = 4294967295
                                _ = v72
                            } else {
                                var v73 uint8 = 85
                                v38 = (v38 - v38)
                                v34 = 0
                                _ = v73
                            }
                            var v74 uint16 = v38
                            _ = v71
                            _ = v74
                        }
                        _ = fi40
                        _ = v66
                    }
                    _ = fi39
                }
                _ = v38
            }
            _ = v37
        }
    case 32768:
        var v75 uint64 = 448
        _ = v75
    case 0:
        if (!true) {
            var v76 uint32 = 42
            _ = v76
        } else {
            var v77 uint64 = 2
            fmt.Println(v26, v26)
            _ = v77
        }
    default:
        if false {
            v27 = (v27 - v27)
            var v78 float64 = 1.5
            v27 -= (v27 * (v27 * (v27 * (v27 | v27))))
            _ = v78
        }
    }
    _ = len(v3)
    v26 -= v26
    if false {
        _ = len(v3)
    } else {
        var v79 uint64 = 42
        _ = v79
    }
    var v80 uint64 = 406
    for fi81 := 0; fi81 < 16; fi81++ {
        v26 = v26
        _ = fi81
    }
    if (!true) {
        var v82 bool = ((("foo bar" + "a") > v1) || false)
        _ = v82
    }
    var v83 uint8 = 1
    v84 := []int{v2, (v2 << ((65535) & 63)), -1, v2}
    var v85 int64 = v26
    v83 = (v83 ^ v83)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v80:", v80)
    fmt.Println("v83:", v83)
    fmt.Println("slicelen v84:", len(v84), v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v80
    _ = v83
    _ = v84
    _ = v85
}
