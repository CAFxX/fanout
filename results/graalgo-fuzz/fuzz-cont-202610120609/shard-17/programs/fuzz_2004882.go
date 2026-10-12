package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 int64 = 1
    fmt.Println(v3)
    v3 -= (-(v3))
    v4 := GenS1{F0: 2147483647}
    v3 = (v3 - (v3 >> ((1) & 63)))
    switch "a" {
    case "":
        switch ((("\n\t" + "foo bar") + "") + "") {
        case "hello":
            if (((("a" + ("foo bar" + "a")) < "\n\t") || true) || true) {
                var v5 int32 = 2147483647
                var v6 float32 = ((-2.25 * 3.14159) + ((1.5 * 1e10) / 3.14159))
                _ = v5
                _ = v6
            }
        case "":
            var v7 int16 = 42
            _ = v7
        default:
            switch -2147483648 {
            case 508:
                v3 = (v3 ^ v3)
            case 0:
                var v8 bool = true
                _ = v8
            case -1:
                fmt.Println(v3)
            default:
                v3 += 0
            }
        }
    case "\n\t":
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                v10 := GenS1{F0: 410}
                _ = v10
            }
            _ = fi9
        }
    default:
        switch "a" {
        case "hello":
            v11 := GenS1{F0: -770}
            _ = v11
        default:
            v12 := []int{32767}
            _ = v12
        }
    }
    v13 := GenS1{F0: 2147483647}
    v3 += (v3 ^ 1)
    _ = len(v1)
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
