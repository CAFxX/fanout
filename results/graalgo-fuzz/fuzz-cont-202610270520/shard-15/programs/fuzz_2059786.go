package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 int32 }
type GenS9 struct { F0 uint; F1 uint64; F2 uint }
func genFn0(p0 string) uint8 {
    var v6 int32 = 0
    var v7 uint8 = 1
    v7 -= v7
    v8 := []int{498, -9223372036854775808, 966, 2147483647}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (v7 ^ 127)
}
func genFn1(p0 int32) uint32 {
    if false {
        var v9 int16 = -32768
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                switch "hello" {
                case "":
                    v9 = (v9 - v9)
                case "hello":
                    if true {
                        var v12 uint16 = 790
                        fmt.Println(fi11, v12, fi11)
                        if (150 == 18446744073709551615) {
                            p0 += 1
                            v13 := GenS1{F0: 125}
                            v14 := GenS2{F0: v9, F1: false}
                            p0 -= (p0 >> ((591) & 63))
                            p0 = -2147483648
                            _ = v13
                            _ = v14
                        }
                        _ = v12
                    } else {
                        var v15 uint16 = 1
                        switch 9223372036854775807 {
                        case 32767:
                            if false {
                                v16 := GenS1{F0: 277}
                                v17 := []int{(fi10 ^ fi11), 714, -231}
                                switch int(v9) {
                                case -9223372036854775808:
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                case 127:
                                    if (1 <= 128) {
                                        var v18 uint32 = 0
                                        _ = v18
                                    }
                                case 32768:
                                    switch (("foo bar" + "") + "\n\t") {
                                    case "foo bar":
                                        v9 = v9
                                        var v19 int16 = v9
                                        if false {
                                            if (65535 >= 4294967295) {
                                                var v20 int64 = 911
                                                _ = v20
                                            }
                                            var v21 int = fi10
                                            if len(v17) > 0 {
                                                v17[0] = v21
                                            }
                                            _ = v21
                                        }
                                        _ = v19
                                    }
                                    v22 := GenS2{F0: 786, F1: false}
                                    var v23 uint16 = (v15 | (v15 - v15))
                                    var v24 float64 = 3.14159
                                    var v25 bool = (((1 >= 42) || true) || (false || true))
                                    fmt.Println(v25, fi11)
                                    v26 := GenS2{F0: 1000, F1: false}
                                    var v27 uint = 42
                                    v28 := GenS3{F0: 100, F1: int64(v23)}
                                    p0 = (p0 ^ (p0 ^ (p0 % (((p0 ^ -1)) | 1))))
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                var v29 float64 = (-2.25 / (0.0 - (3.14159 - -2.25)))
                                v30 := GenS4{F0: "\n\t", F1: v9}
                                _ = v16
                                _ = v17
                                _ = v29
                                _ = v30
                            } else {
                                p0 = (p0 << ((255) & 63))
                            }
                        case 2147483647:
                            v31 := []int{-1}
                            _ = v31
                        case 9223372036854775807:
                            if false {
                                switch "a" {
                                case "hello":
                                    if true {
                                        var v32 uint32 = 1
                                        v33 := []int{fi10, fi10, (fi10 / ((1) | 1))}
                                        _ = v32
                                        _ = v33
                                    }
                                case "foo bar":
                                    var v34 uint32 = 2
                                    _ = v34
                                case "a":
                                    var v35 int64 = 1
                                    _ = v35
                                default:
                                    v36 := GenS4{F0: ("\n\t" + "foo bar"), F1: -32768}
                                    _ = v36
                                }
                                v15 -= (^(v15))
                                v37 := GenS5{F0: 0}
                                _ = v37
                            } else {
                                fmt.Println(v15)
                                v15 = v15
                                v38 := GenS6{F0: 19, F1: 18, F2: 2147483648}
                                _ = v38
                            }
                        default:
                            v15 = 65535
                        }
                        _ = v15
                    }
                case "a":
                    var v39 uint64 = 2
                    _ = v39
                default:
                    fmt.Println(p0)
                }
                var v40 bool = false
                v41 := GenS7{F0: v40}
                _ = fi11
                _ = v40
                _ = v41
            }
            _ = fi10
        }
        v42 := []int{9223372036854775807, 859, 1}
        _ = v42
        _ = v9
    } else {
        var v43 int64 = 0
        _ = v43
    }
    var v44 float64 = -2.25
    v44 -= (1e10 / v44)
    _ = p0
    _ = v44
    return 42
}

func main() {
    var v1 int64 = -796
    v2 := make(map[string]uint32)
    v3 := []int{}
    var v4 int16 = 0
    v5 := []uint32{23, 4294967295, 2147483647, 42}
    v1 = v1
    v45 := GenS8{F0: 2147483648, F1: 685}
    if false {
        fmt.Println(v4, v4)
    }
    for fi46 := 0; fi46 < 5; fi46++ {
        var v47 uint32 = 831
        if len(v5) > 0 {
            v5[0] = v47
        }
        if true {
            fmt.Println(v47, fi46, v47)
            var v48 int64 = 42
            _ = v48
        }
        _ = fi46
        _ = v47
    }
    v1 = v1
    for fi49 := 0; fi49 < 17; fi49++ {
        delete(v2, "foo bar")
        _ = fi49
    }
    if false {
        if true {
            if (v4 == v4) {
                v4 += (v4 + (-(v4)))
                var v50 uint8 = genFn0("foo bar")
                _ = v50
            } else {
                var v51 bool = false
                if true {
                    fmt.Println(v51)
                    v1 -= v1
                    v4 = -32768
                } else {
                    if (2 != 26) {
                        fmt.Println(v51)
                        var v52 bool = true
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v52
                    }
                }
                _ = v51
            }
            v53 := GenS2{F0: v4, F1: (!false)}
            var v54 int = -242
            _ = v53
            _ = v54
        } else {
            delete(v2, "foo bar")
            v4 -= (-(v4))
            if false {
                var v55 float32 = 0.1
                _ = v55
            }
        }
    } else {
        v56 := GenS3{F0: 42, F1: (v1 << ((18446744073709551615) & 63))}
        _ = v56
    }
    var v57 uint64 = 2
    v58 := GenS9{F0: 18446744073709551615, F1: 633, F2: 18446744073709551615}
    if (false || (-1 == 77)) {
        var v59 uint32 = 1
        _ = v59
    } else {
        if (!false) {
            v1 += (-(v1))
            var v60 int32 = 0
            _ = v60
        } else {
            v61 := GenS7{F0: ("\n\t" > (("a" + (("hello" + "foo bar") + "")) + "\n\t"))}
            _ = v61
        }
        v62 := GenS1{F0: v57}
        _ = v62
    }
    var v63 float64 = 0.0
    v64 := GenS8{F0: 1, F1: 0}
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v57
    _ = v58
    _ = v63
    _ = v64
}
