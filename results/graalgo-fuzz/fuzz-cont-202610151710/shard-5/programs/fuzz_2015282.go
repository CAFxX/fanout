package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 string; F2 string }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64; F1 int8; F2 int8 }
type GenS10 struct { F0 int16; F1 uint32; F2 uint }

func main() {
    var v1 float64 = 0.0
    var v2 bool = true
    var v3 int64 = 42
    var v4 uint8 = 2
    var v5 uint = 65535
    v6 := []uint32{39, 364}
    var v7 int64 = 42
    fmt.Println(v5, v3, v4)
    var v8 string = "foo bar"
    if (v8 < v8) {
        v5 -= 42
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 int8 = 0
                switch 247 {
                case 734:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case 0:
                    var v12 uint16 = 65535
                    _ = v12
                default:
                    var v13 uint = v5
                    _ = v13
                }
                var v14 int = (fi9 & fi10)
                if v2 {
                    v15 := GenS1{F0: 2, F1: "\n\t"}
                    if v2 {
                        fmt.Println(fi9, v5, v1)
                        v16 := GenS2{F0: v2, F1: 928}
                        _ = len(v6)
                        if (!(!(2 > 18446744073709551615))) {
                            v14 = 127
                            var v17 uint32 = 0
                            var v18 int32 = -2147483648
                            switch (fi9 >> (((v5 | v5)) & 63)) {
                            case 93:
                                var v19 int64 = int64(v5)
                                _ = v19
                            }
                            if ((!((^(v5)) == (v5 - v5))) || (!(false || true))) {
                                v20 := GenS2{F0: (-1 >= (-(v18))), F1: 2}
                                v3 -= (v7 & 1)
                                switch v8 {
                                case "foo bar":
                                    if (71 != v4) {
                                        v21 := GenS3{F0: 42, F1: 1}
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        fmt.Println(v11)
                                        v22 := GenS4{F0: 884, F1: 255}
                                        _ = v21
                                        _ = v22
                                    } else {
                                        switch (v8 + "foo bar") {
                                        case "":
                                            if len(v6) > 0 {
                                                v6[0] = v17
                                            }
                                        }
                                        fmt.Println(v11, v8)
                                    }
                                case "":
                                    if v2 {
                                        v11 = (-(v11))
                                    }
                                case "\n\t":
                                    v3 += (v7 % ((v7) | 1))
                                default:
                                    v23 := GenS4{F0: 103, F1: 0}
                                    _ = v23
                                }
                                _ = v20
                            }
                            if (445 == 500) {
                                if v2 {
                                    fmt.Println(v18)
                                }
                                v24 := GenS3{F0: (v3 >> ((42) & 63)), F1: (v4 << (((v5 - (v5 << ((v5) & 63)))) & 63))}
                                _ = v24
                            }
                            _ = v17
                            _ = v18
                        } else {
                            v25 := GenS5{F0: v14}
                            if len(v6) > 0 {
                                v6[0] = 1
                            }
                            fmt.Println(fi10, v11, v1)
                            _ = v25
                        }
                        var v26 int64 = v3
                        _ = v16
                        _ = v26
                    }
                    v1 -= v1
                    _ = v15
                }
                v2 = (65535 > 32767)
                v5 = v5
                _ = fi10
                _ = v11
                _ = v14
            }
            _ = fi9
        }
        var v27 bool = true
        _ = v27
    }
    v28 := GenS5{F0: 2147483647}
    v29 := GenS6{F0: 32767, F1: v8, F2: v8}
    for fi30 := 0; fi30 < 9; fi30++ {
        for fi31 := 0; fi31 < 6; fi31++ {
            v5 = 152
            _ = fi31
        }
        _ = fi30
    }
    if v2 {
        if (v8 >= (("\n\t" + "") + "\n\t")) {
            var v32 uint8 = v4
            _ = v32
        } else {
            v2 = v2
            v33 := GenS1{F0: 244, F1: "\n\t"}
            _ = v33
        }
    } else {
        var v34 uint = (^(v5))
        var v35 uint = 970
        v36 := GenS4{F0: (^(v3)), F1: 987}
        _ = v34
        _ = v35
        _ = v36
    }
    var v37 int8 = -41
    v38 := GenS7{F0: 1, F1: ("foo bar" + "")}
    var v39 uint = 18446744073709551615
    if v2 {
        if (0.0 < 0.0) {
            v7 += v3
        } else {
            _ = len(v6)
            switch "" {
            case "a":
                if len(v6) > 0 {
                    v6[0] = 849
                }
            case "hello":
                switch -294 {
                case -657:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case 42:
                    v39 -= v39
                case 233:
                    v37 = v37
                default:
                    v40 := GenS8{F0: (v4 ^ v4)}
                    _ = v40
                }
            default:
                fmt.Println(v2, v7)
            }
        }
        v2 = (!true)
        var v41 int64 = (v3 + (-(v3)))
        _ = v41
    } else {
        for fi42 := 0; fi42 < 9; fi42++ {
            var v43 int8 = (-(v37))
            fmt.Println(v8, v2)
            v44 := GenS3{F0: (v7 & v3), F1: 42}
            _ = fi42
            _ = v43
            _ = v44
        }
        v45 := GenS9{F0: 2, F1: v37, F2: (v37 * (^(v37)))}
        _ = v45
    }
    switch 128 {
    case 42:
        v5 = v5
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v46 int32 = -634
    if (!((!false) && v2)) {
        switch int(v37) {
        case 1:
            v4 = (^(v4))
        case 2147483647:
            switch 4294967295 {
            case -1:
                fmt.Println(v5, v4, v8)
            case -369:
                var v47 int16 = 32767
                _ = v47
            }
        case 74:
            var v48 int16 = -32768
            _ = v48
        }
        var v49 int8 = v37
        if len(v6) > 0 {
            v6[0] = 2
        }
        _ = v49
    } else {
        _ = len(v6)
        v50 := GenS10{F0: 1000, F1: 548, F2: (v39 ^ v39)}
        _ = v50
    }
    if v2 {
        v6 = append(v6, uint32(v39))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
