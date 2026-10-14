package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 bool; F1 uint; F2 int }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 bool; F1 uint64; F2 int32 }
func genFn0(p0 int, p1 uint64, p2 int32) bool {
    v7 := []int{p0, (p0 / ((p0) | 1))}
    p2 -= int32(p1)
    for fi8 := 0; fi8 < 19; fi8++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    if len(v7) > 0 {
        v7[0] = (p0 * (p0 << ((18446744073709551615) & 63)))
    }
    switch "foo bar" {
    case "a":
        p0 -= (p0 ^ (^(p0)))
    case "hello":
        if (p1 != (^(p1))) {
            v9 := GenS1{F0: 318}
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 int64 = 42
                for fi12 := 0; fi12 < 6; fi12++ {
                    fmt.Println(fi12)
                    fmt.Println(fi10, fi12)
                    p2 = p2
                    _ = fi12
                }
                _ = fi10
                _ = v11
            }
            _ = v9
        } else {
            fmt.Println(p0)
            if ("\n\t" < "") {
                fmt.Println(p0, p0, p2)
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                p0 = (p0 | p0)
            }
            p1 -= (^(p1))
        }
    case "foo bar":
        var v13 int64 = 9223372036854775807
        _ = v13
    default:
        v14 := GenS1{F0: 65535}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}
func genFn1(p0 int64, p1 uint32) uint64 {
    for fi15 := 0; fi15 < 19; fi15++ {
        p0 = p0
        v16 := GenS1{F0: (p1 / (((^(p1))) | 1))}
        var v17 int32 = 2147483647
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float64 = 0.1
    for fi19 := 0; fi19 < 13; fi19++ {
        p0 = 0
        if (763 >= 1) {
            var v20 uint16 = 413
            _ = v20
        }
        p1 += p1
        _ = fi19
    }
    p1 -= 526
    _ = p0
    _ = p1
    _ = v18
    return 950
}

func main() {
    var v1 bool = (0 > 0)
    v2 := make(map[int]int)
    var v3 uint8 = 2
    var v4 int8 = -128
    var v5 uint16 = 230
    v6 := make(map[int]bool)
    for fi21 := 0; fi21 < 17; fi21++ {
        v4 = v4
        var v22 int16 = -1
        fmt.Println(v22, v1)
        _ = fi21
        _ = v22
    }
    var v23 uint32 = 883
    v5 += (v5 * v5)
    var v24 int8 = int8(v3)
    v23 = v23
    if false {
        for fi25 := 0; fi25 < 4; fi25++ {
            if (v1 && ((^(v23)) < uint32(v5))) {
                v24 = (v24 - 0)
                if v1 {
                    var v26 float64 = (1e10 / -2.25)
                    _ = v26
                }
                v27 := GenS2{F0: 18446744073709551615}
                _ = v27
            }
            var v28 int16 = int16(v23)
            v28 += (^(v28))
            _ = fi25
            _ = v28
        }
        var v29 uint = 18446744073709551615
        _ = v29
    } else {
        var v30 uint = 42
        _ = v30
    }
    v31 := []int{32768, 1, -846, 0}
    var v32 float32 = ((0.0 + (1.5 * 3.14159)) + (0.1 / 3.14159))
    fmt.Println(v1, v5)
    for fi33 := 0; fi33 < 2; fi33++ {
        fmt.Println(v3, v4)
        if v1 {
            v5 -= v5
            switch "\n\t" {
            case "hello":
                for fi34 := 0; fi34 < 17; fi34++ {
                    fmt.Println(fi33)
                    _ = fi34
                }
            case "\n\t":
                for fi35 := 0; fi35 < 17; fi35++ {
                    v36 := GenS1{F0: v23}
                    v5 += (v5 >> ((773) & 63))
                    _ = fi35
                    _ = v36
                }
            default:
                v23 += v23
            }
            v1 = true
        } else {
            switch (fi33 << ((2) & 63)) {
            case -250:
                switch "\n\t" {
                case "foo bar":
                    switch "a" {
                    case "hello":
                        for fi37 := 0; fi37 < 10; fi37++ {
                            if (v1 && v1) {
                                v24 -= 78
                            } else {
                                if true {
                                    var v38 float32 = -2.25
                                    var v39 int64 = 1
                                    _ = v38
                                    _ = v39
                                } else {
                                    if (v3 == 0) {
                                        v40 := GenS3{F0: uint8(v5), F1: ""}
                                        var v41 uint64 = 4294967295
                                        _ = v40
                                        _ = v41
                                    } else {
                                        var v42 int = fi37
                                        _ = v42
                                    }
                                    if (!(v3 > v3)) {
                                        if v1 {
                                            var v43 uint16 = (v5 << ((2) & 63))
                                            v5 = 0
                                            var v44 bool = (42 > 835)
                                            _ = v43
                                            _ = v44
                                        } else {
                                            v2[42] = int(v23)
                                            v45 := GenS4{F0: false, F1: 114, F2: (fi37 ^ fi37)}
                                            _ = v45
                                        }
                                        if false {
                                            v6[-717] = (float32(0.0) >= v32)
                                            delete(v2, (fi33 % (((fi37 - fi37)) | 1)))
                                            v46 := GenS1{F0: 4294967295}
                                            var v47 int16 = -807
                                            v31 = append(v31, (fi33 << ((2) & 63)))
                                            _ = v46
                                            _ = v47
                                        }
                                    }
                                    v48 := GenS5{F0: (fi37 ^ fi37), F1: "a"}
                                    if v1 {
                                        v23 = (v23 | uint32(v3))
                                        var v49 int64 = 42
                                        if v1 {
                                            if false {
                                                var v50 uint32 = v23
                                                _ = v50
                                            }
                                            var v51 bool = false
                                            v5 += (^(v5))
                                            _ = v51
                                        }
                                        _ = v49
                                    }
                                    _ = v48
                                }
                            }
                            _ = fi37
                        }
                    case "foo bar":
                        if v1 {
                            v52 := GenS6{F0: v3, F1: 32767}
                            _ = v52
                        } else {
                            switch int(v3) {
                            case 924:
                                var v53 int16 = 1
                                _ = v53
                            default:
                                delete(v2, fi33)
                            }
                            v24 += v24
                            v32 += (v32 - float32(1.5))
                        }
                    default:
                        var v54 string = "a"
                        _ = v54
                    }
                case "":
                    v55 := v6[fi33]
                    _ = v55
                default:
                    v4 += (v4 << ((152) & 63))
                }
            case 32768:
                if len(v31) >= 2 {
                    v31 = v31[1:len(v31)]
                }
            default:
                v32 += v32
            }
            var v56 int64 = 9223372036854775807
            if (1 > 32767) {
                v24 += (-(v4))
            } else {
                v57 := GenS7{F0: false, F1: 738, F2: -2147483648}
                _ = v57
            }
            _ = v56
        }
        _ = fi33
    }
    var v58 int16 = -735
    if len(v31) >= 2 {
        v31 = v31[1:len(v31)]
    }
    var v59 uint8 = v3
    _ = len(v31)
    var v60 int16 = -32768
    if len(v31) >= 2 {
        v31 = v31[1:len(v31)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
