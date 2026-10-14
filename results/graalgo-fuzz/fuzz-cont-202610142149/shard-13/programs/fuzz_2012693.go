package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS2 struct { F0 int64 }
func genFn0() int {
    var v7 uint16 = 1
    fmt.Println(v7)
    var v8 int64 = 9223372036854775807
    var v9 int64 = (v8 ^ (^(v8)))
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 uint64) int {
    p0 -= p0
    v10 := []int{}
    _ = p0
    _ = v10
    return genFn0()
}

func main() {
    var v1 int16 = 1000
    var v2 float32 = 3.14159
    var v3 int8 = 100
    v4 := make(map[string]bool)
    var v5 uint32 = 1
    var v6 int8 = (v3 * v3)
    if false {
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 uint = 255
            _ = fi11
            _ = v12
        }
        var v13 int32 = -2147483648
        if false {
            v3 += (v3 / (((-(v6))) | 1))
            switch -216 {
            case 2147483647:
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 float32 = ((float32((0.1 / 0.1)) + ((v2 * (3.14159 / v2)) - float32(-2.25))) / -2.25)
                    _ = fi14
                    _ = v15
                }
            }
        } else {
            for fi16 := 0; fi16 < 9; fi16++ {
                v6 -= (^(v6))
                _ = fi16
            }
        }
        _ = v13
    } else {
        v17 := GenS1{F0: (^(v5)), F1: true, F2: 2}
        _ = v17
    }
    v18 := GenS2{F0: 529}
    var v19 uint16 = 255
    v1 -= -32768
    var v20 float64 = (3.14159 + 1e10)
    var v21 uint16 = v19
    if true {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 int = fi22
            v24 := GenS2{F0: int64(v21)}
            v19 = (v19 | 42)
            _ = fi22
            _ = v23
            _ = v24
        }
        v3 -= -128
        fmt.Println(v21)
    } else {
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 uint32 = v5
            v27 := GenS1{F0: (v26 * 786), F1: (919 <= 55), F2: v19}
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
