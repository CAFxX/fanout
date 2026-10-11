package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 string }
type GenS2 struct { F0 int64; F1 uint16; F2 uint8 }

func main() {
    var v1 uint16 = 42
    var v2 float32 = 0.1
    var v3 bool = false
    var v4 float64 = (((0.1 / 1e10) / float64(float32(float64((((v2 / v2) - v2) * (v2 + 3.14159)))))) - 3.14159)
    var v5 float32 = v2
    var v6 int64 = -1
    v3 = v3
    v1 += v1
    if (!(!v3)) {
        fmt.Println(v3, v1)
        var v7 uint64 = 1
        if true {
            v2 += ((((((3.14159 / -2.25) - float32(v4)) - v5) / v5) * v5) * (v2 - v2))
        }
        _ = v7
    }
    if v3 {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 string = ("" + "foo bar")
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS1{F0: v1, F1: v3, F2: ""}
                var v12 uint8 = 95
                var v13 uint8 = v12
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            if v3 {
                if true {
                    v14 := []int{}
                    for fi15 := 0; fi15 < 6; fi15++ {
                        fmt.Println(fi15, fi15, v4)
                        _ = fi15
                    }
                    _ = v14
                }
            }
            _ = fi8
            _ = v9
        }
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        if v3 {
            var v17 int16 = 1000
            v18 := []int{0, 582}
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = v17
            _ = v18
        } else {
            v2 -= ((v5 / float32(((float64(v5) - (float64(-2.25) / 1e10)) + v4))) / v2)
        }
        v19 := GenS2{F0: v6, F1: 722, F2: 42}
        v6 -= v6
        _ = fi16
        _ = v19
    }
    v3 = (v3 && (-2147483648 != 952))
    v20 := []int{-956}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
