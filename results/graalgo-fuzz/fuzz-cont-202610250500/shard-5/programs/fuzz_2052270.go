package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS2 struct { F0 int; F1 uint; F2 int16 }

func main() {
    var v1 uint16 = 2
    var v2 bool = (32767 != -9223372036854775808)
    v3 := make(map[int]uint32)
    var v4 uint32 = 42
    v5 := []int{9223372036854775807, -691}
    switch -1 {
    case -1:
        fmt.Println(v1, v4)
    case -366:
        fmt.Println(v2, v2)
    default:
        var v6 uint8 = 249
        _ = v6
    }
    v4 += (v4 >> ((255) & 63))
    if ((v1 % ((v1) | 1)) < v1) {
        v1 = 42
        switch 850 {
        case 32767:
            if (128 > 42) {
                v5 = append(v5, 32767)
                var v7 bool = (127 != 42)
                _ = v7
            } else {
                fmt.Println(v4)
                fmt.Println(v2, v2, v2)
            }
        }
        var v8 uint = 18446744073709551615
        _ = v8
    }
    var v9 uint = 934
    var v10 bool = (!(32767 == 42))
    for fi11 := 0; fi11 < 12; fi11++ {
        if false {
            if (false || ("" <= "hello")) {
                switch fi11 {
                case 32768:
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v13 := GenS1{F0: uint16(v9), F1: 807, F2: 42}
                        var v14 uint8 = 1
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case 128:
                    v10 = ((-792 <= -2147483648) || (v9 != (v9 ^ 42)))
                case -441:
                    v15 := GenS1{F0: v1, F1: 1, F2: 12}
                    _ = v15
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
        } else {
            v4 = v4
            v16 := v3[(fi11 + 2147483647)]
            v17 := GenS1{F0: 1, F1: 0, F2: -59}
            _ = v16
            _ = v17
        }
        _ = fi11
    }
    v4 = (v4 >> ((746) & 63))
    switch ("\n\t" + "") {
    case "":
        v5 = append(v5, 1)
    case "a":
        v18 := GenS1{F0: (v1 & v1), F1: 2, F2: 127}
        _ = v18
    case "foo bar":
        v19 := GenS2{F0: 128, F1: (v9 << ((v9) & 63)), F2: int16(v9)}
        _ = v19
    }
    v9 += v9
    fmt.Println(v9, v10, v2)
    for fi20 := 0; fi20 < 12; fi20++ {
        switch "foo bar" {
        case "hello":
            v9 += 18446744073709551615
        default:
            var v21 uint32 = (v4 | (v4 * 2147483648))
            _ = v21
        }
        _ = fi20
    }
    v2 = v2
    var v22 uint8 = 252
    var v23 uint8 = v22
    var v24 uint16 = 2
    var v25 int8 = 122
    var v26 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
