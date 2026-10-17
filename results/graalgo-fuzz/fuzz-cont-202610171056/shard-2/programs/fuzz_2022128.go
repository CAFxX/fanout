package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }

func main() {
    var v1 bool = false
    v2 := []string{("hello" + (("foo bar" + ("hello" + "foo bar")) + "hello")), "foo bar", ("hello" + "hello")}
    var v3 int64 = -96
    var v4 float32 = ((0.0 - -2.25) + 1.5)
    var v5 uint = 2
    v6 := make(map[uint32]uint32)
    switch "\n\t" {
    case "\n\t":
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint64 = 18446744073709551615
            v8 = 627
            _ = fi7
            _ = v8
        }
    case "hello":
        switch 32767 {
        case -217:
            if true {
                if v1 {
                    v4 -= (0.1 - (1.5 / (v4 / v4)))
                }
            } else {
                switch "\n\t" {
                case "hello":
                    if (42 <= 224) {
                        v9 := GenS1{F0: -374, F1: 779}
                        v1 = (-2.25 != (v4 + 1e10))
                        _ = v9
                    }
                }
                v3 -= (v3 * v3)
            }
        }
    default:
        v3 = (v3 & -9223372036854775808)
    }
    var v10 int = 127
    switch v10 {
    case -600:
        v11 := GenS1{F0: 0, F1: 610}
        _ = v11
    case 4294967295:
        v12 := GenS1{F0: 42, F1: 596}
        _ = v12
    case 657:
        switch v10 {
        case 1:
            v4 = 1.5
        default:
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 bool = (int16(v5) != 42)
                if (1 != -2147483648) {
                    var v15 uint8 = 255
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
        }
    default:
        v6[1] = 2
    }
    var v16 uint8 = 149
    var v17 uint16 = 142
    if (v3 >= (v3 % (((v3 ^ (v3 ^ v3))) | 1))) {
        var v18 float32 = v4
        fmt.Println(v1, v4)
        _ = v18
    }
    fmt.Println(v3, v1, v17)
    if v1 {
        switch "a" {
        case "\n\t":
            var v19 string = ""
            _ = v19
        }
        var v20 uint64 = 0
        _ = v20
    }
    v2 = append(v2, "a")
    var v21 int32 = 1
    v22 := v6[65535]
    if ((v3 / ((v3) | 1)) == (v3 / (((-(v3))) | 1))) {
        for fi23 := 0; fi23 < 2; fi23++ {
            fmt.Println(v4)
            _ = fi23
        }
    } else {
        for fi24 := 0; fi24 < 8; fi24++ {
            v21 = v21
            v25 := GenS1{F0: 0, F1: (v22 + uint32(v17))}
            v6[2147483647] = (v22 % (((^(v22))) | 1))
            _ = fi24
            _ = v25
        }
        v16 -= v16
        var v26 int = (v10 / ((-2147483648) | 1))
        _ = v26
    }
    v4 -= 0.0
    fmt.Println(v3, v1)
    var v27 uint8 = v16
    fmt.Println(v1, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
