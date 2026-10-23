package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint64 }
func genFn0() int16 {
    var v6 int32 = 0
    v7 := GenS1{F0: "a", F1: -283, F2: 1}
    v6 = 1
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 uint64) string {
    var v8 uint8 = 128
    p0 = 18446744073709551615
    p0 = (p0 | p0)
    var v9 int64 = -1
    var v10 int = -2147483648
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    var v1 float32 = 0.1
    var v2 uint64 = 98
    v3 := make(map[int]int)
    var v4 int32 = -644
    v5 := make(map[string]uint32)
    fmt.Println(v4)
    switch -806 {
    case 707:
        switch ("foo bar" + ((("" + "foo bar") + "foo bar") + "")) {
        case "foo bar":
            v4 -= 42
        default:
            for fi11 := 0; fi11 < 4; fi11++ {
                fmt.Println(v4, fi11, v1)
                var v12 int8 = -128
                _ = fi11
                _ = v12
            }
        }
    }
    var v13 bool = false
    v1 = v1
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := []int{1, fi14, (fi14 ^ (fi14 - fi14)), fi14}
        var v16 uint64 = (v2 / (((v2 - v2)) | 1))
        v4 = (v4 + -130)
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int = 32767
    v2 = 2
    v2 = 953
    v18 := GenS1{F0: "\n\t", F1: -552, F2: (^(v2))}
    delete(v5, "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
