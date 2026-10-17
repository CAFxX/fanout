package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 string; F2 int }
func genFn0(p0 uint16, p1 uint16, p2 int64) int64 {
    v5 := GenS1{F0: ("" + "a")}
    if (391 == 4294967295) {
        if true {
            p0 += p1
            p0 -= p1
        }
        fmt.Println(p0, p0)
    }
    v6 := GenS1{F0: "hello"}
    var v7 int64 = p2
    var v8 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint64 = 42
    var v2 int32 = 42
    var v3 uint32 = 777
    v4 := []int{32768, -1}
    var v9 int = int(v2)
    v9 -= -791
    v9 = (v9 | (v9 | (v9 << ((1) & 63))))
    for fi10 := 0; fi10 < 11; fi10++ {
        v9 = 1
        v1 = (v1 * v1)
        v11 := GenS2{F0: ((v1 | (v1 % (((^(v1))) | 1))) != v1), F1: "a", F2: v9}
        _ = fi10
        _ = v11
    }
    fmt.Println(v3)
    v2 = v2
    _ = len(v4)
    for fi12 := 0; fi12 < 4; fi12++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi12
    }
    v13 := GenS1{F0: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
