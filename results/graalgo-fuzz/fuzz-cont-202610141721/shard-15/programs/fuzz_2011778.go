package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint32; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 int8; F2 uint16 }

func main() {
    v1 := []string{(("a" + "\n\t") + "\n\t"), "foo bar"}
    var v2 int8 = 127
    v3 := make(map[int]bool)
    var v4 uint = 495
    var v5 int64 = -967
    var v6 int64 = (v5 >> ((v4) & 63))
    v7 := GenS1{F0: false, F1: 4294967295}
    v8 := GenS1{F0: (!true), F1: 161}
    if (!(!(false || ((false || (!(!false))) && true)))) {
        v9 := v3[-958]
        v5 = 9223372036854775807
        fmt.Println(v6)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 12; fi10++ {
            if false {
                v11 := GenS2{F0: -446, F1: 673, F2: 1}
                _ = v11
            }
            _ = fi10
        }
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v3[-113] = false
    v4 = v4
    if true {
        v4 = v4
        if (true && false) {
            var v12 bool = false
            if ((245 <= 1) && v12) {
                for fi13 := 0; fi13 < 18; fi13++ {
                    if (1e10 <= 1.5) {
                        if ((fi13 >= 128) && false) {
                            v14 := GenS2{F0: -156, F1: 270, F2: 42}
                            var v15 uint64 = 352
                            _ = len(v1)
                            _ = v14
                            _ = v15
                        }
                    } else {
                        v16 := GenS2{F0: 0, F1: 965, F2: -2147483648}
                        _ = v16
                    }
                    _ = fi13
                }
                v1 = append(v1, "")
                v5 -= v6
            } else {
                fmt.Println(v4)
                v1 = append(v1, "foo bar")
                var v17 float32 = 0.0
                _ = v17
            }
            _ = v12
        } else {
            if false {
                if ((v5 * v6) >= (v5 + 42)) {
                    v3[4294967295] = false
                    v18 := GenS3{F0: (v6 >> ((v4) & 63))}
                    _ = v18
                } else {
                    switch int(v4) {
                    case 2147483647:
                        if false {
                            _ = len(v1)
                            _ = len(v1)
                        } else {
                            if (-1 == 1) {
                                v2 -= (v2 + 1)
                                if (0 == 42) {
                                    v2 -= v2
                                    var v19 int8 = (v2 << ((2) & 63))
                                    _ = v19
                                } else {
                                    v2 = v2
                                    v5 = 83
                                }
                                var v20 string = ("a" + "a")
                                _ = v20
                            } else {
                                if true {
                                    v21 := GenS1{F0: (245 != -1), F1: 788}
                                    if false {
                                        v5 -= (v5 << (((v4 + v4)) & 63))
                                        fmt.Println(v5, v4, v2)
                                        if (!false) {
                                            var v22 uint32 = 2147483648
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v22
                                        }
                                    } else {
                                        switch ("hello" + ("hello" + "hello")) {
                                        case "hello":
                                            switch 527 {
                                            case -2147483648:
                                                if true {
                                                    for fi23 := 0; fi23 < 9; fi23++ {
                                                        if true {
                                                            for fi24 := 0; fi24 < 8; fi24++ {
                                                                if len(v1) > 0 {
                                                                    v1[0] = "a"
                                                                }
                                                                switch (fi23 % (((fi24 / ((0) | 1))) | 1)) {
                                                                case 0:
                                                                    delete(v3, (fi24 | fi24))
                                                                default:
                                                                    var v25 int = (fi24 | (^(fi23)))
                                                                    _ = v25
                                                                }
                                                                v4 = v4
                                                                _ = fi24
                                                            }
                                                            var v26 uint64 = 0
                                                            v27 := GenS1{F0: true, F1: v26}
                                                            _ = v26
                                                            _ = v27
                                                        }
                                                        _ = len(v3)
                                                        _ = fi23
                                                    }
                                                    v2 -= -1
                                                    v28 := GenS4{F0: "foo bar", F1: ("\n\t" + "foo bar")}
                                                    _ = v28
                                                }
                                            case 1:
                                                _ = len(v3)
                                            }
                                        case "foo bar":
                                            v4 = 323
                                        case "\n\t":
                                            delete(v3, 942)
                                        default:
                                            v6 = v6
                                        }
                                    }
                                    _ = v21
                                } else {
                                    if true {
                                        v6 = (v5 + (v5 % ((v6) | 1)))
                                        fmt.Println(v4)
                                    } else {
                                        fmt.Println(v5, v2)
                                        v29 := GenS1{F0: (0 >= 434), F1: uint64(v4)}
                                        _ = v29
                                    }
                                }
                            }
                            fmt.Println(v4, v5, v2)
                            switch "hello" {
                            case "":
                                var v30 uint64 = 940
                                _ = v30
                            case "hello":
                                var v31 int64 = v6
                                _ = v31
                            }
                        }
                    default:
                        fmt.Println(v5, v4, v6)
                    }
                    switch "foo bar" {
                    case "foo bar":
                        if true {
                            var v32 float64 = 1.5
                            _ = v32
                        } else {
                            var v33 string = "hello"
                            switch -2147483648 {
                            case -1:
                                _ = len(v1)
                            }
                            _ = v33
                        }
                    default:
                        var v34 uint = 368
                        _ = v34
                    }
                }
                var v35 int16 = -32768
                _ = v35
            }
            v6 -= 929
        }
    } else {
        v4 -= (v4 / ((v4) | 1))
        switch 1 {
        case 9223372036854775807:
            switch -9223372036854775808 {
            case 32767:
                if true {
                    for fi36 := 0; fi36 < 4; fi36++ {
                        var v37 uint64 = 2
                        v4 = 0
                        var v38 int8 = -128
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                    switch ("hello" + ("a" + ("" + "\n\t"))) {
                    case "foo bar":
                        switch "" {
                        case "\n\t":
                            for fi39 := 0; fi39 < 9; fi39++ {
                                delete(v3, fi39)
                                v40 := GenS3{F0: -9223372036854775808}
                                _ = fi39
                                _ = v40
                            }
                        case "a":
                            for fi41 := 0; fi41 < 3; fi41++ {
                                v42 := GenS1{F0: true, F1: 617}
                                for fi43 := 0; fi43 < 7; fi43++ {
                                    var v44 uint = v4
                                    _ = len(v3)
                                    _ = fi43
                                    _ = v44
                                }
                                _ = fi41
                                _ = v42
                            }
                        }
                    case "hello":
                        v45 := GenS5{F0: 2}
                        _ = v45
                    }
                }
            case 1:
                v6 += v5
            case 127:
                switch 32767 {
                case 127:
                    fmt.Println(v5, v2)
                }
            }
        default:
            var v46 uint64 = 321
            _ = v46
        }
    }
    v5 = (v6 * 212)
    v47 := GenS6{F0: "\n\t", F1: 0, F2: 255}
    if true {
        v48 := GenS3{F0: v6}
        fmt.Println(v4, v6)
        _ = len(v3)
        _ = v48
    }
    v49 := GenS4{F0: ("" + "foo bar"), F1: ("hello" + ("foo bar" + "a"))}
    var v50 uint = (v4 + v4)
    v5 = v5
    var v51 bool = true
    fmt.Println(v5, v6)
    var v52 string = ("foo bar" + ("foo bar" + "foo bar"))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v7
    _ = v8
}
