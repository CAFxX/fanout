package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint) uint {
    v4 := []int{-9223372036854775808, 2147483647}
    fmt.Println(p0)
    p0 -= 42
    switch "a" {
    case "\n\t":
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 float32 = ((3.14159 + (float32(0.0) - 3.14159)) / 0.1)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi5
            _ = v6
        }
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = v4
    return (^(p0))
}
func genFn1(p0 int8) uint32 {
    v7 := GenS1{F0: -301}
    switch -100 {
    case 0:
        var v8 uint8 = 1
        _ = v8
    default:
        p0 = (p0 & (p0 - (-(p0))))
    }
    switch "" {
    case "a":
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := []int{-166, (fi9 % (((fi9 + (^(fi9)))) | 1)), 0, 953}
            _ = fi9
            _ = v10
        }
    case "":
        v11 := GenS1{F0: 0}
        _ = v11
    case "\n\t":
        v12 := GenS1{F0: 32767}
        _ = v12
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = v7
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 283
    v3 := make(map[string]uint32)
    var v13 int32 = 42
    var v14 uint8 = 42
    fmt.Println(v2, v13, v13)
    var v15 uint32 = 2147483647
    for fi16 := 0; fi16 < 15; fi16++ {
        v17 := []int{(fi16 * (fi16 + fi16)), 375, (fi16 << ((4294967295) & 63)), (fi16 / (((fi16 ^ (^(fi16)))) | 1))}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v14 -= 128
        _ = fi18
    }
    var v19 uint = 0
    v20 := []int{}
    switch (("hello" + "") + "foo bar") {
    case "":
        fmt.Println(v19, v15)
    default:
        v21 := GenS1{F0: 1}
        _ = v21
    }
    var v22 uint8 = (^(v14))
    v14 = v22
    fmt.Println(v2)
    var v23 uint8 = (^(v22))
    if (!true) {
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 float64 = float64((0.1 + (float32(1e10) + -2.25)))
            v15 += uint32(v14)
            switch fi24 {
            case -414:
                var v26 uint16 = (v2 - (v2 - (v2 / (((v2 * v2)) | 1))))
                _ = v26
            case -9223372036854775808:
                v23 = (v14 ^ 42)
            case -941:
                v2 = 508
            default:
                switch (fi24 ^ (fi24 ^ fi24)) {
                case 42:
                    if (!(-203 > fi24)) {
                        _ = len(v20)
                        var v27 uint = 18446744073709551615
                        _ = v27
                    } else {
                        if false {
                            v28 := GenS1{F0: fi24}
                            v15 += 2
                            _ = v28
                        } else {
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            delete(v3, "foo bar")
                            var v29 uint16 = 1
                            _ = v29
                        }
                        var v30 uint8 = uint8(v19)
                        _ = v30
                    }
                case -2147483648:
                    var v31 int8 = 0
                    _ = v31
                case 127:
                    var v32 int8 = 42
                    _ = v32
                }
            }
            _ = fi24
            _ = v25
        }
        var v33 int64 = -91
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
}
