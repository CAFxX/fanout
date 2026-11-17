package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int8, p1 uint) uint64 {
    if (42 == 42) {
        p0 = p0
        for fi7 := 0; fi7 < 4; fi7++ {
            switch ((("" + "") + (("a" + "") + "")) + "a") {
            case "a":
                for fi8 := 0; fi8 < 8; fi8++ {
                    v9 := []int{fi8, 0, (-(fi8)), (fi8 - 1)}
                    if true {
                        p0 = p0
                        var v10 bool = false
                        if v10 {
                            var v11 uint = (p1 * 995)
                            _ = v11
                        } else {
                            v10 = v10
                            var v12 uint8 = 127
                            _ = v12
                        }
                        _ = v10
                    }
                    var v13 uint64 = 441
                    _ = fi8
                    _ = v13
                    _ = v9
                }
            case "":
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := []int{(^(fi7)), int(p0), (fi14 + fi7), (fi14 >> (((p1 + 1)) & 63))}
                    fmt.Println(p0, fi14)
                    _ = fi14
                    _ = v15
                }
            case "foo bar":
                if (!false) {
                    var v16 float64 = 0.1
                    fmt.Println(p0)
                    if (-1 != 42) {
                        var v17 int16 = 0
                        _ = v17
                    }
                    _ = v16
                } else {
                    v18 := []int{}
                    if (!false) {
                        p1 -= (p1 << ((p1) & 63))
                        var v19 int8 = p0
                        for fi20 := 0; fi20 < 4; fi20++ {
                            p0 = (p0 * v19)
                            _ = fi20
                        }
                        _ = v19
                    } else {
                        var v21 uint = p1
                        _ = v21
                    }
                    var v22 uint16 = 65535
                    _ = v18
                    _ = v22
                }
            default:
                p0 += p0
            }
            p0 = (p0 | p0)
            _ = fi7
        }
        for fi23 := 0; fi23 < 11; fi23++ {
            p0 += p0
            var v24 uint8 = 0
            v25 := GenS1{F0: 100, F1: -254}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    v26 := GenS2{F0: 1}
    var v27 uint8 = 128
    if false {
        p0 = -101
        var v28 uint64 = 4294967295
        _ = v28
    } else {
        fmt.Println(p0, p1)
        var v29 int16 = 669
        _ = v29
    }
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 4294967295
}
func genFn1(p0 bool, p1 uint32, p2 int32) int {
    v30 := []int{-1, 42}
    switch -66 {
    case 4294967295:
        v30 = append(v30, 1)
    case 32768:
        v31 := GenS2{F0: 1}
        _ = v31
    case 127:
        var v32 string = "foo bar"
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 127
}

func main() {
    var v1 int8 = -44
    var v2 int8 = v1
    v3 := make(map[int]bool)
    var v4 uint16 = 255
    v5 := make(map[string]uint32)
    v6 := []int{629, 1, 873}
    v33 := GenS3{F0: -502, F1: 971}
    v6 = append(v6, 958)
    var v34 uint8 = 128
    var v35 int32 = 2147483647
    var v36 int8 = (v1 ^ v1)
    v37 := GenS4{F0: 4294967295}
    for fi38 := 0; fi38 < 5; fi38++ {
        for fi39 := 0; fi39 < 19; fi39++ {
            v6 = append(v6, genFn1((((v1 % (((v1 >> ((65535) & 63))) | 1)) <= (v1 ^ v36)) || true), 42, 2147483647))
            var v40 bool = (true || (!((0.1 + 1e10) >= ((1e10 / 3.14159) + (0.1 / 1.0)))))
            v41 := v3[fi39]
            _ = fi39
            _ = v40
            _ = v41
        }
        v6 = append(v6, (fi38 * fi38))
        _ = fi38
    }
    v42 := GenS3{F0: 950, F1: 9223372036854775807}
    switch "\n\t" {
    case "\n\t":
        v36 += v2
    case "a":
        v34 -= (v34 * (v34 - 0))
    }
    fmt.Println(v35)
    var v43 float32 = 0.0
    var v44 int16 = -32768
    v43 += (v43 / v43)
    var v45 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
