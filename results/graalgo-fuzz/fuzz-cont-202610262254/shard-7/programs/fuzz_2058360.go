package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int8; F1 int; F2 string }
type GenS3 struct { F0 int8; F1 int16; F2 bool }
type GenS4 struct { F0 int64; F1 int; F2 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 int64; F1 int64; F2 int }
func genFn0() uint {
    switch "hello" {
    case "":
        var v4 int8 = 127
        _ = v4
    case "a":
        var v5 uint = 42
        _ = v5
    case "\n\t":
        v6 := GenS1{F0: 42, F1: 65535}
        _ = v6
    }
    if true {
        v7 := GenS2{F0: 1, F1: 1, F2: ("a" + ("foo bar" + "\n\t"))}
        _ = v7
    } else {
        if (false && (true || true)) {
            var v8 uint16 = 1
            var v9 bool = false
            v8 -= v8
            _ = v8
            _ = v9
        } else {
            if (true && false) {
                var v10 float32 = 0.0
                _ = v10
            } else {
                var v11 float32 = 0.0
                v11 = v11
                if false {
                    var v12 int64 = -9223372036854775808
                    _ = v12
                } else {
                    v11 = v11
                    switch "foo bar" {
                    case "\n\t":
                        v11 = float32((1.5 + (((float64((((1e10 / ((float32(1.5) / v11) + 1e10)) - v11) + v11)) - float64(0.0)) - (((((1.5 - 1e10) + 0.1) * 1e10) + 1e10) + 1.5)) + (1e10 + 0.0))))
                    }
                    switch "" {
                    case "":
                        for fi13 := 0; fi13 < 18; fi13++ {
                            v11 += v11
                            v14 := GenS1{F0: 2147483647, F1: 32767}
                            for fi15 := 0; fi15 < 9; fi15++ {
                                v16 := []int{(fi13 ^ fi13)}
                                v11 = (v11 + ((((float32(3.14159) - float32((3.14159 - 3.14159))) + v11) / -2.25) - float32(3.14159)))
                                v17 := GenS3{F0: -128, F1: 42, F2: (3.14159 == 1e10)}
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            _ = fi13
                            _ = v14
                        }
                    case "foo bar":
                        if (!true) {
                            if (false && (((1.5 - 0.1) < -2.25) || true)) {
                                v18 := []int{32767, -2147483648, -481}
                                var v19 float32 = 0.1
                                _ = v18
                                _ = v19
                            } else {
                                for fi20 := 0; fi20 < 8; fi20++ {
                                    var v21 bool = false
                                    fmt.Println(fi20, v21)
                                    _ = fi20
                                    _ = v21
                                }
                                var v22 int32 = -879
                                _ = v22
                            }
                            for fi23 := 0; fi23 < 4; fi23++ {
                                v24 := GenS2{F0: 108, F1: -2147483648, F2: ("a" + "foo bar")}
                                _ = fi23
                                _ = v24
                            }
                        }
                    default:
                        for fi25 := 0; fi25 < 1; fi25++ {
                            var v26 uint8 = 60
                            var v27 int64 = 562
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                }
                _ = v11
            }
        }
        v28 := GenS4{F0: -154, F1: -2147483648, F2: 65535}
        for fi29 := 0; fi29 < 1; fi29++ {
            if (!(-1 != -32768)) {
                switch 2147483647 {
                case 565:
                    v30 := []int{}
                    _ = v30
                default:
                    v31 := []int{}
                    _ = v31
                }
                v32 := GenS5{F0: 2}
                var v33 uint64 = 0
                _ = v32
                _ = v33
            } else {
                if true {
                    for fi34 := 0; fi34 < 5; fi34++ {
                        var v35 bool = true
                        v35 = v35
                        fmt.Println(v35, v35)
                        v36 := GenS6{F0: false, F1: 2}
                        if true {
                            if ((v35 || v35) || true) {
                                var v37 uint = 42
                                fmt.Println(fi34, v37)
                                v38 := GenS2{F0: 36, F1: fi34, F2: ""}
                                _ = v37
                                _ = v38
                            } else {
                                v39 := []int{}
                                _ = v39
                            }
                            v35 = (1e10 > 1e10)
                        } else {
                            v35 = (v35 || v35)
                            v35 = true
                            v40 := []int{-2147483648}
                            _ = v40
                        }
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    v41 := GenS6{F0: true, F1: 745}
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 5; fi42++ {
                        v43 := []int{(fi42 + fi29), (-(fi29)), fi29}
                        var v44 int64 = -963
                        v45 := GenS1{F0: -1, F1: 0}
                        _ = fi42
                        _ = v43
                        _ = v44
                        _ = v45
                    }
                }
                v46 := GenS7{F0: 239, F1: 143, F2: fi29}
                fmt.Println(fi29)
                _ = v46
            }
            _ = fi29
        }
        _ = v28
    }
    var v47 int64 = 509
    _ = v47
    return 584
}

func main() {
    var v1 uint = 255
    v2 := make(map[int]bool)
    var v3 float64 = 0.1
    v48 := GenS5{F0: (v1 ^ v1)}
    v1 = v1
    for fi49 := 0; fi49 < 19; fi49++ {
        var v50 uint32 = 65535
        fmt.Println(v3, v3, fi49)
        _ = fi49
        _ = v50
    }
    v1 = genFn0()
    v1 = (v1 | 430)
    v51 := []int{42}
    var v52 uint16 = 42
    v51 = append(v51, -683)
    var v53 int64 = 42
    var v54 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
