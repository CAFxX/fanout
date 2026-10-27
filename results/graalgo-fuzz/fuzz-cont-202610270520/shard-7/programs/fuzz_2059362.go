package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 string, p1 int32, p2 int32) uint {
    switch p0 {
    case "":
        v6 := []int{-348, -708, -486}
        _ = v6
    case "hello":
        v7 := GenS1{F0: 1}
        _ = v7
    case "\n\t":
        var v8 uint = 18446744073709551615
        _ = v8
    }
    var v9 int8 = 10
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2
}
func genFn1(p0 int16, p1 int16, p2 int64) uint64 {
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint64 = 18446744073709551615
        v12 := []int{}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return v13
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 32767
    var v3 uint64 = 0
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int64 = 42
        var v16 uint8 = 2
        _ = fi14
        _ = v15
        _ = v16
    }
    if true {
        fmt.Println(v3, v3)
    } else {
        v3 = (v3 << ((1) & 63))
        delete(v4, "a")
    }
    if (-32768 >= (v2 >> ((0) & 63))) {
        v17 := []int{}
        fmt.Println(v3, v2)
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 int64 = 1
        delete(v5, ((("\n\t" + "\n\t") + "hello") + "\n\t"))
        switch "hello" {
        case "\n\t":
            v20 := v1[(fi18 | fi18)]
            _ = v20
        default:
            for fi21 := 0; fi21 < 7; fi21++ {
                v22 := GenS1{F0: (-(v2))}
                if (!(v2 < (^(v2)))) {
                    if true {
                        var v23 float64 = 1e10
                        v2 = 484
                        _ = v23
                    }
                    var v24 int32 = 0
                    switch fi18 {
                    case -491:
                        if (88 == -128) {
                            v25 := []int{(-(fi18))}
                            switch "\n\t" {
                            case "foo bar":
                                var v26 uint64 = (v3 + genFn1((-(v2)), int16(v24), 838))
                                _ = v26
                            case "\n\t":
                                v19 = (v19 & v19)
                            case "":
                                v3 -= v3
                            default:
                                v3 = (v3 << ((42) & 63))
                            }
                            _ = v25
                        }
                    case -980:
                        v27 := GenS1{F0: v2}
                        _ = v27
                    case -651:
                        var v28 float32 = (1e10 + 3.14159)
                        _ = v28
                    default:
                        var v29 int64 = v19
                        _ = v29
                    }
                    _ = v24
                }
                fmt.Println(v3)
                v3 = 4294967295
                v5[("" + "hello")] = fi18
                _ = fi21
                _ = v22
            }
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v3, v2)
    v2 = 0
    v30 := GenS1{F0: 786}
    switch 32768 {
    case 128:
        for fi31 := 0; fi31 < 16; fi31++ {
            v3 = (v3 | v3)
            v2 = -160
            _ = fi31
        }
    }
    if (true || true) {
        v32 := GenS1{F0: (v2 ^ v2)}
        _ = v32
    }
    v3 -= genFn1((v2 % ((v2) | 1)), (v2 << ((796) & 63)), -9223372036854775808)
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 uint8 = 220
        switch -905 {
        case 629:
            v35 := []int{int(v2), fi33, fi33}
            _ = v35
        case -447:
            v36 := GenS1{F0: (v2 / (((v2 % (((^(v2))) | 1))) | 1))}
            _ = v36
        case -36:
            v37 := GenS1{F0: (v2 ^ v2)}
            _ = v37
        }
        _ = fi33
        _ = v34
    }
    if true {
        v2 = v2
        for fi38 := 0; fi38 < 2; fi38++ {
            v39 := GenS1{F0: v2}
            v40 := GenS1{F0: 42}
            v41 := GenS2{F0: (fi38 - (fi38 << ((0) & 63))), F1: (^(v3))}
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
        delete(v1, 2147483647)
    }
    v42 := v1[-496]
    switch "foo bar" {
    case "a":
        fmt.Println(v2, v2)
    case "\n\t":
        v4[""] = (((v42 || v42) && (516 == 4294967295)) && true)
    default:
        v2 = (v2 + (v2 * v2))
    }
    v3 = v3
    v42 = (!(!(true || ((v3 | v3) != genFn1(-680, (-(v2)), 1)))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v5
}
