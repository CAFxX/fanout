package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }

func main() {
    v1 := []string{(("\n\t" + "a") + "a"), ("a" + ("" + "a")), "a", "a"}
    var v2 int32 = 1
    var v3 bool = false
    v4 := make(map[uint32]uint32)
    var v5 int64 = 0
    v5 += -1
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 957}
        _ = fi6
        _ = v7
    }
    var v8 int32 = v2
    switch 850 {
    case 1:
        switch int(v2) {
        case 32768:
            v8 = (v2 - v8)
        case -83:
            v9 := v4[4294967295]
            _ = v9
        case 128:
            var v10 uint64 = 318
            _ = v10
        }
    case 979:
        var v11 uint64 = 4294967295
        _ = v11
    }
    v12 := GenS2{F0: v3}
    if (false && (!v3)) {
        if v3 {
            v2 = v2
        }
    }
    var v13 uint8 = 1
    v14 := GenS2{F0: (4294967295 > 1)}
    var v15 uint16 = 0
    switch 0 {
    case -620:
        var v16 int = 495
        _ = v16
    default:
        var v17 uint32 = 0
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        switch (fi18 & (fi18 % ((42) | 1))) {
        case 2147483647:
            v19 := GenS1{F0: v15}
            _ = v19
        default:
            if (1 > 960) {
                v8 -= (v2 ^ -2147483648)
            }
        }
        var v20 uint16 = 42
        _ = fi18
        _ = v20
    }
    v21 := GenS2{F0: true}
    v22 := GenS2{F0: ((-2.25 < (((((-2.25 + (0.0 - 3.14159)) - float32(0.0)) * (3.14159 / -2.25)) / (((((((1.5 + (1.5 / 1.5)) / 1.0) / ((((-2.25 + 3.14159) + 1.5) / 1.5) / 1.0)) * (-2.25 * (0.1 - 3.14159))) / 1.5) - (1e10 / (0.1 + float32(1e10)))) * 1.5)) + (1.5 * (1.5 - ((-2.25 / 1e10) - 0.1))))) || v3)}
    v8 = v2
    switch "a" {
    case "":
        for fi23 := 0; fi23 < 18; fi23++ {
            fmt.Println(v8, v8)
            v24 := GenS2{F0: (0.0 > (((0.0 / 1.5) / 1.0) - (0.0 + ((3.14159 - 1.5) / 1e10))))}
            _ = fi23
            _ = v24
        }
    case "hello":
        v25 := GenS3{F0: v3}
        _ = v25
    case "foo bar":
        _ = len(v4)
    default:
        v15 += v15
    }
    fmt.Println(v13, v5, v3)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
