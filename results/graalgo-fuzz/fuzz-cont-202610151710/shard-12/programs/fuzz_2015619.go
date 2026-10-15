package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int64; F2 bool }
func genFn0(p0 uint64, p1 string, p2 uint32) uint64 {
    v7 := GenS1{F0: p1}
    switch "a" {
    case "foo bar":
        if ((true || (255 > 1)) && (-2147483648 >= -1)) {
            var v8 int64 = 42
            _ = v8
        }
    case "\n\t":
        p2 += (p2 & 65535)
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS2{F0: -58}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 731
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []uint32{2147483647, 2147483647, 2147483647}
    v3 := []string{"\n\t", "\n\t", ""}
    var v4 uint8 = 255
    v5 := make(map[string]string)
    v6 := make(map[int]uint32)
    if false {
        v4 += 1
        var v11 int8 = 1
        _ = v11
    }
    v1 = (^(v1))
    if (true && (125 >= 2)) {
        var v12 uint8 = v4
        v12 = v12
        fmt.Println(v12)
        _ = v12
    } else {
        var v13 int16 = -1
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v1 -= (^(v1))
        _ = fi14
    }
    var v15 int8 = 17
    for fi16 := 0; fi16 < 10; fi16++ {
        v4 = 89
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 uint64 = 0
            _ = len(v6)
            if len(v2) > 0 {
                v2[0] = 1
            }
            var v19 int16 = -32768
            v3 = append(v3, "hello")
            var v20 int16 = v19
            if (fi17 < (fi17 % (((-(fi16))) | 1))) {
                v18 += (^(v18))
                v20 += (v19 * (-(v19)))
                var v21 float64 = ((3.14159 / -2.25) - ((1.5 / (float64((-2.25 + 1.5)) - 3.14159)) + 0.0))
                switch 2147483647 {
                case 2147483647:
                    v4 += 173
                case 32767:
                    fmt.Println(fi16, v20, v19)
                case 752:
                    delete(v5, "\n\t")
                }
                switch "\n\t" {
                case "foo bar":
                    v19 = v19
                default:
                    var v22 uint8 = 241
                    _ = v22
                }
                _ = v21
            } else {
                switch (("foo bar" + "a") + "a") {
                case "\n\t":
                    v6[(fi17 * (fi17 - (fi16 / ((fi16) | 1))))] = 2147483647
                case "hello":
                    var v23 uint64 = 130
                    _ = v23
                case "":
                    v2 = append(v2, 2147483648)
                default:
                    v15 += 42
                }
                v24 := GenS3{F0: (v15 >> ((2) & 63)), F1: 0, F2: 2}
                _ = v24
            }
            v5["\n\t"] = "a"
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi16
    }
    if (true && true) {
        for fi25 := 0; fi25 < 17; fi25++ {
            v15 -= int8(v1)
            v4 -= (v4 * 134)
            _ = fi25
        }
    } else {
        var v26 uint8 = (v4 * (v4 << ((2) & 63)))
        _ = v26
    }
    v1 = v1
    v27 := GenS4{F0: 770, F1: v1, F2: true}
    v15 = -1
    delete(v6, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
