package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint8; F2 int }
type GenS4 struct { F0 uint }
func genFn0(p0 bool, p1 int8) int16 {
    for fi4 := 0; fi4 < 3; fi4++ {
        p1 = (p1 >> ((42) & 63))
        var v5 uint = 255
        v6 := []int{fi4, fi4, fi4}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 bool = p0
    v7 = v7
    p1 = 1
    _ = p0
    _ = p1
    _ = v7
    return 496
}
func genFn1(p0 int) string {
    switch p0 {
    case 169:
        fmt.Println(p0)
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := []int{(fi8 / ((4294967295) | 1)), (-(fi8)), (fi8 >> ((4294967295) & 63))}
        _ = fi8
        _ = v9
    }
    fmt.Println(p0)
    p0 -= 127
    var v10 float32 = float32(0.0)
    _ = p0
    _ = v10
    return ""
}

func main() {
    var v1 uint64 = 42
    var v2 uint = 2
    var v3 uint32 = uint32(v1)
    var v11 int32 = -2147483648
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            fmt.Println(v11, v3, v11)
            v14 := GenS1{F0: 532}
            v3 -= (v3 >> ((v2) & 63))
            fmt.Println(fi12, v11)
            _ = fi13
            _ = v14
        }
        var v15 bool = false
        v3 = (^(v3))
        _ = fi12
        _ = v15
    }
    var v16 uint64 = (v1 / ((uint64(v3)) | 1))
    v17 := GenS1{F0: uint16(v3)}
    v18 := GenS1{F0: 1}
    if ((1.5 * 0.0) <= 3.14159) {
        v16 = v1
    } else {
        var v19 uint = (v2 - 814)
        for fi20 := 0; fi20 < 16; fi20++ {
            fmt.Println(v19, v11, v19)
            var v21 int = (-(fi20))
            v19 = v2
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    var v22 uint = 1
    v11 += -2147483648
    var v23 uint8 = 2
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := GenS1{F0: 1}
        v26 := []int{}
        var v27 int32 = (v11 / ((v11) | 1))
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    var v28 uint = uint(v16)
    v29 := []int{42, -690, -2147483648, -337}
    if false {
        var v30 uint32 = (v3 << ((v22) & 63))
        if (!false) {
            fmt.Println(v28, v2, v22)
        }
        switch "\n\t" {
        case "hello":
            v29 = append(v29, 2147483647)
        case "\n\t":
            _ = len(v29)
        case "":
            var v31 int16 = 0
            _ = v31
        default:
            if (false || (!true)) {
                if ((false || false) || true) {
                    if ((v23 | v23) >= (v23 / ((42) | 1))) {
                        var v32 uint = (^(v28))
                        _ = v32
                    }
                    v1 -= v1
                }
                v28 += uint(v30)
            } else {
                var v33 int32 = -2147483648
                _ = v33
            }
        }
        _ = v30
    }
    var v34 float32 = (3.14159 / 3.14159)
    v16 = (v16 ^ 4294967295)
    if true {
        if false {
            for fi35 := 0; fi35 < 10; fi35++ {
                if (v1 >= (^(v16))) {
                    var v36 bool = (((1e10 * 3.14159) < 0.1) || false)
                    _ = v36
                } else {
                    fmt.Println(v23, v16)
                    v16 = (v1 | (v1 * v1))
                }
                switch fi35 {
                case 128:
                    v37 := GenS2{F0: 255}
                    _ = v37
                case 9223372036854775807:
                    if len(v29) > 0 {
                        v29[0] = (fi35 | (fi35 * fi35))
                    }
                }
                v1 += (v1 >> (((v2 << ((v2) & 63))) & 63))
                _ = fi35
            }
            var v38 bool = false
            for fi39 := 0; fi39 < 1; fi39++ {
                v40 := GenS3{F0: -1, F1: 255, F2: fi39}
                _ = fi39
                _ = v40
            }
            _ = v38
        } else {
            var v41 int32 = v11
            v42 := GenS4{F0: v28}
            _ = v41
            _ = v42
        }
    } else {
        v43 := GenS3{F0: 1, F1: (^(v23)), F2: -9223372036854775808}
        _ = v43
    }
    if false {
        var v44 int16 = 42
        _ = v44
    } else {
        var v45 float64 = 0.0
        var v46 uint = (v28 / ((18446744073709551615) | 1))
        _ = v45
        _ = v46
    }
    if false {
        v16 = (v16 * 628)
        v3 = v3
    } else {
        var v47 uint8 = v23
        v1 += 53
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v34
}
