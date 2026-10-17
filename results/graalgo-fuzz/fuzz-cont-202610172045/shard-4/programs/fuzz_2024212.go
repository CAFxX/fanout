package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int; F2 bool }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int32; F1 int }
func genFn0(p0 int64, p1 int8, p2 uint64) int8 {
    for fi4 := 0; fi4 < 5; fi4++ {
        p2 = p2
        var v5 uint8 = 251
        var v6 uint64 = p2
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1() string {
    var v8 int16 = 0
    v9 := []int{-735, -269, -543, 32768}
    var v10 int8 = -1
    v9 = append(v9, -764)
    var v11 float64 = 3.14159
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    v1 := []uint32{948, 0}
    v2 := []uint32{1, 10, 1}
    var v3 int64 = 1
    v2 = append(v2, 42)
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS1{F0: -1, F1: 0, F2: (42 <= -1)}
        for fi14 := 0; fi14 < 2; fi14++ {
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
            v3 = (^(v3))
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            _ = fi14
        }
        v15 := GenS1{F0: -3, F1: 1, F2: (((0.0 / (0.0 + 1.5)) / (3.14159 - 0.1)) >= 1e10)}
        _ = fi12
        _ = v13
        _ = v15
    }
    v16 := GenS2{F0: "a"}
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 int64 = v3
        _ = fi17
        _ = v18
    }
    if false {
        var v19 int8 = -128
        if false {
            var v20 float64 = ((float64(3.14159) - 3.14159) / 0.1)
            v3 = (v3 - v3)
            _ = v20
        }
        _ = v19
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        v22 := GenS2{F0: (("\n\t" + ("" + ("a" + "\n\t"))) + "hello")}
        if true {
            switch ("\n\t" + ("\n\t" + "")) {
            case "\n\t":
                for fi23 := 0; fi23 < 12; fi23++ {
                    if true {
                        switch 42 {
                        case 127:
                            if (!true) {
                                v24 := GenS3{F0: "hello", F1: fi21, F2: ((true && false) || true)}
                                if len(v2) > 0 {
                                    v2[0] = 2
                                }
                                _ = v24
                            }
                        case 32768:
                            v25 := GenS4{F0: 948, F1: 4}
                            _ = v25
                        }
                    } else {
                        var v26 uint16 = 255
                        v27 := GenS5{F0: int32(v3), F1: fi23}
                        _ = v26
                        _ = v27
                    }
                    _ = fi23
                }
            }
            if false {
                v28 := GenS4{F0: fi21, F1: -128}
                _ = v28
            }
            var v29 int8 = genFn0((v3 / (((v3 % ((-344) | 1))) | 1)), 100, 18446744073709551615)
            _ = v29
        } else {
            if (876 <= 2147483648) {
                v3 += (v3 % ((1) | 1))
            }
            v30 := GenS5{F0: 230, F1: fi21}
            _ = v30
        }
        v3 = (v3 - (v3 >> ((543) & 63)))
        _ = fi21
        _ = v22
    }
    v3 = v3
    for fi31 := 0; fi31 < 9; fi31++ {
        for fi32 := 0; fi32 < 19; fi32++ {
            var v33 uint32 = 350
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}
