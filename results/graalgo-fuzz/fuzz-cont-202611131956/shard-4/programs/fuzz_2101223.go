package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 string }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{182}
    v3 := make(map[string]bool)
    var v4 int16 = 1
    var v5 float64 = (-2.25 - 3.14159)
    delete(v3, "foo bar")
    v4 = v4
    switch -603 {
    case -111:
        if (42 < -954) {
            switch 1 {
            case 773:
                v4 = v4
            }
        } else {
            for fi6 := 0; fi6 < 7; fi6++ {
                v4 -= v4
                _ = fi6
            }
            var v7 float64 = 0.1
            for fi8 := 0; fi8 < 18; fi8++ {
                delete(v3, "foo bar")
                v9 := GenS1{F0: "hello", F1: 1, F2: ""}
                _ = fi8
                _ = v9
            }
            _ = v7
        }
    case -949:
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 int8 = -69
            fmt.Println(v11, v5)
            var v12 int = fi10
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v13 := GenS1{F0: "a", F1: -1, F2: ("foo bar" + "foo bar")}
    if (false || true) {
        v4 += (v4 & v4)
    } else {
        v4 += (v4 * (^(v4)))
    }
    _ = len(v2)
    switch 42 {
    case -9223372036854775808:
        var v14 uint32 = 661
        _ = v14
    case 128:
        var v15 uint8 = 127
        _ = v15
    }
    v4 -= v4
    v16 := GenS1{F0: "", F1: 0, F2: ("a" + "\n\t")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
