package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int8; F1 int; F2 int }
type GenS7 struct { F0 uint }
func genFn0() uint64 {
    if true {
        var v4 uint = 4294967295
        v4 = v4
        v5 := GenS1{F0: v4, F1: "a"}
        _ = v4
        _ = v5
    }
    var v6 int16 = 0
    var v7 string = "a"
    var v8 string = (((((v7 + ("a" + ("hello" + (("hello" + v7) + "a")))) + v7) + (v7 + (((("" + v7) + v7) + v7) + ("foo bar" + v7)))) + (v7 + v7)) + (v7 + "a"))
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{2, 1, 2}
    var v3 uint64 = 199
    if (-32768 <= -1) {
        if (42 > 1) {
            var v9 uint = 65535
            v3 += v3
            _ = v9
        } else {
            v3 = (v3 - 2)
            if (-9223372036854775808 > 925) {
                for fi10 := 0; fi10 < 13; fi10++ {
                    v11 := GenS1{F0: 1, F1: ("foo bar" + "a")}
                    if true {
                        for fi12 := 0; fi12 < 7; fi12++ {
                            if true {
                                v13 := v1[(fi12 >> ((1) & 63))]
                                v14 := GenS2{F0: 65535, F1: 918, F2: 795}
                                var v15 uint64 = v3
                                v2 = append(v2, 2147483648)
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                            v3 = (v3 ^ (v3 - (v3 * genFn0())))
                            _ = fi12
                        }
                        if false {
                            if (false || (("hello" + "") <= "a")) {
                                if len(v2) > 0 {
                                    v2[0] = 2147483647
                                }
                            } else {
                                v2 = append(v2, 2)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                    }
                    v16 := GenS2{F0: 29, F1: v3, F2: 32767}
                    _ = fi10
                    _ = v11
                    _ = v16
                }
            }
            if (true && true) {
                if (3.14159 < 0.0) {
                    var v17 uint8 = 2
                    var v18 int32 = 1
                    _ = v17
                    _ = v18
                }
                v3 -= genFn0()
                var v19 uint = 1
                _ = v19
            }
        }
    }
    fmt.Println(v3)
    v20 := v1[1]
    var v21 float64 = float64(-2.25)
    switch (("\n\t" + "a") + "\n\t") {
    case "a":
        var v22 float64 = float64(3.14159)
        _ = v22
    case "hello":
        fmt.Println(v21)
    case "\n\t":
        v23 := GenS3{F0: 2, F1: 1000, F2: 969}
        _ = v23
    default:
        var v24 uint = 255
        _ = v24
    }
    if (1 != 1) {
        v25 := GenS4{F0: uint8(v3)}
        switch 128 {
        case -9223372036854775808:
            if (true && (!(549 > 42))) {
                v20 = (v20 ^ 2)
            }
        default:
            if false {
                switch ("foo bar" + ("a" + ("foo bar" + ("\n\t" + "a")))) {
                case "foo bar":
                    v2 = append(v2, v20)
                default:
                    switch -1 {
                    case 9223372036854775807:
                        for fi26 := 0; fi26 < 1; fi26++ {
                            fmt.Println(v3, v21)
                            for fi27 := 0; fi27 < 9; fi27++ {
                                fmt.Println(v20, fi26, fi27)
                                var v28 uint = 67
                                _ = fi27
                                _ = v28
                            }
                            _ = fi26
                        }
                    }
                }
                var v29 string = ""
                fmt.Println(v21, v21, v20)
                _ = v29
            }
        }
        _ = v25
    } else {
        fmt.Println(v3, v21, v20)
        var v30 uint8 = 128
        var v31 int16 = -623
        _ = v30
        _ = v31
    }
    v20 -= 2147483647
    var v32 string = (("hello" + "\n\t") + "a")
    switch v32 {
    case "\n\t":
        v32 = ("foo bar" + "hello")
    default:
        _ = len(v2)
    }
    delete(v1, -143)
    v20 -= v20
    fmt.Println(v20)
    v21 = float64(1.5)
    if true {
        fmt.Println(v3, v21)
        if (false || true) {
            for fi33 := 0; fi33 < 6; fi33++ {
                v3 += (v3 + (v3 | (v3 / ((0) | 1))))
                v21 += (v21 / v21)
                v34 := GenS5{F0: 1, F1: 310}
                _ = fi33
                _ = v34
            }
            switch 128 {
            case -2147483648:
                var v35 float32 = 3.14159
                _ = v35
            case 4294967295:
                var v36 uint = 255
                _ = v36
            default:
                _ = len(v2)
            }
        }
        switch 42 {
        case 1:
            var v37 string = v32
            _ = v37
        case 42:
            v38 := GenS3{F0: 2, F1: 0, F2: 647}
            _ = v38
        case 241:
            for fi39 := 0; fi39 < 13; fi39++ {
                var v40 int8 = 99
                for fi41 := 0; fi41 < 13; fi41++ {
                    fmt.Println(v3)
                    v21 += (v21 / float64((-2.25 + 1.5)))
                    fmt.Println(v20)
                    var v42 uint64 = v3
                    v43 := GenS6{F0: v40, F1: fi39, F2: (fi41 & fi39)}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = fi39
                _ = v40
            }
        default:
            v44 := GenS6{F0: -99, F1: 127, F2: -1}
            _ = v44
        }
    } else {
        fmt.Println(v32)
        if ((32767 < -38) || (0 == 584)) {
            for fi45 := 0; fi45 < 2; fi45++ {
                v46 := GenS7{F0: 65535}
                _ = fi45
                _ = v46
            }
            v21 = v21
            if len(v2) > 0 {
                v2[0] = uint32(v3)
            }
        } else {
            var v47 int8 = -114
            v48 := v1[-860]
            var v49 string = ""
            _ = v47
            _ = v48
            _ = v49
        }
    }
    fmt.Println(v32, v20)
    var v50 uint = 140
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v50
}
