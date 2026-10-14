package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int16; F2 int32 }
func genFn0(p0 uint64) uint8 {
    v7 := []int{4294967295}
    p0 += p0
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]string)
    v3 := []int{1, 128, 0}
    v4 := make(map[string]int)
    var v5 uint16 = 52
    var v6 bool = ((true || (-41 > 42)) || (862 != 1))
    _ = len(v4)
    v6 = true
    if v6 {
        for fi8 := 0; fi8 < 14; fi8++ {
            v1 -= (-(v1))
            if (2147483647 >= 1) {
                var v9 uint8 = 242
                var v10 uint16 = v5
                v5 = v5
                _ = v10
                _ = v9
            } else {
                v1 = (v1 % ((-1) | 1))
                v3 = append(v3, fi8)
            }
            v1 = v1
            _ = fi8
        }
        if (!v6) {
            var v11 float64 = (-2.25 + 1e10)
            _ = v11
        }
        if (!(!((0.1 == 3.14159) && v6))) {
            v3 = append(v3, 54)
            if (genFn0(2) > 2) {
                switch -934 {
                case 32767:
                    for fi12 := 0; fi12 < 12; fi12++ {
                        v13 := GenS1{F0: 127}
                        var v14 string = ""
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case 4294967295:
                    v6 = (!false)
                }
                v5 += 319
                if (0 >= 7) {
                    for fi15 := 0; fi15 < 2; fi15++ {
                        v16 := GenS1{F0: 184}
                        v5 -= 1
                        _ = fi15
                        _ = v16
                    }
                }
            } else {
                v3 = append(v3, 9223372036854775807)
            }
        } else {
            v17 := GenS2{F0: true, F1: v1, F2: 42}
            v18 := GenS2{F0: v6, F1: 1000, F2: 778}
            _ = v17
            _ = v18
        }
    }
    var v19 float64 = 1.5
    var v20 uint64 = 0
    if (!v6) {
        for fi21 := 0; fi21 < 12; fi21++ {
            _ = len(v3)
            _ = fi21
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            if len(v3) > 0 {
                v3[0] = fi22
            }
            _ = fi22
        }
    }
    fmt.Println(v19, v19)
    if (1 < 42) {
        v23 := v4[""]
        switch v23 {
        case 740:
            fmt.Println(v20, v5)
        case -2147483648:
            fmt.Println(v6, v6)
        default:
            v20 = (v20 >> ((65535) & 63))
        }
        _ = v23
    }
    var v24 int32 = 401
    var v25 int16 = (v1 % (((v1 * -1)) | 1))
    for fi26 := 0; fi26 < 17; fi26++ {
        v1 -= -32768
        _ = fi26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        fmt.Println(v20, v20)
        v28 := v4[(("hello" + "foo bar") + ("\n\t" + "hello"))]
        _ = fi27
        _ = v28
    }
    switch int(v20) {
    case 9223372036854775807:
        if (v6 || false) {
            v5 -= v5
            if (v6 || true) {
                var v29 int16 = (-(v1))
                fmt.Println(v25, v25)
                _ = v29
            }
        }
    default:
        var v30 string = "foo bar"
        _ = v30
    }
    v31 := GenS1{F0: 160}
    for fi32 := 0; fi32 < 9; fi32++ {
        _ = len(v3)
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
