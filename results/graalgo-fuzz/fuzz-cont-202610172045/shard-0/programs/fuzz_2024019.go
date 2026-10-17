package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint16, p1 int64, p2 uint64) bool {
    p0 += 663
    v6 := []int{}
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: 22, F1: (p2 % (((p2 >> ((uint(p2)) & 63))) | 1))}
            var v10 int8 = -1
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 16; fi11++ {
            fmt.Println(fi7)
            var v12 int8 = 1
            var v13 string = "foo bar"
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 int32 = 42
        _ = fi7
        _ = v14
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}

func main() {
    v1 := []int{-762, 4294967295}
    var v2 uint8 = 2
    var v3 int8 = -1
    v4 := []string{}
    var v5 int16 = -676
    if false {
        v5 = v5
        v15 := GenS2{F0: 18446744073709551615, F1: 255}
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 uint8 = v2
            var v18 int32 = 0
            var v19 int = fi16
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = v15
    } else {
        if false {
            for fi20 := 0; fi20 < 11; fi20++ {
                switch "hello" {
                case "foo bar":
                    var v21 bool = ((!(2147483647 >= -2147483648)) || false)
                    _ = v21
                case "":
                    if true {
                        for fi22 := 0; fi22 < 5; fi22++ {
                            v23 := GenS3{F0: ("" + "hello"), F1: 0}
                            var v24 uint = 0
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        v4 = append(v4, ("" + "\n\t"))
                    }
                case "hello":
                    switch -2147483648 {
                    case 883:
                        v4 = append(v4, "a")
                    default:
                        var v25 int64 = 9223372036854775807
                        _ = v25
                    }
                }
                for fi26 := 0; fi26 < 6; fi26++ {
                    fmt.Println(v5, fi20, v3)
                    switch -855 {
                    case -1:
                        var v27 int64 = 1
                        _ = v27
                    default:
                        v5 = (-(v5))
                        fmt.Println(fi26)
                        v28 := GenS4{F0: 0}
                        _ = v28
                    }
                    if (0.0 != -2.25) {
                        v2 -= (v2 % ((v2) | 1))
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = fi26
                }
                _ = fi20
            }
            v2 += v2
            v29 := GenS5{F0: v5, F1: 644}
            _ = v29
        } else {
            v2 -= v2
        }
        fmt.Println(v3)
        if len(v4) > 0 {
            v4[0] = ("" + "foo bar")
        }
    }
    if (!false) {
        var v30 uint16 = 726
        for fi31 := 0; fi31 < 20; fi31++ {
            if (0 == -9223372036854775808) {
                v3 -= 0
                if ((2147483648 != 0) || (42 != 18446744073709551615)) {
                    v32 := GenS6{F0: (v5 >> ((594) & 63))}
                    v30 += v30
                    fmt.Println(v30, v30)
                    _ = v32
                }
            }
            _ = fi31
        }
        v33 := GenS3{F0: "", F1: 42}
        _ = v30
        _ = v33
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v1) > 0 {
        v1[0] = 316
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 = append(v1, 9223372036854775807)
    v1 = append(v1, -9223372036854775808)
    v5 = (v5 % (((^(v5))) | 1))
    v34 := GenS2{F0: 42, F1: 599}
    if len(v1) > 0 {
        v1[0] = 42
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
