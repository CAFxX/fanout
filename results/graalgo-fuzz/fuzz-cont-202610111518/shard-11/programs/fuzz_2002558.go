package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64 }
func genFn0() uint16 {
    if (203 > 18446744073709551615) {
        var v6 uint = 18446744073709551615
        _ = v6
    }
    var v7 string = "\n\t"
    var v8 string = (v7 + ("foo bar" + v7))
    var v9 int = 38
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint8) uint16 {
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int32 = 1
            p0 = (p0 * p0)
            _ = fi11
            _ = v12
        }
        var v13 int32 = 976
        _ = fi10
        _ = v13
    }
    var v14 int = 880
    if false {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 float32 = 1.5
            _ = fi15
            _ = v16
        }
    }
    v14 += 148
    fmt.Println(p0)
    _ = p0
    _ = v14
    return 1
}

func main() {
    var v1 uint = 255
    v2 := []int{-2147483648, 127, 4294967295, 1}
    var v3 string = "\n\t"
    var v4 uint32 = 2147483647
    v5 := []int{-2147483648, 127, 1, 2147483647}
    v4 -= v4
    v17 := GenS1{F0: 128, F1: (0.0 == 0.0)}
    fmt.Println(v1, v4)
    switch int(v4) {
    case 9223372036854775807:
        v4 = v4
    case 2147483647:
        if false {
            v1 = v1
            v1 += v1
        } else {
            if false {
                fmt.Println(v1)
                for fi18 := 0; fi18 < 1; fi18++ {
                    fmt.Println(fi18)
                    v1 -= v1
                    _ = fi18
                }
            }
        }
    case -971:
        v2 = append(v2, 244)
    }
    v19 := GenS2{F0: v4, F1: 2147483647}
    if true {
        v1 = v1
    }
    v20 := GenS1{F0: 209, F1: false}
    v1 = v1
    v3 = (((v3 + v3) + "a") + ("\n\t" + "foo bar"))
    switch "\n\t" {
    case "\n\t":
        v4 += v4
    case "a":
        v2 = append(v2, 42)
    }
    var v21 int = 4294967295
    if false {
        v22 := GenS2{F0: 42, F1: (v21 | v21)}
        v23 := GenS2{F0: (v4 - 42), F1: v21}
        v4 = (v4 / (((v4 << ((65535) & 63))) | 1))
        _ = v22
        _ = v23
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if true {
        fmt.Println(v21, v1)
    }
    v24 := GenS3{F0: 2}
    fmt.Println(v1, v21, v21)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
