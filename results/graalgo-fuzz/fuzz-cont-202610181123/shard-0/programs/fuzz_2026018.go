package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint32; F1 int16; F2 bool }
func genFn0(p0 uint32, p1 int64) uint8 {
    v7 := GenS1{F0: 912, F1: 1}
    var v8 uint64 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 string = ("hello" + ("a" + "hello"))
    var v2 uint64 = 1
    var v3 float64 = (-2.25 * float64((1.5 * (1.5 / (((float32((-2.25 / (1e10 + 0.0))) / -2.25) * 1.5) - 0.1)))))
    v4 := make(map[int]uint32)
    var v5 int16 = -562
    v6 := []string{(v1 + v1)}
    fmt.Println(v3)
    switch 0 {
    case -29:
        if true {
            var v9 uint64 = (^(v2))
            _ = v9
        }
    case -2147483648:
        var v10 bool = false
        _ = v10
    default:
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int64 = 0
            if false {
                fmt.Println(fi11)
                v13 := GenS1{F0: 1, F1: 0}
                v14 := GenS1{F0: 18446744073709551615, F1: uint32(v2)}
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 13; fi15++ {
                    v12 -= v12
                    var v16 int = (^(fi11))
                    switch 2147483647 {
                    case 4294967295:
                        var v17 bool = (((-2147483648 < 1) && (65535 > 1)) && ((!false) || true))
                        _ = v17
                    default:
                        v5 = -32768
                    }
                    v12 += v12
                    if (!true) {
                        v2 = v2
                        var v18 int = 9223372036854775807
                        v19 := GenS2{F0: 89}
                        _ = v18
                        _ = v19
                    }
                    var v20 bool = true
                    _ = fi15
                    _ = v16
                    _ = v20
                }
            }
            _ = fi11
            _ = v12
        }
    }
    if false {
        var v21 uint64 = v2
        _ = v21
    } else {
        v22 := GenS1{F0: 203, F1: 65535}
        for fi23 := 0; fi23 < 20; fi23++ {
            switch (fi23 >> ((583) & 63)) {
            case 0:
                v3 += v3
            case 127:
                var v24 uint64 = v2
                _ = v24
            case 4294967295:
                switch (v1 + "") {
                case "\n\t":
                    v25 := GenS3{F0: 255}
                    _ = v25
                default:
                    if (4294967295 == 0) {
                        _ = len(v4)
                        v3 -= 1.5
                    } else {
                        var v26 int32 = 0
                        _ = v26
                    }
                }
            }
            v27 := v4[-44]
            v28 := GenS2{F0: -1}
            _ = fi23
            _ = v27
            _ = v28
        }
        _ = v22
    }
    if len(v6) > 0 {
        v6[0] = ("foo bar" + (v1 + v1))
    }
    if (0 != -128) {
        v29 := GenS4{F0: v5, F1: (v5 * (^(v5)))}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v29
    } else {
        if ((1 <= -1) || (0 != -128)) {
            v6 = append(v6, v1)
            _ = len(v6)
        }
        v1 = v1
        var v30 int64 = 828
        _ = v30
    }
    v3 = 1.5
    var v31 int32 = 341
    v5 += 42
    var v32 int32 = (v31 ^ (v31 & 134))
    switch 127 {
    case 9223372036854775807:
        var v33 int16 = (^(v5))
        _ = v33
    case -407:
        var v34 int16 = 0
        _ = v34
    }
    switch -2147483648 {
    case -9223372036854775808:
        fmt.Println(v1, v5, v31)
    default:
        for fi35 := 0; fi35 < 13; fi35++ {
            var v36 uint16 = 370
            fmt.Println(v1, v2)
            _ = fi35
            _ = v36
        }
    }
    v37 := v4[4294967295]
    fmt.Println(v1, v2)
    v38 := GenS5{F0: (v37 & v37), F1: v5, F2: false}
    v39 := GenS5{F0: v37, F1: -161, F2: ((false || true) && (3.14159 >= v3))}
    v40 := GenS1{F0: 410, F1: v37}
    for fi41 := 0; fi41 < 2; fi41++ {
        v37 = v37
        v1 = ((("foo bar" + (v1 + "a")) + ((v1 + v1) + "\n\t")) + v1)
        _ = fi41
    }
    var v42 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v6
}
