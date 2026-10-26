package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int8; F1 uint; F2 int }
type GenS3 struct { F0 int8; F1 int16; F2 uint64 }
type GenS4 struct { F0 string; F1 bool; F2 uint32 }
func genFn0(p0 int16, p1 int16) bool {
    fmt.Println(p0, p1)
    var v7 uint32 = 682
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 int64 = 9223372036854775807
            fmt.Println(fi8, fi8, fi8)
            v11 := GenS1{F0: 127, F1: 1}
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := []int{32768, 979}
        _ = len(v12)
        _ = fi8
        _ = v12
    }
    switch 127 {
    case 403:
        v13 := GenS2{F0: -122, F1: 0, F2: -121}
        _ = v13
    case -1:
        switch ("\n\t" + "\n\t") {
        case "":
            for fi14 := 0; fi14 < 17; fi14++ {
                if false {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        fmt.Println(v7, v7)
                        fmt.Println(fi15, fi15)
                        v16 := GenS1{F0: -128, F1: int32(v7)}
                        _ = fi15
                        _ = v16
                    }
                    if false {
                        v17 := []int{(^(fi14)), 2147483647}
                        _ = v17
                    }
                    v18 := []int{(fi14 / ((-1) | 1)), 4294967295}
                    _ = v18
                }
                for fi19 := 0; fi19 < 12; fi19++ {
                    v20 := []int{9223372036854775807}
                    var v21 uint16 = 1
                    v7 -= 642
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = fi14
            }
        case "a":
            var v22 bool = (195 == 0)
            _ = v22
        case "\n\t":
            var v23 uint64 = 4294967295
            _ = v23
        default:
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 uint32 = (v7 & (v7 << ((255) & 63)))
                v26 := GenS3{F0: 100, F1: p0, F2: 1}
                v27 := []int{40, (fi24 & fi24), fi24, (fi24 >> ((812) & 63))}
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return ((1 != 127) && false)
}
func genFn1(p0 uint16, p1 int16) uint64 {
    for fi28 := 0; fi28 < 9; fi28++ {
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 int8 = 11
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v31 := []int{699, 163, 32767, 574}
    _ = p0
    _ = p1
    _ = v31
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = -53
    var v3 string = "\n\t"
    var v4 bool = (false && (713 != 2))
    v5 := []string{"a", v3, "a", v3}
    var v6 float32 = (1.5 * 0.0)
    var v32 bool = (!(42 == 539))
    v33 := GenS3{F0: -128, F1: -1, F2: 540}
    v32 = v32
    fmt.Println(v4, v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v34 uint32 = 451
    v32 = (530 != 245)
    switch v3 {
    case "hello":
        v34 = (v34 * (v34 << ((42) & 63)))
    case "":
        if (v32 || (!v4)) {
            if v4 {
                var v35 int32 = -680
                _ = len(v1)
                _ = v35
            } else {
                fmt.Println(v34, v4, v34)
                v36 := GenS2{F0: v2, F1: 18446744073709551615, F2: 127}
                _ = v36
            }
            var v37 int = 1
            _ = v37
        } else {
            switch v3 {
            case "a":
                if v4 {
                    v4 = v32
                    fmt.Println(v6)
                    var v38 uint8 = 210
                    _ = v38
                } else {
                    v39 := v1[("hello" + ((("a" + (v3 + ((v3 + "\n\t") + v3))) + v3) + v3))]
                    fmt.Println(v2)
                    _ = v39
                }
            }
            if ((-2.25 - v6) >= float32(float64(3.14159))) {
                var v40 bool = v32
                _ = v40
            } else {
                if v4 {
                    for fi41 := 0; fi41 < 18; fi41++ {
                        v6 -= v6
                        v6 -= v6
                        _ = fi41
                    }
                    v42 := GenS4{F0: "", F1: false, F2: (v34 + 1)}
                    var v43 uint16 = 1
                    _ = v42
                    _ = v43
                }
            }
        }
    default:
        for fi44 := 0; fi44 < 11; fi44++ {
            v45 := GenS4{F0: v3, F1: (false && v4), F2: (v34 | (^(v34)))}
            v3 = v3
            _ = fi44
            _ = v45
        }
    }
    var v46 int = 906
    fmt.Println(v2, v4)
    if true {
        v32 = (!(65535 == 255))
    }
    var v47 uint = 18446744073709551615
    _ = len(v5)
    v46 -= 9223372036854775807
    v32 = false
    v48 := GenS1{F0: (v2 ^ (v2 / ((100) | 1))), F1: 905}
    fmt.Println(v32, v2, v47)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
