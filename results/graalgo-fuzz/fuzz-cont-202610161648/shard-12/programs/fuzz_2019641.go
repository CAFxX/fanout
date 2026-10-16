package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int16; F1 int64 }

func main() {
    var v1 float64 = (1.5 - 1e10)
    v2 := []int{127}
    v3 := make(map[string]uint32)
    var v4 int = 313
    fmt.Println(v4, v4)
    v5 := GenS1{F0: 4294967295, F1: (2147483648 >= 380)}
    v6 := GenS2{F0: 541}
    v4 += 32768
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := v3[("" + "\n\t")]
        v8 -= (v8 << ((42) & 63))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS3{F0: ("\n\t" + "a"), F1: -2147483648}
        _ = fi9
        _ = v10
    }
    v4 += (v4 + 2147483647)
    if false {
        for fi11 := 0; fi11 < 5; fi11++ {
            for fi12 := 0; fi12 < 16; fi12++ {
                switch fi11 {
                case 9223372036854775807:
                    v2 = append(v2, v4)
                case -1:
                    if (246 != 1) {
                        v4 = v4
                        switch "a" {
                        case "foo bar":
                            fmt.Println(fi11, fi11, v1)
                        case "\n\t":
                            v1 += v1
                        default:
                            if (4294967295 > 0) {
                                fmt.Println(v1)
                                var v13 int = 42
                                _ = v13
                            } else {
                                v1 += v1
                                var v14 int32 = -542
                                _ = v14
                            }
                            v4 -= (v4 & (fi12 % (((fi11 | (^(fi12)))) | 1)))
                            v15 := GenS4{F0: 1000, F1: 1}
                            _ = v15
                        }
                    }
                }
                _ = fi12
            }
            if len(v2) > 0 {
                v2[0] = fi11
            }
            _ = fi11
        }
    }
    switch 9223372036854775807 {
    case 4294967295:
        v16 := v3["a"]
        _ = v16
    case 9223372036854775807:
        v17 := GenS4{F0: 32767, F1: 9223372036854775807}
        _ = v17
    default:
        for fi18 := 0; fi18 < 6; fi18++ {
            switch (fi18 - v4) {
            case -1:
                fmt.Println(v4, fi18)
            default:
                if len(v2) > 0 {
                    v2[0] = (-(fi18))
                }
            }
            _ = fi18
        }
    }
    v19 := v3[("" + "a")]
    v2 = append(v2, (v4 | (v4 % ((v4) | 1))))
    var v20 uint16 = 32767
    v4 += v4
    var v21 int = 128
    var v22 int = v21
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
