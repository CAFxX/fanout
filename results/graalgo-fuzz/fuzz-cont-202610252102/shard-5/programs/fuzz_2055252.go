package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
func genFn0(p0 int8) int64 {
    v6 := GenS1{F0: 0, F1: -1, F2: (-(p0))}
    v7 := GenS1{F0: 1000, F1: 42, F2: p0}
    _ = p0
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -1
    v3 := make(map[string]string)
    var v4 string = "\n\t"
    v5 := []string{}
    fmt.Println(v4, v4)
    fmt.Println(v4, v2)
    v2 -= 1
    switch ((v4 + v4) + v4) {
    case "":
        delete(v1, 0)
    case "hello":
        fmt.Println(v4, v4)
    default:
        v4 = v4
    }
    v8 := GenS1{F0: -1, F1: v2, F2: 127}
    var v9 int32 = (v2 % (((^(v2))) | 1))
    var v10 uint64 = 860
    for fi11 := 0; fi11 < 11; fi11++ {
        switch fi11 {
        case 9223372036854775807:
            v2 = int32(fi11)
        case 57:
            v4 = "a"
        }
        v12 := GenS1{F0: -561, F1: (-(v2)), F2: 1}
        _ = fi11
        _ = v12
    }
    v9 += 906
    v10 -= (v10 / (((v10 ^ (v10 - v10))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
