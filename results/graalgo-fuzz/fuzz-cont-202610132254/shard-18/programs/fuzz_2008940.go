package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint64; F1 int }
func genFn0(p0 int64, p1 uint) uint64 {
    p0 -= (p0 | (p0 + 9223372036854775807))
    v5 := GenS1{F0: 313, F1: 32767, F2: 1}
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 uint64, p1 uint32, p2 int32) bool {
    var v6 float32 = (3.14159 - 0.0)
    v6 += 3.14159
    if ((^(p0)) <= 0) {
        if false {
            for fi7 := 0; fi7 < 12; fi7++ {
                if (true && (255 >= 1)) {
                    v8 := []int{}
                    _ = len(v8)
                    var v9 float32 = v6
                    _ = v8
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v11 := []int{fi10}
                        if ((v6 + float32(0.1)) > (v6 / v6)) {
                            fmt.Println(fi10, p2)
                        } else {
                            fmt.Println(p1, fi7)
                            var v12 bool = false
                            _ = v12
                        }
                        _ = fi10
                        _ = v11
                    }
                }
                var v13 uint32 = p1
                _ = fi7
                _ = v13
            }
            v14 := GenS1{F0: genFn0(0, 641), F1: 1, F2: -1}
            v6 = float32(3.14159)
            _ = v14
        } else {
            var v15 int = 597
            v15 += int(p1)
            v16 := []int{42, v15, 2147483647, -867}
            _ = v15
            _ = v16
        }
        v17 := []int{128, 9223372036854775807}
        _ = v17
    }
    switch "foo bar" {
    case "a":
        p1 = (p1 * (^(p1)))
    case "\n\t":
        if false {
            v18 := []int{310}
            if (!false) {
                fmt.Println(p1, p0, p1)
                var v19 bool = ((1 >= 2) || false)
                _ = v19
            }
            switch 4294967295 {
            case 0:
                if (((float64(v6) - float64(v6)) != 3.14159) || true) {
                    var v20 int = 9223372036854775807
                    v20 -= (v20 + -9223372036854775808)
                    _ = v20
                } else {
                    p2 = (p2 << ((0) & 63))
                    if ((!false) || (-610 == 42)) {
                        p1 -= uint32(p0)
                        fmt.Println(v6, p1)
                        _ = len(v18)
                    } else {
                        var v21 int32 = (-(p2))
                        _ = v21
                    }
                }
            case 32767:
                if (true || false) {
                    if (495 == 32767) {
                        var v22 string = "hello"
                        var v23 uint16 = 255
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS2{F0: (p0 + genFn0(598, 65535)), F1: -959}
                    _ = v24
                }
            case 1:
                var v25 uint64 = p0
                _ = v25
            default:
                fmt.Println(v6, v6, p1)
            }
            _ = v18
        }
    default:
        if ((false && false) || (!false)) {
            var v26 uint = 0
            _ = v26
        } else {
            var v27 string = ("a" + "hello")
            _ = v27
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ((!false) || false)
}

func main() {
    var v1 uint16 = 42
    var v2 uint64 = 676
    v3 := []int{127, -115, 9223372036854775807}
    var v4 int32 = -1
    if (false && (v1 <= v1)) {
        v4 += v4
    }
    if len(v3) > 0 {
        v3[0] = 128
    }
    var v28 uint = 65535
    v29 := GenS2{F0: 460, F1: -901}
    _ = len(v3)
    v28 = v28
    if len(v3) > 0 {
        v3[0] = -949
    }
    v28 += 270
    for fi30 := 0; fi30 < 8; fi30++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi30
    }
    var v31 int32 = (v4 + v4)
    v32 := GenS1{F0: 480, F1: v1, F2: 1}
    v33 := GenS2{F0: (v2 ^ v2), F1: -800}
    v3 = append(v3, 42)
    fmt.Println(v31)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
