package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int; F2 uint64 }
type GenS3 struct { F0 int32; F1 int64; F2 uint16 }
type GenS4 struct { F0 int16; F1 int32; F2 string }
func genFn0() uint64 {
    v5 := GenS1{F0: "hello"}
    var v6 float64 = (3.14159 + 1.5)
    v7 := GenS1{F0: ""}
    v8 := []int{32768, 0, 9223372036854775807}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 121
}
func genFn1(p0 int8) uint32 {
    v9 := GenS1{F0: "a"}
    p0 = p0
    if (42 == 0) {
        var v10 uint16 = 255
        var v11 string = "a"
        if false {
            if false {
                var v12 uint64 = genFn0()
                v12 = (v12 & 18446744073709551615)
                _ = v12
            } else {
                var v13 uint16 = (v10 - (v10 / (((^(v10))) | 1)))
                for fi14 := 0; fi14 < 13; fi14++ {
                    if false {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            var v16 int64 = -9223372036854775808
                            switch (fi15 + (fi14 >> ((42) & 63))) {
                            case 32768:
                                var v17 float64 = (3.14159 - 1.5)
                                _ = v17
                            }
                            if (!(!(794 == 148))) {
                                v18 := []int{(fi14 ^ fi15), fi15, -790}
                                _ = v18
                            }
                            _ = fi15
                            _ = v16
                        }
                        switch int(p0) {
                        case 42:
                            v19 := GenS1{F0: v11}
                            _ = v19
                        default:
                            fmt.Println(v11)
                        }
                    } else {
                        v20 := GenS2{F0: (p0 != p0), F1: (fi14 * 435), F2: 609}
                        v11 = "hello"
                        var v21 uint8 = 217
                        _ = v20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 20; fi22++ {
                        v23 := GenS1{F0: v11}
                        var v24 int8 = (p0 | p0)
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = fi14
                }
                _ = v13
            }
        }
        _ = v10
        _ = v11
    } else {
        var v25 float32 = float32(-2.25)
        for fi26 := 0; fi26 < 2; fi26++ {
            p0 = (p0 * p0)
            _ = fi26
        }
        v25 = v25
        _ = v25
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        v28 := []int{}
        _ = fi27
        _ = v28
    }
    _ = p0
    _ = v9
    return 863
}

func main() {
    var v1 float32 = (1e10 + 1e10)
    v2 := make(map[uint32]bool)
    var v3 int = 750
    v4 := make(map[int]uint32)
    switch "\n\t" {
    case "":
        v1 = 0.1
    case "\n\t":
        v1 += (-2.25 + 1.5)
    default:
        v1 = v1
    }
    if ((1 == 255) && (796 != -1)) {
        v1 = ((v1 / (((v1 / (v1 * v1)) + 1.5) + 1e10)) - 1.5)
        v29 := GenS2{F0: (18446744073709551615 <= 2), F1: 4294967295, F2: 2}
        _ = v29
    } else {
        _ = len(v4)
    }
    fmt.Println(v3)
    v1 = v1
    var v30 float32 = v1
    for fi31 := 0; fi31 < 10; fi31++ {
        if true {
            v32 := GenS3{F0: 0, F1: 0, F2: 1}
            var v33 float64 = 0.0
            _ = v32
            _ = v33
        }
        _ = fi31
    }
    fmt.Println(v30)
    var v34 int16 = 32767
    v35 := []int{}
    for fi36 := 0; fi36 < 19; fi36++ {
        var v37 int32 = 47
        _ = fi36
        _ = v37
    }
    v38 := GenS4{F0: (v34 * v34), F1: 42, F2: ("" + "")}
    var v39 int64 = 9223372036854775807
    var v40 bool = (2 != 2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
