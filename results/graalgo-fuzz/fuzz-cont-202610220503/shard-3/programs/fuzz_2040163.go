package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int }
func genFn0(p0 int, p1 string) uint8 {
    v5 := GenS1{F0: 57}
    var v6 float64 = -2.25
    var v7 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 82
}
func genFn1(p0 int8, p1 bool, p2 int16) int32 {
    if p1 {
        p1 = (0 > -694)
        if p1 {
            var v8 int16 = -32768
            v9 := []int{}
            _ = v8
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := []int{}
            switch fi10 {
            case 32768:
                var v12 uint16 = 65535
                _ = v12
            }
            _ = fi10
            _ = v11
        }
    }
    var v13 bool = ((p2 & p2) < (p2 / ((p2) | 1)))
    if true {
        if p1 {
            switch ("a" + ("a" + "foo bar")) {
            case "\n\t":
                if (-488 < 42) {
                    for fi14 := 0; fi14 < 17; fi14++ {
                        var v15 int64 = 42
                        v13 = false
                        _ = fi14
                        _ = v15
                    }
                } else {
                    for fi16 := 0; fi16 < 15; fi16++ {
                        var v17 uint16 = 42
                        v18 := []int{(fi16 * (-(fi16))), (fi16 | fi16), fi16, -773}
                        switch fi16 {
                        case 127:
                            v18 = append(v18, fi16)
                        case 369:
                            v19 := GenS1{F0: -2147483648}
                            _ = v19
                        case 32768:
                            for fi20 := 0; fi20 < 16; fi20++ {
                                var v21 uint = 0
                                switch (("hello" + "foo bar") + ("\n\t" + "a")) {
                                case "":
                                    var v22 int16 = p2
                                    _ = v22
                                case "hello":
                                    var v23 int = fi20
                                    _ = v23
                                case "a":
                                    p1 = (0 <= 1)
                                default:
                                    fmt.Println(p2)
                                }
                                _ = fi20
                                _ = v21
                            }
                        default:
                            p0 += (p0 / ((p0) | 1))
                        }
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    var v24 int = int(p2)
                    v25 := GenS1{F0: int32(p0)}
                    _ = v24
                    _ = v25
                }
            }
        }
    }
    v13 = (p1 || ((p1 || p1) || (-803 > 1)))
    switch -1000 {
    case 32767:
        var v26 float32 = 0.0
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -520
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{188, -9223372036854775808, 0, 32768}
    v3 := []uint32{2, 0, 281, 487}
    v4 := make(map[uint32]bool)
    var v27 uint = 255
    if (-2147483648 < -9223372036854775808) {
        v27 += 644
    } else {
        v27 = 755
        var v28 int = -9223372036854775808
        _ = v28
    }
    var v29 int8 = -106
    _ = len(v1)
    for fi30 := 0; fi30 < 6; fi30++ {
        if (52 != 42) {
            v31 := GenS2{F0: -74, F1: -957}
            var v32 bool = true
            _ = v31
            _ = v32
        } else {
            var v33 float32 = -2.25
            v34 := GenS2{F0: 9223372036854775807, F1: 9223372036854775807}
            v33 = 1e10
            _ = v33
            _ = v34
        }
        fmt.Println(fi30)
        _ = fi30
    }
    if true {
        v35 := GenS1{F0: 926}
        if (-1 <= 9223372036854775807) {
            if true {
                for fi36 := 0; fi36 < 14; fi36++ {
                    for fi37 := 0; fi37 < 8; fi37++ {
                        v29 -= -1
                        _ = fi37
                    }
                    _ = fi36
                }
                v38 := GenS1{F0: 1}
                switch 2147483647 {
                case 1:
                    for fi39 := 0; fi39 < 7; fi39++ {
                        v40 := v1[fi39]
                        _ = fi39
                        _ = v40
                    }
                case 32767:
                    v41 := GenS3{F0: 2, F1: true}
                    _ = v41
                case -1:
                    switch 4294967295 {
                    case -2147483648:
                        var v42 uint8 = 0
                        _ = v42
                    case 9223372036854775807:
                        var v43 int8 = (v29 ^ (-(v29)))
                        _ = v43
                    default:
                        v44 := GenS4{F0: v29, F1: 255}
                        _ = v44
                    }
                default:
                    v45 := GenS4{F0: (v29 | v29), F1: 147}
                    _ = v45
                }
                _ = v38
            }
            v46 := GenS3{F0: 42, F1: false}
            _ = v46
        }
        var v47 int8 = 127
        _ = v35
        _ = v47
    } else {
        switch "" {
        case "\n\t":
            fmt.Println(v29, v27)
        }
        v48 := GenS4{F0: (-(v29)), F1: genFn0(4294967295, ("\n\t" + ""))}
        var v49 bool = (!(4294967295 >= 403))
        _ = v48
        _ = v49
    }
    if (("" + ("\n\t" + ("foo bar" + "hello"))) != ("\n\t" + "foo bar")) {
        v50 := v4[292]
        _ = v50
    } else {
        for fi51 := 0; fi51 < 12; fi51++ {
            v52 := GenS5{F0: (fi51 - (fi51 * -1))}
            _ = fi51
            _ = v52
        }
        v1[int(v29)] = false
    }
    fmt.Println(v27)
    v27 += (^(v27))
    delete(v1, 239)
    v53 := GenS5{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v53
}
