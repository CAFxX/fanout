package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 int) string {
    var v6 int8 = -44
    v6 = (v6 | 127)
    fmt.Println(v6)
    _ = p0
    _ = v6
    return "\n\t"
}

func main() {
    var v1 uint8 = 161
    var v2 uint64 = 2
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    var v5 int64 = 55
    fmt.Println(v2, v2, v5)
    fmt.Println(v5, v2, v2)
    fmt.Println(v1, v5, v2)
    v1 -= (v1 >> ((1) & 63))
    var v7 int64 = (^(v5))
    v2 -= v2
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: (v5 & 851), F1: ((("a" + "\n\t") != ("hello" + ("\n\t" + "hello"))) && (0 < -2147483648)), F2: -9223372036854775808}
        if false {
            for fi10 := 0; fi10 < 6; fi10++ {
                fmt.Println(v1)
                _ = fi10
            }
        } else {
            v5 = v5
        }
        var v11 int32 = -1
        _ = fi8
        _ = v11
        _ = v9
    }
    if true {
        var v12 int = 2147483647
        _ = v12
    }
    v1 = 127
    v5 = v7
    switch -173 {
    case -780:
        var v13 uint64 = v2
        _ = v13
    case 32768:
        v1 = 1
    case 2147483647:
        v2 -= 18446744073709551615
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        if (true || (true || (!((!false) && (18446744073709551615 >= 18446744073709551615))))) {
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := GenS2{F0: 2147483647, F1: -128, F2: ((v2 | 18446744073709551615) >= 2)}
                _ = fi15
                _ = v16
            }
        } else {
            var v17 uint16 = 395
            v1 -= (v1 - v1)
            _ = v17
        }
        if (!true) {
            v18 := GenS2{F0: 2147483647, F1: 100, F2: true}
            _ = v18
        }
        _ = fi14
    }
    delete(v4, 2147483647)
    v19 := GenS2{F0: -2147483648, F1: -64, F2: (!true)}
    v20 := GenS3{F0: 735}
    v1 = 2
    var v21 float64 = (0.1 - 3.14159)
    var v22 uint8 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
