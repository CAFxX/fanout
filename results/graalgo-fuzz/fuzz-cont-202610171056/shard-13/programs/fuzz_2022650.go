package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint }
func genFn0(p0 int) int16 {
    p0 += p0
    var v6 uint64 = 2
    fmt.Println(v6)
    var v7 int32 = 42
    _ = p0
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int32) uint8 {
    v8 := GenS1{F0: 961, F1: 886}
    switch 32767 {
    case 32767:
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 bool = (((1.5 < -2.25) || true) || (42 < 0))
            v10 = ((42 >= 341) || v10)
            _ = fi9
            _ = v10
        }
    case 42:
        v11 := GenS1{F0: 0, F1: 36}
        _ = v11
    case 153:
        if true {
            p0 -= p0
        } else {
            var v12 uint8 = 255
            var v13 int64 = -9223372036854775808
            if true {
                v14 := []int{}
                var v15 uint = 65535
                for fi16 := 0; fi16 < 11; fi16++ {
                    v17 := GenS2{F0: 42, F1: 2}
                    _ = fi16
                    _ = v17
                }
                _ = v14
                _ = v15
            }
            _ = v12
            _ = v13
        }
    default:
        p0 = -282
    }
    switch "\n\t" {
    case "":
        var v18 bool = (!true)
        _ = v18
    }
    v19 := []int{4294967295, -764}
    switch -2147483648 {
    case 128:
        fmt.Println(p0)
    case 9223372036854775807:
        var v20 uint8 = 0
        _ = v20
    case 32767:
        p0 = (p0 << ((1) & 63))
    default:
        p0 = (p0 >> ((255) & 63))
    }
    _ = p0
    _ = v19
    _ = v8
    return 127
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[uint32]int)
    var v3 uint16 = 0
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    var v21 uint8 = 0
    var v22 int32 = 1
    fmt.Println(v22)
    v22 = v22
    v22 += (v22 >> ((1) & 63))
    for fi23 := 0; fi23 < 19; fi23++ {
        v24 := GenS1{F0: 9223372036854775807, F1: 0}
        _ = fi23
        _ = v24
    }
    fmt.Println(v1)
    v25 := GenS2{F0: -568, F1: uint(v3)}
    for fi26 := 0; fi26 < 19; fi26++ {
        var v27 int = 32768
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
