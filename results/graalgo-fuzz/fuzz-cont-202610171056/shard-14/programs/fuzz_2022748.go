package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0() int64 {
    switch 9223372036854775807 {
    case 1:
        var v5 uint8 = 1
        _ = v5
    default:
        var v6 uint32 = 2
        _ = v6
    }
    v7 := GenS1{F0: 9223372036854775807, F1: 65535}
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 int) uint8 {
    p0 -= p0
    var v8 int = (p0 << ((0) & 63))
    if (false && true) {
        if (1000 != 42) {
            v9 := GenS1{F0: -1, F1: 104}
            _ = v9
        }
        v8 -= (v8 + v8)
    } else {
        switch v8 {
        case 128:
            p0 += 127
        case 32768:
            var v10 uint32 = 1
            _ = v10
        }
        var v11 uint8 = 0
        _ = v11
    }
    _ = p0
    _ = v8
    return 121
}

func main() {
    v1 := []int{0, 127, 128, 745}
    var v2 int32 = -491
    var v3 uint16 = 255
    var v4 bool = false
    fmt.Println(v4)
    if len(v1) > 0 {
        v1[0] = -1
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 string = ("\n\t" + (("\n\t" + "a") + ("a" + (("hello" + "\n\t") + "hello"))))
        v13 = v13
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 uint16 = v3
            if false {
                switch 879 {
                case 32768:
                    v16 := GenS1{F0: -1, F1: v15}
                    _ = v16
                case 2147483647:
                    var v17 uint64 = 18446744073709551615
                    _ = v17
                }
                v4 = v4
                v18 := GenS1{F0: 1, F1: v15}
                _ = v18
            } else {
                switch (^(fi14)) {
                case 128:
                    var v19 uint16 = 490
                    _ = v19
                case 32768:
                    var v20 float32 = 0.0
                    _ = v20
                case -1:
                    var v21 int16 = -61
                    _ = v21
                default:
                    if v4 {
                        var v22 string = v13
                        if v4 {
                            var v23 int32 = (v2 | -2147483648)
                            _ = v23
                        }
                        if v4 {
                            switch "foo bar" {
                            case "":
                                v2 = v2
                            case "a":
                                var v24 float32 = 3.14159
                                _ = v24
                            }
                            _ = len(v1)
                            _ = len(v1)
                        } else {
                            var v25 string = ((("" + v13) + v13) + v22)
                            fmt.Println(fi14, v3)
                            _ = v25
                        }
                        _ = v22
                    } else {
                        v26 := GenS1{F0: 1, F1: v3}
                        v27 := GenS1{F0: 1, F1: v15}
                        switch v13 {
                        case "":
                            var v28 bool = (v4 && (!v4))
                            _ = v28
                        default:
                            v1 = append(v1, fi12)
                        }
                        _ = v26
                        _ = v27
                    }
                }
                v15 -= 65535
                v2 += 800
                var v29 float32 = 0.0
                fmt.Println(v15)
                var v30 int = (^(fi14))
                _ = v29
                _ = v30
            }
            v3 -= (v15 - (v15 | 942))
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 -= v3
        _ = fi31
    }
    if (1 <= -1) {
        v32 := GenS1{F0: -923, F1: (^(v3))}
        _ = v32
    } else {
        if true {
            if v4 {
                var v33 uint = 1
                var v34 uint32 = 2147483648
                if (v4 || v4) {
                    v34 -= (^(v34))
                    fmt.Println(v2)
                    v34 = (v34 * v34)
                } else {
                    v4 = (32767 >= 42)
                }
                _ = v33
                _ = v34
            }
        }
        var v35 int8 = 0
        _ = v35
    }
    if true {
        if len(v1) > 0 {
            v1[0] = 32767
        }
    } else {
        _ = len(v1)
        v3 += (v3 / ((v3) | 1))
        if len(v1) > 0 {
            v1[0] = 386
        }
    }
    if (!false) {
        var v36 string = "hello"
        if (0 > 18446744073709551615) {
            v1 = append(v1, -389)
            v2 = v2
        } else {
            switch 4294967295 {
            case -2147483648:
                for fi37 := 0; fi37 < 3; fi37++ {
                    var v38 int64 = 1
                    _ = fi37
                    _ = v38
                }
            case 42:
                v3 = (v3 % ((v3) | 1))
            case 128:
                v36 = v36
            default:
                fmt.Println(v36, v4)
            }
            v4 = (0 >= 1)
        }
        _ = v36
    } else {
        var v39 float32 = ((float32((float64(0.1) / 1e10)) * 1.5) + 1e10)
        _ = v39
    }
    if true {
        v3 -= 0
        v3 -= v3
    } else {
        if v4 {
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v40 bool = v4
                _ = v40
            }
        }
        var v41 uint64 = 18446744073709551615
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
