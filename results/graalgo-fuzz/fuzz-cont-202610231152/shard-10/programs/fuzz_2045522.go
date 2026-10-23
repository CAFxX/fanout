package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint16) uint64 {
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int = fi5
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{(fi7 / ((v6) | 1))}
            fmt.Println(fi7, v6)
            v9 := GenS1{F0: 0}
            _ = fi7
            _ = v8
            _ = v9
        }
        if false {
            var v10 uint8 = 2
            v11 := []int{32768, (^(v6)), 9223372036854775807, 4294967295}
            _ = v10
            _ = v11
        } else {
            if (!true) {
                p1 += (p1 & (p1 << ((255) & 63)))
            }
        }
        _ = fi5
        _ = v6
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := []int{}
        _ = fi12
        _ = v13
    }
    var v14 float64 = float64(0.0)
    v14 = v14
    _ = p0
    _ = p1
    _ = v14
    return 1
}
func genFn1(p0 uint32) int32 {
    p0 -= p0
    p0 = p0
    var v15 bool = true
    _ = p0
    _ = v15
    return -2147483648
}

func main() {
    var v1 uint8 = 57
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    var v4 int = 127
    var v16 int = -1
    var v17 uint8 = 42
    for fi18 := 0; fi18 < 16; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            fmt.Println(fi18, fi19, v4)
            if false {
                var v20 uint = 255
                var v21 int32 = -2147483648
                _ = v20
                _ = v21
            }
            v22 := []int{-9223372036854775808, fi18, (v4 >> ((2) & 63)), (v4 / ((128) | 1))}
            var v23 bool = false
            _ = fi19
            _ = v22
            _ = v23
        }
        _ = fi18
    }
    v24 := v2[2]
    v25 := []int{510, (v16 & (v16 % ((380) | 1))), 2147483647, (v4 - (^(v4)))}
    switch (((("a" + ("foo bar" + ("a" + "a"))) + "hello") + "foo bar") + "hello") {
    case "":
        fmt.Println(v16, v1)
    case "\n\t":
        if ((299 == 995) || false) {
            if (1 != -895) {
                var v26 uint32 = 0
                if (-1 > 1) {
                    fmt.Println(v24)
                    _ = len(v25)
                } else {
                    fmt.Println(v17, v16)
                }
                _ = v26
            } else {
                fmt.Println(v24)
                v16 -= v4
                fmt.Println(v16, v1)
            }
            fmt.Println(v1, v17, v16)
        } else {
            fmt.Println(v24)
            for fi27 := 0; fi27 < 4; fi27++ {
                v28 := GenS1{F0: v1}
                if true {
                    if ((^(v24)) >= v16) {
                        var v29 int32 = genFn1(4294967295)
                        v30 := GenS2{F0: (873 < 389)}
                        _ = v29
                        _ = v30
                    } else {
                        if ((!(0 == 127)) || false) {
                            switch "a" {
                            case "\n\t":
                                var v31 uint64 = 18446744073709551615
                                _ = v31
                            case "foo bar":
                                fmt.Println(v16)
                            default:
                                if (false && false) {
                                    fmt.Println(v1)
                                    var v32 uint = 65535
                                    _ = v32
                                }
                            }
                            v24 += (v24 << ((1) & 63))
                        } else {
                            v24 += 32767
                            if true {
                                var v33 float32 = 3.14159
                                var v34 bool = (!false)
                                _ = v33
                                _ = v34
                            }
                        }
                        var v35 float64 = (-2.25 / -2.25)
                        _ = v35
                    }
                    var v36 uint8 = (v17 ^ (v17 + v17))
                    var v37 bool = true
                    _ = v36
                    _ = v37
                } else {
                    v16 -= (v4 / ((32767) | 1))
                    var v38 int64 = 130
                    _ = v38
                }
                v17 = v1
                _ = fi27
                _ = v28
            }
            var v39 int32 = 2147483647
            _ = v39
        }
    case "hello":
        v3["hello"] = (272 <= -2147483648)
    default:
        for fi40 := 0; fi40 < 8; fi40++ {
            for fi41 := 0; fi41 < 18; fi41++ {
                v42 := v2[2147483648]
                switch 127 {
                case 23:
                    if (!false) {
                        var v43 int64 = 42
                        if false {
                            if true {
                                fmt.Println(fi41, v4, v1)
                                if (255 < 736) {
                                    v17 += 0
                                    var v44 uint32 = 806
                                    fmt.Println(fi40, fi41)
                                    _ = v44
                                }
                                var v45 uint32 = 201
                                if false {
                                    v1 = (v17 / ((v1) | 1))
                                    var v46 uint64 = 0
                                    var v47 int16 = 1000
                                    _ = v46
                                    _ = v47
                                }
                                _ = v45
                            } else {
                                if (-2.25 <= 1.5) {
                                    var v48 uint32 = 4294967295
                                    _ = v48
                                } else {
                                    v49 := GenS2{F0: false}
                                    _ = v49
                                }
                                v50 := v3[(("hello" + "") + ((("a" + "foo bar") + "foo bar") + "a"))]
                                var v51 uint16 = 365
                                _ = v50
                                _ = v51
                            }
                            v52 := GenS3{F0: 42}
                            switch "hello" {
                            case "hello":
                                if len(v25) > 0 {
                                    v25[0] = fi41
                                }
                            case "a":
                                var v53 float64 = (((-2.25 + 3.14159) * 0.0) + 1.5)
                                _ = v53
                            }
                            _ = v52
                        } else {
                            v54 := GenS4{F0: 1, F1: 1000, F2: (v1 | 1)}
                            var v55 float64 = 0.0
                            var v56 uint64 = 4294967295
                            v1 = 0
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        if (!(!true)) {
                            switch "" {
                            case "foo bar":
                                fmt.Println(v24)
                            case "":
                                if ((genFn0(543, 42) >= 851) || true) {
                                    v4 = (v24 * 4294967295)
                                }
                            case "hello":
                                fmt.Println(v17, v17, v16)
                                v57 := GenS5{F0: 1}
                                if false {
                                    var v58 uint = 736
                                    var v59 float32 = 0.1
                                    var v60 bool = ((int32(v58) == -2147483648) && (!(!false)))
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                _ = v57
                            }
                            v16 += v4
                        } else {
                            var v61 float32 = 3.14159
                            _ = v61
                        }
                        var v62 uint = 187
                        _ = v43
                        _ = v62
                    } else {
                        var v63 uint = 937
                        var v64 float32 = (1e10 + float32((-2.25 * float64(float32((((0.1 - 1.5) + float64(3.14159)) * (1.5 + 0.0)))))))
                        _ = v63
                        _ = v64
                    }
                    var v65 uint32 = 2147483648
                    _ = v65
                default:
                    var v66 uint = 42
                    _ = v66
                }
                _ = fi41
                _ = v42
            }
            delete(v2, 2147483648)
            _ = fi40
        }
    }
    v17 += v17
    delete(v2, 0)
    fmt.Println(v17, v17)
    var v67 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v67
}
