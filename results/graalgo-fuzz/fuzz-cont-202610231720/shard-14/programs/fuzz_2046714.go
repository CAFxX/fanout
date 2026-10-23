package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8; F1 int16; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int; F1 int16; F2 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16; F1 uint; F2 string }
type GenS7 struct { F0 bool }
func genFn0(p0 string, p1 string) int32 {
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(p0, p1)
        fmt.Println(fi6, p1, p1)
        _ = fi6
    }
    var v7 float64 = ((0.1 * (1.5 - 1.5)) / 1.0)
    if true {
        fmt.Println(p1, p1)
        var v8 uint = 0
        var v9 uint32 = 1
        _ = v8
        _ = v9
    }
    if (1 <= 65535) {
        var v10 float64 = v7
        p0 = p0
        var v11 float32 = (0.1 + (0.0 - 1e10))
        _ = v10
        _ = v11
    } else {
        p1 = p1
    }
    p0 = "foo bar"
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 uint16, p1 int, p2 bool) int16 {
    fmt.Println(p2)
    fmt.Println(p1)
    var v12 uint16 = 1
    v13 := GenS1{F0: p1, F1: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 205
}

func main() {
    var v1 uint16 = 827
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 uint = 1
    v5 := make(map[uint32]bool)
    var v14 uint16 = (^(v1))
    v15 := GenS2{F0: -1, F1: 548, F2: uint32(v4)}
    v16 := []int{-9223372036854775808, -772, 32767, 1}
    v17 := GenS3{F0: 100, F1: (^(v1))}
    v4 += v4
    fmt.Println(v1, v14)
    v18 := GenS4{F0: 567, F1: 0, F2: (v4 - v4)}
    switch -2147483648 {
    case 128:
        if (2 >= 356) {
            if ((true || false) && false) {
                fmt.Println(v14, v1, v4)
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 int8 = -128
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v4 = v4
                        fmt.Println(v4)
                        _ = fi21
                    }
                    if false {
                        v20 = (-(v20))
                    } else {
                        var v22 string = "hello"
                        var v23 int64 = 9223372036854775807
                        if true {
                            v22 = ((("hello" + v22) + ((v22 + ((("\n\t" + v22) + v22) + (("\n\t" + "a") + v22))) + v22)) + "foo bar")
                            v24 := GenS5{F0: v1}
                            _ = v24
                        }
                        _ = v22
                        _ = v23
                    }
                    _ = fi19
                    _ = v20
                }
                if (42 < 224) {
                    if true {
                        if false {
                            for fi25 := 0; fi25 < 1; fi25++ {
                                v5[923] = false
                                _ = fi25
                            }
                            switch -264 {
                            case 2147483647:
                                v14 -= v14
                            case -1:
                                var v26 uint64 = 4294967295
                                _ = v26
                            default:
                                v16 = append(v16, 22)
                            }
                        }
                        fmt.Println(v14)
                    } else {
                        fmt.Println(v14, v4)
                        if true {
                            if true {
                                v27 := v2[1]
                                v28 := GenS3{F0: 1, F1: v14}
                                v29 := GenS6{F0: 255, F1: 18446744073709551615, F2: ""}
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            var v30 int64 = -591
                            _ = v30
                        }
                    }
                    var v31 uint = v4
                    var v32 int = 328
                    _ = v31
                    _ = v32
                } else {
                    v33 := GenS4{F0: -2147483648, F1: genFn1((^(v1)), -510, true), F2: 65535}
                    _ = v33
                }
            } else {
                var v34 uint = 866
                v1 = (^(v1))
                _ = v34
            }
            if false {
                if true {
                    var v35 string = ("\n\t" + "")
                    _ = v35
                }
                v14 -= v14
            }
            var v36 uint16 = v14
            _ = v36
        }
    default:
        var v37 uint8 = 55
        _ = v37
    }
    if len(v16) > 0 {
        v16[0] = 655
    }
    v38 := GenS7{F0: true}
    v39 := GenS5{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
