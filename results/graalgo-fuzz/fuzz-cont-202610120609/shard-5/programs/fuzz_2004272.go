package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint64 }
func genFn0(p0 int64, p1 int32, p2 int) uint32 {
    fmt.Println(p1, p2, p2)
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 1
    v3 := make(map[int]int)
    v2 -= v2
    var v4 uint16 = 2
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(fi5, v2, v2)
        v6 := GenS1{F0: 469, F1: -1}
        v7 := GenS1{F0: fi5, F1: -2147483648}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{266, 2147483647, 9223372036854775807}
    switch "a" {
    case "hello":
        switch -272 {
        case -1:
            fmt.Println(v2)
        default:
            v4 = (v4 - v4)
        }
    case "\n\t":
        for fi9 := 0; fi9 < 17; fi9++ {
            _ = len(v3)
            _ = fi9
        }
    }
    v10 := v1[""]
    v4 -= (v4 ^ v4)
    switch ("" + ("\n\t" + "a")) {
    case "a":
        v11 := GenS1{F0: 0, F1: 1}
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 13; fi12++ {
            if (((1 <= -2147483648) && false) || false) {
                var v13 bool = (-128 <= 58)
                for fi14 := 0; fi14 < 8; fi14++ {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    v10 += v10
                    _ = fi14
                }
                _ = v13
            }
            if true {
                for fi15 := 0; fi15 < 11; fi15++ {
                    v16 := v3[1]
                    switch "a" {
                    case "hello":
                        if (genFn0(696, 42, (fi15 ^ (fi15 + v16))) >= v10) {
                            var v17 int = 42
                            _ = v17
                        }
                    case "\n\t":
                        var v18 int32 = -925
                        _ = v18
                    }
                    _ = fi15
                    _ = v16
                }
                for fi19 := 0; fi19 < 14; fi19++ {
                    v2 = v2
                    _ = fi19
                }
            } else {
                for fi20 := 0; fi20 < 9; fi20++ {
                    fmt.Println(v4, fi20, v2)
                    _ = fi20
                }
                switch int(v10) {
                case 127:
                    var v21 int16 = 0
                    _ = v21
                }
                v10 += v10
            }
            _ = fi12
        }
    }
    if len(v8) > 0 {
        v8[0] = -729
    }
    if true {
        switch 947 {
        case 254:
            v1[("hello" + ("\n\t" + "a"))] = v10
        case 0:
            if true {
                var v22 int = 32767
                var v23 uint32 = genFn0(-9223372036854775808, -1, (v22 + 128))
                _ = v22
                _ = v23
            } else {
                v4 = 227
                if true {
                    v24 := GenS2{F0: "a", F1: 2}
                    switch 128 {
                    case -1:
                        v25 := GenS1{F0: 1, F1: -2147483648}
                        _ = v25
                    case -917:
                        v26 := GenS2{F0: "", F1: (v4 - 65535)}
                        _ = v26
                    case 128:
                        for fi27 := 0; fi27 < 17; fi27++ {
                            for fi28 := 0; fi28 < 3; fi28++ {
                                var v29 uint64 = 0
                                v30 := GenS3{F0: int32(v2), F1: (v29 / ((227) | 1))}
                                v31 := GenS1{F0: -1, F1: 1}
                                _ = fi28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            _ = fi27
                        }
                    }
                    v10 = 2
                    _ = v24
                }
            }
        case -1:
            var v32 int16 = -1
            _ = v32
        }
    } else {
        v4 -= (v4 * (v4 ^ (v4 | v4)))
        switch "foo bar" {
        case "":
            fmt.Println(v4, v2)
        case "\n\t":
            switch "\n\t" {
            case "hello":
                _ = len(v8)
            default:
                var v33 bool = false
                _ = v33
            }
        case "foo bar":
            if (42 != 763) {
                v34 := GenS2{F0: ("" + "\n\t"), F1: v4}
                _ = v34
            }
        }
    }
    fmt.Println(v10, v10)
    v2 -= (v2 | (v2 | (v2 - 86)))
    var v35 int = 9223372036854775807
    var v36 int32 = -1
    switch "hello" {
    case "a":
        var v37 float32 = (1.5 - (1.5 * 1e10))
        _ = v37
    case "foo bar":
        fmt.Println(v35, v4, v10)
    }
    if len(v8) > 0 {
        v8[0] = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v8
}
