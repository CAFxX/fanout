package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int32 }
type GenS2 struct { F0 string; F1 string; F2 int16 }
func genFn0(p0 uint64, p1 string, p2 uint8) int32 {
    var v7 string = (p1 + ("hello" + p1))
    p1 = ""
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint8 = 51
    var v2 uint32 = 4294967295
    var v3 bool = true
    var v4 uint8 = v1
    var v5 bool = (1000 != 192)
    var v6 uint32 = (^(v2))
    fmt.Println(v5, v3, v4)
    switch "a" {
    case "\n\t":
        var v8 int32 = -2147483648
        _ = v8
    case "":
        if v3 {
            v9 := GenS1{F0: 1, F1: 823, F2: 0}
            _ = v9
        } else {
            v10 := GenS1{F0: -427, F1: 255, F2: 110}
            v3 = (-1 == int64(v2))
            v11 := GenS1{F0: -1, F1: 0, F2: 42}
            _ = v10
            _ = v11
        }
    case "foo bar":
        fmt.Println(v2)
    default:
        v12 := GenS2{F0: ("\n\t" + (("foo bar" + ("foo bar" + ("foo bar" + "a"))) + "hello")), F1: ("a" + "hello"), F2: 32767}
        _ = v12
    }
    v2 = (^(v6))
    v4 = (v1 << ((299) & 63))
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int = fi13
        v1 = (v4 * (v4 >> ((42) & 63)))
        v15 := []int{(fi13 * (^(fi13))), 667, (v14 << ((42) & 63))}
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v5, v2, v1)
    v16 := []int{-2147483648, 42, 4294967295}
    if len(v16) > 0 {
        v16[0] = -1
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        if v5 {
            var v18 uint = 510
            var v19 int32 = 42
            if (v5 && (4294967295 != fi17)) {
                var v20 bool = v3
                var v21 uint16 = 42
                v2 = 2
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 8; fi22++ {
                    v19 = 42
                    var v23 uint64 = 815
                    v2 = (^(v6))
                    v4 -= (^(v4))
                    if v3 {
                        v24 := GenS1{F0: 9223372036854775807, F1: v18, F2: v19}
                        _ = v24
                    } else {
                        var v25 int8 = 0
                        v23 = v23
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                }
            }
            _ = v18
            _ = v19
        }
        v5 = false
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
