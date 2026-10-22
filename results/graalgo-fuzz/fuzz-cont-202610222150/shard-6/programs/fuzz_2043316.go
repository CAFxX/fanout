package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int32; F2 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint32 }
func genFn0() uint64 {
    var v5 uint64 = 307
    var v6 string = "a"
    v7 := GenS1{F0: 760, F1: -1}
    _ = v5
    _ = v6
    _ = v7
    return v5
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]int)
    v3 := []string{"\n\t", "", (("\n\t" + "foo bar") + "a"), ""}
    v4 := []uint32{951, 463}
    v8 := GenS2{F0: 4294967295}
    switch "\n\t" {
    case "foo bar":
        var v9 float32 = 1e10
        _ = v9
    case "hello":
        v10 := GenS2{F0: 1}
        _ = v10
    case "":
        v11 := v2[2]
        _ = v11
    default:
        switch "a" {
        case "foo bar":
            var v12 uint16 = 2
            _ = v12
        case "a":
            v13 := GenS2{F0: 4294967295}
            _ = v13
        default:
            fmt.Println(v1)
        }
    }
    v1 -= (v1 % (((v1 % ((v1) | 1))) | 1))
    var v14 uint16 = 255
    if (42 != 53) {
        v15 := GenS3{F0: "hello"}
        _ = v15
    }
    var v16 uint64 = 1
    fmt.Println(v14, v14)
    fmt.Println(v14, v14, v16)
    v17 := GenS4{F0: (v16 >> ((675) & 63)), F1: -1, F2: false}
    var v18 bool = (1 > 65535)
    v19 := GenS5{F0: 2147483647}
    v20 := GenS5{F0: 2147483647}
    v21 := GenS6{F0: 49, F1: 450}
    v16 -= 993
    fmt.Println(v14, v14, v18)
    delete(v2, 30)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v8
}
