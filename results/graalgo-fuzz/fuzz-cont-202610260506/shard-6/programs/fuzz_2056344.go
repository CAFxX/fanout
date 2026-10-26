package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int32; F1 int8; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int8 }
func genFn0() uint32 {
    switch "a" {
    case "\n\t":
        var v4 bool = false
        _ = v4
    case "":
        v5 := GenS1{F0: 2, F1: 1}
        _ = v5
    }
    var v6 float32 = ((((0.0 + 0.1) + 1e10) - (0.1 + (-2.25 * -2.25))) - (0.0 * (0.0 - 0.1)))
    v7 := GenS1{F0: 42, F1: 438}
    v8 := []int{}
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1() int64 {
    v9 := []int{}
    var v10 uint8 = 42
    var v11 uint64 = 18446744073709551615
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 42
    v2 := []uint32{}
    v3 := make(map[uint32]string)
    delete(v3, v1)
    v12 := GenS1{F0: 0, F1: 302}
    v13 := v3[v1]
    switch (v13 + ("" + "")) {
    case "\n\t":
        for fi14 := 0; fi14 < 8; fi14++ {
            _ = len(v3)
            v1 = 865
            _ = fi14
        }
    case "a":
        v15 := GenS1{F0: 2, F1: 42}
        _ = v15
    case "hello":
        if false {
            v13 = "a"
            v16 := v3[genFn0()]
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS1{F0: 2, F1: 947}
            _ = len(v3)
            _ = fi17
            _ = v18
        }
    }
    fmt.Println(v1)
    if (1 != 42) {
        switch -1 {
        case 32768:
            var v19 uint = 65535
            _ = v19
        case 42:
            var v20 uint8 = 150
            _ = v20
        default:
            v1 = 42
        }
        v21 := GenS2{F0: 42, F1: 127, F2: 1}
        for fi22 := 0; fi22 < 19; fi22++ {
            v1 = genFn0()
            var v23 int = (fi22 / (((fi22 << ((604) & 63))) | 1))
            v3[972] = "\n\t"
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        v1 = v1
        var v24 int16 = -525
        _ = v24
    }
    _ = len(v2)
    v2 = append(v2, v1)
    v25 := GenS3{F0: -128}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 9223372036854775807 {
    case 128:
        v1 = genFn0()
    case 9223372036854775807:
        v26 := GenS3{F0: 42}
        _ = v26
    case 127:
        v13 = ((((v13 + "\n\t") + v13) + "a") + v13)
    default:
        if false {
            var v27 int16 = -32768
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        v29 := GenS4{F0: -486, F1: -267}
        if true {
            var v30 int64 = 42
            if (((false && false) && (2 >= 42)) && false) {
                v31 := GenS2{F0: 2147483647, F1: 127, F2: v30}
                _ = v31
            } else {
                if (614 <= 752) {
                    if false {
                        v30 = 42
                        for fi32 := 0; fi32 < 13; fi32++ {
                            v33 := GenS4{F0: -658, F1: 2147483647}
                            _ = fi32
                            _ = v33
                        }
                        v34 := GenS4{F0: 1, F1: 1}
                        _ = v34
                    }
                    if ((fi28 + fi28) == fi28) {
                        for fi35 := 0; fi35 < 8; fi35++ {
                            if ((42 > 712) && true) {
                                fmt.Println(fi28, fi35)
                                if (true && (true && (746 < 974))) {
                                    v2 = append(v2, genFn0())
                                    if false {
                                        if len(v2) > 0 {
                                            v2[0] = 42
                                        }
                                        v36 := GenS3{F0: -128}
                                        _ = v36
                                    } else {
                                        var v37 uint32 = v1
                                        if true {
                                            v38 := GenS1{F0: 42, F1: 0}
                                            v39 := GenS5{F0: 42, F1: true}
                                            v40 := GenS6{F0: -72}
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                        _ = v37
                                    }
                                    _ = len(v2)
                                    if (3.14159 == 1.5) {
                                        var v41 float64 = 0.1
                                        _ = v41
                                    }
                                }
                            }
                            _ = fi35
                        }
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (!((((!(235 != 105)) || true) || (!(!true))) && false)) {
                            v30 = -744
                        } else {
                            v1 += v1
                        }
                    }
                    for fi42 := 0; fi42 < 7; fi42++ {
                        if true {
                            fmt.Println(fi28, fi42, v13)
                        }
                        v1 += (v1 << ((42) & 63))
                        switch (fi28 | 9223372036854775807) {
                        case -9223372036854775808:
                            v1 = (v1 >> ((4294967295) & 63))
                        case 177:
                            v30 += (v30 << ((42) & 63))
                        default:
                            var v43 uint32 = 2147483648
                            _ = v43
                        }
                        _ = fi42
                    }
                } else {
                    v13 = (v13 + "hello")
                    var v44 int8 = 100
                    _ = v44
                }
            }
            fmt.Println(fi28)
            _ = v30
        }
        _ = fi28
        _ = v29
    }
    v1 += (v1 * v1)
    _ = len(v2)
    v13 = ((((v13 + v13) + v13) + v13) + ("" + (("" + "foo bar") + v13)))
    var v45 float64 = (1e10 * 1e10)
    v45 = float64((1e10 - 0.1))
    for fi46 := 0; fi46 < 10; fi46++ {
        var v47 int = (fi46 << ((42) & 63))
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v45
}
