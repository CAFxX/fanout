package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 string; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 int16; F2 int }

func main() {
    v1 := []string{("a" + "a"), "a"}
    var v2 float32 = 1.5
    var v3 float64 = 0.0
    var v4 bool = false
    if (0 >= 1) {
        var v5 uint64 = 243
        v6 := GenS1{F0: 112}
        v3 += ((1.5 / v3) / 1.0)
        _ = v5
        _ = v6
    } else {
        if (!false) {
            for fi7 := 0; fi7 < 4; fi7++ {
                for fi8 := 0; fi8 < 17; fi8++ {
                    if v4 {
                        v2 = v2
                        var v9 uint = 8
                        v10 := GenS1{F0: 232}
                        switch "\n\t" {
                        case "\n\t":
                            var v11 uint8 = 127
                            switch fi8 {
                            case 127:
                                if len(v1) > 0 {
                                    v1[0] = ("foo bar" + "hello")
                                }
                            }
                            if v4 {
                                v1 = append(v1, ("" + "foo bar"))
                            } else {
                                if (0 <= -47) {
                                    v12 := GenS1{F0: 2}
                                    _ = v12
                                }
                                var v13 int64 = 9223372036854775807
                                var v14 uint8 = (v11 / ((v11) | 1))
                                _ = v13
                                _ = v14
                            }
                            var v15 uint = (v9 >> ((18446744073709551615) & 63))
                            if true {
                                fmt.Println(v3)
                                var v16 uint64 = 591
                                _ = v16
                            }
                            v17 := GenS1{F0: 127}
                            if ((v11 >= (v11 % (((v11 << (((v15 + v15)) & 63))) | 1))) && false) {
                                var v18 int = (fi7 ^ (fi7 ^ 0))
                                _ = v18
                            }
                            _ = v11
                            _ = v15
                            _ = v17
                        case "a":
                            fmt.Println(v3)
                        }
                        if v4 {
                            _ = len(v1)
                        }
                        _ = v10
                        _ = v9
                    } else {
                        if v4 {
                            switch fi8 {
                            case -477:
                                v2 += (v2 - -2.25)
                                fmt.Println(fi7)
                                v1 = append(v1, ("a" + ""))
                                fmt.Println(v2)
                                v3 = (3.14159 - v3)
                            default:
                                v19 := GenS2{F0: 65535, F1: 42}
                                _ = v19
                            }
                            v2 -= float32((v3 * v3))
                        }
                        if (!(true || (!false))) {
                            v20 := GenS1{F0: 213}
                            if v4 {
                                v21 := GenS2{F0: 1, F1: 94}
                                _ = v21
                            } else {
                                if v4 {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v2 = ((v2 / (v2 + 0.1)) * (v2 * v2))
                                } else {
                                    fmt.Println(v2)
                                    var v22 float64 = v3
                                    var v23 uint16 = 380
                                    _ = v22
                                    _ = v23
                                }
                            }
                            _ = v20
                        }
                        v24 := GenS3{F0: "hello", F1: 4294967295, F2: 2}
                        _ = v24
                    }
                    _ = fi8
                }
                v1 = append(v1, ((("\n\t" + ("" + "hello")) + "") + (("\n\t" + "foo bar") + "foo bar")))
                if len(v1) > 0 {
                    v1[0] = ""
                }
                _ = fi7
            }
            var v25 bool = true
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        v1 = append(v1, "hello")
        _ = fi26
    }
    if false {
        fmt.Println(v4, v3)
        fmt.Println(v4, v3)
    }
    v3 = 0.1
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 float32 = (1e10 / (1e10 / float32(v3)))
        switch "a" {
        case "a":
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 string = ("foo bar" + (("hello" + ("foo bar" + "")) + "\n\t"))
                _ = fi29
                _ = v30
            }
        case "\n\t":
            v2 += (1.5 / 1.0)
        }
        v31 := GenS2{F0: 774, F1: 131}
        _ = fi27
        _ = v28
        _ = v31
    }
    switch 42 {
    case 127:
        if v4 {
            for fi32 := 0; fi32 < 2; fi32++ {
                v33 := GenS4{F0: 127, F1: 685}
                var v34 uint = 2
                var v35 int8 = 127
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
            v1 = append(v1, ("hello" + "foo bar"))
            fmt.Println(v4, v2, v4)
        } else {
            fmt.Println(v2)
            v36 := GenS5{F0: 42}
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v36
        }
    case 128:
        if (635 < 4294967295) {
            switch 42 {
            case 906:
                fmt.Println(v4)
            case -1:
                v37 := GenS6{F0: 137, F1: -1, F2: 42}
                _ = v37
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        var v38 uint16 = 748
        _ = v38
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        if (("a" == "hello") || v4) {
            var v40 int32 = 1
            _ = v40
        }
        _ = fi39
    }
    v3 += (v3 / float64(((3.14159 * (v2 + v2)) * (v2 / -2.25))))
    switch (((("foo bar" + (("a" + ("a" + "foo bar")) + "a")) + "foo bar") + "\n\t") + "") {
    case "foo bar":
        for fi41 := 0; fi41 < 16; fi41++ {
            v4 = v4
            _ = fi41
        }
    default:
        switch ("" + "") {
        case "foo bar":
            var v42 uint8 = 128
            _ = v42
        case "a":
            v3 += v3
        default:
            for fi43 := 0; fi43 < 5; fi43++ {
                v44 := GenS4{F0: -1, F1: 330}
                _ = fi43
                _ = v44
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
