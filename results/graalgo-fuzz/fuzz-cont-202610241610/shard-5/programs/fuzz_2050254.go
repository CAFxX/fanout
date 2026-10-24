package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint64; F1 int8; F2 int8 }

func main() {
    var v1 float64 = float64(3.14159)
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    v5 := make(map[int]int)
    var v6 uint32 = v2
    v7 := []int{184, -1, 32768, -2147483648}
    v6 = v6
    fmt.Println(v1, v2)
    v4[127] = true
    switch 2147483647 {
    case 2147483647:
        delete(v3, 0)
    default:
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                if true {
                    if (!false) {
                        var v10 string = "foo bar"
                        var v11 uint8 = 1
                        _ = v10
                        _ = v11
                    } else {
                        v1 -= (v1 - float64(0.1))
                        fmt.Println(v6, v2, fi9)
                        fmt.Println(v1)
                    }
                    if false {
                        v6 = 1
                        var v12 float32 = 0.1
                        var v13 uint8 = 42
                        _ = v12
                        _ = v13
                    } else {
                        v5[-2147483648] = (fi8 >> ((65535) & 63))
                        v7 = append(v7, 204)
                        if (!(-32768 >= 0)) {
                            v2 += (^(v6))
                            var v14 uint8 = 255
                            _ = v14
                        } else {
                            v6 += (v2 | v6)
                            fmt.Println(v2, v6, v1)
                        }
                    }
                    if true {
                        v1 -= v1
                    } else {
                        if (true && true) {
                            if true {
                                v1 = (0.1 + 3.14159)
                                v2 = (v6 + (v2 - v2))
                                var v15 int16 = 0
                                _ = v15
                            }
                        }
                        var v16 int16 = 425
                        v1 = -2.25
                        _ = v16
                    }
                }
                var v17 int32 = 818
                _ = fi9
                _ = v17
            }
            var v18 uint32 = v2
            _ = fi8
            _ = v18
        }
    }
    if false {
        var v19 uint16 = 1
        _ = v19
    } else {
        var v20 float64 = 0.0
        var v21 int8 = -1
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS1{F0: 2, F1: 890}
            v2 = v2
            v20 = (-2.25 / (v1 - (1.5 + v20)))
            _ = fi22
            _ = v23
        }
        _ = v20
        _ = v21
    }
    switch "hello" {
    case "foo bar":
        v24 := GenS1{F0: 255, F1: -1}
        _ = v24
    case "hello":
        switch (("" + ("foo bar" + "\n\t")) + "a") {
        case "":
            v6 = v2
        default:
            v25 := GenS1{F0: 21, F1: -675}
            _ = v25
        }
    case "":
        var v26 int8 = -1
        _ = v26
    }
    if ("" >= ("hello" + "foo bar")) {
        v2 -= v6
        switch -218 {
        case 1:
            for fi27 := 0; fi27 < 6; fi27++ {
                _ = len(v4)
                v1 += 0.1
                _ = fi27
            }
        }
    } else {
        v28 := GenS2{F0: 0, F1: 127}
        v29 := GenS1{F0: 232, F1: 1}
        _ = v28
        _ = v29
    }
    v30 := GenS3{F0: 1, F1: 100, F2: 100}
    var v31 int64 = -1
    var v32 bool = true
    v1 -= v1
    v6 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
