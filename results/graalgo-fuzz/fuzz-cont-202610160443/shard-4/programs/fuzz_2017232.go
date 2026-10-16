package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0() int {
    var v5 uint16 = 65535
    v5 += v5
    if (false && (!(!true))) {
        if false {
            v5 += 0
            if true {
                v5 -= 2
                v6 := []int{int(v5), 4294967295}
                _ = v6
            }
            switch "hello" {
            case "\n\t":
                fmt.Println(v5)
            }
        }
        fmt.Println(v5)
    }
    switch ("a" + "\n\t") {
    case "":
        var v7 int8 = 127
        _ = v7
    case "a":
        v8 := []int{4294967295, -928, 32768}
        _ = v8
    case "hello":
        fmt.Println(v5)
    default:
        v5 = v5
    }
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint8 = 42
    var v2 uint8 = 142
    var v3 uint = 18446744073709551615
    v4 := []uint32{2147483647, 65535}
    switch 42 {
    case -91:
        if ((-2.25 - (1.5 / 1.5)) < (1e10 / 1.0)) {
            v9 := GenS1{F0: 32767, F1: 1, F2: 2147483647}
            var v10 int8 = -128
            switch "hello" {
            case "":
                v10 = (v10 - 1)
            case "hello":
                fmt.Println(v1)
            case "\n\t":
                var v11 int16 = 32767
                _ = v11
            default:
                v12 := GenS1{F0: int16(v10), F1: 606, F2: -1}
                _ = v12
            }
            _ = v10
            _ = v9
        }
    case -2147483648:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 32767:
        v2 -= (^(v2))
    default:
        for fi13 := 0; fi13 < 3; fi13++ {
            switch "" {
            case "foo bar":
                var v14 uint64 = 42
                _ = v14
            }
            _ = fi13
        }
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v3 -= (v3 + v3)
        var v16 uint16 = 1
        _ = fi15
        _ = v16
    }
    fmt.Println(v1)
    if (1 < -371) {
        _ = len(v4)
    } else {
        v17 := GenS2{F0: -128, F1: 18446744073709551615}
        _ = v17
    }
    var v18 uint = v3
    switch genFn0() {
    case 42:
        v18 = v3
    default:
        v3 = v18
    }
    var v19 int32 = 0
    v3 = v3
    v2 += v2
    v19 = (v19 & v19)
    if ("\n\t" > "") {
        if false {
            if ((0 == -999) || (-1 >= 0)) {
                var v20 string = "hello"
                fmt.Println(v20, v18)
                _ = v20
            }
        }
        switch "foo bar" {
        case "\n\t":
            v2 -= v1
        case "foo bar":
            switch -356 {
            case 32768:
                var v21 int32 = (-(v19))
                _ = v21
            case -659:
                v3 -= v3
            default:
                _ = len(v4)
            }
        case "":
            var v22 uint64 = 1
            _ = v22
        }
    }
    var v23 string = "a"
    fmt.Println(v23, v18, v19)
    var v24 int64 = -9223372036854775808
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
