package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint64; F1 string; F2 int8 }
type GenS4 struct { F0 string }
func genFn0() uint32 {
    v5 := GenS1{F0: 220, F1: 156, F2: -981}
    if true {
        if (32767 > 436) {
            switch ("foo bar" + "a") {
            case "\n\t":
                var v6 int32 = -1
                _ = v6
            case "a":
                var v7 float32 = 1.5
                _ = v7
            }
        }
    }
    v8 := []int{-1, 9223372036854775807, 127}
    var v9 int16 = 654
    _ = v5
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []string{("hello" + "foo bar"), "hello", "foo bar"}
    var v2 uint16 = 919
    v3 := []int{-577, -452, 9223372036854775807, 42}
    var v4 uint64 = 579
    v2 = v2
    switch 306 {
    case 127:
        v2 += uint16(v4)
    }
    if len(v3) > 0 {
        v3[0] = 127
    }
    var v10 float32 = 0.1
    for fi11 := 0; fi11 < 17; fi11++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            switch (fi11 % ((fi11) | 1)) {
            case -953:
                var v12 int16 = 0
                _ = v12
            }
        } else {
            v1 = append(v1, (("foo bar" + ("hello" + (("a" + "foo bar") + ("\n\t" + ("hello" + "\n\t"))))) + "a"))
        }
        _ = fi11
    }
    if false {
        if true {
            if (127 > -1) {
                var v13 float64 = 1.5
                var v14 uint64 = v4
                _ = v13
                _ = v14
            }
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            var v15 int32 = -252
            _ = v15
        } else {
            if len(v3) > 0 {
                v3[0] = 257
            }
            v4 -= v4
        }
    } else {
        if (25 < 1) {
            var v16 int16 = -1
            v17 := GenS1{F0: (v4 - v4), F1: 33, F2: -2147483648}
            _ = v16
            _ = v17
        }
        v10 = v10
        v4 -= v4
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "" {
    case "":
        v18 := GenS2{F0: 0, F1: (v2 << ((42) & 63))}
        _ = v18
    case "hello":
        for fi19 := 0; fi19 < 19; fi19++ {
            v3 = append(v3, (fi19 & (fi19 >> ((2) & 63))))
            v20 := GenS2{F0: 4294967295, F1: 315}
            v10 -= (float32(float64(0.1)) / v10)
            _ = fi19
            _ = v20
        }
    case "foo bar":
        fmt.Println(v2, v2, v10)
    default:
        fmt.Println(v2)
    }
    if len(v3) > 0 {
        v3[0] = -126
    }
    fmt.Println(v2)
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 uint32 = 2
            v23 -= 763
            if (571 <= 1) {
                v24 := GenS2{F0: 254, F1: v2}
                var v25 uint16 = (v2 / ((541) | 1))
                _ = v24
                _ = v25
            } else {
                v26 := GenS1{F0: (^(v4)), F1: uint(v2), F2: -283}
                if len(v3) > 0 {
                    v3[0] = 0
                }
                _ = v26
            }
            if true {
                v1 = append(v1, "\n\t")
            } else {
                v23 += 42
                if (!(!false)) {
                    v23 = genFn0()
                } else {
                    var v27 uint16 = v2
                    switch fi21 {
                    case 2147483647:
                        v3 = append(v3, fi22)
                    case -1:
                        switch ("hello" + ((("" + "\n\t") + "\n\t") + "a")) {
                        case "hello":
                            var v28 uint64 = 255
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v28
                        case "":
                            var v29 float32 = (0.0 / v10)
                            _ = v29
                        case "foo bar":
                            switch "a" {
                            case "foo bar":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "a":
                                var v30 int32 = -1
                                _ = v30
                            case "\n\t":
                                var v31 uint32 = genFn0()
                                _ = v31
                            default:
                                var v32 float32 = v10
                                _ = v32
                            }
                        default:
                            v2 -= 65535
                        }
                    default:
                        switch ("foo bar" + ("" + "foo bar")) {
                        case "foo bar":
                            v33 := GenS2{F0: 18446744073709551615, F1: v2}
                            _ = v33
                        case "hello":
                            v34 := GenS1{F0: v4, F1: 42, F2: -515}
                            _ = v34
                        case "":
                            fmt.Println(v2, fi21)
                            var v35 int32 = -625
                            v36 := GenS3{F0: (v4 % ((42) | 1)), F1: "hello", F2: -128}
                            _ = v35
                            _ = v36
                        }
                    }
                    fmt.Println(v27, v23, v2)
                    if (-2147483648 >= -1) {
                        fmt.Println(v23, v10)
                        if ((-(fi22)) < (fi21 - (-(fi22)))) {
                            v2 = v27
                            var v37 float32 = float32(0.1)
                            _ = v37
                        } else {
                            v4 += 42
                            v3 = append(v3, 150)
                        }
                    }
                    _ = v27
                }
                _ = len(v1)
            }
            v4 -= v4
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    fmt.Println(v2, v4, v2)
    v38 := GenS4{F0: (("foo bar" + "\n\t") + ("a" + "a"))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
