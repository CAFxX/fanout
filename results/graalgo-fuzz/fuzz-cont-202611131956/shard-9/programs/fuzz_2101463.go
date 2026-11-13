package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 int32; F1 uint8 }
func genFn0() int8 {
    switch "hello" {
    case "a":
        var v4 uint64 = 648
        _ = v4
    default:
        for fi5 := 0; fi5 < 11; fi5++ {
            for fi6 := 0; fi6 < 4; fi6++ {
                v7 := GenS1{F0: 2147483647}
                v8 := []int{}
                if false {
                    var v9 uint16 = 268
                    v10 := GenS1{F0: 2}
                    if len(v8) > 0 {
                        v8[0] = fi6
                    }
                    _ = v10
                    _ = v9
                }
                _ = fi6
                _ = v7
                _ = v8
            }
            var v11 int8 = 42
            _ = fi5
            _ = v11
        }
    }
    if true {
        var v12 float32 = 1e10
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 float32 = 1.5
        _ = fi13
        _ = v14
    }
    return 0
}
func genFn1(p0 bool, p1 uint32) uint {
    if p0 {
        v15 := []int{-565, -9223372036854775808}
        v16 := GenS2{F0: 18446744073709551615, F1: ("a" + "a")}
        _ = v15
        _ = v16
    } else {
        v17 := []int{}
        _ = v17
    }
    p0 = (0 > 2147483647)
    p0 = p0
    _ = p0
    _ = p1
    return 360
}

