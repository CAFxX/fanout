package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint64, p1 int) uint32 {
    var v7 uint = 65535
    if (65535 <= 2) {
        p1 += -9223372036854775808
        v8 := GenS1{F0: 486, F1: 155}
        p1 -= p1
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return 478
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 uint8 = 216
    v4 := make(map[int]bool)
    var v5 uint64 = 1
    var v6 int16 = 1000
    if (v3 != v3) {
        v9 := v4[42]
        for fi10 := 0; fi10 < 13; fi10++ {
            v9 = (!(-128 <= 6))
            _ = fi10
        }
        _ = v9
    }
    v11 := GenS2{F0: 28, F1: v3}
    var v12 int8 = 1
    v6 -= v6
    switch "\n\t" {
    case "hello":
        fmt.Println(v5)
    case "\n\t":
        v6 += (v6 << ((65535) & 63))
    case "foo bar":
        v5 = (v5 | v5)
    default:
        var v13 string = ("hello" + ((("foo bar" + "hello") + "\n\t") + ("foo bar" + "\n\t")))
        _ = v13
    }
    var v14 float64 = 3.14159
    if ((127 != -2147483648) && (-1 >= 825)) {
        fmt.Println(v14, v5)
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 uint64 = v5
            _ = fi16
            _ = v17
        }
        delete(v2, 1)
        _ = fi15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
