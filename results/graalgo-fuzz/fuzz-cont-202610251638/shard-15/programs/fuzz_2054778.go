package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = false
    var v3 string = ("\n\t" + "a")
    v4 := []int{32768, -1}
    var v5 uint64 = 534
    v6 := []int{1, -163, 1}
    var v7 uint = 65535
    var v8 string = ((("hello" + "\n\t") + "") + v3)
    v5 = (v5 / ((v5) | 1))
    switch 865 {
    case 128:
        var v9 float64 = 0.0
        _ = v9
    }
    if v2 {
        fmt.Println(v8)
        var v10 bool = (500 >= (v7 << (((v7 ^ v7)) & 63)))
        _ = v10
    } else {
        var v11 int64 = 275
        v12 := GenS1{F0: (v3 + ((v8 + "hello") + "\n\t")), F1: (v11 + 9223372036854775807)}
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := GenS1{F0: (v3 + "foo bar"), F1: (v11 << ((v7) & 63))}
            var v15 uint8 = 76
            fmt.Println(v3, v2)
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v11
        _ = v12
    }
    switch 628 {
    case 414:
        v5 = (v5 ^ (^(v5)))
    case 705:
        _ = len(v1)
    case -2147483648:
        v16 := GenS1{F0: "foo bar", F1: 42}
        _ = v16
    }
    v7 = v7
    switch v3 {
    case "":
        if v2 {
            var v17 int = -9223372036854775808
            var v18 int64 = 299
            _ = v17
            _ = v18
        } else {
            var v19 int64 = 1
            if (!false) {
                if ((!v2) && v2) {
                    var v20 float64 = 0.1
                    v5 += 0
                    v21 := GenS1{F0: "", F1: 402}
                    _ = v20
                    _ = v21
                } else {
                    v22 := GenS1{F0: "foo bar", F1: -1}
                    v2 = true
                    fmt.Println(v2, v19, v3)
                    _ = v22
                }
            }
            _ = v19
        }
    case "\n\t":
        v8 = ("hello" + ((((v8 + "a") + "") + v3) + v8))
    }
    v23 := GenS2{F0: uint16(v5), F1: 42}
    fmt.Println(v2)
    var v24 string = ((v3 + "hello") + v3)
    fmt.Println(v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
