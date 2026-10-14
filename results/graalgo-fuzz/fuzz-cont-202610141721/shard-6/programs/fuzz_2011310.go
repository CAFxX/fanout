package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int8; F1 int32 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]int)
    var v3 float64 = -2.25
    v4 := make(map[uint32]string)
    var v5 uint16 = 0
    var v6 int = 128
    switch 92 {
    case 9223372036854775807:
        if (false && false) {
            var v7 bool = (false && true)
            _ = v7
        }
    default:
        v5 = (v5 >> ((65535) & 63))
    }
    var v8 int64 = (v1 - (v1 >> ((42) & 63)))
    fmt.Println(v8)
    if true {
        if true {
            if (!(false || (497 > 147))) {
                for fi9 := 0; fi9 < 2; fi9++ {
                    var v10 uint32 = 2
                    v1 -= 817
                    var v11 uint = 4294967295
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            } else {
                var v12 string = ""
                var v13 uint64 = 594
                var v14 uint16 = (^(v5))
                _ = v12
                _ = v13
                _ = v14
            }
            v15 := []int{v6}
            _ = v15
        }
        if (1000 >= -32768) {
            var v16 int16 = 0
            if false {
                fmt.Println(v8, v1, v16)
                for fi17 := 0; fi17 < 15; fi17++ {
                    v18 := GenS1{F0: 272, F1: ("a" + "\n\t"), F2: (v16 | (v16 & (v16 >> ((946) & 63))))}
                    if false {
                        var v19 uint = 42
                        v8 -= (v8 | (v8 + (v8 >> ((18446744073709551615) & 63))))
                        _ = len(v2)
                        _ = v19
                    }
                    v20 := []int{(v6 ^ 32767), 9223372036854775807}
                    _ = fi17
                    _ = v18
                    _ = v20
                }
                fmt.Println(v3, v1, v3)
            } else {
                for fi21 := 0; fi21 < 13; fi21++ {
                    if false {
                        fmt.Println(v1, v6)
                    }
                    if (!true) {
                        v22 := v4[2147483648]
                        _ = v22
                    } else {
                        var v23 int = 1
                        _ = v23
                    }
                    _ = fi21
                }
                v24 := GenS2{F0: (258 >= 2147483647), F1: -1}
                _ = v24
            }
            _ = v16
        } else {
            var v25 float32 = 0.0
            _ = v25
        }
    } else {
        var v26 uint32 = 1
        v27 := GenS3{F0: 27, F1: -234}
        _ = v26
        _ = v27
    }
    var v28 uint = 18446744073709551615
    var v29 int32 = 997
    fmt.Println(v28)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
