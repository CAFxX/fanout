package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (((0.0 * 1.5) - (1e10 - float64(((-2.25 * 0.0) * 0.1)))) * 0.1)
    var v3 string = "foo bar"
    v4 := make(map[uint32]uint32)
    v5 := []int{42}
    if (42 != 0) {
        var v6 uint16 = 0
        var v7 int16 = 946
        _ = v6
        _ = v7
    }
    v3 = ((("hello" + "") + v3) + "\n\t")
    var v8 uint = 1
    v9 := GenS1{F0: 335, F1: 165}
    v3 = "a"
    for fi10 := 0; fi10 < 11; fi10++ {
        _ = len(v5)
        _ = fi10
    }
    if (!(-1 <= -1)) {
        delete(v1, "hello")
        v3 = ((("a" + "a") + "") + ((v3 + (v3 + ((v3 + "a") + v3))) + v3))
        for fi11 := 0; fi11 < 13; fi11++ {
            _ = len(v1)
            _ = fi11
        }
    }
    var v12 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
