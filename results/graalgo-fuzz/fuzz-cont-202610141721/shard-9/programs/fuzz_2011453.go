package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 bool; F2 int }
type GenS4 struct { F0 uint; F1 uint8; F2 int64 }

func main() {
    var v1 uint8 = 219
    v2 := []uint32{2147483647}
    v3 := []uint32{1, 65535}
    v4 := []int{128}
    v5 := make(map[uint32]string)
    var v6 float32 = float32(0.1)
    v7 := v5[0]
    for fi8 := 0; fi8 < 17; fi8++ {
        v4 = append(v4, (-(fi8)))
        v1 = (^(v1))
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 int32 = -219
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v5[1] = "\n\t"
    fmt.Println(v1, v1)
    v11 := GenS1{F0: 4294967295, F1: -2147483648, F2: 0}
    v7 = "foo bar"
    v6 = 3.14159
    switch 624 {
    case 4294967295:
        if false {
            fmt.Println(v6)
        } else {
            v1 += (^(v1))
            v12 := GenS2{F0: 42}
            _ = v12
        }
    case -1:
        var v13 uint16 = 721
        _ = v13
    case -468:
        for fi14 := 0; fi14 < 18; fi14++ {
            if len(v4) > 0 {
                v4[0] = fi14
            }
            _ = fi14
        }
    }
    var v15 int8 = 0
    v16 := GenS3{F0: 14, F1: false, F2: 32767}
    var v17 bool = false
    switch "\n\t" {
    case "foo bar":
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 uint64 = 2
            for fi20 := 0; fi20 < 16; fi20++ {
                if false {
                    var v21 float64 = 3.14159
                    var v22 uint = 255
                    _ = v21
                    _ = v22
                }
                fmt.Println(fi20)
                var v23 uint = uint(v1)
                _ = fi20
                _ = v23
            }
            var v24 uint8 = v1
            _ = fi18
            _ = v19
            _ = v24
        }
    case "":
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS3{F0: 187, F1: (!true), F2: 2147483647}
            for fi27 := 0; fi27 < 18; fi27++ {
                if v17 {
                    v28 := GenS4{F0: 42, F1: v1, F2: 42}
                    v17 = false
                    fmt.Println(fi27)
                    _ = v28
                }
                _ = fi27
            }
            _ = fi25
            _ = v26
        }
    case "hello":
        v17 = (412 <= 0)
    }
    if ((!(v17 && v17)) || v17) {
        fmt.Println(v17, v15)
        fmt.Println(v15, v6)
        v29 := GenS4{F0: 42, F1: (v1 * (^(v1))), F2: 42}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
