package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int8 }
func genFn0() int {
    switch "hello" {
    case "hello":
        for fi4 := 0; fi4 < 17; fi4++ {
            if true {
                var v5 int = (fi4 % (((^(fi4))) | 1))
                var v6 string = "foo bar"
                fmt.Println(v6)
                _ = v5
                _ = v6
            } else {
                var v7 float32 = 3.14159
                v8 := []int{(fi4 - (fi4 & fi4)), fi4, (fi4 + fi4)}
                _ = v7
                _ = v8
            }
            _ = fi4
        }
    case "a":
        v9 := []int{32768}
        _ = v9
    }
    var v10 uint64 = 673
    var v11 float64 = 1.5
    fmt.Println(v10)
    _ = v10
    _ = v11
    return -547
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 32767
    var v3 uint8 = 175
    var v12 bool = false
    if false {
        v3 += (^(v3))
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 int64 = 9223372036854775807
        _ = fi13
        _ = v14
    }
    v15 := []int{9223372036854775807, 2147483647, 127, -313}
    v16 := GenS1{F0: 2147483647, F1: 42, F2: 44}
    v17 := v1[0]
    if len(v15) > 0 {
        v15[0] = v17
    }
    v17 -= (v17 ^ 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
