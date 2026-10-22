package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int64 }
type GenS2 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 bool = (0 > 342)
    var v3 int64 = 9223372036854775807
    for fi4 := 0; fi4 < 2; fi4++ {
        v3 = v3
        if v2 {
            var v5 float64 = float64(3.14159)
            v6 := GenS1{F0: 18446744073709551615, F1: -128, F2: v3}
            var v7 int32 = 0
            _ = v5
            _ = v6
            _ = v7
        }
        switch (fi4 + fi4) {
        case 134:
            var v8 int8 = 19
            _ = v8
        case 128:
            fmt.Println(fi4)
        default:
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 int8 = -11
                v11 := GenS2{F0: 1, F1: v3}
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        _ = fi4
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := GenS1{F0: 792, F1: 33, F2: (v3 >> ((42) & 63))}
            v15 := []int{(fi13 - (fi12 + fi13)), (fi13 - -2147483648), fi12}
            var v16 uint = 97
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi12
    }
    var v17 uint16 = 458
    var v18 float64 = float64(1.5)
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS2{F0: 4294967295, F1: v3}
        _ = fi19
        _ = v20
    }
    fmt.Println(v3, v18)
    fmt.Println(v2)
    var v21 int16 = 1
    switch 32767 {
    case -228:
        v22 := GenS1{F0: 4294967295, F1: -39, F2: -1}
        _ = v22
    case -2147483648:
        v23 := []int{42, 127, 32768, 128}
        _ = v23
    case 978:
        v24 := GenS2{F0: 42, F1: v3}
        _ = v24
    }
    if v2 {
        v18 -= v18
    } else {
        switch ("a" + ("\n\t" + "")) {
        case "foo bar":
            var v25 int16 = v21
            _ = v25
        }
    }
    switch "\n\t" {
    case "a":
        v17 += 48
    case "hello":
        for fi26 := 0; fi26 < 9; fi26++ {
            if v2 {
                v21 = v21
            }
            var v27 int32 = 1
            v28 := GenS1{F0: 454, F1: 100, F2: v3}
            _ = fi26
            _ = v27
            _ = v28
        }
    default:
        fmt.Println(v17, v17, v21)
    }
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
}
