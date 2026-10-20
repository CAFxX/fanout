package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []int{128, 32768, 696, 32767}
    var v2 bool = (-6 != 349)
    var v3 int = 0
    fmt.Println(v2, v2)
    fmt.Println(v3, v3)
    v3 = (^(v3))
    for fi4 := 0; fi4 < 2; fi4++ {
        v2 = false
        _ = fi4
    }
    v5 := GenS1{F0: 9223372036854775807}
    switch "" {
    case "a":
        if (v2 && false) {
            _ = len(v1)
            v2 = v2
        }
    default:
        for fi6 := 0; fi6 < 12; fi6++ {
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 uint8 = 128
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
    }
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
}
