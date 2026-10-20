package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint16 }

func main() {
    var v1 float64 = -2.25
    var v2 uint64 = 730
    var v3 uint8 = 128
    v1 = (((float64(float32(v1)) / float64((3.14159 / (1.5 / ((0.0 - ((0.1 + 0.1) + 1e10)) / 1.0))))) * 1e10) * 1.5)
    var v4 int64 = 0
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := []int{-108, fi5, 4294967295}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int = (fi7 | 204)
        var v9 bool = false
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 65535
    if true {
        switch "foo bar" {
        case "\n\t":
            v11 := []int{-2147483648, 9223372036854775807}
            _ = v11
        case "":
            for fi12 := 0; fi12 < 14; fi12++ {
                if false {
                    var v13 int8 = 100
                    v3 = 131
                    _ = v13
                }
                _ = fi12
            }
        case "a":
            for fi14 := 0; fi14 < 2; fi14++ {
                v15 := []int{(fi14 / (((-(fi14))) | 1)), fi14, (fi14 % (((fi14 << ((0) & 63))) | 1))}
                if len(v15) > 0 {
                    v15[0] = fi14
                }
                for fi16 := 0; fi16 < 3; fi16++ {
                    v10 += 2147483647
                    _ = fi16
                }
                _ = fi14
                _ = v15
            }
        default:
            if false {
                v17 := []int{1, 1, -1}
                _ = v17
            } else {
                var v18 int = 339
                v19 := []int{}
                _ = v18
                _ = v19
            }
        }
    }
    var v20 int = -397
    if false {
        var v21 float32 = ((1.5 + -2.25) + (float32(float64(0.0)) * ((((1e10 + (1e10 - 0.0)) + float32(float64(1e10))) - (1e10 - ((1.5 + ((-2.25 * 1e10) - 3.14159)) / 0.1))) + 1e10)))
        _ = v21
    }
    var v22 bool = true
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 string = ""
        _ = fi23
        _ = v24
    }
    var v25 uint = 215
    switch "" {
    case "\n\t":
        if (3.14159 != ((1.5 * 0.0) * (v1 * 3.14159))) {
            if true {
                v4 -= v4
            } else {
                switch "foo bar" {
                case "hello":
                    if true {
                        v25 += uint(v2)
                    } else {
                        for fi26 := 0; fi26 < 19; fi26++ {
                            v27 := GenS1{F0: 1, F1: v4, F2: 65535}
                            _ = fi26
                            _ = v27
                        }
                    }
                case "\n\t":
                    v2 += (^(v2))
                }
            }
            v22 = (!(250 < 32767))
            var v28 int64 = (v4 / (((v4 * v4)) | 1))
            _ = v28
        } else {
            switch "a" {
            case "hello":
                var v29 int16 = 547
                _ = v29
            }
            v20 = v20
            var v30 uint32 = uint32(v2)
            _ = v30
        }
    case "hello":
        var v31 uint64 = v2
        _ = v31
    case "foo bar":
        v32 := []int{v20, v20, v20, 997}
        _ = v32
    }
    v1 = float64(((0.0 * float32(v1)) / ((float32((v1 / (v1 / v1))) * 0.1) + 0.1)))
    v33 := []int{v20, 32767}
    var v34 int32 = int32(v25)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v22
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
