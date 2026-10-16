package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint8 = 128
    var v2 int8 = 127
    var v3 uint64 = 4294967295
    v4 := GenS1{F0: false, F1: (v1 ^ 208)}
    v5 := GenS1{F0: (9223372036854775807 <= 518), F1: (v1 - (v1 / (((v1 * v1)) | 1)))}
    v1 -= v1
    fmt.Println(v3, v3, v1)
    for fi6 := 0; fi6 < 13; fi6++ {
        if true {
            if true {
                v7 := GenS1{F0: true, F1: v1}
                v2 -= v2
                _ = v7
            } else {
                v8 := []int{(fi6 | (fi6 / (((fi6 ^ (fi6 >> ((18446744073709551615) & 63)))) | 1))), 0}
                var v9 int16 = 895
                _ = v8
                _ = v9
            }
        }
        _ = fi6
    }
    var v10 string = ("a" + "foo bar")
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := []int{}
        var v13 uint8 = v1
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint64 = (v3 % ((v3) | 1))
            _ = fi14
            _ = v15
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    var v16 int32 = 42
    v17 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
