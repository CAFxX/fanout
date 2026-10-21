package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int32; F1 int; F2 int8 }
func genFn0(p0 uint16) uint32 {
    var v4 int16 = int16(p0)
    v4 = v4
    v4 = -655
    switch 9223372036854775807 {
    case -50:
        if false {
            v5 := []int{}
            var v6 uint16 = (^(p0))
            v7 := GenS1{F0: 42, F1: (v4 << ((2) & 63))}
            _ = v5
            _ = v6
            _ = v7
        }
    case 4294967295:
        var v8 int = 0
        _ = v8
    default:
        v9 := GenS1{F0: -558, F1: (v4 / ((v4) | 1))}
        _ = v9
    }
    _ = p0
    _ = v4
    return 2147483647
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint = 1
    var v3 int64 = -9223372036854775808
    v2 -= v2
    if true {
        v10 := []int{-236, -810}
        if (false && false) {
            v1 = (3.14159 / (v1 * 0.0))
        }
        v10 = append(v10, -814)
        _ = v10
    } else {
        v1 = -2.25
        v11 := GenS1{F0: (v3 / ((v3) | 1)), F1: 0}
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := []int{9223372036854775807, fi12, (fi12 * (fi12 | fi12)), fi12}
            var v14 uint64 = 2
            v3 += -9223372036854775808
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    }
    var v15 int8 = 77
    fmt.Println(v1, v1, v15)
    var v16 bool = true
    var v17 int32 = 0
    var v18 uint16 = 2
    v1 = v1
    v3 = (^(v3))
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 int8 = -102
        _ = fi19
        _ = v20
    }
    fmt.Println(v16, v16)
    v17 += (v17 ^ 864)
    if (0 != (v2 >> ((1) & 63))) {
        v2 = (v2 & 513)
    } else {
        fmt.Println(v3, v3)
        fmt.Println(v17)
    }
    if true {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 float64 = v1
            v23 := GenS2{F0: uint32(v18), F1: 100}
            v24 := GenS2{F0: 254, F1: 42}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        v25 := GenS2{F0: 2, F1: 796}
        fmt.Println(v18)
        _ = v25
    } else {
        fmt.Println(v18, v1, v1)
        if (!true) {
            v1 = -2.25
        } else {
            if v16 {
                var v26 int = 407
                switch ((("" + "hello") + "foo bar") + ("a" + ("foo bar" + "foo bar"))) {
                case "a":
                    switch (v26 >> ((65535) & 63)) {
                    case -1:
                        v27 := GenS3{F0: v17, F1: (v26 - (v26 & (v26 | (v26 % ((v26) | 1))))), F2: v15}
                        _ = v27
                    case -2147483648:
                        var v28 uint16 = (^(v18))
                        _ = v28
                    case -392:
                        v16 = (!v16)
                    }
                case "foo bar":
                    v29 := GenS2{F0: 42, F1: uint32(v2)}
                    _ = v29
                case "hello":
                    switch (v26 >> (((v2 - v2)) & 63)) {
                    case 32767:
                        var v30 uint64 = 2
                        _ = v30
                    case -574:
                        fmt.Println(v26, v16)
                    case 139:
                        v2 = 275
                    default:
                        var v31 uint64 = 18446744073709551615
                        _ = v31
                    }
                }
                _ = v26
            } else {
                v32 := []int{42, 32768, 42, -9223372036854775808}
                _ = v32
            }
        }
        var v33 int32 = v17
        _ = v33
    }
    v15 = (v15 | v15)
    for fi34 := 0; fi34 < 7; fi34++ {
        switch 395 {
        case 4294967295:
            fmt.Println(fi34)
        case 9223372036854775807:
            for fi35 := 0; fi35 < 2; fi35++ {
                fmt.Println(v1, v15)
                v36 := GenS3{F0: (v17 & (v17 & v17)), F1: (fi34 & 127), F2: (-(v15))}
                _ = fi35
                _ = v36
            }
        case 0:
            fmt.Println(v17, v16)
        default:
            v18 = 736
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
