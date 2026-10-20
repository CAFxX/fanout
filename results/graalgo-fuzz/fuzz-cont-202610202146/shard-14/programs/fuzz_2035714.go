package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint64 }
func genFn0(p0 bool, p1 uint64, p2 int) int8 {
    if (2 >= 18446744073709551615) {
        v5 := []int{(p2 / ((0) | 1)), -9223372036854775808, p2, (p2 + 0)}
        var v6 uint = 65535
        _ = v5
        _ = v6
    }
    v7 := []int{p2, p2}
    v7 = append(v7, p2)
    var v8 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}
func genFn1() int {
    var v9 int64 = 9223372036854775807
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(v9, v9)
        var v11 uint32 = 395
        v11 -= 0
        _ = fi10
        _ = v11
    }
    fmt.Println(v9)
    v9 = v9
    _ = v9
    return -432
}

func main() {
    var v1 uint32 = 256
    var v2 uint = 1
    var v3 int16 = 42
    var v4 int32 = 2147483647
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(v3, fi12, v2)
        var v13 uint16 = 2
        switch (-(fi12)) {
        case 32767:
            v14 := []int{(fi12 + (fi12 | 42))}
            _ = v14
        case 1:
            v2 = (v2 + 0)
        case -130:
            for fi15 := 0; fi15 < 5; fi15++ {
                v13 -= 42
                v16 := GenS1{F0: (fi12 << ((v2) & 63)), F1: ((v2 ^ 2) != v2), F2: 0}
                _ = fi15
                _ = v16
            }
        default:
            var v17 float32 = (1e10 / -2.25)
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    var v18 string = (("" + "foo bar") + "foo bar")
    fmt.Println(v3)
    var v19 int = 9223372036854775807
    fmt.Println(v3)
    var v20 int64 = 42
    if (!false) {
        fmt.Println(v4, v19)
        v4 = v4
        v21 := []int{v19, v19, (^(v19)), genFn1()}
        _ = v21
    } else {
        v4 += int32(v2)
    }
    var v22 uint32 = 42
    v19 += (v19 % ((v19) | 1))
    switch (v19 << (((v2 / (((^(v2))) | 1))) & 63)) {
    case -1:
        if (false || false) {
            var v23 int16 = (v3 - int16(v19))
            _ = v23
        } else {
            v20 += v20
        }
    case 32768:
        v20 += (v20 >> ((v2) & 63))
    default:
        switch v19 {
        case 1:
            var v24 uint16 = 32767
            _ = v24
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
}
