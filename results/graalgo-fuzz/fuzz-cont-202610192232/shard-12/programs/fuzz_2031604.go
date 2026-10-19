package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }
func genFn0(p0 uint16, p1 string, p2 int8) string {
    v5 := []int{859, 0, 455, 9223372036854775807}
    p0 += 0
    v6 := GenS1{F0: 2, F1: (p0 - p0)}
    fmt.Println(p2, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p1
}
func genFn1(p0 uint8) string {
    var v7 uint16 = 216
    fmt.Println(p0)
    v8 := GenS1{F0: 2, F1: 32767}
    v9 := []int{848}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return ""
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = (-2.25 / 1e10)
    v3 := make(map[string]uint32)
    var v4 uint8 = 127
    switch "foo bar" {
    case "a":
        if true {
            v4 += v4
            v10 := []int{32767, 42, 42, -9223372036854775808}
            _ = v10
        }
    }
    v11 := []int{}
    var v12 uint64 = 18446744073709551615
    for fi13 := 0; fi13 < 17; fi13++ {
        v12 = (^(v12))
        _ = fi13
    }
    v2 = v2
    v4 = (v4 + (v4 & v4))
    v14 := GenS1{F0: 1, F1: 255}
    v12 = v12
    v15 := GenS2{F0: "a", F1: (v4 % ((v4) | 1)), F2: (v4 - uint8(v12))}
    fmt.Println(v12, v4)
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint32 = 2147483648
        _ = fi16
        _ = v17
    }
    fmt.Println(v4)
    if (false || true) {
        _ = len(v11)
        v18 := GenS2{F0: "a", F1: 127, F2: (v4 % ((v4) | 1))}
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
