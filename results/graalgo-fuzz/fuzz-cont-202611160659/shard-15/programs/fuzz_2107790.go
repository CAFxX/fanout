package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint8, p1 bool) uint {
    p0 = p0
    if false {
        for fi7 := 0; fi7 < 13; fi7++ {
            p0 = p0
            for fi8 := 0; fi8 < 14; fi8++ {
                p1 = false
                fmt.Println(fi7)
                v9 := GenS1{F0: 2, F1: 4294967295}
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    }
    fmt.Println(p0)
    if (42 <= 42) {
        var v10 bool = (2 >= 1)
        _ = v10
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 uint32) bool {
    switch "foo bar" {
    case "hello":
        v11 := []int{}
        _ = v11
    case "a":
        v12 := []int{1, 718, 9223372036854775807, -342}
        _ = v12
    case "":
        var v13 uint8 = 42
        _ = v13
    default:
        for fi14 := 0; fi14 < 2; fi14++ {
            if true {
                var v15 float32 = ((3.14159 - -2.25) / ((0.1 - ((0.0 + 3.14159) * (1e10 + 1.5))) / float32(1e10)))
                fmt.Println(p0, p0, fi14)
                _ = v15
            } else {
                p0 = (p0 & p0)
            }
            switch fi14 {
            case 42:
                var v16 float32 = (3.14159 / (0.1 - float32(-2.25)))
                _ = v16
            case 771:
                if (true || (42 != 1)) {
                    p0 -= p0
                    switch fi14 {
                    case -2147483648:
                        fmt.Println(fi14, fi14)
                    }
                    for fi17 := 0; fi17 < 8; fi17++ {
                        v18 := []int{fi14}
                        var v19 uint8 = 1
                        var v20 uint64 = 961
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                } else {
                    p0 = 65535
                    p0 += (p0 + (p0 ^ p0))
                }
            }
            _ = fi14
        }
    }
    v21 := GenS1{F0: 230, F1: 4294967295}
    switch -1 {
    case 874:
        fmt.Println(p0)
    case 674:
        var v22 int32 = 0
        _ = v22
    case 2147483647:
        p0 += (p0 % ((p0) | 1))
    }
    _ = p0
    _ = v21
    return true
}

func main() {
    var v1 uint8 = 255
    var v2 int = -755
    v3 := make(map[int]bool)
    var v4 int = 1
    var v5 string = "hello"
    var v6 uint8 = 128
    v2 -= (v2 % (((v4 << ((0) & 63))) | 1))
    v1 -= v6
    fmt.Println(v6)
    v23 := []int{(v2 + v4), (v2 | v2), v4, 128}
    var v24 int8 = 48
    switch (^(v2)) {
    case 2147483647:
        if false {
            if (0 <= 2) {
                switch 128 {
                case -670:
                    v6 = (v6 * 163)
                case 4294967295:
                    for fi25 := 0; fi25 < 2; fi25++ {
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v24 = v24
                            _ = fi26
                        }
                        fmt.Println(fi25, v5, v2)
                        _ = fi25
                    }
                default:
                    var v27 int16 = int16(v24)
                    _ = v27
                }
                for fi28 := 0; fi28 < 14; fi28++ {
                    var v29 uint16 = 32767
                    _ = fi28
                    _ = v29
                }
            }
            switch v2 {
            case -630:
                v24 = (^(v24))
            default:
                var v30 bool = (!true)
                _ = v30
            }
        }
    case 0:
        delete(v3, 42)
    }
    v31 := GenS1{F0: 35, F1: 0}
    v32 := v3[v2]
    var v33 uint32 = 1
    v34 := GenS1{F0: 255, F1: 1}
    var v35 int = (v2 + (v2 << ((144) & 63)))
    fmt.Println(v32, v1, v24)
    fmt.Println(v24, v6, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
