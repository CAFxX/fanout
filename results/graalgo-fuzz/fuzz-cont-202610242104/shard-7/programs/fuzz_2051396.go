package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int8; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 uint16; F2 uint16 }
func genFn0(p0 bool, p1 uint) int32 {
    var v5 uint64 = 1
    v6 := []int{127, -483}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1() string {
    v7 := []int{-253, 9223372036854775807, 128, 32767}
    if (9223372036854775807 >= 42) {
        var v8 string = ("a" + ("a" + "foo bar"))
        var v9 int32 = genFn0(((!false) || (1 == -2147483648)), 42)
        _ = v8
        _ = v9
    } else {
        var v10 float32 = 0.1
        if (444 == 366) {
            var v11 int16 = 1
            if true {
                for fi12 := 0; fi12 < 1; fi12++ {
                    v10 += float32(0.0)
                    for fi13 := 0; fi13 < 20; fi13++ {
                        _ = len(v7)
                        v11 = (-(v11))
                        if (-44 > -1) {
                            var v14 bool = (0 < 42)
                            v15 := GenS1{F0: 1}
                            v14 = (4294967295 == 2)
                            var v16 float32 = float32((((0.1 / (float64(v10) - float64(v10))) / -2.25) * 0.0))
                            switch "foo bar" {
                            case "foo bar":
                                var v17 int8 = -128
                                _ = v17
                            case "a":
                                var v18 uint64 = 42
                                _ = v18
                            case "hello":
                                if (255 <= 140) {
                                    v19 := GenS1{F0: -1}
                                    _ = v19
                                }
                            default:
                                var v20 uint16 = 255
                                _ = v20
                            }
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            v21 := GenS1{F0: 9223372036854775807}
                            if ("a" >= "a") {
                                v11 += v11
                                fmt.Println(v11)
                            } else {
                                v10 = v10
                                v22 := GenS1{F0: -1}
                                if (true || true) {
                                    v11 += (v11 % ((v11) | 1))
                                }
                                _ = v22
                            }
                            var v23 int16 = -13
                            _ = v21
                            _ = v23
                        }
                        _ = fi13
                    }
                    _ = fi12
                }
                var v24 uint32 = 788
                _ = v24
            } else {
                v25 := GenS2{F0: -2147483648}
                v10 = 1e10
                for fi26 := 0; fi26 < 17; fi26++ {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        v28 := GenS3{F0: -93, F1: 42, F2: -2147483648}
                        switch fi26 {
                        case -1:
                            var v29 uint64 = 1
                            _ = v29
                        default:
                            v11 = (v11 + -1)
                        }
                        v30 := GenS4{F0: (1 < 41)}
                        var v31 float64 = (3.14159 / 1.0)
                        var v32 int64 = -9223372036854775808
                        _ = fi27
                        _ = v28
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    _ = fi26
                }
                _ = v25
            }
            _ = v11
        }
        _ = v10
    }
    if len(v7) > 0 {
        v7[0] = 32767
    }
    var v33 string = "hello"
    _ = v33
    _ = v7
    return ("\n\t" + ((v33 + (v33 + v33)) + ""))
}

func main() {
    v1 := []int{42, -1, -2147483648, 128}
    var v2 int16 = 1
    var v3 string = ""
    v4 := []string{""}
    v3 = v3
    if false {
        var v34 int16 = v2
        if len(v4) > 0 {
            v4[0] = ("foo bar" + v3)
        }
        _ = v34
    }
    v35 := GenS1{F0: 9223372036854775807}
    for fi36 := 0; fi36 < 17; fi36++ {
        v4 = append(v4, (v3 + v3))
        for fi37 := 0; fi37 < 2; fi37++ {
            if (840 <= 2) {
                var v38 int64 = 0
                _ = v38
            }
            var v39 int16 = v2
            if true {
                var v40 uint64 = 18446744073709551615
                var v41 float32 = float32((float64(1e10) + (((1.5 - 3.14159) / 1.5) - (((((float64(0.0) + ((3.14159 - ((0.0 + float64(1e10)) - 1e10)) * (1.5 + -2.25))) / ((-2.25 + 1.5) + ((-2.25 / (1e10 - ((1e10 * (((0.0 + 0.1) * (float64(1.5) + (0.0 + 0.0))) + float64(((0.0 / 1.0) / 1.0)))) - -2.25))) / (1e10 - 0.0)))) * 3.14159) - -2.25) / float64(-2.25)))))
                v41 -= 3.14159
                _ = v40
                _ = v41
            }
            _ = fi37
            _ = v39
        }
        if true {
            if false {
                var v42 int = fi36
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v43 uint8 = 255
                _ = v42
                _ = v43
            }
        } else {
            var v44 int8 = 20
            v45 := GenS1{F0: -9223372036854775808}
            _ = v44
            _ = v45
        }
        _ = fi36
    }
    v46 := GenS5{F0: 9223372036854775807, F1: 32767, F2: 2}
    var v47 string = "hello"
    v48 := GenS3{F0: 100, F1: -128, F2: -1}
    v49 := GenS5{F0: 0, F1: 65535, F2: 445}
    _ = len(v1)
    _ = len(v1)
    var v50 string = (("" + "hello") + (v3 + "foo bar"))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
