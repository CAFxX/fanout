package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int }

func main() {
    v1 := []uint32{570}
    var v2 uint16 = 32767
    var v3 uint16 = (^(v2))
    v2 -= 0
    if (-9223372036854775808 != 128) {
        for fi4 := 0; fi4 < 13; fi4++ {
            if true {
                switch (("a" + "foo bar") + "foo bar") {
                case "foo bar":
                    switch -985 {
                    case 2147483647:
                        v5 := GenS1{F0: 642}
                        _ = v5
                    case 1:
                        var v6 int64 = 42
                        _ = v6
                    }
                case "a":
                    fmt.Println(v3)
                case "":
                    v3 = v2
                }
            }
            fmt.Println(fi4)
            for fi7 := 0; fi7 < 10; fi7++ {
                v2 -= (v2 << ((1) & 63))
                var v8 float64 = ((0.1 - -2.25) / (1.5 * 1e10))
                fmt.Println(v8)
                _ = fi7
                _ = v8
            }
            _ = fi4
        }
        v9 := GenS1{F0: 834}
        switch -792 {
        case 32767:
            for fi10 := 0; fi10 < 2; fi10++ {
                v3 += 42
                v2 = (v2 & 42)
                _ = fi10
            }
        case 630:
            if len(v1) > 0 {
                v1[0] = 636
            }
        }
        _ = v9
    }
    v3 = v3
    fmt.Println(v2, v3)
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v11 := GenS1{F0: -651}
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(v3, v2)
        _ = fi12
    }
    v13 := GenS1{F0: 1}
    v2 -= v2
    fmt.Println(v3, v2)
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int64 = 0
        switch "" {
        case "":
            v15 += 1
        }
        _ = fi14
        _ = v15
    }
    v2 += (v2 & v3)
    var v16 bool = false
    switch -9223372036854775808 {
    case -820:
        var v17 float32 = 3.14159
        _ = v17
    case 32768:
        v18 := GenS2{F0: -32768, F1: (v2 % ((2) | 1))}
        _ = v18
    case 194:
        if (951 == 2) {
            v16 = (!(!false))
        }
    default:
        for fi19 := 0; fi19 < 4; fi19++ {
            if (uint32(v3) > 2) {
                v16 = true
                var v20 uint32 = 0
                var v21 uint = 42
                _ = v20
                _ = v21
            }
            v22 := GenS1{F0: 9223372036854775807}
            _ = fi19
            _ = v22
        }
    }
    var v23 int8 = -128
    if (v16 || true) {
        v24 := GenS3{F0: -9223372036854775808}
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v13
    _ = v16
    _ = v2
    _ = v23
    _ = v3
}
