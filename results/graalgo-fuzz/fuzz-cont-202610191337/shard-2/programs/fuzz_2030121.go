package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int, p2 uint8) int64 {
    v4 := []int{(p1 / ((p1) | 1)), (p1 << ((4294967295) & 63))}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    var v1 float32 = (1.5 * 3.14159)
    var v2 uint8 = 255
    var v3 uint16 = 810
    var v5 uint64 = 569
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(v3, v3)
        _ = fi6
    }
    var v7 int8 = 1
    var v8 uint32 = uint32(v3)
    v7 = v7
    var v9 int64 = -9223372036854775808
    v9 = (v9 * int64(v7))
    var v10 string = "hello"
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 uint = 2
        _ = fi11
        _ = v12
    }
    switch "" {
    case "\n\t":
        v13 := GenS1{F0: (v3 % (((v3 ^ v3)) | 1))}
        _ = v13
    case "":
        v14 := GenS1{F0: 945}
        _ = v14
    case "hello":
        v2 -= uint8(v5)
    }
    v15 := GenS2{F0: -466, F1: v7}
    var v16 string = v10
    v17 := GenS3{F0: 42, F1: v8, F2: v2}
    if false {
        v18 := []int{}
        _ = v18
    }
    var v19 string = "foo bar"
    v20 := GenS2{F0: 507, F1: (v7 ^ (v7 / (((v7 - v7)) | 1)))}
    v21 := GenS4{F0: v5}
    fmt.Println(v3, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
