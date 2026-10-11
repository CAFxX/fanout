package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int; F1 uint8; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
func genFn0(p0 bool, p1 uint, p2 int64) uint {
    v5 := []int{118, 128, -9223372036854775808, 127}
    if p0 {
        fmt.Println(p1, p2, p2)
        var v6 uint8 = 255
        var v7 string = "a"
        _ = v6
        _ = v7
    } else {
        if p0 {
            switch "foo bar" {
            case "\n\t":
                p1 = (^(p1))
            case "foo bar":
                for fi8 := 0; fi8 < 7; fi8++ {
                    var v9 uint = p1
                    _ = fi8
                    _ = v9
                }
            case "a":
                v5 = append(v5, 368)
            default:
                for fi10 := 0; fi10 < 9; fi10++ {
                    p2 -= (p2 >> ((p1) & 63))
                    switch (((("" + "hello") + ("a" + "\n\t")) + "\n\t") + ("\n\t" + "hello")) {
                    case "a":
                        v11 := GenS1{F0: 18446744073709551615, F1: 4294967295}
                        _ = v11
                    case "":
                        for fi12 := 0; fi12 < 11; fi12++ {
                            if true {
                                if ((1 == p1) || true) {
                                    fmt.Println(fi12)
                                } else {
                                    if (p0 && true) {
                                        if len(v5) > 0 {
                                            v5[0] = fi12
                                        }
                                        var v13 int8 = -82
                                        var v14 uint64 = uint64(p1)
                                        _ = v13
                                        _ = v14
                                    }
                                }
                                switch 511 {
                                case 128:
                                    p2 = 1
                                case 9223372036854775807:
                                    var v15 float64 = (float64((float32(0.0) - 3.14159)) / (((-2.25 - (float64((-2.25 / 3.14159)) / 3.14159)) - 3.14159) * (0.1 * -2.25)))
                                    v16 := GenS2{F0: -706, F1: 63, F2: ((p2 + (p2 + p2)) == (p2 ^ p2))}
                                    _ = v15
                                    _ = v16
                                default:
                                    fmt.Println(p0)
                                }
                                v17 := GenS3{F0: 42}
                                _ = v17
                            }
                            _ = fi12
                        }
                    }
                    _ = fi10
                }
            }
        }
        if p0 {
            v18 := GenS1{F0: 4294967295, F1: 1}
            _ = v18
        }
        if len(v5) > 0 {
            v5[0] = int(p1)
        }
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        fmt.Println(p1, p1, p1)
        _ = fi19
    }
    v5 = append(v5, -299)
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}
func genFn1(p0 uint8, p1 int64) uint16 {
    v20 := []int{2147483647, 32768, 245}
    v21 := GenS4{F0: 679}
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"", ("a" + ""), "hello"}
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    var v22 uint32 = 2
    v23 := GenS3{F0: 18446744073709551615}
    v24 := GenS2{F0: 395, F1: 0, F2: (792 <= 219)}
    for fi25 := 0; fi25 < 3; fi25++ {
        delete(v1, (v22 ^ v22))
        _ = fi25
    }
    var v26 bool = (((1e10 / 1.0) == float32((0.1 / 3.14159))) && (0 == 65535))
    v22 = v22
    v2 = append(v2, "foo bar")
    var v27 int16 = 42
    _ = len(v3)
    for fi28 := 0; fi28 < 11; fi28++ {
        _ = len(v2)
        _ = fi28
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    fmt.Println(v27, v26, v27)
    for fi29 := 0; fi29 < 7; fi29++ {
        fmt.Println(v27, v26, v27)
        v27 = (v27 & (v27 * (v27 - (v27 & v27))))
        _ = fi29
    }
    switch "foo bar" {
    case "foo bar":
        v22 += 4294967295
    case "\n\t":
        var v30 int8 = 13
        _ = v30
    }
    var v31 int16 = v27
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
}
