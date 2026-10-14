package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (1.5 - 1.5)
    var v3 int = -9223372036854775808
    v4 := make(map[uint32]uint32)
    var v5 int16 = -1
    var v6 int = (v3 & (v3 & (-(v3))))
    fmt.Println(v3)
    var v7 int16 = 824
    v7 -= (v7 | (v5 >> ((601) & 63)))
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: (v7 ^ v5), F1: 1}
            v6 -= 359
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 uint = 0
    var v12 int32 = -504
    v13 := GenS2{F0: true, F1: (v12 << (((v11 / ((v11) | 1))) & 63))}
    if (true && (9223372036854775807 <= 9223372036854775807)) {
        switch ("foo bar" + ("\n\t" + "a")) {
        case "hello":
            for fi14 := 0; fi14 < 13; fi14++ {
                fmt.Println(v6, v5, v11)
                var v15 uint64 = 18446744073709551615
                v6 += 9223372036854775807
                _ = fi14
                _ = v15
            }
        case "a":
            fmt.Println(v3, v12)
        default:
            v16 := []int{(v3 / ((32768) | 1)), v6, v3}
            _ = v16
        }
    }
    v17 := []int{(v6 * v3), (v6 - (^(v6))), (v6 * v6), v6}
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 float32 = 0.0
        v20 := GenS3{F0: "a", F1: 42}
        var v21 bool = false
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    v3 = 42
    var v22 string = ""
    if true {
        if (2 >= 360) {
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 bool = false
                v11 = v11
                var v25 int64 = -1
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            _ = fi26
        }
    } else {
        for fi27 := 0; fi27 < 3; fi27++ {
            v6 -= (v6 / ((32768) | 1))
            _ = fi27
        }
        fmt.Println(v2, v11, v5)
        if (true && true) {
            switch "hello" {
            case "foo bar":
                var v28 int16 = v5
                _ = v28
            case "a":
                for fi29 := 0; fi29 < 15; fi29++ {
                    for fi30 := 0; fi30 < 18; fi30++ {
                        _ = len(v17)
                        _ = fi30
                    }
                    _ = fi29
                }
            default:
                v31 := GenS4{F0: false}
                _ = v31
            }
        }
    }
    switch ("foo bar" + v22) {
    case "a":
        v5 = (v7 | (v7 * v7))
    case "\n\t":
        var v32 int64 = -1
        _ = v32
    }
    fmt.Println(v7, v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
