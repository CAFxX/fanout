package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint8 = 246
    var v2 uint32 = 162
    var v3 bool = true
    v4 := make(map[int]string)
    var v5 bool = false
    v6 := []int{42, 32768, -1}
    v7 := GenS1{F0: 0, F1: (v1 ^ 216)}
    v8 := GenS2{F0: 1}
    if true {
        fmt.Println(v2, v5)
        if ((!true) || true) {
            v2 += (v2 * 2147483647)
        }
        switch 32767 {
        case 42:
            var v9 int = 32767
            _ = v9
        case 4294967295:
            fmt.Println(v2, v3, v3)
        case 390:
            v1 = v1
        default:
            v5 = (!(((1 == 316) && true) && false))
        }
    } else {
        if len(v6) > 0 {
            v6[0] = 42
        }
        v10 := GenS1{F0: -128, F1: 255}
        _ = v10
    }
    v2 += 0
    v5 = (1.5 == 0.1)
    switch 32767 {
    case 32768:
        var v11 float32 = (((-2.25 * 0.1) * ((0.0 + ((0.1 - 1.5) + -2.25)) / 1.0)) / 1.0)
        _ = v11
    case -9223372036854775808:
        v1 = v1
    }
    v5 = v3
    var v12 int64 = -9223372036854775808
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 string = "a"
        switch -23 {
        case -941:
            if len(v6) > 0 {
                v6[0] = -839
            }
        case 42:
            v5 = v3
        }
        _ = fi13
        _ = v14
    }
    if (v5 || ("hello" >= "foo bar")) {
        switch ("a" + "") {
        case "":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "hello":
            v1 -= 127
        }
        var v15 uint64 = 42
        for fi16 := 0; fi16 < 2; fi16++ {
            for fi17 := 0; fi17 < 15; fi17++ {
                v3 = false
                _ = fi17
            }
            v2 = v2
            v18 := v4[(fi16 - fi16)]
            _ = fi16
            _ = v18
        }
        _ = v15
    } else {
        var v19 int32 = 631
        _ = v19
    }
    v6 = append(v6, 530)
    for fi20 := 0; fi20 < 17; fi20++ {
        delete(v4, (fi20 >> ((42) & 63)))
        var v21 float32 = 3.14159
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
