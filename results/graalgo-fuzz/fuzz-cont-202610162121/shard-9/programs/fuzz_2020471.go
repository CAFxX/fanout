package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 int64; F2 int }
type GenS5 struct { F0 uint8; F1 bool; F2 string }
func genFn0(p0 int32, p1 int8) uint8 {
    var v7 int64 = 154
    if true {
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = v7
    return 100
}
func genFn1() uint64 {
    v8 := GenS1{F0: false, F1: 931}
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    var v10 uint32 = 340
    _ = v10
    _ = v8
    return 204
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 0
    var v3 uint64 = 45
    var v4 int64 = 1
    v5 := []string{"foo bar", ("hello" + ((("a" + "a") + ("\n\t" + "hello")) + ("hello" + ("\n\t" + "a")))), "\n\t"}
    var v6 uint8 = 0
    switch 128 {
    case -727:
        var v11 int = -2147483648
        _ = v11
    case 228:
        if false {
            if ((v6 == (v6 | v6)) || ((v3 >> ((0) & 63)) <= (v3 | (^(v3))))) {
                _ = len(v5)
            } else {
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 string = ("a" + "\n\t")
                    v14 := GenS1{F0: (true && (1.5 <= 1e10)), F1: 2}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 14; fi15++ {
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v17 := GenS1{F0: ((true && true) && (true && false)), F1: uint32(v6)}
                        if false {
                            v6 = v6
                            if len(v5) > 0 {
                                v5[0] = ("\n\t" + ("a" + "\n\t"))
                            }
                        }
                        fmt.Println(fi15)
                        _ = fi16
                        _ = v17
                    }
                    var v18 uint = 4294967295
                    _ = fi15
                    _ = v18
                }
            }
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 uint64 = (^(v3))
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(fi20, fi19)
                    case "\n\t":
                        var v22 float32 = (3.14159 + -2.25)
                        _ = v22
                    default:
                        switch fi20 {
                        case 32768:
                            if true {
                                v21 += genFn1()
                            } else {
                                var v23 int16 = 1
                                switch (((("hello" + "a") + "a") + ((("hello" + "a") + ("a" + ("foo bar" + "a"))) + ("" + ""))) + "") {
                                case "hello":
                                    var v24 int32 = -526
                                    _ = v24
                                case "foo bar":
                                    var v25 bool = false
                                    _ = v25
                                case "":
                                    if false {
                                        var v26 float64 = -2.25
                                        if false {
                                            fmt.Println(v2)
                                            v27 := GenS2{F0: 244}
                                            v6 += v6
                                            if (0 > 457) {
                                                v1[fi20] = "foo bar"
                                                var v28 int32 = -1
                                                var v29 uint = 1
                                                _ = v28
                                                _ = v29
                                            } else {
                                                v6 += (v6 & (v6 - 42))
                                                delete(v1, fi19)
                                                fmt.Println(v4, v6, v4)
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            _ = v27
                                        }
                                        _ = len(v5)
                                        _ = v26
                                    }
                                }
                                var v30 uint16 = v2
                                _ = len(v1)
                                v21 -= genFn1()
                                if ((^(v6)) < v6) {
                                    v31 := GenS3{F0: 2147483647}
                                    _ = v31
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = "\n\t"
                                    }
                                    var v32 int32 = int32(v4)
                                    switch "hello" {
                                    case "foo bar":
                                        v3 = v21
                                    default:
                                        if ((float64(-2.25) - (1e10 - ((-2.25 - (1.5 - 1e10)) / (3.14159 * ((-2.25 + 0.0) / (1e10 / -2.25)))))) > (((((3.14159 * (0.0 / float64(3.14159))) - (1e10 / ((1e10 + 0.0) - (3.14159 * ((((float64(0.0) - 1e10) / ((((float64((1e10 / (1e10 / 0.1))) / 3.14159) - -2.25) - float64(((1.5 * -2.25) - 3.14159))) / 1.5)) + 0.0) / float64(3.14159)))))) / float64(0.1)) * 0.0) * ((1.5 / (1.5 / float64(-2.25))) + 1.5))) {
                                            var v33 int32 = v32
                                            fmt.Println(v21)
                                            _ = v33
                                        } else {
                                            var v34 uint = 255
                                            _ = v34
                                        }
                                    }
                                    _ = v32
                                }
                                if len(v5) > 0 {
                                    v5[0] = "\n\t"
                                }
                                fmt.Println(fi20)
                                var v35 int8 = 100
                                _ = v23
                                _ = v30
                                _ = v35
                            }
                        case 4294967295:
                            if (true || true) {
                                v1[fi19] = "a"
                            }
                        case -886:
                            var v36 int32 = -1
                            _ = v36
                        default:
                            var v37 uint16 = 65535
                            _ = v37
                        }
                        v3 = v21
                        if (false && (true && true)) {
                            fmt.Println(v21)
                            v38 := GenS3{F0: -17}
                            v5 = append(v5, ("hello" + "hello"))
                            _ = v38
                        } else {
                            var v39 int = fi20
                            v4 -= v4
                            _ = v39
                        }
                    }
                    delete(v1, fi19)
                    _ = fi20
                    _ = v21
                }
                _ = fi19
            }
        }
    default:
        fmt.Println(v3)
    }
    if len(v5) > 0 {
        v5[0] = ("a" + ((("hello" + "") + "hello") + ("a" + "hello")))
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        for fi41 := 0; fi41 < 14; fi41++ {
            v3 += v3
            if (-562 <= int16(v6)) {
                if false {
                    v42 := GenS2{F0: -2147483648}
                    _ = v42
                } else {
                    v43 := GenS4{F0: fi41, F1: 719, F2: 32768}
                    _ = v43
                }
                v44 := GenS5{F0: 90, F1: ((v4 >> ((4294967295) & 63)) >= v4), F2: ("" + ("hello" + "foo bar"))}
                _ = v44
            }
            _ = fi41
        }
        v3 += (v3 & (^(v3)))
        _ = fi40
    }
    v2 = (v2 * 682)
    fmt.Println(v2)
    v6 -= (^(v6))
    switch 32768 {
    case 4294967295:
        _ = len(v1)
    }
    if false {
        var v45 bool = (-32768 < -108)
        var v46 string = "hello"
        _ = v45
        _ = v46
    } else {
        switch (("" + ((("hello" + (((("\n\t" + "hello") + "\n\t") + "hello") + "\n\t")) + (("foo bar" + "hello") + "a")) + "a")) + ("a" + ("" + ""))) {
        case "a":
            v6 += 1
        case "foo bar":
            if (!(!false)) {
                var v47 float64 = 0.0
                switch 4294967295 {
                case 821:
                    for fi48 := 0; fi48 < 1; fi48++ {
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                        _ = fi48
                    }
                case 32768:
                    v4 -= (v4 & v4)
                case -2147483648:
                    var v49 string = "\n\t"
                    _ = v49
                }
                var v50 uint64 = (v3 - v3)
                _ = v47
                _ = v50
            }
        case "\n\t":
            v6 = 1
        default:
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
        }
        v4 = 9223372036854775807
        var v51 float32 = (0.0 - (3.14159 - (0.1 / ((3.14159 + 0.0) + 1.5))))
        _ = v51
    }
    v2 -= v2
    for fi52 := 0; fi52 < 12; fi52++ {
        var v53 uint64 = v3
        v3 += (^(v3))
        delete(v1, fi52)
        _ = fi52
        _ = v53
    }
    v54 := GenS5{F0: v6, F1: true, F2: ((("hello" + "\n\t") + ("" + "")) + (("foo bar" + "a") + "hello"))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
