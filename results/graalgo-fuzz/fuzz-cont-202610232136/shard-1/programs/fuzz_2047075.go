package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 string; F1 uint64; F2 int8 }
type GenS7 struct { F0 uint8; F1 bool; F2 string }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint32, p1 uint64) uint8 {
    if false {
        var v4 uint64 = 1
        p1 -= p1
        v5 := []int{458, 32768, -503}
        _ = v4
        _ = v5
    } else {
        var v6 uint64 = p1
        if ((^(p1)) > p1) {
            if true {
                v7 := []int{127}
                _ = v7
            }
        }
        _ = v6
    }
    v8 := GenS1{F0: "hello", F1: uint(p1)}
    var v9 int64 = -9223372036854775808
    switch "a" {
    case "":
        var v10 bool = (false || true)
        _ = v10
    case "a":
        var v11 float64 = float64(1.5)
        _ = v11
    case "\n\t":
        if false {
            v12 := GenS2{F0: 0}
            _ = v12
        }
    }
    v13 := []int{629}
    _ = p0
    _ = p1
    _ = v13
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 int64) int16 {
    var v14 string = "hello"
    v14 = v14
    v15 := []int{}
    var v16 int64 = (^(p0))
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -32768
    v3 := make(map[string]string)
    switch 32768 {
    case 276:
        v17 := []int{0}
        _ = v17
    case 1:
        if false {
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 int64 = -746
                _ = fi18
                _ = v19
            }
        } else {
            var v20 int64 = -9223372036854775808
            v2 = v2
            _ = v20
        }
    default:
        fmt.Println(v2)
    }
    switch -577 {
    case -1:
        for fi21 := 0; fi21 < 10; fi21++ {
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 uint = 18446744073709551615
                _ = fi22
                _ = v23
            }
            if ((v2 - v2) == (v2 % ((178) | 1))) {
                if (!false) {
                    var v24 uint = 4294967295
                    _ = v24
                } else {
                    var v25 uint16 = 1
                    if true {
                        var v26 uint32 = 0
                        _ = v26
                    } else {
                        var v27 int8 = 0
                        var v28 int8 = (^(v27))
                        v25 = (v25 ^ (v25 ^ (^(v25))))
                        _ = v27
                        _ = v28
                    }
                    _ = v25
                }
                for fi29 := 0; fi29 < 11; fi29++ {
                    switch "" {
                    case "":
                        var v30 int32 = -423
                        _ = v30
                    }
                    switch -2147483648 {
                    case 455:
                        v31 := []int{fi21, fi21, 42}
                        _ = v31
                    case 9223372036854775807:
                        v2 += 0
                    case -841:
                        switch (fi21 % (((^(fi21))) | 1)) {
                        case 128:
                            if true {
                                v32 := GenS3{F0: (!false), F1: 65535}
                                if false {
                                    var v33 uint32 = 635
                                    switch (fi21 - (fi21 << ((202) & 63))) {
                                    case -796:
                                        v34 := v1[(v33 - 401)]
                                        _ = v34
                                    case -346:
                                        var v35 int = fi29
                                        _ = v35
                                    default:
                                        var v36 uint8 = 86
                                        _ = v36
                                    }
                                    v37 := GenS2{F0: 128}
                                    _ = v33
                                    _ = v37
                                } else {
                                    v2 = (-(v2))
                                    if (false || true) {
                                        switch "hello" {
                                        case "hello":
                                            v38 := GenS1{F0: "\n\t", F1: 138}
                                            _ = v38
                                        case "a":
                                            if false {
                                                var v39 bool = false
                                                fmt.Println(v2)
                                                _ = v39
                                            } else {
                                                v40 := []int{-9223372036854775808, (fi21 & fi21)}
                                                _ = v40
                                            }
                                        }
                                        v2 -= (v2 + (v2 - (v2 >> ((1) & 63))))
                                    } else {
                                        v2 = (v2 ^ -1)
                                        _ = len(v1)
                                    }
                                    var v41 float64 = 3.14159
                                    v2 = 1
                                    switch "" {
                                    case "hello":
                                        v42 := GenS4{F0: 32767, F1: 4294967295}
                                        _ = v42
                                    case "a":
                                        var v43 int64 = -1
                                        _ = v43
                                    default:
                                        v2 -= (-(v2))
                                    }
                                    if false {
                                        delete(v3, ("" + ""))
                                        var v44 uint64 = 18446744073709551615
                                        _ = v44
                                    }
                                    _ = v41
                                }
                                v45 := GenS1{F0: ("foo bar" + ("foo bar" + ("" + (("a" + "hello") + "a")))), F1: 709}
                                _ = v32
                                _ = v45
                            }
                        default:
                            v3[("a" + "")] = "a"
                        }
                    }
                    _ = fi29
                }
            } else {
                if true {
                    v46 := GenS2{F0: genFn0(65535, 794)}
                    _ = v46
                } else {
                    for fi47 := 0; fi47 < 8; fi47++ {
                        v48 := GenS3{F0: false, F1: 1}
                        v2 += v2
                        v2 = (^(v2))
                        _ = fi47
                        _ = v48
                    }
                    var v49 int = (-(fi21))
                    var v50 int16 = v2
                    _ = v49
                    _ = v50
                }
                fmt.Println(v2, fi21)
            }
            v2 -= v2
            _ = fi21
        }
    default:
        var v51 string = "a"
        _ = v51
    }
    _ = len(v3)
    switch ("\n\t" + "\n\t") {
    case "a":
        v52 := GenS3{F0: false, F1: 4294967295}
        _ = v52
    case "\n\t":
        for fi53 := 0; fi53 < 10; fi53++ {
            v54 := GenS4{F0: 2, F1: 1}
            var v55 uint64 = 18446744073709551615
            for fi56 := 0; fi56 < 8; fi56++ {
                v57 := GenS5{F0: 0, F1: (173 == 42)}
                switch "hello" {
                case "hello":
                    var v58 uint = 18446744073709551615
                    _ = v58
                default:
                    v59 := []int{fi53}
                    _ = v59
                }
                _ = fi56
                _ = v57
            }
            _ = fi53
            _ = v54
            _ = v55
        }
    case "foo bar":
        if true {
            v2 = (v2 & v2)
        } else {
            if true {
                var v60 int8 = -1
                _ = v60
            }
            var v61 uint32 = 574
            var v62 int8 = -128
            _ = v61
            _ = v62
        }
    }
    switch "\n\t" {
    case "\n\t":
        if true {
            v63 := GenS3{F0: false, F1: 186}
            _ = v63
        }
    case "":
        v64 := GenS6{F0: "\n\t", F1: 18446744073709551615, F2: 100}
        _ = v64
    }
    var v65 float64 = float64(float32(1.5))
    switch -1 {
    case 32768:
        v66 := GenS7{F0: genFn0(0, 1), F1: (2 >= 42), F2: "hello"}
        _ = v66
    case 224:
        v2 = (^(v2))
    default:
        v2 -= v2
    }
    switch -616 {
    case 32767:
        fmt.Println(v2)
    case 42:
        v67 := GenS8{F0: 205, F1: (!(false || (true && (255 < 18446744073709551615))))}
        _ = v67
    }
    var v68 bool = (-2147483648 == -1)
    fmt.Println(v2, v68, v68)
    switch (("" + "a") + "") {
    case "a":
        _ = len(v3)
    case "foo bar":
        v69 := GenS9{F0: 831}
        _ = v69
    case "":
        v68 = (!v68)
    default:
        fmt.Println(v68, v65, v65)
    }
    var v70 uint8 = 127
    fmt.Println(v68)
    var v71 int32 = 2147483647
    v72 := GenS10{F0: -128, F1: 0}
    v65 -= ((v65 + v65) + v65)
    for fi73 := 0; fi73 < 20; fi73++ {
        for fi74 := 0; fi74 < 14; fi74++ {
            fmt.Println(v71, v70, v68)
            if v68 {
                var v75 uint16 = 0
                fmt.Println(v71, fi74, v75)
                if v68 {
                    v70 -= v70
                    var v76 string = "\n\t"
                    _ = v76
                }
                _ = v75
            }
            switch "foo bar" {
            case "hello":
                switch (("" + ("a" + "\n\t")) + ("a" + "\n\t")) {
                case "foo bar":
                    v70 = v70
                case "\n\t":
                    v77 := GenS4{F0: 467, F1: 1}
                    _ = v77
                case "a":
                    var v78 uint = 2
                    _ = v78
                }
            default:
                if v68 {
                    v79 := []int{(fi73 * fi73), (fi73 * (fi74 | (fi73 % ((993) | 1)))), fi74}
                    _ = v79
                } else {
                    if false {
                        var v80 int = (fi73 * (fi73 & (fi74 + -382)))
                        if (false || v68) {
                            var v81 bool = true
                            v82 := GenS9{F0: 1}
                            _ = v81
                            _ = v82
                        }
                        _ = v80
                    }
                    fmt.Println(fi74, v68)
                    v83 := []int{-737}
                    _ = v83
                }
            }
            var v84 int = 0
            _ = fi74
            _ = v84
        }
        v70 = v70
        _ = fi73
    }
    var v85 int8 = -98
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
    _ = v68
    _ = v70
    _ = v71
    _ = v72
    _ = v85
}
