package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint }

func main() {
    var v1 uint16 = 884
    var v2 uint8 = 0
    var v3 float32 = float32(0.1)
    switch 9223372036854775807 {
    case 340:
        var v4 int64 = -352
        _ = v4
    case -599:
        if true {
            v5 := []int{}
            if len(v5) > 0 {
                v5[0] = -340
            }
            if len(v5) > 0 {
                v5[0] = 521
            }
            _ = v5
        } else {
            var v6 int32 = -608
            v7 := []int{}
            v8 := GenS1{F0: (v6 & v6), F1: 42}
            _ = v6
            _ = v7
            _ = v8
        }
    case 129:
        var v9 uint8 = uint8(v1)
        _ = v9
    }
    switch 1 {
    case -2147483648:
        if (!true) {
            v3 -= v3
        } else {
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 uint8 = 2
                _ = fi10
                _ = v11
            }
        }
    case 1:
        if (("hello" >= "foo bar") && (1000 >= -663)) {
            fmt.Println(v3, v1)
            v2 -= v2
        }
    }
    v12 := GenS1{F0: 2147483647, F1: 100}
    for fi13 := 0; fi13 < 5; fi13++ {
        switch "a" {
        case "foo bar":
            var v14 uint32 = 517
            _ = v14
        case "":
            var v15 float32 = float32(((0.0 / -2.25) - (0.0 - -2.25)))
            _ = v15
        }
        _ = fi13
    }
    var v16 float32 = v3
    v17 := GenS2{F0: 0, F1: 1, F2: 18446744073709551615}
    fmt.Println(v3, v3)
    var v18 uint32 = 65535
    fmt.Println(v16, v1, v16)
    v19 := []int{}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
