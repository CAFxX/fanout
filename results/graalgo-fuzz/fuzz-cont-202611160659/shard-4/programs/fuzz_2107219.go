package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }
type GenS2 struct { F0 int64; F1 bool; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int32; F2 int8 }
type GenS5 struct { F0 string; F1 uint8 }
func genFn0(p0 uint64, p1 string) uint64 {
    v5 := []int{2147483647, 416}
    var v6 int32 = 762
    v7 := GenS1{F0: 235, F1: 4294967295, F2: "\n\t"}
    for fi8 := 0; fi8 < 2; fi8++ {
        if len(v5) > 0 {
            v5[0] = fi8
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 uint16, p1 bool, p2 int16) uint8 {
    var v9 int32 = 2147483647
    p2 += p2
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 uint16 = (p0 * p0)
            v13 := GenS2{F0: 1, F1: (!(2 > 0)), F2: "foo bar"}
            switch fi11 {
            case -854:
                var v14 int16 = p2
                _ = v14
            case 289:
                var v15 int16 = p2
                _ = v15
            case -943:
                if (p1 && false) {
                    v16 := []int{9223372036854775807}
                    p2 += (^(p2))
                    v17 := GenS3{F0: 2}
                    _ = v16
                    _ = v17
                }
            default:
                var v18 string = ""
                _ = v18
            }
            var v19 uint64 = 163
            var v20 int32 = (-(v9))
            _ = fi11
            _ = v12
            _ = v13
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            p2 = (p2 * (p2 - 750))
            _ = fi21
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 int8 = -86
            v24 := []int{fi22}
            var v25 uint8 = 1
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        _ = fi10
    }
    v26 := GenS3{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v9
    return 143
}

func main() {
    var v1 bool = (32767 != -578)
    var v2 int32 = -1
    v3 := []uint32{680, 2147483648}
    var v4 uint32 = 65535
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "a" {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v27 int16 = 1000
        _ = v27
    }
    v28 := GenS1{F0: 32767, F1: (^(v4)), F2: ("foo bar" + "a")}
    var v29 uint64 = 1
    var v30 int = 9223372036854775807
    switch "hello" {
    case "a":
        if v1 {
            v31 := GenS3{F0: 395}
            _ = v31
        }
    default:
        v32 := GenS4{F0: 128, F1: v2, F2: 0}
        _ = v32
    }
    if false {
        for fi33 := 0; fi33 < 7; fi33++ {
            if false {
                v29 -= (v29 * (v29 / ((v29) | 1)))
                var v34 float32 = 1e10
                _ = v34
            }
            v35 := GenS5{F0: "foo bar", F1: genFn1(32767, v1, -1)}
            if v1 {
                switch "\n\t" {
                case "":
                    if v1 {
                        v30 = v30
                    }
                default:
                    var v36 uint64 = v29
                    _ = v36
                }
            } else {
                fmt.Println(v4)
                v4 = (v4 >> ((194) & 63))
                var v37 int32 = (v2 + (v2 & -1))
                _ = v37
            }
            _ = fi33
            _ = v35
        }
    } else {
        for fi38 := 0; fi38 < 17; fi38++ {
            v3 = append(v3, v4)
            _ = fi38
        }
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
