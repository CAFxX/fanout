package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint; F2 bool }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 uint32; F1 uint8; F2 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 145
    v3 := []uint32{}
    var v4 uint8 = v2
    v2 += v2
    switch -2147483648 {
    case 32768:
        var v5 string = ""
        _ = v5
    case -2147483648:
        v6 := GenS1{F0: ((float64((float32((1.5 * -2.25)) + 1.5)) + 1.5) >= 0.1), F1: (v2 >> ((65535) & 63))}
        _ = v6
    }
    if false {
        v7 := GenS2{F0: "a"}
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            v4 = v4
            v4 -= v2
            var v10 int = fi8
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := GenS3{F0: 209}
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 uint64 = 2
        _ = fi12
        _ = v13
    }
    v2 = (v4 & v4)
    if false {
        v2 -= 0
        v3 = append(v3, 2147483647)
    }
    fmt.Println(v2)
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            v4 -= 0
            if (true && true) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if (127 < 128) {
                    var v16 int64 = -9223372036854775808
                    fmt.Println(v16)
                    v2 += (v4 % ((v4) | 1))
                    _ = v16
                } else {
                    if false {
                        if (true && false) {
                            v2 = (v2 ^ v2)
                            v2 = (v4 - v4)
                        } else {
                            fmt.Println(fi14, fi14)
                            v1[(fi14 >> ((42) & 63))] = -9223372036854775808
                        }
                    } else {
                        if true {
                            v2 = (v2 | 127)
                        }
                    }
                    if true {
                        v4 -= (v2 + v2)
                        var v17 int32 = 574
                        var v18 int32 = v17
                        _ = v17
                        _ = v18
                    }
                    v2 += (v2 * v2)
                }
                v19 := v1[(fi15 / ((42) | 1))]
                _ = v19
            } else {
                v20 := GenS4{F0: 598, F1: 603, F2: true}
                _ = v20
            }
            _ = fi15
        }
        v21 := GenS5{F0: 1, F1: ("" + (((("a" + "\n\t") + "") + "hello") + ("" + ("a" + "hello"))))}
        _ = fi14
        _ = v21
    }
    var v22 int8 = -1
    _ = len(v1)
    v2 = 1
    for fi23 := 0; fi23 < 20; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            v22 = -1
            _ = fi24
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        if true {
            var v26 uint = 2
            _ = v26
        }
        if true {
            if (0.0 > 0.1) {
                v27 := GenS5{F0: -261, F1: ""}
                fmt.Println(v2, v22, v22)
                v4 += v2
                _ = v27
            }
            v2 += (v2 * (v4 & (v4 / (((v4 % ((v4) | 1))) | 1))))
            v3 = append(v3, 84)
        }
        v28 := v1[fi25]
        _ = fi25
        _ = v28
    }
    if false {
        if (2 <= 906) {
            switch (("a" + "foo bar") + "hello") {
            case "\n\t":
                var v29 uint16 = 42
                _ = v29
            default:
                delete(v1, -9223372036854775808)
            }
            v30 := GenS6{F0: 2, F1: (v4 | (v2 + 128)), F2: (v4 >> ((359) & 63))}
            var v31 int8 = v22
            _ = v30
            _ = v31
        } else {
            var v32 float32 = 3.14159
            v4 = v4
            _ = v32
        }
        if true {
            var v33 float32 = ((-2.25 / (0.1 - 1.5)) + 0.0)
            v34 := v1[128]
            var v35 int8 = (v22 | (v22 | (v22 >> ((4294967295) & 63))))
            _ = v33
            _ = v34
            _ = v35
        } else {
            if false {
                if false {
                    var v36 uint = 255
                    for fi37 := 0; fi37 < 5; fi37++ {
                        var v38 uint8 = 2
                        _ = fi37
                        _ = v38
                    }
                    fmt.Println(v2)
                    _ = v36
                }
            } else {
                v22 = (v22 | 100)
                v4 = 1
            }
            switch "foo bar" {
            case "foo bar":
                v39 := GenS4{F0: 2, F1: 255, F2: (!false)}
                _ = v39
            }
            var v40 string = ""
            _ = v40
        }
        v2 = (v4 | (v4 + (v2 | 131)))
    } else {
        v41 := GenS4{F0: 0, F1: 678, F2: (!true)}
        for fi42 := 0; fi42 < 3; fi42++ {
            var v43 uint = 808
            var v44 uint = 4294967295
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
