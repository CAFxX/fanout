package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int16; F2 uint }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int64; F1 uint16; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 int32; F2 uint64 }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 int; F1 uint16; F2 uint8 }
func genFn0(p0 int16) int64 {
    var v5 uint = 18446744073709551615
    v6 := GenS1{F0: 42}
    if ((255 <= 818) || (!true)) {
        v5 = v5
        var v7 uint8 = uint8(v5)
        _ = v7
    } else {
        fmt.Println(v5)
        fmt.Println(p0)
        v8 := GenS2{F0: false, F1: p0, F2: v5}
        _ = v8
    }
    v9 := []int{}
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(v5, fi10, v5)
        var v11 float64 = 0.1
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v9
    return -612
}

func main() {
    v1 := []uint32{65535}
    var v2 string = "a"
    var v3 uint64 = 691
    var v4 uint8 = 0
    if true {
        if true {
            fmt.Println(v2)
            if false {
                _ = len(v1)
                v12 := GenS1{F0: v3}
                _ = v12
            } else {
                var v13 bool = true
                if v13 {
                    v3 = v3
                    var v14 bool = (42 != -32768)
                    var v15 float64 = 1.5
                    _ = v14
                    _ = v15
                } else {
                    v2 = v2
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v17 := GenS1{F0: v3}
                        if v13 {
                            v18 := GenS3{F0: -265, F1: -866}
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 6; fi19++ {
                            v2 = v2
                            if v13 {
                                if true {
                                    v20 := GenS2{F0: v13, F1: -32768, F2: 42}
                                    if (false || (v4 != v4)) {
                                        if false {
                                            var v21 int = fi16
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v21
                                        } else {
                                            var v22 uint32 = 2147483648
                                            _ = v22
                                        }
                                        var v23 bool = false
                                        _ = v23
                                    }
                                    _ = v20
                                }
                                v13 = v13
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = 4294967295
                                }
                                v1 = append(v1, 546)
                                var v24 uint64 = v3
                                v25 := GenS4{F0: 42, F1: 1, F2: ("hello" + ("" + v2))}
                                if (0.0 < 0.0) {
                                    fmt.Println(v2)
                                }
                                _ = v24
                                _ = v25
                            }
                            _ = fi19
                        }
                        _ = fi16
                        _ = v17
                    }
                }
                for fi26 := 0; fi26 < 15; fi26++ {
                    var v27 uint64 = (v3 % (((v3 ^ (v3 / (((v3 + (v3 ^ (v3 ^ (v3 >> ((396) & 63)))))) | 1)))) | 1))
                    fmt.Println(fi26)
                    _ = fi26
                    _ = v27
                }
                _ = v13
            }
            v28 := GenS5{F0: 42}
            _ = v28
        }
        if false {
            v2 = v2
            v2 = "foo bar"
            fmt.Println(v3)
        }
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS6{F0: true}
            v2 = "hello"
            _ = fi29
            _ = v30
        }
    }
    fmt.Println(v4, v4, v3)
    v4 -= (v4 & 128)
    fmt.Println(v4, v3, v3)
    v31 := GenS7{F0: ((v2 + v2) + ("a" + "")), F1: 0, F2: (v3 + (v3 | (v3 - v3)))}
    if (3.14159 != 1.5) {
        _ = len(v1)
        v32 := GenS8{F0: 1000, F1: v3}
        var v33 int16 = -909
        _ = v32
        _ = v33
    } else {
        fmt.Println(v4)
    }
    v2 = v2
    for fi34 := 0; fi34 < 1; fi34++ {
        for fi35 := 0; fi35 < 12; fi35++ {
            if (317 > 4294967295) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v1 = append(v1, 65535)
            }
            _ = fi35
        }
        switch fi34 {
        case 42:
            v36 := GenS2{F0: false, F1: -32768, F2: 409}
            _ = v36
        case -755:
            for fi37 := 0; fi37 < 4; fi37++ {
                if (629 <= 537) {
                    fmt.Println(v4, fi34)
                    switch ("a" + ("hello" + ("" + v2))) {
                    case "\n\t":
                        v38 := GenS9{F0: fi37, F1: 32767, F2: (v4 / (((v4 + v4)) | 1))}
                        _ = v38
                    case "a":
                        v3 -= 2
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = len(v1)
                }
                v3 = 42
                _ = fi37
            }
        default:
            v39 := GenS6{F0: false}
            _ = v39
        }
        _ = fi34
    }
    var v40 int8 = int8(v4)
    if (!true) {
        var v41 int = 4294967295
        v3 -= (v3 - 2)
        fmt.Println(v40, v2, v3)
        _ = v41
    }
    v40 -= (v40 & v40)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
}
