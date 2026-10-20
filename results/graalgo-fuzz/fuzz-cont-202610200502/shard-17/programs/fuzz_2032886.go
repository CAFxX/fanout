package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 int }
func genFn0() uint8 {
    v5 := GenS1{F0: (9223372036854775807 >= 9223372036854775807)}
    var v6 uint8 = 1
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 uint) int16 {
    switch "foo bar" {
    case "hello":
        p0 += (p0 % ((p0) | 1))
    case "":
        v7 := []int{}
        _ = v7
    case "foo bar":
        p0 = p0
    }
    var v8 uint8 = 255
    v9 := GenS2{F0: 1, F1: (p0 + 255)}
    v8 = 255
    _ = p0
    _ = v8
    _ = v9
    return -613
}

func main() {
    var v1 int8 = -9
    var v2 string = ("hello" + ("\n\t" + ("hello" + "\n\t")))
    v3 := []uint32{370, 4294967295, 4294967295, 1}
    v4 := []uint32{1}
    v10 := GenS2{F0: -9223372036854775808, F1: 4294967295}
    switch (v2 + ((v2 + v2) + (v2 + "a"))) {
    case "":
        v11 := GenS3{F0: genFn1(4294967295), F1: 2}
        _ = v11
    default:
        switch 9223372036854775807 {
        case 1:
            switch 792 {
            case 32767:
                var v12 int16 = 0
                _ = v12
            case 128:
                fmt.Println(v1, v1)
            case 42:
                fmt.Println(v2)
            default:
                var v13 int16 = -32768
                _ = v13
            }
        case 9223372036854775807:
            if (255 <= 83) {
                for fi14 := 0; fi14 < 17; fi14++ {
                    if false {
                        var v15 uint = 556
                        v1 = v1
                        if (!((42 >= 128) && true)) {
                            switch (v2 + v2) {
                            case "hello":
                                var v16 int64 = 42
                                _ = v16
                            case "\n\t":
                                v17 := GenS2{F0: 42, F1: v15}
                                _ = v17
                            case "":
                                var v18 float32 = -2.25
                                _ = v18
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            v15 = (v15 - v15)
                            for fi19 := 0; fi19 < 15; fi19++ {
                                switch (fi19 << (((v15 % ((v15) | 1))) & 63)) {
                                case -46:
                                    v3 = append(v3, 2)
                                    var v20 uint16 = 626
                                    v15 = (^(v15))
                                    fmt.Println(v20, fi14, v1)
                                    switch "a" {
                                    case "":
                                        var v21 int8 = v1
                                        _ = v21
                                    }
                                    _ = v20
                                default:
                                    switch "" {
                                    case "":
                                        var v22 int32 = 2147483647
                                        _ = v22
                                    case "foo bar":
                                        v23 := GenS4{F0: -1}
                                        _ = v23
                                    case "a":
                                        var v24 int64 = 468
                                        _ = v24
                                    default:
                                        if (fi14 <= (fi14 << ((v15) & 63))) {
                                            fmt.Println(v1)
                                            v25 := GenS5{F0: -1, F1: 2}
                                            _ = v25
                                        }
                                    }
                                }
                                _ = fi19
                            }
                        } else {
                            v1 += v1
                            v15 = (v15 >> (((v15 | (^(v15)))) & 63))
                            for fi26 := 0; fi26 < 13; fi26++ {
                                switch (fi26 & fi14) {
                                case 0:
                                    v27 := GenS6{F0: -1, F1: 261}
                                    _ = v27
                                case -715:
                                    var v28 string = "\n\t"
                                    _ = v28
                                }
                                _ = fi26
                            }
                        }
                        _ = v15
                    } else {
                        v1 -= (-(v1))
                        v1 -= v1
                    }
                    v29 := GenS4{F0: -2147483648}
                    fmt.Println(v2, v2)
                    _ = fi14
                    _ = v29
                }
                v30 := GenS4{F0: 1}
                v2 = v2
                _ = v30
            }
        default:
            v31 := GenS5{F0: 2147483647, F1: 18446744073709551615}
            _ = v31
        }
    }
    if len(v3) > 0 {
        v3[0] = 471
    }
    _ = len(v3)
    v2 = (v2 + "foo bar")
    var v32 uint32 = 122
    fmt.Println(v2)
    v33 := GenS2{F0: -505, F1: 0}
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 string = (v2 + v2)
        var v36 int16 = -32768
        _ = fi34
        _ = v35
        _ = v36
    }
    if false {
        switch "" {
        case "":
            v37 := GenS7{F0: 32767}
            _ = v37
        }
        v4 = append(v4, (v32 / (((v32 * v32)) | 1)))
    } else {
        fmt.Println(v2, v32, v1)
        v38 := GenS2{F0: 115, F1: 18446744073709551615}
        _ = v38
    }
    var v39 bool = (false && true)
    v40 := GenS1{F0: v39}
    var v41 int16 = 0
    _ = len(v4)
    if ((-1 != -158) && (-977 >= -2147483648)) {
        fmt.Println(v1, v1, v39)
        switch v2 {
        case "a":
            fmt.Println(v1)
        case "foo bar":
            var v42 uint8 = 127
            _ = v42
        case "hello":
            v39 = v39
        }
        v2 = "\n\t"
    } else {
        for fi43 := 0; fi43 < 16; fi43++ {
            var v44 string = v2
            var v45 int = fi43
            _ = fi43
            _ = v44
            _ = v45
        }
        var v46 bool = ((!v39) || v39)
        _ = v46
    }
    var v47 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v47
}
