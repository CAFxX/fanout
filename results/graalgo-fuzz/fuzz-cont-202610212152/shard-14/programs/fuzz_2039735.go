package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int64; F1 int64; F2 uint16 }
type GenS5 struct { F0 int8 }
func genFn0() int16 {
    var v4 uint8 = 1
    for fi5 := 0; fi5 < 3; fi5++ {
        v4 = 2
        _ = fi5
    }
    _ = v4
    return -1
}
func genFn1(p0 uint64, p1 int64, p2 int64) int32 {
    switch "foo bar" {
    case "":
        for fi6 := 0; fi6 < 13; fi6++ {
            if false {
                p0 -= (p0 - 1)
                fmt.Println(p0)
            } else {
                switch fi6 {
                case 42:
                    v7 := []int{4294967295, 42, fi6}
                    _ = v7
                case 127:
                    var v8 int16 = 42
                    _ = v8
                default:
                    var v9 uint64 = 1
                    _ = v9
                }
                p0 -= (p0 | (p0 % ((p0) | 1)))
            }
            v10 := GenS1{F0: 2, F1: 42, F2: 423}
            v11 := GenS1{F0: 255, F1: 127, F2: 42}
            _ = fi6
            _ = v10
            _ = v11
        }
    case "\n\t":
        if true {
            var v12 int16 = 32767
            fmt.Println(p1, v12, p1)
            v12 = v12
            _ = v12
        }
    default:
        var v13 float64 = 0.1
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := []int{933, fi14}
        var v16 int8 = 0
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 uint = 0
        _ = fi17
        _ = v18
    }
    v19 := []int{0}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return -1
}

func main() {
    var v1 float64 = (0.1 * 1.5)
    var v2 uint16 = 0
    var v3 int16 = 42
    fmt.Println(v2, v3, v3)
    v20 := []int{}
    var v21 uint = 2
    v3 -= 42
    v1 = v1
    var v22 uint16 = (v2 ^ v2)
    if (829 > -868) {
        if (false || (false && false)) {
            v23 := GenS1{F0: v22, F1: 0, F2: uint32(v21)}
            _ = v23
        } else {
            v1 += float64(float32(v1))
        }
    } else {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        if true {
            v24 := GenS2{F0: v3, F1: 133, F2: 18446744073709551615}
            v22 += (v2 - 393)
            _ = v24
        } else {
            v25 := GenS1{F0: 0, F1: 82, F2: 509}
            switch -7 {
            case -9223372036854775808:
                v2 = v22
            case 1:
                for fi26 := 0; fi26 < 5; fi26++ {
                    _ = len(v20)
                    _ = fi26
                }
            case 2147483647:
                var v27 float32 = 0.1
                _ = v27
            }
            fmt.Println(v1, v3)
            _ = v25
        }
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        v3 -= v3
        if false {
            if (!(!false)) {
                var v29 int32 = 1
                _ = len(v20)
                for fi30 := 0; fi30 < 13; fi30++ {
                    v21 = (v21 ^ 180)
                    v21 = v21
                    _ = fi30
                }
                _ = v29
            }
            v20 = append(v20, 128)
        }
        for fi31 := 0; fi31 < 5; fi31++ {
            if false {
                switch fi28 {
                case 32767:
                    v32 := GenS3{F0: (3.14159 == v1), F1: 4294967295}
                    _ = v32
                case 9223372036854775807:
                    var v33 int16 = genFn0()
                    _ = v33
                case 4294967295:
                    fmt.Println(fi28, v21, fi28)
                default:
                    switch (("hello" + (("foo bar" + "") + "\n\t")) + ("\n\t" + "")) {
                    case "a":
                        var v34 uint32 = 430
                        _ = v34
                    case "":
                        v35 := GenS4{F0: -476, F1: -1, F2: (v22 % (((v22 & (v2 & v22))) | 1))}
                        _ = v35
                    case "\n\t":
                        v36 := GenS4{F0: int64(fi31), F1: 10, F2: (v22 / ((88) | 1))}
                        v2 += 42
                        _ = v36
                    }
                }
            }
            var v37 int8 = 5
            fmt.Println(v3, v2)
            _ = fi31
            _ = v37
        }
        _ = fi28
    }
    v22 = 65535
    switch 32768 {
    case 32767:
        if false {
            var v38 float32 = 0.1
            fmt.Println(v38, v3)
            for fi39 := 0; fi39 < 19; fi39++ {
                fmt.Println(v21, v22, v3)
                if true {
                    v40 := GenS4{F0: -1, F1: 115, F2: 1}
                    v20 = append(v20, fi39)
                    var v41 float32 = (v38 * 0.1)
                    _ = v40
                    _ = v41
                } else {
                    v42 := GenS1{F0: v22, F1: 2, F2: 42}
                    v3 -= genFn0()
                    fmt.Println(v22, v38)
                    _ = v42
                }
                v43 := GenS1{F0: 452, F1: 128, F2: 4294967295}
                _ = fi39
                _ = v43
            }
            _ = v38
        }
    }
    for fi44 := 0; fi44 < 4; fi44++ {
        var v45 uint32 = 1
        for fi46 := 0; fi46 < 2; fi46++ {
            switch fi44 {
            case 9223372036854775807:
                fmt.Println(fi44, fi46)
            }
            _ = fi46
        }
        _ = fi44
        _ = v45
    }
    v1 -= v1
    fmt.Println(v3, v21)
    v47 := GenS4{F0: 761, F1: 39, F2: v22}
    v48 := GenS5{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v47
    _ = v48
}
