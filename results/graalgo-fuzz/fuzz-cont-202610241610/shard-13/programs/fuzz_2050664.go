package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 bool }
func genFn0() uint {
    v4 := GenS1{F0: 471, F1: 9223372036854775807, F2: true}
    if true {
        v5 := []int{262, 182}
        v6 := GenS1{F0: 2, F1: 42, F2: false}
        var v7 int16 = 32767
        _ = v5
        _ = v6
        _ = v7
    }
    _ = v4
    return 4294967295
}

func main() {
    var v1 bool = (670 != 0)
    var v2 int8 = 100
    v3 := []uint32{220, 42, 360}
    v1 = false
    if (v1 || v1) {
        var v8 uint = genFn0()
        var v9 uint32 = 0
        var v10 float64 = float64(-2.25)
        _ = v10
        _ = v8
        _ = v9
    }
    switch "\n\t" {
    case "hello":
        v2 -= -128
    default:
        v11 := GenS1{F0: 42, F1: -9223372036854775808, F2: (!v1)}
        _ = v11
    }
    fmt.Println(v1, v1)
    v2 -= v2
    switch 9223372036854775807 {
    case 32767:
        if (!v1) {
            v12 := GenS1{F0: 167, F1: 42, F2: (v1 && (799 != 2))}
            _ = len(v3)
            _ = v12
        }
    }
    fmt.Println(v2, v2)
    var v13 int16 = 15
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            v13 -= v13
            v13 += (v13 & (-(v13)))
            _ = fi15
        }
        v3 = append(v3, 1)
        if (!v1) {
            v16 := GenS1{F0: 913, F1: 42, F2: (1 != 42)}
            _ = v16
        }
        _ = fi14
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        if (0.0 < 1.5) {
            if true {
                var v18 int32 = 0
                v3 = append(v3, 263)
                for fi19 := 0; fi19 < 13; fi19++ {
                    v20 := GenS1{F0: 4294967295, F1: 245, F2: (!(65535 <= 0))}
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
        }
        var v21 int32 = -1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi17
        _ = v21
    }
    var v22 int = 748
    var v23 int64 = -1
    switch (v22 % (((v22 >> ((255) & 63))) | 1)) {
    case 32768:
        if (!(42 < 0)) {
            fmt.Println(v1, v13, v22)
            v1 = (!true)
        } else {
            fmt.Println(v23, v23)
            v13 = 1
            v24 := GenS1{F0: 18446744073709551615, F1: (v23 >> ((genFn0()) & 63)), F2: true}
            _ = v24
        }
    case 83:
        for fi25 := 0; fi25 < 17; fi25++ {
            if v1 {
                switch (fi25 / ((9223372036854775807) | 1)) {
                case 28:
                    var v26 uint8 = 0
                    _ = v26
                case 1:
                    for fi27 := 0; fi27 < 5; fi27++ {
                        var v28 uint16 = 0
                        _ = fi27
                        _ = v28
                    }
                }
                v1 = ((v2 % (((v2 * v2)) | 1)) == v2)
            }
            _ = fi25
        }
    }
    v13 += v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
