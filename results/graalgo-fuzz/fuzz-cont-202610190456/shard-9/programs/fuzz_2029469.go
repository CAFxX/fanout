package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int64 }
func genFn0() int8 {
    var v7 uint32 = 4294967295
    for fi8 := 0; fi8 < 4; fi8++ {
        if false {
            v9 := GenS1{F0: 593, F1: 2, F2: 9223372036854775807}
            _ = v9
        } else {
            var v10 int64 = 0
            switch (fi8 ^ -924) {
            case 2147483647:
                var v11 int32 = 42
                _ = v11
            default:
                for fi12 := 0; fi12 < 14; fi12++ {
                    v10 -= v10
                    v13 := []int{(fi12 + (^(fi8))), fi12, 32767}
                    fmt.Println(fi8)
                    _ = fi12
                    _ = v13
                }
            }
            var v14 int16 = -470
            _ = v10
            _ = v14
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 uint32 = 2
            _ = fi15
            _ = v16
        }
        _ = fi8
    }
    _ = v7
    return 1
}
func genFn1(p0 int32, p1 uint64) int32 {
    var v17 int8 = -128
    if true {
        v18 := GenS2{F0: 2147483648}
        _ = v18
    }
    v17 -= (^(v17))
    fmt.Println(v17, p1)
    _ = p0
    _ = p1
    _ = v17
    return (p0 - (p0 ^ 56))
}

func main() {
    var v1 uint16 = 42
    var v2 uint64 = 3
    var v3 int = 508
    var v4 string = ("a" + "a")
    var v5 float32 = 3.14159
    var v6 uint16 = (v1 | (v1 - v1))
    v19 := GenS2{F0: 1}
    v20 := GenS3{F0: (v3 & (v3 & (v3 << ((4294967295) & 63)))), F1: -1}
    v21 := []int{(v3 / ((v3) | 1)), v3, v3, v3}
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 bool = false
        var v24 int8 = -89
        v6 = (v6 << ((965) & 63))
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int = (v3 ^ 2147483647)
    v26 := GenS2{F0: 4294967295}
    var v27 int = v25
    var v28 int32 = 0
    var v29 uint64 = (v2 | (v2 % ((199) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
