package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int8; F2 int32 }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0(p0 int) int64 {
    fmt.Println(p0)
    v4 := GenS1{F0: 2147483647}
    for fi5 := 0; fi5 < 6; fi5++ {
        p0 = (p0 + p0)
        p0 += 4294967295
        _ = fi5
    }
    switch "foo bar" {
    case "hello":
        switch p0 {
        case -9223372036854775808:
            if (255 <= 42) {
                var v6 int16 = 777
                fmt.Println(p0, v6)
                v6 -= (^(v6))
                _ = v6
            } else {
                if (!false) {
                    p0 = (-(p0))
                    v7 := []int{p0}
                    _ = v7
                }
            }
        case -2147483648:
            v8 := GenS2{F0: 9223372036854775807, F1: 127, F2: 2147483647}
            _ = v8
        case 1:
            var v9 int = 127
            _ = v9
        }
    case "":
        v10 := []int{p0, p0, (p0 - (p0 & 32768))}
        _ = v10
    case "foo bar":
        fmt.Println(p0)
    default:
        p0 = p0
    }
    if (false && false) {
        p0 += p0
        var v11 bool = false
        v11 = (p0 != 97)
        _ = v11
    } else {
        p0 += p0
    }
    _ = p0
    _ = v4
    return 437
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 42
    v3 := make(map[string]uint32)
    v2 -= (v2 >> ((65535) & 63))
    delete(v1, 32768)
    v2 -= v2
    switch 0 {
    case 128:
        v2 = 42
    case -2147483648:
        fmt.Println(v2)
    }
    v12 := GenS3{F0: "\n\t", F1: 2147483648}
    v1[127] = true
    v2 = genFn0(127)
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
