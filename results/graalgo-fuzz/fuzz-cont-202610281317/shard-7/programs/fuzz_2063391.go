package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    var v4 uint64 = 18446744073709551615
    var v5 int = -2147483648
    var v6 uint16 = 324
    v6 -= (v6 - (v6 ^ v6))
    v6 -= uint16(v4)
    v4 -= (v4 << ((4294967295) & 63))
    switch (-(v5)) {
    case -198:
        switch "hello" {
        case "a":
            v7 := GenS1{F0: -172}
            _ = v7
        case "hello":
            v8 := GenS1{F0: -543}
            _ = v8
        }
    default:
        v9 := GenS1{F0: 0}
        _ = v9
    }
    var v10 float32 = ((0.1 + 1.5) * 0.1)
    v11 := GenS1{F0: 0}
    var v12 int = (v5 >> ((1) & 63))
    switch "\n\t" {
    case "":
        if (false || true) {
            v10 = 0.0
        }
    }
    if false {
        v13 := GenS2{F0: ((v6 << ((0) & 63)) > v6), F1: 462}
        switch v12 {
        case 32768:
            v14 := GenS2{F0: false, F1: 1}
            _ = v14
        default:
            if ((((v5 - v5) > 967) || (-2147483648 < (v12 | int(v4)))) && (!true)) {
                var v15 int = (v5 << ((44) & 63))
                v16 := GenS3{F0: (v6 + (v6 / ((v6) | 1)))}
                _ = v15
                _ = v16
            }
        }
        _ = v13
    }
    switch "\n\t" {
    case "\n\t":
        v17 := v3[0]
        _ = v17
    default:
        for fi18 := 0; fi18 < 15; fi18++ {
            for fi19 := 0; fi19 < 17; fi19++ {
                v4 += v4
                _ = fi19
            }
            _ = fi18
        }
    }
    v20 := GenS3{F0: v6}
    var v21 string = ""
    fmt.Println(v4, v6)
    fmt.Println(v12)
    if (!(((42 > 255) || false) && (-1 <= 1000))) {
        for fi22 := 0; fi22 < 5; fi22++ {
            v2[v5] = uint32(v4)
            _ = fi22
        }
        switch v21 {
        case "foo bar":
            v5 -= 127
        case "hello":
            switch v5 {
            case -690:
                var v23 int = v12
                _ = v23
            default:
                fmt.Println(v5, v12)
            }
        default:
            for fi24 := 0; fi24 < 18; fi24++ {
                if (2147483647 <= 0) {
                    delete(v1, v5)
                }
                var v25 bool = ((v4 ^ 2) > v4)
                _ = fi24
                _ = v25
            }
        }
    }
    var v26 int64 = -167
    for fi27 := 0; fi27 < 4; fi27++ {
        v28 := GenS1{F0: 32767}
        _ = fi27
        _ = v28
    }
    var v29 int32 = 785
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
