package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 bool; F2 uint8 }
type GenS4 struct { F0 bool; F1 int64; F2 uint8 }
type GenS5 struct { F0 int; F1 int16 }
func genFn0(p0 uint16, p1 int64) uint32 {
    fmt.Println(p1)
    v4 := GenS1{F0: 752}
    var v5 int = 273
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2147483648
}
func genFn1(p0 uint, p1 string) int32 {
    var v6 int = 32767
    var v7 uint8 = 42
    switch "hello" {
    case "\n\t":
        var v8 uint = p0
        _ = v8
    default:
        fmt.Println(p0, v7, p1)
    }
    if false {
        if false {
            if (-270 < 0) {
                var v9 int = v6
                var v10 uint8 = (v7 - 128)
                p1 = p1
                _ = v10
                _ = v9
            } else {
                for fi11 := 0; fi11 < 3; fi11++ {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        v13 := []int{fi12}
                        var v14 float64 = -2.25
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = fi11
                }
            }
            v15 := GenS2{F0: -1}
            _ = v15
        } else {
            v16 := []int{42, (v6 * (-(v6)))}
            if ("\n\t" != p1) {
                var v17 int8 = -1
                _ = v17
            } else {
                v18 := GenS2{F0: 1}
                _ = v18
            }
            _ = v16
        }
    } else {
        if false {
            v6 = -2147483648
            v19 := []int{(^(v6))}
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := []int{689, (fi20 / ((-106) | 1))}
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 409
}

func main() {
    var v1 float32 = ((0.1 * ((float32((-2.25 * float64((float32((-2.25 - (1e10 * 1.5))) / 0.1)))) * ((((-2.25 / 1e10) * (0.1 - 0.1)) - 0.0) + 1e10)) + (0.1 + float32((1e10 / (float64(3.14159) - 1.5)))))) + 3.14159)
    var v2 uint8 = 10
    v3 := make(map[string]uint32)
    v3["\n\t"] = 946
    if false {
        var v22 int = 127
        v23 := v3[("" + "foo bar")]
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 uint = 65535
        _ = fi24
        _ = v25
    }
    switch "a" {
    case "foo bar":
        delete(v3, "hello")
    case "\n\t":
        v26 := []int{-2147483648, 42, 128}
        _ = v26
    }
    v27 := v3["\n\t"]
    var v28 float32 = (v1 / v1)
    v28 += 1.5
    v29 := []int{624}
    v1 = v1
    if ((v27 << ((uint(v27)) & 63)) >= (v27 % (((v27 * v27)) | 1))) {
        v30 := GenS3{F0: 0, F1: true, F2: v2}
        _ = v30
    } else {
        for fi31 := 0; fi31 < 3; fi31++ {
            switch (fi31 % ((42) | 1)) {
            case 0:
                v3["\n\t"] = (v27 % ((65535) | 1))
            case -2147483648:
                v27 -= genFn0(65535, 1)
            case 42:
                if (int32(fi31) > -2147483648) {
                    v28 = (v28 - float32(0.0))
                    for fi32 := 0; fi32 < 13; fi32++ {
                        if true {
                            v1 = v1
                        } else {
                            switch 32767 {
                            case 0:
                                switch (fi32 >> ((459) & 63)) {
                                case 4294967295:
                                    v28 -= 3.14159
                                case 32768:
                                    switch -960 {
                                    case -2147483648:
                                        v33 := GenS4{F0: false, F1: 42, F2: (v2 * 196)}
                                        _ = v33
                                    case 4294967295:
                                        var v34 int64 = 101
                                        _ = v34
                                    case -1:
                                        var v35 uint64 = 2
                                        _ = v35
                                    }
                                case 0:
                                    var v36 uint64 = 42
                                    _ = v36
                                }
                            }
                            delete(v3, ("a" + "foo bar"))
                        }
                        v1 -= v28
                        v2 = uint8(v27)
                        _ = fi32
                    }
                    var v37 string = ("hello" + "foo bar")
                    _ = v37
                }
            default:
                var v38 bool = (42 >= 42)
                _ = v38
            }
            switch 454 {
            case -652:
                var v39 int8 = 127
                _ = v39
            case -2147483648:
                v40 := GenS5{F0: int(v27), F1: 0}
                _ = v40
            default:
                var v41 bool = ((!((v1 / v1) == float32(-2.25))) || false)
                _ = v41
            }
            var v42 uint32 = v27
            _ = fi31
            _ = v42
        }
        var v43 bool = (true && false)
        v2 = v2
        _ = v43
    }
    if len(v29) > 0 {
        v29[0] = 32767
    }
    v2 += 128
    for fi44 := 0; fi44 < 18; fi44++ {
        v1 -= float32(1e10)
        fmt.Println(v2, v2)
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
