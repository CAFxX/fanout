package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint16; F1 int; F2 uint }
type GenS5 struct { F0 bool; F1 uint; F2 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint16, p2 string) uint16 {
    if false {
        switch p2 {
        case "\n\t":
            v5 := GenS1{F0: p2, F1: 1, F2: -615}
            _ = v5
        case "hello":
            var v6 uint16 = p1
            _ = v6
        default:
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 uint64 = 0
                _ = fi7
                _ = v8
            }
        }
        v9 := GenS1{F0: (((p2 + p2) + p2) + p2), F1: 4294967295, F2: 2147483647}
        var v10 uint32 = 4294967295
        _ = v10
        _ = v9
    } else {
        fmt.Println(p0, p2, p1)
        var v11 uint = 255
        var v12 int64 = 502
        _ = v11
        _ = v12
    }
    switch "foo bar" {
    case "a":
        if false {
            if true {
                fmt.Println(p0, p0)
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 ^ (p1 + 137))
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    v4 := []string{}
    var v13 uint16 = v1
    fmt.Println(v1, v13)
    fmt.Println(v1)
    if (!true) {
        for fi14 := 0; fi14 < 12; fi14++ {
            v4 = append(v4, (("" + "") + ""))
            if true {
                fmt.Println(v1, v13, v13)
            }
            switch ((("foo bar" + "foo bar") + ("a" + "a")) + ("" + "a")) {
            case "foo bar":
                var v15 uint8 = 241
                _ = v15
            case "\n\t":
                fmt.Println(fi14, v1)
            case "":
                fmt.Println(v1)
            default:
                fmt.Println(fi14, v1, v1)
            }
            _ = fi14
        }
    }
    if (!(("\n\t" + "foo bar") <= "\n\t")) {
        var v16 uint32 = 65535
        _ = v16
    }
    if false {
        var v17 float64 = 3.14159
        v17 = 1e10
        _ = v17
    }
    v18 := GenS2{F0: 65535, F1: 278, F2: 42}
    for fi19 := 0; fi19 < 1; fi19++ {
        if (32767 < 1000) {
            switch (fi19 >> ((213) & 63)) {
            case 1:
                var v20 int8 = -128
                _ = v20
            case 750:
                var v21 uint = 42
                _ = v21
            case 0:
                fmt.Println(v13)
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                v23 := GenS2{F0: 4294967295, F1: (v13 ^ 225), F2: 185}
                _ = fi22
                _ = v23
            }
        }
        v13 += v13
        _ = fi19
    }
    switch 32767 {
    case 42:
        fmt.Println(v1, v1)
    default:
        var v24 int8 = 1
        _ = v24
    }
    switch "\n\t" {
    case "hello":
        v25 := GenS2{F0: 2147483648, F1: (v1 + (v1 | (^(v13)))), F2: 190}
        _ = v25
    default:
        for fi26 := 0; fi26 < 1; fi26++ {
            v27 := GenS3{F0: 1, F1: 0}
            switch "foo bar" {
            case "\n\t":
                if true {
                    for fi28 := 0; fi28 < 7; fi28++ {
                        fmt.Println(v13, fi28)
                        var v29 int = 32768
                        if ((-125 > -128) && true) {
                            v30 := GenS4{F0: 65535, F1: fi28, F2: 65535}
                            v31 := GenS5{F0: false, F1: 173, F2: 32767}
                            if (0.1 > float32((3.14159 + 3.14159))) {
                                var v32 uint32 = 4294967295
                                switch (v29 << ((41) & 63)) {
                                case 422:
                                    v32 += 4294967295
                                case 128:
                                    v33 := GenS6{F0: (!false)}
                                    _ = v33
                                default:
                                    _ = len(v3)
                                }
                                _ = v32
                            }
                            var v34 float64 = (1e10 / (float64(((1.5 / ((3.14159 - (3.14159 / 3.14159)) + float32(float64(1e10)))) / 0.1)) + -2.25))
                            v29 = -2147483648
                            v1 = (v1 - v13)
                            _ = v30
                            _ = v31
                            _ = v34
                        }
                        v13 -= (v13 % ((v1) | 1))
                        _ = fi28
                        _ = v29
                    }
                } else {
                    if false {
                        if (-39 <= 42) {
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(fi26, v13)
                }
            case "":
                if len(v4) > 0 {
                    v4[0] = ((("a" + "foo bar") + "hello") + "a")
                }
            case "hello":
                switch ("foo bar" + ("a" + "foo bar")) {
                case "":
                    switch (fi26 * fi26) {
                    case 2147483647:
                        for fi35 := 0; fi35 < 9; fi35++ {
                            v36 := GenS6{F0: true}
                            var v37 int8 = 100
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    case 1:
                        var v38 int64 = -558
                        _ = v38
                    }
                case "foo bar":
                    var v39 uint32 = 65535
                    _ = v39
                case "\n\t":
                    var v40 uint8 = 255
                    _ = v40
                }
            }
            _ = fi26
            _ = v27
        }
    }
    v13 = (^(v13))
    delete(v2, 1)
    fmt.Println(v1, v13)
    if true {
        for fi41 := 0; fi41 < 11; fi41++ {
            v42 := GenS7{F0: 1, F1: fi41}
            v1 += v1
            _ = fi41
            _ = v42
        }
        v1 = v13
    }
    if false {
        fmt.Println(v1, v1)
        for fi43 := 0; fi43 < 2; fi43++ {
            var v44 uint = 42
            _ = fi43
            _ = v44
        }
        v45 := GenS8{F0: 2147483647}
        _ = v45
    } else {
        v1 = v13
        var v46 string = ("a" + ("" + "hello"))
        v1 += 961
        _ = v46
    }
    if true {
        v1 = (v13 * 42)
    } else {
        if len(v4) > 0 {
            v4[0] = ("foo bar" + "a")
        }
        v13 += (v1 - 295)
        if true {
            var v47 uint64 = 42
            var v48 float32 = (3.14159 - (-2.25 + 0.0))
            v49 := GenS9{F0: 329}
            _ = v47
            _ = v48
            _ = v49
        } else {
            if len(v4) > 0 {
                v4[0] = ("foo bar" + "")
            }
            v13 += v1
        }
    }
    switch 9223372036854775807 {
    case -1:
        fmt.Println(v1, v13)
    case -2147483648:
        var v50 uint32 = 42
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
