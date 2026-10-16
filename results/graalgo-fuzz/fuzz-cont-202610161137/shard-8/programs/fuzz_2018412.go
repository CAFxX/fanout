package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint; F1 int; F2 bool }
type GenS3 struct { F0 string; F1 int8; F2 int16 }
type GenS4 struct { F0 string; F1 int8; F2 uint }
func genFn0(p0 string, p1 uint64, p2 uint64) uint8 {
    var v5 int = 0
    fmt.Println(v5, p2, p1)
    var v6 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 49
}

func main() {
    var v1 uint16 = 42
    var v2 string = "a"
    v3 := make(map[uint32]bool)
    v4 := []uint32{274, 4294967295}
    v2 = ("a" + v2)
    if false {
        if (true || false) {
            var v7 int64 = 9223372036854775807
            var v8 int64 = 798
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: (v1 * v1), F1: 100}
        fmt.Println(v1)
        _ = v9
    } else {
        v10 := GenS2{F0: 65535, F1: -9223372036854775808, F2: (false || true)}
        _ = v10
    }
    v2 = "hello"
    fmt.Println(v2, v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1)
    if true {
        if (2 <= 254) {
            if ((false && (!(-512 <= 42))) || (1e10 != ((-2.25 + 1.5) + -2.25))) {
                var v11 uint8 = 90
                _ = v11
            } else {
                fmt.Println(v1, v1)
                v12 := GenS3{F0: (v2 + (v2 + v2)), F1: -1, F2: 1}
                for fi13 := 0; fi13 < 6; fi13++ {
                    switch (v2 + v2) {
                    case "hello":
                        v1 = (v1 / (((^(v1))) | 1))
                    case "a":
                        var v14 float32 = 0.1
                        _ = v14
                    }
                    var v15 string = "hello"
                    _ = fi13
                    _ = v15
                }
                _ = v12
            }
            var v16 float32 = 0.1
            switch v2 {
            case "foo bar":
                fmt.Println(v1, v1)
            default:
                fmt.Println(v2, v16, v2)
            }
            _ = v16
        }
    }
    v17 := GenS4{F0: "\n\t", F1: -89, F2: 811}
    var v18 float32 = 0.1
    fmt.Println(v1, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
