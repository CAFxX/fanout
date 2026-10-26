package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []uint32{907}
    v2 := make(map[uint32]uint32)
    var v3 int64 = -1
    v4 := make(map[string]string)
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    v3 += 9223372036854775807
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            delete(v2, 1)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi8
        }
        for fi9 := 0; fi9 < 6; fi9++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi9
        }
        _ = fi7
    }
    v3 = v3
    for fi10 := 0; fi10 < 16; fi10++ {
        v3 = v3
        _ = fi10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v3 = v3
        if (0 < 2) {
            var v12 int64 = v3
            v12 += v3
            _ = v12
        }
        _ = fi11
    }
    var v13 int64 = 60
    var v14 float64 = float64(float32(-2.25))
    v3 = v3
    switch "\n\t" {
    case "":
        v3 -= (v13 | (-(v13)))
    case "foo bar":
        v13 = -1
    case "\n\t":
        v15 := GenS1{F0: 2}
        _ = v15
    }
    v13 = (v3 * 932)
    v13 = (^(v13))
    v3 -= v13
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
