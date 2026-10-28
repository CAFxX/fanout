package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 bool; F1 string; F2 uint32 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int16 = -32768
    var v2 bool = (3.14159 == (0.1 / 0.1))
    v3 := make(map[string]int)
    v4 := GenS1{F0: "\n\t", F1: (302 <= 0)}
    v5 := []int{-135, -205, -9223372036854775808, 32768}
    if false {
        if (!(!true)) {
            if (42 <= -1) {
                v2 = true
                v6 := GenS2{F0: (!v2), F1: "hello", F2: 42}
                v1 -= 707
                _ = v6
            }
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 string = ""
                _ = fi7
                _ = v8
            }
        }
        v2 = false
        v9 := GenS3{F0: (v1 ^ v1)}
        _ = v9
    }
    var v10 uint8 = 1
    v2 = (!(!(9223372036854775807 < 519)))
    v11 := GenS1{F0: ("a" + "foo bar"), F1: ((((int16(v10) >= v1) || (!((v1 - 0) > v1))) && false) && (4294967295 == 0))}
    v5 = append(v5, 611)
    v10 = v10
    switch "hello" {
    case "a":
        for fi12 := 0; fi12 < 2; fi12++ {
            v10 = (v10 / ((42) | 1))
            _ = fi12
        }
    case "foo bar":
        v3[("\n\t" + ("" + "a"))] = 1
    case "hello":
        v1 = v1
    default:
        v13 := GenS3{F0: (v1 - v1)}
        _ = v13
    }
    if v2 {
        delete(v3, "")
        var v14 float64 = (0.0 - (0.0 - (0.0 / 1.5)))
        if false {
            fmt.Println(v14, v10, v14)
        }
        _ = v14
    }
    var v15 int8 = 0
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
