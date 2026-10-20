package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int }
func genFn0() int {
    var v5 int8 = 42
    v6 := GenS1{F0: 0, F1: v5}
    v5 = v5
    var v7 uint64 = 63
    _ = v5
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 float32 = 1e10
    v4 := []uint32{65535}
    fmt.Println(v3)
    v8 := GenS1{F0: 65535, F1: 100}
    _ = len(v4)
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int16 = -32768
        if true {
            v10 -= v10
            for fi11 := 0; fi11 < 20; fi11++ {
                v3 = ((((v3 / 1.5) + v3) / float32(1.5)) - (0.1 - 1e10))
                if (1e10 > v3) {
                    switch (("\n\t" + "foo bar") + "") {
                    case "foo bar":
                        if (65535 <= 0) {
                            if ("a" >= "") {
                                v3 = (v3 / (-2.25 * 3.14159))
                                fmt.Println(fi11)
                            } else {
                                v12 := GenS1{F0: 0, F1: 76}
                                if (1.5 > 1.5) {
                                    var v13 int32 = 2147483647
                                    v10 = int16(v13)
                                    _ = v13
                                } else {
                                    var v14 float64 = 0.0
                                    v10 = 0
                                    if ((42 <= 1) || (127 >= 42)) {
                                        var v15 uint8 = 125
                                        v16 := GenS2{F0: (^(v10)), F1: 173}
                                        _ = v15
                                        _ = v16
                                    }
                                    _ = v14
                                }
                                var v17 float64 = 1.5
                                v18 := GenS3{F0: (v10 & v10)}
                                v19 := GenS4{F0: -2147483648, F1: (v10 & (v10 / (((v10 / ((v10) | 1))) | 1)))}
                                _ = v12
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                            var v20 uint64 = 42
                            if false {
                                v4 = append(v4, 65535)
                            } else {
                                v10 -= v10
                                _ = len(v4)
                                var v21 string = "a"
                                var v22 uint8 = 112
                                v23 := v1[fi11]
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            _ = v20
                        } else {
                            if (false || (1 >= 1)) {
                                switch 9223372036854775807 {
                                case 32767:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                v10 = (v10 | 32767)
                            } else {
                                v3 = ((0.0 - 0.0) * 0.0)
                            }
                            v24 := GenS4{F0: 373, F1: (v10 << ((42) & 63))}
                            if true {
                                v10 = -1
                                v25 := GenS5{F0: 55}
                                v10 += v10
                                v26 := GenS4{F0: 42, F1: v10}
                                v10 += (-(v10))
                                _ = v25
                                _ = v26
                            }
                            _ = v24
                        }
                    default:
                        if (!false) {
                            v3 -= (0.0 / 0.1)
                            v3 -= v3
                            fmt.Println(fi9, fi9)
                        } else {
                            switch fi9 {
                            case -1:
                                var v27 int = (fi9 | (^(fi9)))
                                _ = v27
                            case 42:
                                v10 -= 0
                            case 128:
                                var v28 int8 = -1
                                _ = v28
                            default:
                                var v29 string = ("" + ("a" + (("foo bar" + "foo bar") + "a")))
                                _ = v29
                            }
                        }
                    }
                } else {
                    switch genFn0() {
                    case 32767:
                        v10 += v10
                    case -2147483648:
                        switch "foo bar" {
                        case "":
                            if false {
                                v10 += 32767
                                v30 := GenS3{F0: (-(v10))}
                                _ = v30
                            } else {
                                if (!(976 > 42)) {
                                    v3 = v3
                                } else {
                                    v10 = int16(fi9)
                                }
                                if (true || (-1 >= 42)) {
                                    if false {
                                        v31 := GenS1{F0: 65535, F1: 0}
                                        fmt.Println(v3)
                                        var v32 float64 = 1e10
                                        var v33 uint64 = 355
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    } else {
                                        fmt.Println(v3)
                                    }
                                    delete(v1, (fi11 & (fi11 << ((0) & 63))))
                                    var v34 float64 = float64(1e10)
                                    _ = v34
                                } else {
                                    if true {
                                        fmt.Println(v10, fi11)
                                        var v35 uint64 = 817
                                        _ = v35
                                    }
                                    v3 = ((float32((3.14159 - 0.0)) - v3) + v3)
                                    v36 := GenS6{F0: (-(fi11)), F1: (^(v10))}
                                    v3 -= float32(1e10)
                                    _ = v36
                                }
                                v10 = (v10 | 232)
                                v3 += 0.0
                            }
                        }
                    default:
                        v37 := GenS5{F0: 42}
                        _ = v37
                    }
                    if true {
                        v10 += -1
                        var v38 uint64 = 42
                        fmt.Println(fi9, fi9, v38)
                        _ = v38
                    } else {
                        v4 = append(v4, 2147483647)
                        v3 = -2.25
                        v10 -= (v10 + (v10 & 473))
                    }
                }
                _ = fi11
            }
            var v39 uint = 400
            _ = v39
        } else {
            v10 = 558
            if true {
                for fi40 := 0; fi40 < 14; fi40++ {
                    v3 = v3
                    v41 := GenS2{F0: v10, F1: 1}
                    var v42 uint32 = 2147483648
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
            var v43 uint = 65535
            _ = v43
        }
        _ = fi9
        _ = v10
    }
    switch "\n\t" {
    case "":
        switch ("hello" + "") {
        case "foo bar":
            if true {
                v44 := GenS7{F0: 0}
                _ = v44
            }
        default:
            v3 += (((v3 - (v3 * (v3 - v3))) - 3.14159) + v3)
        }
    case "hello":
        if false {
            if (true && (86 >= 100)) {
                if (!false) {
                    var v45 int = 392
                    _ = v45
                }
            } else {
                var v46 uint64 = 2
                _ = v46
            }
        }
    default:
        v47 := GenS8{F0: 255, F1: 916}
        _ = v47
    }
    var v48 uint16 = 2
    if (!true) {
        for fi49 := 0; fi49 < 7; fi49++ {
            _ = len(v4)
            _ = fi49
        }
        v3 += v3
        var v50 uint16 = (v48 * (v48 ^ v48))
        _ = v50
    }
    if len(v4) > 0 {
        v4[0] = 65535
    }
    if (!true) {
        for fi51 := 0; fi51 < 3; fi51++ {
            for fi52 := 0; fi52 < 11; fi52++ {
                fmt.Println(v3, v48)
                fmt.Println(v3, fi52)
                v53 := GenS8{F0: 122, F1: 667}
                _ = fi52
                _ = v53
            }
            var v54 int8 = 100
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi51
            _ = v54
        }
        var v55 int32 = 0
        _ = v55
    }
    var v56 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v56
    _ = v8
}
