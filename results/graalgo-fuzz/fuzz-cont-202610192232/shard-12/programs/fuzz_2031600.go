package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int64; F2 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint; F1 int32; F2 uint8 }
type GenS10 struct { F0 uint8; F1 uint16 }
type GenS11 struct { F0 int64; F1 uint16 }
type GenS12 struct { F0 int; F1 int16 }
func genFn0(p0 int16, p1 uint16) uint8 {
    switch "" {
    case "\n\t":
        var v4 uint64 = 18446744073709551615
        _ = v4
    case "foo bar":
        v5 := GenS1{F0: 65535}
        _ = v5
    default:
        var v6 int64 = -231
        _ = v6
    }
    if (((false || false) && ("\n\t" < "")) && false) {
        var v7 uint32 = 4294967295
        _ = v7
    } else {
        switch "\n\t" {
        case "\n\t":
            v8 := []int{790, -2147483648, 601}
            _ = v8
        }
        v9 := []int{}
        _ = v9
    }
    _ = p0
    _ = p1
    return 38
}
func genFn1(p0 int16) int16 {
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int8 = 127
        var v12 int = fi10
        v13 := GenS1{F0: 707}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    switch "" {
    case "\n\t":
        p0 = 302
    case "":
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 float64 = float64(3.14159)
            _ = fi14
            _ = v15
        }
    }
    var v16 bool = (true || true)
    _ = p0
    _ = v16
    return -204
}

func main() {
    v1 := []int{9223372036854775807}
    v2 := make(map[string]string)
    var v3 int16 = -907
    var v17 int32 = 0
    if true {
        for fi18 := 0; fi18 < 11; fi18++ {
            delete(v2, "\n\t")
            v3 += (v3 + 42)
            for fi19 := 0; fi19 < 9; fi19++ {
                fmt.Println(v3, v3)
                v20 := GenS1{F0: 1}
                v3 += v3
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    }
    v3 = genFn1((v3 - 42))
    if true {
        for fi21 := 0; fi21 < 11; fi21++ {
            for fi22 := 0; fi22 < 20; fi22++ {
                v3 = v3
                if (!(994 != 40)) {
                    if (true && true) {
                        v23 := GenS2{F0: fi21}
                        v24 := GenS3{F0: 100, F1: 12}
                        _ = v23
                        _ = v24
                    } else {
                        if (!(true || false)) {
                            fmt.Println(fi21, fi22)
                        } else {
                            switch (fi21 % ((fi21) | 1)) {
                            case -637:
                                if false {
                                    v25 := v2[""]
                                    v26 := GenS4{F0: 2, F1: -1}
                                    if false {
                                        v3 = -1
                                        if (!(v25 < v25)) {
                                            var v27 uint64 = 18446744073709551615
                                            v27 -= v27
                                            _ = v27
                                        } else {
                                            v25 = "\n\t"
                                            fmt.Println(v3, v17)
                                            if false {
                                                v28 := GenS5{F0: 417}
                                                _ = v28
                                            }
                                        }
                                        v25 = v25
                                    } else {
                                        v29 := GenS6{F0: 788}
                                        _ = v29
                                    }
                                    _ = v25
                                    _ = v26
                                } else {
                                    v3 = 32767
                                    if (249 == 128) {
                                        v3 += genFn1(v3)
                                        v17 = v17
                                    }
                                    v30 := GenS4{F0: 42, F1: 0}
                                    _ = v30
                                }
                            default:
                                v3 += genFn1((v3 >> ((1) & 63)))
                            }
                        }
                        v31 := GenS4{F0: 2, F1: -361}
                        _ = v31
                    }
                    if ((42 < 65535) && false) {
                        if (true || ((1.5 == ((1e10 + ((0.1 * float32((0.1 + float64((1e10 + -2.25))))) * 0.1)) + 1.5)) && (false && true))) {
                            var v32 int = fi22
                            _ = v32
                        } else {
                            v33 := GenS7{F0: 100, F1: 42, F2: 32767}
                            v34 := GenS8{F0: (fi22 | fi21)}
                            var v35 uint64 = 146
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    }
                    _ = len(v2)
                }
                _ = fi22
            }
            v3 += genFn1(v3)
            _ = fi21
        }
        v3 += (v3 << ((255) & 63))
    } else {
        if false {
            if ((genFn0(v3, 344) < 42) || true) {
                fmt.Println(v3)
            }
            v17 += v17
            fmt.Println(v3, v17)
        }
        if false {
            v36 := GenS4{F0: 2147483648, F1: 9223372036854775807}
            var v37 uint = 0
            if false {
                v37 = v37
                if false {
                    v17 += v17
                }
            } else {
                v38 := GenS9{F0: v37, F1: (v17 / ((-2147483648) | 1)), F2: 255}
                v1 = append(v1, -1)
                var v39 int16 = v3
                _ = v38
                _ = v39
            }
            _ = v36
            _ = v37
        }
    }
    if true {
        var v40 uint = 65535
        _ = v40
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        var v42 float32 = ((-2.25 / (float32((((0.0 + (0.1 / -2.25)) / (3.14159 * -2.25)) / 1.5)) / -2.25)) * 0.0)
        delete(v2, "a")
        v3 += 1000
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 20; fi43++ {
        for fi44 := 0; fi44 < 16; fi44++ {
            v17 = (^(v17))
            var v45 float64 = 1e10
            _ = fi44
            _ = v45
        }
        _ = fi43
    }
    if (0.1 >= 0.1) {
        if len(v1) > 0 {
            v1[0] = -560
        }
    }
    if true {
        for fi46 := 0; fi46 < 2; fi46++ {
            v47 := GenS10{F0: 255, F1: 32767}
            if ((((1.5 * (float32(((((-2.25 * 3.14159) - 3.14159) * 3.14159) + 0.0)) / 3.14159)) / 1.0) - (float32(0.1) * (((1.5 + 0.1) - 0.0) * 0.1))) <= (1.5 / float32((float64(-2.25) + (0.0 * 1.5))))) {
                for fi48 := 0; fi48 < 15; fi48++ {
                    v49 := GenS11{F0: 1, F1: 812}
                    v17 -= -1
                    _ = fi48
                    _ = v49
                }
                switch "a" {
                case "\n\t":
                    var v50 int = (fi46 % (((-(fi46))) | 1))
                    _ = v50
                case "":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            } else {
                switch "" {
                case "\n\t":
                    var v51 float64 = (((0.0 / 1.5) * (((((0.1 - 1e10) - 3.14159) / (-2.25 + ((3.14159 / (0.1 + -2.25)) * 1.5))) + (1e10 - 1e10)) - 1.5)) * -2.25)
                    _ = v51
                case "a":
                    switch (("\n\t" + "a") + (("foo bar" + "") + ("" + ("hello" + ("hello" + "hello"))))) {
                    case "a":
                        v1 = append(v1, (fi46 ^ fi46))
                    default:
                        var v52 int64 = -332
                        _ = v52
                    }
                }
            }
            _ = fi46
            _ = v47
        }
    } else {
        v53 := GenS12{F0: -128, F1: (v3 / ((v3) | 1))}
        if true {
            v54 := GenS12{F0: -125, F1: v3}
            fmt.Println(v3, v3)
            _ = v54
        }
        v3 -= 32767
        _ = v53
    }
    v2["\n\t"] = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
}
