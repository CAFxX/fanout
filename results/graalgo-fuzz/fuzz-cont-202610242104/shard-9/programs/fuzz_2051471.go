package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS9 struct { F0 int; F1 int8; F2 int }

func main() {
    var v1 float32 = 0.1
    v2 := []string{"\n\t", "a"}
    v3 := make(map[int]uint32)
    var v4 int64 = -113
    var v5 int32 = -2147483648
    v6 := []uint32{2, 986, 665}
    switch -530 {
    case -9223372036854775808:
        v7 := GenS1{F0: 362, F1: 128}
        _ = v7
    case 42:
        var v8 uint64 = 354
        _ = v8
    default:
        v4 = v4
    }
    fmt.Println(v4)
    var v9 bool = true
    if v9 {
        delete(v3, 400)
        v10 := GenS2{F0: v9, F1: 116, F2: 321}
        _ = v10
    } else {
        if ("" >= ("\n\t" + "")) {
            var v11 uint16 = 32767
            v4 = (^(v4))
            _ = v11
        }
    }
    v5 = v5
    v5 -= v5
    fmt.Println(v9, v1, v4)
    for fi12 := 0; fi12 < 19; fi12++ {
        switch (("a" + "") + ("hello" + "a")) {
        case "a":
            if v9 {
                v13 := GenS3{F0: -32, F1: 42}
                if v9 {
                    if (0 == 300) {
                        var v14 int8 = -128
                        if v9 {
                            v15 := GenS4{F0: (fi12 % ((fi12) | 1))}
                            v16 := GenS4{F0: fi12}
                            var v17 int64 = v4
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        v18 := v3[(fi12 >> ((2) & 63))]
                        _ = v14
                        _ = v18
                    }
                    v19 := v3[(-(fi12))]
                    v19 = (v19 | (v19 << ((4294967295) & 63)))
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        delete(v3, (fi12 ^ (fi12 ^ fi12)))
                        _ = fi20
                    }
                }
                _ = v13
            }
        case "hello":
            for fi21 := 0; fi21 < 16; fi21++ {
                v22 := GenS5{F0: -128}
                switch (fi12 ^ (fi21 | (fi21 % ((fi21) | 1)))) {
                case -9223372036854775808:
                    v5 += (v5 << ((42) & 63))
                }
                _ = fi21
                _ = v22
            }
        case "":
            for fi23 := 0; fi23 < 18; fi23++ {
                v4 = (v4 ^ 730)
                v5 = v5
                v24 := GenS6{F0: 183}
                _ = fi23
                _ = v24
            }
        }
        if v9 {
            for fi25 := 0; fi25 < 3; fi25++ {
                fmt.Println(v9)
                if ((v9 && v9) || false) {
                    var v26 uint = 42
                    fmt.Println(v1, fi25)
                    v27 := GenS2{F0: (!(2 <= 65535)), F1: 52, F2: 255}
                    var v28 int32 = (v5 / ((v5) | 1))
                    var v29 uint = 255
                    switch "" {
                    case "foo bar":
                        var v30 int16 = -32768
                        _ = v30
                    case "\n\t":
                        fmt.Println(v29, v26)
                    case "a":
                        var v31 float32 = v1
                        _ = v31
                    default:
                        if false {
                            v29 += (v26 % ((42) | 1))
                        } else {
                            var v32 float32 = ((3.14159 * (((v1 + 0.0) - float32(0.0)) + (v1 / v1))) / float32(float64(((v1 + v1) * 3.14159))))
                            switch 32767 {
                            case -422:
                                v28 = (v28 + -1)
                            case 4294967295:
                                var v33 uint8 = 255
                                _ = v33
                            case 0:
                                var v34 uint32 = 1
                                _ = v34
                            }
                            v5 += v5
                            _ = v32
                        }
                    }
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                v35 := GenS7{F0: 1000, F1: 140}
                _ = fi25
                _ = v35
            }
            v9 = false
        } else {
            v5 += 452
            var v36 float64 = 0.1
            v9 = (!((!true) && false))
            _ = v36
        }
        if (!false) {
            switch "\n\t" {
            case "":
                v9 = (int(v5) != (fi12 + (fi12 << ((42) & 63))))
            }
        } else {
            v37 := GenS8{F0: 735, F1: 382, F2: 0}
            if false {
                for fi38 := 0; fi38 < 9; fi38++ {
                    v39 := GenS9{F0: 1, F1: 1, F2: (^(fi12))}
                    v40 := v3[(fi38 / (((fi38 | (fi38 / (((fi38 >> ((255) & 63))) | 1)))) | 1))]
                    if v9 {
                        v9 = v9
                    } else {
                        if len(v6) > 0 {
                            v6[0] = (v40 & (v40 % (((^(v40))) | 1)))
                        }
                        v41 := GenS6{F0: 847}
                        _ = v41
                    }
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            } else {
                v5 += -2147483648
            }
            _ = v37
        }
        _ = fi12
    }
    var v42 uint8 = 2
    var v43 uint8 = (^(v42))
    if v9 {
        switch "" {
        case "a":
            delete(v3, -753)
        default:
            v43 -= (^(v42))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v9
}
