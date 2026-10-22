package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint8 }
type GenS2 struct { F0 int32; F1 int8; F2 bool }
func genFn0(p0 uint) uint {
    if false {
        for fi4 := 0; fi4 < 4; fi4++ {
            switch fi4 {
            case 32767:
                v5 := GenS1{F0: "\n\t", F1: p0, F2: 2}
                _ = v5
            case 42:
                var v6 int = (fi4 & fi4)
                _ = v6
            }
            _ = fi4
        }
    }
    v7 := []int{32767, 32767, -754}
    for fi8 := 0; fi8 < 18; fi8++ {
        if len(v7) > 0 {
            v7[0] = (fi8 >> (((p0 % ((p0) | 1))) & 63))
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 uint16) uint16 {
    p0 += p0
    fmt.Println(p0)
    var v9 string = ("hello" + "\n\t")
    _ = p0
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 int32 = 0
    for fi10 := 0; fi10 < 1; fi10++ {
        _ = len(v1)
        v11 := []int{1, (^(fi10))}
        _ = fi10
        _ = v11
    }
    switch 32767 {
    case -530:
        v12 := []int{-9223372036854775808}
        _ = v12
    case 238:
        v13 := []int{719, 128}
        _ = v13
    case -918:
        v3 += (v3 * v3)
    }
    var v14 uint64 = 1
    v14 += 0
    v3 += -1
    var v15 float32 = 1e10
    switch 2147483647 {
    case 4294967295:
        v14 = (v14 ^ v14)
    default:
        if (uint8(v14) >= 128) {
            delete(v1, 2)
        }
    }
    delete(v2, 65535)
    var v16 float32 = (((1.5 * v15) * 0.0) / v15)
    for fi17 := 0; fi17 < 10; fi17++ {
        fmt.Println(v14, v14)
        for fi18 := 0; fi18 < 3; fi18++ {
            if (!(255 >= 0)) {
                v19 := v1[2147483648]
                var v20 float32 = float32(3.14159)
                v21 := GenS1{F0: ("\n\t" + ("\n\t" + "foo bar")), F1: genFn0(65535), F2: 2}
                _ = v19
                _ = v20
                _ = v21
            }
            switch "a" {
            case "a":
                v22 := GenS1{F0: "a", F1: 2, F2: 197}
                _ = v22
            case "hello":
                v3 += (v3 >> ((428) & 63))
            case "\n\t":
                v23 := GenS1{F0: "a", F1: 42, F2: 255}
                _ = v23
            default:
                v3 = -852
            }
            v3 -= v3
            _ = fi18
        }
        v14 = 893
        _ = fi17
    }
    v15 -= (v15 - v15)
    fmt.Println(v16, v14)
    var v24 int32 = (v3 * (v3 + (v3 >> ((genFn0(1)) & 63))))
    for fi25 := 0; fi25 < 10; fi25++ {
        if (!true) {
            v26 := GenS2{F0: (v24 + -2147483648), F1: -128, F2: (0 > 80)}
            switch -319 {
            case 579:
                v16 += v15
            case 128:
                fmt.Println(v24, fi25)
            }
            _ = v26
        } else {
            var v27 uint8 = 139
            _ = v27
        }
        v24 -= v3
        _ = fi25
    }
    fmt.Println(v24, v14, v16)
    v14 -= (^(v14))
    fmt.Println(v16)
    var v28 float64 = (1.5 / 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v28
    _ = v3
}
