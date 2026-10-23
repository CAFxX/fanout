package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int32 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint) int8 {
    switch -2147483648 {
    case 338:
        p0 += p0
    }
    var v7 int = -93
    _ = p0
    _ = v7
    return -1
}

func main() {
    var v1 bool = (42 > 2)
    var v2 int = 9223372036854775807
    var v3 uint32 = 331
    var v4 int = (v2 - (^(v2)))
    var v5 int32 = 42
    v6 := make(map[uint32]uint32)
    fmt.Println(v5, v3)
    v1 = v1
    v8 := GenS1{F0: 1, F1: -1, F2: 290}
    v9 := GenS1{F0: -6, F1: -1, F2: -272}
    var v10 float32 = -2.25
    v3 += (^(v3))
    v6[(v3 ^ (v3 % ((v3) | 1)))] = (v3 ^ (v3 << ((42) & 63)))
    if (((2 > 89) || true) || (v3 >= (^(v3)))) {
        var v11 uint8 = 188
        for fi12 := 0; fi12 < 4; fi12++ {
            v10 = ((1e10 + 1e10) / (v10 + (v10 - (v10 - v10))))
            v11 += (v11 / ((127) | 1))
            _ = fi12
        }
        _ = v11
    } else {
        v3 = v3
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 uint16 = 722
            for fi15 := 0; fi15 < 13; fi15++ {
                fmt.Println(v1, fi15)
                v16 := GenS2{F0: 42, F1: v1, F2: 179}
                var v17 bool = (42 < 2)
                var v18 uint64 = 1
                var v19 uint = 151
                if v1 {
                    var v20 uint = (v19 ^ 42)
                    switch "foo bar" {
                    case "foo bar":
                        var v21 int32 = v5
                        _ = v21
                    case "\n\t":
                        var v22 int32 = (v5 / ((-485) | 1))
                        _ = v22
                    default:
                        v23 := GenS1{F0: 42, F1: 0, F2: v5}
                        _ = v23
                    }
                    _ = v20
                }
                fmt.Println(v10, v17, v3)
                v3 = (v3 / ((v3) | 1))
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi13
            _ = v14
        }
    }
    v5 -= (v5 << ((42) & 63))
    var v24 uint8 = 2
    v4 -= v4
    for fi25 := 0; fi25 < 2; fi25++ {
        if (!false) {
            fmt.Println(fi25, v1, v4)
            v26 := []int{}
            v27 := v6[v3]
            _ = v26
            _ = v27
        }
        _ = fi25
    }
    var v28 int64 = 102
    for fi29 := 0; fi29 < 1; fi29++ {
        v30 := GenS3{F0: (v5 | (v5 + -760))}
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
