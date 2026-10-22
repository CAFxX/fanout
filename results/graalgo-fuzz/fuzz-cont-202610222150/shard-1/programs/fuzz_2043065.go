package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int, p1 int, p2 uint16) uint {
    var v5 uint8 = 242
    v6 := []int{(^(p0))}
    var v7 int = -2147483648
    if (!false) {
        p0 += p0
        var v8 uint8 = (v5 * 0)
        p0 = p0
        _ = v8
    }
    v5 += (v5 * (v5 / ((53) | 1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 284
}
func genFn1() uint {
    var v9 int8 = -128
    switch ("" + "a") {
    case "":
        v10 := GenS1{F0: 42, F1: 2147483647, F2: -1}
        _ = v10
    case "a":
        if false {
            if true {
                for fi11 := 0; fi11 < 7; fi11++ {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        switch "" {
                        case "hello":
                            switch ("foo bar" + "a") {
                            case "a":
                                if true {
                                    fmt.Println(v9)
                                } else {
                                    switch "" {
                                    case "hello":
                                        if false {
                                            v9 += v9
                                            v9 += v9
                                            var v13 int32 = 0
                                            v14 := []int{(fi12 * 32768)}
                                            var v15 int16 = 1
                                            _ = v13
                                            _ = v14
                                            _ = v15
                                        }
                                    case "\n\t":
                                        switch ("foo bar" + "") {
                                        case "":
                                            v16 := GenS2{F0: 42}
                                            _ = v16
                                        case "foo bar":
                                            v17 := []int{(fi12 % ((4294967295) | 1)), 32767, -1, (fi12 | (fi12 << ((2) & 63)))}
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                            _ = v17
                                        default:
                                            v18 := []int{fi12, 791}
                                            _ = v18
                                        }
                                    default:
                                        v9 = v9
                                    }
                                    var v19 int16 = -545
                                    _ = v19
                                }
                            default:
                                var v20 uint = 696
                                _ = v20
                            }
                        case "a":
                            switch "foo bar" {
                            case "":
                                if (((fi11 - (fi11 ^ 0)) == fi12) || (541 >= 4294967295)) {
                                    v21 := GenS2{F0: -579}
                                    _ = v21
                                }
                            case "hello":
                                v9 += v9
                            case "a":
                                fmt.Println(fi11, fi12)
                            }
                        }
                        _ = fi12
                    }
                    v9 -= (v9 | (v9 & v9))
                    _ = fi11
                }
                switch 0 {
                case 673:
                    var v22 float32 = float32(0.0)
                    _ = v22
                }
            } else {
                v9 = v9
            }
        } else {
            var v23 uint32 = 1
            var v24 uint64 = 4294967295
            _ = v23
            _ = v24
        }
    case "foo bar":
        if true {
            v25 := GenS2{F0: -1}
            for fi26 := 0; fi26 < 12; fi26++ {
                v9 = 0
                for fi27 := 0; fi27 < 8; fi27++ {
                    fmt.Println(v9, fi26, fi26)
                    var v28 uint32 = 354
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
            _ = v25
        } else {
            fmt.Println(v9)
        }
    default:
        if ((4294967295 == 2) && true) {
            v29 := GenS2{F0: 2147483647}
            var v30 int8 = v9
            v31 := []int{}
            _ = v29
            _ = v30
            _ = v31
        }
    }
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = (3.14159 - (1.5 - (0.0 / ((float32(3.14159) / -2.25) + 0.1))))
    var v3 uint8 = 182
    v4 := make(map[int]uint32)
    for fi32 := 0; fi32 < 12; fi32++ {
        v33 := []int{fi32}
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        v1[(fi34 * (fi34 % ((fi34) | 1)))] = 4294967295
        _ = fi34
    }
    fmt.Println(v3, v3)
    if false {
        v3 += v3
        var v35 int8 = 1
        v36 := GenS3{F0: 65535, F1: 1}
        _ = v35
        _ = v36
    } else {
        fmt.Println(v2, v3)
    }
    delete(v1, 42)
    v37 := []int{-496, 795, 788}
    switch "" {
    case "foo bar":
        var v38 int32 = 2147483647
        _ = v38
    case "\n\t":
        v2 = -2.25
    case "a":
        switch "a" {
        case "hello":
            if ((-1 >= -9223372036854775808) || false) {
                var v39 int64 = 599
                if len(v37) > 0 {
                    v37[0] = 960
                }
                _ = v39
            }
        case "":
            var v40 int8 = 42
            _ = v40
        }
    }
    fmt.Println(v2, v3)
    var v41 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
}
