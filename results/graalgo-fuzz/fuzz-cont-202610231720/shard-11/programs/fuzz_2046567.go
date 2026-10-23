package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int; F1 uint64; F2 uint }
type GenS3 struct { F0 int }

func main() {
    var v1 int32 = 42
    var v2 int16 = -676
    var v3 uint16 = 65535
    var v4 uint8 = 1
    var v5 uint16 = v3
    var v6 uint32 = 1
    var v7 int8 = 42
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := []int{(-(fi8)), (fi8 / ((int(v4)) | 1))}
        _ = len(v9)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: (false || (false && (462 <= 18446744073709551615))), F1: 0}
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS2{F0: -1, F1: 42, F2: 4294967295}
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 int64 = int64(fi11)
            var v15 int32 = v1
            v16 := GenS1{F0: (0.0 <= 1.5), F1: 42}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        var v17 uint64 = 42
        _ = fi11
        _ = v12
        _ = v17
    }
    v4 += (v4 * v4)
    var v18 int16 = 1000
    if ((^(v6)) < 360) {
        switch 4294967295 {
        case 9223372036854775807:
            switch "hello" {
            case "\n\t":
                v7 = (v7 >> ((255) & 63))
            }
        case 128:
            if ((^(v4)) == (^(v4))) {
                v5 = v5
                for fi19 := 0; fi19 < 8; fi19++ {
                    var v20 string = (("a" + "") + "hello")
                    _ = fi19
                    _ = v20
                }
            } else {
                if true {
                    var v21 float32 = -2.25
                    v18 = v2
                    v5 = v3
                    _ = v21
                } else {
                    v22 := []int{}
                    for fi23 := 0; fi23 < 19; fi23++ {
                        if true {
                            if (42 > (v5 * (v5 >> ((264) & 63)))) {
                                var v24 int8 = v7
                                _ = v24
                            } else {
                                var v25 uint32 = 605
                                v25 += 419
                                _ = v25
                            }
                        } else {
                            var v26 uint = 0
                            v27 := GenS2{F0: fi23, F1: 2, F2: 42}
                            _ = v26
                            _ = v27
                        }
                        if false {
                            for fi28 := 0; fi28 < 9; fi28++ {
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                                _ = fi28
                            }
                            if len(v22) > 0 {
                                v22[0] = int(v18)
                            }
                            if ((v7 - (v7 / ((v7) | 1))) < (v7 ^ (v7 % ((1) | 1)))) {
                                var v29 int32 = (v1 | (v1 | 42))
                                v4 = (v4 - (^(v4)))
                                _ = v29
                            } else {
                                var v30 float64 = 3.14159
                                _ = v30
                            }
                        }
                        v1 = v1
                        _ = fi23
                    }
                    if len(v22) > 0 {
                        v22[0] = 128
                    }
                    _ = v22
                }
                for fi31 := 0; fi31 < 7; fi31++ {
                    v6 = v6
                    fmt.Println(v18, v4, v7)
                    _ = fi31
                }
                for fi32 := 0; fi32 < 19; fi32++ {
                    if true {
                        for fi33 := 0; fi33 < 11; fi33++ {
                            v34 := []int{fi32}
                            _ = fi33
                            _ = v34
                        }
                        switch fi32 {
                        case 9223372036854775807:
                            var v35 int16 = (v18 >> ((603) & 63))
                            _ = v35
                        case 128:
                            switch -1 {
                            case 9223372036854775807:
                                fmt.Println(v4, v18)
                            }
                        case 1:
                            v36 := GenS1{F0: false, F1: (v4 + v4)}
                            _ = v36
                        default:
                            if (1e10 < ((-2.25 / (float32(3.14159) - ((3.14159 - float32(-2.25)) / 1e10))) * (-2.25 - (1e10 + (0.1 + (1.5 / float32(1e10))))))) {
                                if (3.14159 > 0.0) {
                                    fmt.Println(v6)
                                    if false {
                                        v37 := []int{fi32, (fi32 * fi32)}
                                        fmt.Println(v18, v18)
                                        if len(v37) > 0 {
                                            v37[0] = fi32
                                        }
                                        _ = v37
                                    }
                                    fmt.Println(v4)
                                } else {
                                    for fi38 := 0; fi38 < 6; fi38++ {
                                        v4 -= (v4 ^ 255)
                                        v4 = (v4 * (v4 * v4))
                                        v39 := []int{}
                                        _ = fi38
                                        _ = v39
                                    }
                                    v40 := GenS1{F0: (!true), F1: v4}
                                    v41 := GenS1{F0: (v6 == 936), F1: (v4 ^ (^(v4)))}
                                    _ = v40
                                    _ = v41
                                }
                                var v42 uint32 = 4294967295
                                fmt.Println(v1, v7)
                                _ = v42
                            }
                        }
                    } else {
                        if (!true) {
                            for fi43 := 0; fi43 < 9; fi43++ {
                                v44 := GenS2{F0: 9223372036854775807, F1: 0, F2: 0}
                                _ = fi43
                                _ = v44
                            }
                            v45 := GenS3{F0: 798}
                            _ = v45
                        }
                    }
                    for fi46 := 0; fi46 < 6; fi46++ {
                        var v47 int16 = -1
                        _ = fi46
                        _ = v47
                    }
                    _ = fi32
                }
            }
        default:
            for fi48 := 0; fi48 < 20; fi48++ {
                fmt.Println(v1, v6, v7)
                v49 := GenS3{F0: (fi48 ^ fi48)}
                var v50 int8 = v7
                _ = fi48
                _ = v49
                _ = v50
            }
        }
    }
    v51 := GenS2{F0: -2147483648, F1: 4294967295, F2: 875}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v7
}
