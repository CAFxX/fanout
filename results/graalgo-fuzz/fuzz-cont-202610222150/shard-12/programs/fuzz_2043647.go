package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int16, p1 int8, p2 uint32) string {
    switch "a" {
    case "foo bar":
        v7 := []int{}
        _ = v7
    case "hello":
        switch 590 {
        case -25:
            v8 := []int{32767, -941}
            _ = v8
        case 22:
            if true {
                var v9 string = "foo bar"
                v10 := []int{}
                var v11 float64 = (3.14159 + 0.1)
                _ = v10
                _ = v11
                _ = v9
            } else {
                v12 := []int{127, -726}
                switch "hello" {
                case "a":
                    v13 := GenS1{F0: 556, F1: 2, F2: (p1 % (((^(p1))) | 1))}
                    _ = v13
                default:
                    if (0.0 <= 1.5) {
                        if len(v12) > 0 {
                            v12[0] = -582
                        }
                        v14 := GenS1{F0: 761, F1: 32767, F2: (p1 << ((0) & 63))}
                        _ = v14
                    }
                }
                _ = v12
            }
        }
    default:
        v15 := []int{-9223372036854775808, 0}
        _ = v15
    }
    p0 = (-(p0))
    p1 -= p1
    switch "foo bar" {
    case "a":
        v16 := []int{32768, 32768, 498}
        _ = v16
    case "\n\t":
        for fi17 := 0; fi17 < 4; fi17++ {
            if true {
                var v18 float64 = 3.14159
                _ = v18
            } else {
                switch fi17 {
                case 42:
                    v19 := []int{(^(fi17)), fi17, 0, (-(fi17))}
                    _ = v19
                case -9223372036854775808:
                    var v20 uint8 = 87
                    _ = v20
                }
                p1 -= (p1 | (p1 | p1))
                if (!(0.1 > float64(1e10))) {
                    fmt.Println(p2)
                }
            }
            fmt.Println(p1)
            _ = fi17
        }
    default:
        v21 := []int{-141, 4294967295, 0, -383}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    return (((("a" + ("a" + "a")) + "\n\t") + "") + "")
}
func genFn1(p0 uint8, p1 int) uint8 {
    p1 = p1
    for fi22 := 0; fi22 < 14; fi22++ {
        if (0 != 261) {
            fmt.Println(p0)
            for fi23 := 0; fi23 < 6; fi23++ {
                v24 := GenS2{F0: -463, F1: 2}
                p1 += fi23
                fmt.Println(fi23)
                v25 := GenS3{F0: -63}
                fmt.Println(p0)
                _ = fi23
                _ = v24
                _ = v25
            }
        } else {
            v26 := []int{(p1 & fi22)}
            _ = v26
        }
        p0 = (p0 << ((2) & 63))
        var v27 uint = 4294967295
        _ = fi22
        _ = v27
    }
    var v28 uint32 = 42
    _ = p0
    _ = p1
    _ = v28
    return 42
}

func main() {
    var v1 bool = (!(false && (2 < 18446744073709551615)))
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    v5 := []string{"\n\t", "a"}
    v6 := make(map[string]string)
    v1 = v1
    v5 = append(v5, "")
    if v1 {
        v29 := GenS2{F0: 42, F1: 204}
        switch "\n\t" {
        case "foo bar":
            v6["\n\t"] = "hello"
        }
        if v1 {
            for fi30 := 0; fi30 < 17; fi30++ {
                switch "foo bar" {
                case "foo bar":
                    switch (^(fi30)) {
                    case -244:
                        var v31 int = fi30
                        _ = v31
                    case 4294967295:
                        v1 = (0 == int32(fi30))
                    default:
                        v1 = v1
                    }
                }
                var v32 string = ((("foo bar" + "\n\t") + "hello") + ("a" + "hello"))
                _ = fi30
                _ = v32
            }
        } else {
            var v33 string = ("a" + ("foo bar" + "\n\t"))
            var v34 float32 = (float32((((-2.25 - (-2.25 * 1e10)) / 1e10) / -2.25)) - ((0.1 - (3.14159 / 0.1)) + float32(3.14159)))
            _ = v33
            _ = v34
        }
        _ = v29
    }
    fmt.Println(v1)
    var v35 float32 = 0.1
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    fmt.Println(v1)
    switch "a" {
    case "":
        for fi36 := 0; fi36 < 6; fi36++ {
            if false {
                var v37 string = ("foo bar" + ("hello" + "\n\t"))
                _ = v37
            } else {
                v5 = append(v5, "a")
            }
            var v38 int8 = 127
            _ = fi36
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
