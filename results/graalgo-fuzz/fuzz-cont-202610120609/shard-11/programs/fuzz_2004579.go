package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 bool }
type GenS2 struct { F0 uint32; F1 uint16; F2 int16 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint16 = 2
    var v3 uint = 160
    v4 := GenS1{F0: -9223372036854775808, F1: 32767, F2: false}
    var v5 int16 = -32768
    v5 = 1000
    v5 = 0
    v2 = 1
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: 4294967295, F1: v5, F2: (false && true)}
        if true {
            fmt.Println(v1, v2, v2)
            v8 := []int{(^(fi6)), fi6, fi6, fi6}
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS2{F0: 444, F1: (v2 % ((v2) | 1)), F2: 874}
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            v5 += v5
            if false {
                var v11 string = "foo bar"
                _ = v11
            }
        }
        var v12 int8 = -1
        _ = fi6
        _ = v12
        _ = v7
    }
    var v13 int8 = 100
    var v14 uint64 = 4294967295
    v15 := []int{0, 320, 42, 1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
