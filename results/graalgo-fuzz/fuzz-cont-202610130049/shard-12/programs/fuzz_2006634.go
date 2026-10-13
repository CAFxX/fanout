package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 bool }
type GenS2 struct { F0 int; F1 uint64; F2 int16 }

func main() {
    var v1 float32 = 0.0
    var v2 int64 = 1
    var v3 int32 = -2147483648
    v4 := []string{(("" + "") + ("a" + ("hello" + "hello"))), "\n\t", "hello"}
    var v5 uint = 255
    if (-373 > -385) {
        if (v2 > v2) {
            var v6 float32 = (v1 + ((3.14159 - ((-2.25 * (v1 / v1)) / ((v1 + float32(1e10)) / 1e10))) + float32(((1e10 / 3.14159) * 1e10))))
            _ = v6
        } else {
            if len(v4) > 0 {
                v4[0] = (("" + ("hello" + "\n\t")) + ("" + ""))
            }
        }
        if true {
            var v7 bool = ((v5 << ((776) & 63)) <= 923)
            _ = v7
        } else {
            var v8 uint64 = 808
            var v9 int = 2147483647
            _ = v8
            _ = v9
        }
    } else {
        var v10 uint32 = 2147483647
        _ = v10
    }
    var v11 string = "\n\t"
    var v12 uint32 = 0
    if len(v4) > 0 {
        v4[0] = (("" + ("" + v11)) + "\n\t")
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v12 = 218
    if (1 <= 46) {
        var v13 int64 = -1
        if (v12 > 65535) {
            switch ("" + v11) {
            case "a":
                var v14 int16 = 42
                _ = v14
            default:
                v11 = v11
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                if true {
                    var v16 int8 = 100
                    switch (fi15 & (fi15 >> ((v5) & 63))) {
                    case -1:
                        for fi17 := 0; fi17 < 9; fi17++ {
                            v16 = (-(v16))
                            var v18 bool = false
                            var v19 int = (fi17 | 354)
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    case 33:
                        for fi20 := 0; fi20 < 19; fi20++ {
                            var v21 bool = false
                            _ = fi20
                            _ = v21
                        }
                    default:
                        var v22 int64 = (v13 - (v13 - v2))
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 8; fi23++ {
                        var v24 int = fi15
                        var v25 int8 = (-(v16))
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = v16
                }
                var v26 uint8 = 42
                _ = fi15
                _ = v26
            }
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 int16 = 910
                for fi29 := 0; fi29 < 6; fi29++ {
                    v30 := GenS1{F0: 925, F1: v3, F2: (57 >= 255)}
                    _ = fi29
                    _ = v30
                }
                if (!true) {
                    for fi31 := 0; fi31 < 15; fi31++ {
                        var v32 float64 = (float64(v1) + (0.1 / 3.14159))
                        _ = fi31
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 20; fi33++ {
                        v2 -= v13
                        fmt.Println(v28)
                        _ = fi33
                    }
                }
                _ = fi27
                _ = v28
            }
        } else {
            for fi34 := 0; fi34 < 4; fi34++ {
                var v35 bool = true
                _ = fi34
                _ = v35
            }
        }
        if false {
            var v36 int32 = -2147483648
            var v37 uint16 = 343
            _ = v36
            _ = v37
        } else {
            for fi38 := 0; fi38 < 13; fi38++ {
                var v39 bool = true
                var v40 int16 = 1000
                _ = fi38
                _ = v39
                _ = v40
            }
            switch "foo bar" {
            case "\n\t":
                for fi41 := 0; fi41 < 18; fi41++ {
                    v5 = (v5 >> ((v5) & 63))
                    switch fi41 {
                    case 1:
                        for fi42 := 0; fi42 < 6; fi42++ {
                            fmt.Println(v3, v5)
                            _ = fi42
                        }
                    case 127:
                        switch v11 {
                        case "\n\t":
                            fmt.Println(v12)
                        case "foo bar":
                            v43 := GenS1{F0: 783, F1: v3, F2: false}
                            _ = v43
                        default:
                            for fi44 := 0; fi44 < 14; fi44++ {
                                v45 := GenS1{F0: 32767, F1: (v3 & v3), F2: false}
                                _ = len(v4)
                                _ = fi44
                                _ = v45
                            }
                        }
                    default:
                        for fi46 := 0; fi46 < 14; fi46++ {
                            v47 := GenS2{F0: (fi46 ^ fi41), F1: 930, F2: 0}
                            _ = fi46
                            _ = v47
                        }
                    }
                    _ = fi41
                }
            case "a":
                if false {
                    var v48 uint16 = 545
                    v3 += (v3 | 42)
                    _ = v48
                }
            default:
                var v49 int32 = 467
                _ = v49
            }
        }
        _ = v13
    }
    var v50 int64 = v2
    var v51 int8 = -128
    if true {
        switch ((v11 + v11) + (v11 + (v11 + ((("foo bar" + ("\n\t" + (((v11 + v11) + "") + ""))) + v11) + ((v11 + v11) + v11))))) {
        case "a":
            if false {
                v3 = v3
            }
        case "\n\t":
            var v52 int32 = (v3 - (v3 ^ (^(v3))))
            _ = v52
        }
        v4 = append(v4, v11)
        v1 = (((((1e10 + v1) * (v1 + 3.14159)) * 1.5) - float32(0.0)) / v1)
    }
    v3 += 0
    if len(v4) > 0 {
        v4[0] = ""
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
}
