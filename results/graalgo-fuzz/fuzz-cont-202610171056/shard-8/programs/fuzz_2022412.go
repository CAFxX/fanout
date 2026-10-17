package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 string; F1 string; F2 int }
type GenS3 struct { F0 uint16; F1 int16; F2 bool }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int; F1 int8; F2 int8 }
func genFn0() uint {
    var v7 int16 = -32768
    if ((float32(3.14159) * float32(float64(0.1))) <= float32((1.5 / 0.1))) {
        var v8 uint8 = 127
        _ = v8
    } else {
        var v9 float32 = 1.5
        v10 := GenS1{F0: v7, F1: 65535}
        _ = v10
        _ = v9
    }
    v7 -= (v7 << ((1) & 63))
    fmt.Println(v7)
    _ = v7
    return 2
}
func genFn1(p0 int16) int16 {
    for fi11 := 0; fi11 < 18; fi11++ {
        p0 -= (p0 >> ((1) & 63))
        for fi12 := 0; fi12 < 16; fi12++ {
            fmt.Println(fi12, fi12)
            var v13 uint32 = 485
            p0 = p0
            _ = fi12
            _ = v13
        }
        var v14 uint32 = 2147483648
        _ = fi11
        _ = v14
    }
    if (984 == 1) {
        p0 -= -32768
    } else {
        switch "a" {
        case "a":
            v15 := GenS1{F0: (p0 << ((42) & 63)), F1: 0}
            _ = v15
        }
        switch "a" {
        case "":
            p0 = (p0 ^ (p0 - p0))
        case "a":
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 uint32 = 1
                v18 := []int{(-(fi16)), (fi16 * (fi16 + -9223372036854775808)), (fi16 & (fi16 * fi16))}
                switch fi16 {
                case 781:
                    if (false || true) {
                        v19 := GenS2{F0: "\n\t", F1: "", F2: (fi16 * 9223372036854775807)}
                        _ = v19
                    }
                case 127:
                    v20 := GenS1{F0: 42, F1: 2}
                    _ = v20
                case -932:
                    if false {
                        v17 = v17
                        v21 := GenS2{F0: "", F1: "hello", F2: -2147483648}
                        v18 = append(v18, fi16)
                        _ = v21
                    } else {
                        v17 -= (v17 & v17)
                    }
                default:
                    _ = len(v18)
                }
                _ = fi16
                _ = v17
                _ = v18
            }
        default:
            fmt.Println(p0)
        }
        v22 := GenS3{F0: 0, F1: (p0 & (p0 & (p0 >> ((2) & 63)))), F2: (32767 == 740)}
        _ = v22
    }
    _ = p0
    return (p0 << ((1) & 63))
}

func main() {
    var v1 int16 = -584
    var v2 uint = 10
    v3 := make(map[string]uint32)
    var v4 bool = true
    var v5 uint64 = 4294967295
    v6 := []int{-9223372036854775808, 164, -9223372036854775808, 32767}
    _ = len(v3)
    if false {
        var v23 uint8 = 252
        fmt.Println(v2)
        v3[""] = 2147483647
        _ = v23
    }
    v1 += genFn1((v1 / (((v1 & int16(v5))) | 1)))
    v1 = (v1 + (v1 - (v1 - genFn1(-615))))
    if (false && v4) {
        var v24 string = (("hello" + "a") + "\n\t")
        var v25 uint8 = 234
        _ = v24
        _ = v25
    }
    var v26 int8 = 0
    fmt.Println(v2)
    v27 := GenS4{F0: v2, F1: ((("" + ("\n\t" + "\n\t")) + "foo bar") + "\n\t")}
    if v4 {
        var v28 uint32 = 2147483648
        var v29 uint64 = (v5 ^ 937)
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: v1, F1: uint(v5)}
    for fi31 := 0; fi31 < 18; fi31++ {
        v4 = false
        var v32 uint = 255
        _ = fi31
        _ = v32
    }
    if v4 {
        for fi33 := 0; fi33 < 11; fi33++ {
            fmt.Println(fi33, v26, v4)
            _ = fi33
        }
        if len(v6) > 0 {
            v6[0] = -2147483648
        }
        v34 := GenS5{F0: int(v2), F1: int8(v5), F2: 46}
        _ = v34
    } else {
        for fi35 := 0; fi35 < 3; fi35++ {
            fmt.Println(v5)
            var v36 float64 = 0.0
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 10; fi37++ {
            v1 = 1
            v2 -= (v2 / ((309) | 1))
            v38 := GenS1{F0: -366, F1: 2}
            _ = fi37
            _ = v38
        }
    }
    v39 := GenS5{F0: -2147483648, F1: (v26 ^ v26), F2: v26}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
