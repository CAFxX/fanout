package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int8; F1 bool }
func genFn0(p0 uint8) int8 {
    v5 := GenS1{F0: 858}
    v6 := GenS1{F0: 42}
    v7 := []int{32767, 128, -2147483648}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 100
}
func genFn1(p0 int8, p1 int64, p2 uint) uint16 {
    v8 := []int{}
    switch 128 {
    case -2147483648:
        p1 += p1
    case -710:
        v9 := GenS2{F0: 1}
        _ = v9
    case -405:
        p1 += p1
    default:
        v8 = append(v8, 127)
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        p1 -= p1
        _ = fi10
    }
    var v11 int64 = (p1 | p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return 65535
}

func main() {
    var v1 int32 = 1
    var v2 int = 128
    v3 := []string{"hello"}
    v4 := make(map[string]int)
    switch "hello" {
    case "hello":
        v12 := GenS2{F0: 139}
        _ = v12
    case "a":
        if (42 != 2) {
            v2 -= v2
            if ((true || (false || (-32768 <= 32767))) || (2 <= 668)) {
                v13 := GenS2{F0: 282}
                _ = v13
            } else {
                v14 := GenS3{F0: 1, F1: 65535}
                var v15 string = "foo bar"
                _ = v14
                _ = v15
            }
        } else {
            v2 = (v2 ^ 4294967295)
            var v16 int = 0
            _ = v16
        }
    case "":
        v1 += (v1 * 787)
    }
    v17 := GenS4{F0: 0, F1: true}
    var v18 int = v2
    v1 = (v1 ^ (v1 | v1))
    fmt.Println(v2, v1, v2)
    fmt.Println(v1, v18)
    v18 += v2
    if true {
        var v19 int64 = 0
        v20 := GenS3{F0: 922, F1: 805}
        v1 += 2147483647
        _ = v19
        _ = v20
    } else {
        var v21 string = "a"
        v22 := GenS2{F0: 32767}
        v18 += 1
        _ = v21
        _ = v22
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    var v23 int64 = -788
    var v24 int16 = 596
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
