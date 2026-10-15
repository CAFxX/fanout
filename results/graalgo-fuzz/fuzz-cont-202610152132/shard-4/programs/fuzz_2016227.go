package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]uint32)
    var v3 int16 = -1
    if false {
        v1 = v1
        for fi4 := 0; fi4 < 15; fi4++ {
            switch 430 {
            case 32768:
                v3 += (v3 % (((v3 + (v3 / ((v3) | 1)))) | 1))
            }
            fmt.Println(fi4, v3)
            _ = fi4
        }
    }
    v1 -= v1
    for fi5 := 0; fi5 < 19; fi5++ {
        _ = len(v2)
        v3 = (v3 / ((-1) | 1))
        var v6 float32 = 1e10
        _ = fi5
        _ = v6
    }
    switch 965 {
    case 2147483647:
        v1 = v1
    default:
        if false {
            var v7 int64 = v1
            v8 := GenS1{F0: "foo bar"}
            var v9 int = 672
            _ = v7
            _ = v8
            _ = v9
        } else {
            v1 = v1
            var v10 int32 = 855
            _ = v10
        }
    }
    if true {
        var v11 uint8 = 42
        var v12 uint32 = 42
        switch 998 {
        case -1:
            var v13 int = 42
            _ = v13
        case 2147483647:
            var v14 uint = 4294967295
            _ = v14
        case 42:
            fmt.Println(v12, v11)
        }
        _ = v11
        _ = v12
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(fi15, fi15)
            _ = fi15
        }
    }
    v16 := GenS1{F0: ""}
    fmt.Println(v1)
    delete(v2, 127)
    if true {
        v17 := GenS2{F0: 475}
        for fi18 := 0; fi18 < 15; fi18++ {
            if ("foo bar" >= "foo bar") {
                var v19 int32 = -571
                var v20 float32 = 3.14159
                v21 := GenS3{F0: 1, F1: 2, F2: v1}
                _ = v19
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        _ = v17
    } else {
        var v22 float32 = float32(float64(((0.1 * ((0.0 - ((3.14159 * (1e10 * (0.0 * (1e10 / -2.25)))) * (0.1 / 1.0))) - -2.25)) * 1.5)))
        v23 := []int{32768, 2147483647, 487, 42}
        _ = v22
        _ = v23
    }
    v1 = v1
    v1 -= 1
    fmt.Println(v1)
    v24 := GenS3{F0: 1, F1: 255, F2: 0}
    v25 := GenS4{F0: 671}
    v3 = (v3 * 1000)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
