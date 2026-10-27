package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint; F2 int16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int64, p1 int32, p2 int8) int32 {
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := []int{-881}
        var v7 bool = true
        _ = fi5
        _ = v6
        _ = v7
    }
    if ((-2.25 / 1e10) >= 1.5) {
        fmt.Println(p0)
        fmt.Println(p1)
        var v8 uint16 = 21
        _ = v8
    } else {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{}
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS1{F0: 42, F1: fi9, F2: 139}
                v10 = append(v10, (fi11 << ((849) & 63)))
                if true {
                    fmt.Println(p0, p0, p2)
                }
                _ = fi11
                _ = v12
            }
            p0 += p0
            _ = fi9
            _ = v10
        }
        v13 := GenS2{F0: 1}
        _ = v13
    }
    switch "" {
    case "\n\t":
        v14 := GenS2{F0: p0}
        _ = v14
    case "hello":
        fmt.Println(p0, p1, p0)
    default:
        p1 += 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 + 290)
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 0
    var v3 uint64 = 2
    v4 := []uint32{}
    fmt.Println(v3, v2)
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS3{F0: 988, F1: 2, F2: 723}
        _ = fi15
        _ = v16
    }
    v3 += uint64(v2)
    v17 := GenS3{F0: 2, F1: 42, F2: 635}
    v18 := GenS4{F0: 0}
    var v19 int32 = -2147483648
    if true {
        var v20 uint16 = uint16(v2)
        _ = v20
    }
    fmt.Println(v3, v19, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
