package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint16; F2 int }
type GenS7 struct { F0 uint; F1 uint32; F2 uint32 }

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]int)
    var v3 uint16 = 0
    var v4 uint32 = 725
    if false {
        v4 += v4
        for fi5 := 0; fi5 < 5; fi5++ {
            v4 = (v4 - v4)
            if (1 >= 0) {
                var v6 uint = 175
                delete(v2, (-(fi5)))
                _ = v6
            }
            v7 := GenS1{F0: -1, F1: 127, F2: 1}
            _ = fi5
            _ = v7
        }
    } else {
        v8 := GenS2{F0: v4}
        v9 := GenS2{F0: v4}
        v10 := []int{}
        _ = v10
        _ = v8
        _ = v9
    }
    if false {
        for fi11 := 0; fi11 < 2; fi11++ {
            switch -572 {
            case 0:
                fmt.Println(v3)
            case 9223372036854775807:
                if true {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v13 := []int{}
                        switch "" {
                        case "foo bar":
                            var v14 int8 = -1
                            _ = v14
                        }
                        var v15 float32 = (0.1 - ((1.5 + (0.1 / 0.1)) - 0.0))
                        _ = fi12
                        _ = v13
                        _ = v15
                    }
                    var v16 uint64 = uint64(v4)
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 8; fi17++ {
                    v18 := GenS1{F0: 32767, F1: -26, F2: 2}
                    v3 = (^(v3))
                    v3 += (v3 % ((590) | 1))
                    v19 := GenS3{F0: (v4 % ((v4) | 1)), F1: 18446744073709551615}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            _ = fi11
        }
        v20 := []int{-1, -544}
        _ = v20
    }
    fmt.Println(v4, v3, v1)
    v21 := []int{2147483647, 32768}
    _ = len(v21)
    v3 += v3
    v1 -= -2.25
    var v22 int = int(v3)
    v1 = (0.0 / (float64(0.1) / (0.1 * v1)))
    if true {
        v23 := GenS4{F0: -128, F1: -943}
        if true {
            var v24 uint64 = uint64(v4)
            var v25 int64 = 0
            v26 := GenS5{F0: (^(v3)), F1: (v4 % ((v4) | 1))}
            _ = v24
            _ = v25
            _ = v26
        }
        var v27 int8 = -48
        _ = v23
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        if true {
            var v29 uint64 = 2
            _ = v29
        } else {
            v30 := GenS6{F0: false, F1: (v3 & (v3 % (((v3 ^ 323)) | 1))), F2: fi28}
            _ = v30
        }
        var v31 int8 = 0
        _ = fi28
        _ = v31
    }
    switch v22 {
    case 0:
        v32 := GenS7{F0: 461, F1: (v4 + (v4 * (v4 + (v4 / ((278) | 1))))), F2: v4}
        _ = v32
    }
    if len(v21) > 0 {
        v21[0] = v22
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        v3 -= (v3 % ((254) | 1))
        _ = fi33
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        v2[(v22 >> ((308) & 63))] = fi34
        _ = fi34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        _ = len(v2)
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
