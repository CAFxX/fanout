package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint; F1 bool; F2 string }
type GenS5 struct { F0 uint32; F1 int8; F2 string }
type GenS6 struct { F0 uint32; F1 string; F2 uint64 }
func genFn0() int32 {
    var v4 uint64 = 1
    v4 -= v4
    v5 := []int{-362}
    _ = v4
    _ = v5
    return -156
}
func genFn1(p0 int8, p1 bool) uint32 {
    fmt.Println(p1, p0)
    switch -2147483648 {
    case 1:
        v6 := GenS1{F0: -551, F1: (-(p0))}
        _ = v6
    }
    if p1 {
        if p1 {
            v7 := GenS1{F0: 9223372036854775807, F1: (p0 / ((p0) | 1))}
            _ = v7
        }
    } else {
        p1 = false
        v8 := GenS1{F0: 1, F1: -109}
        _ = v8
    }
    var v9 int32 = -29
    v10 := []int{127, 32768, 2147483647}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 867
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -348
    var v3 int16 = -1
    var v11 int16 = v3
    v11 += (v11 - v11)
    var v12 int16 = (v11 % (((v3 % (((v3 + v3)) | 1))) | 1))
    if true {
        fmt.Println(v12, v2, v12)
        fmt.Println(v3, v11, v2)
        v12 += v12
    } else {
        v13 := []int{774, 4294967295}
        v14 := GenS1{F0: -9223372036854775808, F1: 127}
        v15 := GenS2{F0: 42}
        _ = v13
        _ = v14
        _ = v15
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v11, v3, v2)
    case "\n\t":
        switch "" {
        case "a":
            fmt.Println(v11, v12)
        case "":
            fmt.Println(v11)
        case "foo bar":
            v3 += v11
        default:
            if false {
                v16 := []int{785}
                switch 0 {
                case -1:
                    v3 += (v12 + v12)
                default:
                    v16 = append(v16, 32767)
                }
                _ = v16
            }
        }
    case "foo bar":
        fmt.Println(v12, v2)
    default:
        v17 := []int{-1, 1, 127}
        _ = v17
    }
    if (231 > 802) {
        v11 += v11
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS1{F0: (fi18 << ((2) & 63)), F1: 0}
            v20 := GenS2{F0: -217}
            var v21 int16 = v11
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        var v22 int16 = (v12 % ((v12) | 1))
        _ = v22
    } else {
        for fi23 := 0; fi23 < 20; fi23++ {
            v24 := GenS2{F0: 782}
            _ = fi23
            _ = v24
        }
    }
    v25 := GenS3{F0: 0, F1: v2}
    v26 := []int{42, -97, 1}
    var v27 int64 = (v2 << ((42) & 63))
    switch "hello" {
    case "foo bar":
        v28 := GenS4{F0: 65535, F1: ((1.5 * ((0.0 / 1.0) / 1.5)) >= 0.1), F2: (((("a" + (("foo bar" + "foo bar") + "\n\t")) + (("a" + "") + "\n\t")) + "") + ("\n\t" + "a"))}
        _ = v28
    case "":
        v29 := GenS2{F0: genFn0()}
        _ = v29
    case "hello":
        v30 := GenS3{F0: 1, F1: v27}
        _ = v30
    }
    v3 = v3
    v2 = (v27 / (((v27 / (((v2 | 960)) | 1))) | 1))
    if (!(false && true)) {
        v31 := GenS3{F0: 4294967295, F1: (-(v2))}
        _ = v31
    }
    if false {
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
        for fi32 := 0; fi32 < 8; fi32++ {
            switch (("\n\t" + "a") + "\n\t") {
            case "":
                var v33 uint16 = 0
                _ = v33
            case "foo bar":
                for fi34 := 0; fi34 < 3; fi34++ {
                    var v35 int8 = 100
                    _ = fi34
                    _ = v35
                }
            case "hello":
                v27 = (v27 * 847)
            default:
                for fi36 := 0; fi36 < 6; fi36++ {
                    var v37 float32 = float32(3.14159)
                    fmt.Println(v11, v37)
                    if (1 != 1) {
                        switch "foo bar" {
                        case "foo bar":
                            v37 = v37
                        }
                        v26 = append(v26, (fi36 / ((fi36) | 1)))
                    }
                    _ = fi36
                    _ = v37
                }
            }
            v3 -= v12
            _ = fi32
        }
    }
    v38 := GenS5{F0: 2147483648, F1: 72, F2: "\n\t"}
    v39 := GenS6{F0: 2147483647, F1: ("foo bar" + "\n\t"), F2: 729}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v38
    _ = v39
}
