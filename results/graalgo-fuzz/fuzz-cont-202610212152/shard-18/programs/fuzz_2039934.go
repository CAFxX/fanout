package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64, p1 int64, p2 string) string {
    var v4 uint8 = 128
    var v5 int8 = 5
    p2 = ("" + p2)
    fmt.Println(v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p2 + "\n\t")
}
func genFn1(p0 int16, p1 int64) uint8 {
    if true {
        var v6 uint8 = 128
        var v7 uint32 = 0
        switch -9223372036854775808 {
        case 127:
            switch -267 {
            case -2147483648:
                for fi8 := 0; fi8 < 17; fi8++ {
                    switch "hello" {
                    case "":
                        fmt.Println(p1, v6)
                    case "hello":
                        v9 := GenS1{F0: fi8, F1: true}
                        _ = v9
                    }
                    _ = fi8
                }
            case 9223372036854775807:
                var v10 uint64 = 4294967295
                _ = v10
            }
        case 389:
            switch 635 {
            case 9223372036854775807:
                var v11 uint32 = (v7 % ((4294967295) | 1))
                _ = v11
            }
        }
        _ = v6
        _ = v7
    } else {
        var v12 int16 = p0
        fmt.Println(p1, p1)
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        p1 = p1
        _ = fi13
    }
    _ = p0
    _ = p1
    return 244
}

func main() {
    v1 := []int{0, 0}
    var v2 string = ("\n\t" + "hello")
    v3 := []uint32{2147483647}
    fmt.Println(v2)
    switch "\n\t" {
    case "hello":
        v2 = "hello"
    case "foo bar":
        v14 := GenS2{F0: -1}
        _ = v14
    default:
        if (!true) {
            var v15 bool = false
            _ = v15
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v16 := GenS1{F0: 127, F1: (-1 < 1)}
    var v17 int64 = 0
    switch v2 {
    case "\n\t":
        if true {
            var v18 uint16 = 0
            v17 -= (v17 - (v17 / (((v17 & v17)) | 1)))
            _ = v18
        }
    default:
        if (985 > 785) {
            var v19 float64 = ((3.14159 * (0.0 - float64((0.0 * -2.25)))) + float64(float32(float64(1.5))))
            var v20 int64 = (v17 & v17)
            _ = v19
            _ = v20
        } else {
            fmt.Println(v17, v2)
            for fi21 := 0; fi21 < 8; fi21++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v22 uint64 = 4294967295
                _ = fi21
                _ = v22
            }
            fmt.Println(v2, v2)
        }
    }
    v17 = (v17 % ((v17) | 1))
    var v23 int32 = 0
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
}
