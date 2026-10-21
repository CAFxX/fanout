package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint32; F1 int }

func main() {
    var v1 uint16 = 184
    var v2 uint64 = 0
    v3 := []uint32{103, 65535, 65}
    v4 := []uint32{4294967295, 0}
    fmt.Println(v1, v1)
    for fi5 := 0; fi5 < 19; fi5++ {
        if false {
            v1 += (v1 / ((v1) | 1))
        }
        var v6 uint16 = (v1 >> ((65535) & 63))
        _ = fi5
        _ = v6
    }
    switch 1 {
    case 4294967295:
        switch 2147483647 {
        case 127:
            var v7 uint64 = v2
            _ = v7
        case -9223372036854775808:
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
        case 2147483647:
            v2 += (v2 - 188)
        }
    }
    if false {
        v8 := GenS1{F0: 32768, F1: "\n\t"}
        v9 := GenS2{F0: 0, F1: 9223372036854775807}
        v2 = (v2 & v2)
        _ = v8
        _ = v9
    } else {
        fmt.Println(v2, v2)
        if false {
            if false {
                switch "foo bar" {
                case "a":
                    v10 := GenS1{F0: 0, F1: (("\n\t" + ("foo bar" + "foo bar")) + "foo bar")}
                    _ = v10
                default:
                    v11 := GenS3{F0: 0, F1: -24}
                    _ = v11
                }
                switch -666 {
                case 2147483647:
                    v2 = 18446744073709551615
                case -1:
                    v1 = (v1 - v1)
                case 4294967295:
                    fmt.Println(v2, v1)
                default:
                    var v12 uint32 = 65535
                    _ = v12
                }
                switch "\n\t" {
                case "foo bar":
                    v2 += v2
                case "\n\t":
                    if (42 >= -1) {
                        var v13 uint8 = 127
                        var v14 float32 = 3.14159
                        _ = v13
                        _ = v14
                    }
                case "hello":
                    for fi15 := 0; fi15 < 4; fi15++ {
                        var v16 float64 = -2.25
                        fmt.Println(v16)
                        _ = fi15
                        _ = v16
                    }
                }
            }
        } else {
            v2 += 1
            if false {
                v3 = append(v3, uint32(v1))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v17 float64 = ((3.14159 / 1.5) - -2.25)
                _ = v17
            }
        }
        v1 += (v1 >> ((446) & 63))
    }
    var v18 int32 = 0
    v4 = append(v4, uint32(v2))
    var v19 string = "foo bar"
    var v20 uint64 = v2
    var v21 int = 2147483647
    fmt.Println(v20)
    var v22 float32 = (3.14159 + 0.1)
    _ = len(v4)
    v21 -= (^(v21))
    v21 = int(v1)
    v20 = 1
    for fi23 := 0; fi23 < 3; fi23++ {
        v19 = v19
        v3 = append(v3, 2147483647)
        _ = fi23
    }
    v19 = "hello"
    var v24 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
