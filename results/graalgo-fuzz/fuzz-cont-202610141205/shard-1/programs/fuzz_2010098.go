package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }

func main() {
    var v1 uint16 = 42
    v2 := []int{2147483647, -9223372036854775808, 127}
    v3 := []string{}
    v1 -= v1
    for fi4 := 0; fi4 < 9; fi4++ {
        switch "hello" {
        case "":
            for fi5 := 0; fi5 < 15; fi5++ {
                var v6 string = "\n\t"
                v6 = v6
                v6 = "\n\t"
                fmt.Println(v6, fi4, v6)
                _ = fi5
                _ = v6
            }
        case "foo bar":
            v7 := GenS1{F0: (v1 << ((2) & 63)), F1: -1}
            _ = v7
        }
        if (true || false) {
            _ = len(v2)
            var v8 int8 = 100
            _ = v8
        } else {
            v9 := GenS1{F0: (^(v1)), F1: 32767}
            v10 := GenS1{F0: 1, F1: -1}
            _ = v10
            _ = v9
        }
        v1 += v1
        _ = fi4
    }
    var v11 int16 = -1
    var v12 int64 = 1
    fmt.Println(v11)
    var v13 uint64 = 42
    v3 = append(v3, "hello")
    var v14 bool = true
    for fi15 := 0; fi15 < 20; fi15++ {
        _ = len(v3)
        _ = fi15
    }
    if v14 {
        v1 = (v1 * 378)
    } else {
        if v14 {
            v13 = v13
        }
        v16 := GenS1{F0: v1, F1: v11}
        if len(v2) > 0 {
            v2[0] = 32767
        }
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
