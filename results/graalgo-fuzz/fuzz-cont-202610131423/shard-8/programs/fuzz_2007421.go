package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int16; F1 int16; F2 uint8 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64) uint32 {
    v5 := []int{975, -538, -807, -929}
    v6 := GenS1{F0: 1}
    _ = p0
    _ = v5
    _ = v6
    return 2147483648
}
func genFn1() uint8 {
    switch "a" {
    case "\n\t":
        switch "a" {
        case "\n\t":
            v7 := GenS2{F0: 0, F1: (42 > 0)}
            _ = v7
        case "a":
            v8 := []int{}
            _ = v8
        case "hello":
            if ((!true) && true) {
                var v9 uint32 = 0
                _ = v9
            }
        }
    default:
        var v10 bool = (-903 <= 127)
        _ = v10
    }
    v11 := GenS1{F0: 1}
    _ = v11
    return 128
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    v3 := []int{32768, 531, 4294967295}
    v4 := []uint32{2147483648, 1}
    for fi12 := 0; fi12 < 15; fi12++ {
        if (4294967295 > 18446744073709551615) {
            var v13 uint16 = 2
            _ = v13
        } else {
            for fi14 := 0; fi14 < 9; fi14++ {
                if false {
                    var v15 uint64 = 1
                    var v16 uint = 18446744073709551615
                    _ = v15
                    _ = v16
                }
                _ = fi14
            }
        }
        v17 := GenS2{F0: -191, F1: (-986 == 0)}
        _ = fi12
        _ = v17
    }
    var v18 float32 = 1e10
    switch "a" {
    case "a":
        v19 := GenS3{F0: -1, F1: 42, F2: 1}
        _ = v19
    default:
        for fi20 := 0; fi20 < 17; fi20++ {
            if (0 < genFn1()) {
                v21 := GenS2{F0: 825, F1: (-2.25 == (((float64(0.1) - 1.5) * 0.0) * (0.1 + -2.25)))}
                switch fi20 {
                case 307:
                    fmt.Println(v18, fi20)
                default:
                    if (0.1 != 1.5) {
                        var v22 string = ""
                        v18 += v18
                        fmt.Println(v22)
                        _ = v22
                    }
                }
                _ = v21
            }
            v18 = v18
            fmt.Println(v18, v18)
            _ = fi20
        }
    }
    if false {
        if len(v4) > 0 {
            v4[0] = 1
        }
    } else {
        var v23 string = ("\n\t" + ("" + "foo bar"))
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        fmt.Println(fi24, fi24)
        v2[""] = (fi24 << ((18446744073709551615) & 63))
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 uint64 = 157
            fmt.Println(fi24, v18)
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    v27 := GenS3{F0: 1, F1: 32767, F2: 42}
    delete(v1, "\n\t")
    v4 = append(v4, 325)
    var v28 uint = 1
    if (true || (!false)) {
        for fi29 := 0; fi29 < 11; fi29++ {
            v28 += v28
            var v30 uint64 = 599
            _ = fi29
            _ = v30
        }
        v18 -= float32(3.14159)
    }
    var v31 float64 = (-2.25 - 1e10)
    fmt.Println(v31, v18, v31)
    if false {
        v18 += 1.5
        var v32 int32 = -1
        fmt.Println(v28)
        _ = v32
    } else {
        switch "a" {
        case "foo bar":
            v33 := GenS2{F0: 1, F1: (255 != 42)}
            _ = v33
        case "\n\t":
            var v34 int16 = -258
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        if false {
            fmt.Println(v18, v31, v28)
            v18 -= float32(v31)
            var v36 int64 = -5
            _ = v36
        }
        _ = fi35
    }
    v37 := GenS4{F0: ((((v31 / ((v31 / v31) - (-2.25 * v31))) / (0.1 * v31)) / (v31 - float64(v18))) <= float64(v18))}
    for fi38 := 0; fi38 < 15; fi38++ {
        for fi39 := 0; fi39 < 8; fi39++ {
            var v40 int64 = 1
            _ = len(v3)
            if ((v28 - v28) > (v28 << (((v28 * v28)) & 63))) {
                fmt.Println(fi38, fi39, fi38)
                v41 := GenS1{F0: -2147483648}
                _ = v41
            } else {
                if len(v3) > 0 {
                    v3[0] = (fi39 % ((fi39) | 1))
                }
            }
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    v28 = v28
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v37
    _ = v4
}
