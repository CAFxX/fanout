package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 363
    var v3 uint32 = 0
    v4 := []uint32{v3}
    v5 := make(map[int]bool)
    v6 := GenS1{F0: 2}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v7 uint8 = 42
    v7 = (v7 << ((v2) & 63))
    v3 -= (v3 % ((v3) | 1))
    var v8 float32 = ((-2.25 * (-2.25 / 1e10)) / 1.5)
    var v9 uint32 = v3
    if false {
        switch "" {
        case "a":
            fmt.Println(v9)
        case "":
            var v10 uint = v2
            _ = v10
        default:
            fmt.Println(v2)
        }
        v7 += 127
        v2 += v2
    }
    if (100 != -1) {
        var v11 int16 = 32767
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v11
    } else {
        fmt.Println(v9)
    }
    var v12 uint16 = 2
    v7 = 0
    v2 = v2
    switch "hello" {
    case "a":
        for fi13 := 0; fi13 < 9; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                v15 := GenS1{F0: 777}
                var v16 uint8 = v7
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 uint64 = 42
            _ = fi13
            _ = v17
        }
    case "":
        v18 := v1[269]
        _ = v18
    }
    switch "\n\t" {
    case "foo bar":
        var v19 int32 = 2147483647
        _ = v19
    case "a":
        v9 = (v3 << (((v2 ^ v2)) & 63))
    case "":
        if false {
            var v20 string = "foo bar"
            _ = v20
        } else {
            if (((v7 / ((v7) | 1)) < (v7 | (v7 << (((v2 | (^(v2)))) & 63)))) && (1000 != 893)) {
                v21 := GenS1{F0: v12}
                switch ("" + "") {
                case "a":
                    v4 = append(v4, v3)
                case "hello":
                    delete(v1, v9)
                case "\n\t":
                    delete(v5, 9223372036854775807)
                }
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 int32 = 1
                    _ = fi22
                    _ = v23
                }
                _ = v21
            }
            v24 := GenS1{F0: v12}
            _ = v24
        }
    default:
        for fi25 := 0; fi25 < 11; fi25++ {
            fmt.Println(v3, v8)
            v26 := GenS2{F0: 0, F1: (v7 | v7), F2: int64(v9)}
            _ = fi25
            _ = v26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
