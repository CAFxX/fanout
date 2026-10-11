package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
func genFn0() bool {
    var v7 uint64 = 645
    v8 := GenS1{F0: (0.0 <= (3.14159 / 1e10)), F1: 1}
    v9 := GenS1{F0: true, F1: 65535}
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 0
    v3 := []uint32{1}
    var v4 int8 = 0
    var v5 uint64 = 589
    var v6 float32 = 1.5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v6, v2)
    var v10 int32 = -509
    var v11 int8 = (v4 & (v4 ^ 0))
    var v12 uint64 = v5
    v13 := GenS1{F0: ((1 != -32768) || ("" >= "")), F1: 619}
    for fi14 := 0; fi14 < 4; fi14++ {
        if false {
            v15 := GenS1{F0: (3.14159 != 1e10), F1: 42}
            switch fi14 {
            case 50:
                v16 := GenS1{F0: (v10 >= 2147483647), F1: 2147483648}
                _ = v16
            }
            _ = v15
        }
        v2 = v2
        _ = fi14
    }
    switch "hello" {
    case "foo bar":
        _ = len(v1)
    case "\n\t":
        fmt.Println(v4, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
