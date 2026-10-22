package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int }
func genFn0(p0 int64, p1 string, p2 uint32) int64 {
    p2 = p2
    p2 -= (^(p2))
    var v4 uint = 877
    v5 := GenS1{F0: 42, F1: 0}
    var v6 uint64 = 362
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (^(p0))
}
func genFn1(p0 uint16, p1 uint, p2 int) string {
    var v7 uint16 = p0
    v8 := GenS1{F0: 1, F1: uint16(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    var v1 uint = 1
    v2 := []uint32{}
    var v3 int = 128
    switch ("a" + "a") {
    case "":
        v9 := GenS2{F0: 255, F1: 42, F2: 1}
        _ = v9
    default:
        _ = len(v2)
    }
    v3 -= (v3 + (v3 ^ 4294967295))
    v1 = (^(v1))
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v3 = v3
    fmt.Println(v3, v3)
    switch v3 {
    case -919:
        v3 -= (v3 << (((v1 - v1)) & 63))
    case 42:
        fmt.Println(v3, v3)
    case 496:
        for fi10 := 0; fi10 < 4; fi10++ {
            if false {
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                var v11 string = "hello"
                fmt.Println(v11)
                _ = v11
            }
            v1 = (^(v1))
            _ = fi10
        }
    default:
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS3{F0: fi12}
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS1{F0: -457, F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
