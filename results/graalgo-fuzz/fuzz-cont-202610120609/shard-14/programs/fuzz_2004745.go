package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0(p0 int16, p1 int32) bool {
    var v7 uint8 = 192
    v8 := []int{1, -216}
    v7 += 255
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p1 < 2147483647)
}

func main() {
    var v1 int16 = 42
    var v2 int64 = 42
    var v3 float64 = 3.14159
    var v4 int64 = 9223372036854775807
    var v5 int = -2147483648
    v6 := []string{"a", "\n\t", (("hello" + ("foo bar" + "a")) + "\n\t")}
    v2 -= 9223372036854775807
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 2, F1: 42}
        if (100 == int8(v1)) {
            var v11 int = 1
            _ = v11
        } else {
            v5 -= (fi9 & (fi9 & (v5 - (v5 | (^(v5))))))
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v5, v3)
    var v12 uint64 = 2
    v5 -= (v5 % ((v5) | 1))
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 float64 = ((v3 + float64(1.5)) * v3)
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if ((v4 % ((v4) | 1)) != 1) {
                switch fi13 {
                case 719:
                    var v15 int8 = -1
                    _ = v15
                case 9223372036854775807:
                    switch v5 {
                    case 0:
                        v4 = 0
                    default:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                case 0:
                    var v16 int = fi13
                    _ = v16
                }
                fmt.Println(fi13, v1)
                for fi17 := 0; fi17 < 2; fi17++ {
                    switch fi17 {
                    case -993:
                        if (834 != 732) {
                            var v18 int16 = v1
                            if (false || true) {
                                var v19 uint64 = 1
                                _ = v19
                            }
                            if ((1.5 >= (1e10 - 3.14159)) && false) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            _ = v18
                        }
                    case 27:
                        var v20 float32 = 3.14159
                        _ = v20
                    case 32768:
                        if len(v6) > 0 {
                            v6[0] = "\n\t"
                        }
                    default:
                        if (!(350 == 689)) {
                            if (!(false || (!(v3 <= float64((-2.25 + 0.1)))))) {
                                var v21 int16 = (-(v1))
                                switch "" {
                                case "\n\t":
                                    if false {
                                        var v22 uint32 = 65535
                                        _ = v22
                                    } else {
                                        var v23 uint8 = 42
                                        _ = v23
                                    }
                                default:
                                    var v24 bool = true
                                    _ = v24
                                }
                                _ = v21
                            } else {
                                if (0 < (v1 % ((v1) | 1))) {
                                    _ = len(v6)
                                    fmt.Println(v12, v2, v5)
                                }
                            }
                            v25 := GenS2{F0: (v1 / (((v1 ^ (v1 + (v1 | v1)))) | 1)), F1: 2147483647, F2: v12}
                            switch 344 {
                            case -998:
                                v26 := GenS3{F0: 32767, F1: (fi17 >> ((6) & 63))}
                                _ = v26
                            }
                            _ = v25
                        }
                    }
                    _ = fi17
                }
            }
        } else {
            switch v5 {
            case 42:
                v27 := GenS3{F0: 42, F1: (-(fi13))}
                _ = v27
            default:
                v28 := GenS1{F0: (v12 - 328), F1: 9}
                _ = v28
            }
            for fi29 := 0; fi29 < 4; fi29++ {
                _ = len(v6)
                v14 = 1.5
                fmt.Println(v14)
                var v30 int8 = -48
                _ = fi29
                _ = v30
            }
        }
        v31 := GenS3{F0: 832, F1: -9223372036854775808}
        _ = fi13
        _ = v14
        _ = v31
    }
    var v32 float64 = 0.1
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v33 uint = 2
    var v34 uint16 = 511
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
