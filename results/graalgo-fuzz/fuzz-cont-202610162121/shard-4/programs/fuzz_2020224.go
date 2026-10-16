package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint8 }

func main() {
    var v1 bool = (!true)
    var v2 uint = 1
    var v3 uint64 = 0
    var v4 int64 = 1
    v4 += (v4 * (v4 << ((v2) & 63)))
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: "hello", F1: v2, F2: 42}
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := []int{fi5, fi5, (fi5 * (fi5 * (fi7 & 0))), fi7}
            _ = fi7
            _ = v8
        }
        v9 := []int{fi5, fi5}
        _ = fi5
        _ = v6
        _ = v9
    }
    fmt.Println(v2, v2, v2)
    if (v3 >= 352) {
        v4 = v4
    }
    var v10 uint = (v2 - 1)
    var v11 bool = v1
    v1 = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
