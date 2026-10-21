package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int64; F2 int64 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int32; F1 int8; F2 int }
func genFn0(p0 uint32, p1 uint8) string {
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 bool = false
        for fi9 := 0; fi9 < 6; fi9++ {
            v8 = v8
            fmt.Println(fi9, v8, fi9)
            _ = fi9
        }
        var v10 bool = v8
        _ = fi7
        _ = v10
        _ = v8
    }
    if false {
        var v11 uint16 = 1
        switch 1 {
        case 534:
            v12 := GenS1{F0: true}
            _ = v12
        default:
            var v13 float32 = 1.5
            _ = v13
        }
        if ((p0 / (((^(p0))) | 1)) != (^(p0))) {
            var v14 bool = ((426 >= 95) && false)
            switch 4294967295 {
            case -725:
                if (2 != 994) {
                    fmt.Println(p0, v11, p1)
                    v15 := GenS2{F0: (p0 | 632), F1: -9223372036854775808, F2: 42}
                    v16 := []int{-2147483648, 2147483647, 263}
                    _ = v15
                    _ = v16
                } else {
                    if v14 {
                        var v17 bool = ((false && v14) && ("hello" >= "a"))
                        for fi18 := 0; fi18 < 7; fi18++ {
                            var v19 float64 = 0.0
                            if true {
                                p1 = p1
                            }
                            _ = fi18
                            _ = v19
                        }
                        _ = v17
                    } else {
                        for fi20 := 0; fi20 < 15; fi20++ {
                            if (false && ("\n\t" < ("a" + "hello"))) {
                                v21 := GenS3{F0: 0, F1: (((("a" + "hello") + "hello") + ("foo bar" + "")) + "\n\t")}
                                var v22 float64 = -2.25
                                _ = v21
                                _ = v22
                            } else {
                                if (float64(1e10) >= 1e10) {
                                    v11 = v11
                                } else {
                                    var v23 uint = 65535
                                    v24 := []int{-903, (fi20 | (fi20 - (fi20 | 0))), int(v11)}
                                    v14 = false
                                    _ = v23
                                    _ = v24
                                }
                                var v25 int16 = -353
                                _ = v25
                            }
                            _ = fi20
                        }
                        for fi26 := 0; fi26 < 7; fi26++ {
                            var v27 uint64 = 0
                            v27 += (v27 % ((2) | 1))
                            _ = fi26
                            _ = v27
                        }
                    }
                    switch -1 {
                    case -9223372036854775808:
                        for fi28 := 0; fi28 < 11; fi28++ {
                            v14 = true
                            _ = fi28
                        }
                    default:
                        v29 := GenS2{F0: 42, F1: 9223372036854775807, F2: 1}
                        _ = v29
                    }
                }
            case 1:
                v30 := GenS3{F0: 953, F1: "foo bar"}
                _ = v30
            default:
                var v31 uint = 255
                _ = v31
            }
            _ = v14
        }
        _ = v11
    } else {
        p0 -= (p0 | 28)
        var v32 int8 = -67
        var v33 int64 = 1
        _ = v32
        _ = v33
    }
    fmt.Println(p0, p0)
    p0 += (p0 << ((487) & 63))
    p1 += (p1 | 25)
    _ = p0
    _ = p1
    return "\n\t"
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]string)
    v3 := []int{207}
    v4 := []int{-1, -1, 140}
    var v5 uint32 = 0
    var v6 int8 = 100
    v6 = v6
    if false {
        var v34 int8 = v6
        fmt.Println(v34, v5)
        if true {
            v1 = 1e10
            v35 := GenS4{F0: 392, F1: (v6 * v6), F2: 1}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v35
        }
        _ = v34
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        v5 += (^(v5))
        var v37 int = (-(fi36))
        v2[("" + "")] = "foo bar"
        _ = fi36
        _ = v37
    }
    v6 += (^(v6))
    var v38 int8 = (v6 << ((1) & 63))
    fmt.Println(v6)
    fmt.Println(v6, v5, v1)
    v39 := GenS2{F0: (v5 * v5), F1: 859, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
