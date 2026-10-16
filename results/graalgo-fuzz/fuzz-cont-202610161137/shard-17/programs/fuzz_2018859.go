package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 int64, p1 uint64, p2 int) uint64 {
    fmt.Println(p2, p0, p1)
    switch p2 {
    case 1:
        p1 = (p1 + p1)
    case -9223372036854775808:
        switch "\n\t" {
        case "hello":
            fmt.Println(p0)
        default:
            var v4 float64 = 1.5
            _ = v4
        }
    case 42:
        if true {
            v5 := GenS1{F0: 32767, F1: -609}
            v6 := []int{(p2 / ((p2) | 1)), 331, p2}
            var v7 int8 = 1
            _ = v5
            _ = v6
            _ = v7
        } else {
            p1 = 2
            if true {
                v8 := GenS1{F0: 32767, F1: -142}
                _ = v8
            } else {
                if (0.0 > (1.5 * -2.25)) {
                    var v9 uint = 809
                    switch (^(p2)) {
                    case 32768:
                        v10 := GenS1{F0: 1000, F1: 2147483647}
                        _ = v10
                    case -9223372036854775808:
                        var v11 float32 = 0.0
                        _ = v11
                    }
                    _ = v9
                }
                v12 := GenS1{F0: 0, F1: 731}
                _ = v12
            }
            p0 += p0
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 bool = ((305 < -631) && false)
    var v2 uint64 = 4294967295
    var v3 int = -2147483648
    fmt.Println(v2, v2)
    v13 := []int{0, v3, v3}
    v3 = (v3 / ((128) | 1))
    for fi14 := 0; fi14 < 7; fi14++ {
        fmt.Println(v1, v3, v2)
        _ = fi14
    }
    var v15 uint32 = 823
    switch v3 {
    case -327:
        if true {
            v13 = append(v13, v3)
            _ = len(v13)
        } else {
            v2 = v2
            v13 = append(v13, (v3 * v3))
        }
    default:
        if (!v1) {
            switch (-(v3)) {
            case -9223372036854775808:
                var v16 int32 = -1
                _ = v16
            case 65:
                v3 -= (v3 >> ((18446744073709551615) & 63))
            case -631:
                if (962 <= v2) {
                    if len(v13) > 0 {
                        v13[0] = (v3 >> ((18446744073709551615) & 63))
                    }
                } else {
                    var v17 bool = v1
                    if (-9223372036854775808 == -1) {
                        v17 = (42 > 423)
                        v13 = append(v13, (v3 / (((v3 % (((^(v3))) | 1))) | 1)))
                    }
                    var v18 float64 = (float64((float32((3.14159 / float64(-2.25))) * ((1.5 * float32(1e10)) + 0.1))) * float64(0.0))
                    _ = v17
                    _ = v18
                }
            }
            fmt.Println(v15, v1, v15)
            for fi19 := 0; fi19 < 13; fi19++ {
                if len(v13) > 0 {
                    v13[0] = (^(fi19))
                }
                _ = fi19
            }
        } else {
            var v20 uint64 = v2
            v15 = v15
            for fi21 := 0; fi21 < 4; fi21++ {
                v1 = (!((float64(3.14159) + 1.5) == -2.25))
                for fi22 := 0; fi22 < 8; fi22++ {
                    var v23 int = fi21
                    _ = fi22
                    _ = v23
                }
                v13 = append(v13, (v3 & v3))
                _ = fi21
            }
            _ = v20
        }
    }
    v24 := GenS1{F0: 26, F1: 1}
    v25 := GenS2{F0: 65535}
    v26 := GenS1{F0: 447, F1: -2147483648}
    v1 = (!false)
    for fi27 := 0; fi27 < 20; fi27++ {
        _ = len(v13)
        _ = fi27
    }
    v3 += v3
    v1 = (true && v1)
    var v28 int16 = 374
    v13 = append(v13, (v3 % (((v3 / (((v3 & 0)) | 1))) | 1)))
    if len(v13) > 0 {
        v13[0] = 856
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
}
