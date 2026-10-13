package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
func genFn0(p0 uint64) int {
    p0 = p0
    var v6 string = "\n\t"
    _ = p0
    _ = v6
    return 1
}
func genFn1() uint16 {
    var v7 uint8 = 128
    if (!true) {
        v7 += v7
        if (true || (-2147483648 < 0)) {
            v8 := []int{890}
            _ = v8
        } else {
            v9 := GenS1{F0: 700, F1: -9223372036854775808}
            v10 := []int{}
            if len(v10) > 0 {
                v10[0] = -9223372036854775808
            }
            _ = v10
            _ = v9
        }
    } else {
        var v11 float32 = (1.5 + (0.0 + -2.25))
        v12 := GenS1{F0: 102, F1: -232}
        v7 -= v7
        _ = v11
        _ = v12
    }
    switch ("a" + "") {
    case "\n\t":
        if (false || (255 <= 42)) {
            fmt.Println(v7)
        }
    case "foo bar":
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := GenS1{F0: 598, F1: 42}
            _ = fi13
            _ = v14
        }
    case "a":
        if (false || false) {
            var v15 bool = (true || false)
            switch genFn0(18446744073709551615) {
            case 9223372036854775807:
                for fi16 := 0; fi16 < 1; fi16++ {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        var v18 uint64 = 1
                        fmt.Println(v18)
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
            case -475:
                if false {
                    switch 32768 {
                    case 32767:
                        switch 32768 {
                        case -492:
                            var v19 uint16 = 255
                            _ = v19
                        case 32767:
                            if false {
                                v15 = (784 == int64(v7))
                            }
                        case 2147483647:
                            var v20 int32 = -1
                            _ = v20
                        }
                    case -809:
                        if false {
                            var v21 uint8 = (v7 * (v7 / ((2) | 1)))
                            var v22 int64 = 9223372036854775807
                            _ = v21
                            _ = v22
                        }
                    }
                    v23 := []int{0, -142, 4294967295}
                    _ = v23
                }
            }
            for fi24 := 0; fi24 < 20; fi24++ {
                var v25 int8 = 103
                _ = fi24
                _ = v25
            }
            _ = v15
        }
    }
    var v26 int = 938
    v27 := GenS1{F0: 0, F1: 135}
    _ = v26
    _ = v27
    _ = v7
    return 42
}

func main() {
    var v1 int16 = -1
    var v2 bool = ((563 == 22) || (1 < 4294967295))
    var v3 uint64 = 1
    var v4 int32 = -1
    v5 := []int{136, 908, 46, 0}
    switch -414 {
    case 4294967295:
        var v28 int = -1
        _ = v28
    default:
        v4 -= int32(v3)
    }
    if v2 {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 int16 = -1
            switch fi29 {
            case -2147483648:
                v31 := GenS1{F0: 0, F1: 9223372036854775807}
                _ = v31
            case -1:
                var v32 int32 = (v4 + v4)
                _ = v32
            }
            v4 = 42
            _ = fi29
            _ = v30
        }
    }
    if v2 {
        if (!v2) {
            var v33 uint = 42
            _ = v33
        } else {
            var v34 uint64 = (v3 - (v3 % ((v3) | 1)))
            v3 = (^(v34))
            _ = v34
        }
        var v35 int16 = v1
        v36 := GenS1{F0: 42, F1: 42}
        _ = v35
        _ = v36
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        v5 = append(v5, genFn0((v3 >> ((255) & 63))))
        var v37 uint32 = 0
        v4 = (v4 / ((0) | 1))
        _ = v37
    } else {
        if len(v5) > 0 {
            v5[0] = genFn0((^(v3)))
        }
        v5 = append(v5, -9223372036854775808)
        var v38 int = 193
        _ = v38
    }
    v4 -= (v4 % ((v4) | 1))
    v3 -= v3
    fmt.Println(v4, v1, v4)
    fmt.Println(v2)
    for fi39 := 0; fi39 < 7; fi39++ {
        v5 = append(v5, fi39)
        _ = fi39
    }
    var v40 string = "a"
    v2 = ((!true) && ((int(v4) == 4294967295) || (1 != genFn1())))
    v41 := GenS1{F0: 0, F1: 834}
    v3 = (v3 % ((v3) | 1))
    if v2 {
        if (!(42 == (v4 << ((42) & 63)))) {
            var v42 int64 = 42
            _ = v42
        }
    }
    var v43 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v5
}
