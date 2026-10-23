package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int16 }
type GenS2 struct { F0 uint; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint64, p1 uint32, p2 uint64) int16 {
    if (!false) {
        fmt.Println(p0)
    }
    var v7 int = 32768
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 0
    var v2 string = "a"
    var v3 int16 = -32768
    var v4 int8 = -128
    v5 := make(map[uint32]string)
    var v6 uint16 = v1
    v5[253] = v2
    v9 := GenS1{F0: 693, F1: -9223372036854775808, F2: (v3 + int16(v6))}
    var v10 int8 = (v4 & v4)
    v2 = v2
    var v11 bool = false
    switch v2 {
    case "":
        v12 := GenS1{F0: 435, F1: 32768, F2: genFn0(0, 0, 748)}
        _ = v12
    case "hello":
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 int32 = -90
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS1{F0: 422, F1: (fi15 * int(v14)), F2: (^(v3))}
                fmt.Println(v3, fi15)
                _ = fi15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
    }
    switch v2 {
    case "":
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := GenS2{F0: 18446744073709551615, F1: 0, F2: 18446744073709551615}
                var v20 int = 42
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi17
        }
    case "\n\t":
        v5[4294967295] = v2
    case "foo bar":
        if ((127 <= 127) || (776 < 0)) {
            v21 := GenS3{F0: 150, F1: 127}
            v22 := GenS3{F0: 1, F1: 0}
            v10 = 42
            _ = v21
            _ = v22
        } else {
            var v23 float64 = (1.5 - 3.14159)
            fmt.Println(v3)
            _ = v23
        }
    default:
        v3 -= 0
    }
    v5[65535] = v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
