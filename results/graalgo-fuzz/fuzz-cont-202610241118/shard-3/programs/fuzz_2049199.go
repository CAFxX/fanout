package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }

func main() {
    var v1 int32 = 42
    var v2 string = "foo bar"
    v3 := make(map[uint32]int)
    var v4 uint = 255
    v5 := []int{128, 32768, -1, 77}
    switch v2 {
    case "a":
        fmt.Println(v2, v1, v2)
    case "foo bar":
        var v6 int16 = 775
        _ = v6
    case "hello":
        var v7 uint16 = 0
        _ = v7
    default:
        switch 128 {
        case 9223372036854775807:
            fmt.Println(v4, v2)
        }
    }
    v1 = v1
    var v8 int8 = 127
    fmt.Println(v2, v2, v2)
    v2 = v2
    v1 = (v1 + (v1 * -1))
    switch -2147483648 {
    case 128:
        v4 = 479
    case 56:
        var v9 int64 = 0
        _ = v9
    case 9223372036854775807:
        v1 += (v1 << (((v4 ^ v4)) & 63))
    default:
        var v10 string = ((v2 + v2) + ((("hello" + v2) + "\n\t") + v2))
        _ = v10
    }
    if (!false) {
        switch -433 {
        case 20:
            v1 = 42
        }
    } else {
        v11 := GenS1{F0: v4, F1: true, F2: 0}
        var v12 int64 = int64(v4)
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(fi13, v4, v1)
        fmt.Println(v8)
        if (!false) {
            v4 -= (v4 + (v4 + v4))
        } else {
            for fi14 := 0; fi14 < 3; fi14++ {
                if (18446744073709551615 != 42) {
                    v5 = append(v5, (fi13 ^ (fi13 | fi13)))
                }
                v4 += 4294967295
                var v15 uint32 = 2147483647
                _ = fi14
                _ = v15
            }
        }
        _ = fi13
    }
    v16 := GenS1{F0: 42, F1: (((!((v1 == -2147483648) || true)) && (!(!(!true)))) || true), F2: 65535}
    if (true && false) {
        var v17 string = ("\n\t" + "\n\t")
        _ = v17
    } else {
        v18 := v3[65535]
        fmt.Println(v2, v8, v4)
        for fi19 := 0; fi19 < 6; fi19++ {
            v1 = -985
            _ = fi19
        }
        _ = v18
    }
    var v20 float32 = 0.0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v21 := v3[0]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
