package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint64; F1 int; F2 int16 }
func genFn0() int32 {
    v7 := GenS1{F0: 18446744073709551615, F1: -214, F2: 2}
    if (16 <= 1) {
        var v8 int = 52
        v8 += 127
        _ = v8
    } else {
        var v9 uint64 = 0
        _ = v9
    }
    var v10 bool = false
    var v11 int = 127
    v12 := []int{v11, -1}
    _ = v10
    _ = v11
    _ = v12
    _ = v7
    return -1
}
func genFn1(p0 bool, p1 uint16, p2 bool) int8 {
    v13 := []int{32767}
    p2 = (127 != 100)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    var v1 uint64 = 647
    v2 := []string{"foo bar"}
    var v3 uint8 = 190
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    var v6 bool = false
    var v14 int32 = -607
    v15 := GenS2{F0: v1, F1: 42}
    var v16 uint64 = (v1 | (v1 & v1))
    for fi17 := 0; fi17 < 10; fi17++ {
        _ = len(v5)
        v18 := GenS3{F0: ((v6 && (-105 > -1)) && v6), F1: -667}
        _ = fi17
        _ = v18
    }
    fmt.Println(v14, v3)
    var v19 int64 = 0
    v4["\n\t"] = v6
    if (255 != 0) {
        if true {
            v19 = (v19 * (^(v19)))
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 uint8 = (v3 ^ (v3 * v3))
            if (false || v6) {
                if v6 {
                    var v22 uint16 = 1
                    var v23 uint = uint(v22)
                    var v24 uint64 = 569
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    fmt.Println(v21)
                }
                v21 = 149
                v2 = append(v2, "")
            }
            _ = fi20
            _ = v21
        }
    } else {
        v25 := GenS4{F0: 0, F1: -32768, F2: (v3 | (v3 ^ (v3 % ((v3) | 1))))}
        switch "hello" {
        case "":
            for fi26 := 0; fi26 < 15; fi26++ {
                switch (fi26 % ((fi26) | 1)) {
                case 969:
                    switch "hello" {
                    case "\n\t":
                        v2 = append(v2, "foo bar")
                    default:
                        var v27 uint16 = 0
                        _ = v27
                    }
                case 42:
                    v19 += 42
                }
                var v28 uint16 = 255
                _ = fi26
                _ = v28
            }
        case "hello":
            if (-1 == (v19 % (((v19 << ((uint(v1)) & 63))) | 1))) {
                for fi29 := 0; fi29 < 7; fi29++ {
                    for fi30 := 0; fi30 < 18; fi30++ {
                        var v31 int8 = -128
                        _ = fi30
                        _ = v31
                    }
                    if len(v2) > 0 {
                        v2[0] = (("foo bar" + (("foo bar" + "a") + "a")) + "\n\t")
                    }
                    _ = fi29
                }
            }
        case "\n\t":
            delete(v4, ("foo bar" + "\n\t"))
        }
        var v32 int16 = -263
        _ = v25
        _ = v32
    }
    v33 := GenS5{F0: 825, F1: 0, F2: -1}
    delete(v5, -9223372036854775808)
    if v6 {
        var v34 int16 = 0
        if v6 {
            if (v34 <= (v34 / ((v34) | 1))) {
                var v35 uint = 18446744073709551615
                var v36 int64 = -9223372036854775808
                _ = v35
                _ = v36
            } else {
                fmt.Println(v3)
                var v37 uint = 444
                _ = v37
            }
        }
        var v38 uint16 = 688
        _ = v34
        _ = v38
    }
    v39 := GenS1{F0: 184, F1: -355, F2: v3}
    var v40 int32 = (-(v14))
    for fi41 := 0; fi41 < 6; fi41++ {
        v6 = v6
        if (genFn0() != (v14 & (v40 / (((v14 >> ((1) & 63))) | 1)))) {
            var v42 float32 = 1.5
            v19 -= v19
            var v43 bool = (65535 >= 2)
            _ = v42
            _ = v43
        } else {
            var v44 int32 = 1
            v45 := GenS4{F0: 4294967295, F1: 1000, F2: v3}
            _ = v44
            _ = v45
        }
        _ = fi41
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
