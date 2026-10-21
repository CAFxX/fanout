package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32, p1 uint8, p2 uint32) uint32 {
    switch (("a" + "") + ("" + "a")) {
    case "a":
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(p1, p0, p2)
            _ = fi6
        }
    case "foo bar":
        p1 = (p1 % ((p1) | 1))
    case "hello":
        v7 := []int{127, 1, -9223372036854775808}
        _ = v7
    }
    switch "" {
    case "":
        for fi8 := 0; fi8 < 8; fi8++ {
            switch (("hello" + "foo bar") + "") {
            case "":
                var v9 uint32 = 2
                _ = v9
            case "a":
                var v10 uint64 = 1
                _ = v10
            default:
                v11 := []int{32767}
                _ = v11
            }
            p1 = (p1 | 65)
            _ = fi8
        }
    case "\n\t":
        if false {
            var v12 float32 = (-2.25 / 0.1)
            var v13 int = 32767
            _ = v12
            _ = v13
        }
    case "hello":
        if true {
            var v14 float64 = 1e10
            v15 := GenS1{F0: true, F1: 42, F2: -1}
            var v16 bool = (!(!true))
            _ = v14
            _ = v15
            _ = v16
        }
    default:
        var v17 float32 = 0.0
        _ = v17
    }
    if true {
        var v18 float32 = 1.5
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 * 4294967295)
}
func genFn1(p0 int, p1 uint64, p2 int32) int {
    v19 := []int{p0, p0, (p0 % ((p0) | 1))}
    if true {
        var v20 float32 = 0.1
        _ = len(v19)
        for fi21 := 0; fi21 < 19; fi21++ {
            p1 -= p1
            fmt.Println(p1, v20)
            if false {
                if false {
                    v22 := GenS1{F0: true, F1: -18, F2: (p2 + 526)}
                    for fi23 := 0; fi23 < 8; fi23++ {
                        fmt.Println(p0)
                        if true {
                            var v24 int32 = 0
                            _ = v24
                        }
                        _ = fi23
                    }
                    for fi25 := 0; fi25 < 9; fi25++ {
                        p2 = (^(p2))
                        switch "hello" {
                        case "":
                            v19 = append(v19, fi21)
                        }
                        var v26 int = (p0 << ((4294967295) & 63))
                        _ = fi25
                        _ = v26
                    }
                    _ = v22
                } else {
                    if (993 >= 190) {
                        fmt.Println(p1, fi21)
                        var v27 int64 = -122
                        fmt.Println(v27, p0)
                        _ = v27
                    }
                    if (221 == p1) {
                        var v28 float64 = (1.5 * 0.1)
                        _ = v28
                    } else {
                        p2 -= (p2 / (((p2 / ((p2) | 1))) | 1))
                    }
                }
            } else {
                if false {
                    switch ("\n\t" + "\n\t") {
                    case "":
                        _ = len(v19)
                    default:
                        fmt.Println(p0)
                    }
                    switch "" {
                    case "a":
                        for fi29 := 0; fi29 < 15; fi29++ {
                            switch -868 {
                            case 2147483647:
                                fmt.Println(v20, fi21)
                            case -799:
                                var v30 int8 = -66
                                if false {
                                    p0 = (p0 * fi29)
                                    var v31 uint = 0
                                    _ = v31
                                }
                                _ = v30
                            case 42:
                                var v32 uint32 = 102
                                _ = v32
                            default:
                                if len(v19) > 0 {
                                    v19[0] = p0
                                }
                            }
                            v33 := GenS1{F0: true, F1: -32768, F2: 447}
                            _ = fi29
                            _ = v33
                        }
                    }
                    if (!false) {
                        var v34 int8 = -1
                        var v35 int64 = 500
                        _ = v34
                        _ = v35
                    }
                }
            }
            _ = fi21
        }
        _ = v20
    } else {
        _ = len(v19)
        for fi36 := 0; fi36 < 1; fi36++ {
            if (302 != 342) {
                switch (p0 * 127) {
                case 2147483647:
                    switch "a" {
                    case "foo bar":
                        p1 += (p1 | (p1 * p1))
                    case "\n\t":
                        v37 := GenS1{F0: (!true), F1: -387, F2: (^(p2))}
                        _ = v37
                    }
                }
                v38 := GenS1{F0: (fi36 >= fi36), F1: 0, F2: 359}
                _ = v38
            }
            _ = fi36
        }
    }
    p0 += 32767
    var v39 uint64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v39
    return (p0 - 32767)
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    var v5 float32 = float32(3.14159)
    v40 := GenS1{F0: false, F1: 1000, F2: 501}
    var v41 uint32 = 0
    fmt.Println(v5, v41)
    if len(v1) > 0 {
        v1[0] = v41
    }
    var v42 uint16 = 65535
    fmt.Println(v5, v41)
    v43 := v2[(^(v41))]
    v44 := GenS2{F0: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
