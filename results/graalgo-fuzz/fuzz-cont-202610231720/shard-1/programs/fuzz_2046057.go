package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }
func genFn0() uint64 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 bool = (42 < 42)
        v7 := []int{(fi5 + (fi5 % ((fi5) | 1))), (^(fi5)), fi5, -2147483648}
        for fi8 := 0; fi8 < 17; fi8++ {
            v6 = (0 > 151)
            v6 = false
            if len(v7) > 0 {
                v7[0] = -2147483648
            }
            v6 = false
            switch -558 {
            case 127:
                var v9 int64 = 1
                _ = v9
            case 35:
                var v10 uint8 = 255
                _ = v10
            }
            _ = fi8
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v11 float32 = -2.25
    v12 := GenS1{F0: -147, F1: "\n\t", F2: true}
    _ = v11
    _ = v12
    return 873
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 int64 = -716
    var v4 uint32 = 4294967295
    v2["hello"] = 4294967295
    fmt.Println(v3, v4)
    for fi13 := 0; fi13 < 2; fi13++ {
        v3 += v3
        var v14 int16 = 1000
        if true {
            v15 := GenS1{F0: (^(v14)), F1: "a", F2: false}
            _ = v15
        } else {
            v16 := GenS1{F0: int16(v4), F1: "hello", F2: (!false)}
            var v17 string = "foo bar"
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    var v18 float64 = (3.14159 + float64((1e10 * float32(-2.25))))
    if false {
        v3 += (-(v3))
        var v19 uint = 212
        v19 -= v19
        _ = v19
    }
    v3 = (v3 * v3)
    if false {
        fmt.Println(v4)
    } else {
        var v20 float32 = ((-2.25 - (1.5 - (0.0 / 3.14159))) * 1.5)
        _ = len(v1)
        _ = v20
    }
    v18 -= v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
