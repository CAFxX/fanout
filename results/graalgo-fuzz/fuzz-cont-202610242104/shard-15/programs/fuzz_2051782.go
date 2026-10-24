package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 string; F1 int; F2 int32 }
type GenS4 struct { F0 int; F1 uint32; F2 uint16 }
func genFn0(p0 uint8, p1 int64) int8 {
    p1 = (p1 * p1)
    p0 = (p0 + p0)
    switch 693 {
    case 32767:
        switch 127 {
        case -2147483648:
            v4 := []int{4294967295, -9223372036854775808}
            _ = v4
        case 42:
            if (p1 >= -830) {
                v5 := []int{}
                _ = v5
            }
        case -1:
            p1 = p1
        default:
            v6 := GenS1{F0: p1, F1: 145}
            _ = v6
        }
    default:
        if true {
            p1 -= p1
            p0 += 42
            fmt.Println(p1, p0)
        } else {
            var v7 uint32 = 42
            fmt.Println(v7, p0, p0)
            switch "\n\t" {
            case "":
                v7 = 2147483647
            default:
                if (2 <= 32767) {
                    p1 += (p1 + p1)
                    var v8 int32 = 743
                    var v9 uint8 = p0
                    _ = v8
                    _ = v9
                } else {
                    if (!(!true)) {
                        v10 := []int{1, -67, 128}
                        switch 42 {
                        case 32767:
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        case 1:
                            fmt.Println(p1, p1)
                        }
                        p0 = 51
                        _ = v10
                    }
                }
            }
            _ = v7
        }
    }
    v11 := GenS1{F0: p1, F1: 0}
    p0 -= (p0 % ((22) | 1))
    _ = p0
    _ = p1
    _ = v11
    return 0
}
func genFn1(p0 int64, p1 bool) string {
    var v12 uint64 = 902
    v13 := []int{}
    if p1 {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 bool = p1
            p0 -= (p0 * p0)
            _ = fi14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return ("" + "foo bar")
}

func main() {
    var v1 int16 = -1
    var v2 uint16 = 1
    var v3 uint32 = 2147483647
    if false {
        if (true && true) {
            for fi16 := 0; fi16 < 9; fi16++ {
                fmt.Println(v3, v1)
                _ = fi16
            }
        }
    } else {
        switch -9223372036854775808 {
        case 2147483647:
            v17 := []int{128, -1, 127}
            _ = v17
        default:
            v3 = 214
        }
        var v18 float32 = float32(float64(((float32(-2.25) + 1.5) + 0.0)))
        _ = v18
    }
    if (409 != -1) {
        v19 := []int{}
        _ = v19
    } else {
        fmt.Println(v1)
        var v20 float32 = 1.5
        _ = v20
    }
    v21 := []int{int(v2)}
    v1 += (v1 ^ v1)
    fmt.Println(v2, v3)
    v22 := GenS2{F0: -365, F1: 0}
    var v23 uint8 = 2
    v23 -= (v23 >> ((149) & 63))
    if (v3 > v3) {
        if (3.14159 < 1e10) {
            var v24 bool = true
            v25 := GenS2{F0: int(v1), F1: 128}
            var v26 string = ("foo bar" + "hello")
            _ = v24
            _ = v25
            _ = v26
        } else {
            v27 := GenS3{F0: ("\n\t" + ""), F1: -220, F2: 281}
            fmt.Println(v1)
            _ = v27
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            _ = len(v21)
            var v29 int32 = 2147483647
            v30 := GenS4{F0: fi28, F1: (v3 >> ((1) & 63)), F2: 12}
            _ = fi28
            _ = v29
            _ = v30
        }
    } else {
        var v31 uint64 = 743
        _ = v31
    }
    fmt.Println(v3, v1)
    switch "\n\t" {
    case "foo bar":
        var v32 uint64 = 42
        _ = v32
    case "a":
        fmt.Println(v2)
    }
    if ((v23 | v23) > v23) {
        switch -751 {
        case 32768:
            var v33 float64 = 1e10
            _ = v33
        case 4294967295:
            fmt.Println(v2, v3)
        case 0:
            var v34 int = 128
            _ = v34
        }
        fmt.Println(v1)
    } else {
        if false {
            for fi35 := 0; fi35 < 15; fi35++ {
                v3 = v3
                v21 = append(v21, (fi35 << ((557) & 63)))
                v23 += 255
                _ = fi35
            }
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            v21 = append(v21, 127)
        }
        v21 = append(v21, -702)
        var v36 int32 = 1
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
