package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    if (255 >= 65535) {
        v1 -= (^(v1))
        var v4 string = "\n\t"
        v1 -= 338
        _ = v4
    }
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int16 = 32767
        v7 := GenS1{F0: (!true), F1: 392}
        var v8 uint16 = 284
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := v2[2147483647]
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint8 = 164
        v1 = (^(v1))
        _ = fi10
        _ = v11
    }
    fmt.Println(v9, v1)
    v12 := GenS1{F0: false, F1: 2}
    v13 := []int{}
    if (42 > int32(v1)) {
        if true {
            var v14 int16 = 677
            v1 += (v1 << ((655) & 63))
            _ = v14
        } else {
            var v15 float64 = 1.5
            if (!true) {
                v15 -= v15
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            } else {
                if len(v13) > 0 {
                    v13[0] = 185
                }
            }
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 uint64 = 4294967295
                _ = fi16
                _ = v17
            }
            _ = v15
        }
        fmt.Println(v9, v9)
        var v18 string = v9
        _ = v18
    } else {
        _ = len(v2)
        v1 += v1
    }
    switch -814 {
    case 148:
        v19 := GenS1{F0: (-886 < -1), F1: 748}
        _ = v19
    case 2147483647:
        v9 = v9
    default:
        if true {
            var v20 int32 = -2147483648
            v20 -= v20
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 uint64 = v1
        var v23 int32 = 42
        for fi24 := 0; fi24 < 13; fi24++ {
            v23 = (v23 * (-(v23)))
            _ = fi24
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        var v26 int = 9223372036854775807
        var v27 int8 = -128
        _ = fi25
        _ = v26
        _ = v27
    }
    v9 = v9
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
