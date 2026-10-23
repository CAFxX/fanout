package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint16, p1 int64) uint32 {
    var v5 int8 = -62
    if ((p1 | p1) >= p1) {
        var v6 int8 = v5
        if false {
            p1 += 0
            switch "a" {
            case "foo bar":
                var v7 uint8 = 1
                _ = v7
            default:
                var v8 uint32 = uint32(p0)
                _ = v8
            }
            v9 := GenS1{F0: 0, F1: 18446744073709551615}
            _ = v9
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return 2147483647
}

func main() {
    var v1 uint = 1
    var v2 uint16 = 42
    v3 := []uint32{853, 2147483647}
    v4 := make(map[string]int)
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 uint = (^(v1))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi10
        _ = v11
    }
    var v12 int64 = 9223372036854775807
    if ("foo bar" < ("\n\t" + "\n\t")) {
        v1 -= (v1 + v1)
        var v13 uint8 = 255
        v14 := GenS1{F0: (v13 * (v13 % ((127) | 1))), F1: 42}
        _ = v13
        _ = v14
    }
    v2 = v2
    v12 -= 42
    var v15 uint32 = 2147483648
    if (true || (-2147483648 >= -717)) {
        if true {
            if false {
                for fi16 := 0; fi16 < 20; fi16++ {
                    v17 := GenS1{F0: 42, F1: 0}
                    _ = fi16
                    _ = v17
                }
                var v18 int = 127
                _ = v18
            } else {
                for fi19 := 0; fi19 < 17; fi19++ {
                    fmt.Println(v12, v12)
                    var v20 int32 = 331
                    _ = fi19
                    _ = v20
                }
            }
            v1 = (v1 / ((v1) | 1))
        } else {
            v21 := GenS1{F0: 241, F1: 0}
            _ = v21
        }
        var v22 int16 = -1
        fmt.Println(v22, v1, v2)
        _ = v22
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
