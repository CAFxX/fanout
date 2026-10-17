package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 bool; F1 string; F2 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS8 struct { F0 uint8; F1 uint64; F2 uint8 }
func genFn0(p0 int32, p1 uint8) uint64 {
    v4 := GenS1{F0: 128, F1: 4294967295}
    v5 := GenS1{F0: 541, F1: 0}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 819
}
func genFn1() bool {
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: (fi6 & fi6), F1: 17}
        if (!false) {
            var v8 int64 = 925
            fmt.Println(fi6, fi6)
            switch fi6 {
            case -193:
                v9 := GenS2{F0: 1, F1: 32767}
                _ = v9
            default:
                var v10 bool = (!("hello" == "a"))
                _ = v10
            }
            _ = v8
        }
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    v11 := []int{32767, -687, 1, -479}
    _ = v11
    return (298 == 57)
}

func main() {
    var v1 int8 = 6
    var v2 uint16 = 410
    v3 := make(map[string]uint32)
    v12 := GenS1{F0: 42, F1: 2147483647}
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS3{F0: 1, F1: 2}
        v2 -= (v2 << ((433) & 63))
        for fi15 := 0; fi15 < 4; fi15++ {
            fmt.Println(v1, v2, fi15)
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    switch "hello" {
    case "hello":
        for fi16 := 0; fi16 < 14; fi16++ {
            v2 -= v2
            if (int16(fi16) == -1) {
                v17 := GenS2{F0: 32767, F1: (v2 - (^(v2)))}
                fmt.Println(v1, v1, fi16)
                v1 -= int8(fi16)
                _ = v17
            }
            var v18 int = -9223372036854775808
            _ = fi16
            _ = v18
        }
    case "foo bar":
        if true {
            if (1 != 4294967295) {
                if (!(9223372036854775807 > 42)) {
                    var v19 float32 = -2.25
                    switch "a" {
                    case "a":
                        fmt.Println(v1)
                    }
                    _ = v19
                } else {
                    fmt.Println(v2, v1)
                    fmt.Println(v2)
                }
                switch "a" {
                case "":
                    var v20 float32 = 0.0
                    _ = v20
                default:
                    if (0 != 972) {
                        v21 := GenS1{F0: 808, F1: 335}
                        v2 += (^(v2))
                        _ = v21
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            for fi22 := 0; fi22 < 2; fi22++ {
                                if (!(false && (true && true))) {
                                    v23 := GenS3{F0: 91, F1: 18446744073709551615}
                                    v1 += int8(v2)
                                    fmt.Println(v2, v2)
                                    _ = v23
                                }
                                _ = fi22
                            }
                        case "foo bar":
                            switch "foo bar" {
                            case "":
                                v24 := []int{1}
                                _ = v24
                            case "foo bar":
                                fmt.Println(v2, v2)
                            case "\n\t":
                                var v25 bool = true
                                _ = v25
                            default:
                                if true {
                                    var v26 uint8 = 97
                                    var v27 float64 = 1.5
                                    _ = v26
                                    _ = v27
                                } else {
                                    v28 := GenS1{F0: 0, F1: 1}
                                    v1 = v1
                                    fmt.Println(v2)
                                    _ = v28
                                }
                            }
                        default:
                            if (245 > 32767) {
                                switch 1 {
                                case 1:
                                    switch "" {
                                    case "a":
                                        v1 += (v1 << ((737) & 63))
                                    case "hello":
                                        v29 := []int{2147483647}
                                        _ = v29
                                    default:
                                        v30 := GenS4{F0: 295, F1: 1}
                                        _ = v30
                                    }
                                case 9223372036854775807:
                                    v31 := GenS5{F0: ("" != ""), F1: ("\n\t" + "foo bar"), F2: (v2 * v2)}
                                    _ = v31
                                case 2147483647:
                                    var v32 bool = (false || ((1e10 - ((1.5 + float64(1e10)) - -2.25)) >= 3.14159))
                                    _ = v32
                                default:
                                    var v33 bool = (0 <= (-(v1)))
                                    _ = v33
                                }
                            } else {
                                v34 := GenS6{F0: 42}
                                v35 := GenS2{F0: 1000, F1: (v2 + (v2 % ((v2) | 1)))}
                                _ = v34
                                _ = v35
                            }
                        }
                        fmt.Println(v2, v2)
                    }
                }
            } else {
                var v36 string = "foo bar"
                _ = v36
            }
            v37 := []int{42}
            for fi38 := 0; fi38 < 19; fi38++ {
                fmt.Println(fi38, v2)
                var v39 int = (fi38 * fi38)
                _ = fi38
                _ = v39
            }
            _ = v37
        } else {
            var v40 bool = (false || (369 <= -1))
            if (849 >= -9223372036854775808) {
                v1 += v1
                v41 := GenS7{F0: 63, F1: 635, F2: v2}
                var v42 uint64 = uint64(v2)
                _ = v41
                _ = v42
            } else {
                fmt.Println(v2, v2)
                if (291 != 2) {
                    v2 = 0
                } else {
                    for fi43 := 0; fi43 < 8; fi43++ {
                        var v44 int = fi43
                        _ = fi43
                        _ = v44
                    }
                }
            }
            v2 -= (^(v2))
            _ = v40
        }
    case "a":
        for fi45 := 0; fi45 < 6; fi45++ {
            v46 := GenS2{F0: 32767, F1: (v2 & (^(v2)))}
            _ = fi45
            _ = v46
        }
    }
    delete(v3, ("hello" + "\n\t"))
    if (2147483647 == int32(v1)) {
        v3[("a" + "\n\t")] = 2147483647
    } else {
        v2 += v2
        for fi47 := 0; fi47 < 20; fi47++ {
            switch "hello" {
            case "hello":
                v48 := []int{fi47, -58, -9223372036854775808, (fi47 ^ -1)}
                _ = v48
            case "foo bar":
                v2 += 1
            case "\n\t":
                for fi49 := 0; fi49 < 2; fi49++ {
                    v2 += 2
                    _ = fi49
                }
            }
            _ = fi47
        }
    }
    v50 := GenS8{F0: 176, F1: 42, F2: 42}
    v1 -= v1
    for fi51 := 0; fi51 < 19; fi51++ {
        for fi52 := 0; fi52 < 4; fi52++ {
            var v53 int64 = 1
            _ = fi52
            _ = v53
        }
        _ = fi51
    }
    var v54 uint64 = 4294967295
    switch "" {
    case "hello":
        v55 := []int{-9223372036854775808, 9223372036854775807, 482}
        _ = v55
    }
    if true {
        var v56 int8 = v1
        v57 := []int{0}
        _ = v56
        _ = v57
    }
    for fi58 := 0; fi58 < 4; fi58++ {
        var v59 uint8 = 239
        for fi60 := 0; fi60 < 20; fi60++ {
            v61 := GenS3{F0: 1, F1: (v54 % (((^(v54))) | 1))}
            _ = fi60
            _ = v61
        }
        if false {
            v54 = (v54 | 897)
            if true {
                var v62 int16 = 1
                _ = v62
            }
        }
        _ = fi58
        _ = v59
    }
    v63 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v63:", len(v63), v63)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v50
    _ = v54
    _ = v63
}
