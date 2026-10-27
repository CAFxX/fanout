package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 bool, p1 int64) uint8 {
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(p0, p0, fi6)
        var v7 int64 = (p1 / (((p1 ^ p1)) | 1))
        if true {
            if p0 {
                fmt.Println(v7, fi6, p1)
                p0 = (65535 >= 42)
            }
        }
        _ = fi6
        _ = v7
    }
    var v8 uint16 = 2
    _ = p0
    _ = p1
    _ = v8
    return 209
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[uint32]int)
    v3 := []int{-950}
    var v4 string = (("" + "foo bar") + "")
    var v5 uint8 = 128
    var v9 int32 = 2147483647
    var v10 uint16 = 42
    v2[0] = -276
    var v11 uint8 = (^(v5))
    fmt.Println(v9, v9, v5)
    if false {
        _ = len(v3)
    }
    fmt.Println(v9, v5)
    switch (v4 + "\n\t") {
    case "a":
        var v12 int8 = 100
        _ = v12
    }
    switch ("a" + (v4 + v4)) {
    case "hello":
        v1 = 0.1
    case "a":
        if true {
            v11 += v5
            v13 := v2[591]
            _ = v13
        } else {
            v9 += 497
            switch "" {
            case "a":
                var v14 int32 = (^(v9))
                _ = v14
            case "hello":
                if false {
                    var v15 uint64 = 42
                    v16 := v2[0]
                    _ = v15
                    _ = v16
                }
            default:
                var v17 float32 = v1
                _ = v17
            }
            v2[42] = -837
        }
    default:
        v10 = 851
    }
    var v18 int64 = 301
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS1{F0: (-433 <= 32767)}
        _ = fi19
        _ = v20
    }
    var v21 int8 = 1
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 9; fi23++ {
            v24 := GenS1{F0: false}
            v25 := GenS1{F0: (1 == 1)}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    v26 := GenS1{F0: (((v1 > v1) && false) || (!true))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
