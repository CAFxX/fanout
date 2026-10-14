package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint16 = 612
    var v2 uint = 4294967295
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]string)
    var v5 int64 = v3
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 15; fi7++ {
        _ = len(v6)
        fmt.Println(v1, v3, fi7)
        _ = fi7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(v1, v1)
        switch -630 {
        case 2147483647:
            if (false && (128 > 255)) {
                switch (fi8 ^ fi8) {
                case 4294967295:
                    if false {
                        v9 := GenS1{F0: 562, F1: 1, F2: 776}
                        _ = v9
                    } else {
                        v10 := v6[fi8]
                        _ = v10
                    }
                case 9223372036854775807:
                    var v11 uint16 = (v1 ^ (v1 / ((uint16(v2)) | 1)))
                    _ = v11
                case 2147483647:
                    v2 = 255
                }
            }
        default:
            switch fi8 {
            case 4294967295:
                v2 += (v2 - (^(v2)))
            default:
                v12 := GenS1{F0: 2147483648, F1: 428, F2: 0}
                _ = v12
            }
        }
        _ = fi8
    }
    v13 := []int{525}
    if ((1 < (v2 ^ (^(v2)))) || (float32(((float64((1.5 * ((3.14159 - -2.25) - 0.0))) * 3.14159) - (float64(1e10) + 1.5))) > (float32((-2.25 - (0.1 - 1e10))) / 1e10))) {
        if true {
            var v14 uint32 = 42
            _ = v14
        }
        var v15 float64 = (1.5 / 1.0)
        v3 += (v5 / ((-502) | 1))
        _ = v15
    } else {
        var v16 int32 = 42
        _ = v16
    }
    v17 := GenS1{F0: 2147483648, F1: 42, F2: -32768}
    switch 42 {
    case 139:
        var v18 bool = ((-32768 >= -642) && true)
        _ = v18
    default:
        v2 -= v2
    }
    if true {
        if len(v13) > 0 {
            v13[0] = 90
        }
    } else {
        var v19 int = -1
        for fi20 := 0; fi20 < 3; fi20++ {
            for fi21 := 0; fi21 < 8; fi21++ {
                v4[("foo bar" + "")] = ""
                if (0.0 >= float64(float32((((float64((float32(float64((1e10 + -2.25))) / -2.25)) * 0.0) / 1.5) * 1.5)))) {
                    fmt.Println(fi20, fi21, v1)
                    var v22 int8 = -30
                    _ = v22
                }
                _ = fi21
            }
            _ = fi20
        }
        _ = v19
    }
    switch "hello" {
    case "foo bar":
        if (v5 <= v3) {
            var v23 int32 = -2147483648
            _ = v23
        } else {
            var v24 uint32 = 743
            var v25 int = 32767
            var v26 uint8 = 1
            _ = v24
            _ = v25
            _ = v26
        }
    case "\n\t":
        if true {
            v27 := GenS2{F0: (^(v1))}
            _ = v27
        }
    case "":
        var v28 uint = (v2 / ((18446744073709551615) | 1))
        _ = v28
    }
    if false {
        var v29 uint16 = 42
        var v30 int = int(v2)
        _ = v29
        _ = v30
    } else {
        v31 := GenS2{F0: 657}
        v5 -= v5
        switch "a" {
        case "\n\t":
            var v32 int16 = -32768
            _ = v32
        case "a":
            switch -2147483648 {
            case -1:
                v33 := GenS2{F0: v1}
                _ = v33
            default:
                v3 += (v5 / (((v5 ^ v5)) | 1))
            }
        default:
            var v34 bool = true
            _ = v34
        }
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
