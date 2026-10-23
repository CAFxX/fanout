package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint }
type GenS4 struct { F0 int32; F1 string; F2 uint16 }

func main() {
    var v1 uint = 42
    var v2 float32 = 1e10
    v3 := make(map[string]bool)
    var v4 uint = (v1 | (v1 & (v1 % ((v1) | 1))))
    v5 := make(map[int]string)
    v6 := make(map[uint32]int)
    v2 = float32(float64((1.5 / (v2 - float32((float64(v2) - 0.1))))))
    v1 = 1
    v1 += v1
    fmt.Println(v1)
    if true {
        var v7 float64 = (3.14159 - 0.1)
        fmt.Println(v2, v1, v2)
        v8 := GenS1{F0: 2, F1: 1}
        _ = v7
        _ = v8
    }
    if true {
        if (v2 > (v2 / v2)) {
            var v9 int32 = -1
            switch -1 {
            case 2147483647:
                var v10 uint16 = 1
                _ = v10
            }
            v2 -= ((0.0 * 1e10) * (0.0 + (v2 - (((v2 / 1e10) / ((0.0 * v2) * 0.0)) - v2))))
            _ = v9
        }
        v4 += v4
        if (1 < 36) {
            v2 = 0.0
            var v11 float64 = 1e10
            if false {
                switch -963 {
                case 9223372036854775807:
                    var v12 uint8 = 255
                    _ = v12
                default:
                    fmt.Println(v11)
                }
                switch ("foo bar" + "hello") {
                case "hello":
                    delete(v5, 1)
                case "a":
                    var v13 uint = v1
                    _ = v13
                case "":
                    switch "hello" {
                    case "\n\t":
                        v14 := []int{0}
                        _ = v14
                    case "a":
                        v1 = 65535
                    }
                default:
                    v15 := GenS1{F0: 0, F1: 263}
                    _ = v15
                }
            }
            _ = v11
        }
    }
    v16 := []int{128}
    var v17 uint64 = 42
    fmt.Println(v1, v2)
    v17 -= (^(v17))
    v18 := GenS1{F0: 0, F1: 65535}
    for fi19 := 0; fi19 < 6; fi19++ {
        v20 := GenS2{F0: 255}
        v2 += v2
        var v21 int = (fi19 - (fi19 | fi19))
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: v4, F1: 225, F2: v1}
    v23 := GenS4{F0: -1, F1: "", F2: 65535}
    if ((646 == 32767) && true) {
        var v24 int16 = 42
        v25 := GenS4{F0: 2147483647, F1: "\n\t", F2: 114}
        _ = v24
        _ = v25
    } else {
        v1 = v1
        var v26 int16 = -1
        switch "\n\t" {
        case "a":
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 string = ("" + (("hello" + "\n\t") + "a"))
                v26 = v26
                _ = fi27
                _ = v28
            }
        case "foo bar":
            var v29 uint64 = (^(v17))
            _ = v29
        }
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
