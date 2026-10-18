package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 string }

func main() {
    var v1 uint64 = 947
    var v2 uint8 = 87
    v3 := make(map[uint32]int)
    var v4 uint64 = v1
    v5 := make(map[uint32]string)
    v6 := []int{127, -2147483648, 128}
    var v7 int64 = 1
    if (!(42 != -32768)) {
        if false {
            v7 = v7
        }
    }
    switch 1 {
    case -961:
        v2 = v2
    case -1:
        for fi8 := 0; fi8 < 14; fi8++ {
            v7 = int64(fi8)
            _ = fi8
        }
    default:
        var v9 uint8 = (^(v2))
        _ = v9
    }
    v2 -= (v2 - v2)
    var v10 uint16 = 105
    var v11 uint = 42
    v12 := GenS1{F0: ("foo bar" + "\n\t")}
    if true {
        var v13 int32 = 2147483647
        v14 := GenS1{F0: "\n\t"}
        _ = v13
        _ = v14
    } else {
        switch 42 {
        case 1:
            delete(v5, 2147483647)
        default:
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 uint = (v11 & (v11 << (((v11 << ((uint(v10)) & 63))) & 63)))
                var v17 uint32 = 2
                switch (-(fi15)) {
                case 974:
                    v18 := GenS2{F0: v2}
                    _ = v18
                case 4294967295:
                    _ = len(v6)
                case -820:
                    switch "a" {
                    case "hello":
                        v19 := GenS1{F0: ("foo bar" + ("" + ("foo bar" + "a")))}
                        _ = v19
                    case "\n\t":
                        var v20 uint32 = (v17 / ((v17) | 1))
                        _ = v20
                    case "":
                        v21 := GenS1{F0: "foo bar"}
                        _ = v21
                    default:
                        var v22 int16 = -720
                        _ = v22
                    }
                }
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        var v23 int8 = 127
        _ = v23
    }
    v7 -= v7
    if (true || true) {
        fmt.Println(v11, v7, v10)
        switch ("hello" + "a") {
        case "foo bar":
            v4 = 1
        case "\n\t":
            fmt.Println(v2)
        default:
            v24 := GenS3{F0: false, F1: (("\n\t" + ("" + ("\n\t" + "foo bar"))) + (("foo bar" + "a") + "foo bar"))}
            _ = v24
        }
        var v25 int64 = -9223372036854775808
        _ = v25
    }
    var v26 int = 0
    var v27 uint8 = v2
    v26 += int(v7)
    fmt.Println(v2, v26, v7)
    var v28 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
