package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int) int8 {
    var v4 uint16 = 0
    if (v4 <= v4) {
        v4 += (v4 % (((^(v4))) | 1))
        for fi5 := 0; fi5 < 5; fi5++ {
            for fi6 := 0; fi6 < 9; fi6++ {
                p0 = (fi6 ^ (^(fi5)))
                _ = fi6
            }
            p0 += -864
            _ = fi5
        }
    }
    _ = p0
    _ = v4
    return 0
}
func genFn1() uint16 {
    v7 := []int{32767}
    if true {
        var v8 uint16 = 65535
        _ = v8
    }
    _ = v7
    return 174
}

func main() {
    var v1 int = -58
    v2 := make(map[int]uint32)
    v3 := []int{v1, (v1 % (((v1 | (v1 / (((^(v1))) | 1)))) | 1))}
    var v9 uint32 = 2147483647
    v9 += (v9 ^ (v9 >> ((65535) & 63)))
    v1 += v1
    if true {
        v9 = v9
    }
    if (251 > 42) {
        var v10 int32 = 767
        _ = v10
    } else {
        for fi11 := 0; fi11 < 17; fi11++ {
            if true {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := GenS1{F0: 798}
                    switch fi12 {
                    case 128:
                        if (!true) {
                            if false {
                                v14 := GenS1{F0: 2}
                                v9 = 2
                                v15 := GenS1{F0: 675}
                                _ = v14
                                _ = v15
                            }
                        }
                    case 127:
                        if len(v3) > 0 {
                            v3[0] = -9223372036854775808
                        }
                    }
                    if true {
                        var v16 float64 = 3.14159
                        _ = v16
                    } else {
                        var v17 uint8 = 0
                        _ = v17
                    }
                    _ = fi12
                    _ = v13
                }
                var v18 float32 = (0.0 + (-2.25 * 0.1))
                _ = v18
            }
            for fi19 := 0; fi19 < 2; fi19++ {
                delete(v2, -9223372036854775808)
                if false {
                    delete(v2, v1)
                    fmt.Println(v9)
                }
                v9 += v9
                var v20 float64 = 0.1
                v21 := GenS1{F0: 18446744073709551615}
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi11
        }
        for fi22 := 0; fi22 < 9; fi22++ {
            fmt.Println(v1, v1)
            for fi23 := 0; fi23 < 12; fi23++ {
                fmt.Println(fi23, fi23, fi23)
                v24 := GenS1{F0: 958}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        v9 -= v9
        if true {
            for fi26 := 0; fi26 < 3; fi26++ {
                v1 = -978
                fmt.Println(fi26)
                _ = fi26
            }
        }
        _ = fi25
    }
    v9 = 4294967295
    switch "foo bar" {
    case "":
        fmt.Println(v1)
    case "foo bar":
        for fi27 := 0; fi27 < 8; fi27++ {
            v28 := GenS1{F0: 42}
            _ = fi27
            _ = v28
        }
    case "a":
        var v29 float64 = 0.0
        _ = v29
    default:
        var v30 uint = 65535
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
