package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 string }
func genFn0(p0 int64) int32 {
    v4 := []int{42, 128}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, 4294967295)
    _ = p0
    _ = v4
    return -1
}
func genFn1(p0 uint64, p1 uint32) uint64 {
    p1 -= (p1 - (^(p1)))
    v5 := GenS1{F0: 42, F1: 4294967295}
    for fi6 := 0; fi6 < 12; fi6++ {
        p0 -= (p0 >> ((1) & 63))
        fmt.Println(fi6)
        _ = fi6
    }
    var v7 int16 = 42
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 378
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 0
    v3 := make(map[string]bool)
    v2 = (v2 | v2)
    fmt.Println(v2)
    v8 := v3["\n\t"]
    if false {
        if false {
            fmt.Println(v8, v2)
            for fi9 := 0; fi9 < 14; fi9++ {
                for fi10 := 0; fi10 < 3; fi10++ {
                    v11 := v1[2147483648]
                    var v12 uint64 = 601
                    v13 := GenS1{F0: 255, F1: 42}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                fmt.Println(fi9, v2, fi9)
                _ = fi9
            }
            for fi14 := 0; fi14 < 4; fi14++ {
                v15 := GenS1{F0: 32, F1: 18446744073709551615}
                v16 := GenS2{F0: ((3.14159 * (0.0 / 0.1)) < (((0.0 / 1.0) - 1e10) - (-2.25 * (0.1 / (-2.25 - float32(3.14159)))))), F1: 127}
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    } else {
        if false {
            fmt.Println(v8, v2)
        } else {
            var v17 uint64 = 4294967295
            v17 = genFn1(42, 4294967295)
            if v8 {
                v18 := v3["a"]
                _ = v18
            } else {
                _ = len(v1)
                fmt.Println(v17, v2, v8)
            }
            _ = v17
        }
        if (!(v8 || (1 > 42))) {
            v2 += (v2 / ((0) | 1))
        } else {
            if (778 != 0) {
                v8 = ((!((!v8) || ((!(false && (v2 > 777))) && (42 <= 809)))) || (true && (227 >= v2)))
                _ = len(v3)
                if (v2 == v2) {
                    fmt.Println(v2, v8)
                    fmt.Println(v2, v8)
                }
            } else {
                for fi19 := 0; fi19 < 1; fi19++ {
                    switch "hello" {
                    case "hello":
                        fmt.Println(v2, fi19, v2)
                    case "foo bar":
                        for fi20 := 0; fi20 < 19; fi20++ {
                            if v8 {
                                if false {
                                    delete(v3, "a")
                                    v21 := []int{(fi19 & (fi19 & (fi20 / ((fi20) | 1)))), -1}
                                    _ = v21
                                } else {
                                    var v22 uint16 = 1
                                    _ = v22
                                }
                                v23 := GenS2{F0: (42 != 513), F1: 88}
                                _ = v23
                            } else {
                                v8 = false
                            }
                            _ = fi20
                        }
                    case "":
                        v8 = (false || (1 > 255))
                    default:
                        v24 := GenS2{F0: v8, F1: 127}
                        _ = v24
                    }
                    v25 := v1[uint32(v2)]
                    _ = fi19
                    _ = v25
                }
                v26 := []int{-432, 117, -925}
                var v27 int = 593
                _ = v26
                _ = v27
            }
        }
        v28 := GenS3{F0: ("" + "\n\t")}
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 int8 = 127
        _ = fi29
        _ = v30
    }
    v31 := []int{1}
    v32 := GenS2{F0: v8, F1: -11}
    var v33 int8 = 42
    if (v8 && v8) {
        var v34 int = -273
        for fi35 := 0; fi35 < 18; fi35++ {
            if (false && true) {
                v8 = true
                v2 = 4294967295
                var v36 uint32 = 1
                _ = v36
            } else {
                _ = len(v1)
            }
            _ = fi35
        }
        _ = v34
    }
    v8 = (42 > uint32(v2))
    v8 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v8
}
