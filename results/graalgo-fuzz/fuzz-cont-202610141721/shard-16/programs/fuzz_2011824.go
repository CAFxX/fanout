package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int64; F1 int; F2 string }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int }
func genFn0(p0 uint8) int {
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = p0
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: true, F1: 1}
        _ = fi4
        _ = v5
    }
    _ = p0
    return 128
}
func genFn1() int16 {
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int8 = 0
            var v10 bool = (1.5 > float64(0.0))
            v11 := []int{fi8, fi8}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            v6 = (^(v6))
            _ = fi12
        }
        _ = fi7
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        if (false && (0 >= 128)) {
            var v14 float32 = (3.14159 / 0.1)
            if true {
                var v15 string = (("\n\t" + "") + (("\n\t" + "") + (((("hello" + "foo bar") + "") + "foo bar") + "\n\t")))
                var v16 uint = 0
                _ = v15
                _ = v16
            }
            _ = v14
        }
        v6 -= v6
        v17 := GenS2{F0: 4294967295, F1: false}
        _ = fi13
        _ = v17
    }
    if true {
        v18 := []int{-9223372036854775808}
        _ = v18
    } else {
        v19 := []int{42, -1, 606}
        if ((("a" + "a") > "\n\t") || false) {
            var v20 uint32 = 500
            var v21 float32 = 1.5
            _ = v20
            _ = v21
        } else {
            switch -685 {
            case 42:
                var v22 uint8 = 149
                _ = v22
            case 4294967295:
                if false {
                    if (18446744073709551615 >= 616) {
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        v6 = (v6 * v6)
                        v6 = 4294967295
                    } else {
                        v23 := GenS1{F0: (1 >= 465), F1: 127}
                        v6 = (^(v6))
                        v6 -= (v6 - (v6 | (v6 - 590)))
                        _ = v23
                    }
                } else {
                    v6 += v6
                }
            case 128:
                fmt.Println(v6)
            default:
                v24 := GenS3{F0: 658, F1: 128, F2: "a"}
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 uint8 = 127
            _ = fi25
            _ = v26
        }
        _ = v19
    }
    _ = v6
    return 340
}

