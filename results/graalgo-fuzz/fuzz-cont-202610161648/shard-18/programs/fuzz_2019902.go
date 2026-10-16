package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int8; F1 uint64; F2 int32 }
func genFn0(p0 uint64) uint8 {
    p0 += p0
    var v4 string = ""
    v5 := GenS1{F0: 1}
    _ = p0
    _ = v4
    _ = v5
    return 236
}
func genFn1() int32 {
    var v6 int32 = 18
    var v7 int64 = 42
    if false {
        for fi8 := 0; fi8 < 17; fi8++ {
            v7 += v7
            _ = fi8
        }
        v6 += (v6 >> ((611) & 63))
    }
    var v9 int = 2147483647
    v10 := []int{128, -1}
    _ = v10
    _ = v6
    _ = v7
    _ = v9
    return (^(v6))
}

func main() {
    var v1 uint8 = 202
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    v1 -= genFn0(895)
    fmt.Println(v1)
    if true {
        v11 := GenS1{F0: 10}
        v12 := []int{}
        v13 := GenS2{F0: (v1 ^ (^(v1))), F1: ("foo bar" + ("hello" + ("hello" + "\n\t")))}
        _ = v11
        _ = v12
        _ = v13
    }
    v3["a"] = 813
    switch "" {
    case "hello":
        fmt.Println(v1)
    case "":
        v1 = 128
    }
    if (-49 > -37) {
        v14 := GenS3{F0: 13, F1: 995, F2: 1}
        _ = v14
    }
    var v15 string = ""
    v1 -= 1
    v1 -= v1
    v15 = v15
    if (!true) {
        var v16 uint8 = 193
        _ = v16
    }
    delete(v3, ("a" + v15))
    var v17 int32 = 2147483647
    var v18 int16 = 42
    var v19 string = v15
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
