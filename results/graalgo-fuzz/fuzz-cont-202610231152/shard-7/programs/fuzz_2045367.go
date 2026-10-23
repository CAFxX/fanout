package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 42
    v3 := []uint32{v2, (v2 * v2)}
    var v4 int8 = 102
    if (!true) {
        var v5 int = 454
        _ = v5
    } else {
        v6 := GenS1{F0: 65535}
        v7 := GenS1{F0: 2}
        if false {
            if len(v3) > 0 {
                v3[0] = (v2 * (v2 | v2))
            }
            var v8 string = ("\n\t" + "a")
            _ = v8
        } else {
            if true {
                v2 = (v2 << ((42) & 63))
            }
            v4 += (v4 << ((18446744073709551615) & 63))
            var v9 uint32 = 2147483647
            _ = v9
        }
        _ = v6
        _ = v7
    }
    _ = len(v3)
    var v10 float64 = (float64(0.1) + float64((1.5 + (-2.25 * (0.1 - 3.14159)))))
    var v11 uint16 = 0
    if len(v3) > 0 {
        v3[0] = v2
    }
    if true {
        fmt.Println(v4, v10)
        if (!(true || (!((true || (v2 > 0)) && ((v11 < (^(v11))) && ((1 > 110) || (18446744073709551615 > uint64(v2)))))))) {
            delete(v1, ((("hello" + (((("foo bar" + "\n\t") + ("hello" + ("foo bar" + ""))) + "a") + "")) + ("hello" + ("hello" + "hello"))) + (("a" + "\n\t") + "\n\t")))
            if (-9223372036854775808 > 2147483647) {
                for fi12 := 0; fi12 < 2; fi12++ {
                    v4 = v4
                    for fi13 := 0; fi13 < 6; fi13++ {
                        delete(v1, "foo bar")
                        _ = fi13
                    }
                    _ = fi12
                }
                var v14 int8 = v4
                v15 := GenS1{F0: (v11 % (((v11 ^ (v11 << ((18446744073709551615) & 63)))) | 1))}
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 13; fi16++ {
                    _ = len(v1)
                    var v17 float32 = 0.0
                    _ = fi16
                    _ = v17
                }
            }
            if (v4 <= (v4 ^ v4)) {
                v3 = append(v3, 132)
                if (("" + (("\n\t" + "a") + "\n\t")) == ("hello" + "\n\t")) {
                    fmt.Println(v10, v2)
                } else {
                    v18 := v1[""]
                    switch "foo bar" {
                    case "hello":
                        v19 := GenS2{F0: 9223372036854775807, F1: (-1 == -517), F2: 1}
                        _ = v19
                    default:
                        for fi20 := 0; fi20 < 2; fi20++ {
                            var v21 bool = true
                            if (false || (((-1 > -518) || v21) && (v21 && (!(false && (!((0.1 + v10) == -2.25))))))) {
                                v22 := GenS1{F0: v11}
                                _ = v22
                            } else {
                                switch -817 {
                                case -1:
                                    v23 := GenS2{F0: -1, F1: (((false && v21) && v18) || v18), F2: 45}
                                    _ = v23
                                case 305:
                                    var v24 uint8 = 179
                                    _ = v24
                                case 2147483647:
                                    if v21 {
                                        var v25 uint64 = 18446744073709551615
                                        _ = v25
                                    } else {
                                        for fi26 := 0; fi26 < 4; fi26++ {
                                            switch fi20 {
                                            case -665:
                                                if v18 {
                                                    var v27 uint = 18446744073709551615
                                                    var v28 uint32 = (^(v2))
                                                    _ = v27
                                                    _ = v28
                                                }
                                            case 214:
                                                var v29 string = "foo bar"
                                                _ = v29
                                            default:
                                                _ = len(v1)
                                            }
                                            _ = fi26
                                        }
                                        for fi30 := 0; fi30 < 15; fi30++ {
                                            var v31 uint32 = (v2 - v2)
                                            var v32 int32 = 42
                                            var v33 uint = 18
                                            _ = fi30
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                }
                            }
                            _ = fi20
                            _ = v21
                        }
                    }
                    v4 = (v4 % (((v4 / (((v4 / ((v4) | 1))) | 1))) | 1))
                    _ = v18
                }
                switch -9223372036854775808 {
                case -921:
                    if (true || (!true)) {
                        for fi34 := 0; fi34 < 14; fi34++ {
                            delete(v1, (("\n\t" + "a") + "hello"))
                            _ = fi34
                        }
                    } else {
                        for fi35 := 0; fi35 < 8; fi35++ {
                            var v36 int8 = v4
                            _ = fi35
                            _ = v36
                        }
                    }
                }
            }
        }
    } else {
        var v37 uint64 = 971
        if false {
            v38 := GenS1{F0: 423}
            v39 := GenS2{F0: -85, F1: false, F2: 28}
            var v40 bool = (!((v2 + (v2 & 975)) > 2))
            _ = v38
            _ = v39
            _ = v40
        }
        v10 = v10
        _ = v37
    }
    v10 += (1e10 / float64(1.5))
    if (v10 >= (v10 - (v10 / (float64(0.1) - ((v10 * ((0.1 / 1.5) / 1.0)) + v10))))) {
        v41 := GenS3{F0: 42}
        _ = v41
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        if (false && (1e10 > 0.1)) {
            var v43 int32 = -102
            switch ("hello" + "") {
            case "hello":
                fmt.Println(fi42)
            case "":
                v11 = (v11 & (v11 << ((4294967295) & 63)))
            case "foo bar":
                v10 = (float64(1.5) / (float64(1e10) * v10))
            default:
                v10 = v10
            }
            _ = v43
        } else {
            for fi44 := 0; fi44 < 3; fi44++ {
                v45 := GenS2{F0: -343, F1: (!false), F2: 42}
                if false {
                    var v46 uint32 = (v2 >> ((842) & 63))
                    if true {
                        fmt.Println(v10, v2)
                        var v47 uint8 = 42
                        _ = v47
                    }
                    if false {
                        switch (fi44 * (fi44 - fi44)) {
                        case -2147483648:
                            v2 = (v46 ^ (v2 & (^(v2))))
                        case -379:
                            if true {
                                fmt.Println(fi42)
                            }
                        default:
                            var v48 int16 = -32768
                            _ = v48
                        }
                    } else {
                        switch fi44 {
                        case 4294967295:
                            switch -1 {
                            case -2147483648:
                                fmt.Println(v10, v4, v11)
                            case -470:
                                v49 := GenS2{F0: -9223372036854775808, F1: true, F2: 42}
                                _ = v49
                            }
                        case 32767:
                            var v50 uint64 = 933
                            _ = v50
                        default:
                            var v51 uint64 = 1
                            _ = v51
                        }
                    }
                    _ = v46
                } else {
                    if len(v3) > 0 {
                        v3[0] = v2
                    }
                }
                _ = fi44
                _ = v45
            }
            v2 = v2
            switch (fi42 << ((42) & 63)) {
            case 2147483647:
                v52 := v1["a"]
                _ = v52
            }
        }
        if false {
            v1["hello"] = (v2 <= (v2 << ((796) & 63)))
            for fi53 := 0; fi53 < 10; fi53++ {
                switch 127 {
                case 32768:
                    v2 = 576
                    var v54 bool = (!false)
                    var v55 int8 = v4
                    _ = v54
                    _ = v55
                case 406:
                    var v56 float64 = ((0.0 * (v10 - v10)) - v10)
                    _ = v56
                case 302:
                    v57 := v1["hello"]
                    _ = v57
                default:
                    var v58 string = "foo bar"
                    _ = v58
                }
                v59 := GenS4{F0: 413}
                _ = fi53
                _ = v59
            }
        }
        v60 := v1["foo bar"]
        _ = fi42
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
