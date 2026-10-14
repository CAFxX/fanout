package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint32 = 65535
    var v2 uint = 896
    v3 := make(map[uint32]int)
    var v4 int64 = -1
    var v5 int16 = -1
    var v6 int8 = 95
    var v7 uint8 = 128
    fmt.Println(v2, v5)
    if (!false) {
        var v8 string = "\n\t"
        _ = v8
    } else {
        v9 := GenS1{F0: 18446744073709551615, F1: v2, F2: 2}
        _ = v9
    }
    var v10 int = 128
    var v11 uint16 = 255
    v12 := GenS1{F0: (^(v2)), F1: v2, F2: 4294967295}
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS1{F0: v2, F1: uint(v1), F2: 42}
        _ = fi13
        _ = v14
    }
    var v15 uint64 = 18446744073709551615
    var v16 int8 = int8(v1)
    switch ("foo bar" + "\n\t") {
    case "":
        switch "" {
        case "\n\t":
            v17 := GenS2{F0: (v11 & uint16(v2))}
            _ = v17
        case "foo bar":
            var v18 uint64 = 2
            _ = v18
        default:
            switch (^(v10)) {
            case -1:
                v19 := GenS1{F0: 18446744073709551615, F1: v2, F2: (v15 / ((v15) | 1))}
                _ = v19
            case 32767:
                var v20 int32 = 1
                _ = v20
            default:
                v21 := GenS1{F0: (v2 * 1), F1: 18446744073709551615, F2: (v15 / ((1) | 1))}
                _ = v21
            }
        }
    case "hello":
        v1 = (v1 << (((v2 % (((v2 ^ 297)) | 1))) & 63))
    default:
        fmt.Println(v15, v15, v5)
    }
    fmt.Println(v7)
    var v22 uint32 = 4294967295
    if true {
        fmt.Println(v10, v1)
    } else {
        v23 := GenS1{F0: 0, F1: 779, F2: (v15 ^ 2)}
        v2 -= uint(v15)
        _ = v23
    }
    switch (^(v10)) {
    case 572:
        switch "hello" {
        case "foo bar":
            var v24 int32 = 0
            _ = v24
        case "hello":
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v16)
            }
        case "\n\t":
            if (uint64(v1) >= 0) {
                var v25 uint16 = v11
                _ = v25
            }
        default:
            v5 = (v5 << ((v2) & 63))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
