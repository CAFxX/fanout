package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int64 = -729
    var v2 int32 = 0
    var v3 int = 4294967295
    var v4 uint = 0
    v5 := make(map[uint32]uint32)
    var v6 int32 = v2
    switch (v3 ^ -529) {
    case -141:
        switch v3 {
        case -597:
            v7 := GenS1{F0: 2}
            _ = v7
        }
    }
    var v8 uint = (v4 + (v4 / ((v4) | 1)))
    v2 -= v2
    var v9 int16 = 42
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 uint32 = 65535
        v12 := GenS2{F0: (-(v9))}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch v3 {
    case 9223372036854775807:
        v13 := GenS1{F0: 188}
        _ = v13
    case 1:
        v8 -= 576
    default:
        var v14 uint8 = 146
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS2{F0: v9}
        _ = fi15
        _ = v16
    }
    if ((true || true) || (-71 == -1)) {
        delete(v5, 2147483647)
    } else {
        var v17 uint64 = 1
        var v18 int16 = -1
        switch v3 {
        case 4294967295:
            if (false || true) {
                switch "" {
                case "hello":
                    var v19 int = v3
                    _ = v19
                case "a":
                    var v20 uint = (v8 - (v8 / ((v8) | 1)))
                    _ = v20
                }
                var v21 int16 = v9
                _ = v21
            } else {
                var v22 uint8 = 2
                fmt.Println(v4)
                v22 += (^(v22))
                _ = v22
            }
        case 2147483647:
            if (!false) {
                v4 = (^(v4))
            }
        case 1:
            v2 = (v6 - -69)
        }
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
