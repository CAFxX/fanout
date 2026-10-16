package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 string, p1 uint16, p2 int16) int16 {
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            if (1 <= -96) {
                p2 = p2
                p2 = p2
            }
            _ = fi6
        }
        v7 := []int{(fi5 ^ -1), fi5, (fi5 >> ((65535) & 63))}
        _ = fi5
        _ = v7
    }
    fmt.Println(p1, p0)
    v8 := GenS1{F0: 727, F1: 1, F2: 1}
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int8 = 127
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 235
}
func genFn1(p0 int16) int8 {
    var v11 int64 = 0
    var v12 bool = true
    _ = p0
    _ = v11
    _ = v12
    return -1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    var v3 int16 = 64
    v4 := make(map[string]int)
    for fi13 := 0; fi13 < 12; fi13++ {
        if (true || (true || (896 >= 4294967295))) {
            var v14 int16 = 42
            var v15 uint8 = 128
            var v16 float32 = 0.0
            _ = v14
            _ = v15
            _ = v16
        } else {
            v3 = -885
            var v17 int32 = 417
            var v18 float32 = float32((((1.5 - (-2.25 / 1.5)) + ((1.5 - 3.14159) - (1e10 - (0.1 / 3.14159)))) + 3.14159))
            _ = v17
            _ = v18
        }
        v3 += v3
        _ = fi13
    }
    switch "" {
    case "foo bar":
        v19 := []int{4294967295}
        _ = v19
    case "\n\t":
        switch -9223372036854775808 {
        case -992:
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 int64 = -439
                _ = fi20
                _ = v21
            }
        case 1:
            v3 -= (v3 / (((-(v3))) | 1))
        case 32767:
            v3 = -32768
        default:
            switch 127 {
            case -2147483648:
                var v22 string = ("foo bar" + ("hello" + "foo bar"))
                _ = v22
            }
        }
    case "hello":
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS1{F0: 828, F1: 127, F2: 0}
            _ = fi23
            _ = v24
        }
    }
    v25 := []int{2147483647}
    if len(v25) > 0 {
        v25[0] = 127
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        if true {
            v3 += 32767
        } else {
            fmt.Println(v3, fi26)
            fmt.Println(fi26)
            v27 := GenS1{F0: 58, F1: 221, F2: 0}
            _ = v27
        }
        _ = fi26
    }
    v3 += genFn0("\n\t", 42, v3)
    for fi28 := 0; fi28 < 1; fi28++ {
        if (v3 <= v3) {
            v29 := v2[""]
            var v30 uint64 = 18446744073709551615
            _ = v29
            _ = v30
        } else {
            v3 = (v3 / (((v3 - (v3 & (-(v3))))) | 1))
            v31 := GenS1{F0: 4294967295, F1: 2, F2: 0}
            var v32 int32 = -423
            _ = v31
            _ = v32
        }
        v3 = v3
        var v33 float32 = (((1.5 + 1.5) + float32((1.5 / 1.0))) * ((3.14159 - 3.14159) * (-2.25 + -2.25)))
        _ = fi28
        _ = v33
    }
    if (true && (!true)) {
        v34 := GenS2{F0: 635}
        v3 += (v3 ^ (v3 ^ (v3 % ((-577) | 1))))
        fmt.Println(v3)
        _ = v34
    }
    switch "a" {
    case "foo bar":
        var v35 int = -99
        _ = v35
    default:
        v36 := GenS3{F0: -276}
        _ = v36
    }
    var v37 uint8 = 42
    var v38 uint8 = (v37 * v37)
    v37 += v37
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
