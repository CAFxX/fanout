package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int16 }
type GenS2 struct { F0 int; F1 uint8; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 int32; F2 uint }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 int64; F1 int32 }
type GenS9 struct { F0 int8; F1 int }
type GenS10 struct { F0 uint64; F1 uint32; F2 int }
func genFn0() uint {
    var v5 uint = 2
    if (-2147483648 < 722) {
        v5 = v5
        fmt.Println(v5)
        v6 := GenS1{F0: 42, F1: 0, F2: 728}
        _ = v6
    } else {
        v7 := GenS2{F0: 4294967295, F1: 0, F2: -114}
        var v8 uint32 = 178
        _ = v7
        _ = v8
    }
    if (42 <= 65535) {
        v5 -= v5
        var v9 uint16 = 310
        _ = v9
    }
    v10 := []int{1, -9223372036854775808}
    _ = v10
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 2147483647
    var v3 string = ("hello" + "a")
    v4 := []string{}
    v3 = ""
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int32 = -2147483648
        v2 -= (v2 & fi11)
        _ = fi11
        _ = v12
    }
    fmt.Println(v3, v3)
    var v13 string = ("a" + "foo bar")
    if (((((v3 + ((v13 + v3) + ("hello" + ""))) + "foo bar") + v13) + (v3 + v13)) < v13) {
        var v14 uint16 = 1
        _ = v14
    }
    var v15 uint32 = 2
    if (!(1 <= 1)) {
        v1[(v13 + v13)] = (^(v15))
    } else {
        v3 = ""
        var v16 float32 = (0.1 - float32(0.0))
        fmt.Println(v13, v15)
        _ = v16
    }
    v15 = 1
    if ((3.14159 < (float32((1e10 * 0.0)) * 0.0)) || false) {
        v15 = (^(v15))
        switch (v3 + ("a" + ((((v13 + "hello") + (((v13 + (v13 + v13)) + v13) + v13)) + "") + "hello"))) {
        case "a":
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 uint = 198
                _ = fi17
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 int32 = int32(fi19)
            _ = fi19
            _ = v20
        }
    }
    v4 = append(v4, v13)
    var v21 int8 = 42
    if true {
        v13 = v13
    } else {
        if len(v4) > 0 {
            v4[0] = v13
        }
        v2 += (v2 ^ v2)
        for fi22 := 0; fi22 < 8; fi22++ {
            for fi23 := 0; fi23 < 1; fi23++ {
                v15 += (v15 & (v15 % (((v15 & v15)) | 1)))
                _ = fi23
            }
            switch v3 {
            case "\n\t":
                switch v3 {
                case "a":
                    switch v13 {
                    case "\n\t":
                        switch 42 {
                        case 0:
                            if false {
                                v21 = (-(v21))
                                if (!true) {
                                    for fi24 := 0; fi24 < 5; fi24++ {
                                        if (-1 <= 42) {
                                            v25 := GenS3{F0: 1000}
                                            v26 := GenS4{F0: 2}
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = fi24
                                    }
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(v21, v3, v15)
                                    default:
                                        switch v3 {
                                        case "a":
                                            if true {
                                                if false {
                                                    v1["foo bar"] = v15
                                                } else {
                                                    v27 := GenS1{F0: 42, F1: 2147483647, F2: -872}
                                                    var v28 uint16 = 128
                                                    _ = v27
                                                    _ = v28
                                                }
                                                v2 = (-(fi22))
                                                var v29 bool = true
                                                _ = v29
                                            }
                                        case "\n\t":
                                            var v30 uint = 544
                                            _ = v30
                                        case "":
                                            v15 = v15
                                        default:
                                            var v31 bool = false
                                            _ = v31
                                        }
                                    }
                                }
                            } else {
                                var v32 uint64 = 18446744073709551615
                                v33 := GenS1{F0: v32, F1: 0, F2: -32768}
                                _ = v32
                                _ = v33
                            }
                        case 32767:
                            v2 = fi22
                        default:
                            v2 -= (-(fi22))
                        }
                    case "foo bar":
                        if false {
                            var v34 uint16 = 773
                            _ = v34
                        } else {
                            v35 := GenS5{F0: 1}
                            fmt.Println(v21)
                            var v36 string = ("\n\t" + "a")
                            _ = v35
                            _ = v36
                        }
                    case "hello":
                        v13 = ((v13 + (((v3 + v13) + "foo bar") + (v3 + v13))) + v13)
                    default:
                        switch "\n\t" {
                        case "":
                            switch v13 {
                            case "a":
                                v1[((v13 + "hello") + v3)] = (v15 & (v15 & v15))
                            case "hello":
                                var v37 uint16 = 945
                                _ = v37
                            case "foo bar":
                                var v38 int8 = (v21 >> ((37) & 63))
                                _ = v38
                            }
                        }
                    }
                case "":
                    if false {
                        v39 := GenS2{F0: 32767, F1: 127, F2: (v21 ^ (v21 << ((genFn0()) & 63)))}
                        _ = v39
                    }
                default:
                    for fi40 := 0; fi40 < 11; fi40++ {
                        var v41 bool = false
                        if v41 {
                            var v42 int32 = 2147483647
                            var v43 int64 = 1
                            _ = v42
                            _ = v43
                        } else {
                            v44 := v1[v13]
                            v41 = (0 == 42)
                            v45 := GenS4{F0: 550}
                            v2 -= v2
                            _ = v44
                            _ = v45
                        }
                        var v46 float64 = (0.0 / 1.0)
                        fmt.Println(v15)
                        var v47 float32 = 0.0
                        v48 := GenS2{F0: (v2 + (v2 >> ((genFn0()) & 63))), F1: 128, F2: (v21 + v21)}
                        _ = fi40
                        _ = v41
                        _ = v46
                        _ = v47
                        _ = v48
                    }
                }
            case "foo bar":
                v21 = v21
            default:
                if true {
                    v49 := GenS6{F0: 9223372036854775807, F1: 2147483647, F2: 18446744073709551615}
                    fmt.Println(v2, v2)
                    _ = v49
                }
            }
            v2 = (fi22 | v2)
            _ = fi22
        }
    }
    for fi50 := 0; fi50 < 17; fi50++ {
        for fi51 := 0; fi51 < 1; fi51++ {
            v21 = int8(v15)
            var v52 uint16 = 32767
            _ = fi51
            _ = v52
        }
        if false {
            v21 += 12
        }
        fmt.Println(v2, v15, v15)
        _ = fi50
    }
    v13 = (("foo bar" + "\n\t") + (v3 + (("hello" + v13) + "\n\t")))
    v4 = append(v4, v13)
    if (1 >= 2) {
        if (445 != -2147483648) {
            var v53 int = v2
            if false {
                v54 := GenS4{F0: 639}
                v55 := GenS7{F0: 2, F1: "foo bar"}
                _ = v54
                _ = v55
            }
            var v56 bool = false
            _ = v53
            _ = v56
        } else {
            v57 := GenS8{F0: 42, F1: 2147483647}
            switch (v2 >> ((genFn0()) & 63)) {
            case 2147483647:
                if ((1e10 + (1.5 * (0.0 * (float64((float32(1.5) - (1.5 * 0.1))) + (1e10 - float64(0.0)))))) >= -2.25) {
                    v15 += (v15 * (v15 & v15))
                } else {
                    if ((v15 & (v15 << ((18446744073709551615) & 63))) > 42) {
                        v2 += v2
                    }
                    var v58 int64 = -9223372036854775808
                    _ = v58
                }
            case -896:
                switch v3 {
                case "a":
                    for fi59 := 0; fi59 < 1; fi59++ {
                        v60 := GenS9{F0: (v21 + v21), F1: (-(v2))}
                        _ = fi59
                        _ = v60
                    }
                case "hello":
                    switch v13 {
                    case "":
                        if (float32((0.1 / 3.14159)) > ((float32(-2.25) + 0.0) * (-2.25 * float32(1e10)))) {
                            v13 = ("hello" + (v13 + v13))
                            var v61 string = (v3 + "hello")
                            _ = v61
                        }
                    case "a":
                        v62 := GenS10{F0: 46, F1: 5, F2: v2}
                        _ = v62
                    default:
                        fmt.Println(v13, v3, v13)
                    }
                }
            case 32767:
                fmt.Println(v2)
            default:
                for fi63 := 0; fi63 < 15; fi63++ {
                    var v64 uint32 = 1
                    if true {
                        v13 = (v13 + v3)
                    }
                    _ = fi63
                    _ = v64
                }
            }
            _ = v57
        }
    } else {
        fmt.Println(v2, v2, v3)
        switch v2 {
        case -2147483648:
            v65 := GenS6{F0: 1, F1: -580, F2: 815}
            _ = v65
        case 42:
            fmt.Println(v13)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
