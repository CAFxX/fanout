package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 bool; F1 int; F2 int8 }
func genFn0(p0 int64, p1 int16, p2 uint16) uint32 {
    v5 := GenS1{F0: p2, F1: 42, F2: -354}
    p2 -= 141
    p0 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2147483647
}

func main() {
    var v1 float32 = (0.1 + float32((1e10 / 0.1)))
    var v2 string = ("\n\t" + ("foo bar" + "\n\t"))
    var v3 float64 = 0.0
    v4 := make(map[string]int)
    v2 = "hello"
    var v6 int = 128
    _ = len(v4)
    v3 = (float64(v1) / (float64(float32(0.1)) / v3))
    v7 := GenS1{F0: 510, F1: 65535, F2: 42}
    if false {
        switch (v6 << ((1) & 63)) {
        case -9223372036854775808:
            var v8 bool = true
            _ = v8
        }
    }
    var v9 float32 = float32(((-2.25 / (v3 - (v3 / v3))) + float64(v1)))
    v9 = v9
    v1 = (v9 / 1.5)
    if true {
        switch 743 {
        case -9223372036854775808:
            if (v6 <= v6) {
                v10 := []int{-662, (v6 | (v6 * (v6 << ((42) & 63))))}
                _ = v10
            } else {
                fmt.Println(v1)
                v11 := GenS1{F0: 1, F1: 2, F2: 42}
                _ = v11
            }
        case 32768:
            var v12 int64 = -1
            _ = v12
        default:
            delete(v4, v2)
        }
    } else {
        fmt.Println(v1, v6, v6)
        if true {
            var v13 int16 = -32768
            _ = v13
        } else {
            v6 = -2147483648
            v2 = "a"
            var v14 float64 = v3
            _ = v14
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            switch v2 {
            case "foo bar":
                v16 := GenS2{F0: (true || true), F1: 32767, F2: -1}
                _ = v16
            case "a":
                v2 = v2
            case "\n\t":
                v2 = (v2 + v2)
            default:
                if false {
                    delete(v4, v2)
                    switch v2 {
                    case "hello":
                        fmt.Println(v1, v3)
                    default:
                        fmt.Println(v6, v6)
                    }
                } else {
                    for fi17 := 0; fi17 < 10; fi17++ {
                        v18 := []int{}
                        _ = fi17
                        _ = v18
                    }
                    v1 = v1
                    fmt.Println(v6, v2, v6)
                }
            }
            _ = fi15
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
