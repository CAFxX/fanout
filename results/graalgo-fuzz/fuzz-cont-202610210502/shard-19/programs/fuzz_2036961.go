package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int; F1 string }
func genFn0() int {
    var v4 int64 = -1
    switch -83 {
    case 596:
        v5 := GenS1{F0: v4, F1: 0}
        _ = v5
    case -9223372036854775808:
        for fi6 := 0; fi6 < 12; fi6++ {
            v4 -= v4
            for fi7 := 0; fi7 < 16; fi7++ {
                v4 -= (v4 % (((v4 << ((65535) & 63))) | 1))
                if (!true) {
                    var v8 uint8 = 128
                    v4 = (v4 / (((v4 & (v4 % ((0) | 1)))) | 1))
                    _ = v8
                }
                _ = fi7
            }
            _ = fi6
        }
    case -1:
        v4 -= v4
    default:
        var v9 uint16 = 734
        _ = v9
    }
    if (v4 != (v4 << ((928) & 63))) {
        v4 = v4
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 bool = (65535 <= 165)
            for fi12 := 0; fi12 < 9; fi12++ {
                v4 = v4
                if v11 {
                    var v13 float32 = 0.0
                    var v14 uint = 133
                    _ = v13
                    _ = v14
                }
                var v15 float64 = 1.5
                _ = fi12
                _ = v15
            }
            fmt.Println(v11, v11)
            _ = fi10
            _ = v11
        }
    } else {
        v16 := GenS1{F0: v4, F1: 42}
        if (((157 != 4294967295) && ((!(!false)) && true)) || false) {
            v4 = (v4 / ((v4) | 1))
            var v17 int8 = 42
            var v18 uint32 = 1
            _ = v17
            _ = v18
        }
        _ = v16
    }
    _ = v4
    return 42
}

func main() {
    var v1 int64 = 830
    var v2 string = (("a" + "hello") + "hello")
    var v3 uint64 = 1
    v19 := []int{-401, 1, -737, 4294967295}
    if false {
        v20 := GenS2{F0: genFn0(), F1: ""}
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 string = "hello"
            v19 = append(v19, -2147483648)
            _ = fi21
            _ = v22
        }
        v1 += v1
        _ = v20
    }
    var v23 uint16 = 152
    var v24 uint = 749
    v25 := GenS1{F0: (v1 + v1), F1: 759}
    v2 = "a"
    if false {
        v23 = v23
        for fi26 := 0; fi26 < 15; fi26++ {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            var v27 uint64 = (v3 ^ (v3 % ((v3) | 1)))
            v19 = append(v19, genFn0())
            _ = fi26
            _ = v27
        }
    } else {
        var v28 uint32 = 1
        fmt.Println(v24)
        fmt.Println(v23, v28)
        _ = v28
    }
    if (int64(v23) < (v1 + v1)) {
        for fi29 := 0; fi29 < 16; fi29++ {
            v3 = v3
            var v30 int8 = 100
            _ = len(v19)
            _ = fi29
            _ = v30
        }
        v24 = 373
    } else {
        var v31 bool = ((-(v1)) != v1)
        if (-128 >= 1) {
            v32 := GenS1{F0: v1, F1: 153}
            var v33 float32 = (0.0 - float32(((1e10 / float64(-2.25)) + -2.25)))
            if (456 >= 2) {
                v19 = append(v19, 32768)
                if ((int32(v1) == -572) || (!((v2 + v2) >= v2))) {
                    v34 := GenS1{F0: v1, F1: 1}
                    var v35 uint8 = 1
                    if v31 {
                        var v36 int16 = -316
                        v24 = (v24 << (((^(v24))) & 63))
                        _ = v36
                    } else {
                        fmt.Println(v24)
                    }
                    _ = v34
                    _ = v35
                } else {
                    var v37 uint = v24
                    if (!false) {
                        switch 296 {
                        case -2147483648:
                            v38 := GenS1{F0: (v1 * (v1 - v1)), F1: 4294967295}
                            _ = v38
                        case 2147483647:
                            switch "" {
                            case "a":
                                v33 += v33
                            case "\n\t":
                                fmt.Println(v31)
                            default:
                                v31 = (v2 >= v2)
                            }
                        }
                    }
                    var v39 uint = v37
                    _ = v37
                    _ = v39
                }
                v19 = append(v19, int(v23))
            } else {
                var v40 uint = 42
                fmt.Println(v3, v31, v31)
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = v40
            }
            _ = v32
            _ = v33
        }
        if ((v24 - v24) <= v24) {
            var v41 int32 = -371
            _ = v41
        }
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
