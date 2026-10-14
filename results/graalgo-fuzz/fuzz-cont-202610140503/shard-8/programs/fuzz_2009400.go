package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint8; F2 int }
type GenS3 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 bool }
type GenS5 struct { F0 uint; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 float32 = 0.0
    var v2 uint64 = 36
    var v3 int32 = -2147483648
    v4 := []uint32{65535, 701}
    var v5 int64 = -9223372036854775808
    v5 = v5
    for fi6 := 0; fi6 < 17; fi6++ {
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v7 := GenS1{F0: 234}
            var v8 uint64 = v2
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 4; fi9++ {
                v2 = 455
                _ = fi9
            }
            v10 := GenS1{F0: 417}
            _ = v10
        }
        v1 -= (float32(0.1) / (v1 - float32(3.14159)))
        if true {
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 int8 = 4
                v4 = append(v4, 1)
                var v13 int8 = v12
                v14 := GenS1{F0: 0}
                switch (-(fi6)) {
                case 2147483647:
                    var v15 uint = 42
                    _ = v15
                default:
                    if ((v1 - v1) >= v1) {
                        if false {
                            fmt.Println(v2, v1, v3)
                            if true {
                                switch "\n\t" {
                                case "":
                                    var v16 float32 = float32((3.14159 / -2.25))
                                    _ = v16
                                case "a":
                                    if true {
                                        var v17 uint16 = 897
                                        _ = v17
                                    }
                                }
                            } else {
                                var v18 uint8 = 2
                                _ = v18
                            }
                            var v19 float64 = (3.14159 * 3.14159)
                            v20 := GenS2{F0: true, F1: 234, F2: (fi6 & 42)}
                            var v21 uint32 = 2147483647
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        switch 0 {
                        case -919:
                            v3 = v3
                        case 32767:
                            var v22 uint16 = 32767
                            _ = v22
                        }
                    } else {
                        v23 := GenS1{F0: 150}
                        _ = v23
                    }
                }
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        }
        _ = fi6
    }
    v3 = 42
    if true {
        v24 := GenS1{F0: 18446744073709551615}
        if (65535 < 255) {
            v5 = (v5 % (((v5 % ((-9223372036854775808) | 1))) | 1))
            var v25 float32 = v1
            _ = v25
        } else {
            v3 = 2147483647
            if (("foo bar" + (("\n\t" + "foo bar") + "\n\t")) == ("" + (("" + (("" + ((("a" + (("foo bar" + "hello") + "\n\t")) + "hello") + ("\n\t" + ("foo bar" + "a")))) + "\n\t")) + ""))) {
                var v26 uint = 42
                v3 = int32(v26)
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 uint32 = 42
                    for fi29 := 0; fi29 < 7; fi29++ {
                        var v30 uint64 = (^(v2))
                        v31 := GenS3{F0: 2, F1: (v3 % ((v3) | 1)), F2: 32767}
                        if false {
                            var v32 int = fi27
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(v1, v2)
                            default:
                                v33 := GenS4{F0: 255, F1: 65535, F2: false}
                                _ = v33
                            }
                            _ = v32
                        }
                        var v34 uint32 = (v28 + (v28 >> (((v26 - 255)) & 63)))
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi29
                        _ = v30
                        _ = v31
                        _ = v34
                    }
                    switch 2147483647 {
                    case 32767:
                        switch ("hello" + "hello") {
                        case "a":
                            v35 := GenS3{F0: 1, F1: int32(v26), F2: 65535}
                            _ = v35
                        case "foo bar":
                            var v36 int64 = v5
                            _ = v36
                        default:
                            v5 -= (v5 | 42)
                        }
                    case 42:
                        var v37 uint16 = 365
                        _ = v37
                    default:
                        for fi38 := 0; fi38 < 20; fi38++ {
                            _ = len(v4)
                            _ = fi38
                        }
                    }
                    _ = fi27
                    _ = v28
                }
                _ = v26
            } else {
                switch ("a" + "a") {
                case "a":
                    if false {
                        var v39 float64 = 0.1
                        _ = v39
                    }
                default:
                    v5 = v5
                }
            }
        }
        v3 -= (v3 * (v3 & v3))
        _ = v24
    } else {
        if (3.14159 > v1) {
            if ((!("foo bar" >= "")) && (127 > 701)) {
                v2 = v2
            }
            v4 = append(v4, 312)
            var v40 int8 = 19
            _ = v40
        } else {
            if false {
                for fi41 := 0; fi41 < 13; fi41++ {
                    if false {
                        if false {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v3 = v3
                            var v42 int64 = (v5 % ((9223372036854775807) | 1))
                            _ = v42
                        } else {
                            for fi43 := 0; fi43 < 14; fi43++ {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = fi43
                            }
                            v44 := GenS3{F0: 2, F1: v3, F2: 0}
                            _ = v44
                        }
                        fmt.Println(v5, v1)
                    }
                    var v45 int32 = v3
                    _ = fi41
                    _ = v45
                }
                var v46 float64 = (1e10 + 1e10)
                v4 = append(v4, 51)
                _ = v46
            } else {
                v5 += v5
                v5 = 9223372036854775807
            }
            if (!true) {
                if (false && ("a" > "a")) {
                    v47 := GenS5{F0: 447, F1: 42, F2: (v2 & v2)}
                    _ = v47
                }
                var v48 string = ((((("" + "") + "") + "\n\t") + "foo bar") + "\n\t")
                _ = v48
            }
            fmt.Println(v3, v3, v3)
        }
    }
    v4 = append(v4, 0)
    v2 += v2
    var v49 int32 = (^(v3))
    v50 := GenS6{F0: 0, F1: 4294967295}
    v4 = append(v4, 693)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
}
