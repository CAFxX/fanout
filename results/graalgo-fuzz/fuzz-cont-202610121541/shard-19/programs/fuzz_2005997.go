package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
func genFn0(p0 int8) uint16 {
    var v6 int32 = -481
    if false {
        fmt.Println(v6, p0)
        for fi7 := 0; fi7 < 11; fi7++ {
            if false {
                if true {
                    v8 := GenS1{F0: 127, F1: v6, F2: 100}
                    switch "hello" {
                    case "hello":
                        var v9 uint16 = 290
                        _ = v9
                    case "":
                        if false {
                            var v10 bool = true
                            _ = v10
                        } else {
                            fmt.Println(v6)
                            v11 := []int{fi7, 4294967295, (fi7 >> ((255) & 63))}
                            if false {
                                var v12 int8 = -17
                                v11 = append(v11, (fi7 | (fi7 & int(v6))))
                                _ = len(v11)
                                _ = v12
                            }
                            _ = v11
                        }
                    default:
                        p0 -= 0
                    }
                    v13 := []int{(fi7 / ((127) | 1)), (fi7 + fi7)}
                    _ = v13
                    _ = v8
                }
            } else {
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 float64 = ((((((0.0 / ((((1.5 * ((0.1 / (0.1 * 1e10)) * 1e10)) + 0.0) * 1e10) - (-2.25 / (1.5 - -2.25)))) * 3.14159) - 0.1) + 0.0) - float64(3.14159)) - (0.0 + 1.5))
                    _ = fi14
                    _ = v15
                }
                switch (^(fi7)) {
                case 630:
                    var v16 float32 = (0.0 * (1.5 + -2.25))
                    _ = v16
                case -2147483648:
                    v6 = v6
                }
                p0 -= p0
            }
            _ = fi7
        }
        var v17 string = "hello"
        _ = v17
    } else {
        if (2 == 500) {
            if (v6 < v6) {
                var v18 uint16 = 841
                fmt.Println(v6, v18)
                v6 -= v6
                _ = v18
            }
        } else {
            switch 917 {
            case 784:
                p0 = (p0 << ((317) & 63))
            case -9223372036854775808:
                v6 = v6
            }
            v19 := []int{42, -2147483648, -387, -2147483648}
            _ = v19
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 uint8 = 128
            _ = fi20
            _ = v21
        }
    }
    _ = p0
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 2
    var v3 uint8 = 86
    var v4 float32 = -2.25
    var v5 int16 = -610
    if ((v2 - (v2 - 2)) != 65535) {
        for fi22 := 0; fi22 < 11; fi22++ {
            fmt.Println(v4, v4)
            var v23 int8 = 3
            var v24 uint16 = 255
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    if (32767 != 9223372036854775807) {
        fmt.Println(v2, v5)
    }
    var v25 float64 = 1e10
    if (!(-1 != int64(v5))) {
        v26 := []int{}
        fmt.Println(v2, v25)
        if ((v3 * 127) >= (v3 - v3)) {
            v25 = v25
        }
        _ = v26
    } else {
        if (!(v4 > float32(0.0))) {
            fmt.Println(v5, v4, v5)
            v3 = (v3 + v3)
            for fi27 := 0; fi27 < 16; fi27++ {
                v2 = (^(v2))
                _ = fi27
            }
        }
        v28 := []int{9223372036854775807, 32767}
        _ = v28
    }
    switch 0 {
    case 128:
        delete(v1, "a")
    default:
        fmt.Println(v3, v5)
    }
    var v29 uint = v2
    v5 = -756
    for fi30 := 0; fi30 < 17; fi30++ {
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 int64 = -218
            var v33 uint64 = 2
            _ = fi31
            _ = v32
            _ = v33
        }
        v1["foo bar"] = false
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
