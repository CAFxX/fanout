package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
func genFn0(p0 uint64, p1 int8, p2 uint) bool {
    v5 := GenS1{F0: 65535, F1: ""}
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: 0, F1: "foo bar"}
        fmt.Println(p2, fi6)
        v8 := []int{}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    v10 := GenS1{F0: 777, F1: ""}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v9
    return ((p0 | (v9 & uint64(p2))) >= p0)
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    v3 := []uint32{736}
    var v4 int8 = -1
    fmt.Println(v4)
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(fi11)
        _ = fi11
    }
    if true {
        v4 = 127
        v12 := GenS1{F0: 2, F1: ""}
        _ = v12
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 uint64 = 42
    fmt.Println(v13)
    var v14 float64 = 0.0
    switch 191 {
    case -2147483648:
        _ = len(v3)
    }
    if false {
        if (true && true) {
            v15 := GenS1{F0: 2147483648, F1: "hello"}
            var v16 uint = 42
            _ = v15
            _ = v16
        } else {
            v13 += (^(v13))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
