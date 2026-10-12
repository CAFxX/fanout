package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 string }
func genFn0(p0 int16, p1 int16) uint32 {
    var v5 int8 = 100
    if ((166 <= 127) || true) {
        if false {
            v6 := []int{}
            p1 -= 1
            _ = v6
        }
        v7 := GenS1{F0: 230}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint = 546
    var v2 int32 = 915
    var v3 string = ("a" + "")
    v4 := make(map[string]int)
    var v8 uint = (^(v1))
    switch 42 {
    case 386:
        v8 = (v1 >> ((v1) & 63))
    }
    v9 := GenS1{F0: 1}
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int8 = 0
        v12 := GenS2{F0: v1}
        _ = fi10
        _ = v11
        _ = v12
    }
    if (2147483648 > 654) {
        fmt.Println(v1, v1, v1)
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v1 = 593
        switch fi13 {
        case 32767:
            fmt.Println(v2, v2, v2)
        case 9223372036854775807:
            var v14 uint = 823
            _ = v14
        default:
            var v15 int16 = 0
            _ = v15
        }
        if ((false && ((2 == 1) && false)) || false) {
            v16 := []int{(fi13 * fi13), (fi13 >> ((422) & 63)), (^(fi13))}
            _ = v16
        }
        _ = fi13
    }
    v1 = v8
    fmt.Println(v3, v1)
    _ = len(v4)
    v17 := []int{972, 976, 2147483647}
    v18 := GenS3{F0: -2147483648}
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS4{F0: v3, F1: v3}
        if len(v17) > 0 {
            v17[0] = fi19
        }
        _ = fi19
        _ = v20
    }
    v3 = v3
    switch 128 {
    case -272:
        for fi21 := 0; fi21 < 12; fi21++ {
            for fi22 := 0; fi22 < 11; fi22++ {
                if len(v17) > 0 {
                    v17[0] = fi22
                }
                _ = fi22
            }
            v3 = (((v3 + "") + v3) + v3)
            v3 = "foo bar"
            _ = fi21
        }
    }
    v3 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
