package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint8 }
func genFn0(p0 int32) bool {
    v5 := []int{-861, -130}
    var v6 uint8 = 0
    v6 -= v6
    _ = p0
    _ = v5
    _ = v6
    return ((0.0 < 3.14159) && (true || (32768 <= 42)))
}

func main() {
    var v1 uint = 1
    var v2 float32 = 1e10
    v3 := []int{4294967295, -9223372036854775808, 2147483647}
    var v4 uint = (^(v1))
    if (!(!(0 != 203))) {
        v1 = (^(v4))
    } else {
        if (1 >= 4294967295) {
            for fi7 := 0; fi7 < 3; fi7++ {
                v8 := GenS1{F0: true, F1: fi7, F2: 1}
                switch fi7 {
                case 128:
                    _ = len(v3)
                default:
                    if len(v3) > 0 {
                        v3[0] = fi7
                    }
                }
                _ = fi7
                _ = v8
            }
            v9 := GenS2{F0: 42, F1: 535, F2: 127}
            switch (("hello" + ((("" + ("foo bar" + "a")) + "hello") + "\n\t")) + "") {
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = 114
                }
            case "hello":
                v1 = (v1 >> ((v1) & 63))
            case "\n\t":
                fmt.Println(v1)
            }
            _ = v9
        }
        var v10 uint = (v1 * 785)
        v11 := GenS2{F0: 779, F1: 65535, F2: 0}
        _ = v10
        _ = v11
    }
    if true {
        var v12 int32 = -1
        _ = v12
    }
    v4 = v4
    for fi13 := 0; fi13 < 3; fi13++ {
        fmt.Println(v2, fi13)
        for fi14 := 0; fi14 < 13; fi14++ {
            fmt.Println(v2)
            v15 := GenS2{F0: 2, F1: 971, F2: 255}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    if true {
        fmt.Println(v4, v2)
    } else {
        _ = len(v3)
        fmt.Println(v4, v4, v4)
        if len(v3) > 0 {
            v3[0] = -41
        }
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        v4 += v4
        v1 = 662
        _ = fi16
    }
    var v17 bool = (true || (!(true && false)))
    v4 -= 18446744073709551615
    v17 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
