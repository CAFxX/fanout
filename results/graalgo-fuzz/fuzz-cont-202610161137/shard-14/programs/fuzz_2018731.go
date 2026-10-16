package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
func genFn0(p0 int8, p1 uint) uint64 {
    p1 -= (p1 + (p1 | 965))
    v6 := []int{127, 9223372036854775807, 1}
    var v7 int16 = 1
    if true {
        var v8 int64 = 0
        switch "\n\t" {
        case "a":
            v9 := GenS1{F0: (p0 ^ p0), F1: 42}
            _ = v9
        }
        v10 := GenS1{F0: p0, F1: int(p1)}
        _ = v10
        _ = v8
    } else {
        if true {
            var v11 uint64 = 44
            _ = v11
        } else {
            var v12 string = ("a" + ("hello" + "a"))
            _ = v12
        }
        _ = len(v6)
        fmt.Println(p0, p0)
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 uint16) uint32 {
    var v13 int8 = int8(p0)
    v14 := []int{32767, 579, 730}
    fmt.Println(p0)
    if len(v14) > 0 {
        v14[0] = 2147483647
    }
    if true {
        v14 = append(v14, 2147483647)
        if (920 <= 2) {
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 int64 = 0
                _ = fi15
                _ = v16
            }
        }
    } else {
        p0 += 0
        if len(v14) > 0 {
            v14[0] = 585
        }
        var v17 int8 = (v13 << ((733) & 63))
        _ = v17
    }
    _ = p0
    _ = v13
    _ = v14
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 string = ""
    v4 := make(map[int]int)
    var v5 int64 = -1
    switch -435 {
    case 128:
        switch 351 {
        case 32768:
            v3 = "foo bar"
        case 0:
            for fi18 := 0; fi18 < 9; fi18++ {
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 int16 = 752
                    switch ("" + v3) {
                    case "\n\t":
                        if false {
                            v21 := []int{(fi19 % (((fi18 * fi19)) | 1))}
                            delete(v4, 1)
                            _ = v21
                        } else {
                            v22 := GenS1{F0: -57, F1: (fi19 % (((fi18 ^ 1)) | 1))}
                            v3 = v3
                            fmt.Println(fi18)
                            v23 := []int{1, (fi18 % ((fi19) | 1)), fi18, (fi19 / ((fi19) | 1))}
                            _ = v22
                            _ = v23
                        }
                    }
                    var v24 int64 = 9223372036854775807
                    _ = fi19
                    _ = v20
                    _ = v24
                }
                v3 = (v3 + (v3 + v3))
                _ = fi18
            }
        }
    case -1:
        v4[1] = 0
    default:
        var v25 float64 = 1e10
        _ = v25
    }
    fmt.Println(v3)
    switch 536 {
    case 727:
        var v26 int32 = 42
        _ = v26
    case -9223372036854775808:
        if true {
            var v27 float64 = 1.5
            v28 := []int{-937, 4294967295, -9223372036854775808, 1}
            v5 += 42
            _ = v27
            _ = v28
        }
    case -2147483648:
        var v29 int8 = 127
        _ = v29
    }
    v30 := []int{-944, 42}
    v3 = "a"
    fmt.Println(v3, v5)
    for fi31 := 0; fi31 < 2; fi31++ {
        v5 = 42
        if (true || (false && false)) {
            v3 = (v3 + (v3 + ("foo bar" + v3)))
            v5 += (^(v5))
        }
        _ = fi31
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        v3 = ((v3 + (v3 + v3)) + (((v3 + v3) + v3) + v3))
        _ = fi32
    }
    v5 -= (-(v5))
    v3 = "hello"
    v4[-2147483648] = 0
    var v33 int8 = 90
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
}
