package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int8 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint8; F2 int32 }
type GenS5 struct { F0 int64; F1 int64; F2 bool }
func genFn0(p0 bool) uint16 {
    p0 = (!(!(0.0 < 0.0)))
    v4 := []int{127, 2147483647, 201}
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 65535
}

func main() {
    var v1 uint32 = 286
    var v2 uint8 = 255
    var v3 bool = false
    var v5 int8 = -68
    switch -578 {
    case 128:
        if false {
            var v6 int32 = -2147483648
            fmt.Println(v5, v5)
            _ = v6
        } else {
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 int16 = -32768
                v2 += 0
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 18; fi9++ {
                v2 += v2
                v10 := []int{2147483647}
                v11 := GenS1{F0: true, F1: 255, F2: -1}
                _ = fi9
                _ = v10
                _ = v11
            }
        }
    default:
        v12 := GenS2{F0: v2, F1: v3}
        _ = v12
    }
    var v13 int8 = 0
    v3 = (int16(v1) == 0)
    var v14 float64 = (1.5 * (0.0 / -2.25))
    fmt.Println(v2, v14)
    if v3 {
        v15 := []int{}
        var v16 bool = false
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        v13 += (v13 * -128)
        _ = fi17
    }
    var v18 int64 = 0
    switch ("hello" + "") {
    case "":
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 int32 = -419
            for fi21 := 0; fi21 < 7; fi21++ {
                v22 := []int{(^(fi21))}
                var v23 uint8 = (v2 ^ 206)
                var v24 int16 = 1
                if len(v22) > 0 {
                    v22[0] = (fi21 / (((^(fi19))) | 1))
                }
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            fmt.Println(v1, v3)
            _ = fi19
            _ = v20
        }
    case "\n\t":
        v1 += 305
    default:
        fmt.Println(v5)
    }
    v25 := GenS3{F0: 1}
    if (42 > 749) {
        if false {
            if (1 == 457) {
                v13 -= (v5 / (((v5 << ((4294967295) & 63))) | 1))
            }
            if (!(v2 > v2)) {
                for fi26 := 0; fi26 < 10; fi26++ {
                    if (((0.0 * v14) + v14) == float64(1e10)) {
                        fmt.Println(fi26)
                        v3 = v3
                        var v27 uint16 = 42
                        _ = v27
                    } else {
                        fmt.Println(v2)
                    }
                    v2 += (v2 / ((v2) | 1))
                    v28 := GenS3{F0: (v1 % (((v1 + (v1 - v1))) | 1))}
                    _ = fi26
                    _ = v28
                }
            }
        } else {
            if v3 {
                fmt.Println(v18)
            }
        }
    } else {
        if (true || v3) {
            var v29 bool = (float32(0.1) >= (3.14159 / 1.5))
            _ = v29
        }
        v30 := GenS4{F0: -228, F1: 127, F2: 0}
        _ = v30
    }
    switch "a" {
    case "":
        var v31 bool = true
        _ = v31
    default:
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 uint64 = 4294967295
            for fi34 := 0; fi34 < 1; fi34++ {
                v1 -= v1
                _ = fi34
            }
            _ = fi32
            _ = v33
        }
    }
    if (836 < 42) {
        var v35 int8 = (v13 % ((-1) | 1))
        _ = v35
    } else {
        v36 := GenS5{F0: v18, F1: (v18 * (-(v18))), F2: (!((v2 - (^(v2))) >= v2))}
        for fi37 := 0; fi37 < 2; fi37++ {
            v5 -= 100
            v38 := []int{int(v5), int(v2), (-(fi37)), fi37}
            var v39 float32 = 1.5
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = v36
    }
    if v3 {
        var v40 int16 = 565
        var v41 int32 = int32(v2)
        _ = v40
        _ = v41
    }
    v42 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v42
    _ = v5
}
