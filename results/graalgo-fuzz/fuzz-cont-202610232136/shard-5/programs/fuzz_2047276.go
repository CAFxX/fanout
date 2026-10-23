package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
func genFn0(p0 uint, p1 int32) uint32 {
    if (-17 > -358) {
        v4 := []int{998}
        _ = v4
    }
    if (1.5 >= 3.14159) {
        var v5 string = "a"
        fmt.Println(v5, v5, v5)
        _ = v5
    } else {
        if (239 != 42) {
            p0 = (p0 | (p0 / ((p0) | 1)))
            v6 := GenS1{F0: true, F1: 1}
            _ = v6
        } else {
            var v7 int32 = (p1 + p1)
            _ = v7
        }
        p0 = p0
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 128
    v3 := make(map[uint32]uint32)
    switch 127 {
    case 2147483647:
        fmt.Println(v2)
    default:
        var v8 bool = (true || (1 <= 0))
        _ = v8
    }
    var v9 int32 = 1
    var v10 int8 = 127
    var v11 int8 = (v10 * v10)
    v12 := GenS1{F0: (-9223372036854775808 <= -1), F1: 255}
    var v13 int = 4294967295
    v2 += v2
    if true {
        var v14 uint = 42
        fmt.Println(v11)
        _ = v14
    }
    var v15 int64 = 42
    fmt.Println(v2, v2)
    v9 = v9
    v16 := []int{v13}
    if (42 > -2147483648) {
        for fi17 := 0; fi17 < 9; fi17++ {
            for fi18 := 0; fi18 < 4; fi18++ {
                if ((((((-2.25 / -2.25) - 0.0) / 1.5) * -2.25) * 3.14159) == -2.25) {
                    switch "foo bar" {
                    case "":
                        if (false && (3.14159 == float64(1e10))) {
                            if (("foo bar" + ("hello" + "foo bar")) > "") {
                                var v19 int64 = v15
                                var v20 string = "foo bar"
                                _ = v19
                                _ = v20
                            } else {
                                v13 = (fi18 / (((fi17 % ((879) | 1))) | 1))
                                v13 = (v13 >> ((65535) & 63))
                                v2 -= v2
                            }
                        } else {
                            fmt.Println(v11)
                            if true {
                                fmt.Println(v2, v9)
                                v21 := v1["a"]
                                _ = v21
                            } else {
                                if len(v16) > 0 {
                                    v16[0] = fi17
                                }
                            }
                            fmt.Println(fi18)
                            var v22 uint8 = 59
                            _ = v22
                        }
                    default:
                        v11 = (v10 & (v10 >> ((79) & 63)))
                    }
                    v2 = 42
                } else {
                    var v23 uint16 = 1
                    if len(v16) > 0 {
                        v16[0] = fi18
                    }
                    _ = v23
                }
                v24 := GenS1{F0: false, F1: 2}
                _ = fi18
                _ = v24
            }
            v9 = (v9 ^ v9)
            var v25 bool = ((v13 % (((v13 >> ((650) & 63))) | 1)) == fi17)
            _ = fi17
            _ = v25
        }
        var v26 int = (v13 & (-(v13)))
        _ = v26
    }
    var v27 bool = (("" + "") >= ("hello" + ("" + ("" + ("hello" + ("a" + "foo bar"))))))
    if len(v16) > 0 {
        v16[0] = int(v11)
    }
    var v28 int32 = (v9 * v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v9
}
