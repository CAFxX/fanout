package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS5 struct { F0 int64; F1 int; F2 bool }
func genFn0() uint {
    var v6 int = 32768
    var v7 uint16 = 0
    v6 = (v6 % ((v6) | 1))
    v8 := GenS1{F0: -32768, F1: "", F2: 1}
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint32 = 42
    var v2 uint64 = 2
    v3 := make(map[string]string)
    var v4 int32 = -6
    v5 := make(map[uint32]string)
    if false {
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(v1, v4, v1)
            if (!(v2 == v2)) {
                _ = len(v5)
                if (217 != 42) {
                    v2 = (v2 ^ v2)
                }
            } else {
                var v10 int32 = (v4 >> ((2) & 63))
                _ = v10
            }
            _ = fi9
        }
        v2 += v2
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS2{F0: 754, F1: 152}
        for fi13 := 0; fi13 < 13; fi13++ {
            v4 -= v4
            _ = fi13
        }
        v1 += (^(v1))
        _ = fi11
        _ = v12
    }
    v14 := []int{1, 9223372036854775807, -9223372036854775808}
    if (0.1 <= 1.5) {
        switch "\n\t" {
        case "":
            v1 = v1
        case "foo bar":
            v15 := GenS1{F0: 42, F1: "", F2: v4}
            _ = v15
        case "\n\t":
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int = fi16
                var v18 int = v17
                var v19 int8 = -41
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        default:
            if (!false) {
                var v20 float64 = (((((float64(1.5) / 1.0) / float64(-2.25)) * (3.14159 * 0.1)) - 0.0) - 1e10)
                fmt.Println(v1)
                _ = v20
            } else {
                if ((1e10 > 1e10) || (-1 <= 1000)) {
                    switch -375 {
                    case -322:
                        var v21 float32 = -2.25
                        _ = v21
                    case 9223372036854775807:
                        v2 = 2
                    default:
                        v4 = (v4 - (-(v4)))
                    }
                }
            }
        }
    } else {
        v22 := GenS2{F0: 1, F1: 255}
        if true {
            if false {
                var v23 string = ""
                v3[(v23 + v23)] = v23
                if (1 >= 1) {
                    v23 = v23
                    v1 -= v1
                    for fi24 := 0; fi24 < 17; fi24++ {
                        var v25 float64 = 1.5
                        v25 -= v25
                        switch 42 {
                        case 1:
                            v26 := GenS3{F0: v1}
                            _ = v26
                        case -663:
                            v27 := GenS3{F0: (v1 >> ((1) & 63))}
                            _ = v27
                        case 111:
                            switch -236 {
                            case 347:
                                var v28 float32 = (1.5 / 3.14159)
                                _ = v28
                            case 0:
                                _ = len(v14)
                            default:
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            }
                        }
                        _ = fi24
                        _ = v25
                    }
                } else {
                    switch 20 {
                    case -2147483648:
                        switch v23 {
                        case "foo bar":
                            v2 = (v2 << ((255) & 63))
                        default:
                            v23 = v23
                        }
                    }
                    for fi29 := 0; fi29 < 16; fi29++ {
                        var v30 int8 = -128
                        if true {
                            fmt.Println(fi29)
                        } else {
                            var v31 int8 = 1
                            v30 = 100
                            _ = v31
                        }
                        var v32 int16 = 32767
                        _ = fi29
                        _ = v30
                        _ = v32
                    }
                    if true {
                        _ = len(v5)
                        var v33 int32 = (-(v4))
                        _ = v33
                    }
                }
                _ = v23
            }
            v2 += v2
        }
        _ = v22
    }
    v4 -= (v4 * int32(v2))
    for fi34 := 0; fi34 < 4; fi34++ {
        v2 -= v2
        for fi35 := 0; fi35 < 18; fi35++ {
            v36 := GenS4{F0: 0, F1: 127, F2: (v2 & (v2 ^ 18446744073709551615))}
            v37 := GenS5{F0: -979, F1: fi34, F2: false}
            fmt.Println(v1, v2, fi35)
            _ = fi35
            _ = v36
            _ = v37
        }
        delete(v3, "foo bar")
        _ = fi34
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    switch "foo bar" {
    case "":
        v2 -= 2
    case "a":
        if (0 != v4) {
            if (((-2.25 * 0.0) <= float32((-2.25 + ((0.1 / (float64(1e10) * 3.14159)) - 0.0)))) && ("\n\t" > "\n\t")) {
                fmt.Println(v4)
            } else {
                switch "a" {
                case "foo bar":
                    v38 := GenS5{F0: 1, F1: 128, F2: (true && (("a" <= "foo bar") && true))}
                    _ = v38
                case "hello":
                    switch "\n\t" {
                    case "foo bar":
                        v39 := GenS5{F0: 1, F1: -2147483648, F2: true}
                        _ = v39
                    case "\n\t":
                        for fi40 := 0; fi40 < 10; fi40++ {
                            if (fi40 >= (^(fi40))) {
                                for fi41 := 0; fi41 < 14; fi41++ {
                                    switch 9223372036854775807 {
                                    case -175:
                                        var v42 int64 = -9223372036854775808
                                        _ = v42
                                    case 0:
                                        v43 := v5[v1]
                                        _ = v43
                                    default:
                                        v2 += (v2 ^ 402)
                                    }
                                    var v44 int32 = (v4 - v4)
                                    _ = fi41
                                    _ = v44
                                }
                                _ = len(v14)
                            }
                            _ = fi40
                        }
                    case "a":
                        v2 = v2
                    }
                default:
                    v1 = (v1 * (v1 & v1))
                }
            }
        }
    }
    var v45 int8 = -20
    var v46 uint32 = (^(v1))
    for fi47 := 0; fi47 < 12; fi47++ {
        if (!((!(128 < 157)) && false)) {
            if ("foo bar" != "hello") {
                v14 = append(v14, (fi47 >> ((654) & 63)))
                v1 += (v46 & 0)
            }
        } else {
            _ = len(v14)
        }
        v4 += (v4 % (((-(v4))) | 1))
        _ = fi47
    }
    for fi48 := 0; fi48 < 3; fi48++ {
        v46 += (v1 >> ((42) & 63))
        var v49 int64 = -377
        _ = fi48
        _ = v49
    }
    if false {
        v4 = v4
        var v50 int64 = 307
        _ = v50
    }
    for fi51 := 0; fi51 < 15; fi51++ {
        var v52 bool = true
        _ = fi51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
