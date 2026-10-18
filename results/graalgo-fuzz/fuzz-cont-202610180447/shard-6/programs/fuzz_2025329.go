package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint }
type GenS3 struct { F0 uint64; F1 bool; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 string; F2 int }
func genFn0(p0 uint16, p1 uint32) uint64 {
    p0 -= (p0 / ((p0) | 1))
    var v5 uint16 = 0
    _ = p0
    _ = p1
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 uint16, p1 bool, p2 int8) uint32 {
    p0 -= (p0 + (p0 ^ (p0 % ((p0) | 1))))
    var v6 int16 = 639
    v7 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ("foo bar" + "a")
    var v3 int = 4294967295
    var v4 float32 = float32((float64(3.14159) / 1e10))
    v8 := []int{32767}
    v2 = v2
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            switch fi9 {
            case 532:
                var v11 bool = (false || (42 > 66))
                _ = v11
            default:
                var v12 uint32 = 2147483648
                _ = v12
            }
            if true {
                if false {
                    v2 = v2
                    switch (v2 + v2) {
                    case "a":
                        _ = len(v1)
                    }
                } else {
                    fmt.Println(v4)
                    if true {
                        var v13 string = (v2 + "foo bar")
                        switch fi9 {
                        case 32768:
                            var v14 int32 = 892
                            _ = v14
                        case -1:
                            v13 = ""
                        case -2147483648:
                            v2 = v13
                        }
                        _ = v13
                    }
                }
                if (!false) {
                    v4 += 0.0
                    var v15 int32 = 948
                    fmt.Println(fi10, v15, v3)
                    _ = v15
                } else {
                    _ = len(v1)
                    fmt.Println(fi10)
                }
                v16 := GenS1{F0: 2}
                _ = v16
            }
            v17 := GenS2{F0: 714, F1: 139, F2: 0}
            _ = fi10
            _ = v17
        }
        _ = fi9
    }
    fmt.Println(v3, v2, v2)
    fmt.Println(v3)
    var v18 int = (v3 ^ -954)
    v4 -= (0.0 / v4)
    v1[2147483648] = ((("" + v2) + "") + (((v2 + ((v2 + "hello") + v2)) + (v2 + v2)) + (v2 + v2)))
    switch v2 {
    case "a":
        v18 = v18
    default:
        delete(v1, genFn1(162, (!(!((true && (100 <= 0)) && (0.1 > 1e10)))), -1))
    }
    var v19 uint64 = 400
    switch (v2 + "") {
    case "a":
        var v20 int32 = 61
        _ = v20
    case "hello":
        switch ("" + "\n\t") {
        case "\n\t":
            if false {
                var v21 uint64 = genFn0(467, 2)
                if (0 != 1) {
                    v2 = v2
                    fmt.Println(v18)
                }
                _ = v21
            } else {
                var v22 float32 = float32((1e10 + (float64(v4) * float64(v4))))
                v3 += (v18 * -2147483648)
                _ = v22
            }
        case "":
            var v23 int = v3
            _ = v23
        case "foo bar":
            v2 = v2
        default:
            v24 := GenS3{F0: v19, F1: (!(int64(v3) > 42)), F2: 677}
            _ = v24
        }
    case "":
        v2 = ("\n\t" + v2)
    default:
        var v25 float64 = 3.14159
        _ = v25
    }
    var v26 int32 = 80
    if true {
        if true {
            v27 := GenS3{F0: v19, F1: ((188 > 0) || (!true)), F2: 0}
            v28 := GenS2{F0: 4294967295, F1: 96, F2: 792}
            v29 := GenS4{F0: 52}
            _ = v27
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 14; fi30++ {
            v1[0] = v2
            switch "hello" {
            case "\n\t":
                for fi31 := 0; fi31 < 4; fi31++ {
                    var v32 float64 = -2.25
                    v3 -= v18
                    switch v3 {
                    case 127:
                        var v33 bool = true
                        _ = v33
                    }
                    _ = fi31
                    _ = v32
                }
            case "hello":
                delete(v1, 65535)
            case "a":
                var v34 int32 = (-(v26))
                _ = v34
            default:
                if false {
                    var v35 int = (v18 - v3)
                    v36 := GenS4{F0: 2}
                    _ = v35
                    _ = v36
                }
            }
            _ = fi30
        }
    } else {
        _ = len(v8)
    }
    v37 := GenS5{F0: -128, F1: v2, F2: v3}
    var v38 uint = 65535
    v39 := GenS2{F0: 55, F1: 2, F2: 0}
    v40 := GenS2{F0: 4294967295, F1: uint8(v38), F2: 18446744073709551615}
    fmt.Println(v26, v4, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v8
}
