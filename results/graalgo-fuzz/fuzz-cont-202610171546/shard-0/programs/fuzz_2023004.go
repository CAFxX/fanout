package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int32; F1 int32; F2 uint16 }
func genFn0() uint32 {
    var v6 uint64 = 1
    v7 := GenS1{F0: 127, F1: v6}
    _ = v6
    _ = v7
    return 604
}
func genFn1(p0 uint8, p1 bool) int16 {
    for fi8 := 0; fi8 < 18; fi8++ {
        p0 = 9
        _ = fi8
    }
    v9 := []int{}
    var v10 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 uint8 = 127
    var v2 int16 = 0
    var v3 bool = (342 >= 1)
    var v4 int16 = v2
    var v5 uint8 = v1
    fmt.Println(v2, v4)
    switch 0 {
    case 127:
        if true {
            v3 = v3
            v11 := []int{307, -1, -160}
            switch 57 {
            case 650:
                fmt.Println(v5, v3)
            case 605:
                var v12 int8 = -1
                _ = v12
            case -2147483648:
                if ((v1 % ((42) | 1)) != 42) {
                    switch 32767 {
                    case 127:
                        if (v1 >= 2) {
                            for fi13 := 0; fi13 < 16; fi13++ {
                                var v14 bool = v3
                                fmt.Println(fi13, v14, v2)
                                switch "a" {
                                case "":
                                    var v15 int16 = 32767
                                    _ = v15
                                case "a":
                                    if (44 < 128) {
                                        if (v2 != (v2 >> ((42) & 63))) {
                                            if (434 == 692) {
                                                if (false || v14) {
                                                    for fi16 := 0; fi16 < 4; fi16++ {
                                                        var v17 int16 = v2
                                                        _ = fi16
                                                        _ = v17
                                                    }
                                                } else {
                                                    var v18 int = 1
                                                    var v19 uint = 65535
                                                    if len(v11) >= 2 {
                                                        v11 = v11[1:len(v11)]
                                                    }
                                                    _ = v18
                                                    _ = v19
                                                }
                                                switch "foo bar" {
                                                case "a":
                                                    var v20 float64 = 0.1
                                                    _ = v20
                                                case "":
                                                    v2 += (-(v2))
                                                default:
                                                    fmt.Println(v1)
                                                }
                                            }
                                            switch int(v5) {
                                            case 128:
                                                fmt.Println(v4, v3)
                                            default:
                                                switch fi13 {
                                                case 1:
                                                    if ((65535 > 18446744073709551615) || v3) {
                                                        v11 = append(v11, 32768)
                                                    }
                                                }
                                            }
                                        } else {
                                            v3 = (-1 != 42)
                                            v21 := GenS2{F0: 42, F1: v1}
                                            _ = v21
                                        }
                                        v22 := GenS2{F0: 0, F1: v5}
                                        v14 = (815 == 403)
                                        _ = v22
                                    }
                                }
                                _ = fi13
                                _ = v14
                            }
                        }
                    case 32767:
                        v3 = (!v3)
                    case 865:
                        var v23 float32 = (float32(0.0) * (float32(3.14159) / 0.1))
                        _ = v23
                    default:
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    }
                    if ("\n\t" != "") {
                        if len(v11) > 0 {
                            v11[0] = -2147483648
                        }
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        v4 = 53
                    }
                    switch (((("hello" + (((("" + (("\n\t" + "\n\t") + "")) + "hello") + ("\n\t" + "a")) + "a")) + "hello") + "\n\t") + "") {
                    case "\n\t":
                        for fi24 := 0; fi24 < 12; fi24++ {
                            v25 := GenS2{F0: int8(fi24), F1: 127}
                            var v26 int8 = 51
                            if len(v11) > 0 {
                                v11[0] = 128
                            }
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                }
            }
            _ = v11
        } else {
            var v27 int8 = 41
            v28 := GenS2{F0: (v27 ^ 20), F1: 42}
            _ = v27
            _ = v28
        }
    case 128:
        fmt.Println(v1, v3)
    case 883:
        var v29 bool = ((42 == -34) || (1e10 != 0.1))
        _ = v29
    default:
        v30 := GenS3{F0: 479, F1: "foo bar"}
        _ = v30
    }
    var v31 int32 = 1
    v5 = v1
    var v32 int8 = -31
    fmt.Println(v4, v2)
    v2 -= v2
    v33 := []int{-1, -92}
    if len(v33) > 0 {
        v33[0] = -1
    }
    v34 := GenS4{F0: (v31 & (v31 >> ((42) & 63))), F1: v31, F2: 255}
    var v35 int = 42
    v32 = int8(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
