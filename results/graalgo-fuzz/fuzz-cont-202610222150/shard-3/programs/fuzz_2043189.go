package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint32 }
type GenS2 struct { F0 int; F1 int16; F2 int64 }
type GenS3 struct { F0 int64; F1 int8; F2 bool }

func main() {
    var v1 int16 = -32768
    var v2 string = ("\n\t" + "hello")
    var v3 int32 = -89
    var v4 uint = 255
    var v5 uint64 = 18446744073709551615
    var v6 bool = false
    var v7 uint = 4294967295
    if v6 {
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS1{F0: (v5 ^ v5), F1: "foo bar", F2: 2}
            v10 := []int{}
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 int64 = 0
        if ("foo bar" != "\n\t") {
            v1 = int16(v5)
        } else {
            v3 = v3
            v11 += (v11 >> ((v4) & 63))
        }
        _ = v11
    } else {
        switch "\n\t" {
        case "a":
            v12 := []int{4294967295, 0, 128, int(v7)}
            _ = v12
        case "foo bar":
            fmt.Println(v3, v3, v5)
        case "\n\t":
            v13 := []int{-2147483648, 32768, 127, 502}
            _ = v13
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            v7 -= (v7 * uint(v5))
            v15 := GenS1{F0: v5, F1: v2, F2: 65535}
            if true {
                fmt.Println(v1, v5)
                v16 := GenS1{F0: 4294967295, F1: v2, F2: uint32(v5)}
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        v6 = (4294967295 == 0)
    }
    var v17 uint64 = v5
    v18 := GenS1{F0: v17, F1: "", F2: 0}
    if (0 > 129) {
        v4 = (v4 / ((v4) | 1))
    } else {
        for fi19 := 0; fi19 < 12; fi19++ {
            v1 += v1
            _ = fi19
        }
        var v20 bool = v6
        v20 = true
        _ = v20
    }
    switch -9223372036854775808 {
    case 4294967295:
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS2{F0: fi21, F1: 118, F2: 1}
            _ = fi21
            _ = v22
        }
    case 32767:
        var v23 int32 = (^(v3))
        _ = v23
    }
    v24 := []int{}
    v25 := GenS2{F0: -924, F1: v1, F2: 1}
    if len(v24) > 0 {
        v24[0] = 127
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        switch (fi26 - fi26) {
        case 2147483647:
            var v27 uint64 = v17
            _ = v27
        case 42:
            v28 := GenS1{F0: 2, F1: (((("\n\t" + "foo bar") + "foo bar") + ((v2 + v2) + (v2 + ""))) + v2), F2: 2147483648}
            _ = v28
        default:
            _ = len(v24)
        }
        var v29 int32 = (-(v3))
        _ = fi26
        _ = v29
    }
    var v30 int64 = 0
    fmt.Println(v5, v5)
    var v31 int64 = 693
    fmt.Println(v5, v6)
    v24 = append(v24, -460)
    if false {
        v17 += uint64(v7)
        var v32 uint64 = v17
        if len(v24) > 0 {
            v24[0] = 9223372036854775807
        }
        _ = v32
    } else {
        v33 := GenS3{F0: v30, F1: 1, F2: true}
        _ = v33
    }
    var v34 float64 = 0.1
    v35 := GenS1{F0: (v5 & v5), F1: "", F2: 624}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
