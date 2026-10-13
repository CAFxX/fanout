package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint16; F1 int16 }
func genFn0(p0 string, p1 uint64) int64 {
    v5 := GenS1{F0: false, F1: 128, F2: p1}
    v6 := []int{}
    var v7 int = 225
    if false {
        _ = len(v6)
    }
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int16, p1 string) uint8 {
    p0 = (p0 ^ p0)
    v8 := GenS2{F0: (105 <= 58), F1: 18446744073709551615}
    var v9 uint8 = 127
    switch 9223372036854775807 {
    case 491:
        if ((0.0 / -2.25) > -2.25) {
            var v10 uint64 = 0
            if true {
                var v11 int32 = 1
                _ = v11
            } else {
                switch 9223372036854775807 {
                case 32767:
                    p1 = p1
                case 128:
                    v12 := []int{1}
                    _ = v12
                case 0:
                    v10 -= v10
                }
                p1 = (p1 + p1)
                switch p1 {
                case "foo bar":
                    v10 = (v10 >> ((695) & 63))
                }
            }
            _ = v10
        } else {
            switch 32767 {
            case 0:
                switch p1 {
                case "foo bar":
                    switch 128 {
                    case 964:
                        if (42 < -1) {
                            if (!true) {
                                v13 := GenS3{F0: 42}
                                _ = v13
                            }
                            v14 := GenS2{F0: false, F1: 255}
                            v15 := GenS3{F0: (-(p0))}
                            _ = v14
                            _ = v15
                        }
                    case -2147483648:
                        v9 += (v9 * (v9 % ((128) | 1)))
                    case 978:
                        v16 := GenS4{F0: 1, F1: 757}
                        _ = v16
                    default:
                        fmt.Println(p1, v9)
                    }
                default:
                    v17 := GenS2{F0: false, F1: 2}
                    _ = v17
                }
            case -990:
                if (1 > 0) {
                    p0 -= 42
                }
            case 800:
                v9 += v9
            default:
                fmt.Println(p0, p1)
            }
        }
    }
    if true {
        p0 = (^(p0))
        var v18 bool = (!false)
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 bool = true
    v2 := []string{(("\n\t" + ("\n\t" + (("a" + "a") + ""))) + "a"), "foo bar", "foo bar"}
    var v3 uint32 = 65535
    var v4 int16 = 920
    if true {
        v1 = (-1 == 42)
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 int = -466
            switch ((("hello" + "a") + "\n\t") + "hello") {
            case "a":
                v3 -= (^(v3))
            case "foo bar":
                v22 := GenS4{F0: -1, F1: 42}
                _ = v22
            case "":
                v3 = (v3 | v3)
            }
            var v23 uint8 = 137
            v24 := GenS2{F0: (float64(-2.25) < 0.0), F1: 0}
            _ = fi20
            _ = v21
            _ = v23
            _ = v24
        }
        var v25 int16 = v4
        _ = fi19
        _ = v25
    }
    switch 32768 {
    case 4294967295:
        fmt.Println(v4)
    case -505:
        v26 := GenS4{F0: 127, F1: 872}
        _ = v26
    }
    fmt.Println(v1, v3)
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 uint64 = 42
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    v2 = append(v2, "")
    fmt.Println(v3)
    var v30 int = 696
    if v1 {
        v4 = v4
    } else {
        v31 := GenS3{F0: v4}
        _ = v31
    }
    var v32 uint64 = 1
    for fi33 := 0; fi33 < 14; fi33++ {
        v34 := GenS2{F0: (2147483647 <= 42), F1: 2}
        v35 := GenS5{F0: 0, F1: v4}
        if ((v3 - v3) != 2) {
            fmt.Println(v30, v32)
            switch "hello" {
            case "\n\t":
                var v36 uint64 = v32
                _ = v36
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        } else {
            var v37 bool = false
            _ = v37
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println(v30, v4, v32)
    var v38 uint16 = 926
    v1 = true
    var v39 int = (v30 << ((uint(v32)) & 63))
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v38
    _ = v39
    _ = v4
}
