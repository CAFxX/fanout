package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = float32(0.1)
    var v3 float64 = (0.0 / 1.5)
    v4 := []string{"foo bar", "hello"}
    var v5 int32 = -1
    v3 -= v3
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(fi6)
        var v7 uint16 = 0
        _ = fi6
        _ = v7
    }
    var v8 float64 = (v3 / (float64(3.14159) * (1.5 / v3)))
    var v9 float64 = float64(1e10)
    v8 = float64((((v2 * v2) + v2) / ((v2 / (float32(v3) * (float32(-2.25) * ((v2 - v2) - (((v2 / float32(v9)) + (1e10 * v2)) / (((v2 - v2) / v2) - 3.14159)))))) / ((float32(v9) + (v2 / v2)) + v2))))
    if ("" == "a") {
        v9 -= v3
    }
    if (false || true) {
        var v10 int16 = -974
        v1[("a" + (("" + "\n\t") + "\n\t"))] = 379
        _ = v10
    }
    switch "" {
    case "\n\t":
        v9 -= (v3 - float64(float32(1.5)))
    case "":
        if true {
            fmt.Println(v9)
        } else {
            var v11 int64 = -9223372036854775808
            _ = v11
        }
    }
    v5 = v5
    v12 := GenS1{F0: 747}
    v3 = float64(((float32(((v9 * (v9 - v3)) - v3)) - (v2 / 1e10)) - float32((0.0 * float64((v2 - float32(float64(((v2 * (v2 / v2)) / v2)))))))))
    if (127 <= -82) {
        if true {
            v8 -= 0.0
            var v13 int16 = 501
            _ = v13
        }
        v2 -= v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
