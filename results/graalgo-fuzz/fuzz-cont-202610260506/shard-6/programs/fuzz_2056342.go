package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint16; F2 int32 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 4; fi6++ {
        if false {
            var v7 float32 = (-2.25 - (1e10 + -2.25))
            switch "" {
            case "foo bar":
                if true {
                    for fi8 := 0; fi8 < 4; fi8++ {
                        if (1 == -128) {
                            fmt.Println(v7, fi6)
                            if true {
                                fmt.Println(fi8, v7)
                            }
                        }
                        switch ("foo bar" + "hello") {
                        case "hello":
                            var v9 int = (fi8 - (fi6 / ((fi6) | 1)))
                            _ = v9
                        default:
                            var v10 int = (fi6 ^ (fi6 * fi6))
                            _ = v10
                        }
                        _ = fi8
                    }
                    v11 := GenS1{F0: false, F1: 1}
                    fmt.Println(fi6, v7)
                    _ = v11
                }
            }
            var v12 string = "hello"
            _ = v12
            _ = v7
        }
        var v13 int64 = 9223372036854775807
        _ = fi6
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        if true {
            fmt.Println(fi14)
            var v15 int8 = -1
            var v16 int32 = 2147483647
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    return 29
}
func genFn1(p0 uint64, p1 bool) int8 {
    p0 += 0
    var v17 int16 = 0
    if ((65535 == 0) || ((-2.25 * (3.14159 * 3.14159)) != -2.25)) {
        for fi18 := 0; fi18 < 15; fi18++ {
            p0 = (p0 % ((4294967295) | 1))
            var v19 int16 = (v17 * v17)
            for fi20 := 0; fi20 < 19; fi20++ {
                v17 = -32768
                if false {
                    v21 := []int{(fi20 << ((604) & 63)), fi18, fi18}
                    v19 -= v17
                    _ = v21
                }
                if p1 {
                    v22 := GenS2{F0: 116, F1: false, F2: 100}
                    _ = v22
                } else {
                    if p1 {
                        var v23 string = ("a" + ("a" + (("hello" + "\n\t") + "")))
                        _ = v23
                    } else {
                        v24 := []int{fi18, fi18, (fi18 & (fi18 + (fi18 - fi18)))}
                        var v25 int16 = v19
                        if (0.1 == ((0.0 / 1e10) / 1e10)) {
                            var v26 int32 = -418
                            v27 := GenS3{F0: p1}
                            switch "a" {
                            case "a":
                                if p1 {
                                    var v28 float32 = (1e10 - (1e10 + 0.0))
                                    v24 = append(v24, fi20)
                                    v26 += v26
                                    _ = v28
                                }
                            }
                            if false {
                                var v29 int8 = 1
                                _ = v29
                            } else {
                                v26 += v26
                                var v30 uint = 2
                                if len(v24) > 0 {
                                    v24[0] = 32767
                                }
                                _ = v30
                            }
                            if p1 {
                                v24 = append(v24, 4294967295)
                            }
                            var v31 int64 = int64(p0)
                            _ = v26
                            _ = v27
                            _ = v31
                        } else {
                            v25 = 91
                            _ = len(v24)
                            if p1 {
                                var v32 uint32 = 65535
                                _ = v32
                            }
                        }
                        v19 += (v25 >> ((42) & 63))
                        v33 := GenS1{F0: (0 == (p0 | (p0 * (p0 ^ p0)))), F1: -128}
                        var v34 uint = 1
                        _ = v24
                        _ = v25
                        _ = v33
                        _ = v34
                    }
                }
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
        v35 := []int{27}
        _ = v35
    } else {
        var v36 uint32 = 536
        v36 = v36
        if (!p1) {
            v37 := []int{1, -842, -493}
            switch (("" + "\n\t") + "hello") {
            case "":
                v36 = (v36 & 0)
            default:
                v38 := GenS2{F0: 47, F1: true, F2: 42}
                _ = v38
            }
            _ = v37
        } else {
            v39 := GenS2{F0: -1, F1: ((p1 && p1) && p1), F2: 0}
            if p1 {
                fmt.Println(p0, p1)
            }
            _ = v39
        }
        _ = v36
    }
    p1 = p1
    p0 += p0
    _ = p0
    _ = p1
    _ = v17
    return 127
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 2
    var v5 int32 = -2147483648
    var v40 uint64 = 0
    v5 = -107
    for fi41 := 0; fi41 < 8; fi41++ {
        if true {
            v40 += (v40 - 18446744073709551615)
            v42 := GenS1{F0: (!((1.5 + ((3.14159 * (-2.25 * (0.0 * 0.1))) * 0.0)) == ((1e10 / -2.25) + 0.1))), F1: 100}
            for fi43 := 0; fi43 < 3; fi43++ {
                v4 = (^(v4))
                _ = fi43
            }
            _ = v42
        }
        v44 := GenS4{F0: ("foo bar" + ""), F1: uint16(v40), F2: (v5 ^ v5)}
        v4 = v4
        _ = fi41
        _ = v44
    }
    v40 = v40
    delete(v2, 42)
    if (0.1 <= 3.14159) {
        delete(v2, 700)
        switch "foo bar" {
        case "a":
            var v45 int8 = 51
            _ = v45
        }
        var v46 uint = 2
        _ = v46
    } else {
        var v47 int = -525
        var v48 float32 = 0.0
        _ = v47
        _ = v48
    }
    v49 := GenS1{F0: false, F1: 42}
    var v50 float32 = (1.5 - (-2.25 + -2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v49
    _ = v5
    _ = v50
}
