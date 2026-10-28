package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint8 = 73
    var v2 int8 = 100
    v3 := []string{"a"}
    switch "hello" {
    case "\n\t":
        if true {
            v4 := GenS1{F0: -1, F1: 2, F2: 1}
            v2 = (v2 - (v2 + 0))
            _ = v4
        }
    case "foo bar":
        v2 = v2
    case "a":
        for fi5 := 0; fi5 < 7; fi5++ {
            fmt.Println(v1, v2, v2)
            if (true && (true || true)) {
                v1 -= v1
                v1 += v1
            }
            if (true || false) {
                v2 += -64
                var v6 uint8 = (v1 ^ v1)
                _ = v6
            } else {
                _ = len(v3)
            }
            _ = fi5
        }
    default:
        var v7 int8 = v2
        _ = v7
    }
    var v8 bool = false
    switch ((("\n\t" + "a") + "hello") + "hello") {
    case "":
        v9 := GenS2{F0: 722}
        _ = v9
    }
    v3 = append(v3, ("a" + ""))
    if true {
        var v10 float64 = float64((3.14159 / (1.5 * 0.1)))
        _ = v10
    }
    var v11 float32 = -2.25
    var v12 uint64 = 208
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v13 := GenS3{F0: (^(v1))}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, "foo bar")
    if (v8 || true) {
        v14 := GenS4{F0: 26}
        _ = v14
    }
    v3 = append(v3, "hello")
    fmt.Println(v11, v2, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
}
