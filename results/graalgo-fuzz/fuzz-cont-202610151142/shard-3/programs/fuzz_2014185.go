package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 22
    v3 := []int{}
    v4 := make(map[uint32]string)
    var v5 int64 = 1
    var v6 int32 = 2147483647
    switch "\n\t" {
    case "\n\t":
        var v7 int8 = 0
        _ = v7
    default:
        fmt.Println(v2, v5, v2)
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    var v9 uint64 = 619
    var v10 int16 = -396
    var v11 float64 = (0.0 * (float64(1e10) - 1e10))
    fmt.Println(v9)
    fmt.Println(v11, v5)
    if true {
        var v12 int64 = 42
        v13 := GenS1{F0: (v12 == 0), F1: (v9 >> ((v2) & 63)), F2: 2147483648}
        _ = v12
        _ = v13
    }
    var v14 uint16 = 437
    var v15 uint16 = v14
    switch -2147483648 {
    case 2147483647:
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS1{F0: (((((((("hello" + "\n\t") + "hello") + "hello") + (("hello" + "foo bar") + "foo bar")) + "hello") + (("foo bar" + ("hello" + (("foo bar" + "") + ("\n\t" + (("foo bar" + "") + ("" + ("" + "a"))))))) + "foo bar")) + "hello") != ("hello" + ("" + "foo bar"))), F1: v9, F2: 4294967295}
            v10 = v10
            fmt.Println(v11)
            _ = fi16
            _ = v17
        }
    case -9223372036854775808:
        for fi18 := 0; fi18 < 15; fi18++ {
            v19 := GenS1{F0: false, F1: v9, F2: 1}
            var v20 uint8 = 128
            _ = fi18
            _ = v19
            _ = v20
        }
    case -58:
        v11 = v11
    default:
        v21 := GenS1{F0: true, F1: v9, F2: 2147483648}
        _ = v21
    }
    if false {
        switch -939 {
        case 119:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 448:
            fmt.Println(v11)
        case -178:
            var v22 string = "a"
            _ = v22
        default:
            v15 = 2
        }
    } else {
        fmt.Println(v14, v9, v11)
        var v23 int8 = 100
        _ = v23
    }
    var v24 int8 = -92
    var v25 uint64 = (v9 * 1)
    v26 := GenS1{F0: (1 == 0), F1: (v9 ^ v25), F2: 728}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
