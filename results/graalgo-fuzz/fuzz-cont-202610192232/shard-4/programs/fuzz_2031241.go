package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 165
    var v3 int32 = 681
    switch ("a" + "hello") {
    case "foo bar":
        v4 := []int{-9223372036854775808, -159, 32767}
        _ = v4
    case "hello":
        var v5 float64 = float64((1e10 * float32((1e10 / float64(1e10)))))
        _ = v5
    case "\n\t":
        if (true || (1 <= 0)) {
            for fi6 := 0; fi6 < 11; fi6++ {
                v2 -= 2
                v3 = 2147483647
                _ = fi6
            }
            for fi7 := 0; fi7 < 5; fi7++ {
                fmt.Println(v3, v3)
                _ = fi7
            }
        }
    default:
        v8 := GenS1{F0: (v3 & (v3 & -1))}
        _ = v8
    }
    var v9 uint32 = 515
    var v10 float32 = 0.0
    fmt.Println(v2, v10)
    fmt.Println(v3, v3, v2)
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(v10, v3, v9)
        v9 = v9
        _ = fi11
    }
    v12 := []int{-9223372036854775808, -874, 4294967295}
    if (!(true || (false || true))) {
        v3 += -740
    }
    fmt.Println(v9, v10, v3)
    v13 := GenS1{F0: v3}
    v10 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
