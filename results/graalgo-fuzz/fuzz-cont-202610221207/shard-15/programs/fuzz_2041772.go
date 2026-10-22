package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 uint; F2 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS11 struct { F0 uint; F1 int64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint; F1 uint }
func genFn0(p0 string, p1 int64) int64 {
    p0 = (p0 + (p0 + ""))
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            p1 += p1
            if (102 != -215) {
                var v6 int8 = 42
                _ = v6
            }
            p0 = ((p0 + p0) + "a")
            _ = fi5
        }
        var v7 int64 = (p1 | p1)
        _ = fi4
        _ = v7
    }
    v8 := GenS1{F0: "hello", F1: 1, F2: 32767}
    _ = p0
    _ = p1
    _ = v8
    return 42
}
func genFn1(p0 uint32, p1 int) uint16 {
    if ((p0 << ((1) & 63)) != p0) {
        if false {
            for fi9 := 0; fi9 < 14; fi9++ {
                fmt.Println(p0, p0)
                if (2147483647 < -259) {
                    p1 = fi9
                    for fi10 := 0; fi10 < 19; fi10++ {
                        v11 := GenS2{F0: 2, F1: p0, F2: 399}
                        if false {
                            v12 := GenS3{F0: 573, F1: false}
                            v13 := GenS4{F0: 42}
                            p1 -= (fi9 * p1)
                            _ = v12
                            _ = v13
                        }
                        switch 9223372036854775807 {
                        case 128:
                            switch ("foo bar" + "hello") {
                            case "\n\t":
                                var v14 int16 = 1
                                _ = v14
                            case "a":
                                v15 := []int{4294967295, (fi10 >> ((4294967295) & 63)), fi9}
                                _ = v15
                            }
                        case 746:
                            if (!true) {
                                var v16 int = fi9
                                if false {
                                    var v17 int8 = 0
                                    v18 := GenS2{F0: 255, F1: p0, F2: 2}
                                    _ = v17
                                    _ = v18
                                } else {
                                    p1 = 32768
                                    p1 = 2147483647
                                    p1 -= 32768
                                }
                                _ = v16
                            } else {
                                var v19 float32 = 0.1
                                _ = v19
                            }
                        default:
                            v20 := []int{}
                            _ = v20
                        }
                        _ = fi10
                        _ = v11
                    }
                }
                _ = fi9
            }
        } else {
            if (4294967295 < 4294967295) {
                switch "\n\t" {
                case "foo bar":
                    v21 := GenS2{F0: 27, F1: 979, F2: 594}
                    _ = v21
                default:
                    var v22 float32 = 0.0
                    _ = v22
                }
                if (!false) {
                    var v23 int32 = 120
                    for fi24 := 0; fi24 < 1; fi24++ {
                        var v25 bool = true
                        fmt.Println(p1)
                        v26 := []int{p1, (^(p1)), (^(p1))}
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    _ = v23
                } else {
                    v27 := GenS3{F0: 4294967295, F1: (0 == 42)}
                    for fi28 := 0; fi28 < 5; fi28++ {
                        var v29 int64 = 0
                        _ = fi28
                        _ = v29
                    }
                    p0 = p0
                    _ = v27
                }
            }
        }
        p0 = p0
    } else {
        v30 := GenS5{F0: 2147483647, F1: 255, F2: 295}
        _ = v30
    }
    if false {
        for fi31 := 0; fi31 < 15; fi31++ {
            v32 := []int{634, fi31}
            for fi33 := 0; fi33 < 20; fi33++ {
                if len(v32) > 0 {
                    v32[0] = p1
                }
                var v34 bool = (((-126 > -491) && (3.14159 >= ((3.14159 - 0.1) - 3.14159))) && false)
                _ = fi33
                _ = v34
            }
            _ = fi31
            _ = v32
        }
        var v35 uint = 2
        v35 -= (v35 & v35)
        _ = v35
    }
    p0 -= p0
    p1 -= (p1 - 305)
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 string = ""
    var v2 uint32 = 2147483647
    v3 := make(map[string]bool)
    if (1e10 < ((1.5 / (-2.25 - 3.14159)) / 0.1)) {
        if (!(!(!false))) {
            fmt.Println(v1, v1)
            if (!true) {
                v36 := GenS3{F0: 390, F1: (32768 <= 1)}
                for fi37 := 0; fi37 < 19; fi37++ {
                    v38 := GenS6{F0: v1}
                    fmt.Println(v1, v1, fi37)
                    for fi39 := 0; fi39 < 3; fi39++ {
                        v2 -= v2
                        v40 := []int{(fi37 << ((0) & 63)), fi39, fi39, fi39}
                        delete(v3, v1)
                        if ((!true) || ((genFn1((v2 | v2), -9223372036854775808) != 32767) && true)) {
                            var v41 uint8 = 255
                            if (true && false) {
                                v41 = (v41 >> ((255) & 63))
                            } else {
                                delete(v3, "\n\t")
                            }
                            _ = v41
                        }
                        v42 := GenS7{F0: -32768, F1: 82}
                        _ = fi39
                        _ = v40
                        _ = v42
                    }
                    _ = fi37
                    _ = v38
                }
                _ = v36
            }
            for fi43 := 0; fi43 < 17; fi43++ {
                for fi44 := 0; fi44 < 7; fi44++ {
                    v45 := []int{(fi44 * fi44)}
                    var v46 float32 = ((1.5 / 1e10) * 0.1)
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                fmt.Println(fi43, v2, v2)
                _ = fi43
            }
        } else {
            fmt.Println(v1, v2)
            v1 = ("foo bar" + "a")
            if (true && (0 <= -248)) {
                v2 = 2
                for fi47 := 0; fi47 < 3; fi47++ {
                    v1 = ("\n\t" + v1)
                    if false {
                        var v48 int16 = 0
                        _ = v48
                    }
                    var v49 uint = 65535
                    _ = fi47
                    _ = v49
                }
                var v50 int32 = 608
                _ = v50
            } else {
                v51 := GenS1{F0: ((v1 + (v1 + (v1 + v1))) + v1), F1: -128, F2: 0}
                switch "foo bar" {
                case "a":
                    v52 := GenS2{F0: 255, F1: v2, F2: 2}
                    _ = v52
                case "hello":
                    v2 = (^(v2))
                case "":
                    v1 = ("" + (v1 + v1))
                }
                for fi53 := 0; fi53 < 13; fi53++ {
                    var v54 float64 = (float64(-2.25) - 3.14159)
                    if true {
                        if true {
                            v55 := []int{(fi53 - (^(fi53))), fi53, (fi53 / ((fi53) | 1))}
                            _ = v55
                        } else {
                            v54 -= (1e10 + 0.1)
                            for fi56 := 0; fi56 < 13; fi56++ {
                                var v57 int16 = -756
                                if true {
                                    v58 := GenS7{F0: (v57 >> ((2) & 63)), F1: 42}
                                    var v59 string = "hello"
                                    fmt.Println(v59)
                                    _ = v58
                                    _ = v59
                                }
                                fmt.Println(fi53, v2)
                                _ = fi56
                                _ = v57
                            }
                        }
                        var v60 uint32 = v2
                        _ = v60
                    }
                    _ = fi53
                    _ = v54
                }
                _ = v51
            }
        }
        var v61 int16 = 42
        _ = v61
    }
    fmt.Println(v1, v1)
    v62 := GenS8{F0: 422}
    var v63 uint16 = genFn1(v2, 267)
    v64 := GenS4{F0: 1}
    v65 := v3[("foo bar" + "hello")]
    v66 := GenS2{F0: 127, F1: v2, F2: 51}
    var v67 int32 = 42
    var v68 string = "foo bar"
    fmt.Println(v68)
    fmt.Println(v67, v63)
    switch "a" {
    case "\n\t":
        var v69 int16 = -32768
        _ = v69
    }
    _ = len(v3)
    if v65 {
        switch -689 {
        case 127:
            v70 := []int{824}
            _ = v70
        default:
            v63 += (v63 + v63)
        }
        if (32768 == 32767) {
            var v71 string = ""
            v72 := v3[v71]
            _ = v71
            _ = v72
        } else {
            for fi73 := 0; fi73 < 17; fi73++ {
                for fi74 := 0; fi74 < 2; fi74++ {
                    v1 = "foo bar"
                    _ = fi74
                }
                for fi75 := 0; fi75 < 9; fi75++ {
                    if (v68 > (v68 + v68)) {
                        if false {
                            var v76 string = (v1 + ((v68 + (v68 + "foo bar")) + v1))
                            v77 := GenS9{F0: fi75}
                            _ = v76
                            _ = v77
                        }
                        fmt.Println(fi73, v68)
                    }
                    switch (^(fi75)) {
                    case 128:
                        v63 += genFn1((v2 - (^(v2))), 32767)
                    case -461:
                        switch 633 {
                        case 127:
                            if v65 {
                                var v78 int8 = 100
                                v79 := GenS10{F0: (^(v78)), F1: 127, F2: 223}
                                v80 := []int{fi73}
                                _ = v78
                                _ = v79
                                _ = v80
                            } else {
                                v81 := GenS11{F0: 1, F1: int64(fi73)}
                                _ = v81
                            }
                        case 128:
                            var v82 uint64 = 4294967295
                            _ = v82
                        }
                    default:
                        switch fi75 {
                        case 32768:
                            v3[v1] = (!((4294967295 > 4294967295) || (42 != -28)))
                        }
                    }
                    v83 := GenS12{F0: 843}
                    _ = fi75
                    _ = v83
                }
                fmt.Println(v63, v68)
                _ = fi73
            }
            fmt.Println(v68, v67)
        }
        if (0.0 < 0.0) {
            var v84 uint8 = 1
            v84 = (v84 | v84)
            _ = v84
        }
    }
    if v65 {
        v85 := []int{68, 932}
        for fi86 := 0; fi86 < 18; fi86++ {
            v87 := GenS11{F0: 1, F1: 1}
            var v88 uint8 = 210
            v89 := GenS4{F0: (v67 - 42)}
            _ = fi86
            _ = v87
            _ = v88
            _ = v89
        }
        delete(v3, v1)
        _ = v85
    } else {
        if v65 {
            for fi90 := 0; fi90 < 8; fi90++ {
                var v91 int8 = 42
                _ = fi90
                _ = v91
            }
            switch (v68 + v68) {
            case "\n\t":
                _ = len(v3)
            case "a":
                v92 := []int{}
                _ = v92
            case "foo bar":
                var v93 uint64 = 42
                _ = v93
            }
        }
    }
    _ = len(v3)
    v94 := GenS13{F0: 18446744073709551615, F1: 255}
    switch "" {
    case "a":
        for fi95 := 0; fi95 < 14; fi95++ {
            v63 = (v63 % ((v63) | 1))
            var v96 int64 = -217
            _ = fi95
            _ = v96
        }
    case "\n\t":
        var v97 float32 = ((1.5 / (3.14159 * 1e10)) * (1e10 - -2.25))
        _ = v97
    case "":
        switch ((v68 + (("foo bar" + v1) + (("\n\t" + (v1 + "\n\t")) + (v68 + v68)))) + v1) {
        case "":
            var v98 int64 = 1
            _ = v98
        case "foo bar":
            v99 := []int{2147483647}
            _ = v99
        default:
            var v100 uint = 271
            _ = v100
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v62
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v94
}
