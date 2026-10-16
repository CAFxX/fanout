package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int16 }
type GenS2 struct { F0 int; F1 int8; F2 uint64 }

func main() {
    v1 := []string{"a", "\n\t", "", "a"}
    v2 := make(map[string]bool)
    v3 := []int{-175, 1, 214}
    v4 := make(map[string]string)
    var v5 string = "hello"
    switch "foo bar" {
    case "a":
        _ = len(v2)
    case "\n\t":
        var v6 int32 = 2147483647
        _ = v6
    default:
        fmt.Println(v5)
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        switch fi7 {
        case 127:
            v5 = v5
        }
        v8 := GenS1{F0: 86, F1: fi7, F2: 42}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 255, F1: 1, F2: 1}
    for fi10 := 0; fi10 < 9; fi10++ {
        _ = len(v4)
        _ = fi10
    }
    v5 = "\n\t"
    v5 = "\n\t"
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            v5 = ("\n\t" + "\n\t")
            _ = fi12
        }
        v13 := GenS2{F0: fi11, F1: -1, F2: 0}
        _ = fi11
        _ = v13
    }
    var v14 int32 = 42
    if false {
        v15 := GenS1{F0: 350, F1: 2147483647, F2: -1}
        var v16 uint32 = 2
        _ = v15
        _ = v16
    }
    var v17 int = 128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
