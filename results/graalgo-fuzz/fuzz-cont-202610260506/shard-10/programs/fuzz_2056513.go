package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint32 }

func main() {
    var v1 int = 42
    var v2 bool = true
    v3 := make(map[string]bool)
    v4 := []string{("\n\t" + "foo bar"), "", "\n\t", "foo bar"}
    v1 -= 128
    fmt.Println(v2)
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 float32 = 1.5
            _ = fi6
            _ = v7
        }
        v1 = fi5
        v8 := v3["foo bar"]
        _ = fi5
        _ = v8
    }
    v1 += v1
    var v9 uint64 = 42
    v4 = append(v4, "hello")
    v10 := GenS1{F0: 2, F1: 128, F2: 0}
    v2 = (v2 || (!((v9 << ((950) & 63)) > v9)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
