package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint; F1 uint32; F2 uint }
type GenS5 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint64, p1 int, p2 int64) uint32 {
    if (false || true) {
        fmt.Println(p0, p2, p0)
    } else {
        fmt.Println(p1, p1)
    }
    p0 = p0
    fmt.Println(p1)
    p1 = (p1 ^ p1)
    for fi5 := 0; fi5 < 20; fi5++ {
        p1 = (fi5 & (p1 << ((4294967295) & 63)))
        var v6 int32 = -2147483648
        var v7 bool = false
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 565
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int = 9223372036854775807
    var v3 int64 = 42
    var v4 int64 = (-(v3))
    if (float32(0.0) <= (0.0 * 0.1)) {
        switch "" {
        case "a":
            switch "foo bar" {
            case "hello":
                fmt.Println(v3)
            case "":
                var v8 int = v2
                _ = v8
            }
        case "hello":
            v1 = 2147483647
        case "":
            v9 := GenS1{F0: ("\n\t" + "\n\t"), F1: true}
            _ = v9
        }
    } else {
        if true {
            switch "\n\t" {
            case "":
                v4 -= (v4 - (v3 & (v3 ^ int64(v2))))
            case "hello":
                if (!true) {
                    var v10 float32 = 1e10
                    _ = v10
                }
            }
            if true {
                var v11 int64 = v4
                _ = v11
            }
            v12 := []int{v2, v2}
            _ = v12
        } else {
            var v13 string = (("" + ((("" + (("a" + ("foo bar" + "foo bar")) + "a")) + "hello") + "")) + "a")
            v3 = -9223372036854775808
            _ = v13
        }
    }
    var v14 int64 = v4
    if false {
        v4 = 9223372036854775807
    }
    v1 = genFn0(0, (v2 << ((20) & 63)), v4)
    v15 := []int{(v2 / ((1) | 1)), (-(v2)), (-(v2)), (-(v2))}
    var v16 bool = (1e10 > ((0.0 - float64(0.0)) * 3.14159))
    v17 := GenS2{F0: 18446744073709551615, F1: 123}
    var v18 bool = ((v1 - 0) < (v1 / ((genFn0(722, 1, 0)) | 1)))
    var v19 uint32 = (v1 >> ((160) & 63))
    if false {
        fmt.Println(v18)
        for fi20 := 0; fi20 < 16; fi20++ {
            for fi21 := 0; fi21 < 17; fi21++ {
                if (2 > 42) {
                    v14 += v3
                    _ = len(v15)
                }
                switch (v2 / (((v2 | fi20)) | 1)) {
                case 1:
                    var v22 float64 = 0.1
                    _ = v22
                case 0:
                    if (32767 > 42) {
                        var v23 int32 = 42
                        v24 := GenS3{F0: -112, F1: 100}
                        switch -2147483648 {
                        case -9223372036854775808:
                            var v25 bool = (false || v18)
                            _ = v25
                        case -1:
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        case 4294967295:
                            var v26 float64 = (3.14159 - 1.5)
                            _ = v26
                        }
                        _ = v23
                        _ = v24
                    }
                }
                if len(v15) > 0 {
                    v15[0] = (^(fi20))
                }
                switch "foo bar" {
                case "":
                    fmt.Println(v2)
                case "foo bar":
                    v27 := GenS4{F0: 39, F1: 252, F2: 0}
                    _ = v27
                default:
                    v18 = v16
                    var v28 uint64 = 0
                    if v16 {
                        v18 = v16
                        var v29 float64 = 0.0
                        var v30 uint = 514
                        _ = v29
                        _ = v30
                    }
                    v18 = v16
                    if v18 {
                        var v31 uint = 255
                        v32 := GenS5{F0: uint32(v28), F1: 0}
                        _ = v31
                        _ = v32
                    }
                    _ = v28
                }
                _ = fi21
            }
            v4 = (v14 & (^(v14)))
            v2 -= (v2 / ((fi20) | 1))
            _ = fi20
        }
    }
    _ = len(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
