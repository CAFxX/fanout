package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint32 }
func genFn0(p0 uint16, p1 int) bool {
    switch -501 {
    case 851:
        var v6 int32 = 0
        _ = v6
    default:
        v7 := GenS1{F0: (p1 < 2147483647)}
        _ = v7
    }
    var v8 uint16 = p0
    var v9 uint16 = (^(p0))
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (362 <= 2147483648)
}

func main() {
    var v1 uint8 = 42
    var v2 uint8 = v1
    var v3 uint = 65535
    var v4 float64 = float64(0.1)
    var v5 string = ""
    var v10 int8 = 1
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := []int{(fi11 << (((v3 - v3)) & 63)), -97}
            v14 := GenS2{F0: 1}
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    var v15 int8 = (v10 % (((v10 & v10)) | 1))
    fmt.Println(v5, v15, v15)
    v16 := GenS2{F0: -1}
    var v17 uint64 = 831
    v18 := []int{42, 2147483647}
    var v19 string = v5
    for fi20 := 0; fi20 < 5; fi20++ {
        v17 = 4294967295
        _ = fi20
    }
    v21 := GenS3{F0: 2, F1: 32767, F2: (^(v17))}
    var v22 uint32 = 2147483648
    v23 := GenS4{F0: (^(v2)), F1: v22, F2: 42}
    v24 := GenS5{F0: uint32(v1), F1: v22, F2: uint32(v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
