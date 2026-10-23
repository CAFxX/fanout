package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint32) uint32 {
    var v5 float64 = (1e10 / ((((-2.25 + -2.25) - (1.5 / float64((1.5 / -2.25)))) / 1.0) + -2.25))
    v6 := []int{}
    _ = p0
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 int32 = 293
    v2 := []int{32768}
    var v3 int8 = -1
    v4 := []int{}
    var v7 uint16 = 463
    v7 -= v7
    var v8 string = "\n\t"
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int8 = (v3 + 127)
        var v11 bool = (!false)
        _ = fi9
        _ = v10
        _ = v11
    }
    v1 += v1
    v7 -= v7
    for fi12 := 0; fi12 < 2; fi12++ {
        if ((v1 << ((255) & 63)) == 547) {
            var v13 uint8 = 127
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint64 = 940
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        v16 := GenS1{F0: 0, F1: 0, F2: 1}
        v17 := GenS1{F0: -1, F1: 42, F2: 255}
        _ = fi12
        _ = v16
        _ = v17
    }
    var v18 float64 = (1.5 / (float64((0.1 / 3.14159)) + 1e10))
    var v19 uint = 4294967295
    fmt.Println(v1, v3)
    v3 -= (v3 & v3)
    if true {
        v20 := GenS1{F0: 32767, F1: (v19 << (((v19 | v19)) & 63)), F2: 255}
        _ = v20
    } else {
        if ((!(!true)) && (2147483647 > 2147483648)) {
            v21 := GenS2{F0: 618, F1: 702}
            v1 = (v1 - 127)
            _ = v21
        }
        switch 950 {
        case 0:
            for fi22 := 0; fi22 < 17; fi22++ {
                switch fi22 {
                case -632:
                    fmt.Println(v1, v1, fi22)
                }
                _ = fi22
            }
        case 654:
            if len(v2) > 0 {
                v2[0] = -9223372036854775808
            }
        default:
            v23 := GenS1{F0: 815, F1: (v19 | (^(v19))), F2: 216}
            _ = v23
        }
    }
    var v24 uint32 = 42
    var v25 uint16 = (v7 / ((255) | 1))
    var v26 string = v8
    for fi27 := 0; fi27 < 6; fi27++ {
        v8 = v26
        _ = fi27
    }
    if (255 <= 1) {
        v3 = (^(v3))
        var v28 int = 42
        for fi29 := 0; fi29 < 17; fi29++ {
            if true {
                switch -1 {
                case -9223372036854775808:
                    fmt.Println(v24, v25, v8)
                default:
                    for fi30 := 0; fi30 < 7; fi30++ {
                        v31 := GenS1{F0: -330, F1: 65535, F2: 2}
                        switch fi30 {
                        case -9223372036854775808:
                            if len(v4) > 0 {
                                v4[0] = v28
                            }
                        default:
                            switch v26 {
                            case "hello":
                                var v32 int = v28
                                _ = v32
                            case "":
                                var v33 int16 = 665
                                _ = v33
                            case "\n\t":
                                v18 = (3.14159 / float64(float32((v18 + 1.5))))
                            }
                        }
                        _ = fi30
                        _ = v31
                    }
                }
            } else {
                switch (fi29 + 32767) {
                case 127:
                    var v34 float64 = v18
                    _ = v34
                case 959:
                    v7 = v7
                case -1:
                    if ((9223372036854775807 != -1) || (65535 < v7)) {
                        v35 := GenS2{F0: 1000, F1: 18446744073709551615}
                        _ = v35
                    }
                default:
                    if len(v4) > 0 {
                        v4[0] = (fi29 * (fi29 | 1))
                    }
                }
                v3 += v3
                v7 = v25
            }
            for fi36 := 0; fi36 < 3; fi36++ {
                var v37 string = "foo bar"
                switch 42 {
                case 32768:
                    switch 610 {
                    case 128:
                        v38 := GenS1{F0: -32768, F1: v19, F2: 255}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v38
                    case 42:
                        v7 = uint16(v19)
                        var v39 int64 = -524
                        v39 -= (v39 / (((v39 ^ 1)) | 1))
                        v3 = 42
                        _ = v39
                    default:
                        v7 += (^(v7))
                    }
                }
                switch fi36 {
                case 32768:
                    var v40 string = ("\n\t" + (("foo bar" + v8) + "hello"))
                    _ = v40
                case 127:
                    var v41 uint64 = 2
                    _ = v41
                default:
                    var v42 uint = (v19 | (v19 + 255))
                    _ = v42
                }
                _ = fi36
                _ = v37
            }
            _ = fi29
        }
        _ = v28
    } else {
        var v43 uint = (v19 + 388)
        v43 = (v43 & v19)
        _ = v43
    }
    v2 = append(v2, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
