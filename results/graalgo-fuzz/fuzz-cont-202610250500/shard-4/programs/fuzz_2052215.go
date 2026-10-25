package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int }
func genFn0(p0 int16) uint {
    for fi4 := 0; fi4 < 13; fi4++ {
        if true {
            p0 = p0
        }
        fmt.Println(fi4, fi4)
        _ = fi4
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        p0 = p0
        _ = fi5
    }
    var v6 int64 = 850
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = 249
    v3 := make(map[string]uint32)
    if false {
        var v7 bool = false
        _ = v7
    }
    var v8 uint = 0
    v1 = (v1 + "hello")
    switch 1 {
    case 1:
        v9 := GenS1{F0: 4294967295, F1: 1, F2: 42}
        _ = v9
    }
    v1 = (v1 + ("" + "foo bar"))
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 uint32 = 464
        var v12 int64 = -43
        var v13 int16 = (v2 / (((v2 * (v2 - (v2 & v2)))) | 1))
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v8 = (v8 + v8)
    v8 = v8
    v14 := GenS1{F0: 383, F1: 1, F2: 127}
    v8 = (v8 ^ v8)
    var v15 string = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v8
}
