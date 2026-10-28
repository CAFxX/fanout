package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int64 = -1
    var v2 bool = (false || false)
    v3 := []int{}
    var v4 int64 = (v1 ^ v1)
    _ = len(v3)
    v4 = v1
    v4 -= v4
    v1 = (^(v4))
    v5 := GenS1{F0: 42}
    for fi6 := 0; fi6 < 13; fi6++ {
        v3 = append(v3, fi6)
        for fi7 := 0; fi7 < 18; fi7++ {
            v1 = -1
            if v2 {
                v1 += v4
            }
            v1 -= (v1 / ((v4) | 1))
            _ = fi7
        }
        var v8 uint32 = 2
        _ = fi6
        _ = v8
    }
    v1 += 0
    switch (("hello" + ("a" + "foo bar")) + "hello") {
    case "a":
        var v9 uint = 2
        _ = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
