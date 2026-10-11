package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }
func genFn0(p0 uint, p1 int8, p2 uint32) uint32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := GenS1{F0: 1000, F1: fi4, F2: (p2 / ((738) | 1))}
        _ = fi4
        _ = v5
    }
    switch 226 {
    case 1:
        for fi6 := 0; fi6 < 10; fi6++ {
            p2 = 65535
            var v7 int16 = 42
            _ = fi6
            _ = v7
        }
    case -1:
        if false {
            v8 := GenS1{F0: 646, F1: 0, F2: 1}
            _ = v8
        }
    case 393:
        var v9 int16 = -32768
        _ = v9
    }
    v10 := []int{-9223372036854775808, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p2
}

func main() {
    var v1 uint64 = 4294967295
    v2 := []uint32{449, 2, 469, 997}
    var v3 uint64 = (v1 + v1)
    v11 := GenS1{F0: int16(v3), F1: 42, F2: genFn0(2, -127, 2)}
    for fi12 := 0; fi12 < 20; fi12++ {
        v1 += (^(v1))
        _ = fi12
    }
    fmt.Println(v1)
    v1 = 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "\n\t" {
    case "foo bar":
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS1{F0: -1, F1: (fi13 << ((42) & 63)), F2: 4294967295}
            v1 += 1
            if len(v2) > 0 {
                v2[0] = 0
            }
            _ = fi13
            _ = v14
        }
    case "a":
        if (0 > 2147483648) {
            fmt.Println(v3)
        }
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    var v16 float32 = (((1e10 / (3.14159 / ((3.14159 + (-2.25 + 3.14159)) * (3.14159 / -2.25)))) / (0.0 - 3.14159)) * (0.0 - (1.5 * float32(3.14159))))
    v2 = append(v2, genFn0(18446744073709551615, 127, 669))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
}
