package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint, p1 int32) uint32 {
    v4 := GenS1{F0: (!(255 == 1)), F1: 42}
    switch -797 {
    case 42:
        switch 9223372036854775807 {
        case 1:
            switch (("hello" + "foo bar") + "hello") {
            case "a":
                if ((127 < 900) && true) {
                    v5 := []int{2147483647, 430, 2147483647, 783}
                    var v6 float64 = float64(0.0)
                    if (!true) {
                        var v7 string = "foo bar"
                        var v8 int64 = -9223372036854775808
                        _ = v7
                        _ = v8
                    }
                    _ = v5
                    _ = v6
                }
            default:
                v9 := []int{32768, -754, 42}
                _ = v9
            }
        case -113:
            switch -2147483648 {
            case 343:
                if true {
                    var v10 uint32 = 749
                    _ = v10
                }
            case 1:
                var v11 bool = false
                _ = v11
            case 19:
                var v12 int = 1
                _ = v12
            }
        case 334:
            v13 := []int{}
            _ = v13
        }
    case -9223372036854775808:
        v14 := []int{9223372036854775807, 42, -358, -1}
        _ = v14
    default:
        v15 := []int{-1}
        _ = v15
    }
    var v16 uint8 = 1
    switch (("" + "a") + "") {
    case "\n\t":
        var v17 uint8 = v16
        _ = v17
    case "foo bar":
        var v18 bool = (2 >= 2)
        _ = v18
    case "":
        v19 := GenS1{F0: false, F1: -1}
        _ = v19
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v16
    _ = v4
    return 65535
}

func main() {
    var v1 bool = false
    var v2 uint16 = 1
    v3 := []uint32{2147483648, 959}
    v2 = 42
    switch "" {
    case "hello":
        v20 := GenS2{F0: 177, F1: -63}
        _ = v20
    case "":
        v1 = (!((899 <= 255) || (v1 && v1)))
    default:
        var v21 int8 = -99
        _ = v21
    }
    var v22 uint16 = (v2 / (((v2 >> ((18446744073709551615) & 63))) | 1))
    v1 = (!true)
    v22 += v22
    var v23 int8 = 0
    var v24 int32 = 0
    v23 -= 42
    var v25 uint64 = 2
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 int64 = 9223372036854775807
        _ = fi26
        _ = v27
    }
    v2 = (v2 ^ v2)
    var v28 uint32 = 65535
    switch "foo bar" {
    case "hello":
        if len(v3) > 0 {
            v3[0] = (v28 + 65535)
        }
    case "a":
        var v29 uint16 = (^(v22))
        _ = v29
    case "":
        var v30 bool = (42 > -730)
        _ = v30
    }
    var v31 int8 = 42
    v25 = v25
    for fi32 := 0; fi32 < 20; fi32++ {
        switch ("\n\t" + ("foo bar" + ("" + "foo bar"))) {
        case "":
            v33 := GenS3{F0: v25}
            _ = v33
        case "foo bar":
            v3 = append(v3, 660)
        case "hello":
            v28 -= (v28 & (v28 ^ genFn0(884, -2147483648)))
        default:
            v3 = append(v3, (v28 * 2147483648))
        }
        _ = fi32
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        v2 = (^(v2))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi34
    }
    var v35 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v31
    _ = v35
}
