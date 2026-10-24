package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint }
type GenS2 struct { F0 int32; F1 uint; F2 int }
func genFn0(p0 int, p1 int) uint8 {
    v7 := []int{p1, p0, -21, p1}
    if false {
        p1 = -9223372036854775808
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            if false {
                var v9 uint = 468
                fmt.Println(p0, v9)
                p1 += fi8
                _ = v9
            }
            if len(v7) > 0 {
                v7[0] = (p1 / (((p1 ^ -9223372036854775808)) | 1))
            }
            var v10 uint8 = 127
            _ = fi8
            _ = v10
        }
        if ((255 < 0) || (!false)) {
            if len(v7) > 0 {
                v7[0] = -2147483648
            }
            if (0.1 < (0.0 - -2.25)) {
                fmt.Println(p0, p1)
            }
        }
        var v11 bool = ((("a" + ("" + (("" + "hello") + ""))) + "\n\t") >= ("a" + (("a" + "hello") + ("a" + "\n\t"))))
        _ = v11
    }
    if len(v7) > 0 {
        v7[0] = (p1 ^ -1)
    }
    if false {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            if false {
                fmt.Println(p1, p1, fi12)
            }
            var v13 int8 = -20
            fmt.Println(p1, p1)
            _ = fi12
            _ = v13
        }
    } else {
        switch "hello" {
        case "hello":
            switch p0 {
            case 42:
                var v14 int8 = 1
                _ = v14
            case -2147483648:
                for fi15 := 0; fi15 < 4; fi15++ {
                    v16 := GenS1{F0: 818, F1: 4294967295, F2: 246}
                    _ = fi15
                    _ = v16
                }
            default:
                var v17 uint = 4294967295
                _ = v17
            }
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if (false && true) {
            for fi18 := 0; fi18 < 20; fi18++ {
                for fi19 := 0; fi19 < 10; fi19++ {
                    var v20 bool = (128 == 128)
                    var v21 float32 = 1.5
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = fi18
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 uint = 42
                var v24 int16 = 1
                _ = len(v7)
                _ = fi22
                _ = v23
                _ = v24
            }
            var v25 uint = 1
            _ = v25
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 127
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]string)
    var v3 bool = (1 >= -9223372036854775808)
    var v4 int32 = 0
    v5 := []uint32{2147483647}
    var v6 int16 = -32768
    for fi26 := 0; fi26 < 18; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            if v3 {
                v28 := GenS1{F0: 42, F1: uint(v1), F2: 42}
                v1 = v1
                v6 = (^(v6))
                _ = v28
            } else {
                v29 := GenS2{F0: int32(fi26), F1: 65535, F2: (fi26 >> ((4294967295) & 63))}
                v6 = v6
                var v30 float32 = (((((-2.25 - (((0.0 / ((-2.25 - 1.5) + (-2.25 * 3.14159))) + ((0.1 - 1.5) - (0.1 + ((-2.25 - 0.1) * (1e10 * ((((3.14159 * 1e10) - (1.5 * ((((-2.25 * ((3.14159 - 0.1) / 3.14159)) * float32((-2.25 * 1e10))) * float32(0.1)) - (-2.25 / 1.5)))) * 0.0) * 3.14159)))))) - 0.1)) / float32(1.5)) - -2.25) + 3.14159) - 0.0)
                _ = v29
                _ = v30
            }
            _ = fi27
        }
        var v31 int32 = (v4 ^ v4)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi26
        _ = v31
    }
    fmt.Println(v1, v6)
    v1 = v1
    v6 -= (^(v6))
    v4 += v4
    if (836 == 42) {
        if len(v5) > 0 {
            v5[0] = 65535
        }
        v5 = append(v5, 2147483647)
        fmt.Println(v1, v3)
    } else {
        var v32 int8 = 1
        fmt.Println(v3, v6, v4)
        v6 = 1000
        _ = v32
    }
    fmt.Println(v4)
    var v33 float64 = 1e10
    fmt.Println(v1)
    v33 += v33
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
