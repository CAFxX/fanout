package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int64, p1 uint8) uint {
    v4 := GenS1{F0: 1, F1: p1}
    switch 4294967295 {
    case -1:
        var v5 uint16 = 2
        _ = v5
    case 382:
        v6 := GenS2{F0: -32768, F1: 1}
        _ = v6
    case 32768:
        v7 := GenS2{F0: 1, F1: 42}
        _ = v7
    }
    if (-6 >= -128) {
        var v8 int64 = p0
        _ = v8
    } else {
        if (((p1 / (((p1 << ((65535) & 63))) | 1)) > 42) || (!(p1 >= p1))) {
            p1 = (p1 + (p1 * p1))
        }
        p0 += (p0 | p0)
    }
    p1 += 128
    var v9 uint64 = 996
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return 350
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[int]bool)
    var v3 int = v1
    if true {
        if true {
            if false {
                for fi10 := 0; fi10 < 11; fi10++ {
                    switch "\n\t" {
                    case "foo bar":
                        if false {
                            for fi11 := 0; fi11 < 2; fi11++ {
                                v1 = 147
                                _ = fi11
                            }
                            fmt.Println(v3, fi10, fi10)
                        }
                    default:
                        v1 = -391
                    }
                    _ = fi10
                }
            }
        }
        switch (("foo bar" + "hello") + "") {
        case "foo bar":
            var v12 int64 = -1
            _ = v12
        case "hello":
            v13 := GenS3{F0: 42, F1: (v1 + v3)}
            _ = v13
        default:
            v1 += 32767
        }
        var v14 float64 = float64(float32((0.1 - (-2.25 - 1.5))))
        _ = v14
    }
    var v15 int32 = -952
    fmt.Println(v1, v1)
    v1 += 958
    fmt.Println(v3)
    v16 := GenS4{F0: 18446744073709551615}
    var v17 uint32 = 2147483647
    fmt.Println(v17, v15, v3)
    if ((((v15 | int32(v1)) > v15) || (-1 == 1)) && (v3 == (v3 + int(v17)))) {
        v18 := v2[v3]
        if ("" >= ("a" + "a")) {
            var v19 int64 = 0
            v20 := GenS3{F0: 1, F1: (v1 - (v1 / (((v3 * (v3 << ((65535) & 63)))) | 1)))}
            _ = v19
            _ = v20
        }
        _ = v18
    } else {
        v3 = (v3 << ((4294967295) & 63))
    }
    v21 := GenS4{F0: 143}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
}
