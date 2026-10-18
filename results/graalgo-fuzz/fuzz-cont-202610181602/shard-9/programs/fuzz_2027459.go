package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint, p1 uint32) uint16 {
    switch "a" {
    case "foo bar":
        if true {
            fmt.Println(p1, p1)
        } else {
            p0 = p0
            p0 = p0
            fmt.Println(p1, p0)
        }
    case "a":
        for fi6 := 0; fi6 < 16; fi6++ {
            if (false && (2147483647 < 390)) {
                if ((((0.1 * 0.0) / (1.5 - 3.14159)) / ((((3.14159 - 3.14159) * float32(float64(0.1))) / float32(1.5)) + 0.1)) > (1.5 + float32(-2.25))) {
                    v7 := []int{0}
                    var v8 uint16 = 912
                    switch (-(fi6)) {
                    case 32768:
                        fmt.Println(p1)
                    case -2147483648:
                        if len(v7) > 0 {
                            v7[0] = fi6
                        }
                    case 1:
                        if (0 <= (fi6 | fi6)) {
                            switch "foo bar" {
                            case "hello":
                                var v9 float64 = -2.25
                                _ = v9
                            default:
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                            for fi10 := 0; fi10 < 1; fi10++ {
                                v8 = 32767
                                _ = fi10
                            }
                        }
                    default:
                        v7 = append(v7, fi6)
                    }
                    _ = v7
                    _ = v8
                }
                var v11 uint32 = 2
                fmt.Println(p1, p0, fi6)
                _ = v11
            }
            var v12 int32 = 42
            for fi13 := 0; fi13 < 5; fi13++ {
                p1 += uint32(p0)
                v14 := GenS1{F0: (p0 * 1), F1: 45}
                var v15 uint8 = 42
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi6
            _ = v12
        }
    }
    switch 127 {
    case -9223372036854775808:
        fmt.Println(p0, p1)
    case 127:
        v16 := []int{-253, -487}
        _ = v16
    case 9223372036854775807:
        for fi17 := 0; fi17 < 14; fi17++ {
            if (!true) {
                p0 += 443
            }
            _ = fi17
        }
    }
    _ = p0
    _ = p1
    return uint16(p1)
}
func genFn1(p0 int16) uint16 {
    for fi18 := 0; fi18 < 1; fi18++ {
        p0 += (p0 % (((p0 % (((p0 - (p0 << ((1) & 63)))) | 1))) | 1))
        fmt.Println(p0, p0)
        v19 := []int{(fi18 >> ((602) & 63)), (fi18 - fi18), (fi18 << ((18446744073709551615) & 63)), fi18}
        _ = fi18
        _ = v19
    }
    fmt.Println(p0)
    fmt.Println(p0)
    if ((0 <= 0) && (true && false)) {
        v20 := []int{int(p0), 0}
        p0 -= -32768
        var v21 uint16 = 1
        _ = v20
        _ = v21
    }
    _ = p0
    return 32767
}

func main() {
    var v1 int16 = -909
    v2 := []string{"foo bar", "foo bar", "\n\t"}
    var v3 float32 = 1e10
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    for fi22 := 0; fi22 < 13; fi22++ {
        v3 = 3.14159
        _ = fi22
    }
    var v23 int = 128
    if true {
        var v24 int64 = 9223372036854775807
        _ = len(v2)
        _ = v24
    }
    var v25 int64 = -1
    v25 -= (v25 >> ((255) & 63))
    var v26 string = ("hello" + "")
    fmt.Println(v1)
    for fi27 := 0; fi27 < 19; fi27++ {
        v25 -= (v25 ^ v25)
        var v28 int32 = -2147483648
        v29 := GenS1{F0: 2, F1: 1}
        _ = fi27
        _ = v28
        _ = v29
    }
    v1 = v1
    var v30 uint = 390
    var v31 uint32 = 2147483647
    var v32 uint = v30
    if (1 < genFn0(v32, 2147483647)) {
        if true {
            v33 := GenS2{F0: 42}
            var v34 int = (v23 & (v23 / (((-(v23))) | 1)))
            _ = v33
            _ = v34
        } else {
            v35 := GenS1{F0: v32, F1: 0}
            _ = v35
        }
    } else {
        switch 32767 {
        case 32768:
            v26 = v26
        case 914:
            for fi36 := 0; fi36 < 2; fi36++ {
                var v37 bool = (2 >= 32767)
                _ = fi36
                _ = v37
            }
        default:
            fmt.Println(v3, v32, v31)
        }
        v25 -= (^(v25))
        v26 = ""
    }
    var v38 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v38
    _ = v4
    _ = v5
}
