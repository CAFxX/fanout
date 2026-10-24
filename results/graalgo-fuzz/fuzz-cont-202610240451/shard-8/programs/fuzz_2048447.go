package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint; F1 uint; F2 uint8 }

func main() {
    v1 := []int{983, -684, 1}
    var v2 float32 = 1e10
    var v3 int = 0
    v4 := make(map[string]string)
    v5 := []string{"hello", "a", "\n\t"}
    var v6 int8 = 100
    v7 := GenS1{F0: 371, F1: 902, F2: 2}
    if ((2 != 2) && (!false)) {
        v8 := GenS2{F0: 0, F1: 255}
        switch (^(v3)) {
        case -648:
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 float64 = -2.25
                _ = fi9
                _ = v10
            }
        default:
            v2 -= (((v2 + v2) + 0.0) * (v2 + v2))
        }
        switch "foo bar" {
        case "":
            v3 += (v3 / ((1) | 1))
        }
        _ = v8
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint32 = 156
        v12 = v12
        switch ("" + ((("" + ("\n\t" + "\n\t")) + "a") + "hello")) {
        case "\n\t":
            for fi13 := 0; fi13 < 11; fi13++ {
                v3 = -9223372036854775808
                _ = fi13
            }
        case "":
            v12 += v12
        }
        _ = fi11
        _ = v12
    }
    if false {
        var v14 float32 = ((float32(1.5) * v2) - v2)
        v5 = append(v5, "foo bar")
        if true {
            _ = len(v4)
        }
        _ = v14
    } else {
        for fi15 := 0; fi15 < 10; fi15++ {
            switch "a" {
            case "hello":
                for fi16 := 0; fi16 < 7; fi16++ {
                    if true {
                        var v17 uint = 255
                        if true {
                            var v18 uint32 = 811
                            _ = v18
                        } else {
                            _ = len(v4)
                        }
                        _ = v17
                    } else {
                        delete(v4, "foo bar")
                        v3 = fi15
                        v19 := GenS3{F0: (v6 ^ v6), F1: 9223372036854775807, F2: 65535}
                        _ = v19
                    }
                    _ = fi16
                }
            case "a":
                var v20 uint = 4294967295
                _ = v20
            }
            v21 := GenS4{F0: 93}
            _ = fi15
            _ = v21
        }
    }
    if ((2 == 0) || (!(!(!false)))) {
        v22 := GenS5{F0: 0}
        _ = v22
    }
    v6 = (^(v6))
    var v23 int = 42
    v24 := GenS2{F0: 608, F1: 207}
    v25 := GenS5{F0: 4294967295}
    var v26 uint = 42
    if false {
        switch "a" {
        case "":
            switch "a" {
            case "hello":
                fmt.Println(v26, v26)
            case "a":
                switch "" {
                case "hello":
                    v3 = v3
                default:
                    v4[""] = ("\n\t" + "\n\t")
                }
            }
        case "hello":
            for fi27 := 0; fi27 < 15; fi27++ {
                v2 -= 1e10
                fmt.Println(v26, fi27)
                _ = fi27
            }
        default:
            for fi28 := 0; fi28 < 2; fi28++ {
                v2 = float32((float64(((3.14159 + (v2 + 3.14159)) / ((float32(1.5) - v2) * (v2 * (v2 + 1e10))))) + 0.1))
                _ = fi28
            }
        }
    } else {
        if true {
            v26 = v26
            for fi29 := 0; fi29 < 15; fi29++ {
                v5 = append(v5, "")
                _ = len(v5)
                _ = fi29
            }
            var v30 uint16 = 255
            _ = v30
        }
        fmt.Println(v26)
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        if true {
            switch ((("hello" + "hello") + "hello") + ("foo bar" + ("hello" + (("a" + "a") + "foo bar")))) {
            case "":
                v32 := v4["\n\t"]
                _ = v32
            case "a":
                if false {
                    v3 += (fi31 << (((v26 + v26)) & 63))
                    var v33 uint16 = 65535
                    if (!false) {
                        v33 = (v33 % ((v33) | 1))
                        v23 = v3
                    } else {
                        _ = len(v4)
                        v23 += fi31
                        v34 := v4[""]
                        _ = v34
                    }
                    _ = v33
                } else {
                    var v35 int = (v3 | v3)
                    for fi36 := 0; fi36 < 17; fi36++ {
                        if true {
                            switch ("hello" + "foo bar") {
                            case "":
                                if true {
                                    switch "hello" {
                                    case "foo bar":
                                        v35 -= v35
                                    case "":
                                        if (!(((145 != 1) && (!(!true))) || true)) {
                                            fmt.Println(v3, v3)
                                            v5 = append(v5, (("foo bar" + "") + ("foo bar" + "a")))
                                        } else {
                                            fmt.Println(fi31)
                                        }
                                    case "a":
                                        v3 = v23
                                    default:
                                        v37 := GenS3{F0: (v6 & (v6 % (((-(v6))) | 1))), F1: 9223372036854775807, F2: 692}
                                        _ = v37
                                    }
                                    v4["hello"] = "\n\t"
                                    v3 = fi36
                                } else {
                                    v38 := GenS4{F0: v6}
                                    v3 = v35
                                    _ = v38
                                }
                            case "hello":
                                switch "" {
                                case "a":
                                    var v39 uint64 = 619
                                    _ = v39
                                default:
                                    fmt.Println(v6, v35)
                                }
                            }
                            v40 := GenS6{F0: 4294967295}
                            _ = v40
                        }
                        switch "\n\t" {
                        case "hello":
                            delete(v4, "a")
                        }
                        var v41 bool = (255 <= 178)
                        _ = fi36
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 9; fi42++ {
                        v43 := GenS3{F0: v6, F1: -1, F2: 568}
                        switch (^(v3)) {
                        case 9223372036854775807:
                            if (("hello" + "a") >= "") {
                                var v44 int16 = -168
                                _ = v44
                            }
                        }
                        _ = fi42
                        _ = v43
                    }
                    _ = v35
                }
            case "\n\t":
                if true {
                    var v45 int8 = v6
                    fmt.Println(fi31)
                    _ = v45
                }
            default:
                for fi46 := 0; fi46 < 10; fi46++ {
                    fmt.Println(v6, fi46, fi31)
                    _ = fi46
                }
            }
            var v47 float32 = (-2.25 + 0.1)
            _ = v47
        }
        _ = fi31
    }
    for fi48 := 0; fi48 < 18; fi48++ {
        v2 += v2
        switch 127 {
        case -228:
            if true {
                v26 += v26
            }
        case 4294967295:
            fmt.Println(v2, v23)
        default:
            for fi49 := 0; fi49 < 14; fi49++ {
                v50 := GenS5{F0: 2147483648}
                var v51 bool = ((2 <= 532) || (!(42 <= 344)))
                if v51 {
                    var v52 string = "foo bar"
                    var v53 uint64 = 2
                    _ = v52
                    _ = v53
                } else {
                    v23 += (fi49 ^ v3)
                    var v54 int8 = (^(v6))
                    if (!(false && (!(true && v51)))) {
                        var v55 int32 = -2147483648
                        var v56 uint8 = 3
                        _ = v55
                        _ = v56
                    } else {
                        var v57 int8 = v54
                        var v58 float32 = (v2 + -2.25)
                        if v51 {
                            var v59 uint8 = 1
                            _ = v59
                        }
                        _ = v57
                        _ = v58
                    }
                    _ = v54
                }
                _ = fi49
                _ = v50
                _ = v51
            }
        }
        _ = fi48
    }
    switch 32768 {
    case 4294967295:
        v1 = append(v1, (v3 + v3))
    }
    v5 = append(v5, (("hello" + "a") + "a"))
    for fi60 := 0; fi60 < 17; fi60++ {
        v61 := GenS7{F0: (v26 & v26), F1: (v26 << (((v26 & v26)) & 63)), F2: 154}
        _ = fi60
        _ = v61
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
