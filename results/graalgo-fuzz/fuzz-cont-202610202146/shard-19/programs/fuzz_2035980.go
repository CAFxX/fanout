package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 string }
func genFn0(p0 int, p1 int32, p2 uint8) uint32 {
    fmt.Println(p2, p1, p2)
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float32 = (1.5 + 1e10)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 970
}

func main() {
    v1 := []string{}
    var v2 uint16 = 42
    v3 := make(map[string]uint32)
    var v4 int64 = 9223372036854775807
    v5 := make(map[int]int)
    v4 += v4
    var v8 uint8 = 42
    var v9 uint64 = 2
    v9 -= v9
    for fi10 := 0; fi10 < 4; fi10++ {
        if false {
            v2 = v2
            for fi11 := 0; fi11 < 15; fi11++ {
                if true {
                    var v12 uint16 = 32767
                    v4 = v4
                    _ = v12
                }
                var v13 float64 = (float64(float32(3.14159)) / 1.0)
                v8 -= (^(v8))
                _ = fi11
                _ = v13
            }
        } else {
            v14 := GenS1{F0: -494, F1: -535, F2: fi10}
            _ = v14
        }
        switch "foo bar" {
        case "":
            fmt.Println(v8)
        case "foo bar":
            v15 := v5[fi10]
            _ = v15
        default:
            switch "" {
            case "foo bar":
                v16 := v5[42]
                _ = v16
            case "":
                var v17 uint64 = (v9 ^ v9)
                _ = v17
            }
        }
        var v18 int = 127
        _ = fi10
        _ = v18
    }
    v2 -= 1
    v19 := GenS2{F0: (v9 | 0), F1: ((0 == int(v2)) && (-128 <= 127))}
    v8 -= uint8(v2)
    var v20 int64 = 42
    v21 := GenS3{F0: ("\n\t" + ("hello" + "foo bar"))}
    var v22 uint64 = (^(v9))
    v5[491] = 128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
