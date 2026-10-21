package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint16 }

func main() {
    var v1 uint8 = 166
    var v2 int = 127
    v3 := make(map[string]int)
    var v4 int64 = -248
    var v5 float32 = 1e10
    var v6 int16 = -32768
    var v7 bool = ((884 == 4294967295) && false)
    if (((4294967295 == 2) && true) || false) {
        var v8 int8 = 100
        v7 = v7
        _ = v8
    } else {
        v9 := []int{(v2 & (v2 ^ v2)), v2, 128}
        if v7 {
            v9 = append(v9, v2)
        }
        var v10 int32 = -96
        _ = v10
        _ = v9
    }
    var v11 string = (((("\n\t" + "foo bar") + "hello") + "") + "hello")
    v1 = v1
    v4 = v4
    for fi12 := 0; fi12 < 9; fi12++ {
        fmt.Println(v1)
        _ = fi12
    }
    v13 := GenS1{F0: (v6 ^ (v6 ^ (v6 - (v6 / (((v6 << ((65535) & 63))) | 1))))), F1: 127, F2: 255}
    v3["\n\t"] = int(v1)
    for fi14 := 0; fi14 < 14; fi14++ {
        v7 = v7
        _ = fi14
    }
    v15 := GenS1{F0: (v6 ^ -32768), F1: -128, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
