package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int32; F1 int32; F2 uint8 }
type GenS4 struct { F0 int8; F1 int16 }
func genFn0(p0 bool, p1 string) int16 {
    v7 := GenS1{F0: 42, F1: 1}
    if p0 {
        var v8 string = (p1 + (((((p1 + "foo bar") + p1) + ((p1 + (((p1 + (((("hello" + ("a" + "hello")) + ((p1 + "hello") + "a")) + (((p1 + "foo bar") + ((p1 + (p1 + (p1 + ("foo bar" + p1)))) + ("foo bar" + ""))) + (((((p1 + (p1 + ("hello" + p1))) + ((p1 + ("foo bar" + p1)) + p1)) + p1) + "\n\t") + (p1 + ("hello" + (p1 + "foo bar")))))) + (p1 + p1))) + "") + p1)) + "a")) + p1) + (p1 + p1)))
        _ = v8
    } else {
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 11; fi10++ {
                fmt.Println(fi10, p1)
                p0 = (!(42 > 100))
                var v11 uint16 = 84
                _ = fi10
                _ = v11
            }
            if p0 {
                if (-35 < -75) {
                    for fi12 := 0; fi12 < 18; fi12++ {
                        var v13 int8 = -61
                        _ = fi12
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 12; fi14++ {
                        switch p1 {
                        case "":
                            var v15 uint32 = 7
                            _ = v15
                        case "a":
                            v16 := GenS1{F0: -655, F1: 4294967295}
                            _ = v16
                        case "hello":
                            var v17 bool = p0
                            _ = v17
                        }
                        _ = fi14
                    }
                }
            }
            fmt.Println(p0)
            _ = fi9
        }
    }
    if ((true && true) || true) {
        if true {
            switch -428 {
            case 32767:
                var v18 uint8 = 127
                _ = v18
            case 926:
                var v19 int8 = 0
                _ = v19
            }
        }
        if (351 <= 2) {
            var v20 uint16 = 0
            var v21 float32 = 1.5
            _ = v20
            _ = v21
        } else {
            v22 := GenS2{F0: 1, F1: 541}
            var v23 float64 = 1e10
            _ = v22
            _ = v23
        }
        p1 = ""
    }
    _ = p0
    _ = p1
    _ = v7
    return 1000
}
func genFn1(p0 uint32, p1 int, p2 uint16) string {
    var v24 uint32 = 42
    if (!true) {
        var v25 int64 = 42
        if false {
            var v26 uint = 255
            v27 := []int{}
            _ = v26
            _ = v27
        } else {
            v24 += (p0 >> ((18446744073709551615) & 63))
        }
        _ = v25
    } else {
        var v28 float32 = (1.5 * 3.14159)
        _ = v28
    }
    v29 := GenS1{F0: 963, F1: (v24 << ((259) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v29
    return "foo bar"
}

func main() {
    var v1 uint32 = 83
    var v2 uint16 = 0
    v3 := []uint32{(v1 & (v1 | (v1 % (((v1 * (v1 % ((v1) | 1)))) | 1)))), v1, 1, v1}
    var v4 int64 = -293
    var v5 uint32 = 65535
    var v6 int64 = (v4 * 1)
    if false {
        v6 = (^(v6))
    }
    var v30 int64 = -967
    if true {
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 uint64 = 2
            v33 := GenS3{F0: 1, F1: -2147483648, F2: uint8(v2)}
            _ = fi31
            _ = v32
            _ = v33
        }
        var v34 uint8 = 255
        v5 = (v5 << ((0) & 63))
        _ = v34
    }
    v35 := GenS3{F0: 2147483647, F1: -256, F2: 2}
    fmt.Println(v30, v5, v30)
    var v36 uint64 = 18446744073709551615
    var v37 uint8 = 100
    if (v30 <= (v4 / ((v6) | 1))) {
        var v38 string = "hello"
        for fi39 := 0; fi39 < 18; fi39++ {
            if false {
                for fi40 := 0; fi40 < 3; fi40++ {
                    switch (v38 + v38) {
                    case "":
                        if false {
                            v30 += 3
                            if true {
                                fmt.Println(v30)
                                var v41 float32 = 1e10
                                if ((("" + "a") + v38) == v38) {
                                    switch v38 {
                                    case "\n\t":
                                        if true {
                                            if (fi39 != int(v6)) {
                                                v2 = (v2 + v2)
                                                var v42 bool = true
                                                _ = v42
                                            } else {
                                                if (true || (1 >= 417)) {
                                                    if true {
                                                        if (true && (v36 < v36)) {
                                                            var v43 uint8 = (v37 / (((^(v37))) | 1))
                                                            _ = v43
                                                        }
                                                    }
                                                }
                                                v2 -= (v2 % ((v2) | 1))
                                                v5 = v5
                                            }
                                            fmt.Println(v30, v4)
                                            if len(v3) > 0 {
                                                v3[0] = 2147483648
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = 2
                                            }
                                            if false {
                                                v3 = append(v3, v1)
                                            } else {
                                                if false {
                                                    _ = len(v3)
                                                    var v44 uint8 = (v37 + uint8(v2))
                                                    v30 = (v30 ^ v4)
                                                    var v45 uint = 255
                                                    _ = len(v3)
                                                    _ = v44
                                                    _ = v45
                                                }
                                            }
                                        }
                                    }
                                    switch fi40 {
                                    case 1:
                                        v41 = 0.1
                                    default:
                                        fmt.Println(v5, v37)
                                    }
                                }
                                var v46 string = v38
                                _ = v41
                                _ = v46
                            } else {
                                v2 += v2
                                if (false && (607 == 255)) {
                                    v47 := GenS4{F0: 42, F1: -1}
                                    _ = v47
                                }
                                var v48 float64 = (float64(0.0) / (0.0 - (0.1 + ((1e10 - 0.0) + ((1.5 - 3.14159) - (0.0 + -2.25))))))
                                var v49 int = (fi39 / ((fi40) | 1))
                                v3 = append(v3, v5)
                                _ = v48
                                _ = v49
                            }
                            var v50 uint16 = v2
                            _ = v50
                        } else {
                            v51 := GenS4{F0: 100, F1: 1}
                            v3 = append(v3, (^(v1)))
                            v37 = 42
                            switch fi39 {
                            case 42:
                                v2 = 255
                            case -988:
                                _ = len(v3)
                            }
                            _ = v51
                        }
                    case "foo bar":
                        var v52 string = v38
                        _ = v52
                    case "\n\t":
                        if (true || true) {
                            v38 = "hello"
                            fmt.Println(v5, v5, v2)
                            var v53 int16 = 42
                            if true {
                                v2 = v2
                            }
                            _ = v53
                        }
                    default:
                        var v54 uint16 = 42
                        _ = v54
                    }
                    _ = fi40
                }
                var v55 float32 = (1e10 + ((float32((((1e10 - 0.0) + 3.14159) - (0.0 * (float64(float32(1e10)) + 3.14159)))) - (1e10 / 1.5)) - float32((0.0 * 0.1))))
                v1 = v5
                _ = v55
            } else {
                var v56 int8 = 100
                var v57 uint = 4294967295
                _ = v56
                _ = v57
            }
            _ = len(v3)
            _ = fi39
        }
        _ = v38
    } else {
        var v58 int = 42
        _ = v58
    }
    if true {
        var v59 int16 = -324
        _ = v59
    }
    var v60 uint32 = v1
    v61 := GenS4{F0: 90, F1: -756}
    v1 -= 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
