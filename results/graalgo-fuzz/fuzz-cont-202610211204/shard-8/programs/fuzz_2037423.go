package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int32; F2 bool }
func genFn0(p0 int64, p1 int64) bool {
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: 586, F1: 1, F2: 65535}
        for fi6 := 0; fi6 < 12; fi6++ {
            switch -494 {
            case 2147483647:
                var v7 bool = (!true)
                _ = v7
            case 591:
                var v8 uint64 = 832
                p1 = (p0 << ((617) & 63))
                v8 = (^(v8))
                _ = v8
            default:
                fmt.Println(fi6)
            }
            _ = fi6
        }
        p0 -= 0
        _ = fi4
        _ = v5
    }
    v9 := []int{2147483647, 0, 1}
    var v10 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (float32(1.5) <= (float32(v10) / 1.5))
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    fmt.Println(v1)
    fmt.Println(v1)
    if (2147483647 > 0) {
        v11 := GenS1{F0: 461, F1: 127, F2: 2147483647}
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 uint8 = 178
            fmt.Println(v1, v1, fi12)
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    if false {
        v1 = v1
        if false {
            var v14 uint64 = v1
            _ = v14
        } else {
            v1 += 4294967295
        }
        switch ("" + (("" + ("hello" + "")) + "\n\t")) {
        case "a":
            fmt.Println(v1)
        case "":
            v15 := []int{4294967295, -9223372036854775808, 261}
            _ = v15
        case "\n\t":
            fmt.Println(v1)
        }
    }
    if false {
        if (!false) {
            v1 += v1
            v16 := GenS1{F0: 4294967295, F1: 42, F2: 2147483648}
            v17 := []int{0, -248, -577, -1}
            _ = v16
            _ = v17
        }
        v1 += (v1 >> ((240) & 63))
    } else {
        v1 = (v1 | 4294967295)
    }
    fmt.Println(v1)
    var v18 int = 2147483647
    if ("a" >= "\n\t") {
        v19 := GenS2{F0: 18446744073709551615}
        var v20 uint = 255
        _ = v19
        _ = v20
    }
    v18 -= (v18 % ((v18) | 1))
    for fi21 := 0; fi21 < 9; fi21++ {
        fmt.Println(v1, fi21, fi21)
        switch "foo bar" {
        case "a":
            var v22 uint16 = 747
            _ = v22
        case "foo bar":
            for fi23 := 0; fi23 < 13; fi23++ {
                v1 += (^(v1))
                _ = fi23
            }
        default:
            v18 = 498
        }
        v24 := GenS3{F0: (fi21 * 32767), F1: -2147483648, F2: true}
        _ = fi21
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
}
