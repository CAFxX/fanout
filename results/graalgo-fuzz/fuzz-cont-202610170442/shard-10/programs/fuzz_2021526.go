package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int; F1 int64; F2 int16 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 string }
func genFn0(p0 uint32, p1 uint32, p2 int8) int16 {
    var v6 int32 = -358
    v7 := []int{127, 4294967295, -352}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int8, p1 bool, p2 bool) int8 {
    fmt.Println(p1, p0, p0)
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            if (!(!true)) {
                if (!false) {
                    var v10 uint8 = 240
                    var v11 float32 = 0.0
                    v12 := []int{fi8, 32767, fi8, (fi8 - 32768)}
                    _ = v10
                    _ = v11
                    _ = v12
                }
                if p1 {
                    switch (fi9 / (((-(fi9))) | 1)) {
                    case -355:
                        fmt.Println(p0, fi9)
                        p2 = p1
                        var v13 float32 = 0.1
                        _ = v13
                    default:
                        v14 := GenS1{F0: (^(fi8))}
                        _ = v14
                    }
                    var v15 int32 = 2147483647
                    v15 -= 212
                    _ = v15
                } else {
                    switch ("" + "") {
                    case "\n\t":
                        p2 = false
                    default:
                        fmt.Println(fi9)
                    }
                }
                fmt.Println(fi8, fi9, p0)
            } else {
                v16 := []int{(fi9 % (((fi8 >> ((0) & 63))) | 1))}
                p1 = (!(0 <= 255))
                _ = v16
            }
            p2 = false
            _ = fi9
        }
        p2 = p1
        var v17 string = "a"
        _ = fi8
        _ = v17
    }
    p0 -= p0
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 uint64 = 270
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{487, 619}
    v3 := make(map[int]bool)
    var v4 int64 = -1
    var v5 int = -876
    v4 = v4
    var v20 uint32 = 432
    for fi21 := 0; fi21 < 2; fi21++ {
        delete(v3, (v5 >> ((2) & 63)))
        _ = fi21
    }
    v20 = (^(v20))
    var v22 int64 = v4
    v23 := GenS1{F0: -104}
    v22 += (v4 << ((1) & 63))
    if (!(!false)) {
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := GenS2{F0: fi24, F1: 42}
            var v26 int8 = 123
            var v27 float32 = 0.0
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 20; fi28++ {
            v29 := GenS3{F0: (fi28 & (^(v5))), F1: 42, F2: 42}
            _ = fi28
            _ = v29
        }
        switch "foo bar" {
        case "\n\t":
            v30 := GenS4{F0: 1, F1: 2147483648}
            _ = v30
        default:
            var v31 float32 = 1.5
            _ = v31
        }
    } else {
        if (!true) {
            var v32 int16 = 0
            _ = v32
        }
    }
    if false {
        v33 := GenS5{F0: v20, F1: 62}
        var v34 int8 = 100
        v35 := GenS6{F0: "foo bar"}
        _ = v33
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        delete(v1, "\n\t")
        if (18446744073709551615 != 42) {
            v4 -= -1
            _ = len(v2)
        }
        _ = fi36
    }
    var v37 float32 = 1.5
    var v38 int64 = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
