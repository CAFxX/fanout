package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint32; F2 int64 }
func genFn0(p0 int16, p1 int8, p2 int) int16 {
    p2 -= 42
    switch (p2 % (((p2 + 128)) | 1)) {
    case 4294967295:
        switch p2 {
        case -9223372036854775808:
            v7 := GenS1{F0: 1}
            _ = v7
        }
    default:
        v8 := []int{32767, (p2 & (-(p2)))}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 | (p0 * (p0 + (^(p0)))))
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    var v4 float64 = 1.5
    var v5 uint = 1
    v6 := make(map[uint32]bool)
    var v9 uint = (v5 >> ((v5) & 63))
    v9 += v5
    fmt.Println(v4, v5, v4)
    v5 += (v5 ^ 203)
    v9 = (v9 / (((v9 / ((4294967295) | 1))) | 1))
    var v10 uint32 = 4294967295
    switch "a" {
    case "foo bar":
        switch "\n\t" {
        case "a":
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 int32 = -2147483648
                switch "foo bar" {
                case "a":
                    if (-9223372036854775808 >= fi11) {
                        _ = len(v2)
                        delete(v2, v10)
                    } else {
                        for fi13 := 0; fi13 < 4; fi13++ {
                            switch -778 {
                            case -55:
                                v14 := GenS2{F0: 1}
                                _ = v14
                            case 128:
                                v2[v10] = (fi13 * fi11)
                            case -2147483648:
                                v10 = (v10 & v10)
                            default:
                                var v15 float32 = 0.1
                                _ = v15
                            }
                            var v16 int16 = v1
                            _ = fi13
                            _ = v16
                        }
                        fmt.Println(v12, v10)
                    }
                case "\n\t":
                    fmt.Println(v5, v10, v9)
                }
                var v17 int16 = v1
                _ = fi11
                _ = v12
                _ = v17
            }
        case "\n\t":
            v18 := GenS3{F0: v10, F1: (!true)}
            _ = v18
        case "hello":
            v19 := GenS2{F0: 42}
            _ = v19
        default:
            var v20 int = int(v1)
            _ = v20
        }
    default:
        if true {
            var v21 int32 = -2147483648
            v5 -= (v5 >> ((424) & 63))
            _ = v21
        } else {
            var v22 uint = (v5 << ((255) & 63))
            _ = v22
        }
    }
    v9 -= v5
    v5 = 4294967295
    fmt.Println(v5)
    v4 = -2.25
    var v23 int16 = (-(v1))
    v24 := []int{9223372036854775807, 2147483647}
    v10 -= (^(v10))
    for fi25 := 0; fi25 < 13; fi25++ {
        v4 = (((v4 / (1.5 + v4)) + 1e10) / v4)
        var v26 string = ""
        _ = fi25
        _ = v26
    }
    if (9223372036854775807 > 1) {
        var v27 uint64 = 680
        for fi28 := 0; fi28 < 4; fi28++ {
            v1 -= (v23 | (v1 ^ genFn0((v1 / ((genFn0(42, 42, (fi28 | fi28))) | 1)), 84, fi28)))
            v2[1] = fi28
            switch fi28 {
            case -744:
                var v29 int8 = 118
                _ = v29
            case 32767:
                if false {
                    if (v27 <= v27) {
                        fmt.Println(v1)
                        v2[(v10 - 65535)] = fi28
                    }
                } else {
                    for fi30 := 0; fi30 < 18; fi30++ {
                        var v31 int8 = 42
                        _ = len(v24)
                        _ = fi30
                        _ = v31
                    }
                    fmt.Println(v10, v9)
                }
            case -487:
                v1 -= (v1 >> ((65535) & 63))
            }
            _ = fi28
        }
        if (0 != int8(v27)) {
            v4 += v4
            switch -317 {
            case 128:
                if false {
                    var v32 uint = v5
                    v33 := GenS1{F0: 128}
                    fmt.Println(v10, v23)
                    _ = v32
                    _ = v33
                } else {
                    var v34 uint16 = 626
                    var v35 uint32 = v10
                    _ = v34
                    _ = v35
                }
            case 42:
                var v36 uint8 = 42
                _ = v36
            case -1:
                for fi37 := 0; fi37 < 16; fi37++ {
                    switch "\n\t" {
                    case "hello":
                        if false {
                            v38 := GenS2{F0: 42}
                            var v39 string = ("hello" + (("a" + "foo bar") + ((("\n\t" + ("a" + ("" + ("a" + "")))) + "foo bar") + ("foo bar" + "a"))))
                            _ = v38
                            _ = v39
                        }
                    default:
                        switch fi37 {
                        case 127:
                            var v40 int = 760
                            _ = v40
                        case 2147483647:
                            v9 -= (v9 % ((v9) | 1))
                        default:
                            if false {
                                switch fi37 {
                                case -9223372036854775808:
                                    var v41 uint = 1
                                    _ = v41
                                case -698:
                                    if true {
                                        v23 -= genFn0((^(v23)), -128, fi37)
                                        v4 = ((1.5 + v4) - v4)
                                        v42 := GenS4{F0: (v5 ^ uint(v27)), F1: (v10 / ((302) | 1)), F2: 0}
                                        _ = v42
                                    }
                                case -956:
                                    var v43 int8 = 127
                                    _ = v43
                                }
                                switch (-(fi37)) {
                                case 0:
                                    v10 += (v10 & (v10 / ((v10) | 1)))
                                default:
                                    var v44 int16 = 42
                                    _ = v44
                                }
                                if (!((true || (fi37 != (-(fi37)))) || false)) {
                                    if ((v27 > v27) || ((255 != 255) && true)) {
                                        var v45 int8 = 100
                                        fmt.Println(v5)
                                        _ = v45
                                    } else {
                                        var v46 string = ("\n\t" + "\n\t")
                                        v47 := GenS1{F0: 255}
                                        _ = v46
                                        _ = v47
                                    }
                                    fmt.Println(v10)
                                    v48 := GenS5{F0: 596}
                                    _ = v48
                                }
                            }
                        }
                    }
                    switch "" {
                    case "":
                        v49 := v2[37]
                        _ = v49
                    case "foo bar":
                        switch fi37 {
                        case 4294967295:
                            _ = len(v24)
                        }
                    default:
                        v50 := GenS3{F0: v10, F1: false}
                        _ = v50
                    }
                    _ = fi37
                }
            }
        } else {
            v51 := GenS6{F0: ("hello" + "\n\t"), F1: v10, F2: 9223372036854775807}
            v4 -= ((v4 - (v4 / ((v4 + (v4 * float64((1e10 * 3.14159)))) / ((1e10 - float64(1e10)) + v4)))) + ((v4 / (v4 / (v4 + 1.5))) - v4))
            _ = v51
        }
        _ = v27
    } else {
        v9 = v5
        v9 = 42
        var v52 float32 = 0.1
        _ = v52
    }
    v9 = 174
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
