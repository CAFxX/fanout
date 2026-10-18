package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint64) uint {
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{fi7, fi7, fi7, fi7}
        p0 = (p0 ^ 4294967295)
        fmt.Println(p0, p0)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 float32 = ((((0.1 - 0.1) - 0.0) - (-2.25 / (-2.25 * 1e10))) * (0.0 * -2.25))
            var v12 string = "foo bar"
            v11 += (0.1 / 0.1)
            if false {
                v13 := GenS1{F0: 42, F1: -2147483648, F2: -414}
                v12 = (v12 + v12)
                switch v12 {
                case "hello":
                    switch (fi10 ^ fi9) {
                    case -1:
                        var v14 uint8 = 93
                        _ = v14
                    default:
                        switch -2147483648 {
                        case -9223372036854775808:
                            v11 += 3.14159
                        }
                    }
                case "a":
                    v12 = (("hello" + v12) + v12)
                }
                _ = v13
            } else {
                var v15 float32 = v11
                _ = v15
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        var v16 float32 = 3.14159
        p0 -= (p0 | p0)
        _ = fi9
        _ = v16
    }
    if true {
        p0 += 2
        p0 += 556
        var v17 int32 = -679
        _ = v17
    } else {
        if (725 == 2) {
            for fi18 := 0; fi18 < 6; fi18++ {
                for fi19 := 0; fi19 < 5; fi19++ {
                    p0 -= p0
                    _ = fi19
                }
                for fi20 := 0; fi20 < 12; fi20++ {
                    var v21 uint64 = (^(p0))
                    _ = fi20
                    _ = v21
                }
                _ = fi18
            }
            v22 := GenS2{F0: 127}
            p0 = 4294967295
            _ = v22
        }
        var v23 int8 = 127
        _ = v23
    }
    _ = p0
    return 1
}
func genFn1() int32 {
    v24 := []int{240}
    var v25 int32 = -247
    _ = v24
    _ = v25
    return (v25 >> ((18446744073709551615) & 63))
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]bool)
    v3 := []uint32{2147483647, 121, 79, 900}
    v4 := []int{2147483647, 32768}
    var v5 uint64 = 18446744073709551615
    v6 := make(map[uint32]bool)
    switch -9223372036854775808 {
    case 1:
        var v26 string = "a"
        _ = v26
    case 42:
        v1 = ("foo bar" + "foo bar")
    case 32767:
        v5 -= v5
    default:
        fmt.Println(v5, v1)
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        for fi28 := 0; fi28 < 1; fi28++ {
            v1 = ("a" + "")
            _ = fi28
        }
        fmt.Println(v1, v5, v1)
        _ = fi27
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1, v1)
    if (18446744073709551615 > 2) {
        var v29 uint64 = (v5 + (v5 | v5))
        v30 := v2[2147483647]
        _ = v29
        _ = v30
    }
    var v31 float32 = float32((-2.25 + -2.25))
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 uint64 = 42
        switch (fi32 | (fi32 | (fi32 | fi32))) {
        case -467:
            v5 -= (v33 & 858)
        }
        if ((-2.25 + (0.1 / 1.0)) < 1.5) {
            v34 := GenS1{F0: 197, F1: 4294967295, F2: 990}
            if ((-9223372036854775808 > 0) && (1000 == -1)) {
                v31 = v31
            }
            v35 := GenS3{F0: true, F1: fi32, F2: 0}
            _ = v34
            _ = v35
        } else {
            var v36 int64 = 42
            for fi37 := 0; fi37 < 14; fi37++ {
                v38 := GenS4{F0: v36}
                if (!true) {
                    if true {
                        switch 4294967295 {
                        case 888:
                            if true {
                                fmt.Println(v5, fi32)
                                v39 := GenS1{F0: 236, F1: fi32, F2: 32767}
                                v1 = v1
                                if false {
                                    v40 := GenS5{F0: 2147483648}
                                    v31 = 0.0
                                    v36 -= (v36 >> ((3) & 63))
                                    _ = v40
                                } else {
                                    var v41 bool = true
                                    if len(v4) > 0 {
                                        v4[0] = 127
                                    }
                                    fmt.Println(fi37, v36)
                                    _ = v41
                                }
                                _ = v39
                            }
                        default:
                            if false {
                                v36 -= v36
                                v42 := GenS4{F0: v36}
                                _ = v42
                            } else {
                                switch fi37 {
                                case 2147483647:
                                    v3 = append(v3, 65535)
                                case 4294967295:
                                    v43 := GenS5{F0: 65535}
                                    _ = v43
                                case 0:
                                    v44 := GenS2{F0: 1}
                                    _ = v44
                                default:
                                    var v45 int = (fi37 + fi32)
                                    _ = v45
                                }
                            }
                        }
                        v1 = v1
                        v46 := GenS4{F0: (v36 + (v36 % (((-(v36))) | 1)))}
                        _ = v46
                    }
                    v6[130] = (42 > 80)
                    fmt.Println(v31, v5)
                    _ = len(v2)
                    fmt.Println(v31, v31, v1)
                } else {
                    v47 := GenS3{F0: true, F1: 0, F2: 0}
                    fmt.Println(fi32)
                    _ = v47
                }
                var v48 float64 = 0.0
                _ = fi37
                _ = v38
                _ = v48
            }
            _ = v36
        }
        _ = fi32
        _ = v33
    }
    v5 = v5
    var v49 uint = genFn0(1)
    v49 = (v49 - v49)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
