package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 string }
func genFn0(p0 int16, p1 int64, p2 int32) int16 {
    if ((42 >= 42) && (!true)) {
        fmt.Println(p2, p0)
        p2 = (p2 << ((2) & 63))
        if true {
            switch -9223372036854775808 {
            case -1:
                p2 -= (p2 % ((-1) | 1))
            case 32768:
                v6 := GenS1{F0: 0, F1: 786}
                _ = v6
            default:
                p0 += (p0 << ((255) & 63))
            }
            var v7 float64 = (1e10 * -2.25)
            _ = v7
        }
    }
    var v8 uint16 = 825
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p0 % ((int16(v8)) | 1))
}

func main() {
    var v1 uint8 = 2
    var v2 uint = 42
    var v3 uint8 = uint8(v2)
    v4 := make(map[string]uint32)
    var v5 uint64 = 291
    v9 := GenS1{F0: 482, F1: v5}
    v5 = v5
    v10 := []int{42, 118, 4294967295, -559}
    var v11 float64 = (float64(0.0) / float64(1e10))
    if false {
        for fi12 := 0; fi12 < 20; fi12++ {
            v3 = v1
            fmt.Println(v3, v3, v5)
            fmt.Println(v11)
            _ = fi12
        }
        v11 = ((0.1 / 0.1) * v11)
        v2 -= 4294967295
    }
    v10 = append(v10, 1)
    v13 := GenS1{F0: 1, F1: (^(v5))}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    switch "\n\t" {
    case "foo bar":
        v14 := GenS1{F0: 262, F1: (v5 << ((65535) & 63))}
        _ = v14
    case "\n\t":
        _ = len(v4)
    default:
        if false {
            switch (("\n\t" + (("a" + ("" + "\n\t")) + "\n\t")) + "a") {
            case "\n\t":
                v11 = 1.5
            default:
                switch "" {
                case "":
                    delete(v4, ("" + ""))
                case "\n\t":
                    for fi15 := 0; fi15 < 8; fi15++ {
                        var v16 float64 = 1e10
                        if (((1 <= (fi15 - (^(fi15)))) || (false && true)) && ((v5 ^ (v5 ^ v5)) == (v5 * uint64(v2)))) {
                            fmt.Println(v1, v5)
                        }
                        _ = fi15
                        _ = v16
                    }
                }
            }
            v11 += v11
            var v17 uint32 = 4294967295
            _ = v17
        } else {
            var v18 int = 4294967295
            var v19 uint = v2
            _ = v18
            _ = v19
        }
    }
    v20 := GenS2{F0: -738, F1: 1, F2: "hello"}
    switch 855 {
    case -757:
        var v21 uint8 = v3
        _ = v21
    case 2147483647:
        for fi22 := 0; fi22 < 5; fi22++ {
            var v23 uint8 = v1
            var v24 int16 = int16(v5)
            _ = fi22
            _ = v23
            _ = v24
        }
    default:
        var v25 int64 = -361
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
