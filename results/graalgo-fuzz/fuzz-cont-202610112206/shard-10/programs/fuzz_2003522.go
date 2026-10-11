package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16; F1 int8; F2 int32 }
func genFn0(p0 uint64) int32 {
    var v4 string = "foo bar"
    v4 = v4
    switch 42 {
    case 2147483647:
        if (9223372036854775807 >= -9223372036854775808) {
            if true {
                p0 = 592
                var v5 uint64 = p0
                v5 -= (p0 >> ((39) & 63))
                _ = v5
            } else {
                var v6 int64 = -1
                _ = v6
            }
            for fi7 := 0; fi7 < 13; fi7++ {
                if false {
                    switch fi7 {
                    case -697:
                        v8 := []int{fi7, fi7, (fi7 >> ((674) & 63))}
                        _ = v8
                    }
                    var v9 int32 = -1
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 4; fi10++ {
                        fmt.Println(v4)
                        _ = fi10
                    }
                    if true {
                        fmt.Println(v4, p0, v4)
                        var v11 uint8 = 42
                        _ = v11
                    }
                }
                _ = fi7
            }
            if (!false) {
                v12 := []int{32768, 42}
                _ = v12
            } else {
                var v13 uint8 = 75
                _ = v13
            }
        }
    case 23:
        v4 = v4
    }
    var v14 int = -705
    var v15 int64 = 0
    _ = p0
    _ = v14
    _ = v15
    _ = v4
    return 536
}

func main() {
    var v1 int32 = 1
    var v2 int8 = 25
    var v3 string = "\n\t"
    v2 -= 0
    var v16 int16 = 1000
    var v17 float64 = -2.25
    var v18 int8 = 1
    switch 2147483647 {
    case -872:
        v19 := []int{42, 418, -9223372036854775808, 953}
        _ = v19
    case -115:
        for fi20 := 0; fi20 < 2; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                fmt.Println(v3)
                switch v3 {
                case "hello":
                    fmt.Println(fi21)
                case "foo bar":
                    v22 := []int{fi21, (-(fi21))}
                    _ = v22
                case "":
                    v16 -= (^(v16))
                }
                _ = fi21
            }
            for fi23 := 0; fi23 < 11; fi23++ {
                v2 -= -78
                var v24 float32 = (0.0 / 1e10)
                _ = fi23
                _ = v24
            }
            _ = fi20
        }
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        v16 -= 1000
        fmt.Println(v2, v2, v17)
        _ = fi25
    }
    fmt.Println(v2, v17, v1)
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := []int{(fi26 * (fi26 / ((2147483647) | 1))), -626, (fi26 ^ 128)}
        var v28 uint16 = 936
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: (1 >= 2), F1: v16}
    if true {
        v30 := []int{2147483647, int(v18), 499}
        var v31 bool = (v16 < v16)
        _ = v30
        _ = v31
    }
    v32 := []int{761, -9223372036854775808}
    _ = len(v32)
    fmt.Println(v18, v3)
    var v33 uint32 = 840
    switch 1 {
    case 32767:
        var v34 float64 = v17
        _ = v34
    case -2147483648:
        v35 := GenS2{F0: 255, F1: (v18 | v2), F2: genFn0(809)}
        _ = v35
    default:
        v36 := GenS1{F0: false, F1: 32767}
        _ = v36
    }
    var v37 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v37
}
