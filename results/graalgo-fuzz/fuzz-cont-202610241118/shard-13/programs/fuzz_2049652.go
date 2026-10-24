package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint8 }

func main() {
    var v1 uint = 809
    v2 := []int{2147483647}
    var v3 int64 = 0
    var v4 uint32 = 442
    v3 -= (v3 % (((v3 * (v3 - v3))) | 1))
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(v1, fi6)
            v4 += 563
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi6
        }
        _ = fi5
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v4, v4, v1)
    switch "foo bar" {
    case "hello":
        if (("" + "") <= ("foo bar" + "\n\t")) {
            var v7 uint8 = 218
            if true {
                v8 := GenS1{F0: v7}
                v7 += (v7 % (((v7 * v7)) | 1))
                if len(v2) > 0 {
                    v2[0] = 1
                }
                _ = v8
            } else {
                fmt.Println(v4, v4, v1)
                var v9 float32 = 1e10
                v2 = append(v2, 580)
                _ = v9
            }
            _ = v7
        } else {
            fmt.Println(v1, v3)
        }
    case "\n\t":
        v10 := GenS1{F0: 2}
        _ = v10
    case "foo bar":
        v3 = 9223372036854775807
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v3 += -9223372036854775808
        if (true && false) {
            if false {
                var v12 int = int(v4)
                var v13 uint64 = 42
                v3 = -939
                _ = v12
                _ = v13
            }
            if false {
                var v14 float64 = (1e10 + (0.1 + float64(0.0)))
                var v15 uint32 = v4
                var v16 int16 = -32768
                _ = v14
                _ = v15
                _ = v16
            } else {
                if len(v2) > 0 {
                    v2[0] = fi11
                }
            }
        }
        _ = len(v2)
        _ = fi11
    }
    switch 4294967295 {
    case 127:
        var v17 uint8 = 128
        _ = v17
    case 42:
        var v18 int64 = v3
        _ = v18
    default:
        fmt.Println(v3, v3, v3)
    }
    v1 -= (v1 % ((18446744073709551615) | 1))
    _ = len(v2)
    var v19 uint = (v1 << (((v1 - v1)) & 63))
    for fi20 := 0; fi20 < 5; fi20++ {
        v3 = v3
        v2 = append(v2, (fi20 >> ((800) & 63)))
        _ = fi20
    }
    if true {
        fmt.Println(v1)
        v21 := GenS1{F0: 230}
        _ = v21
    }
    fmt.Println(v1)
    var v22 uint64 = 0
    switch "" {
    case "hello":
        v23 := GenS2{F0: -159, F1: 2}
        _ = v23
    case "foo bar":
        var v24 uint64 = v22
        _ = v24
    }
    var v25 uint64 = (v22 / ((v22) | 1))
    for fi26 := 0; fi26 < 13; fi26++ {
        switch "\n\t" {
        case "hello":
            var v27 uint32 = (v4 / ((65535) | 1))
            _ = v27
        }
        v25 -= (v25 - 4294967295)
        _ = fi26
    }
    switch ("hello" + "\n\t") {
    case "hello":
        v1 -= (v1 ^ v19)
    case "":
        var v28 uint16 = 255
        _ = v28
    default:
        v29 := GenS1{F0: 128}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
