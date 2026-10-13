package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint64 }
func genFn0(p0 int32) uint32 {
    var v4 int = 128
    switch v4 {
    case -793:
        v5 := GenS1{F0: (p0 ^ 42)}
        _ = v5
    default:
        var v6 uint = 4294967295
        _ = v6
    }
    _ = p0
    _ = v4
    return 65535
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    switch 605 {
    case 2147483647:
        var v7 uint8 = 236
        _ = v7
    default:
        if ((1000 <= -32768) || (0.1 >= 1.5)) {
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := GenS2{F0: 42, F1: 255, F2: 18446744073709551615}
                var v10 uint32 = 1
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            var v11 float32 = 1.5
            _ = len(v1)
            v12 := v1[-9223372036854775808]
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        fmt.Println(fi13)
        _ = fi13
    }
    switch "a" {
    case "hello":
        var v14 int32 = -848
        _ = v14
    case "\n\t":
        var v15 uint = 255
        _ = v15
    case "foo bar":
        v16 := GenS1{F0: 2147483647}
        _ = v16
    default:
        var v17 uint32 = genFn0(-924)
        _ = v17
    }
    v18 := GenS1{F0: -1}
    for fi19 := 0; fi19 < 18; fi19++ {
        delete(v3, 89)
        var v20 uint64 = 4294967295
        v20 = (^(v20))
        _ = fi19
        _ = v20
    }
    switch ("hello" + "\n\t") {
    case "":
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 string = ""
            _ = fi21
            _ = v22
        }
    case "\n\t":
        for fi23 := 0; fi23 < 8; fi23++ {
            v24 := GenS2{F0: 1000, F1: 0, F2: 667}
            var v25 int32 = 1
            _ = fi23
            _ = v24
            _ = v25
        }
    case "a":
        var v26 float32 = 3.14159
        _ = v26
    }
    var v27 float64 = (float64(((float32(((((3.14159 / float64(3.14159)) / -2.25) + float64(1.5)) + (float64(1.5) + 1e10))) + 1.5) / -2.25)) + 1e10)
    var v28 float32 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
