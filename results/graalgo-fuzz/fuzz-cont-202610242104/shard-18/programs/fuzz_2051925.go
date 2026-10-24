package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint16; F2 int16 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 string; F1 string; F2 uint }
type GenS6 struct { F0 uint; F1 uint }
func genFn0() int16 {
    var v4 int64 = 1
    v5 := GenS1{F0: (v4 >> ((255) & 63))}
    var v6 float64 = 1.5
    switch -15 {
    case -2147483648:
        v7 := []int{32767}
        _ = v7
    case 429:
        v8 := []int{0, 127}
        _ = v8
    }
    fmt.Println(v4, v6)
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int8) int64 {
    var v9 float32 = float32(float64(-2.25))
    var v10 uint16 = 42
    _ = p0
    _ = v10
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    v11 := GenS1{F0: 9223372036854775807}
    v12 := []int{}
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 int32 = 42
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 1
    v16 := GenS1{F0: genFn1(127)}
    var v17 uint16 = (v15 ^ v15)
    v17 = 255
    if ((v15 * 42) >= (v17 * (v15 | v15))) {
        var v18 int32 = 42
        _ = v18
    } else {
        v19 := v3[2147483647]
        v17 += (v17 & v17)
        for fi20 := 0; fi20 < 6; fi20++ {
            v21 := GenS1{F0: -1}
            for fi22 := 0; fi22 < 4; fi22++ {
                switch (fi20 ^ fi20) {
                case -649:
                    var v23 int16 = 502
                    _ = v23
                default:
                    v24 := GenS2{F0: v19}
                    var v25 uint16 = v15
                    _ = v24
                    _ = v25
                }
                delete(v3, 863)
                var v26 bool = true
                _ = fi22
                _ = v26
            }
            var v27 int = 747
            _ = fi20
            _ = v21
            _ = v27
        }
        _ = v19
    }
    if (331 == 255) {
        var v28 int = -9223372036854775808
        switch (v28 << ((42) & 63)) {
        case 42:
            _ = len(v12)
        default:
            v28 = v28
        }
        _ = v28
    } else {
        if (32767 < genFn0()) {
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 int16 = 1
                delete(v1, fi29)
                for fi31 := 0; fi31 < 9; fi31++ {
                    v30 = genFn0()
                    var v32 int16 = genFn0()
                    var v33 float32 = float32((((float64(float32((1e10 + (0.0 / 0.1)))) + (1e10 + (1e10 - (float64(-2.25) * 1e10)))) * -2.25) * 0.0))
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = fi29
                _ = v30
            }
        } else {
            for fi34 := 0; fi34 < 1; fi34++ {
                if false {
                    var v35 int32 = -1
                    _ = v35
                } else {
                    v2[fi34] = fi34
                }
                if false {
                    switch (fi34 - -9223372036854775808) {
                    case 32767:
                        switch fi34 {
                        case 1:
                            fmt.Println(v17)
                        default:
                            v36 := GenS3{F0: -32768, F1: v17, F2: 32767}
                            _ = v36
                        }
                    case 472:
                        v17 = 0
                    }
                    var v37 int64 = 9223372036854775807
                    _ = v37
                } else {
                    if (((0 <= 604) || (230 > 1000)) || (18446744073709551615 <= 42)) {
                        v17 = (^(v15))
                        if false {
                            v17 -= (v15 ^ v15)
                            switch fi34 {
                            case 85:
                                var v38 uint8 = 255
                                _ = v38
                            case -1:
                                v12 = append(v12, fi34)
                            case -252:
                                _ = len(v12)
                            }
                        } else {
                            for fi39 := 0; fi39 < 7; fi39++ {
                                v40 := GenS4{F0: "", F1: 0}
                                if ((!false) || (false || (((-2.25 + (1e10 * 0.1)) + (-2.25 + 1.5)) == 1e10))) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v17 -= v17
                                    default:
                                        v41 := GenS4{F0: ("a" + "hello"), F1: 42}
                                        _ = v41
                                    }
                                } else {
                                    switch fi39 {
                                    case -61:
                                        delete(v2, 184)
                                    default:
                                        v42 := GenS5{F0: "hello", F1: "", F2: 4294967295}
                                        _ = v42
                                    }
                                    var v43 uint = 374
                                    var v44 float32 = 1e10
                                    _ = v43
                                    _ = v44
                                }
                                _ = fi39
                                _ = v40
                            }
                            var v45 bool = false
                            _ = v45
                        }
                    } else {
                        var v46 uint64 = 42
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 14; fi47++ {
                        var v48 uint = 4294967295
                        if (!(!false)) {
                            var v49 uint16 = v17
                            var v50 uint64 = 4294967295
                            _ = v49
                            _ = v50
                        } else {
                            v12 = append(v12, fi47)
                            v17 -= (^(v17))
                            v51 := GenS6{F0: (v48 - 2), F1: (^(v48))}
                            _ = v51
                        }
                        if (!true) {
                            fmt.Println(fi34, v17)
                            v52 := GenS1{F0: -1}
                            _ = v52
                        }
                        _ = fi47
                        _ = v48
                    }
                    var v53 uint8 = 42
                    _ = v53
                }
                _ = fi34
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
