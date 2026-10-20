package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int }
type GenS4 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int64; F2 int8 }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 uint8; F1 int16; F2 uint }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 4294967295
    v3 := []int{32768, -134}
    v4 := make(map[string]string)
    if (((false && (0 > 42)) && (401 == 1)) || false) {
        v2 = v2
    } else {
        v2 = 2
        v5 := GenS1{F0: -893, F1: (0.0 <= 1e10)}
        _ = v5
    }
    if true {
        v6 := GenS2{F0: 779, F1: -4, F2: 0}
        _ = v6
    } else {
        v2 = v2
        var v7 int64 = -345
        _ = v7
    }
    if true {
        var v8 uint8 = 2
        _ = v8
    } else {
        switch -932 {
        case 42:
            if len(v3) > 0 {
                v3[0] = 1
            }
        case 0:
            v9 := GenS3{F0: 65535, F1: 18446744073709551615, F2: 0}
            _ = v9
        default:
            var v10 uint64 = 92
            _ = v10
        }
    }
    var v11 int16 = 42
    v11 -= v11
    if (1 <= 524) {
        if false {
            var v12 uint8 = 255
            var v13 uint8 = 127
            _ = v12
            _ = v13
        }
        if true {
            if true {
                for fi14 := 0; fi14 < 10; fi14++ {
                    v1["foo bar"] = "\n\t"
                    v11 += v11
                    v2 += (v2 / (((v2 >> ((4294967295) & 63))) | 1))
                    _ = fi14
                }
                v11 -= -32768
                fmt.Println(v2)
            } else {
                v15 := GenS2{F0: (v2 / (((v2 + 979)) | 1)), F1: -1, F2: 1}
                if (!(v11 != 42)) {
                    if false {
                        v16 := GenS1{F0: (-(v11)), F1: (-431 <= 32767)}
                        for fi17 := 0; fi17 < 20; fi17++ {
                            if (65535 > 237) {
                                fmt.Println(fi17)
                            } else {
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    if (true || true) {
                                        v19 := GenS3{F0: 1, F1: 42, F2: (fi18 + fi17)}
                                        v2 -= (v2 * (v2 / ((2) | 1)))
                                        _ = v19
                                    }
                                    switch ("" + ("foo bar" + "hello")) {
                                    case "":
                                        fmt.Println(v11, fi18, fi18)
                                    case "foo bar":
                                        switch ("a" + "foo bar") {
                                        case "":
                                            var v20 int64 = 9223372036854775807
                                            _ = v20
                                        case "foo bar":
                                            v2 -= (v2 >> ((65535) & 63))
                                        default:
                                            fmt.Println(fi17, fi18, v2)
                                        }
                                    }
                                    var v21 bool = false
                                    v21 = (255 != 1)
                                    _ = fi18
                                    _ = v21
                                }
                                if len(v3) > 0 {
                                    v3[0] = fi17
                                }
                                v11 = -1
                            }
                            v11 = 673
                            _ = fi17
                        }
                        _ = v16
                    } else {
                        delete(v4, ((((("hello" + "hello") + "") + "foo bar") + "foo bar") + ("\n\t" + "foo bar")))
                    }
                }
                for fi22 := 0; fi22 < 19; fi22++ {
                    var v23 uint16 = 325
                    _ = fi22
                    _ = v23
                }
                _ = v15
            }
            fmt.Println(v11, v11)
        } else {
            for fi24 := 0; fi24 < 16; fi24++ {
                fmt.Println(v2, v2)
                _ = fi24
            }
        }
    }
    switch ("hello" + (("a" + ("hello" + "hello")) + "a")) {
    case "":
        fmt.Println(v11, v2)
    case "foo bar":
        if true {
            fmt.Println(v11)
        }
    default:
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 int = fi25
            var v27 bool = false
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    if (true || ((v2 + 399) == 4294967295)) {
        v28 := GenS4{F0: ((1.5 / 0.1) < 1e10), F1: 128, F2: 42}
        if (!false) {
            var v29 uint8 = 0
            _ = v29
        }
        _ = v28
    } else {
        var v30 uint16 = 862
        v30 -= 65535
        _ = len(v4)
        _ = v30
    }
    v4[((("" + "hello") + "\n\t") + "")] = "\n\t"
    for fi31 := 0; fi31 < 19; fi31++ {
        v32 := GenS5{F0: 127}
        for fi33 := 0; fi33 < 19; fi33++ {
            v34 := GenS1{F0: v11, F1: (false && false)}
            fmt.Println(fi31, v2)
            var v35 bool = true
            switch -765 {
            case 42:
                if (0 <= 42) {
                    _ = len(v4)
                    v36 := GenS6{F0: 56, F1: -1, F2: 1}
                    _ = v36
                }
            default:
                v11 = -32768
            }
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi31
        _ = v32
    }
    v37 := v4["foo bar"]
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 uint64 = 940
        v39 = 4294967295
        v40 := GenS7{F0: -9223372036854775808, F1: 65535}
        _ = fi38
        _ = v39
        _ = v40
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (195 <= 1) {
        for fi41 := 0; fi41 < 18; fi41++ {
            v42 := GenS2{F0: v2, F1: 127, F2: (v11 & v11)}
            for fi43 := 0; fi43 < 20; fi43++ {
                v11 = v11
                var v44 uint8 = 42
                _ = fi43
                _ = v44
            }
            v37 = ""
            _ = fi41
            _ = v42
        }
        for fi45 := 0; fi45 < 16; fi45++ {
            switch fi45 {
            case -587:
                for fi46 := 0; fi46 < 4; fi46++ {
                    var v47 bool = true
                    v2 = v2
                    _ = fi46
                    _ = v47
                }
            case 663:
                v3 = append(v3, fi45)
            default:
                if ((false || false) && false) {
                    v48 := GenS8{F0: 2, F1: (v11 - v11), F2: v2}
                    var v49 int32 = 1
                    var v50 int32 = (v49 << ((v2) & 63))
                    _ = v48
                    _ = v49
                    _ = v50
                }
            }
            if true {
                v11 = v11
                v2 = v2
            }
            v51 := GenS1{F0: v11, F1: (false && ((v11 / ((v11) | 1)) != (^(v11))))}
            _ = fi45
            _ = v51
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
