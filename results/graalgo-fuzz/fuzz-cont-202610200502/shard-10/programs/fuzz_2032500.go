package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int32 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint32; F1 string; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint32; F1 uint8 }
func genFn0() uint32 {
    var v6 uint8 = 128
    v6 = (v6 & (v6 ^ (v6 - 127)))
    v7 := []int{0}
    if true {
        v6 -= (^(v6))
    }
    var v8 int8 = -87
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 bool = (981 != 1)
    var v2 uint = 1
    var v3 uint8 = 128
    var v4 uint16 = 32767
    var v5 int64 = 684
    v9 := GenS1{F0: 910, F1: (v5 - (^(v5))), F2: int32(v4)}
    for fi10 := 0; fi10 < 12; fi10++ {
        switch (fi10 ^ 127) {
        case -605:
            if false {
                switch (("" + "a") + "foo bar") {
                case "":
                    fmt.Println(v1)
                case "foo bar":
                    v11 := GenS1{F0: uint(v4), F1: 0, F2: -1}
                    _ = v11
                }
                if (-112 == 42) {
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v5 = (v5 / ((1) | 1))
                        v13 := GenS1{F0: v2, F1: (v5 % ((9223372036854775807) | 1)), F2: 42}
                        _ = fi12
                        _ = v13
                    }
                    fmt.Println(v1, v5, v5)
                }
            }
        case 1:
            switch "hello" {
            case "hello":
                var v14 string = "hello"
                _ = v14
            case "":
                v5 += v5
            default:
                v15 := GenS1{F0: (^(v2)), F1: (v5 | v5), F2: -2147483648}
                _ = v15
            }
        case 32767:
            var v16 int16 = -931
            _ = v16
        default:
            fmt.Println(v3, v1, v2)
        }
        v17 := GenS1{F0: 18446744073709551615, F1: (v5 | int64(v3)), F2: -697}
        v4 += (v4 >> ((v2) & 63))
        _ = fi10
        _ = v17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        v1 = ((^(v4)) >= v4)
        v19 := GenS2{F0: 2, F1: 0}
        v20 := GenS3{F0: -490, F1: 83}
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 uint16 = 456
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        v24 := []int{(fi23 | fi23), (fi23 % ((fi23) | 1))}
        if v1 {
            if len(v24) > 0 {
                v24[0] = -9223372036854775808
            }
        }
        switch (fi23 | (fi23 >> ((uint(v3)) & 63))) {
        case 128:
            var v25 float64 = 1e10
            _ = v25
        case 33:
            var v26 int = fi23
            _ = v26
        default:
            if false {
                v3 = (^(v3))
                if ((v1 || (v1 && (!v1))) && ((!((((0.0 + ((3.14159 - 3.14159) - float64((((1e10 / -2.25) - float32(3.14159)) / (1.5 - 0.0))))) / float64(float32((0.1 + 1.5)))) * float64((0.0 + 0.0))) >= -2.25)) && v1)) {
                    v27 := GenS4{F0: genFn0(), F1: ("" + (("foo bar" + "\n\t") + "foo bar")), F2: 285}
                    var v28 int = (^(fi23))
                    if (false || false) {
                        v2 -= 4294967295
                        switch fi23 {
                        case 42:
                            var v29 int16 = 482
                            _ = v29
                        default:
                            if len(v24) > 0 {
                                v24[0] = v28
                            }
                        }
                    } else {
                        switch ("" + ("a" + "\n\t")) {
                        case "\n\t":
                            v1 = v1
                        case "foo bar":
                            var v30 uint = v2
                            _ = v30
                        }
                        var v31 int = (fi23 - (v28 * 127))
                        for fi32 := 0; fi32 < 17; fi32++ {
                            v33 := GenS5{F0: -1, F1: (v3 + (v3 + 180)), F2: 65535}
                            fmt.Println(v31)
                            _ = fi32
                            _ = v33
                        }
                        _ = v31
                    }
                    _ = v27
                    _ = v28
                }
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
            }
        }
        _ = fi23
        _ = v24
    }
    v34 := GenS1{F0: (v2 & v2), F1: v5, F2: -2147483648}
    switch 71 {
    case 128:
        if ((true || (!false)) || (128 > 128)) {
            v5 -= (^(v5))
        }
    case 0:
        v3 = (^(v3))
    case -959:
        v3 = (v3 % ((v3) | 1))
    }
    var v35 uint = 42
    v36 := GenS6{F0: 858, F1: 1}
    v37 := []int{830, -125}
    v37 = append(v37, -437)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v9
}
