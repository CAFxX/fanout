package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 string; F1 uint8; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int16; F2 bool }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint16; F1 int32 }
func genFn0(p0 int, p1 int32, p2 uint32) uint64 {
    switch (p0 % ((p0) | 1)) {
    case 4294967295:
        fmt.Println(p0, p2, p2)
    case -1:
        var v7 uint16 = 42
        _ = v7
    case 2147483647:
        var v8 int32 = (p1 + p1)
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        p1 += (p1 << ((65535) & 63))
        v10 := GenS1{F0: 18446744073709551615}
        _ = fi9
        _ = v10
    }
    switch 127 {
    case 128:
        v11 := GenS2{F0: (!false), F1: 42}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 160
}

func main() {
    v1 := []int{128, -900, 42}
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    v4 := []uint32{2147483648, 465, 199, 0}
    var v5 float64 = 3.14159
    v6 := []int{}
    v5 += 1e10
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 6; fi13++ {
            v5 = v5
            v14 := GenS3{F0: ("hello" + "\n\t"), F1: 37, F2: 0}
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            v5 += ((((((v5 - 0.1) + float64(3.14159)) + 1.5) - 0.1) - ((v5 - v5) - (-2.25 / 1.5))) + v5)
            fmt.Println(fi12, fi12, fi12)
            v5 += (((v5 - (float64(3.14159) * v5)) + (1e10 - float64((3.14159 * 0.1)))) * -2.25)
            _ = fi15
        }
        v5 -= (v5 / (float64(-2.25) * (1e10 - (v5 * -2.25))))
        _ = fi12
    }
    var v16 bool = (694 == 370)
    switch 32768 {
    case -759:
        v17 := GenS3{F0: "\n\t", F1: 1, F2: 255}
        _ = v17
    case -9223372036854775808:
        if false {
            v5 -= v5
            if v16 {
                v5 += 3.14159
                v18 := GenS4{F0: 3}
                v16 = (((3.14159 / 3.14159) / float32(v5)) > float32((v5 / 1.5)))
                _ = v18
            } else {
                var v19 float64 = -2.25
                _ = v19
            }
        } else {
            if ((false || false) || false) {
                fmt.Println(v5)
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 string = "foo bar"
                    if v16 {
                        for fi22 := 0; fi22 < 6; fi22++ {
                            v23 := GenS5{F0: v21, F1: -32768, F2: (127 <= -96)}
                            fmt.Println(v5)
                            if v16 {
                                var v24 int64 = -9223372036854775808
                                v25 := GenS4{F0: fi22}
                                v24 = 1
                                _ = v24
                                _ = v25
                            }
                            _ = fi22
                            _ = v23
                        }
                        v21 = v21
                        v26 := GenS6{F0: 255, F1: 127}
                        _ = v26
                    } else {
                        v16 = (-1 == 42)
                    }
                    _ = fi20
                    _ = v21
                }
                fmt.Println(v16, v16)
            }
        }
    }
    if (-1 > 42) {
        var v27 uint32 = 264
        switch "hello" {
        case "":
            switch -237 {
            case 1:
                v1 = append(v1, -1)
            }
        default:
            v16 = (0 != 4294967295)
        }
        switch "" {
        case "a":
            for fi28 := 0; fi28 < 11; fi28++ {
                for fi29 := 0; fi29 < 6; fi29++ {
                    v5 -= (1e10 + float64(-2.25))
                    _ = fi29
                }
                if v16 {
                    v27 = (v27 - v27)
                } else {
                    var v30 int = fi28
                    if (true || true) {
                        v31 := GenS7{F0: 186}
                        _ = v31
                    } else {
                        v30 -= (v30 % ((v30) | 1))
                    }
                    if ((("foo bar" > "a") && (-2147483648 == (v30 * fi28))) || (!v16)) {
                        if (!false) {
                            if len(v6) > 0 {
                                v6[0] = (-(fi28))
                            }
                        }
                        if v16 {
                            for fi32 := 0; fi32 < 14; fi32++ {
                                var v33 int = (^(fi32))
                                v34 := GenS8{F0: 2, F1: -1}
                                var v35 uint64 = 81
                                _ = fi32
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            fmt.Println(fi28, v5)
                        } else {
                            if (!true) {
                                v5 -= ((float64(1e10) - float64(-2.25)) / (float64(0.1) - v5))
                            } else {
                                v27 += v27
                            }
                        }
                    }
                    _ = v30
                }
                fmt.Println(v27)
                _ = fi28
            }
        case "hello":
            v4 = append(v4, v27)
        default:
            v36 := GenS4{F0: 127}
            _ = v36
        }
        _ = v27
    }
    fmt.Println(v5, v16)
    fmt.Println(v5)
    var v37 float32 = ((0.0 / 3.14159) * 3.14159)
    fmt.Println(v16, v5, v5)
    v4 = append(v4, 2147483647)
    switch "\n\t" {
    case "a":
        var v38 string = "foo bar"
        _ = v38
    }
    if v16 {
        var v39 float64 = -2.25
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
