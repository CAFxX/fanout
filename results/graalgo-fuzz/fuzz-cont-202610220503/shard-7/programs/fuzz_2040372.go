package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 bool }
type GenS2 struct { F0 uint; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint; F1 int64; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 int32; F2 int64 }
func genFn0() bool {
    v7 := []int{}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int16 = -1
        if true {
            if (false || (!false)) {
                v9 += (v9 + (^(v9)))
                v9 = (v9 ^ (-(v9)))
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            if true {
                if len(v7) > 0 {
                    v7[0] = fi8
                }
            }
            v9 = v9
        }
        _ = fi8
        _ = v9
    }
    _ = v7
    return false
}
func genFn1(p0 uint32, p1 int8, p2 int16) uint8 {
    v10 := []int{}
    var v11 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 2
}

func main() {
    var v1 uint16 = 255
    var v2 bool = (123 >= 0)
    var v3 int16 = 32767
    var v4 int = -1
    v5 := make(map[uint32]uint32)
    var v6 bool = v2
    v12 := GenS1{F0: 9223372036854775807, F1: 524, F2: true}
    var v13 int64 = -9223372036854775808
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int8 = -1
        for fi16 := 0; fi16 < 1; fi16++ {
            v5[735] = 60
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    if (!((v13 * -9223372036854775808) >= 731)) {
        if v6 {
            var v17 float64 = -2.25
            _ = v17
        } else {
            fmt.Println(v4, v6)
            v6 = (9223372036854775807 >= (v13 & v13))
        }
        fmt.Println(v3)
        var v18 string = (("foo bar" + "hello") + (("a" + "a") + "hello"))
        _ = v18
    } else {
        fmt.Println(v6)
        v19 := GenS1{F0: 173, F1: v13, F2: v6}
        var v20 int = -244
        _ = v19
        _ = v20
    }
    v21 := v5[42]
    v21 = v21
    if v2 {
        delete(v5, (v21 % (((v21 << ((872) & 63))) | 1)))
    } else {
        v22 := []int{(v4 | -9223372036854775808), (v4 * 32768), 128}
        v23 := GenS1{F0: 778, F1: v13, F2: false}
        _ = v22
        _ = v23
    }
    v24 := []int{(v4 & (v4 / ((v4) | 1))), (v4 % (((^(v4))) | 1)), (v4 << ((2) & 63))}
    v25 := GenS2{F0: 559, F1: (v1 | 2), F2: -1}
    v2 = v6
    v5[(v21 * v21)] = (v21 >> ((4294967295) & 63))
    if (988 == 255) {
        if false {
            for fi26 := 0; fi26 < 17; fi26++ {
                v2 = v2
                if (-1 <= 0) {
                    if (false && (v1 < (^(v1)))) {
                        v27 := GenS3{F0: 4294967295, F1: (-(v13)), F2: "foo bar"}
                        var v28 int8 = 100
                        _ = v27
                        _ = v28
                    } else {
                        if ((-(v3)) > -318) {
                            for fi29 := 0; fi29 < 20; fi29++ {
                                v13 -= (v13 * (-(v13)))
                                v30 := GenS1{F0: 558, F1: (-(v13)), F2: false}
                                if v6 {
                                    fmt.Println(v13)
                                }
                                if len(v24) > 0 {
                                    v24[0] = v4
                                }
                                fmt.Println(fi29)
                                _ = fi29
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 17; fi31++ {
                                fmt.Println(v4, fi31, v4)
                                var v32 int = 43
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                                _ = fi31
                                _ = v32
                            }
                            _ = len(v24)
                        }
                    }
                } else {
                    _ = len(v24)
                }
                if (v6 || v2) {
                    switch fi26 {
                    case 9223372036854775807:
                        v1 = 365
                    case 93:
                        switch fi26 {
                        case -953:
                            _ = len(v24)
                        case -2147483648:
                            switch (v4 - v4) {
                            case 127:
                                v5[(^(v21))] = (v21 & (v21 % ((v21) | 1)))
                            default:
                                var v33 float64 = -2.25
                                _ = v33
                            }
                        }
                    }
                } else {
                    var v34 uint16 = v1
                    _ = v34
                }
                _ = fi26
            }
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 int = int(v21)
                var v37 uint16 = 851
                var v38 bool = false
                _ = fi35
                _ = v36
                _ = v37
                _ = v38
            }
            v39 := GenS3{F0: 255, F1: (v13 / (((v13 % ((v13) | 1))) | 1)), F2: ("a" + "foo bar")}
            _ = v39
        }
    } else {
        v3 += v3
        if true {
            var v40 int32 = 2147483647
            delete(v5, 0)
            for fi41 := 0; fi41 < 19; fi41++ {
                v42 := GenS4{F0: 0}
                _ = fi41
                _ = v42
            }
            _ = v40
        } else {
            v43 := GenS5{F0: (v1 - v1), F1: 491, F2: 144}
            var v44 int16 = (v3 << ((uint(v21)) & 63))
            _ = v43
            _ = v44
        }
        if v2 {
            fmt.Println(v4, v21)
            if (305 > 0) {
                var v45 uint = 18446744073709551615
                v4 = v4
                _ = v45
            } else {
                var v46 uint = 65535
                v24 = append(v24, 128)
                _ = v46
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
