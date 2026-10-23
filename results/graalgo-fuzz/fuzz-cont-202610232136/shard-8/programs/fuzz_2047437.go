package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint8; F1 int; F2 int }
type GenS6 struct { F0 bool; F1 string; F2 uint }

func main() {
    v1 := []uint32{2, 320, 42, 2147483647}
    var v2 int16 = -955
    v3 := []string{"foo bar", "a", "\n\t", ("foo bar" + "hello")}
    v4 := []uint32{2, 0}
    var v5 float64 = 0.0
    v6 := []string{("" + ""), ("\n\t" + "a")}
    var v7 string = "hello"
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int16 = (v2 | 1000)
        var v10 int32 = -2147483648
        for fi11 := 0; fi11 < 11; fi11++ {
            fmt.Println(fi11, fi8)
            var v12 float64 = v5
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int = (fi13 << ((42) & 63))
        switch (v14 * (v14 / ((fi13) | 1))) {
        case -506:
            var v15 int64 = 42
            _ = v15
        }
        var v16 uint64 = 1
        _ = fi13
        _ = v14
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS1{F0: fi17, F1: 1, F2: (^(v2))}
        _ = fi17
        _ = v18
    }
    var v19 uint = 2
    for fi20 := 0; fi20 < 14; fi20++ {
        if true {
            if (!((false && (((((v2 << ((v19) & 63)) <= (v2 / (((^(v2))) | 1))) && (42 > 255)) || (!false)) && false)) && true)) {
                var v21 uint64 = 861
                if false {
                    switch 318 {
                    case 5:
                        v22 := GenS1{F0: -497, F1: 32767, F2: v2}
                        _ = v22
                    case -531:
                        v19 = 2
                    default:
                        v19 += 1
                    }
                    var v23 bool = false
                    v19 = v19
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 9; fi24++ {
                        var v25 uint64 = (v21 * v21)
                        switch "" {
                        case "\n\t":
                            v26 := GenS1{F0: (fi20 >> (((^(v19))) & 63)), F1: 2, F2: (v2 / ((v2) | 1))}
                            _ = v26
                        case "foo bar":
                            v27 := GenS2{F0: 42}
                            _ = v27
                        case "a":
                            var v28 uint16 = 255
                            _ = v28
                        default:
                            if len(v6) > 0 {
                                v6[0] = "a"
                            }
                        }
                        var v29 int16 = (v2 & 525)
                        _ = fi24
                        _ = v25
                        _ = v29
                    }
                    if (v19 == 2) {
                        if ((32767 < 1) || (65535 <= 460)) {
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                        } else {
                            switch v7 {
                            case "foo bar":
                                if ((fi20 - fi20) == 42) {
                                    v21 += v21
                                }
                            case "a":
                                switch 2147483647 {
                                case 32767:
                                    v19 = v19
                                case 2147483647:
                                    v30 := GenS3{F0: 108, F1: 2}
                                    _ = v30
                                }
                            default:
                                v7 = "hello"
                            }
                            v2 = -690
                            fmt.Println(v7, fi20)
                        }
                        for fi31 := 0; fi31 < 3; fi31++ {
                            v2 = v2
                            _ = fi31
                        }
                    }
                }
                for fi32 := 0; fi32 < 9; fi32++ {
                    fmt.Println(v7)
                    _ = fi32
                }
                _ = v21
            }
            v7 = (v7 + (v7 + ""))
        } else {
            var v33 uint8 = 42
            v7 = ((v7 + "\n\t") + ("" + ""))
            _ = v33
        }
        v7 = v7
        _ = fi20
    }
    v34 := GenS3{F0: 1, F1: 0}
    v19 += 19
    if false {
        v5 = 0.1
    } else {
        v5 = v5
        if true {
            if (0 == 65535) {
                v2 -= v2
                fmt.Println(v2, v7)
                var v35 int16 = (v2 & v2)
                _ = v35
            }
            switch v7 {
            case "":
                if true {
                    if (!(v5 == 0.1)) {
                        v36 := GenS2{F0: 2147483647}
                        _ = v36
                    }
                    var v37 int32 = -736
                    if (!true) {
                        v38 := GenS2{F0: 0}
                        v39 := GenS4{F0: 4294967295, F1: 127, F2: 852}
                        _ = len(v6)
                        _ = v38
                        _ = v39
                    }
                    _ = v37
                }
            }
        } else {
            _ = len(v6)
            _ = len(v4)
            v2 = (-(v2))
        }
        v19 = (v19 ^ (v19 & (^(v19))))
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        var v41 uint32 = 1
        _ = fi40
        _ = v41
    }
    var v42 uint = v19
    if ((!((590 == 0) || true)) || (42 != 2147483647)) {
        var v43 float32 = 3.14159
        _ = v43
    } else {
        if true {
            v2 = (v2 ^ (v2 - v2))
            v7 = ("a" + v7)
        }
        v7 = v7
        for fi44 := 0; fi44 < 8; fi44++ {
            v5 -= v5
            v5 = (v5 * (v5 + v5))
            switch ((v7 + ("a" + v7)) + v7) {
            case "":
                if len(v3) > 0 {
                    v3[0] = ("" + v7)
                }
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                v45 := GenS5{F0: 33, F1: 1, F2: -2147483648}
                _ = v45
            }
            _ = fi44
        }
    }
    var v46 uint32 = 2
    _ = len(v6)
    var v47 uint = 304
    fmt.Println(v42)
    var v48 uint8 = 255
    if true {
        for fi49 := 0; fi49 < 12; fi49++ {
            v50 := GenS6{F0: false, F1: v7, F2: (v47 >> (((v42 & (v19 << ((uint(v46)) & 63)))) & 63))}
            _ = fi49
            _ = v50
        }
        v51 := GenS5{F0: 66, F1: 42, F2: -2147483648}
        fmt.Println(v46)
        _ = v51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
