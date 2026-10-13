package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int8 }
type GenS2 struct { F0 uint }
func genFn0(p0 int16) uint64 {
    v4 := GenS1{F0: -1, F1: ("\n\t" + "a"), F2: 42}
    fmt.Println(p0)
    fmt.Println(p0)
    v5 := GenS1{F0: -128, F1: "foo bar", F2: 42}
    var v6 int = 42
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 570
}
func genFn1() uint16 {
    var v7 uint64 = 18446744073709551615
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := []int{fi8, (fi8 - fi8)}
        v7 = v7
        switch ("a" + ("foo bar" + ("a" + "\n\t"))) {
        case "hello":
            fmt.Println(v7, fi8)
        case "\n\t":
            fmt.Println(fi8, v7)
        }
        _ = fi8
        _ = v9
    }
    var v10 uint16 = 2
    v11 := []int{1, 127, 983}
    fmt.Println(v7, v7)
    _ = v10
    _ = v11
    _ = v7
    return v10
}

func main() {
    var v1 uint8 = 97
    var v2 int = 0
    v3 := []uint32{4294967295, 0, 525, 0}
    if false {
        if (true || (54 == 65535)) {
            for fi12 := 0; fi12 < 1; fi12++ {
                v2 = (v2 - (v2 ^ fi12))
                v2 += (^(fi12))
                _ = len(v3)
                _ = fi12
            }
            v1 = v1
        }
        switch (v2 << ((65535) & 63)) {
        case 0:
            var v13 float64 = 0.1
            _ = v13
        }
    } else {
        var v14 uint8 = 0
        _ = v14
    }
    v3 = append(v3, 2)
    switch ("\n\t" + ("\n\t" + "a")) {
    case "":
        fmt.Println(v2, v2)
    default:
        v2 = 32768
    }
    v15 := GenS2{F0: 140}
    var v16 int16 = 0
    v17 := GenS2{F0: 0}
    if (732 <= 0) {
        v2 = (v2 / ((9223372036854775807) | 1))
        v18 := GenS2{F0: 255}
        _ = v18
    } else {
        fmt.Println(v1, v16, v2)
        v16 += 32767
        v16 += (v16 ^ v16)
    }
    _ = len(v3)
    var v19 float32 = float32(3.14159)
    v3 = append(v3, 3)
    v20 := GenS1{F0: -115, F1: ("a" + (("\n\t" + ((((("\n\t" + ("a" + "a")) + "a") + "foo bar") + "\n\t") + "foo bar")) + "hello")), F2: -35}
    v21 := GenS1{F0: 0, F1: "", F2: -127}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
