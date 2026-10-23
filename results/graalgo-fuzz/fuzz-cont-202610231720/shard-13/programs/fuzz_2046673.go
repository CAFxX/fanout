package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint }

func main() {
    var v1 int32 = 0
    var v2 uint8 = 57
    var v3 int64 = 1
    v4 := []string{("\n\t" + "foo bar"), ("hello" + (("" + ("foo bar" + "")) + "foo bar"))}
    var v5 uint8 = (v2 / (((v2 - v2)) | 1))
    for fi6 := 0; fi6 < 9; fi6++ {
        switch fi6 {
        case 311:
            var v7 uint32 = 4294967295
            _ = v7
        case 32768:
            if (!false) {
                if false {
                    if false {
                        var v8 uint16 = 32767
                        _ = v8
                    }
                } else {
                    var v9 bool = (!(97 != 32767))
                    switch "" {
                    case "a":
                        if v9 {
                            v10 := GenS1{F0: v1, F1: 1, F2: v3}
                            v11 := GenS2{F0: v3}
                            var v12 int32 = v1
                            _ = v10
                            _ = v11
                            _ = v12
                        } else {
                            switch "a" {
                            case "\n\t":
                                var v13 float64 = -2.25
                                _ = v13
                            case "a":
                                v14 := GenS3{F0: (v1 & v1)}
                                _ = v14
                            }
                        }
                    case "foo bar":
                        v15 := GenS4{F0: "\n\t", F1: v2}
                        _ = v15
                    case "hello":
                        fmt.Println(v5, v3)
                    }
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                    _ = v9
                }
                fmt.Println(v3, v1)
                var v16 int8 = 63
                _ = v16
            } else {
                if false {
                    var v17 uint = 65535
                    _ = v17
                } else {
                    v1 = 42
                }
            }
        case 22:
            v3 -= 654
        default:
            var v18 uint8 = 0
            _ = v18
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := GenS5{F0: 32767, F1: 310}
            _ = fi19
            _ = v20
        }
        _ = fi6
    }
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    if true {
        fmt.Println(v5, v3)
    } else {
        switch "a" {
        case "hello":
            if (!false) {
                switch -472 {
                case 128:
                    var v21 int8 = 73
                    _ = v21
                default:
                    var v22 int64 = (v3 >> ((459) & 63))
                    _ = v22
                }
            } else {
                for fi23 := 0; fi23 < 13; fi23++ {
                    v2 += v5
                    fmt.Println(v5, v1)
                    _ = fi23
                }
                switch 9223372036854775807 {
                case 89:
                    if true {
                        v24 := GenS4{F0: ("hello" + "hello"), F1: v5}
                        if ((!true) && false) {
                            var v25 float64 = 3.14159
                            v26 := GenS6{F0: 2147483648, F1: 32767, F2: 1}
                            _ = v25
                            _ = v26
                        } else {
                            v3 += (v3 - v3)
                            if (!(!false)) {
                                v3 = 381
                                if (true || false) {
                                    if ((("hello" + "") + "a") < "\n\t") {
                                        v5 -= v2
                                        fmt.Println(v1, v3, v3)
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = "foo bar"
                                        }
                                        var v27 uint32 = 65535
                                        v28 := GenS3{F0: v1}
                                        _ = v27
                                        _ = v28
                                    }
                                    var v29 int8 = 42
                                    _ = v29
                                }
                                if (94 == (v2 % ((127) | 1))) {
                                    v5 = 128
                                    if ("\n\t" != ("hello" + "hello")) {
                                        fmt.Println(v3)
                                        var v30 bool = ((!false) || (42 >= 4294967295))
                                        _ = v30
                                    } else {
                                        if false {
                                            _ = len(v4)
                                        } else {
                                            fmt.Println(v2)
                                            if len(v4) > 0 {
                                                v4[0] = ("\n\t" + ((("foo bar" + (("\n\t" + ("" + "foo bar")) + "a")) + "foo bar") + "foo bar"))
                                            }
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                } else {
                                    var v31 int32 = 0
                                    v5 = (v5 & 42)
                                    v32 := GenS2{F0: (^(v3))}
                                    _ = v31
                                    _ = v32
                                }
                            }
                        }
                        _ = v24
                    }
                case -1:
                    if true {
                        var v33 bool = false
                        _ = v33
                    }
                }
                fmt.Println(v1, v5, v1)
            }
        default:
            for fi34 := 0; fi34 < 10; fi34++ {
                if (((-2.25 - (1.5 / (-2.25 * (1.5 * -2.25)))) < 0.0) && true) {
                    v35 := GenS7{F0: 1}
                    _ = v35
                } else {
                    v36 := GenS6{F0: 0, F1: -758, F2: 454}
                    _ = v36
                }
                fmt.Println(v5, fi34)
                v37 := GenS3{F0: (v1 >> ((255) & 63))}
                _ = fi34
                _ = v37
            }
        }
    }
    var v38 int32 = (-(v1))
    fmt.Println(v3, v3, v38)
    fmt.Println(v38)
    var v39 float32 = 3.14159
    if false {
        v39 = v39
        if (-1 >= v38) {
            for fi40 := 0; fi40 < 5; fi40++ {
                for fi41 := 0; fi41 < 3; fi41++ {
                    v3 = -9223372036854775808
                    v42 := GenS8{F0: 919}
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
            v43 := GenS6{F0: 65535, F1: 947, F2: 0}
            _ = v43
        }
        var v44 uint64 = 0
        _ = v44
    } else {
        for fi45 := 0; fi45 < 19; fi45++ {
            if len(v4) > 0 {
                v4[0] = "a"
            }
            if ((((false && (!((-128 > 105) || (!true)))) || (("foo bar" + "a") > (("" + ("" + "")) + "a"))) && false) || false) {
                fmt.Println(v3)
                v4 = append(v4, "\n\t")
                v2 = (v2 << ((799) & 63))
            } else {
                var v46 uint8 = 0
                _ = v46
            }
            for fi47 := 0; fi47 < 18; fi47++ {
                if (!(true && false)) {
                    v48 := GenS6{F0: 2147483647, F1: -1, F2: 41}
                    if ("\n\t" <= "\n\t") {
                        fmt.Println(v38)
                        switch 1 {
                        case 127:
                            var v49 uint64 = 916
                            _ = v49
                        case -9223372036854775808:
                            var v50 uint64 = 42
                            _ = v50
                        default:
                            v2 = (^(v5))
                        }
                    }
                    _ = v48
                } else {
                    v51 := GenS8{F0: 4294967295}
                    v38 += v38
                    var v52 int64 = v3
                    _ = v51
                    _ = v52
                }
                _ = fi47
            }
            _ = fi45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
