package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := []int{32767}
    v3 := make(map[int]bool)
    for fi4 := 0; fi4 < 8; fi4++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v5 := GenS1{F0: 2147483647, F1: 255, F2: ("" + "\n\t")}
        if len(v2) > 0 {
            v2[0] = fi4
        }
        _ = fi4
        _ = v5
    }
    delete(v3, 9223372036854775807)
    if ((1.5 + 0.0) < ((float64(float32(-2.25)) + 0.1) / 1e10)) {
        var v6 string = ((("hello" + ("" + "\n\t")) + "") + "foo bar")
        _ = v6
    }
    var v7 int8 = 0
    if (!(false && (663 != 42))) {
        if len(v2) > 0 {
            v2[0] = 32767
        }
        for fi8 := 0; fi8 < 15; fi8++ {
            if false {
                v9 := GenS1{F0: int32(v7), F1: 0, F2: ("\n\t" + ("hello" + ""))}
                _ = v9
            } else {
                if false {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        v7 -= 1
                        v7 = v7
                        _ = fi10
                    }
                    if false {
                        var v11 float32 = (1e10 + (-2.25 - (((0.1 / 0.1) * 1.5) - float32(0.0))))
                        var v12 uint8 = 2
                        v3[128] = true
                        _ = v11
                        _ = v12
                    }
                    var v13 float64 = float64(1.5)
                    _ = v13
                } else {
                    v14 := GenS1{F0: 42, F1: 128, F2: "a"}
                    for fi15 := 0; fi15 < 15; fi15++ {
                        var v16 uint16 = 554
                        _ = fi15
                        _ = v16
                    }
                    v17 := GenS1{F0: 569, F1: 14, F2: ("a" + "\n\t")}
                    _ = v14
                    _ = v17
                }
                delete(v3, (fi8 | 32767))
            }
            _ = fi8
        }
    }
    if true {
        var v18 int16 = 247
        if true {
            delete(v1, "")
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                if true {
                    v20 := GenS2{F0: true}
                    v21 := GenS2{F0: false}
                    switch fi19 {
                    case 2147483647:
                        var v22 float64 = (1.5 - (float64((float32((float64((-2.25 + (1.5 + 3.14159))) - 1e10)) / 1e10)) * float64(0.0)))
                        _ = v22
                    case 733:
                        v2 = append(v2, (fi19 | (^(fi19))))
                    case 42:
                        var v23 uint = 18446744073709551615
                        _ = v23
                    default:
                        v24 := GenS2{F0: (true && true)}
                        _ = v24
                    }
                    _ = v20
                    _ = v21
                } else {
                    switch fi19 {
                    case 32767:
                        v25 := GenS2{F0: ((-2.25 * -2.25) != 3.14159)}
                        _ = v25
                    case -1:
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v7 -= -1
                            _ = fi26
                        }
                    default:
                        v18 = (v18 << ((65535) & 63))
                    }
                    var v27 uint32 = 2147483648
                    _ = v27
                }
                switch "hello" {
                case "foo bar":
                    switch (fi19 + -2147483648) {
                    case 127:
                        var v28 float64 = 0.0
                        _ = v28
                    case 32768:
                        for fi29 := 0; fi29 < 17; fi29++ {
                            v7 += (v7 | v7)
                            var v30 uint64 = 201
                            v2 = append(v2, fi19)
                            _ = fi29
                            _ = v30
                        }
                    default:
                        switch "\n\t" {
                        case "\n\t":
                            if ("" > "foo bar") {
                                switch fi19 {
                                case 4294967295:
                                    v31 := v1[((("foo bar" + "hello") + ("hello" + "hello")) + "foo bar")]
                                    _ = v31
                                case 312:
                                    v32 := GenS1{F0: -2147483648, F1: 2, F2: (("a" + ("a" + ("" + "a"))) + ("hello" + ("hello" + "hello")))}
                                    _ = v32
                                case 568:
                                    fmt.Println(v18)
                                default:
                                    switch "a" {
                                    case "":
                                        v1[("foo bar" + ((((("\n\t" + "") + "a") + ("a" + "a")) + (("hello" + (("" + ("a" + ("" + (("a" + ("" + ("\n\t" + ("foo bar" + ("foo bar" + ((("" + "\n\t") + "") + "")))))) + ("hello" + (("hello" + "foo bar") + (("foo bar" + "a") + ((((("hello" + (("foo bar" + ("\n\t" + "\n\t")) + ("" + "\n\t"))) + ("" + "a")) + ("hello" + "foo bar")) + ("foo bar" + ("" + "hello"))) + "a")))))))) + "foo bar")) + "hello")) + "foo bar"))] = fi19
                                    case "\n\t":
                                        var v33 bool = (!(!((v7 * (v7 >> ((0) & 63))) < 127)))
                                        _ = v33
                                    default:
                                        v18 = v18
                                    }
                                }
                            } else {
                                switch "hello" {
                                case "a":
                                    if true {
                                        v7 = (^(v7))
                                        if false {
                                            var v34 int32 = 1
                                            _ = v34
                                        } else {
                                            for fi35 := 0; fi35 < 11; fi35++ {
                                                v36 := GenS2{F0: (!(!true))}
                                                var v37 string = "\n\t"
                                                _ = fi35
                                                _ = v36
                                                _ = v37
                                            }
                                            switch "a" {
                                            case "hello":
                                                var v38 float32 = 1.5
                                                _ = v38
                                            default:
                                                var v39 int64 = 136
                                                _ = v39
                                            }
                                            _ = len(v3)
                                        }
                                    }
                                case "hello":
                                    var v40 int8 = (v7 << ((255) & 63))
                                    _ = v40
                                case "\n\t":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    _ = len(v2)
                                }
                            }
                        case "a":
                            var v41 int = fi19
                            _ = v41
                        default:
                            for fi42 := 0; fi42 < 11; fi42++ {
                                fmt.Println(fi19, fi19, v18)
                                _ = fi42
                            }
                        }
                    }
                default:
                    for fi43 := 0; fi43 < 3; fi43++ {
                        fmt.Println(v7, v7, fi43)
                        v7 += (v7 / ((v7) | 1))
                        if len(v2) > 0 {
                            v2[0] = (fi43 << ((65535) & 63))
                        }
                        _ = fi43
                    }
                }
                var v44 int64 = 0
                _ = fi19
                _ = v44
            }
            v18 = v18
            v18 += (v18 >> ((4294967295) & 63))
        }
        v7 += (^(v7))
        _ = v18
    } else {
        for fi45 := 0; fi45 < 3; fi45++ {
            v7 = -127
            var v46 string = ("" + "foo bar")
            delete(v3, (-(fi45)))
            _ = fi45
            _ = v46
        }
    }
    v47 := GenS3{F0: 854, F1: v7}
    if (!true) {
        var v48 int8 = (v7 % ((v7) | 1))
        _ = v48
    } else {
        if (!true) {
            var v49 int16 = -1
            switch "\n\t" {
            case "a":
                switch (("foo bar" + ("foo bar" + "foo bar")) + "foo bar") {
                case "hello":
                    for fi50 := 0; fi50 < 12; fi50++ {
                        var v51 int16 = v49
                        v49 = 705
                        _ = fi50
                        _ = v51
                    }
                case "a":
                    v49 += v49
                }
            case "\n\t":
                fmt.Println(v7)
            case "foo bar":
                _ = len(v3)
            default:
                if (!false) {
                    for fi52 := 0; fi52 < 3; fi52++ {
                        v49 -= v49
                        _ = fi52
                    }
                    for fi53 := 0; fi53 < 8; fi53++ {
                        v54 := GenS3{F0: 585, F1: (v7 << ((25) & 63))}
                        _ = fi53
                        _ = v54
                    }
                }
            }
            _ = v49
        }
        var v55 uint8 = 2
        _ = v55
    }
    v56 := GenS2{F0: false}
    var v57 bool = (455 > 2147483647)
    var v58 uint16 = 65535
    var v59 uint64 = 42
    v7 -= v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v7
}
