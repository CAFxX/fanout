package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 string; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 string; F1 uint16; F2 string }
func genFn0(p0 int, p1 uint16) int64 {
    fmt.Println(p0, p1)
    switch -657 {
    case 128:
        switch "" {
        case "foo bar":
            v7 := []int{(-(p0)), p0, p0}
            _ = v7
        case "\n\t":
            var v8 bool = (1 != 467)
            _ = v8
        case "hello":
            p1 = p1
        default:
            v9 := GenS1{F0: 2147483647, F1: p1}
            _ = v9
        }
    }
    p1 = 32767
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 uint) int {
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 int16 = 42
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    fmt.Println(p0)
    p0 = p0
    p0 -= p0
    _ = p0
    return 127
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 4294967295
    v3 := make(map[int]uint32)
    v4 := []uint32{(v2 / (((v2 >> ((683) & 63))) | 1)), (v2 - (v2 ^ 0)), 42}
    var v5 uint8 = 147
    v6 := []int{}
    if false {
        if (!true) {
            v13 := GenS2{F0: 42}
            if (genFn0(-1, 65535) != -1) {
                v5 += v5
            } else {
                switch 265 {
                case 779:
                    if true {
                        v14 := GenS3{F0: 2147483647, F1: "hello", F2: 2}
                        var v15 uint8 = (v5 | v5)
                        _ = v14
                        _ = v15
                    }
                case -1:
                    fmt.Println(v5)
                case 4294967295:
                    if false {
                        for fi16 := 0; fi16 < 10; fi16++ {
                            switch "\n\t" {
                            case "a":
                                var v17 uint = 4294967295
                                _ = v17
                            case "hello":
                                switch "a" {
                                case "foo bar":
                                    var v18 bool = (false || false)
                                    _ = v18
                                default:
                                    if (65535 >= 4294967295) {
                                        var v19 int8 = 0
                                        _ = v19
                                    }
                                }
                            default:
                                v20 := GenS4{F0: 127, F1: 2, F2: 42}
                                _ = v20
                            }
                            if (!false) {
                                var v21 float64 = 0.1
                                var v22 int16 = 0
                                for fi23 := 0; fi23 < 13; fi23++ {
                                    v5 += 255
                                    _ = fi23
                                }
                                _ = v21
                                _ = v22
                            }
                            _ = fi16
                        }
                    } else {
                        v24 := GenS5{F0: v5, F1: 278}
                        for fi25 := 0; fi25 < 6; fi25++ {
                            if (true && (!((fi25 >> ((4294967295) & 63)) < -768))) {
                                _ = len(v4)
                                v5 += (^(v5))
                            }
                            v2 += v2
                            v26 := GenS6{F0: "\n\t", F1: 812, F2: "hello"}
                            _ = fi25
                            _ = v26
                        }
                        _ = v24
                    }
                default:
                    v2 += (v2 >> ((0) & 63))
                }
                var v27 int16 = -32768
                _ = v27
            }
            _ = v13
        }
    }
    var v28 int = 9223372036854775807
    v28 -= v28
    for fi29 := 0; fi29 < 2; fi29++ {
        var v30 uint = 0
        _ = fi29
        _ = v30
    }
    if false {
        v31 := GenS6{F0: "a", F1: 255, F2: "\n\t"}
        _ = v31
    }
    v32 := GenS3{F0: 1, F1: ("" + "hello"), F2: (v2 + (^(v2)))}
    var v33 string = "foo bar"
    v34 := GenS3{F0: -2147483648, F1: "a", F2: 159}
    var v35 int32 = 305
    if (v2 <= (v2 << ((255) & 63))) {
        v36 := v3[v28]
        _ = v36
    }
    if false {
        if (v28 > v28) {
            switch ("" + v33) {
            case "":
                v35 = v35
            case "foo bar":
                for fi37 := 0; fi37 < 11; fi37++ {
                    v2 += (^(v2))
                    if (!false) {
                        var v38 float32 = (float32(0.0) * float32(0.1))
                        v1[(v33 + v33)] = (v28 >> ((2) & 63))
                        var v39 int32 = (v35 % (((v35 >> ((2) & 63))) | 1))
                        _ = v38
                        _ = v39
                    }
                    _ = fi37
                }
            default:
                for fi40 := 0; fi40 < 4; fi40++ {
                    v2 = (v2 / ((v2) | 1))
                    _ = fi40
                }
            }
            v2 += v2
            v35 += -742
        }
        if (v28 != -1) {
            var v41 float32 = (1e10 + 1.5)
            fmt.Println(v41, v33)
            for fi42 := 0; fi42 < 19; fi42++ {
                _ = len(v6)
                fmt.Println(v28, v28, v5)
                _ = fi42
            }
            _ = v41
        }
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
