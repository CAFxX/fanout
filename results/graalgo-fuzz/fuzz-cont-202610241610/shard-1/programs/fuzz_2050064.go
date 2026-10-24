package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 int64; F1 int64; F2 int8 }
type GenS3 struct { F0 int16; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint16; F1 int16; F2 int }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 bool; F1 int; F2 uint8 }
func genFn0(p0 uint16, p1 int64) uint16 {
    v6 := []int{127, 9223372036854775807}
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint32 = 2
        v8 += v8
        _ = fi7
        _ = v8
    }
    var v9 bool = (true && true)
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return p0
}
func genFn1(p0 int) uint16 {
    v10 := []int{p0, (p0 & p0)}
    switch 32768 {
    case 152:
        if false {
            fmt.Println(p0)
            var v11 uint32 = 4294967295
            _ = v11
        } else {
            var v12 float32 = (3.14159 * 0.1)
            fmt.Println(p0)
            _ = v12
        }
    default:
        switch "a" {
        case "a":
            if (1 >= 1) {
                switch -795 {
                case 32767:
                    for fi13 := 0; fi13 < 16; fi13++ {
                        p0 = (-(p0))
                        _ = fi13
                    }
                case 127:
                    if (231 <= -185) {
                        v14 := GenS1{F0: 53, F1: 128, F2: -1}
                        _ = v14
                    }
                case 0:
                    v15 := GenS1{F0: 42, F1: 2, F2: 42}
                    _ = v15
                }
            }
        case "":
            var v16 int = -1
            _ = v16
        case "\n\t":
            var v17 float32 = (0.0 * 0.0)
            _ = v17
        }
    }
    _ = p0
    _ = v10
    return genFn0(350, 0)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    var v4 uint = 669
    v5 := []uint32{841, 2147483648, 752, 2147483647}
    v4 -= 57
    if false {
        var v18 uint16 = 65535
        _ = v18
    } else {
        v19 := GenS1{F0: 0, F1: 26, F2: 31}
        v4 -= v4
        _ = v19
    }
    v5 = append(v5, 0)
    if (0 <= -459) {
        fmt.Println(v4)
    } else {
        var v20 uint16 = 38
        _ = v20
    }
    var v21 int = 128
    v2["hello"] = 330
    v22 := GenS2{F0: 0, F1: 0, F2: -128}
    var v23 int32 = 42
    v4 = (v4 + 42)
    if true {
        var v24 string = "hello"
        _ = v24
    } else {
        var v25 int32 = (^(v23))
        var v26 uint32 = 896
        _ = v25
        _ = v26
    }
    if (false || (!(false && true))) {
        var v27 int64 = -77
        _ = v27
    } else {
        switch "hello" {
        case "hello":
            var v28 bool = true
            _ = v28
        default:
            var v29 uint64 = 107
            _ = v29
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            if true {
                v31 := v3[2147483648]
                v31 = (!v31)
                _ = v31
            } else {
                var v32 float32 = 0.1
                switch (v21 >> ((v4) & 63)) {
                case 203:
                    var v33 float64 = 1e10
                    _ = v33
                case 0:
                    if false {
                        for fi34 := 0; fi34 < 6; fi34++ {
                            v4 = 65535
                            if (4294967295 < 737) {
                                v23 = (v23 / (((v23 ^ (v23 & -841))) | 1))
                            } else {
                                v32 += v32
                                switch -1 {
                                case 127:
                                    if true {
                                        var v35 int32 = v23
                                        _ = len(v5)
                                        fmt.Println(v32)
                                        _ = v35
                                    } else {
                                        v5 = append(v5, 2147483648)
                                        v21 += -160
                                        v36 := GenS3{F0: 324, F1: 128, F2: -64}
                                        v23 -= (v23 * -2147483648)
                                        var v37 int32 = (-(v23))
                                        v21 = (fi30 / (((fi30 + (fi34 % ((9223372036854775807) | 1)))) | 1))
                                        switch fi30 {
                                        case 32768:
                                            fmt.Println(v4, v4, fi30)
                                        case 2147483647:
                                            v38 := GenS4{F0: genFn1(v21), F1: 42, F2: fi30}
                                            _ = v38
                                        case 675:
                                            delete(v3, 2147483647)
                                        default:
                                            fmt.Println(v4)
                                        }
                                        _ = v36
                                        _ = v37
                                    }
                                }
                                v39 := GenS1{F0: 456, F1: uint8(v4), F2: 127}
                                _ = v39
                            }
                            _ = fi34
                        }
                    }
                }
                _ = v32
            }
            var v40 int = (v21 * -87)
            fmt.Println(v23)
            _ = fi30
            _ = v40
        }
    }
    switch v21 {
    case 435:
        switch "\n\t" {
        case "foo bar":
            delete(v2, ("" + "foo bar"))
        case "a":
            v41 := GenS5{F0: true, F1: v4}
            _ = v41
        }
    case -649:
        v42 := GenS5{F0: false, F1: 93}
        _ = v42
    case 710:
        v4 += 699
    default:
        var v43 int8 = -27
        _ = v43
    }
    var v44 int32 = (v23 << ((42) & 63))
    v45 := GenS6{F0: 42, F1: 2147483647}
    for fi46 := 0; fi46 < 8; fi46++ {
        switch ("a" + ("a" + "")) {
        case "a":
            v5 = append(v5, 909)
        case "hello":
            var v47 float64 = 1.5
            _ = v47
        case "":
            v48 := GenS7{F0: (42 == (v4 >> (((^(v4))) & 63))), F1: fi46, F2: 39}
            _ = v48
        default:
            fmt.Println(v4, v21, fi46)
        }
        _ = fi46
    }
    fmt.Println(v44)
    fmt.Println(v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
