package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint64, p1 uint32, p2 uint8) uint16 {
    fmt.Println(p2)
    var v6 uint32 = (p1 | p1)
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 bool = true
        _ = fi7
        _ = v8
    }
    var v9 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 308
}
func genFn1() int16 {
    v10 := []int{983, -2147483648, 2147483647, 42}
    var v11 uint8 = 42
    switch "" {
    case "foo bar":
        if ((v11 | 50) <= 200) {
            switch 127 {
            case -265:
                var v12 bool = (genFn0(42, 2147483647, (v11 * (v11 | (v11 % ((v11) | 1))))) >= 42)
                _ = v12
            default:
                var v13 uint = 42
                _ = v13
            }
        } else {
            fmt.Println(v11)
            for fi14 := 0; fi14 < 16; fi14++ {
                for fi15 := 0; fi15 < 2; fi15++ {
                    v16 := GenS1{F0: -1, F1: 0}
                    var v17 uint16 = 42
                    if len(v10) > 0 {
                        v10[0] = 127
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = fi14
            }
        }
    default:
        switch ("" + ("hello" + (("foo bar" + "hello") + ("hello" + "\n\t")))) {
        case "a":
            var v18 int64 = 0
            _ = v18
        case "\n\t":
            v11 += (v11 % (((v11 / ((v11) | 1))) | 1))
        case "hello":
            v19 := GenS2{F0: 42, F1: 9223372036854775807, F2: -424}
            _ = v19
        }
    }
    if true {
        if (!true) {
            if ("hello" != ((((("\n\t" + "hello") + "") + ("hello" + "a")) + ("" + "\n\t")) + (("a" + "") + "foo bar"))) {
                fmt.Println(v11)
            } else {
                var v20 float64 = float64(0.1)
                var v21 float64 = v20
                v11 += 203
                _ = v20
                _ = v21
            }
            var v22 float64 = 1e10
            _ = v22
        } else {
            v11 -= (v11 >> ((42) & 63))
        }
        v23 := GenS3{F0: 2, F1: 0}
        _ = v23
    } else {
        v10 = append(v10, 768)
        fmt.Println(v11)
    }
    fmt.Println(v11)
    _ = v10
    _ = v11
    return -1
}

func main() {
    v1 := []uint32{2147483648, 42, 698}
    var v2 uint8 = 2
    var v3 int8 = 42
    v4 := make(map[uint32]int)
    v5 := []int{128, -66, -9223372036854775808, -9223372036854775808}
    var v24 int = 0
    var v25 int8 = (^(v3))
    fmt.Println(v25, v2, v3)
    var v26 uint32 = 65535
    v24 = v24
    v27 := GenS4{F0: (v2 ^ (v2 - 243))}
    v28 := GenS4{F0: (v2 >> ((65535) & 63))}
    var v29 int64 = int64(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
