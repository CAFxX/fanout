package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 bool }
func genFn0(p0 bool, p1 int16) uint {
    v4 := GenS1{F0: -899}
    p1 -= p1
    _ = p0
    _ = p1
    _ = v4
    return 161
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{(("a" + "a") + "hello"), "hello"}
    var v3 int64 = 9223372036854775807
    if (false || true) {
        delete(v1, "foo bar")
        v3 = (-(v3))
        switch 0 {
        case -2147483648:
            fmt.Println(v3)
        case -1:
            if (1 <= 1) {
                var v5 uint32 = 42
                fmt.Println(v5, v5)
                _ = v5
            } else {
                var v6 int32 = 2147483647
                _ = v6
            }
        }
    } else {
        switch "a" {
        case "a":
            v3 -= v3
        case "foo bar":
            var v7 uint16 = 65535
            _ = v7
        case "":
            v8 := GenS1{F0: (v3 + 42)}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        _ = len(v2)
        _ = fi9
    }
    v10 := GenS2{F0: false, F1: (1 != 1)}
    var v11 float32 = 0.0
    var v12 int64 = (v3 - (v3 + v3))
    v12 -= 42
    v3 -= -678
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
