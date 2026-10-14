package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
func genFn0(p0 uint) int32 {
    v6 := []int{4294967295, 1, 85}
    var v7 int8 = -125
    _ = p0
    _ = v6
    _ = v7
    return -603
}

func main() {
    var v1 uint8 = 128
    var v2 uint16 = 255
    v3 := []uint32{2, 42, 1, 2147483647}
    var v4 float64 = ((float64(1.5) + -2.25) / ((3.14159 * float64((0.1 + 0.1))) + 1e10))
    var v5 uint = 1
    if false {
        var v8 int16 = 32767
        _ = v8
    }
    v5 = 255
    switch 630 {
    case 128:
        fmt.Println(v5, v5, v1)
    case -770:
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 uint16 = (v2 & (^(v2)))
            v11 := GenS1{F0: (v5 / ((4294967295) | 1)), F1: v5}
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        if true {
            if false {
                v5 -= v5
                switch ("\n\t" + "hello") {
                case "foo bar":
                    switch ("foo bar" + "hello") {
                    case "foo bar":
                        v3 = append(v3, 2147483648)
                    case "hello":
                        fmt.Println(v2, v2)
                    case "":
                        fmt.Println(v4, v4, v2)
                    }
                default:
                    var v12 uint32 = 579
                    _ = v12
                }
                fmt.Println(v1, v1, v5)
            } else {
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 int16 = 32767
                    fmt.Println(v2, v5)
                    _ = fi13
                    _ = v14
                }
            }
            var v15 int32 = 1
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        v1 -= v1
        var v17 bool = ((!(!((714 < 742) || false))) || true)
        var v18 uint8 = v1
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 bool = true
    var v20 int8 = 42
    if (42 != 1) {
        if len(v3) > 0 {
            v3[0] = 1
        }
        var v21 uint64 = uint64(v1)
        _ = v21
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v22 uint32 = 964
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 uint8 = v1
        for fi25 := 0; fi25 < 9; fi25++ {
            v22 += 4294967295
            if v19 {
                if false {
                    v5 = (v5 % (((v5 - v5)) | 1))
                    var v26 uint32 = (v22 / ((uint32(v5)) | 1))
                    if v19 {
                        var v27 int16 = 0
                        fmt.Println(fi23, v2)
                        _ = v27
                    }
                    _ = v26
                } else {
                    v3 = append(v3, v22)
                    var v28 int64 = 314
                    _ = v28
                }
            } else {
                v19 = (!true)
            }
            if len(v3) > 0 {
                v3[0] = v22
            }
            _ = fi25
        }
        for fi29 := 0; fi29 < 8; fi29++ {
            fmt.Println(v1)
            var v30 uint32 = (v22 << ((4294967295) & 63))
            _ = fi29
            _ = v30
        }
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
