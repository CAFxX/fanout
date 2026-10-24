package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int32 }
type GenS2 struct { F0 int32; F1 uint; F2 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 uint16; F2 string }
type GenS6 struct { F0 uint64; F1 int64 }

func main() {
    var v1 string = "hello"
    var v2 uint = 1
    v3 := make(map[string]uint32)
    var v4 int64 = 0
    delete(v3, ("\n\t" + v1))
    var v5 float64 = float64(3.14159)
    if (false && ((1.5 >= (1.5 + (((float32(3.14159) - (3.14159 * 3.14159)) / 0.1) / 1e10))) && false)) {
        v2 = 2
    }
    switch "" {
    case "":
        v6 := []int{}
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        v1 = ("hello" + v1)
        v8 := []int{fi7}
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 float32 = -2.25
            v11 := GenS1{F0: fi9, F1: 42, F2: -1}
            v12 := GenS2{F0: 928, F1: v2, F2: false}
            var v13 float32 = ((-2.25 * 3.14159) + (((v10 * (v10 / float32(v5))) * -2.25) / (0.0 * v10)))
            fmt.Println(v1, v2)
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := []int{(fi14 - 32767), (^(fi14)), fi14, fi14}
        _ = fi14
        _ = v15
    }
    switch 698 {
    case 387:
        var v16 uint = v2
        _ = v16
    case 0:
        v17 := GenS3{F0: 153, F1: 233, F2: v2}
        _ = v17
    case -9223372036854775808:
        for fi18 := 0; fi18 < 9; fi18++ {
            if (((^(v2)) != v2) || (!((v2 - (v2 % ((453) | 1))) != 1))) {
                if (!((float32((v5 / ((v5 + 0.1) * ((1.5 * v5) / v5)))) + (0.0 - ((-2.25 + (1e10 * 0.1)) / 1e10))) != 3.14159)) {
                    v2 -= (v2 - (v2 << ((4294967295) & 63)))
                    delete(v3, v1)
                    if (fi18 < (fi18 - fi18)) {
                        fmt.Println(fi18, fi18)
                    }
                }
                var v19 uint8 = 255
                _ = v19
            } else {
                for fi20 := 0; fi20 < 19; fi20++ {
                    var v21 int8 = 42
                    var v22 uint32 = 42
                    switch (fi18 * fi18) {
                    case 216:
                        switch (fi20 + -243) {
                        case 1:
                            v2 += (v2 | v2)
                        case 790:
                            switch 331 {
                            case 127:
                                v1 = (((v1 + v1) + "foo bar") + "a")
                            case -456:
                                fmt.Println(fi20, v22)
                            default:
                                delete(v3, v1)
                            }
                        }
                    }
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v2 = 18446744073709551615
                v23 := GenS3{F0: 0, F1: 2147483647, F2: (v2 & 792)}
                _ = v23
            }
            switch fi18 {
            case 128:
                v5 -= float64(0.1)
            case 4294967295:
                v2 = v2
            case -1:
                if false {
                    fmt.Println(v2, v2, v4)
                }
            }
            _ = fi18
        }
    }
    v5 -= float64(float32(v5))
    v24 := GenS4{F0: -84}
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := v3[(v1 + v1)]
        _ = fi25
        _ = v26
    }
    fmt.Println(v5, v5, v2)
    var v27 int = 2147483647
    var v28 uint32 = 533
    v29 := GenS5{F0: true, F1: 2, F2: v1}
    if (!(false && false)) {
        var v30 int64 = -9223372036854775808
        _ = v30
    } else {
        if false {
            v31 := GenS6{F0: 4294967295, F1: (v4 / ((v4) | 1))}
            _ = v31
        }
        switch int(v28) {
        case -242:
            v32 := v3[v1]
            _ = v32
        }
        fmt.Println(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
