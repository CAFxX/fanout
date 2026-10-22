package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    v8 := []int{-2147483648, -2147483648, 844, -999}
    for fi9 := 0; fi9 < 13; fi9++ {
        switch "hello" {
        case "a":
            switch "hello" {
            case "foo bar":
                if (true || (42 <= -128)) {
                    v8 = append(v8, fi9)
                    _ = len(v8)
                    var v10 int = (fi9 & (-(fi9)))
                    _ = v10
                } else {
                    var v11 int8 = -1
                    _ = v11
                }
            case "a":
                var v12 bool = false
                _ = v12
            default:
                _ = len(v8)
            }
        case "hello":
            _ = len(v8)
        default:
            v13 := GenS1{F0: 74, F1: fi9, F2: fi9}
            _ = v13
        }
        _ = fi9
    }
    _ = v8
    return 354
}

func main() {
    var v1 float64 = (3.14159 * (float64(1.5) - 3.14159))
    var v2 int64 = -393
    v3 := []int{}
    var v4 float32 = 1.5
    v5 := make(map[uint32]uint32)
    var v6 int32 = -1
    v14 := GenS1{F0: 841, F1: -9223372036854775808, F2: -468}
    var v15 uint16 = 499
    if (false && false) {
        if (((!false) || (true || false)) || (!true)) {
            fmt.Println(v2, v1, v2)
            var v16 int16 = int16(v6)
            _ = v16
        } else {
            var v17 uint32 = 65535
            _ = v17
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS1{F0: 0, F1: (-(fi18)), F2: (fi18 / ((fi18) | 1))}
            _ = fi18
            _ = v19
        }
    } else {
        var v20 uint64 = 1
        if true {
            v20 += v20
            switch 2147483647 {
            case -2147483648:
                v20 += 42
            default:
                v15 = 42
            }
            switch "" {
            case "":
                var v21 uint = 376
                _ = v21
            case "\n\t":
                delete(v5, 2147483648)
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        v2 = (-(v2))
        _ = v20
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        v23 := GenS2{F0: false}
        switch (fi22 * (fi22 + (fi22 ^ fi22))) {
        case 498:
            if (55 > 127) {
                fmt.Println(v4)
                v24 := GenS2{F0: true}
                v3 = append(v3, 128)
                _ = v24
            } else {
                for fi25 := 0; fi25 < 6; fi25++ {
                    switch fi22 {
                    case 2147483647:
                        v15 -= v15
                    default:
                        var v26 uint = 65535
                        _ = v26
                    }
                    var v27 string = "foo bar"
                    _ = fi25
                    _ = v27
                }
                _ = len(v3)
            }
        case 1:
            v28 := GenS3{F0: -535}
            _ = v28
        default:
            var v29 int16 = -462
            _ = v29
        }
        _ = fi22
        _ = v23
    }
    v30 := GenS2{F0: true}
    var v31 int8 = 0
    var v32 uint64 = 187
    fmt.Println(v32, v6, v2)
    for fi33 := 0; fi33 < 4; fi33++ {
        v6 -= v6
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
