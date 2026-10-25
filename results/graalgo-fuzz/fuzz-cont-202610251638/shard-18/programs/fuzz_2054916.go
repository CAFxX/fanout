package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 string }
type GenS2 struct { F0 int32; F1 string; F2 int8 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint16, p1 bool, p2 int32) int {
    p2 = -800
    var v4 uint = 2
    v4 += v4
    if p1 {
        v5 := []int{2147483647, 1}
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 uint16 = p0
            _ = fi6
            _ = v7
        }
        if p1 {
            for fi8 := 0; fi8 < 4; fi8++ {
                if len(v5) > 0 {
                    v5[0] = fi8
                }
                fmt.Println(p2)
                _ = fi8
            }
        } else {
            p1 = false
            p0 = (p0 - (p0 * 42))
            if len(v5) > 0 {
                v5[0] = 178
            }
        }
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2147483647
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 float32 = 1e10
    v3 := make(map[uint32]int)
    for fi9 := 0; fi9 < 15; fi9++ {
        v1 = v1
        v10 := GenS1{F0: "\n\t", F1: 1, F2: "foo bar"}
        _ = fi9
        _ = v10
    }
    _ = len(v3)
    v11 := []int{}
    v2 += float32((((((1.5 + 0.1) + (3.14159 + 3.14159)) - float64((3.14159 - 0.0))) + 1.5) + float64(1e10)))
    switch 32768 {
    case -834:
        var v12 bool = false
        _ = v12
    case 2147483647:
        v1 = v1
    case -796:
        v13 := v3[2]
        _ = v13
    }
    v14 := GenS1{F0: "", F1: 0, F2: ""}
    switch -9223372036854775808 {
    case -1:
        if false {
            _ = len(v11)
            for fi15 := 0; fi15 < 10; fi15++ {
                switch "hello" {
                case "foo bar":
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(fi15, v1)
                    case "\n\t":
                        var v16 float32 = ((v2 - (v2 * 0.1)) / v2)
                        _ = v16
                    }
                case "a":
                    if false {
                        for fi17 := 0; fi17 < 2; fi17++ {
                            fmt.Println(fi17)
                            _ = fi17
                        }
                    }
                default:
                    v11 = append(v11, genFn0(1, false, 154))
                }
                if (uint8(v1) <= 1) {
                    for fi18 := 0; fi18 < 6; fi18++ {
                        if false {
                            var v19 int64 = 9223372036854775807
                            var v20 int16 = -1
                            _ = v19
                            _ = v20
                        } else {
                            _ = len(v11)
                        }
                        v1 = v1
                        v21 := GenS2{F0: 2147483647, F1: "hello", F2: -122}
                        _ = fi18
                        _ = v21
                    }
                    v22 := GenS3{F0: 42, F1: 42}
                    _ = v22
                } else {
                    var v23 float32 = (v2 + (float32((((((0.1 + 3.14159) + float64(v2)) - float64(((float32((float64(v2) / 1e10)) * 0.0) - 3.14159))) - -2.25) - float64(0.1))) - v2))
                    _ = v23
                }
                _ = fi15
            }
        }
    default:
        switch genFn0(255, (42 < 42), 2147483647) {
        case 634:
            if (!((("foo bar" < "hello") && false) || (false && (!true)))) {
                v1 -= 65535
            }
        case 1:
            v11 = append(v11, 127)
        case 4294967295:
            v24 := GenS3{F0: 18446744073709551615, F1: 65}
            _ = v24
        default:
            v25 := GenS2{F0: 2147483647, F1: ("hello" + "hello"), F2: 100}
            _ = v25
        }
    }
    v2 = 3.14159
    v3[942] = genFn0(993, false, -1)
    switch "" {
    case "\n\t":
        var v26 int64 = 42
        _ = v26
    }
    v27 := GenS1{F0: "a", F1: genFn0(32767, false, 0), F2: "a"}
    if len(v11) > 0 {
        v11[0] = 9223372036854775807
    }
    var v28 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
