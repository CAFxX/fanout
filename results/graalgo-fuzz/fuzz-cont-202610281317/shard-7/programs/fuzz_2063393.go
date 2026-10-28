package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 string; F1 int16; F2 int }
func genFn0(p0 bool, p1 int8, p2 uint64) int64 {
    fmt.Println(p0, p0, p1)
    var v4 int = 32768
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -9223372036854775808
}
func genFn1(p0 int32, p1 uint32, p2 int64) int32 {
    v5 := []int{1, int(p2)}
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint32 = (p1 / (((p1 << ((18446744073709551615) & 63))) | 1))
            p0 = int32(fi6)
            switch "a" {
            case "":
                fmt.Println(fi7, fi7, v8)
            case "hello":
                p0 = (p0 >> ((670) & 63))
            default:
                var v9 int64 = (p2 ^ -1)
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v10 string = "\n\t"
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 uint8 = 255
        if (false || true) {
            var v13 uint = 1
            var v14 bool = true
            var v15 int = fi11
            _ = v13
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                v10 = ((v10 + "\n\t") + "hello")
                if (p2 <= (p2 - (p2 & genFn0((!true), 34, 2)))) {
                    var v17 uint16 = 663
                    var v18 int8 = 100
                    var v19 float64 = 0.1
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    switch fi16 {
                    case 847:
                        var v20 int16 = -1
                        _ = v20
                    }
                    v21 := GenS1{F0: 0, F1: 1}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v21
                }
                var v22 string = (("foo bar" + v10) + ("" + v10))
                _ = fi16
                _ = v22
            }
            if len(v5) > 0 {
                v5[0] = fi11
            }
            _ = len(v5)
        }
        var v23 uint8 = 255
        _ = fi11
        _ = v12
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    return (p0 | p0)
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []int{42}
    var v3 int = -1
    v3 = v3
    if false {
        _ = len(v2)
    }
    v24 := GenS1{F0: v1, F1: v3}
    var v25 bool = true
    v26 := GenS2{F0: "hello", F1: 32767, F2: -1}
    if (false && (true || (-32768 >= -1))) {
        var v27 int16 = -612
        _ = v27
    }
    if (v3 < (^(v3))) {
        v25 = (100 <= 1)
        v3 = (-(v3))
        v3 = (v3 * (v3 % ((v3) | 1)))
    }
    fmt.Println(v3, v1)
    var v28 int16 = 999
    var v29 int8 = 0
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        fmt.Println(v3, v1, v3)
        var v30 uint16 = 255
        _ = v30
    }
    v25 = false
    switch "" {
    case "foo bar":
        switch -9223372036854775808 {
        case 80:
            switch "" {
            case "hello":
                _ = len(v2)
            case "a":
                fmt.Println(v28)
            case "":
                if (18446744073709551615 == 78) {
                    for fi31 := 0; fi31 < 2; fi31++ {
                        _ = len(v2)
                        _ = fi31
                    }
                    fmt.Println(v28)
                    v28 = (v28 - (v28 >> ((405) & 63)))
                }
            default:
                v3 = v3
            }
        case -9223372036854775808:
            v28 = (v28 << ((554) & 63))
        default:
            v29 = (v29 * -1)
        }
    case "":
        v3 = (v3 + 1)
    default:
        v25 = false
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
}
