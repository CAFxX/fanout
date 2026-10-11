package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint }
func genFn0(p0 int32, p1 int64, p2 int8) int {
    var v4 uint16 = 65535
    var v5 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}
func genFn1() uint8 {
    if (true || true) {
        var v6 int8 = -1
        v6 -= v6
        switch "" {
        case "\n\t":
            v6 = (v6 / (((v6 + v6)) | 1))
        case "":
            v7 := GenS1{F0: 42}
            _ = v7
        default:
            switch "hello" {
            case "":
                if true {
                    var v8 float32 = (0.1 * 1e10)
                    v9 := []int{-1, 9223372036854775807, -844}
                    var v10 uint = 255
                    _ = v10
                    _ = v8
                    _ = v9
                }
            case "a":
                v11 := GenS1{F0: 42}
                _ = v11
            }
        }
        _ = v6
    } else {
        var v12 bool = true
        fmt.Println(v12)
        v13 := []int{}
        _ = v12
        _ = v13
    }
    if (0 > 4294967295) {
        if ("\n\t" != ("" + "a")) {
            if (2 <= 1) {
                v14 := []int{-9223372036854775808, -498, 127}
                v15 := GenS1{F0: 9223372036854775807}
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 uint64 = 42
                    _ = fi16
                    _ = v17
                }
            }
        }
    } else {
        v18 := GenS1{F0: -1}
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 12; fi20++ {
            if false {
                v21 := []int{0, 32767, fi19, 896}
                var v22 float64 = 3.14159
                _ = v21
                _ = v22
            } else {
                var v23 float64 = (-2.25 / 3.14159)
                switch fi19 {
                case 2147483647:
                    v24 := GenS2{F0: 0, F1: 919, F2: 65535}
                    _ = v24
                case -1:
                    v23 -= (v23 * ((v23 / 0.1) * (v23 / (((((v23 - ((v23 + v23) * (0.0 - 1e10))) + (v23 - v23)) * (v23 - v23)) + 3.14159) / 1.0))))
                case -2147483648:
                    v25 := GenS2{F0: 882, F1: 915, F2: 255}
                    _ = v25
                default:
                    var v26 int64 = -1
                    _ = v26
                }
                _ = v23
            }
            _ = fi20
        }
        _ = fi19
    }
    var v27 uint64 = 710
    _ = v27
    return 2
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    var v3 string = ((("foo bar" + "") + ("\n\t" + "a")) + "a")
    var v28 int8 = -128
    var v29 float32 = 0.1
    var v30 uint32 = 42
    v31 := GenS1{F0: 499}
    v28 += v28
    v32 := v2[(v3 + (v3 + (v3 + v3)))]
    switch -360 {
    case 696:
        fmt.Println(v28, v29, v29)
    case 127:
        fmt.Println(v29, v3)
    case 397:
        if (100 > (v28 | -128)) {
            fmt.Println(v29, v3, v29)
            if v32 {
                for fi33 := 0; fi33 < 11; fi33++ {
                    fmt.Println(v32)
                    v30 = (^(v30))
                    for fi34 := 0; fi34 < 8; fi34++ {
                        v29 += v29
                        v35 := []int{}
                        var v36 int = (fi34 % (((fi33 << ((994) & 63))) | 1))
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    _ = fi33
                }
                switch (v3 + v3) {
                case "\n\t":
                    v32 = (false && (!true))
                }
            }
            var v37 int32 = -778
            _ = v37
        }
    default:
        v29 -= v29
    }
    v38 := []int{-750}
    _ = len(v1)
    v32 = (genFn0(-603, 0, (v28 << ((0) & 63))) == 4294967295)
    var v39 string = v3
    v3 = ""
    for fi40 := 0; fi40 < 7; fi40++ {
        if true {
            if len(v38) >= 2 {
                v38 = v38[1:len(v38)]
            }
            if (v32 || v32) {
                if len(v38) > 0 {
                    v38[0] = (fi40 >> ((255) & 63))
                }
                v3 = v39
            }
        }
        _ = fi40
    }
    v30 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v38
    _ = v39
}
