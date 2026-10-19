package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 string) int64 {
    v5 := GenS1{F0: 1, F1: -563}
    v6 := GenS2{F0: 42, F1: -469}
    v7 := []int{}
    fmt.Println(p0)
    for fi8 := 0; fi8 < 3; fi8++ {
        if true {
            var v9 int64 = 0
            v10 := GenS1{F0: 0, F1: fi8}
            p0 = "\n\t"
            _ = v10
            _ = v9
        }
        v7 = append(v7, (fi8 ^ 2147483647))
        for fi11 := 0; fi11 < 15; fi11++ {
            fmt.Println(p0, p0, fi11)
            p0 = "a"
            var v12 string = "hello"
            _ = fi11
            _ = v12
        }
        _ = fi8
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int32 = 367
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    var v4 int32 = (v1 + v1)
    _ = len(v2)
    if (1.5 < ((1e10 / 1.5) - 0.1)) {
        switch "" {
        case "":
            switch 127 {
            case 556:
                fmt.Println(v4, v1)
            case 32768:
                v13 := GenS3{F0: 209, F1: true}
                _ = v13
            default:
                switch -2147483648 {
                case 443:
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v4, v1)
                    }
                case 0:
                    fmt.Println(v4, v1)
                case 725:
                    var v14 uint16 = 65535
                    _ = v14
                }
            }
        default:
            v15 := []int{-465, 795, 0}
            _ = v15
        }
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            v16 := []int{}
            _ = v16
        case "hello":
            for fi17 := 0; fi17 < 9; fi17++ {
                v4 = -59
                if false {
                    v4 -= v1
                }
                _ = fi17
            }
        }
    }
    var v18 uint = 255
    v19 := GenS4{F0: 65535}
    v20 := GenS5{F0: 253, F1: -9223372036854775808, F2: 18446744073709551615}
    v18 = 2
    fmt.Println(v1)
    v21 := []int{9223372036854775807}
    delete(v3, (("\n\t" + "a") + "a"))
    var v22 float32 = 3.14159
    v23 := GenS6{F0: 479}
    var v24 int64 = -9223372036854775808
    v4 = v4
    _ = len(v21)
    if false {
        v4 = (-(v4))
        v1 = 2147483647
        v4 -= (v4 << (((v18 ^ v18)) & 63))
    }
    var v25 int16 = 624
    var v26 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
