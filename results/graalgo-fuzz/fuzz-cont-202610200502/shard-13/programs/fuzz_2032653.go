package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 int32; F1 int8; F2 uint }
type GenS6 struct { F0 uint16; F1 uint8 }
func genFn0(p0 bool, p1 uint32) uint8 {
    fmt.Println(p0, p0)
    switch -1 {
    case 2147483647:
        var v7 int64 = 408
        _ = v7
    }
    v8 := []int{}
    if true {
        fmt.Println(p1, p0)
        var v9 int32 = 1
        var v10 uint32 = 2147483648
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: "hello"}
        if true {
            var v12 uint = 65535
            p0 = (-85 >= 72)
            switch -9223372036854775808 {
            case -575:
                for fi13 := 0; fi13 < 17; fi13++ {
                    var v14 int64 = -9223372036854775808
                    if (p0 || true) {
                        p1 -= p1
                        for fi15 := 0; fi15 < 11; fi15++ {
                            if (fi13 <= (fi15 % ((fi15) | 1))) {
                                v8 = append(v8, (fi15 + int(p1)))
                                fmt.Println(fi13, v14, p0)
                            } else {
                                fmt.Println(fi13)
                                if (!(477 > -241)) {
                                    v14 = v14
                                    switch "foo bar" {
                                    case "":
                                        p0 = (!true)
                                    case "a":
                                        fmt.Println(v12, fi15)
                                    case "hello":
                                        var v16 int8 = -114
                                        if true {
                                            var v17 int8 = v16
                                            v8 = append(v8, 783)
                                            v8 = append(v8, (fi15 * fi15))
                                            _ = v17
                                        }
                                        _ = v16
                                    }
                                    if p0 {
                                        v12 += (^(v12))
                                        fmt.Println(v14, fi15, v12)
                                    } else {
                                        switch fi13 {
                                        case 870:
                                            v18 := GenS1{F0: "hello"}
                                            _ = v18
                                        case 1:
                                            if p0 {
                                                p0 = ((42 > -128) || p0)
                                            }
                                        default:
                                            var v19 int16 = -32768
                                            _ = v19
                                        }
                                        if (((!(((!(p0 || (1 == 94))) || ((p0 && true) && p0)) && true)) && true) || (191 >= 127)) {
                                            v20 := GenS2{F0: (fi13 & fi13), F1: 32767}
                                            v12 = (^(v12))
                                            _ = v20
                                        } else {
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(p0)
                                            case "\n\t":
                                                if len(v8) > 0 {
                                                    v8[0] = (fi15 % (((fi13 & (fi15 ^ -2147483648))) | 1))
                                                }
                                            case "":
                                                var v21 uint64 = 2
                                                _ = v21
                                            }
                                            p0 = (false || p0)
                                        }
                                        v22 := GenS2{F0: fi15, F1: 65535}
                                        _ = v22
                                    }
                                }
                                if (342 < (p1 | p1)) {
                                    if p0 {
                                        var v23 int32 = 0
                                        _ = v23
                                    } else {
                                        var v24 uint16 = 42
                                        v14 = (v14 - (-(v14)))
                                        v24 += 76
                                        var v25 uint64 = 18446744073709551615
                                        fmt.Println(p1)
                                        _ = v24
                                        _ = v25
                                    }
                                } else {
                                    p0 = (1 == -1)
                                }
                                switch "" {
                                case "":
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                }
                                v12 -= v12
                                fmt.Println(v12)
                            }
                            v26 := GenS3{F0: v14}
                            var v27 int32 = -810
                            if len(v8) > 0 {
                                v8[0] = 1
                            }
                            var v28 string = (("a" + "foo bar") + (("" + "hello") + ("" + "hello")))
                            _ = fi15
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        p0 = ((!((!(!(false || (((-128 >= 42) && p0) && p0)))) && p0)) && p0)
                    } else {
                        for fi29 := 0; fi29 < 5; fi29++ {
                            p0 = (186 <= 320)
                            switch fi13 {
                            case -856:
                                p1 = p1
                            case 9223372036854775807:
                                if len(v8) > 0 {
                                    v8[0] = int(v14)
                                }
                            default:
                                var v30 float32 = float32(0.0)
                                _ = v30
                            }
                            _ = fi29
                        }
                    }
                    _ = fi13
                    _ = v14
                }
            case -655:
                switch "foo bar" {
                case "a":
                    v31 := GenS3{F0: 0}
                    _ = v31
                case "\n\t":
                    if len(v8) > 0 {
                        v8[0] = 9223372036854775807
                    }
                }
            case 599:
                _ = len(v8)
            default:
                v32 := GenS2{F0: -445, F1: 2}
                _ = v32
            }
            _ = v12
        }
        v33 := GenS3{F0: -1}
        _ = v11
        _ = v33
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v8
    return 6
}
func genFn1(p0 int64, p1 int16) int8 {
    fmt.Println(p1, p0)
    fmt.Println(p1, p0)
    v34 := GenS4{F0: 814, F1: p0}
    _ = p0
    _ = p1
    _ = v34
    return 62
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 float64 = float64((1.5 * 1.5))
    var v4 int16 = 0
    var v5 uint64 = 978
    var v6 int = 2147483647
    v35 := GenS4{F0: 42, F1: 0}
    v36 := v1[(v6 << ((4294967295) & 63))]
    v4 = (v4 ^ (v4 & v4))
    if true {
        switch "a" {
        case "\n\t":
            var v37 int = 127
            _ = v37
        default:
            fmt.Println(v36)
        }
        var v38 int8 = genFn1(int64(v5), 1000)
        var v39 uint8 = 42
        _ = v38
        _ = v39
    }
    v40 := GenS5{F0: 0, F1: -11, F2: 65535}
    var v41 uint16 = 32767
    v42 := GenS6{F0: (v41 | v41), F1: uint8(v5)}
    switch (v36 * (v6 << ((42) & 63))) {
    case -330:
        for fi43 := 0; fi43 < 19; fi43++ {
            v36 = (v36 / ((42) | 1))
            v44 := []int{v6, fi43, 0, (fi43 << ((4294967295) & 63))}
            var v45 int8 = -17
            _ = fi43
            _ = v44
            _ = v45
        }
    default:
        v2[32767] = "hello"
    }
    v3 += v3
    v41 += (v41 / ((v41) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
