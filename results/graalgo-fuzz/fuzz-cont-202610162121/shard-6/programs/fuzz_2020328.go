package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 string; F1 int8; F2 string }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 int32; F1 bool; F2 int16 }
type GenS8 struct { F0 uint16; F1 uint64 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 bool; F1 int; F2 int }
type GenS11 struct { F0 uint64; F1 string }

func main() {
    v1 := []uint32{336, 632}
    v2 := []string{"foo bar", ("a" + "a")}
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    v5 := []string{"\n\t", (("" + ("a" + "\n\t")) + ("" + "hello")), ("hello" + ("foo bar" + (("foo bar" + ((("" + "foo bar") + "hello") + ("foo bar" + ""))) + "a"))), "hello"}
    v6 := GenS1{F0: (("" + ("" + "hello")) + "a"), F1: "hello"}
    switch 32767 {
    case 0:
        var v7 bool = false
        _ = v7
    case 149:
        var v8 uint = 572
        _ = v8
    case 128:
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(fi9)
            _ = fi9
        }
    default:
        if (true || (-1 < 7)) {
            for fi10 := 0; fi10 < 20; fi10++ {
                switch "hello" {
                case "\n\t":
                    v11 := GenS2{F0: 255, F1: 127}
                    _ = v11
                case "a":
                    if false {
                        if ((^(fi10)) <= fi10) {
                            var v12 int = 32768
                            _ = v12
                        }
                        var v13 uint64 = 18446744073709551615
                        switch ("foo bar" + "a") {
                        case "hello":
                            fmt.Println(v13)
                        default:
                            var v14 uint8 = uint8(v13)
                            _ = v14
                        }
                        _ = v13
                    }
                }
                fmt.Println(fi10)
                _ = fi10
            }
            var v15 int16 = 42
            _ = v15
        }
    }
    var v16 bool = true
    v16 = (!(2 == 905))
    v17 := GenS1{F0: (("\n\t" + (("" + ("hello" + "hello")) + "a")) + ("\n\t" + "a")), F1: "a"}
    v18 := GenS1{F0: "foo bar", F1: "foo bar"}
    v16 = true
    _ = len(v1)
    v16 = v16
    v4[-2147483648] = 4294967295
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := GenS3{F0: (65535 != 65535), F1: -29}
        _ = fi19
        _ = v20
    }
    if true {
        var v21 uint = 484
        switch 32767 {
        case 1:
            var v22 uint8 = 137
            _ = v22
        case -1:
            if true {
                v23 := GenS4{F0: false, F1: -686}
                _ = v23
            }
        case 4294967295:
            v24 := GenS3{F0: v16, F1: 30}
            _ = v24
        }
        if (!true) {
            for fi25 := 0; fi25 < 15; fi25++ {
                v21 += 0
                _ = fi25
            }
        } else {
            for fi26 := 0; fi26 < 7; fi26++ {
                for fi27 := 0; fi27 < 4; fi27++ {
                    v4[42] = (-(fi27))
                    v21 += 18446744073709551615
                    if (v16 && v16) {
                        if len(v5) > 0 {
                            v5[0] = ("\n\t" + ("\n\t" + "hello"))
                        }
                        v21 = v21
                    }
                    _ = fi27
                }
                for fi28 := 0; fi28 < 9; fi28++ {
                    v16 = v16
                    switch "" {
                    case "\n\t":
                        var v29 uint32 = 843
                        v30 := GenS5{F0: ("a" + "\n\t"), F1: 106, F2: "a"}
                        if (214 < 255) {
                            v31 := GenS6{F0: (fi28 + fi28), F1: 42}
                            if (-128 <= -54) {
                                if v16 {
                                    switch "hello" {
                                    case "hello":
                                        var v32 float32 = (-2.25 - 1.5)
                                        var v33 int = fi28
                                        v21 += 799
                                        _ = v32
                                        _ = v33
                                    }
                                    if v16 {
                                        var v34 string = ("a" + "\n\t")
                                        v35 := GenS7{F0: 1, F1: ((253 > 0) || v16), F2: 581}
                                        v34 = v34
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v4[(fi26 * fi26)] = fi28
                                        switch -324 {
                                        case -98:
                                            _ = len(v2)
                                        default:
                                            if (v16 && false) {
                                                if v16 {
                                                    _ = len(v5)
                                                } else {
                                                    v2 = append(v2, ("\n\t" + "foo bar"))
                                                }
                                                v36 := GenS5{F0: "foo bar", F1: -1, F2: "\n\t"}
                                                v29 = 2147483648
                                                if ((2 == 988) && (v16 && (!(2147483647 == 42)))) {
                                                    v16 = false
                                                    var v37 uint = 42
                                                    _ = v37
                                                } else {
                                                    v38 := GenS1{F0: "foo bar", F1: (("a" + (("hello" + ("" + "a")) + "a")) + "hello")}
                                                    v21 = v21
                                                    v39 := GenS8{F0: 1, F1: 1}
                                                    _ = v38
                                                    _ = v39
                                                }
                                                _ = v36
                                            } else {
                                                switch (fi26 % ((-1) | 1)) {
                                                case 13:
                                                    fmt.Println(v29, v21)
                                                case 42:
                                                    v40 := GenS6{F0: (fi26 << ((556) & 63)), F1: 2}
                                                    _ = v40
                                                default:
                                                    var v41 uint = (v21 * v21)
                                                    _ = v41
                                                }
                                                var v42 uint = v21
                                                _ = v42
                                            }
                                        }
                                    }
                                    v43 := GenS9{F0: -2147483648}
                                    _ = v43
                                }
                            } else {
                                var v44 int32 = 1
                                v44 = -2147483648
                                v45 := GenS9{F0: v44}
                                _ = len(v2)
                                _ = v44
                                _ = v45
                            }
                            v21 -= v21
                            _ = v31
                        } else {
                            var v46 uint8 = 98
                            fmt.Println(v16)
                            _ = v46
                        }
                        _ = v29
                        _ = v30
                    case "foo bar":
                        switch 0 {
                        case -791:
                            v47 := GenS1{F0: ("hello" + "foo bar"), F1: ("a" + "\n\t")}
                            _ = v47
                        case 42:
                            v48 := GenS10{F0: (-1 <= -79), F1: fi28, F2: -9223372036854775808}
                            _ = v48
                        default:
                            v16 = true
                        }
                    }
                    _ = fi28
                }
                var v49 uint16 = 255
                _ = fi26
                _ = v49
            }
            var v50 uint64 = 2
            var v51 uint8 = 40
            _ = v50
            _ = v51
        }
        _ = v21
    } else {
        switch ("" + "foo bar") {
        case "a":
            v16 = (true || false)
        case "foo bar":
            if v16 {
                v16 = v16
            } else {
                delete(v3, ("hello" + ("\n\t" + "")))
                if ((!((!(305 >= -2147483648)) || v16)) || ((((0.0 + 3.14159) < float64(0.0)) && (!(false && (!(455 != -839))))) && v16)) {
                    _ = len(v3)
                    v16 = (42 != 1)
                    switch ("hello" + (("a" + ("\n\t" + "hello")) + (("foo bar" + ((("\n\t" + "") + "foo bar") + ("\n\t" + ""))) + "a"))) {
                    case "foo bar":
                        var v52 int64 = 1
                        _ = v52
                    case "":
                        v53 := GenS11{F0: 1, F1: ("foo bar" + (("\n\t" + "\n\t") + ""))}
                        _ = v53
                    case "\n\t":
                        if (!(!v16)) {
                            var v54 float32 = (((1.5 * 0.0) * (0.1 / (1.5 + 1e10))) * 1.5)
                            if (1.5 == float64(float32(1.5))) {
                                fmt.Println(v54)
                                switch 448 {
                                case 127:
                                    v55 := GenS9{F0: 42}
                                    _ = v55
                                case 99:
                                    v54 = 0.1
                                }
                            }
                            fmt.Println(v16, v54)
                            _ = v54
                        } else {
                            v16 = v16
                            var v56 int8 = 100
                            _ = v56
                        }
                    }
                } else {
                    v3[(("" + "") + "")] = ("foo bar" + "")
                    v4[128] = -641
                }
            }
        default:
            var v57 int32 = 140
            _ = v57
        }
        v3["\n\t"] = "a"
        fmt.Println(v16)
    }
    if (1 >= 42) {
        v1 = append(v1, 4294967295)
        if (2147483647 == 65535) {
            v16 = false
            if (true || v16) {
                var v58 uint64 = 4294967295
                fmt.Println(v16)
                _ = v58
            }
        }
        var v59 int16 = 1
        _ = v59
    }
    var v60 uint = 56
    var v61 uint32 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
