package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int16; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint }
func genFn0() int64 {
    var v4 float32 = (1e10 * -2.25)
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 int8 = -128
        _ = fi5
        _ = v6
    }
    v4 -= v4
    _ = v4
    return 149
}
func genFn1() uint32 {
    var v7 uint32 = 2
    var v8 uint16 = 34
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 int = -9223372036854775808
    var v2 float32 = 0.1
    var v3 uint64 = 4294967295
    if true {
        if true {
            switch "" {
            case "a":
                v9 := GenS1{F0: 127}
                _ = v9
            case "":
                var v10 uint32 = uint32(v3)
                _ = v10
            case "\n\t":
                var v11 uint64 = 2
                _ = v11
            default:
                fmt.Println(v1, v1, v3)
            }
            switch "hello" {
            case "\n\t":
                fmt.Println(v3, v3)
            }
        } else {
            v12 := GenS2{F0: 2, F1: v1}
            _ = v12
        }
        var v13 int = (v1 & v1)
        v14 := GenS1{F0: -128}
        _ = v13
        _ = v14
    } else {
        switch "\n\t" {
        case "foo bar":
            var v15 uint8 = 42
            _ = v15
        default:
            switch v1 {
            case -9223372036854775808:
                var v16 int16 = 412
                _ = v16
            case 0:
                v3 = (v3 << ((65535) & 63))
            case 42:
                var v17 uint64 = 18446744073709551615
                _ = v17
            default:
                fmt.Println(v3, v3)
            }
        }
        switch "hello" {
        case "\n\t":
            v18 := []int{v1}
            _ = v18
        case "foo bar":
            for fi19 := 0; fi19 < 1; fi19++ {
                v20 := GenS1{F0: -115}
                v3 += (v3 - 4294967295)
                v1 -= -906
                _ = fi19
                _ = v20
            }
        case "hello":
            v21 := []int{(v1 ^ v1), (v1 << ((277) & 63)), 1}
            _ = v21
        default:
            var v22 uint = 65535
            _ = v22
        }
        fmt.Println(v1)
    }
    v2 = (v2 * v2)
    v23 := []int{9223372036854775807}
    switch "" {
    case "\n\t":
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
    default:
        _ = len(v23)
    }
    v1 = 1
    v1 -= (v1 % ((-22) | 1))
    switch ((("foo bar" + "a") + "a") + "a") {
    case "\n\t":
        if false {
            fmt.Println(v2, v1, v3)
        }
    case "foo bar":
        fmt.Println(v1, v2, v2)
    default:
        switch (v1 << ((459) & 63)) {
        case 128:
            switch (v1 / ((v1) | 1)) {
            case 893:
                var v24 int8 = 1
                _ = v24
            case 466:
                var v25 string = "hello"
                _ = v25
            case 539:
                var v26 int32 = -2147483648
                _ = v26
            }
        default:
            _ = len(v23)
        }
    }
    v27 := GenS1{F0: 127}
    v2 = v2
    v3 += (v3 * (v3 % ((0) | 1)))
    var v28 int64 = 480
    v29 := GenS3{F0: -32768, F1: 32767, F2: 100}
    v30 := GenS4{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
