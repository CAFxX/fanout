package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint }
func genFn0(p0 int) uint64 {
    for fi4 := 0; fi4 < 1; fi4++ {
        if (2 >= 1) {
            v5 := GenS1{F0: fi4}
            p0 = fi4
            var v6 uint32 = 0
            _ = v5
            _ = v6
        }
        p0 = (p0 / ((fi4) | 1))
        var v7 float32 = (1e10 * (((-2.25 / 1e10) - (0.1 + ((float32(0.0) / (-2.25 * (0.1 / ((3.14159 + ((float32(-2.25) - -2.25) * 0.1)) - 0.1)))) - 3.14159))) - ((3.14159 - 1.5) - float32((-2.25 / 1e10)))))
        _ = fi4
        _ = v7
    }
    p0 += (p0 & (^(p0)))
    if false {
        v8 := GenS1{F0: p0}
        v9 := []int{(p0 >> ((332) & 63))}
        _ = v8
        _ = v9
    } else {
        switch ("" + "\n\t") {
        case "a":
            switch "hello" {
            case "hello":
                for fi10 := 0; fi10 < 8; fi10++ {
                    p0 = (fi10 % (((p0 - fi10)) | 1))
                    _ = fi10
                }
            case "a":
                p0 -= p0
            case "":
                switch (p0 + 1) {
                case 0:
                    p0 = (p0 | p0)
                case 2147483647:
                    switch 9223372036854775807 {
                    case -360:
                        p0 -= p0
                    default:
                        v11 := GenS2{F0: 110, F1: 4294967295}
                        _ = v11
                    }
                default:
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v13 := []int{}
                        v14 := GenS2{F0: 46, F1: 65535}
                        switch "\n\t" {
                        case "hello":
                            for fi15 := 0; fi15 < 6; fi15++ {
                                fmt.Println(p0, p0)
                                var v16 string = ""
                                p0 = (p0 & p0)
                                _ = fi15
                                _ = v16
                            }
                        case "a":
                            switch "" {
                            case "\n\t":
                                var v17 float64 = -2.25
                                _ = v17
                            default:
                                switch "a" {
                                case "foo bar":
                                    var v18 int16 = -398
                                    _ = v18
                                }
                            }
                        case "\n\t":
                            if len(v13) > 0 {
                                v13[0] = -197
                            }
                        default:
                            var v19 float32 = 1e10
                            _ = v19
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
            default:
                var v20 string = "a"
                _ = v20
            }
        case "":
            p0 += (p0 * 9223372036854775807)
        }
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        if false {
            var v22 uint = 255
            _ = v22
        } else {
            p0 = (p0 | (fi21 | (p0 / (((fi21 & (p0 ^ fi21))) | 1))))
        }
        v23 := []int{fi21, fi21, 4294967295, (p0 / ((-496) | 1))}
        var v24 float32 = 1.5
        _ = fi21
        _ = v23
        _ = v24
    }
    var v25 float32 = (0.0 - 0.1)
    _ = p0
    _ = v25
    return 757
}

func main() {
    v1 := []uint32{2147483647, 2147483647, 42}
    v2 := []int{}
    var v3 float32 = (3.14159 * 0.0)
    if len(v1) > 0 {
        v1[0] = 200
    }
    fmt.Println(v3)
    for fi26 := 0; fi26 < 16; fi26++ {
        fmt.Println(fi26)
        v27 := GenS1{F0: (fi26 * 2147483647)}
        _ = fi26
        _ = v27
    }
    v28 := GenS2{F0: 255, F1: 6}
    v1 = append(v1, 1)
    v3 -= (0.0 + 0.1)
    if len(v2) > 0 {
        v2[0] = 770
    }
    var v29 uint16 = 2
    var v30 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
