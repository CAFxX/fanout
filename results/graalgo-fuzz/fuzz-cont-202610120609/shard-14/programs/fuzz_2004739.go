package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 string; F1 int8 }

func main() {
    var v1 float32 = (((-2.25 + (1e10 + -2.25)) / ((0.1 / -2.25) / 3.14159)) / -2.25)
    var v2 float64 = (1.5 * float64(v1))
    var v3 int32 = 2147483647
    var v4 uint = 0
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: uint64(v4), F1: 0}
    if true {
        v7 := GenS1{F0: 42, F1: 100}
        v3 = -2147483648
        _ = v7
    } else {
        var v8 int8 = -22
        if (-1 > -523) {
            if true {
                var v9 uint8 = 2
                _ = v9
            }
        } else {
            v1 += (v1 - 1.5)
            v2 = v2
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 bool = true
            switch 1 {
            case 127:
                v12 := GenS2{F0: -644}
                _ = v12
            case -488:
                var v13 uint8 = 238
                _ = v13
            default:
                v3 = v3
            }
            _ = fi10
            _ = v11
        }
        _ = v8
    }
    if true {
        if true {
            var v14 int64 = 42
            _ = v14
        }
    }
    if (32767 == 9223372036854775807) {
        var v15 uint16 = 0
        if (76 < -128) {
            v16 := GenS1{F0: 2, F1: -128}
            var v17 uint = 1
            _ = v16
            _ = v17
        } else {
            switch "\n\t" {
            case "hello":
                switch 32768 {
                case 127:
                    var v18 float32 = v1
                    _ = v18
                case 1:
                    if false {
                        v2 -= 1e10
                    }
                default:
                    if (v3 >= 2147483647) {
                        for fi19 := 0; fi19 < 6; fi19++ {
                            v20 := GenS3{F0: 0, F1: "hello"}
                            var v21 uint16 = (v15 ^ v15)
                            fmt.Println(v4)
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    } else {
                        v22 := []int{977, -2147483648, 1}
                        _ = v22
                    }
                }
            case "":
                v23 := GenS4{F0: "hello", F1: 127}
                _ = v23
            default:
                fmt.Println(v2, v1)
            }
            var v24 string = "foo bar"
            _ = v24
        }
        v2 = v2
        _ = v15
    }
    switch "a" {
    case "":
        _ = len(v5)
    }
    var v25 uint8 = 164
    for fi26 := 0; fi26 < 18; fi26++ {
        if false {
            v1 = (v1 * v1)
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 uint = (v4 * 970)
                var v29 uint32 = 4294967295
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        _ = fi26
    }
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
