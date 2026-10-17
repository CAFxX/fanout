package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 bool = (1.5 == -2.25)
    for fi4 := 0; fi4 < 13; fi4++ {
        v3 = false
        for fi5 := 0; fi5 < 9; fi5++ {
            v3 = (849 == 885)
            var v6 uint8 = 2
            var v7 bool = (!v3)
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    var v8 int8 = 1
    fmt.Println(v3, v3)
    _ = len(v1)
    v9 := GenS1{F0: v8}
    v8 -= (v8 + v8)
    var v10 int64 = 153
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
