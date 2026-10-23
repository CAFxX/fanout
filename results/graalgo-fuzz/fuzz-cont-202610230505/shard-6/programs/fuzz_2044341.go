package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 bool; F1 string; F2 int64 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 bool }
func genFn0(p0 bool, p1 int, p2 uint16) int64 {
    var v6 int = (p1 * (p1 ^ p1))
    v7 := GenS1{F0: 0, F1: 128}
    p1 = -9223372036854775808
    v8 := []int{}
    v9 := GenS1{F0: 1, F1: 252}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() int32 {
    v10 := GenS2{F0: false, F1: "a", F2: 1}
    v11 := []int{42, -684}
    var v12 int = -743
    if (true || false) {
        _ = len(v11)
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        _ = len(v11)
        v14 := GenS2{F0: false, F1: "a", F2: genFn0(false, (v12 | v12), 42)}
        _ = fi13
        _ = v14
    }
    _ = v10
    _ = v11
    _ = v12
    return 516
}

func main() {
    var v1 float32 = (0.0 + -2.25)
    v2 := make(map[uint32]int)
    var v3 string = ""
    var v4 uint64 = 42
    v5 := make(map[string]bool)
    if (719 < 860) {
        var v15 uint64 = 561
        _ = v15
    } else {
        delete(v5, (v3 + v3))
        v3 = "a"
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        if true {
            v17 := GenS3{F0: 0, F1: 0}
            for fi18 := 0; fi18 < 14; fi18++ {
                v3 = ""
                fmt.Println(v1)
                v4 += v4
                v19 := GenS4{F0: v4, F1: (-2.25 < (0.0 / (((1.5 * float32((3.14159 - 0.0))) * (float32((((-2.25 / -2.25) + 3.14159) * 3.14159)) - v1)) * 1.5)))}
                var v20 int16 = -32768
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 int64 = -159
            _ = v17
            _ = v21
        } else {
            if (true && ((^(v4)) < (^(v4)))) {
                var v22 uint16 = 0
                v3 = v3
                _ = v22
            }
            var v23 float64 = 1e10
            _ = v23
        }
        _ = fi16
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 uint16 = 255
        if true {
            var v26 string = v3
            var v27 int8 = 94
            v27 = v27
            _ = v26
            _ = v27
        }
        if true {
            if true {
                for fi28 := 0; fi28 < 14; fi28++ {
                    v25 = 42
                    _ = fi28
                }
                if true {
                    var v29 uint16 = (^(v25))
                    var v30 int16 = 0
                    switch fi24 {
                    case -661:
                        for fi31 := 0; fi31 < 1; fi31++ {
                            v25 = (v25 << ((0) & 63))
                            _ = fi31
                        }
                    case 4294967295:
                        v32 := []int{fi24, fi24}
                        _ = v32
                    case -9223372036854775808:
                        v33 := GenS1{F0: int64(v4), F1: 127}
                        _ = v33
                    }
                    _ = v29
                    _ = v30
                } else {
                    fmt.Println(v4, v25)
                    switch fi24 {
                    case -613:
                        var v34 int64 = genFn0(true, (fi24 / (((fi24 << ((402) & 63))) | 1)), 727)
                        _ = v34
                    case 969:
                        v25 = 255
                    case 4294967295:
                        if true {
                            v35 := []int{fi24}
                            v1 = v1
                            if (0 >= 9223372036854775807) {
                                var v36 float64 = (0.1 * float64(1.5))
                                switch ("\n\t" + v3) {
                                case "a":
                                    v25 = v25
                                default:
                                    v37 := GenS2{F0: ((false && false) || false), F1: (v3 + v3), F2: 0}
                                    _ = v37
                                }
                                _ = v36
                            }
                            _ = v35
                        } else {
                            if (("hello" + "") == v3) {
                                fmt.Println(fi24, v4, fi24)
                            } else {
                                v25 -= 255
                                switch (fi24 + 127) {
                                case 1:
                                    v38 := GenS3{F0: 0, F1: 2}
                                    _ = v38
                                case 32768:
                                    v39 := []int{(-(fi24)), (fi24 - (fi24 / ((fi24) | 1)))}
                                    _ = v39
                                default:
                                    for fi40 := 0; fi40 < 13; fi40++ {
                                        var v41 float32 = v1
                                        fmt.Println(fi24, fi40, v25)
                                        _ = fi40
                                        _ = v41
                                    }
                                }
                                if (!(-219 > -2147483648)) {
                                    switch fi24 {
                                    case 2147483647:
                                        if (!(true || (v4 >= (v4 >> ((0) & 63))))) {
                                            for fi42 := 0; fi42 < 4; fi42++ {
                                                v1 += (v1 * ((-2.25 / v1) + (float32(float64((0.1 / ((3.14159 / v1) / (-2.25 - v1))))) / (-2.25 / 1e10))))
                                                _ = fi42
                                            }
                                            v4 -= 4294967295
                                        }
                                    default:
                                        var v43 uint32 = 2147483647
                                        _ = v43
                                    }
                                }
                            }
                            v44 := []int{fi24, (fi24 ^ 1), 2147483647}
                            fmt.Println(v4)
                            _ = v44
                        }
                    default:
                        v45 := v2[65535]
                        _ = v45
                    }
                }
                if false {
                    fmt.Println(v1, v4)
                    switch -433 {
                    case 2147483647:
                        switch v3 {
                        case "a":
                            v1 = float32((1e10 + -2.25))
                        case "\n\t":
                            fmt.Println(v25)
                        }
                    case 471:
                        v2[42] = fi24
                    case 127:
                        v1 = (float32(1e10) * -2.25)
                    }
                    v3 = (v3 + v3)
                } else {
                    v25 += (v25 | v25)
                    v4 = v4
                }
            }
            var v46 uint16 = (^(v25))
            var v47 int = 32768
            _ = v46
            _ = v47
        } else {
            switch ("foo bar" + (v3 + (("a" + "\n\t") + v3))) {
            case "\n\t":
                var v48 int32 = -1
                _ = v48
            case "a":
                fmt.Println(fi24, v4, v1)
            case "":
                _ = len(v5)
            }
            v49 := GenS5{F0: (!false)}
            _ = v49
        }
        _ = fi24
        _ = v25
    }
    if (v3 <= "") {
        v50 := v2[169]
        _ = v50
    } else {
        var v51 uint16 = 255
        _ = v51
    }
    v52 := []int{-563}
    v3 = "hello"
    var v53 uint16 = 255
    fmt.Println(v53, v3, v3)
    if len(v52) >= 2 {
        v52 = v52[1:len(v52)]
    }
    v54 := GenS4{F0: 146, F1: false}
    var v55 uint8 = 240
    var v56 uint8 = (v55 / (((^(v55))) | 1))
    v57 := GenS3{F0: -2147483648, F1: 7}
    if len(v52) >= 2 {
        v52 = v52[1:len(v52)]
    }
    var v58 int16 = 0
    if (!true) {
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
