package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 string }

func main() {
    var v1 uint16 = 32767
    v2 := []string{("a" + "hello"), "a", ""}
    var v3 int = -625
    v4 := make(map[int]int)
    var v5 float32 = float32(float64(float32(1e10)))
    v6 := []int{}
    v4[(v3 - v3)] = (^(v3))
    v5 = float32(-2.25)
    var v7 float32 = v5
    v8 := GenS1{F0: 1, F1: 2}
    if true {
        v1 = v1
    } else {
        var v9 int8 = -1
        if true {
            var v10 uint64 = 1
            v6 = append(v6, (v3 + -545))
            _ = v10
        } else {
            fmt.Println(v7, v7)
        }
        _ = v9
    }
    switch "foo bar" {
    case "a":
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v5 = ((1.5 * 3.14159) - v5)
            v3 += v3
            v1 -= (v1 ^ v1)
        }
    default:
        fmt.Println(v5, v7)
    }
    var v11 bool = false
    fmt.Println(v1, v7)
    for fi12 := 0; fi12 < 11; fi12++ {
        if v11 {
            v1 += 2
            var v13 bool = (true || v11)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v13
        }
        _ = fi12
    }
    var v14 int32 = -109
    v15 := GenS2{F0: "hello"}
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint = 18446744073709551615
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: "\n\t"}
    if len(v2) > 0 {
        v2[0] = (("\n\t" + "a") + "")
    }
    v19 := v4[9223372036854775807]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
