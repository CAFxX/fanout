package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    v1 := []uint32{4294967295, 0, 65535}
    var v2 uint32 = 0
    var v3 uint64 = 1
    v4 := make(map[uint32]int)
    var v5 int16 = -32768
    v3 = (v3 - (v3 - v3))
    v1 = append(v1, (v2 & 1))
    var v6 int = -329
    v7 := GenS1{F0: -410}
    switch (("\n\t" + "") + "a") {
    case "hello":
        var v8 uint64 = v3
        _ = v8
    case "\n\t":
        if true {
            var v9 int64 = 42
            _ = v9
        } else {
            v6 = (v6 << ((4294967295) & 63))
        }
    case "foo bar":
        var v10 int16 = (-(v5))
        _ = v10
    default:
        var v11 uint8 = 128
        _ = v11
    }
    var v12 uint32 = (v2 - (^(v2)))
    if (1 >= 224) {
        var v13 int16 = v5
        _ = v13
    }
    v4[(^(v2))] = (v6 - v6)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch v6 {
    case 0:
        var v14 uint8 = 1
        _ = v14
    case -2147483648:
        switch ("hello" + "a") {
        case "foo bar":
            var v15 uint16 = 144
            _ = v15
        case "a":
            fmt.Println(v2, v3, v12)
        default:
            v16 := GenS1{F0: 0}
            _ = v16
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v5, v3, v12)
    v17 := GenS1{F0: -103}
    if (false || false) {
        var v18 int64 = 0
        switch "a" {
        case "":
            fmt.Println(v6)
        case "\n\t":
            v2 -= 2147483648
        default:
            switch (("foo bar" + ("" + "a")) + "\n\t") {
            case "a":
                var v19 int32 = 1
                _ = v19
            case "hello":
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 uint64 = v3
                    v5 -= -140
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi20
                    _ = v21
                }
            }
        }
        _ = v18
    }
    v22 := GenS2{F0: -1, F1: 259}
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
