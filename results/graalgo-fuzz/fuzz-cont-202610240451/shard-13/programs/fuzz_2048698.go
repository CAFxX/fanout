package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 string }
func genFn0(p0 uint64, p1 uint8, p2 string) int16 {
    p0 = 608
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: (((((0.0 * 0.0) / (((3.14159 + ((((3.14159 - 1.5) + (1.5 - float32(((float64(1e10) - (0.1 + -2.25)) - 1.5)))) * (3.14159 + (-2.25 / -2.25))) - 3.14159)) * 1e10) + -2.25)) > ((0.1 * -2.25) / -2.25)) && false) && true), F1: p1}
        var v9 bool = true
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 2147483647 {
    case 127:
        v10 := []int{32767, -2147483648}
        _ = v10
    }
    switch ("hello" + ((p2 + p2) + (((p2 + "foo bar") + p2) + ""))) {
    case "":
        var v11 uint8 = (p1 / (((p1 * p1)) | 1))
        _ = v11
    case "\n\t":
        var v12 float32 = -2.25
        _ = v12
    default:
        for fi13 := 0; fi13 < 16; fi13++ {
            p2 = p2
            var v14 uint = 526
            if false {
                p2 = p2
                v15 := []int{fi13, -384, fi13, (fi13 * int(p1))}
                var v16 float64 = 1.5
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 8; fi17++ {
                    if (true || (!((^(fi17)) < (-(fi17))))) {
                        switch int(p1) {
                        case 1:
                            switch (-(fi13)) {
                            case -2147483648:
                                v18 := GenS1{F0: true, F1: (p1 + (p1 | (p1 - 1)))}
                                _ = v18
                            default:
                                var v19 bool = false
                                _ = v19
                            }
                        default:
                            v20 := []int{(fi13 | (fi17 + (-(fi17)))), (fi13 ^ fi17)}
                            _ = v20
                        }
                        var v21 int = 767
                        var v22 bool = (!true)
                        p2 = (p2 + p2)
                        _ = v21
                        _ = v22
                    } else {
                        fmt.Println(p2)
                    }
                    _ = fi17
                }
            }
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -377
}
func genFn1(p0 uint16, p1 int64) uint16 {
    fmt.Println(p1, p1)
    var v23 uint8 = 2
    var v24 int64 = 42
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return 2
}

func main() {
    var v1 int32 = -1
    var v2 uint16 = 690
    var v3 float64 = (0.0 / 1.0)
    var v4 uint = uint(v2)
    v5 := []int{}
    var v6 int = -9223372036854775808
    for fi25 := 0; fi25 < 16; fi25++ {
        if (0 > 0) {
            switch "a" {
            case "hello":
                v6 = -2147483648
            case "":
                var v26 float32 = (1.5 * 3.14159)
                _ = v26
            case "a":
                for fi27 := 0; fi27 < 9; fi27++ {
                    var v28 uint16 = (v2 & v2)
                    _ = fi27
                    _ = v28
                }
            default:
                var v29 uint16 = (v2 % ((188) | 1))
                _ = v29
            }
        } else {
            v6 += (v6 - -949)
        }
        v30 := GenS1{F0: true, F1: 2}
        _ = fi25
        _ = v30
    }
    if (1 > 128) {
        fmt.Println(v4, v4, v3)
        v31 := GenS1{F0: false, F1: 180}
        v1 -= (-(v1))
        _ = v31
    } else {
        if (genFn0(42, 255, "foo bar") >= -32768) {
            var v32 int = v6
            _ = v32
        } else {
            v3 += float64((0.0 * float32(1.5)))
            v2 = genFn1(v2, -711)
            var v33 uint32 = 42
            _ = v33
        }
    }
    if len(v5) > 0 {
        v5[0] = (v6 * (v6 + v6))
    }
    v3 = v3
    for fi34 := 0; fi34 < 2; fi34++ {
        v35 := GenS2{F0: "foo bar"}
        fmt.Println(v1)
        v4 += (v4 & (v4 << ((v4) & 63)))
        _ = fi34
        _ = v35
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v36 float64 = 1.5
    v36 = v3
    for fi37 := 0; fi37 < 16; fi37++ {
        v1 -= (v1 ^ (-(v1)))
        _ = fi37
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
