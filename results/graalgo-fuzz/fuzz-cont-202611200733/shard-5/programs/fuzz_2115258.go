package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int8; F1 int8; F2 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 uint16 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 int32; F1 uint32; F2 string }
func genFn0(p0 int, p1 string, p2 int8) string {
    fmt.Println(p0, p1)
    switch p0 {
    case 32768:
        var v5 int = (p0 * -9223372036854775808)
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1() uint32 {
    v6 := []int{1}
    var v7 float32 = 0.0
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: "", F1: 243}
        v7 += v7
        _ = fi8
        _ = v9
    }
    if true {
        v7 += v7
    } else {
        var v10 int = -482
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS2{F0: 18446744073709551615, F1: 0}
        _ = fi11
        _ = v12
    }
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint = 1
    v2 := []int{}
    v3 := []uint32{0}
    var v4 uint16 = 0
    switch 32768 {
    case -9223372036854775808:
        v1 = v1
    case 2147483647:
        if ((!false) || true) {
            v13 := GenS1{F0: "foo bar", F1: 2}
            v14 := GenS3{F0: -128, F1: 127, F2: v4}
            if len(v3) > 0 {
                v3[0] = 144
            }
            _ = v13
            _ = v14
        }
    case 127:
        fmt.Println(v4)
    default:
        var v15 int64 = -9223372036854775808
        _ = v15
    }
    _ = len(v2)
    for fi16 := 0; fi16 < 2; fi16++ {
        switch fi16 {
        case -9223372036854775808:
            var v17 float32 = (1e10 + -2.25)
            _ = v17
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi16
    }
    v4 += v4
    var v18 float64 = 0.0
    if (0 == 2) {
        v19 := GenS4{F0: 42, F1: -1, F2: v4}
        var v20 uint8 = 164
        switch "" {
        case "hello":
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        default:
            if (!((float32(v18) / float32(v18)) < float32((0.1 / v18)))) {
                if false {
                    var v21 uint16 = (v4 >> (((v1 ^ v1)) & 63))
                    if (32767 <= 127) {
                        if (false || (-1 == 2147483647)) {
                            fmt.Println(v18)
                        }
                        fmt.Println(v18, v4, v20)
                    } else {
                        var v22 int64 = 42
                        v21 = (v4 % ((uint16(v20)) | 1))
                        var v23 int8 = -2
                        _ = v22
                        _ = v23
                    }
                    if (!true) {
                        v20 = v20
                    }
                    _ = v21
                }
                _ = len(v2)
                if false {
                    var v24 float32 = 0.0
                    if true {
                        _ = len(v2)
                    } else {
                        if (!false) {
                            fmt.Println(v20)
                            for fi25 := 0; fi25 < 14; fi25++ {
                                if false {
                                    v4 = v4
                                    var v26 uint = v1
                                    v27 := GenS1{F0: ("" + (("" + "") + "")), F1: 42}
                                    _ = v26
                                    _ = v27
                                }
                                switch (^(fi25)) {
                                case 0:
                                    v24 = float32(v18)
                                case 32767:
                                    v20 += 127
                                case -9223372036854775808:
                                    for fi28 := 0; fi28 < 8; fi28++ {
                                        v18 = (float64(v24) + (v18 + (0.1 / v18)))
                                        _ = fi28
                                    }
                                default:
                                    var v29 uint32 = 2
                                    _ = v29
                                }
                                _ = fi25
                            }
                            v20 = (v20 ^ (v20 - (v20 % (((^(v20))) | 1))))
                        } else {
                            fmt.Println(v20)
                        }
                        var v30 string = "foo bar"
                        _ = v30
                    }
                    _ = v24
                } else {
                    v1 -= (v1 & 65535)
                    var v31 uint16 = v4
                    _ = v31
                }
            } else {
                var v32 string = ("hello" + "hello")
                switch (v32 + ("foo bar" + "a")) {
                case "a":
                    v33 := GenS5{F0: v32, F1: 2147483647}
                    _ = v33
                }
                _ = v32
            }
        }
        _ = v19
        _ = v20
    }
    if false {
        var v34 string = ("foo bar" + "foo bar")
        _ = v34
    }
    if false {
        if true {
            v35 := GenS3{F0: 42, F1: 100, F2: (v4 / (((v4 - (v4 + v4))) | 1))}
            _ = v35
        } else {
            v36 := GenS1{F0: ("foo bar" + "\n\t"), F1: 400}
            _ = v36
        }
        v4 -= v4
    }
    var v37 string = ("" + "a")
    var v38 int16 = 42
    switch 190 {
    case -50:
        var v39 int8 = 1
        _ = v39
    default:
        switch (("a" + "") + v37) {
        case "":
            v18 += (v18 + 3.14159)
        case "a":
            var v40 int32 = -2147483648
            _ = v40
        default:
            if (!(true || (-32768 == (v38 % (((v38 / ((v38) | 1))) | 1))))) {
                v3 = append(v3, 595)
            }
        }
    }
    fmt.Println(v18, v18, v1)
    for fi41 := 0; fi41 < 18; fi41++ {
        v4 = (v4 & 2)
        _ = fi41
    }
    v42 := GenS6{F0: -675, F1: 41, F2: v37}
    switch 243 {
    case 9223372036854775807:
        v37 = v37
    default:
        fmt.Println(v1, v37, v18)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
}
