package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS3 struct { F0 string; F1 string; F2 int16 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int64) int32 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 14; fi4++ {
        p0 += p0
        _ = fi4
    }
    var v5 uint16 = 65535
    var v6 uint32 = 2
    v7 := []int{}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 42
    var v3 uint8 = 177
    v2 += (v2 & v2)
    if (true && true) {
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS1{F0: 842, F1: 2, F2: 42}
            _ = fi8
            _ = v9
        }
        var v10 float64 = 0.1
        v10 = 3.14159
        _ = v10
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v2 = v2
        _ = fi11
    }
    v12 := GenS1{F0: 2, F1: 671, F2: 1}
    v13 := GenS2{F0: (v3 >> ((2) & 63)), F1: -80, F2: 0}
    if false {
        if (1 > 806) {
            if (float32(1.5) != 1.5) {
                v3 -= 42
                fmt.Println(v2)
                for fi14 := 0; fi14 < 9; fi14++ {
                    v15 := GenS1{F0: 671, F1: 65535, F2: -26}
                    v3 = v3
                    switch (fi14 + (fi14 + 9223372036854775807)) {
                    case -1:
                        if (fi14 != (fi14 & (-(fi14)))) {
                            for fi16 := 0; fi16 < 12; fi16++ {
                                var v17 float64 = 0.1
                                _ = fi16
                                _ = v17
                            }
                        }
                    case 2147483647:
                        if true {
                            if true {
                                v18 := []int{}
                                _ = v18
                            } else {
                                v2 = v2
                                var v19 uint16 = 32767
                                switch ("a" + "") {
                                case "a":
                                    v3 = (v3 + uint8(v19))
                                case "\n\t":
                                    var v20 int32 = 0
                                    _ = v20
                                case "hello":
                                    for fi21 := 0; fi21 < 4; fi21++ {
                                        v2 -= 198
                                        var v22 float64 = float64((0.1 + -2.25))
                                        v2 += v2
                                        if (577 >= uint(v3)) {
                                            var v23 int64 = 0
                                            v24 := GenS1{F0: 2147483647, F1: 0, F2: 0}
                                            if true {
                                                v19 -= uint16(v3)
                                                switch int(v3) {
                                                case 9223372036854775807:
                                                    var v25 uint = 2
                                                    _ = v25
                                                default:
                                                    v19 = (v19 | (^(v19)))
                                                }
                                                if (fi14 == (fi21 / ((fi21) | 1))) {
                                                    var v26 float32 = 3.14159
                                                    v3 = 232
                                                    _ = v26
                                                } else {
                                                    if true {
                                                        var v27 int32 = genFn0(v23)
                                                        v19 = 255
                                                        _ = v27
                                                    }
                                                    v2 += v2
                                                    var v28 int64 = (^(v23))
                                                    v23 = (v28 * (v23 + 0))
                                                    switch fi14 {
                                                    case 756:
                                                        var v29 int8 = 100
                                                        _ = v29
                                                    case 2147483647:
                                                        if true {
                                                            v30 := v1["\n\t"]
                                                            _ = v30
                                                        } else {
                                                            if false {
                                                                v31 := GenS2{F0: (v3 << ((18446744073709551615) & 63)), F1: 127, F2: 65535}
                                                                v3 = (^(v3))
                                                                v32 := GenS3{F0: "foo bar", F1: "a", F2: -667}
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                            v3 = (v3 * (v3 - v3))
                                                            if false {
                                                                var v33 uint32 = 2147483648
                                                                _ = v33
                                                            } else {
                                                                v34 := GenS4{F0: v3, F1: int(v19)}
                                                                v1[("\n\t" + "")] = "hello"
                                                                var v35 uint32 = 2147483648
                                                                v28 = v23
                                                                v3 = v3
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                        }
                                                    }
                                                    _ = v28
                                                }
                                            }
                                            _ = v23
                                            _ = v24
                                        } else {
                                            v36 := GenS5{F0: v19, F1: (("hello" + "") + "")}
                                            _ = v36
                                        }
                                        _ = fi21
                                        _ = v22
                                    }
                                default:
                                    delete(v1, ("\n\t" + "hello"))
                                }
                                _ = v19
                            }
                            fmt.Println(fi14, v3, v3)
                            var v37 uint = 581
                            _ = v37
                        } else {
                            v38 := GenS4{F0: v3, F1: fi14}
                            for fi39 := 0; fi39 < 10; fi39++ {
                                v40 := []int{}
                                v41 := GenS5{F0: 496, F1: ("\n\t" + "foo bar")}
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                            v3 = v3
                            _ = v38
                        }
                    }
                    _ = fi14
                    _ = v15
                }
            } else {
                v42 := []int{-2147483648, 0}
                for fi43 := 0; fi43 < 13; fi43++ {
                    v44 := GenS6{F0: 2, F1: (v3 & v3)}
                    _ = fi43
                    _ = v44
                }
                _ = v42
            }
        }
        v3 = (v3 ^ v3)
    } else {
        v3 = v3
        v45 := v1[""]
        for fi46 := 0; fi46 < 6; fi46++ {
            v47 := []int{fi46}
            _ = fi46
            _ = v47
        }
        _ = v45
    }
    for fi48 := 0; fi48 < 8; fi48++ {
        if false {
            var v49 int8 = 0
            _ = v49
        }
        v2 = v2
        _ = fi48
    }
    switch int(v3) {
    case 672:
        var v50 bool = false
        _ = v50
    case 311:
        var v51 float32 = -2.25
        _ = v51
    default:
        var v52 uint = 18446744073709551615
        _ = v52
    }
    for fi53 := 0; fi53 < 5; fi53++ {
        var v54 uint8 = 0
        v55 := GenS6{F0: 1, F1: v54}
        v56 := GenS5{F0: 32767, F1: "foo bar"}
        _ = fi53
        _ = v54
        _ = v55
        _ = v56
    }
    if false {
        fmt.Println(v3, v3)
        var v57 float64 = (1e10 * 1e10)
        fmt.Println(v2, v57, v3)
        _ = v57
    } else {
        var v58 uint64 = v2
        fmt.Println(v2, v2, v58)
        _ = v58
    }
    fmt.Println(v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
