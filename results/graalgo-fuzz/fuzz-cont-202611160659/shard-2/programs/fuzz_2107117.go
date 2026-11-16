package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int16 }
func genFn0(p0 int64, p1 bool, p2 bool) int64 {
    v7 := []int{-361, 42, -1, -1}
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 / ((p0) | 1))
}
func genFn1(p0 uint32, p1 bool) uint64 {
    var v8 uint64 = 394
    v9 := []int{-2147483648}
    fmt.Println(p1, p1)
    switch 127 {
    case 1:
        switch ("a" + "foo bar") {
        case "a":
            p0 -= p0
        case "foo bar":
            v10 := GenS1{F0: (!p1)}
            _ = v10
        default:
            p0 = (p0 << ((948) & 63))
        }
    case 128:
        var v11 int16 = -1
        _ = v11
    default:
        v8 = v8
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (v8 + (v8 / ((320) | 1)))
}

func main() {
    v1 := []uint32{2147483648, 42}
    var v2 int64 = 0
    v3 := []int{4294967295}
    var v4 uint64 = 4294967295
    var v5 int64 = (v2 ^ v2)
    v6 := make(map[uint32]bool)
    if true {
        if (!(522 <= 237)) {
            if false {
                var v12 uint64 = v4
                v12 += 1
                _ = v12
            }
            v5 = 1
        }
        for fi13 := 0; fi13 < 6; fi13++ {
            if (42 > 1000) {
                switch fi13 {
                case 42:
                    v5 = (v2 << ((42) & 63))
                case -9223372036854775808:
                    switch "a" {
                    case "":
                        if false {
                            var v14 int = 0
                            _ = v14
                        }
                    case "a":
                        for fi15 := 0; fi15 < 1; fi15++ {
                            v16 := GenS1{F0: false}
                            _ = fi15
                            _ = v16
                        }
                    default:
                        v4 += (v4 / ((v4) | 1))
                    }
                default:
                    fmt.Println(v5)
                }
                var v17 int32 = -1
                _ = v17
            }
            _ = fi13
        }
    } else {
        var v18 uint = 784
        v2 = (^(v5))
        if false {
            v2 = 42
        }
        _ = v18
    }
    v2 = (v5 ^ (v5 << ((255) & 63)))
    v19 := v6[410]
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 uint32 = 2147483648
        for fi22 := 0; fi22 < 11; fi22++ {
            fmt.Println(v5, v5, v5)
            fmt.Println(v5, v5)
            v23 := GenS1{F0: v19}
            _ = fi22
            _ = v23
        }
        if (!(false && ((v4 % ((v4) | 1)) <= 2))) {
            var v24 int = (fi20 | -1)
            v19 = v19
            _ = v24
        } else {
            _ = len(v1)
            var v25 int = fi20
            _ = v25
        }
        _ = fi20
        _ = v21
    }
    if (1 != -1) {
        var v26 uint8 = 186
        fmt.Println(v2, v26)
        for fi27 := 0; fi27 < 14; fi27++ {
            for fi28 := 0; fi28 < 9; fi28++ {
                fmt.Println(v2)
                _ = fi28
            }
            _ = fi27
        }
        _ = v26
    }
    var v29 int64 = genFn0(v2, v19, v19)
    var v30 float64 = float64(((-2.25 / 1.5) / (0.1 / 1.5)))
    v2 = v5
    for fi31 := 0; fi31 < 10; fi31++ {
        v32 := GenS2{F0: 1, F1: (v4 | v4)}
        if (65535 < 4294967295) {
            for fi33 := 0; fi33 < 14; fi33++ {
                var v34 int32 = 0
                var v35 int32 = (v34 * 42)
                v29 += v5
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = len(v6)
        } else {
            v2 -= genFn0(v5, (!v19), v19)
        }
        _ = fi31
        _ = v32
    }
    if (v30 <= (v30 - ((v30 - v30) * v30))) {
        v36 := GenS2{F0: 0, F1: (v4 ^ v4)}
        fmt.Println(v30, v5, v4)
        var v37 float32 = 1.5
        _ = v36
        _ = v37
    }
    var v38 uint32 = 1
    v39 := GenS3{F0: v38, F1: 255, F2: 42}
    switch ("foo bar" + "") {
    case "a":
        v19 = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
