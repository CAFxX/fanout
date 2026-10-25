package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 string; F2 uint }
type GenS4 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int32, p1 int32) int32 {
    v7 := GenS1{F0: 65535, F1: 255, F2: 4294967295}
    v8 := GenS1{F0: 804, F1: 252, F2: 358}
    v9 := GenS1{F0: 174, F1: 32767, F2: 20}
    switch 1 {
    case 203:
        var v10 uint32 = 0
        _ = v10
    case -914:
        if (3.14159 == 1e10) {
            p1 = p0
        } else {
            for fi11 := 0; fi11 < 20; fi11++ {
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 uint8 = 0
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
            var v14 int64 = -9223372036854775808
            _ = v14
        }
    default:
        v15 := GenS1{F0: 282, F1: 872, F2: 281}
        _ = v15
    }
    var v16 uint32 = 0
    _ = p0
    _ = p1
    _ = v16
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 uint32 = 0
    var v2 int16 = -32768
    var v3 uint = 52
    v4 := []uint32{v1, v1, 2147483648, v1}
    var v5 int64 = -938
    v6 := make(map[uint32]int)
    var v17 int16 = -533
    v18 := GenS2{F0: (1e10 <= 3.14159)}
    var v19 bool = (128 <= 0)
    v3 -= (v3 & v3)
    v20 := GenS3{F0: (-(v2)), F1: "a", F2: v3}
    _ = len(v4)
    if false {
        var v21 int64 = (v5 * (v5 << (((v3 << ((v3) & 63))) & 63)))
        if ((!false) && (-2.25 > 1.5)) {
            if v19 {
                var v22 int32 = 42
                switch -9223372036854775808 {
                case 875:
                    fmt.Println(v2, v22)
                case 32767:
                    v4 = append(v4, 65535)
                default:
                    fmt.Println(v19, v17)
                }
                _ = v22
            } else {
                v23 := GenS3{F0: (v17 % ((382) | 1)), F1: "\n\t", F2: v3}
                _ = v23
            }
            var v24 uint = (v3 * v3)
            v3 = 864
            _ = v24
        } else {
            v5 = (v5 * int64(v1))
            v17 = v2
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            v19 = ((-(v2)) <= v2)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi25
        }
        _ = v21
    } else {
        if ((0 == 4294967295) && v19) {
            v1 = v1
            v26 := GenS4{F0: 18446744073709551615, F1: 1}
            _ = v26
        }
    }
    var v27 uint16 = 32767
    delete(v6, v1)
    v27 = 2
    for fi28 := 0; fi28 < 16; fi28++ {
        fmt.Println(v19)
        _ = fi28
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        fmt.Println(v3)
        _ = fi29
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        v4 = append(v4, v1)
        switch fi30 {
        case -9223372036854775808:
            switch 4294967295 {
            case -763:
                var v31 uint = (v3 - (v3 | (v3 >> ((v3) & 63))))
                _ = v31
            case 2147483647:
                v6[uint32(v3)] = fi30
            default:
                _ = len(v6)
            }
        default:
            if len(v4) > 0 {
                v4[0] = 863
            }
        }
        _ = fi30
    }
    switch "a" {
    case "a":
        v32 := GenS3{F0: (v17 << (((v3 | (v3 >> (((v3 * 108)) & 63)))) & 63)), F1: (("" + "a") + "a"), F2: (v3 | (v3 & v3))}
        _ = v32
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
