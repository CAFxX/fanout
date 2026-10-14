package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint }
func genFn0(p0 uint32, p1 int8, p2 int) int {
    p1 += p1
    p1 = (p1 ^ (p1 >> ((501) & 63)))
    _ = p0
    _ = p1
    _ = p2
    return (p2 << ((540) & 63))
}
func genFn1() uint8 {
    var v6 int16 = 979
    switch ("foo bar" + ("" + ("foo bar" + "a"))) {
    case "a":
        v7 := []int{9223372036854775807, -1}
        _ = v7
    case "foo bar":
        if false {
            var v8 int32 = -1
            v8 = v8
            _ = v8
        }
    case "\n\t":
        var v9 uint32 = 2147483647
        _ = v9
    default:
        var v10 int = genFn0(2147483647, -128, -9223372036854775808)
        _ = v10
    }
    var v11 uint32 = 2147483647
    _ = v11
    _ = v6
    return 127
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int = 316
    var v3 uint16 = 0
    v4 := make(map[string]int)
    v5 := []uint32{}
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(v2)
        _ = fi12
    }
    var v13 uint16 = (^(v3))
    var v14 float32 = ((0.1 - 1.5) - float32(1e10))
    var v15 uint64 = 435
    fmt.Println(v1, v1, v15)
    v3 -= 32767
    for fi16 := 0; fi16 < 2; fi16++ {
        switch ("foo bar" + "") {
        case "foo bar":
            if (true && true) {
                v1 = v1
                for fi17 := 0; fi17 < 20; fi17++ {
                    fmt.Println(v15, v2, v2)
                    v18 := GenS1{F0: 42, F1: 178, F2: 2}
                    if (false || true) {
                        v19 := GenS1{F0: 427, F1: 127, F2: 825}
                        _ = v19
                    }
                    if false {
                        var v20 uint16 = v13
                        fmt.Println(fi16, fi17)
                        v2 += int(v1)
                        if (!(!false)) {
                            v15 = v15
                            delete(v4, "foo bar")
                        } else {
                            v4[(("" + ("foo bar" + "")) + "foo bar")] = (fi17 << ((0) & 63))
                            if (-1 <= -1) {
                                var v21 int32 = 0
                                var v22 int8 = 0
                                var v23 float64 = ((0.1 / 1.0) - (((((0.1 + -2.25) / 3.14159) / (((0.1 * (float64(v14) * 0.1)) - 1e10) * 0.1)) * 0.0) / 1.5))
                                _ = v21
                                _ = v22
                                _ = v23
                            } else {
                                var v24 int16 = -1
                                v3 += (v13 << ((4294967295) & 63))
                                _ = v24
                            }
                            fmt.Println(fi16, fi17, v20)
                        }
                        var v25 uint64 = 18446744073709551615
                        var v26 int16 = -981
                        var v27 int16 = int16(fi17)
                        _ = v20
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = fi17
                    _ = v18
                }
            }
        case "":
            var v28 bool = ((v1 > (v1 ^ 0)) && false)
            _ = v28
        case "a":
            delete(v4, "a")
        default:
            var v29 uint32 = (v1 | v1)
            _ = v29
        }
        _ = fi16
    }
    var v30 int32 = 1
    switch (v2 * (-(v2))) {
    case 128:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        if len(v5) > 0 {
            v5[0] = v1
        }
    }
    if false {
        v14 = ((v14 + v14) - 0.1)
        if ((!(!true)) && true) {
            var v31 uint32 = (^(v1))
            _ = len(v4)
            var v32 uint8 = 87
            _ = v31
            _ = v32
        }
        var v33 uint64 = (v15 % ((176) | 1))
        _ = v33
    } else {
        v14 = v14
        switch "" {
        case "":
            for fi34 := 0; fi34 < 9; fi34++ {
                fmt.Println(v30, v2)
                if (!(false || ((!(-128 == 42)) || (uint8(v1) > 0)))) {
                    var v35 uint16 = v13
                    fmt.Println(v35)
                    _ = v35
                }
                for fi36 := 0; fi36 < 18; fi36++ {
                    v3 = v13
                    v37 := GenS1{F0: 65535, F1: 128, F2: 2}
                    _ = fi36
                    _ = v37
                }
                _ = fi34
            }
        default:
            _ = len(v5)
        }
    }
    fmt.Println(v3, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
