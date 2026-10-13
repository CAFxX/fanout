package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }
type GenS3 struct { F0 int8; F1 uint16; F2 int }
type GenS4 struct { F0 uint16; F1 int16; F2 uint16 }
func genFn0(p0 int64, p1 string, p2 uint32) uint32 {
    fmt.Println(p1, p2, p1)
    switch p1 {
    case "a":
        v6 := GenS1{F0: 255, F1: "\n\t", F2: 748}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 / ((2) | 1))
}
func genFn1() string {
    var v7 int32 = 0
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := []int{fi8, fi8, (fi8 * 42)}
        var v10 int32 = v7
        v11 := GenS2{F0: 0, F1: 2, F2: "a"}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := []int{856, 32768, 128, 4294967295}
    if true {
        if (true && (255 == 604)) {
            v12 = append(v12, 32768)
            switch 2147483647 {
            case 716:
                for fi13 := 0; fi13 < 15; fi13++ {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        var v15 uint64 = 4294967295
                        v15 = (^(v15))
                        _ = fi14
                        _ = v15
                    }
                    fmt.Println(fi13, fi13)
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v7 = (-(v7))
                        switch fi16 {
                        case 2147483647:
                            switch "\n\t" {
                            case "\n\t":
                                v7 = -988
                            case "a":
                                v17 := GenS1{F0: 1, F1: "\n\t", F2: 2}
                                _ = v17
                            case "hello":
                                v12 = append(v12, (-(fi13)))
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                if true {
                                    _ = len(v12)
                                } else {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                }
                            default:
                                switch ("hello" + "foo bar") {
                                case "":
                                    var v18 float64 = (1e10 * 0.1)
                                    _ = v18
                                case "hello":
                                    fmt.Println(v7)
                                case "a":
                                    var v19 int8 = -57
                                    _ = v19
                                default:
                                    var v20 uint32 = 2147483647
                                    _ = v20
                                }
                            }
                        default:
                            v7 = (v7 << ((1) & 63))
                        }
                        var v21 float32 = 3.14159
                        if len(v12) > 0 {
                            v12[0] = (fi13 % ((fi13) | 1))
                        }
                        if ((!true) || false) {
                            fmt.Println(v7, v21, fi16)
                            fmt.Println(v7, v7, v7)
                            var v22 uint = 255
                            _ = v22
                        }
                        _ = fi16
                        _ = v21
                    }
                    _ = fi13
                }
            default:
                v23 := GenS3{F0: 42, F1: 91, F2: 1}
                _ = v23
            }
            v7 -= -950
        } else {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
    }
    fmt.Println(v7)
    _ = v12
    _ = v7
    return "\n\t"
}

func main() {
    v1 := []int{32768, 9223372036854775807}
    var v2 int16 = 0
    var v3 uint32 = 65535
    var v4 uint = 42
    var v5 uint32 = (^(v3))
    v3 += uint32(v4)
    if (!(!("hello" < "hello"))) {
        if ((v4 & (v4 ^ (v4 >> (((v4 << ((uint(v3)) & 63))) & 63)))) < (^(v4))) {
            v5 -= v3
            v2 += 401
            v24 := GenS2{F0: 18446744073709551615, F1: 65535, F2: "\n\t"}
            _ = v24
        } else {
            if (!false) {
                fmt.Println(v2, v2, v4)
            } else {
                var v25 bool = true
                var v26 int8 = 100
                _ = v25
                _ = v26
            }
            v4 = v4
        }
        fmt.Println(v5, v4)
    } else {
        v1 = append(v1, 128)
    }
    v1 = append(v1, 32767)
    var v27 int = 879
    v4 = v4
    v27 = v27
    switch "" {
    case "foo bar":
        v28 := GenS2{F0: 1, F1: 1, F2: "\n\t"}
        _ = v28
    case "hello":
        v29 := GenS4{F0: 1, F1: (v2 << ((42) & 63)), F2: 0}
        _ = v29
    case "a":
        _ = len(v1)
    default:
        var v30 int16 = (^(v2))
        _ = v30
    }
    switch (^(v27)) {
    case 400:
        v1 = append(v1, v27)
    case 32768:
        for fi31 := 0; fi31 < 1; fi31++ {
            for fi32 := 0; fi32 < 9; fi32++ {
                _ = len(v1)
                _ = fi32
            }
            _ = fi31
        }
    }
    var v33 uint8 = 2
    switch "foo bar" {
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            var v34 uint = (v4 << ((142) & 63))
            _ = v34
        case "a":
            for fi35 := 0; fi35 < 9; fi35++ {
                _ = len(v1)
                v27 = v27
                _ = fi35
            }
        case "hello":
            v5 -= (v3 * 764)
        }
    case "a":
        var v36 int8 = 0
        _ = v36
    case "foo bar":
        v1 = append(v1, (-(v27)))
    default:
        v33 = (v33 * uint8(v4))
    }
    v4 -= 18446744073709551615
    v2 += (v2 * v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
