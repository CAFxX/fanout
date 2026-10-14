package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64, p1 bool) uint32 {
    v4 := []int{-2147483648, 32767, -9223372036854775808}
    fmt.Println(p0, p0)
    p0 += p0
    var v5 uint8 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1(p0 uint64, p1 string, p2 uint32) int32 {
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(p1, p2, fi6)
        _ = fi6
    }
    p0 += 0
    if false {
        var v7 bool = true
        var v8 uint32 = 2147483648
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    var v1 int8 = 100
    var v2 int64 = 9223372036854775807
    var v3 string = ("\n\t" + "\n\t")
    fmt.Println(v1, v3)
    var v9 uint16 = 1
    fmt.Println(v3, v1, v3)
    v1 = 100
    var v10 uint32 = 42
    v1 = (-(v1))
    v1 = (^(v1))
    v11 := GenS1{F0: 255, F1: 0, F2: v1}
    fmt.Println(v3, v10)
    if true {
        for fi12 := 0; fi12 < 2; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                v10 += v10
                var v14 uint = 255
                _ = fi13
                _ = v14
            }
            if true {
                var v15 float64 = 1e10
                _ = v15
            }
            _ = fi12
        }
    } else {
        var v16 string = ((("" + ((((v3 + ("" + "hello")) + v3) + (v3 + v3)) + ("hello" + "\n\t"))) + "hello") + "\n\t")
        v17 := GenS2{F0: 964}
        _ = v16
        _ = v17
    }
    v18 := []int{-673, -973, 2147483647}
    if true {
        if false {
            var v19 int = 128
            _ = v19
        }
        var v20 int8 = (v1 >> ((65535) & 63))
        _ = v20
    }
    v21 := GenS1{F0: 0, F1: 433, F2: (v1 ^ v1)}
    v18 = append(v18, 153)
    v10 = v10
    v2 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v9
}
