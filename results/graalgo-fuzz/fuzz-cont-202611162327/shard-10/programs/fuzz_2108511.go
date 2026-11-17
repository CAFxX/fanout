package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS3 struct { F0 int; F1 uint; F2 bool }
type GenS4 struct { F0 uint64 }
func genFn0(p0 string, p1 uint8, p2 uint) int {
    fmt.Println(p0)
    var v6 uint = (p2 / (((p2 / ((p2) | 1))) | 1))
    var v7 int = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 uint8) int16 {
    if true {
        var v8 bool = true
        if (("" != (("hello" + "a") + "foo bar")) && (73 != 2)) {
            var v9 bool = v8
            _ = v9
        } else {
            v10 := GenS1{F0: p0}
            _ = v10
        }
        v11 := GenS1{F0: (^(p0))}
        _ = v11
        _ = v8
    }
    v12 := GenS2{F0: 2, F1: p0, F2: -512}
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 int16 = -32768
            _ = fi14
            _ = v15
        }
        switch (fi13 << ((504) & 63)) {
        case 4294967295:
            p0 = (p0 & (p0 ^ p0))
        }
        _ = fi13
    }
    if true {
        var v16 float32 = 1.5
        v17 := GenS2{F0: 0, F1: 128, F2: -2147483648}
        _ = v16
        _ = v17
    }
    var v18 uint16 = 32767
    _ = p0
    _ = v12
    _ = v18
    return 32767
}

