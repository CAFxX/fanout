package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 bool; F1 int16; F2 int32 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint64 = 19
    var v2 int = -2147483648
    var v3 float64 = 1.5
    var v4 float64 = (float64((1.5 - ((1e10 * 0.1) * 0.0))) * 1.5)
    var v5 int16 = -536
    v6 := []string{""}
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int64 = -416
        var v9 string = "hello"
        var v10 uint = 1
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    fmt.Println(v4, v5)
    switch (("\n\t" + "") + "") {
    case "foo bar":
        if true {
            for fi11 := 0; fi11 < 6; fi11++ {
                if true {
                    v6 = append(v6, "\n\t")
                } else {
                    v5 = v5
                    var v12 float64 = float64(3.14159)
                    v13 := GenS1{F0: fi11, F1: (("\n\t" + "a") + (("hello" + "a") + ("foo bar" + "foo bar")))}
                    _ = v12
                    _ = v13
                }
                if len(v6) > 0 {
                    v6[0] = ""
                }
                for fi14 := 0; fi14 < 7; fi14++ {
                    switch (v2 % ((fi11) | 1)) {
                    case 473:
                        if (2 > 255) {
                            _ = len(v6)
                        }
                    }
                    _ = len(v6)
                    _ = fi14
                }
                _ = fi11
            }
            for fi15 := 0; fi15 < 13; fi15++ {
                v4 = float64(0.0)
                if len(v6) > 0 {
                    v6[0] = ((("hello" + "hello") + "\n\t") + "a")
                }
                var v16 uint8 = 2
                _ = fi15
                _ = v16
            }
            if ((((v2 <= v2) || true) && ((^(v1)) < 733)) || false) {
                v3 = (v4 / v4)
                v5 = (v5 & v5)
                for fi17 := 0; fi17 < 13; fi17++ {
                    if ((0.1 + 3.14159) == -2.25) {
                        v18 := GenS1{F0: v2, F1: (("hello" + "foo bar") + "\n\t")}
                        var v19 uint16 = 65535
                        _ = v18
                        _ = v19
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    if ((!false) || (42 == -1)) {
                        _ = len(v6)
                    } else {
                        v4 = 3.14159
                        v1 = v1
                        if false {
                            var v20 int = (fi17 << ((255) & 63))
                            v21 := GenS2{F0: 702, F1: 611}
                            _ = v20
                            _ = v21
                        } else {
                            v2 += (^(v2))
                            var v22 float32 = 0.0
                            _ = v22
                        }
                    }
                    _ = fi17
                }
            } else {
                v23 := GenS2{F0: 2, F1: 4294967295}
                _ = v23
            }
        }
    case "a":
        var v24 bool = true
        _ = v24
    default:
        for fi25 := 0; fi25 < 18; fi25++ {
            v26 := GenS2{F0: 2147483647, F1: 0}
            v27 := GenS1{F0: (v2 * (v2 | (v2 % (((-(v2))) | 1)))), F1: "hello"}
            v2 -= (-(fi25))
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    v4 -= v4
    v28 := GenS2{F0: 2147483647, F1: 2147483647}
    if true {
        v5 += 1000
        switch v2 {
        case 322:
            v6 = append(v6, ("\n\t" + "hello"))
        case 9223372036854775807:
            if (!(("\n\t" + "foo bar") <= (("foo bar" + "") + ""))) {
                var v29 float32 = ((-2.25 * 3.14159) + 1.5)
                _ = v29
            } else {
                for fi30 := 0; fi30 < 16; fi30++ {
                    var v31 bool = (!false)
                    v32 := GenS1{F0: v2, F1: "hello"}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                for fi33 := 0; fi33 < 11; fi33++ {
                    var v34 bool = true
                    var v35 int16 = v5
                    var v36 string = "foo bar"
                    _ = fi33
                    _ = v34
                    _ = v35
                    _ = v36
                }
            }
        case 0:
            if false {
                var v37 int16 = 1
                _ = len(v6)
                _ = v37
            }
        }
    }
    if true {
        if false {
            for fi38 := 0; fi38 < 16; fi38++ {
                var v39 int16 = (v5 - (v5 + 0))
                if ((v1 & v1) >= 4294967295) {
                    v40 := GenS2{F0: 4294967295, F1: 636}
                    v41 := GenS1{F0: 127, F1: "foo bar"}
                    for fi42 := 0; fi42 < 3; fi42++ {
                        var v43 uint8 = 128
                        var v44 uint16 = 42
                        switch "" {
                        case "hello":
                            v6 = append(v6, "a")
                        case "a":
                            v1 = (v1 % (((v1 & v1)) | 1))
                        case "":
                            var v45 float32 = 1.5
                            _ = v45
                        default:
                            v6 = append(v6, "\n\t")
                        }
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    _ = v40
                    _ = v41
                } else {
                    v46 := GenS1{F0: v2, F1: "a"}
                    _ = v46
                }
                _ = fi38
                _ = v39
            }
            v1 = (^(v1))
        }
        v47 := GenS3{F0: false, F1: int16(v2), F2: -1}
        if true {
            for fi48 := 0; fi48 < 3; fi48++ {
                var v49 uint64 = (^(v1))
                _ = len(v6)
                _ = fi48
                _ = v49
            }
            if true {
                fmt.Println(v3, v5)
                v50 := GenS4{F0: 2}
                _ = v50
            } else {
                for fi51 := 0; fi51 < 18; fi51++ {
                    if false {
                        fmt.Println(v5, v2)
                        for fi52 := 0; fi52 < 15; fi52++ {
                            v53 := GenS1{F0: int(v5), F1: (("a" + "foo bar") + "foo bar")}
                            var v54 float64 = ((v4 + (0.1 + v3)) - 1.5)
                            fmt.Println(v5, v2)
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                        v55 := GenS1{F0: fi51, F1: ("\n\t" + (("hello" + "a") + ("\n\t" + "foo bar")))}
                        _ = v55
                    } else {
                        v4 = (-2.25 * (v3 - ((v3 + ((3.14159 - 1.5) * 1.5)) - -2.25)))
                        v56 := GenS3{F0: (42 > 196), F1: 1, F2: 577}
                        _ = v56
                    }
                    _ = fi51
                }
            }
            var v57 int32 = -1
            _ = v57
        }
        _ = v47
    }
    for fi58 := 0; fi58 < 3; fi58++ {
        var v59 int8 = 127
        _ = fi58
        _ = v59
    }
    for fi60 := 0; fi60 < 11; fi60++ {
        fmt.Println(v1, v2, v3)
        _ = fi60
    }
    v3 = (v3 * v3)
    if len(v6) > 0 {
        v6[0] = "foo bar"
    }
    fmt.Println(v4, v3, v2)
    if (false && true) {
        var v61 bool = true
        _ = v61
    } else {
        _ = len(v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
