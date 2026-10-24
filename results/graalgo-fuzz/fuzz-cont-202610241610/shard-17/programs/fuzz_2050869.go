package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 bool; F2 int16 }

func main() {
    var v1 bool = (32767 < 86)
    var v2 float64 = 0.0
    var v3 uint8 = 74
    for fi4 := 0; fi4 < 10; fi4++ {
        v1 = v1
        _ = fi4
    }
    if (!v1) {
        if true {
            v5 := GenS1{F0: 42}
            switch 128 {
            case 4294967295:
                v1 = true
            }
            _ = v5
        }
    } else {
        var v6 bool = (!v1)
        _ = v6
    }
    fmt.Println(v1, v3, v3)
    fmt.Println(v1, v1, v3)
    v3 = 143
    fmt.Println(v1, v2, v1)
    fmt.Println(v2)
    if (9223372036854775807 <= -1) {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := GenS1{F0: 2147483647}
            v9 := GenS2{F0: fi7, F1: true, F2: 57}
            v10 := GenS1{F0: 2147483647}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(v2, v2)
            _ = fi11
        }
    }
    var v12 float64 = (v2 + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
