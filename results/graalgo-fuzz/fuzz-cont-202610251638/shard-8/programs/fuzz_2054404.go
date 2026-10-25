package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int64; F1 string; F2 uint }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 int16 = 59
    v4 := []uint32{179, 0, 4294967295}
    v5 := []string{"", "\n\t", "", ""}
    var v6 int16 = v3
    switch "hello" {
    case "a":
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 bool = false
            _ = fi7
            _ = v8
        }
    }
    if len(v5) > 0 {
        v5[0] = ("" + ((("a" + ("a" + "foo bar")) + "") + "\n\t"))
    }
    fmt.Println(v6, v6)
    _ = len(v1)
    if (-340 > (v6 % ((v3) | 1))) {
        _ = len(v5)
        if false {
            var v9 int64 = -1
            _ = v9
        }
    } else {
        if (127 <= 197) {
            v6 = v6
            v6 -= -929
            fmt.Println(v6, v6)
        }
        fmt.Println(v6, v6)
    }
    var v10 uint32 = 2147483647
    var v11 bool = false
    var v12 int = -386
    v6 = (v6 % ((-32768) | 1))
    v13 := GenS1{F0: 861, F1: 100}
    if (2 > 18446744073709551615) {
        fmt.Println(v6, v6)
    } else {
        switch "foo bar" {
        case "\n\t":
            var v14 int32 = 0
            _ = v14
        case "hello":
            v11 = v11
        case "foo bar":
            switch v12 {
            case 0:
                v2[(v12 + v12)] = "hello"
            case -1:
                v15 := GenS1{F0: 310, F1: 78}
                _ = v15
            }
        default:
            v16 := GenS2{F0: -432, F1: "a", F2: 1}
            _ = v16
        }
        v11 = ((("" > "") && v11) && (56 < 100))
        delete(v2, 4294967295)
    }
    var v17 int16 = v6
    fmt.Println(v3)
    fmt.Println(v3, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
