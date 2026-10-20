package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = -2147483648
    v3 := []string{("" + (((("hello" + ("a" + "a")) + "foo bar") + "\n\t") + "foo bar")), "a"}
    v4 := []string{"", ("\n\t" + "\n\t"), "\n\t", ""}
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]int)
    v2 += v2
    v2 -= v2
    if false {
        v7 := GenS1{F0: 1}
        var v8 int16 = -1
        v9 := v6[1]
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint8 = 2
    v11 := GenS1{F0: 0}
    v3 = append(v3, "foo bar")
    for fi12 := 0; fi12 < 15; fi12++ {
        switch "a" {
        case "":
            v2 += (v2 + v2)
        }
        _ = fi12
    }
    var v13 int32 = -271
    if (!(255 >= 0)) {
        if (!false) {
            if true {
                var v14 float64 = ((1e10 - (float64(0.1) / -2.25)) * 0.0)
                _ = v14
            }
            v15 := GenS2{F0: v13}
            _ = v15
        }
        fmt.Println(v13)
        var v16 int = (v2 & v2)
        _ = v16
    } else {
        v13 += (v13 ^ (v13 / ((-1) | 1)))
        switch "" {
        case "":
            v2 = (v2 & (v2 % ((v2) | 1)))
        }
    }
    var v17 uint8 = (v10 + 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
