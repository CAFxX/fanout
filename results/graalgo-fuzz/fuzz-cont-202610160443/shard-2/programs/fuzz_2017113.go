package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint32) int64 {
    if (!true) {
        v7 := []int{844, 128}
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        p0 -= 258
        fmt.Println(fi8, fi8, fi8)
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 float64 = (1e10 / 1.5)
            fmt.Println(fi8, v10, fi9)
            fmt.Println(p0)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = v11
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 255
    var v2 float32 = (-2.25 - 0.0)
    var v3 uint8 = v1
    var v4 uint64 = uint64(v3)
    var v5 uint = 1
    v6 := make(map[string]string)
    v12 := GenS2{F0: 127, F1: 4294967295, F2: (v5 ^ v5)}
    v3 -= (v1 - (v3 / ((uint8(v5)) | 1)))
    if true {
        switch 32767 {
        case 32768:
            if false {
                fmt.Println(v4)
            } else {
                v1 -= (v3 >> (((v5 % (((^(v5))) | 1))) & 63))
            }
        }
        var v13 int32 = 503
        v13 -= v13
        _ = v13
    } else {
        var v14 int16 = -251
        _ = v14
    }
    v4 = v4
    v15 := GenS1{F0: 0}
    var v16 uint8 = (v1 - 2)
    switch 32768 {
    case -61:
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := []int{fi17}
            v19 := GenS2{F0: 0, F1: 42, F2: 42}
            _ = len(v18)
            _ = fi17
            _ = v18
            _ = v19
        }
    case 2147483647:
        switch 4294967295 {
        case 0:
            v5 = v5
        case 128:
            v5 = (v5 - (v5 * 42))
        }
    case 4294967295:
        v4 = 0
    default:
        switch ("" + "foo bar") {
        case "\n\t":
            v4 = (v4 + (^(v4)))
        case "foo bar":
            switch "foo bar" {
            case "hello":
                v20 := GenS3{F0: -2147483648}
                _ = v20
            }
        }
    }
    v1 -= (v3 + 0)
    var v21 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
