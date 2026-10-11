package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint32; F1 int8; F2 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 string }

func main() {
    var v1 bool = true
    var v2 int = 32767
    v3 := []uint32{}
    _ = len(v3)
    var v4 float32 = 1.5
    switch v2 {
    case -9223372036854775808:
        v3 = append(v3, 512)
    case 32768:
        v5 := GenS1{F0: "hello"}
        _ = v5
    default:
        v4 += 3.14159
    }
    var v6 int64 = 302
    var v7 uint8 = 0
    for fi8 := 0; fi8 < 3; fi8++ {
        if ((v1 || (32767 != 42)) && true) {
            for fi9 := 0; fi9 < 1; fi9++ {
                switch -437 {
                case 2147483647:
                    var v10 uint8 = (v7 >> ((476) & 63))
                    _ = v10
                case 697:
                    v4 = v4
                default:
                    v11 := GenS2{F0: 100, F1: "foo bar"}
                    _ = v11
                }
                fmt.Println(v1, v1)
                _ = fi9
            }
            switch fi8 {
            case 2147483647:
                v12 := GenS3{F0: 2, F1: -128, F2: -128}
                _ = v12
            default:
                v1 = true
            }
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 string = "\n\t"
                v15 := GenS2{F0: -128, F1: (v14 + v14)}
                switch v14 {
                case "hello":
                    fmt.Println(fi13, v6)
                default:
                    switch ("\n\t" + ("" + (v14 + (v14 + v14)))) {
                    case "\n\t":
                        var v16 string = v14
                        _ = v16
                    case "":
                        var v17 float64 = 1.5
                        _ = v17
                    case "foo bar":
                        var v18 uint8 = 1
                        if v1 {
                            v2 = 128
                            var v19 int32 = 2147483647
                            _ = v19
                        }
                        _ = v18
                    }
                }
                var v20 int8 = 100
                _ = fi13
                _ = v14
                _ = v15
                _ = v20
            }
        } else {
            v1 = ((v6 << ((24) & 63)) != int64(v2))
        }
        if v1 {
            var v21 int = fi8
            fmt.Println(v1, fi8, v4)
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 float64 = -2.25
                v24 := GenS4{F0: (v7 ^ (v7 & (v7 % ((v7) | 1))))}
                var v25 uint = 65535
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = v21
        }
        v26 := GenS3{F0: 2147483648, F1: -92, F2: 0}
        _ = fi8
        _ = v26
    }
    var v27 uint = 65535
    for fi28 := 0; fi28 < 6; fi28++ {
        v6 -= (v6 | v6)
        _ = fi28
    }
    fmt.Println(v4, v7)
    for fi29 := 0; fi29 < 19; fi29++ {
        v7 = (v7 - (v7 | (v7 >> (((v27 << (((v27 & v27)) & 63))) & 63))))
        v7 = (^(v7))
        _ = fi29
    }
    var v30 uint = v27
    switch v2 {
    case 2147483647:
        v31 := GenS5{F0: 42}
        _ = v31
    case 0:
        v32 := GenS6{F0: v1, F1: ("\n\t" + "a")}
        _ = v32
    case -2147483648:
        switch "hello" {
        case "\n\t":
            v1 = v1
        case "a":
            _ = len(v3)
        }
    }
    v3 = append(v3, uint32(v27))
    if len(v3) > 0 {
        v3[0] = 391
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        for fi34 := 0; fi34 < 8; fi34++ {
            _ = len(v3)
            var v35 float64 = 0.1
            _ = fi34
            _ = v35
        }
        if (!(v1 || (282 == 17))) {
            var v36 float64 = 3.14159
            if false {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                var v37 bool = ((v4 >= (3.14159 / (v4 + float32(((((0.0 * v36) + 1e10) - 1.5) / 1.5))))) || (v6 == int64(v7)))
                _ = v37
            }
            _ = v36
        }
        _ = fi33
    }
    v1 = (2147483647 < 2147483648)
    if (v2 != (v2 + 32767)) {
        fmt.Println(v2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v6
    _ = v7
}
