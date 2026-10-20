package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS5 struct { F0 int16; F1 int32; F2 int }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 uint }
func genFn0(p0 string, p1 uint) uint8 {
    v7 := GenS1{F0: 1000, F1: 2147483647, F2: p1}
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 int8, p1 int32) string {
    v8 := []int{-558, 2147483647, 32768}
    if true {
        p0 = (-(p0))
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        p1 -= (p1 % (((^(p1))) | 1))
        fmt.Println(p0, fi9, p0)
        _ = fi9
    }
    v10 := GenS2{F0: 590, F1: 504}
    p1 -= (p1 >> ((354) & 63))
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    return ""
}

func main() {
    var v1 uint64 = 4294967295
    var v2 string = ""
    v3 := make(map[string]string)
    var v4 uint = 42
    v5 := make(map[uint32]bool)
    v6 := make(map[int]uint32)
    v1 += (v1 & 42)
    var v11 int = 0
    v12 := GenS1{F0: 32767, F1: 0, F2: v4}
    for fi13 := 0; fi13 < 15; fi13++ {
        if false {
            v14 := GenS2{F0: 155, F1: 32767}
            v4 = v4
            switch v2 {
            case "":
                switch v11 {
                case 127:
                    fmt.Println(v11, fi13)
                case 1:
                    fmt.Println(fi13, v2)
                case 0:
                    for fi15 := 0; fi15 < 11; fi15++ {
                        v16 := []int{fi13}
                        var v17 string = ((v2 + v2) + (("\n\t" + v2) + v2))
                        if true {
                            if ((!false) && true) {
                                _ = len(v5)
                                var v18 float64 = (float64(((0.1 - ((((1e10 + (float32((0.0 / 3.14159)) / 1.0)) * float32(1.5)) * 0.1) * (float32(float64(1.5)) / (1e10 / 1.5)))) / 1e10)) - (((1.5 * 0.0) + 1e10) / -2.25))
                                var v19 int32 = 2147483647
                                _ = v18
                                _ = v19
                            }
                            v1 -= v1
                        } else {
                            var v20 uint8 = 128
                            _ = v20
                        }
                        var v21 int8 = -128
                        if (4294967295 >= (v1 | 2)) {
                            var v22 uint8 = 1
                            var v23 int32 = 0
                            var v24 int64 = 42
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            v1 += (v1 & v1)
                            var v25 uint32 = uint32(v1)
                            _ = v25
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v21
                    }
                }
            }
            _ = v14
        }
        _ = len(v3)
        _ = fi13
    }
    v26 := GenS2{F0: (v4 * 831), F1: uint16(v1)}
    var v27 int = v11
    v28 := GenS2{F0: uint(v1), F1: uint16(v1)}
    v27 += -1
    if (!false) {
        fmt.Println(v1, v4)
        for fi29 := 0; fi29 < 19; fi29++ {
            switch (v2 + (v2 + "")) {
            case "":
                if (65 >= 1) {
                    fmt.Println(v4, fi29, v1)
                    if (-601 == 32767) {
                        v30 := GenS2{F0: v4, F1: 32767}
                        if false {
                            var v31 int = -424
                            for fi32 := 0; fi32 < 14; fi32++ {
                                v33 := GenS2{F0: v4, F1: 294}
                                v34 := GenS1{F0: 0, F1: 176, F2: (v4 | (^(v4)))}
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                            _ = v31
                        } else {
                            fmt.Println(fi29)
                        }
                        v35 := GenS2{F0: v4, F1: 32767}
                        _ = v30
                        _ = v35
                    }
                    v36 := []int{82, (fi29 | fi29), v27}
                    _ = v36
                }
            case "foo bar":
                var v37 bool = true
                _ = v37
            case "\n\t":
                fmt.Println(v27)
            }
            if ((0 < 32767) && (true || (true && true))) {
                v2 = v2
                if ((0.1 / 0.1) <= (0.0 - 0.1)) {
                    v38 := []int{(fi29 >> (((v4 * v4)) & 63)), v11}
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 14; fi39++ {
                        fmt.Println(v11, v4, v2)
                        var v40 uint = 0
                        _ = fi39
                        _ = v40
                    }
                }
            }
            for fi41 := 0; fi41 < 16; fi41++ {
                v42 := GenS3{F0: 2147483647}
                var v43 string = ""
                if true {
                    var v44 uint8 = 0
                    _ = v44
                } else {
                    v6[v27] = 4294967295
                }
                v45 := GenS4{F0: 2, F1: (v1 >> (((v4 | v4)) & 63)), F2: 83}
                fmt.Println(fi29, v11, fi41)
                _ = fi41
                _ = v42
                _ = v43
                _ = v45
            }
            _ = fi29
        }
        _ = len(v5)
    } else {
        switch v2 {
        case "hello":
            v46 := []int{}
            _ = v46
        case "a":
            var v47 int16 = 42
            _ = v47
        case "\n\t":
            v48 := GenS5{F0: -32768, F1: 199, F2: v11}
            _ = v48
        }
    }
    v49 := v3[v2]
    if (!true) {
        var v50 int16 = 256
        v51 := GenS3{F0: 2147483647}
        for fi52 := 0; fi52 < 20; fi52++ {
            switch "" {
            case "a":
                for fi53 := 0; fi53 < 19; fi53++ {
                    fmt.Println(v2)
                    var v54 uint = v4
                    _ = fi53
                    _ = v54
                }
            default:
                var v55 int32 = 0
                _ = v55
            }
            v56 := GenS2{F0: v4, F1: 42}
            _ = fi52
            _ = v56
        }
        _ = v50
        _ = v51
    }
    for fi57 := 0; fi57 < 14; fi57++ {
        if ((42 > 32767) && true) {
            if (v4 >= (v4 - v4)) {
                v5[615] = false
            }
            switch (v11 + -418) {
            case 15:
                var v58 bool = (!false)
                _ = v58
            case 2147483647:
                var v59 bool = (0.1 > ((float32((1.5 / 1e10)) - -2.25) + 1.5))
                _ = v59
            case -2147483648:
                if false {
                    for fi60 := 0; fi60 < 8; fi60++ {
                        delete(v5, 2147483648)
                        _ = fi60
                    }
                    switch (^(v11)) {
                    case 209:
                        if ((false && (-89 != 100)) || ((128 >= 128) && false)) {
                            v49 = (v49 + "\n\t")
                        }
                    case -1:
                        v61 := GenS6{F0: (!true), F1: false}
                        _ = v61
                    case -2147483648:
                        switch -684 {
                        case -2147483648:
                            var v62 int16 = 0
                            _ = v62
                        case 4294967295:
                            var v63 int = fi57
                            _ = v63
                        default:
                            var v64 int = (v27 & v11)
                            _ = v64
                        }
                    default:
                        if true {
                            fmt.Println(v1)
                            fmt.Println(fi57, v1)
                        }
                    }
                    v65 := []int{(fi57 ^ 13)}
                    _ = v65
                }
            }
        }
        _ = fi57
    }
    if true {
        v66 := v6[42]
        v66 = v66
        fmt.Println(v11, v11)
        _ = v66
    } else {
        switch v2 {
        case "hello":
            fmt.Println(v27)
        case "foo bar":
            v27 = v11
        default:
            for fi67 := 0; fi67 < 5; fi67++ {
                v49 = v49
                v68 := []int{}
                _ = fi67
                _ = v68
            }
        }
    }
    if false {
        v49 = "a"
        v69 := GenS7{F0: (v4 << (((v4 + (^(v4)))) & 63))}
        _ = v69
    } else {
        _ = len(v6)
        if (1 < 42) {
            v1 -= (v1 % ((v1) | 1))
            fmt.Println(v2, v49)
            v70 := GenS5{F0: 42, F1: 1, F2: v11}
            _ = v70
        } else {
            v71 := v3["a"]
            v49 = "a"
            _ = len(v3)
            _ = v71
        }
    }
    for fi72 := 0; fi72 < 4; fi72++ {
        v1 = (v1 - v1)
        v73 := GenS2{F0: (^(v4)), F1: 0}
        _ = fi72
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
