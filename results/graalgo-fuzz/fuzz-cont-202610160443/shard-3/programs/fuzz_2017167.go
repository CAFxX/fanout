package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 bool }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 uint; F2 uint64 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS9 struct { F0 uint; F1 uint; F2 uint64 }
func genFn0(p0 int8, p1 int, p2 int64) uint16 {
    fmt.Println(p2, p2)
    v6 := []int{}
    var v7 int = p1
    var v8 uint64 = 933
    if (true || ((v8 <= 4294967295) && false)) {
        p1 += -493
        for fi9 := 0; fi9 < 1; fi9++ {
            fmt.Println(v8)
            var v10 float64 = float64(float32(1e10))
            var v11 uint16 = 525
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 uint8 = 42
    v2 := []int{9223372036854775807, 2147483647}
    var v3 int32 = -1
    var v4 int32 = (-(v3))
    v5 := make(map[int]uint32)
    v2 = append(v2, -9223372036854775808)
    if false {
        var v12 uint16 = 0
        _ = len(v5)
        _ = v12
    }
    v4 += v3
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    switch "" {
    case "foo bar":
        if (!true) {
            v3 -= v3
            var v13 uint64 = 4294967295
            _ = v13
        }
    }
    v14 := GenS1{F0: v1, F1: 0, F2: true}
    var v15 int32 = v3
    v16 := GenS1{F0: (v1 + (v1 << ((141) & 63))), F1: 2, F2: ("\n\t" == "a")}
    v17 := GenS1{F0: v1, F1: 2147483647, F2: ((0.0 / 1.5) < -2.25)}
    if true {
        v3 = 0
        if false {
            if (672 <= 260) {
                v3 -= (v3 << ((1) & 63))
                for fi18 := 0; fi18 < 8; fi18++ {
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v20 := GenS1{F0: v1, F1: 42, F2: ((("foo bar" + "foo bar") + "a") != (("" + ((("hello" + ("foo bar" + ("foo bar" + "hello"))) + "foo bar") + "")) + "hello"))}
                        _ = fi19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 10; fi21++ {
                        if ((^(v1)) < (v1 + v1)) {
                            _ = len(v5)
                            v4 -= (v15 % ((v15) | 1))
                            var v22 float64 = ((-2.25 - (3.14159 + -2.25)) + (-2.25 + float64(-2.25)))
                            _ = v22
                        } else {
                            v3 -= (v15 + v15)
                            switch "" {
                            case "hello":
                                if ((-1 <= 42) || true) {
                                    var v23 uint16 = 0
                                    v23 = (^(v23))
                                    fmt.Println(fi21, v23)
                                    _ = v23
                                } else {
                                    v3 += (v4 ^ (-(v3)))
                                }
                            case "foo bar":
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(v15)
                                case "foo bar":
                                    if ((!(((0.0 / (float32((0.1 * 0.1)) - float32(3.14159))) / (0.1 + (0.0 / ((0.0 / (1.5 - ((0.1 - (3.14159 * (((3.14159 - (1e10 * 0.0)) * 1e10) + float32(1e10)))) * (1e10 * 0.0)))) - -2.25)))) <= 3.14159)) && false) {
                                        v1 = v1
                                        v4 += v4
                                    }
                                default:
                                    v15 -= (v3 << ((571) & 63))
                                }
                            case "":
                                v24 := GenS1{F0: (v1 - 128), F1: 42, F2: (true && true)}
                                _ = v24
                            default:
                                switch (("hello" + ("" + ((((("foo bar" + (((("a" + "") + "foo bar") + "") + ("a" + ""))) + "foo bar") + "") + ("\n\t" + ("\n\t" + (("hello" + "foo bar") + "")))) + "foo bar"))) + ("" + "foo bar")) {
                                case "":
                                    if (!true) {
                                        v2 = append(v2, fi18)
                                    } else {
                                        v25 := GenS1{F0: v1, F1: 732, F2: (821 < 42)}
                                        v26 := GenS2{F0: v3, F1: genFn0(100, fi18, 9223372036854775807)}
                                        fmt.Println(v3, v3, v3)
                                        _ = v25
                                        _ = v26
                                    }
                                case "a":
                                    fmt.Println(v15, fi18)
                                case "hello":
                                    v27 := GenS1{F0: (v1 & v1), F1: 154, F2: false}
                                    _ = v27
                                }
                                v1 = 1
                                fmt.Println(v15, v4)
                            }
                            v28 := GenS3{F0: 71}
                            _ = v28
                        }
                        _ = fi21
                    }
                    var v29 int64 = -9223372036854775808
                    _ = fi18
                    _ = v29
                }
                v30 := GenS4{F0: 42}
                _ = v30
            }
            v31 := GenS5{F0: 18446744073709551615, F1: 947, F2: 577}
            _ = v31
        }
        if true {
            var v32 bool = true
            _ = v32
        } else {
            v5[0] = 70
        }
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        if false {
            if (true || (830 < 0)) {
                v34 := GenS6{F0: true, F1: 1000}
                _ = v34
            }
            v3 = int32(v1)
            v3 += (v3 - int32(v1))
        }
        if len(v2) > 0 {
            v2[0] = fi33
        }
        var v35 uint32 = 2147483647
        _ = fi33
        _ = v35
    }
    switch ("a" + "hello") {
    case "foo bar":
        switch "a" {
        case "hello":
            fmt.Println(v15, v4, v15)
        case "a":
            v5[42] = 65535
        }
    case "a":
        switch -2147483648 {
        case -1:
            if (v1 < v1) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v36 uint = 0
                for fi37 := 0; fi37 < 7; fi37++ {
                    switch (fi37 << ((v36) & 63)) {
                    case -357:
                        switch -770 {
                        case -1:
                            var v38 int8 = -1
                            _ = v38
                        case 918:
                            v36 -= 18446744073709551615
                        case 2147483647:
                            switch (("hello" + "hello") + "\n\t") {
                            case "foo bar":
                                v36 = (v36 + 531)
                            case "a":
                                v1 += v1
                            default:
                                fmt.Println(v4)
                            }
                        default:
                            switch fi37 {
                            case -2147483648:
                                v36 += (v36 - v36)
                            default:
                                var v39 int8 = -1
                                _ = v39
                            }
                        }
                    case 4294967295:
                        if (1 == 827) {
                            var v40 uint = (^(v36))
                            _ = len(v2)
                            _ = v40
                        }
                    case 589:
                        switch "" {
                        case "foo bar":
                            v3 -= 42
                        case "\n\t":
                            for fi41 := 0; fi41 < 14; fi41++ {
                                if (212 >= v1) {
                                    if false {
                                        var v42 int8 = -128
                                        if true {
                                            v43 := GenS7{F0: 2147483647, F1: 0}
                                            _ = v43
                                        } else {
                                            v3 += (v15 & (v4 & v15))
                                        }
                                        v4 = v15
                                        _ = v42
                                    }
                                    v44 := GenS8{F0: 0, F1: 42, F2: 0}
                                    v45 := GenS9{F0: 42, F1: (v36 + v36), F2: 4294967295}
                                    _ = v44
                                    _ = v45
                                }
                                var v46 string = ("" + "hello")
                                _ = fi41
                                _ = v46
                            }
                        case "hello":
                            v5[fi37] = 2147483647
                        default:
                            if false {
                                delete(v5, 128)
                                var v47 int16 = 1000
                                v48 := GenS8{F0: 127, F1: 0, F2: 2147483648}
                                _ = v47
                                _ = v48
                            }
                        }
                    }
                    _ = fi37
                }
                _ = v36
            }
        case 42:
            var v49 uint8 = 1
            _ = v49
        case 127:
            var v50 float32 = (1e10 + 1e10)
            _ = v50
        }
    default:
        var v51 uint = 968
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
