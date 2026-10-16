package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }

func main() {
    var v1 int64 = 342
    var v2 uint = 732
    var v3 bool = true
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := []int{}
        _ = fi4
        _ = v5
    }
    var v6 int8 = 42
    var v7 uint32 = 883
    v1 += (v1 | -867)
    switch 128 {
    case -32:
        v8 := GenS1{F0: v6, F1: 788}
        _ = v8
    }
    fmt.Println(v3)
    var v9 float32 = (1e10 - 0.1)
    v2 = v2
    v7 = 2
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := GenS1{F0: v6, F1: 1000}
        if true {
            var v12 float32 = v9
            if (1 > 0) {
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 uint16 = 338
                    var v15 int64 = v1
                    v14 = 182
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 bool = (v3 && (v12 < v12))
                if v3 {
                    var v17 uint = (v2 / (((v2 / ((v2) | 1))) | 1))
                    _ = v17
                }
                _ = v16
            }
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    if (!((v7 >> ((42) & 63)) == v7)) {
        v18 := []int{-9223372036854775808, 1, 0, -2147483648}
        v19 := GenS1{F0: v6, F1: -328}
        _ = v18
        _ = v19
    } else {
        switch 1 {
        case 32767:
            if v3 {
                v9 = (v9 / 3.14159)
                var v20 int16 = 903
                _ = v20
            }
        case 0:
            fmt.Println(v7, v9, v1)
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            v9 = v9
            v22 := []int{}
            var v23 uint16 = 918
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v3 = false
    v9 = (0.0 / v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
