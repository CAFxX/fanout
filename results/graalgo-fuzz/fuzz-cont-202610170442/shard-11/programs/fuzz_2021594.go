package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint8 }

func main() {
    var v1 int8 = 9
    v2 := make(map[uint32]string)
    v3 := []int{}
    v4 := []int{840}
    var v5 bool = (127 < 32767)
    switch 1 {
    case -859:
        v6 := v2[4294967295]
        _ = v6
    case 256:
        v7 := GenS1{F0: 191, F1: 2}
        _ = v7
    case 32767:
        var v8 bool = v5
        _ = v8
    }
    v1 = v1
    var v9 int32 = 1
    if (42 >= 1) {
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := GenS2{F0: int8(fi10), F1: 128}
            fmt.Println(v5, v5, fi10)
            _ = fi10
            _ = v11
        }
        var v12 float64 = 3.14159
        _ = v12
    }
    v9 = int32(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint16 = 32767
        var v15 float64 = 3.14159
        _ = fi13
        _ = v14
        _ = v15
    }
    if len(v3) > 0 {
        v3[0] = 465
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 float64 = 3.14159
        for fi18 := 0; fi18 < 9; fi18++ {
            if v5 {
                var v19 int16 = 0
                var v20 uint = 42
                if ((v19 ^ 1000) <= v19) {
                    v1 = v1
                    _ = len(v4)
                    v20 = 1
                }
                _ = v19
                _ = v20
            }
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    var v21 string = "hello"
    v21 = v21
    v21 = "\n\t"
    v9 = (v9 & (v9 + v9))
    v22 := v2[1]
    var v23 float32 = (1e10 + 3.14159)
    v2[4294967295] = "a"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
