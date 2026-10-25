package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 string }
type GenS2 struct { F0 int; F1 int; F2 uint8 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint, p2 string) int16 {
    var v5 int = 0
    v6 := []int{v5}
    v5 = v5
    var v7 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 uint = 1
    var v2 float32 = float32(3.14159)
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    var v8 uint = (^(v1))
    fmt.Println(v2)
    fmt.Println(v2)
    if (((true && (v2 < (v2 * (1.5 - v2)))) && (true && true)) && false) {
        var v9 bool = false
        _ = v9
    } else {
        if false {
            var v10 int8 = 126
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 uint16 = 42
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    }
    switch 32768 {
    case -9223372036854775808:
        v13 := GenS1{F0: (("hello" + "a") + "\n\t"), F1: (1 <= 385), F2: "hello"}
        _ = v13
    case 4294967295:
        var v14 int64 = 0
        _ = v14
    case -964:
        fmt.Println(v2, v8, v8)
    default:
        switch -213 {
        case 9223372036854775807:
            var v15 uint64 = 840
            _ = v15
        default:
            var v16 int16 = -1
            _ = v16
        }
    }
    fmt.Println(v1)
    for fi17 := 0; fi17 < 3; fi17++ {
        if true {
            v8 = v8
            v18 := GenS2{F0: fi17, F1: 9223372036854775807, F2: 255}
            _ = v18
        }
        _ = fi17
    }
    var v19 uint64 = 1
    switch ("foo bar" + (("\n\t" + "foo bar") + "hello")) {
    case "\n\t":
        v1 = 1
    case "hello":
        fmt.Println(v1, v19)
    case "":
        fmt.Println(v8, v1, v8)
    }
    v20 := []int{-556, -742}
    var v21 int64 = -9223372036854775808
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 uint64 = v19
        _ = fi22
        _ = v23
    }
    v1 = (^(v1))
    if false {
        fmt.Println(v1, v8)
        v20 = append(v20, 9223372036854775807)
    }
    var v24 float64 = -2.25
    if true {
        var v25 uint = v1
        _ = v25
    } else {
        if (!(-9223372036854775808 >= 1)) {
            for fi26 := 0; fi26 < 18; fi26++ {
                v1 = (v8 - (v1 | 2))
                var v27 uint64 = (v19 | 840)
                _ = fi26
                _ = v27
            }
            if true {
                fmt.Println(v21, v2, v24)
                var v28 uint8 = 128
                _ = v28
            } else {
                switch 720 {
                case 1:
                    v1 = (v1 - (^(v1)))
                case 32768:
                    var v29 uint64 = (v19 >> (((v8 / (((v1 & (v8 / (((v1 % (((v8 + v1)) | 1))) | 1)))) | 1))) & 63))
                    _ = v29
                }
            }
            var v30 uint16 = 2
            _ = v30
        }
        switch -606 {
        case 4294967295:
            var v31 uint32 = 4294967295
            _ = v31
        case 127:
            for fi32 := 0; fi32 < 19; fi32++ {
                if false {
                    switch "" {
                    case "foo bar":
                        for fi33 := 0; fi33 < 9; fi33++ {
                            v8 = 171
                            _ = fi33
                        }
                    case "a":
                        v34 := GenS3{F0: 0, F1: ("\n\t" + "hello")}
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 17; fi35++ {
                        v36 := GenS4{F0: fi32, F1: 65535}
                        v37 := GenS4{F0: (fi32 << (((v8 * v1)) & 63)), F1: 2147483647}
                        v38 := GenS2{F0: fi32, F1: (-(fi35)), F2: 2}
                        fmt.Println(v1, v19, v1)
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    v21 = v21
                }
                if true {
                    v2 += 0.0
                    var v39 int32 = 2147483647
                    v40 := GenS2{F0: fi32, F1: fi32, F2: 255}
                    _ = v39
                    _ = v40
                } else {
                    v19 += v19
                    if false {
                        switch "foo bar" {
                        case "":
                            v1 += v1
                        case "hello":
                            if (!(!true)) {
                                v41 := GenS5{F0: 42}
                                for fi42 := 0; fi42 < 10; fi42++ {
                                    var v43 int64 = -9223372036854775808
                                    v44 := GenS2{F0: -9223372036854775808, F1: fi32, F2: 255}
                                    _ = fi42
                                    _ = v43
                                    _ = v44
                                }
                                _ = v41
                            }
                        default:
                            if (0.0 < (v24 - v24)) {
                                if true {
                                    for fi45 := 0; fi45 < 4; fi45++ {
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                        v2 -= (float32((v24 + v24)) + v2)
                                        fmt.Println(v8, v21, v1)
                                        _ = fi45
                                    }
                                    switch "a" {
                                    case "":
                                        var v46 uint = 78
                                        _ = v46
                                    case "a":
                                        v47 := GenS5{F0: 889}
                                        _ = v47
                                    }
                                } else {
                                    for fi48 := 0; fi48 < 5; fi48++ {
                                        if true {
                                            if false {
                                                v49 := GenS1{F0: "", F1: ("" != "foo bar"), F2: "a"}
                                                if (!(!false)) {
                                                    v50 := GenS2{F0: fi48, F1: fi48, F2: 1}
                                                    var v51 uint32 = 2
                                                    _ = v50
                                                    _ = v51
                                                }
                                                _ = v49
                                            } else {
                                                var v52 int16 = 32767
                                                var v53 float64 = (v24 * (0.1 - (1.5 - v24)))
                                                if ((!true) && ("a" < ("" + ""))) {
                                                    switch fi32 {
                                                    case 1:
                                                        if len(v20) >= 2 {
                                                            v20 = v20[1:len(v20)]
                                                        }
                                                    default:
                                                        var v54 uint8 = 127
                                                        _ = v54
                                                    }
                                                    v20 = append(v20, fi32)
                                                } else {
                                                    _ = len(v20)
                                                    var v55 float32 = (1.5 + (float32(-2.25) - (3.14159 * v2)))
                                                    v53 = (v24 / (1e10 * v24))
                                                    _ = v55
                                                }
                                                var v56 uint32 = 590
                                                var v57 uint = (v1 - (v1 * (v8 << ((42) & 63))))
                                                _ = v52
                                                _ = v53
                                                _ = v56
                                                _ = v57
                                            }
                                            if false {
                                                var v58 int = (fi32 + (fi48 * (fi32 * fi32)))
                                                _ = v58
                                            } else {
                                                v21 = v21
                                                fmt.Println(v19, v24, v8)
                                            }
                                            fmt.Println(v21, v2, v24)
                                        }
                                        v21 = v21
                                        _ = fi48
                                    }
                                    v59 := GenS4{F0: -935, F1: 2}
                                    var v60 float32 = (v2 / v2)
                                    _ = v59
                                    _ = v60
                                }
                                fmt.Println(v8, v8, v2)
                            } else {
                                v21 += v21
                                var v61 int8 = 0
                                _ = v61
                            }
                        }
                    }
                    switch fi32 {
                    case 9223372036854775807:
                        v1 -= (v1 >> ((v1) & 63))
                    case 42:
                        fmt.Println(v21, fi32, v1)
                    default:
                        v8 -= 42
                    }
                }
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                _ = fi32
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v8
}
