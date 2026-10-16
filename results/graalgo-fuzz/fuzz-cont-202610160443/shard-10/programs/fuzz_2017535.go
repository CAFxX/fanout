package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint, p1 int8, p2 int32) uint {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: true}
        if false {
            if (true && false) {
                v8 := GenS2{F0: 743, F1: -9223372036854775808}
                for fi9 := 0; fi9 < 18; fi9++ {
                    switch ("" + ((("foo bar" + "foo bar") + "\n\t") + "hello")) {
                    case "foo bar":
                        p0 = p0
                    case "a":
                        p1 -= (p1 & p1)
                    default:
                        switch ("" + ("hello" + "hello")) {
                        case "a":
                            if false {
                                fmt.Println(fi6, p2, fi9)
                                var v10 string = ("foo bar" + ("hello" + "a"))
                                v11 := []int{(fi9 >> (((p0 % ((p0) | 1))) & 63))}
                                _ = v10
                                _ = v11
                            } else {
                                v12 := []int{1, fi9}
                                _ = v12
                            }
                        case "":
                            var v13 uint8 = 0
                            _ = v13
                        default:
                            var v14 uint32 = 2
                            _ = v14
                        }
                    }
                    p1 -= p1
                    v15 := []int{}
                    _ = fi9
                    _ = v15
                }
                _ = v8
            } else {
                var v16 int16 = 1000
                v17 := GenS1{F0: ("" <= ((("hello" + ("hello" + "foo bar")) + "") + "foo bar"))}
                _ = v16
                _ = v17
            }
            switch "" {
            case "":
                fmt.Println(p0, p0)
            case "hello":
                p2 = p2
            default:
                var v18 int16 = 1
                _ = v18
            }
        } else {
            switch 4294967295 {
            case 32768:
                fmt.Println(p0)
            default:
                var v19 float64 = 1e10
                _ = v19
            }
        }
        switch ("foo bar" + (("foo bar" + "") + "hello")) {
        case "":
            fmt.Println(fi6, fi6)
        case "hello":
            var v20 uint8 = 197
            _ = v20
        }
        _ = fi6
        _ = v7
    }
    switch 32767 {
    case 45:
        for fi21 := 0; fi21 < 15; fi21++ {
            if (!false) {
                for fi22 := 0; fi22 < 9; fi22++ {
                    if (!(0 < 2147483648)) {
                        var v23 int = 9223372036854775807
                        var v24 float32 = -2.25
                        _ = v23
                        _ = v24
                    } else {
                        if (((p0 >= (p0 | p0)) || false) || true) {
                            var v25 uint16 = 255
                            if true {
                                var v26 uint16 = v25
                                _ = v26
                            } else {
                                v27 := GenS1{F0: false}
                                _ = v27
                            }
                            p1 = p1
                            var v28 float64 = (3.14159 / 0.1)
                            _ = v25
                            _ = v28
                        }
                    }
                    p1 += -1
                    _ = fi22
                }
            }
            var v29 int32 = (p2 / (((p2 << (((p0 ^ p0)) & 63))) | 1))
            _ = fi21
            _ = v29
        }
    case 803:
        v30 := GenS2{F0: 65535, F1: 0}
        _ = v30
    case -106:
        v31 := []int{-194, 1}
        _ = v31
    default:
        p1 -= 127
    }
    var v32 uint16 = 2
    if false {
        var v33 int32 = (p2 << ((p0) & 63))
        switch "\n\t" {
        case "\n\t":
            for fi34 := 0; fi34 < 4; fi34++ {
                fmt.Println(p0, v33)
                for fi35 := 0; fi35 < 3; fi35++ {
                    v36 := []int{(fi34 * 32768), (fi35 ^ (fi34 / ((-976) | 1))), (fi35 / ((-9223372036854775808) | 1))}
                    v37 := GenS2{F0: v32, F1: -9223372036854775808}
                    var v38 float32 = 0.0
                    _ = fi35
                    _ = v36
                    _ = v37
                    _ = v38
                }
                v39 := GenS2{F0: 32767, F1: int64(p1)}
                _ = fi34
                _ = v39
            }
        case "a":
            v40 := GenS1{F0: (!(1 == 42))}
            _ = v40
        }
        _ = v33
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return (p0 % ((p0) | 1))
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 int8 = 100
    var v5 uint64 = 42
    fmt.Println(v4, v1)
    v1 -= v1
    v4 = 100
    v1 = float64(1e10)
    var v41 uint64 = (v5 - (^(v5)))
    v42 := GenS1{F0: true}
    v5 += v41
    if true {
        v1 = (v1 * float64(-2.25))
        var v43 uint64 = 427
        var v44 float32 = float32(v1)
        _ = v43
        _ = v44
    } else {
        v45 := []int{-2147483648}
        if false {
            var v46 int = -2147483648
            _ = v46
        } else {
            var v47 uint16 = 222
            v48 := GenS1{F0: true}
            _ = v47
            _ = v48
        }
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
