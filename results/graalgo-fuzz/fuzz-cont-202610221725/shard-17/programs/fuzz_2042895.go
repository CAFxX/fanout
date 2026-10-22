package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint8; F2 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int32, p1 uint) bool {
    v5 := []int{470}
    switch -156 {
    case 9223372036854775807:
        if (-32768 != 977) {
            v6 := GenS1{F0: 42, F1: true}
            switch 42 {
            case -2147483648:
                p0 += -1
            default:
                _ = len(v5)
            }
            _ = v6
        }
    }
    var v7 int64 = -795
    if len(v5) > 0 {
        v5[0] = 127
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return true
}
func genFn1(p0 int8, p1 int32, p2 uint64) uint {
    p0 = p0
    if false {
        var v8 float64 = 0.0
        if (true || true) {
            for fi9 := 0; fi9 < 13; fi9++ {
                p2 -= p2
                _ = fi9
            }
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 uint16 = 32767
                for fi12 := 0; fi12 < 12; fi12++ {
                    switch 127 {
                    case 370:
                        switch "a" {
                        case "hello":
                            v13 := []int{fi12, (-(fi12))}
                            _ = v13
                        default:
                            if false {
                                var v14 int16 = -955
                                _ = v14
                            } else {
                                var v15 float64 = v8
                                _ = v15
                            }
                        }
                    case 448:
                        p1 = p1
                    case 2147483647:
                        v16 := []int{}
                        _ = v16
                    }
                    _ = fi12
                }
                var v17 bool = (fi10 > 924)
                _ = fi10
                _ = v11
                _ = v17
            }
            for fi18 := 0; fi18 < 2; fi18++ {
                if true {
                    p2 = (p2 | p2)
                    fmt.Println(v8, p1)
                } else {
                    p0 -= -1
                    switch fi18 {
                    case 42:
                        switch "\n\t" {
                        case "a":
                            v19 := GenS2{F0: 98, F1: 4294967295, F2: 0}
                            _ = v19
                        case "\n\t":
                            p1 -= (p1 << ((801) & 63))
                        default:
                            v20 := GenS3{F0: 2, F1: 54, F2: ((fi18 * (fi18 & 0)) > (fi18 ^ fi18))}
                            _ = v20
                        }
                    case -314:
                        var v21 int8 = (p0 & (p0 * (p0 | p0)))
                        _ = v21
                    }
                    p1 = 236
                }
                _ = fi18
            }
        }
        _ = v8
    } else {
        switch "\n\t" {
        case "foo bar":
            if false {
                var v22 int = 251
                fmt.Println(p0, p2)
                _ = v22
            }
        default:
            var v23 uint = 550
            _ = v23
        }
        for fi24 := 0; fi24 < 14; fi24++ {
            v25 := []int{fi24, (-(fi24))}
            _ = fi24
            _ = v25
        }
    }
    if false {
        p0 = 100
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 uint64 = 2
    v2 := []int{-1}
    var v3 bool = false
    var v4 string = "a"
    switch "" {
    case "a":
        var v26 int32 = 42
        _ = v26
    case "foo bar":
        if (v1 >= v1) {
            if ((-2.25 != 0.1) || false) {
                v27 := GenS2{F0: 2, F1: 833, F2: 128}
                _ = v27
            } else {
                v4 = "a"
                var v28 int64 = 1
                _ = v28
            }
        } else {
            switch 42 {
            case -966:
                v3 = v3
            case 42:
                v29 := GenS4{F0: 2147483647}
                _ = v29
            case 128:
                v4 = v4
            default:
                var v30 uint8 = 1
                _ = v30
            }
            fmt.Println(v3, v4)
        }
    default:
        var v31 int16 = 32767
        _ = v31
    }
    var v32 bool = (("foo bar" + v4) == v4)
    if (!(128 != -605)) {
        fmt.Println(v1)
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        for fi34 := 0; fi34 < 12; fi34++ {
            if v32 {
                fmt.Println(v1)
            } else {
                v3 = (float64(1e10) < 0.1)
                if (-1 > -718) {
                    var v35 int32 = -2147483648
                    v36 := GenS5{F0: -7}
                    var v37 float32 = ((((1.5 * 3.14159) * 1e10) + (((0.0 * (3.14159 * 1.5)) * 0.1) + -2.25)) * float32(float64(3.14159)))
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
            fmt.Println(v1, fi34)
            _ = fi34
        }
        if (false || (!v3)) {
            fmt.Println(v3, v1, v3)
        } else {
            v32 = (255 != 221)
        }
        v1 = v1
        _ = fi33
    }
    v3 = (!(1 == -1))
    var v38 int64 = 9223372036854775807
    _ = len(v2)
    switch 32767 {
    case 262:
        v39 := GenS4{F0: 2}
        _ = v39
    default:
        v1 += (^(v1))
    }
    v4 = (v4 + v4)
    var v40 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v40
}
