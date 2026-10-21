package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint64 }
func genFn0(p0 bool, p1 int32, p2 uint8) bool {
    if false {
        p2 -= 1
    } else {
        var v4 float32 = 0.0
        v5 := []int{-2147483648, 52, 128}
        v6 := GenS1{F0: -2147483648, F1: p2}
        _ = v4
        _ = v5
        _ = v6
    }
    fmt.Println(p1, p2, p1)
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(p2, p1, p1)
            var v9 int16 = 1000
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := []int{0, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p0
}
func genFn1(p0 uint32, p1 int64) uint {
    for fi11 := 0; fi11 < 4; fi11++ {
        p1 = 42
        v12 := GenS1{F0: 128, F1: 35}
        var v13 bool = (1 > -2147483648)
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(p1, p1)
    var v14 uint32 = 645
    if false {
        var v15 uint16 = 264
        _ = v15
    } else {
        if true {
            switch 1 {
            case 1:
                var v16 string = "foo bar"
                _ = v16
            case 42:
                var v17 float32 = float32(0.1)
                _ = v17
            case -9223372036854775808:
                var v18 uint8 = 127
                _ = v18
            }
            switch -1 {
            case 0:
                if (-2147483648 != 2147483647) {
                    switch -2147483648 {
                    case 42:
                        if false {
                            p1 += 9223372036854775807
                            if (false && true) {
                                v14 -= (v14 >> ((255) & 63))
                            }
                        }
                    case 32768:
                        for fi19 := 0; fi19 < 19; fi19++ {
                            v20 := []int{}
                            _ = fi19
                            _ = v20
                        }
                    case 746:
                        var v21 bool = ((!(1 > 255)) && ((!(1 == 42)) && true))
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 9; fi22++ {
                            v23 := []int{32768}
                            if len(v23) > 0 {
                                v23[0] = (fi22 | (fi22 / ((4294967295) | 1)))
                            }
                            _ = fi22
                            _ = v23
                        }
                    }
                    v24 := GenS2{F0: 624, F1: -128, F2: 18446744073709551615}
                    fmt.Println(v14, p0, p1)
                    _ = v24
                } else {
                    var v25 int8 = -128
                    var v26 int = 0
                    _ = v25
                    _ = v26
                }
            }
        } else {
            switch "a" {
            case "foo bar":
                for fi27 := 0; fi27 < 8; fi27++ {
                    if false {
                        fmt.Println(fi27, p1, v14)
                        var v28 uint32 = (p0 >> ((255) & 63))
                        var v29 int8 = -128
                        _ = v28
                        _ = v29
                    } else {
                        var v30 float64 = 3.14159
                        _ = v30
                    }
                    _ = fi27
                }
            case "a":
                v14 = (p0 - 0)
            }
            v14 += 1
        }
    }
    _ = p0
    _ = p1
    _ = v14
    return 4294967295
}

func main() {
    var v1 uint64 = 2
    v2 := []int{}
    v3 := make(map[uint32]int)
    var v31 string = ("hello" + "hello")
    v32 := GenS1{F0: 1, F1: 1}
    switch 127 {
    case -32:
        for fi33 := 0; fi33 < 18; fi33++ {
            v1 += v1
            _ = fi33
        }
    case 2147483647:
        v1 = (v1 / ((v1) | 1))
    }
    if true {
        v31 = ""
        fmt.Println(v31, v1)
        v31 = (((("hello" + "") + ((((v31 + ("" + v31)) + "a") + "a") + (((v31 + v31) + (v31 + "\n\t")) + v31))) + v31) + (("hello" + v31) + (v31 + v31)))
    } else {
        v34 := v3[2]
        for fi35 := 0; fi35 < 15; fi35++ {
            v31 = v31
            _ = fi35
        }
        _ = v34
    }
    var v36 int = -565
    for fi37 := 0; fi37 < 11; fi37++ {
        v36 = fi37
        _ = fi37
    }
    v31 = v31
    fmt.Println(v31)
    var v38 string = (v31 + (v31 + v31))
    v31 = (v38 + (v38 + (v38 + v31)))
    delete(v3, 0)
    v31 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v38
}
