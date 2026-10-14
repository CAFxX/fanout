package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint8 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]int)
    var v3 bool = (32767 == 1000)
    v4 := []int{}
    var v5 uint64 = 1
    switch -1 {
    case -1:
        v6 := GenS1{F0: 2, F1: -1, F2: -372}
        _ = v6
    case 0:
        v7 := GenS1{F0: 209, F1: -9223372036854775808, F2: 278}
        _ = v7
    case 32768:
        if (!(802 <= 93)) {
            for fi8 := 0; fi8 < 11; fi8++ {
                switch "hello" {
                case "\n\t":
                    for fi9 := 0; fi9 < 20; fi9++ {
                        switch (fi8 ^ fi8) {
                        case 2147483647:
                            v5 += v5
                        case 343:
                            v4 = append(v4, (fi8 >> ((177) & 63)))
                        case -1:
                            delete(v2, 65535)
                        }
                        var v10 int32 = -700
                        _ = fi9
                        _ = v10
                    }
                case "hello":
                    v3 = v3
                case "":
                    if v3 {
                        var v11 uint = 4294967295
                        v4 = append(v4, (fi8 % ((fi8) | 1)))
                        _ = v11
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v3)
                    }
                }
                _ = fi8
            }
            var v12 bool = v3
            _ = v12
        }
    }
    switch 32768 {
    case 4294967295:
        fmt.Println(v3, v3)
    }
    v3 = true
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int64 = -9223372036854775808
        fmt.Println(fi13)
        _ = fi13
        _ = v14
    }
    v1 += v1
    var v15 uint16 = 1
    if v3 {
        if len(v4) > 0 {
            v4[0] = -2147483648
        }
        var v16 uint16 = v15
        _ = v16
    } else {
        switch "a" {
        case "":
            for fi17 := 0; fi17 < 16; fi17++ {
                v5 -= 0
                fmt.Println(fi17, fi17)
                var v18 uint64 = (v5 * 2)
                _ = fi17
                _ = v18
            }
        case "hello":
            var v19 int32 = 170
            _ = v19
        case "a":
            fmt.Println(v3, v5, v15)
        default:
            fmt.Println(v15, v5)
        }
        var v20 int = 0
        _ = v20
    }
    v21 := GenS2{F0: 32768}
    fmt.Println(v5, v3, v1)
    if v3 {
        delete(v2, 508)
        var v22 bool = (255 > 0)
        var v23 int8 = -1
        _ = v22
        _ = v23
    } else {
        v24 := GenS2{F0: -811}
        v4 = append(v4, 2147483647)
        _ = len(v4)
        _ = v24
    }
    switch 1 {
    case -2147483648:
        var v25 string = "foo bar"
        _ = v25
    }
    v26 := GenS3{F0: 65535, F1: (^(v5)), F2: 80}
    v3 = v3
    v5 += (v5 >> ((767) & 63))
    v1 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
