package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{990, -9223372036854775808, -9223372036854775808, -806}
    var v3 int16 = -1
    var v4 uint32 = 2147483647
    var v5 uint16 = 32767
    var v6 uint64 = 4294967295
    v6 -= v6
    for fi7 := 0; fi7 < 7; fi7++ {
        v1[v4] = 65535
        _ = fi7
    }
    if len(v2) > 0 {
        v2[0] = 66
    }
    if (((1.5 / 1.5) == 1e10) && false) {
        _ = len(v2)
        v4 = v4
        for fi8 := 0; fi8 < 6; fi8++ {
            v5 -= (v5 >> ((384) & 63))
            switch "a" {
            case "a":
                var v9 uint8 = 127
                _ = v9
            case "hello":
                var v10 bool = false
                _ = v10
            default:
                if (!false) {
                    v11 := GenS1{F0: (("foo bar" + "") + "a"), F1: 474, F2: 69}
                    v3 -= v3
                    for fi12 := 0; fi12 < 5; fi12++ {
                        var v13 string = ""
                        switch "a" {
                        case "\n\t":
                            fmt.Println(fi8)
                        case "hello":
                            v14 := GenS2{F0: ("" + (v13 + "hello"))}
                            _ = v14
                        }
                        if true {
                            v15 := GenS3{F0: 871, F1: 42, F2: 550}
                            _ = v15
                        }
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                }
            }
            _ = fi8
        }
    }
    v16 := GenS2{F0: "foo bar"}
    fmt.Println(v4, v4)
    if false {
        switch "\n\t" {
        case "a":
            v17 := GenS3{F0: 1, F1: int64(v4), F2: v6}
            _ = v17
        case "hello":
            v5 += (v5 >> ((0) & 63))
        default:
            for fi18 := 0; fi18 < 19; fi18++ {
                v19 := GenS2{F0: ""}
                v20 := GenS3{F0: 42, F1: -9223372036854775808, F2: (v6 + 307)}
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    }
    switch 42 {
    case 2147483647:
        fmt.Println(v6, v5, v6)
    case 1:
        var v21 int = -330
        _ = v21
    case -1:
        for fi22 := 0; fi22 < 19; fi22++ {
            _ = len(v1)
            _ = fi22
        }
    }
    switch ("a" + "foo bar") {
    case "":
        var v23 float32 = 0.0
        _ = v23
    case "\n\t":
        v4 = uint32(v6)
    default:
        var v24 uint8 = 0
        _ = v24
    }
    v6 = v6
    fmt.Println(v6, v5, v5)
    for fi25 := 0; fi25 < 10; fi25++ {
        v4 = v4
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
