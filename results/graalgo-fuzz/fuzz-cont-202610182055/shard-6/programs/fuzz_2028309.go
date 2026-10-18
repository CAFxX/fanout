package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 string; F2 uint }

func main() {
    v1 := []string{}
    v2 := []int{490, 566, 21, 32768}
    v3 := []int{32768}
    v4 := GenS1{F0: 540}
    var v5 float32 = (3.14159 - 0.1)
    v5 += -2.25
    v5 = v5
    _ = len(v2)
    switch "\n\t" {
    case "":
        v5 = v5
    case "hello":
        if true {
            v5 = 1e10
        }
    case "a":
        var v6 uint64 = 1
        _ = v6
    }
    v7 := GenS2{F0: 1000}
    var v8 string = (("hello" + "a") + "foo bar")
    v8 = "\n\t"
    var v9 float64 = (3.14159 * 1e10)
    if ((42 == 33) || false) {
        if len(v3) > 0 {
            v3[0] = 576
        }
        if (!true) {
            if true {
                var v10 float32 = (v5 / 3.14159)
                _ = v10
            }
            switch v8 {
            case "":
                if false {
                    v11 := GenS3{F0: v8}
                    _ = v11
                }
            case "\n\t":
                var v12 float32 = v5
                _ = v12
            default:
                if (9223372036854775807 < -183) {
                    var v13 uint64 = 418
                    switch "" {
                    case "hello":
                        fmt.Println(v5, v13)
                    }
                    v5 -= ((float32(v9) / (float32(float64(float32(v9))) - 0.0)) - v5)
                    _ = v13
                } else {
                    v14 := GenS4{F0: ((v8 + (("" + "foo bar") + ((v8 + v8) + (v8 + v8)))) + v8), F1: "hello", F2: 1}
                    fmt.Println(v8)
                    _ = v14
                }
            }
        }
    }
    var v15 uint8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
