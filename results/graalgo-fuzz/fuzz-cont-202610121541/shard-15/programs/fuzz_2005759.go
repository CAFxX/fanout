package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int64 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int64; F1 string }
func genFn0() uint8 {
    var v6 int16 = 1000
    v7 := GenS1{F0: 16, F1: "\n\t", F2: 42}
    for fi8 := 0; fi8 < 1; fi8++ {
        v6 -= v6
        _ = fi8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        if (float32(0.0) < (-2.25 * -2.25)) {
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: 128, F1: "a", F2: 42}
                _ = fi10
                _ = v11
            }
            v6 += v6
            for fi12 := 0; fi12 < 14; fi12++ {
                if true {
                    if false {
                        switch "" {
                        case "foo bar":
                            var v13 uint32 = 42
                            _ = v13
                        default:
                            var v14 int = (fi12 ^ (^(fi9)))
                            _ = v14
                        }
                        if false {
                            if (!true) {
                                var v15 bool = (!false)
                                v16 := []int{}
                                if (((false && (46 > 127)) || true) && v15) {
                                    v17 := GenS1{F0: 0, F1: ("foo bar" + "hello"), F2: -9223372036854775808}
                                    var v18 string = "\n\t"
                                    _ = v17
                                    _ = v18
                                }
                                _ = v15
                                _ = v16
                            } else {
                                var v19 uint64 = 72
                                var v20 uint16 = 255
                                switch "\n\t" {
                                case "a":
                                    if false {
                                        v21 := GenS1{F0: 128, F1: ("" + "hello"), F2: 567}
                                        _ = v21
                                    }
                                default:
                                    v22 := []int{fi9, fi9}
                                    if ((^(v20)) <= v20) {
                                        v23 := GenS1{F0: 104, F1: ("" + "\n\t"), F2: 289}
                                        var v24 int64 = -1
                                        _ = v23
                                        _ = v24
                                    }
                                    var v25 uint32 = 1
                                    _ = v22
                                    _ = v25
                                }
                                _ = v19
                                _ = v20
                            }
                            if (!false) {
                                if true {
                                    if false {
                                        var v26 uint16 = 0
                                        _ = v26
                                    } else {
                                        v27 := []int{fi9, (fi9 ^ 823), fi9, (-(fi9))}
                                        _ = v27
                                    }
                                }
                                if true {
                                    v28 := GenS1{F0: 223, F1: ("hello" + ((("" + "foo bar") + "foo bar") + "\n\t")), F2: -249}
                                    var v29 float32 = -2.25
                                    _ = v28
                                    _ = v29
                                } else {
                                    fmt.Println(fi12)
                                }
                            } else {
                                var v30 uint64 = 601
                                _ = v30
                            }
                            if (1e10 < 1e10) {
                                if false {
                                    var v31 int16 = (-(v6))
                                    var v32 float64 = 1.5
                                    fmt.Println(v6, v31)
                                    _ = v31
                                    _ = v32
                                }
                            } else {
                                v6 -= (-(v6))
                            }
                            v6 += 1
                        }
                        if true {
                            v6 = v6
                            v6 = 42
                            v33 := []int{}
                            _ = v33
                        }
                    }
                    v34 := GenS1{F0: 0, F1: "\n\t", F2: 9223372036854775807}
                    var v35 bool = false
                    var v36 float64 = (((((0.0 * 1e10) + 0.0) + (1.5 / float64(0.1))) + 1.5) / (1e10 - 3.14159))
                    if (!((v6 / ((v6) | 1)) <= v6)) {
                        v37 := []int{(fi12 | (fi9 & (^(fi9))))}
                        v38 := GenS2{F0: 65535, F1: 220, F2: 9}
                        var v39 int = fi9
                        _ = v37
                        _ = v38
                        _ = v39
                    } else {
                        v36 = v36
                    }
                    _ = v34
                    _ = v35
                    _ = v36
                }
                v6 += (v6 ^ (v6 + (v6 >> ((4294967295) & 63))))
                _ = fi12
            }
        }
        switch "" {
        case "hello":
            var v40 uint8 = 1
            _ = v40
        }
        var v41 float64 = (-2.25 + 0.0)
        _ = fi9
        _ = v41
    }
    v42 := []int{9223372036854775807, 2147483647}
    _ = v42
    _ = v6
    _ = v7
    return 75
}

func main() {
    var v1 int32 = 495
    v2 := []uint32{42, 376}
    v3 := make(map[uint32]bool)
    var v4 int64 = 88
    var v5 int32 = -1
    var v43 uint = 634
    var v44 int8 = 42
    var v45 int = 1
    v46 := GenS1{F0: 173, F1: "hello", F2: v4}
    v47 := GenS3{F0: (v4 >> ((v43) & 63)), F1: ("foo bar" + "")}
    var v48 uint64 = 2
    if false {
        var v49 uint8 = 127
        _ = v49
    } else {
        fmt.Println(v44, v4)
        fmt.Println(v1, v4, v1)
    }
    v50 := GenS2{F0: 646, F1: 42, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v50
}
