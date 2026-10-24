package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 bool; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int8; F2 int32 }
type GenS5 struct { F0 int8; F1 bool; F2 uint }
func genFn0() int {
    v7 := []int{0}
    _ = len(v7)
    _ = v7
    return 306
}
func genFn1(p0 int8, p1 int32) int8 {
    v8 := GenS1{F0: p0}
    p0 = -28
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            p1 = (p1 * (-(p1)))
            var v11 int64 = -1
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return p0
}

func main() {
    var v1 int64 = 42
    v2 := []string{}
    var v3 uint = 98
    var v4 uint8 = 128
    v5 := []string{"hello", "foo bar", "\n\t", ("" + "a")}
    v6 := make(map[string]uint32)
    if false {
        for fi12 := 0; fi12 < 14; fi12++ {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 uint16 = 255
                var v15 int = genFn0()
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 uint = v3
            _ = fi12
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            if (-690 > (^(fi17))) {
                switch "hello" {
                case "\n\t":
                    if (3.14159 == 3.14159) {
                        _ = len(v5)
                        v1 = 949
                    }
                case "a":
                    v3 = v3
                default:
                    for fi18 := 0; fi18 < 15; fi18++ {
                        v4 += (v4 % (((v4 % (((v4 & 128)) | 1))) | 1))
                        _ = fi18
                    }
                }
                var v19 float32 = (1.5 + (0.1 - float32(1e10)))
                _ = v19
            }
            _ = fi17
        }
        if len(v5) > 0 {
            v5[0] = "a"
        }
    }
    var v20 int32 = 0
    fmt.Println(v20)
    if false {
        fmt.Println(v4, v4, v1)
    } else {
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
        fmt.Println(v4, v4, v20)
        for fi21 := 0; fi21 < 16; fi21++ {
            if (!false) {
                var v22 uint = 2
                v23 := GenS1{F0: 1}
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 15; fi24++ {
                _ = len(v2)
                v20 -= v20
                v25 := GenS1{F0: 0}
                var v26 int32 = (v20 >> (((v3 & (v3 | (v3 / ((v3) | 1))))) & 63))
                var v27 uint8 = 0
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            _ = fi21
        }
    }
    var v28 int = genFn0()
    if true {
        if (!true) {
            v4 -= (v4 + (v4 / (((v4 & v4)) | 1)))
            delete(v6, "")
            var v29 uint32 = 4294967295
            _ = v29
        }
        var v30 int16 = 296
        for fi31 := 0; fi31 < 7; fi31++ {
            var v32 bool = true
            v3 -= (^(v3))
            _ = fi31
            _ = v32
        }
        _ = v30
    }
    delete(v6, ((("a" + "\n\t") + ("" + "a")) + "\n\t"))
    var v33 int = (v28 % (((-(v28))) | 1))
    if false {
        v33 += genFn0()
        v5 = append(v5, "foo bar")
    } else {
        for fi34 := 0; fi34 < 9; fi34++ {
            v35 := GenS2{F0: (fi34 ^ fi34), F1: (!true), F2: (!true)}
            v36 := GenS2{F0: -9223372036854775808, F1: (float64(((float32(float64(float32(-2.25))) * 0.1) - -2.25)) == 1e10), F2: true}
            switch (v28 + (-(v33))) {
            case 127:
                v28 = (^(fi34))
            case 32767:
                var v37 float64 = (3.14159 / 0.1)
                _ = v37
            default:
                switch "\n\t" {
                case "":
                    for fi38 := 0; fi38 < 16; fi38++ {
                        var v39 int8 = 42
                        _ = fi38
                        _ = v39
                    }
                case "foo bar":
                    var v40 bool = (false || (((1.5 - 0.0) / float64(1e10)) > (0.1 + (1.5 + 1.5))))
                    _ = v40
                case "hello":
                    v20 = -263
                }
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        switch (v33 / ((4294967295) | 1)) {
        case 4294967295:
            v4 = (v4 << (((v3 | 2)) & 63))
        case 32768:
            var v41 int8 = -128
            _ = v41
        case 42:
            v42 := GenS3{F0: 255}
            _ = v42
        default:
            switch genFn0() {
            case 193:
                if true {
                    var v43 uint32 = 65535
                    v1 += v1
                    _ = v43
                } else {
                    if (3.14159 > ((0.0 - ((float64(((1e10 + float32(1.5)) + 1.5)) / float64(float32(((0.0 / (0.0 + (0.1 + (((0.0 + 0.1) * 3.14159) - ((((0.0 / (1e10 + float64(1.5))) - float64(0.0)) + ((float64(0.1) + (3.14159 + 1.5)) - (0.0 - 1.5))) + 1e10))))) + 1e10)))) * -2.25)) * (0.0 - 0.0))) {
                        fmt.Println(v20)
                        for fi44 := 0; fi44 < 13; fi44++ {
                            v45 := GenS4{F0: v20, F1: genFn1(-1, (v20 & (v20 * (-(v20))))), F2: (v20 & (^(v20)))}
                            v2 = append(v2, "")
                            var v46 int = (v28 << (((v3 & (v3 & 65535))) & 63))
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                    }
                }
            case 127:
                var v47 bool = true
                _ = v47
            case 1:
                var v48 uint32 = 697
                _ = v48
            }
        }
    }
    v4 = uint8(v3)
    v49 := GenS5{F0: 100, F1: false, F2: (v3 ^ v3)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
