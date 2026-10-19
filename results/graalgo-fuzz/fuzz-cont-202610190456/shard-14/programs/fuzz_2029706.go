package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int, p1 uint32) int64 {
    v6 := GenS1{F0: (p1 % (((p1 - p1)) | 1)), F1: 0, F2: 255}
    if true {
        var v7 float32 = 1e10
        _ = v7
    } else {
        var v8 int32 = 1
        v9 := GenS1{F0: p1, F1: v8, F2: 1}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 uint, p1 uint32) uint32 {
    if (1 == 0) {
        var v10 float32 = float32(1e10)
        fmt.Println(v10, p1)
        _ = v10
    } else {
        if false {
            for fi11 := 0; fi11 < 6; fi11++ {
                p0 += 18446744073709551615
                var v12 uint64 = 139
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                for fi14 := 0; fi14 < 1; fi14++ {
                    switch 2147483647 {
                    case 2147483647:
                        v15 := GenS2{F0: (!false)}
                        _ = v15
                    case 0:
                        switch "hello" {
                        case "hello":
                            var v16 uint64 = 418
                            _ = v16
                        }
                    case 32768:
                        v17 := []int{2147483647, (fi13 + (fi13 ^ (fi14 >> ((4294967295) & 63)))), fi13}
                        _ = v17
                    }
                    _ = fi14
                }
                var v18 string = "\n\t"
                _ = fi13
                _ = v18
            }
        } else {
            v19 := []int{213}
            _ = v19
        }
        v20 := GenS2{F0: ((p0 << ((953) & 63)) < p0)}
        var v21 float32 = 0.0
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: (153 == 255)}
    if (42 != -67) {
        if true {
            var v23 uint = (p0 & p0)
            switch "hello" {
            case "foo bar":
                var v24 float32 = 0.1
                _ = v24
            default:
                p0 += (v23 >> (((v23 % ((p0) | 1))) & 63))
            }
            _ = v23
        }
    }
    v25 := []int{-686, 0, 1}
    var v26 int32 = -2147483648
    _ = p0
    _ = p1
    _ = v22
    _ = v25
    _ = v26
    return (p1 % (((^(p1))) | 1))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 uint32 = 770
    v4 := []uint32{42}
    v5 := make(map[int]int)
    v27 := GenS2{F0: (!true)}
    var v28 int32 = 0
    v1[127] = 4294967295
    v5[-345] = 9223372036854775807
    v28 += (v28 + (-(v28)))
    switch -9223372036854775808 {
    case -959:
        var v29 uint8 = 198
        _ = v29
    default:
        v28 -= (^(v28))
    }
    var v30 float64 = -2.25
    switch -1 {
    case -9223372036854775808:
        var v31 int8 = 1
        _ = v31
    default:
        for fi32 := 0; fi32 < 13; fi32++ {
            if (!false) {
                if false {
                    v33 := GenS3{F0: 42}
                    _ = v33
                }
                if false {
                    v34 := GenS2{F0: false}
                    fmt.Println(v30, fi32, v30)
                    switch 200 {
                    case -2147483648:
                        _ = len(v4)
                    case 9223372036854775807:
                        var v35 uint = 65535
                        _ = v35
                    default:
                        switch "a" {
                        case "\n\t":
                            switch fi32 {
                            case 912:
                                v3 -= (v3 & 65535)
                            case -337:
                                v36 := GenS4{F0: -898}
                                _ = v36
                            }
                        case "hello":
                            fmt.Println(fi32)
                        }
                    }
                    _ = v34
                } else {
                    var v37 int32 = int32(v3)
                    switch (-(fi32)) {
                    case 2147483647:
                        v38 := GenS4{F0: int64(v28)}
                        _ = v38
                    case 4294967295:
                        if (!(0 > 0)) {
                            fmt.Println(v37, v37)
                        } else {
                            fmt.Println(v30)
                            fmt.Println(v30)
                            switch fi32 {
                            case 128:
                                var v39 float64 = (-2.25 - v30)
                                _ = v39
                            case 830:
                                switch fi32 {
                                case 32767:
                                    var v40 uint64 = 1
                                    _ = v40
                                case 9223372036854775807:
                                    var v41 int16 = -377
                                    _ = v41
                                }
                            case -14:
                                switch 0 {
                                case 0:
                                    v28 = (v37 | (v37 >> ((1) & 63)))
                                case 275:
                                    fmt.Println(v3, v28)
                                case -1:
                                    fmt.Println(v28, v37, v3)
                                }
                            default:
                                v5[(fi32 * (fi32 / ((fi32) | 1)))] = fi32
                            }
                        }
                    default:
                        v3 = v3
                    }
                    v3 = 2
                    _ = v37
                }
            }
            _ = fi32
        }
    }
    switch ("" + "") {
    case "\n\t":
        v42 := GenS1{F0: v3, F1: v28, F2: 0}
        _ = v42
    default:
        _ = len(v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
