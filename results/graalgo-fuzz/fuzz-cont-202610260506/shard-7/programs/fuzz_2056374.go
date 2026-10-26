package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int8; F1 uint16; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 int32; F2 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 int64; F2 string }
type GenS8 struct { F0 bool; F1 uint8; F2 int16 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 int64 }
func genFn0(p0 uint, p1 int32) int16 {
    p0 = p0
    p1 = (-(p1))
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 bool = true
    v2 := []string{"", "a", "hello", ("a" + "a")}
    var v3 float32 = (-2.25 * 1e10)
    var v4 float64 = 1.5
    v2 = append(v2, "foo bar")
    fmt.Println(v4)
    v3 = v3
    for fi5 := 0; fi5 < 2; fi5++ {
        switch "hello" {
        case "\n\t":
            v6 := GenS1{F0: fi5, F1: (!v1)}
            _ = v6
        case "a":
            if (42 <= -1) {
                if (42 <= 0) {
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    v7 := GenS2{F0: 897, F1: 0}
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 8; fi8++ {
                        v3 = ((v3 * (0.1 * v3)) - (v3 - v3))
                        _ = fi8
                    }
                    v3 = 0.1
                    v4 = float64(float32(v4))
                }
            }
        case "foo bar":
            var v9 string = ""
            _ = v9
        }
        fmt.Println(v3, v1, v4)
        _ = fi5
    }
    v10 := GenS3{F0: 36, F1: 65535, F2: ("\n\t" + "hello")}
    var v11 int = -591
    var v12 int16 = 1000
    v1 = v1
    v4 = ((v4 / (-2.25 / v4)) * (0.1 - float64(v3)))
    fmt.Println(v11, v3)
    switch "hello" {
    case "":
        switch ((("foo bar" + "\n\t") + "a") + ((("foo bar" + "\n\t") + "\n\t") + "a")) {
        case "foo bar":
            switch "\n\t" {
            case "a":
                var v13 int16 = (v12 * v12)
                _ = v13
            }
        case "\n\t":
            if v1 {
                v11 = v11
            }
        case "hello":
            switch (("hello" + "\n\t") + "a") {
            case "a":
                switch -9223372036854775808 {
                case -185:
                    var v14 int32 = 1
                    _ = v14
                case 42:
                    fmt.Println(v12, v3)
                }
            case "foo bar":
                var v15 int8 = 42
                _ = v15
            case "\n\t":
                var v16 uint8 = 255
                _ = v16
            default:
                for fi17 := 0; fi17 < 2; fi17++ {
                    if (-2147483648 >= -357) {
                        var v18 float64 = float64(0.1)
                        for fi19 := 0; fi19 < 13; fi19++ {
                            switch "" {
                            case "\n\t":
                                v4 = v4
                            case "a":
                                v20 := GenS3{F0: -1, F1: 2, F2: "foo bar"}
                                _ = v20
                            case "":
                                v21 := GenS4{F0: true}
                                _ = v21
                            default:
                                fmt.Println(v3)
                            }
                            v11 = 128
                            var v22 int16 = (v12 / ((v12) | 1))
                            v23 := GenS5{F0: 683, F1: 0, F2: 0}
                            _ = fi19
                            _ = v22
                            _ = v23
                        }
                        if false {
                            fmt.Println(v18)
                        } else {
                            v24 := GenS6{F0: 2147483647}
                            _ = len(v2)
                            _ = v24
                        }
                        _ = v18
                    } else {
                        for fi25 := 0; fi25 < 20; fi25++ {
                            _ = len(v2)
                            _ = fi25
                        }
                    }
                    _ = fi17
                }
            }
        }
    case "a":
        v3 = (v3 - v3)
    case "hello":
        v26 := GenS7{F0: (-(v11)), F1: -9223372036854775808, F2: "hello"}
        _ = v26
    }
    v1 = (!(!(v3 <= (v3 - (1e10 * v3)))))
    v3 += (v3 - 0.0)
    switch v11 {
    case -197:
        if (((v12 & v12) > v12) || v1) {
            var v27 int32 = 42
            var v28 int16 = v12
            switch (v11 ^ (v11 | -1)) {
            case 2147483647:
                switch "foo bar" {
                case "\n\t":
                    v29 := GenS5{F0: 42, F1: (-(v27)), F2: (-(v27))}
                    _ = v29
                }
            default:
                for fi30 := 0; fi30 < 14; fi30++ {
                    if (v1 || (((v3 + 0.1) * (v3 * (v3 - v3))) <= 1e10)) {
                        fmt.Println(v3, v11, v11)
                        switch (-(v11)) {
                        case 2147483647:
                            if v1 {
                                var v31 uint = 255
                                if (!(true || ((v1 && (127 < 1)) || ((v28 - (v12 / ((v12) | 1))) < v12)))) {
                                    if ((!(v3 > 1e10)) && v1) {
                                        v32 := GenS8{F0: v1, F1: 42, F2: (v28 << (((^(v31))) & 63))}
                                        v28 -= 42
                                        _ = v32
                                    } else {
                                        var v33 int = (fi30 * (v11 << (((v31 % ((v31) | 1))) & 63)))
                                        _ = v33
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                } else {
                                    if (!v1) {
                                        if v1 {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    }
                                }
                                _ = v31
                            } else {
                                if ((v28 / ((-32768) | 1)) < 0) {
                                    _ = len(v2)
                                    var v34 int16 = (-(v28))
                                    var v35 bool = v1
                                    _ = v34
                                    _ = v35
                                }
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                    } else {
                        fmt.Println(v4)
                        switch 32768 {
                        case 2147483647:
                            var v36 int64 = 1
                            _ = v36
                        }
                        if (1 < 896) {
                            v1 = (0.0 < v4)
                            v37 := GenS9{F0: v27, F1: 128}
                            _ = v37
                        } else {
                            if (42 > -9223372036854775808) {
                                if (("hello" + "foo bar") < "") {
                                    for fi38 := 0; fi38 < 17; fi38++ {
                                        if (int32(v28) > int32(fi30)) {
                                            if len(v2) > 0 {
                                                v2[0] = "foo bar"
                                            }
                                            if true {
                                                v28 += (v28 >> ((1) & 63))
                                                var v39 int8 = 100
                                                _ = v39
                                            }
                                            var v40 uint = 42
                                            _ = v40
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi38
                                    }
                                    if v1 {
                                        v12 -= (v12 & (^(v28)))
                                    } else {
                                        fmt.Println(v11)
                                    }
                                }
                                var v41 uint16 = 1
                                v42 := GenS4{F0: (!((!(1 < 255)) || ((v1 && (!((1.5 != v4) && ((^(v12)) >= genFn0(560, (^(v27))))))) || v1)))}
                                _ = v41
                                _ = v42
                            } else {
                                if true {
                                    for fi43 := 0; fi43 < 18; fi43++ {
                                        fmt.Println(fi30, fi43)
                                        v44 := GenS2{F0: 233, F1: 0}
                                        var v45 uint16 = 59
                                        _ = fi43
                                        _ = v44
                                        _ = v45
                                    }
                                    v11 = fi30
                                } else {
                                    for fi46 := 0; fi46 < 8; fi46++ {
                                        var v47 int16 = 1
                                        fmt.Println(v47)
                                        _ = fi46
                                        _ = v47
                                    }
                                    if (v1 || v1) {
                                        var v48 uint8 = 117
                                        v49 := GenS3{F0: 42, F1: 149, F2: ("" + "foo bar")}
                                        if (127 < -103) {
                                            fmt.Println(fi30)
                                        } else {
                                            v50 := GenS4{F0: v1}
                                            var v51 uint8 = 42
                                            _ = v50
                                            _ = v51
                                        }
                                        _ = v48
                                        _ = v49
                                    } else {
                                        v11 -= v11
                                        switch "a" {
                                        case "\n\t":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case "hello":
                                            for fi52 := 0; fi52 < 13; fi52++ {
                                                v12 = v28
                                                v53 := GenS4{F0: (42 > 1)}
                                                _ = fi52
                                                _ = v53
                                            }
                                        default:
                                            fmt.Println(v27, v27, v11)
                                        }
                                        v54 := GenS4{F0: v1}
                                        _ = v54
                                    }
                                }
                                v12 += (v28 | (v28 | (v28 ^ genFn0(644, -2147483648))))
                            }
                        }
                    }
                    var v55 int32 = v27
                    switch ("hello" + ("foo bar" + "")) {
                    case "foo bar":
                        v28 = genFn0(65535, v55)
                    default:
                        var v56 uint = 619
                        _ = v56
                    }
                    _ = fi30
                    _ = v55
                }
            }
            _ = v27
            _ = v28
        }
    }
    if v1 {
        v57 := GenS10{F0: 287}
        for fi58 := 0; fi58 < 1; fi58++ {
            v12 += genFn0(42, 514)
            _ = fi58
        }
        _ = v57
    } else {
        for fi59 := 0; fi59 < 9; fi59++ {
            var v60 int32 = 0
            var v61 float32 = 0.0
            _ = fi59
            _ = v60
            _ = v61
        }
        var v62 bool = v1
        _ = v62
    }
    v11 = (-(v11))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v63 uint16 = 866
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v63
}
