package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int16 }

func main() {
    var v1 int64 = 42
    var v2 uint = 2
    var v3 uint32 = 2147483648
    var v4 bool = (v2 == v2)
    var v5 int64 = v1
    var v6 int8 = 42
    v7 := GenS1{F0: 886, F1: -13, F2: 1}
    v8 := GenS1{F0: v3, F1: -1, F2: 1}
    var v9 int16 = 32767
    switch ("\n\t" + "\n\t") {
    case "hello":
        v1 = v1
    default:
        switch 127 {
        case 1:
            for fi10 := 0; fi10 < 2; fi10++ {
                v4 = false
                fmt.Println(v3, v6, v4)
                _ = fi10
            }
        case 42:
            fmt.Println(v4, v3, v1)
        }
    }
    fmt.Println(v1)
    var v11 float32 = (1e10 * 0.0)
    if true {
        v2 = (v2 % ((v2) | 1))
        v12 := []int{}
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = v12
    }
    v4 = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
