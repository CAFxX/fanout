package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint) int8 {
    switch -28 {
    case 32767:
        for fi4 := 0; fi4 < 4; fi4++ {
            var v5 uint32 = 1
            if false {
                var v6 uint8 = 145
                var v7 uint64 = 615
                _ = v6
                _ = v7
            } else {
                var v8 uint = p0
                v5 -= (v5 * v5)
                _ = v8
            }
            _ = fi4
            _ = v5
        }
    }
    fmt.Println(p0)
    v9 := []int{-988}
    p0 = (p0 ^ (p0 | p0))
    if false {
        v10 := GenS1{F0: 1, F1: 1, F2: 1}
        var v11 uint16 = 255
        _ = v10
        _ = v11
    } else {
        v9 = append(v9, -930)
        v12 := GenS2{F0: 2}
        _ = v12
    }
    _ = p0
    _ = v9
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 string = "\n\t"
    var v3 int = 4294967295
    switch v3 {
    case 144:
        if true {
            var v13 uint64 = 249
            v14 := []int{(^(v3)), v3, v3, v3}
            fmt.Println(v13)
            _ = v13
            _ = v14
        }
    default:
        v2 = ("\n\t" + v2)
    }
    if false {
        v15 := []int{(v3 + v3)}
        _ = v15
    } else {
        var v16 bool = false
        fmt.Println(v3, v1, v1)
        _ = v16
    }
    v3 = v3
    fmt.Println(v3)
    var v17 int32 = 1
    for fi18 := 0; fi18 < 18; fi18++ {
        v17 -= v17
        _ = fi18
    }
    if (597 == 42) {
        v19 := GenS1{F0: -487, F1: -9223372036854775808, F2: (v3 + int(v17))}
        var v20 uint8 = 128
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 float32 = 1e10
            fmt.Println(v2)
            v23 := GenS3{F0: true, F1: 641}
            _ = fi21
            _ = v22
            _ = v23
        }
        if (true && (2147483647 <= 42)) {
            switch "\n\t" {
            case "a":
                v24 := []int{}
                _ = v24
            default:
                v2 = v2
            }
            var v25 uint8 = 255
            v26 := GenS1{F0: v17, F1: (v1 ^ (v1 << ((1) & 63))), F2: (-(v3))}
            _ = v25
            _ = v26
        }
        var v27 uint16 = 42
        _ = v27
    }
    v17 = (v17 & (v17 % (((v17 - 547)) | 1)))
    v2 = v2
    if (0.0 < 1e10) {
        v28 := GenS2{F0: 0}
        if (((v1 >> ((18446744073709551615) & 63)) != -155) && true) {
            var v29 float64 = ((float64(1e10) / 1.0) + (1e10 + float64(1.5)))
            _ = v29
        } else {
            switch 32767 {
            case 9223372036854775807:
                v30 := []int{}
                _ = v30
            case -1:
                v3 += -9223372036854775808
            }
        }
        _ = v28
    }
    switch "" {
    case "":
        fmt.Println(v17, v17)
    case "foo bar":
        var v31 float64 = float64(3.14159)
        _ = v31
    }
    v32 := []int{}
    v1 -= (v1 << ((562) & 63))
    fmt.Println(v2)
    if (1 != 850) {
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 float32 = (-2.25 + 3.14159)
            var v35 float32 = v34
            for fi36 := 0; fi36 < 14; fi36++ {
                v37 := GenS3{F0: true, F1: 506}
                v38 := GenS1{F0: (v17 | v17), F1: v1, F2: v3}
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = fi33
            _ = v34
            _ = v35
        }
        if true {
            v3 = (v3 * (^(v3)))
            fmt.Println(v1)
        } else {
            var v39 float64 = float64(1.5)
            var v40 bool = true
            _ = v39
            _ = v40
        }
    } else {
        v17 += 42
        for fi41 := 0; fi41 < 1; fi41++ {
            for fi42 := 0; fi42 < 8; fi42++ {
                if true {
                    var v43 bool = true
                    _ = v43
                } else {
                    v1 = (v1 * v1)
                    v44 := GenS1{F0: (v17 ^ 909), F1: (-(v1)), F2: (v3 - 9223372036854775807)}
                    _ = v44
                }
                v1 += (v1 ^ (v1 / (((v1 % ((v1) | 1))) | 1)))
                v32 = append(v32, (^(v3)))
                var v45 float64 = ((0.1 * 1.5) - 1.5)
                _ = fi42
                _ = v45
            }
            _ = fi41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
}
