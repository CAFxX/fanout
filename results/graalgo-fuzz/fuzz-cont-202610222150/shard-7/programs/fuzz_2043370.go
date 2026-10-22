package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS6 struct { F0 bool; F1 uint; F2 int16 }
func genFn0(p0 int16, p1 int8, p2 int16) string {
    v4 := GenS1{F0: -683, F1: 940}
    if ((0 > 4294967295) || false) {
        switch "foo bar" {
        case "":
            var v5 float64 = 3.14159
            _ = v5
        case "foo bar":
            v6 := GenS1{F0: 42, F1: 1}
            _ = v6
        }
    }
    v7 := []int{2147483647, 127, 127}
    if true {
        var v8 uint = 0
        _ = v8
    } else {
        if true {
            v9 := GenS1{F0: 0, F1: 1}
            v10 := GenS1{F0: 9223372036854775807, F1: 2147483647}
            _ = v10
            _ = v9
        }
        switch -160 {
        case 507:
            p2 = 0
        case 4294967295:
            v11 := GenS2{F0: p2}
            _ = v11
        case 2147483647:
            p1 = (-(p1))
        }
        var v12 int8 = -128
        _ = v12
    }
    switch 32767 {
    case 2147483647:
        if true {
            for fi13 := 0; fi13 < 17; fi13++ {
                if len(v7) > 0 {
                    v7[0] = (fi13 ^ -1)
                }
                for fi14 := 0; fi14 < 8; fi14++ {
                    v15 := GenS1{F0: int64(p2), F1: 1}
                    var v16 float64 = (1.5 + -2.25)
                    v16 -= v16
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi13
            }
        } else {
            var v17 int16 = (^(p0))
            for fi18 := 0; fi18 < 12; fi18++ {
                p2 = int16(p1)
                v17 = (p0 >> ((18446744073709551615) & 63))
                v19 := GenS3{F0: 1, F1: 32767}
                _ = fi18
                _ = v19
            }
            v7 = append(v7, 1)
            _ = v17
        }
    default:
        switch 726 {
        case 32767:
            var v20 int8 = (p1 ^ p1)
            _ = v20
        case 127:
            var v21 int64 = 0
            _ = v21
        case 857:
            for fi22 := 0; fi22 < 14; fi22++ {
                if ((false || (4294967295 == 1)) || false) {
                    var v23 int8 = p1
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        if len(v7) > 0 {
                            v7[0] = -1
                        }
                        _ = fi24
                    }
                    _ = len(v7)
                }
                v25 := GenS2{F0: (p0 * p0)}
                if ((fi22 >= -2147483648) || false) {
                    v26 := GenS2{F0: (p0 ^ -852)}
                    v27 := GenS4{F0: (fi22 - (fi22 << ((314) & 63))), F1: 32767}
                    v28 := GenS4{F0: fi22, F1: 0}
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 int16 = (p2 + -784)
                        var v31 uint64 = 1
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                }
                _ = fi22
                _ = v25
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return ("foo bar" + "a")
}
func genFn1() string {
    var v32 uint16 = 505
    v33 := []int{}
    _ = v32
    _ = v33
    return "hello"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 uint8 = 2
    v34 := GenS5{F0: 65535, F1: (!((true || true) || true)), F2: 42}
    var v35 int32 = 86
    var v36 uint8 = (v3 - v3)
    v37 := []int{32767, 32767, -70, -9223372036854775808}
    fmt.Println(v35, v3, v36)
    fmt.Println(v35)
    fmt.Println(v3, v3, v3)
    _ = len(v37)
    for fi38 := 0; fi38 < 18; fi38++ {
        v3 = v3
        _ = fi38
    }
    var v39 uint16 = 1
    v40 := GenS6{F0: (!true), F1: 635, F2: 877}
    v35 -= (v35 ^ -1)
    for fi41 := 0; fi41 < 10; fi41++ {
        delete(v1, "\n\t")
        v39 = v39
        _ = fi41
    }
    var v42 int16 = 267
    v2[-87] = uint32(v39)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v39:", v39)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v40
    _ = v42
}
