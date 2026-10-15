package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
type GenS2 struct { F0 uint; F1 string; F2 int64 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int64 = 0
    var v3 int = 179
    v4 := make(map[int]uint32)
    if true {
        v1 = v2
    }
    v3 = v3
    v5 := GenS1{F0: 9223372036854775807, F1: true, F2: 42}
    fmt.Println(v2, v1)
    v6 := []int{(v3 / ((int(v2)) | 1))}
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v1, v3, v3)
        v2 -= 0
        for fi8 := 0; fi8 < 8; fi8++ {
            switch 1 {
            case 42:
                _ = len(v4)
            case 127:
                var v9 int64 = (v1 | 9223372036854775807)
                _ = v9
            case 264:
                v6 = append(v6, (^(fi7)))
            }
            _ = fi8
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 int8 = -54
        fmt.Println(v1, fi10)
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 255, F1: ("foo bar" + (("\n\t" + (("hello" + "hello") + "hello")) + "hello")), F2: v1}
    fmt.Println(v2)
    switch "" {
    case "a":
        _ = len(v4)
    case "hello":
        switch ("hello" + "foo bar") {
        case "\n\t":
            if true {
                if ((((0.1 / 0.1) / (float32(3.14159) * 1.5)) + 3.14159) > float32(float64((-2.25 * (0.0 - -2.25))))) {
                    v2 = (v1 + (-(v2)))
                    v6 = append(v6, v3)
                    switch v3 {
                    case 42:
                        v6 = append(v6, v3)
                    }
                } else {
                    v13 := GenS2{F0: 274, F1: "\n\t", F2: v1}
                    v14 := GenS2{F0: 42, F1: "a", F2: (^(v2))}
                    _ = v13
                    _ = v14
                }
                var v15 float32 = ((float32(0.1) + -2.25) / 0.1)
                _ = v15
            } else {
                var v16 uint64 = 4294967295
                _ = v16
            }
        }
    case "foo bar":
        var v17 float64 = -2.25
        _ = v17
    }
    v18 := GenS2{F0: 42, F1: (((("a" + "") + "a") + (("" + "hello") + "")) + ""), F2: (-(v2))}
    var v19 float32 = 0.1
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
