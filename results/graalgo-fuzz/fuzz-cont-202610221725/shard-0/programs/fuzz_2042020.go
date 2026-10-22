package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint32 }
type GenS2 struct { F0 int16; F1 string; F2 int }
type GenS3 struct { F0 int; F1 bool; F2 int16 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int16) int64 {
    switch 127 {
    case 0:
        var v5 uint16 = 255
        _ = v5
    }
    switch 715 {
    case -56:
        var v6 int8 = 100
        _ = v6
    case 32767:
        switch 42 {
        case -690:
            v7 := []int{int(p0)}
            _ = v7
        default:
            v8 := []int{127, -1}
            _ = v8
        }
    default:
        p0 = 1
    }
    _ = p0
    return 42
}
func genFn1(p0 uint32) int8 {
    fmt.Println(p0)
    var v9 int = 9223372036854775807
    var v10 int16 = 1000
    _ = p0
    _ = v10
    _ = v9
    return -1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{""}
    v3 := []uint32{}
    var v4 uint = 65535
    v11 := GenS1{F0: (v4 << (((v4 - (v4 | 4294967295))) & 63)), F1: -1, F2: 65535}
    fmt.Println(v4)
    v12 := GenS1{F0: v4, F1: 1, F2: 0}
    var v13 uint8 = 127
    if (false && false) {
        var v14 uint16 = 1
        _ = v14
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        if ("" != "\n\t") {
            v15 := GenS2{F0: -622, F1: ("a" + ""), F2: -9223372036854775808}
            _ = v15
        }
        if (-61 == 42) {
            if (("a" + "\n\t") < "") {
                v13 += v13
                if false {
                    var v16 int8 = 42
                    switch 2147483647 {
                    case 841:
                        for fi17 := 0; fi17 < 15; fi17++ {
                            for fi18 := 0; fi18 < 10; fi18++ {
                                var v19 int = (fi17 >> ((v4) & 63))
                                if true {
                                    if (326 < 0) {
                                        _ = len(v3)
                                        var v20 bool = true
                                        fmt.Println(v19)
                                        _ = v20
                                    }
                                }
                                switch ("\n\t" + "foo bar") {
                                case "":
                                    _ = len(v1)
                                case "foo bar":
                                    if false {
                                        fmt.Println(v19, v16)
                                    } else {
                                        fmt.Println(v16, v19)
                                        switch "\n\t" {
                                        case "hello":
                                            v21 := GenS2{F0: -292, F1: (("\n\t" + "hello") + ("a" + "foo bar")), F2: (v19 % (((-(v19))) | 1))}
                                            _ = v21
                                        case "\n\t":
                                            v13 += v13
                                            v22 := GenS3{F0: fi18, F1: true, F2: -1}
                                            _ = v22
                                        case "foo bar":
                                            v16 += -66
                                        default:
                                            v4 += 189
                                        }
                                        fmt.Println(v13)
                                        v13 -= 127
                                        if (!(true || false)) {
                                            if ((v16 + genFn1(256)) != (^(v16))) {
                                                delete(v1, "hello")
                                                if (true && (474 >= 2)) {
                                                    var v23 int16 = 1000
                                                    if (fi17 >= -1) {
                                                        var v24 uint16 = 750
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v25 uint64 = 2
                                                            _ = v25
                                                        case "\n\t":
                                                            v23 = (v23 & (v23 - (v23 + (v23 + v23))))
                                                        case "hello":
                                                            var v26 int64 = int64(v23)
                                                            _ = v26
                                                        default:
                                                            var v27 uint32 = 65535
                                                            _ = v27
                                                        }
                                                        _ = v24
                                                    } else {
                                                        fmt.Println(v4)
                                                        switch 32767 {
                                                        case 2147483647:
                                                            v23 += (v23 & v23)
                                                        case 651:
                                                            if (((-2.25 / 1e10) - 1e10) <= (1.5 * 0.1)) {
                                                                delete(v1, "a")
                                                                fmt.Println(v4)
                                                            } else {
                                                                if false {
                                                                    if false {
                                                                        var v28 int16 = -32768
                                                                        _ = v28
                                                                    } else {
                                                                        v4 += v4
                                                                        v29 := GenS1{F0: (v4 & (v4 - v4)), F1: v16, F2: 42}
                                                                        _ = len(v1)
                                                                        _ = v29
                                                                    }
                                                                    v30 := GenS1{F0: (v4 ^ 65535), F1: v16, F2: 47}
                                                                    _ = v30
                                                                }
                                                                if false {
                                                                    if (!true) {
                                                                        v23 = v23
                                                                        delete(v1, ("\n\t" + ""))
                                                                    }
                                                                    var v31 uint64 = 18446744073709551615
                                                                    v32 := v1["hello"]
                                                                    v33 := GenS1{F0: 4294967295, F1: v16, F2: (v32 / ((v32) | 1))}
                                                                    _ = v31
                                                                    _ = v32
                                                                    _ = v33
                                                                } else {
                                                                    v34 := GenS4{F0: 0}
                                                                    v35 := GenS2{F0: (v23 << (((v4 % ((792) | 1))) & 63)), F1: ("\n\t" + ("foo bar" + (("" + ("a" + ("foo bar" + ""))) + ""))), F2: (fi18 | fi18)}
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                            }
                                                        default:
                                                            var v36 uint64 = 4294967295
                                                            _ = v36
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = 4294967295
                                                        }
                                                    }
                                                    fmt.Println(v4, fi17)
                                                    _ = v23
                                                } else {
                                                    if (-2.25 == (0.1 + (-2.25 / -2.25))) {
                                                        v13 += 1
                                                        v4 = v4
                                                        var v37 uint64 = 18446744073709551615
                                                        _ = v37
                                                    } else {
                                                        switch "\n\t" {
                                                        case "":
                                                            fmt.Println(v19)
                                                        case "a":
                                                            v38 := GenS3{F0: fi17, F1: ((true && false) && true), F2: -32768}
                                                            _ = v38
                                                        default:
                                                            var v39 uint = 334
                                                            _ = v39
                                                        }
                                                    }
                                                }
                                            }
                                            fmt.Println(v19, fi17, v16)
                                            v4 = (v4 ^ (v4 << ((568) & 63)))
                                            if len(v2) > 0 {
                                                v2[0] = ("a" + ("\n\t" + ("foo bar" + ("" + "hello"))))
                                            }
                                            var v40 uint32 = 65535
                                            _ = v40
                                        }
                                    }
                                case "a":
                                    fmt.Println(fi18)
                                    v4 += v4
                                    if (!true) {
                                        switch 128 {
                                        case -9223372036854775808:
                                            var v41 uint16 = 42
                                            _ = v41
                                        case 4294967295:
                                            if false {
                                                v42 := GenS2{F0: -1, F1: "foo bar", F2: int(v4)}
                                                switch ("\n\t" + "\n\t") {
                                                case "foo bar":
                                                    fmt.Println(fi17, v16)
                                                case "\n\t":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                default:
                                                    if len(v3) > 0 {
                                                        v3[0] = 1
                                                    }
                                                }
                                                _ = v42
                                            }
                                        case -1:
                                            if true {
                                                var v43 float32 = float32(1.5)
                                                _ = v43
                                            }
                                        }
                                        var v44 uint = v4
                                        switch (fi17 / ((4294967295) | 1)) {
                                        case 0:
                                            var v45 uint = (v44 ^ (v4 + (^(v4))))
                                            _ = v45
                                        case 4294967295:
                                            v44 = 42
                                        default:
                                            switch v19 {
                                            case -792:
                                                _ = len(v2)
                                                v19 = -9223372036854775808
                                                v16 = (v16 | (v16 << (((v4 ^ v44)) & 63)))
                                            case 32767:
                                                switch -1 {
                                                case 32768:
                                                    v13 = v13
                                                case -260:
                                                    var v46 int = (fi18 ^ fi18)
                                                    _ = v46
                                                case -9223372036854775808:
                                                    var v47 int8 = (-(v16))
                                                    _ = v47
                                                }
                                            default:
                                                if true {
                                                    var v48 int = (v19 / ((-709) | 1))
                                                    _ = v48
                                                }
                                            }
                                        }
                                        _ = v44
                                    } else {
                                        var v49 uint = v4
                                        _ = v49
                                    }
                                }
                                _ = fi18
                                _ = v19
                            }
                            delete(v1, ("" + "hello"))
                            _ = fi17
                        }
                    default:
                        var v50 string = "foo bar"
                        _ = v50
                    }
                    _ = v16
                } else {
                    if false {
                        v1["hello"] = 359
                    } else {
                        v51 := GenS3{F0: -809, F1: (18446744073709551615 <= 42), F2: 333}
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 19; fi52++ {
                        v1[(((("a" + "foo bar") + "") + "foo bar") + ("hello" + ""))] = 676
                        _ = fi52
                    }
                }
                if (18446744073709551615 == 18446744073709551615) {
                    v13 = v13
                    v13 = 255
                    var v53 int16 = 1
                    _ = v53
                }
            }
            v13 -= 138
        } else {
            var v54 string = ((("hello" + "a") + "") + "foo bar")
            _ = v54
        }
    }
    var v55 uint32 = 4294967295
    delete(v1, "hello")
    for fi56 := 0; fi56 < 6; fi56++ {
        switch fi56 {
        case 2147483647:
            if ((v13 - (v13 * (v13 & v13))) <= 255) {
                fmt.Println(v55, v55, fi56)
            }
        case -9223372036854775808:
            fmt.Println(fi56)
        default:
            if true {
                var v57 string = "foo bar"
                _ = len(v1)
                _ = v57
            }
        }
        v55 = v55
        _ = fi56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v55
}
