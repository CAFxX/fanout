package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint64 }

func main() {
    v1 := []uint32{2, 408, 2147483648, 42}
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    v5 := []uint32{492, 994}
    var v6 int64 = 0
    var v7 string = "\n\t"
    if true {
        v6 -= -100
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(v6, fi8, v7)
            _ = fi8
        }
        v7 = "\n\t"
    }
    var v9 int64 = (v6 << ((298) & 63))
    if ("hello" > "foo bar") {
        for fi10 := 0; fi10 < 5; fi10++ {
            switch ("" + v7) {
            case "a":
                switch ("foo bar" + v7) {
                case "\n\t":
                    fmt.Println(fi10, v9, fi10)
                case "foo bar":
                    v6 += -802
                case "hello":
                    v11 := GenS1{F0: 18446744073709551615, F1: -106, F2: 29}
                    _ = v11
                }
            case "hello":
                v12 := GenS1{F0: 58, F1: -1, F2: 18446744073709551615}
                _ = v12
            }
            _ = len(v1)
            v7 = (("" + "\n\t") + "a")
            _ = fi10
        }
        var v13 string = "foo bar"
        if true {
            var v14 int32 = -326
            _ = v14
        } else {
            v1 = append(v1, 2147483648)
            if (false && false) {
                var v15 int8 = 42
                _ = v15
            }
        }
        _ = v13
    } else {
        if (0.0 == -2.25) {
            delete(v3, "hello")
        }
    }
    switch -547 {
    case 2147483647:
        v9 = v6
    }
    switch -2147483648 {
    case 533:
        var v16 int64 = v6
        _ = v16
    default:
        if true {
            v6 = -595
        } else {
            v9 -= 299
        }
    }
    v17 := GenS1{F0: 65535, F1: -1, F2: 1}
    var v18 uint16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
