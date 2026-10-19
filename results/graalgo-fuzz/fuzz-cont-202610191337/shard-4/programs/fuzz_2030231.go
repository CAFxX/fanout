package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint, p1 uint8, p2 uint16) int8 {
    switch -702 {
    case 128:
        if (-9223372036854775808 < -1) {
            if (p1 < 42) {
                p2 -= (p2 | p2)
                if true {
                    var v7 int16 = -32768
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 9; fi8++ {
                        v9 := GenS1{F0: p2, F1: true, F2: 1}
                        var v10 bool = (!false)
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    fmt.Println(p1, p2)
                }
            }
            for fi11 := 0; fi11 < 5; fi11++ {
                switch fi11 {
                case 0:
                    p2 = (p2 | (^(p2)))
                }
                var v12 int16 = 256
                var v13 float64 = (((0.1 - (-2.25 * 0.1)) * 3.14159) + (0.0 - (1e10 + (1e10 * 0.0))))
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 uint64 = 644
            _ = v14
        } else {
            var v15 uint64 = 419
            fmt.Println(v15)
            _ = v15
        }
    case 127:
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := []int{237, (fi16 / ((32768) | 1)), (fi16 * 4294967295)}
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 int8 = -1
                var v20 uint64 = 150
                var v21 int32 = -1
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 uint8 = (^(p1))
            _ = fi16
            _ = v17
            _ = v22
        }
    }
    var v23 int8 = 127
    for fi24 := 0; fi24 < 6; fi24++ {
        if (678 >= 1) {
            v25 := []int{}
            v26 := GenS1{F0: (p2 % (((p2 % ((p2) | 1))) | 1)), F1: (18446744073709551615 <= 0), F2: 939}
            v25 = append(v25, fi24)
            _ = v25
            _ = v26
        } else {
            v27 := GenS2{F0: ("hello" + ("" + "a")), F1: (fi24 << ((255) & 63))}
            v28 := GenS3{F0: 2, F1: 634}
            var v29 uint16 = 65535
            _ = v27
            _ = v28
            _ = v29
        }
        p2 = 2
        fmt.Println(p2, p2, v23)
        _ = fi24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return (v23 - 92)
}

func main() {
    var v1 int = 4294967295
    var v2 bool = true
    var v3 float32 = 0.0
    var v4 bool = (!(-2.25 == (((1e10 + 0.1) * (0.0 / -2.25)) + 3.14159)))
    v5 := make(map[string]bool)
    var v6 uint32 = 2
    v30 := GenS4{F0: 1, F1: v1}
    var v31 uint8 = 127
    if (v2 && false) {
        v32 := GenS3{F0: 18446744073709551615, F1: 0}
        for fi33 := 0; fi33 < 20; fi33++ {
            v4 = true
            v34 := []int{}
            fmt.Println(fi33, v2, fi33)
            _ = fi33
            _ = v34
        }
        var v35 uint16 = 323
        _ = v32
        _ = v35
    }
    var v36 float64 = 1e10
    var v37 float32 = (-2.25 * ((0.0 * v3) + float32(((v36 * v36) - (v36 / (v36 * (float64((v3 * (float32(((float64(v3) - (v36 / v36)) + (v36 - (v36 * (0.1 - 1.5))))) / v3))) * v36)))))))
    v31 += 255
    fmt.Println(v4, v4)
    var v38 uint8 = (^(v31))
    fmt.Println(v2)
    v6 = v6
    v39 := []int{(v1 | (v1 | (v1 + (v1 * v1)))), (v1 >> ((2) & 63)), v1, (v1 % (((v1 + (^(v1)))) | 1))}
    fmt.Println(v3)
    if true {
        v40 := GenS5{F0: 140, F1: v31}
        _ = v40
    }
    var v41 uint16 = 255
    fmt.Println(v38, v3)
    var v42 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
