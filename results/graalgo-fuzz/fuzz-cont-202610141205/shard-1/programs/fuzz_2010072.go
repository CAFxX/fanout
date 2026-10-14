package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint64; F2 bool }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int; F1 uint16; F2 int64 }
func genFn0(p0 uint32, p1 uint32) int {
    var v7 int8 = -128
    v8 := GenS1{F0: 32767, F1: false, F2: (v7 & (v7 * v7))}
    v7 = (v7 >> ((uint(p1)) & 63))
    if (2 == 867) {
        v9 := []int{127, 32768}
        _ = v9
    } else {
        v10 := []int{-695}
        _ = v10
    }
    if true {
        v11 := GenS1{F0: 42, F1: false, F2: (v7 ^ v7)}
        fmt.Println(p0, p1)
        var v12 uint8 = 1
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 int, p1 uint32) uint16 {
    if (85 < 128) {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := []int{9223372036854775807}
            var v15 float64 = -2.25
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := []int{(p0 - p0), (p0 - (-(p0))), (-(p0)), -2147483648}
        _ = v16
    }
    v17 := GenS1{F0: 32767, F1: true, F2: -48}
    var v18 bool = false
    for fi19 := 0; fi19 < 2; fi19++ {
        if ((127 != 127) && (v18 && ((0.1 == 3.14159) || v18))) {
            v20 := []int{(fi19 - genFn0(p1, (p1 / ((42) | 1)))), 42}
            _ = v20
        }
        fmt.Println(v18, p1)
        v21 := []int{genFn0(p1, (^(p1))), fi19, (p0 % (((fi19 % ((genFn0((p1 & p1), (p1 >> ((2) & 63)))) | 1))) | 1))}
        _ = fi19
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 65535
}

func main() {
    var v1 uint32 = 343
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 42
    v4 := []int{127}
    var v5 int8 = 42
    var v6 float32 = 0.0
    var v22 int8 = (-(v5))
    delete(v2, 4294967295)
    fmt.Println(v22, v1, v5)
    if true {
        if (!false) {
            switch 0 {
            case -90:
                if false {
                    v2[(v1 & (v1 - (v1 - v1)))] = v1
                } else {
                    var v23 string = "foo bar"
                    fmt.Println(v6, v1)
                    var v24 uint16 = genFn1(127, 2)
                    _ = v23
                    _ = v24
                }
            default:
                v22 = v5
            }
        }
    } else {
        var v25 int16 = 1
        v4 = append(v4, 77)
        if (0 == 1) {
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            for fi26 := 0; fi26 < 13; fi26++ {
                delete(v2, uint32(v3))
                _ = fi26
            }
            v2[(^(v1))] = (v1 * (v1 & (^(v1))))
        } else {
            v25 += v25
            var v27 uint64 = v3
            _ = len(v2)
            _ = v27
        }
        _ = v25
    }
    var v28 int8 = 1
    switch genFn0(669, v1) {
    case -9223372036854775808:
        for fi29 := 0; fi29 < 15; fi29++ {
            v30 := GenS1{F0: 42, F1: (-774 == 1), F2: int8(fi29)}
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 float64 = ((0.1 / -2.25) * 0.1)
                _ = fi31
                _ = v32
            }
            _ = fi29
            _ = v30
        }
    default:
        if false {
            if (32768 != genFn0(v1, v1)) {
                v1 += (v1 ^ (v1 / ((v1) | 1)))
                var v33 int = -784
                _ = v33
            } else {
                var v34 int32 = -1
                _ = v34
            }
            fmt.Println(v28, v5, v6)
            var v35 bool = true
            _ = v35
        } else {
            var v36 int64 = -1
            v37 := GenS2{F0: v3}
            _ = v36
            _ = v37
        }
    }
    switch "\n\t" {
    case "":
        if ((false && ((-1 >= -836) && ("foo bar" <= (("foo bar" + ("\n\t" + "foo bar")) + "")))) || (false && (genFn0(v1, 65535) > 1))) {
            fmt.Println(v1, v1)
            if true {
                fmt.Println(v6, v3)
                v3 -= v3
                for fi38 := 0; fi38 < 5; fi38++ {
                    v3 += 665
                    _ = fi38
                }
            }
            v39 := GenS3{F0: genFn0((v1 - v1), v1), F1: (v3 << ((42) & 63)), F2: false}
            _ = v39
        }
    }
    if true {
        v40 := GenS4{F0: genFn0((v1 >> ((255) & 63)), (v1 << ((65535) & 63))), F1: (v3 - (v3 | (^(v3))))}
        var v41 uint32 = (v1 * (v1 | v1))
        switch "" {
        case "\n\t":
            fmt.Println(v5)
        case "":
            if true {
                var v42 uint32 = v41
                if true {
                    var v43 uint16 = 501
                    var v44 int64 = -9223372036854775808
                    fmt.Println(v44, v22, v41)
                    _ = v43
                    _ = v44
                }
                for fi45 := 0; fi45 < 20; fi45++ {
                    if true {
                        for fi46 := 0; fi46 < 1; fi46++ {
                            if (-1 >= 1) {
                                var v47 uint8 = 228
                                var v48 float64 = 0.0
                                _ = v47
                                _ = v48
                            }
                            _ = fi46
                        }
                    } else {
                        v49 := GenS5{F0: (fi45 ^ (fi45 % (((fi45 % ((genFn0((v42 - (v41 | (v41 ^ v1))), 145)) | 1))) | 1))), F1: 501, F2: -1}
                        _ = v49
                    }
                    _ = fi45
                }
                _ = v42
            }
        case "foo bar":
            for fi50 := 0; fi50 < 9; fi50++ {
                if true {
                    fmt.Println(v22)
                }
                _ = fi50
            }
        }
        _ = v40
        _ = v41
    }
    var v51 bool = false
    fmt.Println(v6)
    fmt.Println(v5, v51, v28)
    v28 = (v5 | 1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
