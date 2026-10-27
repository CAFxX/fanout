package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int16; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint8; F1 uint; F2 uint }
type GenS7 struct { F0 uint; F1 int16; F2 int16 }
func genFn0(p0 int64) uint {
    fmt.Println(p0)
    p0 = p0
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return 456
}
func genFn1(p0 string, p1 uint64, p2 int) string {
    p0 = p0
    fmt.Println(p2)
    p2 = p2
    var v5 string = (p0 + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}

func main() {
    v1 := []uint32{47, 509, 0, 0}
    v2 := []string{"hello", ("" + "foo bar"), "hello"}
    var v3 int = -975
    var v4 int = v3
    if (42 <= 42) {
        _ = len(v1)
    } else {
        v4 = 1
        if (false && true) {
            if true {
                if true {
                    fmt.Println(v3)
                    for fi6 := 0; fi6 < 4; fi6++ {
                        v7 := GenS1{F0: 42}
                        var v8 int8 = -1
                        _ = fi6
                        _ = v7
                        _ = v8
                    }
                } else {
                    v3 += -9223372036854775808
                    if true {
                        v1 = append(v1, 1)
                        var v9 int16 = -618
                        var v10 string = (("\n\t" + "a") + (((("\n\t" + "\n\t") + ("foo bar" + "")) + "") + ""))
                        _ = v10
                        _ = v9
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
                var v11 uint16 = 815
                for fi12 := 0; fi12 < 5; fi12++ {
                    switch v3 {
                    case 32767:
                        v11 = (v11 * v11)
                    case 128:
                        _ = len(v1)
                    default:
                        v13 := GenS1{F0: v11}
                        _ = v13
                    }
                    switch (-(v3)) {
                    case 4294967295:
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    case 32767:
                        for fi14 := 0; fi14 < 18; fi14++ {
                            switch (^(fi14)) {
                            case -1:
                                switch "\n\t" {
                                case "hello":
                                    v3 -= 4294967295
                                case "":
                                    fmt.Println(fi12)
                                case "\n\t":
                                    if false {
                                        if true {
                                            v15 := GenS2{F0: 4294967295, F1: -1, F2: 1}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v15
                                        }
                                        var v16 uint8 = 57
                                        v16 = v16
                                        _ = v16
                                    } else {
                                        v11 -= 476
                                        fmt.Println(v4, v11, fi12)
                                        v11 = v11
                                    }
                                }
                                v3 += (v3 + (fi12 >> ((42) & 63)))
                            case 128:
                                v1 = append(v1, 167)
                            case 1:
                                if (true && true) {
                                    v17 := GenS2{F0: 1, F1: 1000, F2: 1}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v17
                                }
                            default:
                                v3 -= fi14
                            }
                            var v18 int8 = -66
                            _ = fi14
                            _ = v18
                        }
                    }
                    _ = fi12
                }
                _ = v11
            }
            var v19 uint32 = 404
            _ = v19
        } else {
            v1 = append(v1, 2147483647)
            v20 := GenS1{F0: 261}
            v4 -= (v3 % (((v4 % ((638) | 1))) | 1))
            _ = v20
        }
        var v21 int32 = int32(v4)
        _ = v21
    }
    fmt.Println(v4, v4)
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 float32 = ((0.1 + 1.5) + float32(float64(3.14159)))
        for fi24 := 0; fi24 < 14; fi24++ {
            if (2 >= 42) {
                v23 += 1.5
                if true {
                    v3 = (v3 * (-(v3)))
                } else {
                    var v25 string = "hello"
                    switch 0 {
                    case 1:
                        if true {
                            var v26 int = v4
                            _ = len(v2)
                            _ = v26
                        } else {
                            switch (("foo bar" + ("\n\t" + v25)) + "a") {
                            case "":
                                v27 := GenS3{F0: 669}
                                _ = v27
                            }
                            if false {
                                _ = len(v2)
                                var v28 bool = (true && false)
                                _ = v28
                            } else {
                                var v29 int = (v3 * v3)
                                _ = v29
                            }
                        }
                    case 42:
                        switch 0 {
                        case -18:
                            fmt.Println(fi22, v25)
                        }
                    case -904:
                        v25 = v25
                    default:
                        v30 := GenS2{F0: 65535, F1: 1, F2: 42}
                        _ = v30
                    }
                    _ = v25
                }
                var v31 uint = 18446744073709551615
                _ = v31
            } else {
                fmt.Println(v4, fi22)
                switch ("a" + "a") {
                case "a":
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                case "":
                    v3 -= (fi24 ^ (fi24 << ((1) & 63)))
                default:
                    fmt.Println(v4, v4)
                }
                var v32 uint16 = 255
                v3 = fi22
                var v33 uint64 = 18446744073709551615
                v23 -= -2.25
                _ = v32
                _ = v33
            }
            v3 = 9223372036854775807
            var v34 string = "hello"
            _ = fi24
            _ = v34
        }
        v35 := GenS4{F0: -2147483648}
        _ = fi22
        _ = v23
        _ = v35
    }
    fmt.Println(v4, v4)
    switch (v3 + (v3 % ((v3) | 1))) {
    case 849:
        v3 = (v4 % (((-(v3))) | 1))
    }
    var v36 int8 = -128
    v37 := GenS5{F0: 65535, F1: -32768}
    var v38 int64 = -256
    if true {
        v4 = v3
        v36 += v36
        v36 += (v36 % ((v36) | 1))
    } else {
        var v39 uint = 0
        v40 := GenS3{F0: 1}
        _ = v39
        _ = v40
    }
    v41 := GenS3{F0: 2}
    var v42 bool = false
    for fi43 := 0; fi43 < 5; fi43++ {
        fmt.Println(v36, v4, v4)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v44 string = ""
        _ = fi43
        _ = v44
    }
    var v45 string = "\n\t"
    if true {
        var v46 uint32 = 42
        _ = v46
    }
    if ((v38 < int64(v36)) && true) {
        v4 -= v3
        var v47 int16 = -517
        fmt.Println(v3, v36)
        _ = v47
    } else {
        switch v45 {
        case "a":
            v36 = 0
        case "\n\t":
            _ = len(v1)
        default:
            if (!(v42 || v42)) {
                v48 := GenS6{F0: 163, F1: 449, F2: 18446744073709551615}
                _ = v48
            } else {
                if true {
                    if (((-2.25 / 1.0) + 1.5) >= -2.25) {
                        v42 = (v42 && v42)
                        var v49 uint = 444
                        switch (v3 - (v4 + v4)) {
                        case 42:
                            for fi50 := 0; fi50 < 12; fi50++ {
                                for fi51 := 0; fi51 < 9; fi51++ {
                                    if (2 < 0) {
                                        var v52 uint32 = 4294967295
                                        var v53 uint16 = 255
                                        _ = v52
                                        _ = v53
                                    } else {
                                        v3 = int(v49)
                                    }
                                    var v54 int8 = v36
                                    _ = fi51
                                    _ = v54
                                }
                                _ = fi50
                            }
                        case 128:
                            var v55 string = "a"
                            _ = v55
                        default:
                            v36 -= 1
                        }
                        _ = v49
                    }
                }
                for fi56 := 0; fi56 < 7; fi56++ {
                    for fi57 := 0; fi57 < 11; fi57++ {
                        var v58 uint64 = 2
                        var v59 uint64 = v58
                        v4 -= (v3 | fi57)
                        _ = len(v2)
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                    _ = fi56
                }
                for fi60 := 0; fi60 < 11; fi60++ {
                    v61 := GenS7{F0: 1, F1: 0, F2: -1}
                    fmt.Println(v4)
                    _ = fi60
                    _ = v61
                }
            }
        }
        v4 += 32767
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v45
}
