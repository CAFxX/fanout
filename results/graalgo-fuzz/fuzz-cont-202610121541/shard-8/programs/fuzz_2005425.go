package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int }
type GenS3 struct { F0 int32 }
func genFn0() uint32 {
    var v7 uint32 = 2147483647
    var v8 uint16 = 2
    for fi9 := 0; fi9 < 18; fi9++ {
        if (!false) {
            fmt.Println(fi9, fi9)
            switch -154 {
            case 128:
                v7 -= 720
            default:
                v10 := []int{(-(fi9)), (fi9 << ((382) & 63)), fi9}
                _ = v10
            }
        } else {
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := []int{fi9}
                _ = fi11
                _ = v12
            }
            switch "hello" {
            case "\n\t":
                var v13 uint8 = 41
                _ = v13
            case "a":
                switch ("a" + "\n\t") {
                case "":
                    v14 := GenS1{F0: v8, F1: "a", F2: (^(v8))}
                    _ = v14
                case "\n\t":
                    v7 += 554
                case "foo bar":
                    fmt.Println(v8, fi9, v8)
                default:
                    if ((v8 << ((368) & 63)) == v8) {
                        v7 = 831
                    } else {
                        v15 := []int{}
                        if (0.1 < (-2.25 - 0.1)) {
                            var v16 int8 = -128
                            _ = v16
                        }
                        if len(v15) > 0 {
                            v15[0] = (fi9 * fi9)
                        }
                        _ = v15
                    }
                }
            }
            fmt.Println(fi9, v7)
        }
        _ = fi9
    }
    v17 := []int{128, -9223372036854775808, 42, -2147483648}
    v7 = (^(v7))
    _ = v17
    _ = v7
    _ = v8
    return v7
}
func genFn1(p0 uint16) uint64 {
    var v18 uint64 = 18446744073709551615
    p0 = 65535
    _ = p0
    _ = v18
    return uint64(p0)
}

func main() {
    v1 := []int{0, 9223372036854775807, 0}
    var v2 int = -2147483648
    var v3 int32 = 0
    var v4 int8 = 87
    var v5 uint16 = 884
    v6 := make(map[uint32]bool)
    v3 += 726
    if ((false || true) || true) {
        if ((0 < 889) || true) {
            var v19 string = "foo bar"
            if len(v1) > 0 {
                v1[0] = (-(v2))
            }
            for fi20 := 0; fi20 < 16; fi20++ {
                if (true || true) {
                    if false {
                        v19 = "\n\t"
                        var v21 int32 = 2147483647
                        _ = v21
                    } else {
                        fmt.Println(v3, v19, v5)
                        fmt.Println(v19)
                        _ = len(v1)
                    }
                }
                _ = fi20
            }
            _ = v19
        }
        v3 -= 42
    } else {
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 uint32 = 296
            v4 = (v4 / ((v4) | 1))
            _ = fi22
            _ = v23
        }
        if false {
            v24 := GenS1{F0: 831, F1: "\n\t", F2: 1}
            v25 := v6[genFn0()]
            v4 = v4
            _ = v24
            _ = v25
        }
    }
    if (!true) {
        switch 0 {
        case 198:
            for fi26 := 0; fi26 < 9; fi26++ {
                var v27 float32 = 3.14159
                _ = fi26
                _ = v27
            }
        }
    }
    if ((v3 >> ((94) & 63)) <= v3) {
        var v28 int32 = 2147483647
        var v29 float64 = (0.1 - 3.14159)
        _ = v28
        _ = v29
    } else {
        for fi30 := 0; fi30 < 3; fi30++ {
            if (!false) {
                if (!true) {
                    fmt.Println(v3, v3, fi30)
                    v2 += -2
                } else {
                    for fi31 := 0; fi31 < 19; fi31++ {
                        v3 = 42
                        _ = fi31
                    }
                    fmt.Println(v5, fi30, v2)
                }
                var v32 float64 = -2.25
                var v33 float64 = v32
                _ = v32
                _ = v33
            }
            var v34 uint32 = 2
            v35 := GenS1{F0: (v5 ^ v5), F1: "hello", F2: 1}
            _ = fi30
            _ = v34
            _ = v35
        }
        v36 := GenS1{F0: 65535, F1: "foo bar", F2: (v5 - (^(v5)))}
        _ = v36
    }
    var v37 int = v2
    fmt.Println(v4, v4)
    if true {
        var v38 float32 = 1.5
        v39 := GenS1{F0: v5, F1: "a", F2: v5}
        _ = v38
        _ = v39
    }
    v40 := GenS1{F0: 2, F1: "hello", F2: (v5 >> ((561) & 63))}
    v41 := GenS2{F0: 0, F1: (v5 / ((v5) | 1)), F2: 2147483647}
    if (false && (true || ((-(v4)) != (^(v4))))) {
        v5 += (v5 ^ 255)
        v4 = int8(v2)
        var v42 float32 = 1e10
        _ = v42
    } else {
        if len(v1) > 0 {
            v1[0] = (v2 % ((v2) | 1))
        }
        v43 := GenS3{F0: v3}
        _ = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
