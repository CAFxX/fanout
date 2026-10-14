package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint; F2 string }

func main() {
    var v1 float32 = 0.0
    var v2 int8 = -1
    var v3 int16 = 32767
    var v4 uint64 = 2
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(fi5, v2, v1)
        _ = fi5
    }
    var v6 uint8 = 0
    for fi7 := 0; fi7 < 13; fi7++ {
        v1 -= v1
        _ = fi7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int16 = (^(v3))
        fmt.Println(fi8, v2, v4)
        _ = fi8
        _ = v9
    }
    switch 779 {
    case 191:
        var v10 int = 98
        _ = v10
    default:
        if (3.14159 < (v1 - (((v1 / v1) - v1) - float32(0.1)))) {
            switch "" {
            case "foo bar":
                v11 := []int{32767, 0, 42, 128}
                _ = v11
            case "a":
                for fi12 := 0; fi12 < 1; fi12++ {
                    v1 = 0.0
                    v13 := []int{fi12, -351, -9223372036854775808, (fi12 | fi12)}
                    _ = fi12
                    _ = v13
                }
            case "\n\t":
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 int8 = 78
                    var v16 int = (fi14 | (fi14 + int(v4)))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
        } else {
            if (false || true) {
                v17 := GenS1{F0: -1, F1: v4, F2: 858}
                _ = v17
            } else {
                fmt.Println(v2, v3, v6)
            }
            if (-2.25 < float32(1.5)) {
                v18 := GenS1{F0: -462, F1: v4, F2: -958}
                v2 = (v2 % ((-55) | 1))
                _ = v18
            } else {
                for fi19 := 0; fi19 < 13; fi19++ {
                    v4 += uint64(v6)
                    if (v1 != v1) {
                        switch fi19 {
                        case -1:
                            v20 := []int{}
                            _ = v20
                        default:
                            if false {
                                if true {
                                    fmt.Println(v4)
                                    var v21 int = fi19
                                    _ = v21
                                }
                                var v22 int8 = -128
                                _ = v22
                            } else {
                                v23 := []int{}
                                v24 := GenS1{F0: 375, F1: 197, F2: 767}
                                switch fi19 {
                                case -262:
                                    if true {
                                        v1 += v1
                                        for fi25 := 0; fi25 < 14; fi25++ {
                                            v26 := GenS2{F0: 65535, F1: uint(v6), F2: "a"}
                                            var v27 int16 = 42
                                            switch fi25 {
                                            case -903:
                                                fmt.Println(v3)
                                            default:
                                                var v28 int16 = v3
                                                _ = v28
                                            }
                                            _ = fi25
                                            _ = v26
                                            _ = v27
                                        }
                                        var v29 uint8 = 128
                                        _ = v29
                                    }
                                default:
                                    var v30 string = ""
                                    _ = v30
                                }
                                _ = v23
                                _ = v24
                            }
                        }
                        if (!(-26 > v3)) {
                            fmt.Println(v3, v6)
                            v6 += v6
                        }
                        v4 += v4
                    }
                    _ = fi19
                }
                v31 := []int{127, 128, 32767}
                fmt.Println(v3)
                _ = v31
            }
            v32 := []int{0, -9223372036854775808, 9223372036854775807}
            _ = v32
        }
    }
    var v33 bool = (4294967295 != 51)
    v1 = v1
    if (!v33) {
        var v34 uint = uint(v4)
        v35 := GenS2{F0: 802, F1: v34, F2: ("foo bar" + "\n\t")}
        v36 := GenS2{F0: 1, F1: v34, F2: "\n\t"}
        _ = v34
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v6
}
