package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 bool }
func genFn0() int16 {
    var v5 string = "\n\t"
    fmt.Println(v5)
    fmt.Println(v5)
    _ = v5
    return -32768
}

func main() {
    var v1 int64 = -1
    var v2 uint32 = 707
    var v3 uint64 = 42
    v4 := []uint32{42, 834, 2147483647}
    v2 = v2
    var v6 string = ("hello" + "\n\t")
    fmt.Println(v1, v6)
    v3 -= v3
    v7 := GenS1{F0: 324}
    if (0 != -2147483648) {
        switch 9223372036854775807 {
        case 0:
            var v8 int = 682
            _ = v8
        }
        var v9 bool = false
        v2 = uint32(v3)
        _ = v9
    }
    switch "" {
    case "foo bar":
        v4 = append(v4, v2)
    case "\n\t":
        if (42 != 677) {
            v6 = "\n\t"
            _ = len(v4)
        }
    }
    var v10 bool = true
    v11 := GenS1{F0: 65535}
    for fi12 := 0; fi12 < 5; fi12++ {
        if len(v4) > 0 {
            v4[0] = (v2 - (v2 + (v2 - 259)))
        }
        var v13 string = (v6 + v6)
        v6 = (v13 + v6)
        _ = fi12
        _ = v13
    }
    switch ("" + v6) {
    case "foo bar":
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 uint8 = 128
            if ((^(v15)) == (v15 << ((1) & 63))) {
                var v16 float64 = 0.1
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 float32 = (0.1 / 1.0)
                    _ = fi17
                    _ = v18
                }
                var v19 uint8 = 212
                _ = v16
                _ = v19
            }
            _ = len(v4)
            _ = fi14
            _ = v15
        }
    case "":
        v20 := GenS2{F0: (^(v3))}
        _ = v20
    case "a":
        var v21 uint = 476
        _ = v21
    default:
        v22 := GenS3{F0: 32767, F1: false}
        _ = v22
    }
    switch (v6 + "a") {
    case "":
        _ = len(v4)
    case "a":
        v3 -= (v3 ^ 4294967295)
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v23 int16 = -1
    v10 = v10
    v24 := GenS3{F0: (v23 - (v23 << ((4294967295) & 63))), F1: v10}
    v10 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
