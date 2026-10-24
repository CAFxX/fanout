package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int16, p1 uint32, p2 string) uint64 {
    if (498 < uint16(p1)) {
        fmt.Println(p1, p0)
    } else {
        v5 := GenS1{F0: 42}
        p0 += p0
        _ = v5
    }
    v6 := GenS1{F0: -108}
    for fi7 := 0; fi7 < 3; fi7++ {
        p2 = ("hello" + p2)
        _ = fi7
    }
    switch 0 {
    case 1:
        p1 -= (p1 ^ 2147483648)
    case 4294967295:
        p1 = 0
    case -2147483648:
        v8 := []int{9223372036854775807, -1, 0, -953}
        _ = v8
    default:
        if (false || false) {
            switch p2 {
            case "":
                v9 := []int{2147483647}
                _ = v9
            case "hello":
                v10 := GenS1{F0: -90}
                _ = v10
            case "foo bar":
                for fi11 := 0; fi11 < 6; fi11++ {
                    p0 = p0
                    for fi12 := 0; fi12 < 14; fi12++ {
                        var v13 uint = 4294967295
                        _ = fi12
                        _ = v13
                    }
                    if true {
                        p1 = p1
                        v14 := []int{fi11, (fi11 % ((fi11) | 1)), fi11}
                        _ = v14
                    } else {
                        var v15 string = "a"
                        _ = v15
                    }
                    _ = fi11
                }
            }
        } else {
            v16 := []int{32767, -2147483648, 0}
            v17 := GenS1{F0: 100}
            if ("a" != p2) {
                if len(v16) > 0 {
                    v16[0] = 9223372036854775807
                }
                if true {
                    v18 := GenS1{F0: -128}
                    fmt.Println(p2, p2)
                    _ = v18
                } else {
                    var v19 uint8 = 255
                    p2 = p2
                    if len(v16) > 0 {
                        v16[0] = 906
                    }
                    _ = v19
                }
                switch 127 {
                case 0:
                    v20 := GenS2{F0: p0, F1: -78}
                    _ = v20
                case 32767:
                    if (((((!((p0 == -1) && false)) || true) || (0 == 2)) || false) && false) {
                        var v21 float32 = (0.0 / 1.5)
                        v16 = append(v16, -342)
                        _ = v21
                    }
                case -9223372036854775808:
                    v22 := GenS3{F0: 42}
                    _ = v22
                }
            } else {
                switch -633 {
                case 128:
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                default:
                    v16 = append(v16, 11)
                }
            }
            _ = v16
            _ = v17
        }
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 uint64 = 4294967295
    v23 := []int{-585}
    var v24 int32 = 1
    var v25 uint = 255
    var v26 uint8 = 128
    v27 := GenS2{F0: (v1 | v1), F1: 34}
    var v28 uint = v25
    for fi29 := 0; fi29 < 13; fi29++ {
        v25 = v28
        v4 = v4
        delete(v2, fi29)
        _ = fi29
    }
    switch -275 {
    case 1:
        v1 += v1
    case 339:
        if true {
            v30 := GenS1{F0: 127}
            _ = v30
        }
    case 2147483647:
        v31 := GenS1{F0: int8(v24)}
        _ = v31
    }
    v26 = 2
    fmt.Println(v26)
    v1 = (v1 - int16(v26))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
