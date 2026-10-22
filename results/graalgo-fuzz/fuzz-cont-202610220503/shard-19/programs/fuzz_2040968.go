package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
func genFn0() int16 {
    var v5 uint32 = 2
    if false {
        v6 := GenS1{F0: 9223372036854775807, F1: false}
        v5 = v5
        v5 += 2147483647
        _ = v6
    } else {
        fmt.Println(v5)
    }
    _ = v5
    return -32768
}

func main() {
    v1 := []uint32{65535, 1, 2147483648}
    v2 := []string{"\n\t", "", "foo bar"}
    var v3 bool = (!true)
    v4 := make(map[uint32]int)
    if len(v1) > 0 {
        v1[0] = 130
    }
    var v7 float32 = 0.1
    v3 = true
    switch "a" {
    case "":
        if (100 == 100) {
            var v8 bool = ("" >= "")
            v8 = v8
            if v8 {
                v7 = float32((float64(v7) - -2.25))
                v3 = v3
            }
            _ = v8
        }
    case "foo bar":
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := v4[666]
            _ = fi9
            _ = v10
        }
    }
    v11 := GenS1{F0: -1, F1: v3}
    var v12 uint64 = 4294967295
    v13 := GenS1{F0: 42, F1: v3}
    var v14 uint = 65535
    var v15 uint32 = 65535
    fmt.Println(v3, v12, v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
