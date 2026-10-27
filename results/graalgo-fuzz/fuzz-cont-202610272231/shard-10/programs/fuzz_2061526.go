package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 bool, p1 uint8) int64 {
    p1 = 182
    var v7 int64 = -815
    switch ("foo bar" + "a") {
    case "a":
        for fi8 := 0; fi8 < 15; fi8++ {
            p0 = ((p0 || p0) && false)
            if (("foo bar" == "\n\t") || (1.5 == 1.5)) {
                fmt.Println(fi8)
            } else {
                if p0 {
                    v9 := []int{(^(fi8))}
                    _ = v9
                } else {
                    if p0 {
                        var v10 uint = 508
                        _ = v10
                    } else {
                        p1 = 54
                        var v11 float64 = 0.1
                        _ = v11
                    }
                    v12 := GenS1{F0: 976, F1: 42, F2: 284}
                    _ = v12
                }
                var v13 uint32 = 984
                _ = v13
            }
            _ = fi8
        }
    case "\n\t":
        var v14 uint32 = 4294967295
        _ = v14
    }
    v15 := GenS2{F0: v7}
    _ = p0
    _ = p1
    _ = v15
    _ = v7
    return v7
}
func genFn1(p0 uint32, p1 int32, p2 int) uint16 {
    var v16 uint32 = p0
    v17 := []int{(^(p2)), (p2 / (((p2 | (^(p2)))) | 1))}
    var v18 int = (^(p2))
    var v19 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 128
    var v3 int64 = -9223372036854775808
    v4 := make(map[string]int)
    var v5 int = 0
    v6 := make(map[uint32]uint32)
    v2 = (v2 % (((v2 << ((769) & 63))) | 1))
    for fi20 := 0; fi20 < 15; fi20++ {
        v3 = 1
        fmt.Println(v5, fi20, v5)
        _ = fi20
    }
    var v21 uint32 = 2147483647
    v1[("a" + (("foo bar" + "hello") + ("\n\t" + "\n\t")))] = (v5 << ((255) & 63))
    if ((v2 | v2) >= (v2 * (v2 >> ((4294967295) & 63)))) {
        var v22 uint = 42
        fmt.Println(v3, v22, v22)
        var v23 uint64 = 18446744073709551615
        _ = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 uint16 = 1
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 float64 = -2.25
                _ = fi26
                _ = v27
            }
            if true {
                v28 := GenS2{F0: (v3 >> ((255) & 63))}
                for fi29 := 0; fi29 < 5; fi29++ {
                    var v30 int16 = int16(v21)
                    switch ("" + "foo bar") {
                    case "a":
                        var v31 bool = false
                        _ = v31
                    }
                    var v32 uint64 = 286
                    _ = fi29
                    _ = v30
                    _ = v32
                }
                v5 = int(v3)
                _ = v28
            }
            _ = fi24
            _ = v25
        }
        fmt.Println(v21, v5)
        _ = len(v6)
    }
    fmt.Println(v3, v21, v5)
    var v33 string = ""
    v34 := GenS3{F0: false, F1: -11}
    var v35 uint32 = (v21 << ((2) & 63))
    v36 := GenS4{F0: ("" + "\n\t"), F1: 394}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
