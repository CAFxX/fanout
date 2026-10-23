package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int16; F1 int64; F2 uint16 }
func genFn0(p0 int, p1 uint) string {
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: 128, F1: 0}
        var v9 uint16 = 42
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint8 = 42
    _ = p0
    _ = p1
    _ = v10
    return ""
}
func genFn1(p0 int8) int16 {
    if true {
        p0 = (p0 - (-(p0)))
        var v11 float64 = 1e10
        _ = v11
    } else {
        p0 -= -1
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        p0 += 0
        p0 = (p0 & (^(p0)))
        v13 := GenS2{F0: (246 <= 0), F1: (fi12 << ((18446744073709551615) & 63))}
        _ = fi12
        _ = v13
    }
    v14 := GenS3{F0: 42, F1: 0, F2: 65535}
    _ = p0
    _ = v14
    return 32767
}

func main() {
    var v1 int16 = 42
    var v2 uint16 = 0
    var v3 int = 42
    var v4 uint32 = 4294967295
    var v5 uint32 = 2147483648
    var v6 int32 = 0
    var v15 int64 = -9223372036854775808
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint16 = 1
        for fi18 := 0; fi18 < 3; fi18++ {
            if (!((((((3.14159 + 0.1) + -2.25) + ((1.5 * float64(1e10)) * (1.5 - (((0.0 * float64(1e10)) * (((((1.5 - float64((-2.25 - 3.14159))) + 1e10) / (1.5 - 0.0)) / -2.25) - 0.0)) * float64(3.14159))))) / -2.25) == 1.5) && (!(("\n\t" + "hello") == "hello")))) {
                v6 = v6
                if ((v15 | 9223372036854775807) > v15) {
                    v19 := GenS3{F0: v1, F1: v15, F2: v17}
                    var v20 uint8 = 2
                    _ = v19
                    _ = v20
                }
            }
            fmt.Println(fi18, v15)
            v6 += v6
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        switch -9223372036854775808 {
        case -534:
            v22 := []int{(v3 + fi21)}
            _ = v22
        }
        _ = fi21
    }
    v1 += genFn1(-1)
    var v23 uint8 = 0
    if false {
        v6 += -1
        var v24 uint = 487
        _ = v24
    }
    v25 := GenS2{F0: false, F1: v3}
    for fi26 := 0; fi26 < 20; fi26++ {
        v27 := GenS1{F0: v23, F1: 1}
        _ = fi26
        _ = v27
    }
    v28 := []int{(v3 % (((v3 - v3)) | 1)), (v3 | -294), v3}
    fmt.Println(v6, v4)
    var v29 int64 = (^(v15))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
