package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint }
type GenS2 struct { F0 int; F1 bool; F2 int32 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int32; F1 bool }

func main() {
    var v1 int8 = 0
    var v2 string = "a"
    v3 := []uint32{115}
    var v4 uint64 = 18446744073709551615
    var v5 int32 = -2147483648
    if (true || true) {
        var v6 float32 = 3.14159
        v3 = append(v3, 4294967295)
        switch ((((v2 + "hello") + ("hello" + ((v2 + "a") + (("a" + ("foo bar" + (v2 + ("foo bar" + v2)))) + v2)))) + v2) + (v2 + "hello")) {
        case "\n\t":
            fmt.Println(v4, v4)
        case "hello":
            v1 = -83
        default:
            var v7 uint64 = (v4 ^ v4)
            _ = v7
        }
        _ = v6
    }
    fmt.Println(v5)
    var v8 uint8 = 127
    for fi9 := 0; fi9 < 14; fi9++ {
        switch int(v8) {
        case 4294967295:
            for fi10 := 0; fi10 < 17; fi10++ {
                if false {
                    var v11 int64 = -1
                    fmt.Println(v8, v11, v2)
                    v2 = v2
                    _ = v11
                } else {
                    fmt.Println(fi10, v1, v5)
                    _ = len(v3)
                }
                _ = fi10
            }
        case -9223372036854775808:
            var v12 float64 = 1.5
            _ = v12
        case -385:
            switch (fi9 + (fi9 ^ 127)) {
            case -1:
                switch fi9 {
                case 828:
                    v8 = (v8 ^ v8)
                case 1:
                    fmt.Println(v2, fi9, v1)
                default:
                    v3 = append(v3, 2147483647)
                }
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        default:
            v13 := GenS1{F0: 242, F1: ("a" + v2), F2: 963}
            _ = v13
        }
        v14 := GenS2{F0: 288, F1: (!((fi9 * (fi9 - 9223372036854775807)) == (fi9 * fi9))), F2: v5}
        _ = fi9
        _ = v14
    }
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    if true {
        var v15 bool = (2 <= 523)
        var v16 uint16 = 846
        _ = v15
        _ = v16
    }
    v1 = (v1 * 1)
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS3{F0: -9223372036854775808, F1: fi17}
        v5 -= v5
        v3 = append(v3, 0)
        _ = fi17
        _ = v18
    }
    v19 := GenS4{F0: -1, F1: (!false)}
    for fi20 := 0; fi20 < 13; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            v3 = append(v3, 2)
            v1 += v1
            _ = fi21
        }
        v1 = 100
        _ = fi20
    }
    v22 := GenS1{F0: 381, F1: "hello", F2: 2}
    if (!true) {
        var v23 int8 = (-(v1))
        _ = v23
    } else {
        v5 = v5
    }
    v5 = v5
    _ = len(v3)
    var v24 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
