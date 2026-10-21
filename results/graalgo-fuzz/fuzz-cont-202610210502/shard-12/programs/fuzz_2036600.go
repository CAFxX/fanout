package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int8; F2 int }
func genFn0(p0 uint32, p1 int8) int16 {
    var v7 bool = ((!(-1 == 730)) && true)
    var v8 uint16 = 255
    fmt.Println(v7)
    var v9 uint = 42
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -791
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 uint = 255
    v4 := make(map[uint32]bool)
    v5 := []uint32{}
    v6 := make(map[int]uint32)
    fmt.Println(v3)
    if (false && ((!true) || true)) {
        if true {
            for fi10 := 0; fi10 < 12; fi10++ {
                v3 = 4294967295
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi10
            }
            var v11 int32 = 42
            _ = v11
        }
    }
    v3 += 2
    if ((!true) || true) {
        var v12 uint16 = 735
        v13 := GenS1{F0: ((!((3.14159 * 3.14159) >= 1e10)) || true), F1: 1}
        for fi14 := 0; fi14 < 20; fi14++ {
            _ = len(v5)
            _ = fi14
        }
        _ = v12
        _ = v13
    }
    v15 := GenS1{F0: true, F1: 4294967295}
    var v16 float32 = -2.25
    v17 := GenS2{F0: 458, F1: -34, F2: -1}
    var v18 int64 = 1
    v19 := GenS2{F0: 18446744073709551615, F1: 127, F2: -373}
    v18 = v18
    if false {
        v5 = append(v5, 42)
        var v20 uint8 = 1
        v3 -= 943
        _ = v20
    }
    v16 -= float32(((float64(float32(1.5)) * 0.0) - 0.0))
    var v21 uint16 = 0
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
