package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int32 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{127, 2147483647, 28, 32767}
    var v3 int = 4294967295
    var v4 float32 = (1.5 * 0.0)
    v5 := make(map[string]bool)
    var v6 int8 = -128
    switch (v3 * (^(v3))) {
    case -1:
        v6 += v6
    case -2147483648:
        v7 := GenS1{F0: (v6 ^ (v6 >> ((976) & 63))), F1: 655, F2: 741}
        _ = v7
    case 2147483647:
        var v8 float32 = (v4 + 0.1)
        _ = v8
    }
    v9 := GenS1{F0: (v6 & (^(v6))), F1: 42, F2: 2147483647}
    var v10 bool = (!(int32(v3) == 2147483647))
    v6 = (v6 & (v6 % (((v6 ^ v6)) | 1)))
    if v10 {
        if v10 {
            var v11 int32 = 2147483647
            fmt.Println(v10, v3)
            _ = v11
        }
        var v12 uint16 = 65535
        _ = v12
    } else {
        v2 = append(v2, 42)
        if true {
            var v13 uint64 = 348
            _ = v13
        } else {
            if len(v2) > 0 {
                v2[0] = (v3 / (((-(v3))) | 1))
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 float32 = 3.14159
                v10 = (1 <= 1)
                v16 := GenS2{F0: (42 == 230), F1: (v6 | 100)}
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 float64 = 1e10
            _ = v17
        }
        v6 -= v6
    }
    v3 = 9223372036854775807
    v18 := v5["a"]
    fmt.Println(v18)
    v18 = (!(0 >= 400))
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 uint16 = 42
        _ = fi19
        _ = v20
    }
    _ = len(v1)
    v6 = v6
    v4 -= v4
    v5["\n\t"] = (32767 >= 255)
    v6 += v6
    var v21 uint32 = 1
    if v18 {
        switch (v3 << ((2) & 63)) {
        case -9223372036854775808:
            delete(v5, "a")
        case 128:
            if (v6 <= v6) {
                fmt.Println(v18)
            }
        case 2147483647:
            if (-1 != -9223372036854775808) {
                if v18 {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v23 := GenS3{F0: 761}
                        v24 := GenS3{F0: 2}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                } else {
                    for fi25 := 0; fi25 < 12; fi25++ {
                        v26 := GenS1{F0: 2, F1: 18446744073709551615, F2: 42}
                        _ = fi25
                        _ = v26
                    }
                }
            } else {
                var v27 string = "\n\t"
                _ = v27
            }
        default:
            v21 -= 4294967295
        }
        v21 -= v21
        v28 := GenS4{F0: v6, F1: "a"}
        _ = v28
    }
    var v29 float64 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
