package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS5 struct { F0 int16; F1 int32; F2 uint16 }
func genFn0(p0 string) uint16 {
    v7 := []int{-755}
    var v8 float32 = 3.14159
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"hello"}
    v3 := []int{}
    v4 := []string{"hello", "a"}
    var v5 uint = 338
    v6 := make(map[string]bool)
    for fi9 := 0; fi9 < 20; fi9++ {
        v1[12] = 4294967295
        if false {
            var v10 int16 = 32767
            v5 -= v5
            _ = v10
        }
        _ = fi9
    }
    if false {
        v5 -= (^(v5))
        switch "hello" {
        case "\n\t":
            v11 := GenS1{F0: 4294967295, F1: 42}
            _ = v11
        case "a":
            v5 -= v5
        }
    }
    switch ("hello" + (("hello" + "a") + "a")) {
    case "\n\t":
        if true {
            switch -256 {
            case 9223372036854775807:
                var v12 int = -652
                _ = v12
            case -226:
                v13 := GenS1{F0: 0, F1: 86}
                _ = v13
            }
            v14 := GenS1{F0: 1, F1: 190}
            var v15 float64 = (float64(float32((-2.25 + 1e10))) - 0.0)
            _ = v14
            _ = v15
        }
    }
    var v16 uint16 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v17 uint64 = 2
    switch "\n\t" {
    case "a":
        var v18 string = "hello"
        _ = v18
    case "hello":
        v17 = (v17 * (v17 / ((4294967295) | 1)))
    default:
        v1[768] = 1
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := v6[(("foo bar" + "foo bar") + "a")]
        _ = fi19
        _ = v20
    }
    if true {
        v21 := GenS2{F0: v5, F1: false}
        _ = v21
    } else {
        if len(v2) > 0 {
            v2[0] = ""
        }
    }
    fmt.Println(v5, v5, v5)
    fmt.Println(v17)
    for fi22 := 0; fi22 < 18; fi22++ {
        if ((true && false) || true) {
            var v23 float32 = ((-2.25 - (3.14159 / 1e10)) - (0.0 / -2.25))
            if false {
                if (347 < v17) {
                    fmt.Println(v16, v5, v17)
                } else {
                    var v24 float64 = -2.25
                    _ = v24
                }
            } else {
                v25 := GenS1{F0: 2, F1: (v17 ^ (^(v17)))}
                if (0 < -11) {
                    switch 306 {
                    case 4294967295:
                        if (!((false || (!true)) && true)) {
                            delete(v6, (("a" + ("\n\t" + "foo bar")) + "hello"))
                            fmt.Println(v17, fi22, v17)
                            if (false && ((1e10 * 1.5) > 1.5)) {
                                var v26 bool = true
                                var v27 uint64 = 0
                                _ = v26
                                _ = v27
                            } else {
                                var v28 uint32 = 998
                                _ = v28
                            }
                        } else {
                            for fi29 := 0; fi29 < 18; fi29++ {
                                switch "hello" {
                                case "hello":
                                    v30 := GenS1{F0: 1, F1: (v17 + 808)}
                                    v17 += v17
                                    if (2147483647 >= 2147483647) {
                                        var v31 float64 = 3.14159
                                        v32 := v6[("foo bar" + "foo bar")]
                                        fmt.Println(v17, v32)
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v33 uint32 = 2147483648
                                        v34 := GenS3{F0: v17}
                                        v17 -= uint64(v5)
                                        _ = v33
                                        _ = v34
                                    }
                                    v5 += 944
                                    _ = v30
                                case "a":
                                    var v35 float32 = float32(((1e10 + (1e10 + 0.1)) * -2.25))
                                    _ = v35
                                case "foo bar":
                                    var v36 uint32 = 4294967295
                                    _ = v36
                                default:
                                    switch "" {
                                    case "":
                                        v17 = (v17 * (v17 << ((v5) & 63)))
                                    }
                                }
                                _ = fi29
                            }
                            var v37 int16 = int16(v17)
                            _ = v37
                        }
                    case -298:
                        v38 := GenS4{F0: v17, F1: -9223372036854775808, F2: 0}
                        _ = v38
                    default:
                        var v39 uint16 = (v16 << (((v5 % (((v5 + (v5 << ((v5) & 63)))) | 1))) & 63))
                        _ = v39
                    }
                    _ = len(v6)
                } else {
                    v17 += 528
                    var v40 int32 = -721
                    _ = v40
                }
                var v41 uint32 = 42
                _ = v25
                _ = v41
            }
            _ = v23
        } else {
            for fi42 := 0; fi42 < 14; fi42++ {
                var v43 int = fi22
                _ = fi42
                _ = v43
            }
            var v44 uint64 = v17
            _ = v44
        }
        _ = fi22
    }
    v16 = genFn0("hello")
    v45 := GenS5{F0: 1, F1: -1, F2: v16}
    v16 += v16
    fmt.Println(v17, v17, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
