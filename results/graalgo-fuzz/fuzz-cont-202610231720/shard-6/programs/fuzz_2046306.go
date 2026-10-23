package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 string }
func genFn0(p0 uint8) int32 {
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = (p0 << ((303) & 63))
    _ = p0
    return -382
}

func main() {
    var v1 float32 = (((float32(0.1) + float32(3.14159)) - ((-2.25 * 0.0) * float32(1e10))) * 0.1)
    v2 := make(map[string]bool)
    v3 := []uint32{}
    var v4 uint64 = 18446744073709551615
    for fi5 := 0; fi5 < 6; fi5++ {
        if true {
            switch "a" {
            case "a":
                if false {
                    for fi6 := 0; fi6 < 6; fi6++ {
                        _ = len(v2)
                        v7 := GenS1{F0: 644, F1: v4}
                        _ = fi6
                        _ = v7
                    }
                    _ = len(v2)
                }
            }
            v8 := GenS2{F0: 0}
            v4 = v4
            _ = v8
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(v1, v4, fi9)
            v10 := GenS2{F0: 0}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 15; fi11++ {
            if true {
                var v12 string = "\n\t"
                switch fi11 {
                case 977:
                    var v13 bool = (false || true)
                    _ = v13
                }
                var v14 int64 = 42
                _ = v12
                _ = v14
            }
            v4 -= (v4 + (v4 % (((v4 * (v4 | (v4 % ((v4) | 1))))) | 1)))
            _ = fi11
        }
        _ = fi5
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        switch fi15 {
        case 42:
            v4 -= 4294967295
        default:
            switch "hello" {
            case "foo bar":
                v1 -= ((v1 / v1) * ((((v1 - v1) * float32(((1.5 / 1.0) / float64(0.1)))) / -2.25) + 3.14159))
            case "a":
                fmt.Println(v4, v4, fi15)
            default:
                if (2 < 217) {
                    v16 := GenS2{F0: -9223372036854775808}
                    fmt.Println(fi15)
                    _ = v16
                }
            }
        }
        _ = fi15
    }
    if true {
        var v17 uint = 1
        _ = v17
    }
    switch "" {
    case "hello":
        var v18 uint8 = 159
        _ = v18
    default:
        v4 += v4
    }
    v1 = v1
    v19 := GenS3{F0: "", F1: 255}
    v20 := GenS2{F0: 1}
    if (32767 >= 255) {
        v4 -= v4
        var v21 string = ("foo bar" + "\n\t")
        v1 += float32(-2.25)
        _ = v21
    }
    v22 := GenS4{F0: "\n\t"}
    switch "\n\t" {
    case "":
        delete(v2, "foo bar")
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
}
