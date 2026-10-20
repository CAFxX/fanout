package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 int = 0
    v2 := []uint32{2147483648, 65535}
    var v3 int64 = 42
    var v4 int16 = 42
    v5 := make(map[uint32]uint32)
    v6 := []string{}
    fmt.Println(v3, v1)
    v1 += v1
    if (!(("" != "") && true)) {
        var v7 int = 32768
        _ = v7
    } else {
        for fi8 := 0; fi8 < 20; fi8++ {
            v1 -= 32767
            delete(v5, 0)
            _ = fi8
        }
    }
    switch "" {
    case "":
        if true {
            fmt.Println(v3, v1, v4)
            v9 := GenS1{F0: 230, F1: 2}
            for fi10 := 0; fi10 < 4; fi10++ {
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 uint64 = 81
                    v1 -= fi11
                    _ = fi11
                    _ = v12
                }
                var v13 uint64 = 18446744073709551615
                _ = fi10
                _ = v13
            }
            _ = v9
        }
    case "\n\t":
        var v14 float32 = 0.0
        _ = v14
    case "a":
        v4 -= (^(v4))
    }
    switch (v1 ^ v1) {
    case -347:
        var v15 int16 = (v4 ^ v4)
        _ = v15
    case -972:
        var v16 bool = false
        _ = v16
    case 32768:
        fmt.Println(v3)
    default:
        if (!(v4 > int16(v1))) {
            _ = len(v5)
            fmt.Println(v4, v1, v1)
        }
    }
    switch v1 {
    case -291:
        var v17 int = (v1 >> ((18446744073709551615) & 63))
        _ = v17
    default:
        v4 = int16(v3)
    }
    switch 9223372036854775807 {
    case -9223372036854775808:
        switch "" {
        case "a":
            var v18 int32 = -117
            _ = v18
        case "hello":
            var v19 int64 = 663
            _ = v19
        }
    default:
        delete(v5, 2147483648)
    }
    if (3.14159 != 0.1) {
        v4 = 1
        if len(v6) > 0 {
            v6[0] = "foo bar"
        }
        var v20 uint = 564
        _ = v20
    }
    var v21 int64 = v3
    v1 += v1
    v22 := v5[743]
    var v23 int32 = 0
    var v24 uint8 = 0
    switch 197 {
    case 0:
        v25 := GenS1{F0: 65535, F1: v24}
        _ = v25
    case -635:
        var v26 int = (-(v1))
        _ = v26
    }
    fmt.Println(v21, v3)
    var v27 uint16 = 32767
    v5[(v22 << ((1) & 63))] = (v22 << ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
