package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int32 }
func genFn0(p0 int16, p1 uint16) int32 {
    switch "foo bar" {
    case "\n\t":
        v7 := GenS1{F0: 255, F1: 949, F2: -543}
        _ = v7
    default:
        var v8 string = ("\n\t" + ("a" + (("hello" + "hello") + ("hello" + ""))))
        _ = v8
    }
    fmt.Println(p1, p1)
    var v9 uint64 = 71
    var v10 int32 = -2147483648
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return v10
}
func genFn1(p0 uint16, p1 uint16, p2 uint64) string {
    var v11 uint64 = (^(p2))
    p0 = (p0 >> ((577) & 63))
    p0 = 413
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return "a"
}

func main() {
    var v1 int8 = 100
    var v2 int64 = 0
    v3 := make(map[uint32]uint32)
    var v4 float32 = 1.5
    v5 := make(map[int]bool)
    var v6 uint32 = 2
    var v12 uint8 = 2
    v3[(v6 >> ((402) & 63))] = v6
    v2 += (v2 + (v2 * v2))
    if false {
        for fi13 := 0; fi13 < 16; fi13++ {
            fmt.Println(v1)
            _ = fi13
        }
        v1 = v1
    }
    switch -1 {
    case 32767:
        switch ("a" + "") {
        case "":
            if false {
                v14 := GenS1{F0: (v12 << ((885) & 63)), F1: 4294967295, F2: 2147483647}
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 int32 = -567
                    _ = fi15
                    _ = v16
                }
                v2 += (v2 % ((v2) | 1))
                _ = v14
            }
        }
    default:
        var v17 float32 = v4
        _ = v17
    }
    fmt.Println(v2, v12, v4)
    fmt.Println(v2, v6)
    fmt.Println(v12)
    var v18 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
