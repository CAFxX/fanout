package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int; F1 int64; F2 uint64 }
type GenS6 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    v3 := []uint32{2147483648}
    v4 := make(map[int]int)
    v5 := []string{("\n\t" + "foo bar")}
    v3 = append(v3, 316)
    var v6 int64 = 42
    v6 -= (^(v6))
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int64 = v6
        v9 := GenS1{F0: -2147483648}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float64 = ((1.5 - (0.0 / (1e10 / 0.1))) + -2.25)
    var v11 int16 = 32767
    switch 2147483647 {
    case 32767:
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 float64 = v10
            _ = fi12
            _ = v13
        }
    case -95:
        v14 := GenS2{F0: 1, F1: 977}
        _ = v14
    case 128:
        fmt.Println(v11)
    default:
        v6 -= (v6 ^ v6)
    }
    if true {
        var v15 int = 42
        v16 := GenS3{F0: false}
        _ = v15
        _ = v16
    }
    v17 := GenS4{F0: 2, F1: 13}
    v5 = append(v5, (("" + "\n\t") + ("\n\t" + ("a" + "a"))))
    switch (("\n\t" + "") + "hello") {
    case "hello":
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS5{F0: (fi18 / ((fi18) | 1)), F1: (-(v6)), F2: 18446744073709551615}
            _ = fi18
            _ = v19
        }
    case "a":
        var v20 int16 = v11
        _ = v20
    case "\n\t":
        switch (("" + "hello") + "foo bar") {
        case "\n\t":
            _ = len(v1)
        case "":
            v10 = v10
        }
    default:
        v10 += ((float64(float32((v10 * 0.0))) * 3.14159) * 0.0)
    }
    var v21 string = "foo bar"
    v21 = ("foo bar" + (v21 + ((("" + ("hello" + v21)) + ("\n\t" + (v21 + "a"))) + v21)))
    delete(v4, -860)
    v22 := GenS6{F0: (v11 / (((v11 & (-(v11)))) | 1)), F1: ""}
    var v23 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
