package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int64 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int64; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 bool }
func genFn0() int16 {
    var v5 float64 = (((0.1 * -2.25) + ((-2.25 + -2.25) + -2.25)) - 3.14159)
    for fi6 := 0; fi6 < 14; fi6++ {
        v5 -= v5
        v7 := []int{(fi6 % ((127) | 1)), fi6, fi6, (fi6 * fi6)}
        switch fi6 {
        case 9223372036854775807:
            var v8 int16 = -32768
            _ = v8
        case 602:
            var v9 uint = 436
            _ = v9
        default:
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := GenS1{F0: ((float32(v5) - 3.14159) >= (float32(v5) * 0.0)), F1: ((v5 * float64(float32(1.5))) >= 1.5), F2: 1}
                fmt.Println(fi10, v5)
                if true {
                    v12 := GenS1{F0: (true || (false || true)), F1: (!true), F2: -227}
                    if (!("\n\t" > "a")) {
                        fmt.Println(fi6, v5, fi6)
                        v5 -= v5
                    } else {
                        v13 := GenS2{F0: false, F1: ""}
                        v5 -= (v5 - v5)
                        var v14 string = "hello"
                        _ = v13
                        _ = v14
                    }
                    fmt.Println(fi10, fi10)
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
        }
        _ = fi6
        _ = v7
    }
    _ = v5
    return 186
}

func main() {
    var v1 uint = 42
    var v2 uint8 = 2
    var v3 string = "foo bar"
    v4 := make(map[string]string)
    fmt.Println(v1, v1)
    var v15 int32 = -1
    for fi16 := 0; fi16 < 15; fi16++ {
        if (4294967295 != 1) {
            for fi17 := 0; fi17 < 7; fi17++ {
                v15 = (v15 << ((42) & 63))
                var v18 uint16 = 32767
                v15 += 42
                v1 -= (v1 * v1)
                if false {
                    if (false && false) {
                        v19 := GenS2{F0: ((true && false) && (true || true)), F1: "a"}
                        v20 := GenS3{F0: 764, F1: (1 >= 764), F2: 165}
                        var v21 int64 = -9223372036854775808
                        var v22 float64 = (((0.0 * -2.25) / 1.0) - -2.25)
                        _ = v19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                } else {
                    fmt.Println(v15, v3)
                    var v23 float32 = (3.14159 * -2.25)
                    v15 = (v15 * v15)
                    _ = v23
                }
                var v24 bool = false
                _ = fi17
                _ = v18
                _ = v24
            }
        }
        fmt.Println(fi16, v1, v3)
        switch "hello" {
        case "":
            for fi25 := 0; fi25 < 19; fi25++ {
                v26 := []int{int(v1), fi16}
                v1 -= 0
                switch 1 {
                case 32768:
                    switch v3 {
                    case "\n\t":
                        var v27 float32 = (float32(1.5) * 0.0)
                        var v28 int16 = 1
                        switch v3 {
                        case "hello":
                            v2 = v2
                        default:
                            v28 = (v28 >> ((v1) & 63))
                        }
                        _ = v27
                        _ = v28
                    case "hello":
                        if false {
                            v29 := v4["foo bar"]
                            _ = v29
                        }
                    default:
                        if (-2147483648 >= 42) {
                            switch -594 {
                            case -1:
                                if (4294967295 != (fi16 / (((fi25 * (fi16 | (fi25 + fi25)))) | 1))) {
                                    v3 = ""
                                    if len(v26) > 0 {
                                        v26[0] = 9223372036854775807
                                    }
                                } else {
                                    v30 := GenS3{F0: 0, F1: (("" + v3) < v3), F2: 42}
                                    if ((1.5 * 0.0) <= (0.1 + (float32(1e10) * 0.0))) {
                                        v31 := GenS1{F0: (9223372036854775807 != -866), F1: false, F2: 615}
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        var v32 uint8 = (v2 | 128)
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = v30
                                }
                            }
                        }
                    }
                default:
                    fmt.Println(fi25, v2)
                }
                _ = fi25
                _ = v26
            }
        default:
            v33 := GenS1{F0: (false && true), F1: false, F2: -606}
            _ = v33
        }
        _ = fi16
    }
    v15 -= (v15 % ((1) | 1))
    if true {
        var v34 float32 = 0.0
        _ = v34
    } else {
        delete(v4, ("foo bar" + v3))
        v1 -= v1
        if true {
            v4[("" + "\n\t")] = v3
            fmt.Println(v2, v3)
            v3 = ("foo bar" + (((((v3 + v3) + (("foo bar" + "hello") + "")) + (v3 + "foo bar")) + ("foo bar" + v3)) + "hello"))
        } else {
            for fi35 := 0; fi35 < 8; fi35++ {
                v36 := []int{fi35}
                v37 := GenS4{F0: 42, F1: 32767}
                for fi38 := 0; fi38 < 8; fi38++ {
                    var v39 uint = 62
                    v40 := GenS5{F0: true}
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                _ = fi35
                _ = v36
                _ = v37
            }
        }
    }
    v1 = v1
    v1 -= 228
    v41 := GenS6{F0: 513, F1: (42 < 550)}
    var v42 int16 = -88
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
