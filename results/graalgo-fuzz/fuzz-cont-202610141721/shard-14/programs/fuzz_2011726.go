package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int32; F1 int64; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 int16; F2 uint }
func genFn0(p0 uint8, p1 uint8) uint {
    if false {
        for fi6 := 0; fi6 < 5; fi6++ {
            switch -9223372036854775808 {
            case 82:
                for fi7 := 0; fi7 < 11; fi7++ {
                    var v8 int32 = 1
                    _ = fi7
                    _ = v8
                }
            default:
                var v9 bool = true
                _ = v9
            }
            _ = fi6
        }
    }
    p0 = (p1 * 2)
    fmt.Println(p0, p1)
    for fi10 := 0; fi10 < 18; fi10++ {
        p1 = (p0 + p1)
        var v11 uint32 = 0
        var v12 uint8 = 89
        _ = fi10
        _ = v11
        _ = v12
    }
    p0 -= (p0 << ((255) & 63))
    _ = p0
    _ = p1
    return 103
}
func genFn1() int64 {
    v13 := GenS1{F0: 2, F1: 1}
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 string = ("hello" + "\n\t")
            v16 = v16
            v17 := []int{(fi15 % (((fi15 % ((-945) | 1))) | 1))}
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := []int{0, (fi14 * (fi14 ^ (fi14 | fi14)))}
        var v19 uint16 = 255
        _ = fi14
        _ = v18
        _ = v19
    }
    v20 := []int{-283, 42, 32768, 9223372036854775807}
    _ = v13
    _ = v20
    return -1
}

func main() {
    var v1 uint = 495
    var v2 bool = false
    v3 := make(map[string]bool)
    var v4 uint = 255
    var v5 uint32 = 2147483648
    v5 = (v5 & (v5 | v5))
    if v2 {
        var v21 int64 = 9223372036854775807
        v22 := GenS2{F0: (v21 + (v21 << ((v4) & 63))), F1: 787, F2: 42}
        var v23 int32 = 2147483647
        _ = v21
        _ = v22
        _ = v23
    }
    v2 = (233 > 128)
    v24 := GenS1{F0: 0, F1: -1}
    v25 := []int{128, 954, 2147483647, 128}
    v26 := v3["foo bar"]
    var v27 int = -725
    if ((("\n\t" + "\n\t") + (("a" + (("" + "hello") + ("\n\t" + "\n\t"))) + "a")) == "a") {
        v28 := GenS3{F0: -466, F1: 1, F2: -1}
        var v29 float32 = float32(1.5)
        _ = v28
        _ = v29
    }
    v5 = v5
    var v30 int16 = 1000
    v31 := GenS4{F0: 492}
    var v32 int32 = -1
    v33 := GenS5{F0: 438, F1: int16(v27), F2: (v1 << ((439) & 63))}
    var v34 uint32 = (v5 | (^(v5)))
    switch v27 {
    case 127:
        v34 = 2147483648
    case 4294967295:
        var v35 float32 = (((-2.25 - float32(-2.25)) - float32(1e10)) + 1e10)
        _ = v35
    default:
        v27 = 1
    }
    v32 = (v32 / (((v32 + v32)) | 1))
    var v36 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
}
