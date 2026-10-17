package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint8 }
func genFn0(p0 bool) uint16 {
    fmt.Println(p0)
    v5 := GenS1{F0: (p0 || ((228 <= 269) && false)), F1: 17, F2: 255}
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int32 = -2147483648
        var v8 int8 = 127
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := []int{128}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 uint8 = 42
    var v2 float64 = ((float64(1e10) / (((1.5 + (float64(0.0) * (float64((1e10 + 0.0)) + 1.5))) + (0.0 * (float64(((1e10 - 1.5) - 0.1)) - -2.25))) / 1.0)) / -2.25)
    var v3 uint64 = 963
    var v4 uint16 = uint16(v1)
    v11 := GenS1{F0: (681 != 2), F1: (v1 % (((v1 - 255)) | 1)), F2: (^(v1))}
    if true {
        v2 = v2
    } else {
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 int = fi12
            fmt.Println(v2, fi12)
            _ = fi12
            _ = v13
        }
    }
    var v14 uint8 = v1
    v15 := GenS1{F0: false, F1: 211, F2: (^(v14))}
    var v16 int32 = 2147483647
    var v17 int64 = 0
    var v18 uint32 = 2
    for fi19 := 0; fi19 < 4; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := GenS1{F0: false, F1: (v1 >> ((608) & 63)), F2: 127}
            v1 = (v1 | (^(v14)))
            v14 = v1
            _ = fi20
            _ = v21
        }
        v18 = (v18 * (v18 * 2))
        v3 += (v3 ^ v3)
        _ = fi19
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        v23 := GenS1{F0: false, F1: (^(v1)), F2: 128}
        var v24 int = (^(fi22))
        if true {
            v18 = (v18 ^ v18)
        } else {
            var v25 uint8 = v1
            fmt.Println(v24, v2)
            for fi26 := 0; fi26 < 13; fi26++ {
                v4 += 894
                switch (^(fi22)) {
                case 470:
                    fmt.Println(fi26, v14, v25)
                case -984:
                    var v27 uint32 = v18
                    _ = v27
                default:
                    switch (fi22 >> ((42) & 63)) {
                    case 32768:
                        if true {
                            if true {
                                var v28 uint16 = genFn0(true)
                                _ = v28
                            }
                        }
                    }
                    v3 -= v3
                }
                var v29 int = 128
                var v30 int8 = 0
                _ = fi26
                _ = v29
                _ = v30
            }
            _ = v25
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    if (!true) {
        v2 += 0.0
        for fi31 := 0; fi31 < 11; fi31++ {
            var v32 uint32 = (v18 / ((v18) | 1))
            v16 = (-(v16))
            var v33 uint = 642
            _ = fi31
            _ = v32
            _ = v33
        }
        var v34 int = 2147483647
        _ = v34
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        fmt.Println(v1, v2)
        v36 := GenS1{F0: true, F1: (v1 % ((v1) | 1)), F2: (v1 % (((v14 | (v1 ^ v1))) | 1))}
        _ = fi35
        _ = v36
    }
    v1 = uint8(v18)
    var v37 string = "a"
    var v38 uint64 = (v3 % (((v3 | 4294967295)) | 1))
    v1 += (^(v14))
    for fi39 := 0; fi39 < 6; fi39++ {
        v40 := GenS1{F0: (true || true), F1: 127, F2: 127}
        _ = fi39
        _ = v40
    }
    v16 -= 1
    var v41 uint32 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v41
}
