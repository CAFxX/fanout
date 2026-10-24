package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int64 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 string; F1 uint; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 int8; F1 int; F2 bool }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]string)
    var v3 int32 = 2147483647
    fmt.Println(v3)
    v1 += v1
    v1 += v1
    if (42 > 1) {
        v2[2147483647] = "\n\t"
    }
    v4 := GenS1{F0: 42, F1: "foo bar", F2: 1}
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            v3 = (-(v3))
            switch (-(fi5)) {
            case -44:
                var v7 string = ""
                _ = v7
            case -418:
                if true {
                    if false {
                        var v8 int32 = -55
                        v8 = 2147483647
                        _ = v8
                    }
                }
            }
            _ = fi6
        }
        _ = fi5
    }
    v9 := []int{555, 737}
    v3 -= v3
    fmt.Println(v1)
    switch ("" + ("\n\t" + "hello")) {
    case "foo bar":
        delete(v2, 742)
    }
    v10 := GenS2{F0: 1000, F1: (true || (false || true))}
    for fi11 := 0; fi11 < 19; fi11++ {
        if ((4294967295 != uint64(v1)) && false) {
            if (-2.25 <= float64(((((float32(1.5) * (0.1 - 1.5)) / 1.5) * (0.1 * 0.0)) - (float32(float64((float32(((0.0 + 0.1) + 0.0)) * ((1.5 / 3.14159) + float32(1.5))))) + (0.1 - 0.0))))) {
                for fi12 := 0; fi12 < 1; fi12++ {
                    if len(v9) > 0 {
                        v9[0] = fi11
                    }
                    if (true || (((((("hello" + "hello") + "hello") + "hello") + "\n\t") + ("foo bar" + ("foo bar" + "foo bar"))) != "hello")) {
                        v1 += (^(v1))
                        if false {
                            switch fi12 {
                            case -2147483648:
                                var v13 string = ""
                                _ = v13
                            case 423:
                                var v14 uint = v1
                                _ = v14
                            case 394:
                                _ = len(v9)
                            }
                            v15 := GenS2{F0: 1, F1: (65535 > 977)}
                            _ = v15
                        } else {
                            var v16 int = fi12
                            _ = v16
                        }
                    } else {
                        var v17 float32 = -2.25
                        var v18 int16 = 1
                        var v19 int16 = (v18 ^ (v18 << (((v1 >> ((v1) & 63))) & 63)))
                        switch "a" {
                        case "a":
                            if true {
                                switch ("\n\t" + "a") {
                                case "":
                                    fmt.Println(v17, fi11)
                                    v1 = v1
                                case "a":
                                    if (!((true && false) && (336 < (v1 ^ 2)))) {
                                        var v20 int16 = v18
                                        v17 = 0.1
                                        _ = v20
                                    } else {
                                        var v21 float32 = ((((1.5 / float32((0.0 - -2.25))) * v17) + v17) * ((v17 * v17) * v17))
                                        if true {
                                            v22 := GenS3{F0: "", F1: (v1 * (v1 * 42)), F2: -225}
                                            _ = v22
                                        }
                                        v23 := v2[606]
                                        if (!false) {
                                            var v24 uint = (v1 + v1)
                                            _ = v24
                                        } else {
                                            fmt.Println(v21)
                                            v18 -= v19
                                            var v25 uint64 = 18446744073709551615
                                            v25 = (^(v25))
                                            fmt.Println(fi12, v23, v21)
                                            var v26 uint16 = 255
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = v21
                                        _ = v23
                                    }
                                case "hello":
                                    v1 = v1
                                }
                                if true {
                                    v19 -= -1
                                    var v27 float32 = v17
                                    _ = v27
                                }
                                if true {
                                    v18 = v18
                                    v28 := GenS2{F0: -1, F1: ((fi11 >= (fi11 * fi11)) && true)}
                                    _ = v28
                                }
                                var v29 float32 = v17
                                v30 := GenS1{F0: 0, F1: ("a" + (("a" + ("" + "\n\t")) + "\n\t")), F2: -9223372036854775808}
                                _ = v29
                                _ = v30
                            } else {
                                var v31 uint32 = 42
                                if len(v9) > 0 {
                                    v9[0] = (fi12 | fi11)
                                }
                                _ = v31
                            }
                        case "":
                            var v32 uint = (v1 & v1)
                            _ = v32
                        case "foo bar":
                            if len(v9) > 0 {
                                v9[0] = 32768
                            }
                        }
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    v3 += v3
                    _ = fi12
                }
                fmt.Println(fi11, v1, v3)
                fmt.Println(v1, v3)
            } else {
                v3 -= v3
                var v33 int16 = 570
                var v34 uint32 = 986
                _ = v33
                _ = v34
            }
            for fi35 := 0; fi35 < 4; fi35++ {
                v9 = append(v9, fi35)
                if (!(0.1 >= 0.1)) {
                    v36 := GenS4{F0: 105}
                    if ((fi11 | 32768) < (^(fi35))) {
                        v37 := GenS4{F0: 109}
                        if ((fi11 == (fi11 & fi11)) || true) {
                            v38 := GenS5{F0: (!(1 <= -1))}
                            v2[0] = ""
                            v1 = 887
                            _ = v38
                        } else {
                            v39 := GenS6{F0: v1}
                            _ = v39
                        }
                        var v40 bool = (true || ((305 <= 1) && false))
                        var v41 int64 = 9223372036854775807
                        _ = v37
                        _ = v40
                        _ = v41
                    }
                    switch ("" + "") {
                    case "hello":
                        switch 2147483647 {
                        case 590:
                            fmt.Println(fi35, fi35, fi35)
                        default:
                            switch (fi11 ^ fi35) {
                            case -2147483648:
                                switch fi11 {
                                case 2147483647:
                                    var v42 uint8 = 0
                                    _ = v42
                                default:
                                    if false {
                                        v1 += 399
                                    } else {
                                        v1 = (v1 % (((^(v1))) | 1))
                                        delete(v2, 42)
                                        switch "hello" {
                                        case "":
                                            v43 := GenS7{F0: 210, F1: 255}
                                            _ = v43
                                        case "a":
                                            var v44 int16 = -1
                                            _ = v44
                                        case "foo bar":
                                            v1 = v1
                                        default:
                                            v45 := GenS8{F0: 88, F1: 127, F2: true}
                                            _ = v45
                                        }
                                    }
                                }
                            case 513:
                                fmt.Println(v1, v1)
                            default:
                                fmt.Println(fi11)
                            }
                        }
                    case "foo bar":
                        v46 := GenS8{F0: 127, F1: fi11, F2: (0 >= 42)}
                        _ = v46
                    case "\n\t":
                        v3 += 549
                    default:
                        var v47 int64 = 0
                        var v48 float64 = ((-2.25 + 1e10) + (1e10 + ((1e10 / ((-2.25 * -2.25) + float64((-2.25 * 0.1)))) / 1e10)))
                        v48 += v48
                        var v49 bool = true
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                    _ = v36
                } else {
                    v1 += 42
                    fmt.Println(v3, fi11)
                }
                _ = fi35
            }
        }
        switch 128 {
        case -568:
            fmt.Println(fi11, v1, v1)
        case 32768:
            v9 = append(v9, fi11)
        case 128:
            v50 := GenS8{F0: 42, F1: fi11, F2: (true && (-85 >= -1))}
            _ = v50
        default:
            v1 += v1
        }
        var v51 int16 = -287
        _ = fi11
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
