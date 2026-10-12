package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int; F2 uint32 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 uint8 = 10
    var v4 int8 = 42
    _ = len(v2)
    switch -2147483648 {
    case 2147483647:
        var v5 int8 = -67
        _ = v5
    case 128:
        v4 += 127
    case 32768:
        switch "hello" {
        case "":
            switch "foo bar" {
            case "a":
                if false {
                    for fi6 := 0; fi6 < 9; fi6++ {
                        fmt.Println(v3, v3)
                        fmt.Println(fi6, v4)
                        var v7 uint = 18446744073709551615
                        _ = fi6
                        _ = v7
                    }
                    v4 = 42
                } else {
                    v4 = (v4 << ((4294967295) & 63))
                    for fi8 := 0; fi8 < 3; fi8++ {
                        v9 := v2["a"]
                        var v10 uint8 = (v3 | v3)
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    v3 = 141
                }
            case "hello":
                if (true || ((!false) || false)) {
                    if (0 < 278) {
                        for fi11 := 0; fi11 < 16; fi11++ {
                            if true {
                                var v12 uint = 0
                                var v13 uint16 = 145
                                _ = v12
                                _ = v13
                            } else {
                                if true {
                                    var v14 uint8 = v3
                                    var v15 uint32 = 2147483647
                                    if true {
                                        v16 := []int{fi11, -5}
                                        var v17 string = "hello"
                                        if true {
                                            var v18 int = (fi11 - (fi11 & (fi11 / ((fi11) | 1))))
                                            v14 += (v3 - 128)
                                            var v19 uint16 = 871
                                            _ = v18
                                            _ = v19
                                        } else {
                                            v15 = v15
                                            v20 := v2[((v17 + (v17 + "")) + "")]
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                            _ = v20
                                        }
                                        _ = v16
                                        _ = v17
                                    } else {
                                        for fi21 := 0; fi21 < 20; fi21++ {
                                            var v22 int8 = v4
                                            v23 := GenS1{F0: 0, F1: 9223372036854775807}
                                            _ = fi21
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                    _ = v14
                                    _ = v15
                                }
                                var v24 int16 = 32767
                                var v25 int64 = 9223372036854775807
                                _ = v24
                                _ = v25
                            }
                            _ = fi11
                        }
                        v3 = (v3 % (((v3 | (^(v3)))) | 1))
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v27 := GenS2{F0: (^(fi26)), F1: 401, F2: 0}
                            _ = fi26
                            _ = v27
                        }
                    }
                    v28 := []int{}
                    _ = v28
                }
            default:
                if true {
                    v29 := GenS1{F0: 635, F1: 841}
                    for fi30 := 0; fi30 < 10; fi30++ {
                        for fi31 := 0; fi31 < 18; fi31++ {
                            v3 -= v3
                            var v32 int32 = 2147483647
                            _ = fi31
                            _ = v32
                        }
                        switch 32767 {
                        case -728:
                            v4 = (v4 % (((v4 / ((104) | 1))) | 1))
                        default:
                            for fi33 := 0; fi33 < 2; fi33++ {
                                fmt.Println(fi33)
                                fmt.Println(fi30, fi33, fi33)
                                v4 = (v4 + -1)
                                var v34 float64 = float64((3.14159 / 0.1))
                                _ = fi33
                                _ = v34
                            }
                        }
                        v35 := GenS3{F0: 141, F1: int(v4), F2: 243}
                        _ = fi30
                        _ = v35
                    }
                    _ = v29
                }
            }
        }
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        v37 := GenS2{F0: (fi36 ^ -9223372036854775808), F1: -347, F2: 18446744073709551615}
        fmt.Println(fi36, v3, fi36)
        v4 = v4
        _ = fi36
        _ = v37
    }
    fmt.Println(v4, v4)
    v3 += 128
    fmt.Println(v3, v3)
    v4 += v4
    if (143 <= 127) {
        var v38 uint64 = 2
        v39 := v2[("hello" + "")]
        _ = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 15; fi40++ {
        for fi41 := 0; fi41 < 2; fi41++ {
            v42 := []int{(fi40 - fi41), fi40}
            _ = fi41
            _ = v42
        }
        _ = len(v1)
        _ = fi40
    }
    if (127 >= 9223372036854775807) {
        v43 := []int{}
        var v44 int16 = 639
        _ = v43
        _ = v44
    }
    fmt.Println(v4)
    for fi45 := 0; fi45 < 13; fi45++ {
        var v46 uint8 = 129
        _ = fi45
        _ = v46
    }
    var v47 bool = (112 >= 176)
    if (v47 || false) {
        v4 = v4
    }
    for fi48 := 0; fi48 < 18; fi48++ {
        for fi49 := 0; fi49 < 10; fi49++ {
            v50 := GenS1{F0: (fi49 - 128), F1: -72}
            _ = fi49
            _ = v50
        }
        if ((v3 + v3) != (v3 - v3)) {
            var v51 string = "a"
            _ = v51
        }
        if v47 {
            var v52 uint32 = uint32(v3)
            if v47 {
                var v53 float64 = ((1.5 * 1.5) / 1.0)
                _ = v53
            } else {
                var v54 uint = 1
                _ = v54
            }
            _ = v52
        } else {
            var v55 int = (fi48 & fi48)
            switch "foo bar" {
            case "hello":
                v55 = (fi48 % ((664) | 1))
            case "foo bar":
                v56 := GenS3{F0: 109, F1: v55, F2: 214}
                _ = v56
            }
            v57 := v1[v55]
            _ = v55
            _ = v57
        }
        _ = fi48
    }
    v4 = (v4 & 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
}
