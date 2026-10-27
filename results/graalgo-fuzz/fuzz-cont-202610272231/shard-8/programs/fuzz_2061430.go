package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 string; F1 bool; F2 int16 }
func genFn0(p0 uint) uint {
    switch ("foo bar" + "\n\t") {
    case "":
        if (!(!true)) {
            if (!false) {
                var v5 string = "\n\t"
                var v6 uint32 = 2147483647
                _ = v5
                _ = v6
            }
        }
    }
    fmt.Println(p0)
    p0 += (p0 / (((p0 / ((1) | 1))) | 1))
    p0 -= 65535
    _ = p0
    return p0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    switch 32767 {
    case 32767:
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int16 = 996
            _ = fi7
            _ = v8
        }
    case 127:
        var v9 int8 = 100
        _ = v9
    }
    var v10 int8 = 1
    switch ("foo bar" + ("hello" + ("" + "foo bar"))) {
    case "foo bar":
        var v11 float32 = float32(0.1)
        _ = v11
    }
    var v12 int64 = 42
    if (42 >= 1) {
        if (4294967295 >= 999) {
            switch "" {
            case "":
                for fi13 := 0; fi13 < 5; fi13++ {
                    v14 := GenS1{F0: -1, F1: 42}
                    _ = fi13
                    _ = v14
                }
            case "a":
                if true {
                    if true {
                        var v15 int8 = v10
                        _ = v15
                    }
                    var v16 uint32 = 2
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 4; fi17++ {
                    switch ("a" + "hello") {
                    case "a":
                        fmt.Println(fi17, v12)
                    case "hello":
                        v12 += (v12 - 0)
                    }
                    v18 := []int{fi17}
                    var v19 float64 = 1e10
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
        } else {
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 int32 = -734
                var v22 int8 = v10
                v23 := GenS2{F0: 65535}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        }
        var v24 int = 9223372036854775807
        v25 := GenS3{F0: 127, F1: 255}
        _ = v24
        _ = v25
    }
    if (!false) {
        for fi26 := 0; fi26 < 10; fi26++ {
            fmt.Println(fi26, fi26)
            var v27 uint8 = 128
            v27 -= v27
            _ = fi26
            _ = v27
        }
        var v28 float64 = 1.5
        v12 -= (-(v12))
        _ = v28
    }
    var v29 float64 = 0.0
    for fi30 := 0; fi30 < 3; fi30++ {
        if true {
            v29 = v29
            switch fi30 {
            case 778:
                for fi31 := 0; fi31 < 9; fi31++ {
                    v32 := GenS2{F0: 2}
                    v33 := []int{0}
                    fmt.Println(v12)
                    v34 := GenS4{F0: v10, F1: 42}
                    switch (fi31 % (((-(fi30))) | 1)) {
                    case -9223372036854775808:
                        var v35 int64 = 42
                        _ = v35
                    case 32767:
                        v10 = (-(v10))
                    case 1:
                        var v36 int = 388
                        _ = v36
                    }
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
            default:
                for fi37 := 0; fi37 < 5; fi37++ {
                    switch 4294967295 {
                    case 509:
                        var v38 int32 = 2147483647
                        _ = v38
                    case 4294967295:
                        switch fi37 {
                        case 128:
                            v10 = 0
                        }
                    case 0:
                        _ = len(v3)
                    default:
                        if false {
                            v39 := []int{}
                            v12 -= (v12 ^ v12)
                            v2[2] = (fi37 | (fi37 ^ fi30))
                            v39 = append(v39, (-(fi37)))
                            v40 := GenS4{F0: v10, F1: 0}
                            fmt.Println(v29, fi37)
                            var v41 uint16 = 121
                            var v42 float32 = (float32((v29 + ((0.0 * 1e10) - v29))) - ((-2.25 - (float32(-2.25) - 0.1)) / 1.5))
                            _ = v39
                            _ = v40
                            _ = v41
                            _ = v42
                        } else {
                            var v43 float64 = 1.5
                            _ = v43
                        }
                    }
                    fmt.Println(v29)
                    var v44 string = ("" + "a")
                    v3[(fi37 ^ fi37)] = 65535
                    delete(v2, 144)
                    _ = fi37
                    _ = v44
                }
            }
            v45 := GenS5{F0: "", F1: (1000 > 0), F2: -456}
            _ = v45
        }
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
