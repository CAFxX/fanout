package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int32 }

func main() {
    var v1 uint16 = 1
    v2 := []uint32{191, 0, 65535, 2147483647}
    v3 := make(map[uint32]string)
    var v4 int = 32768
    v5 := GenS1{F0: 0, F1: 0, F2: 1}
    v1 -= v1
    var v6 bool = (1 >= 699)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v4)
    v1 += v1
    delete(v3, 964)
    if true {
        v7 := GenS1{F0: 607, F1: 2, F2: 42}
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        if (2 <= 1) {
            v3[2147483647] = ((("hello" + "") + "foo bar") + ("\n\t" + "foo bar"))
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi8
    }
    fmt.Println(v6, v1, v6)
    if (v6 || v6) {
        v6 = (65 == 246)
        var v9 uint8 = 1
        _ = v9
    } else {
        fmt.Println(v6)
    }
    fmt.Println(v4, v1, v4)
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 float64 = 3.14159
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 2147483648, F1: 599, F2: -607}
    v13 := GenS1{F0: 507, F1: 255, F2: 42}
    v4 -= 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
