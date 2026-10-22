package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint16 }

func main() {
    var v1 string = ("foo bar" + "\n\t")
    var v2 uint16 = 417
    v3 := []string{v1, ((v1 + "foo bar") + ""), v1, v1}
    var v4 float32 = 1e10
    v5 := make(map[uint32]int)
    v6 := make(map[int]bool)
    if false {
        var v7 uint64 = 18446744073709551615
        var v8 float32 = v4
        var v9 int8 = -128
        _ = v7
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 59, F1: -9223372036854775808, F2: (v2 + (v2 & (v2 << ((68) & 63))))}
        v4 += ((v4 - (0.1 / (v4 * 0.1))) * (-2.25 * v4))
        _ = v10
    }
    if false {
        if (v4 > v4) {
            var v11 uint32 = 65535
            _ = v11
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            v1 = "hello"
            var v13 uint8 = 0
            _ = fi12
            _ = v13
        }
        var v14 int = 128
        _ = v14
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v15 float32 = v4
        v6[-2147483648] = (false || false)
        _ = v15
    }
    var v16 int64 = 0
    v16 = (v16 * (v16 >> ((255) & 63)))
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS1{F0: -2147483648, F1: v16, F2: v2}
        _ = fi17
        _ = v18
    }
    fmt.Println(v2, v2, v4)
    fmt.Println(v1)
    var v19 bool = ((!false) || false)
    v20 := GenS1{F0: 1, F1: v16, F2: v2}
    v19 = v19
    var v21 int = 0
    switch "hello" {
    case "\n\t":
        v2 -= (v2 >> ((2) & 63))
    case "hello":
        v3 = append(v3, (v1 + "\n\t"))
    case "a":
        v22 := GenS1{F0: -305, F1: v16, F2: 32767}
        _ = v22
    }
    var v23 string = v1
    var v24 bool = v19
    if ((v19 || false) || v19) {
        for fi25 := 0; fi25 < 16; fi25++ {
            v23 = v23
            _ = fi25
        }
    }
    v26 := GenS1{F0: -2147483648, F1: (v16 + v16), F2: 1}
    v24 = v24
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
