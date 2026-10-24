package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint8 }

func main() {
    v1 := []string{"a", ("foo bar" + ("\n\t" + "foo bar"))}
    var v2 int = 76
    v3 := []string{"foo bar"}
    var v4 float32 = 3.14159
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v5 uint8 = 42
        var v6 uint16 = 65535
        _ = v5
        _ = v6
    } else {
        if len(v1) > 0 {
            v1[0] = ""
        }
        var v7 int32 = -449
        _ = v7
    }
    v8 := GenS1{F0: 18446744073709551615}
    fmt.Println(v2, v4)
    var v9 uint = 993
    if (0 != 42) {
        var v10 uint16 = 2
        if false {
            switch (v2 * v2) {
            case 757:
                v9 += 18446744073709551615
            default:
                v11 := GenS1{F0: 1}
                _ = v11
            }
        } else {
            var v12 uint8 = 255
            if (false && false) {
                switch "\n\t" {
                case "":
                    var v13 uint32 = 65535
                    _ = v13
                case "a":
                    var v14 int64 = 9223372036854775807
                    _ = v14
                case "hello":
                    var v15 int8 = 42
                    _ = v15
                }
            }
            _ = v12
        }
        switch "" {
        case "foo bar":
            v16 := GenS2{F0: 1, F1: v9, F2: 0}
            _ = v16
        case "a":
            v17 := GenS1{F0: 903}
            _ = v17
        case "\n\t":
            v18 := GenS2{F0: 129, F1: 2, F2: 240}
            _ = v18
        }
        _ = v10
    } else {
        var v19 uint8 = 0
        _ = v19
    }
    var v20 float32 = 1e10
    var v21 uint32 = 2147483647
    fmt.Println(v4)
    var v22 int8 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
