package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }
type GenS4 struct { F0 string; F1 string; F2 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 int32; F2 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 uint; F2 bool }
type GenS9 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int, p2 uint) int32 {
    var v7 string = "hello"
    var v8 float64 = 0.1
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 int32) uint64 {
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint = 712
        v10 += 25
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 9223372036854775807}
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 int64 = 42
        v14 := GenS2{F0: 0, F1: 127}
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v11
    return 1
}

func main() {
    var v1 float32 = (1.5 * (float32((1.5 + 3.14159)) + ((float32(3.14159) * 1.5) - (float32((1.5 - 1e10)) + (3.14159 - 0.1)))))
    var v2 int64 = 1
    var v3 bool = false
    var v4 int64 = v2
    var v5 string = "\n\t"
    var v6 float32 = (float32((-2.25 + float64((v1 / (v1 / (((0.1 * v1) / (-2.25 / (0.0 - v1))) - v1)))))) * (v1 / -2.25))
    v15 := GenS3{F0: v3, F1: v3, F2: 86}
    fmt.Println(v4, v4, v4)
    v16 := []int{127, 32768, -2147483648}
    v2 -= v2
    _ = len(v16)
    v17 := GenS1{F0: 800}
    fmt.Println(v4, v6)
    var v18 uint8 = 42
    for fi19 := 0; fi19 < 14; fi19++ {
        switch "hello" {
        case "hello":
            if (18446744073709551615 < 884) {
                v18 = v18
                v6 -= 0.1
            } else {
                if false {
                    var v20 int = fi19
                    switch fi19 {
                    case -2147483648:
                        v2 = (v2 % (((v2 + -9223372036854775808)) | 1))
                    default:
                        fmt.Println(v6)
                    }
                    for fi21 := 0; fi21 < 15; fi21++ {
                        v22 := GenS3{F0: (65535 == 707), F1: (0 != (v18 ^ (^(v18)))), F2: (v18 / ((v18) | 1))}
                        _ = len(v16)
                        _ = fi21
                        _ = v22
                    }
                    _ = v20
                } else {
                    var v23 uint8 = (^(v18))
                    switch fi19 {
                    case 32768:
                        switch ((v5 + ((v5 + v5) + v5)) + "foo bar") {
                        case "hello":
                            var v24 float64 = 1e10
                            _ = v24
                        }
                    case 42:
                        if (fi19 > fi19) {
                            if v3 {
                                fmt.Println(v5, v4, v2)
                                switch fi19 {
                                case 32768:
                                    switch 1 {
                                    case -1:
                                        switch (fi19 / (((fi19 % ((-985) | 1))) | 1)) {
                                        case 128:
                                            v16 = append(v16, fi19)
                                        }
                                    case -885:
                                        v18 = (v23 | 32)
                                    }
                                case 405:
                                    fmt.Println(v2)
                                default:
                                    var v25 int16 = -32768
                                    _ = v25
                                }
                            }
                            v2 = (-(v2))
                        } else {
                            v26 := GenS4{F0: (v5 + v5), F1: (v5 + (((((v5 + "\n\t") + v5) + v5) + v5) + v5)), F2: 42}
                            for fi27 := 0; fi27 < 1; fi27++ {
                                var v28 float64 = 1.5
                                v3 = v3
                                v29 := GenS2{F0: 2, F1: (^(v18))}
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            _ = v26
                        }
                    }
                    v4 += (-(v2))
                    _ = v23
                }
            }
        }
        if v3 {
            var v30 int32 = genFn0(uint64(v18), fi19, 18446744073709551615)
            if v3 {
                v2 += v4
            } else {
                fmt.Println(v6, v3, v4)
            }
            _ = v30
        } else {
            if (984 < 1) {
                switch v5 {
                case "foo bar":
                    var v31 float64 = (((float64((v1 / v1)) - (1.5 - float64(0.1))) / 3.14159) + 3.14159)
                    _ = v31
                }
                v32 := GenS5{F0: 42}
                fmt.Println(fi19, fi19)
                _ = v32
            }
        }
        if false {
            if len(v16) > 0 {
                v16[0] = (fi19 | fi19)
            }
        } else {
            for fi33 := 0; fi33 < 9; fi33++ {
                if ((2 < 806) || (-1 != (v4 / ((v4) | 1)))) {
                    v6 -= 0.0
                    var v34 int = fi33
                    switch (v5 + "a") {
                    case "a":
                        v4 += 9223372036854775807
                    case "":
                        v35 := GenS2{F0: 2, F1: (v18 - (v18 % ((v18) | 1)))}
                        _ = v35
                    case "foo bar":
                        switch v5 {
                        case "foo bar":
                            v36 := GenS6{F0: 0, F1: -595, F2: -128}
                            _ = v36
                        case "hello":
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        case "a":
                            switch "hello" {
                            case "\n\t":
                                if len(v16) > 0 {
                                    v16[0] = v34
                                }
                            default:
                                v37 := GenS5{F0: -1}
                                _ = v37
                            }
                        }
                    default:
                        v38 := GenS3{F0: (-545 >= -1), F1: (0.1 > -2.25), F2: v18}
                        _ = v38
                    }
                    _ = v34
                }
                if ((1e10 / (0.1 / float64(float32(((float64((1e10 - float32(-2.25))) + (float64(v6) + 1e10)) - 0.0))))) == 0.0) {
                    if true {
                        if false {
                            if (0 == genFn0(121, fi19, uint(v18))) {
                                if true {
                                    var v39 float64 = (1.5 + (0.1 * (0.0 - 1.5)))
                                    v40 := GenS7{F0: fi19}
                                    _ = v39
                                    _ = v40
                                }
                            }
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            if (uint(v18) > 18446744073709551615) {
                                v4 -= (^(v4))
                                v41 := GenS8{F0: v18, F1: 4294967295, F2: (!v3)}
                                _ = v41
                            }
                        }
                        if v3 {
                            fmt.Println(v3, v5)
                            v16 = append(v16, (fi19 & (fi33 / ((485) | 1))))
                        }
                    } else {
                        if (1 > 65535) {
                            v3 = true
                            fmt.Println(v4)
                            v2 = 9223372036854775807
                        } else {
                            v5 = v5
                            v42 := GenS9{F0: (v18 + 127)}
                            _ = v42
                        }
                    }
                    v1 = v6
                } else {
                    v4 -= v2
                    v16 = append(v16, (^(fi33)))
                    if ((genFn0(4294967295, fi33, 0) == 781) && (42 != 42)) {
                        var v43 uint8 = 1
                        _ = v43
                    }
                }
                var v44 uint = 4294967295
                _ = fi33
                _ = v44
            }
            v1 -= (0.0 - v1)
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
