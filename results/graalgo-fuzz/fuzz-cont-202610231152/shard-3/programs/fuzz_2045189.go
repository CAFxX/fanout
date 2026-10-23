package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
func genFn0(p0 int16, p1 uint64) bool {
    fmt.Println(p0, p0)
    var v7 string = "\n\t"
    _ = p0
    _ = p1
    _ = v7
    return (!true)
}
func genFn1(p0 uint8) int32 {
    for fi8 := 0; fi8 < 19; fi8++ {
        switch fi8 {
        case -698:
            if false {
                switch (fi8 ^ fi8) {
                case -221:
                    v9 := GenS1{F0: 254, F1: (fi8 % (((fi8 >> ((275) & 63))) | 1))}
                    _ = v9
                case 2147483647:
                    var v10 string = ("hello" + "")
                    _ = v10
                default:
                    if (2 < (p0 * (p0 / (((p0 / ((p0) | 1))) | 1)))) {
                        p0 = p0
                    }
                }
                fmt.Println(fi8)
            } else {
                v11 := GenS1{F0: -993, F1: 1}
                _ = v11
            }
        default:
            for fi12 := 0; fi12 < 10; fi12++ {
                p0 = (p0 & (p0 << ((0) & 63)))
                v13 := []int{-696}
                _ = fi12
                _ = v13
            }
        }
        var v14 int64 = 42
        switch "\n\t" {
        case "\n\t":
            var v15 bool = false
            _ = v15
        default:
            if (false && true) {
                if true {
                    v16 := []int{fi8}
                    _ = v16
                } else {
                    switch (fi8 & fi8) {
                    case 1:
                        var v17 uint8 = (p0 & p0)
                        _ = v17
                    default:
                        v14 = (v14 >> ((42) & 63))
                    }
                    switch fi8 {
                    case -9223372036854775808:
                        if false {
                            var v18 int64 = (v14 % ((0) | 1))
                            var v19 uint32 = 65535
                            _ = v18
                            _ = v19
                        } else {
                            var v20 uint8 = (p0 + p0)
                            _ = v20
                        }
                    case 32767:
                        p0 = p0
                    case 1:
                        var v21 uint = 65535
                        _ = v21
                    }
                }
            }
        }
        _ = fi8
        _ = v14
    }
    p0 += (p0 % (((p0 - (p0 << ((4294967295) & 63)))) | 1))
    p0 -= 76
    _ = p0
    return 1
}

func main() {
    var v1 int8 = 127
    v2 := []string{(("" + "\n\t") + "foo bar")}
    var v3 float32 = (1.5 * -2.25)
    var v4 uint = 348
    v5 := []string{"", "foo bar", (("" + "foo bar") + "")}
    v6 := []string{"", "hello", "\n\t", "\n\t"}
    var v22 uint32 = 4294967295
    switch "" {
    case "a":
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    default:
        v4 = v4
    }
    if len(v6) > 0 {
        v6[0] = "foo bar"
    }
    v23 := GenS1{F0: 0, F1: -511}
    var v24 uint16 = 42
    var v25 uint64 = 4294967295
    var v26 uint16 = 32767
    switch -9223372036854775808 {
    case -645:
        if (!true) {
            var v27 int16 = 32767
            var v28 uint16 = v24
            _ = v27
            _ = v28
        }
    case 0:
        v26 = 0
    }
    switch -766 {
    case 127:
        for fi29 := 0; fi29 < 6; fi29++ {
            v26 = (v26 << ((v4) & 63))
            v2 = append(v2, ("a" + ("\n\t" + ("hello" + "\n\t"))))
            _ = fi29
        }
    default:
        var v30 uint32 = (v22 ^ v22)
        _ = v30
    }
    var v31 float64 = (((3.14159 * float64((v3 - 1e10))) * 3.14159) / (-2.25 - (0.0 + float64((float32(0.0) - 0.1)))))
    var v32 int8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