func main() {
    v1 := []string{}
    var v2 float64 = -2.25
    var v3 int16 = 24
    v4 := make(map[int]uint32)
    var v5 int = 1
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 uint16 = 0
        var v21 int16 = (v3 | 311)
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v2)
    v22 := GenS2{F0: 2, F1: 2, F2: 2147483647}
    if true {
        if false {
            var v23 bool = false
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            _ = v23
        } else {
            v5 = v5
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            v24 := GenS3{F0: (v5 + 0), F1: 713, F2: false}
            _ = v24
        }
        v5 = v5
    }
    fmt.Println(v3)
    v3 = (v3 ^ v3)
    v25 := GenS4{F0: 876}
    v5 -= (v5 * (v5 / (((v5 >> ((255) & 63))) | 1)))
    fmt.Println(v3, v5)
    v5 -= (^(v5))
    v26 := GenS2{F0: 2, F1: 127, F2: -1}
    if ((0.1 / (3.14159 - 0.1)) < 3.14159) {
        if (!true) {
            var v27 uint8 = 1
            _ = v27
        } else {
            if true {
                if true {
                    v5 -= (^(v5))
                    v28 := v4[v5]
                    _ = v28
                }
                v29 := GenS1{F0: 2}
                _ = v29
            } else {
                switch "hello" {
                case "\n\t":
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                    case "":
                        switch "" {
                        case "":
                            if false {
                                if false {
                                    var v30 int32 = 1
                                    _ = v30
                                } else {
                                    v2 -= (1.5 + 0.1)
                                    for fi31 := 0; fi31 < 11; fi31++ {
                                        for fi32 := 0; fi32 < 7; fi32++ {
                                            switch (v5 << ((4294967295) & 63)) {
                                            case -1:
                                                var v33 int = 1
                                                _ = v33
                                            default:
                                                v5 = 898
                                            }
                                            _ = fi32
                                        }
                                        _ = fi31
                                    }
                                }
                                v5 += (v5 >> ((4294967295) & 63))
                                v1 = append(v1, "a")
                            }
                        case "a":
                            v34 := v4[(^(v5))]
                            _ = v34
                        case "hello":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case "a":
                        v35 := GenS1{F0: 255}
                        _ = v35
                    }
                case "foo bar":
                    switch "hello" {
                    case "":
                        switch (v5 - genFn0("a", 255, 42)) {
                        case 0:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case "foo bar":
                        if (9223372036854775807 != 997) {
                            v3 = 380
                            var v36 uint8 = 127
                            var v37 bool = true
                            _ = v36
                            _ = v37
                        } else {
                            fmt.Println(v3, v5, v5)
                        }
                    }
                case "hello":
                    if ((255 >= 2) || (545 < 42)) {
                        v3 += v3
                    }
                default:
                    v38 := GenS3{F0: genFn0("hello", 42, 950), F1: 387, F2: true}
                    _ = v38
                }
                for fi39 := 0; fi39 < 10; fi39++ {
                    var v40 uint8 = 1
                    v3 = v3
                    _ = fi39
                    _ = v40
                }
                switch (v5 * (v5 / (((v5 & (-(v5)))) | 1))) {
                case 856:
                    for fi41 := 0; fi41 < 6; fi41++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v42 int16 = 42
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi41
                        _ = v42
                    }
                case 75:
                    v43 := GenS1{F0: 0}
                    _ = v43
                }
            }
        }
    }
    if true {
        v44 := GenS3{F0: int(v3), F1: 60, F2: true}
        v2 = 0.1
        _ = v44
    } else {
        v3 -= (v3 >> ((255) & 63))
        v45 := GenS3{F0: (v5 ^ v5), F1: 18446744073709551615, F2: (1 <= -128)}
        _ = v45
    }
    if true {
        for fi46 := 0; fi46 < 13; fi46++ {
            for fi47 := 0; fi47 < 20; fi47++ {
                var v48 string = ""
                _ = len(v4)
                v48 = "hello"
                _ = fi47
                _ = v48
            }
            fmt.Println(v5, v3, v5)
            _ = fi46
        }
        if (107 != 310) {
            fmt.Println(v5, v5)
            _ = len(v4)
            switch v5 {
            case -1:
                var v49 bool = (32768 < (^(v5)))
                _ = v49
            }
        } else {
            v2 = float64((0.1 / 3.14159))
            switch "foo bar" {
            case "\n\t":
                for fi50 := 0; fi50 < 5; fi50++ {
                    for fi51 := 0; fi51 < 12; fi51++ {
                        fmt.Println(fi51, v3)
                        v3 = v3
                        _ = fi51
                    }
                    switch genFn0((("hello" + "foo bar") + "hello"), 0, 255) {
                    case 962:
                        switch (v5 / (((-(fi50))) | 1)) {
                        case -804:
                            var v52 int8 = 127
                            _ = v52
                        case 127:
                            v2 = float64((-2.25 / 1e10))
                        default:
                            v5 += 2147483647
                        }
                    case 4294967295:
                        switch "a" {
                        case "hello":
                            v3 = v3
                        case "foo bar":
                            if (!true) {
                                var v53 float64 = 0.0
                                if true {
                                    v5 = (v5 % (((-(fi50))) | 1))
                                    switch genFn0((("" + "hello") + "hello"), 42, 409) {
                                    case 42:
                                        v3 += genFn1(2)
                                    default:
                                        if true {
                                            fmt.Println(v53, v3)
                                            v5 += (v5 | 676)
                                        } else {
                                            if true {
                                                if len(v1) > 0 {
                                                    v1[0] = "hello"
                                                }
                                            } else {
                                                v53 = v2
                                            }
                                            if false {
                                                v53 = v2
                                                v3 -= genFn1(207)
                                            }
                                        }
                                    }
                                }
                                _ = v53
                            } else {
                                if ((4294967295 > 1) && (!((!true) && (true || true)))) {
                                    v54 := GenS2{F0: 1, F1: 58, F2: -845}
                                    delete(v4, (fi50 + v5))
                                    for fi55 := 0; fi55 < 5; fi55++ {
                                        if len(v1) > 0 {
                                            v1[0] = ("foo bar" + "")
                                        }
                                        var v56 int64 = 44
                                        v3 = v3
                                        _ = fi55
                                        _ = v56
                                    }
                                    _ = v54
                                }
                                for fi57 := 0; fi57 < 4; fi57++ {
                                    v3 = (^(v3))
                                    if ((false || (9223372036854775807 != 40)) && false) {
                                        var v58 uint8 = 36
                                        var v59 int32 = 42
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = fi57
                                }
                            }
                        case "\n\t":
                            v60 := GenS2{F0: 42, F1: 42, F2: 2147483647}
                            _ = v60
                        }
                    case 42:
                        v2 = v2
                    }
                    fmt.Println(fi50)
                    _ = fi50
                }
            }
        }
    }
    if true {
        v61 := GenS2{F0: 42, F1: 1, F2: 42}
        var v62 int = (v5 % (((v5 / (((v5 / ((v5) | 1))) | 1))) | 1))
        _ = v61
        _ = v62
    } else {
        var v63 float64 = float64((((-2.25 - float32(v2)) / 3.14159) - ((1e10 / 1.0) / (1.5 + (3.14159 + float32((float64(((((1e10 / ((float32((3.14159 + (float64(1.5) * -2.25))) - 1.5) / float32(v2))) + 1e10) + 1e10) - ((3.14159 - 1e10) - float32((0.1 * 1e10))))) + float64(-2.25))))))))
        v63 += 0.1
        var v64 int16 = (v3 & v3)
        _ = v63
        _ = v64
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
