package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }
type GenS2 struct { F0 int16; F1 bool }

func main() {
    var v1 float32 = (0.0 - float32(((-2.25 - float64(1.5)) - (1.5 - 0.1))))
    var v2 string = "\n\t"
    var v3 int64 = 13
    if (!false) {
        v2 = v2
        if true {
            fmt.Println(v3, v1)
            for fi4 := 0; fi4 < 16; fi4++ {
                switch v2 {
                case "foo bar":
                    var v5 uint16 = 656
                    _ = v5
                }
                _ = fi4
            }
        } else {
            v6 := GenS1{F0: 2147483647, F1: 42, F2: 42}
            var v7 uint32 = 559
            _ = v6
            _ = v7
        }
        v8 := []int{}
        _ = v8
    }
    var v9 int32 = 0
    switch -1 {
    case -321:
        var v10 float32 = v1
        _ = v10
    }
    fmt.Println(v2)
    fmt.Println(v2, v9, v3)
    var v11 string = v2
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(v11, v9)
        _ = fi12
    }
    v3 -= (v3 * 42)
    v13 := []int{0, 127, 4294967295}
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint16 = 0
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 1, F1: (v9 >= (v9 | (v9 * (v9 << ((861) & 63)))))}
    v3 = 783
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v17 uint32 = 2147483648
    var v18 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
