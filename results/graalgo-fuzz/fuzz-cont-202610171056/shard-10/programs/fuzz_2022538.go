package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 string; F1 uint }

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 4294967295
    var v3 uint64 = v2
    var v4 float64 = (1e10 * -2.25)
    var v5 uint = 0
    var v6 uint64 = (v2 >> (((v5 * (v5 & (^(v5))))) & 63))
    if (true || true) {
        for fi7 := 0; fi7 < 15; fi7++ {
            v3 += 0
            var v8 int8 = 127
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := []int{}
                v8 = v8
                fmt.Println(fi9, v6, v5)
                if (((3.14159 < (((1.5 * (3.14159 / float32(v4))) - float32((v4 / (v4 * 1.5)))) / 1e10)) && (!true)) || false) {
                    v11 := GenS1{F0: 42, F1: (v4 <= (v4 - (1.5 / v4))), F2: v1}
                    _ = v11
                }
                v3 = v3
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        v12 := GenS2{F0: ("foo bar" + ("foo bar" + ("foo bar" + "a"))), F1: 0}
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v4 = ((float64(-2.25) + float64(1.5)) + (1e10 / v4))
        v4 = -2.25
        v14 := []int{fi13, 315, 0, 128}
        _ = fi13
        _ = v14
    }
    v15 := []int{2147483647}
    v4 = float64(3.14159)
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := GenS2{F0: "foo bar", F1: 850}
        _ = fi16
        _ = v17
    }
    if true {
        _ = len(v15)
        v1 -= (v1 % ((v1) | 1))
        fmt.Println(v5, v3)
    }
    v5 += 1
    _ = len(v15)
    fmt.Println(v2, v3, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
