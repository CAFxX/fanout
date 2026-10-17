package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int64; F2 int8 }
type GenS3 struct { F0 uint64; F1 int32; F2 int64 }
type GenS4 struct { F0 uint8; F1 int8; F2 int }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS7 struct { F0 uint }
func genFn0(p0 uint64) uint8 {
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 0}
        _ = fi6
        _ = v7
    }
    var v8 uint16 = 39
    v8 += (v8 % ((v8) | 1))
    v9 := []int{591}
    var v10 int32 = -1
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 8
}
func genFn1(p0 uint64, p1 int8, p2 string) int16 {
    switch -843 {
    case 2147483647:
        switch ("hello" + p2) {
        case "":
            if true {
                p0 = p0
                switch -1 {
                case 1:
                    var v11 uint = 1
                    _ = v11
                }
            }
        case "a":
            var v12 float64 = float64(((3.14159 + 0.1) * 1e10))
            _ = v12
        }
    case 32768:
        v13 := GenS1{F0: 32767}
        _ = v13
    case -832:
        v14 := GenS2{F0: -1, F1: -9223372036854775808, F2: -1}
        _ = v14
    }
    v15 := GenS3{F0: (p0 + (p0 / ((p0) | 1))), F1: 2147483647, F2: -979}
    switch 127 {
    case 42:
        var v16 int8 = (^(p1))
        _ = v16
    case 496:
        var v17 int32 = 0
        _ = v17
    case 32767:
        v18 := GenS4{F0: 86, F1: p1, F2: -9223372036854775808}
        _ = v18
    default:
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 int8 = p1
            _ = fi19
            _ = v20
        }
    }
    var v21 int8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v21
    return -32768
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 4294967295
    var v3 uint16 = 858
    v4 := []uint32{585, 65535, 0}
    var v5 int32 = 0
    var v22 uint16 = (v3 & (v3 | 2))
    if (4294967295 > 633) {
        v5 = v5
    }
    v3 = v22
    fmt.Println(v22)
    v22 = 65535
    if ((v5 << ((65535) & 63)) >= v5) {
        var v23 uint64 = v2
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(v22, v5, fi24)
            _ = fi24
        }
        _ = v23
    }
    if (1 >= 42) {
        v25 := GenS5{F0: ("foo bar" + "a"), F1: 2}
        _ = v25
    }
    v5 = v5
    v26 := GenS6{F0: false, F1: 73, F2: 32767}
    fmt.Println(v3, v3, v2)
    var v27 float64 = 0.1
    v28 := GenS7{F0: uint(v2)}
    switch 827 {
    case 720:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 292:
        if ((-2.25 - ((0.1 - 0.1) / 1.0)) < float32((v27 - v27))) {
            v22 = (v3 << ((0) & 63))
        } else {
            v4 = append(v4, uint32(v2))
            for fi29 := 0; fi29 < 1; fi29++ {
                fmt.Println(fi29)
                var v30 int64 = 148
                if true {
                    if true {
                        for fi31 := 0; fi31 < 20; fi31++ {
                            var v32 string = "a"
                            _ = fi31
                            _ = v32
                        }
                    }
                    fmt.Println(v22, v30, v2)
                    fmt.Println(v27)
                }
                _ = fi29
                _ = v30
            }
            fmt.Println(v5)
        }
    case 0:
        v33 := GenS7{F0: 775}
        _ = v33
    }
    v22 += (^(v3))
    v34 := GenS3{F0: v2, F1: (v5 & 2147483647), F2: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
