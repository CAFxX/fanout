package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int32 {
    var v7 uint64 = 626
    v8 := []int{9223372036854775807, -543, 942, 32767}
    v7 -= v7
    _ = v7
    _ = v8
    return 445
}

func main() {
    var v1 float64 = (1.5 - (0.1 / (1.5 * ((3.14159 + ((1e10 * (0.0 + -2.25)) + 3.14159)) + (((((0.1 / -2.25) / -2.25) / (-2.25 - (3.14159 / 1.0))) / 1.0) / (float64((3.14159 + (1.5 * 3.14159))) + (0.1 - 1e10)))))))
    v2 := make(map[int]bool)
    var v3 uint32 = 2
    v4 := []uint32{}
    var v5 int64 = 63
    var v6 float64 = 1e10
    var v9 int8 = 42
    v10 := GenS1{F0: 42}
    switch "\n\t" {
    case "":
        for fi11 := 0; fi11 < 18; fi11++ {
            switch (^(fi11)) {
            case 32768:
                var v12 int64 = 42
                _ = v12
            case 9223372036854775807:
                for fi13 := 0; fi13 < 20; fi13++ {
                    v14 := GenS1{F0: -1}
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi11
        }
    case "a":
        var v15 uint64 = 2
        _ = v15
    case "hello":
        switch 89 {
        case 128:
            var v16 uint8 = 213
            _ = v16
        default:
            v1 -= float64(float32(float64(1e10)))
        }
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        fmt.Println(v6, v3)
        v3 = (v3 - v3)
        _ = fi17
    }
    fmt.Println(v6)
    fmt.Println(v6, v9)
    v1 -= v6
    delete(v2, 9223372036854775807)
    for fi18 := 0; fi18 < 7; fi18++ {
        v1 += ((1.5 * (v6 * 1e10)) / v1)
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
