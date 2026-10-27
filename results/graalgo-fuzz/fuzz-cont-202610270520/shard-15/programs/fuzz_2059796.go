package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]string)
    var v3 int = 32768
    var v4 int = v3
    v1 = float32((1.5 + -2.25))
    if false {
        switch v4 {
        case 128:
            var v5 int32 = -322
            _ = v5
        case 0:
            v6 := []int{(v3 - v3), 4294967295, (^(v3)), 32767}
            _ = v6
        case 754:
            v7 := GenS1{F0: -485}
            _ = v7
        }
    }
    var v8 float32 = v1
    var v9 int16 = 1000
    v10 := GenS1{F0: v4}
    switch "hello" {
    case "\n\t":
        v11 := []int{v3, 960}
        _ = v11
    default:
        fmt.Println(v1)
    }
    v1 = v1
    var v12 int = (v4 % (((v3 >> ((296) & 63))) | 1))
    v12 += (v12 + v12)
    delete(v2, 945)
    fmt.Println(v9)
    switch (-(v4)) {
    case 2147483647:
        if true {
            v3 = 1
        }
    case 0:
        v13 := []int{}
        _ = v13
    case 4294967295:
        if (((!false) && (0 >= 0)) && (!true)) {
            fmt.Println(v12)
            var v14 bool = (2 == 952)
            var v15 int8 = 100
            _ = v14
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 9; fi16++ {
            for fi17 := 0; fi17 < 13; fi17++ {
                v4 += 4294967295
                var v18 uint = 251
                _ = fi17
                _ = v18
            }
            v9 = (v9 - (v9 + (v9 / ((v9) | 1))))
            _ = fi16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
