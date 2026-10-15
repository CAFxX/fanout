package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 uint16; F2 string }
func genFn0(p0 uint, p1 int, p2 uint32) int8 {
    v5 := []int{0, 1}
    for fi6 := 0; fi6 < 2; fi6++ {
        if true {
            switch p1 {
            case 333:
                var v7 int = (fi6 << (((^(p0))) & 63))
                _ = v7
            case -2147483648:
                fmt.Println(p1, p0)
            case 128:
                var v8 int32 = -60
                _ = v8
            default:
                if true {
                    if (!(-79 < -124)) {
                        if len(v5) > 0 {
                            v5[0] = 127
                        }
                        p2 = p2
                        switch "hello" {
                        case "\n\t":
                            p1 += 42
                        case "foo bar":
                            switch ("hello" + "a") {
                            case "foo bar":
                                var v9 int64 = 1
                                _ = v9
                            case "hello":
                                _ = len(v5)
                            }
                        }
                    }
                    if true {
                        if false {
                            var v10 int64 = -1
                            p2 -= (p2 % ((p2) | 1))
                            fmt.Println(p1, p1)
                            _ = v10
                        } else {
                            var v11 float32 = ((0.1 - 0.1) + 3.14159)
                            v12 := GenS1{F0: 942, F1: 42}
                            v13 := GenS1{F0: 255, F1: 0}
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 18; fi14++ {
                            fmt.Println(p2, p0, p2)
                            _ = fi14
                        }
                    } else {
                        switch "hello" {
                        case "":
                            if false {
                                fmt.Println(p0, p0)
                                p1 = 0
                                for fi15 := 0; fi15 < 20; fi15++ {
                                    p0 = p0
                                    _ = fi15
                                }
                            } else {
                                var v16 string = ((("" + "a") + "foo bar") + "\n\t")
                                _ = v16
                            }
                        case "foo bar":
                            var v17 uint32 = (p2 % ((p2) | 1))
                            _ = v17
                        case "\n\t":
                            p0 -= 4294967295
                        default:
                            if len(v5) > 0 {
                                v5[0] = (fi6 - fi6)
                            }
                        }
                        if ((-2.25 * -2.25) >= 1.5) {
                            v18 := GenS1{F0: uint16(p0), F1: 42}
                            p2 = (p2 / ((p2) | 1))
                            p1 = -1
                            _ = v18
                        }
                        fmt.Println(p2)
                    }
                    switch 438 {
                    case 9223372036854775807:
                        v19 := GenS2{F0: 0, F1: 861, F2: 4294967295}
                        _ = v19
                    case 464:
                        var v20 uint8 = 42
                        _ = v20
                    }
                } else {
                    _ = len(v5)
                }
            }
            v21 := GenS3{F0: (!false)}
            _ = v21
        }
        if true {
            if (42 > 789) {
                p0 = 121
                fmt.Println(p1, p0, p1)
                fmt.Println(p2)
            } else {
                v22 := GenS1{F0: 42, F1: -797}
                for fi23 := 0; fi23 < 19; fi23++ {
                    var v24 int16 = 42
                    v25 := GenS4{F0: 206}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = v22
            }
            var v26 uint16 = 851
            _ = v26
        } else {
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 uint8 = uint8(p2)
                _ = fi27
                _ = v28
            }
        }
        _ = fi6
    }
    var v29 float64 = (1e10 / (1.5 + (-2.25 - float64(3.14159))))
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v5
    return -1
}
func genFn1(p0 string, p1 string) uint16 {
    var v30 int16 = -736
    var v31 int = 0
    fmt.Println(v30, p0, v31)
    v30 -= -1
    _ = p0
    _ = p1
    _ = v30
    _ = v31
    return 650
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 977
    var v3 uint = 42
    var v4 string = v1
    v32 := []int{}
    switch 984 {
    case -1:
        fmt.Println(v3, v1)
    case 2147483647:
        v33 := GenS5{F0: v2}
        _ = v33
    default:
        fmt.Println(v2, v2, v2)
    }
    v34 := GenS1{F0: v2, F1: 818}
    if (false && ((v3 * 924) < v3)) {
        v35 := GenS2{F0: 2, F1: 0, F2: v3}
        _ = v35
    }
    var v36 bool = false
    if len(v32) > 0 {
        v32[0] = -1
    }
    switch "foo bar" {
    case "foo bar":
        var v37 int64 = -9223372036854775808
        _ = v37
    case "":
        v38 := GenS6{F0: 0, F1: (v2 << (((v3 / ((4294967295) | 1))) & 63)), F2: v1}
        _ = v38
    default:
        if (float64(1.5) != -2.25) {
            v39 := GenS1{F0: v2, F1: 129}
            _ = v39
        }
    }
    if len(v32) > 0 {
        v32[0] = -238
    }
    fmt.Println(v1, v36, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v36
    _ = v4
}