func main() {
    v1 := make(map[string]int)
    var v2 string = "hello"
    var v3 uint8 = 4
    v2 = (v2 + v2)
    switch -9223372036854775808 {
    case 128:
        v18 := []int{-484, 42}
        _ = v18
    case -1:
        if true {
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := v1[v2]
                if ((-471 == 703) && (-494 < 9223372036854775807)) {
                    var v21 uint32 = 903
                    v22 := GenS1{F0: v21}
                    v23 := []int{fi19}
                    _ = v21
                    _ = v22
                    _ = v23
                }
                v24 := GenS1{F0: 0}
                _ = fi19
                _ = v20
                _ = v24
            }
            for fi25 := 0; fi25 < 9; fi25++ {
                if false {
                    v26 := GenS3{F0: 0}
                    v3 = v3
                    var v27 int8 = -1
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 8; fi28++ {
                    if true {
                        var v29 uint64 = 744
                        var v30 uint32 = 2147483647
                        _ = v29
                        _ = v30
                    } else {
                        if (943 != 544) {
                            var v31 int64 = -9223372036854775808
                            if false {
                                fmt.Println(fi25, v3)
                                if false {
                                    v31 = v31
                                } else {
                                    switch (fi25 & (fi25 << ((genFn1(true, 1)) & 63))) {
                                    case -1:
                                        var v32 float64 = 3.14159
                                        _ = v32
                                    default:
                                        v33 := GenS1{F0: 184}
                                        _ = v33
                                    }
                                }
                            } else {
                                v31 = v31
                            }
                            v2 = "a"
                            _ = v31
                        } else {
                            v34 := GenS4{F0: 1000, F1: 411}
                            v35 := GenS5{F0: 4294967295, F1: v3}
                            _ = v34
                            _ = v35
                        }
                        v36 := []int{}
                        _ = v36
                    }
                    if false {
                        if (42 <= 0) {
                            switch (v2 + "foo bar") {
                            case "a":
                                var v37 uint32 = 845
                                _ = v37
                            default:
                                if true {
                                    var v38 int32 = 1
                                    fmt.Println(v2, fi28)
                                    _ = v38
                                } else {
                                    fmt.Println(v2, v2)
                                    switch ("" + v2) {
                                    case "foo bar":
                                        var v39 uint32 = 1
                                        _ = v39
                                    case "hello":
                                        var v40 float32 = 3.14159
                                        _ = v40
                                    case "a":
                                        switch v2 {
                                        case "a":
                                            if true {
                                                v41 := GenS3{F0: -1}
                                                fmt.Println(v2, v3)
                                                v3 -= v3
                                                _ = v41
                                            } else {
                                                v42 := GenS2{F0: 18446744073709551615, F1: ("a" + "foo bar")}
                                                _ = v42
                                            }
                                        case "foo bar":
                                            var v43 float32 = -2.25
                                            _ = v43
                                        case "\n\t":
                                            fmt.Println(v2, v2)
                                        }
                                    default:
                                        var v44 float32 = (((1e10 - (1e10 * (((0.1 - 1.5) / float32(3.14159)) / 1.0))) / (3.14159 - 0.1)) + -2.25)
                                        _ = v44
                                    }
                                }
                            }
                            delete(v1, "foo bar")
                            if (("hello" + v2) < v2) {
                                v2 = (v2 + v2)
                                v45 := []int{fi28, -9223372036854775808, fi25}
                                v2 = ((("hello" + (v2 + v2)) + "a") + "\n\t")
                                _ = v45
                            } else {
                                v2 = (v2 + ((v2 + v2) + v2))
                                if true {
                                    v2 = v2
                                    v2 = v2
                                } else {
                                    fmt.Println(fi25, v2, v2)
                                    v2 = v2
                                    if false {
                                        v46 := []int{fi28, (fi28 | fi25), fi28, (fi25 + fi28)}
                                        _ = v46
                                    }
                                }
                                if (fi28 <= fi28) {
                                    v47 := []int{}
                                    v1[(v2 + "hello")] = (fi28 >> ((710) & 63))
                                    switch -136 {
                                    case -2147483648:
                                        v48 := GenS6{F0: 456}
                                        v49 := GenS2{F0: 42, F1: "hello"}
                                        v50 := GenS5{F0: 1, F1: (v3 % ((v3) | 1))}
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    }
                                    _ = v47
                                } else {
                                    v51 := GenS4{F0: 209, F1: 2147483647}
                                    if true {
                                        fmt.Println(v3, fi28, fi28)
                                        var v52 int32 = 477
                                        var v53 float64 = -2.25
                                        _ = v52
                                        _ = v53
                                    } else {
                                        if true {
                                            var v54 float64 = -2.25
                                            _ = v54
                                        } else {
                                            fmt.Println(fi25, v2, v2)
                                        }
                                        v3 = v3
                                        var v55 uint32 = 2147483647
                                        v56 := GenS3{F0: int64(v55)}
                                        var v57 int16 = 32767
                                        _ = v55
                                        _ = v56
                                        _ = v57
                                    }
                                    _ = v51
                                }
                            }
                        }
                    }
                    v2 = "a"
                    var v58 int32 = 483
                    if (1 >= 1) {
                        if false {
                            v3 -= 0
                            var v59 uint64 = 2
                            _ = v59
                        }
                    } else {
                        var v60 int16 = -32768
                        v58 -= v58
                        _ = v60
                    }
                    v61 := GenS6{F0: 4294967295}
                    _ = fi28
                    _ = v58
                    _ = v61
                }
                fmt.Println(v3, fi25)
                _ = fi25
            }
        }
    }
    v3 -= (v3 >> ((65535) & 63))
    var v62 int32 = 670
    fmt.Println(v2, v62)
    if (255 == 0) {
        fmt.Println(v62, v62, v2)
    } else {
        if ((true || false) || true) {
            var v63 uint16 = 65535
            _ = v63
        } else {
            for fi64 := 0; fi64 < 18; fi64++ {
                v3 = 194
                v65 := GenS3{F0: 1}
                _ = fi64
                _ = v65
            }
        }
    }
    for fi66 := 0; fi66 < 20; fi66++ {
        if true {
            v2 = ""
            v2 = v2
        } else {
            for fi67 := 0; fi67 < 19; fi67++ {
                _ = len(v1)
                v3 += 220
                _ = fi67
            }
            v62 = v62
        }
        v62 += (v62 * int32(v3))
        _ = fi66
    }
    for fi68 := 0; fi68 < 16; fi68++ {
        for fi69 := 0; fi69 < 7; fi69++ {
            v2 = v2
            v70 := []int{fi69, (^(fi68)), (fi68 * (fi69 - fi68)), fi69}
            if true {
                if ((!((genFn1(false, 65535) != 0) && false)) || (-1 != 127)) {
                    if len(v70) >= 2 {
                        v70 = v70[1:len(v70)]
                    }
                    v71 := GenS7{F0: 4294967295, F1: 0}
                    var v72 uint = 667
                    _ = v71
                    _ = v72
                }
                var v73 uint = 65535
                var v74 int16 = -32768
                _ = v73
                _ = v74
            } else {
                if true {
                    fmt.Println(fi69, v62)
                } else {
                    var v75 int64 = 459
                    _ = v75
                }
                if true {
                    fmt.Println(v62, v2, fi69)
                    var v76 float32 = 1.5
                    var v77 uint32 = 233
                    _ = v76
                    _ = v77
                } else {
                    v2 = ""
                    var v78 int16 = -939
                    var v79 float64 = -2.25
                    v3 = 42
                    _ = v78
                    _ = v79
                }
                v80 := v1[(v2 + (v2 + (((((v2 + v2) + "") + v2) + v2) + v2)))]
                v81 := GenS8{F0: 42, F1: v80}
                _ = v80
                _ = v81
            }
            fmt.Println(v62, fi69, v62)
            _ = fi69
            _ = v70
        }
        var v82 int64 = -1
        var v83 float64 = 3.14159
        _ = fi68
        _ = v82
        _ = v83
    }
    if true {
        switch ("a" + v2) {
        case "":
            delete(v1, "foo bar")
        }
    }
    v2 = ("" + ((v2 + (("\n\t" + ("" + (v2 + (v2 + v2)))) + v2)) + ((v2 + ((((v2 + v2) + ("a" + "\n\t")) + "hello") + v2)) + "\n\t")))
    v84 := []int{int(v3)}
    if (!(!(0.1 < float64(((3.14159 / 1.0) - 1.5))))) {
        for fi85 := 0; fi85 < 14; fi85++ {
            for fi86 := 0; fi86 < 17; fi86++ {
                switch (v2 + v2) {
                case "":
                    v87 := GenS9{F0: (v62 ^ (v62 << ((4294967295) & 63))), F1: v3}
                    _ = v87
                case "hello":
                    v88 := v1[v2]
                    _ = v88
                case "foo bar":
                    switch v2 {
                    case "":
                        if len(v84) > 0 {
                            v84[0] = fi86
                        }
                    case "a":
                        v2 = "hello"
                    default:
                        var v89 bool = false
                        _ = v89
                    }
                default:
                    v90 := GenS5{F0: 65535, F1: (v3 * v3)}
                    _ = v90
                }
                var v91 uint8 = 1
                if true {
                    if true {
                        if true {
                            var v92 int32 = (v62 << ((42) & 63))
                            _ = v92
                        } else {
                            var v93 int8 = -128
                            switch (fi86 >> ((1) & 63)) {
                            case -275:
                                v93 = (-(v93))
                            default:
                                v93 += genFn0()
                            }
                            var v94 uint16 = 65535
                            _ = v93
                            _ = v94
                        }
                        fmt.Println(v62, v3, fi85)
                    }
                    v95 := GenS3{F0: int64(fi85)}
                    if true {
                        if ((v62 | v62) < v62) {
                            switch 2147483647 {
                            case 9223372036854775807:
                                var v96 int8 = genFn0()
                                _ = v96
                            case 4294967295:
                                if true {
                                    v62 += (v62 - v62)
                                    fmt.Println(v91, v3, v62)
                                } else {
                                    v1[v2] = 127
                                }
                            case -478:
                                var v97 int8 = 102
                                _ = v97
                            }
                        }
                    } else {
                        var v98 uint64 = 2
                        _ = v98
                    }
                    _ = v95
                }
                _ = fi86
                _ = v91
            }
            for fi99 := 0; fi99 < 11; fi99++ {
                fmt.Println(fi99, fi99, fi99)
                var v100 bool = (false && false)
                _ = fi99
                _ = v100
            }
            _ = fi85
        }
    }
    for fi101 := 0; fi101 < 18; fi101++ {
        v3 -= 127
        _ = fi101
    }
    fmt.Println(v2, v62, v2)
    fmt.Println(v2)
    var v102 int = -981
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v84:", len(v84), v84)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v62
    _ = v84
}
