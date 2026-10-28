package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int16; F2 int }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint8 = 60
    v2 := []uint32{2147483648, 693}
    v3 := []int{470, 1, 32767}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 := GenS1{F0: 1}
    var v5 int64 = -609
    if (true && (-128 != 100)) {
        v6 := GenS1{F0: -125}
        var v7 uint32 = 4294967295
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v6
        _ = v7
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        if ((fi9 * fi9) >= fi9) {
            fmt.Println(v5)
        }
        _ = fi9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: -76}
        _ = fi10
        _ = v11
    }
    var v12 int = 316
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS2{F0: false, F1: 1, F2: (-(fi13))}
        var v15 string = ("" + "a")
        v3 = append(v3, -9223372036854775808)
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint8 = (v1 >> ((42) & 63))
    switch v12 {
    case 42:
        fmt.Println(v1, v16, v12)
    case 32767:
        fmt.Println(v5, v1)
    case 32768:
        var v17 int16 = 32767
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := GenS2{F0: (-1 == 127), F1: 1, F2: (fi18 | -2147483648)}
        _ = fi18
        _ = v19
    }
    v20 := GenS3{F0: 948}
    var v21 uint16 = 255
    v3 = append(v3, 9223372036854775807)
    v3 = append(v3, v12)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
