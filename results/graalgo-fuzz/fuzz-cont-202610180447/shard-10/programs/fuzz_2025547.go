package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int16; F2 int64 }
type GenS3 struct { F0 int16; F1 int; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS5 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 0
    v3 := make(map[string]uint32)
    var v4 float64 = 1e10
    if (false || false) {
        for fi5 := 0; fi5 < 19; fi5++ {
            fmt.Println(fi5)
            _ = fi5
        }
        v4 -= (3.14159 + v4)
    } else {
        switch ((("" + "hello") + (("hello" + "") + "a")) + ("hello" + ((("a" + "hello") + "foo bar") + "foo bar"))) {
        case "\n\t":
            if ((true && true) || true) {
                v4 += v4
                if false {
                    for fi6 := 0; fi6 < 2; fi6++ {
                        v2 = v2
                        v7 := GenS1{F0: 2, F1: (fi6 | (fi6 | (-(fi6)))), F2: uint64(v2)}
                        _ = fi6
                        _ = v7
                    }
                    if false {
                        v8 := []int{}
                        _ = v8
                    } else {
                        v9 := GenS2{F0: 1, F1: 0, F2: 42}
                        var v10 int8 = -1
                        v4 += v4
                        _ = v10
                        _ = v9
                    }
                    v11 := GenS1{F0: 18446744073709551615, F1: 0, F2: uint64(v2)}
                    _ = v11
                } else {
                    v12 := []int{2147483647, -81, 4294967295}
                    if true {
                        fmt.Println(v4, v4)
                        if (true && (1 <= 0)) {
                            if (2 != 0) {
                                var v13 uint32 = 4294967295
                                _ = v13
                            }
                            for fi14 := 0; fi14 < 11; fi14++ {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                v2 -= 83
                                fmt.Println(v4)
                                _ = fi14
                            }
                            v4 -= 0.0
                        }
                    } else {
                        fmt.Println(v4)
                    }
                    _ = v12
                }
                for fi15 := 0; fi15 < 16; fi15++ {
                    if ((145 >= 584) || false) {
                        v4 += 0.1
                        if (((v2 ^ (v2 / (((v2 >> ((42) & 63))) | 1))) >= (v2 ^ (v2 - v2))) && true) {
                            var v16 int8 = 15
                            for fi17 := 0; fi17 < 9; fi17++ {
                                v18 := GenS3{F0: -1, F1: (fi17 * (^(fi15))), F2: 65535}
                                _ = fi17
                                _ = v18
                            }
                            v16 -= v16
                            _ = v16
                        }
                        if (42 <= 0) {
                            if (0 >= 1) {
                                v19 := v3["a"]
                                v20 := GenS4{F0: v19, F1: 514, F2: (v2 - (v2 % (((v2 | v2)) | 1)))}
                                _ = v19
                                _ = v20
                            } else {
                                v2 += 1
                            }
                        } else {
                            var v21 uint64 = 1
                            for fi22 := 0; fi22 < 4; fi22++ {
                                var v23 int16 = 1000
                                v24 := v3["foo bar"]
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            _ = v21
                        }
                    }
                    _ = fi15
                }
            } else {
                v4 = 1e10
            }
        case "hello":
            v2 -= (v2 & (v2 << ((255) & 63)))
        default:
            v4 = (0.1 + 0.1)
        }
        var v25 int = -1
        var v26 uint8 = v2
        _ = v25
        _ = v26
    }
    v1[801] = ("" + "a")
    var v27 int32 = 42
    var v28 int64 = 9223372036854775807
    var v29 uint64 = 1
    var v30 float32 = -2.25
    var v31 int64 = (v28 ^ v28)
    v32 := GenS5{F0: (v27 / ((v27) | 1))}
    switch "" {
    case "\n\t":
        var v33 uint16 = 255
        _ = v33
    case "":
        var v34 int = 9223372036854775807
        _ = v34
    case "foo bar":
        v2 = (v2 << ((65535) & 63))
    default:
        for fi35 := 0; fi35 < 4; fi35++ {
            var v36 uint16 = 42
            v29 = v29
            v3[("" + (("a" + "foo bar") + "\n\t"))] = 2147483648
            _ = fi35
            _ = v36
        }
    }
    v37 := GenS1{F0: 465, F1: 4294967295, F2: 18446744073709551615}
    delete(v3, ("hello" + ("hello" + ((("" + (((("hello" + "hello") + "") + "foo bar") + "\n\t")) + "foo bar") + ("a" + "")))))
    v38 := GenS4{F0: uint32(v2), F1: 255, F2: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v4
}
