package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 int64; F1 bool; F2 uint8 }
type GenS8 struct { F0 int32; F1 int8; F2 int32 }
func genFn0(p0 uint64, p1 uint8) int32 {
    if false {
        v6 := GenS1{F0: 100}
        v7 := []int{4294967295, -698, 783, -198}
        switch "\n\t" {
        case "\n\t":
            v8 := GenS1{F0: -115}
            _ = v8
        default:
            for fi9 := 0; fi9 < 17; fi9++ {
                if true {
                    var v10 string = ("foo bar" + "foo bar")
                    fmt.Println(p0)
                    _ = v10
                } else {
                    v11 := GenS2{F0: 1}
                    _ = v11
                }
                var v12 uint = 594
                _ = fi9
                _ = v12
            }
        }
        _ = v6
        _ = v7
    } else {
        p1 -= p1
        var v13 uint8 = p1
        _ = v13
    }
    if (2 != 65535) {
        switch "a" {
        case "hello":
            v14 := []int{834, -1, 128}
            _ = v14
        case "foo bar":
            var v15 int16 = 91
            _ = v15
        case "a":
            if ((1 != 2) || (true && false)) {
                if true {
                    v16 := GenS2{F0: 1000}
                    _ = v16
                }
            }
        }
        p1 += (p1 % ((178) | 1))
    } else {
        var v17 float32 = 0.1
        _ = v17
    }
    var v18 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v18
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 0
    var v3 int64 = -524
    var v4 uint8 = 205
    var v5 uint16 = 65535
    v19 := v1[32767]
    if false {
        for fi20 := 0; fi20 < 17; fi20++ {
            for fi21 := 0; fi21 < 3; fi21++ {
                v22 := GenS1{F0: -1}
                v4 = 121
                if ((0.1 == 0.1) || false) {
                    fmt.Println(v19, v5)
                } else {
                    v23 := GenS3{F0: fi20, F1: -1}
                    _ = v23
                }
                _ = fi21
                _ = v22
            }
            v3 -= v3
            _ = fi20
        }
    } else {
        v24 := GenS3{F0: 32767, F1: 0}
        _ = v24
    }
    switch "hello" {
    case "":
        v3 = (v3 % (((v3 | (-(v3)))) | 1))
    default:
        var v25 string = "hello"
        _ = v25
    }
    var v26 uint8 = 2
    var v27 int64 = v3
    for fi28 := 0; fi28 < 3; fi28++ {
        fmt.Println(v19, v27, v3)
        if false {
            var v29 int16 = -753
            if (true || (-766 > -668)) {
                for fi30 := 0; fi30 < 16; fi30++ {
                    if (!true) {
                        var v31 uint = 172
                        v31 = (v31 | v31)
                        var v32 uint16 = uint16(v19)
                        v29 = -1
                        _ = v31
                        _ = v32
                    }
                    if false {
                        var v33 bool = ((1.5 - 0.0) == ((-2.25 / 0.1) - float64(0.1)))
                        v34 := GenS3{F0: 128, F1: 1}
                        fmt.Println(v19, v27, v2)
                        _ = v33
                        _ = v34
                    }
                    if false {
                        v35 := GenS3{F0: fi30, F1: -107}
                        switch "\n\t" {
                        case "hello":
                            var v36 uint64 = (v2 - v2)
                            _ = v36
                        default:
                            var v37 uint = 1
                            _ = v37
                        }
                        _ = v35
                    } else {
                        if ((v26 | 2) <= (^(v4))) {
                            if true {
                                var v38 uint = 393
                                v27 = v27
                                _ = v38
                            }
                            v3 += (v3 % (((v27 % ((v27) | 1))) | 1))
                            v2 = v2
                            v39 := GenS4{F0: 127, F1: v19, F2: v5}
                            fmt.Println(v3, v19, v27)
                            v40 := GenS3{F0: int(v2), F1: 53}
                            _ = v39
                            _ = v40
                        }
                    }
                    _ = fi30
                }
            }
            var v41 uint32 = v19
            _ = v29
            _ = v41
        }
        _ = fi28
    }
    if (true || (false || true)) {
        for fi42 := 0; fi42 < 6; fi42++ {
            v43 := GenS5{F0: -59}
            var v44 bool = true
            var v45 int16 = -145
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
        var v46 uint8 = 2
        _ = v46
    }
    fmt.Println(v3, v27)
    v47 := v1[-9223372036854775808]
    fmt.Println(v26, v19)
    v48 := GenS6{F0: "\n\t", F1: v2}
    v26 = 128
    v26 += v26
    if (true || (true && true)) {
        for fi49 := 0; fi49 < 7; fi49++ {
            v50 := []int{fi49}
            var v51 uint = 0
            _ = fi49
            _ = v50
            _ = v51
        }
        v52 := GenS7{F0: int64(v26), F1: true, F2: 225}
        _ = v52
    }
    v53 := []int{128}
    v54 := GenS8{F0: genFn0(uint64(v19), v26), F1: 19, F2: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v53:", len(v53), v53)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v53
    _ = v54
}
