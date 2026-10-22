package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint32; F1 int16; F2 int16 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint16; F1 int16; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 int8 }
func genFn0(p0 uint32, p1 uint32) int {
    v7 := GenS1{F0: 0, F1: 1}
    var v8 float64 = 3.14159
    if (42 != 0) {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 float64 = v8
            if (-1 != 0) {
                if true {
                    v11 := GenS1{F0: 242, F1: 4294967295}
                    v12 := []int{127, fi9, fi9, fi9}
                    _ = v11
                    _ = v12
                }
                v8 += v10
                var v13 int = (fi9 ^ 42)
                _ = v13
            }
            var v14 int32 = 981
            _ = fi9
            _ = v10
            _ = v14
        }
    } else {
        var v15 uint64 = 0
        _ = v15
    }
    var v16 float64 = float64(3.14159)
    var v17 int16 = -1
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 float32 = (0.1 * (float32(((((float64(1e10) + 3.14159) * -2.25) + (((3.14159 - -2.25) + 0.1) / 1.0)) - (3.14159 - ((((((((1e10 - -2.25) - 1e10) - (((((1e10 - 3.14159) / (float64((0.0 - (float32((3.14159 * 3.14159)) / 1.5))) - float64(float32(float64(((((0.1 - 0.1) * (3.14159 - 0.0)) + 3.14159) / 1e10)))))) / ((1.5 + float64((float32(-2.25) - 0.1))) + (3.14159 * 3.14159))) - 0.1) + 1.5)) + (1.5 - float64(1e10))) + -2.25) + float64(0.0)) / float64((3.14159 / 1.0))) * float64(0.0))))) + 1.5))
    var v2 float64 = 1e10
    var v3 uint32 = 4294967295
    v4 := []int{4294967295, 737, 4294967295, 32767}
    var v5 float32 = v1
    v6 := make(map[int]bool)
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS2{F0: (v3 & (v3 ^ v3)), F1: 42, F2: int16(v3)}
        var v20 uint64 = 4294967295
        var v21 string = "hello"
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    var v22 int32 = 2147483647
    var v23 uint = 52
    var v24 float32 = ((v1 - (v1 - 3.14159)) / v1)
    v25 := GenS3{F0: (false && (!true)), F1: v3, F2: 1}
    v3 -= v3
    _ = len(v4)
    if (((42 < 128) && false) && (!false)) {
        switch ("" + ("foo bar" + ("a" + ("foo bar" + "foo bar")))) {
        case "hello":
            if false {
                v1 += v1
                v2 += -2.25
                v26 := GenS4{F0: v22, F1: v3, F2: 1}
                _ = v26
            } else {
                _ = len(v4)
                v27 := GenS5{F0: uint16(v23), F1: -32768, F2: -499}
                v5 -= (v1 * float32(v2))
                _ = v27
            }
        default:
            switch 2147483647 {
            case 127:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 32767:
                fmt.Println(v23, v2, v3)
            case 426:
                v28 := GenS4{F0: (v22 * v22), F1: (v3 ^ (v3 | v3)), F2: -681}
                _ = v28
            }
        }
    } else {
        if true {
            var v29 string = ""
            _ = v29
        }
        for fi30 := 0; fi30 < 8; fi30++ {
            v3 += (v3 & (^(v3)))
            switch fi30 {
            case 937:
                for fi31 := 0; fi31 < 17; fi31++ {
                    v23 -= v23
                    var v32 uint32 = 1
                    _ = fi31
                    _ = v32
                }
            default:
                var v33 uint = 953
                _ = v33
            }
            for fi34 := 0; fi34 < 13; fi34++ {
                v35 := GenS4{F0: v22, F1: (^(v3)), F2: 1000}
                if ((true || false) && false) {
                    if (!false) {
                        var v36 float64 = -2.25
                        v23 = v23
                        v37 := GenS6{F0: (v23 | (v23 ^ (v23 / ((v23) | 1))))}
                        _ = v36
                        _ = v37
                    } else {
                        v38 := GenS4{F0: v22, F1: v3, F2: 1}
                        if true {
                            v2 = float64((3.14159 / v5))
                            v24 = 1.5
                        } else {
                            v6[9223372036854775807] = true
                        }
                        var v39 uint = v23
                        v3 = (v3 % ((459) | 1))
                        _ = v38
                        _ = v39
                    }
                    v40 := GenS7{F0: 256, F1: 127}
                    _ = v40
                } else {
                    v41 := GenS7{F0: 32767, F1: 127}
                    _ = v41
                }
                _ = fi34
                _ = v35
            }
            _ = fi30
        }
        var v42 uint32 = (^(v3))
        _ = v42
    }
    var v43 uint32 = 1
    v44 := GenS2{F0: v43, F1: -550, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
