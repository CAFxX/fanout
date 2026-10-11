package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint16; F1 bool; F2 int }
type GenS4 struct { F0 int; F1 uint16; F2 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 uint; F2 int16 }
func genFn0(p0 int8, p1 uint64) uint32 {
    v5 := GenS1{F0: 32767, F1: -770, F2: -2147483648}
    p1 = p1
    p0 -= 127
    var v6 uint = 565
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2147483648
}

func main() {
    v1 := []uint32{42, 621, 0, 2147483648}
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v4 int32 = 42
    var v7 uint8 = 2
    var v8 int32 = (v4 ^ 2147483647)
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS2{F0: 42, F1: v8}
        if true {
            v7 = v7
        }
        _ = fi9
        _ = v10
    }
    v11 := GenS3{F0: 65535, F1: ((-2.25 / 1.0) < ((-2.25 / -2.25) / 1e10)), F2: 4294967295}
    switch "foo bar" {
    case "":
        v12 := GenS4{F0: 32767, F1: 1, F2: false}
        _ = v12
    case "hello":
        var v13 int64 = 1
        _ = v13
    case "a":
        for fi14 := 0; fi14 < 2; fi14++ {
            switch fi14 {
            case 4294967295:
                for fi15 := 0; fi15 < 5; fi15++ {
                    v4 = 42
                    fmt.Println(v7, fi15, fi15)
                    v4 -= 42
                    _ = fi15
                }
            case 282:
                v8 += 446
            }
            _ = fi14
        }
    default:
        var v16 bool = false
        _ = v16
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v17 := GenS2{F0: 4294967295, F1: -2147483648}
    v8 += (v8 & (v4 << ((4294967295) & 63)))
    fmt.Println(v8)
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int = fi18
        _ = fi18
        _ = v19
    }
    v20 := GenS5{F0: 18446744073709551615}
    v21 := GenS6{F0: 127, F1: 18446744073709551615, F2: 32767}
    fmt.Println(v8, v4, v7)
    v7 += v7
    var v22 float32 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
