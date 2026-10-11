package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 int16 = -55
    var v2 uint = 568
    v3 := []string{}
    var v4 uint32 = 65535
    v2 -= v2
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(v4, v1)
        _ = fi5
    }
    v2 = 65535
    if false {
        v6 := GenS1{F0: -1}
        _ = len(v3)
        _ = v6
    }
    fmt.Println(v2, v2, v1)
    var v7 uint32 = (v4 << ((v2) & 63))
    var v8 bool = (1e10 < (1e10 - (0.0 - 1.5)))
    switch 4294967295 {
    case 9223372036854775807:
        v7 += v4
    }
    if (-2.25 >= 0.0) {
        var v9 bool = v8
        v10 := GenS2{F0: -927, F1: 94, F2: 861}
        _ = v10
        _ = v9
    }
    var v11 int64 = 570
    switch "foo bar" {
    case "foo bar":
        var v12 uint16 = 953
        _ = v12
    case "":
        fmt.Println(v2)
    default:
        fmt.Println(v1)
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 int16 = (v1 / ((v1) | 1))
        var v15 string = "foo bar"
        switch "" {
        case "\n\t":
            v7 += 289
        case "foo bar":
            v7 = (v4 * (^(v7)))
        case "hello":
            if ((!(!v8)) || true) {
                if len(v3) > 0 {
                    v3[0] = ((v15 + ("" + "foo bar")) + v15)
                }
            } else {
                v16 := GenS1{F0: -9223372036854775808}
                v17 := GenS3{F0: 1}
                v15 = ((v15 + v15) + "hello")
                _ = v16
                _ = v17
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        switch "hello" {
        case "\n\t":
            var v19 uint = 1
            _ = v19
        }
        v11 += v11
        var v20 uint32 = v7
        _ = fi18
        _ = v20
    }
    if true {
        if (255 != 0) {
            v21 := GenS4{F0: (v11 - v11)}
            fmt.Println(v8)
            var v22 int8 = -96
            _ = v21
            _ = v22
        }
        fmt.Println(v2)
    } else {
        var v23 int = 32767
        _ = v23
    }
    fmt.Println(v8, v4)
    fmt.Println(v11, v8)
    v8 = (v4 > (v7 >> (((v2 - (v2 % (((v2 / (((v2 >> ((v2) & 63))) | 1))) | 1)))) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
