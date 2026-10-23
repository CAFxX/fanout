package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 int; F1 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 0
    var v3 float64 = 0.1
    v4 := []string{}
    var v5 float32 = 1.5
    v5 += (float32(((v3 / 0.1) / -2.25)) * v5)
    v5 = -2.25
    var v6 int = -2147483648
    _ = len(v1)
    v6 -= -48
    var v7 int64 = (v2 >> ((255) & 63))
    switch (v6 + v6) {
    case 32767:
        if len(v4) > 0 {
            v4[0] = ("\n\t" + ("foo bar" + ("foo bar" + "\n\t")))
        }
    case -916:
        switch v6 {
        case 668:
            v8 := GenS1{F0: 4294967295, F1: 42, F2: 1}
            _ = v8
        case -1:
            var v9 uint = 263
            _ = v9
        case 1:
            fmt.Println(v5, v2)
        }
    case -397:
        var v10 uint64 = 2
        _ = v10
    }
    var v11 int8 = 42
    fmt.Println(v2, v7, v11)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "a" {
    case "hello":
        v12 := GenS2{F0: (v6 >> ((429) & 63)), F1: 824}
        _ = v12
    case "\n\t":
        for fi13 := 0; fi13 < 15; fi13++ {
            v4 = append(v4, "")
            if false {
                v1[724] = "hello"
                delete(v1, 0)
            }
            var v14 int8 = (^(v11))
            _ = fi13
            _ = v14
        }
    case "foo bar":
        v1[42] = ""
    default:
        fmt.Println(v5, v2)
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(v3)
        _ = fi15
    }
    var v16 uint64 = 42
    var v17 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
