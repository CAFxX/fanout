package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int32; F1 int }
func genFn0(p0 int64) uint {
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(fi5, p0)
        _ = fi5
    }
    var v6 int32 = 2147483647
    p0 = p0
    var v7 int8 = 33
    if false {
        p0 += p0
    }
    _ = p0
    _ = v6
    _ = v7
    return 669
}
func genFn1() uint64 {
    var v8 uint = 0
    for fi9 := 0; fi9 < 2; fi9++ {
        if (!false) {
            var v10 string = "\n\t"
            _ = v10
        }
        var v11 int16 = 1000
        v11 = 605
        _ = fi9
        _ = v11
    }
    v12 := []int{440, 142}
    v8 += 18446744073709551615
    _ = v12
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := []int{127, 127, -141, 128}
    v2 := make(map[int]uint32)
    var v3 uint16 = 42
    v4 := make(map[string]string)
    var v13 uint = 324
    var v14 uint16 = 558
    switch "\n\t" {
    case "":
        if false {
            var v15 uint32 = 42
            var v16 int8 = 127
            v14 = v14
            _ = v15
            _ = v16
        }
    case "a":
        var v17 int = 4294967295
        _ = v17
    case "hello":
        var v18 int32 = 1
        _ = v18
    }
    delete(v2, 2147483647)
    fmt.Println(v13)
    var v19 uint8 = 42
    var v20 uint16 = v3
    v21 := GenS1{F0: (v13 + 44), F1: v3, F2: 2}
    v1 = append(v1, 42)
    for fi22 := 0; fi22 < 18; fi22++ {
        fmt.Println(v14, v14)
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 bool = true
            var v25 uint64 = genFn1()
            fmt.Println(v25)
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        var v27 bool = true
        v20 -= (^(v20))
        _ = fi26
        _ = v27
    }
    v19 -= 2
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 uint8 = 164
        v30 := GenS2{F0: 253, F1: (fi28 << ((v13) & 63))}
        _ = fi28
        _ = v29
        _ = v30
    }
    if true {
        fmt.Println(v19, v20)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
