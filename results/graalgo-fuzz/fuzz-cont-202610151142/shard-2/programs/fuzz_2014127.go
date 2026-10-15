package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []int{128, -2147483648}
    var v2 uint8 = 128
    v3 := []uint32{65535}
    v4 := []uint32{65535}
    switch "" {
    case "\n\t":
        var v5 int64 = -114
        _ = v5
    case "foo bar":
        if false {
            var v6 int = 32768
            _ = v6
        } else {
            v2 = (v2 - 38)
            if (!(0 <= 128)) {
                for fi7 := 0; fi7 < 9; fi7++ {
                    var v8 int64 = 42
                    _ = fi7
                    _ = v8
                }
                v2 -= (^(v2))
                if (true && false) {
                    v9 := GenS1{F0: ((18446744073709551615 != 140) || false)}
                    _ = v9
                }
            }
        }
    case "a":
        var v10 float64 = float64(float32(0.1))
        _ = v10
    }
    v2 -= v2
    v1 = append(v1, 1)
    v11 := GenS1{F0: true}
    var v12 int64 = 150
    v12 += v12
    v12 = (v12 << ((1) & 63))
    v13 := GenS1{F0: (v2 > v2)}
    var v14 uint32 = 65535
    fmt.Println(v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
