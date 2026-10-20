package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int }
func genFn0(p0 bool, p1 int8, p2 bool) int64 {
    switch "hello" {
    case "foo bar":
        v5 := []int{0, 4294967295, 9223372036854775807, 128}
        _ = v5
    case "a":
        v6 := []int{}
        _ = v6
    default:
        var v7 uint32 = 617
        _ = v7
    }
    fmt.Println(p0, p1)
    var v8 float64 = 1.5
    v9 := []int{32768}
    v10 := GenS1{F0: "", F1: -284}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint) uint32 {
    if (float64(1e10) > float64((-2.25 * 0.1))) {
        v11 := GenS1{F0: "a", F1: 9223372036854775807}
        _ = v11
    } else {
        var v12 uint8 = 128
        var v13 uint64 = 18446744073709551615
        v12 = (v12 + (^(v12)))
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        if (false && (0 != 887)) {
            p0 += (^(p0))
            for fi15 := 0; fi15 < 10; fi15++ {
                if false {
                    if (81 > -60) {
                        v16 := GenS1{F0: ("a" + ""), F1: 1}
                        _ = v16
                    } else {
                        switch (-(fi14)) {
                        case 128:
                            p0 = (p0 ^ 2)
                        default:
                            p0 = 42
                        }
                        fmt.Println(fi15, p0, fi15)
                        p0 -= p0
                    }
                } else {
                    p0 = 255
                    var v17 int8 = 41
                    _ = v17
                }
                p0 -= 894
                var v18 float32 = 0.0
                _ = fi15
                _ = v18
            }
            p0 = (^(p0))
        } else {
            v19 := []int{}
            _ = v19
        }
        _ = fi14
    }
    v20 := GenS1{F0: "hello", F1: 0}
    _ = p0
    _ = v20
    return 886
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    var v3 int = 42
    var v4 bool = true
    var v21 uint32 = 65535
    var v22 uint32 = 0
    switch "a" {
    case "\n\t":
        if (false || v4) {
            v23 := []int{v3, v3, v3}
            if ((128 > uint8(v21)) && v4) {
                switch v3 {
                case 32767:
                    v24 := GenS2{F0: 127}
                    _ = v24
                case 4294967295:
                    var v25 int16 = -1
                    _ = v25
                case -79:
                    for fi26 := 0; fi26 < 14; fi26++ {
                        var v27 float32 = 0.1
                        _ = fi26
                        _ = v27
                    }
                }
            } else {
                for fi28 := 0; fi28 < 7; fi28++ {
                    var v29 string = "a"
                    _ = fi28
                    _ = v29
                }
                if v4 {
                    if true {
                        v23 = append(v23, v3)
                        for fi30 := 0; fi30 < 8; fi30++ {
                            v21 -= (v21 >> ((1) & 63))
                            if (!(v4 || v4)) {
                                var v31 int16 = 57
                                _ = v31
                            } else {
                                v21 += v21
                                v4 = (565 < 42)
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                            }
                            v23 = append(v23, (v3 & (v3 - v3)))
                            _ = fi30
                        }
                    }
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    var v32 int = (v3 | 4294967295)
                    _ = v32
                } else {
                    var v33 int32 = 17
                    _ = v33
                }
                fmt.Println(v3, v3, v4)
            }
            _ = v23
        } else {
            v34 := []int{(v3 % ((4294967295) | 1)), v3, (v3 >> ((0) & 63)), (v3 << ((598) & 63))}
            for fi35 := 0; fi35 < 4; fi35++ {
                v4 = (0 >= 234)
                var v36 int32 = 2147483647
                v37 := GenS1{F0: (("foo bar" + "a") + "\n\t"), F1: -9223372036854775808}
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = v34
        }
    default:
        if v4 {
            v38 := []int{(v3 >> ((42) & 63)), 750, v3, -1}
            v39 := GenS1{F0: "\n\t", F1: -9223372036854775808}
            v4 = v4
            _ = v38
            _ = v39
        } else {
            v1[v22] = (^(v22))
        }
    }
    v21 = v21
    var v40 int64 = genFn0(v4, 32, v4)
    v4 = (42 == 568)
    var v41 uint = 2
    switch "\n\t" {
    case "":
        v42 := GenS2{F0: (v3 | v3)}
        _ = v42
    case "a":
        v41 = v41
    case "hello":
        switch "hello" {
        case "":
            v41 -= 65535
        }
    }
    var v43 int32 = -22
    v43 += (v43 / (((v43 << (((^(v41))) & 63))) | 1))
    v40 -= v40
    for fi44 := 0; fi44 < 9; fi44++ {
        fmt.Println(v40)
        var v45 float32 = 0.0
        _ = fi44
        _ = v45
    }
    v21 = (^(v22))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v43
}
