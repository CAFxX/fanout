package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS8 struct { F0 string; F1 bool; F2 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint; F1 int64; F2 uint }

func main() {
    var v1 int64 = 169
    var v2 int16 = 561
    v3 := make(map[string]string)
    v1 -= 42
    fmt.Println(v2, v1)
    var v4 uint16 = 0
    v5 := []int{32768, -835, 9223372036854775807, -945}
    for fi6 := 0; fi6 < 17; fi6++ {
        v1 += 480
        v1 += 9223372036854775807
        if true {
            v7 := GenS1{F0: 1}
            v4 = 65535
            _ = v7
        } else {
            _ = len(v5)
        }
        _ = fi6
    }
    if true {
        var v8 int16 = v2
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
        var v9 bool = (-128 <= 42)
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 float32 = 1e10
            switch fi10 {
            case -9223372036854775808:
                v1 = (^(v1))
            }
            for fi12 := 0; fi12 < 14; fi12++ {
                v1 += v1
                v11 += 0.0
                switch (("\n\t" + ("a" + ("" + "hello"))) + (("\n\t" + ("hello" + "hello")) + "hello")) {
                case "\n\t":
                    if false {
                        v13 := GenS2{F0: 65535, F1: 127}
                        var v14 int = (fi10 * 0)
                        var v15 float32 = 1.5
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        if (fi10 >= 37) {
                            v3[("a" + "hello")] = ("a" + ("a" + (("foo bar" + ("hello" + "\n\t")) + "foo bar")))
                            _ = len(v3)
                        } else {
                            _ = len(v5)
                            var v16 uint64 = 302
                            _ = v16
                        }
                        switch "hello" {
                        case "hello":
                            if (18 < 255) {
                                v17 := GenS1{F0: 2}
                                v18 := GenS2{F0: 0, F1: 127}
                                _ = v17
                                _ = v18
                            } else {
                                v11 = v11
                                v3[""] = (("a" + ("a" + ("hello" + "hello"))) + "\n\t")
                                v4 -= (v4 >> ((52) & 63))
                            }
                            switch fi12 {
                            case -2147483648:
                                v19 := v3["foo bar"]
                                _ = v19
                            case 647:
                                switch ("foo bar" + "a") {
                                case "":
                                    var v20 bool = true
                                    _ = v20
                                }
                            }
                            v1 += -25
                        case "":
                            v21 := GenS3{F0: false, F1: fi12}
                            v22 := GenS3{F0: false, F1: -216}
                            if (0 != v4) {
                                v11 = 0.0
                            }
                            _ = v21
                            _ = v22
                        }
                    }
                }
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            if false {
                var v24 string = "a"
                var v25 int8 = 81
                _ = v24
                _ = v25
            }
            _ = fi23
        }
    }
    if true {
        v4 = 83
        if true {
            v2 = v2
            var v26 string = "hello"
            _ = v26
        } else {
            for fi27 := 0; fi27 < 15; fi27++ {
                v28 := GenS4{F0: 0, F1: 0}
                _ = fi27
                _ = v28
            }
        }
        fmt.Println(v2)
    } else {
        v2 = (v2 & v2)
        if (1 == 20) {
            if false {
                for fi29 := 0; fi29 < 19; fi29++ {
                    var v30 int32 = -1
                    v5 = append(v5, (fi29 % ((0) | 1)))
                    _ = fi29
                    _ = v30
                }
                v31 := GenS5{F0: 4294967295}
                _ = v31
            } else {
                var v32 int = 4294967295
                if (true || true) {
                    delete(v3, ("foo bar" + ""))
                    switch (v32 << ((4294967295) & 63)) {
                    case -1:
                        var v33 string = "\n\t"
                        _ = v33
                    case 483:
                        v4 -= (v4 | (v4 * (v4 / ((1) | 1))))
                    default:
                        for fi34 := 0; fi34 < 4; fi34++ {
                            var v35 string = ("hello" + (("hello" + (("" + "\n\t") + "a")) + ((("a" + ("a" + (("foo bar" + ("a" + "hello")) + "foo bar"))) + "hello") + "a")))
                            _ = fi34
                            _ = v35
                        }
                    }
                    var v36 uint32 = 795
                    _ = v36
                } else {
                    v2 = v2
                    var v37 int64 = 0
                    v4 -= 226
                    _ = v37
                }
                var v38 uint16 = v4
                _ = v32
                _ = v38
            }
            switch 128 {
            case -948:
                if (1 == -2147483648) {
                    v39 := GenS3{F0: false, F1: 4294967295}
                    if (!(true || true)) {
                        fmt.Println(v4)
                        if len(v5) > 0 {
                            v5[0] = -172
                        }
                    }
                    var v40 uint16 = (v4 | (v4 | (v4 >> ((65535) & 63))))
                    _ = v39
                    _ = v40
                }
            case 0:
                switch -1 {
                case -272:
                    var v41 uint64 = uint64(v4)
                    _ = v41
                default:
                    v42 := GenS3{F0: (100 < -19), F1: 1}
                    _ = v42
                }
            }
        }
        if (2 == 42) {
            v43 := GenS6{F0: v4, F1: 1}
            var v44 bool = (4294967295 >= -2147483648)
            _ = v43
            _ = v44
        } else {
            if true {
                switch ("" + "hello") {
                case "a":
                    v1 += v1
                default:
                    v4 = (v4 / ((255) | 1))
                }
                v1 = v1
                if false {
                    var v45 int8 = -3
                    _ = v45
                }
            } else {
                v46 := GenS7{F0: 65535, F1: (v4 | v4), F2: 42}
                fmt.Println(v4, v4)
                _ = v46
            }
            v47 := GenS2{F0: 1, F1: 0}
            fmt.Println(v4, v1)
            _ = v47
        }
    }
    v1 -= -222
    var v48 uint16 = 218
    if true {
        if (1 >= 4294967295) {
            v4 -= (v48 & (v48 << ((4294967295) & 63)))
            var v49 uint8 = 128
            _ = v49
        } else {
            var v50 uint8 = 1
            var v51 uint8 = v50
            v52 := GenS4{F0: v4, F1: -9223372036854775808}
            _ = v50
            _ = v51
            _ = v52
        }
        for fi53 := 0; fi53 < 17; fi53++ {
            v54 := GenS6{F0: (v4 % ((65535) | 1)), F1: 42}
            _ = fi53
            _ = v54
        }
        if ((!true) || (!false)) {
            if true {
                fmt.Println(v1, v4, v1)
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if false {
                    v55 := GenS5{F0: 0}
                    _ = v55
                } else {
                    for fi56 := 0; fi56 < 5; fi56++ {
                        var v57 float32 = 1.5
                        if len(v5) > 0 {
                            v5[0] = 128
                        }
                        for fi58 := 0; fi58 < 3; fi58++ {
                            v2 -= v2
                            _ = fi58
                        }
                        _ = fi56
                        _ = v57
                    }
                }
            }
            var v59 int64 = (^(v1))
            if (2147483647 != 42) {
                v60 := GenS1{F0: 674}
                var v61 uint8 = 127
                _ = v60
                _ = v61
            }
            _ = v59
        }
    } else {
        var v62 float32 = -2.25
        _ = v62
    }
    v48 += v48
    v63 := GenS8{F0: ("a" + ("" + "")), F1: ((^(v2)) >= -698), F2: "a"}
    switch "\n\t" {
    case "foo bar":
        v4 = (v4 / ((0) | 1))
    case "\n\t":
        v5 = append(v5, -265)
    case "hello":
        v48 = v48
    default:
        v5 = append(v5, 32768)
    }
    v64 := GenS9{F0: -706}
    switch 2147483647 {
    case 42:
        if (!false) {
            v65 := GenS10{F0: 0, F1: 0, F2: 2}
            v5 = append(v5, 2147483647)
            _ = v65
        }
    default:
        if (!true) {
            var v66 int8 = 1
            _ = v66
        }
    }
    v4 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v63
    _ = v64
}
