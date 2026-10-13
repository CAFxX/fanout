package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint64 = 0
    var v3 uint64 = (^(v2))
    v4 := make(map[int]uint32)
    v5 := []int{0, 42, 4294967295}
    if true {
        var v6 float32 = 1.5
        var v7 int = 127
        _ = v6
        _ = v7
    }
    var v8 int = 641
    fmt.Println(v1)
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 float32 = (float32(0.1) + 1e10)
        _ = fi9
        _ = v10
    }
    v8 = v8
    if false {
        for fi11 := 0; fi11 < 9; fi11++ {
            v5 = append(v5, -9223372036854775808)
            _ = fi11
        }
        if (false && true) {
            v12 := GenS1{F0: 32767, F1: 1}
            fmt.Println(v3, v1)
            var v13 uint16 = 255
            _ = v12
            _ = v13
        }
    } else {
        var v14 uint32 = 810
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 17; fi16++ {
                if ((((!(!false)) && true) || (!(!(32767 >= -32768)))) || (v1 > ((v1 / 3.14159) / float64(3.14159)))) {
                    fmt.Println(v1, v8)
                }
                switch (^(fi16)) {
                case -9223372036854775808:
                    fmt.Println(fi16, v14, v8)
                }
                _ = fi16
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi15
        }
        _ = v14
    }
    v3 = (v2 - v3)
    var v17 int32 = 1
    fmt.Println(v17, v1, v17)
    var v18 uint64 = v3
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS1{F0: 0, F1: 255}
        var v21 int = (-(v8))
        delete(v4, (-(v8)))
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 32767, F1: 909}
    var v23 float32 = 1e10
    v17 -= (v17 & v17)
    v23 -= -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
