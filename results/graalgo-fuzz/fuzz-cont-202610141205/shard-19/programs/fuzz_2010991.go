package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int32 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint32 = 2147483647
    var v2 float32 = 1.5
    var v3 uint32 = (v1 / (((v1 | (v1 >> ((949) & 63)))) | 1))
    var v4 int8 = 42
    v4 = int8(v1)
    var v5 int = 2147483647
    v6 := []int{(v5 + 128), v5, (v5 % (((v5 | v5)) | 1))}
    var v7 int8 = v4
    switch v5 {
    case 1:
        if ((v5 + (v5 >> ((4294967295) & 63))) <= v5) {
            var v8 uint16 = 42
            fmt.Println(v3, v5, v1)
            _ = v8
        } else {
            switch v5 {
            case 2147483647:
                v6 = append(v6, v5)
            case 744:
                for fi9 := 0; fi9 < 4; fi9++ {
                    var v10 uint64 = 42
                    _ = fi9
                    _ = v10
                }
            }
        }
    case -721:
        var v11 float64 = 0.0
        _ = v11
    default:
        var v12 int32 = -465
        _ = v12
    }
    v6 = append(v6, v5)
    switch "\n\t" {
    case "foo bar":
        v13 := GenS1{F0: 1000, F1: 9223372036854775807, F2: 626}
        _ = v13
    case "a":
        switch "\n\t" {
        case "hello":
            for fi14 := 0; fi14 < 2; fi14++ {
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 float64 = 0.1
                    var v17 float32 = v2
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                fmt.Println(v3, v7, v7)
                for fi18 := 0; fi18 < 1; fi18++ {
                    if (true || false) {
                        v2 = (1e10 * v2)
                        switch v5 {
                        case 127:
                            v7 += 1
                        case 499:
                            v5 = (fi14 - (fi14 - (v5 + fi14)))
                        default:
                            if (true || false) {
                                if false {
                                    v1 = (v3 - (v3 - (v1 & v3)))
                                    v6 = append(v6, fi14)
                                }
                                v7 += (v4 & (v7 + (v4 - (v7 << ((502) & 63)))))
                            } else {
                                fmt.Println(fi18, fi18)
                                v6 = append(v6, v5)
                                var v19 int16 = 32767
                                _ = v19
                            }
                        }
                    } else {
                        var v20 float64 = 1.5
                        v21 := GenS2{F0: 42}
                        _ = v20
                        _ = v21
                    }
                    var v22 uint = 2
                    _ = fi18
                    _ = v22
                }
                _ = fi14
            }
        case "a":
            var v23 int = (v5 / ((v5) | 1))
            _ = v23
        case "":
            var v24 uint32 = (^(v3))
            _ = v24
        }
    case "hello":
        var v25 bool = (1 == -364)
        _ = v25
    }
    var v26 bool = true
    fmt.Println(v5, v4)
    var v27 int64 = 1
    v5 -= (v5 >> ((402) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
