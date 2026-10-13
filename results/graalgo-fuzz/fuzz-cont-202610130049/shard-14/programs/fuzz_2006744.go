package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int16; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 int16 }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{}
    v3 := make(map[int]int)
    var v4 int8 = -128
    v5 := []int{32767, 4294967295, 32767}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        fmt.Println(v4, v1)
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := v3[2147483647]
            v8 := GenS1{F0: 18446744073709551615}
            v1 = ("a" + "")
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint8 = 1
            v4 += -128
            _ = fi9
            _ = v10
        }
    }
    if len(v5) > 0 {
        v5[0] = 32767
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 int = 32768
    var v12 uint32 = 0
    switch (v11 >> ((603) & 63)) {
    case -2147483648:
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(v4)
            _ = fi13
        }
    case 0:
        var v14 uint = 0
        _ = v14
    default:
        switch "a" {
        case "foo bar":
            var v15 int = (v11 + (^(v11)))
            _ = v15
        case "hello":
            for fi16 := 0; fi16 < 11; fi16++ {
                fmt.Println(v12)
                v17 := v3[v11]
                v18 := GenS2{F0: 0}
                _ = fi16
                _ = v17
                _ = v18
            }
        default:
            if ((!false) || true) {
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 int64 = 9223372036854775807
                    for fi21 := 0; fi21 < 10; fi21++ {
                        v22 := GenS2{F0: (v20 | 0)}
                        v23 := GenS3{F0: 4294967295, F1: 18446744073709551615, F2: (^(v12))}
                        switch (^(fi21)) {
                        case 4294967295:
                            if len(v5) > 0 {
                                v5[0] = (fi21 * 9223372036854775807)
                            }
                        case -744:
                            v24 := GenS4{F0: 1, F1: 283, F2: v12}
                            _ = v24
                        case -657:
                            var v25 int8 = v4
                            _ = v25
                        default:
                            switch v1 {
                            case "hello":
                                var v26 uint8 = 255
                                _ = v26
                            case "foo bar":
                                var v27 uint = 469
                                _ = v27
                            case "a":
                                switch (^(v11)) {
                                case 1:
                                    v20 = v20
                                    var v28 string = ""
                                    _ = v28
                                case 32768:
                                    switch (v1 + v1) {
                                    case "":
                                        delete(v3, -706)
                                    case "foo bar":
                                        var v29 int8 = -8
                                        _ = v29
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "\n\t":
                                            v2 = append(v2, 2147483647)
                                        case "hello":
                                            fmt.Println(fi19)
                                            v12 += v12
                                        }
                                    default:
                                        var v30 int16 = -288
                                        _ = v30
                                    }
                                case -2147483648:
                                    var v31 float32 = 1.5
                                    _ = v31
                                }
                            }
                        }
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(v4, v20)
                    _ = fi19
                    _ = v20
                }
                for fi32 := 0; fi32 < 18; fi32++ {
                    var v33 float64 = float64(3.14159)
                    _ = fi32
                    _ = v33
                }
                v11 -= v11
            } else {
                for fi34 := 0; fi34 < 3; fi34++ {
                    fmt.Println(v11)
                    _ = fi34
                }
                v11 = v11
                var v35 uint = 2
                _ = v35
            }
        }
    }
    v1 = (((("a" + v1) + "foo bar") + (("" + v1) + v1)) + v1)
    var v36 uint32 = (v12 / (((v12 % (((v12 ^ (v12 >> ((18446744073709551615) & 63)))) | 1))) | 1))
    v37 := GenS5{F0: 42, F1: -1, F2: v12}
    v1 = v1
    if true {
        v3[v11] = (v11 / ((32768) | 1))
        v38 := GenS6{F0: 100}
        v11 -= -9223372036854775808
        _ = v38
    } else {
        for fi39 := 0; fi39 < 20; fi39++ {
            v36 += 1
            var v40 int8 = (v4 & 0)
            _ = fi39
            _ = v40
        }
    }
    var v41 int16 = -1
    if (0 == 255) {
        if true {
            if true {
                fmt.Println(v4, v4, v36)
                v42 := GenS7{F0: 902, F1: -93}
                v11 += (v11 >> ((601) & 63))
                _ = v42
            } else {
                for fi43 := 0; fi43 < 14; fi43++ {
                    fmt.Println(v11, v12)
                    _ = fi43
                }
                var v44 bool = false
                var v45 string = v1
                _ = v44
                _ = v45
            }
            v41 -= (^(v41))
        }
        for fi46 := 0; fi46 < 9; fi46++ {
            v11 = 42
            for fi47 := 0; fi47 < 8; fi47++ {
                v48 := GenS3{F0: 18446744073709551615, F1: 713, F2: 4294967295}
                v3[(fi47 * (v11 ^ fi47))] = (fi46 >> ((42) & 63))
                v4 = (v4 * -128)
                _ = fi47
                _ = v48
            }
            fmt.Println(v11, v1)
            _ = fi46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v5
}
