package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int; F1 uint8; F2 string }

func main() {
    var v1 int8 = 13
    var v2 float32 = (float32(0.1) + 1e10)
    v3 := []int{4294967295}
    var v4 int64 = -9223372036854775808
    var v5 uint64 = 2
    v1 = -79
    switch 304 {
    case 127:
        for fi6 := 0; fi6 < 19; fi6++ {
            switch fi6 {
            case 127:
                v4 += v4
            case 128:
                v2 = v2
            }
            _ = fi6
        }
    case 32767:
        fmt.Println(v2, v4, v5)
    case 9223372036854775807:
        if true {
            var v7 uint32 = 484
            fmt.Println(v4, v5, v1)
            _ = v7
        }
    }
    switch 42 {
    case -616:
        if (32767 <= 42) {
            _ = len(v3)
            var v8 int64 = 896
            _ = v8
        }
    }
    var v9 bool = (!((v4 << ((18446744073709551615) & 63)) > (^(v4))))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v10 := GenS1{F0: "\n\t", F1: int32(v4)}
    for fi11 := 0; fi11 < 15; fi11++ {
        switch -9223372036854775808 {
        case 1:
            switch "" {
            case "\n\t":
                var v12 int64 = (v4 >> ((255) & 63))
                _ = v12
            case "":
                fmt.Println(v2)
            default:
                fmt.Println(fi11)
            }
        case -9223372036854775808:
            if false {
                var v13 float64 = (-2.25 / -2.25)
                for fi14 := 0; fi14 < 11; fi14++ {
                    v4 += v4
                    switch (-(fi11)) {
                    case 127:
                        if (v9 || (!false)) {
                            fmt.Println(fi11, v2)
                            v15 := GenS1{F0: (((("\n\t" + "") + "foo bar") + "hello") + "a"), F1: 2147483647}
                            _ = v15
                        }
                    default:
                        fmt.Println(fi14)
                    }
                    var v16 float64 = v13
                    _ = fi14
                    _ = v16
                }
                _ = v13
            }
        default:
            switch "a" {
            case "a":
                var v17 float64 = float64(((((0.1 / v2) * v2) + (((v2 + 0.0) - float32((((1e10 + -2.25) * 1.5) + (1e10 * float64((v2 - v2)))))) * ((v2 / 1e10) - (v2 - (v2 - -2.25))))) / 1.5))
                _ = v17
            case "":
                v2 += v2
            }
        }
        v1 = (^(v1))
        _ = fi11
    }
    var v18 float64 = float64(1.5)
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS2{F0: (fi19 / ((fi19) | 1)), F1: 1, F2: "a"}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi19
        _ = v20
    }
    var v21 string = "\n\t"
    switch 32768 {
    case -742:
        v3 = append(v3, 128)
    case -313:
        v2 = 0.0
    case -36:
        var v22 bool = (v9 && (true && false))
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        if len(v3) > 0 {
            v3[0] = fi23
        }
        var v24 uint64 = v5
        _ = fi23
        _ = v24
    }
    var v25 uint64 = v5
    for fi26 := 0; fi26 < 13; fi26++ {
        v27 := GenS2{F0: fi26, F1: 2, F2: ("" + "")}
        v5 += (v5 % ((386) | 1))
        v9 = v9
        _ = fi26
        _ = v27
    }
    var v28 uint64 = v25
    var v29 uint64 = v28
    v2 += v2
    for fi30 := 0; fi30 < 7; fi30++ {
        var v31 int16 = int16(v4)
        v32 := GenS1{F0: "a", F1: 42}
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v21
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
