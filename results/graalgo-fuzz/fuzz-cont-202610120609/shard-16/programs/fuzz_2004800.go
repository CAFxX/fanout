package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int = 4294967295
    var v2 uint64 = 966
    v3 := []string{}
    v4 := make(map[int]uint32)
    switch (v1 | v1) {
    case -1:
        var v5 uint8 = 128
        _ = v5
    case 522:
        var v6 float64 = (0.1 + 1.5)
        _ = v6
    case 4294967295:
        v7 := GenS1{F0: 42, F1: 128, F2: (0.1 != -2.25)}
        _ = v7
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v1 += (v1 | 647)
    if (!true) {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int16 = -246
            var v10 float64 = 1.5
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        fmt.Println(v1, v2)
        for fi11 := 0; fi11 < 1; fi11++ {
            if true {
                var v12 uint8 = 49
                var v13 int8 = -1
                _ = v12
                _ = v13
            }
            v2 = v2
            _ = fi11
        }
    }
    _ = len(v4)
    var v14 string = (("a" + "a") + "foo bar")
    var v15 int32 = 2147483647
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 uint = 42
        if true {
            for fi18 := 0; fi18 < 14; fi18++ {
                switch (v1 - fi18) {
                case 4294967295:
                    v19 := v4[fi18]
                    _ = v19
                default:
                    v15 = v15
                }
                v4[fi16] = 2147483647
                delete(v4, v1)
                _ = fi18
            }
            var v20 int8 = 0
            _ = v20
        } else {
            if len(v3) > 0 {
                v3[0] = v14
            }
            v21 := GenS2{F0: 0}
            if (42 < 94) {
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 uint16 = 32767
                    var v24 uint8 = 42
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            } else {
                var v25 float32 = 0.1
                for fi26 := 0; fi26 < 20; fi26++ {
                    v1 += fi26
                    var v27 uint32 = 544
                    var v28 int = (fi16 + 32767)
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = v25
            }
            _ = v21
        }
        _ = len(v3)
        _ = fi16
        _ = v17
    }
    fmt.Println(v1)
    v29 := GenS3{F0: v2, F1: (223 >= 0)}
    var v30 int16 = -1
    var v31 uint = 2
    v32 := GenS2{F0: 620}
    var v33 uint64 = v2
    if (145 < 252) {
        var v34 uint = v31
        var v35 int32 = v15
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        v37 := GenS4{F0: 128}
        v2 = (v33 << (((v31 / ((v31) | 1))) & 63))
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
