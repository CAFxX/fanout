package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint32 = 36
    var v2 float64 = 0.0
    v3 := make(map[int]int)
    var v4 int16 = 42
    v5 := []uint32{v1}
    v6 := make(map[int]uint32)
    var v7 uint64 = 0
    if (670 >= 166) {
        v5 = append(v5, (^(v1)))
        fmt.Println(v4, v2, v1)
    }
    v8 := GenS1{F0: true}
    for fi9 := 0; fi9 < 12; fi9++ {
        fmt.Println(v4, v4, fi9)
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: false}
            var v12 int32 = 42
            v4 -= (v4 << ((4294967295) & 63))
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    var v13 uint = 496
    v6[-1] = 2
    var v14 float32 = 0.0
    if true {
        if (uint32(v7) != (v1 + v1)) {
            switch 1 {
            case 2147483647:
                switch int(v13) {
                case -1:
                    var v15 uint64 = v7
                    _ = v15
                default:
                    var v16 uint = v13
                    _ = v16
                }
            case 127:
                fmt.Println(v13, v13)
            }
            var v17 uint32 = (v1 - v1)
            _ = v17
        } else {
            fmt.Println(v13, v14)
            for fi18 := 0; fi18 < 11; fi18++ {
                _ = len(v5)
                switch (fi18 & (fi18 + (^(fi18)))) {
                case 570:
                    var v19 int16 = -1
                    _ = v19
                case 32767:
                    v4 += 1000
                case -99:
                    switch "hello" {
                    case "\n\t":
                        var v20 uint8 = 2
                        _ = v20
                    default:
                        var v21 int32 = 207
                        _ = v21
                    }
                default:
                    v6[(fi18 % ((int(v7)) | 1))] = uint32(v7)
                }
                fmt.Println(fi18)
                _ = fi18
            }
        }
        if (!(0 > 32767)) {
            v4 = 0
            v13 = v13
        } else {
            if false {
                var v22 uint = (v13 / (((v13 >> ((v13) & 63))) | 1))
                _ = v22
            }
            var v23 int16 = 966
            v23 = v23
            _ = v23
        }
        v24 := v3[1]
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        fmt.Println(v1, v1)
        v14 += (v14 - 1.5)
        _ = fi25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 uint32 = v1
        var v28 float32 = 1.5
        var v29 int64 = -9223372036854775808
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: false}
    var v31 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
