package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 int32 }

func main() {
    var v1 uint16 = 31
    var v2 string = ("a" + "foo bar")
    var v3 int16 = 0
    v4 := []int{4294967295, 2147483647, 128, 100}
    v5 := []uint32{2, 164, 9}
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            if (false || (154 != 252)) {
                v8 := GenS1{F0: 127}
                v9 := GenS2{F0: 82, F1: v2}
                if true {
                    if (-757 >= -167) {
                        if (42 != 42) {
                            v10 := GenS2{F0: 42, F1: (v2 + (v2 + v2))}
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v10
                        } else {
                            var v11 float64 = -2.25
                            _ = v11
                        }
                    }
                    v2 = ("a" + "a")
                    if true {
                        if true {
                            v1 -= v1
                            fmt.Println(v3)
                            v2 = v2
                            v12 := GenS2{F0: 0, F1: (v2 + v2)}
                            _ = v12
                        }
                    } else {
                        v5 = append(v5, 1)
                        var v13 bool = true
                        var v14 int = (fi6 & (^(fi6)))
                        _ = v13
                        _ = v14
                    }
                }
                _ = v8
                _ = v9
            } else {
                switch fi6 {
                case 32767:
                    if false {
                        if ((true && (!true)) && false) {
                            switch 9223372036854775807 {
                            case -189:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case -800:
                                fmt.Println(fi7, fi6)
                            case -67:
                                var v15 int = 4294967295
                                if false {
                                    v2 = (v2 + v2)
                                } else {
                                    v1 = (^(v1))
                                    var v16 string = v2
                                    v2 = "\n\t"
                                    _ = v16
                                }
                                switch (fi7 + fi6) {
                                case 127:
                                    fmt.Println(v1, fi6, fi7)
                                case -645:
                                    var v17 string = v2
                                    _ = v17
                                }
                                _ = v15
                            }
                            v5 = append(v5, 859)
                            if false {
                                v18 := GenS3{F0: 0, F1: 0}
                                switch -9223372036854775808 {
                                case 4294967295:
                                    fmt.Println(fi7, v1)
                                default:
                                    var v19 bool = ((0.1 * 3.14159) < (1e10 / -2.25))
                                    _ = v19
                                }
                                v2 = (v2 + "hello")
                                _ = v18
                            } else {
                                switch "" {
                                case "a":
                                    v2 = "hello"
                                case "foo bar":
                                    v1 = v1
                                default:
                                    v5 = append(v5, 42)
                                }
                                v20 := GenS4{F0: (^(v1)), F1: v3}
                                var v21 int8 = int8(v3)
                                _ = v20
                                _ = v21
                            }
                        }
                    } else {
                        var v22 int32 = -2147483648
                        if true {
                            var v23 uint = 42
                            fmt.Println(v2, fi7)
                            fmt.Println(fi6, fi6, v2)
                            _ = v23
                        } else {
                            if true {
                                v4 = append(v4, fi6)
                            } else {
                                var v24 int64 = 1
                                v24 += v24
                                _ = v24
                            }
                        }
                        _ = v22
                    }
                case 128:
                    if false {
                        fmt.Println(v1)
                    }
                    var v25 int = (fi7 >> ((746) & 63))
                    _ = v25
                case 599:
                    if (965 < 4294967295) {
                        v2 = v2
                        var v26 uint64 = 591
                        var v27 uint32 = 2147483647
                        var v28 bool = (v3 <= (v3 & v3))
                        v1 -= v1
                        _ = v26
                        _ = v27
                        _ = v28
                    } else {
                        fmt.Println(v2, fi7)
                    }
                default:
                    if true {
                        if (v1 < (^(v1))) {
                            if false {
                                v2 = (("hello" + v2) + v2)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v2 = "\n\t"
                            }
                            v29 := GenS5{F0: -128}
                            _ = v29
                        }
                        if len(v4) > 0 {
                            v4[0] = fi6
                        }
                        var v30 uint32 = 4294967295
                        _ = v30
                    } else {
                        if (fi7 <= 9223372036854775807) {
                            var v31 float64 = (0.0 * ((1.5 * float64(float32(float64(0.0)))) + float64(0.1)))
                            _ = len(v4)
                            if false {
                                var v32 bool = false
                                var v33 float64 = v31
                                v31 -= 0.1
                                _ = v32
                                _ = v33
                            }
                            var v34 uint32 = 54
                            _ = v31
                            _ = v34
                        } else {
                            v35 := GenS4{F0: v1, F1: -1}
                            _ = v35
                        }
                        switch v2 {
                        case "foo bar":
                            fmt.Println(fi7, v2, v1)
                        case "hello":
                            v2 = "\n\t"
                        case "":
                            v4 = append(v4, (fi7 & (fi6 - (fi7 % (((fi7 + (fi6 & (^(fi7))))) | 1)))))
                        default:
                            v1 = (v1 - v1)
                        }
                        var v36 int16 = v3
                        _ = v36
                    }
                }
                fmt.Println(v3, v3)
            }
            v3 = 42
            _ = fi7
        }
        for fi37 := 0; fi37 < 6; fi37++ {
            v38 := GenS1{F0: -128}
            fmt.Println(v2, fi37)
            var v39 bool = true
            fmt.Println(v1, fi37, v2)
            v1 = v1
            var v40 float32 = 1e10
            var v41 uint16 = v1
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
            _ = v41
        }
        _ = fi6
    }
    fmt.Println(v3)
    var v42 int16 = v3
    fmt.Println(v1)
    v43 := GenS6{F0: 827, F1: -764}
    v1 = 32767
    fmt.Println(v2, v3)
    var v44 int64 = int64(v42)
    switch v2 {
    case "foo bar":
        v44 = (v44 % ((v44) | 1))
    default:
        fmt.Println(v42, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
