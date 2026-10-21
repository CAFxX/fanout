package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 string }

func main() {
    var v1 float64 = float64(3.14159)
    v2 := make(map[int]bool)
    v3 := []string{"\n\t", "foo bar"}
    v4 := []uint32{65535}
    var v5 int32 = 617
    fmt.Println(v1)
    var v6 string = ("hello" + "")
    for fi7 := 0; fi7 < 15; fi7++ {
        if true {
            v8 := GenS1{F0: 1, F1: -9223372036854775808, F2: (v6 + v6)}
            v5 = (v5 * v5)
            if (false && (!true)) {
                v3 = append(v3, (("foo bar" + "hello") + v6))
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 int32 = (v5 + (^(v5)))
                    _ = fi9
                    _ = v10
                }
                var v11 int16 = -927
                _ = v11
            }
            _ = v8
        }
        _ = fi7
    }
    fmt.Println(v6, v6)
    switch (v6 + (v6 + "a")) {
    case "\n\t":
        var v12 float32 = 0.0
        _ = v12
    }
    switch v6 {
    case "hello":
        v1 += ((((3.14159 - v1) + v1) * (v1 - 3.14159)) / (-2.25 - (3.14159 * v1)))
    case "":
        v13 := v2[1]
        _ = v13
    case "foo bar":
        v1 -= float64((0.1 - (1.5 / 3.14159)))
    default:
        switch -9223372036854775808 {
        case 32767:
            var v14 float32 = float32(1.5)
            _ = v14
        case 0:
            v15 := GenS1{F0: 4294967295, F1: 42, F2: ("\n\t" + "hello")}
            _ = v15
        case -645:
            var v16 float32 = float32(((v1 + ((v1 - v1) + float64(3.14159))) / v1))
            _ = v16
        default:
            v17 := GenS2{F0: -128}
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS1{F0: 901, F1: 0, F2: v6}
        _ = fi18
        _ = v19
    }
    _ = len(v2)
    var v20 int32 = v5
    v1 = float64(-2.25)
    fmt.Println(v20)
    v20 += -772
    var v21 uint8 = 42
    if (0 <= 979) {
        v6 = v6
        switch 9223372036854775807 {
        case 9223372036854775807:
            v22 := v2[128]
            _ = v22
        case 127:
            _ = len(v4)
        case 32768:
            v21 = (v21 << ((255) & 63))
        default:
            v23 := GenS3{F0: 816, F1: 849}
            _ = v23
        }
        var v24 bool = true
        _ = v24
    }
    v25 := GenS4{F0: (v6 + "\n\t")}
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 int16 = -252
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
