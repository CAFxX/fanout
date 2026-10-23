package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint8 }
type GenS2 struct { F0 bool; F1 int32; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint32; F2 int16 }
func genFn0(p0 int8, p1 uint16, p2 bool) string {
    p0 = -104
    if true {
        v6 := []int{2147483647, 128, -1, -2147483648}
        for fi7 := 0; fi7 < 19; fi7++ {
            p0 += 127
            var v8 int32 = -610
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 972, F1: -2147483648, F2: 56}
        _ = v6
        _ = v9
    } else {
        v10 := []int{}
        var v11 uint32 = 91
        p2 = p2
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        p0 = p0
        _ = fi12
    }
    v13 := GenS2{F0: (!(((true && p2) || p2) || (42 <= 2147483647))), F1: 0, F2: 2147483647}
    var v14 uint64 = 130
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{-1, -9223372036854775808}
    v3 := []uint32{778, 2147483648}
    v4 := []string{"foo bar", "", "", "\n\t"}
    v5 := []int{}
    var v15 uint64 = 42
    if false {
        v2 = append(v2, 127)
    } else {
        var v16 string = "\n\t"
        switch -625 {
        case 272:
            var v17 uint = 436
            _ = v17
        case 32767:
            if (0 > 627) {
                var v18 int8 = -115
                fmt.Println(v15, v18, v15)
                fmt.Println(v18, v18, v15)
                _ = v18
            } else {
                switch v16 {
                case "hello":
                    fmt.Println(v16)
                case "a":
                    for fi19 := 0; fi19 < 20; fi19++ {
                        v20 := v1[0]
                        _ = fi19
                        _ = v20
                    }
                default:
                    v1[0] = true
                }
                var v21 bool = true
                v15 = (v15 ^ (v15 / ((1) | 1)))
                _ = v21
            }
        case 881:
            _ = len(v4)
        }
        _ = v16
    }
    v22 := v1[1]
    v23 := GenS2{F0: (688 > 454), F1: 0, F2: 32767}
    for fi24 := 0; fi24 < 11; fi24++ {
        fmt.Println(v22, v22)
        for fi25 := 0; fi25 < 15; fi25++ {
            v22 = ((!(!(255 >= 127))) && v22)
            _ = fi25
        }
        for fi26 := 0; fi26 < 15; fi26++ {
            v1[2147483648] = (42 < 18446744073709551615)
            v27 := GenS2{F0: v22, F1: -1, F2: 1}
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            _ = fi26
            _ = v27
        }
        _ = fi24
    }
    v15 = (v15 * (v15 & v15))
    var v28 uint16 = 372
    var v29 bool = (0 == -9223372036854775808)
    v30 := GenS3{F0: 76}
    v31 := GenS3{F0: 128}
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 int64 = 42
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 float64 = -2.25
        var v36 uint32 = 310
        v22 = false
        _ = fi34
        _ = v35
        _ = v36
    }
    v37 := GenS4{F0: -9223372036854775808, F1: 343, F2: 32767}
    for fi38 := 0; fi38 < 2; fi38++ {
        var v39 int64 = 811
        _ = fi38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
    _ = v5
}
