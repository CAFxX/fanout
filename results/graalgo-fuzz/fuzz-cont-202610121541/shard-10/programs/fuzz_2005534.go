package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int64 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint32; F2 int64 }
func genFn0(p0 int32, p1 uint64) uint32 {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 uint64 = p1
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v8 string = ((("\n\t" + "hello") + "hello") + (("\n\t" + ("a" + (("\n\t" + "hello") + ("hello" + ("\n\t" + "foo bar"))))) + "hello"))
    p1 = 2
    _ = p0
    _ = p1
    _ = v8
    return 991
}

func main() {
    v1 := []uint32{2147483648, 65535, 42, 0}
    var v2 int32 = 1
    var v3 uint8 = 128
    v4 := make(map[uint32]uint32)
    switch "a" {
    case "\n\t":
        v2 -= v2
    }
    var v9 int32 = (v2 * (v2 ^ -1))
    fmt.Println(v3)
    for fi10 := 0; fi10 < 7; fi10++ {
        switch "\n\t" {
        case "hello":
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS1{F0: 1, F1: 4294967295, F2: 42}
                v13 := GenS2{F0: -128, F1: 708}
                v2 = v2
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        if (65535 >= 65535) {
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 float64 = float64(-2.25)
                v3 = 182
                _ = fi14
                _ = v15
            }
            v16 := GenS3{F0: 974, F1: 927, F2: -9223372036854775808}
            var v17 int16 = 1000
            _ = v16
            _ = v17
        }
        _ = fi10
    }
    if true {
        var v18 uint = 65535
        var v19 uint8 = 38
        v20 := GenS1{F0: 1, F1: 0, F2: int64(v18)}
        _ = v18
        _ = v19
        _ = v20
    }
    v9 += (v2 << ((930) & 63))
    v9 -= v9
    v2 += (v9 << ((2) & 63))
    v2 = (v9 << ((0) & 63))
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 string = ("a" + "a")
        _ = fi21
        _ = v22
    }
    var v23 uint32 = 37
    var v24 int16 = 930
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
