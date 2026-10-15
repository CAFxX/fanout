package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint; F2 int }
type GenS4 struct { F0 int16; F1 int8; F2 bool }
type GenS5 struct { F0 int32; F1 int16; F2 bool }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 int64 }
func genFn0(p0 uint8, p1 uint32) bool {
    v5 := []int{}
    fmt.Println(p0, p0)
    if len(v5) > 0 {
        v5[0] = 1
    }
    var v6 int8 = 100
    v7 := GenS1{F0: 32767, F1: p1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (0.1 <= 0.1)
}

func main() {
    var v1 int16 = 762
    var v2 float64 = 3.14159
    v3 := []uint32{446, 590, 219}
    v4 := make(map[string]int)
    if false {
        if len(v3) > 0 {
            v3[0] = 417
        }
        if false {
            for fi8 := 0; fi8 < 18; fi8++ {
                v1 = v1
                switch "a" {
                case "\n\t":
                    var v9 string = ""
                    _ = v9
                default:
                    for fi10 := 0; fi10 < 5; fi10++ {
                        var v11 int32 = 1
                        var v12 bool = ((v11 % (((v11 >> ((0) & 63))) | 1)) <= (v11 | (-(v11))))
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
                _ = fi8
            }
            v2 -= v2
            var v13 int64 = -9223372036854775808
            _ = v13
        } else {
            v14 := GenS1{F0: 221, F1: 518}
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS1{F0: 1, F1: 619}
        _ = fi15
        _ = v16
    }
    switch 1 {
    case 128:
        for fi17 := 0; fi17 < 5; fi17++ {
            if false {
                var v18 uint = 65535
                _ = v18
            } else {
                switch fi17 {
                case -149:
                    v2 -= v2
                case 569:
                    if len(v3) > 0 {
                        v3[0] = 101
                    }
                }
                if (276 == -1) {
                    v19 := GenS1{F0: 255, F1: 145}
                    var v20 int64 = 42
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(v1, fi17)
                }
                if false {
                    v2 = float64(1e10)
                }
            }
            _ = fi17
        }
    case 32768:
        if true {
            v1 = v1
        }
    default:
        if (0.0 < (v2 * 0.1)) {
            v3 = append(v3, 2)
        }
    }
    if (42 < 1) {
        v2 = (-2.25 - (((v2 - (v2 / 1.5)) - ((1e10 - 1e10) - 1e10)) * 3.14159))
    }
    if true {
        fmt.Println(v1, v1)
    } else {
        switch 362 {
        case -1:
            var v21 uint64 = 2
            _ = v21
        }
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    switch "\n\t" {
    case "\n\t":
        if false {
            v3 = append(v3, 4294967295)
            v2 += v2
        }
    case "hello":
        _ = len(v4)
    default:
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        v1 += v1
        switch (fi22 % ((32768) | 1)) {
        case 128:
            if (42 < 2147483647) {
                if (2 <= 411) {
                    switch 2147483647 {
                    case 696:
                        var v23 uint = 558
                        _ = v23
                    case -9223372036854775808:
                        switch -1 {
                        case 0:
                            fmt.Println(v1, v2, v1)
                        }
                    }
                } else {
                    switch fi22 {
                    case 128:
                        v24 := GenS2{F0: 2}
                        _ = v24
                    case -2147483648:
                        switch (fi22 * (fi22 << ((104) & 63))) {
                        case 127:
                            switch (fi22 >> ((255) & 63)) {
                            case 0:
                                v1 = (v1 + (v1 - (^(v1))))
                            }
                        case -1:
                            fmt.Println(v1, v2)
                        default:
                            _ = len(v4)
                        }
                    case 2147483647:
                        v1 += v1
                    default:
                        for fi25 := 0; fi25 < 3; fi25++ {
                            v26 := GenS3{F0: 65535, F1: 373, F2: (fi22 >> ((805) & 63))}
                            var v27 uint32 = 4294967295
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                    var v28 uint32 = 0
                    var v29 float64 = v2
                    _ = v28
                    _ = v29
                }
                switch "a" {
                case "\n\t":
                    if true {
                        if false {
                            fmt.Println(fi22, fi22)
                            for fi30 := 0; fi30 < 14; fi30++ {
                                v31 := GenS1{F0: 1, F1: 304}
                                v32 := GenS2{F0: 0}
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                    } else {
                        switch "foo bar" {
                        case "hello":
                            v1 = v1
                        }
                        fmt.Println(v2, v1)
                        v2 += v2
                    }
                }
                v33 := GenS1{F0: 0, F1: 42}
                _ = v33
            } else {
                v1 = (-(v1))
                if (4294967295 < 1) {
                    fmt.Println(fi22, fi22, v2)
                    v34 := GenS4{F0: (v1 ^ (v1 << ((144) & 63))), F1: 18, F2: (!false)}
                    var v35 uint64 = 999
                    _ = v34
                    _ = v35
                } else {
                    if (904 > 42) {
                        for fi36 := 0; fi36 < 5; fi36++ {
                            v1 = v1
                            var v37 uint32 = 2147483648
                            _ = fi36
                            _ = v37
                        }
                        if ((false || false) && (-1 == fi22)) {
                            if true {
                                _ = len(v4)
                            } else {
                                v2 -= v2
                            }
                            var v38 bool = (!(-2.25 < (1.5 + 1.5)))
                            fmt.Println(v38, v2, fi22)
                            _ = v38
                        }
                        if true {
                            v39 := GenS5{F0: 2147483647, F1: v1, F2: (fi22 >= (fi22 >> ((4294967295) & 63)))}
                            var v40 bool = true
                            v41 := GenS2{F0: 1}
                            _ = v39
                            _ = v40
                            _ = v41
                        } else {
                            v2 = v2
                            v42 := GenS2{F0: 2}
                            var v43 uint64 = 4294967295
                            _ = v42
                            _ = v43
                        }
                    } else {
                        var v44 uint8 = 2
                        fmt.Println(fi22, v2, fi22)
                        _ = v44
                    }
                }
                for fi45 := 0; fi45 < 11; fi45++ {
                    v46 := GenS6{F0: v1, F1: 879}
                    v2 += 3.14159
                    _ = fi45
                    _ = v46
                }
            }
        default:
            v47 := GenS7{F0: 4294967295, F1: (28 > 42)}
            _ = v47
        }
        _ = fi22
    }
    v48 := GenS3{F0: 2147483648, F1: 1, F2: 9223372036854775807}
    _ = len(v4)
    var v49 int32 = 1
    v1 += (v1 | v1)
    var v50 uint32 = 42
    v49 -= v49
    if true {
        v51 := GenS4{F0: (v1 & (v1 ^ v1)), F1: 118, F2: false}
        _ = v51
    }
    v52 := GenS8{F0: -9223372036854775808}
    var v53 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v50
    _ = v52
    _ = v53
}
