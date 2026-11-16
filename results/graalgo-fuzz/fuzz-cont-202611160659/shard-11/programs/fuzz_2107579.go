package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint; F2 int32 }
func genFn0() uint32 {
    if false {
        var v4 uint64 = 2
        _ = v4
    } else {
        var v5 bool = ((65535 != 42) || (!(-794 >= 9223372036854775807)))
        var v6 int64 = 0
        fmt.Println(v6)
        _ = v5
        _ = v6
    }
    var v7 bool = true
    v7 = false
    var v8 bool = ((42 <= 255) && (v7 && v7))
    v8 = false
    _ = v7
    _ = v8
    return 2
}
func genFn1() int64 {
    v9 := []int{203, -950, -459}
    v9 = append(v9, 32767)
    var v10 uint = 208
    v10 = v10
    _ = v10
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = -1
    v3 := make(map[string]uint32)
    v2 -= 32767
    v2 -= 1
    var v11 string = "\n\t"
    var v12 string = (v11 + "\n\t")
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint = 401
            var v16 int64 = 1
            var v17 float64 = float64(1.5)
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        if (42 <= 42) {
            v18 := GenS1{F0: 705, F1: 0, F2: 100}
            _ = v18
        }
        if false {
            v19 := GenS1{F0: 1, F1: 42, F2: -90}
            if false {
                v12 = v11
                for fi20 := 0; fi20 < 6; fi20++ {
                    v12 = v11
                    var v21 uint32 = 290
                    v12 = ("" + ((v12 + "foo bar") + ("foo bar" + "")))
                    _ = fi20
                    _ = v21
                }
            } else {
                var v22 uint32 = 2
                if (1 == 128) {
                    for fi23 := 0; fi23 < 3; fi23++ {
                        v24 := GenS2{F0: (v2 - v2), F1: 255}
                        _ = fi23
                        _ = v24
                    }
                    v25 := GenS1{F0: 18446744073709551615, F1: 2, F2: 100}
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 20; fi26++ {
                        var v27 int = (fi13 + fi26)
                        if (!true) {
                            fmt.Println(v22, v11, fi26)
                            v11 = v11
                        }
                        var v28 int8 = -1
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
                _ = v22
            }
            switch v12 {
            case "":
                v2 = (v2 | (v2 - (-(v2))))
            case "foo bar":
                v29 := GenS1{F0: 65535, F1: 710, F2: -1}
                _ = v29
            default:
                v30 := []int{(fi13 & -2147483648), (fi13 | fi13), (fi13 / ((-921) | 1))}
                _ = v30
            }
            _ = v19
        } else {
            v11 = ((v11 + v11) + v11)
        }
        _ = fi13
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        v11 = ("foo bar" + v12)
        switch v11 {
        case "hello":
            if (100 != 87) {
                v32 := []int{fi31, (^(fi31)), -2147483648}
                _ = v32
            }
        }
        _ = fi31
    }
    if false {
        switch 131 {
        case -187:
            fmt.Println(v12, v11, v11)
        }
    } else {
        _ = len(v1)
        for fi33 := 0; fi33 < 4; fi33++ {
            if false {
                v34 := GenS3{F0: 42}
                v35 := []int{}
                _ = v34
                _ = v35
            }
            v36 := GenS3{F0: 65535}
            _ = fi33
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        v2 = v2
        v11 = v12
        var v38 int = (-(fi37))
        _ = fi37
        _ = v38
    }
    if (!false) {
        fmt.Println(v12, v11)
        v39 := []int{4294967295}
        for fi40 := 0; fi40 < 7; fi40++ {
            var v41 uint16 = 1
            _ = fi40
            _ = v41
        }
        _ = v39
    }
    v12 = v11
    v12 = ""
    switch 128 {
    case 0:
        fmt.Println(v2, v11)
    default:
        v2 -= (-(v2))
    }
    v42 := []int{-2147483648, 32767}
    fmt.Println(v12, v2, v2)
    v43 := GenS2{F0: -1, F1: 287}
    for fi44 := 0; fi44 < 11; fi44++ {
        if len(v42) > 0 {
            v42[0] = fi44
        }
        v11 = (v11 + v12)
        _ = fi44
    }
    var v45 int = 9223372036854775807
    for fi46 := 0; fi46 < 14; fi46++ {
        switch fi46 {
        case 4294967295:
            var v47 int = (v45 * (fi46 >> ((255) & 63)))
            _ = v47
        case -39:
            fmt.Println(fi46, v12, fi46)
        case 128:
            var v48 uint32 = genFn0()
            _ = v48
        }
        for fi49 := 0; fi49 < 11; fi49++ {
            v50 := GenS4{F0: 356, F1: 865, F2: -1}
            _ = fi49
            _ = v50
        }
        _ = fi46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v45
}
