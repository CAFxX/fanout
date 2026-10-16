package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int8; F1 int; F2 int32 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 bool; F1 int64; F2 uint32 }
func genFn0() int {
    var v5 uint64 = 425
    v5 = v5
    _ = v5
    return -1
}
func genFn1() int16 {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 bool = true
        if false {
            v8 := GenS1{F0: 1}
            v9 := []int{fi6, genFn0(), 4294967295, fi6}
            v7 = v7
            _ = v8
            _ = v9
        }
        v7 = v7
        _ = fi6
        _ = v7
    }
    var v10 uint = 185
    fmt.Println(v10)
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint64 = 18446744073709551615
        v13 := GenS2{F0: "\n\t", F1: 255}
        var v14 string = "a"
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    _ = v10
    return 32767
}

func main() {
    var v1 uint16 = 42
    var v2 uint = 4294967295
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    if (2147483647 < 42) {
        v15 := []int{genFn0()}
        var v16 bool = (0 != 309)
        v2 += v2
        _ = v15
        _ = v16
    }
    switch "a" {
    case "a":
        _ = len(v3)
    case "\n\t":
        for fi17 := 0; fi17 < 5; fi17++ {
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 int32 = -1
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 int = -144
                switch v21 {
                case 42:
                    switch genFn0() {
                    case 32767:
                        v22 := []int{128}
                        _ = v22
                    }
                }
                v1 = uint16(v2)
                _ = fi20
                _ = v21
            }
            v2 += (v2 % ((v2) | 1))
            _ = fi17
        }
    default:
        v23 := GenS1{F0: 0}
        _ = v23
    }
    v1 = v1
    v24 := GenS3{F0: -1, F1: -911, F2: 42}
    v2 -= (v2 % (((^(v2))) | 1))
    var v25 uint8 = 31
    var v26 uint16 = v1
    fmt.Println(v1, v26)
    if (0 == -1) {
        v26 = 1
    } else {
        var v27 int = -333
        for fi28 := 0; fi28 < 14; fi28++ {
            fmt.Println(v25, fi28)
            v26 += 32767
            v29 := GenS4{F0: (v2 ^ (v2 / ((66) | 1))), F1: ""}
            _ = fi28
            _ = v29
        }
        v1 = v1
        _ = v27
    }
    fmt.Println(v25, v1, v2)
    switch 1 {
    case 4294967295:
        fmt.Println(v26, v25)
    case -1:
        v25 -= (v25 << ((v2) & 63))
    default:
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 uint64 = 4294967295
            v32 := GenS5{F0: true, F1: int64(v25), F2: 1}
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    v33 := []int{}
    v33 = append(v33, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v4
}
