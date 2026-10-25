package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
func genFn0(p0 uint32, p1 int32, p2 bool) int16 {
    fmt.Println(p0, p1, p1)
    p2 = p2
    v4 := GenS1{F0: 1, F1: "foo bar"}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    v1 := []int{2147483647, 0, 0}
    v2 := make(map[int]bool)
    var v3 int16 = -32768
    var v5 string = "\n\t"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 := GenS1{F0: 42, F1: (v5 + "hello")}
    var v7 float64 = float64((3.14159 + 0.1))
    switch 42 {
    case 32767:
        v8 := v2[9223372036854775807]
        _ = v8
    case 0:
        for fi9 := 0; fi9 < 18; fi9++ {
            v7 += (1e10 - (float64(0.1) * (((v7 + (v7 - float64((0.1 - 0.1)))) * 3.14159) * v7)))
            _ = fi9
        }
    case 127:
        v7 = 1e10
    }
    _ = len(v1)
    switch v5 {
    case "foo bar":
        v5 = ((("a" + "foo bar") + (v5 + ("a" + ("foo bar" + "")))) + ((v5 + v5) + v5))
    case "a":
        if true {
            if len(v1) > 0 {
                v1[0] = -1
            }
        } else {
            v7 -= (3.14159 - (v7 + (((v7 * (v7 / v7)) - (float64(1e10) / (((((v7 + v7) - v7) - v7) * v7) / 1.0))) / float64((float32(((v7 + float64(3.14159)) * v7)) - (float32(float64((0.0 * float32(v7)))) - (0.0 * 3.14159)))))))
        }
    default:
        v3 += genFn0(1, -1, (false || (true || false)))
    }
    _ = len(v2)
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(v7, v3, fi10)
        var v11 int32 = 0
        _ = fi10
        _ = v11
    }
    v3 -= v3
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
