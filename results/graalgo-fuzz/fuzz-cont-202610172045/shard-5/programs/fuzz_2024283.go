package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int; F2 bool }
type GenS4 struct { F0 uint; F1 int8 }

func main() {
    var v1 float64 = ((3.14159 - (1e10 * -2.25)) / ((0.0 + (0.1 * 0.0)) - 0.1))
    var v2 uint8 = 59
    var v3 int16 = 1000
    v4 := make(map[string]int)
    v5 := []int{380, 127, 42, int(v2)}
    v6 := make(map[string]bool)
    v7 := GenS1{F0: 18446744073709551615, F1: 2147483647}
    switch "\n\t" {
    case "hello":
        var v8 uint8 = v2
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 17; fi9++ {
            delete(v6, "hello")
            _ = fi9
        }
    }
    v10 := GenS1{F0: uint64(v2), F1: 65535}
    delete(v4, "\n\t")
    var v11 bool = (!(535 > 42))
    var v12 uint32 = 2147483648
    if ((v2 << ((4294967295) & 63)) > v2) {
        for fi13 := 0; fi13 < 12; fi13++ {
            v11 = (2 == 186)
            var v14 int64 = 1
            if false {
                var v15 int8 = 116
                _ = v15
            } else {
                var v16 int32 = -1
                v6[("\n\t" + "a")] = (((v1 / ((float64((1.5 * 1.5)) * 1e10) * 1e10)) / v1) >= 3.14159)
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        v17 := GenS2{F0: (v12 & v12)}
        delete(v4, "a")
        _ = v17
    } else {
        for fi18 := 0; fi18 < 8; fi18++ {
            if ((v2 >> ((4294967295) & 63)) >= 255) {
                var v19 uint8 = v2
                v20 := GenS3{F0: 2, F1: (fi18 * (fi18 / ((fi18) | 1))), F2: false}
                _ = v19
                _ = v20
            } else {
                fmt.Println(v3, v1, v1)
            }
            v21 := GenS4{F0: 0, F1: 100}
            _ = fi18
            _ = v21
        }
        if (1 == -490) {
            _ = len(v5)
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS2{F0: v12}
            v24 := GenS1{F0: 18446744073709551615, F1: v12}
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    v25 := v4[""]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
