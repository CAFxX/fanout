package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 string }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int32 = 354
    var v2 int32 = v1
    var v3 uint32 = 1
    for fi4 := 0; fi4 < 7; fi4++ {
        switch fi4 {
        case 32768:
            if true {
                fmt.Println(v2, fi4, v3)
                fmt.Println(v1, v3, v3)
            }
        }
        switch (fi4 % ((2147483647) | 1)) {
        case 0:
            v3 -= 0
        case 128:
            fmt.Println(v2, v3)
        }
        fmt.Println(fi4)
        _ = fi4
    }
    v5 := GenS1{F0: v3, F1: v3, F2: "foo bar"}
    v3 = v3
    var v6 float32 = float32((-2.25 - 0.0))
    v7 := GenS2{F0: 1000}
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := []int{fi9}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := []int{}
    for fi12 := 0; fi12 < 17; fi12++ {
        fmt.Println(fi12, v3)
        v11 = append(v11, -9223372036854775808)
        _ = fi12
    }
    v3 -= (v3 & (v3 >> ((4294967295) & 63)))
    switch 32768 {
    case 0:
        v13 := GenS2{F0: -708}
        _ = v13
    }
    v2 = (v2 | 0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
