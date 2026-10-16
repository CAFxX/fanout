package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int16; F1 uint32; F2 int16 }

func main() {
    var v1 bool = (2147483647 == 2147483647)
    var v2 int8 = 1
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    var v5 uint32 = 42
    var v6 int16 = 42
    v7 := []int{2147483647}
    v8 := v3[v5]
    var v9 uint64 = 2
    v2 -= -21
    for fi10 := 0; fi10 < 17; fi10++ {
        v7 = append(v7, fi10)
        v8 = ((v5 - 2147483648) > 2)
        _ = fi10
    }
    switch 901 {
    case 128:
        if false {
            v11 := GenS1{F0: 0, F1: -9223372036854775808}
            var v12 float32 = (3.14159 - 3.14159)
            _ = v11
            _ = v12
        }
    case 4294967295:
        var v13 int32 = -1
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 uint = 305
        var v16 float32 = float32(3.14159)
        v17 := GenS1{F0: -1, F1: -9223372036854775808}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 string = "\n\t"
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := v4[v18]
        v18 = "a"
        if len(v7) > 0 {
            v7[0] = 4294967295
        }
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: -1, F1: (^(v5)), F2: (v6 ^ -32768)}
    var v22 int32 = 0
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 uint16 = 205
        _ = fi23
        _ = v24
    }
    var v25 float64 = 3.14159
    v26 := GenS2{F0: v6, F1: 2147483648, F2: (v6 ^ v6)}
    switch v18 {
    case "":
        var v27 int = int(v6)
        _ = v27
    }
    v9 = (v9 % (((v9 / ((18446744073709551615) | 1))) | 1))
    v25 += v25
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
