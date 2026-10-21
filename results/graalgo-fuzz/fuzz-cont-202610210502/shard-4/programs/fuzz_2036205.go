package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int64; F1 int64; F2 uint8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 string; F1 int; F2 int }
func genFn0(p0 int, p1 uint32, p2 uint64) int64 {
    for fi4 := 0; fi4 < 8; fi4++ {
        if true {
            var v5 uint8 = uint8(p1)
            _ = v5
        } else {
            if (!false) {
                var v6 int64 = -1
                v7 := []int{2147483647, 4294967295, (-(p0)), p0}
                var v8 int8 = 42
                _ = v6
                _ = v7
                _ = v8
            }
        }
        if false {
            v9 := GenS1{F0: 65535}
            v10 := []int{(fi4 % (((fi4 - (fi4 % ((p0) | 1)))) | 1)), p0}
            switch (("\n\t" + "a") + ("foo bar" + ("foo bar" + "hello"))) {
            case "hello":
                p1 = p1
            case "foo bar":
                p1 = (^(p1))
            default:
                for fi11 := 0; fi11 < 11; fi11++ {
                    if true {
                        _ = len(v10)
                        v12 := GenS1{F0: 42}
                        var v13 float64 = -2.25
                        _ = v12
                        _ = v13
                    } else {
                        if true {
                            var v14 uint16 = 17
                            _ = v14
                        } else {
                            if false {
                                v15 := GenS1{F0: uint16(p2)}
                                var v16 int16 = 1
                                p1 -= (^(p1))
                                p0 = (p0 % ((128) | 1))
                                _ = v15
                                _ = v16
                            } else {
                                _ = len(v10)
                                switch p0 {
                                case 128:
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                case 0:
                                    v17 := GenS2{F0: 2, F1: (^(p1))}
                                    _ = v17
                                default:
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                }
                            }
                            v18 := GenS1{F0: 255}
                            _ = v18
                        }
                    }
                    if true {
                        v19 := GenS3{F0: -1, F1: -9223372036854775808, F2: 94}
                        _ = len(v10)
                        _ = v19
                    } else {
                        var v20 uint8 = 75
                        if len(v10) > 0 {
                            v10[0] = int(p2)
                        }
                        _ = v20
                    }
                    _ = fi11
                }
            }
            _ = v10
            _ = v9
        } else {
            var v21 int8 = -1
            var v22 uint8 = 255
            switch fi4 {
            case 128:
                var v23 int32 = -2147483648
                _ = v23
            case 0:
                v24 := []int{127, (fi4 & (p0 >> ((2) & 63))), p0, (fi4 - int(v22))}
                _ = v24
            default:
                p1 += 4294967295
            }
            _ = v21
            _ = v22
        }
        _ = fi4
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        if (-128 != 127) {
            var v26 string = ""
            fmt.Println(v26)
            if ((p0 != (fi25 ^ (fi25 | p0))) || (892 >= 32767)) {
                v27 := []int{(p0 ^ (p0 & 42))}
                v28 := GenS4{F0: 2147483647, F1: v26}
                _ = v27
                _ = v28
            }
            _ = v26
        }
        _ = fi25
    }
    v29 := GenS3{F0: 1, F1: 84, F2: 152}
    var v30 uint = 350
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v30
    return 0
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    v31 := GenS5{F0: (!false)}
    v32 := GenS1{F0: 255}
    fmt.Println(v1)
    fmt.Println(v1)
    var v33 bool = false
    var v34 float64 = -2.25
    for fi35 := 0; fi35 < 1; fi35++ {
        if (v33 && true) {
            var v36 string = "\n\t"
            v37 := GenS6{F0: v33, F1: 539}
            _ = v36
            _ = v37
        }
        delete(v3, 42)
        v38 := GenS1{F0: 65535}
        _ = fi35
        _ = v38
    }
    for fi39 := 0; fi39 < 20; fi39++ {
        switch fi39 {
        case -733:
            for fi40 := 0; fi40 < 12; fi40++ {
                fmt.Println(fi39, fi39)
                _ = fi40
            }
        case 127:
            _ = len(v3)
        case 2147483647:
            var v41 string = ("\n\t" + (("foo bar" + ("" + ("foo bar" + "a"))) + "a"))
            _ = v41
        }
        var v42 bool = v33
        _ = fi39
        _ = v42
    }
    v43 := GenS7{F0: "hello", F1: 127, F2: 201}
    v44 := GenS2{F0: 1, F1: 4294967295}
    if (!(65535 <= 18446744073709551615)) {
        v45 := []int{-649}
        _ = v45
    } else {
        v46 := GenS4{F0: -2147483648, F1: "hello"}
        v47 := []int{128, -9223372036854775808, 95}
        v34 = (-2.25 * 1.5)
        _ = v46
        _ = v47
    }
    if v33 {
        delete(v2, -9223372036854775808)
    }
    v48 := v2[32768]
    v2[0] = 65535
    switch "foo bar" {
    case "":
        if (42 == 42) {
            if (42 >= -2147483648) {
                v49 := []int{-450}
                v48 = 2147483648
                fmt.Println(v33, v1)
                _ = v49
            }
            var v50 int32 = 1
            fmt.Println(v34, v34)
            _ = v50
        }
    }
    if (v33 && v33) {
        if v33 {
            for fi51 := 0; fi51 < 16; fi51++ {
                switch -1 {
                case -2147483648:
                    var v52 uint = 65535
                    _ = v52
                case -1:
                    for fi53 := 0; fi53 < 12; fi53++ {
                        if true {
                            v34 -= 1.5
                            switch "hello" {
                            case "hello":
                                v34 = v34
                            case "":
                                var v54 int64 = (v1 ^ 42)
                                _ = v54
                            case "a":
                                switch "foo bar" {
                                case "\n\t":
                                    v33 = ((0 != 42) || v33)
                                    switch "foo bar" {
                                    case "a":
                                        if v33 {
                                            v55 := GenS2{F0: 1, F1: v48}
                                            _ = v55
                                        }
                                    case "":
                                        switch fi51 {
                                        case 127:
                                            v1 += 42
                                        case 2147483647:
                                            v33 = (!true)
                                        case 42:
                                            v1 += (v1 >> ((597) & 63))
                                        default:
                                            var v56 string = ""
                                            _ = v56
                                        }
                                    default:
                                        v57 := []int{(fi53 - fi53), fi51, 32767, -576}
                                        _ = v57
                                    }
                                    v58 := GenS7{F0: "\n\t", F1: (fi51 + (fi53 & (fi53 ^ fi53))), F2: (^(fi53))}
                                    _ = v58
                                case "foo bar":
                                    var v59 bool = v33
                                    _ = v59
                                case "":
                                    fmt.Println(v1)
                                }
                                v1 = v1
                                _ = len(v2)
                                var v60 uint = 485
                                _ = v60
                            }
                        }
                        _ = fi53
                    }
                default:
                    v61 := []int{}
                    _ = v61
                }
                var v62 uint = 320
                for fi63 := 0; fi63 < 13; fi63++ {
                    v64 := GenS5{F0: v33}
                    _ = fi63
                    _ = v64
                }
                _ = fi51
                _ = v62
            }
        }
    } else {
        v34 = v34
        var v65 uint16 = 0
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v43
    _ = v44
    _ = v48
}
