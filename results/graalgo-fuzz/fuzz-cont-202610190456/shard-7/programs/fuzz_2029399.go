package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint8 }
func genFn0(p0 int8) int64 {
    var v7 float64 = (1e10 / (1.5 * float64((((1.5 / ((((((((1.5 / float32(3.14159)) * (1.5 - 0.1)) - 1.5) * 0.1) + (0.0 + float32(-2.25))) / 0.1) / 3.14159) + (0.0 / 1e10))) - float32((0.0 * -2.25))) + 0.0))))
    v8 := []int{9223372036854775807, -9223372036854775808, 127, 322}
    _ = p0
    _ = v7
    _ = v8
    return 460
}
func genFn1(p0 string, p1 int8) int64 {
    var v9 int8 = p1
    p0 = ""
    v10 := GenS1{F0: 2147483648, F1: -2147483648}
    for fi11 := 0; fi11 < 16; fi11++ {
        if true {
            for fi12 := 0; fi12 < 2; fi12++ {
                v13 := []int{(^(fi12)), fi11, (fi11 / (((fi11 % ((fi12) | 1))) | 1)), fi12}
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := GenS2{F0: 32767}
                var v16 float64 = 0.0
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 uint32 = 412
                v19 := GenS3{F0: 42, F1: uint16(v18)}
                var v20 int8 = p1
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            v21 := GenS1{F0: 2147483647, F1: -225}
            v22 := []int{fi11, fi11, 32768, (-(fi11))}
            _ = v21
            _ = v22
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 551
}

func main() {
    var v1 int32 = 42
    var v2 string = ("hello" + "a")
    var v3 int = 4294967295
    v4 := make(map[string]bool)
    var v5 bool = false
    var v6 float64 = 3.14159
    fmt.Println(v1)
    v5 = (571 != -102)
    var v23 int8 = 1
    if true {
        fmt.Println(v2)
    } else {
        fmt.Println(v23, v2, v23)
        var v24 uint = 2
        v25 := GenS4{F0: (v24 * (v24 + v24)), F1: 42, F2: 2147483648}
        _ = v24
        _ = v25
    }
    var v26 uint64 = 158
    v3 = (v3 * v3)
    if false {
        v26 += (v26 + v26)
        var v27 int64 = 42
        _ = v27
    } else {
        v3 = v3
    }
    delete(v4, "\n\t")
    v28 := []int{(-(v3)), v3}
    var v29 int16 = 1000
    var v30 int = (-(v3))
    v5 = ((!true) || true)
    var v31 float64 = 3.14159
    if (1 != 255) {
        if (float64(-2.25) == v31) {
            var v32 int = v3
            var v33 int = (v3 % ((862) | 1))
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 7; fi34++ {
            if ((!(v5 || v5)) && true) {
                var v35 float32 = 1e10
                if ((1 != 42) && v5) {
                    v36 := GenS3{F0: 255, F1: 42}
                    if (v1 <= -2147483648) {
                        for fi37 := 0; fi37 < 8; fi37++ {
                            _ = len(v4)
                            switch (v3 + -697) {
                            case 2147483647:
                                var v38 bool = v5
                                _ = v38
                            case 1:
                                fmt.Println(v26)
                            default:
                                v39 := GenS5{F0: 973, F1: 138, F2: 2}
                                _ = v39
                            }
                            _ = fi37
                        }
                        fmt.Println(v6, v35)
                    }
                    _ = v36
                } else {
                    for fi40 := 0; fi40 < 14; fi40++ {
                        if false {
                            fmt.Println(v1, v26, v5)
                        }
                        var v41 uint = 255
                        _ = fi40
                        _ = v41
                    }
                    fmt.Println(v6, v5)
                }
                _ = v35
            }
            _ = fi34
        }
    } else {
        var v42 float64 = v6
        _ = v42
    }
    if len(v28) > 0 {
        v28[0] = int(v29)
    }
    v26 += v26
    var v43 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
