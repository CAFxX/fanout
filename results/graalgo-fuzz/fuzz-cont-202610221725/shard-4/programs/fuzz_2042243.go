package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 int32 = 0
    v2 := make(map[string]int)
    var v3 bool = (1e10 >= -2.25)
    _ = len(v2)
    v3 = (("a" + "\n\t") != (("a" + "") + (("foo bar" + "a") + "hello")))
    v4 := []int{}
    v5 := GenS1{F0: "\n\t"}
    fmt.Println(v1)
    if v3 {
        for fi6 := 0; fi6 < 2; fi6++ {
            v2["\n\t"] = (^(fi6))
            _ = fi6
        }
        v3 = (v3 && v3)
    } else {
        v3 = v3
        _ = len(v2)
        delete(v2, "a")
    }
    v1 -= (v1 * 943)
    v7 := GenS2{F0: 2147483648, F1: 0}
    if len(v4) > 0 {
        v4[0] = 486
    }
    var v8 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
