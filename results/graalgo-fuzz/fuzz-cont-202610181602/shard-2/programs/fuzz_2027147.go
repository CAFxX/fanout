package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint32 }
func genFn0(p0 int32, p1 bool) int32 {
    var v7 float64 = float64(1.5)
    var v8 float64 = v7
    switch ("foo bar" + "") {
    case "a":
        p1 = p1
    case "\n\t":
        if (!((p0 % ((p0) | 1)) != p0)) {
            v9 := []int{32768, 356}
            v7 -= v8
            _ = v9
        }
    case "foo bar":
        switch "foo bar" {
        case "\n\t":
            v10 := []int{}
            _ = v10
        case "":
            switch "" {
            case "a":
                fmt.Println(v7, p0, v8)
            case "hello":
                fmt.Println(v8)
            default:
                var v11 uint = 1
                _ = v11
            }
        default:
            if p1 {
                var v12 int16 = -1
                _ = v12
            }
        }
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p0 | p0)
}

func main() {
    var v1 int = -428
    v2 := make(map[uint32]bool)
    var v3 int = v1
    var v4 uint16 = 65535
    var v5 uint16 = (v4 & v4)
    var v6 uint64 = 4294967295
    for fi13 := 0; fi13 < 11; fi13++ {
        if (552 > 42) {
            var v14 string = "\n\t"
            v14 = v14
            _ = v14
        } else {
            var v15 uint32 = 2147483648
            v1 -= fi13
            _ = v15
        }
        v16 := GenS1{F0: 751, F1: 886, F2: (v6 >> ((65535) & 63))}
        for fi17 := 0; fi17 < 3; fi17++ {
            fmt.Println(v3)
            v6 -= (v6 ^ (v6 << ((255) & 63)))
            _ = fi17
        }
        _ = fi13
        _ = v16
    }
    if (!(!false)) {
        v1 -= v1
        var v18 int64 = 9223372036854775807
        v19 := GenS1{F0: 18446744073709551615, F1: 0, F2: (v6 * 18446744073709551615)}
        _ = v18
        _ = v19
    } else {
        switch ("hello" + "") {
        case "hello":
            var v20 int32 = genFn0(319, (float64(0.1) > (3.14159 / (1.5 * 1.5))))
            _ = v20
        case "foo bar":
            v4 = 0
        default:
            v21 := GenS2{F0: v6, F1: 1, F2: 2}
            _ = v21
        }
    }
    var v22 float64 = float64(1.5)
    v4 = 261
    var v23 int = -9223372036854775808
    var v24 int8 = int8(v6)
    delete(v2, 469)
    v4 = v4
    delete(v2, 2147483648)
    v24 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
