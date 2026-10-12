package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint32 }
func genFn0(p0 string, p1 int8) uint8 {
    if true {
        switch p0 {
        case "a":
            var v6 uint32 = 1
            _ = v6
        case "\n\t":
            p1 -= (p1 & p1)
        case "foo bar":
            var v7 int32 = 1
            _ = v7
        }
    }
    var v8 uint16 = 42
    _ = p0
    _ = p1
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 0.0
    v3 := make(map[string]uint32)
    var v4 float32 = 0.1
    var v5 int64 = 241
    var v9 bool = ((v4 + float32(1.5)) == 1.5)
    if true {
        for fi10 := 0; fi10 < 2; fi10++ {
            if true {
                fmt.Println(v4)
                v11 := GenS1{F0: 18, F1: 0, F2: 0}
                _ = v11
            }
            var v12 float32 = (v4 + v4)
            _ = fi10
            _ = v12
        }
        switch 9223372036854775807 {
        case -2147483648:
            switch -489 {
            case 9223372036854775807:
                fmt.Println(v9, v5, v9)
            case 42:
                v9 = v9
            case -656:
                var v13 int = 621
                _ = v13
            default:
                v14 := GenS1{F0: 477, F1: 32767, F2: 2}
                _ = v14
            }
        case 32768:
            var v15 string = "\n\t"
            _ = v15
        case 105:
            for fi16 := 0; fi16 < 13; fi16++ {
                if v9 {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        v4 += ((v4 + v4) + v4)
                        _ = fi17
                    }
                    fmt.Println(v2, v5, v9)
                } else {
                    for fi18 := 0; fi18 < 14; fi18++ {
                        v19 := []int{(fi16 ^ fi16), fi18, (fi16 << ((1) & 63))}
                        _ = fi18
                        _ = v19
                    }
                }
                _ = fi16
            }
        default:
            v20 := []int{2147483647}
            _ = v20
        }
        delete(v3, "foo bar")
    }
    v5 = -989
    var v21 int = -221
    v22 := []int{}
    if len(v22) > 0 {
        v22[0] = (v21 & (v21 + (v21 >> ((18446744073709551615) & 63))))
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        switch "a" {
        case "a":
            v21 = (v21 ^ fi23)
        }
        _ = fi23
    }
    var v24 uint64 = 994
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 string = ("\n\t" + ("" + "\n\t"))
        _ = fi25
        _ = v26
    }
    var v27 int64 = (v5 - 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
