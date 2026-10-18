package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int32 }
type GenS2 struct { F0 int; F1 uint32; F2 int32 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int16; F1 string; F2 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 int64 }
func genFn0(p0 int32, p1 int8, p2 int16) uint64 {
    var v7 uint8 = 0
    v8 := GenS1{F0: (0.0 != (0.0 - (-2.25 * ((0.0 * float64(1.5)) - 1e10)))), F1: p1, F2: p0}
    v9 := GenS2{F0: 32767, F1: 2147483647, F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 2147483648
    var v2 string = "\n\t"
    var v3 int16 = -1
    var v4 int32 = -608
    v5 := make(map[string]string)
    v6 := make(map[string]int)
    if false {
        v10 := []int{}
        _ = v10
    }
    if (!false) {
        var v11 uint16 = 65535
        var v12 string = "a"
        _ = v11
        _ = v12
    } else {
        switch -966 {
        case 128:
            fmt.Println(v1, v2)
        case 2147483647:
            for fi13 := 0; fi13 < 15; fi13++ {
                for fi14 := 0; fi14 < 13; fi14++ {
                    if (true || true) {
                        var v15 bool = false
                        v16 := GenS1{F0: true, F1: -82, F2: (v4 << ((0) & 63))}
                        if (-941 <= v4) {
                            v15 = v15
                            var v17 uint32 = 65535
                            _ = v17
                        }
                        _ = v15
                        _ = v16
                    }
                    var v18 float32 = -2.25
                    var v19 string = (v2 + v2)
                    _ = fi14
                    _ = v18
                    _ = v19
                }
                _ = fi13
            }
        case 42:
            for fi20 := 0; fi20 < 3; fi20++ {
                switch fi20 {
                case 89:
                    switch (fi20 & (fi20 >> ((271) & 63))) {
                    case 333:
                        v21 := GenS2{F0: (fi20 * (fi20 % ((fi20) | 1))), F1: (v1 / ((167) | 1)), F2: 0}
                        _ = v21
                    case 225:
                        v1 = v1
                    case 341:
                        var v22 int32 = (v4 % (((v4 >> ((787) & 63))) | 1))
                        _ = v22
                    }
                case -227:
                    v4 = (v4 + v4)
                case 128:
                    v1 = v1
                default:
                    if false {
                        if true {
                            for fi23 := 0; fi23 < 14; fi23++ {
                                fmt.Println(fi20, fi20, v2)
                                var v24 uint32 = v1
                                _ = fi23
                                _ = v24
                            }
                        } else {
                            fmt.Println(fi20, v1, v4)
                        }
                        if true {
                            var v25 int64 = -236
                            for fi26 := 0; fi26 < 16; fi26++ {
                                v3 = v3
                                _ = fi26
                            }
                            _ = v25
                        }
                    } else {
                        switch (fi20 | fi20) {
                        case 944:
                            var v27 float32 = 1.5
                            _ = v27
                        case 32768:
                            var v28 uint64 = 0
                            _ = v28
                        case 127:
                            var v29 int16 = v3
                            _ = v29
                        default:
                            v30 := []int{42, fi20}
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 16; fi31++ {
                            v32 := GenS3{F0: 1, F1: v1}
                            if true {
                                v2 = "hello"
                            } else {
                                switch ("foo bar" + "hello") {
                                case "a":
                                    if true {
                                        if false {
                                            v1 = v1
                                            var v33 int64 = 0
                                            _ = v33
                                        } else {
                                            v3 = 32767
                                            var v34 int32 = v4
                                            _ = v34
                                        }
                                        var v35 uint8 = 127
                                        _ = v35
                                    }
                                default:
                                    var v36 string = "a"
                                    _ = v36
                                }
                                if false {
                                    switch "hello" {
                                    case "hello":
                                        v37 := []int{int(v1), fi20}
                                        _ = v37
                                    default:
                                        if true {
                                            _ = len(v5)
                                            fmt.Println(fi31, v2, v4)
                                        } else {
                                            if (((fi31 * fi31) != (fi20 % (((-(fi20))) | 1))) || false) {
                                                v4 = (^(v4))
                                                var v38 uint32 = v1
                                                fmt.Println(v2, fi20)
                                                _ = v38
                                            }
                                        }
                                    }
                                } else {
                                    v39 := GenS4{F0: (v3 + int16(v1)), F1: "\n\t", F2: 42}
                                    v3 -= (v3 & v3)
                                    _ = v39
                                }
                                v40 := GenS5{F0: 65535}
                                _ = v40
                            }
                            v41 := v6[v2]
                            _ = fi31
                            _ = v32
                            _ = v41
                        }
                    }
                }
                v42 := GenS6{F0: uint64(v1), F1: 31}
                _ = fi20
                _ = v42
            }
        }
    }
    fmt.Println(v1, v1)
    for fi43 := 0; fi43 < 18; fi43++ {
        fmt.Println(v4)
        var v44 bool = (!false)
        var v45 uint16 = 42
        _ = fi43
        _ = v44
        _ = v45
    }
    switch v2 {
    case "\n\t":
        v46 := GenS3{F0: -2147483648, F1: 4294967295}
        _ = v46
    case "a":
        for fi47 := 0; fi47 < 7; fi47++ {
            v4 += -922
            var v48 uint32 = v1
            _ = fi47
            _ = v48
        }
    case "foo bar":
        v49 := []int{0, 743, 4294967295}
        _ = v49
    }
    var v50 uint64 = 42
    var v51 float64 = 1e10
    v1 = uint32(v50)
    var v52 int32 = v4
    _ = len(v5)
    for fi53 := 0; fi53 < 6; fi53++ {
        v54 := GenS7{F0: 42}
        _ = fi53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
