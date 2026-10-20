package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16; F1 int8; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }

func main() {
    var v1 uint32 = 42
    var v2 uint8 = 255
    v3 := []string{((("" + "") + "a") + ("foo bar" + "a"))}
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 uint16 = 255
        _ = fi4
        _ = v5
    }
    if true {
        if (uint16(v1) < 0) {
            var v6 uint64 = 2
            v1 = v1
            fmt.Println(v2, v2, v6)
            _ = v6
        }
        var v7 string = ""
        _ = v7
    }
    v8 := GenS1{F0: -1, F1: 0}
    v9 := GenS2{F0: 1000, F1: 1, F2: 1000}
    var v10 int16 = -516
    fmt.Println(v2, v1, v1)
    v2 = v2
    v11 := GenS2{F0: 255, F1: 127, F2: v10}
    for fi12 := 0; fi12 < 1; fi12++ {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = fi12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        if (42 > v10) {
            var v14 uint32 = uint32(v2)
            for fi15 := 0; fi15 < 17; fi15++ {
                if false {
                    v16 := GenS3{F0: 42}
                    _ = v16
                }
                _ = fi15
            }
            _ = len(v3)
            _ = v14
        } else {
            switch "" {
            case "hello":
                for fi17 := 0; fi17 < 6; fi17++ {
                    v10 = v10
                    var v18 int16 = 531
                    _ = fi17
                    _ = v18
                }
            default:
                if (-895 >= -316) {
                    v19 := GenS2{F0: 82, F1: 46, F2: (v10 * (v10 >> ((255) & 63)))}
                    var v20 uint32 = v1
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS4{F0: -1}
                    v1 = 815
                    _ = v21
                }
            }
        }
        _ = fi13
    }
    switch -416 {
    case 32768:
        var v22 uint16 = 65535
        _ = v22
    case 9223372036854775807:
        fmt.Println(v2, v2, v2)
    case -592:
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS1{F0: 42, F1: 1}
            v25 := GenS4{F0: (fi23 / ((fi23) | 1))}
            for fi26 := 0; fi26 < 11; fi26++ {
                fmt.Println(fi26, v1)
                v27 := GenS2{F0: 65535, F1: -128, F2: v10}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi26
                _ = v27
            }
            _ = fi23
            _ = v24
            _ = v25
        }
    default:
        var v28 float32 = (float32(1e10) - -2.25)
        _ = v28
    }
    v10 += (v10 / (((v10 >> ((255) & 63))) | 1))
    _ = len(v3)
    for fi29 := 0; fi29 < 16; fi29++ {
        if false {
            var v30 uint64 = 2
            _ = v30
        } else {
            var v31 float64 = ((1e10 * float64(float32(0.0))) - 1e10)
            _ = v31
        }
        v3 = append(v3, ("\n\t" + "hello"))
        var v32 string = ""
        _ = fi29
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
