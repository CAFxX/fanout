package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int16; F2 uint64 }
func genFn0(p0 int64, p1 uint16, p2 bool) uint16 {
    v5 := GenS1{F0: -788, F1: 45, F2: 1000}
    for fi6 := 0; fi6 < 20; fi6++ {
        p2 = (49 != -1)
        fmt.Println(p1, p1)
        _ = fi6
    }
    if p2 {
        v7 := []int{-429, -167}
        _ = v7
    } else {
        var v8 string = "a"
        var v9 uint64 = 42
        _ = v8
        _ = v9
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}
func genFn1(p0 string) int {
    if (42 <= -783) {
        var v10 int = -1
        fmt.Println(v10)
        _ = v10
    }
    v11 := []int{}
    var v12 int32 = 1
    if ((42 < 589) && (2147483648 <= 65535)) {
        switch 32767 {
        case 4294967295:
            v13 := GenS1{F0: v12, F1: 0, F2: 32767}
            _ = v13
        case 127:
            _ = len(v11)
        }
        p0 = "a"
        if true {
            var v14 uint32 = 42
            v11 = append(v11, -2147483648)
            _ = v14
        } else {
            var v15 uint16 = 0
            _ = v15
        }
    }
    _ = p0
    _ = v11
    _ = v12
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 2147483647
    var v3 int16 = 0
    v4 := make(map[string]int)
    for fi16 := 0; fi16 < 14; fi16++ {
        v17 := GenS2{F0: "a"}
        v18 := GenS1{F0: 42, F1: 127, F2: v3}
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v2, v3)
    switch 32767 {
    case -119:
        v3 = v3
    case 128:
        fmt.Println(v2, v2)
    default:
        var v19 int16 = v3
        _ = v19
    }
    v20 := GenS3{F0: 65535, F1: 1000, F2: 1}
    var v21 int32 = -695
    var v22 uint32 = v2
    v21 -= (v21 / (((-(v21))) | 1))
    v22 = 4294967295
    v2 -= v2
    delete(v1, "\n\t")
    fmt.Println(v2)
    v3 = (v3 & 32767)
    switch "hello" {
    case "hello":
        fmt.Println(v2, v2, v22)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
