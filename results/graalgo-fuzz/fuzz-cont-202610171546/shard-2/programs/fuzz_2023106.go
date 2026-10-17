package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint32; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8 }
func genFn0(p0 bool, p1 uint8, p2 int) string {
    v6 := []int{(p2 ^ 283), 9223372036854775807, 2147483647}
    if len(v6) > 0 {
        v6[0] = (p2 % ((p2) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "foo bar"
}
func genFn1(p0 uint64, p1 int32) int32 {
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{(fi7 >> ((605) & 63)), (fi7 << ((59) & 63))}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint64 = p0
        v10 -= (^(v10))
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]uint32)
    v3 := []int{0, -1, -380, 9223372036854775807}
    v4 := make(map[string]string)
    v5 := make(map[uint32]uint32)
    v1 -= v1
    var v11 float32 = 0.1
    v12 := GenS1{F0: -671}
    v13 := GenS2{F0: false}
    if false {
        var v14 uint64 = v1
        _ = v14
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(fi15, v11, v11)
            var v16 uint16 = 1
            _ = fi15
            _ = v16
        }
        if ((v1 ^ v1) < (v1 ^ 498)) {
            if false {
                v1 += (v1 ^ (^(v1)))
            } else {
                v17 := GenS3{F0: 127, F1: 0, F2: -128}
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 uint = 42
                    var v20 float32 = v11
                    fmt.Println(v20)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v21 := GenS3{F0: 127, F1: 2, F2: -128}
                _ = v17
                _ = v21
            }
        }
    }
    switch "\n\t" {
    case "a":
        for fi22 := 0; fi22 < 9; fi22++ {
            fmt.Println(v1)
            _ = fi22
        }
    case "hello":
        var v23 uint64 = (v1 * v1)
        _ = v23
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = 421
        }
    default:
        v24 := GenS3{F0: -796, F1: 65535, F2: 42}
        _ = v24
    }
    switch "foo bar" {
    case "hello":
        v25 := GenS4{F0: 42}
        _ = v25
    case "a":
        if false {
            switch "\n\t" {
            case "a":
                var v26 int64 = 473
                _ = v26
            }
            v27 := GenS5{F0: 2}
            v28 := GenS2{F0: (42 <= 0)}
            _ = v27
            _ = v28
        } else {
            v11 = 0.1
            v11 = (1e10 + v11)
            switch "\n\t" {
            case "":
                var v29 int32 = genFn1(v1, 1)
                _ = v29
            }
        }
    }
    v2["foo bar"] = 42
    fmt.Println(v1, v1)
    var v30 uint8 = 255
    v31 := GenS1{F0: -425}
    var v32 int16 = -1
    for fi33 := 0; fi33 < 1; fi33++ {
        var v34 int = (fi33 >> ((0) & 63))
        if ((false && true) && true) {
            var v35 int = fi33
            if len(v3) > 0 {
                v3[0] = fi33
            }
            _ = v35
        }
        if false {
            if ((("hello" + "") + ("hello" + "hello")) >= "foo bar") {
                fmt.Println(v30, v30, v1)
                v36 := GenS3{F0: fi33, F1: 0, F2: 98}
                if true {
                    switch "\n\t" {
                    case "hello":
                        if true {
                            switch fi33 {
                            case -666:
                                if len(v3) > 0 {
                                    v3[0] = (v34 - (fi33 - (^(v34))))
                                }
                            case 283:
                                var v37 int = fi33
                                _ = v37
                            case -734:
                                v30 += (^(v30))
                            default:
                                var v38 bool = (false || ((-(v34)) <= fi33))
                                _ = v38
                            }
                        } else {
                            var v39 uint8 = (v30 | v30)
                            v40 := GenS3{F0: v34, F1: 1, F2: -1}
                            _ = v39
                            _ = v40
                        }
                    case "":
                        v41 := GenS6{F0: -796}
                        _ = v41
                    case "a":
                        v11 -= (v11 * 1e10)
                    default:
                        var v42 int8 = 42
                        _ = v42
                    }
                }
                _ = v36
            }
        } else {
            v43 := GenS7{F0: 42}
            if false {
                var v44 int16 = 42
                v45 := GenS3{F0: fi33, F1: 42, F2: 88}
                _ = v44
                _ = v45
            } else {
                switch "\n\t" {
                case "hello":
                    if (0 > 65535) {
                        if true {
                            v46 := GenS6{F0: 9223372036854775807}
                            v34 = (fi33 & int(v1))
                            for fi47 := 0; fi47 < 1; fi47++ {
                                var v48 int32 = -306
                                var v49 float32 = ((v11 * (v11 / (((v11 * v11) + v11) + 0.1))) - 0.0)
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                            _ = v46
                        } else {
                            for fi50 := 0; fi50 < 8; fi50++ {
                                var v51 int64 = -485
                                v51 = (^(v51))
                                v11 = ((float32(-2.25) * v11) + float32(3.14159))
                                _ = fi50
                                _ = v51
                            }
                        }
                        v52 := GenS5{F0: 632}
                        _ = v52
                    }
                case "a":
                    if (!true) {
                        var v53 float64 = 0.0
                        _ = v53
                    }
                case "\n\t":
                    v11 -= 1e10
                default:
                    delete(v5, 437)
                }
                fmt.Println(v32)
                var v54 uint32 = 65535
                _ = v54
            }
            fmt.Println(v1, v32, v34)
            _ = v43
        }
        _ = fi33
        _ = v34
    }
    v11 -= (3.14159 / float32(3.14159))
    var v55 int8 = -19
    var v56 int8 = (v55 / (((^(v55))) | 1))
    v57 := GenS3{F0: 42, F1: 696, F2: 29}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
}
