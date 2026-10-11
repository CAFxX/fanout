package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int8; F1 int32; F2 int8 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16; F1 int32; F2 int32 }
type GenS7 struct { F0 int64 }
func genFn0() int16 {
    v5 := []int{-879, 42, -118}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v6 int16 = 42
    switch "" {
    case "foo bar":
        var v7 uint8 = 2
        _ = v7
    default:
        for fi8 := 0; fi8 < 1; fi8++ {
            v6 -= -32768
            v9 := GenS1{F0: int8(v6), F1: 255}
            _ = fi8
            _ = v9
        }
    }
    if len(v5) > 0 {
        v5[0] = -835
    }
    _ = v5
    _ = v6
    return (v6 % (((v6 % ((v6) | 1))) | 1))
}
func genFn1(p0 bool, p1 int32) bool {
    switch "hello" {
    case "foo bar":
        v10 := []int{1}
        _ = v10
    case "hello":
        v11 := []int{}
        _ = v11
    case "":
        var v12 int64 = 988
        _ = v12
    default:
        var v13 int16 = 1
        _ = v13
    }
    v14 := []int{-261, -266, -1, 1}
    p0 = (127 != 42)
    _ = p0
    _ = p1
    _ = v14
    return (0 <= 1000)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 float32 = -2.25
    var v4 int32 = -23
    delete(v1, 4294967295)
    v3 = v3
    if (!(4294967295 >= 2147483647)) {
        if (false || true) {
            var v15 bool = (906 == 2)
            v16 := GenS2{F0: 41, F1: v4, F2: -1}
            _ = v15
            _ = v16
        }
    } else {
        if (1.5 < 1e10) {
            if false {
                for fi17 := 0; fi17 < 12; fi17++ {
                    for fi18 := 0; fi18 < 19; fi18++ {
                        var v19 float32 = v3
                        v4 += (-(v4))
                        if (-2.25 < ((-2.25 - v19) / (v3 - ((((v19 / (v19 - 1e10)) + 1e10) + 1e10) / (0.1 / ((float32((float64(v3) * (-2.25 / 1.5))) + ((((v3 * float32(0.0)) + (0.1 / v3)) * (v19 + ((-2.25 - (v19 + 0.0)) + 0.1))) + 3.14159)) + float32(1e10))))))) {
                            fmt.Println(fi18, fi17, v19)
                        }
                        v19 += ((float32((((3.14159 * -2.25) - (-2.25 - 1e10)) - -2.25)) / (float32((3.14159 / (0.1 + float64(((1e10 - (0.1 * 1e10)) + 1e10))))) * -2.25)) + 0.0)
                        v20 := GenS3{F0: 408, F1: 4294967295, F2: 836}
                        v4 = v4
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    _ = fi17
                }
            } else {
                fmt.Println(v3)
                v21 := GenS1{F0: 42, F1: 32767}
                if true {
                    v22 := GenS2{F0: 1, F1: -2147483648, F2: 1}
                    v23 := []int{484, 127}
                    v3 = 1.5
                    _ = v22
                    _ = v23
                }
                _ = v21
            }
            if (-368 <= -922) {
                v24 := GenS4{F0: 4294967295, F1: 0, F2: 4294967295}
                if (0 != 2) {
                    if (!true) {
                        v3 = float32(((((-2.25 - 1e10) / 1.0) / 1e10) / 0.1))
                        v4 += (v4 ^ -373)
                        v25 := []int{2147483647, 42}
                        _ = v25
                    } else {
                        v26 := GenS5{F0: 42}
                        _ = v26
                    }
                    var v27 string = ("" + ("" + ""))
                    fmt.Println(v4, v3)
                    _ = v27
                }
                _ = v24
            } else {
                v4 = v4
                var v28 uint64 = 0
                _ = v28
            }
        } else {
            v3 += v3
            for fi29 := 0; fi29 < 4; fi29++ {
                fmt.Println(v4, v3)
                _ = fi29
            }
        }
        var v30 int32 = (-(v4))
        var v31 uint32 = 884
        _ = v30
        _ = v31
    }
    if (849 < 2147483647) {
        var v32 uint16 = 32767
        _ = v32
    }
    fmt.Println(v3)
    fmt.Println(v4, v4)
    v4 -= v4
    var v33 float32 = v3
    v34 := GenS6{F0: 42, F1: v4, F2: (v4 * v4)}
    var v35 bool = (967 == 4294967295)
    fmt.Println(v33, v3, v3)
    var v36 int64 = 42
    v36 -= 42
    if v35 {
        for fi37 := 0; fi37 < 16; fi37++ {
            for fi38 := 0; fi38 < 18; fi38++ {
                v39 := GenS3{F0: (v36 / (((v36 % (((v36 >> ((42) & 63))) | 1))) | 1)), F1: 65535, F2: 361}
                _ = fi38
                _ = v39
            }
            fmt.Println(v4)
            _ = fi37
        }
    } else {
        for fi40 := 0; fi40 < 10; fi40++ {
            v3 = v3
            var v41 int8 = 127
            _ = fi40
            _ = v41
        }
        var v42 uint8 = 0
        v35 = v35
        _ = v42
    }
    var v43 float64 = ((1.5 / float64(v33)) / 3.14159)
    if v35 {
        if v35 {
            v44 := []int{25, 128, 448}
            v36 += (v36 & (v36 >> ((0) & 63)))
            _ = v44
        } else {
            v45 := []int{127, 0}
            switch 2147483647 {
            case 403:
                v3 -= (0.1 * ((1.5 * (0.1 / (v3 - ((v33 * -2.25) - float32(0.0))))) / v33))
            case 952:
                fmt.Println(v35)
            }
            _ = v45
        }
        delete(v1, -156)
        v35 = false
    } else {
        var v46 int16 = -536
        var v47 uint = 700
        v48 := GenS7{F0: v36}
        _ = v46
        _ = v47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v43
}
