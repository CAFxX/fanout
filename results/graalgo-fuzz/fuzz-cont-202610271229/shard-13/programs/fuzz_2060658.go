package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint64; F1 bool; F2 bool }
func genFn0() bool {
    for fi4 := 0; fi4 < 15; fi4++ {
        var v5 bool = (0.1 > 1e10)
        if (0 > 243) {
            if ((-91 != (fi4 ^ fi4)) || (float64(-2.25) == 1e10)) {
                for fi6 := 0; fi6 < 15; fi6++ {
                    switch -96 {
                    case 32768:
                        v5 = (((fi6 * (fi6 % ((-308) | 1))) < (fi4 / ((fi4) | 1))) || (v5 || v5))
                    default:
                        v5 = false
                    }
                    if (!(false && v5)) {
                        v7 := GenS1{F0: fi4, F1: 4294967295, F2: 1000}
                        switch "" {
                        case "foo bar":
                            var v8 int64 = 862
                            _ = v8
                        }
                        _ = v7
                    } else {
                        switch (fi6 + (fi4 ^ (fi4 / ((fi4) | 1)))) {
                        case 32768:
                            var v9 float64 = 0.1
                            _ = v9
                        default:
                            if v5 {
                                v5 = v5
                                var v10 float64 = -2.25
                                _ = v10
                            }
                        }
                        v5 = v5
                    }
                    v11 := []int{-9223372036854775808}
                    _ = fi6
                    _ = v11
                }
                v12 := []int{}
                _ = v12
            }
            for fi13 := 0; fi13 < 16; fi13++ {
                v5 = (-2.25 >= 1.5)
                v5 = v5
                _ = fi13
            }
            if (42 < 42) {
                fmt.Println(fi4)
                for fi14 := 0; fi14 < 16; fi14++ {
                    v15 := []int{fi4, fi14, fi4, 1}
                    _ = len(v15)
                    v16 := GenS1{F0: fi14, F1: 4294967295, F2: 841}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            } else {
                fmt.Println(fi4, v5)
                fmt.Println(fi4, fi4)
            }
        }
        v5 = (456 != 734)
        _ = fi4
        _ = v5
    }
    var v17 int16 = 940
    v18 := GenS1{F0: 2147483647, F1: 1, F2: (-(v17))}
    switch "\n\t" {
    case "":
        fmt.Println(v17)
    default:
        v17 += v17
    }
    switch 0 {
    case 826:
        for fi19 := 0; fi19 < 5; fi19++ {
            v17 = -1
            _ = fi19
        }
    case 2147483647:
        if (!(!false)) {
            v20 := GenS2{F0: 2, F1: false, F2: ((!false) && true)}
            v21 := GenS1{F0: int(v17), F1: 42, F2: (v17 * v17)}
            var v22 int8 = 100
            _ = v20
            _ = v21
            _ = v22
        } else {
            v23 := GenS2{F0: 0, F1: (false || false), F2: true}
            fmt.Println(v17)
            var v24 int32 = -924
            _ = v23
            _ = v24
        }
    default:
        var v25 uint = 643
        _ = v25
    }
    _ = v17
    _ = v18
    return (42 == 922)
}
func genFn1() int8 {
    var v26 uint8 = 2
    var v27 int32 = 889
    v26 -= v26
    _ = v26
    _ = v27
    return -120
}

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = 637
    var v3 int = 32768
    var v28 uint16 = 1
    for fi29 := 0; fi29 < 7; fi29++ {
        switch (fi29 | v3) {
        case 32767:
            v1 = v1
        case -462:
            var v30 int32 = -143
            _ = v30
        }
        _ = fi29
    }
    var v31 uint64 = 654
    v32 := []int{(v3 ^ v3), v3}
    v33 := GenS1{F0: v3, F1: 4294967295, F2: 42}
    fmt.Println(v2, v3, v28)
    v2 = int32(v28)
    v3 -= (-(v3))
    v34 := GenS2{F0: 4294967295, F1: (-9223372036854775808 > 42), F2: (0.0 != (1.5 * 1e10))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
