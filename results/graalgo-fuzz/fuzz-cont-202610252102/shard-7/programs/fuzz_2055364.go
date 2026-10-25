package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint64; F1 string; F2 int32 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int64 }
func genFn0(p0 int8) int16 {
    var v6 uint8 = 0
    v7 := GenS1{F0: 0, F1: v6}
    _ = p0
    _ = v6
    _ = v7
    return -32768
}

func main() {
    var v1 int64 = 371
    v2 := []string{("foo bar" + "foo bar"), ("foo bar" + "hello")}
    var v3 float64 = float64(3.14159)
    var v4 string = ("foo bar" + "")
    var v5 string = "a"
    var v8 int = -341
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 32767, F1: 127}
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 32767, F1: (-2147483648 > -1)}
    v12 := GenS3{F0: 1, F1: "hello", F2: 2147483647}
    v4 = ((((v5 + "hello") + v4) + "a") + "\n\t")
    var v13 int64 = (v1 ^ 1)
    if (502 >= 2) {
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 int64 = (v13 ^ (v1 % ((v1) | 1)))
            if (true && true) {
                v16 := GenS1{F0: 392, F1: 127}
                _ = v16
            }
            _ = fi14
            _ = v15
        }
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int = fi17
            _ = fi17
            _ = v18
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v19 float64 = (float64((1.5 * -2.25)) - float64(0.0))
    switch "foo bar" {
    case "a":
        var v20 uint = 939
        _ = v20
    case "hello":
        var v21 int32 = 2147483647
        _ = v21
    case "":
        switch "" {
        case "a":
            fmt.Println(v4)
        case "\n\t":
            switch v4 {
            case "a":
                fmt.Println(v13, v3)
            default:
                var v22 int32 = -2147483648
                _ = v22
            }
        default:
            v23 := GenS4{F0: 42, F1: 2147483647, F2: (v13 | (v13 - v13))}
            _ = v23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
