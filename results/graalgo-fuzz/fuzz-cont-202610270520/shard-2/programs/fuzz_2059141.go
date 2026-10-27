package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 string }
func genFn0(p0 uint64) uint64 {
    var v5 uint8 = 127
    fmt.Println(p0, p0)
    _ = p0
    _ = v5
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    v3 := []string{("hello" + ""), ("a" + "hello"), "foo bar", "\n\t"}
    var v4 int32 = 1
    v4 -= 498
    fmt.Println(v4)
    if false {
        v6 := v2[128]
        _ = v6
    } else {
        v4 += 42
        var v7 int16 = -32768
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        v3 = append(v3, "a")
        _ = fi8
    }
    var v9 uint64 = 154
    delete(v2, -9223372036854775808)
    v4 -= (v4 & 1)
    v1[2] = ""
    v4 = (-(v4))
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v9, v4)
        _ = fi10
    }
    if false {
        v11 := GenS1{F0: v9}
        var v12 float64 = 0.1
        _ = v11
        _ = v12
    }
    fmt.Println(v4)
    fmt.Println(v4, v4)
    if (!(!(300 < 0))) {
        if (42 >= 128) {
            v13 := GenS1{F0: 18446744073709551615}
            _ = v13
        } else {
            fmt.Println(v4)
            v14 := GenS2{F0: -1, F1: 0}
            for fi15 := 0; fi15 < 13; fi15++ {
                fmt.Println(v4, v9)
                var v16 uint32 = 2147483647
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        var v17 uint16 = 984
        _ = v17
    }
    var v18 float64 = (-2.25 * (-2.25 / 1e10))
    switch -999 {
    case 0:
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 string = "\n\t"
                _ = fi20
                _ = v21
            }
            v18 = v18
            v9 -= v9
            _ = fi19
        }
    case 733:
        if (91 < 42) {
            v4 -= int32(v9)
            switch -301 {
            case 9223372036854775807:
                for fi22 := 0; fi22 < 5; fi22++ {
                    if false {
                        v23 := GenS2{F0: -40, F1: 1}
                        switch "hello" {
                        case "\n\t":
                            if true {
                                if false {
                                    var v24 uint64 = genFn0(genFn0(18446744073709551615))
                                    _ = v24
                                }
                                if (false && true) {
                                    var v25 uint = uint(v9)
                                    _ = v25
                                }
                            } else {
                                v3 = append(v3, "")
                            }
                        case "foo bar":
                            var v26 int8 = 127
                            _ = v26
                        case "a":
                            switch ("" + "foo bar") {
                            case "a":
                                v27 := GenS3{F0: 2, F1: 255, F2: "a"}
                                _ = v27
                            case "\n\t":
                                var v28 uint = 646
                                _ = v28
                            case "":
                                var v29 uint64 = 0
                                _ = v29
                            }
                        default:
                            var v30 float64 = v18
                            _ = v30
                        }
                        var v31 bool = (!true)
                        _ = v23
                        _ = v31
                    } else {
                        v32 := v2[fi22]
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 10; fi33++ {
                        var v34 uint8 = 2
                        _ = fi33
                        _ = v34
                    }
                    _ = fi22
                }
            case 397:
                v35 := GenS4{F0: 151}
                _ = v35
            }
        } else {
            v36 := GenS5{F0: "hello", F1: "foo bar"}
            _ = v36
        }
    default:
        var v37 int64 = 0
        _ = v37
    }
    v18 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
