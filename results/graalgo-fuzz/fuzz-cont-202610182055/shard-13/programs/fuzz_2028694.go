package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0(p0 int32, p1 uint16) int8 {
    var v4 int8 = 100
    v5 := GenS1{F0: 4294967295, F1: 65535}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 65
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]string)
    var v3 bool = ((false && true) || (!(!false)))
    var v6 uint8 = 128
    v1 = v1
    delete(v2, 2147483647)
    fmt.Println(v6, v1, v3)
    for fi7 := 0; fi7 < 13; fi7++ {
        v2[449] = "a"
        _ = fi7
    }
    if (!(226 <= 33)) {
        var v8 uint32 = 42
        fmt.Println(v8)
        if (!(!v3)) {
            v8 = v8
            var v9 int32 = -909
            _ = v9
        } else {
            if v3 {
                if (!(891 <= 0)) {
                    for fi10 := 0; fi10 < 20; fi10++ {
                        v11 := v2[2147483647]
                        v12 := GenS1{F0: uint64(v8), F1: 773}
                        fmt.Println(v8, v8)
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    v6 = v6
                }
                v13 := GenS1{F0: 704, F1: 4294967295}
                var v14 int16 = 335
                _ = v13
                _ = v14
            }
        }
        _ = v8
    }
    var v15 int16 = 32767
    for fi16 := 0; fi16 < 17; fi16++ {
        if v3 {
            if (uint16(v6) >= 255) {
                v17 := []int{}
                _ = v17
            } else {
                var v18 uint8 = (v6 / ((v6) | 1))
                fmt.Println(v15, v18, v6)
                _ = v18
            }
        } else {
            v19 := []int{v1, (fi16 / (((v1 << ((42) & 63))) | 1))}
            _ = v19
        }
        _ = fi16
    }
    v20 := GenS1{F0: 0, F1: 4294967295}
    switch (((("hello" + "hello") + "a") + ("a" + "a")) + ("foo bar" + "\n\t")) {
    case "a":
        v21 := v2[42]
        _ = v21
    case "hello":
        for fi22 := 0; fi22 < 19; fi22++ {
            fmt.Println(v6)
            for fi23 := 0; fi23 < 1; fi23++ {
                fmt.Println(fi23, v3)
                v6 = (v6 + v6)
                v15 += (v15 * (v15 - 42))
                _ = fi23
            }
            fmt.Println(v3)
            _ = fi22
        }
    case "foo bar":
        var v24 float64 = 1e10
        _ = v24
    default:
        v25 := []int{}
        _ = v25
    }
    v1 = v1
    if v3 {
        fmt.Println(v1, v15)
    }
    fmt.Println(v15, v1)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v6
}
