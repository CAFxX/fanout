package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32, p1 uint16) uint64 {
    var v4 float32 = (1e10 - -2.25)
    v5 := []int{32767, int(p0), 42}
    fmt.Println(v4)
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return uint64(p1)
}
func genFn1(p0 int32) uint16 {
    if true {
        if false {
            p0 -= -2147483648
            switch "hello" {
            case "":
                var v6 string = ("" + "hello")
                _ = v6
            }
        }
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := []int{(fi7 ^ 32767), fi7}
            if len(v8) > 0 {
                v8[0] = (fi7 % (((fi7 - (fi7 % ((fi7) | 1)))) | 1))
            }
            for fi9 := 0; fi9 < 10; fi9++ {
                p0 -= p0
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        if true {
            p0 -= (^(p0))
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint8 = 0
        fmt.Println(p0, v12, v12)
        _ = fi11
        _ = v12
    }
    v13 := []int{128, 127}
    if len(v13) > 0 {
        v13[0] = -860
    }
    _ = p0
    _ = v13
    return 255
}

func main() {
    var v1 uint16 = 336
    var v2 int16 = 42
    var v3 int16 = -542
    v14 := GenS1{F0: 1, F1: true}
    v1 -= (^(v1))
    if true {
        var v15 int64 = 0
        _ = v15
    } else {
        if (127 != 255) {
            fmt.Println(v1, v1, v3)
        } else {
            if true {
                v16 := GenS2{F0: v3}
                v17 := []int{42, -9223372036854775808, 32768}
                if false {
                    v2 = v2
                } else {
                    if false {
                        switch "\n\t" {
                        case "":
                            if (int(v1) != -221) {
                                v2 -= (v2 ^ v2)
                            } else {
                                fmt.Println(v2, v1)
                            }
                        default:
                            var v18 uint64 = 2
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 8; fi19++ {
                            var v20 int32 = -1
                            fmt.Println(v2, v20)
                            _ = len(v17)
                            _ = fi19
                            _ = v20
                        }
                    }
                    if len(v17) > 0 {
                        v17[0] = -9223372036854775808
                    }
                }
                _ = v16
                _ = v17
            } else {
                var v21 uint16 = v1
                v22 := GenS2{F0: v2}
                _ = v21
                _ = v22
            }
            v23 := []int{-623, -90, -568}
            var v24 float32 = ((1e10 + float32(3.14159)) * 1e10)
            _ = v23
            _ = v24
        }
        var v25 uint16 = v1
        v26 := []int{4294967295}
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 bool = (("" + ("\n\t" + "\n\t")) < "foo bar")
        var v29 bool = (false || true)
        _ = fi27
        _ = v28
        _ = v29
    }
    switch "a" {
    case "":
        var v30 int16 = (v2 << ((uint(v1)) & 63))
        _ = v30
    case "hello":
        v31 := []int{}
        _ = v31
    default:
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 int64 = 1
            v1 = (v1 & (v1 & (v1 / ((2) | 1))))
            var v34 int32 = 0
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    fmt.Println(v2, v3, v3)
    v2 -= (v3 / (((v2 / ((-142) | 1))) | 1))
    fmt.Println(v2, v3)
    for fi35 := 0; fi35 < 8; fi35++ {
        fmt.Println(v2, v1, fi35)
        _ = fi35
    }
    if false {
        fmt.Println(v1, v3)
        switch 32768 {
        case 32767:
            fmt.Println(v3)
        case -9223372036854775808:
            for fi36 := 0; fi36 < 15; fi36++ {
                v2 = (v2 % (((v3 / ((v3) | 1))) | 1))
                _ = fi36
            }
        case 9223372036854775807:
            v37 := []int{0, 42, 127}
            _ = v37
        default:
            v1 = v1
        }
        if ((v3 / (((v3 % ((v2) | 1))) | 1)) >= v3) {
            var v38 float32 = 3.14159
            v39 := GenS2{F0: (v2 + (v2 >> ((815) & 63)))}
            _ = v38
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
