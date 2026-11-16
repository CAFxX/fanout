package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8; F1 bool; F2 int }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 uint8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 38
    v3 := []uint32{927, 2147483647}
    fmt.Println(v2)
    var v4 string = "\n\t"
    fmt.Println(v4)
    _ = len(v3)
    if (false && ((1 <= -320) || true)) {
        if true {
            if ((4294967295 > 1) && ((^(v2)) > (v2 >> ((2) & 63)))) {
                var v5 uint16 = 243
                if false {
                    v6 := GenS1{F0: 255}
                    _ = v6
                } else {
                    v7 := GenS2{F0: 305, F1: 0}
                    var v8 bool = true
                    _ = v7
                    _ = v8
                }
                if (true || ((v2 >= v5) && (!(2 >= 4294967295)))) {
                    var v9 float64 = float64(0.0)
                    var v10 int8 = 100
                    v11 := GenS3{F0: 0, F1: true}
                    _ = v10
                    _ = v11
                    _ = v9
                } else {
                    for fi12 := 0; fi12 < 14; fi12++ {
                        v2 += 425
                        v1[fi12] = 906
                        _ = fi12
                    }
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v5 = (v2 ^ v5)
                        v5 += (^(v5))
                        for fi14 := 0; fi14 < 11; fi14++ {
                            switch fi13 {
                            case 127:
                                var v15 uint32 = 2147483648
                                _ = v15
                            }
                            v16 := GenS1{F0: v5}
                            if true {
                                v17 := GenS2{F0: fi13, F1: 701}
                                var v18 uint8 = 1
                                _ = v17
                                _ = v18
                            } else {
                                if true {
                                    switch (fi13 >> ((18446744073709551615) & 63)) {
                                    case 4294967295:
                                        var v19 int32 = 747
                                        _ = v19
                                    case -628:
                                        v20 := GenS2{F0: -746, F1: 255}
                                        _ = v20
                                    }
                                    v21 := GenS2{F0: (fi13 % (((-(fi13))) | 1)), F1: 0}
                                    fmt.Println(fi14)
                                    _ = v21
                                } else {
                                    v22 := GenS1{F0: 703}
                                    if false {
                                        delete(v1, 4294967295)
                                    }
                                    _ = v22
                                }
                                var v23 uint16 = (^(v5))
                                if (-1 == -38) {
                                    v24 := GenS4{F0: false}
                                    var v25 uint8 = 127
                                    if len(v3) > 0 {
                                        v3[0] = 536
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                    _ = v24
                                    _ = v25
                                }
                                _ = v23
                            }
                            if (false && ((672 < 42) && true)) {
                                var v26 string = "hello"
                                _ = v26
                            }
                            if ((!(!((true || (9223372036854775807 != 42)) && (65535 < 1)))) || true) {
                                var v27 uint64 = 4294967295
                                if len(v3) > 0 {
                                    v3[0] = 226
                                }
                                _ = v27
                            } else {
                                var v28 uint8 = 97
                                switch fi14 {
                                case -9223372036854775808:
                                    v29 := GenS2{F0: -694, F1: 0}
                                    _ = v29
                                }
                                _ = v28
                            }
                            _ = fi14
                            _ = v16
                        }
                        _ = fi13
                    }
                    v4 = v4
                }
                _ = v5
            }
            if (63 > 31) {
                v2 = v2
                for fi30 := 0; fi30 < 14; fi30++ {
                    var v31 bool = (-1 <= -9223372036854775808)
                    v2 -= (v2 & (^(v2)))
                    _ = fi30
                    _ = v31
                }
                if len(v3) > 0 {
                    v3[0] = 0
                }
            } else {
                v32 := GenS2{F0: 32767, F1: 65535}
                _ = v32
            }
            v2 = v2
        } else {
            if false {
                fmt.Println(v2, v4)
                v33 := GenS1{F0: 707}
                _ = v33
            }
            var v34 uint = 4294967295
            for fi35 := 0; fi35 < 1; fi35++ {
                var v36 int8 = -128
                _ = fi35
                _ = v36
            }
            _ = v34
        }
        for fi37 := 0; fi37 < 16; fi37++ {
            if (1 >= 0) {
                for fi38 := 0; fi38 < 3; fi38++ {
                    v2 += v2
                    v39 := GenS2{F0: fi38, F1: 757}
                    fmt.Println(v2, fi37)
                    _ = fi38
                    _ = v39
                }
            } else {
                var v40 int16 = 0
                _ = len(v3)
                var v41 uint = 65535
                _ = v40
                _ = v41
            }
            for fi42 := 0; fi42 < 16; fi42++ {
                if true {
                    v43 := GenS1{F0: v2}
                    _ = v43
                }
                if (false || false) {
                    v44 := GenS5{F0: -45, F1: true, F2: fi42}
                    var v45 uint64 = 634
                    _ = v44
                    _ = v45
                } else {
                    v2 = v2
                    if true {
                        if ((fi42 | (fi42 + (fi42 * fi37))) > (fi42 ^ fi42)) {
                            var v46 string = (("\n\t" + "") + "hello")
                            fmt.Println(v2, v46)
                            var v47 int16 = 42
                            v47 -= (v47 % ((v47) | 1))
                            _ = v46
                            _ = v47
                        }
                        var v48 string = v4
                        var v49 int = fi37
                        _ = v48
                        _ = v49
                    }
                    fmt.Println(v4, v2)
                }
                _ = fi42
            }
            _ = fi37
        }
        v1[2147483647] = 65535
    } else {
        var v50 int64 = 0
        if (0 > 12) {
            var v51 float32 = (3.14159 / float32((0.1 * 1e10)))
            _ = v51
        }
        _ = v50
    }
    if true {
        for fi52 := 0; fi52 < 20; fi52++ {
            if false {
                if (((100 >= 1) && true) || true) {
                    if ((1.5 - -2.25) < 3.14159) {
                        var v53 int = -811
                        if ((v2 * (v2 % ((789) | 1))) <= v2) {
                            v2 = 729
                            for fi54 := 0; fi54 < 13; fi54++ {
                                v2 -= v2
                                fmt.Println(v2, v2, v53)
                                switch fi52 {
                                case 127:
                                    v2 = 380
                                case 42:
                                    v55 := GenS3{F0: 1, F1: false}
                                    _ = v55
                                case 9223372036854775807:
                                    v2 = 471
                                default:
                                    var v56 uint8 = 128
                                    _ = v56
                                }
                                _ = fi54
                            }
                            v4 = (v4 + v4)
                        }
                        _ = v53
                    }
                    var v57 uint8 = 115
                    _ = v57
                }
            }
            v1[(fi52 | -578)] = 2147483647
            _ = fi52
        }
    }
    if (2 > 42) {
        v58 := GenS6{F0: 4294967295, F1: 607}
        v1[-604] = 65535
        _ = v58
    }
    var v59 uint8 = 2
    v60 := GenS7{F0: 127}
    var v61 int = 32767
    fmt.Println(v2)
    var v62 uint64 = 682
    v4 = "\n\t"
    v63 := GenS7{F0: uint8(v62)}
    for fi64 := 0; fi64 < 2; fi64++ {
        if false {
            if true {
                v2 = v2
                v62 = v62
                v3 = append(v3, 4294967295)
            } else {
                fmt.Println(v61, v59, v62)
            }
        }
        _ = fi64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
