package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 bool; F1 int8; F2 int64 }
func genFn0(p0 int8, p1 int32, p2 uint64) uint {
    fmt.Println(p0, p2, p0)
    v5 := GenS1{F0: int(p0), F1: 1}
    var v6 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint64, p1 int16, p2 int16) int {
    p1 = 224
    p0 += p0
    v7 := GenS1{F0: 32767, F1: (^(p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]int)
    var v3 int32 = -1
    var v4 uint32 = 0
    v1 = (v1 & 2)
    if false {
        v8 := GenS1{F0: 0, F1: 636}
        var v9 int = -2147483648
        _ = v8
        _ = v9
    }
    v2[v4] = 32768
    switch -9223372036854775808 {
    case 32768:
        v4 = v4
    case 769:
        if (!(2147483647 == v4)) {
            var v10 uint = genFn0(0, v3, 380)
            for fi11 := 0; fi11 < 11; fi11++ {
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 int8 = 127
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
            switch "" {
            case "foo bar":
                if false {
                    var v14 uint32 = v4
                    var v15 int = 32768
                    fmt.Println(v14)
                    _ = v14
                    _ = v15
                }
            default:
                switch genFn1(18446744073709551615, 42, 32767) {
                case -476:
                    fmt.Println(v4)
                default:
                    v10 += (v10 % ((v10) | 1))
                }
            }
            _ = v10
        }
    case 42:
        if (true || true) {
            var v16 uint16 = 14
            v17 := GenS2{F0: (((("foo bar" != "hello") || true) && true) || ((v16 >> ((42) & 63)) == (v16 >> ((401) & 63)))), F1: 42, F2: 42}
            _ = v16
            _ = v17
        } else {
            var v18 int32 = (v3 % (((v3 >> ((126) & 63))) | 1))
            _ = v18
        }
    default:
        var v19 float32 = 3.14159
        _ = v19
    }
    v20 := GenS2{F0: true, F1: 127, F2: 9223372036854775807}
    v4 = v4
    v4 = v4
    var v21 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
