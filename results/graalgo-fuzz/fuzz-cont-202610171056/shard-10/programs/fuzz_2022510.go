package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64) int64 {
    v5 := []int{70}
    var v6 int16 = 0
    if len(v5) > 0 {
        v5[0] = 1
    }
    _ = p0
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 uint, p1 uint16) uint16 {
    v7 := GenS1{F0: -1}
    var v8 float32 = 1e10
    var v9 bool = (234 > 0)
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(p1, p0, fi10)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 int8 = 1
    var v2 bool = (0 == 4294967295)
    var v3 int = 1
    var v4 int8 = (v1 - (v1 >> ((988) & 63)))
    v11 := []int{v3}
    var v12 uint = 736
    switch v3 {
    case 9223372036854775807:
        var v13 string = "a"
        _ = v13
    case 0:
        _ = len(v11)
    case -1:
        var v14 int8 = (v4 - v4)
        _ = v14
    default:
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 uint = 255
                var v18 float32 = 1.5
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = len(v11)
            _ = fi15
        }
    }
    v19 := GenS1{F0: 9223372036854775807}
    fmt.Println(v4)
    if len(v11) > 0 {
        v11[0] = (v3 & v3)
    }
    fmt.Println(v12)
    var v20 uint8 = 70
    var v21 float32 = -2.25
    if (65535 <= 42) {
        for fi22 := 0; fi22 < 13; fi22++ {
            v12 -= (v12 | v12)
            v20 = (v20 >> (((v12 ^ 65535)) & 63))
            v23 := GenS1{F0: 1}
            _ = fi22
            _ = v23
        }
        var v24 uint = v12
        _ = v24
    } else {
        for fi25 := 0; fi25 < 18; fi25++ {
            v1 -= v4
            var v26 int32 = 2147483647
            v11 = append(v11, v3)
            _ = fi25
            _ = v26
        }
    }
    if (v2 || (v20 > (v20 << (((^(v12))) & 63)))) {
        v2 = v2
        v27 := GenS1{F0: 42}
        v20 += (v20 >> (((v12 ^ (v12 / ((v12) | 1)))) & 63))
        _ = v27
    }
    var v28 string = "\n\t"
    switch 1 {
    case 128:
        var v29 int = (v3 + (v3 * 2147483647))
        _ = v29
    case 9223372036854775807:
        v21 += (0.1 - (v21 * v21))
    case 905:
        v4 -= (v4 | 0)
    default:
        var v30 uint16 = 0
        _ = v30
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 int64 = 1
        v33 := GenS1{F0: v32}
        _ = fi31
        _ = v32
        _ = v33
    }
    if (1 >= 660) {
        var v34 uint = v12
        fmt.Println(v34)
        _ = v34
    } else {
        v35 := GenS2{F0: uint64(v20)}
        for fi36 := 0; fi36 < 8; fi36++ {
            v1 += v4
            v21 -= (-2.25 * v21)
            _ = fi36
        }
        _ = v35
    }
    var v37 bool = (v4 < v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v37
    _ = v4
}
