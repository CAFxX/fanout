package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint32 = 65535
    var v2 int = 322
    v3 := make(map[int]bool)
    var v4 bool = true
    var v5 uint8 = 23
    var v6 bool = v4
    var v7 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 18; fi8++ {
        v7 += (v7 & v7)
        _ = fi8
    }
    var v9 uint32 = v1
    v10 := GenS1{F0: 1}
    var v11 float64 = 0.0
    var v12 float64 = v11
    if (((v11 * (v12 + v12)) - 0.0) > (0.0 + 0.1)) {
        var v13 int16 = int16(v7)
        for fi14 := 0; fi14 < 2; fi14++ {
            fmt.Println(v13, v12, v5)
            v9 = (v9 ^ 2)
            v15 := []int{(fi14 - (-(fi14))), fi14, fi14, (-(fi14))}
            _ = fi14
            _ = v15
        }
        delete(v3, v2)
        _ = v13
    } else {
        v16 := GenS2{F0: v2}
        var v17 int16 = 1
        _ = v16
        _ = v17
    }
    v6 = (!v6)
    if v4 {
        v18 := GenS1{F0: v7}
        var v19 uint8 = 127
        _ = v18
        _ = v19
    } else {
        var v20 int64 = (v7 << ((18446744073709551615) & 63))
        v21 := []int{(v2 & -813)}
        if v6 {
            for fi22 := 0; fi22 < 12; fi22++ {
                fmt.Println(v6, v4)
                if (!(v4 && v4)) {
                    v23 := GenS3{F0: 81}
                    _ = v23
                }
                _ = fi22
            }
            fmt.Println(v9)
        }
        _ = v20
        _ = v21
    }
    if (432 >= v7) {
        fmt.Println(v6)
        var v24 float32 = 0.0
        fmt.Println(v24, v11)
        _ = v24
    } else {
        var v25 float32 = -2.25
        fmt.Println(v4, v7, v6)
        _ = v25
    }
    switch (^(v2)) {
    case 128:
        v3[(v2 + (v2 & (-(v2))))] = false
    case -623:
        var v26 uint16 = 32767
        _ = v26
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 int64 = (-(v7))
        _ = fi27
        _ = v28
    }
    switch "a" {
    case "\n\t":
        v4 = false
    case "":
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
