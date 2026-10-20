package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int8; F2 bool }

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = 77
    var v3 int64 = -9223372036854775808
    var v4 bool = false
    v5 := make(map[uint32]int)
    v6 := []uint32{2147483647}
    if (1 == 706) {
        v2 -= (^(v2))
        fmt.Println(v1, v3, v4)
    } else {
        var v7 bool = v4
        _ = v7
    }
    v1 = v1
    if (((!v4) || v4) || v4) {
        v8 := GenS1{F0: 2}
        for fi9 := 0; fi9 < 14; fi9++ {
            v2 += v2
            _ = fi9
        }
        var v10 bool = (!(1 != -32768))
        _ = v10
        _ = v8
    } else {
        var v11 uint32 = 4294967295
        var v12 uint8 = v2
        if v4 {
            v2 += 58
            v5[2147483648] = 860
            if v4 {
                _ = len(v6)
                fmt.Println(v2, v12, v2)
            }
        }
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        if v4 {
            switch ("" + ("foo bar" + "")) {
            case "":
                for fi14 := 0; fi14 < 18; fi14++ {
                    _ = len(v5)
                    _ = fi14
                }
            default:
                v4 = (("foo bar" == ((("a" + "a") + ("foo bar" + "hello")) + "")) || false)
            }
            v6 = append(v6, 42)
        }
        fmt.Println(v3, fi13)
        v15 := v5[1]
        _ = fi13
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := GenS2{F0: false, F1: 0, F2: (-706 >= 2147483647)}
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v4 = ((-9223372036854775808 == 128) || v4)
    fmt.Println(v4)
    var v19 int16 = -774
    var v20 uint8 = v2
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