func main() {
    var v1 int = 839
    var v2 uint = 712
    var v3 int16 = 42
    var v27 int32 = 700
    switch v1 {
    case -630:
        v28 := []int{(v1 << (((v2 & v2)) & 63))}
        _ = v28
    }
    v1 = (v1 + (^(v1)))
    v29 := GenS4{F0: 4294967295, F1: true}
    if true {
        v30 := []int{-1, (v1 ^ (v1 - (v1 ^ v1))), (v1 >> (((v2 << (((v2 / ((v2) | 1))) & 63))) & 63)), v1}
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
        var v31 uint = v2
        _ = v30
        _ = v31
    } else {
        v3 += 611
        fmt.Println(v1)
        if (!true) {
            v1 += -1
            var v32 uint16 = 0
            switch genFn0(42) {
            case 53:
                fmt.Println(v1, v2, v1)
            }
            _ = v32
        }
    }
    if false {
        var v33 int16 = genFn1()
        if false {
            var v34 uint = (^(v2))
            var v35 int64 = 0
            _ = v34
            _ = v35
        }
        _ = v33
    } else {
        if ((1 < int8(v2)) && false) {
            if false {
                fmt.Println(v1, v1)
                switch "a" {
                case "foo bar":
                    v36 := GenS1{F0: (906 >= 186), F1: -50}
                    _ = v36
                case "":
                    v37 := []int{v1, (v1 & (v1 - v1)), 4294967295, (v1 << ((1) & 63))}
                    _ = v37
                }
                if ("" > "foo bar") {
                    for fi38 := 0; fi38 < 19; fi38++ {
                        if ((v2 & (v2 ^ v2)) != (v2 & v2)) {
                            v39 := []int{(^(v1))}
                            var v40 int8 = int8(v3)
                            for fi41 := 0; fi41 < 7; fi41++ {
                                _ = len(v39)
                                _ = fi41
                            }
                            _ = v39
                            _ = v40
                        }
                        v27 += 0
                        _ = fi38
                    }
                    v42 := []int{}
                    for fi43 := 0; fi43 < 7; fi43++ {
                        if len(v42) >= 2 {
                            v42 = v42[1:len(v42)]
                        }
                        var v44 int8 = 0
                        var v45 uint64 = 0
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    _ = v42
                }
            }
        } else {
            for fi46 := 0; fi46 < 10; fi46++ {
                var v47 bool = false
                v3 += v3
                _ = fi46
                _ = v47
            }
            fmt.Println(v1, v2)
        }
        for fi48 := 0; fi48 < 10; fi48++ {
            fmt.Println(v2, v2, fi48)
            if (true && true) {
                switch 42 {
                case 42:
                    v1 -= (fi48 * fi48)
                case 32767:
                    if true {
                        v1 = v1
                        v49 := []int{(v1 / ((v1) | 1)), 563, v1, (v1 / (((v1 * 1)) | 1))}
                        if len(v49) > 0 {
                            v49[0] = (fi48 & -211)
                        }
                        _ = v49
                    }
                case 4294967295:
                    var v50 string = "foo bar"
                    _ = v50
                }
            } else {
                switch "\n\t" {
                case "a":
                    for fi51 := 0; fi51 < 5; fi51++ {
                        switch fi48 {
                        case 32768:
                            var v52 int64 = -1
                            _ = v52
                        case -9223372036854775808:
                            var v53 int32 = (v27 ^ v27)
                            _ = v53
                        case 128:
                            switch 127 {
                            case 1:
                                v3 -= v3
                            case 127:
                                fmt.Println(v2, v27, v2)
                            case -751:
                                v54 := []int{(-(fi51)), v1}
                                _ = v54
                            }
                        default:
                            if (!true) {
                                switch "foo bar" {
                                case "":
                                    v2 = (v2 << (((^(v2))) & 63))
                                case "a":
                                    v55 := GenS2{F0: 364, F1: (false || false)}
                                    _ = v55
                                default:
                                    v3 = v3
                                }
                                v56 := GenS4{F0: 2, F1: (!false)}
                                var v57 int64 = int64(v2)
                                fmt.Println(fi48, v2, v1)
                                if false {
                                    v3 = (v3 + int16(v2))
                                    v3 += genFn1()
                                } else {
                                    v58 := []int{fi48, (fi48 - fi48)}
                                    v59 := GenS5{F0: (v1 / (((fi48 >> (((^(v2))) & 63))) | 1))}
                                    v3 -= (v3 >> ((v2) & 63))
                                    _ = v58
                                    _ = v59
                                }
                                v60 := GenS2{F0: 2, F1: false}
                                _ = v56
                                _ = v57
                                _ = v60
                            }
                        }
                        _ = fi51
                    }
                default:
                    var v61 float64 = 3.14159
                    _ = v61
                }
                switch "a" {
                case "foo bar":
                    v62 := GenS5{F0: fi48}
                    _ = v62
                case "\n\t":
                    var v63 int = (fi48 + v1)
                    _ = v63
                default:
                    switch "a" {
                    case "a":
                        v64 := []int{(fi48 + fi48), v1, (fi48 >> ((500) & 63))}
                        _ = v64
                    case "hello":
                        for fi65 := 0; fi65 < 14; fi65++ {
                            v66 := GenS2{F0: 2, F1: false}
                            _ = fi65
                            _ = v66
                        }
                    default:
                        fmt.Println(v3, v2)
                    }
                }
                var v67 uint16 = 440
                _ = v67
            }
            var v68 string = ("" + "foo bar")
            _ = fi48
            _ = v68
        }
        fmt.Println(v27, v2, v1)
    }
    var v69 string = "\n\t"
    if false {
        v27 += v27
    }
    if true {
        v70 := []int{v1, v1, genFn0(215)}
        v1 = (v1 - v1)
        _ = len(v70)
        _ = v70
    } else {
        v71 := []int{(-(v1)), 0}
        v27 = -512
        _ = v71
    }
    v72 := []int{v1, genFn0(180), v1, 42}
    var v73 float64 = 0.1
    v74 := GenS3{F0: int64(v1), F1: v1, F2: "a"}
    if len(v72) > 0 {
        v72[0] = (v1 >> (((v2 * v2)) & 63))
    }
    fmt.Println(v3, v3, v69)
    for fi75 := 0; fi75 < 1; fi75++ {
        if (((fi75 != (v1 >> ((v2) & 63))) && (0.0 != ((1.5 - ((-2.25 + (-2.25 + 3.14159)) * 3.14159)) * 3.14159))) && false) {
            switch v69 {
            case "\n\t":
                if len(v72) >= 2 {
                    v72 = v72[1:len(v72)]
                }
            case "":
                var v76 int = genFn0(127)
                _ = v76
            }
            v73 = float64(0.1)
        }
        v2 += 306
        var v77 uint32 = 1
        _ = fi75
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v69:", v69)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v69
    _ = v72
    _ = v73
    _ = v74
}
