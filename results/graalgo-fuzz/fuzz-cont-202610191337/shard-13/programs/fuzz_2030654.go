package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int64; F2 string }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 0
    var v3 uint8 = 128
    var v4 float64 = (float64(-2.25) - (1.5 + (1.5 / 3.14159)))
    v4 += float64((3.14159 / -2.25))
    v3 = 127
    v4 = v4
    var v5 uint32 = 4294967295
    v4 += v4
    fmt.Println(v5)
    v6 := []int{0}
    var v7 float32 = float32(-2.25)
    v8 := GenS1{F0: -1}
    for fi9 := 0; fi9 < 19; fi9++ {
        v6 = append(v6, (fi9 * 32767))
        if (-9223372036854775808 >= (fi9 % (((^(fi9))) | 1))) {
            var v10 uint32 = (v5 >> ((2) & 63))
            _ = v10
        }
        v11 := GenS2{F0: v2, F1: -1, F2: ("hello" + "hello")}
        _ = fi9
        _ = v11
    }
    v5 += (v5 >> ((42) & 63))
    _ = len(v1)
    for fi12 := 0; fi12 < 4; fi12++ {
        switch "hello" {
        case "":
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 int = -986
                var v15 int16 = 42
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            v1[-1] = false
        }
        var v16 float32 = v7
        switch (^(fi12)) {
        case 32767:
            v17 := GenS3{F0: -9223372036854775808}
            _ = v17
        case 928:
            switch "a" {
            case "":
                var v18 uint8 = 128
                _ = v18
            default:
                v2 += v2
            }
        case 42:
            var v19 float64 = v4
            _ = v19
        default:
            fmt.Println(v4)
        }
        _ = fi12
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
