package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint8; F1 int16; F2 int16 }
func genFn0(p0 int16, p1 bool, p2 uint64) int32 {
    p1 = true
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 float32 = (1.5 - (((((3.14159 / -2.25) - (1e10 * float32((3.14159 / 0.1)))) / -2.25) + (3.14159 / 1.0)) - 3.14159))
            _ = fi6
            _ = v7
        }
        var v8 int64 = 9223372036854775807
        _ = fi5
        _ = v8
    }
    var v9 int16 = (p0 >> ((255) & 63))
    switch 32767 {
    case 32767:
        for fi10 := 0; fi10 < 15; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                fmt.Println(p2, p2, p1)
                v12 := GenS1{F0: 42, F1: p2}
                _ = fi11
                _ = v12
            }
            var v13 uint64 = 0
            _ = fi10
            _ = v13
        }
    case 82:
        for fi14 := 0; fi14 < 14; fi14++ {
            fmt.Println(v9)
            _ = fi14
        }
    case 0:
        for fi15 := 0; fi15 < 2; fi15++ {
            if true {
                v16 := []int{fi15, (^(fi15))}
                var v17 uint64 = p2
                _ = v16
                _ = v17
            }
            if true {
                v18 := []int{(fi15 << ((1) & 63)), fi15}
                v19 := GenS2{F0: uint8(p2), F1: ("" + ("foo bar" + "hello"))}
                p1 = (92 <= 1)
                _ = v18
                _ = v19
            }
            _ = fi15
        }
    default:
        if (65535 == 4294967295) {
            p2 = (p2 & p2)
            if true {
                if (!p1) {
                    var v20 uint = 1
                    v21 := []int{32767, 127}
                    _ = v20
                    _ = v21
                }
            } else {
                fmt.Println(p0, p0)
                v22 := []int{9223372036854775807, 4294967295}
                v23 := GenS1{F0: 9223372036854775807, F1: (^(p2))}
                _ = v22
                _ = v23
            }
        }
    }
    var v24 bool = (0 != 112)
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{}
    v2 := []int{-398, -430}
    v3 := make(map[uint32]int)
    var v4 int32 = 0
    v4 += (^(v4))
    v25 := GenS3{F0: 67, F1: 0, F2: 42}
    v4 -= (v4 & (v4 & v4))
    var v26 int16 = 1
    for fi27 := 0; fi27 < 18; fi27++ {
        switch (-(fi27)) {
        case 32767:
            v28 := GenS3{F0: 255, F1: 1, F2: (^(v26))}
            _ = v28
        }
        v29 := GenS3{F0: 65535, F1: 255, F2: (v26 / ((0) | 1))}
        var v30 int32 = v4
        _ = fi27
        _ = v29
        _ = v30
    }
    if (!(((true || ((-133 <= 2147483647) || false)) && (false && true)) && (-1 != 1))) {
        for fi31 := 0; fi31 < 8; fi31++ {
            if false {
                if (62 <= 127) {
                    delete(v3, 0)
                }
                v26 -= (-(v26))
            } else {
                var v32 int32 = (v4 % ((v4) | 1))
                switch fi31 {
                case 128:
                    v4 += -960
                case 2147483647:
                    v2 = append(v2, (fi31 * fi31))
                }
                v33 := GenS4{F0: 0, F1: (v26 << ((0) & 63)), F2: 42}
                _ = v32
                _ = v33
            }
            v26 = -1
            _ = fi31
        }
    } else {
        v34 := GenS3{F0: 65535, F1: 42, F2: (v26 / ((v26) | 1))}
        v26 = (v26 >> ((18446744073709551615) & 63))
        _ = v34
    }
    fmt.Println(v26, v26)
    var v35 float64 = -2.25
    var v36 int16 = v26
    v36 = v26
    var v37 float64 = ((v35 - v35) * v35)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
