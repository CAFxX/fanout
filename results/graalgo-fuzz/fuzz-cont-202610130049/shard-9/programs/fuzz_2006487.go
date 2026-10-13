package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint8; F1 bool; F2 int16 }
func genFn0() int64 {
    switch "\n\t" {
    case "foo bar":
        var v7 float32 = 0.1
        _ = v7
    case "\n\t":
        var v8 float32 = (((0.0 * (-2.25 * 0.1)) * 3.14159) / 1.0)
        _ = v8
    }
    var v9 int = 9223372036854775807
    if (316 >= 255) {
        var v10 int8 = 100
        _ = v10
    }
    v11 := []int{(v9 / ((v9) | 1)), (-(v9)), v9}
    _ = v11
    _ = v9
    return 0
}
func genFn1(p0 int64, p1 uint) int {
    if (42 <= 2147483647) {
        fmt.Println(p1)
        var v12 int16 = 0
        v13 := []int{32767, 914, 121, 9223372036854775807}
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int16 = 1
            _ = fi14
            _ = v15
        }
        fmt.Println(p0, p1)
    }
    if (!(-260 < -209)) {
        p1 += (p1 ^ p1)
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    var v1 int = 2147483647
    v2 := []string{}
    var v3 int64 = 0
    var v4 uint16 = 1
    v5 := make(map[uint32]bool)
    var v6 uint32 = 440
    v2 = append(v2, "foo bar")
    if true {
        var v16 float32 = (((-2.25 - float32(float64(1e10))) - 1.5) - ((1.5 * 1e10) / (float32(((1e10 / 1e10) + 0.1)) * float32(3.14159))))
        _ = len(v5)
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        if false {
            v18 := v5[2147483648]
            _ = v18
        } else {
            var v19 bool = false
            var v20 int = (^(fi17))
            v19 = false
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: 0}
        _ = fi17
        _ = v21
    }
    if (true || false) {
        for fi22 := 0; fi22 < 4; fi22++ {
            v5[(v6 % ((v6) | 1))] = false
            for fi23 := 0; fi23 < 12; fi23++ {
                v24 := GenS1{F0: 455}
                v3 -= (v3 * (v3 << ((733) & 63)))
                fmt.Println(fi23)
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    } else {
        v25 := GenS1{F0: -1}
        var v26 uint16 = 65535
        v5[(v6 + (^(v6)))] = false
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: 2147483647}
    v1 += (v1 / ((-933) | 1))
    var v28 uint32 = 956
    for fi29 := 0; fi29 < 3; fi29++ {
        delete(v5, (v6 >> ((255) & 63)))
        _ = fi29
    }
    delete(v5, (v28 * v28))
    switch "hello" {
    case "":
        v30 := GenS2{F0: 4294967295, F1: -798, F2: 2147483647}
        _ = v30
    case "foo bar":
        switch genFn1(genFn0(), uint(v4)) {
        case -9223372036854775808:
            fmt.Println(v28, v1)
        }
    case "\n\t":
        var v31 int8 = 43
        _ = v31
    default:
        v32 := GenS2{F0: 0, F1: 0, F2: (v6 * v6)}
        _ = v32
    }
    v33 := GenS1{F0: -319}
    if ((!(!false)) && true) {
        var v34 uint64 = 4294967295
        var v35 int16 = 0
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        _ = v34
        _ = v35
    }
    var v36 uint32 = 2147483647
    delete(v5, v28)
    v1 = v1
    _ = len(v5)
    for fi37 := 0; fi37 < 1; fi37++ {
        fmt.Println(v36, v4)
        v38 := GenS3{F0: 178, F1: true, F2: 32767}
        _ = fi37
        _ = v38
    }
    if ((true && true) && false) {
        switch (-(v1)) {
        case -2147483648:
            for fi39 := 0; fi39 < 17; fi39++ {
                fmt.Println(v36, v1)
                switch (("\n\t" + "a") + "a") {
                case "":
                    if (!false) {
                        var v40 bool = (0.1 < (-2.25 * (3.14159 + (-2.25 - 1.5))))
                        v3 -= v3
                        _ = v40
                    } else {
                        _ = len(v2)
                        v1 = fi39
                    }
                default:
                    v41 := GenS1{F0: 42}
                    _ = v41
                }
                var v42 uint64 = 42
                _ = fi39
                _ = v42
            }
        case -982:
            v28 = 2147483648
        case 624:
            var v43 string = ""
            _ = v43
        default:
            fmt.Println(v4, v6)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
