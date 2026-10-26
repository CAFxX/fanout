package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint }
func genFn0() bool {
    var v6 int = 128
    fmt.Println(v6)
    _ = v6
    return true
}

func main() {
    var v1 uint = 308
    v2 := []uint32{4294967295}
    v3 := []string{("foo bar" + ""), "", (((("\n\t" + ("" + "foo bar")) + "hello") + "") + "")}
    var v4 int = -714
    var v5 int = v4
    switch (((("\n\t" + ("" + "")) + "foo bar") + (("foo bar" + ("" + "")) + "hello")) + ("\n\t" + ("a" + ((("foo bar" + "a") + ("foo bar" + "\n\t")) + ("a" + "\n\t"))))) {
    case "":
        v1 = (v1 * v1)
    case "a":
        v7 := GenS1{F0: 42, F1: v4}
        _ = v7
    default:
        if false {
            v5 -= v5
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    var v8 int32 = 0
    var v9 uint32 = 565
    if (((v9 & 962) <= v9) || (18446744073709551615 >= 0)) {
        v1 -= (v1 & (v1 + (v1 << (((^(v1))) & 63))))
        v4 -= (-(v4))
        if (!false) {
            for fi10 := 0; fi10 < 15; fi10++ {
                for fi11 := 0; fi11 < 14; fi11++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if false {
                        var v12 int = fi10
                        var v13 float32 = (0.1 + (float32((((float64(float32((3.14159 - float64(-2.25)))) / (((((3.14159 - 1e10) * 0.1) - 0.1) / 1.5) + (float64(0.1) * 0.0))) - ((0.0 * ((((1.5 - ((0.0 - 1.5) + -2.25)) - 0.0) + -2.25) * (float64((3.14159 + 0.1)) - 3.14159))) / 1.0)) * (1.5 * (-2.25 / 1.5)))) * 1.5))
                        var v14 int = v12
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            v15 := GenS2{F0: 128, F1: v9, F2: v8}
                            _ = v15
                        default:
                            if (!(!((162 < 42) && (!true)))) {
                                v4 -= (fi11 | 32768)
                            }
                        }
                    }
                    fmt.Println(v8, v1)
                    v16 := GenS1{F0: 32767, F1: v4}
                    v4 -= (v4 / (((fi10 + (^(fi11)))) | 1))
                    var v17 int8 = -1
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    _ = fi11
                    _ = v16
                    _ = v17
                }
                _ = fi10
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v3 = append(v3, "a")
        } else {
            v4 += 168
        }
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        v1 += v1
        fmt.Println(v5)
        _ = fi18
    }
    v2 = append(v2, v9)
    var v19 uint64 = 791
    if true {
        fmt.Println(v1, v5, v9)
        v5 -= -612
    }
    var v20 int64 = -1
    _ = len(v2)
    v21 := GenS2{F0: 127, F1: v9, F2: int32(v19)}
    var v22 string = "foo bar"
    for fi23 := 0; fi23 < 11; fi23++ {
        fmt.Println(v20)
        switch ((("a" + v22) + v22) + (((((v22 + v22) + v22) + (("\n\t" + (v22 + "a")) + "\n\t")) + "foo bar") + ("\n\t" + v22))) {
        case "hello":
            v24 := GenS1{F0: 78, F1: (v5 ^ (fi23 / (((fi23 + 452)) | 1)))}
            _ = v24
        case "":
            switch fi23 {
            case 491:
                if false {
                    if true {
                        v19 -= (v19 * (v19 | (v19 >> ((v1) & 63))))
                        v25 := GenS3{F0: (v8 << ((v1) & 63))}
                        _ = v25
                    } else {
                        var v26 float32 = ((3.14159 - 0.0) * 1.5)
                        _ = v26
                    }
                } else {
                    v19 -= (v19 - 517)
                    var v27 int16 = -1
                    var v28 uint16 = 612
                    _ = v27
                    _ = v28
                }
            case 32767:
                switch v22 {
                case "":
                    var v29 int16 = 0
                    _ = v29
                default:
                    if len(v3) > 0 {
                        v3[0] = ("a" + "foo bar")
                    }
                }
            }
        default:
            _ = len(v2)
        }
        _ = fi23
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        v31 := GenS4{F0: (v9 - v9), F1: (true && ((!true) || (uint16(v19) < 0))), F2: v1}
        fmt.Println(v19, v19, fi30)
        _ = fi30
        _ = v31
    }
    var v32 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
