package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int32; F2 bool }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int32 }
func genFn0() uint64 {
    if true {
        for fi6 := 0; fi6 < 10; fi6++ {
            fmt.Println(fi6)
            v7 := []int{-9223372036854775808, (fi6 * (fi6 | fi6))}
            _ = fi6
            _ = v7
        }
        var v8 uint = 42
        _ = v8
    } else {
        switch 260 {
        case 127:
            if (!((!true) || (false && (!false)))) {
                if ("foo bar" != (("a" + ("a" + "")) + "\n\t")) {
                    var v9 float64 = float64(((-2.25 / (-2.25 + (1e10 - (1.5 + 3.14159)))) - ((float32((((3.14159 + 3.14159) - 0.1) - 1e10)) - 0.1) + 0.0)))
                    v10 := GenS1{F0: 1}
                    v11 := GenS1{F0: -864}
                    _ = v10
                    _ = v11
                    _ = v9
                }
                v12 := GenS2{F0: 254, F1: 0, F2: (0 == -32768)}
                if true {
                    if (-287 > 1) {
                        switch 32767 {
                        case -9223372036854775808:
                            v13 := GenS2{F0: -409, F1: -2147483648, F2: true}
                            _ = v13
                        case 42:
                            switch -817 {
                            case -90:
                                var v14 bool = (false && true)
                                _ = v14
                            case 128:
                                var v15 int32 = 42
                                _ = v15
                            case -2147483648:
                                var v16 uint32 = 2147483648
                                _ = v16
                            default:
                                for fi17 := 0; fi17 < 12; fi17++ {
                                    var v18 int16 = 1000
                                    var v19 int64 = 148
                                    v20 := GenS2{F0: (fi17 % ((fi17) | 1)), F1: 1, F2: false}
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                            }
                        case 32768:
                            if false {
                                if true {
                                    v21 := GenS3{F0: 0, F1: 18446744073709551615, F2: 65535}
                                    var v22 int64 = 1
                                    v22 = (-(v22))
                                    _ = v21
                                    _ = v22
                                } else {
                                    var v23 uint8 = 255
                                    _ = v23
                                }
                            } else {
                                v24 := []int{0, 0, 347, -358}
                                var v25 uint16 = 839
                                v25 = 1
                                _ = v24
                                _ = v25
                            }
                        }
                    } else {
                        for fi26 := 0; fi26 < 9; fi26++ {
                            v27 := []int{(fi26 / (((fi26 & (fi26 + fi26))) | 1))}
                            _ = fi26
                            _ = v27
                        }
                        v28 := []int{-1}
                        _ = v28
                    }
                }
                _ = v12
            }
        case 32768:
            var v29 uint16 = 65535
            _ = v29
        default:
            var v30 int8 = 10
            _ = v30
        }
    }
    var v31 int = -396
    v32 := GenS3{F0: 356, F1: 342, F2: 818}
    v31 = (v31 % (((v31 & (v31 >> ((960) & 63)))) | 1))
    var v33 string = (("hello" + "\n\t") + "hello")
    _ = v31
    _ = v32
    _ = v33
    return 474
}
func genFn1(p0 int, p1 int8, p2 int64) int8 {
    for fi34 := 0; fi34 < 20; fi34++ {
        p0 -= int(p1)
        p2 = (^(p2))
        p2 = (^(p2))
        _ = fi34
    }
    switch (("hello" + (("hello" + (((("foo bar" + ("foo bar" + ("foo bar" + ""))) + ("a" + "\n\t")) + "a") + "a")) + "\n\t")) + "") {
    case "\n\t":
        var v35 float32 = float32(-2.25)
        _ = v35
    case "hello":
        v36 := GenS4{F0: -206}
        _ = v36
    default:
        v37 := GenS1{F0: 295}
        _ = v37
    }
    var v38 int8 = p1
    var v39 float64 = float64(1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v38
    _ = v39
    return (v38 + -128)
}

func main() {
    var v1 uint8 = 127
    var v2 int8 = 127
    var v3 float32 = 0.0
    v4 := make(map[int]uint32)
    var v5 bool = (166 <= 0)
    delete(v4, 304)
    v5 = (!(!v5))
    delete(v4, 293)
    var v40 int16 = int16(v2)
    v5 = (v1 != v1)
    fmt.Println(v40)
    v1 -= v1
    _ = len(v4)
    var v41 int16 = v40
    v3 += v3
    v4[-2147483648] = 371
    var v42 int64 = -1
    v43 := GenS2{F0: 2147483647, F1: 1, F2: true}
    v44 := v4[1]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
