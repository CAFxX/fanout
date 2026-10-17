package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint64; F1 int64; F2 uint }
func genFn0(p0 int, p1 int8, p2 uint16) string {
    fmt.Println(p0, p1)
    var v6 int8 = 0
    v7 := []int{}
    v8 := GenS1{F0: 913}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return ((("\n\t" + "") + "foo bar") + ("foo bar" + "hello"))
}
func genFn1() int64 {
    if false {
        switch ("" + "") {
        case "":
            var v9 string = ("foo bar" + "a")
            _ = v9
        case "hello":
            var v10 float64 = (-2.25 - (-2.25 - (0.1 - (1e10 / (float64(3.14159) + ((((-2.25 / 0.1) * 0.1) / (-2.25 / 3.14159)) + 3.14159))))))
            _ = v10
        case "a":
            v11 := GenS1{F0: -493}
            _ = v11
        }
        var v12 float32 = -2.25
        _ = v12
    }
    if (!true) {
        var v13 uint = 255
        fmt.Println(v13)
        _ = v13
    }
    v14 := GenS2{F0: 42, F1: 65535}
    v15 := GenS2{F0: 0, F1: 255}
    var v16 float32 = (3.14159 + ((0.0 * -2.25) / 1.0))
    _ = v14
    _ = v15
    _ = v16
    return -1
}

func main() {
    var v1 int = 983
    var v2 uint64 = 4294967295
    var v3 int32 = 1
    var v4 uint16 = 314
    var v5 int = (v1 - int(v3))
    v2 = uint64(v4)
    fmt.Println(v1, v5, v1)
    if ("\n\t" == "foo bar") {
        if false {
            v17 := []int{v5, v5}
            _ = len(v17)
            v2 -= (v2 - v2)
            _ = v17
        }
    } else {
        if true {
            var v18 uint = 154
            v19 := []int{(v1 * v1), (-(v1)), (v1 ^ (v5 | v1)), v1}
            if false {
                fmt.Println(v3, v4, v5)
                if (766 < 151) {
                    fmt.Println(v3, v4, v1)
                    for fi20 := 0; fi20 < 6; fi20++ {
                        switch "" {
                        case "\n\t":
                            var v21 int32 = 1
                            _ = v21
                        case "foo bar":
                            v22 := GenS2{F0: -128, F1: (v18 + 4294967295)}
                            _ = v22
                        case "a":
                            v1 -= 32768
                        }
                        if false {
                            fmt.Println(v5)
                        } else {
                            v2 += v2
                            v2 += 4294967295
                            for fi23 := 0; fi23 < 18; fi23++ {
                                v19 = append(v19, -2147483648)
                                if len(v19) > 0 {
                                    v19[0] = fi23
                                }
                                if (4294967295 != 944) {
                                    v1 += 42
                                    v3 = (v3 ^ (v3 & (v3 + (v3 ^ (v3 << ((18446744073709551615) & 63))))))
                                    var v24 int8 = 127
                                    _ = v24
                                } else {
                                    var v25 float64 = 1e10
                                    var v26 uint8 = 0
                                    _ = v25
                                    _ = v26
                                }
                                _ = fi23
                            }
                        }
                        switch (^(fi20)) {
                        case 32767:
                            if true {
                                v27 := GenS3{F0: 42, F1: 0, F2: (v18 | v18)}
                                _ = v27
                            } else {
                                var v28 uint16 = 1
                                v18 -= (v18 ^ v18)
                                if false {
                                    var v29 int16 = 1000
                                    _ = v29
                                }
                                _ = v28
                            }
                        case 2147483647:
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        case 17:
                            v5 = (v1 % (((v1 | (fi20 | (^(v1))))) | 1))
                        default:
                            if len(v19) > 0 {
                                v19[0] = 4294967295
                            }
                        }
                        _ = fi20
                    }
                } else {
                    if len(v19) > 0 {
                        v19[0] = int(v2)
                    }
                    v19 = append(v19, (-(v5)))
                    v2 = v2
                }
            }
            _ = v18
            _ = v19
        }
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        fmt.Println(v1, v4)
        switch "\n\t" {
        case "a":
            v31 := GenS3{F0: v2, F1: -778, F2: 255}
            _ = v31
        case "\n\t":
            v2 = (^(v2))
        default:
            var v32 int32 = v3
            _ = v32
        }
        for fi33 := 0; fi33 < 11; fi33++ {
            v34 := []int{}
            v5 = fi33
            _ = fi33
            _ = v34
        }
        _ = fi30
    }
    var v35 uint32 = 2147483648
    var v36 int64 = -9223372036854775808
    var v37 int32 = 0
    var v38 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
