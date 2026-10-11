package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint8 = 2
    var v2 uint16 = uint16(v1)
    v3 := []int{}
    var v4 uint32 = 0
    var v5 int = 2147483647
    var v6 float32 = 0.0
    var v7 string = "a"
    v4 = v4
    var v8 int64 = -9223372036854775808
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint = 2
        v1 = v1
        _ = len(v3)
        _ = fi9
        _ = v10
    }
    var v11 int16 = 583
    for fi12 := 0; fi12 < 13; fi12++ {
        if (!(202 <= 4294967295)) {
            v13 := GenS1{F0: v1, F1: ((v11 != v11) && (true && (v6 > (v6 + float32((1.5 - ((1e10 / (float64((-2.25 + v6)) / 1.5)) + ((((0.1 + -2.25) + -2.25) - float64(v6)) + 3.14159))))))))}
            v7 = (v7 + "a")
            v5 += int(v1)
            _ = v13
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v4 += v4
            _ = fi14
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        v8 += v8
        v6 = 1e10
        _ = fi15
    }
    v7 = ("a" + v7)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
