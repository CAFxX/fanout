package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 string; F2 uint16 }
func genFn0(p0 uint) int32 {
    fmt.Println(p0)
    if false {
        var v5 int8 = -128
        var v6 uint = (p0 * (p0 & p0))
        if false {
            if (false && ((v5 / (((v5 - (v5 + v5))) | 1)) <= v5)) {
                var v7 uint32 = 42
                v7 -= (v7 & (v7 * v7))
                _ = v7
            }
            for fi8 := 0; fi8 < 3; fi8++ {
                v5 = -80
                for fi9 := 0; fi9 < 10; fi9++ {
                    v6 = (p0 - (^(p0)))
                    _ = fi9
                }
                var v10 int32 = 2147483647
                _ = fi8
                _ = v10
            }
            var v11 float32 = (0.0 / 1.5)
            _ = v11
        } else {
            fmt.Println(v6, p0)
        }
        _ = v5
        _ = v6
    }
    v12 := GenS1{F0: 100, F1: 32767, F2: (-483 >= 2147483647)}
    var v13 uint = (p0 << ((4294967295) & 63))
    _ = p0
    _ = v12
    _ = v13
    return -1
}
func genFn1() int {
    if (!true) {
        var v14 string = "a"
        _ = v14
    }
    var v15 int = 2147483647
    v16 := []int{127, v15, v15, v15}
    if false {
        if len(v16) > 0 {
            v16[0] = (v15 >> ((2) & 63))
        }
    } else {
        switch v15 {
        case 32767:
            v17 := GenS1{F0: -74, F1: 727, F2: true}
            _ = v17
        case -1:
            var v18 uint16 = 0
            _ = v18
        case -785:
            var v19 uint16 = 32767
            _ = v19
        }
    }
    _ = v15
    _ = v16
    return v15
}

func main() {
    var v1 bool = (!true)
    var v2 uint16 = 2
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    v20 := GenS2{F0: 127, F1: 438}
    var v21 int32 = genFn0(65535)
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            v21 -= v21
            v24 := []int{-9223372036854775808, -668, fi23, fi22}
            var v25 uint32 = 208
            _ = fi23
            _ = v24
            _ = v25
        }
        v26 := GenS2{F0: 175, F1: 995}
        _ = fi22
        _ = v26
    }
    var v27 int8 = -1
    for fi28 := 0; fi28 < 3; fi28++ {
        if true {
            switch genFn1() {
            case 2147483647:
                var v29 uint32 = 812
                _ = v29
            case -9223372036854775808:
                v30 := GenS3{F0: 2}
                _ = v30
            case -953:
                for fi31 := 0; fi31 < 2; fi31++ {
                    v2 -= v2
                    if true {
                        var v32 float64 = 1.5
                        _ = v32
                    }
                    if v1 {
                        v33 := v3[("hello" + "foo bar")]
                        var v34 int8 = (v27 >> ((180) & 63))
                        v35 := GenS1{F0: (^(v34)), F1: 32767, F2: (1 < genFn0(65535))}
                        _ = v33
                        _ = v34
                        _ = v35
                    } else {
                        v21 += 0
                    }
                    v21 = 0
                    v27 -= v27
                    _ = fi31
                }
            default:
                var v36 int32 = int32(fi28)
                _ = v36
            }
        }
        _ = fi28
    }
    var v37 int16 = 32767
    v1 = ((^(v21)) >= v21)
    var v38 uint = 4294967295
    if true {
        fmt.Println(v38)
    } else {
        v39 := GenS4{F0: "a", F1: ("\n\t" + (((("foo bar" + "hello") + "foo bar") + "") + ("" + "foo bar"))), F2: (v2 ^ v2)}
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 int64 = -203
            _ = fi40
            _ = v41
        }
        _ = v39
    }
    for fi42 := 0; fi42 < 12; fi42++ {
        v38 += 0
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
