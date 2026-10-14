package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v4 uint32 = 2147483648
    v5 := []string{("\n\t" + "a"), ("" + "hello")}
    var v6 float64 = ((1.5 / 0.1) + (float64(((((1.5 - 1e10) + 0.1) + float32(3.14159)) / (0.0 + float32(3.14159)))) / (((0.0 * 1e10) - -2.25) - ((1.5 * 3.14159) / ((0.1 - ((float64(3.14159) * 3.14159) / 1.0)) - 1e10)))))
    switch -9223372036854775808 {
    case -272:
        v4 = (v4 % ((2147483648) | 1))
    case -9223372036854775808:
        if len(v5) > 0 {
            v5[0] = "a"
        }
    default:
        v6 -= v6
    }
    v4 += (v4 + 660)
    var v7 uint = 18446744073709551615
    if len(v5) > 0 {
        v5[0] = (("foo bar" + "\n\t") + "\n\t")
    }
    v5 = append(v5, "foo bar")
    _ = len(v3)
    v3[886] = "foo bar"
    v6 += -2.25
    v4 += (v4 - v4)
    v4 -= uint32(v7)
    switch 0 {
    case 0:
        var v8 int = 32768
        _ = v8
    case 9223372036854775807:
        fmt.Println(v4)
    default:
        v6 = 3.14159
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        switch fi9 {
        case -765:
            v6 -= float64(((-2.25 * -2.25) + (float32(v6) + float32((v6 - (v6 * 1e10))))))
        case -1:
            for fi10 := 0; fi10 < 19; fi10++ {
                fmt.Println(v7, v7)
                v6 -= 1.5
                _ = fi10
            }
        default:
            v4 += (v4 % ((v4) | 1))
        }
        _ = fi9
    }
    var v11 string = "foo bar"
    for fi12 := 0; fi12 < 7; fi12++ {
        switch v11 {
        case "hello":
            var v13 int = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            if len(v5) > 0 {
                v5[0] = ("hello" + (v11 + v11))
            }
            v15 := GenS1{F0: 42, F1: 1}
            _ = fi14
            _ = v15
        }
        switch "" {
        case "":
            delete(v1, v11)
        }
        _ = fi12
    }
    var v16 uint64 = 233
    if true {
        v3[-472] = (v11 + v11)
        for fi17 := 0; fi17 < 8; fi17++ {
            v7 += uint(v4)
            v6 = (v6 / float64((0.1 * 1e10)))
            _ = fi17
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
