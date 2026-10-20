package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int32; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint; F2 int }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 int, p1 bool, p2 int64) uint32 {
    var v7 float64 = -2.25
    switch "\n\t" {
    case "":
        if p1 {
            var v8 float64 = -2.25
            p2 -= 284
            _ = v8
        } else {
            switch (p0 - -9223372036854775808) {
            case 127:
                for fi9 := 0; fi9 < 10; fi9++ {
                    for fi10 := 0; fi10 < 13; fi10++ {
                        if (!((v7 / 1.0) > (v7 / v7))) {
                            v11 := []int{p0, (^(fi10)), (fi9 / ((fi10) | 1)), fi9}
                            _ = v11
                        }
                        _ = fi10
                    }
                    v12 := []int{(-(p0))}
                    _ = fi9
                    _ = v12
                }
            case 0:
                v13 := GenS1{F0: p2, F1: 1000}
                _ = v13
            default:
                fmt.Println(p0, p2, p1)
            }
            if (float32(v7) >= (float32(1.5) - 0.1)) {
                v7 = v7
            }
            if ((p2 / ((p2) | 1)) > -9223372036854775808) {
                if (128 <= 255) {
                    v14 := []int{}
                    v7 -= v7
                    _ = v14
                }
                var v15 int32 = 1
                _ = v15
            }
        }
    default:
        p1 = (!p1)
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 string = ("" + "")
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{}
    var v3 float32 = v1
    var v4 int = 1
    v5 := make(map[int]string)
    v6 := make(map[string]uint32)
    v2 = append(v2, 2)
    var v18 uint32 = 2
    v1 = v3
    var v19 int16 = 0
    var v20 bool = true
    for fi21 := 0; fi21 < 13; fi21++ {
        v22 := GenS2{F0: 524, F1: ((1e10 - float64((0.0 - ((3.14159 + ((v3 - 1e10) - float32(0.1))) + (3.14159 * v1))))) >= 1e10), F2: 128}
        if v20 {
            var v23 int = (v4 * 32767)
            _ = v23
        } else {
            for fi24 := 0; fi24 < 1; fi24++ {
                _ = len(v2)
                var v25 uint32 = (v18 >> ((18446744073709551615) & 63))
                _ = fi24
                _ = v25
            }
        }
        if (!(((v18 >> ((255) & 63)) > (v18 >> ((42) & 63))) && (!v20))) {
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 uint8 = 2
                _ = fi26
                _ = v27
            }
            switch ("a" + ("" + ("\n\t" + "foo bar"))) {
            case "hello":
                v28 := v6[(((("" + "foo bar") + (("\n\t" + "") + ("a" + "foo bar"))) + ("" + "foo bar")) + ("foo bar" + "hello"))]
                _ = v28
            case "\n\t":
                v29 := GenS2{F0: 42, F1: (v20 || (0 <= 32767)), F2: 127}
                _ = v29
            case "foo bar":
                for fi30 := 0; fi30 < 19; fi30++ {
                    if len(v2) > 0 {
                        v2[0] = (v18 ^ (^(v18)))
                    }
                    v20 = v20
                    v1 = 0.0
                    _ = fi30
                }
            default:
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 int32 = -1
                    var v33 float32 = (-2.25 + v3)
                    var v34 int = v4
                    switch -651 {
                    case -841:
                        if len(v2) > 0 {
                            v2[0] = (v18 / ((v18) | 1))
                        }
                    case -1:
                        if v20 {
                            v2 = append(v2, (v18 / ((2) | 1)))
                        }
                        fmt.Println(v4, v1, v20)
                    default:
                        v35 := GenS3{F0: 42}
                        _ = v35
                    }
                    v33 = (v3 / ((((0.1 + float32((((((float64(v1) * 3.14159) + (1.5 + (1.5 / 1.5))) / (((0.0 - (float64(((v1 * (float32(-2.25) + (v3 - float32(1e10)))) / (v3 + v33))) / -2.25)) + (0.1 + 0.0)) * 0.1)) / (-2.25 * 0.1)) * 1.5))) * (1e10 / (-2.25 + (0.1 / (v3 / v3))))) + v33) + 0.1))
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
            }
        }
        _ = fi21
        _ = v22
    }
    if false {
        for fi36 := 0; fi36 < 13; fi36++ {
            v19 += 42
            if len(v2) > 0 {
                v2[0] = (v18 / ((817) | 1))
            }
            v18 += v18
            _ = fi36
        }
    } else {
        if len(v2) > 0 {
            v2[0] = v18
        }
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        v38 := GenS1{F0: -9223372036854775808, F1: 1}
        v39 := GenS3{F0: 0}
        v40 := GenS4{F0: -1, F1: 509, F2: int(v19)}
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    v41 := GenS5{F0: 2, F1: 255}
    var v42 string = "\n\t"
    fmt.Println(v3, v18)
    if (!(v20 && v20)) {
        if (((128 != 241) || (!true)) || (genFn0(128, ((-2.25 - ((0.1 * (((((((3.14159 - 0.0) / ((3.14159 - 3.14159) + (float64(v1) - 0.0))) + (1e10 + 1e10)) * float64(v1)) / 1.0) - -2.25) / ((-2.25 - 0.0) + float64(v1)))) * (3.14159 - -2.25))) != 1e10), 42) >= v18)) {
            if false {
                var v43 int32 = 133
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                delete(v5, (v4 << ((156) & 63)))
                _ = v43
            }
            v42 = (("\n\t" + ("foo bar" + "")) + v42)
        } else {
            for fi44 := 0; fi44 < 9; fi44++ {
                if false {
                    v45 := GenS2{F0: 1, F1: (2147483647 != 2147483647), F2: 65}
                    v4 += fi44
                    _ = v45
                }
                var v46 uint8 = 0
                _ = fi44
                _ = v46
            }
        }
        switch v42 {
        case "\n\t":
            var v47 uint = 2
            _ = v47
        case "foo bar":
            v20 = true
        }
    }
    v48 := GenS3{F0: 42}
    var v49 float32 = 3.14159
    var v50 uint32 = v18
    for fi51 := 0; fi51 < 20; fi51++ {
        v52 := GenS1{F0: -9223372036854775808, F1: (v19 % ((v19) | 1))}
        for fi53 := 0; fi53 < 15; fi53++ {
            var v54 uint8 = 1
            if false {
                v54 -= 238
                fmt.Println(v19, v4, v20)
                if v20 {
                    switch ((v42 + ("" + ((v42 + v42) + "hello"))) + ("\n\t" + v42)) {
                    case "hello":
                        v55 := GenS6{F0: 105}
                        _ = v55
                    case "foo bar":
                        var v56 uint = 0
                        _ = v56
                    default:
                        v54 += v54
                    }
                    v57 := GenS7{F0: 255}
                    v58 := GenS5{F0: 0, F1: (^(v54))}
                    _ = v57
                    _ = v58
                } else {
                    fmt.Println(v18)
                }
                v54 = v54
                v59 := GenS2{F0: int32(v50), F1: false, F2: (v54 & v54)}
                _ = v59
            }
            var v60 float64 = 0.0
            _ = fi53
            _ = v54
            _ = v60
        }
        _ = fi51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
