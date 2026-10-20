package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 int) uint16 {
    p0 = p0
    switch -9223372036854775808 {
    case 9223372036854775807:
        fmt.Println(p0)
    case -873:
        if (18446744073709551615 > 1) {
            var v7 uint32 = 2
            _ = v7
        } else {
            for fi8 := 0; fi8 < 8; fi8++ {
                switch (fi8 | p0) {
                case -1:
                    if false {
                        v9 := []int{0, fi8}
                        _ = v9
                    } else {
                        v10 := []int{-1, (fi8 >> ((0) & 63)), (fi8 >> ((641) & 63))}
                        _ = v10
                    }
                case 32767:
                    var v11 int8 = 0
                    _ = v11
                default:
                    var v12 int = (fi8 + 2147483647)
                    _ = v12
                }
                var v13 uint16 = 32767
                var v14 bool = (!false)
                _ = fi8
                _ = v13
                _ = v14
            }
        }
    default:
        v15 := []int{}
        _ = v15
    }
    switch "hello" {
    case "a":
        switch -341 {
        case -46:
            v16 := GenS1{F0: 0, F1: (42 != p0), F2: 276}
            _ = v16
        case 9223372036854775807:
            if false {
                var v17 bool = true
                _ = v17
            }
        }
    case "\n\t":
        switch p0 {
        case 4294967295:
            v18 := []int{(p0 % ((p0) | 1)), p0, (p0 % (((p0 ^ p0)) | 1)), 1}
            _ = v18
        default:
            v19 := GenS1{F0: 636, F1: true, F2: -83}
            _ = v19
        }
    case "hello":
        p0 += 276
    default:
        var v20 uint = 18446744073709551615
        _ = v20
    }
    _ = p0
    return 32767
}

func main() {
    var v1 int16 = 1
    v2 := []string{"hello", ("" + "\n\t"), ""}
    var v3 uint16 = 65535
    var v4 int8 = 0
    var v5 uint32 = 4294967295
    var v6 float32 = -2.25
    var v21 int8 = 5
    if false {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 uint = 528
            v5 = v5
            _ = fi22
            _ = v23
        }
        v2 = append(v2, ("\n\t" + "a"))
        v3 = uint16(v5)
    }
    fmt.Println(v1)
    var v24 int8 = (v21 << ((540) & 63))
    var v25 int32 = 627
    var v26 bool = false
    var v27 float64 = float64(v6)
    v28 := GenS1{F0: 475, F1: v26, F2: v25}
    var v29 uint = 1
    switch "a" {
    case "":
        v30 := GenS1{F0: -1, F1: (0 >= 0), F2: v25}
        _ = v30
    case "a":
        v6 += 1.5
    default:
        fmt.Println(v6, v1, v21)
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 float64 = (v27 * v27)
        if len(v2) > 0 {
            v2[0] = "a"
        }
        switch (fi31 + fi31) {
        case 4294967295:
            if (v26 || true) {
                v26 = v26
                var v33 uint64 = 143
                v2 = append(v2, ("foo bar" + "a"))
                _ = v33
            } else {
                v34 := GenS2{F0: (((fi31 ^ 9223372036854775807) < (-(fi31))) && (!v26))}
                var v35 bool = (!v26)
                v36 := GenS1{F0: 1, F1: true, F2: v25}
                _ = v34
                _ = v35
                _ = v36
            }
        case -775:
            var v37 uint8 = 42
            _ = v37
        case 127:
            fmt.Println(v21)
        }
        _ = fi31
        _ = v32
    }
    switch 42 {
    case -2147483648:
        v29 += (v29 ^ 273)
    default:
        if (false && v26) {
            var v38 int16 = 1
            _ = v38
        } else {
            var v39 int16 = (v1 + int16(v24))
            _ = v39
        }
    }
    var v40 uint16 = (v3 & (v3 & (^(v3))))
    v41 := GenS3{F0: (v1 >> ((v29) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
