package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
func genFn0(p0 int8) int16 {
    p0 -= (p0 << ((340) & 63))
    var v6 int = 9223372036854775807
    p0 -= int8(v6)
    v7 := []int{v6}
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 string, p1 uint16, p2 uint8) int32 {
    fmt.Println(p1)
    v8 := GenS1{F0: 1, F1: 9223372036854775807}
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int8 = -128
        _ = fi9
        _ = v10
    }
    var v11 int = -220
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return -558
}

func main() {
    v1 := []int{4294967295}
    var v2 uint16 = 2
    v3 := []uint32{863, 193}
    var v4 uint32 = 656
    var v5 uint = 2
    for fi12 := 0; fi12 < 9; fi12++ {
        v1 = append(v1, 0)
        var v13 int8 = 42
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v5 = (v5 >> ((v5) & 63))
        _ = fi14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        switch "\n\t" {
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "a":
            v5 += (v5 * 357)
        }
        _ = fi15
    }
    fmt.Println(v2, v2, v5)
    switch 2147483647 {
    case 32767:
        v1 = append(v1, 658)
    default:
        fmt.Println(v5)
    }
    switch "hello" {
    case "a":
        for fi16 := 0; fi16 < 18; fi16++ {
            if (!false) {
                v2 += 896
                if ((int8(v2) > 0) && (!(-32768 <= 42))) {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        if (!("\n\t" >= "hello")) {
                            v18 := GenS1{F0: 1, F1: fi16}
                            v4 -= (v4 / ((v4) | 1))
                            var v19 int = (fi17 + -2147483648)
                            v4 += v4
                            var v20 int32 = genFn1(("hello" + "a"), 0, 128)
                            _ = v18
                            _ = v19
                            _ = v20
                        } else {
                            if false {
                                _ = len(v3)
                                fmt.Println(v5, fi17)
                                fmt.Println(v2, v4, fi16)
                            }
                            var v21 float64 = 3.14159
                            v21 = (((((v21 / 1.5) * v21) + (v21 / v21)) + v21) * v21)
                            _ = v21
                        }
                        if true {
                            if false {
                                var v22 int16 = 0
                                var v23 int64 = 0
                                _ = v22
                                _ = v23
                            }
                            if len(v3) > 0 {
                                v3[0] = (v4 + (v4 - 2147483647))
                            }
                        } else {
                            var v24 int = fi17
                            _ = len(v3)
                            _ = v24
                        }
                        v4 -= 411
                        _ = fi17
                    }
                }
                var v25 uint8 = 44
                _ = v25
            } else {
                v4 = uint32(v2)
                v26 := GenS1{F0: 1, F1: fi16}
                v4 -= (v4 >> ((v5) & 63))
                _ = v26
            }
            _ = fi16
        }
    case "foo bar":
        v27 := GenS1{F0: -2147483648, F1: 127}
        _ = v27
    }
    v2 = v2
    v2 -= v2
    if false {
        var v28 string = "\n\t"
        v29 := GenS1{F0: 127, F1: 128}
        _ = v28
        _ = v29
    }
    var v30 int16 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
