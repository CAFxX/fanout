package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint64; F1 string }
func genFn0(p0 uint32) uint32 {
    fmt.Println(p0)
    if (0.1 > float64(1.5)) {
        var v6 uint32 = (p0 | p0)
        v7 := []int{1, 4294967295, -9223372036854775808, 32768}
        _ = v6
        _ = v7
    } else {
        var v8 int16 = -1
        _ = v8
    }
    fmt.Println(p0)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint8 = 42
        v10 -= 127
        p0 += p0
        _ = fi9
        _ = v10
    }
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    v3 := []uint32{42, 0, 42}
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]bool)
    var v11 uint = 42
    v11 -= v11
    var v12 float32 = 0.0
    v13 := GenS1{F0: (^(v11)), F1: false}
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint8 = 255
        switch (fi14 % ((fi14) | 1)) {
        case -9223372036854775808:
            if false {
                v15 = v15
                v16 := GenS1{F0: 560, F1: ((1 < (v11 & (v11 ^ (v11 * (v11 ^ (v11 + 42)))))) && (1.5 > 1.5))}
                var v17 uint32 = 0
                _ = v16
                _ = v17
            } else {
                var v18 int = -255
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v15)
                _ = v18
            }
        case -2147483648:
            if (true || ((4294967295 == 1) || false)) {
                v12 = (v12 + 1e10)
                var v19 bool = (!true)
                v11 = v11
                _ = v19
            } else {
                v20 := GenS1{F0: (v11 << (((^(v11))) & 63)), F1: (160 == 2)}
                _ = v20
            }
        }
        var v21 int8 = -128
        _ = fi14
        _ = v15
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := GenS2{F0: 549, F1: ("\n\t" < "a")}
        v12 = ((0.1 * v12) - 1e10)
        _ = fi22
        _ = v23
    }
    v12 = (v12 - v12)
    for fi24 := 0; fi24 < 8; fi24++ {
        v12 = v12
        var v25 int16 = -1
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v12, v11)
    fmt.Println(v11)
    switch ("hello" + "\n\t") {
    case "":
        fmt.Println(v12, v11)
    default:
        if (131 > 1) {
            var v26 bool = false
            var v27 uint16 = 94
            fmt.Println(v12)
            _ = v26
            _ = v27
        }
    }
    if (!false) {
        for fi28 := 0; fi28 < 7; fi28++ {
            for fi29 := 0; fi29 < 20; fi29++ {
                fmt.Println(v11)
                v30 := GenS3{F0: "foo bar"}
                fmt.Println(v12, v11, v12)
                _ = fi29
                _ = v30
            }
            var v31 bool = false
            if (true && (-2.25 <= (((-2.25 / 1.5) * 0.0) * (((float64(v12) + 0.1) / (1.5 * -2.25)) * 1.5)))) {
                switch "hello" {
                case "foo bar":
                    var v32 int = fi28
                    _ = v32
                case "a":
                    v33 := GenS1{F0: (v11 / ((v11) | 1)), F1: (386 == 18446744073709551615)}
                    _ = v33
                default:
                    var v34 uint32 = 1
                    _ = v34
                }
                var v35 float64 = (0.1 + 0.1)
                _ = v35
            }
            _ = fi28
            _ = v31
        }
        v11 = 475
        v36 := GenS3{F0: "a"}
        _ = v36
    } else {
        if true {
            for fi37 := 0; fi37 < 4; fi37++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v38 := v5[65535]
                _ = fi37
                _ = v38
            }
            fmt.Println(v12, v11)
        } else {
            if false {
                var v39 float64 = (1.5 * 0.1)
                for fi40 := 0; fi40 < 13; fi40++ {
                    v41 := GenS4{F0: 32767, F1: (v11 | (v11 & (v11 >> ((v11) & 63))))}
                    _ = fi40
                    _ = v41
                }
                _ = v39
            }
        }
        var v42 uint64 = 357
        _ = v42
    }
    v43 := GenS5{F0: 18446744073709551615, F1: "\n\t"}
    for fi44 := 0; fi44 < 14; fi44++ {
        delete(v1, 2147483648)
        _ = fi44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
