package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 255
    v3 := make(map[int]uint32)
    var v4 int8 = -128
    var v5 string = "a"
    var v6 float32 = 0.0
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int16 = int16(v2)
        v9 := GenS1{F0: 42, F1: 815}
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 uint = 368
            if true {
                v11 -= (^(v11))
                v1[4294967295] = "a"
            } else {
                v12 := GenS1{F0: 2, F1: 65535}
                v13 := []int{fi7}
                _ = v12
                _ = v13
            }
            fmt.Println(v4)
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v14 := GenS2{F0: v5, F1: -16}
    fmt.Println(v2)
    v5 = ((("foo bar" + (v5 + v5)) + v5) + "\n\t")
    var v15 int32 = 835
    v16 := GenS1{F0: 65535, F1: 42}
    v17 := []int{}
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
