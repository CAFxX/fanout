package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }

func main() {
    v1 := []int{2147483647, 127, -1}
    var v2 int = -807
    var v3 int32 = 0
    v2 = (v2 & 32767)
    v2 = v2
    v2 -= (v2 % ((v2) | 1))
    v1 = append(v1, (v2 << ((4294967295) & 63)))
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := GenS1{F0: -1, F1: 100}
        var v6 string = "foo bar"
        fmt.Println(v3)
        _ = fi4
        _ = v5
        _ = v6
    }
    if (("foo bar" + "") > "\n\t") {
        v7 := GenS1{F0: -243, F1: -1}
        _ = len(v1)
        if len(v1) > 0 {
            v1[0] = v2
        }
        _ = v7
    } else {
        if len(v1) > 0 {
            v1[0] = (v2 | v2)
        }
        v3 -= v3
    }
    v1 = append(v1, 146)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 uint32 = 4294967295
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
