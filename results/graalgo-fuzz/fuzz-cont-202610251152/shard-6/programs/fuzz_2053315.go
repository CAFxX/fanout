package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint }

func main() {
    var v1 int16 = -652
    var v2 bool = (0 <= 42)
    var v3 int16 = v1
    var v4 int64 = -1
    var v5 int = -886
    v6 := make(map[int]bool)
    if (100 > 47) {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int = v5
            if (2 > 2147483648) {
                v9 := GenS1{F0: true, F1: 197, F2: 0}
                switch v8 {
                case 863:
                    for fi10 := 0; fi10 < 17; fi10++ {
                        v4 -= v4
                        fmt.Println(fi10, v3, v2)
                        _ = fi10
                    }
                }
                var v11 int16 = 32767
                _ = v11
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    }
    var v12 string = (("hello" + ("" + "\n\t")) + "\n\t")
    var v13 int64 = -1
    var v14 uint8 = 127
    v13 -= v4
    v5 = (v5 << ((42) & 63))
    var v15 uint8 = 128
    v4 -= (v4 * -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
