package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 string }
type GenS2 struct { F0 int; F1 string; F2 uint64 }
type GenS3 struct { F0 int; F1 int16; F2 int }
type GenS4 struct { F0 uint64; F1 int8; F2 uint16 }
func genFn0(p0 string, p1 uint8, p2 int16) int8 {
    v7 := GenS1{F0: 42, F1: -1, F2: (p0 + (("a" + p0) + "foo bar"))}
    var v8 int8 = -35
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (-(v8))
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]int)
    var v3 int16 = -1
    v4 := make(map[string]uint32)
    var v5 float32 = (1e10 + -2.25)
    var v6 float32 = ((v5 / v5) * 1.5)
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int = fi9
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        if (!((1e10 - -2.25) != (-2.25 + 0.1))) {
            if true {
                var v12 uint64 = 2
                v13 := GenS2{F0: fi11, F1: "", F2: (^(v12))}
                _ = v12
                _ = v13
            }
        }
        v14 := []int{}
        _ = fi11
        _ = v14
    }
    _ = len(v4)
    v15 := v4["foo bar"]
    v16 := []int{932}
    var v17 bool = (0 == 128)
    v18 := GenS3{F0: -2147483648, F1: v3, F2: 506}
    v19 := GenS4{F0: 0, F1: 0, F2: 0}
    v4[(("hello" + "hello") + "")] = v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
