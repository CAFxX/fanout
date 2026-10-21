package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint16 }
func genFn0() bool {
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := []int{}
        var v7 float64 = -2.25
        _ = fi5
        _ = v6
        _ = v7
    }
    if (true || false) {
        if (128 == 42) {
            v8 := []int{0, 127, 9223372036854775807, -576}
            var v9 int = -9223372036854775808
            switch v9 {
            case 42:
                v9 = (v9 ^ v9)
            case 1:
                if false {
                    var v10 bool = (float32(0.1) > (3.14159 - 3.14159))
                    v9 = v9
                    var v11 uint = 680
                    _ = v10
                    _ = v11
                } else {
                    switch ("foo bar" + "a") {
                    case "\n\t":
                        v9 -= -2147483648
                    case "hello":
                        for fi12 := 0; fi12 < 6; fi12++ {
                            v9 = (fi12 / (((v9 - 9223372036854775807)) | 1))
                            _ = fi12
                        }
                    case "":
                        v13 := GenS1{F0: 0, F1: 0}
                        _ = v13
                    }
                    v9 += -1
                }
            }
            _ = v8
            _ = v9
        } else {
            v14 := []int{587, 127, -2147483648}
            for fi15 := 0; fi15 < 7; fi15++ {
                fmt.Println(fi15)
                for fi16 := 0; fi16 < 6; fi16++ {
                    switch 4294967295 {
                    case 42:
                        fmt.Println(fi16)
                    case 160:
                        switch "foo bar" {
                        case "hello":
                            var v17 uint8 = 128
                            _ = v17
                        }
                    case 207:
                        fmt.Println(fi16, fi15)
                    default:
                        var v18 bool = false
                        _ = v18
                    }
                    fmt.Println(fi15)
                    _ = fi16
                }
                _ = fi15
            }
            _ = v14
        }
        var v19 uint16 = 255
        _ = v19
    } else {
        var v20 uint = 0
        v20 -= v20
        v20 = 65535
        _ = v20
    }
    return (float64(1.5) == 0.1)
}
func genFn1(p0 int16, p1 uint64, p2 int64) uint {
    p2 += p2
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 int64 = 0
        v23 := GenS1{F0: 370, F1: 2147483647}
        var v24 uint8 = 171
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := []uint32{164}
    v4 := []int{-2147483648}
    v25 := GenS1{F0: 0, F1: 138}
    v26 := GenS2{F0: 65535, F1: 42, F2: 255}
    var v27 uint32 = 0
    v28 := v2[127]
    v29 := GenS2{F0: 65535, F1: 1, F2: 216}
    fmt.Println(v27)
    _ = len(v3)
    v27 = v27
    v27 -= 350
    v27 -= 255
    delete(v2, -327)
    v28 = (^(v27))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
