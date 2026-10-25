package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 bool = false
    var v2 uint32 = 774
    var v3 int8 = -122
    var v4 float32 = float32(1.5)
    var v5 uint16 = 255
    fmt.Println(v4, v3)
    if v1 {
        switch 803 {
        case -9223372036854775808:
            var v6 uint = 42
            _ = v6
        case 2147483647:
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 bool = v1
                _ = fi7
                _ = v8
            }
        case 32768:
            v5 = (v5 ^ (v5 * (v5 + v5)))
        }
    }
    v2 = 394
    v9 := []int{32767}
    v10 := GenS1{F0: 269}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v11 uint = 1
    if true {
        _ = len(v9)
    }
    if (-1 != -9223372036854775808) {
        var v12 uint64 = 4294967295
        _ = v12
    }
    v4 += float32(((3.14159 / 1.5) / -2.25))
    var v13 int8 = v3
    switch 474 {
    case -520:
        if true {
            v11 -= (v11 | (v11 + uint(v5)))
            var v14 bool = (42 > 1)
            switch "a" {
            case "a":
                var v15 int = 32767
                _ = v15
            case "hello":
                fmt.Println(v14, v5, v13)
            case "\n\t":
                v5 += 65535
            }
            _ = v14
        }
    }
    if (-9223372036854775808 != 0) {
        fmt.Println(v1)
        v4 += float32(((float64(0.1) - 1.5) + 0.0))
    }
    v9 = append(v9, -648)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
