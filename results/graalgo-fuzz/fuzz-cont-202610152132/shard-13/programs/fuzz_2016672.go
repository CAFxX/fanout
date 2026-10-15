package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    var v3 int16 = 1
    var v4 bool = true
    v5 := make(map[int]uint32)
    v3 = (v3 << ((255) & 63))
    v4 = (!(!(4294967295 != 127)))
    var v6 uint8 = 128
    var v7 float32 = 1.5
    var v8 bool = true
    switch 42 {
    case 1:
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := []int{(fi9 >> ((18446744073709551615) & 63)), fi9, (fi9 >> ((18446744073709551615) & 63)), 845}
            _ = fi9
            _ = v10
        }
    case 108:
        v4 = (42 != (v3 / (((-(v3))) | 1)))
    case -9223372036854775808:
        _ = len(v2)
    default:
        var v11 bool = false
        _ = v11
    }
    var v12 float32 = ((1.5 - -2.25) + 0.1)
    if v8 {
        for fi13 := 0; fi13 < 5; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                v8 = (!v8)
                _ = fi14
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                fmt.Println(v6)
                if (true && v8) {
                    switch "foo bar" {
                    case "hello":
                        v3 += (v3 ^ v3)
                    }
                    var v16 int64 = -286
                    var v17 int8 = 0
                    var v18 float32 = v7
                    v7 = 3.14159
                    _ = v16
                    _ = v17
                    _ = v18
                }
                if v4 {
                    _ = len(v2)
                } else {
                    var v19 float64 = float64(1e10)
                    fmt.Println(v3, v3)
                    switch 691 {
                    case -453:
                        switch 2147483647 {
                        case -9223372036854775808:
                            v20 := []int{(fi15 % (((fi15 & fi13)) | 1)), (fi13 >> ((979) & 63))}
                            _ = v20
                        case 42:
                            v21 := GenS1{F0: v4, F1: 32767, F2: (^(v3))}
                            _ = v21
                        default:
                            var v22 uint8 = v6
                            _ = v22
                        }
                    default:
                        var v23 int = (fi15 & 554)
                        _ = v23
                    }
                    _ = v19
                }
                var v24 uint32 = 2147483648
                _ = fi15
                _ = v24
            }
            if true {
                var v25 int8 = 100
                _ = v25
            }
            _ = fi13
        }
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        fmt.Println(v6, fi26)
        fmt.Println(fi26, v12, v3)
        _ = fi26
    }
    var v27 bool = (425 <= 0)
    v8 = v4
    fmt.Println(v12, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
