package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
func genFn0(p0 int16, p1 int32, p2 int16) uint32 {
    v4 := []int{32767}
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(p0, p0, p0)
        switch "a" {
        case "hello":
            var v6 bool = (65535 >= 42)
            _ = v6
        case "\n\t":
            v7 := GenS1{F0: 0, F1: 42}
            _ = v7
        case "":
            if true {
                if len(v4) > 0 {
                    v4[0] = (fi5 >> ((255) & 63))
                }
            }
        }
        p0 -= -32768
        _ = fi5
    }
    _ = len(v4)
    v8 := GenS1{F0: 1, F1: -673}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return 0
}

func main() {
    var v1 uint8 = 128
    var v2 uint16 = 255
    var v3 bool = true
    fmt.Println(v1)
    fmt.Println(v3, v3, v1)
    v2 = v2
    if (v2 >= (^(v2))) {
        v9 := GenS1{F0: (v1 & 196), F1: 9223372036854775807}
        _ = v9
    } else {
        v10 := []int{32767, -934, 2147483647, 9223372036854775807}
        v2 -= 42
        _ = v10
    }
    var v11 int16 = -331
    var v12 uint64 = 173
    v12 = v12
    for fi13 := 0; fi13 < 14; fi13++ {
        if false {
            v14 := []int{fi13, 601, fi13, fi13}
            v11 -= (v11 | 42)
            _ = v14
        } else {
            v15 := GenS1{F0: (^(v1)), F1: (fi13 * (fi13 % ((fi13) | 1)))}
            fmt.Println(v3)
            _ = v15
        }
        for fi16 := 0; fi16 < 3; fi16++ {
            v2 += v2
            v12 -= v12
            _ = fi16
        }
        _ = fi13
    }
    if ((3.14159 * -2.25) > 1e10) {
        fmt.Println(v11, v12)
        var v17 int32 = 2147483647
        v1 = v1
        _ = v17
    }
    var v18 float64 = 0.1
    if (0.0 == (-2.25 + 0.1)) {
        v19 := GenS1{F0: 2, F1: int(v12)}
        var v20 float64 = v18
        _ = v19
        _ = v20
    } else {
        var v21 float64 = v18
        _ = v21
    }
    switch "\n\t" {
    case "hello":
        v1 = (v1 % (((v1 + v1)) | 1))
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := []int{(^(fi22)), fi22}
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: uint8(v12), F1: 128}
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := GenS1{F0: v1, F1: (fi25 + (fi25 << ((819) & 63)))}
        _ = fi25
        _ = v26
    }
    fmt.Println(v18)
    switch "\n\t" {
    case "hello":
        var v27 int16 = (-(v11))
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v24
    _ = v3
}
