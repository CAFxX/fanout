package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint64; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint64, p2 int) int {
    var v4 string = "a"
    v5 := []int{}
    v6 := GenS1{F0: 42}
    v7 := GenS1{F0: 177}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 bool = ("a" < ("" + "\n\t"))
    var v2 int = 127
    var v3 int16 = -842
    v8 := []int{(v2 | int(v3)), v2}
    if ((v3 != (v3 - 1)) && v1) {
        fmt.Println(v2, v1, v2)
        var v9 int = (v2 / ((v2) | 1))
        _ = v9
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            v3 = 32767
            _ = fi10
        }
    }
    v3 = (^(v3))
    for fi11 := 0; fi11 < 9; fi11++ {
        v2 = v2
        v12 := GenS1{F0: 1}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi11
        _ = v12
    }
    var v13 int8 = 42
    if v1 {
        var v14 uint64 = 42
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 uint64 = (v14 ^ v14)
            for fi17 := 0; fi17 < 16; fi17++ {
                v14 += v16
                _ = fi17
            }
            var v18 uint64 = (v14 << ((255) & 63))
            _ = fi15
            _ = v16
            _ = v18
        }
        if v1 {
            if (100 == v13) {
                var v19 string = ""
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = v19
            } else {
                for fi20 := 0; fi20 < 1; fi20++ {
                    switch v2 {
                    case -640:
                        v21 := GenS1{F0: 2}
                        _ = v21
                    }
                    _ = fi20
                }
            }
            if len(v8) > 0 {
                v8[0] = 9223372036854775807
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        _ = v14
    }
    if (-1 != -611) {
        v3 = (v3 % ((v3) | 1))
        fmt.Println(v13, v2)
    }
    if (genFn0(v1, 295, v2) != -73) {
        var v22 int64 = 9223372036854775807
        if len(v8) > 0 {
            v8[0] = (v2 - genFn0((!v1), 0, v2))
        }
        _ = v22
    }
    v13 = v13
    var v23 bool = (v13 < (v13 - v13))
    _ = len(v8)
    v24 := GenS1{F0: 103}
    if (v23 && true) {
        var v25 float32 = ((float32(0.1) + (3.14159 * (1.5 + -2.25))) - 0.0)
        for fi26 := 0; fi26 < 9; fi26++ {
            if (108 <= 255) {
                v27 := GenS2{F0: true, F1: 1, F2: 412}
                _ = v27
            } else {
                switch v2 {
                case 9223372036854775807:
                    switch 951 {
                    case 750:
                        switch (-(v2)) {
                        case -9223372036854775808:
                            fmt.Println(fi26, v13)
                        case 32768:
                            v2 = v2
                        default:
                            if (v13 != (v13 | (v13 | v13))) {
                                if (!v23) {
                                    v13 -= 1
                                } else {
                                    v2 -= (-(fi26))
                                    v28 := GenS2{F0: ((128 < 127) && ((1.5 / float64(v25)) == 0.0)), F1: 1, F2: (^(fi26))}
                                    _ = v28
                                }
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                var v29 float64 = 0.1
                                _ = v29
                            } else {
                                v30 := GenS3{F0: 128}
                                var v31 int16 = 42
                                v31 = (v31 % ((v3) | 1))
                                _ = v30
                                _ = v31
                            }
                        }
                    case 0:
                        v2 = v2
                    case -957:
                        for fi32 := 0; fi32 < 9; fi32++ {
                            fmt.Println(v1)
                            var v33 int = v2
                            v34 := GenS4{F0: 127}
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    default:
                        if (!false) {
                            var v35 int32 = -172
                            var v36 int32 = v35
                            var v37 uint16 = 1
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                    }
                case 128:
                    v13 -= (v13 ^ v13)
                default:
                    var v38 int64 = 958
                    _ = v38
                }
            }
            v3 -= (v3 - (^(v3)))
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi26
        }
        _ = v25
    } else {
        var v39 uint = 18446744073709551615
        _ = v39
    }
    v2 = genFn0((!v23), 42, v2)
    if len(v8) > 0 {
        v8[0] = 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v8
}
