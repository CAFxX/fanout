package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    var v3 uint64 = 4294967295
    var v4 uint16 = 2
    var v5 uint16 = 65535
    var v6 int32 = -75
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint32 = 2147483648
        _ = fi7
        _ = v8
    }
    v5 += v5
    v6 += v6
    v9 := GenS1{F0: 127, F1: (v6 % ((-2147483648) | 1))}
    var v10 string = ("hello" + ("\n\t" + "a"))
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(v5, fi11)
        _ = len(v1)
        if false {
            v3 = v3
            if false {
                v12 := v1[(fi11 >> ((65535) & 63))]
                var v13 int8 = 42
                v5 += 0
                _ = v12
                _ = v13
            } else {
                fmt.Println(v5, v3)
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 uint8 = 128
                v15 = 127
                var v16 uint16 = (v5 % ((v4) | 1))
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        _ = fi11
    }
    fmt.Println(v3, v2)
    var v17 int = 1
    fmt.Println(v17)
    v18 := GenS1{F0: v17, F1: 1}
    v17 = v17
    fmt.Println(v17)
    v19 := GenS1{F0: -9223372036854775808, F1: (v6 * (v6 ^ (v6 * v6)))}
    fmt.Println(v2, v10, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
