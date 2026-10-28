package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 string; F1 int16; F2 uint8 }
type GenS4 struct { F0 bool }
func genFn0() uint16 {
    if false {
        var v6 float32 = 1.5
        v7 := GenS1{F0: 1, F1: -314, F2: 0}
        var v8 int8 = 0
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 int64 = 159
        switch "hello" {
        case "":
            var v10 int64 = (-(v9))
            _ = v10
        default:
            var v11 uint = 1
            _ = v11
        }
        _ = v9
    }
    v12 := []int{4294967295, 2147483647, 42, -1}
    switch 197 {
    case 127:
        var v13 uint8 = 128
        _ = v13
    }
    var v14 bool = (53 < 194)
    _ = v12
    _ = v14
    return 65535
}
func genFn1(p0 uint, p1 int32) int32 {
    v15 := GenS2{F0: ("a" + "foo bar"), F1: 105}
    v16 := []int{132, 2147483647, 9223372036854775807, 4294967295}
    switch -254 {
    case -9223372036854775808:
        var v17 uint32 = 2147483648
        _ = v17
    case -933:
        if false {
            p0 -= 2
        } else {
            v18 := GenS3{F0: "hello", F1: -656, F2: 58}
            for fi19 := 0; fi19 < 11; fi19++ {
                fmt.Println(p0)
                var v20 int = fi19
                _ = fi19
                _ = v20
            }
            _ = v18
        }
    case 459:
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return p1
}

func main() {
    var v1 bool = false
    var v2 float64 = (1.5 / 0.1)
    var v3 uint64 = 18446744073709551615
    var v4 int16 = 1000
    var v5 int64 = -861
    v21 := []int{32768}
    if (!(54 >= 151)) {
        var v22 int16 = (-(v4))
        _ = v22
    }
    if len(v21) > 0 {
        v21[0] = -2147483648
    }
    v5 = (v5 >> ((255) & 63))
    fmt.Println(v1, v2, v5)
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 uint = 0
        if (v3 > (v3 / ((v3) | 1))) {
            for fi25 := 0; fi25 < 2; fi25++ {
                switch "foo bar" {
                case "a":
                    v2 = (v2 - float64((float32(1e10) * ((1e10 - (1.5 * 3.14159)) / (1.5 + (-2.25 * (1e10 * 0.0)))))))
                case "foo bar":
                    v26 := GenS1{F0: v4, F1: fi23, F2: 929}
                    if v1 {
                        fmt.Println(v24, v3, v5)
                    }
                    _ = v26
                default:
                    v27 := GenS3{F0: "a", F1: 109, F2: 144}
                    v28 := GenS4{F0: v1}
                    switch (^(fi23)) {
                    case 204:
                        fmt.Println(fi25)
                    case -1:
                        v2 -= float64(3.14159)
                    default:
                        v29 := GenS1{F0: 1000, F1: 9223372036854775807, F2: 333}
                        _ = v29
                    }
                    _ = v27
                    _ = v28
                }
                _ = fi25
            }
        }
        _ = fi23
        _ = v24
    }
    var v30 bool = ("foo bar" != "")
    var v31 string = "\n\t"
    v4 -= (v4 - (v4 % ((1) | 1)))
    var v32 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
