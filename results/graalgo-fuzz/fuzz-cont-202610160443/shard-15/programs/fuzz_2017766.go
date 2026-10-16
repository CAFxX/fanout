package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []uint32{42, 255}
    var v2 int64 = -294
    var v3 uint64 = 0
    var v4 float32 = 0.0
    v5 := GenS1{F0: v3, F1: 1}
    _ = len(v1)
    var v6 uint64 = (v3 % (((v3 | v3)) | 1))
    v2 += (v2 | 9223372036854775807)
    var v7 uint = 0
    v3 -= 2
    var v8 int = -9223372036854775808
    fmt.Println(v2, v8, v7)
    if false {
        var v9 uint = 255
        _ = v9
    } else {
        v7 -= (^(v7))
        v6 = (v3 << ((719) & 63))
        var v10 uint64 = v6
        _ = v10
    }
    switch ("foo bar" + "") {
    case "hello":
        if len(v1) > 0 {
            v1[0] = 1
        }
    case "foo bar":
        v11 := GenS1{F0: 0, F1: 42}
        _ = v11
    case "\n\t":
        if (!true) {
            switch v8 {
            case -1:
                fmt.Println(v2)
            default:
                var v12 int64 = (v2 ^ (v2 >> (((v7 / (((^(v7))) | 1))) & 63)))
                _ = v12
            }
            v13 := GenS2{F0: 2147483647, F1: ((v3 > v6) && true)}
            for fi14 := 0; fi14 < 7; fi14++ {
                switch ("hello" + "a") {
                case "a":
                    fmt.Println(v3, v2)
                case "hello":
                    if (("" + (("\n\t" + "a") + "")) <= "a") {
                        v15 := GenS1{F0: v6, F1: 373}
                        var v16 uint32 = 4294967295
                        _ = v15
                        _ = v16
                    }
                }
                _ = fi14
            }
            _ = v13
        }
    default:
        _ = len(v1)
    }
    var v17 string = "foo bar"
    for fi18 := 0; fi18 < 10; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(v7)
            switch "" {
            case "hello":
                if len(v1) > 0 {
                    v1[0] = 1
                }
            case "a":
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
            }
            _ = fi19
        }
        for fi20 := 0; fi20 < 10; fi20++ {
            if false {
                switch "\n\t" {
                case "foo bar":
                    v1 = append(v1, 42)
                case "\n\t":
                    var v21 int8 = 100
                    _ = v21
                case "hello":
                    v6 = v6
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                var v22 uint64 = (v3 % (((v6 ^ (v6 << ((v7) & 63)))) | 1))
                v3 = 4294967295
                if len(v1) > 0 {
                    v1[0] = 0
                }
                var v23 int8 = 1
                _ = v22
                _ = v23
            }
            v24 := GenS3{F0: v2}
            _ = fi20
            _ = v24
        }
        _ = fi18
    }
    var v25 int8 = 80
    v2 += v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
