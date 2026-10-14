package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int; F2 int16 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := []uint32{1, 2147483648, 2147483647}
    var v2 string = "a"
    var v3 uint64 = 308
    var v4 uint8 = 226
    var v5 int8 = 1
    var v6 uint16 = 0
    v7 := GenS1{F0: 886}
    v8 := GenS1{F0: 2}
    v2 = v2
    v1 = append(v1, 2147483648)
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS2{F0: 415, F1: 707}
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 uint64 = (v3 + v3)
            if len(v1) > 0 {
                v1[0] = 2
            }
            _ = fi11
            _ = v12
        }
        v5 -= (v5 * (^(v5)))
        _ = fi9
        _ = v10
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        if false {
            var v14 string = v2
            fmt.Println(v2, v14, v3)
            _ = v14
        }
        if false {
            fmt.Println(v2)
            fmt.Println(v5, fi13, v6)
            if (!false) {
                var v15 uint16 = 585
                switch (fi13 & fi13) {
                case 798:
                    var v16 uint16 = 63
                    _ = v16
                default:
                    v6 -= v15
                }
                _ = v15
            } else {
                if (950 != -32768) {
                    if true {
                        if false {
                            fmt.Println(fi13)
                            fmt.Println(fi13, v2, fi13)
                        }
                        for fi17 := 0; fi17 < 8; fi17++ {
                            var v18 uint32 = 42
                            v6 -= (v6 >> ((135) & 63))
                            _ = fi17
                            _ = v18
                        }
                        fmt.Println(v4, v5)
                    } else {
                        fmt.Println(fi13, v5)
                        if false {
                            v2 = v2
                        }
                        var v19 string = (v2 + v2)
                        _ = v19
                    }
                }
                var v20 uint = 201
                _ = v20
            }
        }
        v21 := GenS3{F0: v4, F1: 681, F2: 32767}
        _ = fi13
        _ = v21
    }
    var v22 uint = 1
    v23 := GenS4{F0: -939}
    var v24 uint8 = v4
    var v25 string = (((("" + (v2 + (v2 + ("foo bar" + "hello")))) + "a") + (v2 + "")) + ("" + v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
