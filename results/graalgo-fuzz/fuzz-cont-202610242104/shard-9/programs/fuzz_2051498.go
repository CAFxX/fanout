package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool; F1 int32; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint32 }

func main() {
    var v1 int64 = 0
    v2 := make(map[string]int)
    var v3 int64 = (v1 ^ v1)
    var v4 uint16 = 0
    v5 := GenS1{F0: 2, F1: 1, F2: 42}
    fmt.Println(v1, v4, v3)
    v4 = v4
    if false {
        v1 -= (v3 / (((v3 - (v1 | v1))) | 1))
    }
    var v6 bool = false
    switch "" {
    case "a":
        var v7 uint32 = 2
        _ = v7
    case "":
        v8 := GenS2{F0: "foo bar", F1: 680}
        _ = v8
    default:
        switch "foo bar" {
        case "a":
            if (true && (!(42 < 42))) {
                if (!v6) {
                    var v9 int32 = 42
                    _ = v9
                } else {
                    switch 42 {
                    case -9223372036854775808:
                        delete(v2, "a")
                    default:
                        var v10 bool = (42 == -667)
                        _ = v10
                    }
                    var v11 uint16 = (v4 << ((65535) & 63))
                    var v12 uint32 = 991
                    _ = v11
                    _ = v12
                }
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 float64 = ((0.1 + float64(1.5)) - 3.14159)
                    _ = fi13
                    _ = v14
                }
                switch "a" {
                case "foo bar":
                    var v15 uint16 = (^(v4))
                    _ = v15
                case "a":
                    v16 := GenS3{F0: true, F1: 2147483647, F2: 65535}
                    _ = v16
                }
            } else {
                if v6 {
                    if v6 {
                        v17 := GenS4{F0: false, F1: 65535}
                        var v18 uint64 = 42
                        _ = v17
                        _ = v18
                    }
                    v1 += (v3 * v3)
                    var v19 int32 = 2147483647
                    _ = v19
                }
                v4 -= (v4 / ((255) | 1))
                v20 := []int{0, 1}
                _ = v20
            }
        case "\n\t":
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(fi21, v6)
                _ = fi21
            }
        }
    }
    v1 = -1
    v3 = (-(v3))
    var v22 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
