package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint; F1 uint64; F2 int32 }
type GenS3 struct { F0 string }

func main() {
    var v1 uint = 255
    var v2 int32 = 1
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    v5 := []int{32767}
    var v6 uint16 = 65535
    fmt.Println(v1)
    for fi7 := 0; fi7 < 9; fi7++ {
        v1 = v1
        _ = fi7
    }
    var v8 float32 = (((1e10 + (0.0 - float32(float64((0.1 * (-2.25 / 0.1)))))) * 1.5) - ((1e10 * (((-2.25 * (float32(((3.14159 + float64(1.5)) * (-2.25 - (1e10 * 1.5)))) - float32(1.5))) - 0.0) + ((1e10 * (1e10 / 0.1)) - (0.0 * float32(1e10))))) - 1.5))
    v8 = 3.14159
    v9 := GenS1{F0: -128, F1: ("a" + "")}
    for fi10 := 0; fi10 < 2; fi10++ {
        if (fi10 > (^(fi10))) {
            fmt.Println(fi10, v6, v6)
            _ = len(v5)
            if false {
                var v11 uint32 = 2147483648
                v12 := GenS2{F0: v1, F1: 2, F2: v2}
                _ = v11
                _ = v12
            }
        } else {
            var v13 uint16 = v6
            _ = len(v3)
            _ = v13
        }
        fmt.Println(v1, v2)
        var v14 int = 42
        _ = fi10
        _ = v14
    }
    var v15 uint8 = 128
    switch "a" {
    case "\n\t":
        var v16 uint16 = (v6 | (^(v6)))
        _ = v16
    case "a":
        v8 = v8
    case "foo bar":
        switch 42 {
        case 4294967295:
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 uint16 = 235
                var v19 int32 = (v2 ^ int32(v6))
                _ = fi17
                _ = v18
                _ = v19
            }
        default:
            v2 = (v2 + v2)
        }
    }
    v20 := GenS2{F0: v1, F1: 743, F2: 42}
    switch "foo bar" {
    case "\n\t":
        var v21 int8 = -13
        _ = v21
    case "a":
        fmt.Println(v2)
    case "hello":
        v15 += v15
    }
    if ((-1 > 1) && (true && (117 == (v15 - v15)))) {
        fmt.Println(v8, v1)
        switch 115 {
        case -1:
            switch "\n\t" {
            case "":
                v8 = v8
            case "a":
                switch ("a" + "a") {
                case "foo bar":
                    v22 := GenS3{F0: ""}
                    _ = v22
                case "hello":
                    v6 = v6
                case "a":
                    v1 = (v1 >> ((2) & 63))
                default:
                    if (true && true) {
                        if (v6 < 0) {
                            v3[-857] = ("hello" + "foo bar")
                            v6 += 42
                            v2 = (-(v2))
                        }
                    }
                }
            default:
                if false {
                    var v23 bool = (0.0 > (0.0 + 0.1))
                    var v24 uint32 = 420
                    _ = v23
                    _ = v24
                } else {
                    v2 = (v2 ^ v2)
                    var v25 uint8 = (^(v15))
                    _ = v25
                }
            }
        default:
            v8 += v8
        }
    }
    var v26 int32 = (v2 | v2)
    var v27 uint8 = v15
    v8 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
