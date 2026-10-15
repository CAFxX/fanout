package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint64, p1 uint16) uint {
    var v7 int8 = 127
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: -1, F1: "", F2: 510}
        v10 := []int{fi8}
        var v11 int = fi8
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint32 = 1
    switch "" {
    case "\n\t":
        v13 := GenS1{F0: 1, F1: "foo bar", F2: (v12 % ((v12) | 1))}
        _ = v13
    default:
        switch 21 {
        case 127:
            for fi14 := 0; fi14 < 19; fi14++ {
                v7 = (v7 + 42)
                var v15 uint = 0
                _ = fi14
                _ = v15
            }
        case 391:
            v16 := []int{}
            _ = v16
        default:
            var v17 string = "a"
            _ = v17
        }
    }
    v18 := GenS2{F0: 0, F1: 0}
    _ = p0
    _ = p1
    _ = v12
    _ = v18
    _ = v7
    return 1
}
func genFn1(p0 uint) uint64 {
    p0 = genFn0(1, 42)
    var v19 int8 = -128
    _ = p0
    _ = v19
    return 42
}

func main() {
    var v1 uint16 = 255
    v2 := []int{}
    v3 := []string{"hello", ("" + "foo bar")}
    v4 := make(map[uint32]string)
    var v5 int64 = 674
    var v6 uint16 = 1
    if (!true) {
        v6 = v6
        if true {
            var v20 float64 = 0.0
            for fi21 := 0; fi21 < 9; fi21++ {
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 uint = 65535
                    _ = fi22
                    _ = v23
                }
                var v24 uint = 65535
                _ = fi21
                _ = v24
            }
            _ = v20
        }
        v25 := GenS1{F0: 578, F1: "", F2: 660}
        _ = v25
    }
    fmt.Println(v5, v6, v6)
    var v26 uint16 = (v6 << ((766) & 63))
    _ = len(v2)
    if false {
        for fi27 := 0; fi27 < 16; fi27++ {
            v1 += 0
            _ = fi27
        }
        switch "hello" {
        case "":
            var v28 uint32 = 0
            _ = v28
        }
    } else {
        v29 := GenS1{F0: 0, F1: "", F2: 42}
        for fi30 := 0; fi30 < 9; fi30++ {
            for fi31 := 0; fi31 < 10; fi31++ {
                v32 := GenS3{F0: "a"}
                switch "a" {
                case "\n\t":
                    var v33 int64 = v5
                    _ = v33
                case "":
                    if (!false) {
                        var v34 int8 = 127
                        var v35 string = "hello"
                        fmt.Println(v5)
                        _ = v34
                        _ = v35
                    } else {
                        v4[560] = "foo bar"
                        v36 := GenS2{F0: v5, F1: v5}
                        var v37 float32 = 1.5
                        v1 -= (^(v1))
                        _ = v36
                        _ = v37
                    }
                }
                v38 := GenS4{F0: v5, F1: 18446744073709551615, F2: 1}
                v39 := v4[1]
                delete(v4, 754)
                _ = fi31
                _ = v32
                _ = v38
                _ = v39
            }
            if (!(!false)) {
                v40 := GenS2{F0: v5, F1: -1}
                _ = v40
            }
            _ = fi30
        }
        _ = v29
    }
    v41 := GenS2{F0: (-(v5)), F1: v5}
    if true {
        var v42 uint8 = 2
        v43 := GenS4{F0: (^(v5)), F1: 622, F2: 1}
        switch 127 {
        case -917:
            var v44 bool = (!true)
            _ = v44
        }
        _ = v42
        _ = v43
    } else {
        var v45 int32 = 2147483647
        _ = v45
    }
    v6 -= v6
    var v46 uint64 = 42
    v47 := GenS5{F0: 559}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
