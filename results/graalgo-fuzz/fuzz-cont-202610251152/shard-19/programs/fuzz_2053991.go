package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 bool }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 string; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int64) uint16 {
    p0 -= (p0 >> ((65535) & 63))
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := GenS1{F0: "a", F1: 65535, F2: false}
        if (208 >= 128) {
            var v6 uint = 666
            v7 := []int{0, (fi4 ^ (fi4 << ((v6) & 63)))}
            var v8 int64 = 42
            _ = v6
            _ = v7
            _ = v8
        }
        var v9 float32 = 3.14159
        _ = fi4
        _ = v5
        _ = v9
    }
    _ = p0
    return 0
}
func genFn1(p0 uint32) string {
    p0 = p0
    var v10 int16 = 1000
    _ = p0
    _ = v10
    return "a"
}

func main() {
    var v1 int = -1
    v2 := make(map[string]int)
    var v3 int32 = -762
    v3 = v3
    v2["hello"] = (v1 >> ((18446744073709551615) & 63))
    var v11 int64 = -1
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            v3 -= v3
            fmt.Println(v3)
            _ = fi13
        }
        switch "hello" {
        case "foo bar":
            if true {
                if true {
                    var v14 int32 = v3
                    for fi15 := 0; fi15 < 15; fi15++ {
                        v16 := GenS1{F0: "a", F1: 565, F2: false}
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                } else {
                    v17 := []int{9223372036854775807, v1, int(v3)}
                    v18 := GenS1{F0: "a", F1: 42, F2: (402 != 42)}
                    _ = v17
                    _ = v18
                }
                v19 := GenS1{F0: "hello", F1: 2147483647, F2: true}
                _ = v19
            }
        default:
            v11 -= 1
        }
        _ = fi12
    }
    v20 := v2[""]
    if false {
        var v21 uint64 = 333
        if true {
            var v22 uint8 = 4
            var v23 uint32 = 136
            var v24 int = v20
            _ = v22
            _ = v23
            _ = v24
        } else {
            if true {
                v21 -= v21
            } else {
                v21 += (v21 ^ 4294967295)
            }
        }
        if (300 < 18446744073709551615) {
            v25 := GenS1{F0: "hello", F1: 4294967295, F2: true}
            for fi26 := 0; fi26 < 4; fi26++ {
                v3 += (v3 + v3)
                v27 := GenS1{F0: "hello", F1: 687, F2: false}
                _ = fi26
                _ = v27
            }
            _ = v25
        } else {
            v28 := []int{(v20 % ((1) | 1)), (v20 | (^(v1))), v1}
            v21 -= (v21 & 821)
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 int8 = 42
                if true {
                    delete(v2, "foo bar")
                    v11 = -1
                    var v31 uint64 = (^(v21))
                    _ = v31
                }
                fmt.Println(v1, fi29, v21)
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        _ = v21
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        var v33 uint16 = genFn0((v11 / (((v11 ^ (v11 - (v11 + -873)))) | 1)))
        for fi34 := 0; fi34 < 6; fi34++ {
            v35 := []int{}
            if false {
                var v36 uint16 = (v33 - v33)
                _ = v36
            }
            if true {
                var v37 uint16 = (v33 >> ((386) & 63))
                var v38 float64 = 0.0
                _ = v37
                _ = v38
            }
            if false {
                delete(v2, "foo bar")
            } else {
                if true {
                    if false {
                        var v39 float64 = 3.14159
                        if false {
                            v2["foo bar"] = fi32
                        } else {
                            fmt.Println(v20, v20, v3)
                            v40 := GenS1{F0: "hello", F1: 65535, F2: false}
                            if true {
                                if false {
                                    switch fi34 {
                                    case 734:
                                        v35 = append(v35, 452)
                                    case -104:
                                        v33 = 42
                                    }
                                }
                                v41 := GenS1{F0: "a", F1: 42, F2: true}
                                var v42 uint64 = 1
                                v43 := GenS1{F0: "a", F1: 42, F2: (!(4294967295 > 65535))}
                                v35 = append(v35, (v20 | (v1 / ((fi34) | 1))))
                                if true {
                                    fmt.Println(v11, v1, v1)
                                    var v44 float64 = 1e10
                                    _ = v44
                                }
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            _ = v40
                        }
                        if (0 <= 128) {
                            v45 := GenS2{F0: true, F1: v11}
                            if (25 > -82) {
                                v46 := GenS2{F0: false, F1: v11}
                                var v47 uint32 = 739
                                _ = v46
                                _ = v47
                            }
                            _ = v45
                        }
                        _ = v39
                    } else {
                        fmt.Println(v20, v20)
                        var v48 string = "foo bar"
                        v49 := GenS2{F0: false, F1: (v11 + (^(v11)))}
                        if (v48 >= v48) {
                            v50 := GenS3{F0: 867, F1: 1}
                            fmt.Println(v3, v48, fi34)
                            _ = v50
                        } else {
                            v51 := GenS2{F0: true, F1: v11}
                            v11 += v11
                            v3 -= -758
                            v33 = (^(v33))
                            if false {
                                var v52 uint = 18446744073709551615
                                var v53 uint8 = 31
                                v1 = int(v3)
                                _ = v52
                                _ = v53
                            } else {
                                v20 = v20
                                v20 = 32767
                            }
                            v20 = fi32
                            _ = v51
                        }
                        _ = v48
                        _ = v49
                    }
                    var v54 string = "foo bar"
                    _ = v54
                }
                v1 = v1
            }
            _ = len(v35)
            _ = fi34
            _ = v35
        }
        v11 = 0
        _ = fi32
        _ = v33
    }
    switch "foo bar" {
    case "a":
        v3 += v3
    }
    if true {
        switch "\n\t" {
        case "":
            for fi55 := 0; fi55 < 13; fi55++ {
                switch "a" {
                case "":
                    v56 := GenS4{F0: "foo bar", F1: 0, F2: v3}
                    _ = v56
                }
                _ = fi55
            }
        case "a":
            v11 += v11
        }
        var v57 uint64 = 426
        v58 := GenS5{F0: 4294967295}
        _ = v57
        _ = v58
    } else {
        if true {
            if false {
                var v59 uint16 = 382
                var v60 uint32 = 43
                if ((v60 / ((v60) | 1)) < v60) {
                    v60 += (v60 / (((^(v60))) | 1))
                } else {
                    v61 := []int{v20, 1}
                    fmt.Println(v1, v1)
                    _ = v61
                }
                _ = v59
                _ = v60
            } else {
                v3 += v3
            }
            v2["\n\t"] = (v20 << ((260) & 63))
        } else {
            v3 = -643
            var v62 float64 = 0.0
            v63 := GenS4{F0: "\n\t", F1: 4294967295, F2: v3}
            _ = v62
            _ = v63
        }
        var v64 int8 = -1
        _ = v64
    }
    switch v20 {
    case 128:
        fmt.Println(v1, v3)
    case 32767:
        var v65 int = (-(v1))
        _ = v65
    case -2147483648:
        v3 = v3
    }
    v66 := GenS4{F0: "a", F1: 1, F2: (-(v3))}
    var v67 uint = 255
    v68 := v2[""]
    var v69 int64 = v11
    var v70 int8 = 0
    v71 := []int{v20, (v1 * (v1 >> (((^(v67))) & 63)))}
    v11 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
    _ = v71
}
