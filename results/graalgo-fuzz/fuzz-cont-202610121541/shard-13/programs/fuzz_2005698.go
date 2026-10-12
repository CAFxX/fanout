package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    var v3 uint = 490
    var v4 int64 = 111
    v5 := make(map[int]uint32)
    v4 -= (v4 << ((0) & 63))
    if true {
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 int64 = (v4 - 9223372036854775807)
            v4 -= 617
            _ = fi6
            _ = v7
        }
    }
    v4 += v4
    var v8 bool = false
    fmt.Println(v8, v4)
    if true {
        v8 = (("\n\t" + "foo bar") <= "foo bar")
        switch "" {
        case "a":
            v9 := GenS1{F0: 1}
            _ = v9
        default:
            fmt.Println(v8, v8)
        }
    }
    v8 = v8
    if (353 >= 2147483647) {
        if true {
            var v10 int16 = 32767
            v4 = (v4 * v4)
            _ = v10
        } else {
            v3 -= 320
            v8 = true
            if false {
                switch 4294967295 {
                case 879:
                    var v11 uint = v3
                    _ = v11
                default:
                    v8 = (!(-1 != 32767))
                }
                var v12 int16 = 1
                for fi13 := 0; fi13 < 13; fi13++ {
                    for fi14 := 0; fi14 < 19; fi14++ {
                        var v15 uint64 = 18446744073709551615
                        _ = fi14
                        _ = v15
                    }
                    v16 := GenS1{F0: v4}
                    switch "a" {
                    case "foo bar":
                        v12 = 1000
                    case "a":
                        var v17 uint8 = 128
                        _ = v17
                    }
                    _ = fi13
                    _ = v16
                }
                _ = v12
            } else {
                var v18 int32 = 0
                fmt.Println(v18, v18)
                _ = v18
            }
        }
        var v19 float32 = ((0.1 / 3.14159) / 1.0)
        switch "\n\t" {
        case "a":
            var v20 int64 = (v4 % ((v4) | 1))
            _ = v20
        case "":
            if v8 {
                v8 = ((0 > 42) || ((v4 << ((65) & 63)) <= v4))
            }
        default:
            v19 += 3.14159
        }
        _ = v19
    } else {
        v21 := v5[1]
        v4 -= v4
        v2[(v21 ^ v21)] = 2147483647
        _ = v21
    }
    v4 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
