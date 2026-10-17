package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint }

func main() {
    var v1 int32 = -819
    v2 := make(map[uint32]string)
    var v3 uint64 = 794
    var v4 float64 = (0.1 - -2.25)
    switch 695 {
    case 4294967295:
        v4 = ((float64((3.14159 - 0.0)) / v4) / ((3.14159 / (((((float64(0.1) + v4) * 1.5) + v4) - float64(1e10)) + -2.25)) / -2.25))
    case 0:
        if false {
            v5 := GenS1{F0: 477, F1: int64(v3), F2: 4294967295}
            _ = v5
        } else {
            v6 := GenS1{F0: -799, F1: 0, F2: 65535}
            var v7 float32 = (1e10 / 3.14159)
            for fi8 := 0; fi8 < 4; fi8++ {
                v2[65535] = ("hello" + "a")
                _ = fi8
            }
            _ = v6
            _ = v7
        }
    case 42:
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := []int{2147483647, 24, fi9, -1}
            _ = fi9
            _ = v10
        }
    }
    switch -9223372036854775808 {
    case 32767:
        _ = len(v2)
    case 32768:
        v11 := GenS1{F0: 32768, F1: 275, F2: 4294967295}
        _ = v11
    case 9223372036854775807:
        v2[0] = "\n\t"
    default:
        v3 = 767
    }
    if false {
        v4 += 1.5
        v4 -= float64(0.1)
        if (1.5 >= -2.25) {
            v1 = v1
        }
    } else {
        v1 += (v1 % ((658) | 1))
        for fi12 := 0; fi12 < 14; fi12++ {
            fmt.Println(fi12)
            _ = fi12
        }
        var v13 uint16 = 1
        _ = v13
    }
    var v14 int = 32767
    v15 := GenS1{F0: v14, F1: -94, F2: 4294967295}
    var v16 uint32 = 42
    v14 = (v14 >> ((979) & 63))
    for fi17 := 0; fi17 < 5; fi17++ {
        if false {
            v18 := GenS1{F0: fi17, F1: 853, F2: 2}
            if true {
                var v19 string = "a"
                _ = v19
            } else {
                if true {
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v21 := v2[uint32(v3)]
                        _ = fi20
                        _ = v21
                    }
                    var v22 float64 = (v4 - 0.1)
                    _ = v22
                }
                v1 = v1
            }
            _ = v18
        }
        _ = fi17
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        v3 += v3
        v24 := GenS1{F0: v14, F1: 334, F2: 311}
        v25 := v2[(v16 / ((v16) | 1))]
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        if true {
            v16 -= (v16 << ((18446744073709551615) & 63))
        }
        var v27 float32 = float32(v4)
        _ = fi26
        _ = v27
    }
    if ((0.0 / 1.5) == (0.1 - 0.0)) {
        v4 += v4
    }
    var v28 float64 = v4
    if (!true) {
        fmt.Println(v28)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
