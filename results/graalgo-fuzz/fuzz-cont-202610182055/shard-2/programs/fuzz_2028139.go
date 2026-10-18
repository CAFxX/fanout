package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int64 }

func main() {
    var v1 float64 = (0.1 + (0.1 - (1e10 / 1.5)))
    var v2 uint16 = 477
    var v3 int8 = 127
    for fi4 := 0; fi4 < 6; fi4++ {
        v5 := []int{fi4}
        var v6 int32 = -1
        v7 := GenS1{F0: (^(v2))}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 float32 = 1.5
    if false {
        v9 := []int{2147483647, 922, 894, 1}
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 float64 = float64(0.0)
            if true {
                v1 = v11
                fmt.Println(v3)
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        v12 := []int{-66, 2147483647, 158, 4294967295}
        _ = v12
    }
    fmt.Println(v8, v2, v2)
    if false {
        if true {
            v13 := GenS2{F0: -1, F1: -9223372036854775808}
            _ = v13
        }
        v14 := []int{}
        var v15 uint64 = 0
        _ = v14
        _ = v15
    }
    var v16 int8 = -96
    var v17 uint = 42
    v18 := GenS1{F0: (^(v2))}
    if (((v8 - float32(3.14159)) / 1.5) != v8) {
        v2 -= 32767
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v8
}
