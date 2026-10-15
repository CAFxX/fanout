package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint }
func genFn0(p0 int8, p1 int64) uint64 {
    switch 32768 {
    case -972:
        fmt.Println(p1, p0)
    }
    p0 = -128
    var v5 int64 = (p1 - 566)
    if true {
        if (!((v5 - v5) == (^(p1)))) {
            for fi6 := 0; fi6 < 7; fi6++ {
                var v7 string = "a"
                var v8 bool = (!true)
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 int32 = 533
            _ = v9
        }
        var v10 uint64 = 154
        v11 := []int{-9223372036854775808}
        _ = v10
        _ = v11
    }
    switch 1 {
    case 0:
        var v12 uint64 = 326
        _ = v12
    case -1:
        v13 := []int{}
        _ = v13
    default:
        p0 -= (-(p0))
    }
    _ = p0
    _ = p1
    _ = v5
    return 706
}
func genFn1(p0 uint8) uint {
    var v14 uint8 = (p0 - (^(p0)))
    var v15 int = -458
    v16 := GenS1{F0: "\n\t", F1: (false || true), F2: 2147483647}
    v14 -= (p0 << ((509) & 63))
    v14 += v14
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return 4294967295
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    var v3 int = 128
    v4 := make(map[string]string)
    v1["hello"] = 2147483647
    v3 += 968
    for fi17 := 0; fi17 < 2; fi17++ {
        v3 -= (v3 & (^(fi17)))
        _ = fi17
    }
    v3 -= (v3 | v3)
    v3 -= (v3 ^ -426)
    switch "" {
    case "hello":
        v3 = v3
    case "a":
        v3 = (v3 - 32767)
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := GenS2{F0: "hello"}
        switch fi18 {
        case 32768:
            v20 := GenS3{F0: 811, F1: 435}
            _ = v20
        }
        fmt.Println(v3)
        _ = fi18
        _ = v19
    }
    v21 := GenS4{F0: 4294967295}
    v22 := []int{}
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
