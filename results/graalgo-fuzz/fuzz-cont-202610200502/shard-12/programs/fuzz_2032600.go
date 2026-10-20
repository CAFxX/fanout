package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 563
    var v3 uint16 = 65535
    var v4 uint16 = v3
    v5 := []int{}
    v6 := GenS1{F0: v2, F1: "a"}
    var v7 float32 = -2.25
    switch "" {
    case "foo bar":
        switch -9223372036854775808 {
        case 32768:
            fmt.Println(v2, v4, v3)
        case 9223372036854775807:
            if true {
                v8 := GenS2{F0: true, F1: 73}
                if false {
                    v2 = v2
                    fmt.Println(v4)
                } else {
                    delete(v1, 2147483647)
                    if (!false) {
                        var v9 uint8 = 128
                        if len(v5) > 0 {
                            v5[0] = 4294967295
                        }
                        _ = v9
                    } else {
                        v1[2147483648] = ("foo bar" + ("a" + "foo bar"))
                        fmt.Println(v3, v7, v4)
                        if false {
                            var v10 string = "foo bar"
                            _ = v10
                        }
                    }
                    switch "hello" {
                    case "hello":
                        if false {
                            v11 := GenS3{F0: 32767, F1: (v2 + v2)}
                            v2 = 1
                            _ = v11
                        } else {
                            v7 -= (v7 / 0.1)
                        }
                    case "a":
                        fmt.Println(v7)
                    }
                }
                _ = v8
            }
        }
    case "hello":
        v5 = append(v5, 42)
    default:
        for fi12 := 0; fi12 < 11; fi12++ {
            fmt.Println(v4, fi12)
            _ = fi12
        }
    }
    v5 = append(v5, 668)
    var v13 string = "foo bar"
    switch "foo bar" {
    case "hello":
        _ = len(v1)
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        fmt.Println(v3, v13)
        var v15 bool = false
        if v15 {
            v2 = (v2 - (v2 & v2))
            var v16 int8 = 1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v16
        } else {
            fmt.Println(v3)
        }
        _ = fi14
        _ = v15
    }
    v7 -= v7
    var v17 uint16 = (v3 * 558)
    fmt.Println(v7, v7)
    var v18 int16 = 0
    v13 = (v13 + v13)
    if false {
        var v19 int16 = (-(v18))
        v2 = v2
        _ = v19
    } else {
        v17 -= (^(v3))
        v2 = (v2 << ((18446744073709551615) & 63))
        fmt.Println(v7, v2, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
