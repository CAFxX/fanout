package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
func genFn0(p0 int8, p1 int32) int8 {
    fmt.Println(p1, p0)
    if false {
        for fi4 := 0; fi4 < 10; fi4++ {
            if (p1 == (p1 % ((638) | 1))) {
                v5 := GenS1{F0: 42}
                _ = v5
            }
            var v6 uint64 = 18446744073709551615
            _ = fi4
            _ = v6
        }
        var v7 int64 = 934
        _ = v7
    }
    var v8 int = 42
    _ = p0
    _ = p1
    _ = v8
    return p0
}
func genFn1(p0 uint) uint16 {
    v9 := GenS1{F0: 42}
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 1}
    v13 := GenS2{F0: (p0 * (p0 - p0))}
    var v14 int = 4294967295
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 1
    var v3 uint8 = 2
    v3 += (v3 - (v3 * v3))
    v15 := []int{}
    v16 := GenS3{F0: true}
    v3 = v3
    var v17 int8 = -128
    var v18 uint64 = 1
    fmt.Println(v17, v17)
    if len(v15) > 0 {
        v15[0] = 596
    }
    var v19 int8 = v17
    switch "\n\t" {
    case "hello":
        v20 := GenS4{F0: 9223372036854775807}
        _ = v20
    case "a":
        var v21 uint = 472
        _ = v21
    default:
        delete(v1, "hello")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
