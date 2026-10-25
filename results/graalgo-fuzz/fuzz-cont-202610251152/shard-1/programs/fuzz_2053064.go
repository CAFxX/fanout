package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }

func main() {
    var v1 uint8 = 255
    var v2 int16 = 0
    var v3 int = -898
    if false {
        v2 += v2
        fmt.Println(v1, v2)
        v2 = v2
    }
    var v4 int8 = 127
    fmt.Println(v2, v1)
    v5 := []int{(-(v3)), (^(v3)), v3}
    fmt.Println(v1, v1, v1)
    v3 = v3
    var v6 float32 = (3.14159 - 1e10)
    v7 := GenS1{F0: v4, F1: ("hello" + "foo bar")}
    v8 := GenS1{F0: -128, F1: "foo bar"}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v9 int8 = (v4 ^ 100)
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(v9, v2, v6)
        for fi11 := 0; fi11 < 13; fi11++ {
            fmt.Println(fi11, v9)
            if true {
                v12 := GenS1{F0: -46, F1: ("" + ("hello" + "\n\t"))}
                _ = v12
            } else {
                v9 = (v4 - 1)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi11
        }
        _ = fi10
    }
    var v13 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
