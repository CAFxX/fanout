package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = (float32(3.14159) / float32(3.14159))
    var v3 float64 = -2.25
    v4 := []int{-635, 42, 88}
    switch ("\n\t" + "foo bar") {
    case "a":
        v2 += v2
    case "foo bar":
        fmt.Println(v2, v2)
    case "\n\t":
        v5 := GenS1{F0: 32767}
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int = -144
        var v8 uint8 = 2
        switch fi6 {
        case 4294967295:
            fmt.Println(v2, v3, v7)
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            if true {
                var v11 uint16 = 66
                if false {
                    v12 := GenS1{F0: (v11 << ((45) & 63))}
                    _ = v12
                } else {
                    v3 = v3
                    v2 = (v2 * v2)
                }
                switch "" {
                case "hello":
                    v3 += 3.14159
                default:
                    delete(v1, 714)
                }
                _ = v11
            }
            _ = fi10
        }
        _ = fi9
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        if true {
            v3 = v3
        } else {
            v14 := GenS1{F0: 65535}
            v15 := GenS1{F0: 65535}
            var v16 uint32 = 1
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    fmt.Println(v3)
    v2 = v2
    v4 = append(v4, -881)
    for fi17 := 0; fi17 < 20; fi17++ {
        if false {
            if false {
                var v18 int64 = -266
                _ = v18
            }
            if ((!((-1 < -1) || ((2 >= 286) && false))) || false) {
                switch (("" + ("hello" + "\n\t")) + (("hello" + ("a" + ("" + "hello"))) + "a")) {
                case "\n\t":
                    fmt.Println(v2)
                default:
                    if (!(!(2147483648 >= 210))) {
                        v2 -= v2
                        v19 := GenS1{F0: 255}
                        switch ("hello" + ("a" + "hello")) {
                        case "foo bar":
                            v20 := GenS1{F0: 32767}
                            _ = v20
                        }
                        _ = v19
                    }
                }
            }
            fmt.Println(fi17)
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            _ = len(v4)
            v22 := GenS1{F0: 255}
            var v23 float64 = (float64(v2) * ((v3 + -2.25) + v3))
            _ = fi21
            _ = v22
            _ = v23
        }
        v2 = (float32(0.1) - -2.25)
        _ = fi17
    }
    v24 := GenS2{F0: -1, F1: (!(255 != 42)), F2: 0}
    if true {
        var v25 uint32 = 0
        v25 = 211
        v25 = v25
        _ = v25
    }
    v2 += (1e10 - (0.1 + (-2.25 * 3.14159)))
    if (32767 > 989) {
        for fi26 := 0; fi26 < 6; fi26++ {
            if (!(2147483647 <= 42)) {
                for fi27 := 0; fi27 < 5; fi27++ {
                    var v28 uint = 4294967295
                    switch (fi27 | 42) {
                    case 4294967295:
                        v29 := GenS3{F0: 859}
                        _ = v29
                    default:
                        v2 = 0.0
                    }
                    _ = fi27
                    _ = v28
                }
                fmt.Println(v3, v2, fi26)
                switch -244 {
                case 127:
                    switch ((((("foo bar" + "foo bar") + "a") + ("foo bar" + "a")) + ("\n\t" + ("" + ""))) + ("a" + ("" + ("hello" + "a")))) {
                    case "\n\t":
                        var v30 int16 = 1
                        _ = v30
                    }
                }
            } else {
                v2 = (-2.25 + v2)
                v31 := GenS2{F0: -128, F1: (153 == 57), F2: 2}
                if ((372 >= -309) && (819 > 4294967295)) {
                    if true {
                        v2 = 3.14159
                    } else {
                        for fi32 := 0; fi32 < 13; fi32++ {
                            _ = len(v1)
                            var v33 uint = 515
                            _ = fi32
                            _ = v33
                        }
                    }
                } else {
                    for fi34 := 0; fi34 < 8; fi34++ {
                        var v35 float32 = v2
                        if (!(2 != 2147483648)) {
                            fmt.Println(v2, fi26, v3)
                            var v36 uint = 0
                            v4 = append(v4, fi34)
                            _ = v36
                        } else {
                            v37 := v1[4294967295]
                            var v38 float32 = 3.14159
                            v39 := GenS3{F0: 127}
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        _ = fi34
                        _ = v35
                    }
                    var v40 int16 = 1
                    _ = v40
                }
                _ = v31
            }
            switch 32767 {
            case 2147483647:
                if true {
                    if false {
                        if (-1 == 32767) {
                            _ = len(v4)
                            _ = len(v4)
                            var v41 uint64 = 411
                            _ = v41
                        }
                    }
                    v42 := GenS4{F0: 65535}
                    _ = v42
                }
            case -86:
                v43 := GenS2{F0: -89, F1: true, F2: 42}
                _ = v43
            case -2147483648:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                v2 -= float32(float64(v2))
            }
            for fi44 := 0; fi44 < 11; fi44++ {
                v2 = ((3.14159 + 0.0) / v2)
                var v45 int = fi44
                if len(v4) > 0 {
                    v4[0] = fi26
                }
                v3 -= ((((1e10 + ((((1.5 - (v3 * (-2.25 + v3))) + float64(float32(v3))) * 1e10) - v3)) / float64(float32(v3))) / (((0.1 * 1.5) + 1e10) / v3)) * v3)
                _ = fi44
                _ = v45
            }
            _ = fi26
        }
    } else {
        v46 := GenS5{F0: 42, F1: 1, F2: 127}
        v3 += v3
        _ = v46
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        v48 := GenS6{F0: 32767}
        _ = len(v4)
        delete(v1, 527)
        _ = fi47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
