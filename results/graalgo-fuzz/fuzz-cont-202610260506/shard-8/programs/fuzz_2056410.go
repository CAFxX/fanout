package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint; F2 int32 }
type GenS4 struct { F0 uint64; F1 bool }
func genFn0(p0 int8) string {
    p0 -= (^(p0))
    if false {
        var v4 uint = 0
        _ = v4
    }
    if false {
        var v5 uint16 = 51
        var v6 uint8 = 127
        v7 := []int{2147483647, 990}
        _ = v5
        _ = v6
        _ = v7
    } else {
        switch -865 {
        case 128:
            p0 = p0
        case 1:
            p0 = (p0 * p0)
        default:
            switch "a" {
            case "":
                switch "a" {
                case "a":
                    v8 := GenS1{F0: "foo bar", F1: 138}
                    _ = v8
                case "":
                    if false {
                        v9 := []int{1, -1, 32767, 128}
                        _ = v9
                    } else {
                        v10 := GenS2{F0: 0, F1: 1, F2: 1}
                        if (!(2147483648 >= 2147483648)) {
                            var v11 int64 = 42
                            v11 -= v11
                            _ = v11
                        } else {
                            var v12 int64 = -9223372036854775808
                            _ = v12
                        }
                        switch 127 {
                        case 2147483647:
                            p0 -= p0
                        default:
                            v13 := GenS1{F0: "a", F1: 127}
                            _ = v13
                        }
                        _ = v10
                    }
                case "foo bar":
                    if true {
                        p0 += (p0 >> ((18446744073709551615) & 63))
                    }
                }
            default:
                for fi14 := 0; fi14 < 5; fi14++ {
                    if true {
                        for fi15 := 0; fi15 < 11; fi15++ {
                            var v16 bool = true
                            _ = fi15
                            _ = v16
                        }
                        switch (("" + "\n\t") + "hello") {
                        case "foo bar":
                            fmt.Println(fi14)
                        case "hello":
                            p0 += (p0 >> ((0) & 63))
                        case "\n\t":
                            v17 := []int{fi14, 128, (fi14 / (((fi14 / ((fi14) | 1))) | 1))}
                            _ = v17
                        default:
                            for fi18 := 0; fi18 < 2; fi18++ {
                                var v19 int8 = p0
                                v20 := GenS2{F0: 42, F1: 32767, F2: 42}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                        }
                        for fi21 := 0; fi21 < 1; fi21++ {
                            p0 = p0
                            fmt.Println(p0)
                            _ = fi21
                        }
                    }
                    _ = fi14
                }
            }
        }
    }
    v22 := GenS3{F0: 342, F1: 214, F2: 337}
    _ = p0
    _ = v22
    return "hello"
}
func genFn1(p0 int, p1 uint64) uint {
    for fi23 := 0; fi23 < 3; fi23++ {
        p1 += p1
        fmt.Println(p0)
        v24 := []int{}
        _ = fi23
        _ = v24
    }
    p1 = p1
    var v25 uint16 = 364
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 float64 = (float64((((1.5 + ((float32(0.1) + 0.1) + (1e10 - 0.1))) + 1.5) - 0.1)) - 3.14159)
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        p0 = fi28
        _ = fi28
    }
    _ = p0
    _ = p1
    _ = v25
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 865
    v3 := []int{32767, 9223372036854775807, -176, 1}
    if ("" == "hello") {
        v2 = v2
    }
    v29 := GenS3{F0: 2147483648, F1: (v2 - (v2 * 3)), F2: 0}
    fmt.Println(v2)
    var v30 int64 = 1
    fmt.Println(v30)
    var v31 uint64 = 1
    if (1.5 >= 0.1) {
        v31 = (v31 + 18446744073709551615)
    } else {
        switch 9223372036854775807 {
        case 9223372036854775807:
            for fi32 := 0; fi32 < 14; fi32++ {
                fmt.Println(v30, v2, v30)
                _ = fi32
            }
        }
        for fi33 := 0; fi33 < 2; fi33++ {
            if (1000 <= -32768) {
                for fi34 := 0; fi34 < 17; fi34++ {
                    v31 = 18446744073709551615
                    v2 += v2
                    _ = fi34
                }
                v35 := GenS2{F0: 42, F1: 255, F2: 9223372036854775807}
                var v36 int16 = -219
                _ = v35
                _ = v36
            } else {
                v37 := GenS4{F0: v31, F1: (1 >= uint8(v31))}
                _ = v37
            }
            _ = fi33
        }
        _ = len(v3)
    }
    v30 -= v30
    var v38 float64 = float64(float32(0.1))
    _ = len(v1)
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
}
