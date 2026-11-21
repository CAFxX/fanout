package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint64; F1 int16; F2 bool }
type GenS7 struct { F0 uint16; F1 bool }
func genFn0(p0 uint8) int {
    if true {
        p0 = 0
    } else {
        v7 := GenS1{F0: (p0 + p0), F1: 0, F2: (p0 & (p0 % ((p0) | 1)))}
        _ = v7
    }
    v8 := GenS1{F0: p0, F1: -128, F2: p0}
    v9 := GenS2{F0: -9223372036854775808}
    for fi10 := 0; fi10 < 18; fi10++ {
        p0 = (p0 ^ p0)
        v11 := []int{(fi10 ^ fi10), (fi10 | fi10), (fi10 >> ((472) & 63)), (fi10 - (-(fi10)))}
        if false {
            if true {
                fmt.Println(p0, p0)
                p0 -= 2
                var v12 int16 = 32767
                _ = v12
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                switch fi10 {
                case 24:
                    v14 := GenS2{F0: 928}
                    v15 := GenS2{F0: 1}
                    if (366 < 0) {
                        if false {
                            fmt.Println(fi13, fi10)
                        }
                    }
                    var v16 float64 = 0.1
                    _ = v14
                    _ = v15
                    _ = v16
                case -998:
                    var v17 int16 = -1
                    _ = v17
                default:
                    v18 := GenS3{F0: (!(!((-2147483648 == -391) || (!(!false))))), F1: 2, F2: 1}
                    _ = v18
                }
                if len(v11) > 0 {
                    v11[0] = (fi10 - (^(fi10)))
                }
                _ = fi13
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                switch fi19 {
                case -1:
                    if len(v11) > 0 {
                        v11[0] = (fi10 & (fi10 % (((fi19 / ((-794) | 1))) | 1)))
                    }
                case 9223372036854775807:
                    p0 -= 128
                    _ = len(v11)
                    var v20 uint = 880
                    _ = v20
                case 4294967295:
                    var v21 uint64 = 813
                    _ = v21
                }
                fmt.Println(fi19, fi10)
                var v22 uint16 = 278
                _ = fi19
                _ = v22
            }
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v8
    _ = v9
    return 0
}
func genFn1() int64 {
    var v23 uint32 = 964
    v24 := []int{42, -2147483648}
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    _ = v23
    _ = v24
    return -9223372036854775808
}

func main() {
    v1 := []string{}
    var v2 uint = 2
    var v3 uint32 = 2147483648
    v4 := make(map[int]string)
    var v5 uint8 = 127
    var v6 int32 = 120
    v25 := GenS4{F0: -1, F1: (v5 & v5), F2: (v3 | 0)}
    fmt.Println(v6, v5)
    v6 = v6
    v1 = append(v1, "foo bar")
    v26 := GenS5{F0: 127, F1: 2, F2: 0}
    var v27 uint = (v2 + v2)
    v2 += (v2 - 255)
    var v28 uint64 = 370
    switch ("" + (("\n\t" + "a") + (("\n\t" + "foo bar") + ("\n\t" + "\n\t")))) {
    case "a":
        var v29 int32 = (v6 & v6)
        _ = v29
    case "":
        for fi30 := 0; fi30 < 12; fi30++ {
            for fi31 := 0; fi31 < 7; fi31++ {
                if true {
                    delete(v4, (fi30 >> ((914) & 63)))
                    switch ("hello" + "") {
                    case "":
                        if true {
                            v32 := GenS6{F0: (v28 / ((v28) | 1)), F1: 1000, F2: (42 == uint8(v2))}
                            _ = v32
                        } else {
                            switch ("" + "") {
                            case "a":
                                v28 += 1
                            default:
                                if (!false) {
                                    v33 := GenS7{F0: uint16(v5), F1: (1.5 >= 1e10)}
                                    _ = len(v1)
                                    _ = v33
                                } else {
                                    v34 := GenS6{F0: (v28 % ((v28) | 1)), F1: 880, F2: (623 <= 2147483647)}
                                    v27 = (v27 ^ v27)
                                    _ = len(v4)
                                    v27 = 0
                                    _ = v34
                                }
                                switch "a" {
                                case "a":
                                    var v35 uint64 = v28
                                    _ = v35
                                default:
                                    v2 = 404
                                }
                                v28 -= 18446744073709551615
                                var v36 bool = false
                                fmt.Println(v3)
                                _ = v36
                            }
                            _ = len(v1)
                            v3 += v3
                        }
                    }
                    var v37 bool = false
                    var v38 string = "foo bar"
                    _ = v37
                    _ = v38
                }
                _ = fi31
            }
            _ = fi30
        }
    default:
        for fi39 := 0; fi39 < 4; fi39++ {
            var v40 int16 = 32767
            _ = fi39
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 10; fi41++ {
        var v42 uint = uint(v3)
        var v43 uint64 = (^(v28))
        v5 = 0
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println(v28, v5, v5)
    v6 += -70
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
