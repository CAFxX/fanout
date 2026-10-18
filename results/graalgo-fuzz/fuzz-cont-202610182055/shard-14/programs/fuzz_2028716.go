package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[uint32]bool)
    var v3 uint32 = 159
    v4 := make(map[string]string)
    var v5 int8 = 127
    v6 := GenS1{F0: 2}
    v5 = v5
    for fi7 := 0; fi7 < 18; fi7++ {
        v4[""] = "hello"
        _ = fi7
    }
    switch 127 {
    case 128:
        v8 := GenS1{F0: 127}
        _ = v8
    case 9:
        v3 -= v3
    case 626:
        v3 = 847
    }
    var v9 int8 = v5
    v3 += (v3 ^ (v3 * (v3 ^ 706)))
    v5 = 42
    v3 -= (v3 * v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
