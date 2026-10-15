package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int64; F1 int16; F2 bool }
type GenS6 struct { F0 int8; F1 int64; F2 int64 }
func genFn0(p0 int64) int16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: 2, F1: 0}
        _ = fi6
        _ = v7
    }
    p0 = p0
    if false {
        p0 = 860
        v8 := GenS2{F0: 87}
        _ = v8
    }
    v9 := GenS2{F0: 82}
    _ = p0
    _ = v9
    return 237
}

func main() {
    var v1 float64 = 0.0
    var v2 float64 = float64(0.0)
    v3 := make(map[int]uint32)
    var v4 uint16 = 32767
    v5 := make(map[string]uint32)
    var v10 int16 = 980
    fmt.Println(v4, v4, v10)
    v11 := []int{9223372036854775807}
    var v12 uint = 42
    if true {
        var v13 uint64 = 741
        _ = v13
    } else {
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint = (v12 & v12)
            if false {
                v16 := GenS3{F0: 65535, F1: ((("hello" + "\n\t") + (("" + (("\n\t" + "") + ("foo bar" + ("" + "")))) + "\n\t")) + "hello")}
                v4 = (^(v4))
                v15 = (v12 - 4294967295)
                _ = v16
            } else {
                if true {
                    v15 += v15
                    if false {
                        var v17 int8 = 1
                        v18 := GenS2{F0: v12}
                        if (false || (-2.25 > float32(v1))) {
                            var v19 int16 = (v10 - (v10 % ((v10) | 1)))
                            _ = v19
                        } else {
                            var v20 int = (fi14 - fi14)
                            fmt.Println(v20)
                            var v21 int16 = v10
                            _ = v20
                            _ = v21
                        }
                        _ = v17
                        _ = v18
                    } else {
                        v2 -= v2
                        if (2 >= 172) {
                            var v22 uint32 = 2147483647
                            _ = v22
                        } else {
                            if true {
                                var v23 uint8 = 139
                                fmt.Println(v15, fi14, v15)
                                _ = v23
                            }
                        }
                    }
                } else {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v12 += v12
                        v15 = (v12 & v12)
                        var v25 int64 = 9223372036854775807
                        _ = fi24
                        _ = v25
                    }
                }
                var v26 int = (^(fi14))
                _ = v26
            }
            v27 := GenS4{F0: 422, F1: -2147483648}
            _ = fi14
            _ = v15
            _ = v27
        }
        v2 = v2
        fmt.Println(v10, v4, v4)
    }
    if true {
        for fi28 := 0; fi28 < 8; fi28++ {
            v29 := GenS5{F0: 0, F1: genFn0(-52), F2: false}
            for fi30 := 0; fi30 < 5; fi30++ {
                _ = len(v11)
                _ = fi30
            }
            _ = fi28
            _ = v29
        }
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        delete(v3, fi31)
        v4 -= 399
        switch (fi31 % ((fi31) | 1)) {
        case 664:
            v32 := GenS4{F0: -1, F1: -633}
            _ = v32
        case 128:
            if false {
                if len(v11) > 0 {
                    v11[0] = (fi31 >> (((v12 - 42)) & 63))
                }
                v11 = append(v11, fi31)
                v10 -= v10
            }
        case 32768:
            if (0 >= 94) {
                for fi33 := 0; fi33 < 4; fi33++ {
                    v10 += v10
                    var v34 uint = 0
                    v4 = (v4 * (v4 % ((65535) | 1)))
                    _ = fi33
                    _ = v34
                }
                if (!false) {
                    var v35 uint64 = 0
                    v1 -= -2.25
                    var v36 uint = (v12 * (v12 + 18446744073709551615))
                    _ = v35
                    _ = v36
                }
                var v37 float32 = 0.1
                _ = v37
            }
        default:
            v2 = (v2 / (v1 - ((v1 + (((1e10 / (1e10 * v2)) + float64(1e10)) + (v1 + 0.1))) * (((v2 + (1.5 * 0.0)) - (v2 + v1)) - float64(3.14159)))))
        }
        _ = fi31
    }
    v38 := GenS6{F0: 100, F1: 926, F2: 42}
    var v39 uint64 = 233
    var v40 uint32 = 65535
    var v41 string = "\n\t"
    v3[0] = uint32(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
