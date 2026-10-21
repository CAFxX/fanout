package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{}
    v3 := []uint32{1, 2147483647, 2}
    var v4 int32 = -1
    var v5 int32 = 452
    if false {
        var v6 string = "foo bar"
        _ = v6
    } else {
        v5 += (v4 * v4)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v7 := GenS1{F0: (42 <= 1)}
    v5 += (-(v5))
    delete(v1, -609)
    v5 -= (v5 / ((v4) | 1))
    switch -438 {
    case -2147483648:
        for fi8 := 0; fi8 < 16; fi8++ {
            v5 = v5
            if false {
                if len(v2) > 0 {
                    v2[0] = 813
                }
            }
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 int16 = -1
                v4 += v5
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    case 63:
        if false {
            v5 = v4
        }
    case 128:
        var v11 int8 = -72
        _ = v11
    default:
        if (802 < 2) {
            if len(v3) > 0 {
                v3[0] = 2
            }
            switch (("foo bar" + "hello") + "hello") {
            case "a":
                for fi12 := 0; fi12 < 17; fi12++ {
                    for fi13 := 0; fi13 < 15; fi13++ {
                        v14 := GenS2{F0: 42, F1: 2}
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
            case "":
                v4 = (v4 | v5)
            case "foo bar":
                for fi15 := 0; fi15 < 7; fi15++ {
                    var v16 uint16 = 659
                    _ = fi15
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 19; fi17++ {
                    if true {
                        var v18 uint64 = 379
                        _ = v18
                    }
                    _ = fi17
                }
            }
            var v19 int8 = 1
            _ = v19
        }
    }
    var v20 uint8 = 1
    v5 += 2147483647
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
