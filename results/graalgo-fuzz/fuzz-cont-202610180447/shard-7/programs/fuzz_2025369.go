package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
func genFn0() uint64 {
    v5 := GenS1{F0: 1, F1: (-566 > 32767), F2: 128}
    v6 := GenS1{F0: 42, F1: (128 < 1), F2: 157}
    var v7 int = 1
    if false {
        var v8 int64 = -1
        v7 += (v7 << ((299) & 63))
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    _ = v5
    _ = v6
    _ = v7
    _ = v9
    return (v9 / ((248) | 1))
}

func main() {
    var v1 int = 32767
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 2
    var v4 int32 = -1
    switch "a" {
    case "foo bar":
        v4 = 42
    case "":
        fmt.Println(v3, v3)
    case "a":
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := []int{}
            _ = fi10
            _ = v11
        }
    default:
        v12 := GenS2{F0: v3}
        _ = v12
    }
    v13 := GenS3{F0: 42}
    var v14 string = "foo bar"
    if false {
        fmt.Println(v1, v4)
    } else {
        var v15 int = (v1 | 1)
        for fi16 := 0; fi16 < 6; fi16++ {
            v17 := GenS4{F0: ""}
            v4 = -1
            _ = fi16
            _ = v17
        }
        v3 = v3
        _ = v15
    }
    delete(v2, 42)
    switch v14 {
    case "a":
        var v18 float32 = 0.0
        _ = v18
    case "":
        v4 -= (v4 % (((v4 << ((261) & 63))) | 1))
    case "foo bar":
        if false {
            var v19 uint64 = v3
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 float64 = 3.14159
                v22 := GenS3{F0: 42}
                v23 := GenS5{F0: v4}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            if false {
                var v24 int16 = 32767
                _ = v24
            }
            _ = v19
        }
    default:
        v25 := GenS3{F0: -483}
        _ = v25
    }
    v3 = v3
    v26 := []int{v1, 42}
    v4 = (v4 / ((v4) | 1))
    if (!true) {
        v14 = ("\n\t" + (((v14 + v14) + "a") + ("foo bar" + (("" + v14) + "foo bar"))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
