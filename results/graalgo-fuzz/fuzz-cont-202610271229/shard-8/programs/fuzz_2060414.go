package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 bool }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint64, p1 bool, p2 uint64) uint16 {
    var v6 float64 = (0.0 / 1.5)
    v7 := GenS1{F0: 42, F1: p0, F2: false}
    fmt.Println(p2)
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(p0, fi8, fi8)
        v9 := GenS2{F0: ("foo bar" + "a"), F1: 823}
        var v10 float32 = -2.25
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2
    var v3 float32 = 0.0
    var v4 int8 = 100
    var v5 int32 = -737
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS1{F0: -757, F1: uint64(v2), F2: ((1e10 == ((0.1 + 1.5) * (0.0 - float64((v3 / -2.25))))) || (128 > uint8(v2)))}
        v13 := []int{(fi11 - (fi11 | -1)), fi11}
        fmt.Println(v3)
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: "foo bar", F1: 1}
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(v5, v5)
        fmt.Println(v3, v5)
        switch 0 {
        case -1:
            if (1.5 >= ((0.0 * v3) / 1.0)) {
                var v16 int32 = 1
                if false {
                    var v17 bool = (!(1 > 2))
                    _ = v17
                }
                fmt.Println(v2, v2)
                _ = v16
            }
        case 32768:
            for fi18 := 0; fi18 < 18; fi18++ {
                if (!(-582 != 42)) {
                    var v19 uint16 = 135
                    _ = v19
                }
                _ = fi18
            }
        case 9223372036854775807:
            v20 := GenS3{F0: (fi15 << ((800) & 63))}
            _ = v20
        default:
            v5 = v5
        }
        _ = fi15
    }
    if true {
        v21 := []int{334, 4294967295, 9223372036854775807}
        _ = v21
    }
    switch ((("\n\t" + "a") + ((((("" + ("foo bar" + "a")) + "") + "hello") + "\n\t") + "a")) + "hello") {
    case "hello":
        v22 := []int{}
        _ = v22
    case "\n\t":
        fmt.Println(v2, v2, v2)
    default:
        _ = len(v1)
    }
    var v23 int8 = 74
    fmt.Println(v2, v2, v3)
    for fi24 := 0; fi24 < 7; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            v4 = (-(v23))
            fmt.Println(v4, v5)
            var v26 int16 = 388
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
