package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint; F2 int64 }
type GenS3 struct { F0 uint; F1 uint8; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 uint16; F2 uint }
type GenS6 struct { F0 uint16; F1 uint8; F2 int32 }
func genFn0(p0 bool, p1 uint, p2 int) uint {
    switch p2 {
    case 0:
        v4 := []int{(p2 << ((2) & 63)), (p2 & p2)}
        _ = v4
    default:
        p2 = -777
    }
    var v5 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 705
}
func genFn1(p0 int32) uint64 {
    switch "a" {
    case "":
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 int16 = -32768
            _ = fi6
            _ = v7
        }
    case "hello":
        var v8 uint8 = 1
        _ = v8
    case "a":
        var v9 uint8 = 128
        _ = v9
    default:
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := []int{fi10}
            if (false && false) {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                switch -1 {
                case -1:
                    if len(v11) > 0 {
                        v11[0] = (fi10 & (-(fi10)))
                    }
                case 42:
                    switch 9223372036854775807 {
                    case -174:
                        var v12 int8 = 1
                        _ = v12
                    case 1:
                        var v13 int16 = -32768
                        _ = v13
                    case -188:
                        if len(v11) > 0 {
                            v11[0] = fi10
                        }
                    default:
                        if false {
                            var v14 uint64 = 18446744073709551615
                            var v15 uint16 = 692
                            if (((!false) && true) || true) {
                                v16 := GenS1{F0: (p0 / (((p0 / ((p0) | 1))) | 1))}
                                fmt.Println(fi10)
                                for fi17 := 0; fi17 < 5; fi17++ {
                                    if len(v11) > 0 {
                                        v11[0] = (-(fi10))
                                    }
                                    _ = fi17
                                }
                                _ = v16
                            }
                            _ = v14
                            _ = v15
                        } else {
                            var v18 uint64 = 4294967295
                            _ = v18
                        }
                    }
                }
                if (true || (3.14159 < -2.25)) {
                    fmt.Println(fi10, p0)
                    switch ("" + "") {
                    case "\n\t":
                        if false {
                            fmt.Println(fi10, fi10)
                            fmt.Println(fi10, p0)
                            p0 = 42
                        }
                    case "hello":
                        p0 = (^(p0))
                    default:
                        for fi19 := 0; fi19 < 10; fi19++ {
                            fmt.Println(p0, fi19, fi19)
                            p0 = (-(p0))
                            if true {
                                var v20 uint64 = 1
                                _ = v20
                            } else {
                                p0 = (p0 ^ p0)
                            }
                            _ = fi19
                        }
                    }
                    v21 := GenS2{F0: 4294967295, F1: genFn0(false, genFn0((42 > 42), 1, fi10), fi10), F2: -67}
                    _ = v21
                } else {
                    var v22 bool = false
                    _ = v22
                }
            } else {
                p0 += (^(p0))
            }
            _ = fi10
            _ = v11
        }
    }
    var v23 float64 = 1e10
    var v24 int = 2147483647
    _ = p0
    _ = v23
    _ = v24
    return 2
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    if (!("a" <= "\n\t")) {
        if false {
            var v25 uint8 = v1
            if true {
                _ = len(v2)
                v26 := GenS3{F0: 2, F1: v1, F2: (int64(v25) > 42)}
                _ = v26
            }
            _ = v25
        }
        v27 := GenS2{F0: 18446744073709551615, F1: 42, F2: 1}
        _ = v27
    }
    switch "\n\t" {
    case "a":
        var v28 float64 = float64(float32(float64(0.0)))
        _ = v28
    case "hello":
        var v29 uint = 0
        _ = v29
    }
    var v30 uint32 = 1
    v30 = v30
    v30 += 1
    v1 = (v1 & 2)
    fmt.Println(v30)
    v31 := v2[("a" + "\n\t")]
    switch v31 {
    case "\n\t":
        fmt.Println(v1, v1, v31)
    case "hello":
        var v32 int8 = -75
        _ = v32
    }
    v33 := []int{128, -895, 0}
    switch v31 {
    case "hello":
        v34 := GenS4{F0: 118}
        _ = v34
    case "\n\t":
        if false {
            v35 := GenS4{F0: 255}
            _ = v35
        } else {
            switch v31 {
            case "hello":
                if (!true) {
                    var v36 uint64 = 102
                    _ = v36
                }
            case "foo bar":
                v1 -= (v1 - v1)
            case "a":
                var v37 int8 = 127
                _ = v37
            default:
                var v38 float32 = 1e10
                _ = v38
            }
            v39 := GenS2{F0: 457, F1: uint(v1), F2: 9223372036854775807}
            _ = v39
        }
    case "foo bar":
        var v40 int64 = 42
        _ = v40
    default:
        if ((v1 >= v1) || false) {
            v1 += (v1 / (((v1 >> ((2) & 63))) | 1))
            var v41 uint64 = 2
            var v42 float32 = (3.14159 * 3.14159)
            _ = v41
            _ = v42
        } else {
            switch 127 {
            case -347:
                switch 1 {
                case -43:
                    v43 := GenS4{F0: 295}
                    _ = v43
                case -9223372036854775808:
                    var v44 uint16 = 171
                    _ = v44
                default:
                    var v45 bool = (false || (2 <= 189))
                    _ = v45
                }
            case -268:
                if true {
                    fmt.Println(v1, v1)
                } else {
                    for fi46 := 0; fi46 < 20; fi46++ {
                        v47 := GenS4{F0: 4294967295}
                        _ = fi46
                        _ = v47
                    }
                    v48 := GenS5{F0: 42, F1: 878, F2: 255}
                    fmt.Println(v31, v1)
                    _ = v48
                }
            default:
                if (32767 >= 162) {
                    if (!true) {
                        v30 = v30
                    } else {
                        v30 = v30
                        var v49 int16 = -793
                        if len(v33) >= 2 {
                            v33 = v33[1:len(v33)]
                        }
                        _ = v49
                    }
                    v33 = append(v33, 4294967295)
                }
            }
            v30 -= (v30 << ((926) & 63))
        }
    }
    v1 = (v1 / ((255) | 1))
    fmt.Println(v1, v30, v31)
    switch (v31 + (v31 + (v31 + v31))) {
    case "a":
        v50 := GenS3{F0: 18446744073709551615, F1: v1, F2: (uint(v30) < genFn0((0 == 42), genFn0((!((v30 % ((v30) | 1)) == v30)), 18446744073709551615, -9223372036854775808), 9223372036854775807))}
        _ = v50
    case "foo bar":
        v30 = (v30 + v30)
    case "\n\t":
        v33 = append(v33, 942)
    default:
        v30 = (v30 / ((v30) | 1))
    }
    var v51 int16 = 1
    v1 += (v1 ^ (v1 * (^(v1))))
    v52 := GenS6{F0: 1, F1: v1, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v51
    _ = v52
}
