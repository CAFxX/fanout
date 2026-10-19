package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]uint32)
    var v3 string = "foo bar"
    v1 -= v1
    v4 := []int{}
    for fi5 := 0; fi5 < 9; fi5++ {
        if ((!(!(128 < 1))) && true) {
            var v6 int16 = 295
            _ = v6
        }
        _ = fi5
    }
    fmt.Println(v1, v3)
    v3 = ""
    v3 = "a"
    var v7 float32 = ((0.1 / v1) - v1)
    var v8 float64 = 1e10
    if false {
        v9 := GenS1{F0: 0, F1: 255}
        _ = v9
    }
    v3 = v3
    if (false || (!(127 >= 1))) {
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int32 = 2147483647
            var v12 uint64 = 18446744073709551615
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS1{F0: 803, F1: 0}
            switch (fi13 ^ fi13) {
            case 127:
                v2["hello"] = 2147483647
            case 0:
                v15 := GenS1{F0: 2147483648, F1: 42}
                _ = v15
            case 128:
                fmt.Println(v8, v3)
            default:
                if true {
                    var v16 bool = false
                    delete(v2, (((v3 + "\n\t") + "\n\t") + (v3 + (v3 + v3))))
                    _ = v16
                }
            }
            v2[v3] = 534
            _ = fi13
            _ = v14
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
