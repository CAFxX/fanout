package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 bool }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int }
func genFn0(p0 int64, p1 uint, p2 int16) int16 {
    v4 := []int{1}
    v5 := GenS1{F0: 415, F1: -2147483648, F2: true}
    if false {
        v6 := GenS1{F0: 1, F1: 0, F2: false}
        if len(v4) > 0 {
            v4[0] = -1
        }
        _ = v6
    }
    if true {
        if false {
            p0 = (-(p0))
            p1 -= (p1 % (((^(p1))) | 1))
        } else {
            for fi7 := 0; fi7 < 15; fi7++ {
                switch "hello" {
                case "foo bar":
                    var v8 uint = (^(p1))
                    _ = v8
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi7
            }
        }
        v9 := GenS1{F0: (p1 & 666), F1: 0, F2: ((false || (!false)) && false)}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            for fi11 := 0; fi11 < 15; fi11++ {
                switch fi10 {
                case 9223372036854775807:
                    var v12 int = (-(fi10))
                    _ = v12
                case 635:
                    var v13 float64 = 0.0
                    _ = v13
                default:
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(p1, fi10)
                    default:
                        if true {
                            v14 := GenS1{F0: p1, F1: 1, F2: false}
                            var v15 uint = p1
                            v15 -= p1
                            _ = v14
                            _ = v15
                        }
                    }
                }
                if len(v4) > 0 {
                    v4[0] = fi11
                }
                v16 := GenS1{F0: (p1 & (p1 * 1)), F1: -510, F2: ((-2.25 == 1.5) && false)}
                _ = fi11
                _ = v16
            }
            var v17 int8 = 42
            _ = fi10
            _ = v17
        }
    }
    if false {
        v18 := GenS2{F0: -1, F1: 65535}
        _ = v18
    } else {
        var v19 int8 = 100
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 32767
}

func main() {
    var v1 float32 = 3.14159
    var v2 int32 = -1
    v3 := make(map[uint32]int)
    switch ("\n\t" + "") {
    case "a":
        v20 := GenS1{F0: 65535, F1: -563, F2: true}
        _ = v20
    case "":
        var v21 uint64 = 924
        _ = v21
    default:
        switch "" {
        case "foo bar":
            for fi22 := 0; fi22 < 17; fi22++ {
                for fi23 := 0; fi23 < 15; fi23++ {
                    if (-128 != 127) {
                        var v24 uint32 = 102
                        _ = v24
                    } else {
                        if (!(fi23 > 66)) {
                            v25 := GenS2{F0: 127, F1: 2147483647}
                            v1 -= (float32(-2.25) + v1)
                            _ = v25
                        } else {
                            fmt.Println(v1, fi23, fi23)
                        }
                        v3[209] = 2147483647
                        var v26 int16 = -32768
                        _ = v26
                    }
                    fmt.Println(v1, fi23)
                    fmt.Println(fi23, fi22)
                    v1 -= v1
                    _ = fi23
                }
                v2 -= (v2 - v2)
                fmt.Println(fi22)
                _ = fi22
            }
        }
    }
    v27 := []int{4294967295, 32768, 422}
    fmt.Println(v2, v1)
    var v28 int16 = genFn0(42, 255, -32768)
    v2 += 83
    fmt.Println(v28)
    var v29 uint32 = 42
    for fi30 := 0; fi30 < 5; fi30++ {
        var v31 float64 = 1.5
        v28 = (^(v28))
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        v33 := GenS3{F0: 32767}
        _ = fi32
        _ = v33
    }
    if true {
        if true {
            v1 -= float32((1e10 * (1.5 * (-2.25 * -2.25))))
        }
    }
    if len(v27) > 0 {
        v27[0] = 32767
    }
    delete(v3, (v29 / ((2) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
