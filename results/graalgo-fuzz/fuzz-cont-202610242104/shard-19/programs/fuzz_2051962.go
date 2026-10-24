package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 int8; F2 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 int64 }

func main() {
    var v1 int16 = -32768
    var v2 int64 = 42
    var v3 float64 = 1.5
    var v4 uint16 = 32767
    var v5 int8 = 127
    var v6 uint32 = 2147483647
    fmt.Println(v2)
    v7 := GenS1{F0: 0}
    if (true || true) {
        v8 := GenS2{F0: 418}
        var v9 uint8 = 101
        switch 42 {
        case 127:
            v10 := GenS3{F0: -9223372036854775808, F1: v1}
            _ = v10
        case 42:
            fmt.Println(v4, v4)
        case -747:
            v11 := GenS4{F0: (v4 * (v4 - (v4 - 255)))}
            _ = v11
        }
        _ = v8
        _ = v9
    } else {
        switch 1 {
        case 4294967295:
            v12 := GenS5{F0: 32767, F1: v5, F2: v4}
            _ = v12
        case -163:
            v13 := []int{}
            _ = v13
        }
        v5 = v5
    }
    v14 := []int{int(v1), 0, 32768, 9223372036854775807}
    if len(v14) > 0 {
        v14[0] = 673
    }
    switch "a" {
    case "hello":
        v15 := GenS1{F0: 389}
        _ = v15
    case "\n\t":
        switch 4294967295 {
        case 32767:
            var v16 float32 = 1.5
            _ = v16
        case 955:
            var v17 int16 = v1
            _ = v17
        case 4294967295:
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 uint = 495
                if false {
                    if true {
                        if (!false) {
                            var v20 float64 = 3.14159
                            _ = v20
                        } else {
                            v1 += v1
                            fmt.Println(v2, v6, v1)
                        }
                    } else {
                        v4 += 255
                        fmt.Println(v3, v4, v2)
                        v2 = 1
                    }
                    for fi21 := 0; fi21 < 16; fi21++ {
                        var v22 int32 = -1
                        v2 = v2
                        _ = fi21
                        _ = v22
                    }
                    v23 := GenS1{F0: 18446744073709551615}
                    _ = v23
                }
                v24 := GenS6{F0: (v2 - v2)}
                _ = fi18
                _ = v19
                _ = v24
            }
        default:
            fmt.Println(v6)
        }
    case "":
        switch "foo bar" {
        case "":
            v25 := GenS6{F0: (^(v2))}
            _ = v25
        case "a":
            v14 = append(v14, 1)
        case "foo bar":
            if (((float64(((0.0 - (0.1 - ((1.5 * 0.1) + ((-2.25 - (3.14159 * ((3.14159 + -2.25) / float32((0.1 + v3))))) - 3.14159)))) / 1.0)) < 0.1) || true) || (!(false && (0 == 4294967295)))) {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            } else {
                var v26 int32 = 2147483647
                v27 := GenS2{F0: 2147483647}
                _ = v26
                _ = v27
            }
        default:
            if (18446744073709551615 > 788) {
                switch (((("hello" + ("" + ((("hello" + ("foo bar" + (("\n\t" + "a") + ""))) + "\n\t") + ""))) + "") + "hello") + ((("" + ("\n\t" + ("hello" + "hello"))) + "hello") + (("foo bar" + (("hello" + "foo bar") + (((("" + "hello") + "a") + ("\n\t" + (((("foo bar" + "a") + (("a" + "\n\t") + ((("" + ("" + "")) + ("" + "hello")) + "hello"))) + "hello") + "hello"))) + ""))) + "\n\t"))) {
                case "\n\t":
                    for fi28 := 0; fi28 < 20; fi28++ {
                        fmt.Println(v5, v1)
                        fmt.Println(v4)
                        v3 = ((0.1 * v3) + float64((((1.5 / (((((1e10 / float32((v3 / 1.0))) / (0.1 * 1e10)) * 1.5) / 0.1) / float32(v3))) + (3.14159 / (0.1 - 0.0))) - float32(v3))))
                        _ = fi28
                    }
                case "hello":
                    for fi29 := 0; fi29 < 14; fi29++ {
                        if ("foo bar" > "\n\t") {
                            if ((true && false) && (42 >= -2147483648)) {
                                var v30 uint64 = 2
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                _ = v30
                            }
                            if ((v4 - 0) >= v4) {
                                switch "" {
                                case "foo bar":
                                    var v31 int8 = v5
                                    _ = v31
                                default:
                                    fmt.Println(v1, fi29, v3)
                                }
                                switch int(v2) {
                                case 32768:
                                    for fi32 := 0; fi32 < 14; fi32++ {
                                        var v33 float64 = 1e10
                                        _ = fi32
                                        _ = v33
                                    }
                                case -2147483648:
                                    switch (fi29 & int(v5)) {
                                    case 42:
                                        if true {
                                            var v34 uint64 = 18446744073709551615
                                            _ = v34
                                        }
                                    case 218:
                                        switch fi29 {
                                        case 4294967295:
                                            var v35 uint16 = 1
                                            _ = v35
                                        default:
                                            var v36 int8 = -1
                                            _ = v36
                                        }
                                    case -199:
                                        v6 = (v6 / (((v6 % (((v6 ^ (^(v6)))) | 1))) | 1))
                                    default:
                                        if (v2 < v2) {
                                            if true {
                                                v3 = v3
                                                var v37 int8 = (v5 - v5)
                                                _ = v37
                                            }
                                            switch fi29 {
                                            case 710:
                                                v4 = 871
                                            default:
                                                fmt.Println(v4, v3, v2)
                                            }
                                        } else {
                                            for fi38 := 0; fi38 < 12; fi38++ {
                                                var v39 float32 = 1.5
                                                var v40 uint = uint(v4)
                                                _ = fi38
                                                _ = v39
                                                _ = v40
                                            }
                                            var v41 int64 = (v2 ^ 1)
                                            _ = len(v14)
                                            _ = v41
                                        }
                                    }
                                case 9223372036854775807:
                                    if true {
                                        v14 = append(v14, (fi29 / (((-(fi29))) | 1)))
                                    } else {
                                        for fi42 := 0; fi42 < 3; fi42++ {
                                            v6 = v6
                                            var v43 uint32 = v6
                                            var v44 float64 = v3
                                            if len(v14) > 0 {
                                                v14[0] = fi29
                                            }
                                            _ = fi42
                                            _ = v43
                                            _ = v44
                                        }
                                    }
                                default:
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                }
                            }
                        } else {
                            v2 += 0
                            if len(v14) > 0 {
                                v14[0] = (^(fi29))
                            }
                            v5 += -1
                        }
                        _ = fi29
                    }
                case "a":
                    fmt.Println(v1)
                }
            } else {
                var v45 bool = ((!((((("a" + "\n\t") + "a") + "a") != ("foo bar" + ("\n\t" + ""))) || true)) && true)
                _ = v45
            }
        }
    default:
        switch (("a" + "\n\t") + (("" + "") + "a")) {
        case "a":
            var v46 int = 42
            _ = v46
        default:
            fmt.Println(v5, v1)
        }
    }
    _ = len(v14)
    for fi47 := 0; fi47 < 12; fi47++ {
        var v48 int = fi47
        v49 := GenS7{F0: (v4 + 65535), F1: (true || false)}
        v48 += (v48 / (((v48 & -689)) | 1))
        _ = fi47
        _ = v48
        _ = v49
    }
    v50 := GenS3{F0: -141, F1: v1}
    var v51 uint = 18446744073709551615
    for fi52 := 0; fi52 < 14; fi52++ {
        fmt.Println(v1, v6, v51)
        var v53 int = fi52
        _ = fi52
        _ = v53
    }
    v54 := GenS8{F0: (v2 & -9223372036854775808)}
    var v55 uint8 = uint8(v51)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v54
    _ = v55
    _ = v6
    _ = v7
}
