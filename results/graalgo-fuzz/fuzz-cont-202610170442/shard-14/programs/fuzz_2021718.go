package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint; F2 uint64 }

func main() {
    var v1 bool = true
    var v2 int = 128
    var v3 uint16 = 65535
    var v4 int64 = -1
    var v5 uint32 = 1
    var v6 uint16 = (v3 << ((2) & 63))
    v7 := []int{}
    fmt.Println(v1, v4)
    v8 := GenS1{F0: (v6 | v3)}
    for fi9 := 0; fi9 < 16; fi9++ {
        if v1 {
            var v10 bool = v1
            fmt.Println(v5, v4, v6)
            switch ("\n\t" + "hello") {
            case "hello":
                switch (fi9 / ((fi9) | 1)) {
                case 193:
                    v2 += (v2 << ((65535) & 63))
                default:
                    v11 := GenS2{F0: 42}
                    _ = v11
                }
            }
            _ = v10
        } else {
            v12 := GenS1{F0: (v6 * v6)}
            v13 := GenS2{F0: -1}
            _ = v12
            _ = v13
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi9
    }
    v4 += v4
    switch (v2 - (v2 - (v2 << ((255) & 63)))) {
    case 4294967295:
        switch "" {
        case "hello":
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 float32 = ((1.5 / 1.5) - ((1e10 - ((1e10 - 0.0) * 0.0)) + 1e10))
                _ = fi14
                _ = v15
            }
        default:
            fmt.Println(v4, v2, v2)
        }
    case 108:
        v2 = -474
    default:
        var v16 int8 = 0
        _ = v16
    }
    v17 := GenS3{F0: (^(v2)), F1: 1, F2: 4294967295}
    v18 := GenS2{F0: 42}
    v7 = append(v7, v2)
    switch 42 {
    case 128:
        switch (v2 & 9223372036854775807) {
        case 42:
            v4 -= (v4 / ((v4) | 1))
        }
    case 0:
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 uint = 0
            var v21 uint32 = (v5 | (v5 ^ v5))
            for fi22 := 0; fi22 < 6; fi22++ {
                v4 = -1
                _ = fi22
            }
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v7 = append(v7, 32768)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        fmt.Println(v3, v4)
        _ = fi23
    }
    v24 := GenS1{F0: (v3 >> ((0) & 63))}
    var v25 int16 = -30
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
