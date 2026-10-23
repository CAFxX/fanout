package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := []int{660, -51}
    var v2 uint8 = 255
    var v3 uint16 = 0
    _ = len(v1)
    var v4 uint32 = 981
    var v5 int32 = -2147483648
    if false {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v5, v2)
        }
        for fi6 := 0; fi6 < 10; fi6++ {
            v7 := GenS1{F0: 42}
            fmt.Println(v4)
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 int = int(v4)
                var v10 uint32 = v4
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    } else {
        v11 := GenS2{F0: v2, F1: v3}
        _ = v11
    }
    var v12 string = "a"
    for fi13 := 0; fi13 < 11; fi13++ {
        v2 += v2
        v14 := GenS2{F0: v2, F1: (v3 >> ((2) & 63))}
        if (3.14159 < (0.1 - 1e10)) {
            var v15 int16 = -32768
            _ = v15
        } else {
            if len(v1) > 0 {
                v1[0] = (fi13 ^ (fi13 & (fi13 << ((1) & 63))))
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 uint32 = (v4 + (v4 / (((v4 ^ 2)) | 1)))
                if true {
                    switch v12 {
                    case "":
                        var v18 bool = false
                        _ = v18
                    case "\n\t":
                        v3 += (v3 | v3)
                    case "hello":
                        var v19 uint16 = 120
                        _ = v19
                    }
                    if false {
                        switch (fi16 + fi13) {
                        case -2147483648:
                            v1 = append(v1, fi13)
                            if (-128 == -128) {
                                if false {
                                    var v20 int8 = -37
                                    v21 := GenS1{F0: v20}
                                    var v22 int8 = (v20 - v20)
                                    if (-1 < 1) {
                                        v17 -= (v4 + 65535)
                                        var v23 uint8 = (v2 + v2)
                                        var v24 float64 = 1e10
                                        _ = v23
                                        _ = v24
                                    }
                                    v25 := GenS1{F0: (v22 ^ v20)}
                                    var v26 uint16 = (v3 | (v3 * v3))
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                    _ = v25
                                    _ = v26
                                }
                            }
                            v17 -= (v4 * 2)
                        default:
                            switch v12 {
                            case "foo bar":
                                v5 += v5
                            }
                        }
                        var v27 uint32 = (v17 << ((663) & 63))
                        _ = v27
                    }
                    var v28 int16 = 1
                    _ = v28
                }
                v4 += v17
                _ = fi16
                _ = v17
            }
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v12, v3)
    v5 += -1
    v4 += v4
    v5 += v5
    v1 = append(v1, 1)
    var v29 uint16 = v3
    v30 := GenS1{F0: 1}
    v4 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
