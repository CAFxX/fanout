package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint16; F1 string }
func genFn0(p0 bool, p1 uint16, p2 int64) uint {
    var v7 uint16 = 32767
    v7 = (v7 << ((54) & 63))
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 float32 = 0.1
        _ = fi8
        _ = v9
    }
    v7 += 2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 97
}
func genFn1(p0 bool, p1 int64, p2 uint) int16 {
    switch 1 {
    case -2147483648:
        fmt.Println(p2, p0, p1)
    }
    switch ("" + "a") {
    case "foo bar":
        p0 = true
    }
    v10 := GenS1{F0: 32767}
    p2 = (p2 % ((genFn0((p0 && p0), 65535, p1)) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}

func main() {
    var v1 uint = 725
    v2 := make(map[uint32]bool)
    v3 := []int{-625, -932}
    var v4 int32 = -1
    var v5 int16 = 734
    var v6 uint32 = 2
    v3 = append(v3, 1)
    if false {
        switch 373 {
        case 32768:
            v11 := GenS1{F0: (v5 << ((v1) & 63))}
            _ = v11
        case -133:
            v12 := GenS1{F0: v5}
            _ = v12
        default:
            if (!true) {
                _ = len(v3)
            } else {
                fmt.Println(v6)
                v6 += v6
            }
        }
        v1 = 255
        for fi13 := 0; fi13 < 3; fi13++ {
            if false {
                var v14 int32 = v4
                if (!false) {
                    switch "hello" {
                    case "hello":
                        var v15 uint8 = 51
                        _ = v15
                    case "a":
                        _ = len(v2)
                    }
                } else {
                    _ = len(v2)
                }
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 int64 = 42
                    fmt.Println(v4)
                    if (true && (4294967295 == v1)) {
                        if true {
                            v18 := GenS2{F0: 42, F1: 127}
                            _ = v18
                        }
                        if ((243 != 42) && true) {
                            v5 += (v5 << ((v1) & 63))
                            if (v6 != (v6 << ((v1) & 63))) {
                                v1 += v1
                            }
                        } else {
                            v19 := GenS1{F0: int16(v17)}
                            if (true || false) {
                                v20 := GenS2{F0: 127, F1: 0}
                                if (int8(fi16) < 49) {
                                    _ = len(v3)
                                    _ = len(v3)
                                } else {
                                    v3 = append(v3, -9223372036854775808)
                                }
                                if (!false) {
                                    var v21 float64 = 0.1
                                    var v22 bool = ((v5 + 0) != -593)
                                    _ = v21
                                    _ = v22
                                } else {
                                    var v23 int8 = -1
                                    if true {
                                        fmt.Println(v6, fi16)
                                    } else {
                                        var v24 int8 = (^(v23))
                                        _ = v24
                                    }
                                    _ = v23
                                }
                                _ = v20
                            }
                            _ = v19
                        }
                    }
                    _ = fi16
                    _ = v17
                }
                _ = v14
            }
            _ = fi13
        }
    }
    _ = len(v3)
    v6 = v6
    if false {
        var v25 uint = (v1 * v1)
        v5 += v5
        for fi26 := 0; fi26 < 19; fi26++ {
            if false {
                if ((255 <= 42) || false) {
                    var v27 float64 = ((3.14159 - 3.14159) - (1.5 + 0.0))
                    v25 -= 0
                    fmt.Println(v27)
                    _ = v27
                }
            } else {
                fmt.Println(fi26, v25, v6)
                v4 -= (v4 * (v4 % ((v4) | 1)))
                if len(v3) > 0 {
                    v3[0] = 241
                }
            }
            v4 -= (v4 >> ((v1) & 63))
            v4 -= v4
            _ = fi26
        }
        _ = v25
    }
    if false {
        for fi28 := 0; fi28 < 19; fi28++ {
            if false {
                v29 := GenS3{F0: (v6 >> (((v1 & (v1 / ((v1) | 1)))) & 63)), F1: 0}
                var v30 float64 = (float64(0.1) + -2.25)
                var v31 int8 = 42
                _ = v29
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 2; fi32++ {
                v4 += (v4 & v4)
                var v33 int64 = 42
                _ = fi32
                _ = v33
            }
            _ = fi28
        }
        var v34 uint8 = 62
        var v35 float32 = 1e10
        _ = v34
        _ = v35
    } else {
        var v36 uint8 = 102
        var v37 float32 = 0.1
        _ = v36
        _ = v37
    }
    fmt.Println(v5)
    for fi38 := 0; fi38 < 2; fi38++ {
        v39 := GenS1{F0: (v5 ^ (v5 & (v5 >> ((genFn0((true && ((32767 >= 32767) && true)), 778, -661)) & 63))))}
        v4 -= (v4 - (v4 - v4))
        _ = fi38
        _ = v39
    }
    v1 = (^(v1))
    v40 := GenS3{F0: (v6 >> ((v1) & 63)), F1: 9223372036854775807}
    var v41 uint8 = 1
    v42 := GenS4{F0: 743, F1: ("a" + "\n\t")}
    for fi43 := 0; fi43 < 4; fi43++ {
        fmt.Println(v41)
        var v44 int = int(v5)
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 12; fi45++ {
        v46 := v2[(^(v6))]
        var v47 int16 = (v5 * 1000)
        _ = fi45
        _ = v46
        _ = v47
    }
    if true {
        v6 -= (v6 / ((v6) | 1))
        v48 := GenS2{F0: v41, F1: v41}
        _ = v48
    } else {
        v1 += (v1 ^ (v1 - v1))
        v4 -= 0
        var v49 int32 = v4
        _ = v49
    }
    var v50 bool = (v4 < 42)
    if true {
        v51 := GenS2{F0: 42, F1: (v41 >> ((v1) & 63))}
        for fi52 := 0; fi52 < 15; fi52++ {
            for fi53 := 0; fi53 < 1; fi53++ {
                v2[v6] = (42 != 120)
                _ = fi53
            }
            _ = fi52
        }
        fmt.Println(v6, v5, v5)
        _ = v51
    } else {
        if (v5 != v5) {
            switch "foo bar" {
            case "\n\t":
                var v54 uint64 = 2
                _ = v54
            case "foo bar":
                var v55 int32 = v4
                _ = v55
            case "":
                v50 = (false || v50)
            }
        } else {
            v5 -= v5
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v50
    _ = v6
}
