package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 uint64; F2 int8 }
func genFn0(p0 uint8, p1 uint8, p2 int) int64 {
    var v6 float64 = float64(0.0)
    v7 := []int{(^(p2)), (^(p2)), -569}
    var v8 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (v8 * (v8 - (v8 ^ v8)))
}
func genFn1(p0 bool) bool {
    v9 := []int{4294967295, 968, 32768}
    var v10 uint64 = 0
    _ = p0
    _ = v10
    _ = v9
    return false
}

func main() {
    var v1 int32 = -898
    v2 := make(map[string]int)
    var v3 int64 = 9223372036854775807
    v4 := []int{}
    var v5 int = 75
    delete(v2, "foo bar")
    fmt.Println(v3, v5)
    switch "a" {
    case "\n\t":
        switch v5 {
        case -1:
            v5 = v5
        case -9223372036854775808:
            if (true && true) {
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 int32 = (v1 - (v1 / ((2147483647) | 1)))
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 7; fi13++ {
                    v3 = v3
                    _ = fi13
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                switch "" {
                case "\n\t":
                    if true {
                        v14 := GenS1{F0: true, F1: 42}
                        v15 := v2[((("" + "a") + (("" + "") + "")) + "")]
                        _ = v14
                        _ = v15
                    } else {
                        var v16 string = "\n\t"
                        fmt.Println(v16, v16, v1)
                        fmt.Println(v16, v16)
                        _ = v16
                    }
                case "hello":
                    if false {
                        var v17 float32 = (float32(-2.25) - 0.1)
                        if (!false) {
                            v18 := GenS1{F0: false, F1: 448}
                            fmt.Println(v1, v3)
                            _ = v18
                        }
                        v5 += -1
                        _ = v17
                    }
                case "foo bar":
                    fmt.Println(v1)
                default:
                    for fi19 := 0; fi19 < 19; fi19++ {
                        for fi20 := 0; fi20 < 2; fi20++ {
                            v1 = (v1 + (v1 + (v1 << ((18446744073709551615) & 63))))
                            v21 := v2["\n\t"]
                            v1 -= v1
                            switch "hello" {
                            case "hello":
                                if (true || true) {
                                    v4 = append(v4, (fi19 << ((18446744073709551615) & 63)))
                                    v5 += fi19
                                    fmt.Println(v21, v21, fi19)
                                }
                            case "a":
                                fmt.Println(v1, fi19, v5)
                            default:
                                var v22 int64 = (v3 % ((int64(v1)) | 1))
                                _ = v22
                            }
                            var v23 int64 = (v3 ^ genFn0(227, 255, v21))
                            v1 = int32(fi20)
                            switch "" {
                            case "a":
                                delete(v2, (("" + (("hello" + ("a" + ("" + "a"))) + "a")) + ("hello" + "")))
                            }
                            _ = fi20
                            _ = v21
                            _ = v23
                        }
                        _ = fi19
                    }
                }
            }
        default:
            if false {
                for fi24 := 0; fi24 < 15; fi24++ {
                    fmt.Println(v5, fi24, v1)
                    _ = fi24
                }
                v3 -= genFn0(2, 255, 785)
                v5 += (^(v5))
            } else {
                v25 := GenS2{F0: 0, F1: "foo bar"}
                for fi26 := 0; fi26 < 14; fi26++ {
                    v1 += -553
                    v27 := GenS3{F0: 1, F1: 1, F2: 952}
                    fmt.Println(v3)
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
        }
    case "":
        for fi28 := 0; fi28 < 9; fi28++ {
            v29 := GenS1{F0: false, F1: 4294967295}
            v30 := GenS2{F0: 1, F1: "hello"}
            _ = fi28
            _ = v29
            _ = v30
        }
    default:
        v31 := v2["hello"]
        _ = v31
    }
    switch "hello" {
    case "":
        if ((0.0 + 1e10) != 0.1) {
            var v32 int = v5
            v3 = -949
            _ = v32
        }
    case "foo bar":
        v5 += (v5 ^ (-(v5)))
    }
    var v33 uint = 18446744073709551615
    switch v5 {
    case -1:
        if false {
            v34 := GenS4{F0: 32767}
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 float64 = (1e10 / (((1e10 * float64((3.14159 + 0.0))) / (0.1 / float64(1.5))) / ((float64(0.1) * float64(3.14159)) + (3.14159 / 1.0))))
                var v37 uint64 = 42
                switch v5 {
                case 451:
                    v5 = v5
                case -390:
                    v33 += (^(v33))
                }
                _ = fi35
                _ = v36
                _ = v37
            }
            fmt.Println(v5, v1)
            _ = v34
        } else {
            v38 := GenS5{F0: v5, F1: 1, F2: -1}
            fmt.Println(v33)
            _ = v38
        }
    case -582:
        fmt.Println(v5, v5, v5)
    case 483:
        _ = len(v4)
    }
    v39 := GenS4{F0: 32767}
    if false {
        var v40 int8 = 1
        _ = v40
    }
    v1 = (v1 ^ -367)
    v33 -= v33
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v5
}
