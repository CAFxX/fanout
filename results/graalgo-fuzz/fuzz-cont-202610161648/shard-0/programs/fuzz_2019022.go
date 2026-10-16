package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int64 }

func main() {
    var v1 uint = 42
    var v2 int8 = 127
    var v3 int16 = 0
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    var v6 uint = 4294967295
    if (false && true) {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 uint64 = 0
            switch "" {
            case "hello":
                v9 := []int{(fi7 | (fi7 << (((^(v1))) & 63))), (fi7 / ((fi7) | 1))}
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        var v10 uint16 = 42
        _ = v10
    }
    switch 9223372036854775807 {
    case 153:
        var v11 int64 = 1
        _ = v11
    case 42:
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 float64 = 3.14159
            v6 += (v1 / ((v6) | 1))
            _ = fi12
            _ = v13
        }
    case 135:
        if true {
            var v14 int8 = v2
            for fi15 := 0; fi15 < 17; fi15++ {
                v1 = (v1 + 614)
                _ = fi15
            }
            _ = v14
        }
    default:
        v16 := []int{4294967295, -982, -740}
        _ = v16
    }
    v17 := GenS1{F0: (v6 - (^(v6)))}
    if true {
        var v18 uint8 = 42
        _ = v18
    } else {
        fmt.Println(v6)
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        if (true && true) {
            switch (fi19 | -9223372036854775808) {
            case 42:
                if true {
                    var v20 uint16 = 296
                    v21 := GenS2{F0: 42}
                    _ = v20
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v23 := GenS2{F0: 217}
                        _ = fi22
                        _ = v23
                    }
                    var v24 int64 = -35
                    v25 := []int{0, 4294967295, (fi19 + (fi19 << ((v1) & 63))), fi19}
                    _ = v24
                    _ = v25
                }
            default:
                var v26 uint16 = 528
                _ = v26
            }
            v27 := GenS2{F0: 219}
            _ = len(v5)
            _ = v27
        }
        for fi28 := 0; fi28 < 16; fi28++ {
            switch (fi19 & fi28) {
            case 42:
                if true {
                    v6 += v1
                    v1 = (v6 << ((v1) & 63))
                } else {
                    v29 := GenS3{F0: true, F1: 0}
                    _ = v29
                }
            case -9223372036854775808:
                v5[42] = 408
            }
            _ = fi28
        }
        v30 := []int{}
        _ = fi19
        _ = v30
    }
    var v31 uint = (v1 | (v1 >> ((18446744073709551615) & 63)))
    var v32 int = 127
    v1 += (v1 << (((v1 & (v31 << ((255) & 63)))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
