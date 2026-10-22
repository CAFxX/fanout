package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 float64 = (float64(float32((((-2.25 / (1.5 / ((-2.25 + (-2.25 - -2.25)) - ((0.1 / (3.14159 / 0.1)) / (float64(float32(0.0)) + 0.1))))) - 3.14159) - (-2.25 / 1.5)))) * float64(((-2.25 - (-2.25 - float32(0.0))) - 0.1)))
    var v2 uint64 = 0
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    for fi5 := 0; fi5 < 13; fi5++ {
        v4[2] = ""
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int16 = -748
            _ = fi6
            _ = v7
        }
        if (1 < v2) {
            var v8 uint64 = 504
            v1 = 0.0
            v9 := GenS1{F0: 697, F1: -1, F2: false}
            _ = v8
            _ = v9
        } else {
            v10 := []int{}
            _ = v10
        }
        _ = fi5
    }
    if false {
        v11 := GenS1{F0: 1000, F1: 100, F2: true}
        v12 := []int{994, 4294967295}
        var v13 float64 = v1
        _ = v11
        _ = v12
        _ = v13
    }
    fmt.Println(v1, v1)
    fmt.Println(v2, v1)
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 string = "\n\t"
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        v4[42] = ("foo bar" + (("" + "\n\t") + "a"))
        if false {
            v17 := GenS2{F0: -1}
            v2 = v2
            var v18 string = "a"
            _ = v17
            _ = v18
        } else {
            v2 -= (v2 % ((v2) | 1))
            if false {
                v1 += (v1 + 1e10)
                v1 = ((v1 / v1) - v1)
                fmt.Println(fi16)
            } else {
                var v19 uint64 = (v2 % ((22) | 1))
                _ = v19
            }
        }
        _ = fi16
    }
    v20 := GenS2{F0: 1}
    if false {
        v21 := []int{584}
        v1 += v1
        _ = v21
    } else {
        v22 := []int{95}
        var v23 uint32 = 65535
        v22 = append(v22, 127)
        _ = v22
        _ = v23
    }
    delete(v3, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
