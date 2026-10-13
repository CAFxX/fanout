package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint32; F1 int; F2 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 uint; F2 int }
func genFn0(p0 uint16, p1 string, p2 int) uint {
    p2 = (p2 / ((p2) | 1))
    var v6 bool = (!false)
    switch (((("hello" + (p1 + "\n\t")) + "a") + ((("" + p1) + p1) + "hello")) + p1) {
    case "a":
        v7 := []int{}
        _ = v7
    default:
        v8 := GenS1{F0: (p2 | p2), F1: "a"}
        _ = v8
    }
    var v9 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 1
}

func main() {
    var v1 int64 = 558
    var v2 uint64 = 0
    var v3 float64 = 3.14159
    v4 := []uint32{88, 2, 634}
    var v5 float32 = ((0.0 - ((1.5 * ((0.1 * 0.1) / 0.1)) * float32(v3))) / (((0.0 / 3.14159) - (3.14159 + 1.5)) + 0.0))
    if false {
        v2 += (^(v2))
        if (42 > v2) {
            if true {
                if false {
                    fmt.Println(v5, v1)
                } else {
                    if true {
                        for fi10 := 0; fi10 < 11; fi10++ {
                            switch (^(fi10)) {
                            case 32767:
                                v11 := GenS1{F0: fi10, F1: ""}
                                _ = v11
                            case 584:
                                if len(v4) > 0 {
                                    v4[0] = 2147483648
                                }
                            case -776:
                                v3 = v3
                            default:
                                for fi12 := 0; fi12 < 4; fi12++ {
                                    var v13 bool = false
                                    _ = fi12
                                    _ = v13
                                }
                            }
                            if (v2 < 304) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v14 uint64 = v2
                                _ = v14
                            } else {
                                v15 := GenS2{F0: 42, F1: 0}
                                var v16 float32 = -2.25
                                v17 := GenS2{F0: -1, F1: 868}
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            if false {
                                v3 -= v3
                            }
                            _ = fi10
                        }
                        v2 = v2
                        v18 := GenS3{F0: 757, F1: 32768, F2: -9223372036854775808}
                        _ = v18
                    } else {
                        var v19 float32 = 1e10
                        var v20 int = 0
                        var v21 int64 = v1
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    if (((-527 != -402) || (!((!false) || true))) && (!(0.0 != v5))) {
                        for fi22 := 0; fi22 < 7; fi22++ {
                            switch fi22 {
                            case 42:
                                if len(v4) > 0 {
                                    v4[0] = 65535
                                }
                            default:
                                var v23 int64 = (^(v1))
                                _ = v23
                            }
                            var v24 int64 = -9223372036854775808
                            _ = fi22
                            _ = v24
                        }
                        if true {
                            var v25 uint16 = 1
                            _ = v25
                        }
                    }
                    v5 = (float32(((float64(v5) - (v3 * (v3 - -2.25))) * 0.0)) * 0.0)
                }
            }
        } else {
            var v26 uint = 2
            var v27 int = 127
            _ = v26
            _ = v27
        }
        var v28 bool = (61 > 18446744073709551615)
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        v4 = append(v4, 2147483648)
        switch "a" {
        case "":
            for fi30 := 0; fi30 < 2; fi30++ {
                v2 = v2
                v2 -= (v2 - (^(v2)))
                _ = fi30
            }
        }
        for fi31 := 0; fi31 < 3; fi31++ {
            if true {
                var v32 int32 = 2147483647
                if true {
                    v5 += 1e10
                    v5 += v5
                }
                _ = v32
            } else {
                var v33 uint32 = 65535
                var v34 float32 = (((v5 / ((float32(((v3 + 3.14159) * (v3 + v3))) - v5) - (1e10 - 0.1))) * v5) * v5)
                v2 -= 845
                v35 := GenS4{F0: -32768}
                var v36 int16 = -49
                _ = v33
                _ = v34
                _ = v35
                _ = v36
            }
            v37 := GenS1{F0: fi29, F1: "\n\t"}
            _ = fi31
            _ = v37
        }
        _ = fi29
    }
    if (!false) {
        v4 = append(v4, 1)
        var v38 int8 = 0
        _ = v38
    }
    _ = len(v4)
    v3 = 1.5
    var v39 float32 = v5
    v40 := GenS4{F0: -32768}
    var v41 int16 = -398
    var v42 uint64 = 500
    v39 += 1.5
    v3 += float64(((v39 * v39) + ((v39 * (v5 - (v5 * ((-2.25 - (3.14159 / (v39 - v39))) + float32(float64(1.5)))))) / v5)))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v43 := GenS4{F0: 32767}
    v44 := GenS5{F0: 127, F1: 2, F2: 660}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
