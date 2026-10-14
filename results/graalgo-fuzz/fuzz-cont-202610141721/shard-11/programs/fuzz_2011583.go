package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 string, p1 bool, p2 bool) int32 {
    v4 := []int{32767, 42}
    if p1 {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        var v5 uint8 = 149
        fmt.Println(v5)
        _ = v5
    }
    switch 180 {
    case 712:
        for fi6 := 0; fi6 < 3; fi6++ {
            fmt.Println(p0, p1, p2)
            p1 = ((-9223372036854775808 < 9223372036854775807) || p1)
            _ = fi6
        }
    case -687:
        v4 = append(v4, 42)
    }
    switch "foo bar" {
    case "\n\t":
        for fi7 := 0; fi7 < 5; fi7++ {
            if (0 == -605) {
                var v8 int16 = -1
                v4 = append(v4, (fi7 & (fi7 * fi7)))
                _ = v8
            } else {
                var v9 string = (((p0 + "\n\t") + "") + "\n\t")
                v4 = append(v4, (fi7 % ((fi7) | 1)))
                p1 = (128 < 2)
                _ = v9
            }
            var v10 uint8 = 255
            _ = fi7
            _ = v10
        }
    default:
        v11 := GenS1{F0: 4294967295, F1: 409}
        _ = v11
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -723
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{165, 65535, 2147483647, 42}
    v3 := []string{}
    var v12 float64 = ((0.0 + 0.1) * 0.1)
    var v13 float32 = (3.14159 * 1e10)
    fmt.Println(v13, v13)
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int = (-(fi14))
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 18446744073709551615, F1: 2}
    for fi17 := 0; fi17 < 14; fi17++ {
        v18 := GenS1{F0: 2, F1: 65535}
        v19 := v1[fi17]
        _ = fi17
        _ = v18
        _ = v19
    }
    if true {
        v1[461] = 42
        var v20 uint64 = 4294967295
        switch "foo bar" {
        case "":
            _ = len(v3)
        case "foo bar":
            v21 := GenS1{F0: v20, F1: 0}
            _ = v21
        case "\n\t":
            var v22 uint = 76
            _ = v22
        }
        _ = v20
    }
    v12 += (float64(1.5) / (float64((v13 * v13)) / v12))
    var v23 float64 = v12
    v24 := GenS1{F0: 930, F1: 208}
    if true {
        v12 = v23
        switch 2147483647 {
        case 292:
            if (42 != 0) {
                v23 = 1.5
                fmt.Println(v13, v12, v23)
            }
        case -798:
            if (4294967295 != 302) {
                for fi25 := 0; fi25 < 7; fi25++ {
                    if (!true) {
                        if false {
                            fmt.Println(fi25)
                        } else {
                            if (1 < 1) {
                                v23 = float64(v13)
                                for fi26 := 0; fi26 < 12; fi26++ {
                                    var v27 uint32 = 850
                                    v28 := GenS2{F0: 130, F1: -122, F2: 144}
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                }
                                fmt.Println(v23, v23, v13)
                            } else {
                                v29 := GenS2{F0: 255, F1: 127, F2: 42}
                                var v30 uint = 558
                                _ = v29
                                _ = v30
                            }
                            var v31 int32 = -54
                            _ = v31
                        }
                    }
                    _ = fi25
                }
                var v32 uint32 = 423
                _ = v32
            }
        case 974:
            v13 = (0.0 / v13)
        default:
            var v33 int16 = 1000
            _ = v33
        }
    } else {
        _ = len(v1)
        for fi34 := 0; fi34 < 3; fi34++ {
            if false {
                switch (fi34 * -9223372036854775808) {
                case 32768:
                    switch "hello" {
                    case "hello":
                        v23 += 0.0
                    case "\n\t":
                        var v35 float64 = 1e10
                        _ = v35
                    default:
                        var v36 int16 = 1000
                        _ = v36
                    }
                case 0:
                    var v37 uint = 2
                    _ = v37
                }
            } else {
                fmt.Println(fi34, v23, v23)
            }
            var v38 int32 = 2147483647
            _ = len(v2)
            _ = fi34
            _ = v38
        }
        v39 := GenS3{F0: 42}
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
